// SPDX-License-Identifier: GPL-3.0-only
pragma solidity 0.8.24;

import { OwnableUpgradeable } from "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import { PausableUpgradeable } from "@openzeppelin/contracts-upgradeable/utils/PausableUpgradeable.sol";
import { TransparentUpgradeableProxy } from "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import { FeeOracleV1, IFeeOracleV1 } from "src/xchain/FeeOracleV1.sol";
import { MockPortal, IOmniPortal } from "test/utils/MockPortal.sol";
import { ConfLevel } from "src/libraries/ConfLevel.sol";
import { OmniGasStation } from "src/token/OmniGasStation.sol";
import { OmniGasPump } from "src/token/OmniGasPump.sol";
import { Test } from "forge-std/Test.sol";

/**
 * @title OmniGasPump_Test
 * @notice Test suite for OmniGasPump
 */
contract OmniGasPump_Test is Test {
    OmniGasPump pump;
    MockPortal portal;
    FeeOracleV1 feeOracle;

    address owner;
    address gasStation;
    address feeOracleMngr;
    uint256 maxSwap = 2 ether;
    uint256 toll = 100; // 10%

    function setUp() public {
        portal = new MockPortal();
        owner = makeAddr("owner");
        gasStation = makeAddr("gasStation");
        feeOracleMngr = makeAddr("feeOracleMngr");
        address proxyAdminOwner = makeAddr("padmin");

        // We use a FeeOracleV1 as our IConversionRateOracle for the gas pump
        address feeOracleImpl = address(new FeeOracleV1());
        IFeeOracleV1.ChainFeeParams[] memory chainFeeParams = new IFeeOracleV1.ChainFeeParams[](1);

        // Omni's fee params - only one that's needed
        chainFeeParams[0] = IFeeOracleV1.ChainFeeParams({
            chainId: portal.omniChainId(),
            gasPrice: 1e9, // 1 Gwei
            postsTo: portal.omniChainId(),
            toNativeRate: 1e5 // 10 OMNI * (1e5 / 1e6) = 1 ETH
        });

        feeOracle = FeeOracleV1(
            address(
                new TransparentUpgradeableProxy(
                    feeOracleImpl,
                    proxyAdminOwner,
                    abi.encodeCall(
                        FeeOracleV1.initialize,
                        (
                            owner,
                            feeOracleMngr,
                            100_000, // baseGasLimit
                            0, // protocolFee
                            chainFeeParams
                        )
                    )
                )
            )
        );

        address pumpImpl = address(new OmniGasPump());
        pump = OmniGasPump(
            payable(address(
                    new TransparentUpgradeableProxy(
                        pumpImpl,
                        proxyAdminOwner,
                        abi.encodeCall(
                            OmniGasPump.initialize,
                            (OmniGasPump.InitParams({
                                    oracle: address(feeOracle),
                                    gasStation: gasStation,
                                    maxSwap: maxSwap,
                                    toll: toll,
                                    portal: address(portal),
                                    owner: owner
                                }))
                        )
                    )
                ))
        );
    }

    function test_fillUp() public {
        address recipient = makeAddr("recipient");
        uint256 initialBalance = 1000 ether;
        uint256 fee = pump.xfee();
        uint64 omniChainId = portal.omniChainId();
        vm.deal(recipient, initialBalance);

        // no zero recipient
        vm.expectRevert("OmniGasPump: no zero addr");
        pump.fillUp(address(0));

        // requires fee
        vm.expectRevert("OmniGasPump: insufficient fee");
        pump.fillUp(recipient);

        // requires < maxSwap
        vm.expectRevert("OmniGasPump: over max");
        vm.prank(recipient);
        pump.fillUp{ value: fee + maxSwap + 1 }(recipient);

        // takes toll, updates owed, emits xcall
        uint256 swapAmt = 1 ether;
        uint256 expectedOwedETH = swapAmt - (swapAmt * toll / pump.TOLL_DENOM());
        uint256 expectedOwedOMNI = expectedOwedETH * 10; // 1 ETH == 10 OMNI

        vm.expectCall(
            address(portal),
            fee,
            abi.encodeCall(
                IOmniPortal.xcall,
                (
                    omniChainId,
                    ConfLevel.Latest,
                    gasStation,
                    abi.encodeWithSelector(OmniGasStation.settleUp.selector, recipient, expectedOwedOMNI),
                    pump.SETTLE_GAS()
                )
            )
        );
        vm.prank(recipient);

        // fillUp returns amount swapped for, not total
        uint256 actualOwedOMNI = pump.fillUp{ value: fee + swapAmt }(recipient);

        assertEq(expectedOwedOMNI, actualOwedOMNI);
        assertEq(expectedOwedOMNI, pump.owed(recipient));

        // fillUp again, assert owed accumulates
        swapAmt = 2 ether;
        expectedOwedETH += swapAmt - (swapAmt * toll / pump.TOLL_DENOM());
        expectedOwedOMNI = expectedOwedETH * 10; // 1 ETH == 10 OMNI

        vm.expectCall(
            address(portal),
            fee,
            abi.encodeCall(
                IOmniPortal.xcall,
                (
                    omniChainId,
                    ConfLevel.Latest,
                    gasStation,
                    abi.encodeWithSelector(OmniGasStation.settleUp.selector, recipient, expectedOwedOMNI),
                    pump.SETTLE_GAS()
                )
            )
        );
        vm.prank(recipient);

        // fillUp returns amount swapped for, not total
        actualOwedOMNI += pump.fillUp{ value: fee + swapAmt }(recipient);

        assertEq(expectedOwedOMNI, actualOwedOMNI);
        assertEq(expectedOwedOMNI, pump.owed(recipient));
    }

    function test_setMaxSwap() public {
        uint256 newMaxSwap = 2e18 ether;

        // only owner
        address notOwner = address(0x456);
        vm.prank(notOwner);
        vm.expectRevert(abi.encodeWithSelector(OwnableUpgradeable.OwnableUnauthorizedAccount.selector, notOwner));
        pump.setMaxSwap(newMaxSwap);

        // success
        vm.prank(owner);
        pump.setMaxSwap(newMaxSwap);
        assertEq(pump.maxSwap(), newMaxSwap);
    }

    function test_setOmniGasStation() public {
        address newGasStation = makeAddr("newGasStation");

        // only owner
        address notOwner = address(0x456);
        vm.prank(notOwner);
        vm.expectRevert(abi.encodeWithSelector(OwnableUpgradeable.OwnableUnauthorizedAccount.selector, notOwner));
        pump.setGasStation(newGasStation);

        // success
        vm.prank(owner);
        pump.setGasStation(newGasStation);
        assertEq(pump.gasStation(), newGasStation);
    }

    function test_setToll() public {
        uint256 newToll = 100; // 10%

        // only owner
        address notOwner = address(0x456);
        vm.prank(notOwner);
        vm.expectRevert(abi.encodeWithSelector(OwnableUpgradeable.OwnableUnauthorizedAccount.selector, notOwner));
        pump.setToll(newToll);

        // success
        vm.prank(owner);
        pump.setToll(newToll);
        assertEq(pump.toll(), newToll);
    }

    function test_setOracle() public {
        address newOracle = address(0x123);

        // only owner
        address notOwner = address(0x456);
        vm.prank(notOwner);
        vm.expectRevert(abi.encodeWithSelector(OwnableUpgradeable.OwnableUnauthorizedAccount.selector, notOwner));
        pump.setOracle(newOracle);

        // success
        vm.prank(owner);
        pump.setOracle(newOracle);
        assertEq(address(pump.oracle()), newOracle);
    }

    function test_pause() public {
        // only owner can pause
        address notOwner = address(0x456);
        vm.prank(notOwner);
        vm.expectRevert(abi.encodeWithSelector(OwnableUpgradeable.OwnableUnauthorizedAccount.selector, notOwner));
        pump.pause();

        // owner can pause
        vm.prank(owner);
        pump.pause();

        assertTrue(pump.paused());

        // fillUp is paused
        address recipient = makeAddr("recipient");
        vm.expectRevert(abi.encodeWithSelector(PausableUpgradeable.EnforcedPause.selector));
        pump.fillUp(recipient);

        // only owner can unpause
        vm.prank(notOwner);
        vm.expectRevert(abi.encodeWithSelector(OwnableUpgradeable.OwnableUnauthorizedAccount.selector, notOwner));
        pump.unpause();

        // owner can unpause
        vm.prank(owner);
        pump.unpause();

        assertFalse(pump.paused());

        // fillUp is unpaused
        vm.expectRevert("OmniGasPump: insufficient fee"); // reverts, but not becasue its paused
        pump.fillUp(recipient);
    }

    function test_withdraw() public {
        vm.deal(address(pump), 10 ether);

        // no zero address
        vm.prank(owner);
        vm.expectRevert("OmniGasPump: no zero addr");
        pump.withdraw(address(0));

        // only owner
        address notOwner = address(0x456);
        vm.prank(notOwner);
        vm.expectRevert(abi.encodeWithSelector(OwnableUpgradeable.OwnableUnauthorizedAccount.selector, notOwner));
        pump.withdraw(owner);

        // success
        address to = makeAddr("to");
        vm.prank(owner);
        pump.withdraw(to);
        assertEq(address(pump).balance, 0);
        assertEq(to.balance, 10 ether);
    }

    /// @notice Test that GasPump.quote is accurate
    function testFuzz_quote(uint32 _targetOMNI) public view {
        uint256 targetOMNI = bound(_targetOMNI, 0.1 ether, 10 ether);
        uint256 neededETH = pump.quote(targetOMNI);

        (uint256 dryRunOMNI, bool wouldSucceed, string memory reason) = pump.dryFillUp(neededETH);

        assertTrue(wouldSucceed, reason);

        // assert quoted and actual within 10 wei of each other
        // allows for rounding errors in different between taking / undoing toll
        _assertWithinEpsilon(dryRunOMNI, targetOMNI, 10);
    }

    function _assertWithinEpsilon(uint256 a, uint256 b, uint256 epsilon) internal pure {
        assertTrue(a >= b - epsilon && a <= b + epsilon);
    }
}
