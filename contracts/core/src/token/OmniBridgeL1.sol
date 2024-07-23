// SPDX-License-Identifier: GPL-3.0-only
pragma solidity 0.8.24;

import { OwnableUpgradeable } from "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import { IERC20 } from "@openzeppelin/contracts/token/ERC20/IERC20.sol";
import { IOmniPortal } from "../interfaces/IOmniPortal.sol";
import { XTypes } from "../libraries/XTypes.sol";
import { ConfLevel } from "../libraries/ConfLevel.sol";
import { Predeploys } from "../libraries/Predeploys.sol";
import { OmniBridgeNative } from "./OmniBridgeNative.sol";

/**
 * @title OmniBridgeL1
 * @notice The Ethereum side of Omni's native token bridge. Partner to OmniBridgeNative, which is
 *         deployed to Omni's EVM.
 * @dev We currently do now have any onlyOwner methods, but we inherit from OwnableUpgradeable let us
 *      add them in the future.
 */
contract OmniBridgeL1 is OwnableUpgradeable {
    /**
     * @notice Emitted when an account deposits OMNI, bridging it to Ethereum.
     */
    event Bridge(address indexed payor, address indexed to, uint256 amount);

    /**
     * @notice Emitted when OMNI tokens are withdrawn for an account.
     */
    event Withdraw(address indexed to, uint256 amount);

    /**
     * @notice Total supply of OMNI tokens minted on L1.
     */
    uint256 public constant totalL1Supply = 100_000_000 * 10 ** 18;

    /**
     * @notice xcall gas limit for OmniBridgeNative.withdraw
     */
    uint64 public constant XCALL_WITHDRAW_GAS_LIMIT = 75_000;

    /**
     * @notice The OMNI token contract.
     */
    IERC20 public immutable token;

    /**
     * @notice The OmniPortal contract.
     */
    IOmniPortal public omni;

    constructor(address token_) {
        token = IERC20(token_);
        _disableInitializers();
    }

    function initialize(address owner_, address omni_) external initializer {
        __Ownable_init(owner_);
        omni = IOmniPortal(omni_);
    }

    /**
     * @notice Withdraw `amount` L1 OMNI to `to`. Onyl callable via xcall from OmniBridgeNative.
     * @dev Omni native <> L1 bridge accounting rules ensure that this contract will always
     *     have enough balance to cover the withdrawal.
     */
    function withdraw(address to, uint256 amount) external {
        XTypes.MsgContext memory xmsg = omni.xmsg();

        require(msg.sender == address(omni), "OmniBridge: not xcall");
        require(xmsg.sender == Predeploys.OmniBridgeNative, "OmniBridge: not bridge");
        require(xmsg.sourceChainId == omni.omniChainId(), "OmniBridge: not omni");

        token.transfer(to, amount);

        emit Withdraw(to, amount);
    }

    /**
     * @notice Bridge `amount` OMNI to `to` on Omni's EVM.
     */
    function bridge(address to, uint256 amount) external payable {
        _bridge(msg.sender, to, amount);
    }

    /**
     * @dev Trigger a withdraw of `amount` OMNI to `to` on Omni's EVM, via xcall.
     */
    function _bridge(address payor, address to, uint256 amount) internal {
        require(amount > 0, "OmniBridge: amount must be > 0");
        require(to != address(0), "OmniBridge: no bridge to zero");
        require(msg.value == bridgeFee(payor, to, amount), "OmniBridge: incorrect fee");
        require(token.transferFrom(payor, address(this), amount), "OmniBridge: transfer failed");

        omni.xcall{ value: msg.value }(
            omni.omniChainId(),
            ConfLevel.Finalized,
            Predeploys.OmniBridgeNative,
            abi.encodeWithSelector(OmniBridgeNative.withdraw.selector, payor, to, amount),
            XCALL_WITHDRAW_GAS_LIMIT
        );

        emit Bridge(payor, to, amount);
    }

    /**
     * @notice Return the xcall fee required to bridge `amount` to `to`.
     */
    function bridgeFee(address payor, address to, uint256 amount) public view returns (uint256) {
        return omni.feeFor(
            omni.omniChainId(),
            abi.encodeWithSelector(OmniBridgeNative.withdraw.selector, payor, to, amount),
            XCALL_WITHDRAW_GAS_LIMIT
        );
    }
}
