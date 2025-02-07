// SPDX-License-Identifier: GPL-3.0-only
pragma solidity 0.8.24;

import { Script } from "forge-std/Script.sol";
import { CompleteMerkle } from "murky/src/CompleteMerkle.sol";

import { Create3 } from "src/deploy/Create3.sol";
import { IERC20 } from "@openzeppelin/contracts/interfaces/IERC20.sol";
import { IOmniPortal } from "src/interfaces/IOmniPortal.sol";
import { ISolverNetInbox } from "solve/src/interfaces/ISolverNetInbox.sol";

import { GenesisStake } from "src/token/GenesisStake.sol";
import { MerkleDistributorWithDeadline } from "src/token/MerkleDistributorWithDeadline.sol";
import { TransparentUpgradeableProxy } from "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";

contract GenesisStakeScript is Script {
    CompleteMerkle internal m;

    Create3 internal create3 = Create3(0xd64EdA3D758944d62C4c94042DAf41b3a405A94d);
    IERC20 internal omni = IERC20(0xB50029Dc0DF4Db0193F25a8E41DEa207c13D09BB);
    IOmniPortal internal portal = IOmniPortal(0x67F639c83d7F6Ab75bf561f8cE27f23c80A3c5Bd);
    ISolverNetInbox internal inbox = ISolverNetInbox(0xf8C498bCB5fA0FAb64FdB2cDaD863d6A3b48FEC4);

    GenesisStake internal genesisStake;
    MerkleDistributorWithDeadline internal merkleDistributor;

    uint256 internal endTime = block.timestamp + 30 days;
    uint256 internal depositAmount = 80 ether;
    uint256 internal rewardAmount = 20 ether;
    bytes32[] internal leaves = new bytes32[](2);
    bytes32[][] internal proofs = new bytes32[][](2);
    bytes32 internal root;

    /**
     * @dev This assumes the four relevant addresses above have been set and that a new GenesisStake contract should be
     * deployed. It also assumes that the broadcaster has 100 OMNI ERC20 tokens to spend on the network.
     */
    function doAll() public {
        vm.startBroadcast();
        _prepMerkleTree();
        _deployContracts();
        _approveStakeAndFund();
        _migrateToOmni();
        vm.stopBroadcast();
    }

    function _prepMerkleTree() internal {
        m = new CompleteMerkle();

        leaves[0] = keccak256(abi.encodePacked(uint256(0), msg.sender, rewardAmount));
        leaves[1] = keccak256(abi.encodePacked(uint256(1), address(0xdead), uint256(1))); // Can't have one addr in a merkle tree
        proofs[0] = m.getProof(leaves, 0);
        proofs[1] = m.getProof(leaves, 1);
        root = m.getRoot(leaves);

        require(m.verifyProof(root, proofs[0], leaves[0]), "Proof 0 is invalid");
        require(m.verifyProof(root, proofs[1], leaves[1]), "Proof 1 is invalid");
    }

    function _deployContracts() internal {
        address genesisStakeAddr = create3.getDeployed(msg.sender, keccak256("genesisStake"));
        address merkleDistributorAddr = create3.getDeployed(msg.sender, keccak256("merkleDistributor"));

        address genesisStakeImpl = address(new GenesisStake(address(omni), merkleDistributorAddr));
        genesisStake = GenesisStake(
            create3.deploy(
                keccak256("genesisStake"),
                abi.encodePacked(
                    type(TransparentUpgradeableProxy).creationCode,
                    abi.encode(
                        genesisStakeImpl, msg.sender, abi.encodeCall(GenesisStake.initialize, (msg.sender, 30 days))
                    )
                )
            )
        );
        merkleDistributor = MerkleDistributorWithDeadline(
            create3.deploy(
                keccak256("merkleDistributor"),
                abi.encodePacked(
                    type(MerkleDistributorWithDeadline).creationCode,
                    abi.encode(address(omni), root, endTime, address(portal), genesisStakeAddr, address(inbox))
                )
            )
        );

        require(address(genesisStake) == genesisStakeAddr, "GenesisStake addr mismatch");
        require(address(merkleDistributor) == merkleDistributorAddr, "MerkleDistributor addr mismatch");
    }

    function _approveStakeAndFund() internal {
        omni.approve(address(genesisStake), type(uint256).max);
        genesisStake.stake(depositAmount);
        omni.transfer(address(merkleDistributor), rewardAmount);
    }

    function _migrateToOmni() internal {
        merkleDistributor.migrateToOmni(0, rewardAmount, proofs[0]);
    }
}
