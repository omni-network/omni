// SPDX-License-Identifier: GPL-3.0-only
/* solhint-disable no-console */
pragma solidity 0.8.24;

import { Script, console2 } from "forge-std/Script.sol";
import { CompleteMerkle } from "murky/src/CompleteMerkle.sol";
import { LibString } from "solady/src/utils/LibString.sol";
import { JSONParserLib } from "solady/src/utils/JSONParserLib.sol";

import { ICreateX } from "createx/src/ICreateX.sol";
import { IERC20 } from "@openzeppelin/contracts/interfaces/IERC20.sol";
import { IOmniPortal } from "src/interfaces/IOmniPortal.sol";
import { ISolverNetInbox } from "solve/src/interfaces/ISolverNetInbox.sol";

import { GenesisStake } from "src/token/GenesisStake.sol";
import { DebugMerkleDistributorWithDeadline } from "./DebugMerkleDistributorWithDeadline.sol";
import { TransparentUpgradeableProxy } from "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";

contract StagingGenesisStakeScript is Script {
    CompleteMerkle internal m;

    ICreateX internal createX = ICreateX(0xba5Ed099633D3B313e4D5F7bdc1305d3c28ba5Ed);
    IERC20 internal omni;
    IOmniPortal internal portal;
    ISolverNetInbox internal inbox;

    address internal validator = 0xD6CD71dF91a6886f69761826A9C4D123178A8d9D;

    GenesisStake internal genesisStake;
    DebugMerkleDistributorWithDeadline internal merkleDistributor;

    uint256 internal endTime = block.timestamp + 30 days;
    uint256 internal depositAmount = 100 ether;
    uint256 internal rewardAmount = 5 ether;
    bytes32[] internal leaves = new bytes32[](64);
    bytes32[][] internal proofs = new bytes32[][](64);
    bytes32 internal root;

    function setUp() public {
        string memory stagingAddrsJson = _getStagingAddresses();
        _setStagingAddresses(stagingAddrsJson);
    }

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
        merkleDistributor.upgradeStake(validator, 0, rewardAmount, proofs[0]);

        vm.stopBroadcast();
    }

    function merkleTree() public {
        _prepMerkleTree();

        console2.log("Merkle root:");
        console2.logBytes32(root);
        console2.log("Merkle proofs:");
        console2.logBytes(abi.encode(proofs));
    }

    function merkleTest() public {
        m = new CompleteMerkle();

        leaves[0] = keccak256(
            abi.encodePacked(
                uint256(7956), 0xA779fC675Db318dab004Ab8D538CB320D0013F42, uint256(1_305_710_306_500_000_000)
            )
        );
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

        m.verifyProof(root, proofs[0], leaves[0]);
    }

    function _prepMerkleTree() internal {
        m = new CompleteMerkle();

        leaves[0] = keccak256(abi.encodePacked(uint256(0), 0xA779fC675Db318dab004Ab8D538CB320D0013F42, rewardAmount));
        leaves[1] = keccak256(abi.encodePacked(uint256(1), 0x2e6d9f2CA3E25b4216a1430208046f965bdc1f75, rewardAmount));
        leaves[2] = keccak256(abi.encodePacked(uint256(2), 0xF8d882Bc33D4a2b327E7C4d3cA7AaA325b5591Ea, rewardAmount));
        leaves[3] = keccak256(abi.encodePacked(uint256(3), 0xF52015A4777aE31e0540441b345a9dC17111a956, rewardAmount));
        leaves[4] = keccak256(abi.encodePacked(uint256(4), 0x66196e04eAaCE82A1bCFF133296b5A8e08C34292, rewardAmount));
        leaves[5] = keccak256(abi.encodePacked(uint256(5), 0x15f8F98625d5Ac3D52A964Fdb9dd200253A1c618, rewardAmount));
        leaves[6] = keccak256(abi.encodePacked(uint256(6), 0xA1080Ad10671c39A3d66E781bEcAD38006b66C2f, rewardAmount));
        leaves[7] = keccak256(abi.encodePacked(uint256(7), 0x3253Fe68a22ea73e6b4E4FfdE6552fa72D0773Ef, rewardAmount));
        leaves[8] = keccak256(abi.encodePacked(uint256(8), 0xBE91C9Ca683897DB87651283529CB0D0c0D8c030, rewardAmount));
        leaves[9] = keccak256(abi.encodePacked(uint256(9), 0xbD438e226c31dE1D15E51C2976CA9d0e5A01c3A9, rewardAmount));
        leaves[10] = keccak256(abi.encodePacked(uint256(10), 0x50E9E406234A724Dad241Fb8348C5e1e33D7804f, rewardAmount));
        leaves[11] = keccak256(abi.encodePacked(uint256(11), 0xafb26D25A86a80A78Fd610dF587983a0c662d412, rewardAmount));
        leaves[12] = keccak256(abi.encodePacked(uint256(12), 0xFDe6B312737c624123C857D0e6a8b030B5AC4701, rewardAmount));
        leaves[13] = keccak256(abi.encodePacked(uint256(13), 0x755f3b851a538303fd0dD55c1Bd07280d6617558, rewardAmount));
        leaves[14] = keccak256(abi.encodePacked(uint256(14), 0x41FdB83422432e2f31542327242D5eA6445Ca0E1, rewardAmount));
        leaves[15] = keccak256(abi.encodePacked(uint256(15), 0x60E627F7Ca6B03A0AEC2193956242dD690834E9a, rewardAmount));
        leaves[16] = keccak256(abi.encodePacked(uint256(16), 0xA84f84c377D3fE8541A1C742108865Ae885cC6dE, rewardAmount));
        leaves[17] = keccak256(abi.encodePacked(uint256(17), 0x0F388963c4d6AfD7f221AEacACCC6A170B8b03E4, rewardAmount));
        leaves[18] = keccak256(abi.encodePacked(uint256(18), 0x86E28B1125Fe97033A9A5A88e0987e7a9cfBd64F, rewardAmount));
        leaves[19] = keccak256(abi.encodePacked(uint256(19), 0x3e16929c908Aae64793792fA56E65C538D406C83, rewardAmount));
        leaves[20] = keccak256(abi.encodePacked(uint256(20), 0x602eD22b7A9c08979a61c488E8c2A2630502CF85, rewardAmount));
        leaves[21] = keccak256(abi.encodePacked(uint256(21), 0x4D1f7e1a6390fe2D33afC3895f13f1223a7C4b07, rewardAmount));
        leaves[22] = keccak256(abi.encodePacked(uint256(22), 0x789f42F61d2773519b6B5efbD69Ffe0f18d3D58F, rewardAmount));
        leaves[23] = keccak256(abi.encodePacked(uint256(23), 0x883D8acb3435e283089D28a54EC3116A787b39B7, rewardAmount));
        leaves[24] = keccak256(abi.encodePacked(uint256(24), 0x469c91d58e70174cF5205AF43C49dC34273F7FcC, rewardAmount));
        leaves[25] = keccak256(abi.encodePacked(uint256(25), 0x2b23e3cAE1187422ffD6d737AA51dD2321d27fe6, rewardAmount));
        leaves[26] = keccak256(abi.encodePacked(uint256(26), 0xAC1e04ecaCf5622ef9F67ceD6c7d675d2e074Fd4, rewardAmount));
        leaves[27] = keccak256(abi.encodePacked(uint256(27), 0xe352D7e2EB502fc691b20B90cb1667Ac81Dae31D, rewardAmount));
        leaves[28] = keccak256(abi.encodePacked(uint256(28), 0x69D18C72a35Fbee774ec0DF9A96786A644393946, rewardAmount));
        leaves[29] = keccak256(abi.encodePacked(uint256(29), 0x9FE50c320b5f65473a8D595ECD0401a32113b527, rewardAmount));
        leaves[30] = keccak256(abi.encodePacked(uint256(30), 0x14C97Ca130Cf8cB5b5AED9d6EFC61AdE2d41479C, rewardAmount));
        leaves[31] = keccak256(abi.encodePacked(uint256(31), 0xc59cB2C577e5910B09F22542A895bE3F38a094E3, rewardAmount));
        leaves[32] = keccak256(abi.encodePacked(uint256(32), 0x4D0165Cbd5f4a3135eAbd8FB0a3E86C2D24E235a, rewardAmount));
        leaves[33] = keccak256(abi.encodePacked(uint256(33), 0x2503961F86534Ae0b6b348deb0bE56e920f0a2AA, rewardAmount));
        leaves[34] = keccak256(abi.encodePacked(uint256(34), 0xAECfbced9e27662f6f29d77fB21cbe55bf79beaF, rewardAmount));
        leaves[35] = keccak256(abi.encodePacked(uint256(35), 0xF6a1Ed7fA2DB22c801ffd17d2e861CBdDF1304a7, rewardAmount));
        leaves[36] = keccak256(abi.encodePacked(uint256(36), 0x74a5CE18b168dc0496870967F5D72640c1Cb233a, rewardAmount));
        leaves[37] = keccak256(abi.encodePacked(uint256(37), 0x78fC8d86d9a969c9a6E481BCc3980E4b11777B69, rewardAmount));
        leaves[38] = keccak256(abi.encodePacked(uint256(38), 0xe431ca9d69d14e7306948349D99CA85Fb3D2E989, rewardAmount));
        leaves[39] = keccak256(abi.encodePacked(uint256(39), 0x706CC438Ee2FE01860E80BD6754d7D0f6De49ca9, rewardAmount));
        leaves[40] = keccak256(abi.encodePacked(uint256(40), 0x32Aa390F5ba2dc11dAf4e00F828C26e9bb03b1C2, rewardAmount));
        leaves[41] = keccak256(abi.encodePacked(uint256(41), 0x8883644304aA2938B1C5E4464469bDE034A80AAa, rewardAmount));
        leaves[42] = keccak256(abi.encodePacked(uint256(42), 0x50C5559d0d3a8F11Fc4454C074E64fc0de289512, rewardAmount));
        leaves[43] = keccak256(abi.encodePacked(uint256(43), 0x477F5e11157CBF8A6E41D07241b8e4914bA850fA, rewardAmount));
        leaves[44] = keccak256(abi.encodePacked(uint256(44), 0x5cE4376Ed699E65e31c0FcbaDCf98e20eA2Fa20a, rewardAmount));
        leaves[45] = keccak256(abi.encodePacked(uint256(45), 0xC2Bb30cB65c4F1181ff1393415C13Baa7F1C8fB8, rewardAmount));
        leaves[46] = keccak256(abi.encodePacked(uint256(46), 0xC1278555d08C99f55Cbe19678BfCc61fA5DFf6eF, rewardAmount));
        leaves[47] = keccak256(abi.encodePacked(uint256(47), 0x64f2059fC811706776fB6A5f4cd62215b89f2F5C, rewardAmount));
        leaves[48] = keccak256(abi.encodePacked(uint256(48), 0x0AC94Fa7C75a48a3bbCDd1210882FcE6391fCE58, rewardAmount));
        leaves[49] = keccak256(abi.encodePacked(uint256(49), 0x70bD9E2297edC6b4712ad51E37454504b936BBe3, rewardAmount));
        leaves[50] = keccak256(abi.encodePacked(uint256(50), 0xfaB5c13854603F439bEc15aB4FC756D81ed058C2, rewardAmount));
        leaves[51] = keccak256(abi.encodePacked(uint256(51), 0x2D61bE547b365BD5CdCc02920818492Fb7bdb765, rewardAmount));
        leaves[52] = keccak256(abi.encodePacked(uint256(52), 0x3ddB180d96C98e77A8F20aC456B6764B4D478A8a, rewardAmount));
        leaves[53] = keccak256(abi.encodePacked(uint256(53), 0x62398788692aDed44638F8b9F3eE4B977C78ff46, rewardAmount));
        leaves[54] = keccak256(abi.encodePacked(uint256(54), 0x38E2a3FC1923767F74d2308a529a353e91763EBF, rewardAmount));
        leaves[55] = keccak256(abi.encodePacked(uint256(55), 0xe3481474b23f88a8917DbcB4cBC55Efcf0f68CC7, rewardAmount));
        leaves[56] = keccak256(abi.encodePacked(uint256(56), 0xD9c0BB3476CE2AD2102D3AC07287BB802EeA98f1, rewardAmount));
        leaves[57] = keccak256(abi.encodePacked(uint256(57), 0xDEdDf2DA39E0E39469a28F5A0392DcFbe40323de, rewardAmount));
        leaves[58] = keccak256(abi.encodePacked(uint256(58), 0x9474d842BaCa1fe809810dF4fe4D194Dae83f9d6, rewardAmount));
        leaves[59] = keccak256(abi.encodePacked(uint256(59), 0xf41c4c528E06020Ccc1FC738398f26e7334854b3, rewardAmount));
        leaves[60] = keccak256(abi.encodePacked(uint256(60), 0xA6C9c842dc0C9C16338444e8bB77b885986Ef38b, rewardAmount));
        leaves[61] = keccak256(abi.encodePacked(uint256(61), 0xc83629D6A24851b7B90A2fa7f63a762dFE1021BC, rewardAmount));
        leaves[62] = keccak256(abi.encodePacked(uint256(62), 0xF6CDB1E733EA00D0eEa1A32F218B0ec76ABF1517, rewardAmount));
        leaves[63] = keccak256(abi.encodePacked(uint256(63), 0xBeD17aa3E1c99ea86e19e7B38356C54007BB6CDe, rewardAmount));

        // Generate the Merkle root
        root = m.getRoot(leaves);

        // Generate proofs for each leaf
        for (uint256 i; i < leaves.length; ++i) {
            proofs[i] = m.getProof(leaves, i);
        }

        // Verify all proofs
        for (uint256 i; i < leaves.length; ++i) {
            require(
                m.verifyProof(root, proofs[i], leaves[i]),
                string(abi.encodePacked("Proof ", LibString.toString(i), " is invalid"))
            );
        }
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

    function _getStagingAddresses() internal returns (string memory) {
        string[] memory inputs = new string[](3);
        inputs[0] = "go";
        inputs[1] = "run";
        inputs[2] = "../../scripts/stagingaddrs/main.go";

        bytes memory stagingAddrsJson = vm.ffi(inputs);
        return string(stagingAddrsJson);
    }

    function _setStagingAddresses(string memory stagingAddrsJson) internal {
        JSONParserLib.Item memory object = JSONParserLib.parse(stagingAddrsJson);
        /* solhint-disable quotes */
        JSONParserLib.Item memory omniItem = JSONParserLib.at(object, '"token"');
        JSONParserLib.Item memory portalItem = JSONParserLib.at(object, '"portal"');
        JSONParserLib.Item memory inboxItem = JSONParserLib.at(object, '"solvernetinbox"');
        /* solhint-enable quotes */

        string memory omniAddr = JSONParserLib.value(omniItem);
        omniAddr = LibString.slice(omniAddr, 1, bytes(omniAddr).length - 1);
        omni = IERC20(vm.parseAddress(omniAddr));

        string memory portalAddr = JSONParserLib.value(portalItem);
        portalAddr = LibString.slice(portalAddr, 1, bytes(portalAddr).length - 1);
        portal = IOmniPortal(vm.parseAddress(portalAddr));

        string memory inboxAddr = JSONParserLib.value(inboxItem);
        inboxAddr = LibString.slice(inboxAddr, 1, bytes(inboxAddr).length - 1);
        inbox = ISolverNetInbox(vm.parseAddress(inboxAddr));
    }
}
