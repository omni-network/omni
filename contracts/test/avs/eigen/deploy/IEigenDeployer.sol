// SPDX-License-Identifier: BUSL-1.1
pragma solidity =0.8.12;

interface IEigenDeployer {
    struct Deployments {
        address proxyAdminOwner;
        address proxyAdmin;
        address pauserRegistry;
        address delegationManager;
        address eigenPodManager;
        address strategyManager;
        address slasher;
        address avsDirectory;
        address[] strategies;
    }

    function deploy() external returns (Deployments memory);
}
