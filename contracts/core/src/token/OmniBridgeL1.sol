// SPDX-License-Identifier: GPL-3.0-only
pragma solidity 0.8.24;

import { IERC20 } from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import { OmniBridgeCommon } from "./OmniBridgeCommon.sol";
import { IOmniPortal } from "../interfaces/IOmniPortal.sol";
import { XTypes } from "../libraries/XTypes.sol";
import { ConfLevel } from "../libraries/ConfLevel.sol";
import { Predeploys } from "../libraries/Predeploys.sol";
import { OmniBridgeNative } from "./OmniBridgeNative.sol";

/**
 * @title OmniBridgeL1
 * @notice The Ethereum side of Omni's native token bridge. Partner to OmniBridgeNative, which is
 *         deployed to Omni's EVM.
 */
contract OmniBridgeL1 is OmniBridgeCommon {
    /**
     * @notice Emitted when an account deposits OMNI, bridging it to Ethereum.
     */
    event Bridge(address indexed payor, address indexed to, uint256 amount);

    /**
     * @notice Emitted when OMNI tokens are withdrawn for an account.
     */
    event Withdraw(address indexed to, uint256 amount);

    /**
     * @notice xcall gas limit for OmniBridgeNative.withdraw
     */
    uint64 public constant XCALL_WITHDRAW_GAS_LIMIT = 80_000;

    /**
     * @notice The OMNI token contract.
     */
    IERC20 public immutable token;

    /**
     * @notice The OmniPortal contract.
     */
    IOmniPortal public portal;

    constructor(address token_) {
        token = IERC20(token_);
        _disableInitializers();
    }

    function initialize(address owner_, address portal_) external initializer {
        require(portal_ != address(0), "OmniBridge: no zero addr");
        __Ownable_init(owner_);
        portal = IOmniPortal(portal_);
    }

    /**
     * @notice Withdraw `amount` L1 OMNI to `to`. Only callable via xcall from OmniBridgeNative.
     * @dev Omni native <> L1 bridge accounting rules ensure that this contract will always
     *     have enough balance to cover the withdrawal.
     */
    function withdraw(address to, uint256 amount) external whenNotPaused(ACTION_WITHDRAW) {
        XTypes.MsgContext memory xmsg = portal.xmsg();

        require(msg.sender == address(portal), "OmniBridge: not xcall");
        require(xmsg.sender == Predeploys.OmniBridgeNative, "OmniBridge: not bridge");
        require(xmsg.sourceChainId == portal.omniChainId(), "OmniBridge: not omni portal");

        token.transfer(to, amount);

        emit Withdraw(to, amount);
    }

    /**
     * @notice Bridge `amount` OMNI to `to` on Omni's EVM.
     */
    function bridge(address to, uint256 amount) external payable whenNotPaused(ACTION_BRIDGE) {
        _bridge(msg.sender, to, amount);
    }

    /**
     * @dev Trigger a withdraw of `amount` OMNI to `to` on Omni's EVM, via xcall.
     */
    function _bridge(address payor, address to, uint256 amount) internal {
        require(amount > 0, "OmniBridge: amount must be > 0");
        require(to != address(0), "OmniBridge: no bridge to zero");

        uint64 omniChainId = portal.omniChainId();
        bytes memory xcalldata = abi.encodeCall(OmniBridgeNative.withdraw, (payor, to, amount));

        require(
            msg.value >= portal.feeFor(omniChainId, xcalldata, XCALL_WITHDRAW_GAS_LIMIT), "OmniBridge: insufficient fee"
        );
        require(token.transferFrom(payor, address(this), amount), "OmniBridge: transfer failed");

        portal.xcall{ value: msg.value }(
            omniChainId, ConfLevel.Finalized, Predeploys.OmniBridgeNative, xcalldata, XCALL_WITHDRAW_GAS_LIMIT
        );

        emit Bridge(payor, to, amount);
    }

    /**
     * @notice Return the xcall fee required to bridge `amount` to `to`.
     */
    function bridgeFee(address payor, address to, uint256 amount) public view returns (uint256) {
        return portal.feeFor(
            portal.omniChainId(),
            abi.encodeCall(OmniBridgeNative.withdraw, (payor, to, amount)),
            XCALL_WITHDRAW_GAS_LIMIT
        );
    }
}
