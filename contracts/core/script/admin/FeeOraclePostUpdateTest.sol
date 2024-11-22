// SPDX-License-Identifier: GPL-3.0-only
pragma solidity 0.8.24;

import { FeeOracleV2 } from "src/xchain/FeeOracleV2.sol";
// import { IFeeOracleV2 } from "src/interfaces/IFeeOracleV2.sol";
import { Test } from "forge-std/Test.sol";
// import { VmSafe } from "forge-std/Vm.sol";
import { EnumerableSetLib } from "solady/src/utils/EnumerableSetLib.sol";

contract FeeOraclePostUpdateTest is Test {
    using EnumerableSetLib for EnumerableSetLib.Uint256Set;

    FeeOracleV2 feeOracle;

    EnumerableSetLib.Uint256Set testnetChainIds;
    EnumerableSetLib.Uint256Set mainnetChainIds;

    enum Network {
        Invalid,
        Testnet,
        Mainnet
    }

    /*
    function run(address oracle) public {
        (VmSafe.CallerMode mode,,) = vm.readCallers();
        require(mode == VmSafe.CallerMode.None, "no broadcast");

        Network network = _setup(oracle);
        _testFeeParams(network);
    }

    function _setup(address oracle) internal returns (Network) {
        feeOracle = FeeOracleV2(oracle);

        // Testnet chain ids
        testnetChainIds.add(164);
        testnetChainIds.add(17_000);
        testnetChainIds.add(84_532);
        testnetChainIds.add(421_614);
        testnetChainIds.add(11_155_420);

        // Mainnet chain ids
        mainnetChainIds.add(1);
        mainnetChainIds.add(166);
        mainnetChainIds.add(10);
        mainnetChainIds.add(8453);
        mainnetChainIds.add(42_161);

        if (testnetChainIds.contains(block.chainid)) {
            return Network.Testnet;
        } else if (mainnetChainIds.contains(block.chainid)) {
            return Network.Mainnet;
        } else {
            revert("invalid network");
        }
    }

    function _testFeeParams(Network network) internal pure {
        if (network == Network.Testnet) _testTestnetFeeParams();
        else _testMainnetFeeParams();
    }

    function _testTestnetFeeParams() internal view {
        for (uint256 i; i < testnetChainIds.length(); i++) {
            uint64 chainId = uint64(testnetChainIds.at(i));
            IFeeOracleV2.FeeParams memory feeParams = feeOracle.feeParams(chainId);
            // TODO: test fee params once updated
        }
    }

    function _testMainnetFeeParams() internal view {
        for (uint256 i; i < mainnetChainIds.length(); i++) {
            uint64 chainId = uint64(mainnetChainIds.at(i));
            IFeeOracleV2.FeeParams memory feeParams = feeOracle.feeParams(chainId);
            // TODO: test fee params once updated
        }
    }
    */
}
