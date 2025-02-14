// SPDX-License-Identifier: GPL-3.0-only
pragma solidity 0.8.24;

import { MerkleDistributor } from "./MerkleDistributor.sol";
import { Ownable } from "solady/src/auth/Ownable.sol";
import { EIP712 } from "solady/src/utils/EIP712.sol";
import { SignatureCheckerLib } from "solady/src/utils/SignatureCheckerLib.sol";
import { SafeTransferLib } from "solady/src/utils/SafeTransferLib.sol";
import { MerkleProofLib } from "solady/src/utils/MerkleProofLib.sol";
import { LibBitmap } from "solady/src/utils/LibBitmap.sol";
import { IOmniPortal } from "../interfaces/IOmniPortal.sol";
import { IGenesisStake } from "../interfaces/IGenesisStake.sol";
import { IERC7683, IOriginSettler } from "solve/src/erc7683/IOriginSettler.sol";
import { SolverNet } from "solve/src/lib/SolverNet.sol";

contract MerkleDistributorWithDeadline is MerkleDistributor, Ownable, EIP712 {
    using LibBitmap for LibBitmap.Bitmap;
    using SafeTransferLib for address;

    error Expired();
    error EndTimeInPast();
    error InvalidSignature();
    error NothingToMigrate();
    error ClaimWindowFinished();
    error NoWithdrawDuringClaim();

    bytes32 internal constant ORDERDATA_TYPEHASH = keccak256(
        "OrderData(address owner,uint64 destChainId,Deposit deposit,Call[] calls,Expense[] expenses)Deposit(address token,uint96 amount)Call(address target,bytes4 selector,uint256 value,bytes params)Expense(address spender,address token,uint96 amount)"
    );
    bytes32 internal constant MIGRATION_TYPEHASH = keccak256("Migration(address user,uint256 nonce,uint256 expiry)");

    uint256 public immutable endTime;
    IOmniPortal public immutable omniPortal;
    IGenesisStake public immutable genesisStaking;
    IOriginSettler public immutable solvernetInbox;

    mapping(address account => uint256) public nonces;

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
     * @notice Get the EIP-712 digest for a migration signature
     * @param account  Address of the user migrating
     * @param expiry   Signature expiry
     * @return _       Migration digest
     */
    function getMigrationDigest(address account, uint256 expiry) public view returns (bytes32) {
        if (expiry != 0 && block.timestamp > expiry) revert Expired();
        bytes32 migrationHash = keccak256(abi.encode(MIGRATION_TYPEHASH, account, nonces[account], expiry));
        return _hashTypedData(migrationHash);
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
     * @notice Claim rewards and migrate stake to Omni
     * @dev Triggers a SolverNet order to generate a subsidized order for deposited tokens on Omni 1:1
     *      If the user has already claimed rewards, they can still migrate their stake to Omni
     * @param index        Index of the claim
     * @param amount       Amount of tokens to claim
     * @param merkleProof  Merkle proof for the claim
     */
    function migrateToOmni(uint256 index, uint256 amount, bytes32[] calldata merkleProof) external {
        _migrate(msg.sender, index, amount, merkleProof);
    }

    /**
     * @notice Claim rewards and migrate stake to Omni on behalf of a user
     * @dev Triggers a SolverNet order to generate a subsidized order for deposited tokens on Omni 1:1
     *      If the user has already claimed rewards, they can still migrate their stake to Omni
     * @param account      Address of the user migrating
     * @param index        Index of the claim
     * @param amount       Amount of tokens to claim
     * @param merkleProof  Merkle proof for the claim
     * @param v            Signature v
     * @param r            Signature r
     * @param s            Signature s
     * @param expiry       Signature expiry
     */
    function migrateUserToOmni(
        address account,
        uint256 index,
        uint256 amount,
        bytes32[] calldata merkleProof,
        uint8 v,
        bytes32 r,
        bytes32 s,
        uint256 expiry
    ) external {
        // If the user isn't the caller, verify the signature
        if (account != msg.sender) {
            bytes32 digest = getMigrationDigest(account, expiry);

            if (!SignatureCheckerLib.isValidSignatureNow(account, digest, v, r, s)) {
                if (!SignatureCheckerLib.isValidERC1271SignatureNow(account, digest, v, r, s)) {
                    revert InvalidSignature();
                }
            }

            unchecked {
                ++nonces[account];
            }
        }

        _migrate(account, index, amount, merkleProof);
    }

    /**
     * @notice Withdraw tokens from the contract
     * @dev Reverts if called before the claim window has ended
     */
    function withdraw() external onlyOwner {
        if (block.timestamp < endTime) revert NoWithdrawDuringClaim();
        token.safeTransfer(msg.sender, token.balanceOf(address(this)));
    }

    /**
     * @notice Migrate stake to Omni
     * @param account      Address of the user migrating
     * @param index        Index of the claim
     * @param amount       Amount of tokens to claim
     * @param merkleProof  Merkle proof for the claim
     */
    function _migrate(address account, uint256 index, uint256 amount, bytes32[] calldata merkleProof) internal {
        if (block.timestamp > endTime) revert ClaimWindowFinished();

        // Migrate user's stake, if any
        uint256 stake = IGenesisStake(genesisStaking).migrateStake(account);

        // If the user has unclaimed rewards, add them to their stake
        if (_claimRewards(account, index, amount, merkleProof)) {
            unchecked {
                stake += amount;
            }
        }

        // Block zero value migrations
        if (stake == 0) revert NothingToMigrate();

        // Generate and send the order
        IERC7683.OnchainCrossChainOrder memory order = _generateOrder(account, stake);
        solvernetInbox.open(order);
    }

    /**
     * @notice Claim rewards
     * @dev Reverts if merkle proofs are invalid
     * @param account  Address of the user claiming
     * @param index    Index of the claim
     * @param amount   Amount of tokens to claim
     * @return _       True if the claim was processed, false if already claimed
     */
    function _claimRewards(address account, uint256 index, uint256 amount, bytes32[] calldata merkleProof)
        internal
        returns (bool)
    {
        // If rewards are unclaimed, verify the proof and process the claim
        if (!isClaimed(index)) {
            bytes32 leaf = keccak256(abi.encodePacked(index, account, amount));
            if (!MerkleProofLib.verifyCalldata(merkleProof, merkleRoot, leaf)) revert InvalidProof();
            claimedBitMap.set(index);
            return true;
        }
        return false;
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

    function _domainNameAndVersion() internal pure override returns (string memory name, string memory version) {
        name = "MerkleDistributorWithDeadline";
        version = "1";
    }
}
