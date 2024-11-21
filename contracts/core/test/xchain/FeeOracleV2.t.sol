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

    uint64 chainDId = 4;
    uint64 chainEId = 5;

    uint64 gasLimit = 200_000;
    bytes data = abi.encodeWithSignature("test()");

    function setUp() public {
        manager = makeAddr("manager");
        owner = makeAddr("owner");

        IFeeOracleV2.ExecFeeParams[] memory execFeeParams = new IFeeOracleV2.ExecFeeParams[](3);
        execFeeParams[0] = IFeeOracleV2.ExecFeeParams({
            chainId: chainAId,
            postsTo: chainDId,
            execGasPrice: 2 gwei,
            toNativeRate: 1e6
        });
        execFeeParams[1] = IFeeOracleV2.ExecFeeParams({
            chainId: chainBId,
            postsTo: chainDId,
            execGasPrice: 4 gwei,
            toNativeRate: 3e6
        });
        execFeeParams[2] = IFeeOracleV2.ExecFeeParams({
            chainId: chainCId,
            postsTo: chainEId,
            execGasPrice: 6 gwei,
            toNativeRate: 5e3
        });

        IFeeOracleV2.DataFeeParams[] memory dataFeeParams = new IFeeOracleV2.DataFeeParams[](2);
        dataFeeParams[0] =
            IFeeOracleV2.DataFeeParams({ chainId: chainDId, sizeBuffer: 100, dataGasPrice: 1 gwei, toNativeRate: 2e6 });
        dataFeeParams[1] =
            IFeeOracleV2.DataFeeParams({ chainId: chainEId, sizeBuffer: 200, dataGasPrice: 2 gwei, toNativeRate: 4e3 });

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
                        baseGasLimit,
                        protocolFee,
                        execFeeParams,
                        dataFeeParams
                    )
                )
            )
        );
    }

    function test_feeFor() public {
        uint64 destChainId = chainBId;
        uint64 dataChainId = feeOracle.execPostsTo(destChainId);

        vm.startPrank(manager);

        // test feeFor supported chain is expected
        // we do not duplicate feeFor logic here, rather we test that the fee is
        // reasonable, the function does not revert, and changes when fee params change

        uint256 fee = feeOracle.feeFor(destChainId, data, gasLimit);
        assertEq(fee > feeOracle.execGasPrice(destChainId) * gasLimit, true); // should be higher thatn just gas price
        assertEq(fee, feeOracle.feeFor(destChainId, data, gasLimit)); // should be stable

        // change exec gas price
        uint64 execGasPrice = feeOracle.execGasPrice(destChainId);

        feeOracle.setExecGasPrice(destChainId, execGasPrice * 2);
        assertEq(fee < feeOracle.feeFor(destChainId, data, gasLimit), true); // should be higher now

        feeOracle.setExecGasPrice(destChainId, execGasPrice / 2);
        assertEq(fee > feeOracle.feeFor(destChainId, data, gasLimit), true); // should be lower now

        feeOracle.setExecGasPrice(destChainId, execGasPrice); // reset

        // change exec to native rate
        uint64 execToNativeRate = uint64(feeOracle.execToNativeRate(destChainId));

        feeOracle.setExecToNativeRate(destChainId, execToNativeRate * 2);
        assertEq(fee < feeOracle.feeFor(destChainId, data, gasLimit), true); // should be higher

        feeOracle.setExecToNativeRate(destChainId, execToNativeRate / 2);
        assertEq(fee > feeOracle.feeFor(destChainId, data, gasLimit), true); // should be lower

        feeOracle.setExecToNativeRate(destChainId, execToNativeRate); // reset

        // change data size buffer
        uint64 sizeBuffer = feeOracle.dataSizeBuffer(dataChainId);

        feeOracle.setDataSizeBuffer(dataChainId, sizeBuffer * 2);
        assertEq(fee < feeOracle.feeFor(destChainId, data, gasLimit), true); // should be higher

        feeOracle.setDataSizeBuffer(dataChainId, sizeBuffer / 2);
        assertEq(fee > feeOracle.feeFor(destChainId, data, gasLimit), true); // should be lower

        feeOracle.setDataSizeBuffer(dataChainId, sizeBuffer); // reset

        // change data gas price
        uint64 dataGasPrice = feeOracle.dataGasPrice(dataChainId);

        feeOracle.setDataGasPrice(dataChainId, dataGasPrice * 2);
        assertEq(fee < feeOracle.feeFor(destChainId, data, gasLimit), true); // should be higher

        feeOracle.setDataGasPrice(dataChainId, dataGasPrice / 2);
        assertEq(fee > feeOracle.feeFor(destChainId, data, gasLimit), true); // should be lower

        feeOracle.setDataGasPrice(dataChainId, dataGasPrice); // reset

        // change data to native rate
        uint64 dataToNativeRate = uint64(feeOracle.dataToNativeRate(dataChainId));

        feeOracle.setDataToNativeRate(dataChainId, dataToNativeRate * 2);
        assertEq(fee < feeOracle.feeFor(destChainId, data, gasLimit), true); // should be higher

        feeOracle.setDataToNativeRate(dataChainId, dataToNativeRate / 2);
        assertEq(fee > feeOracle.feeFor(destChainId, data, gasLimit), true); // should be lower

        feeOracle.setDataToNativeRate(dataChainId, dataToNativeRate); // reset

        // increaes with calldata length
        bytes memory data2 = abi.encodeWithSignature("test(uint256)", 123);
        assertEq(feeOracle.feeFor(destChainId, data2, gasLimit) > fee, true); // should be higher

        // reverts for unsupported exec chain
        vm.expectRevert("FeeOracleV2: no exec fee params");
        feeOracle.feeFor(123_456, data, gasLimit);

        // reverts for unsupported data chain
        feeOracle.setExecPostsTo(destChainId, 123_456);
        vm.expectRevert("FeeOracleV2: no data fee params");
        feeOracle.feeFor(destChainId, data, gasLimit);

        vm.stopPrank();
    }

    function test_bulkSetExecFeeParams() public {
        IFeeOracleV2.ExecFeeParams[] memory feeParams = new IFeeOracleV2.ExecFeeParams[](4);

        feeParams[0] = IFeeOracleV2.ExecFeeParams({
            chainId: chainAId,
            postsTo: chainEId,
            execGasPrice: feeOracle.execGasPrice(chainAId) + 1 gwei,
            toNativeRate: uint64(feeOracle.execToNativeRate(chainAId) + 1)
        });

        feeParams[1] = IFeeOracleV2.ExecFeeParams({
            chainId: chainBId,
            postsTo: chainEId,
            execGasPrice: feeOracle.execGasPrice(chainBId) + 2 gwei,
            toNativeRate: uint64(feeOracle.execToNativeRate(chainBId) + 2)
        });

        feeParams[2] = IFeeOracleV2.ExecFeeParams({
            chainId: chainCId,
            postsTo: chainDId,
            execGasPrice: feeOracle.execGasPrice(chainCId) + 3 gwei,
            toNativeRate: uint64(feeOracle.execToNativeRate(chainCId) + 3)
        });

        feeParams[3] = IFeeOracleV2.ExecFeeParams({
            chainId: 123_456, // new chain id
            postsTo: chainDId,
            execGasPrice: 123 gwei,
            toNativeRate: 2e18
        });

        // only manager can set bulk exec fee params
        vm.expectRevert("FeeOracleV2: not manager");
        feeOracle.bulkSetExecFeeParams(feeParams);

        // set bulk exec fee params
        vm.prank(manager);
        feeOracle.bulkSetExecFeeParams(feeParams);

        for (uint256 i = 0; i < feeParams.length; i++) {
            IFeeOracleV2.ExecFeeParams memory p = feeParams[i];
            assertEq(feeOracle.execPostsTo(p.chainId), p.postsTo);
            assertEq(feeOracle.execGasPrice(p.chainId), p.execGasPrice);
            assertEq(feeOracle.execToNativeRate(p.chainId), p.toNativeRate);
        }
    }

    function test_bulkSetDataFeeParams() public {
        IFeeOracleV2.DataFeeParams[] memory feeParams = new IFeeOracleV2.DataFeeParams[](3);

        feeParams[0] = IFeeOracleV2.DataFeeParams({
            chainId: chainDId,
            sizeBuffer: feeOracle.dataSizeBuffer(chainDId) + 1,
            dataGasPrice: feeOracle.dataGasPrice(chainDId) + 1 gwei,
            toNativeRate: uint64(feeOracle.dataToNativeRate(chainDId) + 1)
        });

        feeParams[1] = IFeeOracleV2.DataFeeParams({
            chainId: chainEId,
            sizeBuffer: feeOracle.dataSizeBuffer(chainEId) + 2,
            dataGasPrice: feeOracle.dataGasPrice(chainEId) + 2 gwei,
            toNativeRate: uint64(feeOracle.dataToNativeRate(chainEId) + 2)
        });

        feeParams[2] = IFeeOracleV2.DataFeeParams({
            chainId: 123_456, // new chain id
            sizeBuffer: 123,
            dataGasPrice: 123 gwei,
            toNativeRate: 2e18
        });

        // only manager can set bulk data fee params
        vm.expectRevert("FeeOracleV2: not manager");
        feeOracle.bulkSetDataFeeParams(feeParams);

        // set bulk data fee params
        vm.prank(manager);
        feeOracle.bulkSetDataFeeParams(feeParams);

        for (uint256 i = 0; i < feeParams.length; i++) {
            IFeeOracleV2.DataFeeParams memory p = feeParams[i];
            assertEq(feeOracle.dataSizeBuffer(p.chainId), p.sizeBuffer);
            assertEq(feeOracle.dataGasPrice(p.chainId), p.dataGasPrice);
            assertEq(feeOracle.dataToNativeRate(p.chainId), p.toNativeRate);
        }
    }

    function test_setExecPostsTo() public {
        uint64 destChainId = chainAId;
        uint64 newPostsTo = chainEId;

        // only manager can set exec posts to
        vm.expectRevert("FeeOracleV2: not manager");
        feeOracle.setExecPostsTo(destChainId, newPostsTo);

        // no zero chain id
        vm.expectRevert("FeeOracleV2: no zero chain id");
        vm.prank(manager);
        feeOracle.setExecPostsTo(0, newPostsTo);

        vm.expectRevert("FeeOracleV2: no zero chain id");
        vm.prank(manager);
        feeOracle.setExecPostsTo(destChainId, 0);

        // set exec posts to
        vm.prank(manager);
        feeOracle.setExecPostsTo(destChainId, newPostsTo);
        assertEq(feeOracle.execPostsTo(destChainId), newPostsTo);
    }

    function test_setDataSizeBuffer() public {
        uint64 destChainId = chainAId;
        uint64 dataChainId = feeOracle.execPostsTo(destChainId);
        uint64 sizeBuffer = feeOracle.dataSizeBuffer(dataChainId);

        // only manager can set data size buffer
        vm.expectRevert("FeeOracleV2: not manager");
        feeOracle.setDataSizeBuffer(destChainId, sizeBuffer + 1);

        // no zero chain id
        vm.expectRevert("FeeOracleV2: no zero chain id");
        vm.prank(manager);
        feeOracle.setDataSizeBuffer(0, sizeBuffer + 1);

        // set data size buffer
        vm.prank(manager);
        feeOracle.setDataSizeBuffer(dataChainId, sizeBuffer + 1);
        assertEq(feeOracle.dataSizeBuffer(dataChainId), sizeBuffer + 1);
    }

    function test_setExecGasPrice() public {
        uint64 destChainId = chainAId;
        uint64 newGasPrice = feeOracle.execGasPrice(destChainId) + 1 gwei;

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
        uint64 newGasPrice = feeOracle.dataGasPrice(destChainId) + 1 gwei;

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

    function test_setExecToNativeRate() public {
        uint64 destChainId = chainAId;
        uint64 newExecToNativeRate = uint64(feeOracle.execToNativeRate(destChainId) + 1);

        // only manager can set to native rate
        vm.expectRevert("FeeOracleV2: not manager");
        feeOracle.setExecToNativeRate(destChainId, newExecToNativeRate);

        // no zero rate
        vm.expectRevert("FeeOracleV2: no zero rate");
        vm.prank(manager);
        feeOracle.setExecToNativeRate(destChainId, 0);

        // no zero chain id
        vm.expectRevert("FeeOracleV2: no zero chain id");
        vm.prank(manager);
        feeOracle.setExecToNativeRate(0, newExecToNativeRate);

        // set exec to native rate
        vm.prank(manager);
        feeOracle.setExecToNativeRate(destChainId, newExecToNativeRate);
        assertEq(feeOracle.execToNativeRate(destChainId), newExecToNativeRate);
    }

    function test_setDataToNativeRate() public {
        uint64 destChainId = chainAId;
        uint64 newDataToNativeRate = uint64(feeOracle.dataToNativeRate(destChainId) + 1);

        // only manager can set data to native rate
        vm.expectRevert("FeeOracleV2: not manager");
        feeOracle.setDataToNativeRate(destChainId, newDataToNativeRate);

        // no zero rate
        vm.expectRevert("FeeOracleV2: no zero rate");
        vm.prank(manager);
        feeOracle.setDataToNativeRate(destChainId, 0);

        // no zero chain id
        vm.expectRevert("FeeOracleV2: no zero chain id");
        vm.prank(manager);
        feeOracle.setDataToNativeRate(0, newDataToNativeRate);

        // set data to native rate
        vm.prank(manager);
        feeOracle.setDataToNativeRate(destChainId, newDataToNativeRate);
        assertEq(feeOracle.dataToNativeRate(destChainId), newDataToNativeRate);
    }

    function test_setProtocolFee() public {
        uint72 newProtocolFee = feeOracle.protocolFee() + 1 gwei;

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
        uint24 newBaseGasLimit = feeOracle.baseGasLimit() + 10_000;

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
}
