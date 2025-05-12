// SPDX-License-Identifier: GPL-3.0-only
pragma solidity ^0.8.12;

import { IOmniPortal } from "core/src/interfaces/IOmniPortal.sol";
import { XTypes } from "core/src/libraries/XTypes.sol";
import { ConfLevel } from "core/src/libraries/ConfLevel.sol";

/**
 * @title XAppBase
 * @notice Shared XApp storage and logic.
 */
abstract contract XAppBase {
    /**
     * @notice Deprecated `_omni` variable in favor of an immutable equivalent.
     * @dev This variable is used to avoid a storage slot collision.
     */
    IOmniPortal private _deprecatedOmni;

    /**
     * @notice The OmniPortal contract
     */
    IOmniPortal public immutable omni;

    /**
     * @notice Deprecated `_defaultConfLevel` variable as we do not use the xcall method that retrieves it.
     * @dev This variable is used to avoid a storage slot collision.
     */
    uint8 private _deprecatedDefaultConfLevel;

    /**
     * @notice Transient storage for the current xmsg
     */
    XTypes.MsgContext internal xmsg;

    /**
     * @notice Read current xmsg into storage before execution, delete it afterwards
     * @dev If `omni` is not set, xmsg is not modified.
     */
    modifier xrecv() {
        if (address(omni) != address(0)) {
            xmsg = omni.xmsg();
            _;
            delete xmsg;
        } else {
            _;
        }
    }

    constructor(address omni_) {
        omni = IOmniPortal(omni_);
    }

    /**
     * @notice Returns the fee for calling a contract on another chain, with the specified gas limit
     */
    function feeFor(uint64 destChainId, bytes memory data, uint64 gasLimit) internal view returns (uint256) {
        return omni.feeFor(destChainId, data, gasLimit);
    }

    /**
     * @notice Call a contract on another chain. (Only Finalized ConfLevel)
     * @dev This function pays the xcall fee to the OmniPortal from the
     *      contract's balance. If you would prefer charge callers for the
     *      fee, you must check msg.value.
     *      Ex:
     *          uint256 fee = xcall(...)
     *          require(msg.value >= fee)
     *
     * @param destChainId   Destination chain ID
     * @param to            Address of contract to call on destination chain
     * @param data          ABI Encoded function calldata
     * @param gasLimit      Execution gas limit, enforced on destination chain
     */
    function xcall(uint64 destChainId, address to, bytes memory data, uint64 gasLimit) internal returns (uint256) {
        uint256 fee = omni.feeFor(destChainId, data, gasLimit);
        require(address(this).balance >= fee, "XApp: insufficient funds");
        omni.xcall{ value: fee }(destChainId, ConfLevel.Finalized, to, data, gasLimit);
        return fee;
    }
}
