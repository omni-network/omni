// SPDX-License-Identifier: GPL-3.0-only
pragma solidity ^0.8.12;

import { XTypes } from "../../src/libraries/XTypes.sol";
import { OmniPortalConstants } from "../../src/protocol/OmniPortalConstants.sol";

/**
 * @title MockPortal
 * @dev A mock OmniPortal contract, for testing purposes.
 */
contract MockPortal is OmniPortalConstants {
    /// @notice The current XMsg being executed, exposed via xmsg() getter
    XTypes.Msg private _currentXMsg;

    /// @notice Returns current xmsg, same as OmniPortal.xmsg()
    function xmsg() external view returns (XTypes.Msg memory) {
        return _currentXMsg;
    }

    /// @notice Stubbed feeFor, default gas limit
    function feeFor(uint64, /* destChainId */ bytes calldata /* data */ ) external pure returns (uint256) {
        return 1 gwei;
    }

    /// @notice Stubbed feeFor, custom gas limit
    function feeFor(uint64, /* destChainId */ bytes calldata, /* data */ uint64 /* gasLimit */ )
        external
        pure
        returns (uint256)
    {
        return 1 gwei;
    }

    /// @notice Stubbed xcall, default gas limit
    function xcall(uint64, /* destChainId */ address, /* to */ bytes calldata /* data */ ) external payable { }

    /// @notice Stubbed xcall, custom gas limit
    function xcall(uint64, /* destChainId */ address, /* to */ bytes calldata, /* data */ uint64 /* gasLimit */ )
        external
        payable
    { }

    /// @notice Returns true if the current call is an xcall, same as OmniPortal.isXCall()
    function isXCall() external view returns (bool) {
        return _currentXMsg.sourceChainId != 0;
    }

    /// @notice Execute a mock xcall, default gas limit. Reverts if the call fails, or if the gas limit is too low
    function mockXCall(uint64 sourceChainId, address sender, address to, bytes calldata data) external {
        mockXCall(sourceChainId, sender, to, data, XMSG_DEFAULT_GAS_LIMIT);
    }

    /// @dev Execute a mock xcall, custom gas limit. Reverts if the call fails, or if the gas limit is too low
    function mockXCall(uint64 sourceChainId, address sender, address to, bytes calldata data, uint64 gasLimit) public {
        require(gasLimit <= XMSG_MAX_GAS_LIMIT, "MockPortal: gasLimit too high");
        require(gasLimit >= XMSG_MIN_GAS_LIMIT, "MockPortal: gasLimit too low");

        _currentXMsg = XTypes.Msg({ sourceChainId: sourceChainId, sender: sender });

        uint256 gasUsed = gasleft();
        (bool success, bytes memory returnData) = to.call{ gas: gasLimit }(data);
        gasUsed = gasUsed - gasleft();

        if (!success && gasUsed >= gasLimit) revert("MockPortal: out of gas");
        require(success, string(returnData));
    }
}
