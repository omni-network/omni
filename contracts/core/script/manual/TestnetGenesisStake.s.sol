// SPDX-License-Identifier: GPL-3.0-only
/* solhint-disable no-console */
pragma solidity 0.8.24;

import { Script, console2 } from "forge-std/Script.sol";
import { CompleteMerkle } from "murky/src/CompleteMerkle.sol";

import { ICreateX } from "createx/src/ICreateX.sol";
import { IERC20 } from "@openzeppelin/contracts/interfaces/IERC20.sol";
import { IOmniPortal } from "src/interfaces/IOmniPortal.sol";
import { ISolverNetInbox } from "solve/src/interfaces/ISolverNetInbox.sol";

import { GenesisStake } from "src/token/GenesisStake.sol";
import { MerkleDistributorWithDeadline } from "src/token/MerkleDistributorWithDeadline.sol";
import { TransparentUpgradeableProxy } from "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";

contract TestnetGenesisStakeScript is Script {
    CompleteMerkle internal m;

    ICreateX internal createX = ICreateX(0xba5Ed099633D3B313e4D5F7bdc1305d3c28ba5Ed);
    // IERC20 internal omni = IERC20(0xD036C60f46FF51dd7Fbf6a819b5B171c8A076b07); // Holesky Omni
    IERC20 internal omni = IERC20(0x8FbF29dE613Fb25F33587083b89A38baF150ca32); // DebugMockOmni
    IOmniPortal internal portal = IOmniPortal(0xcB60A0451831E4865bC49f41F9C67665Fc9b75C3);
    ISolverNetInbox internal inbox = ISolverNetInbox(0x80b6Ed465241a17080DC4A68bE42e80FEa1214DD);

    bytes32 internal genesisStakeSalt = 0xa779fc675db318dab004ab8d538cb320d0013f42006fda006bab5cd1034643cc;
    bytes32 internal merkleDistributorSalt = 0xa779fc675db318dab004ab8d538cb320d0013f4200205ad9f17a619b0079c949;

    address internal expectedGenesisStakeAddr = 0x00000000000063B7931226e67CeF52d86085154d;
    address internal expectedMerkleDistributorAddr = 0x0000000000003E960a909Ef4F0E77699a4286711;

    GenesisStake internal genesisStake;
    MerkleDistributorWithDeadline internal merkleDistributor;

    uint256 internal endTime = block.timestamp + 30 days;
    uint256 internal depositAmount = 100 ether;
    uint256 internal rewardAmount = 100 ether;
    bytes32[] internal leaves = new bytes32[](5);
    bytes32[][] internal proofs = new bytes32[][](5);
    bytes32 internal root;

    function deployTestnet() public {
        vm.startBroadcast();

        _prepMerkleTree();
        _deployDeterministicContracts();

        omni.transfer(address(merkleDistributor), 10_000 ether);

        vm.stopBroadcast();
    }

    /**
     * @dev This assumes the four relevant addresses above have been set and that a new GenesisStake contract should be
     * deployed. It also assumes that the broadcaster has 200 OMNI ERC20 tokens to spend on the network.
     */
    function freshDeployAndTest() public {
        vm.startBroadcast();

        _prepMerkleTree();
        _deployFreshContracts();
        _approveStakeAndFund();

        merkleDistributor.migrateToOmni(0, rewardAmount, proofs[0]);

        vm.stopBroadcast();
    }

    function _prepMerkleTree() internal {
        m = new CompleteMerkle();

        leaves[0] = keccak256(abi.encodePacked(uint256(0), msg.sender, rewardAmount));
        leaves[1] = keccak256(abi.encodePacked(uint256(1), 0xF6CDB1E733EA00D0eEa1A32F218B0ec76ABF1517, rewardAmount));
        leaves[2] = keccak256(abi.encodePacked(uint256(2), 0xBeD17aa3E1c99ea86e19e7B38356C54007BB6CDe, rewardAmount));
        leaves[3] = keccak256(abi.encodePacked(uint256(3), 0x2D61bE547b365BD5CdCc02920818492Fb7bdb765, rewardAmount));
        leaves[4] = keccak256(abi.encodePacked(uint256(4), 0xA6C9c842dc0C9C16338444e8bB77b885986Ef38b, rewardAmount));

        proofs[0] = m.getProof(leaves, 0);
        proofs[1] = m.getProof(leaves, 1);
        proofs[2] = m.getProof(leaves, 2);
        proofs[3] = m.getProof(leaves, 3);
        proofs[4] = m.getProof(leaves, 4);

        root = m.getRoot(leaves);

        require(m.verifyProof(root, proofs[0], leaves[0]), "Proof 0 is invalid");
        require(m.verifyProof(root, proofs[1], leaves[1]), "Proof 1 is invalid");
        require(m.verifyProof(root, proofs[2], leaves[2]), "Proof 2 is invalid");
        require(m.verifyProof(root, proofs[3], leaves[3]), "Proof 3 is invalid");
        require(m.verifyProof(root, proofs[4], leaves[4]), "Proof 4 is invalid");
    }

    function _deployFreshContracts() internal {
        genesisStakeSalt = keccak256(abi.encodePacked("genesisStake", block.timestamp));
        merkleDistributorSalt = keccak256(abi.encodePacked("merkleDistributor", block.timestamp));

        address genesisStakeAddr = createX.computeCreate3Address(keccak256(abi.encodePacked(genesisStakeSalt)));
        address merkleDistributorAddr =
            createX.computeCreate3Address(keccak256(abi.encodePacked(merkleDistributorSalt)));

        address genesisStakeImpl = address(new GenesisStake(address(omni), merkleDistributorAddr));
        genesisStake = GenesisStake(
            createX.deployCreate3(
                genesisStakeSalt,
                abi.encodePacked(
                    type(TransparentUpgradeableProxy).creationCode,
                    abi.encode(
                        genesisStakeImpl, msg.sender, abi.encodeCall(GenesisStake.initialize, (msg.sender, 30 days))
                    )
                )
            )
        );
        merkleDistributor = MerkleDistributorWithDeadline(
            createX.deployCreate3(
                merkleDistributorSalt,
                abi.encodePacked(
                    type(MerkleDistributorWithDeadline).creationCode,
                    abi.encode(address(omni), root, endTime, address(portal), genesisStakeAddr, address(inbox))
                )
            )
        );

        require(address(genesisStake) == genesisStakeAddr, "GenesisStake addr mismatch");
        require(address(merkleDistributor) == merkleDistributorAddr, "MerkleDistributor addr mismatch");

        console2.log("GenesisStake implementation:", address(genesisStakeImpl));
        console2.log("GenesisStake implementation constructor args:");
        console2.logBytes(abi.encode(address(omni), merkleDistributorAddr));
        console2.log("GenesisStake proxy address:", address(genesisStake));
        console2.log("GenesisStake proxy constructor args:");
        console2.logBytes(
            abi.encode(genesisStakeImpl, msg.sender, abi.encodeCall(GenesisStake.initialize, (msg.sender, 30 days)))
        );
        console2.log("");
        console2.log("MerkleDistributor address:", address(merkleDistributor));
        console2.log("MerkleDistributor constructor args:");
        console2.logBytes(abi.encode(address(omni), root, endTime, address(portal), genesisStakeAddr, address(inbox)));
    }

    function _deployDeterministicContracts() internal {
        address genesisStakeImpl = address(new GenesisStake(address(omni), expectedMerkleDistributorAddr));
        genesisStake = GenesisStake(
            createX.deployCreate3(
                genesisStakeSalt,
                abi.encodePacked(
                    type(TransparentUpgradeableProxy).creationCode,
                    abi.encode(
                        genesisStakeImpl, msg.sender, abi.encodeCall(GenesisStake.initialize, (msg.sender, 30 days))
                    )
                )
            )
        );
        merkleDistributor = MerkleDistributorWithDeadline(
            createX.deployCreate3(
                merkleDistributorSalt,
                abi.encodePacked(
                    type(MerkleDistributorWithDeadline).creationCode,
                    abi.encode(address(omni), root, endTime, address(portal), expectedGenesisStakeAddr, address(inbox))
                )
            )
        );

        require(address(genesisStake) == expectedGenesisStakeAddr, "GenesisStake addr mismatch");
        require(address(merkleDistributor) == expectedMerkleDistributorAddr, "MerkleDistributor addr mismatch");

        console2.log("GenesisStake implementation:", address(genesisStakeImpl));
        console2.log("GenesisStake implementation constructor args:");
        console2.logBytes(abi.encode(address(omni), expectedMerkleDistributorAddr));
        console2.log("GenesisStake proxy address:", address(genesisStake));
        console2.log("GenesisStake proxy constructor args:");
        console2.logBytes(
            abi.encode(genesisStakeImpl, msg.sender, abi.encodeCall(GenesisStake.initialize, (msg.sender, 30 days)))
        );
        console2.log("");
        console2.log("MerkleDistributor address:", address(merkleDistributor));
        console2.log("MerkleDistributor constructor args:");
        console2.logBytes(
            abi.encode(address(omni), root, endTime, address(portal), expectedGenesisStakeAddr, address(inbox))
        );
    }

    function _approveStakeAndFund() internal {
        omni.approve(address(genesisStake), type(uint256).max);
        genesisStake.stake(depositAmount);
        omni.transfer(address(merkleDistributor), rewardAmount);
    }
}
