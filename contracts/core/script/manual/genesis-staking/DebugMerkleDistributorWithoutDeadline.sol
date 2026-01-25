// SPDX-License-Identifier: GPL-3.0-only
pragma solidity 0.8.24;

import { MerkleDistributor } from "src/token/MerkleDistributor.sol";
import { OwnableUpgradeable } from "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import { EIP712 } from "solady/src/utils/EIP712.sol";
import { SignatureCheckerLib } from "solady/src/utils/SignatureCheckerLib.sol";
import { SafeTransferLib } from "solady/src/utils/SafeTransferLib.sol";
import { MerkleProofLib } from "solady/src/utils/MerkleProofLib.sol";
import { LibBitmap } from "solady/src/utils/LibBitmap.sol";
import { IStaking } from "src/interfaces/IStaking.sol";
import { IOmniPortal } from "src/interfaces/IOmniPortal.sol";
import { IGenesisStakeV2 } from "src/interfaces/IGenesisStakeV2.sol";
import { IERC7683, IOriginSettler } from "solve/src/erc7683/IOriginSettler.sol";
import { SolverNet } from "solve/src/lib/SolverNet.sol";
import { IERC20, SafeERC20 } from "@openzeppelin/contracts/token/ERC20/utils/SafeERC20.sol";

contract DebugMerkleDistributorWithoutDeadline is MerkleDistributor, OwnableUpgradeable, EIP712 {
    using LibBitmap for LibBitmap.Bitmap;
    using SafeTransferLib for address;
    using SafeERC20 for IERC20;

    error ZeroAddress();
    error InvalidSignature();
    error InsufficientAmount();
    error ManualClaimDisabled();
    error Expired();

    bytes32 internal constant ORDERDATA_TYPEHASH = keccak256(
        "OrderData(address owner,uint64 destChainId,Deposit deposit,Call[] calls,TokenExpense[] expenses)Deposit(address token,uint96 amount)Call(address target,bytes4 selector,uint256 value,bytes params)TokenExpense(address spender,address token,uint96 amount)"
    );
    bytes32 internal constant UPGRADE_TYPEHASH =
        keccak256("Upgrade(address user,address validator,uint256 nonce,uint256 expiry)");

    address internal constant STAKING = 0xCCcCcC0000000000000000000000000000000001;

    IOmniPortal public immutable omniPortal;
    IGenesisStakeV2 public immutable genesisStaking;
    IOriginSettler public immutable solvernetInbox;

    mapping(uint256 version => LibBitmap.Bitmap) internal _claimedBitmaps;
    mapping(address account => uint256) public nonces;
    uint256 public bitmapVersion;

    constructor(
        address token_,
        bytes32 merkleRoot_,
        address omniPortal_,
        address genesisStaking_,
        address solverNetInbox_
    ) MerkleDistributor(token_, merkleRoot_) {
        _disableInitializers();
        if (omniPortal_ == address(0) || genesisStaking_ == address(0) || solverNetInbox_ == address(0)) {
            revert ZeroAddress();
        }

        omniPortal = IOmniPortal(omniPortal_);
        genesisStaking = IGenesisStakeV2(genesisStaking_);
        solvernetInbox = IOriginSettler(solverNetInbox_);
    }

    /**
     * @notice Initialize the contract
     * @param admin_            The admin of the contract
     */
    function initialize(address admin_) external initializer {
        __Ownable_init(admin_);

        token.safeApprove(address(solvernetInbox), type(uint256).max);
    }

    /**
     * @notice Resets claim status for all users
     */
    function resetClaims() external onlyOwner {
        ++bitmapVersion;
    }

    /**
     * @notice Check if a claim has been made
     * @dev Debug override to allow resetting claim status
     * @param index Index of the claim
     * @return _ True if the claim has been made, false otherwise
     */
    function isClaimed(uint256 index) public view virtual override returns (bool) {
        return _claimedBitmaps[bitmapVersion].get(index);
    }

    /**
     * @notice Override to prevent manual claims - use upgradeStake or unstake instead
     * @dev This function always reverts to prevent manual claims
     */
    function claim(uint256, address, uint256, bytes32[] calldata) public pure override {
        revert ManualClaimDisabled();
    }

    /**
     * @notice Withdraw tokens from the contract
     */
    function withdraw(address to) external onlyOwner {
        token.safeTransferAll(to);
    }

    /**
     * @notice Get the EIP-712 digest for a stake upgrade signature
     * @param account   Address of the user upgrading
     * @param validator Validator to delegate to
     * @param expiry    Signature expiry
     * @return _        Upgrade digest
     */
    function getUpgradeDigest(address account, address validator, uint256 expiry) public view returns (bytes32) {
        if (expiry != 0 && block.timestamp > expiry) revert Expired();
        bytes32 migrationHash = keccak256(abi.encode(UPGRADE_TYPEHASH, account, validator, nonces[account], expiry));
        return _hashTypedData(migrationHash);
    }

    /**
     * @notice Claim rewards and upgrade stake to Omni
     * @dev Triggers a SolverNet order to generate a subsidized order for deposited tokens on Omni 1:1
     *      If the user has already claimed rewards, they can still upgrade their stake to Omni
     * @param validator    Validator to delegate to
     * @param index        Index of the claim
     * @param amount       Amount of tokens to claim
     * @param merkleProof  Merkle proof for the claim
     */
    function upgradeStake(address validator, uint256 index, uint256 amount, bytes32[] calldata merkleProof) external {
        unchecked {
            ++nonces[msg.sender];
        }
        _upgrade(msg.sender, validator, index, amount, merkleProof);
    }

    /**
     * @notice Claim rewards and upgrade stake to Omni on behalf of a user
     * @dev Triggers a SolverNet order to generate a subsidized order for deposited tokens on Omni 1:1
     *      If the user has already claimed rewards, they can still upgrade their stake to Omni
     * @param account      Address of the user upgrading
     * @param validator    Validator to delegate to
     * @param index        Index of the claim
     * @param amount       Amount of tokens to claim
     * @param merkleProof  Merkle proof for the claim
     * @param v            Signature v
     * @param r            Signature r
     * @param s            Signature s
     * @param expiry       Signature expiry
     */
    function upgradeUserStake(
        address account,
        address validator,
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
            bytes32 digest = getUpgradeDigest(account, validator, expiry);

            if (!SignatureCheckerLib.isValidSignatureNow(account, digest, v, r, s)) {
                if (!SignatureCheckerLib.isValidERC1271SignatureNow(account, digest, v, r, s)) {
                    revert InvalidSignature();
                }
            }

            unchecked {
                ++nonces[account];
            }
        }

        _upgrade(account, validator, index, amount, merkleProof);
    }

    /**
     * @notice Unstake from Genesis Staking and claim rewards
     * @param index        Index of the claim
     * @param amount       Amount of tokens to claim
     * @param merkleProof  Merkle proof for the claim
     */
    function unstake(uint256 index, uint256 amount, bytes32[] calldata merkleProof) external {
        unchecked {
            ++nonces[msg.sender];
        }
        _unstake(msg.sender, index, amount, merkleProof);
    }

    /**
     * @notice Upgrade stake to Omni
     * @param account      Address of the user upgrading
     * @param validator    Validator to delegate to
     * @param index        Index of the claim
     * @param amount       Amount of tokens to claim
     * @param merkleProof  Merkle proof for the claim
     */
    function _upgrade(
        address account,
        address validator,
        uint256 index,
        uint256 amount,
        bytes32[] calldata merkleProof
    ) internal {
        if (validator == address(0)) revert ZeroAddress();

        // Migrate user's stake, if any
        uint256 stake = IGenesisStakeV2(genesisStaking).migrateStake(account);

        // If proofs are provided, check if the user is eligible for rewards and add them to their stake
        if (merkleProof.length > 0) {
            if (_claimRewards(account, index, amount, merkleProof)) {
                unchecked {
                    stake += amount;
                }
            }
        }

        // Block insufficient stake migrations
        if (stake < 1 ether) revert InsufficientAmount();

        // Generate and send the order
        IERC7683.OnchainCrossChainOrder memory order = _generateOrder(account, validator, stake);
        solvernetInbox.open(order);
    }

    /**
     * @notice Unstake from Genesis Staking and potentially claim rewards
     * @param account      Address of the user unstaking
     * @param index        Index of the claim (if claiming)
     * @param amount       Amount of tokens to claim (if claiming)
     * @param merkleProof  Merkle proof for the claim (if claiming)
     */
    function _unstake(address account, uint256 index, uint256 amount, bytes32[] calldata merkleProof) internal {
        uint256 totalAmount = genesisStaking.migrateStake(account);

        // if proofs provided and rewards not already claimed, add them to total
        if (merkleProof.length > 0 && _claimRewards(account, index, amount, merkleProof)) {
            totalAmount += amount;
            emit Claimed(index, account, amount);
        }

        // transfer total amount if greater than 0
        if (totalAmount > 0) {
            token.safeTransfer(account, totalAmount);
        }
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
            _claimedBitmaps[bitmapVersion].set(index);
            return true;
        }
        return false;
    }

    /**
     * @notice Generate a SolverNet order that generates a subsidized order for deposited tokens on Omni 1:1
     * @param account   Address of the user claiming
     * @param validator Validator to delegate to
     * @param amount    Amount of tokens to claim
     * @return         SolverNet order
     */
    function _generateOrder(address account, address validator, uint256 amount)
        internal
        view
        returns (IERC7683.OnchainCrossChainOrder memory)
    {
        SolverNet.Deposit memory deposit = SolverNet.Deposit({ token: token, amount: uint96(amount) });

        SolverNet.Call[] memory call = new SolverNet.Call[](1);
        call[0] = SolverNet.Call({
            target: STAKING,
            selector: IStaking.delegateFor.selector,
            value: amount,
            params: abi.encode(account, validator)
        });

        SolverNet.OrderData memory orderData = SolverNet.OrderData({
            owner: account,
            destChainId: omniPortal.omniChainId(),
            deposit: deposit,
            calls: call,
            expenses: new SolverNet.TokenExpense[](0)
        });

        return IERC7683.OnchainCrossChainOrder({
            fillDeadline: uint32(block.timestamp + 6 hours),
            orderDataType: ORDERDATA_TYPEHASH,
            orderData: abi.encode(orderData)
        });
    }

    function _domainNameAndVersion() internal pure override returns (string memory name, string memory version) {
        name = "DebugMerkleDistributorWithoutDeadline";
        version = "1";
    }
}
