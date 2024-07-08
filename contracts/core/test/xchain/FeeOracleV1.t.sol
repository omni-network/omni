// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.24;

import { OwnableUpgradeable } from "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import { Base } from "./common/Base.sol";
import { IFeeOracleV1 } from "src/interfaces/IFeeOracleV1.sol";
import { console } from "forge-std/console.sol";

/**
 * @title FeeOracleV1_Test
 * @dev Test of FeeOracleV1
 */
contract FeeOracleV1_Test is Base {
    function test_feeFor() public {
        uint64 destChainId = chainBId; // using chain b, because postsTo is chain a (see Fixtures.sol)
        uint64 gasLimit = 200_000;
        bytes memory data = abi.encodeWithSignature("test()");

        vm.startPrank(feeOracleManager);

        // test feeFor supported chain is expected
        // we do not duplicate feeFor logic here, rather we test that the fee is
        // reasonable, the function does not revert, and changes when fee params change

        uint256 fee = feeOracle.feeFor(destChainId, data, gasLimit);
        assertEq(fee > feeOracle.gasPriceOn(destChainId) * gasLimit, true); // should be higher thatn just gas price
        assertEq(fee, feeOracle.feeFor(destChainId, data, gasLimit)); // should be stable

        // change gas price
        uint256 gasPrice = feeOracle.gasPriceOn(destChainId);

        feeOracle.setGasPrice(destChainId, gasPrice * 2);
        assertEq(fee < feeOracle.feeFor(destChainId, data, gasLimit), true); // should be higher now

        feeOracle.setGasPrice(destChainId, gasPrice / 2);
        assertEq(fee > feeOracle.feeFor(destChainId, data, gasLimit), true); // should be lower now

        feeOracle.setGasPrice(destChainId, gasPrice); // reset

        // change to native rate
        uint256 toNativeRate = feeOracle.toNativeRate(destChainId);

        feeOracle.setToNativeRate(destChainId, toNativeRate * 2);
        assertEq(fee < feeOracle.feeFor(destChainId, data, gasLimit), true); // should be higher

        feeOracle.setToNativeRate(destChainId, toNativeRate / 2);
        assertEq(fee > feeOracle.feeFor(destChainId, data, gasLimit), true); // should be lower

        feeOracle.setToNativeRate(destChainId, toNativeRate); // reset

        // change postsTo gas price
        uint64 postsTo = feeOracle.feeParams(destChainId).postsTo;
        require(postsTo != destChainId, "test_feeFor: postsTo == destChainId");

        uint256 postsToGasPrice = feeOracle.gasPriceOn(postsTo);

        feeOracle.setGasPrice(postsTo, postsToGasPrice * 2);
        assertEq(fee < feeOracle.feeFor(destChainId, data, gasLimit), true); // should be higher

        feeOracle.setGasPrice(postsTo, postsToGasPrice / 2);
        assertEq(fee > feeOracle.feeFor(destChainId, data, gasLimit), true); // should be lower

        // increaes with calldata length
        bytes memory data2 = abi.encodeWithSignature("test(uint256)", 123);
        assertEq(feeOracle.feeFor(destChainId, data2, gasLimit) > fee, true); // should be higher

        // reverts for unsupported chain
        vm.expectRevert("FeeOracleV1: no fee params");
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
            postsTo: thisChainId,
            gasPrice: feeOracle.gasPriceOn(thisChainId) + 1 gwei,
            toNativeRate: feeOracle.toNativeRate(thisChainId) + 1
        });

        feeParams[1] = IFeeOracleV1.ChainFeeParams({
            chainId: chainAId,
            postsTo: chainAId,
            gasPrice: feeOracle.gasPriceOn(chainAId) + 2 gwei,
            toNativeRate: feeOracle.toNativeRate(chainAId) + 2
        });

        feeParams[2] = IFeeOracleV1.ChainFeeParams({
            chainId: chainBId,
            postsTo: chainAId,
            gasPrice: feeOracle.gasPriceOn(chainBId) + 3 gwei,
            toNativeRate: feeOracle.toNativeRate(chainBId) + 3
        });

        feeParams[3] = IFeeOracleV1.ChainFeeParams({
            chainId: 123_456, // new chain id
            postsTo: 1,
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
            assertEq(feeOracle.gasPriceOn(p.chainId), p.gasPrice);
            assertEq(feeOracle.toNativeRate(p.chainId), p.toNativeRate);
        }
    }
}
