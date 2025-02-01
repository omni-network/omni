// SPDX-License-Identifier: GPL-3.0-only
pragma solidity 0.8.24;

import { MerkleDistributor } from "../libraries/MerkleDistributor.sol";
import { Ownable } from "solady/src/auth/Ownable.sol";
import { SafeTransferLib } from "solady/src/utils/SafeTransferLib.sol";
import { MerkleProofLib } from "solady/src/utils/MerkleProofLib.sol";
import { LibBitmap } from "solady/src/utils/LibBitmap.sol";
import { AddrUtils } from "solve/lib/AddrUtils.sol";
import { IGenesisStake } from "../interfaces/IGenesisStake.sol";
import { IERC7683, IOriginSettler } from "solve/erc7683/IOriginSettler.sol";
import { ISolverNet } from "solve/interfaces/ISolverNet.sol";

contract MerkleDistributorWithDeadline is MerkleDistributor, Ownable {
    using LibBitmap for LibBitmap.Bitmap;
    using SafeTransferLib for address;
    using AddrUtils for address;

    error EndTimeInPast();
    error ClaimWindowFinished();
    error NoWithdrawDuringClaim();

    bytes32 internal constant ORDER_DATA_TYPEHASH = keccak256(
        "OrderData(address owner,Call call,Deposit[] deposits)Call(uint64 chainId,bytes32 target,uint256 value,bytes data,TokenExpense[] expenses)TokenExpense(bytes32 token,bytes32 spender,uint256 amount)Deposit(bytes32 token,uint256 amount)"
    );
    uint64 internal constant OMNI_CHAINID = 166;

    uint256 public immutable endTime;
    address public immutable genesisStaking;
    address public immutable solverNetInbox;

    constructor(address token_, bytes32 merkleRoot_, uint256 endTime_, address genesisStaking_, address solverNetInbox_)
        MerkleDistributor(token_, merkleRoot_)
    {
        if (endTime_ <= block.timestamp) revert EndTimeInPast();

        _initializeOwner(msg.sender);
        token_.safeApprove(solverNetInbox_, type(uint256).max);

        endTime = endTime_;
        genesisStaking = genesisStaking_;
        solverNetInbox = solverNetInbox_;
    }

    /**
     * @notice Claim rewards
     * @dev Does not trigger any changes on the GenesisStake contract
     * @param index        Index of the claim
     * @param account      Address of the user claiming
     * @param amount       Amount of tokens to claim
     * @param merkleProof  Merkle proof for the claim
     */
    function claim(uint256 index, address account, uint256 amount, bytes32[] calldata merkleProof) public override {
        if (block.timestamp > endTime) revert ClaimWindowFinished();
        super.claim(index, account, amount, merkleProof);
    }

    /**
     * @notice Claim rewards and/or migrate stake to Omni
     * @dev Triggers a SolverNet order to generate a subsidized order for deposited tokens on Omni 1:1
     *      If the user has already claimed rewards, they can still migrate their stake to Omni
     * @param index        Index of the claim
     * @param amount       Amount of tokens to claim
     * @param merkleProof  Merkle proof for the claim
     */
    function migrateToOmni(uint256 index, uint256 amount, bytes32[] calldata merkleProof) external {
        if (block.timestamp > endTime) revert ClaimWindowFinished();

        // Verify the merkle proof.
        bytes32 node = keccak256(abi.encodePacked(index, msg.sender, amount));
        if (!MerkleProofLib.verifyCalldata(merkleProof, merkleRoot, node)) revert InvalidProof();

        // Migrate user's stake to Omni, if any
        uint256 stake = IGenesisStake(genesisStaking).migrateStake(msg.sender);

        // Mark reward distribution as claimed and add it to user's stake
        if (!isClaimed(index)) {
            claimedBitMap.set(index);
            unchecked {
                stake += amount;
            }
        }

        // Generate and send the order
        IERC7683.OnchainCrossChainOrder memory order = _generateOrder(msg.sender, stake);
        IOriginSettler(solverNetInbox).open(order);
    }

    function withdraw() external onlyOwner {
        if (block.timestamp < endTime) revert NoWithdrawDuringClaim();
        token.safeTransfer(msg.sender, token.balanceOf(address(this)));
    }

    /**
     * @notice Generate a SolverNet order that generates a subsidized order for deposited tokens on Omni 1:1
     * @param account  Address of the user claiming
     * @param amount   Amount of tokens to claim
     * @return         SolverNet order
     */
    function _generateOrder(address account, uint256 amount)
        internal
        view
        returns (IERC7683.OnchainCrossChainOrder memory)
    {
        ISolverNet.Deposit[] memory deposits = new ISolverNet.Deposit[](1);
        deposits[0] = ISolverNet.Deposit({ token: token.toBytes32(), amount: amount });

        ISolverNet.TokenExpense[] memory expenses = new ISolverNet.TokenExpense[](1);
        expenses[0] =
            ISolverNet.TokenExpense({ token: address(0).toBytes32(), spender: address(0).toBytes32(), amount: amount });

        ISolverNet.Call memory call = ISolverNet.Call({
            chainId: OMNI_CHAINID,
            target: account.toBytes32(),
            value: amount,
            data: "",
            expenses: expenses
        });

        ISolverNet.OrderData memory orderData = ISolverNet.OrderData({ owner: account, call: call, deposits: deposits });

        IERC7683.OnchainCrossChainOrder memory order = IERC7683.OnchainCrossChainOrder({
            fillDeadline: 0,
            orderDataType: ORDER_DATA_TYPEHASH,
            orderData: abi.encode(orderData)
        });
        return order;
    }
}
