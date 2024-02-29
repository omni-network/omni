// SPDX-License-Identifier: GPL-3.0-only
pragma solidity ^0.8.12;

import { XTypes } from "../../src/libraries/XTypes.sol";

/**
 * @title MockPortal
 * @dev A mock OmniPortal contract, for testing purposes.
 */
contract MockPortal {
    /// @dev Matches OmniPortal.XMSG_DEFAULT_GAS_LIMIT
    uint64 public constant XMSG_DEFAULT_GAS_LIMIT = 200_000;

    XTypes.Msg private _currentXMsg;

    /// @dev Returns current xmsg, same as OmniPortal.xmsg()
    function xmsg() external view returns (XTypes.Msg memory) {
        return _currentXMsg;
    }

    /// @dev Stubbed feeFor, default gas limit
    function feeFor(uint64, /* destChainId */ bytes calldata /* data */ ) external pure returns (uint256) {
        return 1 gwei;
    }

    /// @dev Stubbed feeFor, custom gas limit
    function feeFor(uint64, /* destChainId */ bytes calldata, /* data */ uint64 /* gasLimit */ )
        external
        pure
        returns (uint256)
    {
        return 1 gwei;
    }

    /// @dev Stubbed xcall, default gas limit
    function xcall(uint64, /* destChainId */ address, /* to */ bytes calldata /* data */ ) external payable { }

    /// @dev Stubbed xcall, custom gas limit
    function xcall(uint64, /* destChainId */ address, /* to */ bytes calldata, /* data */ uint64 /* gasLimit */ )
        external
        payable
    { }

    /// @dev Returns true if the current call is an xcall, same as OmniPortal.isXCall()
    function isXCall() external view returns (bool) {
        return _currentXMsg.sourceChainId != 0;
    }

    /// @dev Execute a mock xcall, default gas limit
    ///      Reverts if the call fails
    function mockXCall(uint64 sourceChainId, address sender, address to, bytes calldata data) external {
        mockXCall(sourceChainId, sender, to, data, XMSG_DEFAULT_GAS_LIMIT);
    }

    /// @dev Execute a mock xcall, custom gas limit
    ///     Reverts if the call fails
    function mockXCall(uint64 sourceChainId, address sender, address to, bytes calldata data, uint64 gasLimit) public {
        _currentXMsg = XTypes.Msg({
            sourceChainId: sourceChainId,
            destChainId: uint64(block.chainid),
            streamOffset: 0, // doesn't matter
            sender: sender,
            to: to,
            data: data,
            gasLimit: gasLimit
        });

        (bool success, bytes memory returnData) = to.call(data);
        require(success, string(returnData));
    }
}
