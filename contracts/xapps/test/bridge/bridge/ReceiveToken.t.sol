// SPDX-License-Identifier: GPL-3.0-only
pragma solidity 0.8.26;

import "../TestBase.sol";

contract ReceiveTokenTest is TestBase {
    function test_receiveToken_reverts() public {
        bytes memory data = abi.encodeCall(Bridge.receiveToken, (user, 1));
        uint64 unknownChainId = DEST_CHAIN_ID + 1;
        address unknownSender = makeAddr("unknownSender");
        uint64 lockboxGasLimit = _getGasLimit(Bridge(bridgeWithLockbox));

        // Unknown source chain ID
        vm.expectRevert(abi.encodeWithSelector(IBridge.Unauthorized.selector, unknownChainId, address(bridgeNoLockbox)));
        omni.mockXCall({
            sourceChainId: unknownChainId,
            sender: address(bridgeNoLockbox),
            to: address(bridgeWithLockbox),
            data: data,
            gasLimit: lockboxGasLimit
        });

        // Unknown source chain sender
        vm.expectRevert(abi.encodeWithSelector(IBridge.Unauthorized.selector, DEST_CHAIN_ID, unknownSender));
        omni.mockXCall({
            sourceChainId: DEST_CHAIN_ID,
            sender: unknownSender,
            to: address(bridgeWithLockbox),
            data: data,
            gasLimit: lockboxGasLimit
        });

        // Unauthorized direct call
        vm.expectRevert(abi.encodeWithSelector(IBridge.Unauthorized.selector, SRC_CHAIN_ID, user));
        vm.prank(user);
        bridgeWithLockbox.receiveToken(user, 1);
    }

    function test_receiveToken_succeeds_solvent_lockbox() public {
        bytes memory data = abi.encodeCall(Bridge.receiveToken, (user, INITIAL_USER_BALANCE));

        vm.prank(user);
        lockbox.deposit(INITIAL_USER_BALANCE);

        omni.mockXCall({
            sourceChainId: DEST_CHAIN_ID,
            sender: address(bridgeNoLockbox),
            to: address(bridgeWithLockbox),
            data: data,
            gasLimit: _getGasLimit(Bridge(bridgeWithLockbox))
        });

        _assertBalances({
            addr: user, tokenUserBal: INITIAL_USER_BALANCE, tokenLockboxBal: 0, wrapperUserBal: INITIAL_USER_BALANCE
        });
    }

    function test_receiveToken_succeeds_insolvent_lockbox() public {
        bytes memory data = abi.encodeCall(Bridge.receiveToken, (user, INITIAL_USER_BALANCE));

        omni.mockXCall({
            sourceChainId: DEST_CHAIN_ID,
            sender: address(bridgeNoLockbox),
            to: address(bridgeWithLockbox),
            data: data,
            gasLimit: _getGasLimit(Bridge(bridgeWithLockbox))
        });

        _assertBalances({
            addr: user, tokenUserBal: INITIAL_USER_BALANCE, tokenLockboxBal: 0, wrapperUserBal: INITIAL_USER_BALANCE
        });
    }

    function test_receiveToken_succeeds_no_lockbox() public {
        bytes memory data = abi.encodeCall(Bridge.receiveToken, (user, INITIAL_USER_BALANCE));

        omni.mockXCall({
            sourceChainId: SRC_CHAIN_ID,
            sender: address(bridgeWithLockbox),
            to: address(bridgeNoLockbox),
            data: data,
            gasLimit: _getGasLimit(Bridge(bridgeNoLockbox))
        });

        _assertBalances({
            addr: user, tokenUserBal: INITIAL_USER_BALANCE, tokenLockboxBal: 0, wrapperUserBal: INITIAL_USER_BALANCE
        });
    }

    function test_receiveToken_succeeds_paused() public prank(pauser) {
        bytes memory data = abi.encodeCall(Bridge.receiveToken, (user, INITIAL_USER_BALANCE));

        bridgeWithLockbox.pause();
        bridgeNoLockbox.pause();

        omni.mockXCall({
            sourceChainId: DEST_CHAIN_ID,
            sender: address(bridgeNoLockbox),
            to: address(bridgeWithLockbox),
            data: data,
            gasLimit: _getGasLimit(Bridge(bridgeWithLockbox))
        });
        omni.mockXCall({
            sourceChainId: SRC_CHAIN_ID,
            sender: address(bridgeWithLockbox),
            to: address(bridgeNoLockbox),
            data: data,
            gasLimit: _getGasLimit(Bridge(bridgeNoLockbox))
        });

        _assertBalances({
            addr: user, tokenUserBal: INITIAL_USER_BALANCE, tokenLockboxBal: 0, wrapperUserBal: INITIAL_USER_BALANCE * 2
        });
    }
}
