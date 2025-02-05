// SPDX-License-Identifier: GPL-3.0-only
pragma solidity 0.8.24;

import { MerkleDistributor } from "./MerkleDistributor.sol";
import { Ownable } from "solady/src/auth/Ownable.sol";
import { SafeTransferLib } from "solady/src/utils/SafeTransferLib.sol";
import { MerkleProofLib } from "solady/src/utils/MerkleProofLib.sol";
import { LibBitmap } from "solady/src/utils/LibBitmap.sol";
import { IOmniPortal } from "../interfaces/IOmniPortal.sol";
import { IGenesisStake } from "../interfaces/IGenesisStake.sol";
import { IERC7683, IOriginSettler } from "solve/src/erc7683/IOriginSettler.sol";
import { SolverNet } from "solve/src/lib/SolverNet.sol";

contract MerkleDistributorWithDeadline is MerkleDistributor, Ownable {
    using LibBitmap for LibBitmap.Bitmap;
    using SafeTransferLib for address;

    error EndTimeInPast();
    error NothingToMigrate();
    error ClaimWindowFinished();
    error NoWithdrawDuringClaim();

    bytes32 internal constant ORDERDATA_TYPEHASH = keccak256(
        "OrderData(address owner,uint64 destChainId,Deposit deposit,Call[] calls,Expense[] expenses)Deposit(address token,uint96 amount)Call(address target,bytes4 selector,uint256 value,bytes params)Expense(address spender,address token,uint96 amount)"
    );

    uint256 public immutable endTime;
    IOmniPortal public immutable omniPortal;
    IGenesisStake public immutable genesisStaking;
    IOriginSettler public immutable solvernetInbox;

    constructor(
        address token_,
        bytes32 merkleRoot_,
        uint256 endTime_,
        address omniPortal_,
        address genesisStaking_,
        address solverNetInbox_
    ) MerkleDistributor(token_, merkleRoot_) {
        if (endTime_ <= block.timestamp) revert EndTimeInPast();

        _initializeOwner(msg.sender);
        token_.safeApprove(solverNetInbox_, type(uint256).max);

        endTime = endTime_;
        omniPortal = IOmniPortal(omniPortal_);
        genesisStaking = IGenesisStake(genesisStaking_);
        solvernetInbox = IOriginSettler(solverNetInbox_);
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

        // Migrate user's stake to Omni, if any
        uint256 stake = IGenesisStake(genesisStaking).migrateStake(msg.sender);

        // Mark reward distribution as claimed and add it to user's stake
        if (!isClaimed(index)) {
            // Verify the merkle proof.
            bytes32 node = keccak256(abi.encodePacked(index, msg.sender, amount));
            if (!MerkleProofLib.verifyCalldata(merkleProof, merkleRoot, node)) revert InvalidProof();

            // Update bitmap and add claim amount to stake amount
            claimedBitMap.set(index);
            unchecked {
                stake += amount;
            }
        }

        // Cannot migrate if user has no stake to migrate
        if (stake == 0) revert NothingToMigrate();

        // Generate and send the order
        IERC7683.OnchainCrossChainOrder memory order = _generateOrder(msg.sender, stake);
        solvernetInbox.open(order);
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
        SolverNet.Deposit memory deposit = SolverNet.Deposit({ token: token, amount: uint96(amount) });

        SolverNet.Call[] memory call = new SolverNet.Call[](1);
        call[0] = SolverNet.Call({ target: account, selector: bytes4(0), value: amount, params: "" });

        SolverNet.OrderData memory orderData = SolverNet.OrderData({
            owner: account,
            destChainId: omniPortal.omniChainId(),
            deposit: deposit,
            calls: call,
            expenses: new SolverNet.Expense[](0)
        });

        return IERC7683.OnchainCrossChainOrder({
            fillDeadline: 0,
            orderDataType: ORDERDATA_TYPEHASH,
            orderData: abi.encode(orderData)
        });
    }
}
