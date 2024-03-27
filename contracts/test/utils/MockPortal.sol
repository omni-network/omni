// SPDX-License-Identifier: GPL-3.0-only
pragma solidity ^0.8.12;

import { XTypes } from "../../src/libraries/XTypes.sol";
import { OmniPortalConstants } from "../../src/protocol/OmniPortalConstants.sol";
import { IFeeOracle } from "../../src/interfaces/IFeeOracle.sol";
import { MockFeeOracle } from "./MockFeeOracle.sol";

/**
 * @title MockPortal
 * @notice A mock portal, used for testing.
 *         - Matches real portal functionality for user facing functions (xcall, feeFor, and xmsg),
 *           so that user unit tests consume gas as expected.
 *         - All non-user facing functions & state are not included.
 *         - Provides a mockXCall function for testing xcall execution.
 */
contract MockPortal is OmniPortalConstants {
    event XMsg(
        uint64 indexed destChainId, uint64 indexed streamOffset, address sender, address to, bytes data, uint64 gasLimit
    );

    event XReceipt(
        uint64 indexed sourceChainId, uint64 indexed streamOffset, uint256 gasUsed, address relayer, bool success
    );

    uint64 public immutable chainId;

    address public feeOracle;

    mapping(uint64 => uint64) public outXStreamOffset;

    XTypes.MsgShort private _currentXmsg;

    constructor() {
        chainId = uint64(block.chainid);
        feeOracle = address(new MockFeeOracle(1 gwei));
    }

    //////////////////////////////////////////////////////////////////////////////
    //                      Standard Portal Functions                           //
    //////////////////////////////////////////////////////////////////////////////

    function xcall(uint64 destChainId, address to, bytes calldata data) external payable {
        _xcall(destChainId, msg.sender, to, data, XMSG_DEFAULT_GAS_LIMIT);
    }

    function xcall(uint64 destChainId, address to, bytes calldata data, uint64 gasLimit) external payable {
        _xcall(destChainId, msg.sender, to, data, gasLimit);
    }

    function feeFor(uint64 destChainId, bytes calldata data) public view returns (uint256) {
        return IFeeOracle(feeOracle).feeFor(destChainId, data, XMSG_DEFAULT_GAS_LIMIT);
    }

    function feeFor(uint64 destChainId, bytes calldata data, uint64 gasLimit) public view returns (uint256) {
        return IFeeOracle(feeOracle).feeFor(destChainId, data, gasLimit);
    }

    function _xcall(uint64 destChainId, address sender, address to, bytes calldata data, uint64 gasLimit) private {
        require(msg.value >= feeFor(destChainId, data, gasLimit), "OmniPortal: insufficient fee");
        require(gasLimit <= XMSG_MAX_GAS_LIMIT, "OmniPortal: gasLimit too high");
        require(gasLimit >= XMSG_MIN_GAS_LIMIT, "OmniPortal: gasLimit too low");
        require(destChainId != chainId, "OmniPortal: no same-chain xcall");

        outXStreamOffset[destChainId] += 1;

        emit XMsg(destChainId, outXStreamOffset[destChainId], sender, to, data, gasLimit);
    }

    function xmsg() external view returns (XTypes.MsgShort memory) {
        return _currentXmsg;
    }

    function isXCall() external view returns (bool) {
        return _currentXmsg.sourceChainId != 0;
    }

    //////////////////////////////////////////////////////////////////////////////
    //                              Portal Mocks                                //
    //////////////////////////////////////////////////////////////////////////////

    /// @notice Execute a mock xcall, default gas limit. Passes the revert for call fails or too low gas limit
    function mockXCall(uint64 sourceChainId, address to, bytes calldata data) external {
        mockXCall(sourceChainId, msg.sender, to, data, XMSG_DEFAULT_GAS_LIMIT);
    }

    /// @dev Execute a mock xcall, custom gas limit. Passes the revert for call fails or too low gas limit
    function mockXCall(uint64 sourceChainId, address sender, address to, bytes calldata data, uint64 gasLimit) public {
        _mockXCall(sourceChainId, sender, to, data, gasLimit);
    }

    /// @dev Execute a mock xcall, custom gas limit, passing the revert message if the call fails
    function _mockXCall(uint64 sourceChainId, address sender, address to, bytes calldata data, uint64 gasLimit)
        private
    {
        require(gasLimit <= XMSG_MAX_GAS_LIMIT, "OmniPortal: gasLimit too high");
        require(gasLimit >= XMSG_MIN_GAS_LIMIT, "OmniPortal: gasLimit too low");

        _currentXmsg = XTypes.MsgShort({ sourceChainId: sourceChainId, sender: sender });

        uint256 gasUsed = gasleft();
        (bool success, bytes memory returnData) = to.call{ gas: gasLimit }(data);
        gasUsed = gasUsed - gasleft();

        if (!success && gasUsed >= gasLimit) revert("MockPortal: out of gas");
        if (!success) {
            assembly {
                revert(add(returnData, 32), mload(returnData))
            }
        }
    }
}
