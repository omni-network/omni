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

    uint96 protocolFee = 1 gwei;

    uint16 gasTokenA = 1;
    uint16 gasTokenB = 2;

    uint64 chainAId = 1;
    uint64 chainBId = 2;
    uint64 chainCId = 3;

    uint64 dataCostAId = 1;
    uint64 dataCostBId = 2;

    uint64 gasLimit = 200_000;
    bytes data = abi.encodeWithSignature("test()");

    function setUp() public {
        manager = makeAddr("manager");
        owner = makeAddr("owner");

        IFeeOracleV2.FeeParams[] memory feeParams = new IFeeOracleV2.FeeParams[](3);
        feeParams[0] = IFeeOracleV2.FeeParams({
            gasToken: gasTokenA, baseGasLimit: 100_000, chainId: chainAId, gasPrice: 2 gwei, dataCostId: dataCostAId
        });
        feeParams[1] = IFeeOracleV2.FeeParams({
            gasToken: gasTokenB, baseGasLimit: 100_000, chainId: chainBId, gasPrice: 4 gwei, dataCostId: dataCostBId
        });
        feeParams[2] = IFeeOracleV2.FeeParams({
            gasToken: gasTokenA, baseGasLimit: 100_000, chainId: chainCId, gasPrice: 6 gwei, dataCostId: dataCostBId
        });

        IFeeOracleV2.DataCostParams[] memory dataCostParams = new IFeeOracleV2.DataCostParams[](2);
        dataCostParams[0] = IFeeOracleV2.DataCostParams({
            gasToken: gasTokenA, baseBytes: 100, id: dataCostAId, gasPrice: 1 gwei, gasPerByte: 1e2
        });
        dataCostParams[1] = IFeeOracleV2.DataCostParams({
            gasToken: gasTokenB, baseBytes: 200, id: dataCostBId, gasPrice: 2 gwei, gasPerByte: 2e2
        });

        IFeeOracleV2.ToNativeRateParams[] memory toNativeRateParams = new IFeeOracleV2.ToNativeRateParams[](2);
        toNativeRateParams[0] = IFeeOracleV2.ToNativeRateParams({ gasToken: gasTokenA, nativeRate: 1e18 });
        toNativeRateParams[1] = IFeeOracleV2.ToNativeRateParams({ gasToken: gasTokenB, nativeRate: 2e18 });

        address impl = address(new FeeOracleV2());
        feeOracle = FeeOracleV2(
            address(
                new TransparentUpgradeableProxy(
                    address(impl),
                    owner,
                    abi.encodeWithSelector(
                        FeeOracleV2.initialize.selector,
                        owner,
                        manager,
                        protocolFee,
                        feeParams,
                        dataCostParams,
                        toNativeRateParams
                    )
                )
            )
        );
    }

    function test_feeFor() public {
        uint64 dataCostId = feeOracle.execDataCostId(chainBId);

        uint16 execGasToken = feeOracle.execGasToken(chainBId);
        uint16 dataGasToken = feeOracle.dataGasToken(dataCostId);

        vm.startPrank(manager);

        // test feeFor supported chain is expected
        // we do not duplicate feeFor logic here, rather we test that the fee is
        // reasonable, the function does not revert, and changes when fee params change

        uint256 fee = feeOracle.feeFor(chainBId, data, gasLimit);
        assertEq(fee > feeOracle.execGasPrice(chainBId) * gasLimit, true); // should be higher than just gas price
        assertEq(fee, feeOracle.feeFor(chainBId, data, gasLimit)); // should be stable

        // change exec gas price
        uint64 execGasPrice = feeOracle.execGasPrice(chainBId);

        feeOracle.setExecGasPrice(chainBId, execGasPrice * 2);
        assertEq(fee < feeOracle.feeFor(chainBId, data, gasLimit), true); // should be higher now

        feeOracle.setExecGasPrice(chainBId, execGasPrice / 2);
        assertEq(fee > feeOracle.feeFor(chainBId, data, gasLimit), true); // should be lower now

        feeOracle.setExecGasPrice(chainBId, execGasPrice); // reset

        // change data gas price
        uint64 dataGasPrice = feeOracle.dataGasPrice(dataCostId);

        feeOracle.setDataGasPrice(dataCostId, dataGasPrice * 2);
        assertEq(fee < feeOracle.feeFor(chainBId, data, gasLimit), true); // should be higher

        feeOracle.setDataGasPrice(dataCostId, dataGasPrice / 2);
        assertEq(fee > feeOracle.feeFor(chainBId, data, gasLimit), true); // should be lower

        feeOracle.setDataGasPrice(dataCostId, dataGasPrice); // reset

        // change base gas limit
        uint32 baseGasLimit = feeOracle.baseGasLimit(chainBId);

        feeOracle.setBaseGasLimit(chainBId, baseGasLimit * 2);
        assertEq(fee < feeOracle.feeFor(chainBId, data, gasLimit), true); // should be higher

        feeOracle.setBaseGasLimit(chainBId, baseGasLimit / 2);
        assertEq(fee > feeOracle.feeFor(chainBId, data, gasLimit), true); // should be lower

        feeOracle.setBaseGasLimit(chainBId, baseGasLimit); // reset

        // change base data bytes buffer
        uint32 baseBytes = feeOracle.baseBytes(dataCostId);

        feeOracle.setBaseBytes(dataCostId, baseBytes * 2);
        assertEq(fee < feeOracle.feeFor(chainBId, data, gasLimit), true); // should be higher

        feeOracle.setBaseBytes(dataCostId, baseBytes / 2);
        assertEq(fee > feeOracle.feeFor(chainBId, data, gasLimit), true); // should be lower

        feeOracle.setBaseBytes(dataCostId, baseBytes); // reset

        // change exec to native rate
        uint256 execToNativeRate = feeOracle.toNativeRate(chainBId);

        feeOracle.setToNativeRate(execGasToken, execToNativeRate * 2);
        assertEq(fee < feeOracle.feeFor(chainBId, data, gasLimit), true); // should be higher

        feeOracle.setToNativeRate(execGasToken, execToNativeRate / 2);
        assertEq(fee > feeOracle.feeFor(chainBId, data, gasLimit), true); // should be lower

        feeOracle.setToNativeRate(execGasToken, execToNativeRate); // reset

        // change data to native rate
        uint256 dataToNativeRate = feeOracle.tokenToNativeRate(dataGasToken);

        feeOracle.setToNativeRate(dataGasToken, dataToNativeRate * 2);
        assertEq(fee < feeOracle.feeFor(chainBId, data, gasLimit), true); // should be higher

        feeOracle.setToNativeRate(dataGasToken, dataToNativeRate / 2);
        assertEq(fee > feeOracle.feeFor(chainBId, data, gasLimit), true); // should be lower

        feeOracle.setToNativeRate(dataGasToken, dataToNativeRate); // reset

        // change gas per byte
        uint64 gasPerByte = feeOracle.dataGasPerByte(dataCostId);

        feeOracle.setGasPerByte(dataCostId, gasPerByte * 2);
        assertEq(fee < feeOracle.feeFor(chainBId, data, gasLimit), true); // should be higher

        feeOracle.setGasPerByte(dataCostId, gasPerByte / 2);
        assertEq(fee > feeOracle.feeFor(chainBId, data, gasLimit), true); // should be lower

        feeOracle.setGasPerByte(dataCostId, gasPerByte); // reset

        // increaes with calldata length
        bytes memory data2 = abi.encodeWithSignature("test(uint256)", 123);
        assertEq(feeOracle.feeFor(chainBId, data2, gasLimit) > fee, true); // should be higher

        // reverts for unsupported exec chain
        vm.expectRevert(IFeeOracleV2.NoFeeParams.selector);
        feeOracle.feeFor(123_456, data, gasLimit);

        // reverts for unsupported data chain
        feeOracle.setDataCostId(chainBId, 123_456);
        vm.expectRevert(IFeeOracleV2.NoFeeParams.selector);
        feeOracle.feeFor(chainBId, data, gasLimit);

        vm.stopPrank();
    }

    function test_bulkSetExecFeeParams() public {
        IFeeOracleV2.FeeParams[] memory feeParams = new IFeeOracleV2.FeeParams[](4);

        feeParams[0] = IFeeOracleV2.FeeParams({
            gasToken: gasTokenA,
            baseGasLimit: 100_000,
            chainId: chainAId,
            gasPrice: feeOracle.execGasPrice(chainAId) + 1 gwei,
            dataCostId: dataCostAId
        });

        feeParams[1] = IFeeOracleV2.FeeParams({
            gasToken: gasTokenB,
            baseGasLimit: 100_000,
            chainId: chainBId,
            gasPrice: feeOracle.execGasPrice(chainBId) + 2 gwei,
            dataCostId: dataCostBId
        });

        feeParams[2] = IFeeOracleV2.FeeParams({
            gasToken: gasTokenA,
            baseGasLimit: 100_000,
            chainId: chainCId,
            gasPrice: feeOracle.execGasPrice(chainCId) + 3 gwei,
            dataCostId: dataCostBId
        });

        feeParams[3] = IFeeOracleV2.FeeParams({
            gasToken: gasTokenB,
            baseGasLimit: 100_000,
            chainId: 123_456, // new chain id
            gasPrice: 123 gwei,
            dataCostId: dataCostBId
        });

        // only manager can set bulk exec fee params
        vm.expectRevert(IFeeOracleV2.NotManager.selector);
        feeOracle.bulkSetFeeParams(feeParams);

        // set bulk exec fee params
        vm.prank(manager);
        feeOracle.bulkSetFeeParams(feeParams);

        for (uint256 i = 0; i < feeParams.length; i++) {
            IFeeOracleV2.FeeParams memory p = feeParams[i];
            assertEq(feeOracle.execGasToken(p.chainId), p.gasToken);
            assertEq(feeOracle.baseGasLimit(p.chainId), p.baseGasLimit);
            assertEq(feeOracle.execGasPrice(p.chainId), p.gasPrice);
            assertEq(feeOracle.execDataCostId(p.chainId), p.dataCostId);
        }
    }

    function test_bulkSetDataCostParams() public {
        IFeeOracleV2.DataCostParams[] memory feeParams = new IFeeOracleV2.DataCostParams[](3);

        feeParams[0] = IFeeOracleV2.DataCostParams({
            gasToken: gasTokenA,
            baseBytes: feeOracle.baseBytes(dataCostAId) + 1,
            id: dataCostAId,
            gasPrice: feeOracle.dataGasPrice(dataCostAId) + 1 gwei,
            gasPerByte: 1e2
        });

        feeParams[1] = IFeeOracleV2.DataCostParams({
            gasToken: gasTokenB,
            baseBytes: feeOracle.baseBytes(dataCostBId) + 2,
            id: dataCostBId,
            gasPrice: feeOracle.dataGasPrice(dataCostBId) + 2 gwei,
            gasPerByte: 2e2
        });

        feeParams[2] = IFeeOracleV2.DataCostParams({
            gasToken: gasTokenA,
            baseBytes: 123,
            id: 123_456, // new data cost id
            gasPrice: 123 gwei,
            gasPerByte: 3e2
        });

        // only manager can set bulk data fee params
        vm.expectRevert(IFeeOracleV2.NotManager.selector);
        feeOracle.bulkSetDataCostParams(feeParams);

        // set bulk data fee params
        vm.prank(manager);
        feeOracle.bulkSetDataCostParams(feeParams);

        for (uint256 i = 0; i < feeParams.length; i++) {
            IFeeOracleV2.DataCostParams memory p = feeParams[i];
            assertEq(feeOracle.dataGasToken(p.id), p.gasToken);
            assertEq(feeOracle.baseBytes(p.id), p.baseBytes);
            assertEq(feeOracle.dataGasPrice(p.id), p.gasPrice);
            assertEq(feeOracle.dataGasPerByte(p.id), p.gasPerByte);
        }
    }

    function test_bulkSetToNativeRate() public {
        IFeeOracleV2.ToNativeRateParams[] memory toNativeRateParams = new IFeeOracleV2.ToNativeRateParams[](3);

        toNativeRateParams[0] = IFeeOracleV2.ToNativeRateParams({
            gasToken: gasTokenA, nativeRate: feeOracle.tokenToNativeRate(gasTokenA) + 1
        });

        toNativeRateParams[1] = IFeeOracleV2.ToNativeRateParams({
            gasToken: gasTokenB, nativeRate: feeOracle.tokenToNativeRate(gasTokenB) + 2
        });

        toNativeRateParams[2] = IFeeOracleV2.ToNativeRateParams({ gasToken: 3, nativeRate: 3e18 });

        // only manager can set bulk to native rate
        vm.expectRevert(IFeeOracleV2.NotManager.selector);
        feeOracle.bulkSetToNativeRate(toNativeRateParams);

        // set bulk to native rate params
        vm.prank(manager);
        feeOracle.bulkSetToNativeRate(toNativeRateParams);

        for (uint256 i = 0; i < toNativeRateParams.length; i++) {
            IFeeOracleV2.ToNativeRateParams memory p = toNativeRateParams[i];
            assertEq(feeOracle.tokenToNativeRate(p.gasToken), p.nativeRate);
        }
    }

    function test_setExecGasPrice() public {
        uint64 destChainId = chainAId;
        uint64 newGasPrice = feeOracle.execGasPrice(destChainId) + 1 gwei;

        // only manager can set gas price
        vm.expectRevert(IFeeOracleV2.NotManager.selector);
        feeOracle.setExecGasPrice(destChainId, newGasPrice);

        // no zero gas price
        vm.expectRevert(IFeeOracleV2.ZeroGasPrice.selector);
        vm.prank(manager);
        feeOracle.setExecGasPrice(destChainId, 0);

        // no zero chain id
        vm.expectRevert(IFeeOracleV2.ZeroChainId.selector);
        vm.prank(manager);
        feeOracle.setExecGasPrice(0, newGasPrice);

        // set gas price
        vm.prank(manager);
        feeOracle.setExecGasPrice(destChainId, newGasPrice);
        assertEq(feeOracle.execGasPrice(destChainId), newGasPrice);
    }

    function test_setDataGasPrice() public {
        uint64 dataCostId = dataCostAId;
        uint64 newGasPrice = feeOracle.dataGasPrice(dataCostId) + 1 gwei;

        // only manager can set gas price
        vm.expectRevert(IFeeOracleV2.NotManager.selector);
        feeOracle.setDataGasPrice(dataCostId, newGasPrice);

        // no zero gas price
        vm.expectRevert(IFeeOracleV2.ZeroGasPrice.selector);
        vm.prank(manager);
        feeOracle.setDataGasPrice(dataCostId, 0);

        // no zero chain id
        vm.expectRevert(IFeeOracleV2.ZeroDataCostId.selector);
        vm.prank(manager);
        feeOracle.setDataGasPrice(0, newGasPrice);

        // set gas price
        vm.prank(manager);
        feeOracle.setDataGasPrice(dataCostId, newGasPrice);
        assertEq(feeOracle.dataGasPrice(dataCostId), newGasPrice);
    }

    function test_setBaseGasLimit() public {
        uint64 destChainId = chainCId;
        uint32 newBaseGasLimit = feeOracle.baseGasLimit(destChainId) + 10_000;

        // only manager can set gas price
        vm.expectRevert(IFeeOracleV2.NotManager.selector);
        feeOracle.setBaseGasLimit(destChainId, newBaseGasLimit);

        // no zero chain id
        vm.expectRevert(IFeeOracleV2.ZeroChainId.selector);
        vm.prank(manager);
        feeOracle.setBaseGasLimit(0, newBaseGasLimit);

        // set base gas limit
        vm.prank(manager);
        feeOracle.setBaseGasLimit(destChainId, newBaseGasLimit);
        assertEq(feeOracle.baseGasLimit(destChainId), newBaseGasLimit);
    }

    function test_setBaseBytes() public {
        uint64 dataCostId = dataCostBId;
        uint32 newBaseBytes = feeOracle.baseBytes(dataCostId) + 100;

        // only manager can set data size buffer
        vm.expectRevert(IFeeOracleV2.NotManager.selector);
        feeOracle.setBaseBytes(dataCostId, newBaseBytes);

        // no zero chain id
        vm.expectRevert(IFeeOracleV2.ZeroDataCostId.selector);
        vm.prank(manager);
        feeOracle.setBaseBytes(0, newBaseBytes);

        // set data size buffer
        vm.prank(manager);
        feeOracle.setBaseBytes(dataCostId, newBaseBytes);
        assertEq(feeOracle.baseBytes(dataCostId), newBaseBytes);
    }

    function test_setDataCostId() public {
        uint64 destChainId = chainAId;
        uint64 newDataCostId = dataCostBId;

        // only manager can set data cost id
        vm.expectRevert(IFeeOracleV2.NotManager.selector);
        feeOracle.setDataCostId(destChainId, newDataCostId);

        // no zero chain id
        vm.expectRevert(IFeeOracleV2.ZeroChainId.selector);
        vm.prank(manager);
        feeOracle.setDataCostId(0, newDataCostId);

        // no zero data cost id
        vm.expectRevert(IFeeOracleV2.ZeroDataCostId.selector);
        vm.prank(manager);
        feeOracle.setDataCostId(destChainId, 0);

        // set data cost id
        vm.prank(manager);
        feeOracle.setDataCostId(destChainId, newDataCostId);
        assertEq(feeOracle.execDataCostId(destChainId), newDataCostId);
    }

    function test_setGasPerByte() public {
        uint64 dataCostId = dataCostAId;
        uint64 newGasPerByte = feeOracle.dataGasPerByte(dataCostId) + 16;

        // only manager can set gas per byte
        vm.expectRevert(IFeeOracleV2.NotManager.selector);
        feeOracle.setGasPerByte(dataCostId, newGasPerByte);

        // no zero data cost id
        vm.expectRevert(IFeeOracleV2.ZeroDataCostId.selector);
        vm.prank(manager);
        feeOracle.setGasPerByte(0, newGasPerByte);

        // no zero gas per byte
        vm.expectRevert(IFeeOracleV2.ZeroGasPerByte.selector);
        vm.prank(manager);
        feeOracle.setGasPerByte(dataCostId, 0);

        // set gas per byte
        vm.prank(manager);
        feeOracle.setGasPerByte(dataCostId, newGasPerByte);
        assertEq(feeOracle.dataGasPerByte(dataCostId), newGasPerByte);
    }

    function test_setToNativeRate() public {
        uint16 gasToken = gasTokenB;
        uint256 newToNativeRate = feeOracle.tokenToNativeRate(gasToken) + 1e6;

        // only manager can set to native rate
        vm.expectRevert(IFeeOracleV2.NotManager.selector);
        feeOracle.setToNativeRate(gasToken, newToNativeRate);

        // no zero rate
        vm.expectRevert(IFeeOracleV2.ZeroNativeRate.selector);
        vm.prank(manager);
        feeOracle.setToNativeRate(gasToken, 0);

        // no zero chain id
        vm.expectRevert(IFeeOracleV2.ZeroGasToken.selector);
        vm.prank(manager);
        feeOracle.setToNativeRate(0, newToNativeRate);

        // set to native rate
        vm.prank(manager);
        feeOracle.setToNativeRate(gasToken, newToNativeRate);
        assertEq(feeOracle.tokenToNativeRate(gasToken), newToNativeRate);
    }

    function test_setProtocolFee() public {
        uint96 newProtocolFee = feeOracle.protocolFee() + 1 gwei;

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

    function test_setManager() public {
        address newManager = makeAddr("newManager");

        // only owner can set manager
        address notOwner = address(0x456);
        vm.expectRevert(abi.encodeWithSelector(OwnableUpgradeable.OwnableUnauthorizedAccount.selector, notOwner));
        vm.prank(notOwner);
        feeOracle.setManager(newManager);

        // cannot set zero manager
        vm.expectRevert(IFeeOracleV2.ZeroAddress.selector);
        vm.prank(owner);
        feeOracle.setManager(address(0));

        // set manager
        vm.prank(owner);
        feeOracle.setManager(newManager);
        assertEq(feeOracle.manager(), newManager);
    }
}
