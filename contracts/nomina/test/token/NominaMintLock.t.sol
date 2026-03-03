// SPDX-License-Identifier: GPL-3.0-only
pragma solidity 0.8.24;

import { Test } from "forge-std/Test.sol";
import { Nomina } from "src/token/Nomina.sol";
import { NominaMintLock } from "src/token/NominaMintLock.sol";
import { MockOmni } from "test/utils/MockOmni.sol";

contract NominaMintLock_Test is Test {
    Nomina public nomina;
    NominaMintLock public lock;
    MockOmni public omni;

    address public mintAuthority = makeAddr("mintAuthority");
    address public minter = makeAddr("minter");
    address public user = makeAddr("user");

    function setUp() public {
        omni = new MockOmni(1_000_000 ether, user);
        nomina = new Nomina(address(omni), mintAuthority);
        lock = new NominaMintLock(nomina);

        vm.prank(mintAuthority);
        nomina.setMinter(minter);
    }

    function test_acceptMintAuthority() public {
        // Queue the lock contract as pending mint authority
        vm.prank(mintAuthority);
        nomina.setMintAuthority(address(lock));

        // Accept mint authority via the lock contract
        lock.acceptMintAuthority();

        // Verify the lock contract is now the mint authority
        assertEq(nomina.mintAuthority(), address(lock), "mint authority mismatch");
        assertEq(nomina.pendingMintAuthority(), address(0), "pending mint authority not cleared");
    }

    function test_acceptMintAuthority_reverts_notPending() public {
        // Revert when lock is not the pending mint authority
        vm.expectRevert(Nomina.Unauthorized.selector);
        lock.acceptMintAuthority();
    }

    function test_mintLocked() public {
        // Transfer mint authority to the lock contract
        vm.prank(mintAuthority);
        nomina.setMintAuthority(address(lock));
        lock.acceptMintAuthority();

        // The lock contract cannot set a minter, so minting is permanently locked.
        // Existing minter still works until mint authority sets a new one,
        // but the lock contract has no way to call setMinter or setMintAuthority.

        // Verify the lock contract has no setMinter function by checking
        // that the mint authority (lock) cannot set a new minter.
        // Since NominaMintLock has no setMinter or setMintAuthority functions,
        // the mint authority is permanently locked.
        (bool success,) = address(lock).call(abi.encodeWithSignature("setMinter(address)", user));
        assertFalse(success, "lock should not have setMinter");

        (success,) = address(lock).call(abi.encodeWithSignature("setMintAuthority(address)", user));
        assertFalse(success, "lock should not have setMintAuthority");
    }
}
