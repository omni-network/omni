// SPDX-License-Identifier: GPL-3.0-only
pragma solidity ^0.8.12;

import { IOmniPortal } from "../interfaces/IOmniPortal.sol";
import { XTypes } from "../libraries/XTypes.sol";

/**
 * @title XApp
 * @dev Base contract for Omni cross-chain applications
 */
contract XApp {
    /// @notice The OmniPortal contract
    IOmniPortal internal immutable omni;

    /// @notice Transient storage for the current xmsg
    XTypes.MsgShort internal xmsg;

    /// @notice Read current xmsg into storage before execution, delete it afterwards
    modifier xrecv() {
        xmsg = omni.xmsg();
        _;
        delete xmsg;
    }

    constructor(address _omni) {
        omni = IOmniPortal(_omni);
    }

    /**
     * @notice Return true if the current call is an xcall from the OmniPortal
     */
    function isXCall() internal view returns (bool) {
        return omni.isXCall() && msg.sender == address(omni);
    }

    /**
     * @notice Returns the fee for calling a contract on another chain. Uses OmniPortal.XMSG_DEFAULT_GAS_LIMIT
     */
    function feeFor(uint64 destChainId, bytes memory data) internal view returns (uint256) {
        return omni.feeFor(destChainId, data);
    }

    /**
     * @notice Retruns the fee for calling a contract on another chain, with the specified gas limit
     */
    function feeFor(uint64 destChainId, bytes memory data, uint64 gasLimit) internal view returns (uint256) {
        return omni.feeFor(destChainId, data, gasLimit);
    }

    /**
     * @notice Call a contract on another chain. Uses OmniPortal.XMSG_DEFAULT_GAS_LIMIT
     * @return fee The fee for the xcall
     */
    function xcall(uint64 destChainId, address to, bytes memory data) internal returns (uint256) {
        uint256 fee = omni.feeFor(destChainId, data);
        require(address(this).balance >= fee || msg.value >= fee, "XApp: insufficient funds");
        omni.xcall{ value: fee }(destChainId, to, data);
        return fee;
    }

    /**
     * @notice Call a contract on another chain, with the specified gas limit
     * @return fee The fee, denominated in wei
     */
    function xcall(uint64 destChainId, address to, bytes memory data, uint64 gasLimit) internal returns (uint256) {
        uint256 fee = omni.feeFor(destChainId, data, gasLimit);
        require(address(this).balance >= fee || msg.value >= fee, "XApp: insufficient funds");
        omni.xcall{ value: fee }(destChainId, to, data, gasLimit);
        return fee;
    }
}
