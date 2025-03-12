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

contract OmegaGenesisStakeQAScript is Script {
    CompleteMerkle internal m;

    ICreateX internal createX = ICreateX(0xba5Ed099633D3B313e4D5F7bdc1305d3c28ba5Ed);
    IERC20 internal omni = IERC20(0xD036C60f46FF51dd7Fbf6a819b5B171c8A076b07); // Holesky
    IOmniPortal internal portal = IOmniPortal(0xcB60A0451831E4865bC49f41F9C67665Fc9b75C3);
    ISolverNetInbox internal inbox = ISolverNetInbox(0x7EA0CeB70D5Df75a730E6cB3EADeC12EfdFe80a1);

    address internal validator = 0xdBd26a685DB4475b6c58ADEC0DE06c6eE387EAa8;

    GenesisStake internal genesisStake = GenesisStake(0xf4A841bcaD26a4c4fD36B30369252f6A49Bf6E4d);
    MerkleDistributorWithDeadline internal merkleDistributor =
        MerkleDistributorWithDeadline(0xCD1111A338B5E797AaC30a370ba49EA611bb1D43);

    uint256 internal endTime = block.timestamp + 30 days;
    uint256 internal index = 7987;
    address internal claimer = 0xA779fC675Db318dab004Ab8D538CB320D0013F42;
    uint256 internal rewardAmount = 843_524_630_884_109_952;
    uint256 internal depositAmount = 100 ether;
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

        leaves[0] = keccak256(abi.encodePacked(index, claimer, rewardAmount));
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

    function _approveStakeAndFund() internal {
        if (omni.allowance(msg.sender, address(genesisStake)) < depositAmount) {
            omni.approve(address(genesisStake), type(uint256).max);
        }
        genesisStake.stake(depositAmount);
        omni.transfer(address(merkleDistributor), rewardAmount);
    }
}
