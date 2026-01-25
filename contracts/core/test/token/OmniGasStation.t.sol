// SPDX-License-Identifier: GPL-3.0-only
pragma solidity 0.8.24;

import { OwnableUpgradeable } from "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import { PausableUpgradeable } from "@openzeppelin/contracts-upgradeable/utils/PausableUpgradeable.sol";
import { TransparentUpgradeableProxy } from "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import { MockPortal } from "test/utils/MockPortal.sol";
import { OmniGasStation } from "src/token/OmniGasStation.sol";
import { OmniGasPump } from "src/token/OmniGasPump.sol";
import { Test } from "forge-std/Test.sol";

/**
 * @title OmniGasStation_Test
 * @notice Test suite for OmniGasStation
 */
contract OmniGasStation_Test is Test {
    OmniGasStation station;
    MockPortal portal;
    address owner;

    function setUp() public {
        address impl = address(new OmniGasStation());

        portal = new MockPortal();
        owner = makeAddr("owner");

        station = OmniGasStation(
            payable(address(
                    new TransparentUpgradeableProxy(
                        impl,
                        makeAddr("admin"),
                        abi.encodeCall(
                            OmniGasStation.initialize, (address(portal), owner, new OmniGasStation.GasPump[](0))
                        )
                    )
                ))
        );
    }

    function test_settleUp() public {
        address recipient = makeAddr("recipient");
        uint256 owed = 10 ether;
        uint64 chainId = 1;
        address pump = makeAddr("pump");
        uint64 gasLimit = 100_000; // TODO: test and trim if possible

        // only xcall
        vm.expectRevert("GasStation: unauthorized");
        station.settleUp(recipient, owed);

        // only pump (pump not set yet)
        vm.expectRevert("GasStation: unauthorized");
        portal.mockXCall({
            sourceChainId: chainId,
            sender: pump,
            data: abi.encodeCall(OmniGasStation.settleUp, (recipient, owed)),
            to: address(station),
            gasLimit: gasLimit
        });

        // add pump
        vm.prank(owner);
        station.setPump(chainId, pump);

        // handles out of funds
        portal.mockXCall({
            sourceChainId: chainId,
            sender: pump,
            data: abi.encodeCall(OmniGasStation.settleUp, (recipient, owed)),
            to: address(station),
            gasLimit: gasLimit
        });
        assertEq(station.fueled(recipient, chainId), 0);

        // fund station
        uint256 initialSupply = 1000 ether;
        (bool success,) = address(station).call{ value: initialSupply }("");
        assertTrue(success);

        // settles up
        portal.mockXCall({
            sourceChainId: chainId,
            sender: pump,
            data: abi.encodeCall(OmniGasStation.settleUp, (recipient, owed)),
            to: address(station),
            gasLimit: gasLimit
        });

        // check settled
        assertEq(station.fueled(recipient, chainId), owed);
        assertEq(address(station).balance, initialSupply - owed);
        assertEq(recipient.balance, owed);

        // doesn't transfer again
        vm.expectRevert("GasStation: already funded");
        portal.mockXCall({
            sourceChainId: chainId,
            sender: pump,
            data: abi.encodeCall(OmniGasStation.settleUp, (recipient, owed)),
            to: address(station),
            gasLimit: gasLimit
        });

        // transfers more when owed
        uint256 more = 5 ether;
        owed += more;
        portal.mockXCall({
            sourceChainId: chainId,
            sender: pump,
            data: abi.encodeCall(OmniGasStation.settleUp, (recipient, owed)),
            to: address(station),
            gasLimit: gasLimit
        });

        // check settled
        assertEq(station.fueled(recipient, chainId), owed);
        assertEq(address(station).balance, initialSupply - owed);
        assertEq(recipient.balance, owed);
    }

    function test_setPump() public {
        uint64 chainId = 1;
        address pump = makeAddr("pump");

        // only owner
        address notOwner = address(0x456);
        vm.prank(notOwner);
        vm.expectRevert(abi.encodeWithSelector(OwnableUpgradeable.OwnableUnauthorizedAccount.selector, notOwner));
        station.setPump(chainId, pump);

        // no zero addr
        vm.prank(owner);
        vm.expectRevert("GasStation: zero addr");
        station.setPump(chainId, address(0));

        // no zero chainID
        vm.prank(owner);
        vm.expectRevert("GasStation: zero chainId");
        station.setPump(0, pump);

        // owner can set
        vm.prank(owner);
        station.setPump(chainId, pump);

        assertEq(station.pumps(chainId), pump);
        assertEq(station.isPump(chainId, pump), true);

        // isPump false for other chainID
        assertEq(station.isPump(chainId + 1, pump), false);

        // isPump false for other pump
        assertEq(station.isPump(chainId, makeAddr("other")), false);
    }

    function test_pause() public {
        assertFalse(station.paused());

        // only owner can pause
        address notOwner = address(0x456);
        vm.prank(notOwner);
        vm.expectRevert(abi.encodeWithSelector(OwnableUpgradeable.OwnableUnauthorizedAccount.selector, notOwner));
        station.pause();

        // owner can pause
        vm.prank(owner);
        station.pause();

        assertTrue(station.paused());

        // settledUp is paused
        address recipient = makeAddr("recipient");
        uint256 owed = 1e18;
        vm.expectRevert(abi.encodeWithSelector(PausableUpgradeable.EnforcedPause.selector));
        station.settleUp(recipient, owed);

        // only owner can unpause
        vm.prank(notOwner);
        vm.expectRevert(abi.encodeWithSelector(OwnableUpgradeable.OwnableUnauthorizedAccount.selector, notOwner));
        station.unpause();

        // owner can unpause
        vm.prank(owner);
        station.unpause();

        assertFalse(station.paused());

        // settledUp is unpaused
        vm.prank(owner);
        vm.expectRevert("GasStation: unauthorized"); // reverts,  but not because its paused
        station.settleUp(recipient, owed);
    }
}
