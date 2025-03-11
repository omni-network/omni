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
import { DebugMerkleDistributorWithDeadline } from "./DebugMerkleDistributorWithDeadline.sol";
import { TransparentUpgradeableProxy } from "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";

contract OmegaGenesisStakeQAScript is Script {
    CompleteMerkle internal m;

    ICreateX internal createX = ICreateX(0xba5Ed099633D3B313e4D5F7bdc1305d3c28ba5Ed);
    IERC20 internal omni = IERC20(0xD036C60f46FF51dd7Fbf6a819b5B171c8A076b07); // Holesky
    IOmniPortal internal portal = IOmniPortal(0xcB60A0451831E4865bC49f41F9C67665Fc9b75C3);
    ISolverNetInbox internal inbox = ISolverNetInbox(0x7EA0CeB70D5Df75a730E6cB3EADeC12EfdFe80a1);

    address internal validator = 0xdBd26a685DB4475b6c58ADEC0DE06c6eE387EAa8;

    GenesisStake internal genesisStake = GenesisStake(0x2A20996AaaE4456707E6f6e8187d07459b67A93c);
    DebugMerkleDistributorWithDeadline internal merkleDistributor =
        DebugMerkleDistributorWithDeadline(0xBD9bE69985ca96FD601DC89F04Ed3b7bf52Fe038);

    uint256 internal endTime = block.timestamp + 30 days;
    uint256 internal index = 7956;
    uint256 internal depositAmount = 100 ether;
    uint256 internal rewardAmount = 1_305_710_306_500_000_000;
    bytes32[] internal leaves = new bytes32[](1);
    bytes32[][] internal proofs = new bytes32[][](1);
    bytes32 internal root;

    function deploy() public {
        vm.startBroadcast();

        _prepMerkleTree();
        _deployContracts();

        omni.transfer(address(merkleDistributor), 10_000 ether);

        vm.stopBroadcast();
    }

    /**
     * @dev This assumes the four relevant addresses above have been set and that a new GenesisStake contract should be
     * deployed. It also assumes that the broadcaster has 200 OMNI ERC20 tokens to spend on the network.
     */
    function deployAndTest() public {
        vm.startBroadcast();

        _prepMerkleTree();
        _deployContracts();
        _approveStakeAndFund();

        // Change index values according to deployer/caller address in merkle tree
        merkleDistributor.migrateToOmni(validator, index, rewardAmount, proofs[0]);

        vm.stopBroadcast();
    }

    function migrate() public {
        vm.startBroadcast();

        _prepMerkleTree();
        _approveStakeAndFund();

        // Change index values according to deployer/caller address in merkle tree
        merkleDistributor.migrateToOmni(validator, index, rewardAmount, proofs[0]);

        vm.stopBroadcast();
    }

    function _prepMerkleTree() internal {
        m = new CompleteMerkle();

        leaves[0] = keccak256(abi.encodePacked(uint256(7956), 0xA779fC675Db318dab004Ab8D538CB320D0013F42, rewardAmount));
        proofs[0].push(hex"00e7dad07181b59b1f038181844e40254aa180d3b04b7e473faf4c960a99eded");
        proofs[0].push(hex"463978574b00aa05219cd93135c1d787a4565ba1a5001b7063e48419e9d6719d");
        proofs[0].push(hex"fb0c0e87305db129fb1acb5aa5e75e055c584c3d7109754b343438f0256a7c96");
        proofs[0].push(hex"893f87ead4b105b5af0389dfbc354e06c38026ece5c94411de9d78e5f928e294");
        proofs[0].push(hex"cbf387b1cd966ed9aa8549af125c0039b96062a13e3f7053d245635c19d004fb");
        proofs[0].push(hex"12a296264c3483bf23456a745feb1f0afbb423b50ed04800b3f4bf3203f58104");
        proofs[0].push(hex"4d5f5a35000501ba7c713fbe0663ef0325fdeaca3de2ff79c16f1193aa2d3688");
        proofs[0].push(hex"00599b87e2278d1778c5872df387af6530dbfcc7d9c56601a8af14452d6a7c52");
        proofs[0].push(hex"3dc163521a23ea985c1908f9fef2267cd3be23452f92218f8d72362e9efa6159");
        proofs[0].push(hex"866bf35be46662c6de366bcdd17c89f6fa4167528f599bcd9dc6fed4a1216697");
        proofs[0].push(hex"108e8afc22ef50935d515b90522e4155a96d762d018fb11275bde0719608fb08");
        proofs[0].push(hex"cf736f928c4819ad0219d7c32179e6cddc474132d775792b206b804e42ababba");
        proofs[0].push(hex"ebf7bdce15089351b96d633be463d1e6d327624db7ae7e8c94d3fe27a7cb34ba");
        proofs[0].push(hex"0c17dbcd20b4f15264b9b382cc1cc57b05c2e0dd309a23d7f08e35743a3f5dde");

        root = hex"2fe059bec6dd8491e5aa41329711a14bce8108f8b00aa4b9cb8795579836c5c2";

        // Verify proof
        require(m.verifyProof(root, proofs[0], leaves[0]), "Proof is invalid");
    }

    function _deployContracts() internal {
        bytes32 genesisStakeSalt = keccak256(abi.encodePacked("genesisStake", block.timestamp));
        bytes32 merkleDistributorSalt = keccak256(abi.encodePacked("merkleDistributor", block.timestamp));

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
        merkleDistributor = DebugMerkleDistributorWithDeadline(
            createX.deployCreate3(
                merkleDistributorSalt,
                abi.encodePacked(
                    type(DebugMerkleDistributorWithDeadline).creationCode,
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

    function _approveStakeAndFund() internal {
        if (omni.allowance(msg.sender, address(genesisStake)) < depositAmount) {
            omni.approve(address(genesisStake), type(uint256).max);
        }
        genesisStake.stake(depositAmount);
        omni.transfer(address(merkleDistributor), rewardAmount);
    }
}
