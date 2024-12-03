// SPDX-License-Identifier: GPL-3.0-only
pragma solidity 0.8.24;

import { FeeOracleV2 } from "src/xchain/FeeOracleV2.sol";
import { IFeeOracleV2 } from "src/interfaces/IFeeOracleV2.sol";
import { Test } from "forge-std/Test.sol";
import { VmSafe } from "forge-std/Vm.sol";
import { EnumerableSetLib } from "solady/src/utils/EnumerableSetLib.sol";

contract FeeOracleV2PostUpdateTest is Test {
    using EnumerableSetLib for EnumerableSetLib.Uint256Set;

    FeeOracleV2 feeOracle;

    EnumerableSetLib.Uint256Set testnetChainIds;
    EnumerableSetLib.Uint256Set mainnetChainIds;

    // Make sure these match the gas token IDs in `lib/contracts/feeoraclev2/gastokens.go`
    uint16 constant OMNI = 1;
    uint16 constant ETH = 2;

    enum Network {
        Invalid,
        Testnet,
        Mainnet
    }

    function run(address oracle) public {
        (VmSafe.CallerMode mode,,) = vm.readCallers();
        require(mode == VmSafe.CallerMode.None, "no broadcast");

        Network network = _setup(oracle);
        _testFeeParams(network);
        _testDataCostParams(network);
        _testToNativeRateParams();
    }

    function _setup(address oracle) internal returns (Network) {
        feeOracle = FeeOracleV2(oracle);

        // Testnet chain ids
        testnetChainIds.add(164); // Omni Omega
        testnetChainIds.add(17_000); // Ethereum Holesky
        testnetChainIds.add(84_532); // Base Sepolia
        testnetChainIds.add(421_614); // Arbitrum Sepolia
        testnetChainIds.add(11_155_420); // Optimism Sepolia

        // Mainnet chain ids
        mainnetChainIds.add(1); // Ethereum Mainnet
        mainnetChainIds.add(10); // Optimism Mainnet
        mainnetChainIds.add(166); // Omni Mainnet
        mainnetChainIds.add(8453); // Base Mainnet
        mainnetChainIds.add(42_161); // Arbitrum Mainnet

        if (testnetChainIds.contains(block.chainid)) {
            return Network.Testnet;
        } else if (mainnetChainIds.contains(block.chainid)) {
            return Network.Mainnet;
        } else {
            revert("invalid network");
        }
    }

    function _testFeeParams(Network network) internal view {
        uint256[] memory chainIds = network == Network.Testnet ? testnetChainIds.values() : mainnetChainIds.values();
        for (uint256 i; i < chainIds.length; i++) {
            uint64 chainId = uint64(chainIds[i]);
            IFeeOracleV2.FeeParams memory feeParams = feeOracle.feeParams(chainId);
            assertGt(feeParams.gasToken, 0, "gas token must be set");
            assertGt(feeParams.chainId, 0, "chain id must be set");
            assertGt(feeParams.gasPrice, 0, "gas price must be set");
            assertGt(feeParams.dataCostId, 0, "data cost id must be set");
        }
    }

    function _testDataCostParams(Network network) internal view {
        uint256[] memory chainIds = network == Network.Testnet ? testnetChainIds.values() : mainnetChainIds.values();
        for (uint256 i; i < chainIds.length; i++) {
            uint64 chainId = uint64(chainIds[i]);
            IFeeOracleV2.DataCostParams memory dataCostParams = feeOracle.dataCostParams(chainId);
            assertGt(dataCostParams.gasToken, 0, "gas token must be set");
            assertGt(dataCostParams.id, 0, "data cost id must be set");
            assertGt(dataCostParams.gasPrice, 0, "gas price must be set");
            assertGt(dataCostParams.gasPerByte, 0, "gas per byte must be set");
        }
    }

    function _testToNativeRateParams() internal view {
        uint256 omniNativeRate = feeOracle.tokenToNativeRate(uint16(OMNI));
        uint256 ethNativeRate = feeOracle.tokenToNativeRate(uint16(ETH));

        assertGt(omniNativeRate, 0, "omni native rate must be set");
        assertGt(ethNativeRate, 0, "eth native rate must be set");
    }
}
