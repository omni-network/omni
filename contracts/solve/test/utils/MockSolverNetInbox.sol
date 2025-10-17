// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.24;

import { ReentrancyGuard } from "solady/src/utils/ReentrancyGuard.sol";
import { DeployedAt } from "../../src/util/DeployedAt.sol";
import { XAppBase } from "../../src/lib/XAppBase.sol";
import { MailboxClient } from "../../src/ext/MailboxClient.sol";
import { IERC7683 } from "../../src/erc7683/IERC7683.sol";
import { ISolverNetInbox } from "../../src/interfaces/ISolverNetInbox.sol";
import { IMessageRecipient } from "@hyperlane-xyz/core/contracts/interfaces/IMessageRecipient.sol";
import { SafeTransferLib } from "solady/src/utils/SafeTransferLib.sol";
import { SolverNet } from "../../src/lib/SolverNet.sol";
import { AddrUtils } from "../../src/lib/AddrUtils.sol";
import { HashLib } from "../../src/lib/HashLib.sol";
import { IPermit2, ISignatureTransfer } from "@uniswap/permit2/src/interfaces/IPermit2.sol";

/**
 * @title MockSolverNetInbox
 * @notice Entrypoint and alt-mempool for user solve orders.
 * @dev This mock implementation is ownerless and is not initialized.
 */
contract MockSolverNetInbox is
    ReentrancyGuard,
    DeployedAt,
    XAppBase,
    MailboxClient,
    ISolverNetInbox,
    IMessageRecipient
{
    using SafeTransferLib for address;
    using AddrUtils for address;
    using AddrUtils for bytes32;

    /**
     * @notice Maximum number of calls and expenses in an order.
     */
    uint8 internal constant MAX_ARRAY_SIZE = 32;

    /**
     * @notice Buffer for closing orders after fill deadline to give Omni Core relayer time to act.
     */
    uint256 internal constant CLOSE_BUFFER = 6 hours;

    /**
     * @notice Action ID for xsubmissions, used as Pauseable key in OmniPortal
     */
    bytes32 internal constant ACTION_XSUBMIT = keccak256("xsubmit");

    /**
     * @notice Key for pausing the `open` function.
     */
    bytes32 internal constant OPEN = keccak256("OPEN");

    /**
     * @notice Key for pausing the `close` function.
     */
    bytes32 internal constant CLOSE = keccak256("CLOSE");

    uint8 internal constant NONE_PAUSED = 0;
    uint8 internal constant OPEN_PAUSED = 1;
    uint8 internal constant CLOSE_PAUSED = 2;
    uint8 internal constant ALL_PAUSED = 3;

    /**
     * @notice Canonical permit2 contract address.
     */
    IPermit2 internal constant PERMIT2 = IPermit2(0x000000000022D473030F116dDEE9F6B43aC78BA3);

    /**
     * @dev Incremental order offset for source inbox orders.
     */
    uint248 internal _offset;

    /**
     * @notice Pause state.
     * @dev 0 = no pause, 1 = open paused, 2 = close paused, 3 = all paused.
     */
    uint8 public pauseState;

    /**
     * @notice Addresses of the outbox contracts.
     */
    mapping(uint64 chainId => address outbox) internal _outboxes;

    /**
     * @notice Map order ID to header parameters.
     * @dev (owner, destChainId, fillDeadline)
     */
    mapping(bytes32 id => SolverNet.Header) internal _orderHeader;

    /**
     * @notice Map order ID to deposit parameters.
     * @dev (token, amount)
     */
    mapping(bytes32 id => SolverNet.Deposit) internal _orderDeposit;

    /**
     * @notice Map order ID to call parameters.
     * @dev (target, selector, value, params)
     */
    mapping(bytes32 id => SolverNet.Call[]) internal _orderCalls;

    /**
     * @notice Map order ID to expense parameters.
     * @dev (spender, token, amount)
     */
    mapping(bytes32 id => SolverNet.TokenExpense[]) internal _orderExpenses;

    /**
     * @notice Map order ID to order parameters.
     */
    mapping(bytes32 id => OrderState) internal _orderState;

    /**
     * @notice Map order ID to order offset.
     */
    mapping(bytes32 id => uint248) internal _orderOffset;

    /**
     * @notice Map user address to onchain order nonce.
     */
    mapping(address user => uint256 nonce) internal _onchainUserNonce;

    /**
     * @notice Modifier to ensure contract functions are not paused.
     */
    modifier whenNotPaused(bytes32 pauseKey) {
        uint8 _pauseState = pauseState;
        if (_pauseState != NONE_PAUSED) {
            if (_pauseState == OPEN_PAUSED && pauseKey == OPEN) revert IsPaused();
            if (_pauseState == CLOSE_PAUSED && pauseKey == CLOSE) revert IsPaused();
            if (_pauseState == ALL_PAUSED) revert AllPaused();
        }
        _;
    }

    /**
     * @notice Constructor sets the OmniPortal, and Hyperlane Mailbox contract addresses.
     * @param omni_     Address of the OmniPortal.
     * @param mailbox_  Address of the Hyperlane Mailbox.
     */
    constructor(address omni_, address mailbox_) XAppBase(omni_) MailboxClient(mailbox_) { }

    /**
     * @notice Pause the `open` function, preventing new orders from being opened.
     * @dev Cannot override ALL_PAUSED state.
     * @param pause True to pause, false to unpause.
     */
    function pauseOpen(bool pause) external {
        _setPauseState(OPEN, pause);
    }

    /**
     * @notice Pause the `close` function, preventing orders from being closed by users.
     * @dev `close` should only be paused if the Omni Core relayer is not available.
     * @dev Cannot override ALL_PAUSED state.
     * @param pause True to pause, false to unpause.
     */
    function pauseClose(bool pause) external {
        _setPauseState(CLOSE, pause);
    }

    /**
     * @notice Pause open and close functions.
     * @dev Can override OPEN_PAUSED or CLOSE_PAUSED states.
     * @param pause True to pause, false to unpause.
     */
    function pauseAll(bool pause) external {
        pause ? pauseState = ALL_PAUSED : pauseState = NONE_PAUSED;
        emit Paused(OPEN, pause, pauseState);
        emit Paused(CLOSE, pause, pauseState);
    }

    /**
     * @notice Set the outbox addresses for the given chain IDs.
     * @param chainIds IDs of the chains.
     * @param outboxes Addresses of the outboxes.
     */
    function setOutboxes(uint64[] calldata chainIds, address[] calldata outboxes) external {
        if (chainIds.length != outboxes.length) revert InvalidArrayLength();
        for (uint256 i; i < chainIds.length; ++i) {
            _outboxes[chainIds[i]] = outboxes[i];
            emit OutboxSet(chainIds[i], outboxes[i]);
        }
    }

    /**
     * @notice Returns the outbox address for the given chain ID.
     * @param chainId ID of the chain.
     * @return outbox Outbox address.
     */
    function getOutbox(uint64 chainId) external view returns (address) {
        return _outboxes[chainId];
    }

    /**
     * @notice Returns the order, its state, and offset with the given ID.
     * @param id ID of the order.
     */
    function getOrder(bytes32 id)
        external
        view
        returns (ResolvedCrossChainOrder memory resolved, OrderState memory state, uint248 offset)
    {
        SolverNet.Order memory order = _getOrder(id);
        return (_resolve({ order: order, id: id, openDeadline: 0 }), _orderState[id], _orderOffset[id]);
    }

    /**
     * @notice Returns the order ID for the given user and nonce.
     * @param user  Address of the user.
     * @param nonce Nonce of the order.
     * @param gasless Whether the order is gasless.
     */
    function getOrderId(address user, uint256 nonce, bool gasless) external view returns (bytes32) {
        return _getOrderId({ user: user, nonce: nonce, gasless: gasless });
    }

    /**
     * @notice Returns the next onchain order ID for the given user.
     * @param user Address of the user the order is opened for.
     */
    function getNextOnchainOrderId(address user) external view returns (bytes32) {
        return _getOrderId({ user: user, nonce: _onchainUserNonce[user], gasless: false });
    }

    /**
     * @notice Returns the onchain nonce for the given user.
     * @param user Address of the user the order is opened for.
     */
    function getOnchainUserNonce(address user) external view returns (uint256) {
        return _onchainUserNonce[user];
    }

    /**
     * @notice Returns the order offset of the latest order opened at this inbox.
     */
    function getLatestOrderOffset() external view returns (uint248) {
        return _offset;
    }

    /**
     * @dev Validate the onchain order.
     * @param order OnchainCrossChainOrder to validate.
     */
    function validate(OnchainCrossChainOrder calldata order) external view returns (bool) {
        _validate(order);
        return true;
    }

    /**
     * @dev Validate the gasless order.
     * @param order GaslessCrossChainOrder to validate.
     */
    function validateFor(GaslessCrossChainOrder calldata order) external view returns (bool) {
        _validateFor(order);
        return true;
    }

    /**
     * @notice Resolve the onchain order with validation.
     * @param order OnchainCrossChainOrder to resolve.
     */
    function resolve(OnchainCrossChainOrder calldata order) public view returns (ResolvedCrossChainOrder memory) {
        SolverNet.Order memory _order = _validate(order);
        address user = _order.header.owner;
        return _resolve({
            order: _order,
            id: _getOrderId({ user: user, nonce: _onchainUserNonce[user], gasless: false }),
            openDeadline: 0
        });
    }

    /**
     * @notice Resolve the gasless order with validation.
     * @dev The `bytes calldata originFillerData` param is currently unused.
     * @param order GaslessCrossChainOrder to resolve.
     */
    function resolveFor(GaslessCrossChainOrder calldata order, bytes calldata)
        public
        view
        returns (ResolvedCrossChainOrder memory)
    {
        (, SolverNet.Order memory _order) = _validateFor(order);
        address user = order.user;
        return _resolve({
            order: _order,
            id: _getOrderId({ user: user, nonce: order.nonce, gasless: true }),
            openDeadline: order.openDeadline
        });
    }

    /**
     * @notice Open an order to execute a call on another chain, backed by deposits.
     * @dev Token deposits are transferred from msg.sender to this inbox.
     * @param order OnchainCrossChainOrder to open.
     */
    function open(OnchainCrossChainOrder calldata order) external payable whenNotPaused(OPEN) nonReentrant {
        SolverNet.Order memory _order = _validate(order);
        address user = _order.header.owner;
        bytes32 id = _getOrderId({ user: user, nonce: _onchainUserNonce[user]++, gasless: false });

        _onchainDeposit(_order.deposit);
        _openOrder({ order: _order, id: id, openDeadline: 0 });
    }

    /**
     * @notice Open a gasless order to execute a call on another chain, backed by deposits.
     * @dev The `bytes calldata originFillerData` param is currently unused.
     * @dev Token deposits are transferred from order.user to this inbox, native deposits are not supported.
     * @param order GaslessCrossChainOrder to open.
     * @param signature Signature from order.user.
     */
    function openFor(GaslessCrossChainOrder calldata order, bytes calldata signature, bytes calldata)
        external
        whenNotPaused(OPEN)
        nonReentrant
    {
        (SolverNet.OrderData memory orderData, SolverNet.Order memory _order) = _validateFor(order);
        address user = order.user;
        bytes32 id = _getOrderId({ user: user, nonce: order.nonce, gasless: true });

        _gaslessDeposit(order, orderData, signature);
        _openOrder({ order: _order, id: id, openDeadline: order.openDeadline });
    }

    /**
     * @notice Reject an open order and refund deposits.
     * @dev Only a whitelisted solver can reject.
     * @param id     ID of the order.
     * @param reason Reason code for rejection.
     */
    function reject(bytes32 id, uint8 reason) external nonReentrant {
        OrderState memory state = _orderState[id];

        if (reason == 0) revert InvalidReason();
        if (state.status != Status.Pending) revert OrderNotPending();

        _upsertOrder({ id: id, status: Status.Rejected, rejectReason: reason, updatedBy: msg.sender });
        _transferDeposit({ id: id, to: _orderHeader[id].owner });
        _purgeState({ id: id, status: Status.Rejected });

        emit Rejected(id, msg.sender, reason);
    }

    /**
     * @notice Close order and refund deposits after fill deadline has elapsed.
     * @dev Only order initiator can close.
     * @param id ID of the order.
     */
    function close(bytes32 id) external whenNotPaused(CLOSE) nonReentrant {
        OrderState memory state = _orderState[id];
        SolverNet.Header memory header = _orderHeader[id];

        uint256 buffer = header.destChainId == block.chainid ? 0 : CLOSE_BUFFER;
        if (state.status != Status.Pending) revert OrderNotPending();
        if (header.owner != msg.sender) revert Unauthorized();
        if (header.fillDeadline + buffer >= block.timestamp) revert OrderStillValid();

        _upsertOrder({ id: id, status: Status.Closed, rejectReason: 0, updatedBy: msg.sender });
        _transferDeposit({ id: id, to: header.owner });
        _purgeState({ id: id, status: Status.Closed });

        emit Closed(id);
    }

    /**
     * @notice Fill an order via Omni Core.
     * @dev Only callable by the outbox.
     * @param id         ID of the order.
     * @param fillHash   Hash of fill instructions origin data.
     * @param creditedTo Address deposits are credited to, provided by the filler.
     */
    function markFilled(bytes32 id, bytes32 fillHash, address creditedTo) external xrecv {
        uint64 origin = xmsg.sourceChainId == 0 ? uint64(block.chainid) : xmsg.sourceChainId;
        address sender = xmsg.sender == address(0) ? msg.sender : xmsg.sender;
        _markFilled({ id: id, fillHash: fillHash, creditedTo: creditedTo, origin: origin, sender: sender });
    }

    /**
     * @notice Fill an order via Hyperlane.
     * @dev Only callable by the outbox.
     * @param origin  The origin domain
     * @param sender  The sender address
     * @param message The message
     */
    function handle(uint32 origin, bytes32 sender, bytes calldata message) external payable override onlyMailbox {
        (bytes32 id, bytes32 fillHash, address creditedTo) = abi.decode(message, (bytes32, bytes32, address));
        _markFilled({ id: id, fillHash: fillHash, creditedTo: creditedTo, origin: origin, sender: sender.toAddress() });
    }

    /**
     * @notice Claim deposits for a filled order.
     * @param id ID of the order.
     * @param to Address to send deposits to.
     */
    function claim(bytes32 id, address to) external nonReentrant {
        OrderState memory state = _orderState[id];

        if (state.status != Status.Filled) revert OrderNotFilled();
        if (state.updatedBy != msg.sender) revert Unauthorized();

        _upsertOrder({ id: id, status: Status.Claimed, rejectReason: 0, updatedBy: msg.sender });
        _transferDeposit({ id: id, to: to });
        _purgeState({ id: id, status: Status.Claimed });

        emit Claimed(id, msg.sender, to);
    }

    /**
     * @dev Return the order for the given ID.
     * @param id ID of the order.
     */
    function _getOrder(bytes32 id) internal view returns (SolverNet.Order memory) {
        return SolverNet.Order({
            header: _orderHeader[id], calls: _orderCalls[id], deposit: _orderDeposit[id], expenses: _orderExpenses[id]
        });
    }

    /**
     * @dev Validate and parse OnchainCrossChainOrder parameters
     * @param order OnchainCrossChainOrder to validate
     */
    function _validate(OnchainCrossChainOrder calldata order) internal view returns (SolverNet.Order memory) {
        _validateOnchainOrder(order);
        (, SolverNet.Order memory _order) = _validateOrderData({
            orderDataBytes: order.orderData, fillDeadline: order.fillDeadline, user: address(0), gasless: false
        });
        return _order;
    }

    /**
     * @dev Validate and parse GaslessCrossChainOrder parameters
     * @param order GaslessCrossChainOrder to validate
     */
    function _validateFor(GaslessCrossChainOrder calldata order)
        internal
        view
        returns (SolverNet.OrderData memory, SolverNet.Order memory)
    {
        _validateGaslessOrder(order);
        return _validateOrderData({
            orderDataBytes: order.orderData, fillDeadline: order.fillDeadline, user: order.user, gasless: true
        });
    }

    /**
     * @dev Validate OnchainCrossChainOrder parameters
     * @param order OnchainCrossChainOrder to validate
     */
    function _validateOnchainOrder(OnchainCrossChainOrder calldata order) internal view {
        if (order.fillDeadline <= block.timestamp) revert InvalidFillDeadline();
        if (
            order.orderDataType != HashLib.OLD_ORDERDATA_TYPEHASH
                && order.orderDataType != HashLib.OMNIORDERDATA_TYPEHASH
        ) {
            revert InvalidOrderTypehash();
        }
        if (order.orderData.length == 0) revert InvalidOrderData();
    }

    /**
     * @dev Validate GaslessCrossChainOrder parameters
     * @param order GaslessCrossChainOrder to validate
     */
    function _validateGaslessOrder(GaslessCrossChainOrder calldata order) internal view {
        if (order.originSettler != address(this)) revert InvalidOriginSettler();
        if (order.user == address(0)) revert InvalidUser();
        if (order.nonce == 0) revert InvalidNonce();
        if (order.originChainId != block.chainid) revert InvalidOriginChainId();
        if (order.openDeadline < block.timestamp) revert InvalidOpenDeadline();
        if (order.fillDeadline <= order.openDeadline) revert InvalidFillDeadline();
        if (
            order.orderDataType != HashLib.OLD_ORDERDATA_TYPEHASH
                && order.orderDataType != HashLib.OMNIORDERDATA_TYPEHASH
        ) {
            revert InvalidOrderTypehash();
        }
        if (order.orderData.length == 0) revert InvalidOrderData();
    }

    /**
     * @dev Validate SolverNet.OrderData
     * @param orderDataBytes Undecoded SolverNet.OrderData to validate
     * @param fillDeadline Fill deadline of the order
     * @param user Address to override a missing order owner with
     * @param gasless Whether the order is gasless
     */
    function _validateOrderData(bytes calldata orderDataBytes, uint32 fillDeadline, address user, bool gasless)
        internal
        view
        returns (SolverNet.OrderData memory, SolverNet.Order memory)
    {
        SolverNet.OrderData memory orderData = abi.decode(orderDataBytes, (SolverNet.OrderData));

        // Validate SolverNet.OrderData.Header fields
        if (orderData.owner == address(0)) {
            if (user == address(0)) orderData.owner = msg.sender;
            else orderData.owner = user;
        }
        if (orderData.destChainId == 0) revert InvalidDestinationChainId();

        // Gasless orders do not support native deposits
        if (gasless) {
            if (orderData.deposit.token == address(0) && orderData.deposit.amount > 0) revert InvalidNativeDeposit();
        }

        SolverNet.Header memory header = SolverNet.Header({
            owner: orderData.owner, destChainId: orderData.destChainId, fillDeadline: fillDeadline
        });

        // Validate SolverNet.OrderData.Call
        SolverNet.Call[] memory calls = orderData.calls;
        if (calls.length == 0) revert InvalidMissingCalls();
        if (calls.length > MAX_ARRAY_SIZE) revert InvalidArrayLength();

        // Validate SolverNet.OrderData.Expenses
        SolverNet.TokenExpense[] memory expenses = orderData.expenses;
        if (expenses.length > MAX_ARRAY_SIZE) revert InvalidArrayLength();
        for (uint256 i; i < expenses.length; ++i) {
            if (expenses[i].token == address(0)) revert InvalidExpenseToken();
            if (expenses[i].amount == 0) revert InvalidExpenseAmount();
        }

        SolverNet.Order memory order =
            SolverNet.Order({ header: header, calls: calls, deposit: orderData.deposit, expenses: expenses });

        return (orderData, order);
    }

    /**
     * @dev Derive the maxSpent Output for the order.
     * @param order Order data to derive from.
     */
    function _deriveMaxSpent(SolverNet.Order memory order) internal pure returns (IERC7683.Output[] memory) {
        SolverNet.Header memory header = order.header;
        SolverNet.Call[] memory calls = order.calls;
        SolverNet.TokenExpense[] memory expenses = order.expenses;

        uint256 totalNativeValue;
        for (uint256 i; i < calls.length; ++i) {
            if (calls[i].value > 0) totalNativeValue += calls[i].value;
        }

        // maxSpent recipient field is unused in SolverNet, was previously set to outbox as a placeholder
        IERC7683.Output[] memory maxSpent =
            new IERC7683.Output[](totalNativeValue > 0 ? expenses.length + 1 : expenses.length);
        for (uint256 i; i < expenses.length; ++i) {
            maxSpent[i] = IERC7683.Output({
                token: expenses[i].token.toBytes32(),
                amount: expenses[i].amount,
                recipient: bytes32(0),
                chainId: header.destChainId
            });
        }
        if (totalNativeValue > 0) {
            maxSpent[expenses.length] = IERC7683.Output({
                token: bytes32(0), amount: totalNativeValue, recipient: bytes32(0), chainId: header.destChainId
            });
        }

        return maxSpent;
    }

    /**
     * @dev Derive the minReceived Output for the order.
     * @param order Order data to derive from.
     */
    function _deriveMinReceived(SolverNet.Order memory order) internal view returns (IERC7683.Output[] memory) {
        SolverNet.Deposit memory deposit = order.deposit;

        IERC7683.Output[] memory minReceived = new IERC7683.Output[](deposit.amount > 0 ? 1 : 0);
        if (deposit.amount > 0) {
            minReceived[0] = IERC7683.Output({
                token: deposit.token.toBytes32(), amount: deposit.amount, recipient: bytes32(0), chainId: block.chainid
            });
        }

        return minReceived;
    }

    /**
     * @dev Derive the fillInstructions for the order.
     * @param order Order data to derive from.
     */
    function _deriveFillInstructions(SolverNet.Order memory order)
        internal
        view
        returns (IERC7683.FillInstruction[] memory)
    {
        SolverNet.Header memory header = order.header;
        SolverNet.Call[] memory calls = order.calls;
        SolverNet.TokenExpense[] memory expenses = order.expenses;

        IERC7683.FillInstruction[] memory fillInstructions = new IERC7683.FillInstruction[](1);
        fillInstructions[0] = IERC7683.FillInstruction({
            destinationChainId: header.destChainId,
            destinationSettler: _outboxes[header.destChainId].toBytes32(),
            originData: abi.encode(
                SolverNet.FillOriginData({
                    srcChainId: uint64(block.chainid),
                    destChainId: header.destChainId,
                    fillDeadline: header.fillDeadline,
                    calls: calls,
                    expenses: expenses
                })
            )
        });

        return fillInstructions;
    }

    /**
     * @dev Resolve the order without validation.
     * @param order SolverNet order to resolve.
     * @param id ID of the order.
     * @param openDeadline Open deadline of the order.
     */
    function _resolve(SolverNet.Order memory order, bytes32 id, uint32 openDeadline)
        internal
        view
        returns (ResolvedCrossChainOrder memory)
    {
        SolverNet.Header memory header = order.header;

        IERC7683.Output[] memory maxSpent = _deriveMaxSpent(order);
        IERC7683.Output[] memory minReceived = _deriveMinReceived(order);
        IERC7683.FillInstruction[] memory fillInstructions = _deriveFillInstructions(order);

        return ResolvedCrossChainOrder({
            user: header.owner,
            originChainId: block.chainid,
            openDeadline: openDeadline,
            fillDeadline: header.fillDeadline,
            orderId: id,
            maxSpent: maxSpent,
            minReceived: minReceived,
            fillInstructions: fillInstructions
        });
    }

    /**
     * @notice Validate and intake an ERC20 or native deposit.
     * @param deposit Deposit to process.
     */
    function _onchainDeposit(SolverNet.Deposit memory deposit) internal {
        if (deposit.token == address(0)) {
            if (msg.value != deposit.amount) revert InvalidNativeDeposit();
        } else {
            uint256 balance = deposit.token.balanceOf(address(this));
            deposit.token.safeTransferFrom(msg.sender, address(this), deposit.amount);
            // If we received less tokens than expected (max transfer value override or fee on transfer), revert
            if (deposit.token.balanceOf(address(this)) < balance + deposit.amount) revert InvalidERC20Deposit();
        }
    }

    /**
     * @dev Validate the signature for the given gasless order and retrieve order deposits.
     * @param order GaslessCrossChainOrder to validate.
     * @param orderData Order data to validate.
     * @param signature Signature to validate.
     */
    function _gaslessDeposit(
        GaslessCrossChainOrder calldata order,
        SolverNet.OrderData memory orderData,
        bytes calldata signature
    ) internal {
        if (msg.value > 0) revert InvalidNativeDeposit();

        ISignatureTransfer.TokenPermissions memory perms =
            ISignatureTransfer.TokenPermissions({ token: orderData.deposit.token, amount: orderData.deposit.amount });
        ISignatureTransfer.PermitTransferFrom memory permit = ISignatureTransfer.PermitTransferFrom({
            permitted: perms, nonce: order.nonce, deadline: order.openDeadline
        });
        ISignatureTransfer.SignatureTransferDetails memory details = ISignatureTransfer.SignatureTransferDetails({
            to: address(this), requestedAmount: orderData.deposit.amount
        });

        uint256 balance = orderData.deposit.token.balanceOf(address(this));
        PERMIT2.permitWitnessTransferFrom(
            permit,
            details,
            order.user,
            HashLib.witnessHashCalldata(order, orderData),
            HashLib.PERMIT2_WITNESS_TYPE_STRING,
            signature
        );
        // If we received less tokens than expected (max transfer value override or fee on transfer), revert
        if (orderData.deposit.token.balanceOf(address(this)) < balance + orderData.deposit.amount) {
            revert InvalidERC20Deposit();
        }
    }

    /**
     * @dev Transfer deposit to recipient. Used for both refunds and claims.
     * @param id ID of the order.
     * @param to Address to send deposits to.
     */
    function _transferDeposit(bytes32 id, address to) internal {
        SolverNet.Deposit memory deposit = _orderDeposit[id];

        if (deposit.amount > 0) {
            if (deposit.token == address(0)) to.safeTransferETH(deposit.amount);
            else deposit.token.safeTransfer(to, deposit.amount);
        }
    }

    /**
     * @dev Opens a new order by initializing its state.
     * @param order SolverNet order to open.
     * @param id ID of the order.
     * @param openDeadline Open deadline of the order.
     */
    function _openOrder(SolverNet.Order memory order, bytes32 id, uint32 openDeadline) internal {
        _orderHeader[id] = order.header;
        _orderDeposit[id] = order.deposit;
        _orderOffset[id] = _incrementOffset();
        for (uint256 i; i < order.calls.length; ++i) {
            _orderCalls[id].push(order.calls[i]);
        }
        for (uint256 i; i < order.expenses.length; ++i) {
            _orderExpenses[id].push(order.expenses[i]);
        }

        _upsertOrder({ id: id, status: Status.Pending, rejectReason: 0, updatedBy: msg.sender });
        ResolvedCrossChainOrder memory resolved = _resolve(order, id, openDeadline);
        SolverNet.FillOriginData memory fillOriginData =
            abi.decode(resolved.fillInstructions[0].originData, (SolverNet.FillOriginData));

        emit FillOriginData(resolved.orderId, fillOriginData);
        emit Open(resolved.orderId, resolved);
    }

    /**
     * @dev Mark an order as filled.
     * @param id ID of the order.
     * @param fillHash Hash of fill instructions origin data.
     * @param creditedTo Address deposits are credited to, provided by the filler.
     * @param origin Origin chain ID.
     * @param sender Remote sender address.
     */
    function _markFilled(bytes32 id, bytes32 fillHash, address creditedTo, uint64 origin, address sender) internal {
        SolverNet.Header memory header = _orderHeader[id];
        OrderState memory state = _orderState[id];

        if (state.status != Status.Pending) revert OrderNotPending();
        if (origin != header.destChainId) revert WrongSourceChain();
        if (sender != _outboxes[origin]) revert Unauthorized();

        // Ensure reported fill hash matches origin data
        // We are temporarily supporting both so no in-flight order settlements are rejected during the transition
        (bytes32 fillhash, SolverNet.FillOriginData memory fillOriginData) = _deprecatedFillHash(id);
        if (fillHash != fillhash) {
            if (fillHash != _fillHash(id, fillOriginData)) {
                revert WrongFillHash();
            }
        }

        _upsertOrder(id, Status.Filled, 0, creditedTo);
        _purgeState(id, Status.Filled);

        emit Filled(id, fillHash, creditedTo);
    }

    /**
     * @dev Update or insert order state by id.
     * @param id           ID of the order.
     * @param status       Status to upsert.
     * @param rejectReason Reason code for rejecting the order, if rejected.
     * @param updatedBy    Address updating the order.
     */
    function _upsertOrder(bytes32 id, Status status, uint8 rejectReason, address updatedBy) internal {
        uint8 _rejectReason = _orderState[id].rejectReason;
        _orderState[id] = OrderState({
            status: status,
            rejectReason: rejectReason > 0 ? rejectReason : _rejectReason,
            timestamp: uint32(block.timestamp),
            updatedBy: updatedBy
        });
    }

    /**
     * @dev Purge order state after it is no longer needed.
     * @param id     ID of the order.
     * @param status Status of the order.
     */
    function _purgeState(bytes32 id, Status status) internal {
        if (status == Status.Pending) return;
        if (status != Status.Filled) delete _orderDeposit[id];
        if (status != Status.Claimed) {
            delete _orderHeader[id];
            delete _orderCalls[id];
            delete _orderExpenses[id];
        }
    }

    /**
     * @dev Derive order ID from user and nonce.
     * @param user  Address of the user.
     * @param nonce Nonce of the order.
     * @param gasless Whether the order is gasless.
     */
    function _getOrderId(address user, uint256 nonce, bool gasless) internal view returns (bytes32) {
        return keccak256(abi.encode(user, nonce, gasless, block.chainid));
    }

    /**
     * @dev Increment and return the next order offset.
     */
    function _incrementOffset() internal returns (uint248) {
        return ++_offset;
    }

    /**
     * @dev Returns old fill hash. Used to discern fulfillment.
     * @param orderId ID of the order.
     */
    function _deprecatedFillHash(bytes32 orderId) internal view returns (bytes32, SolverNet.FillOriginData memory) {
        SolverNet.Header memory header = _orderHeader[orderId];
        SolverNet.Call[] memory calls = _orderCalls[orderId];
        SolverNet.TokenExpense[] memory expenses = _orderExpenses[orderId];

        SolverNet.FillOriginData memory fillOriginData = SolverNet.FillOriginData({
            srcChainId: uint64(block.chainid),
            destChainId: header.destChainId,
            fillDeadline: header.fillDeadline,
            calls: calls,
            expenses: expenses
        });

        return (keccak256(abi.encode(orderId, abi.encode(fillOriginData))), fillOriginData);
    }

    /**
     * @dev Return fill hash without unnecessary double encoding.
     *      We will transition to this hash method, where this will eventually incorporate state logic from _deprecatedFillHash.
     * @param orderId ID of the order.
     * @param fillOriginData Fill origin data to hash. (prevents having to derive it again)
     */
    function _fillHash(bytes32 orderId, SolverNet.FillOriginData memory fillOriginData)
        internal
        pure
        returns (bytes32)
    {
        return keccak256(abi.encode(orderId, fillOriginData));
    }

    /**
     * @notice Pause the `open` or `close` function
     * @dev Cannot override ALL_PAUSED state
     * @param key OPEN or CLOSE pause key
     * @param pause True to pause, false to unpause
     */
    function _setPauseState(bytes32 key, bool pause) internal {
        uint8 _pauseState = pauseState;
        if (_pauseState == ALL_PAUSED) revert AllPaused();

        uint8 targetState = key == OPEN ? OPEN_PAUSED : CLOSE_PAUSED;
        if (pause ? _pauseState == targetState : _pauseState != targetState) revert IsPaused();

        pauseState = pause ? targetState : NONE_PAUSED;

        emit Paused(key, pause, pauseState);
    }
}
