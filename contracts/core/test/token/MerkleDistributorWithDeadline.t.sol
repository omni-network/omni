// SPDX-License-Identifier: GPL-3.0-only
pragma solidity 0.8.24;

import { Test } from "forge-std/Test.sol";
import { CompleteMerkle } from "murky/src/CompleteMerkle.sol";

import { Omni } from "src/token/Omni.sol";
import { Create3 } from "src/deploy/Create3.sol";
import { MockPortal } from "test/utils/MockPortal.sol";
import { MockSolverNetInbox } from "test/utils/MockSolverNetInbox.sol";
import { TransparentUpgradeableProxy } from "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";

import { GenesisStake } from "src/token/GenesisStake.sol";
import { MerkleDistributor } from "src/libraries/MerkleDistributor.sol";
import { MerkleDistributorWithDeadline } from "src/token/MerkleDistributorWithDeadline.sol";

import { IERC7683, IOriginSettler } from "solve/erc7683/IOriginSettler.sol";
import { SolverNet } from "solve/lib/SolverNet.sol";

contract MerkleDistributorWithDeadline_Test is Test {
    CompleteMerkle m;

    Omni omni;
    Create3 create3;
    MockPortal omniPortal;
    MockSolverNetInbox inbox;

    GenesisStake genesisStake;
    MerkleDistributorWithDeadline merkleDistributor;

    address proxyAdmin = makeAddr("proxyAdmin");
    address outbox = makeAddr("outbox");

    uint256 endTime = block.timestamp + 30 days;
    uint256 initialSupply = 1_000_000 ether;
    uint256 addrCount = 32;

    bytes32 internal constant ORDER_DATA_TYPEHASH = keccak256(
        "OrderData(address owner,uint64 destChainId,Deposit deposit,Call[] calls,Expense[] expenses)Deposit(address token,uint96 amount)Call(address target,bytes4 selector,uint256 value,bytes params)Expense(address spender,address token,uint96 amount)"
    );

    address[] stakers = new address[](addrCount);
    uint256[] amounts = new uint256[](addrCount);
    bytes32[] leaves = new bytes32[](addrCount);
    bytes32[][] proofs = new bytes32[][](addrCount);
    bytes32 root;

    function setUp() public {
        _setupEnvironment();
        _prepMerkleTree();
        _deployContracts();
        _fundEverything();
        _stakeAll();
    }

    // Setup environment contracts
    function _setupEnvironment() internal {
        m = new CompleteMerkle();

        omni = new Omni(initialSupply, address(this));
        create3 = new Create3();
        omniPortal = new MockPortal();
        inbox = new MockSolverNetInbox(outbox);
    }

    // Setup randomly generated merkle tree
    function _prepMerkleTree() internal {
        // Prep leaves
        for (uint256 i; i < addrCount; ++i) {
            uint256 addrSeed = vm.randomUint();
            address addr = address(uint160(addrSeed));
            uint256 amount = vm.randomUint(1 ether, (initialSupply / 2) / addrCount);

            stakers[i] = addr;
            amounts[i] = amount;
            leaves[i] = keccak256(abi.encodePacked(i, addr, amount));
        }

        // Generate proofs
        for (uint256 i; i < addrCount; ++i) {
            proofs[i] = m.getProof(leaves, i);
        }

        // Generate root
        root = m.getRoot(leaves);

        // Verify all proofs
        for (uint256 i; i < addrCount; ++i) {
            assertTrue(m.verifyProof(root, proofs[i], leaves[i]));
        }
    }

    // Deploy contracts and verify contract addresses
    function _deployContracts() internal {
        // Precompute create3 addresses
        address genesisStakeAddr = create3.getDeployed(address(this), keccak256("genesisStake"));
        address merkleDistributorAddr = create3.getDeployed(address(this), keccak256("merkleDistributor"));

        // Deploy contracts
        address genesisStakeImpl = address(new GenesisStake(address(omni), merkleDistributorAddr));
        genesisStake = GenesisStake(
            create3.deploy(
                keccak256("genesisStake"),
                abi.encodePacked(
                    type(TransparentUpgradeableProxy).creationCode,
                    abi.encode(
                        genesisStakeImpl, proxyAdmin, abi.encodeCall(GenesisStake.initialize, (address(this), 30 days))
                    )
                )
            )
        );
        merkleDistributor = MerkleDistributorWithDeadline(
            create3.deploy(
                keccak256("merkleDistributor"),
                abi.encodePacked(
                    type(MerkleDistributorWithDeadline).creationCode,
                    abi.encode(address(omni), root, endTime, address(omniPortal), genesisStakeAddr, address(inbox))
                )
            )
        );

        // Verify precomputed addresses
        require(address(genesisStake) == genesisStakeAddr, "GenesisStake address mismatch");
        require(address(merkleDistributor) == merkleDistributorAddr, "MerkleDistributorWithDeadline address mismatch");
    }

    // Fund stakers and the distributor contract
    function _fundEverything() internal {
        for (uint256 i; i < addrCount; ++i) {
            omni.transfer(stakers[i], amounts[i]);
        }

        omni.transfer(address(merkleDistributor), initialSupply / 2);
    }

    // Prank all stakers and stake all of their tokens
    function _stakeAll() public {
        for (uint256 i; i < addrCount; ++i) {
            vm.startPrank(stakers[i]);
            omni.approve(address(genesisStake), type(uint256).max);
            genesisStake.stake(amounts[i]);
            vm.stopPrank();
        }
    }

    // Generate an ERC7683 order to check merkle distributor's call to inbox against
    function generateERC7683Order(address addr, uint256 amount)
        public
        view
        returns (IERC7683.OnchainCrossChainOrder memory)
    {
        SolverNet.Deposit memory deposit = SolverNet.Deposit({ token: address(omni), amount: uint96(amount * 2) });

        SolverNet.Call[] memory call = new SolverNet.Call[](1);
        call[0] = SolverNet.Call({ target: addr, selector: bytes4(0), value: amount * 2, params: "" });

        SolverNet.OrderData memory orderData = SolverNet.OrderData({
            owner: addr,
            destChainId: omniPortal.omniChainId(),
            deposit: deposit,
            calls: call,
            expenses: new SolverNet.Expense[](0)
        });

        return IERC7683.OnchainCrossChainOrder({
            fillDeadline: 0,
            orderDataType: ORDER_DATA_TYPEHASH,
            orderData: abi.encode(orderData)
        });
    }

    // Simple test to verify that our merkle tree is compatible with the distributor
    function test_claim() public {
        for (uint256 i; i < addrCount; ++i) {
            vm.prank(stakers[i]);
            merkleDistributor.claim(i, stakers[i], amounts[i], proofs[i]);
            assertEq(omni.balanceOf(stakers[i]), amounts[i]);
        }
    }

    function test_migrateToOmni_reverts() public {
        // Cannot migrate after claim window
        vm.warp(endTime + 1);
        vm.prank(stakers[0]);
        vm.expectRevert(MerkleDistributorWithDeadline.ClaimWindowFinished.selector);
        merkleDistributor.migrateToOmni(0, amounts[0], proofs[0]);

        // Cannot migrate if proof is invalid
        bytes32 proof = proofs[0][0];
        proofs[0][0] = bytes32(uint256(proof) + 1);
        vm.warp(endTime - 1);
        vm.prank(stakers[0]);
        vm.expectRevert(MerkleDistributor.InvalidProof.selector);
        merkleDistributor.migrateToOmni(0, amounts[0], proofs[0]);
        proofs[0][0] = proof;

        // Fully claim all stake and rewards
        vm.prank(stakers[0]);
        merkleDistributor.migrateToOmni(0, amounts[0], proofs[0]);

        // Cannot migrate if user has no stake to migrate
        vm.prank(stakers[0]);
        vm.expectRevert(MerkleDistributorWithDeadline.NothingToMigrate.selector);
        merkleDistributor.migrateToOmni(0, amounts[0], proofs[0]);
    }

    // Fully test migrateToOmni for all members of the merkle tree
    function test_migrateToOmni_succeeds() public {
        for (uint256 i; i < addrCount; ++i) {
            vm.startPrank(stakers[i]);
            uint256 inboxBalance = omni.balanceOf(address(inbox));

            // Get IERC7683 order and resolved orders
            IERC7683.OnchainCrossChainOrder memory order = generateERC7683Order(stakers[i], amounts[i]);
            IERC7683.ResolvedCrossChainOrder memory resolved = inbox.resolve(order);

            // Confirm merkleDistributor is calling the inbox with the order and that the resolved order is emitted
            vm.expectCall(address(inbox), abi.encodeCall(MockSolverNetInbox.open, (order)));
            vm.expectEmit(true, true, true, true);
            emit IERC7683.Open(resolved.orderId, resolved);
            merkleDistributor.migrateToOmni(i, amounts[i], proofs[i]);

            // Confirm the inbox balance has increased by the user's staked balance and claim reward
            assertEq(omni.balanceOf(address(inbox)), inboxBalance + amounts[i] * 2);
            assertEq(omni.balanceOf(stakers[i]), 0);

            vm.stopPrank();
        }
    }
}
