// SPDX-License-Identifier: GPL-3.0-only
/* solhint-disable no-console */
pragma solidity 0.8.24;

import { Script, console2 } from "forge-std/Script.sol";
import { CompleteMerkle } from "murky/src/CompleteMerkle.sol";
import { LibString } from "solady/src/utils/LibString.sol";

import { ICreateX } from "createx/src/ICreateX.sol";
import { IERC20 } from "@openzeppelin/contracts/interfaces/IERC20.sol";
import { IOmniPortal } from "src/interfaces/IOmniPortal.sol";
import { ISolverNetInbox } from "solve/src/interfaces/ISolverNetInbox.sol";

import { GenesisStakeV2 } from "src/token/GenesisStakeV2.sol";
import { DebugMerkleDistributorWithoutDeadline } from "./DebugMerkleDistributorWithoutDeadline.sol";
import { TransparentUpgradeableProxy } from "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";

contract OmegaGenesisStakeV2DebugScript is Script {
    CompleteMerkle internal m;

    ICreateX internal createX = ICreateX(0xba5Ed099633D3B313e4D5F7bdc1305d3c28ba5Ed);
    IERC20 internal omni = IERC20(0xD036C60f46FF51dd7Fbf6a819b5B171c8A076b07); // Holesky. Base Sepolia: 0xe4075366F03C286846dECE8AAF104cF2a542294d
    IOmniPortal internal portal = IOmniPortal(0xcB60A0451831E4865bC49f41F9C67665Fc9b75C3);
    ISolverNetInbox internal inbox = ISolverNetInbox(0x7EA0CeB70D5Df75a730E6cB3EADeC12EfdFe80a1);

    address internal validator = 0xdBd26a685DB4475b6c58ADEC0DE06c6eE387EAa8;

    GenesisStakeV2 internal genesisStake = GenesisStakeV2(0xE5728CB5EdDD90DFcE8318783642b399b7d7feE7); // owner
    DebugMerkleDistributorWithoutDeadline internal merkleDistributor =
        DebugMerkleDistributorWithoutDeadline(0x983285184A687A137E5eF3ADAf0B4cC84C8d3a90);

    uint256 internal rewardAmount = 50 ether;
    bytes32[] internal leaves = new bytes32[](64);
    bytes32[][] internal proofs = new bytes32[][](64);
    bytes32 internal root;

    function deploy() public {
        vm.startBroadcast();

        root = hex"a206a8bc03504b83f44bd1cfbc709c454346ca3a2485e29edf9627b61a5c7b89";
        _deployContracts();

        omni.transfer(address(merkleDistributor), 50 ether);

        vm.stopBroadcast();
    }

    function _deployContracts() internal {
        bytes32 genesisStakeSalt = keccak256(abi.encodePacked("genesisStake", block.timestamp));
        bytes32 merkleDistributorSalt = keccak256(abi.encodePacked("merkleDistributor", block.timestamp));

        address genesisStakeAddr = createX.computeCreate3Address(keccak256(abi.encodePacked(genesisStakeSalt)));
        address merkleDistributorAddr =
            createX.computeCreate3Address(keccak256(abi.encodePacked(merkleDistributorSalt)));

        address genesisStakeImpl = address(new GenesisStakeV2(address(omni), merkleDistributorAddr));

        // Deploy proxy without initialization
        address proxyAddress = createX.deployCreate3(
            genesisStakeSalt,
            abi.encodePacked(
                type(TransparentUpgradeableProxy).creationCode,
                abi.encode(
                    genesisStakeImpl,
                    msg.sender, // admin
                    new bytes(0) // empty initialization data
                )
            )
        );

        genesisStake = GenesisStakeV2(proxyAddress);

        // Initialize separately after deployment
        genesisStake.initialize(msg.sender);

        address merkleDistributorImpl = address(
            new DebugMerkleDistributorWithoutDeadline(
                address(omni), root, address(portal), address(genesisStake), address(inbox)
            )
        );

        merkleDistributor = DebugMerkleDistributorWithoutDeadline(
            createX.deployCreate3(
                merkleDistributorSalt,
                abi.encodePacked(
                    type(TransparentUpgradeableProxy).creationCode,
                    abi.encode(
                        merkleDistributorImpl,
                        msg.sender,
                        abi.encodeCall(DebugMerkleDistributorWithoutDeadline.initialize, (msg.sender))
                    )
                )
            )
        );

        require(address(genesisStake) == genesisStakeAddr, "GenesisStakeV2 addr mismatch");
        require(address(merkleDistributor) == merkleDistributorAddr, "DebugMerkleDistributor addr mismatch");

        console2.log("GenesisStakeV2 implementation:", address(genesisStakeImpl));
        console2.log("GenesisStakeV2 implementation constructor args:");
        console2.logBytes(abi.encode(address(omni), merkleDistributorAddr));
        console2.log("GenesisStakeV2 proxy address:", address(genesisStake));
        console2.log("GenesisStakeV2 proxy constructor args:");
        console2.logBytes(new bytes(0));
        console2.log("");
        console2.log("DebugMerkleDistributor implementation:", address(merkleDistributorImpl));
        console2.log("DebugMerkleDistributor implementation constructor args:");
        console2.logBytes(abi.encode(address(omni), root));
        console2.log("DebugMerkleDistributor proxy address:", address(merkleDistributor));
        console2.log("DebugMerkleDistributor proxy constructor args:");
        console2.logBytes(
            abi.encode(
                merkleDistributorImpl,
                msg.sender,
                abi.encodeCall(DebugMerkleDistributorWithoutDeadline.initialize, (msg.sender))
            )
        );
    }
}
