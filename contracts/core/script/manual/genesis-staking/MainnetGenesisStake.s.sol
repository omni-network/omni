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

import { GenesisStakeV2 } from "src/token/GenesisStakeV2.sol";
import { MerkleDistributorWithoutDeadline } from "src/token/MerkleDistributorWithoutDeadline.sol";

import {
    TransparentUpgradeableProxy,
    ITransparentUpgradeableProxy
} from "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import { ProxyAdmin } from "@openzeppelin/contracts/proxy/transparent/ProxyAdmin.sol";

contract MainnetGenesisStakeScript is Script, StdCheats {
    CompleteMerkle internal m;

    address internal deployer = 0xA779fC675Db318dab004Ab8D538CB320D0013F42;
    address internal admin = 0xd09DD1126385877352d24B669Fd68f462200756E;
    address internal adminOfProxy = 0x42A72499eDDB0374ebFba44Fc880F82CCe736614;

    ICreateX internal createX = ICreateX(0xba5Ed099633D3B313e4D5F7bdc1305d3c28ba5Ed);
    IERC20 internal omni = IERC20(0x36E66fbBce51e4cD5bd3C62B637Eb411b18949D4);
    IOmniPortal internal portal = IOmniPortal(0x5e9A8Aa213C912Bf54C86bf64aDB8ed6A79C04d1);
    ISolverNetInbox internal inbox = ISolverNetInbox(0x8FCFcd0B4Fa2cc2965a3c7F27995B0A43F210dB8);

    address internal validator = 0x8be1aBb26435fc1AF39Fc88DF9499f626094f9AF;

    bytes32 internal genesisStakeImplSalt = 0xa779fc675db318dab004ab8d538cb320d0013f4200b8a3e6b8edde1903ec4dde;
    bytes32 internal merkleDistributorImplSalt = 0xa779fc675db318dab004ab8d538cb320d0013f4200897b8566748c57033e1d07;
    bytes32 internal merkleDistributorSalt = 0xa779fc675db318dab004ab8d538cb320d0013f4200fa0b61ab1783e803655002;

    address internal expectedGenesisStakeImplAddr = 0x000000000000A1052e28F4bDc995Db4b8bfA6608;
    address internal expectedMerkleDistributorImplAddr = 0x000000000000726A22A1B8620E5851e24BfE3290;
    address internal expectedMerkleDistributorAddr = 0x00000000000058607f45dD7535ebD9f44ef3766D;

    GenesisStakeV2 internal genesisStake = GenesisStakeV2(0xD2639676dA3dEA5491d27DA19340556b3a7d58B8);
    MerkleDistributorWithoutDeadline internal merkleDistributor;

    address internal user = 0x2bA3dD20Db1fd35D648be1BE582eC563895D78fA;
    uint256 internal index = 1638;
    //uint256 internal depositAmount = 100 ether;
    uint256 internal rewardAmount = 115_068_493_150_684_004_352;
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
        //_upgrade();
        _unstake();

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
            abi.encodePacked(
                type(GenesisStakeV2).creationCode, abi.encode(address(omni), expectedMerkleDistributorAddr)
            )
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

        leaves[0] = keccak256(abi.encodePacked(index, user, rewardAmount));
        proofs[0].push(hex"8718726df34b55313774b7949f1053b74d4ef64af47a3cad5cf04e852ebf50e5");
        proofs[0].push(hex"547fd56a523966816df814126f7bb9fae3d81b093370316387ba2ec7863cf0ad");
        proofs[0].push(hex"ec348c4497e3554cd97e943c6350db42d8bc89d7557ec0659a0cbf6568d281b1");
        proofs[0].push(hex"d5eeeab527c5d81b6fc0e4a9a6cc9ef769ff76a821eb5dcef85e5457ed52aea3");
        proofs[0].push(hex"d06104550a87126a380f9ef5f9b8da72b7473f72edd28b87138153ebfeb3296b");
        proofs[0].push(hex"878ed5c4764e39a8aaf1a134b8cc3c1311ba9fb1ff53fc87eb99f9b26f1cd720");
        proofs[0].push(hex"5f37e6bd3b1cf6221d91c906e3a3183c48d65a38b79ccc456ab223e98ebb0904");
        proofs[0].push(hex"d997217263bb7fdbf05a2a147f1e59986b62675a973696e4b5bd8e012640288c");
        proofs[0].push(hex"6b2f4b175e1598bca1348e47435a0e649374a157caac0187e9914c9067bb90c0");
        proofs[0].push(hex"39150938f4c90eff3ca2487771824053ed7bb28504879cbb50c6dcdcd84542d0");
        proofs[0].push(hex"05101c489c8a6d83f19a031a568d4cc139f8d1e22a4d9aee5d1dff2906e4025d");
        proofs[0].push(hex"77bf903da427e7d7c35cf75efcdf1918cff520b063bc34bbbbe089a9c5390766");
        proofs[0].push(hex"2b497b8e97ccceea2aadcaf8cff19fea1af63a93198c0289c199596f9a11a54b");
        proofs[0].push(hex"8cbb3263d4d5343920e924facf7b3e96edac51fbab1d38705beef1f6c1d4f7f9");

        root = hex"a58fa42aae5003f337b37d078164b635e4b324baca33c129e27ea3d7bec51003";

        require(m.verifyProof(root, proofs[0], leaves[0]), "Proof is invalid");
    }

    function _deployMerkleDistributor() internal returns (bool) {
        if (address(expectedMerkleDistributorAddr).code.length != 0) return false;

        (VmSafe.CallerMode mode,,) = vm.readCallers();
        if (mode == VmSafe.CallerMode.None) vm.startPrank(deployer);

        address merkleDistributorImpl = createX.deployCreate3(
            merkleDistributorImplSalt,
            abi.encodePacked(
                type(MerkleDistributorWithoutDeadline).creationCode,
                abi.encode(address(omni), root, address(portal), address(genesisStake), address(inbox))
            )
        );

        merkleDistributor = MerkleDistributorWithoutDeadline(
            createX.deployCreate3(
                merkleDistributorSalt,
                abi.encodePacked(
                    type(TransparentUpgradeableProxy).creationCode,
                    abi.encode(
                        merkleDistributorImpl,
                        adminOfProxy,
                        abi.encodeCall(MerkleDistributorWithoutDeadline.initialize, (admin))
                    )
                )
            )
        );
        vm.stopPrank();

        console2.log("MerkleDistributor constructor args:");
        console2.logBytes(
            abi.encode(admin, address(omni), root, address(portal), address(genesisStake), address(inbox))
        );

        require(
            address(merkleDistributorImpl) == expectedMerkleDistributorImplAddr,
            "MerkleDistributor implementation addr mismatch"
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
        vm.prank(user);
        merkleDistributor.upgradeStake(validator, index, rewardAmount, proofs[0]);
    }

    function _unstake() internal {
        vm.prank(user);
        merkleDistributor.unstake(index, rewardAmount, proofs[0]);
    }
}
