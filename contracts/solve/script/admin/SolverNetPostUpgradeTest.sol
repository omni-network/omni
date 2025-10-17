// SPDX-License-Identifier: GPL-3.0-only
pragma solidity 0.8.24;

import { Test } from "forge-std/Test.sol";
import { VmSafe } from "forge-std/Vm.sol";
import { MockERC721 } from "test/utils/MockERC721.sol";
import { Refunder } from "test/utils/Refunder.sol";

import { SolverNetInbox, ISolverNetInbox } from "src/SolverNetInbox.sol";
import { SolverNetOutbox, ISolverNetOutbox } from "src/SolverNetOutbox.sol";
import { SolverNetExecutor, ISolverNetExecutor } from "src/SolverNetExecutor.sol";
import { IERC7683 } from "src/erc7683/IERC7683.sol";
import { SolverNet } from "src/lib/SolverNet.sol";

contract SolverNetPostUpgradeTest is Test {
    SolverNetInbox internal inbox;
    SolverNetOutbox internal outbox;
    SolverNetExecutor internal executor;

    Refunder internal refunder;
    MockERC721 internal milady;

    bytes32 constant ORDERDATA_TYPEHASH = keccak256(
        "OrderData(address owner,uint64 destChainId,Deposit deposit,Call[] calls,TokenExpense[] expenses)Deposit(address token,uint96 amount)Call(address target,bytes4 selector,uint256 value,bytes params)TokenExpense(address spender,address token,uint96 amount)"
    );

    address internal owner;
    address internal solver = makeAddr("solver");
    address internal user = makeAddr("user");

    function runInbox(address addr) public {
        (VmSafe.CallerMode mode,,) = vm.readCallers();
        require(mode == VmSafe.CallerMode.None, "no broadcast");

        _setupInbox(addr);
        _checkInboxConfigs();
        _openOrder();
    }

    function runOutbox(address addr, uint64[] calldata chainIds, ISolverNetOutbox.InboxConfig[] calldata configs)
        public
    {
        (VmSafe.CallerMode mode,,) = vm.readCallers();
        require(mode == VmSafe.CallerMode.None, "no broadcast");

        _setupOutbox(addr);
        _checkOutboxConfigs(chainIds, configs);
        _fillOrder(address(milady), MockERC721.mintTo.selector, 0, abi.encode(user), chainIds);

        assertEq(milady.balanceOf(user), chainIds.length, "user should receive 1 NFT per origin chain order");
    }

    function runExecutor(address addr, uint64[] calldata chainIds) public {
        (VmSafe.CallerMode mode,,) = vm.readCallers();
        require(mode == VmSafe.CallerMode.None, "no broadcast");

        _setupExecutor(addr);
        _executeAndTransfer(chainIds);
        _executeAndTransfer721(chainIds);
    }

    function _setupInbox(address addr) internal {
        inbox = SolverNetInbox(addr);
        owner = inbox.owner();
    }

    function _setupOutbox(address addr) internal {
        outbox = SolverNetOutbox(addr);
        owner = outbox.owner();
        vm.prank(owner);
        outbox.grantRoles(solver, 1); // 1 = SOLVER

        refunder = new Refunder();
        milady = new MockERC721("Milady Maker", "MILADY", "https://www.miladymaker.net/milady/json/");
    }

    function _setupExecutor(address addr) internal {
        executor = SolverNetExecutor(payable(addr));
        _setupOutbox(executor.outbox());
    }

    function _checkInboxConfigs() internal view {
        assertTrue(inbox.getOutbox(uint64(block.chainid)) != address(0), "outbox should be set");
    }

    function _checkOutboxConfigs(uint64[] calldata chainIds, ISolverNetOutbox.InboxConfig[] calldata configs)
        internal
        view
    {
        for (uint256 i; i < chainIds.length; ++i) {
            assertEq(keccak256(abi.encode(outbox.getInboxConfig(chainIds[i]))), keccak256(abi.encode(configs[i])));
        }
    }

    function _openOrder() internal {
        SolverNet.Deposit memory deposit = SolverNet.Deposit({ token: address(0), amount: 1 ether });

        SolverNet.Call[] memory calls = new SolverNet.Call[](1);
        calls[0] = SolverNet.Call({ target: user, selector: bytes4(hex"00000000"), value: 1 ether, params: "" });

        SolverNet.OrderData memory orderData = SolverNet.OrderData({
            owner: user,
            destChainId: uint64(block.chainid == 1 ? 10 : 1),
            deposit: deposit,
            calls: calls,
            expenses: new SolverNet.TokenExpense[](0)
        });

        IERC7683.OnchainCrossChainOrder memory order = IERC7683.OnchainCrossChainOrder({
            fillDeadline: type(uint32).max, orderDataType: ORDERDATA_TYPEHASH, orderData: abi.encode(orderData)
        });

        bytes32 id = ISolverNetInboxTemp(address(inbox)).getNextOrderId(user);
        SolverNet.FillOriginData memory fillOriginData = SolverNet.FillOriginData({
            srcChainId: uint64(block.chainid),
            destChainId: uint64(block.chainid == 1 ? 10 : 1),
            fillDeadline: type(uint32).max,
            calls: calls,
            expenses: new SolverNet.TokenExpense[](0)
        });

        vm.deal(user, 1 ether);
        vm.prank(user);
        vm.expectEmit(address(inbox));
        emit ISolverNetInbox.FillOriginData(id, fillOriginData);
        inbox.open{ value: 1 ether }(order);
    }

    // `value` is shared between deposit and call to avoid stack too deep
    function _fillOrder(address target, bytes4 selector, uint256 value, bytes memory params, uint64[] memory chainIds)
        internal
    {
        for (uint256 i; i < chainIds.length; ++i) {
            SolverNet.Call[] memory calls = new SolverNet.Call[](1);
            calls[0] = SolverNet.Call({ target: target, selector: selector, value: value, params: params });
            SolverNet.TokenExpense[] memory expenses = new SolverNet.TokenExpense[](0);
            SolverNet.FillOriginData memory fillOriginData = SolverNet.FillOriginData({
                srcChainId: chainIds[chainIds.length - 1 - i],
                destChainId: uint64(block.chainid),
                fillDeadline: type(uint32).max,
                calls: calls,
                expenses: expenses
            });

            bytes32 orderId = bytes32(type(uint256).max - i);
            uint256 fillFee = outbox.fillFee(abi.encode(fillOriginData));
            bool sameChain = chainIds[chainIds.length - 1 - i] == uint64(block.chainid);

            // If same chain order, it needs to be opened first
            if (sameChain) {
                ISolverNetOutbox.InboxConfig memory config = outbox.getInboxConfig(uint64(block.chainid));
                inbox = SolverNetInbox(config.inbox);
                orderId = ISolverNetInboxTemp(address(inbox)).getNextOrderId(user);
                vm.deal(user, value);

                SolverNet.Deposit memory deposit = SolverNet.Deposit({ token: address(0), amount: uint96(value) });
                SolverNet.OrderData memory orderData = SolverNet.OrderData({
                    owner: user, destChainId: uint64(block.chainid), deposit: deposit, calls: calls, expenses: expenses
                });
                IERC7683.OnchainCrossChainOrder memory order = IERC7683.OnchainCrossChainOrder({
                    fillDeadline: type(uint32).max, orderDataType: ORDERDATA_TYPEHASH, orderData: abi.encode(orderData)
                });
                assertTrue(inbox.validate(order), "order should be valid");

                vm.prank(user);
                inbox.open{ value: value }(order);
            }

            vm.deal(solver, value + fillFee);
            vm.prank(solver);
            outbox.fill{ value: value + fillFee }(orderId, abi.encode(fillOriginData), abi.encode(solver));

            assertTrue(outbox.didFill(orderId, abi.encode(fillOriginData)), "order should be filled");
        }
    }

    function _executeAndTransfer(uint64[] calldata chainIds) internal {
        bytes memory executeAndTransferParams = abi.encode(address(0), user, address(refunder), "");
        for (uint256 i; i < chainIds.length; ++i) {
            uint64[] memory _chainIds = new uint64[](1);
            _chainIds[0] = chainIds[i];

            vm.deal(user, 0); // Wipe user balance before each fill
            _fillOrder(
                address(0), ISolverNetExecutor.executeAndTransfer.selector, 1 ether, executeAndTransferParams, _chainIds
            );

            assertEq(user.balance, 1 ether, "user should have 1 ETH after last order");
        }
    }

    function _executeAndTransfer721(uint64[] calldata chainIds) internal {
        for (uint256 i; i < chainIds.length; ++i) {
            uint64[] memory _chainIds = new uint64[](1);
            _chainIds[0] = chainIds[i];

            bytes memory executeAndTransfer721Params =
                abi.encode(address(milady), i + 1, user, address(milady), abi.encodeCall(MockERC721.mint, ()));
            _fillOrder(
                address(0), ISolverNetExecutor.executeAndTransfer721.selector, 0, executeAndTransfer721Params, _chainIds
            );

            assertEq(milady.ownerOf(i + 1), user, "nft should be owned by user");
        }
    }
}

interface ISolverNetInboxTemp {
    function getNextOrderId(address user) external view returns (bytes32);
}
