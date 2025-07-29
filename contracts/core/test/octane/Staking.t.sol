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

    /// @dev Matches Staking.Undelegate event
    event Undelegate(address indexed delegator, address indexed validator, uint256 amount);

    address owner;
    address validator;
    address[] validators;
    StakingHarness staking;

    function setUp() public {
        owner = makeAddr("owner");
        validator = makeAddr("validator");
        validators.push(validator);
        staking = new StakingHarness(owner);
    }

    function test_createValidator_withoutSignature() public {
        uint256 deposit = staking.MinDeposit();
        bytes memory pubkey = abi.encodePacked(hex"03440d290e4394cd9832cc7025769be18ab7975e34e4514c31c07da3d370fe0b05");
        vm.deal(validator, staking.MinDeposit());

        // allowlist is disabled
        assertFalse(staking.isAllowlistEnabled());

        // enable allowlist
        vm.prank(owner);
        staking.enableAllowlist();
        assertTrue(staking.isAllowlistEnabled());

        // must be in allowlist
        vm.expectRevert("Staking: not allowed");
        staking.createValidator(pubkey);

        // add to allowlist
        vm.prank(owner);
        staking.allowValidators(validators);
        assertTrue(staking.isAllowedValidator(validator));

        // requires minimum deposit
        uint256 insufficientDeposit = staking.MinDeposit() - 1;

        vm.expectRevert("Staking: insufficient deposit");
        vm.prank(validator);
        staking.createValidator{ value: insufficientDeposit }(pubkey);

        // requires 33 byte pubkey
        bytes memory pubkey32 = abi.encodePacked(keccak256("pubkey"));
        vm.expectRevert("Secp256k1: pubkey not 33 bytes");
        vm.prank(validator);
        staking.createValidator{ value: deposit }(pubkey32);

        // requires valid pubkey prefix
        bytes memory badPrefix = abi.encodePacked(hex"01", pubkey32);
        vm.expectRevert("Secp256k1: invalid pubkey prefix");
        vm.prank(validator);
        staking.createValidator{ value: deposit }(badPrefix);

        // requires a valid pubkey on the secp256k1 curve
        bytes memory notOnCurve = abi.encodePacked(hex"02", pubkey32);
        vm.expectRevert("Secp256k1: pubkey not on curve");
        vm.prank(validator);
        staking.createValidator{ value: deposit }(notOnCurve);

        // succeeds with valid deposit and pubkey
        vm.expectEmit();
        emit CreateValidator(validator, pubkey, deposit);

        vm.prank(validator);
        staking.createValidator{ value: deposit }(pubkey);

        // remove from allowlist
        vm.prank(owner);
        staking.disallowValidators(validators);
        assertFalse(staking.isAllowedValidator(validator));

        // must be in allowlist
        vm.expectRevert("Staking: not allowed");
        vm.deal(validator, deposit);
        vm.prank(validator);
        staking.createValidator{ value: deposit }(pubkey);

        // disable allowlist
        vm.prank(owner);
        staking.disableAllowlist();
        assertFalse(staking.isAllowlistEnabled());

        // can create validator with allowlist disabled
        vm.expectEmit();
        emit CreateValidator(validator, pubkey, deposit);

        vm.prank(validator);
        staking.createValidator{ value: deposit }(pubkey);
    }

    function test_createValidator_withSignature() public {
        uint256 privkey = 0x5aae8cd28d4456aba1d24542558bc2fac787e2fdc2210c20f2f3375e82174205;
        uint256 x = 0x534d719d4f56544f42e22cab20886dd64fb713a5c72b31f929d856654a11dc0c;
        uint256 y = 0x5609e3c7f55c46a197ead4a96caa63eeade00b4a775e7709f6e673157a724d6c;
        bytes memory pubkey = Secp256k1.compress(x, y);
        bytes memory compressed = Secp256k1.compress(x, y);
        bytes32 digest = staking.getConsPubkeyDigest(validator);
        bytes memory signature = _sign(digest, privkey);

        vm.deal(validator, staking.MinDeposit());

        // allowlist is disabled
        assertFalse(staking.isAllowlistEnabled());

        // enable allowlist
        vm.prank(owner);
        staking.enableAllowlist();
        assertTrue(staking.isAllowlistEnabled());

        // must be in allowlist
        vm.expectRevert("Staking: not allowed");
        staking.createValidator(pubkey, signature);

        // add to allowlist
        vm.prank(owner);
        staking.allowValidators(validators);
        assertTrue(staking.isAllowedValidator(validator));

        // requires minimum deposit
        uint256 insufficientDeposit = staking.MinDeposit() - 1;

        vm.expectRevert("Staking: insufficient deposit");
        vm.prank(validator);
        staking.createValidator{ value: insufficientDeposit }(pubkey, signature);

        uint256 deposit = staking.MinDeposit();

        // requires 33 byte pubkey
        bytes memory pubkey32 = abi.encodePacked(keccak256("pubkey"));
        vm.expectRevert("Secp256k1: pubkey not 33 bytes");
        vm.prank(validator);
        staking.createValidator{ value: deposit }(pubkey32, signature);

        // requires valid pubkey prefix
        bytes memory badPrefix = abi.encodePacked(hex"01", pubkey32);
        vm.expectRevert("Secp256k1: invalid pubkey prefix");
        vm.prank(validator);
        staking.createValidator{ value: deposit }(badPrefix, signature);

        // requires a valid pubkey on the secp256k1 curve
        bytes memory notOnCurve = abi.encodePacked(hex"02", pubkey32);
        vm.expectRevert("Secp256k1: pubkey not on curve");
        vm.prank(validator);
        staking.createValidator{ value: deposit }(notOnCurve, signature);

        // requires a valid signature
        bytes memory badSignature = abi.encodePacked(signature);
        badSignature[0] = bytes1(uint8(badSignature[0]) + 1);
        vm.expectRevert("Staking: invalid signature");
        vm.prank(validator);
        staking.createValidator{ value: deposit }(pubkey, badSignature);

        // requires signed by msg.sender
        vm.prank(owner);
        staking.disableAllowlist();
        address attacker = makeAddr("attacker");
        vm.deal(attacker, deposit);
        vm.expectRevert("Staking: invalid signature");
        vm.prank(attacker);
        staking.createValidator{ value: deposit }(pubkey, signature);
        vm.prank(owner);
        staking.enableAllowlist();

        // succeeds with valid deposit and pubkey
        vm.expectEmit();
        emit CreateValidator(validator, compressed, deposit);

        vm.prank(validator);
        staking.createValidator{ value: deposit }(pubkey, signature);

        // remove from allowlist
        vm.prank(owner);
        staking.disallowValidators(validators);
        assertFalse(staking.isAllowedValidator(validator));

        // must be in allowlist
        vm.expectRevert("Staking: not allowed");
        vm.deal(validator, deposit);
        vm.prank(validator);
        staking.createValidator{ value: deposit }(pubkey, signature);

        // disable allowlist
        vm.prank(owner);
        staking.disableAllowlist();
        assertFalse(staking.isAllowlistEnabled());

        // can create validator with allowlist disabled
        vm.expectEmit();
        emit CreateValidator(validator, compressed, deposit);

        vm.prank(validator);
        staking.createValidator{ value: deposit }(pubkey, signature);
    }

    /*function test_delegate() public {
        // requires min delegation
        uint256 minDelegation = staking.MinDelegation();

        vm.deal(validator, minDelegation);

        vm.expectRevert("Staking: insufficient deposit");
        staking.delegate{ value: minDelegation - 1 }(validator);

        // if allowlist enabled, must be in allowlist
        vm.prank(owner);
        staking.enableAllowlist();

        vm.expectRevert("Staking: not allowed val");
        vm.prank(validator);
        staking.delegate{ value: minDelegation }(validator);

        // succeeds
        vm.prank(owner);
        staking.allowValidators(validators);

        vm.expectEmit();
        emit Delegate(validator, validator, minDelegation);

        vm.prank(validator);
        staking.delegate{ value: minDelegation }(validator);
    }

    function test_undelegate() public {
        uint256 fee = 0.1 ether;
        uint256 amount = 1 ether;

        vm.expectRevert("Staking: insufficient fee");
        staking.undelegate{ value: 0 }(owner, amount);

        vm.deal(owner, amount + fee);

        // if allowlist enabled, must be in allowlist
        vm.prank(owner);
        staking.enableAllowlist();

        vm.expectRevert("Staking: not allowed val");
        vm.prank(owner);
        staking.undelegate{ value: fee }(validator, amount);

        // succeeds
        vm.prank(owner);
        staking.allowValidators(validators);

        vm.expectEmit();
        emit Undelegate(owner, validator, amount);

        vm.prank(owner);
        staking.undelegate{ value: fee }(validator, amount);
        emit Undelegate(owner, validator, amount);
    }*/

    function test_temporarilyDisabled() public {
        vm.expectRevert(abi.encodeWithSelector(Staking.TemporarilyDisabled.selector));
        staking.delegate(validator);

        vm.expectRevert(abi.encodeWithSelector(Staking.TemporarilyDisabled.selector));
        staking.delegateFor(owner, validator);

        vm.expectRevert(abi.encodeWithSelector(Staking.TemporarilyDisabled.selector));
        staking.undelegate(validator, 1 ether);
    }

    function _sign(bytes32 digest, uint256 privkey) private pure returns (bytes memory) {
        (uint8 v, bytes32 r, bytes32 s) = vm.sign(privkey, digest);
        return abi.encodePacked(r, s, v);
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
