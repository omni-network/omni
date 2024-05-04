// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.24;

import { Base } from "./common/Base.sol";
import { IFeeOracleV1 } from "src/interfaces/IFeeOracleV1.sol";

/**
 * @title FeeOracleV1_Test
 * @dev Test of FeeOracleV1
 */
contract FeeOracleV1_Test is Base {
    function test_feeFor() public {
        uint64 destChainId = chainAId;
        uint64 gasLimit = 200_000;
        bytes memory data = abi.encodeWithSignature("test()");

        // test feeFor supported chain is expected

        uint256 fee = feeOracle.feeFor(destChainId, data, gasLimit);
        uint256 gasPrice =
            feeOracle.gasPriceOn(destChainId) * feeOracle.toNativeRate(destChainId) / feeOracle.CONVERSION_RATE_DENOM();
        assertEq(fee, feeOracle.protocolFee() + (gasPrice * gasLimit) + (feeOracle.baseGasLimit() * gasPrice));

        // test feeFor unsupported chain reverts

        vm.expectRevert("FeeOracleV1: no fee params");
        feeOracle.feeFor(123_456, data, gasLimit);
    }

    function test_setManager() public {
        address newManager = makeAddr("newManager");

        // only owner can set manager
        vm.expectRevert("Ownable: caller is not the owner");
        feeOracle.setManager(newManager);

        // cannot set zero manager
        vm.expectRevert("FeeOracleV1: no zero manager");
        vm.prank(owner);
        feeOracle.setManager(address(0));

        // set manager
        vm.prank(owner);
        feeOracle.setManager(newManager);
        assertEq(feeOracle.manager(), newManager);
    }

    function test_setGasPrice() public {
        uint64 destChainId = chainAId;
        uint256 newGasPrice = feeOracle.gasPriceOn(destChainId) + 1 gwei;

        // only manager can set gas price
        vm.expectRevert("FeeOracleV1: not manager");
        feeOracle.setGasPrice(destChainId, newGasPrice);

        // no zero gas price
        vm.expectRevert("FeeOracleV1: no zero gas price");
        vm.prank(feeOracleManager);
        feeOracle.setGasPrice(destChainId, 0);

        // no zero chain id
        vm.expectRevert("FeeOracleV1: no zero chain id");
        vm.prank(feeOracleManager);
        feeOracle.setGasPrice(0, newGasPrice);

        // set gas price
        vm.prank(feeOracleManager);
        feeOracle.setGasPrice(destChainId, newGasPrice);
        assertEq(feeOracle.gasPriceOn(destChainId), newGasPrice);
    }

    function test_setProtocolFee() public {
        uint256 newProtocolFee = feeOracle.protocolFee() + 1 gwei;

        // only owner can set protocol fee
        vm.expectRevert("Ownable: caller is not the owner");
        feeOracle.setProtocolFee(newProtocolFee);

        // set protocol fee
        vm.prank(owner);
        feeOracle.setProtocolFee(newProtocolFee);
        assertEq(feeOracle.protocolFee(), newProtocolFee);
    }

    function test_setBaseGasLimit() public {
        uint64 newBaseGasLimit = feeOracle.baseGasLimit() + 10_000;

        // only owner can set base gas limit
        vm.expectRevert("Ownable: caller is not the owner");
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
        vm.expectRevert("FeeOracleV1: not manager");
        feeOracle.setToNativeRate(destChainId, newToNativeRate);

        // no zero rate
        vm.expectRevert("FeeOracleV1: no zero rate");
        vm.prank(feeOracleManager);
        feeOracle.setToNativeRate(destChainId, 0);

        // no zero chain id
        vm.expectRevert("FeeOracleV1: no zero chain id");
        vm.prank(feeOracleManager);
        feeOracle.setToNativeRate(0, newToNativeRate);

        // set to native rate
        vm.prank(feeOracleManager);
        feeOracle.setToNativeRate(destChainId, newToNativeRate);
        assertEq(feeOracle.toNativeRate(destChainId), newToNativeRate);
    }

    function test_bulkSetFeeParams() public {
        IFeeOracleV1.ChainFeeParams[] memory feeParams = new IFeeOracleV1.ChainFeeParams[](4);

        feeParams[0] = IFeeOracleV1.ChainFeeParams({
            chainId: thisChainId,
            gasPrice: feeOracle.gasPriceOn(thisChainId) + 1 gwei,
            toNativeRate: feeOracle.toNativeRate(thisChainId) + 1
        });

        feeParams[1] = IFeeOracleV1.ChainFeeParams({
            chainId: chainAId,
            gasPrice: feeOracle.gasPriceOn(chainAId) + 2 gwei,
            toNativeRate: feeOracle.toNativeRate(chainAId) + 2
        });

        feeParams[2] = IFeeOracleV1.ChainFeeParams({
            chainId: chainBId,
            gasPrice: feeOracle.gasPriceOn(chainBId) + 3 gwei,
            toNativeRate: feeOracle.toNativeRate(chainBId) + 3
        });

        feeParams[3] = IFeeOracleV1.ChainFeeParams({
            chainId: 123_456, // new chain id
            gasPrice: 123 gwei,
            toNativeRate: 2e18
        });

        // only manager can set bulk fee params
        vm.expectRevert("FeeOracleV1: not manager");
        feeOracle.bulkSetFeeParams(feeParams);

        // set bulk fee params
        vm.prank(feeOracleManager);
        feeOracle.bulkSetFeeParams(feeParams);

        for (uint256 i = 0; i < feeParams.length; i++) {
            IFeeOracleV1.ChainFeeParams memory p = feeParams[i];
            assertEq(feeOracle.gasPriceOn(p.chainId), p.gasPrice);
            assertEq(feeOracle.toNativeRate(p.chainId), p.toNativeRate);
        }
    }
}
