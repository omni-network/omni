// SPDX-License-Identifier: GPL-3.0-only
pragma solidity 0.8.24;

import { TransparentUpgradeableProxy } from "@openzeppelin/contracts/proxy/transparent/TransparentUpgradeableProxy.sol";
import { MerkleProof } from "@openzeppelin/contracts/utils/cryptography/MerkleProof.sol";
import { MerkleGen } from "multiproof/src/MerkleGen.sol";
import { MockPortal } from "test/utils/MockPortal.sol";
import { IOmniPortal } from "src/interfaces/IOmniPortal.sol";
import { Predeploys } from "src/libraries/nomina/Predeploys.sol";
import { MockOmni } from "nomina/test/utils/MockOmni.sol";
import { Nomina } from "nomina/src/token/Nomina.sol";
import { NominaBridgeNative } from "src/token/nomina/NominaBridgeNative.sol";
import { NominaBridgeL1 } from "src/token/nomina/NominaBridgeL1.sol";
import { ConfLevel } from "src/libraries/ConfLevel.sol";
import { Test } from "forge-std/Test.sol";
import { console } from "forge-std/console.sol";
import { ERC20 } from "solady/src/tokens/ERC20.sol";

/**
 * @title NominaBridgeL1_Test
 * @notice Test suite for NominaBridgeNative contract.
 */
contract NominaBridgeL1_Test is Test {
    // events copied from NominaBridgeL1.sol
    event Bridge(address indexed payor, address indexed to, uint256 amount);
    event Withdraw(address indexed to, uint256 amount);
    event PostHaltWithdraw(address indexed to, uint256 amount);

    MockPortal portal;
    MockOmni omni;
    Nomina nomina;
    NominaBridgeL1Harness b;

    address owner;
    address minter;
    address mintAuthority;
    address proxyAdmin;
    address initialSupplyRecipient;
    uint8 conversionRate = 75;
    uint256 amount = 1 ether;
    uint256 totalSupply = 100_000_000 ether;

    /// @notice Helper struct for withdrawal data
    struct WithdrawalData {
        address account;
        uint256 amount;
    }

    /// @notice Helper function to generate merkle multiproof for withdrawals
    function _generateWithdrawalProof(WithdrawalData[] memory withdrawals, uint256[] memory selectedIndices)
        internal
        pure
        returns (
            address[] memory accounts,
            uint256[] memory amounts,
            bytes32[] memory multiProof,
            bool[] memory multiProofFlags,
            bytes32 root
        )
    {
        // create leaves for all withdrawals
        bytes32[] memory leaves = new bytes32[](withdrawals.length);
        for (uint256 i = 0; i < withdrawals.length; i++) {
            leaves[i] = keccak256(bytes.concat(keccak256(abi.encode(withdrawals[i].account, withdrawals[i].amount))));
        }

        // generate multiproof
        (multiProof, multiProofFlags, root) = MerkleGen.generateMultiproof(leaves, selectedIndices);

        // prepare accounts and amounts arrays for selected indices
        accounts = new address[](selectedIndices.length);
        amounts = new uint256[](selectedIndices.length);
        for (uint256 i = 0; i < selectedIndices.length; i++) {
            accounts[i] = withdrawals[selectedIndices[i]].account;
            amounts[i] = withdrawals[selectedIndices[i]].amount;
        }
    }

    function setUp() public {
        initialSupplyRecipient = makeAddr("initialSupplyRecipient");
        owner = makeAddr("owner");
        proxyAdmin = makeAddr("proxyAdmin");
        minter = makeAddr("minter");
        mintAuthority = makeAddr("mintAuthority");

        portal = new MockPortal();
        omni = new MockOmni(totalSupply, initialSupplyRecipient);

        nomina = new Nomina(address(omni), mintAuthority);
        vm.prank(mintAuthority);
        nomina.setMinter(minter);

        address impl = address(new NominaBridgeL1Harness(address(omni), address(nomina)));
        b = NominaBridgeL1Harness(
            address(
                new TransparentUpgradeableProxy(
                    impl, proxyAdmin, abi.encodeCall(NominaBridgeL1.initialize, (owner, address(portal)))
                )
            )
        );
        b.initializeV2();
    }

    function test_initialize() public {
        address impl = address(new NominaBridgeL1(address(omni), address(nomina)));
        address proxy = address(new TransparentUpgradeableProxy(impl, proxyAdmin, ""));

        // reverts
        vm.expectRevert("NominaBridge: no zero addr");
        NominaBridgeL1(proxy).initialize(owner, address(0));

        // succeeds
        NominaBridgeL1(proxy).initialize(owner, address(portal));

        // initializeV2 converts omni balance to nomina
        vm.prank(initialSupplyRecipient);
        omni.transfer(proxy, amount);
        NominaBridgeL1(proxy).initializeV2();

        // assert balance
        assertEq(nomina.balanceOf(proxy), amount * conversionRate);
        assertEq(omni.balanceOf(proxy), 0);
    }

    function test_bridge() public {
        address to = makeAddr("to");
        address payor = address(this);
        uint256 fee = b.bridgeFee(payor, to, amount);

        // requires amount > 0
        vm.expectRevert("NominaBridge: amount must be > 0");
        b.bridge(to, 0);

        // to must not be zero
        vm.expectRevert("NominaBridge: no bridge to zero");
        b.bridge(address(0), amount);

        // value must be greater than or equal fee
        vm.expectRevert("NominaBridge: insufficient fee");
        b.bridge{ value: fee - 1 }(to, amount);

        // requires allowance
        vm.expectRevert(ERC20.InsufficientAllowance.selector);
        b.bridge{ value: fee }(to, amount);

        nomina.approve(address(b), amount);

        // requires balance
        vm.expectRevert(ERC20.InsufficientBalance.selector);
        b.bridge{ value: fee }(to, amount);

        // succeeds
        //
        // fund payor
        vm.startPrank(initialSupplyRecipient);
        omni.approve(address(nomina), amount);
        nomina.convert(payor, amount);
        vm.stopPrank();

        // emits event
        vm.expectEmit();
        emit Bridge(payor, to, amount);

        // emits xcall
        vm.expectCall(
            address(portal),
            fee,
            abi.encodeCall(
                IOmniPortal.xcall,
                (
                    portal.omniChainId(),
                    ConfLevel.Finalized,
                    Predeploys.NominaBridgeNative,
                    abi.encodeCall(NominaBridgeNative.withdraw, (payor, to, amount)),
                    b.XCALL_WITHDRAW_GAS_LIMIT()
                )
            )
        );
        b.bridge{ value: fee }(to, amount);

        // assert balance change
        assertEq(nomina.balanceOf(address(b)), amount);
        assertEq(nomina.balanceOf(payor), (amount * conversionRate) - amount);
    }

    function test_withdraw() public {
        address to = makeAddr("to");
        uint64 omniChainId = portal.omniChainId();
        uint64 gasLimit = new NominaBridgeNative().XCALL_WITHDRAW_GAS_LIMIT();

        // sender must be portal
        vm.expectRevert("NominaBridge: not xcall");
        b.withdraw(to, amount);

        // xmsg must be from native bridge
        vm.expectRevert("NominaBridge: not bridge");
        portal.mockXCall({
            sourceChainId: omniChainId,
            sender: address(1234), // wrong
            to: address(b),
            data: abi.encodeCall(NominaBridgeL1.withdraw, (to, amount)),
            gasLimit: gasLimit
        });

        // xmsg must be from nomina evm
        vm.expectRevert("NominaBridge: not omni portal");
        portal.mockXCall({
            sourceChainId: omniChainId + 1, // wrong
            sender: Predeploys.NominaBridgeNative,
            to: address(b),
            data: abi.encodeCall(NominaBridgeL1.withdraw, (to, amount)),
            gasLimit: gasLimit
        });

        // succeeds
        //
        // need to fund bridge first
        vm.startPrank(initialSupplyRecipient);
        omni.approve(address(nomina), amount);
        nomina.convert(address(b), amount);
        vm.stopPrank();

        // emit event
        vm.expectEmit();
        emit Withdraw(to, amount);

        // tranfers amount to to
        vm.expectCall(address(nomina), abi.encodeCall(nomina.transfer, (to, amount)));
        uint256 gasUsed = portal.mockXCall({
            sourceChainId: portal.omniChainId(),
            sender: Predeploys.NominaBridgeNative,
            to: address(b),
            data: abi.encodeCall(NominaBridgeL1.withdraw, (to, amount)),
            gasLimit: gasLimit
        });

        // assert balance change
        assertEq(nomina.balanceOf(to), amount);
        assertEq(nomina.balanceOf(address(b)), (amount * conversionRate) - amount);

        // log gas, to inform xcall gas limit
        console.log("NominaBridgeL1.withdraw gas used: ", gasUsed);
    }

    function test_pauseBridging() public {
        address to = makeAddr("to");
        bytes32 action = b.ACTION_BRIDGE();

        // pause bridging
        vm.prank(owner);
        b.pause(action);

        // assert paused
        assertTrue(b.isPaused(action));

        // bridge reverts
        vm.expectRevert("NominaBridge: paused");
        b.bridge(to, amount);

        // unpause bridging
        vm.prank(owner);
        b.unpause(action);

        // assert unpaused
        assertFalse(b.isPaused(action));

        // bridge not paused (reverts, but not due to pause)
        vm.expectRevert("NominaBridge: insufficient fee");
        b.bridge(to, amount);
    }

    function test_pauseWithdraws() public {
        address to = makeAddr("to");
        bytes32 action = b.ACTION_WITHDRAW();

        // pause withdraws
        vm.prank(owner);
        b.pause(action);

        // assert paused
        assertTrue(b.isPaused(action));

        // withdraw reverts
        vm.expectRevert("NominaBridge: paused");
        b.withdraw(to, amount);

        // unpause
        vm.prank(owner);
        b.unpause(action);

        // assert unpaued
        assertFalse(b.isPaused(action));

        // no longer paused
        vm.expectRevert("NominaBridge: not xcall");
        b.withdraw(to, amount);
    }

    function test_pauseAll() public {
        address to = makeAddr("to");

        // pause all
        vm.prank(owner);
        b.pause();

        // assert actions paus
        assertTrue(b.isPaused(b.ACTION_BRIDGE()));
        assertTrue(b.isPaused(b.ACTION_WITHDRAW()));

        // bridge reverts
        vm.expectRevert("NominaBridge: paused");
        b.bridge(to, amount);

        // withdraw reverts
        vm.expectRevert("NominaBridge: paused");
        b.withdraw(to, amount);

        // unpause all
        vm.prank(owner);
        b.unpause();

        assertFalse(b.isPaused(b.ACTION_BRIDGE()));
        assertFalse(b.isPaused(b.ACTION_WITHDRAW()));
    }

    function test_initializeV3() public {
        bytes32 root = keccak256("test root");

        // should not be paused initially
        assertFalse(b.isPaused(b.ACTION_BRIDGE()));
        assertFalse(b.isPaused(b.ACTION_WITHDRAW()));

        // initialize v3
        vm.prank(proxyAdmin);
        b.initializeV3(root);

        // assert root is set
        assertEq(b.postHaltRoot(), root);

        // assert both actions are paused
        assertTrue(b.isPaused(b.ACTION_BRIDGE()));
        assertTrue(b.isPaused(b.ACTION_WITHDRAW()));
    }

    function test_initializeV3_cannotReinitialize() public {
        bytes32 root = keccak256("test root");

        // initialize v3
        vm.prank(proxyAdmin);
        b.initializeV3(root);

        // cannot reinitialize
        vm.expectRevert();
        vm.prank(proxyAdmin);
        b.initializeV3(root);
    }

    function test_postHaltWithdraw_singleAccount() public {
        // setup
        address account = makeAddr("account");
        uint256 withdrawAmount = 5 ether;

        // fund bridge
        vm.startPrank(initialSupplyRecipient);
        omni.approve(address(nomina), withdrawAmount);
        nomina.convert(address(b), withdrawAmount);
        vm.stopPrank();

        // create withdrawal data
        WithdrawalData[] memory withdrawals = new WithdrawalData[](1);
        withdrawals[0] = WithdrawalData(account, withdrawAmount);

        uint256[] memory selectedIndices = new uint256[](1);
        selectedIndices[0] = 0;

        (
            address[] memory accounts,
            uint256[] memory amounts,
            bytes32[] memory multiProof,
            bool[] memory multiProofFlags,
            bytes32 root
        ) = _generateWithdrawalProof(withdrawals, selectedIndices);

        // initialize v3 with root
        vm.prank(proxyAdmin);
        b.initializeV3(root);

        // expect event
        vm.expectEmit();
        emit PostHaltWithdraw(account, withdrawAmount);

        // withdraw
        b.postHaltWithdraw(accounts, amounts, multiProof, multiProofFlags);

        // assert balance
        assertEq(nomina.balanceOf(account), withdrawAmount);
        assertTrue(b.postHaltClaimed(account));
    }

    function test_postHaltWithdraw_twoAccounts() public {
        // setup
        address account1 = makeAddr("account1");
        address account2 = makeAddr("account2");
        uint256 amount1 = 1 ether;
        uint256 amount2 = 2 ether;

        // fund bridge
        vm.startPrank(initialSupplyRecipient);
        omni.approve(address(nomina), amount1 + amount2);
        nomina.convert(address(b), amount1 + amount2);
        vm.stopPrank();

        // create withdrawal data
        WithdrawalData[] memory withdrawals = new WithdrawalData[](2);
        withdrawals[0] = WithdrawalData(account1, amount1);
        withdrawals[1] = WithdrawalData(account2, amount2);

        uint256[] memory selectedIndices = new uint256[](2);
        selectedIndices[0] = 0;
        selectedIndices[1] = 1;

        (
            address[] memory accounts,
            uint256[] memory amounts,
            bytes32[] memory multiProof,
            bool[] memory multiProofFlags,
            bytes32 root
        ) = _generateWithdrawalProof(withdrawals, selectedIndices);

        // initialize v3 with root
        vm.prank(proxyAdmin);
        b.initializeV3(root);

        // expect events
        vm.expectEmit();
        emit PostHaltWithdraw(account1, amount1);
        vm.expectEmit();
        emit PostHaltWithdraw(account2, amount2);

        // withdraw
        b.postHaltWithdraw(accounts, amounts, multiProof, multiProofFlags);

        // assert balances
        assertEq(nomina.balanceOf(account1), amount1);
        assertEq(nomina.balanceOf(account2), amount2);
        assertTrue(b.postHaltClaimed(account1));
        assertTrue(b.postHaltClaimed(account2));
    }

    function test_postHaltWithdraw_multipleAccounts() public {
        // setup
        uint256 numAccounts = 5;
        uint256 totalAmount = 0;

        WithdrawalData[] memory withdrawals = new WithdrawalData[](numAccounts);
        for (uint256 i = 0; i < numAccounts; i++) {
            address account = makeAddr(string(abi.encodePacked("account", i)));
            uint256 amt = (i + 1) * 1 ether;
            withdrawals[i] = WithdrawalData(account, amt);
            totalAmount += amt;
        }

        // fund bridge
        vm.startPrank(initialSupplyRecipient);
        omni.approve(address(nomina), totalAmount);
        nomina.convert(address(b), totalAmount);
        vm.stopPrank();

        // select all accounts
        uint256[] memory selectedIndices = new uint256[](numAccounts);
        for (uint256 i = 0; i < numAccounts; i++) {
            selectedIndices[i] = i;
        }

        (
            address[] memory accounts,
            uint256[] memory amounts,
            bytes32[] memory multiProof,
            bool[] memory multiProofFlags,
            bytes32 root
        ) = _generateWithdrawalProof(withdrawals, selectedIndices);

        // initialize v3 with root
        vm.prank(proxyAdmin);
        b.initializeV3(root);

        // withdraw all
        b.postHaltWithdraw(accounts, amounts, multiProof, multiProofFlags);

        // assert all claimed
        for (uint256 i = 0; i < numAccounts; i++) {
            assertEq(nomina.balanceOf(withdrawals[i].account), withdrawals[i].amount);
            assertTrue(b.postHaltClaimed(withdrawals[i].account));
        }
    }

    function test_postHaltWithdraw_partialClaim() public {
        // setup three accounts
        WithdrawalData[] memory withdrawals = new WithdrawalData[](3);
        withdrawals[0] = WithdrawalData(makeAddr("account1"), 1 ether);
        withdrawals[1] = WithdrawalData(makeAddr("account2"), 2 ether);
        withdrawals[2] = WithdrawalData(makeAddr("account3"), 3 ether);

        // fund bridge
        vm.startPrank(initialSupplyRecipient);
        omni.approve(address(nomina), 6 ether);
        nomina.convert(address(b), 6 ether);
        vm.stopPrank();

        // get root from all withdrawals
        bytes32 root;
        {
            uint256[] memory allIndices = new uint256[](3);
            allIndices[0] = 0;
            allIndices[1] = 1;
            allIndices[2] = 2;
            (,,,, root) = _generateWithdrawalProof(withdrawals, allIndices);
        }

        // initialize v3 with root
        vm.prank(proxyAdmin);
        b.initializeV3(root);

        // first claim account1 and account3 (indices 0 and 2)
        {
            uint256[] memory batch1Indices = new uint256[](2);
            batch1Indices[0] = 0;
            batch1Indices[1] = 2;

            (address[] memory accounts, uint256[] memory amounts, bytes32[] memory proof, bool[] memory flags,) =
                _generateWithdrawalProof(withdrawals, batch1Indices);

            b.postHaltWithdraw(accounts, amounts, proof, flags);

            // assert batch 1 claimed
            assertTrue(b.postHaltClaimed(withdrawals[0].account));
            assertFalse(b.postHaltClaimed(withdrawals[1].account));
            assertTrue(b.postHaltClaimed(withdrawals[2].account));
        }

        // now claim account2 (index 1)
        {
            uint256[] memory batch2Indices = new uint256[](1);
            batch2Indices[0] = 1;

            (address[] memory accounts, uint256[] memory amounts, bytes32[] memory proof, bool[] memory flags,) =
                _generateWithdrawalProof(withdrawals, batch2Indices);

            b.postHaltWithdraw(accounts, amounts, proof, flags);
        }

        // assert all claimed
        assertTrue(b.postHaltClaimed(withdrawals[0].account));
        assertTrue(b.postHaltClaimed(withdrawals[1].account));
        assertTrue(b.postHaltClaimed(withdrawals[2].account));
    }

    function test_postHaltWithdraw_revertsWhenNoRootSet() public {
        address account = makeAddr("account");
        uint256 withdrawAmount = 1 ether;

        address[] memory accounts = new address[](1);
        accounts[0] = account;

        uint256[] memory amounts = new uint256[](1);
        amounts[0] = withdrawAmount;

        bytes32[] memory multiProof = new bytes32[](0);
        bool[] memory multiProofFlags = new bool[](0);

        // should revert when no root set
        vm.expectRevert("NominaBridge: no root set");
        b.postHaltWithdraw(accounts, amounts, multiProof, multiProofFlags);
    }

    function test_postHaltWithdraw_revertsWhenLengthMismatch() public {
        bytes32 root = keccak256("test root");
        vm.prank(proxyAdmin);
        b.initializeV3(root);

        address[] memory accounts = new address[](2);
        accounts[0] = makeAddr("account1");
        accounts[1] = makeAddr("account2");

        uint256[] memory amounts = new uint256[](1); // wrong length
        amounts[0] = 1 ether;

        bytes32[] memory multiProof = new bytes32[](0);
        bool[] memory multiProofFlags = new bool[](0);

        vm.expectRevert("NominaBridge: length mismatch");
        b.postHaltWithdraw(accounts, amounts, multiProof, multiProofFlags);
    }

    function test_postHaltWithdraw_revertsWhenAlreadyClaimed() public {
        address account = makeAddr("account");
        uint256 withdrawAmount = 1 ether;

        // fund bridge
        vm.startPrank(initialSupplyRecipient);
        omni.approve(address(nomina), withdrawAmount * 2);
        nomina.convert(address(b), withdrawAmount * 2);
        vm.stopPrank();

        // create withdrawal data
        WithdrawalData[] memory withdrawals = new WithdrawalData[](1);
        withdrawals[0] = WithdrawalData(account, withdrawAmount);

        uint256[] memory selectedIndices = new uint256[](1);
        selectedIndices[0] = 0;

        (
            address[] memory accounts,
            uint256[] memory amounts,
            bytes32[] memory multiProof,
            bool[] memory multiProofFlags,
            bytes32 root
        ) = _generateWithdrawalProof(withdrawals, selectedIndices);

        vm.prank(proxyAdmin);
        b.initializeV3(root);

        // first claim succeeds
        b.postHaltWithdraw(accounts, amounts, multiProof, multiProofFlags);

        // second claim reverts
        vm.expectRevert("NominaBridge: already claimed");
        b.postHaltWithdraw(accounts, amounts, multiProof, multiProofFlags);
    }

    function test_postHaltWithdraw_revertsWhenZeroAddress() public {
        bytes32 root = keccak256("test root");
        vm.prank(proxyAdmin);
        b.initializeV3(root);

        address[] memory accounts = new address[](1);
        accounts[0] = address(0);

        uint256[] memory amounts = new uint256[](1);
        amounts[0] = 1 ether;

        bytes32[] memory multiProof = new bytes32[](0);
        bool[] memory multiProofFlags = new bool[](0);

        vm.expectRevert("NominaBridge: no zero addr");
        b.postHaltWithdraw(accounts, amounts, multiProof, multiProofFlags);
    }

    function test_postHaltWithdraw_revertsWhenZeroAmount() public {
        bytes32 root = keccak256("test root");
        vm.prank(proxyAdmin);
        b.initializeV3(root);

        address[] memory accounts = new address[](1);
        accounts[0] = makeAddr("account");

        uint256[] memory amounts = new uint256[](1);
        amounts[0] = 0;

        bytes32[] memory multiProof = new bytes32[](0);
        bool[] memory multiProofFlags = new bool[](0);

        vm.expectRevert("NominaBridge: amount must be > 0");
        b.postHaltWithdraw(accounts, amounts, multiProof, multiProofFlags);
    }

    function test_postHaltWithdraw_revertsWhenInvalidProof() public {
        address account = makeAddr("account");
        uint256 withdrawAmount = 1 ether;

        // create valid withdrawal
        WithdrawalData[] memory withdrawals = new WithdrawalData[](1);
        withdrawals[0] = WithdrawalData(account, withdrawAmount);

        uint256[] memory selectedIndices = new uint256[](1);
        selectedIndices[0] = 0;

        (,,,, bytes32 root) = _generateWithdrawalProof(withdrawals, selectedIndices);

        vm.prank(proxyAdmin);
        b.initializeV3(root);

        // try to claim different amount (invalid proof)
        address[] memory accounts = new address[](1);
        accounts[0] = account;

        uint256[] memory amounts = new uint256[](1);
        amounts[0] = withdrawAmount + 1; // wrong amount

        bytes32[] memory multiProof = new bytes32[](0);
        bool[] memory multiProofFlags = new bool[](0);

        vm.expectRevert("NominaBridge: invalid proof");
        b.postHaltWithdraw(accounts, amounts, multiProof, multiProofFlags);
    }

    function test_postHaltWithdraw_revertsWhenWrongAccount() public {
        address account1 = makeAddr("account1");
        address account2 = makeAddr("account2");
        uint256 withdrawAmount = 1 ether;

        // create valid withdrawal for account1
        WithdrawalData[] memory withdrawals = new WithdrawalData[](1);
        withdrawals[0] = WithdrawalData(account1, withdrawAmount);

        uint256[] memory selectedIndices = new uint256[](1);
        selectedIndices[0] = 0;

        (,,,, bytes32 root) = _generateWithdrawalProof(withdrawals, selectedIndices);

        vm.prank(proxyAdmin);
        b.initializeV3(root);

        // try to claim for account2 (invalid proof)
        address[] memory accounts = new address[](1);
        accounts[0] = account2; // wrong account

        uint256[] memory amounts = new uint256[](1);
        amounts[0] = withdrawAmount;

        bytes32[] memory multiProof = new bytes32[](0);
        bool[] memory multiProofFlags = new bool[](0);

        vm.expectRevert("NominaBridge: invalid proof");
        b.postHaltWithdraw(accounts, amounts, multiProof, multiProofFlags);
    }

    function test_bridgeAndWithdrawPausedAfterInitializeV3() public {
        bytes32 root = keccak256("test root");

        // initialize v3
        vm.prank(proxyAdmin);
        b.initializeV3(root);

        // bridge should be paused
        address to = makeAddr("to");

        vm.expectRevert("NominaBridge: paused");
        b.bridge(to, amount);

        // withdraw should be paused
        vm.expectRevert("NominaBridge: paused");
        b.withdraw(to, amount);
    }

    function test_postHaltWithdraw_multipleBatches() public {
        // setup 6 accounts
        WithdrawalData[] memory withdrawals = new WithdrawalData[](6);
        uint256 totalAmount = 21 ether; // 1+2+3+4+5+6

        for (uint256 i = 0; i < 6; i++) {
            withdrawals[i] = WithdrawalData(makeAddr(string(abi.encodePacked("account", i))), (i + 1) * 1 ether);
        }

        // fund bridge
        vm.startPrank(initialSupplyRecipient);
        omni.approve(address(nomina), totalAmount);
        nomina.convert(address(b), totalAmount);
        vm.stopPrank();

        // generate proof for all and get root
        uint256[] memory allIndices = new uint256[](6);
        for (uint256 i = 0; i < 6; i++) {
            allIndices[i] = i;
        }
        (,,,, bytes32 root) = _generateWithdrawalProof(withdrawals, allIndices);

        // initialize v3 with root
        vm.prank(proxyAdmin);
        b.initializeV3(root);

        // claim first batch (accounts 0, 1, 2)
        {
            uint256[] memory batch1Indices = new uint256[](3);
            batch1Indices[0] = 0;
            batch1Indices[1] = 1;
            batch1Indices[2] = 2;

            (address[] memory accounts, uint256[] memory amounts, bytes32[] memory proof, bool[] memory flags,) =
                _generateWithdrawalProof(withdrawals, batch1Indices);

            b.postHaltWithdraw(accounts, amounts, proof, flags);

            // verify first batch claimed
            assertTrue(b.postHaltClaimed(withdrawals[0].account));
            assertTrue(b.postHaltClaimed(withdrawals[1].account));
            assertTrue(b.postHaltClaimed(withdrawals[2].account));
        }

        // claim second batch (accounts 3, 4, 5)
        {
            uint256[] memory batch2Indices = new uint256[](3);
            batch2Indices[0] = 3;
            batch2Indices[1] = 4;
            batch2Indices[2] = 5;

            (address[] memory accounts, uint256[] memory amounts, bytes32[] memory proof, bool[] memory flags,) =
                _generateWithdrawalProof(withdrawals, batch2Indices);

            b.postHaltWithdraw(accounts, amounts, proof, flags);
        }

        // verify all claimed
        for (uint256 i = 0; i < 6; i++) {
            assertTrue(b.postHaltClaimed(withdrawals[i].account));
            assertEq(nomina.balanceOf(withdrawals[i].account), withdrawals[i].amount);
        }
    }
}

contract NominaBridgeL1Harness is NominaBridgeL1 {
    constructor(address omni, address nomina) NominaBridgeL1(omni, nomina) { }
}
