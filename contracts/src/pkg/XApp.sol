// SPDX-License-Identifier: GPL-3.0-only
pragma solidity ^0.8.12;

import { IOmniPortal } from "../interfaces/IOmniPortal.sol";
import { XTypes } from "../libraries/XTypes.sol";

/**
 * @title XApp
 * @dev Base contract for Omni cross-chain applications
 */
contract XApp {
    /// @dev The OmniPortal contract
    IOmniPortal internal immutable omni;

    /// @dev Transient storage for the current xmsg
    XTypes.Msg internal xmsg;

    /// @dev Read current xmsg into storage before execution, delete it afterwards
    modifier setXMsg() {
        xmsg = omni.xmsg();
        _;
        delete xmsg;
    }

    /// @dev Only allow xcalls to call this function
    modifier onlyXCall() {
        require(isXCall(), "XApp: not xcall");
        _;
    }

    constructor(address _omni) {
        omni = IOmniPortal(_omni);
    }

    /// @dev Return true if the current call is an xcall from the OmniPortal
    function isXCall() internal view returns (bool) {
        return omni.isXCall() && msg.sender == address(omni);
    }

    /// @dev Calculate the fee for calling a contract on another chain.
    ///      Uses OmniPortal.XMSG_DEFAULT_GAS_LIMIT
    /// @return fee The fee, denominated in wei
    function feeFor(uint64 destChainId, bytes memory data) internal view returns (uint256) {
        return omni.feeFor(destChainId, data);
    }

    /// @dev Calculate the fee for calling a contract on another chain.
    ///      Uses OmniPortal.XMSG_DEFAULT_GAS_LIMIT
    /// @return fee The fee, denominated in wei
    function feeFor(uint64 destChainId, bytes memory data, uint64 gasLimit) internal view returns (uint256) {
        return omni.feeFor(destChainId, data, gasLimit);
    }

    /// @dev Call a contract on another chain. Uses OmniPortal.XMSG_DEFAULT_GAS_LIMIT
    /// @return fee The fee for the xcall
    function xcall(uint64 destChainId, address to, bytes memory data) internal returns (uint256) {
        uint256 fee = omni.feeFor(destChainId, data);
        omni.xcall{ value: fee }(destChainId, to, data);
        return fee;
    }

    /// @dev Call a contract on another chain
    /// @return fee The fee, denominated in wei
    function xcall(uint64 destChainId, address to, bytes memory data, uint64 gasLimit) internal returns (uint256) {
        uint256 fee = omni.feeFor(destChainId, data, gasLimit);
        omni.xcall{ value: fee }(destChainId, to, data, gasLimit);
        return fee;
    }
}
