// SPDX-License-Identifier: GPL-3.0-only
/* solhint-disable no-console */
pragma solidity 0.8.24;

import { Script, console2 } from "forge-std/Script.sol";
import { StdCheats } from "forge-std/StdCheats.sol";
import { VmSafe } from "forge-std/Vm.sol";

import { CompleteMerkle } from "murky/src/CompleteMerkle.sol";
import { EIP1967Helper } from "script/utils/EIP1967Helper.sol";

import { ICreateX } from "createx/src/ICreateX.sol";
import { IERC20 } from "@openzeppelin/contracts/interfaces/IERC20.sol";
import { IOmniPortal } from "src/interfaces/IOmniPortal.sol";
import { SolverNetInbox, ISolverNetInbox } from "solve/src/SolverNetInbox.sol";

import { GenesisStake } from "src/token/GenesisStake.sol";
import { MerkleDistributorWithDeadline } from "src/token/MerkleDistributorWithDeadline.sol";

import {
    TransparentUpgradeableProxy,
    ITransparentUpgradeableProxy
} from "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import { ProxyAdmin } from "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";

contract MainnetGenesisStakeScript is Script, StdCheats {
    CompleteMerkle internal m;

    address internal deployer = 0xA779fC675Db318dab004Ab8D538CB320D0013F42;

    ICreateX internal createX = ICreateX(0xba5Ed099633D3B313e4D5F7bdc1305d3c28ba5Ed);
    IERC20 internal omni = IERC20(0x36E66fbBce51e4cD5bd3C62B637Eb411b18949D4);
    IOmniPortal internal portal = IOmniPortal(0x5e9A8Aa213C912Bf54C86bf64aDB8ed6A79C04d1);
    ISolverNetInbox internal inbox = ISolverNetInbox(0x8FCFcd0B4Fa2cc2965a3c7F27995B0A43F210dB8);

    address internal validator = 0x8be1aBb26435fc1AF39Fc88DF9499f626094f9AF;

    bytes32 internal genesisStakeImplSalt = 0xa779fc675db318dab004ab8d538cb320d0013f42006fda006bab5cd1034643cc;
    bytes32 internal merkleDistributorSalt = 0xa779fc675db318dab004ab8d538cb320d0013f4200205ad9f17a619b0079c949;

    address internal expectedGenesisStakeImplAddr = 0x00000000000063B7931226e67CeF52d86085154d;
    address internal expectedMerkleDistributorAddr = 0x0000000000003E960a909Ef4F0E77699a4286711;

    GenesisStake internal genesisStake = GenesisStake(0xD2639676dA3dEA5491d27DA19340556b3a7d58B8);
    MerkleDistributorWithDeadline internal merkleDistributor; // = MerkleDistributorWithDeadline(0x5B46d1fA23584c071fa5478D709A189c452Eb050);

    uint256 internal endTime = block.timestamp + 30 days;
    uint256 internal depositAmount = 100 ether;
    uint256 internal rewardAmount = 100 ether;
    bytes32[] internal leaves = new bytes32[](2);
    bytes32[][] internal proofs = new bytes32[][](2);
    bytes32 internal root;

    function fullDryRun() public {
        _preventBroadcast();
        _maybeEtchSolverNetInbox();

        bool genesisStakeImplDeployed = _deployGenesisStakeImpl();
        if (genesisStakeImplDeployed) console2.log("GenesisStake implementation deployed");
        else console2.log("GenesisStake implementation already deployed");

        bool upgraded = _upgradeGenesisStake();
        if (upgraded) console2.log("GenesisStake proxy upgraded");
        else console2.log("GenesisStake proxy already upgraded");

        bool merkleDistributorDeployed = _deployMerkleDistributor();
        if (merkleDistributorDeployed) console2.log("MerkleDistributor deployed");
        else console2.log("MerkleDistributor already deployed");

        _dealApproveStakeAndFund();
        _migrate();
    }

    function deployGenesisStakeImpl() public {
        vm.startBroadcast();
        if (msg.sender != 0xA779fC675Db318dab004Ab8D538CB320D0013F42) revert("Unintended deployer");
        bool deployed = _deployGenesisStakeImpl();
        vm.stopBroadcast();

        if (!deployed) console2.log("GenesisStake implementation already deployed");
        else console2.log("GenesisStake implementation deployed");
    }

    function _preventBroadcast() internal {
        (VmSafe.CallerMode mode,,) = vm.readCallers();
        require(mode == VmSafe.CallerMode.None, "no broadcast");
    }

    function _maybeEtchSolverNetInbox() internal {
        if (address(inbox).code.length != 0) return;

        SolverNetInbox inboxImpl = new SolverNetInbox();
        TransparentUpgradeableProxy proxy = new TransparentUpgradeableProxy(
            address(inboxImpl),
            address(deployer),
            abi.encodeCall(SolverNetInbox.initialize, (deployer, deployer, address(portal)))
        );

        vm.etch(address(inbox), address(proxy).code);
    }

    function _deployGenesisStakeImpl() internal returns (bool) {
        if (address(expectedGenesisStakeImplAddr).code.length != 0) return false;

        (VmSafe.CallerMode mode,,) = vm.readCallers();
        if (mode == VmSafe.CallerMode.None) vm.prank(deployer);
        address genesisStakeImpl = createX.deployCreate3(
            genesisStakeImplSalt,
            abi.encodePacked(type(GenesisStake).creationCode, abi.encode(address(omni), expectedMerkleDistributorAddr))
        );

        console2.log("GenesisStake implementation constructor args:");
        console2.logBytes(abi.encode(address(omni), expectedMerkleDistributorAddr));

        require(genesisStakeImpl == expectedGenesisStakeImplAddr, "GenesisStake implementation addr mismatch");
        return true;
    }

    function _upgradeGenesisStake() internal returns (bool) {
        address impl = EIP1967Helper.getImplementation(address(genesisStake));
        if (impl == expectedGenesisStakeImplAddr) return false;

        ProxyAdmin proxyAdmin = ProxyAdmin(EIP1967Helper.getAdmin(address(genesisStake)));
        address admin = proxyAdmin.owner();

        vm.prank(admin);
        proxyAdmin.upgradeAndCall(ITransparentUpgradeableProxy(address(genesisStake)), expectedGenesisStakeImplAddr, "");

        address newImpl = EIP1967Helper.getImplementation(address(genesisStake));
        require(newImpl == expectedGenesisStakeImplAddr, "Applied GenesisStake implementation addr mismatch");
        return true;
    }

    function _prepMerkleTree() internal {
        m = new CompleteMerkle();

        leaves[0] = keccak256(abi.encodePacked(uint256(0), 0xA779fC675Db318dab004Ab8D538CB320D0013F42, rewardAmount));
        leaves[1] = keccak256(abi.encodePacked(uint256(1), address(0xb00f), rewardAmount));

        proofs[0] = m.getProof(leaves, 0);
        proofs[1] = m.getProof(leaves, 1);

        root = m.getRoot(leaves);

        require(m.verifyProof(root, proofs[0], leaves[0]), "Proof 0 is invalid");
        require(m.verifyProof(root, proofs[1], leaves[1]), "Proof 1 is invalid");
    }

    function _deployMerkleDistributor() internal returns (bool) {
        _prepMerkleTree();
        if (address(expectedMerkleDistributorAddr).code.length != 0) return false;

        vm.prank(deployer);
        merkleDistributor = MerkleDistributorWithDeadline(
            createX.deployCreate3(
                merkleDistributorSalt,
                abi.encodePacked(
                    type(MerkleDistributorWithDeadline).creationCode,
                    abi.encode(address(omni), root, endTime, address(portal), address(genesisStake), address(inbox))
                )
            )
        );

        console2.log("MerkleDistributor constructor args:");
        console2.logBytes(
            abi.encode(address(omni), root, endTime, address(portal), address(genesisStake), address(inbox))
        );

        require(address(merkleDistributor) == expectedMerkleDistributorAddr, "MerkleDistributor addr mismatch");
        return true;
    }

    function _dealApproveStakeAndFund() internal {
        bool isOpen = genesisStake.isOpen();

        deal(address(omni), deployer, (isOpen ? depositAmount : 0) + rewardAmount);

        vm.startPrank(deployer);
        if (isOpen) {
            omni.approve(address(genesisStake), type(uint256).max);
            genesisStake.stake(depositAmount);
        }

        omni.transfer(address(merkleDistributor), rewardAmount);
        vm.stopPrank();
    }

    function _migrate() internal {
        vm.prank(deployer);
        merkleDistributor.migrateToOmni(validator, 0, rewardAmount, proofs[0]);
    }
}
