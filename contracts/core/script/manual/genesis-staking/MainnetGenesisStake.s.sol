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
    address internal admin = 0xd09DD1126385877352d24B669Fd68f462200756E;

    ICreateX internal createX = ICreateX(0xba5Ed099633D3B313e4D5F7bdc1305d3c28ba5Ed);
    IERC20 internal omni = IERC20(0x36E66fbBce51e4cD5bd3C62B637Eb411b18949D4);
    IOmniPortal internal portal = IOmniPortal(0x5e9A8Aa213C912Bf54C86bf64aDB8ed6A79C04d1);
    ISolverNetInbox internal inbox = ISolverNetInbox(0x8FCFcd0B4Fa2cc2965a3c7F27995B0A43F210dB8);

    address internal validator = 0x8be1aBb26435fc1AF39Fc88DF9499f626094f9AF;

    bytes32 internal genesisStakeImplSalt = 0xa779fc675db318dab004ab8d538cb320d0013f4200fc20573d7a708801d08213;
    bytes32 internal merkleDistributorSalt = 0xa779fc675db318dab004ab8d538cb320d0013f42008dbcdb4f2b296d009be3b2;

    address internal expectedGenesisStakeImplAddr = 0x00000000000022bff6Fa8AC173fA1A87b1E4a04A;
    address internal expectedMerkleDistributorAddr = 0x0000000000009bBE7DE32eaF5a88E355907B680E;

    GenesisStake internal genesisStake = GenesisStake(0xD2639676dA3dEA5491d27DA19340556b3a7d58B8);
    MerkleDistributorWithDeadline internal merkleDistributor =
        MerkleDistributorWithDeadline(0x0000000000009bBE7DE32eaF5a88E355907B680E);

    uint256 internal endTime = 1_746_057_599; // 2025-04-29 23:59:59 UTC
    uint256 internal index = 7987;
    uint256 internal depositAmount = 100 ether;
    uint256 internal rewardAmount = 843_524_630_884_109_952;
    bytes32[] internal leaves = new bytes32[](1);
    bytes32[][] internal proofs = new bytes32[][](1);
    bytes32 internal root;

    function fullDryRun() public {
        _prepMerkleTree();
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

        _fundRewards();
        _upgrade();

        require(merkleDistributor.owner() == admin, "MerkleDistributor owner mismatch");
    }

    function deployGenesisStakeImpl() public {
        vm.startBroadcast();
        if (msg.sender != deployer) revert("Unintended deployer");
        bool deployed = _deployGenesisStakeImpl();
        vm.stopBroadcast();

        if (!deployed) console2.log("GenesisStake implementation already deployed");
        else console2.log("GenesisStake implementation deployed");
    }

    function deployMerkleDistributor() public {
        _prepMerkleTree();

        vm.startBroadcast();
        if (msg.sender != deployer) revert("Unintended deployer");
        bool deployed = _deployMerkleDistributor();
        vm.stopBroadcast();

        if (!deployed) console2.log("MerkleDistributor already deployed");
        else console2.log("MerkleDistributor deployed");
    }

    function _preventBroadcast() internal {
        (VmSafe.CallerMode mode,,) = vm.readCallers();
        require(mode == VmSafe.CallerMode.None, "no broadcast");
    }

    function _maybeEtchSolverNetInbox() internal {
        if (address(inbox).code.length != 0) return;

        // NOTE: If this is ever required again, we need to pass in the correct address for the mailbox
        SolverNetInbox inboxImpl = new SolverNetInbox(address(0));
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
        address _admin = proxyAdmin.owner();

        vm.prank(_admin);
        proxyAdmin.upgradeAndCall(ITransparentUpgradeableProxy(address(genesisStake)), expectedGenesisStakeImplAddr, "");

        address newImpl = EIP1967Helper.getImplementation(address(genesisStake));
        require(newImpl == expectedGenesisStakeImplAddr, "Applied GenesisStake implementation addr mismatch");
        return true;
    }

    function _prepMerkleTree() internal {
        m = new CompleteMerkle();

        leaves[0] = keccak256(abi.encodePacked(index, deployer, rewardAmount));
        proofs[0].push(hex"72456ff6c5ae451e3dfe2e76a9c1ae4f316514ef6c5216cd7fc67d84ecbb6474");
        proofs[0].push(hex"9f68cbac4e4c8977504176df4e02ee45338eaf7be24fce9e0c551da8537e68a5");
        proofs[0].push(hex"6b47664d11f9d0de4e89cd2f3ffd553048477292ba4e045f78e1d3b0d363866f");
        proofs[0].push(hex"1682c88343c865daf4dccf035cd064bc5ff350e0ef439d8a21afde30e5e8f117");
        proofs[0].push(hex"6ed83d2dad1ac9577f67d641d1ca1cd5023b8703e5c839afe77ea7bf95844d3e");
        proofs[0].push(hex"9af730042a389ed429eeaf204a31b494d76d855911ccc39f8345a7ee9c5f4e4d");
        proofs[0].push(hex"112d14bdc992bb8b71a7179d91021ef7ebc733925ed5323a05d142af8008f28b");
        proofs[0].push(hex"c9f2448b67173ca1a5b0564c58bf6249ae64a0d3aadd39d99d34e33f2e2e18cd");
        proofs[0].push(hex"b2aa63149d0fca20a0586c7e291ee86fcd6a9ab7ce722d4b5d78f20325a26f60");
        proofs[0].push(hex"aa862cee2ca1c076902205197547d8035cfc82fee22f7722630d0557c3c4f548");
        proofs[0].push(hex"20818a4fe98b62e2c916c9caf037c10ac22e7971394b06f2dc9b64f53f8ea3ad");
        proofs[0].push(hex"021461cfd282c77c6915461a1b8a51cdf020b129cbf561ea89f244a501f6b8b1");
        proofs[0].push(hex"bdcdcfb79c1b7113a7fe38bb82e734e92fc20362c88fc2a0543f89ba49e01edc");
        proofs[0].push(hex"a6e860fc7ca37e6fbb248f16250f1a3c57e70e4354921fa526372b142429b149");

        root = hex"5b95c1e6c7fe3ebe64c4b095810a61feb5b219e91261ad181c3149ff26364c3e";

        require(m.verifyProof(root, proofs[0], leaves[0]), "Proof is invalid");
    }

    function _deployMerkleDistributor() internal returns (bool) {
        if (address(expectedMerkleDistributorAddr).code.length != 0) return false;

        (VmSafe.CallerMode mode,,) = vm.readCallers();
        if (mode == VmSafe.CallerMode.None) vm.prank(deployer);
        merkleDistributor = MerkleDistributorWithDeadline(
            createX.deployCreate3(
                merkleDistributorSalt,
                abi.encodePacked(
                    type(MerkleDistributorWithDeadline).creationCode,
                    abi.encode(
                        admin, address(omni), root, endTime, address(portal), address(genesisStake), address(inbox)
                    )
                )
            )
        );

        console2.log("MerkleDistributor constructor args:");
        console2.logBytes(
            abi.encode(admin, address(omni), root, endTime, address(portal), address(genesisStake), address(inbox))
        );

        require(address(merkleDistributor) == expectedMerkleDistributorAddr, "MerkleDistributor addr mismatch");
        return true;
    }

    function _fundRewards() internal {
        deal(address(omni), deployer, rewardAmount);

        vm.startPrank(deployer);
        omni.transfer(address(merkleDistributor), rewardAmount);
        vm.stopPrank();
    }

    function _upgrade() internal {
        vm.prank(deployer);
        merkleDistributor.upgradeStake(validator, index, rewardAmount, proofs[0]);
    }
}
