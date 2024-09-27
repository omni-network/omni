// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.24;

import { OwnableUpgradeable } from "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import { TransparentUpgradeableProxy } from "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import { FeeOracleV2 } from "src/xchain/FeeOracleV2.sol";
import { IFeeOracleV2 } from "src/interfaces/IFeeOracleV2.sol";
import { Test } from "forge-std/Test.sol";

/**
 * @title FeeOracleV2_Test
 * @dev Test of FeeOracleV2
 */
contract FeeOracleV2_Test is Test {
    IFeeOracleV2 feeOracle;
    address manager;
    address owner;

    uint256 baseGasLimit = 100_000;
    uint256 protocolFee = 1 gwei;

    uint64 chainAId = 1;
    uint64 chainBId = 2;
    uint64 chainCId = 3;

    function setUp() public {
        manager = makeAddr("manager");
        owner = makeAddr("owner");

        IFeeOracleV2.FeeParams[] memory feeParams = new IFeeOracleV2.FeeParams[](3);
        feeParams[0] =
            IFeeOracleV2.FeeParams({ chainId: chainAId, execGasPrice: 1 gwei, dataGasPrice: 2 gwei, toNativeRate: 1e6 });
        feeParams[1] =
            IFeeOracleV2.FeeParams({ chainId: chainBId, execGasPrice: 2 gwei, dataGasPrice: 3 gwei, toNativeRate: 2e6 });
        feeParams[2] =
            IFeeOracleV2.FeeParams({ chainId: chainCId, execGasPrice: 3 gwei, dataGasPrice: 4 gwei, toNativeRate: 5e3 });

        address impl = address(new FeeOracleV2());
        feeOracle = FeeOracleV2(
            address(
                new TransparentUpgradeableProxy(
                    address(impl),
                    owner,
                    abi.encodeWithSelector(
                        FeeOracleV2.initialize.selector, owner, manager, baseGasLimit, protocolFee, feeParams
                    )
                )
            )
        );
    }

    function test_feeFor() public {
        uint64 destChainId = chainBId;
        uint64 gasLimit = 200_000;
        bytes memory data = abi.encodeWithSignature("test()");

        vm.startPrank(manager);

        // test feeFor supported chain is expected
        // we do not duplicate feeFor logic here, rather we test that the fee is
        // reasonable, the function does not revert, and changes when fee params change

        uint256 fee = feeOracle.feeFor(destChainId, data, gasLimit);
        assertEq(fee > feeOracle.execGasPrice(destChainId) * gasLimit, true); // should be higher thatn just gas price
        assertEq(fee, feeOracle.feeFor(destChainId, data, gasLimit)); // should be stable

        // change exec gas price
        uint256 gasPrice = feeOracle.execGasPrice(destChainId);

        feeOracle.setExecGasPrice(destChainId, gasPrice * 2);
        assertEq(fee < feeOracle.feeFor(destChainId, data, gasLimit), true); // should be higher now

        feeOracle.setExecGasPrice(destChainId, gasPrice / 2);
        assertEq(fee > feeOracle.feeFor(destChainId, data, gasLimit), true); // should be lower now

        feeOracle.setExecGasPrice(destChainId, gasPrice); // reset

        // change to native rate
        uint256 toNativeRate = feeOracle.toNativeRate(destChainId);

        feeOracle.setToNativeRate(destChainId, toNativeRate * 2);
        assertEq(fee < feeOracle.feeFor(destChainId, data, gasLimit), true); // should be higher

        feeOracle.setToNativeRate(destChainId, toNativeRate / 2);
        assertEq(fee > feeOracle.feeFor(destChainId, data, gasLimit), true); // should be lower

        feeOracle.setToNativeRate(destChainId, toNativeRate); // reset

        // change data gas price

        uint256 dataGasPrice = feeOracle.dataGasPrice(destChainId);

        feeOracle.setDataGasPrice(destChainId, dataGasPrice * 2);
        assertEq(fee < feeOracle.feeFor(destChainId, data, gasLimit), true); // should be higher

        feeOracle.setDataGasPrice(destChainId, dataGasPrice / 2);
        assertEq(fee > feeOracle.feeFor(destChainId, data, gasLimit), true); // should be lower

        // increaes with calldata length
        bytes memory data2 = abi.encodeWithSignature("test(uint256)", 123);
        assertEq(feeOracle.feeFor(destChainId, data2, gasLimit) > fee, true); // should be higher

        // reverts for unsupported chain
        vm.expectRevert("FeeOracleV2: no fee params");
        feeOracle.feeFor(123_456, data, gasLimit);

        vm.stopPrank();
    }

    function test_setManager() public {
        address newManager = makeAddr("newManager");

        // only owner can set manager
        address notOwner = address(0x456);
        vm.expectRevert(abi.encodeWithSelector(OwnableUpgradeable.OwnableUnauthorizedAccount.selector, notOwner));
        vm.prank(notOwner);
        feeOracle.setManager(newManager);

        // cannot set zero manager
        vm.expectRevert("FeeOracleV2: no zero manager");
        vm.prank(owner);
        feeOracle.setManager(address(0));

        // set manager
        vm.prank(owner);
        feeOracle.setManager(newManager);
        assertEq(feeOracle.manager(), newManager);
    }

    function test_setExecGasPrice() public {
        uint64 destChainId = chainAId;
        uint256 newGasPrice = feeOracle.execGasPrice(destChainId) + 1 gwei;

        // only manager can set gas price
        vm.expectRevert("FeeOracleV2: not manager");
        feeOracle.setExecGasPrice(destChainId, newGasPrice);

        // no zero gas price
        vm.expectRevert("FeeOracleV2: no zero gas price");
        vm.prank(manager);
        feeOracle.setExecGasPrice(destChainId, 0);

        // no zero chain id
        vm.expectRevert("FeeOracleV2: no zero chain id");
        vm.prank(manager);
        feeOracle.setExecGasPrice(0, newGasPrice);

        // set gas price
        vm.prank(manager);
        feeOracle.setExecGasPrice(destChainId, newGasPrice);
        assertEq(feeOracle.execGasPrice(destChainId), newGasPrice);
    }

    function test_setDataGasPrice() public {
        uint64 destChainId = chainAId;
        uint256 newGasPrice = feeOracle.dataGasPrice(destChainId) + 1 gwei;

        // only manager can set gas price
        vm.expectRevert("FeeOracleV2: not manager");
        feeOracle.setDataGasPrice(destChainId, newGasPrice);

        // no zero gas price
        vm.expectRevert("FeeOracleV2: no zero gas price");
        vm.prank(manager);
        feeOracle.setDataGasPrice(destChainId, 0);

        // no zero chain id
        vm.expectRevert("FeeOracleV2: no zero chain id");
        vm.prank(manager);
        feeOracle.setDataGasPrice(0, newGasPrice);

        // set gas price
        vm.prank(manager);
        feeOracle.setDataGasPrice(destChainId, newGasPrice);
        assertEq(feeOracle.dataGasPrice(destChainId), newGasPrice);
    }

    function test_setProtocolFee() public {
        uint256 newProtocolFee = feeOracle.protocolFee() + 1 gwei;

        // only owner can set protocol fee
        address notOwner = address(0x456);
        vm.expectRevert(abi.encodeWithSelector(OwnableUpgradeable.OwnableUnauthorizedAccount.selector, notOwner));
        vm.prank(notOwner);
        feeOracle.setProtocolFee(newProtocolFee);

        // set protocol fee
        vm.prank(owner);
        feeOracle.setProtocolFee(newProtocolFee);
        assertEq(feeOracle.protocolFee(), newProtocolFee);
    }

    function test_setBaseGasLimit() public {
        uint64 newBaseGasLimit = feeOracle.baseGasLimit() + 10_000;

        // only owner can set base gas limit
        address notOwner = address(0x456);
        vm.expectRevert(abi.encodeWithSelector(OwnableUpgradeable.OwnableUnauthorizedAccount.selector, notOwner));
        vm.prank(notOwner);
        feeOracle.setBaseGasLimit(newBaseGasLimit);

        // set base gas limit
        vm.prank(owner);
        feeOracle.setBaseGasLimit(newBaseGasLimit);
        assertEq(feeOracle.baseGasLimit(), newBaseGasLimit);
    }

    function test_setToNativeRate() public {
        uint64 destChainId = chainAId;
        uint256 newToNativeRate = feeOracle.toNativeRate(destChainId) + 1;

        // only manager can set to native rate
        vm.expectRevert("FeeOracleV2: not manager");
        feeOracle.setToNativeRate(destChainId, newToNativeRate);

        // no zero rate
        vm.expectRevert("FeeOracleV2: no zero rate");
        vm.prank(manager);
        feeOracle.setToNativeRate(destChainId, 0);

        // no zero chain id
        vm.expectRevert("FeeOracleV2: no zero chain id");
        vm.prank(manager);
        feeOracle.setToNativeRate(0, newToNativeRate);

        // set to native rate
        vm.prank(manager);
        feeOracle.setToNativeRate(destChainId, newToNativeRate);
        assertEq(feeOracle.toNativeRate(destChainId), newToNativeRate);
    }

    function test_bulkSetFeeParams() public {
        IFeeOracleV2.FeeParams[] memory feeParams = new IFeeOracleV2.FeeParams[](4);

        feeParams[0] = IFeeOracleV2.FeeParams({
            chainId: chainAId,
            execGasPrice: feeOracle.execGasPrice(chainAId) + 1 gwei,
            dataGasPrice: feeOracle.dataGasPrice(chainAId) + 2 gwei,
            toNativeRate: feeOracle.toNativeRate(chainAId) + 1
        });

        feeParams[1] = IFeeOracleV2.FeeParams({
            chainId: chainBId,
            execGasPrice: feeOracle.execGasPrice(chainBId) + 2 gwei,
            dataGasPrice: feeOracle.dataGasPrice(chainBId) + 3 gwei,
            toNativeRate: feeOracle.toNativeRate(chainBId) + 2
        });

        feeParams[2] = IFeeOracleV2.FeeParams({
            chainId: chainCId,
            execGasPrice: feeOracle.execGasPrice(chainCId) + 3 gwei,
            dataGasPrice: feeOracle.dataGasPrice(chainCId) + 4 gwei,
            toNativeRate: feeOracle.toNativeRate(chainCId) + 3
        });

        feeParams[3] = IFeeOracleV2.FeeParams({
            chainId: 123_456, // new chain id
            execGasPrice: 123 gwei,
            dataGasPrice: 456 gwei,
            toNativeRate: 2e18
        });

        // only manager can set bulk fee params
        vm.expectRevert("FeeOracleV2: not manager");
        feeOracle.bulkSetFeeParams(feeParams);

        // set bulk fee params
        vm.prank(manager);
        feeOracle.bulkSetFeeParams(feeParams);

        for (uint256 i = 0; i < feeParams.length; i++) {
            IFeeOracleV2.FeeParams memory p = feeParams[i];
            assertEq(feeOracle.execGasPrice(p.chainId), p.execGasPrice);
            assertEq(feeOracle.dataGasPrice(p.chainId), p.dataGasPrice);
            assertEq(feeOracle.toNativeRate(p.chainId), p.toNativeRate);
        }
    }
}
