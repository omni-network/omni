// SPDX-License-Identifier: GPL-3.0-only
pragma solidity 0.8.24;

import { Predeploys } from "src/libraries/Predeploys.sol";
import { Staking } from "src/octane/Staking.sol";
import { Secp256k1 } from "src/libraries/Secp256k1.sol";
import { Test } from "forge-std/Test.sol";
import { VmSafe } from "forge-std/Vm.sol";

contract StakingPostUpgradeTest is Test {
    Staking staking;
    address owner;
    address validator;

    function run() public {
        (VmSafe.CallerMode mode,,) = vm.readCallers();
        require(mode == VmSafe.CallerMode.None, "no broadcast");

        _setup();
        _testEip712();
        _testAllowlist();
        _testCreateValidator();
        _testDelegate();
    }

    function _setup() internal {
        staking = Staking(Predeploys.Staking);
        owner = staking.owner();
        validator = makeAddr("validator");
    }

    function _testEip712() internal view {
        (
            bytes1 fields,
            string memory name,
            string memory version,
            uint256 chainId,
            address verifyingContract,
            bytes32 salt,
            uint256[] memory extensions
        ) = staking.eip712Domain();

        assertEq(fields, hex"0f", "EIP-712 fields");
        assertEq(name, "Staking", "EIP-712 name");
        assertEq(version, "1", "EIP-712 version");
        assertEq(chainId, block.chainid, "EIP-712 chainId");
        assertEq(verifyingContract, address(staking), "EIP-712 verifyingContract");
        assertEq(salt, bytes32(0), "EIP-712 salt");
        assertEq(extensions.length, 0, "EIP-712 extensions");
    }

    function _testAllowlist() internal {
        vm.startPrank(owner);
        staking.enableAllowlist();
        assertTrue(staking.isAllowlistEnabled(), "allowlist disabled");

        address[] memory validators = new address[](1);
        validators[0] = validator;
        staking.allowValidators(validators);
        assertTrue(staking.isAllowedValidator(validator), "validator not in allowlist");

        staking.disallowValidators(validators);
        assertFalse(staking.isAllowedValidator(validator), "validator in allowlist");

        staking.disableAllowlist();
        assertFalse(staking.isAllowlistEnabled(), "allowlist enabled");
        vm.stopPrank();
    }

    function _testCreateValidator() internal {
        bytes32 privkey = 0xf0e1605dd50ce33553290b778b0f53b2cde5e47a8794c0e7d2815e456e6da3b9;
        bytes32 x = 0x3b12d750493ed6b12b390447f6dd38f587af12ed04ab8d6858e818cf0c63607c;
        bytes32 y = 0x044e0321a3e57de51e95f2b230b9e4ffed2318578baab1a80652234fe0115d13;
        bytes memory pubkey = Secp256k1.compressPublicKey(x, y);
        bytes32 digest = staking.getValidatorPubkeyDigest(x, y);
        (uint8 v, bytes32 r, bytes32 s) = vm.sign(uint256(privkey), digest);
        bytes memory signature = abi.encodePacked(r, s, v);
        uint256 deposit = staking.MinDeposit();
        vm.deal(validator, deposit);

        vm.expectEmit();
        emit Staking.CreateValidator(validator, pubkey, deposit);
        vm.prank(validator);
        staking.createValidator{ value: deposit }(x, y, signature);
    }

    function _testDelegate() internal {
        uint256 deposit = staking.MinDelegation();
        vm.deal(validator, deposit);

        vm.expectEmit();
        emit Staking.Delegate(validator, validator, deposit);
        vm.prank(validator);
        staking.delegate{ value: deposit }(validator, validator);
    }
}
