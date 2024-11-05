// SPDX-License-Identifier: GPL-3.0-only
pragma solidity 0.8.24;

import { OmniBridgeNative } from "src/token/OmniBridgeNative.sol";
import { TransparentUpgradeableProxy } from "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import { MockPortal } from "test/utils/MockPortal.sol";
import { Omni } from "src/token/Omni.sol";
import { Test } from "forge-std/Test.sol";

/// @dev OmniBridgeNative test fixtures
contract OmniBridgeNativeFixtures is Test {
    MockPortal portal;
    OmniBridgeNativeHarness nativebridge;

    address owner;
    address l1bridge;
    uint64 l1ChainId;

    uint256 totalSupply = 100_000_000 * 10 ** 18;

    function setUp() public {
        portal = new MockPortal();
        l1ChainId = 1;
        l1bridge = makeAddr("l1bridge");
        owner = makeAddr("owner");

        address impl = address(new OmniBridgeNativeHarness());
        nativebridge = OmniBridgeNativeHarness(
            address(
                new TransparentUpgradeableProxy(
                    impl, owner, abi.encodeWithSelector(OmniBridgeNative.initialize.selector, (owner))
                )
            )
        );

        vm.prank(owner);
        nativebridge.setup(l1ChainId, address(portal), l1bridge, 0);
        vm.deal(address(nativebridge), totalSupply);
    }
}

/// @dev A wrapper around OmniBridgeNative, with public state setters.
contract OmniBridgeNativeHarness is OmniBridgeNative {
    function setL1Deposits(uint256 balance) public {
        l1Deposits = balance;
    }

    function setClaimable(address claimant, uint256 amount) public {
        claimable[claimant] = amount;
    }
}
