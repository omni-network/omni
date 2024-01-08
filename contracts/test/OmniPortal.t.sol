// SPDX-License-Identifier: No License (None)
pragma solidity 0.8.23;

import {OmniPortal} from "src/OmniPortal.sol";
import {CommonTest} from "test/common/CommonTest.sol";

contract OmniPortal_Test is CommonTest {
    /// @dev Test that xcall emits an XMsg event with default gas limit, and increments outXStreamOffset
    function test_xcall_defaultGasLimit_succeeds() public {
        uint64 destChainId = 2;
        address to = makeAddr("foo-addr-on-dest");
        bytes memory data = abi.encodeWithSignature("foo()");
        uint64 gasLimit = portal.XMSG_DEFAULT_GAS_LIMIT();
        uint64 offset = portal.outXStreamOffset(destChainId);

        // check XMsg event is emitted
        vm.expectEmit();
        emit XMsg(destChainId, offset, xcaller, to, data, gasLimit);

        // make xcall
        vm.prank(xcaller);
        portal.xcall(destChainId, to, data);

        // check outXStreamOffset is incremented
        assertEq(portal.outXStreamOffset(destChainId), offset + 1);
    }
}
