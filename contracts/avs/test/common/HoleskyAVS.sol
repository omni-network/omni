// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.12;

import { IAVSDirectory } from "eigenlayer-contracts/src/contracts/interfaces/IAVSDirectory.sol";
import { IDelegationManager } from "src/ext/IDelegationManager.sol";

import { Create3 } from "core/deploy/Create3.sol";
import { OmniAVS } from "src/OmniAVS.sol";
import { EigenM2HoleskyDeployments } from "test/common/eigen/EigenM2HoleskyDeployments.sol";
import { StrategyParams } from "./StrategyParams.sol";

import { Test } from "forge-std/Test.sol";

contract HoleskyAVS is Test {
    bool allowlistEnabled = false;
    string metadaURI = "https://raw.githubusercontent.com/omni-network/omni/main/static/avs-metadata.json";

    /// @dev defines holesky deployment logic
    function deploy(
        address create3Factory,
        bytes32 create3Salt,
        address owner,
        address proxyAdmin,
        address portal,
        uint64 omniChainId,
        address ethStakeInbox,
        uint96 minOperatorStake,
        uint32 maxOperatorCount
    ) public returns (OmniAVS) {
        Create3 create3 = Create3(create3Factory);

        address impl = address(
            new OmniAVS(
                IDelegationManager(EigenM2HoleskyDeployments.DelegationManager),
                IAVSDirectory(EigenM2HoleskyDeployments.AVSDirectory)
            )
        );

        bytes memory proxyConstructorArgs = abi.encode(
            address(impl),
            proxyAdmin,
            abi.encodeWithSelector(
                OmniAVS.initialize.selector,
                owner,
                portal,
                omniChainId,
                ethStakeInbox,
                minOperatorStake,
                maxOperatorCount,
                StrategyParams.holesky(),
                metadaURI,
                allowlistEnabled
            )
        );

        bytes memory bytecode = abi.encodePacked(
            vm.getCode("TransparentUpgradeableProxy.sol:TransparentUpgradeableProxy.0.8.12"), proxyConstructorArgs
        );

        address proxy = create3.deploy(create3Salt, bytecode);

        return OmniAVS(proxy);
    }
}
