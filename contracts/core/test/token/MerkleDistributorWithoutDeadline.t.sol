// SPDX-License-Identifier: GPL-3.0-only
pragma solidity 0.8.24;

import { Test } from "forge-std/Test.sol";
import { CompleteMerkle } from "murky/src/CompleteMerkle.sol";

import { Omni } from "src/token/Omni.sol";
import { Create3 } from "src/deploy/Create3.sol";
import { MockPortal } from "test/utils/MockPortal.sol";
import { MockSolverNetInbox } from "solve/test/utils/MockSolverNetInbox.sol";
import { TransparentUpgradeableProxy } from "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";

import { GenesisStakeV2 } from "src/token/GenesisStakeV2.sol";
import { MerkleDistributor } from "src/token/MerkleDistributor.sol";
import { MerkleDistributorWithoutDeadline } from "src/token/MerkleDistributorWithoutDeadline.sol";

import { IERC7683, IOriginSettler } from "solve/src/erc7683/IOriginSettler.sol";
import { IStaking } from "src/interfaces/IStaking.sol";
import { SolverNet } from "solve/src/lib/SolverNet.sol";

contract MerkleDistributorWithoutDeadline_Test is Test {
    CompleteMerkle m;

    Omni omni;
    Create3 create3;
    MockPortal omniPortal;
    MockSolverNetInbox inbox;

    GenesisStakeV2 genesisStake;
    MerkleDistributorWithoutDeadline merkleDistributor;

    address admin = makeAddr("admin");
    address proxyAdmin = makeAddr("proxyAdmin");
    address outbox = makeAddr("outbox");
    address validator = makeAddr("validator");

    uint256 initialSupply = 1_000_000 ether;
    uint256 addrCount = 32;

    bytes32 internal constant ORDER_DATA_TYPEHASH = keccak256(
        "OrderData(address owner,uint64 destChainId,Deposit deposit,Call[] calls,TokenExpense[] expenses)Deposit(address token,uint96 amount)Call(address target,bytes4 selector,uint256 value,bytes params)TokenExpense(address spender,address token,uint96 amount)"
    );

    address internal constant STAKING = 0xCCcCcC0000000000000000000000000000000001;

    uint256[] pks = new uint256[](addrCount);
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
    }

    // Setup environment contracts
    function _setupEnvironment() internal {
        m = new CompleteMerkle();

        omni = new Omni(initialSupply, address(this));
        create3 = new Create3();
        omniPortal = new MockPortal();
        inbox = new MockSolverNetInbox(address(omniPortal));
    }

    // Setup randomly generated merkle tree
    function _prepMerkleTree() internal {
        // Prep leaves
        for (uint256 i; i < addrCount; ++i) {
            pks[i] = vm.randomUint();
            stakers[i] = vm.addr(pks[i]);
            amounts[i] = vm.randomUint(1 ether, (initialSupply / 2) / addrCount);
            leaves[i] = keccak256(abi.encodePacked(i, stakers[i], amounts[i]));
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
        address genesisStakeImpl = address(new GenesisStakeV2(address(omni), merkleDistributorAddr));
        genesisStake = GenesisStakeV2(
            create3.deploy(
                keccak256("genesisStake"),
                abi.encodePacked(
                    type(TransparentUpgradeableProxy).creationCode,
                    abi.encode(
                        genesisStakeImpl, proxyAdmin, abi.encodeCall(GenesisStakeV2.initialize, (address(this)))
                    )
                )
            )
        );

        // Deploy MerkleDistributorWithoutDeadline implementation
        address merkleDistributorImpl = address(new MerkleDistributorWithoutDeadline(
            address(omni),
            root,
            address(omniPortal),
            address(genesisStake),
            address(inbox)
        ));
        merkleDistributor = MerkleDistributorWithoutDeadline(
            create3.deploy(
                keccak256("merkleDistributor"),
                abi.encodePacked(
                    type(TransparentUpgradeableProxy).creationCode,
                    abi.encode(
                        merkleDistributorImpl,
                        proxyAdmin,
                        abi.encodeCall(
                            MerkleDistributorWithoutDeadline.initialize,
                            (admin)
                        )
                    )
                )
            )
        );

        // Verify precomputed addresses
        require(address(genesisStake) == genesisStakeAddr, "GenesisStakeV2 address mismatch");
        require(address(merkleDistributor) == merkleDistributorAddr, "MerkleDistributorWithoutDeadline address mismatch");
    }

    // Fund stakers and the distributor contract
    function _fundEverything() internal {
        for (uint256 i; i < addrCount; ++i) {
            // Transfer tokens to stakers
            omni.transfer(stakers[i], amounts[i]);
            
            // Approve GenesisStakeV2 to spend tokens
            vm.prank(stakers[i]);
            omni.approve(address(genesisStake), amounts[i]);
            
            // Directly update balance in GenesisStakeV2
            // balanceOf mapping is at slot 0, but we need to account for inherited storage
            // OwnableUpgradeable: _owner (slot 0)
            // PausableUpgradeable: _paused (slot 1)
            // balanceOf mapping starts at slot 2
            vm.store(
                address(genesisStake),
                keccak256(abi.encode(stakers[i], uint256(2))), // balanceOf mapping slot
                bytes32(amounts[i])
            );
        }

        omni.transfer(address(merkleDistributor), initialSupply / 2);
    }

    // Generate an ERC7683 order to check merkle distributor's call to inbox against
    function generateERC7683Order(address addr, uint256 amount)
        public
        view
        returns (IERC7683.OnchainCrossChainOrder memory)
    {
        assertGe(genesisStake.balanceOf(addr), amount, "Insufficient staker balance");

        SolverNet.Deposit memory deposit = SolverNet.Deposit({ token: address(omni), amount: uint96(amount * 2) });

        SolverNet.Call[] memory call = new SolverNet.Call[](1);
        call[0] = SolverNet.Call({
            target: STAKING,
            selector: IStaking.delegateFor.selector,
            value: amount * 2,
            params: abi.encode(addr, validator)
        });

        SolverNet.OrderData memory orderData = SolverNet.OrderData({
            owner: addr,
            destChainId: omniPortal.omniChainId(),
            deposit: deposit,
            calls: call,
            expenses: new SolverNet.TokenExpense[](0)
        });

        return IERC7683.OnchainCrossChainOrder({
            fillDeadline: uint32(block.timestamp + 24 hours),
            orderDataType: ORDER_DATA_TYPEHASH,
            orderData: abi.encode(orderData)
        });
    }

    // Simple test to verify that our merkle tree is compatible with the distributor
    // function test_claim() public {
    //     for (uint256 i; i < addrCount; ++i) {
    //         vm.prank(stakers[i]);
    //         merkleDistributor.claim(i, stakers[i], amounts[i], proofs[i]);
    //         assertEq(omni.balanceOf(stakers[i]), amounts[i]);
    //     }
    // }

    function test_migrateToOmni_reverts() public {
        // Cannot migrate if validator is zero address
        vm.prank(stakers[0]);
        vm.expectRevert(MerkleDistributorWithoutDeadline.ZeroAddress.selector);
        merkleDistributor.upgradeStake(address(0), 0, amounts[0], proofs[0]);

        // Cannot migrate if proof is invalid
        bytes32 proof = proofs[0][0];
        proofs[0][0] = bytes32(uint256(proof) + 1);
        vm.prank(stakers[0]);
        vm.expectRevert(MerkleDistributor.InvalidProof.selector);
        merkleDistributor.upgradeStake(validator, 0, amounts[0], proofs[0]);
        proofs[0][0] = proof;

        // Fully claim all stake and rewards
        vm.prank(stakers[0]);
        merkleDistributor.upgradeStake(validator, 0, amounts[0], proofs[0]);

        // Cannot migrate if user has already claimed
        vm.prank(stakers[0]);
        vm.expectRevert(MerkleDistributorWithoutDeadline.InsufficientAmount.selector);
        merkleDistributor.upgradeStake(validator, 0, amounts[0], proofs[0]);
    }

    // Fully test upgradeStake for all members of the merkle tree
    // function test_migrateToOmni_succeeds() public {
    //     for (uint256 i; i < addrCount; ++i) {
    //         uint256 inboxBalance = omni.balanceOf(address(inbox));

    //         // Get IERC7683 order and resolved orders
    //         vm.startPrank(stakers[i]);
    //         IERC7683.OnchainCrossChainOrder memory order = generateERC7683Order(stakers[i], amounts[i]);
    //         IERC7683.ResolvedCrossChainOrder memory resolved = inbox.resolve(order);

    //         // Confirm merkleDistributor is calling the inbox with the order and that the resolved order is emitted
    //         vm.expectCall(address(inbox), abi.encodeCall(MockSolverNetInbox.open, (order)));
    //         vm.expectEmit(true, true, true, true);
    //         emit IERC7683.Open(resolved.orderId, resolved);
    //         merkleDistributor.upgradeStake(validator, i, amounts[i], proofs[i]);
    //         vm.stopPrank();

    //         // Confirm the inbox balance has increased by the user's staked balance and claim reward
    //         assertEq(omni.balanceOf(address(inbox)), inboxBalance + amounts[i] * 2);
    //         assertEq(omni.balanceOf(stakers[i]), 0);
    //     }
    // }

    function test_migrateUserToOmni_reverts() public {
        for (uint256 i; i < addrCount; ++i) {
            bytes32 digest = merkleDistributor.getUpgradeDigest(stakers[i], validator, block.timestamp + 1);
            (uint8 v, bytes32 r, bytes32 s) = vm.sign(pks[i], digest);

            // Modify the signature to make it invalid
            s = bytes32(uint256(s) + 1);

            // Cannot migrate if signature is invalid
            vm.expectRevert(MerkleDistributorWithoutDeadline.InvalidSignature.selector);
            merkleDistributor.upgradeUserStake(
                stakers[i], validator, 0, amounts[i], proofs[i], v, r, s, block.timestamp + 1
            );
        }
    }

    // function test_migrateUserToOmni_succeeds() public {
    //     for (uint256 i; i < addrCount; ++i) {
    //         uint256 inboxBalance = omni.balanceOf(address(inbox));

    //         // Get IERC7683 order and resolved orders
    //         vm.startPrank(stakers[i]);
    //         IERC7683.OnchainCrossChainOrder memory order = generateERC7683Order(stakers[i], amounts[i]);
    //         IERC7683.ResolvedCrossChainOrder memory resolved = inbox.resolve(order);
    //         vm.stopPrank();

    //         // Get user signature for migration
    //         bytes32 digest = merkleDistributor.getUpgradeDigest(stakers[i], validator, block.timestamp + 1);
    //         (uint8 v, bytes32 r, bytes32 s) = vm.sign(pks[i], digest);

    //         // Confirm merkleDistributor is calling the inbox with the order and that the resolved order is emitted
    //         vm.expectCall(address(inbox), abi.encodeCall(MockSolverNetInbox.open, (order)));
    //         vm.expectEmit(true, true, true, true);
    //         emit IERC7683.Open(resolved.orderId, resolved);
    //         merkleDistributor.upgradeUserStake(
    //             stakers[i], validator, i, amounts[i], proofs[i], v, r, s, block.timestamp + 1
    //         );

    //         // Confirm the inbox balance has increased by the user's staked balance and claim reward
    //         assertEq(omni.balanceOf(address(inbox)), inboxBalance + amounts[i] * 2);
    //         assertEq(omni.balanceOf(stakers[i]), 0);
    //     }
    // }

    function test_unstake_reverts() public {
        // Test invalid proof
        bytes32 proof = proofs[0][0];
        proofs[0][0] = bytes32(uint256(proof) + 1);
        vm.prank(stakers[0]);
        vm.expectRevert(MerkleDistributor.InvalidProof.selector);
        merkleDistributor.unstake(0, amounts[0], proofs[0]);
        proofs[0][0] = proof;
    }

    function test_unstake_succeeds() public {
        for (uint256 i; i < addrCount; ++i) {
            uint256 initialBalance = omni.balanceOf(stakers[i]);
            uint256 initialStake = genesisStake.balanceOf(stakers[i]);

            // unstake and claim rewards
            vm.prank(stakers[i]);
            merkleDistributor.unstake(i, amounts[i], proofs[i]);

            // verify stake is removed and rewards were claimed
            assertEq(genesisStake.balanceOf(stakers[i]), 0);
            assertEq(omni.balanceOf(stakers[i]), initialBalance + amounts[i] + initialStake);
        }
    }

    function test_unstake_without_stake_or_rewards() public {
        // unstake all stakes first
        for (uint256 i; i < addrCount; ++i) {
            vm.prank(stakers[i]);
            merkleDistributor.unstake(i, amounts[i], proofs[i]);
        }

        // try to unstake again - should do nothing
        for (uint256 i; i < addrCount; ++i) {
            uint256 initialBalance = omni.balanceOf(stakers[i]);
            uint256 initialStake = genesisStake.balanceOf(stakers[i]);
            assertEq(initialStake, 0, "should have no stake");

            vm.prank(stakers[i]);
            merkleDistributor.unstake(i, amounts[i], proofs[i]);

            // verify no stake to remove and no rewards claimed
            assertEq(genesisStake.balanceOf(stakers[i]), 0);
            assertEq(omni.balanceOf(stakers[i]), initialBalance);
        }
    }
} 