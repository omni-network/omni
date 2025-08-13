// SPDX-License-Identifier: GPL-3.0-only
pragma solidity ^0.8.24;

import { XTypes } from "src/libraries/XTypes.sol";
import { ConfLevel } from "src/libraries/ConfLevel.sol";
import { NominaPortalConstants } from "src/xchain/nomina/NominaPortalConstants.sol";
import { IFeeOracle } from "src/interfaces/IFeeOracle.sol";
import { INominaPortal } from "src/interfaces/nomina/INominaPortal.sol";
import { MockFeeOracle } from "../MockFeeOracle.sol";

/**
 * @title MockPortal
 * @notice A mock portal, used for testing.
 *         - Matches real portal functionality for user facing functions (xcall, feeFor, and xmsg),
 *           so that user unit tests consume gas as expected.
 *         - Non-user facing functions & state are stubbed.
 *         - Provides a mockXCall function for testing xcall execution.
 */
contract MockPortal is INominaPortal, NominaPortalConstants {
    uint64 public immutable chainId;
    uint64 public immutable nominaChainId;

    uint64 public xmsgMaxGasLimit = 5_000_000;
    uint64 public xmsgMinGasLimit = 21_000;
    uint16 public xmsgMaxDataSize = 20_000;
    uint16 public xreceiptMaxErrorSize = 256;

    address public feeOracle;

    bool public paused;

    mapping(uint64 => mapping(uint64 => uint64)) public outXMsgOffset;
    mapping(uint64 => mapping(uint64 => uint64)) public inXMsgOffset;
    mapping(uint64 => mapping(uint64 => uint64)) public inXBlockOffset;
    mapping(uint64 => bool) public isSupportedShard;
    mapping(uint64 => bool) public isSupportedDest;

    XTypes.MsgContext internal _xmsg;

    constructor() {
        chainId = uint64(block.chainid);
        nominaChainId = 166;
        feeOracle = address(new MockFeeOracle(1 gwei));
    }

    //////////////////////////////////////////////////////////////////////////////
    //                      Standard Portal Functions                           //
    //////////////////////////////////////////////////////////////////////////////

    function xcall(uint64 destChainId, uint8 conf, address to, bytes calldata data, uint64 gasLimit) external payable {
        require(gasLimit <= xmsgMaxGasLimit, "NominaPortal: gasLimit too high");
        require(gasLimit >= xmsgMinGasLimit, "NominaPortal: gasLimit too low");
        require(destChainId != chainId, "NominaPortal: unsupported dest");
        require(destChainId != BroadcastChainId, "NominaPortal: unsupported dest");
        require(to != VirtualPortalAddress, "NominaPortal: no portal xcall");

        uint256 fee = feeFor(destChainId, data, gasLimit);
        require(msg.value >= fee, "NominaPortal: insufficient fee");

        uint64 shardId = uint64(conf);
        outXMsgOffset[destChainId][shardId] += 1;

        emit XMsg(destChainId, shardId, outXMsgOffset[destChainId][shardId], msg.sender, to, data, gasLimit, fee);
    }

    function feeFor(uint64 destChainId, bytes calldata data, uint64 gasLimit) public view returns (uint256) {
        return IFeeOracle(feeOracle).feeFor(destChainId, data, gasLimit);
    }

    function xmsg() external view returns (XTypes.MsgContext memory) {
        return _xmsg;
    }

    function isXCall() external view returns (bool) {
        return _xmsg.sourceChainId != 0;
    }

    // stub
    function network() external pure returns (XTypes.Chain[] memory) {
        return new XTypes.Chain[](0);
    }

    //////////////////////////////////////////////////////////////////////////////
    //                              Portal Mocks                                //
    //////////////////////////////////////////////////////////////////////////////

    /// @dev Execute a mock xcall, no gas limit. Forwards revert on call fails
    function mockXCall(uint64 sourceChainId, address sender, address to, bytes calldata data)
        public
        returns (uint256 gasUsed)
    {
        _xmsg = XTypes.MsgContext({ sourceChainId: sourceChainId, sender: sender });

        gasUsed = gasleft();
        (bool success, bytes memory returnData) = to.call(data);
        gasUsed = gasUsed - gasleft();

        delete _xmsg;

        if (!success) {
            assembly {
                revert(add(returnData, 32), mload(returnData))
            }
        }

        return gasUsed;
    }

    /// @dev Execute a mock xcall, custom gas limit. Forwards revert on call fails. Reverts on out of gas.
    function mockXCall(uint64 sourceChainId, address sender, address to, bytes calldata data, uint64 gasLimit)
        public
        returns (uint256 gasUsed)
    {
        require(gasLimit <= xmsgMaxGasLimit, "NominaPortal: gasLimit too high");
        require(gasLimit >= xmsgMinGasLimit, "NominaPortal: gasLimit too low");

        _xmsg = XTypes.MsgContext({ sourceChainId: sourceChainId, sender: sender });

        gasUsed = gasleft();
        (bool success, bytes memory returnData) = to.call{ gas: gasLimit }(data);
        gasUsed = gasUsed - gasleft();

        delete _xmsg;

        if (!success && gasUsed >= gasLimit) revert("MockPortal: out of gas");
        if (!success) {
            assembly {
                revert(add(returnData, 32), mload(returnData))
            }
        }

        return gasUsed;
    }

    //////////////////////////////////////////////////////////////////////////////
    //                              Stubs                                       //
    //////////////////////////////////////////////////////////////////////////////

    function xsubmit(XTypes.Submission calldata submit) external override { }

    function pause(bool status) external {
        paused = status;
    }

    function isPaused(bytes32) external view returns (bool) {
        return paused;
    }

    function isPaused(bytes32, uint64) external view returns (bool) {
        return paused;
    }

    function isPaused() external view returns (bool) {
        return paused;
    }
}
