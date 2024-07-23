// SPDX-License-Identifier: GPL-3.0-only
pragma solidity ^0.8.12;

import { IOmniPortal } from "../interfaces/IOmniPortal.sol";
import { XTypes } from "../libraries/XTypes.sol";
import { ConfLevel } from "../libraries/ConfLevel.sol";

/**
 * @title XApp
 * @notice Base contract for Omni cross-chain applications
 */
contract XApp {
    /**
     * @notice Emitted when the OmniPortal contract address is set
     */
    event OmniPortalSet(address omni);

    /**
     * @notice Emitted when the default confirmation level is set
     * @param conf  Confirmation level
     */
    event DefaultConfLevelSet(uint8 conf);

    /**
     * @notice The OmniPortal contract
     */
    IOmniPortal public omni;

    /**
     * @notice Default confirmation level for xcalls
     */
    uint8 public defaultConfLevel;

    /**
     * @notice Transient storage for the current xmsg
     */
    XTypes.MsgContext internal xmsg;

    /**
     * @notice Read current xmsg into storage before execution, delete it afterwards
     */
    modifier xrecv() {
        xmsg = omni.xmsg();
        _;
        delete xmsg;
    }

    // TODO: write XAppUpgradeable that follows initialize pattern
    constructor(address omni_, uint8 defaultConfLevel_) {
        _setOmniPortal(omni_);
        _setDefaultConfLevel(defaultConfLevel_);
    }

    /**
     * @notice Return true if the current call is an xcall from the OmniPortal
     */
    function isXCall() internal view returns (bool) {
        return omni.isXCall() && msg.sender == address(omni);
    }

    /**
     * @notice Retruns the fee for calling a contract on another chain, with the specified gas limit
     */
    function feeFor(uint64 destChainId, bytes memory data, uint64 gasLimit) internal view returns (uint256) {
        return omni.feeFor(destChainId, data, gasLimit);
    }

    /**
     * @notice Call a contract on another. (Default ConfLevel)
     * @param destChainId   Destination chain ID
     * @param to            Address of contract to call on destination chain
     * @param data          ABI Encoded function calldata
     * @param gasLimit      Execution gas limit, enforced on destination chain
     */
    function xcall(uint64 destChainId, address to, bytes memory data, uint64 gasLimit) internal returns (uint256) {
        uint256 fee = omni.feeFor(destChainId, data, gasLimit);
        require(address(this).balance >= fee || msg.value >= fee, "XApp: insufficient funds");
        omni.xcall{ value: fee }(destChainId, defaultConfLevel, to, data, gasLimit);
        return fee;
    }

    /**
     * @notice Call a contract on another chain. (Explicit ConfLevel)
     * @param destChainId   Destination chain ID
     * @param conf          Confirmation level
     * @param to            Address of contract to call on destination chain
     * @param data          ABI Encoded function calldata
     * @param gasLimit      Execution gas limit, enforced on destination chain
     */
    function xcall(uint64 destChainId, uint8 conf, address to, bytes memory data, uint64 gasLimit)
        internal
        returns (uint256)
    {
        uint256 fee = omni.feeFor(destChainId, data, gasLimit);
        require(address(this).balance >= fee || msg.value >= fee, "XApp: insufficient funds");
        omni.xcall{ value: fee }(destChainId, conf, to, data, gasLimit);
        return fee;
    }

    /**
     * @notice Set the default confirmation level for xcalls
     * @param conf  Confirmation level
     */
    function _setDefaultConfLevel(uint8 conf) internal {
        require(ConfLevel.isValid(conf), "XApp: invalid conf level");
        defaultConfLevel = conf;
        emit DefaultConfLevelSet(conf);
    }

    /**
     * @notice Set the OmniPortal contract address
     * @param _omni    The OmniPortal contract address
     */
    function _setOmniPortal(address _omni) internal {
        require(_omni != address(0), "XApp: no zero omni");
        omni = IOmniPortal(_omni);
        emit OmniPortalSet(_omni);
    }
}
