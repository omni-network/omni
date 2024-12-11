// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.24;

import { Staking } from "src/octane/Staking.sol";
import { Test, Vm } from "forge-std/Test.sol";
import { Secp256k1 } from "src/libraries/Secp256k1.sol";

/**
 * @title Staking_Test
 * @notice Test suite for Staking.sol
 */
contract Staking_Test is Test {
    /// @dev Matches Staking.CreateValidator event
    event CreateValidator(address indexed validator, bytes pubkey, uint256 deposit);

    /// @dev Matches Staking.Delegate event
    event Delegate(address indexed delegator, address indexed validator, uint256 amount);

    address owner;
    address validator;
    StakingHarness staking;

    function setUp() public {
        owner = makeAddr("owner");
        validator = makeAddr("validator");
        staking = new StakingHarness(owner);
    }

    function test_createValidator() public {
        address[] memory validators = new address[](1);
        validators[0] = validator;
        bytes32 privkey = 0x5aae8cd28d4456aba1d24542558bc2fac787e2fdc2210c20f2f3375e82174205;
        bytes32 x = 0x534d719d4f56544f42e22cab20886dd64fb713a5c72b31f929d856654a11dc0c;
        bytes32 y = 0x5609e3c7f55c46a197ead4a96caa63eeade00b4a775e7709f6e673157a724d6c;
        bytes memory pubkey = Secp256k1.compressPublicKey(x, y);
        bytes32 validatorPubkeyDigest = staking.getValidatorPubkeyDigest(x, y);
        (uint8 v, bytes32 r, bytes32 s) = vm.sign(uint256(privkey), validatorPubkeyDigest);
        bytes memory signature = abi.encodePacked(r, s, v);
        vm.deal(validator, staking.MinDeposit());

        // allowlist is disabled
        assertFalse(staking.isAllowlistEnabled());

        // enable allowlist
        vm.prank(owner);
        staking.enableAllowlist();
        assertTrue(staking.isAllowlistEnabled());

        // must be in allowlist
        vm.expectRevert("Staking: not allowed");
        staking.createValidator(x, y, signature);

        // add to allowlist
        vm.prank(owner);
        staking.allowValidators(validators);
        assertTrue(staking.isAllowedValidator(validator));

        // requires minimum deposit
        uint256 insufficientDeposit = staking.MinDeposit() - 1;

        vm.expectRevert("Staking: insufficient deposit");
        vm.prank(validator);
        staking.createValidator{ value: insufficientDeposit }(x, y, signature);

        uint256 deposit = staking.MinDeposit();

        // requires a valid pubkey on the secp256k1 curve
        bytes32 badY = bytes32(uint256(y) + 1);
        vm.expectRevert("Staking: invalid pubkey");
        vm.prank(validator);
        staking.createValidator{ value: deposit }(x, badY, signature);

        // requires a valid signature
        bytes memory badSignature = abi.encodePacked(signature);
        badSignature[0] = bytes1(uint8(badSignature[0]) + 1);
        vm.expectRevert("Staking: invalid signature");
        vm.prank(validator);
        staking.createValidator{ value: deposit }(x, y, badSignature);

        // succeeds with valid deposit and pubkey
        vm.expectEmit();
        emit CreateValidator(validator, pubkey, deposit);

        vm.prank(validator);
        staking.createValidator{ value: deposit }(x, y, signature);

        // remove from allowlist
        vm.prank(owner);
        staking.disallowValidators(validators);
        assertFalse(staking.isAllowedValidator(validator));

        // must be in allowlist
        vm.expectRevert("Staking: not allowed");
        vm.deal(validator, deposit);
        vm.prank(validator);
        staking.createValidator{ value: deposit }(x, y, signature);

        // disable allowlist
        vm.prank(owner);
        staking.disableAllowlist();
        assertFalse(staking.isAllowlistEnabled());

        // can create validator with allowlist disabled
        vm.expectEmit();
        emit CreateValidator(validator, pubkey, deposit);

        vm.prank(validator);
        staking.createValidator{ value: deposit }(x, y, signature);
    }

    function test_delegate() public {
        // requires min delegation
        uint256 minDelegation = staking.MinDelegation();

        vm.deal(validator, minDelegation);

        vm.expectRevert("Staking: insufficient deposit");
        staking.delegate{ value: minDelegation - 1 }(validator, validator);

        // if allowlist enabled, must be in allowlist
        vm.prank(owner);
        staking.enableAllowlist();

        vm.expectRevert("Staking: not allowed val");
        vm.prank(validator);
        staking.delegate{ value: minDelegation }(validator, validator);

        // succeeds
        address[] memory validators = new address[](1);
        validators[0] = validator;
        vm.prank(owner);
        staking.allowValidators(validators);

        vm.expectEmit();
        emit Delegate(validator, validator, minDelegation);

        vm.prank(validator);
        staking.delegate{ value: minDelegation }(validator, validator);
    }
}

/**
 * @title StakingHarness
 * @notice Wrapper around Staking.sol that allows setting owner and EIP-712 in constructor
 */
contract StakingHarness is Staking {
    bytes32 private constant EIP712StorageLocation = 0xa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d100;

    function getEIP712Storage() private pure returns (EIP712Storage storage $) {
        assembly {
            $.slot := EIP712StorageLocation
        }
    }

    constructor(address _owner) {
        _transferOwnership(_owner);

        EIP712Storage storage $ = getEIP712Storage();
        $._name = "Staking";
        $._version = "1";
    }
}
