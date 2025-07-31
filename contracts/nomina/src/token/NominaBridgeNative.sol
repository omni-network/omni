// SPDX-License-Identifier: GPL-3.0-only
pragma solidity 0.8.30;

import { NominaBridgeCommon } from "./NominaBridgeCommon.sol";
import { INominaPortal } from "src/interfaces/INominaPortal.sol";
import { NominaBridgeL1 } from "./NominaBridgeL1.sol";
import { ConfLevel } from "src/libraries/ConfLevel.sol";
import { XTypes } from "src/libraries/XTypes.sol";

/**
 * @title NominaBridgeNative
 * @notice The Nomina side of Nomina's native token bridge. Partner to NominaBridgeL1, which is deployed to Ethereum.
 *         This contract is predeployed to Nomina's EVM, prefunded with native NOM tokens to match totalL1Supply, such
 *         that each L1 token has a "sibling" native token on Nomina.
 * @dev This contract is predeployed, and requires storage slots to be set in genesis.
 *      initialize(...) is called pre-deployment, in script/genesis/AllocPredeploys.s.sol
 *      Initializers on the implementation are disabled via manual storage updates, rather than in a constructor.
 *      If a new implementation is required, a constructor should be added.
 */
contract NominaBridgeNative is NominaBridgeCommon {
    /**
     * @notice Emitted when an account deposits NOM, bridging it to Ethereum.
     */
    event Bridge(address indexed payor, address indexed to, uint256 amount);

    /**
     * @notice Emitted when NOM tokens are withdrawn for an account.
     *         If success is false, the amount is claimable by the account.
     */
    event Withdraw(address indexed payor, address indexed to, uint256 amount, bool success);

    /**
     * @notice Emitted when an account claims NOM tokens that failed to be withdrawn.
     */
    event Claimed(address indexed claimant, address indexed to, uint256 amount);

    /**
     * @notice Emitted on setup(...)
     */
    event Setup(uint64 l1ChainId, address portal, address l1Bridge, uint256 l1Deposits);

    /**
     * @notice The conversion rate from OMNI to NOM.
     */
    uint8 private constant _CONVERSION_RATE = 75;

    /**
     * @notice xcall gas limit for NominaBridgeL1.withdraw
     */
    uint64 public constant XCALL_WITHDRAW_GAS_LIMIT = 80_000;

    /**
     * @notice L1 chain id, configurable so that this contract can be used on testnets.
     */
    uint64 public l1ChainId;

    /**
     * @notice The NominaPortal contract.
     */
    INominaPortal public portal;

    /**
     * @notice Total NOM tokens deposited to NominaBridgeL1.
     *
     *         If l1Deposits == totalL1Supply, all NOM tokens are on Nomina's EVM.
     *         If l1Deposits == 0, withdraws to L1 are blocked.
     *
     *         Without validator rewards, l1Deposits == 0 would mean all NOM tokens are on Ethereum.
     *         However with validator rewards, some NOM may remain on Nomina.
     *
     *         This balance is synced on each withdraw to Nomina, and decremented on each bridge to back L1.
     */
    uint256 public l1Deposits;

    /**
     * @notice The address of the NominaBridgeL1 contract deployed to Ethereum.
     */
    address public l1Bridge;

    /**
     * @notice Track claimable amount per address. Claimable amount increases when withdraw transfer fails.
     */
    mapping(address => uint256) public claimable;

    constructor() {
        _disableInitializers();
    }

    function initialize(address owner_) external initializer {
        __Ownable_init(owner_);
    }

    function initializeV2() external reinitializer(2) {
        l1Deposits *= _CONVERSION_RATE;
    }

    /**
     * @notice Withdraw `amount` native NOM to `to`. Only callable via xcall from NominaBridgeL1.
     * @param payor     The address of the account with NOM on L1.
     * @param to        The address to receive the NOM on Nomina.
     * @param amount    The amount of NOM to withdraw.
     */
    function withdraw(address payor, address to, uint256 amount) external whenNotPaused(ACTION_WITHDRAW) {
        XTypes.MsgContext memory xmsg = portal.xmsg();

        require(msg.sender == address(portal), "NominaBridge: not xcall"); // this protects against reentrancy
        require(xmsg.sender == l1Bridge, "NominaBridge: not bridge");
        require(xmsg.sourceChainId == l1ChainId, "NominaBridge: not L1");

        l1Deposits += amount;

        (bool success,) = to.call{ value: amount }("");

        if (!success) claimable[payor] += amount;

        emit Withdraw(payor, to, amount, success);
    }

    /**
     * @notice Bridge `amount` NOM to `to` on L1.
     */
    function bridge(address to, uint256 amount) external payable whenNotPaused(ACTION_BRIDGE) {
        _bridge(to, amount);
    }

    /**
     * @dev Trigger a withdraw of `amount` NOM to `to` on L1, via xcall.
     */
    function _bridge(address to, uint256 amount) internal {
        require(to != address(0), "NominaBridge: no bridge to zero");
        require(amount > 0, "NominaBridge: amount must be > 0");
        require(amount <= l1Deposits, "NominaBridge: no liquidity");
        require(msg.value >= amount + bridgeFee(to, amount), "NominaBridge: insufficient funds");

        l1Deposits -= amount;

        // if fee is overpaid, forward excess to portal.
        // balance of this contract should continue to reflect funds bridged to L1.
        portal.xcall{ value: msg.value - amount }(
            l1ChainId,
            ConfLevel.Finalized,
            l1Bridge,
            abi.encodeCall(NominaBridgeL1.withdraw, (to, amount)),
            XCALL_WITHDRAW_GAS_LIMIT
        );

        emit Bridge(msg.sender, to, amount);
    }

    /**
     * @notice Return the xcall fee required to bridge `amount` to `to`.
     */
    function bridgeFee(address to, uint256 amount) public view returns (uint256) {
        return portal.feeFor(l1ChainId, abi.encodeCall(NominaBridgeL1.withdraw, (to, amount)), XCALL_WITHDRAW_GAS_LIMIT);
    }

    /**
     * @notice Claim NOM tokens that failed to be withdrawn via xmsg from NominaBridgeL1.
     * @dev We require this be made by xcall, because an account on Nomina may not be authorized to
     *      claim for that address on L1. Consider the case wherein the address of the L1 contract that initiated the
     *      withdraw is the same as the address of a contract on Nomina deployed and owned by a malicious actor.
     */
    function claim(address to) external whenNotPaused(ACTION_WITHDRAW) {
        XTypes.MsgContext memory xmsg = portal.xmsg();

        require(msg.sender == address(portal), "NominaBridge: not xcall");
        require(xmsg.sourceChainId == l1ChainId, "NominaBridge: not L1");
        require(to != address(0), "NominaBridge: no claim to zero");

        address claimant = xmsg.sender;
        require(claimable[claimant] > 0, "NominaBridge: nothing to claim");

        uint256 amount = claimable[claimant];
        claimable[claimant] = 0;

        (bool success,) = to.call{ value: amount }("");
        require(success, "NominaBridge: transfer failed");

        emit Claimed(claimant, to, amount);
    }

    //////////////////////////////////////////////////////////////////////////////
    //                          Admin functions                                 //
    //////////////////////////////////////////////////////////////////////////////

    /**
     * @notice Setup core contract parameters, done by owner immediately after pre-deployment.
     * @param l1ChainId_    The chain id of the L1 network.
     * @param portal_       The address of the NominaPortal contract.
     * @param l1Bridge_     The address of the L1 NominaBridge contract.
     * @param l1Deposits_   The number of tokens deposited to L1 bridge contract at setup
     *                      (to account for genesis prefunds)
     */
    function setup(uint64 l1ChainId_, address portal_, address l1Bridge_, uint256 l1Deposits_) external onlyOwner {
        l1ChainId = l1ChainId_;
        portal = INominaPortal(portal_);
        l1Bridge = l1Bridge_;
        l1Deposits = l1Deposits_;
        emit Setup(l1ChainId_, portal_, l1Bridge_, l1Deposits_);
    }
}
