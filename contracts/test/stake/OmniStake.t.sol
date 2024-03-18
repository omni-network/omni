// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.24;

import { OmniStake } from "src/protocol/OmniStake.sol";
import { Test, Vm } from "forge-std/Test.sol";

/**
 * @title OmniStake_Test
 * @notice Test suite for the OmniStake contract
 */
contract OmniStake_Test is Test {
    /// @notice Emitted when a user deposits into OmniStake
    event Deposit(bytes pubkey, uint256 amount);

    OmniStake stake;

    function setUp() public {
        stake = new OmniStake();
    }

    /// @dev Test that a valid deposit succeeds
    function test_deposit_succeeds() public {
        Vm.Wallet memory wallet = vm.createWallet("test val");
        bytes memory pubkey = _pubkey(wallet);

        uint256 amt = 16 ether;
        vm.deal(wallet.addr, amt);

        vm.expectEmit();
        emit Deposit(pubkey, amt);

        vm.prank(wallet.addr);
        stake.deposit{ value: amt }(pubkey);
    }

    /// @dev Test that the pubkey of a deposit must match the sender
    function test_deposit_wrongPubkey_reverts() public {
        Vm.Wallet memory wallet = vm.createWallet("test val");

        Vm.Wallet memory wallet2 = vm.createWallet("test val2");
        bytes memory pubkey2 = _pubkey(wallet2);

        uint256 amt = 16 ether;
        vm.deal(wallet.addr, amt);

        vm.expectRevert("OmniStake: pubkey not sender");
        vm.prank(wallet.addr);
        stake.deposit{ value: amt }(pubkey2);
    }

    /// @dev Test that a deposit with a pubkey of incorrect length reverts
    function test_deposit_invalidLength_reverts() public {
        Vm.Wallet memory wallet = vm.createWallet("test val");
        bytes memory pubkey = bytes.concat(hex"04", _pubkey(wallet));

        uint256 amt = 16 ether;
        vm.deal(wallet.addr, amt);

        vm.expectRevert("Secp256k1: invalid pubkey length");
        vm.prank(wallet.addr);
        stake.deposit{ value: amt }(pubkey);
    }

    /// @dev Test that a deposit below 1 ether reverts
    function test_deposit_below1Ether_reverts() public {
        Vm.Wallet memory wallet = vm.createWallet("test val");
        bytes memory pubkey = _pubkey(wallet);

        uint256 amt = 1 ether - 1;
        vm.deal(wallet.addr, amt);

        vm.expectRevert("OmniStake: deposit amt too low");
        vm.prank(wallet.addr);
        stake.deposit{ value: amt }(pubkey);
    }

    /// @dev Test that a deposit above max uint64 reverts
    function test_deposit_aboveMaxUint64_reverts() public {
        Vm.Wallet memory wallet = vm.createWallet("test val");
        bytes memory pubkey = _pubkey(wallet);

        uint256 amt = uint256(type(uint64).max) + 1;
        vm.deal(wallet.addr, amt);

        vm.expectRevert("OmniStake: deposit amt too high");
        vm.prank(wallet.addr);
        stake.deposit{ value: amt }(pubkey);
    }

    /// @dev Helper function to convert a wallet to a secp256k1 64 uncompressed public key (without 0x04 prefix)
    function _pubkey(Vm.Wallet memory wallet) internal pure returns (bytes memory) {
        return abi.encode(wallet.publicKeyX, wallet.publicKeyY);
    }
}
