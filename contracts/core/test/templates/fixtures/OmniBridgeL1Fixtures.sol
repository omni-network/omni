// SPDX-License-Identifier: GPL-3.0-only
pragma solidity 0.8.24;

import { TransparentUpgradeableProxy } from "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import { MockPortal } from "test/utils/MockPortal.sol";
import { Omni } from "src/token/Omni.sol";
import { OmniBridgeL1 } from "src/token/OmniBridgeL1.sol";
import { Test } from "forge-std/Test.sol";

/// @dev OmniBridgeL1 test fixtures
contract OmniBridgeL1Fixtures is Test {
    OmniBridgeL1 l1bridge;
    Omni token;
    MockPortal portal;

    address owner;
    address proxyAdmin;
    address bank;
    uint256 totalSupply = 100_000_000 * 10 ** 18;

    function setUp() public {
        bank = makeAddr("bank");
        owner = makeAddr("owner");
        proxyAdmin = makeAddr("proxyadmin");

        portal = new MockPortal();
        token = new Omni(totalSupply, bank);

        address impl = address(new OmniBridgeL1(address(token)));
        l1bridge = OmniBridgeL1(
            address(
                new TransparentUpgradeableProxy(
                    impl, proxyAdmin, abi.encodeCall(OmniBridgeL1.initialize, (owner, address(portal)))
                )
            )
        );
    }
}
