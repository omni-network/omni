// SPDX-License-Identifier: GPL-3.0-only
pragma solidity 0.8.24;

import { IERC20 } from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import { INomina } from "src/interfaces/nomina/INomina.sol";
import { NominaBridgeCommon } from "./NominaBridgeCommon.sol";
import { IOmniPortal } from "src/interfaces/IOmniPortal.sol";
import { XTypes } from "src/libraries/XTypes.sol";
import { ConfLevel } from "src/libraries/ConfLevel.sol";
import { Predeploys } from "src/libraries/nomina/Predeploys.sol";
import { NominaBridgeNative } from "./NominaBridgeNative.sol";

/**
 * @title NominaBridgeL1
 * @notice The Ethereum side of Nomina's native token bridge. Partner to NominaBridgeNative, which is
 *         deployed to Nomina's EVM.
 */
contract NominaBridgeL1 is NominaBridgeCommon {
    /**
     * @notice Emitted when an account deposits NOM, bridging it to Ethereum.
     */
    event Bridge(address indexed payor, address indexed to, uint256 amount);

    /**
     * @notice Emitted when NOM tokens are withdrawn for an account.
     */
    event Withdraw(address indexed to, uint256 amount);

    /**
     * @notice xcall gas limit for NominaBridgeNative.withdraw
     */
    uint64 public constant XCALL_WITHDRAW_GAS_LIMIT = 80_000;

    /**
     * @notice The OMNI token contract.
     */
    IERC20 public immutable OMNI;

    /**
     * @notice The NOM token contract.
     */
    IERC20 public immutable NOMINA;

    /**
     * @notice The OmniPortal contract.
     */
    IOmniPortal public portal;

    constructor(address omni, address nomina) {
        OMNI = IERC20(omni);
        NOMINA = IERC20(nomina);
        _disableInitializers();
    }

    function initialize(address owner_, address portal_) external initializer {
        require(portal_ != address(0), "NominaBridge: no zero addr");
        __Ownable_init(owner_);
        portal = IOmniPortal(portal_);
    }

    function initializeV2() external reinitializer(2) {
        OMNI.approve(address(NOMINA), type(uint256).max);
        INomina(address(NOMINA)).convert(address(this), OMNI.balanceOf(address(this)));
    }

    /**
     * @notice Withdraw `amount` L1 NOM to `to`. Only callable via xcall from NominaBridgeNative.
     * @dev Nomina native <> L1 bridge accounting rules ensure that this contract will always
     *     have enough balance to cover the withdrawal.
     */
    function withdraw(address to, uint256 amount) external whenNotPaused(ACTION_WITHDRAW) {
        XTypes.MsgContext memory xmsg = portal.xmsg();

        require(msg.sender == address(portal), "NominaBridge: not xcall");
        require(xmsg.sender == Predeploys.NominaBridgeNative, "NominaBridge: not bridge");
        require(xmsg.sourceChainId == portal.omniChainId(), "NominaBridge: not omni portal");

        NOMINA.transfer(to, amount);

        emit Withdraw(to, amount);
    }

    /**
     * @notice Bridge `amount` NOM to `to` on Nomina's EVM.
     */
    function bridge(address to, uint256 amount) external payable whenNotPaused(ACTION_BRIDGE) {
        _bridge(msg.sender, to, amount);
    }

    /**
     * @dev Trigger a withdraw of `amount` NOM to `to` on Nomina's EVM, via xcall.
     */
    function _bridge(address payor, address to, uint256 amount) internal {
        require(amount > 0, "NominaBridge: amount must be > 0");
        require(to != address(0), "NominaBridge: no bridge to zero");

        uint64 omniChainId = portal.omniChainId();
        bytes memory xcalldata = abi.encodeCall(NominaBridgeNative.withdraw, (payor, to, amount));

        require(
            msg.value >= portal.feeFor(omniChainId, xcalldata, XCALL_WITHDRAW_GAS_LIMIT),
            "NominaBridge: insufficient fee"
        );
        require(NOMINA.transferFrom(payor, address(this), amount), "NominaBridge: transfer failed");

        portal.xcall{ value: msg.value }(
            omniChainId, ConfLevel.Finalized, Predeploys.NominaBridgeNative, xcalldata, XCALL_WITHDRAW_GAS_LIMIT
        );

        emit Bridge(payor, to, amount);
    }

    /**
     * @notice Return the xcall fee required to bridge `amount` to `to`.
     */
    function bridgeFee(address payor, address to, uint256 amount) public view returns (uint256) {
        return portal.feeFor(
            portal.omniChainId(),
            abi.encodeCall(NominaBridgeNative.withdraw, (payor, to, amount)),
            XCALL_WITHDRAW_GAS_LIMIT
        );
    }
}
