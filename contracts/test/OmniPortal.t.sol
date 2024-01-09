// SPDX-License-Identifier: GPL-3.0-only
pragma solidity 0.8.23;

import {OmniPortal} from "src/OmniPortal.sol";
import {CommonTest} from "test/common/CommonTest.sol";

contract OmniPortal_Test is CommonTest {
    /// @dev Test that xcall with default gas limt emits XMsg event and increments outXStreamOffset
    function test_xcall_defaultGasLimit_succeeds() public {
        (uint64 destChainId, uint64 offset, address to, bytes memory data) = _xfoo();
        uint64 gasLimit = portal.XMSG_DEFAULT_GAS_LIMIT();

        // check XMsg event is emitted
        vm.expectEmit();
        emit XMsg(destChainId, offset, xcaller, to, data, gasLimit);

        // make xcall
        vm.prank(xcaller);
        portal.xcall(destChainId, to, data);

        // check outXStreamOffset is incremented
        assertEq(portal.outXStreamOffset(destChainId), offset + 1);
    }

    /// @dev Test that xcall with explicit gas limt emits XMsg event and increments outXStreamOffset
    function test_xcall_explicitGasLimit_succeeds() public {
        (uint64 destChainId, uint64 offset, address to, bytes memory data) = _xfoo();
        uint64 gasLimit = portal.XMSG_DEFAULT_GAS_LIMIT();

        // check XMsg event is emitted
        vm.expectEmit();
        emit XMsg(destChainId, offset, xcaller, to, data, gasLimit);

        // make xcall
        vm.prank(xcaller);
        portal.xcall(destChainId, to, data);

        // check outXStreamOffset is incremented
        assertEq(portal.outXStreamOffset(destChainId), offset + 1);
    }

    /// @dev Get test foo() xcall params
    function _xfoo() private returns (uint64 destChainId, uint64 offset, address to, bytes memory data) {
        return (2, portal.outXStreamOffset(2), makeAddr("foo-addr-on-dest"), abi.encodeWithSignature("foo()"));
    }
}
