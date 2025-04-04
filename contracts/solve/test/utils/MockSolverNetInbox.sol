// SPDX-License-Identifier: GPL-3.0-only
pragma solidity ^0.8.24;

import { ReentrancyGuard } from "solady/src/utils/ReentrancyGuard.sol";
import { EIP712 } from "solady/src/utils/EIP712.sol";
import { DeployedAt } from "../../src/util/DeployedAt.sol";
import { XAppBase } from "core/src/pkg/XAppBase.sol";
import { IERC7683 } from "../../src/erc7683/IERC7683.sol";
import { ISolverNetInbox } from "../../src/interfaces/ISolverNetInbox.sol";
import { LibBytes } from "solady/src/utils/LibBytes.sol";
import { SafeTransferLib } from "solady/src/utils/SafeTransferLib.sol";
import { SignatureCheckerLib } from "solady/src/utils/SignatureCheckerLib.sol";
import { SolverNet } from "../../src/lib/SolverNet.sol";
import { AddrUtils } from "../../src/lib/AddrUtils.sol";
import { IPermit2 } from "../../src/ext/IPermit2.sol";
import { IOmniPortalPausable } from "core/src/interfaces/IOmniPortalPausable.sol";

/**
 * @title MockSolverNetInbox
 * @notice Entrypoint and alt-mempool for user solve orders.
 * @dev This mock implementation is ownerless and is not initialized.
 */
contract MockSolverNetInbox is ReentrancyGuard, EIP712, DeployedAt, XAppBase, ISolverNetInbox {
    using SafeTransferLib for address;
    using AddrUtils for address;

    error Unauthorized();

    /**
     * @notice Maximum number of calls and expenses in an order.
     */
    uint8 internal constant MAX_ARRAY_SIZE = 32;

    /**
     * @notice Maximum allowed manual gasless nonce step.
     * @dev Prevents a user's nonce from being maxed out, blocking future orders.
     */
    uint16 internal constant MAX_NONCE_STEP = 10_000;

    /**
     * @notice Buffer for closing orders after fill deadline to give Omni Core relayer time to act.
     */
    uint256 internal constant CLOSE_BUFFER = 6 hours;

    /**
     * @notice Typehash for the OrderData struct.
     */
    bytes32 internal constant ORDERDATA_TYPEHASH = keccak256(
        "OrderData(address owner,uint64 destChainId,Deposit deposit,Call[] calls,TokenExpense[] expenses)Deposit(address token,uint96 amount)Call(address target,bytes4 selector,uint256 value,bytes params)TokenExpense(address spender,address token,uint96 amount)"
    );

    /**
     * @notice Typehash for the GaslessCrossChainOrder struct.
     */
    bytes32 internal constant GASLESS_ORDER_TYPEHASH = keccak256(
        "GaslessCrossChainOrder(address originSettler,address user,uint256 nonce,uint256 originChainId,uint32 openDeadline,uint32 fillDeadline,bytes32 orderDataType,bytes orderData)"
    );

    /**
     * @notice Typehash for the PermitDetails struct.
     */
    bytes32 internal constant PERMITDETAILS_TYPEHASH =
        keccak256("PermitDetails(address token,uint160 amount,uint48 expiration,uint48 nonce)");

    /**
     * @notice Typehash for the PermitSingle struct.
     */
    bytes32 internal constant PERMITSINGLE_TYPEHASH = keccak256(
        "PermitSingle(PermitDetails details,address spender,uint256 sigDeadline)PermitDetails(address token,uint160 amount,uint48 expiration,uint48 nonce)"
    );

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
     * @notice The canonical Permit2 contract.
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
     * @notice Map user to onchain nonce.
     */
    mapping(address user => uint256 nonce) internal _onchainUserNonce;

    /**
     * @notice Map user to gasless nonce.
     */
    mapping(address user => uint256 nonce) internal _gaslessUserNonce;

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
     * @notice Initialize the contract's OmniPortal.
     * @param omni_ Address of the OmniPortal.
     */
    constructor(address omni_) {
        _setOmniPortal(omni_);
    }

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
     * @notice Returns the order, its state, and offset with the given ID.
     * @param id ID of the order.
     */
    function getOrder(bytes32 id)
        external
        view
        returns (ResolvedCrossChainOrder memory resolved, OrderState memory state, uint248 offset)
    {
        SolverNet.Order memory orderData = _getOrder(id);
        return (_resolve(orderData, id, 0), _orderState[id], _orderOffset[id]);
    }

    /**
     * @notice Returns the order ID for the given user and nonce.
     * @param user  Address of the user.
     * @param nonce Nonce of the order.
     */
    function getOrderId(address user, uint256 nonce) external view returns (bytes32) {
        return _getOrderId(user, nonce);
    }

    /**
     * @notice Returns the next onchain order ID for the given user.
     * @param user Address of the user the order is opened for.
     */
    function getNextOnchainOrderId(address user) external view returns (bytes32) {
        return _getOrderId(user, _onchainUserNonce[user]);
    }

    /**
     * @notice Returns the next gasless order ID for the given user.
     * @param user Address of the user paying for the order.
     */
    function getNextGaslessOrderId(address user) external view returns (bytes32) {
        return _getOrderId(user, _gaslessUserNonce[user]);
    }

    /**
     * @notice Returns the onchain nonce for the given user.
     * @param user Address of the user the order is opened for.
     */
    function getOnchainUserNonce(address user) external view returns (uint256) {
        return _onchainUserNonce[user];
    }

    /**
     * @notice Returns the gasless nonce for the given user.
     * @param user Address of the user paying for the order.
     */
    function getGaslessUserNonce(address user) external view returns (uint256) {
        return _gaslessUserNonce[user];
    }

    /**
     * @notice Returns the order offset of the latest order opened at this inbox.
     */
    function getLatestOrderOffset() external view returns (uint248) {
        return _offset;
    }

    /**
     * @notice Returns the EIP-712 digest for the given gasless order.
     * @param order GaslessCrossChainOrder being signed.
     * @return _ EIP-712 digest for the given gasless order.
     */
    function getGaslessCrossChainOrderDigest(GaslessCrossChainOrder calldata order) public view returns (bytes32) {
        bytes32 orderHash = keccak256(abi.encode(GASLESS_ORDER_TYPEHASH, order));
        return _hashTypedData(orderHash);
    }

    /**
     * @notice Returns the EIP-712 digest for the given Permit2 data.
     * @param order GaslessCrossChainOrder containing user, deposit token/amount, and deadline.
     * @return _ EIP-712 digest for the given Permit2 data.
     */
    function getPermit2Digest(GaslessCrossChainOrder calldata order) external view returns (bytes32) {
        SolverNet.Order memory orderData = _validateFor(order, LibBytes.emptyCalldata());
        address user = order.user;
        address token = orderData.deposit.token;
        address spender = address(this);
        uint160 amount = orderData.deposit.amount;
        uint48 expiration = type(uint48).max; // Set max value for expiration, consistent with SafeTransferLib's approach
        uint256 deadline = order.openDeadline;

        (,, uint48 nonce) = PERMIT2.allowance(user, token, spender);
        bytes32 permit2DomainSeparator = PERMIT2.DOMAIN_SEPARATOR();

        bytes32 permitDetailsHash = keccak256(abi.encode(PERMITDETAILS_TYPEHASH, token, amount, expiration, nonce));
        bytes32 permitSingleHash = keccak256(abi.encode(PERMITSINGLE_TYPEHASH, permitDetailsHash, spender, deadline));

        bytes32 digest = keccak256(abi.encodePacked("\x19\x01", permit2DomainSeparator, permitSingleHash));
        return digest;
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
     * @param originFillerData Permit2 data for the origin settler.
     */
    function validateFor(GaslessCrossChainOrder calldata order, bytes calldata originFillerData)
        external
        view
        returns (bool)
    {
        _validateFor(order, originFillerData);
        return true;
    }

    /**
     * @notice Resolve the onchain order with validation.
     * @param order OnchainCrossChainOrder to resolve.
     */
    function resolve(OnchainCrossChainOrder calldata order) public view returns (ResolvedCrossChainOrder memory) {
        SolverNet.Order memory orderData = _validate(order);
        address user = orderData.header.owner;
        return _resolve(orderData, _getOrderId(user, _onchainUserNonce[user]), 0);
    }

    /**
     * @notice Resolve the gasless order with validation.
     * @param order GaslessCrossChainOrder to resolve.
     * @param originFillerData Permit2 data for the origin settler.
     */
    function resolveFor(GaslessCrossChainOrder calldata order, bytes calldata originFillerData)
        public
        view
        returns (ResolvedCrossChainOrder memory)
    {
        SolverNet.Order memory orderData = _validateFor(order, originFillerData);
        address user = order.user;
        return _resolve(orderData, _getOrderId(user, _gaslessUserNonce[user]), order.openDeadline);
    }

    /**
     * @notice Open an order to execute a call on another chain, backed by deposits.
     * @dev Token deposits are transferred from msg.sender to this inbox.
     * @param order OnchainCrossChainOrder to open.
     */
    function open(OnchainCrossChainOrder calldata order) external payable whenNotPaused(OPEN) nonReentrant {
        SolverNet.Order memory orderData = _validate(order);
        address user = orderData.header.owner;
        bytes32 id = _getOrderId(user, _onchainUserNonce[user]++);

        _open(orderData, id, msg.sender, 0);
    }

    /**
     * @notice Open a gasless order to execute a call on another chain, backed by deposits.
     * @dev Token deposits are transferred from order.user to this inbox.
     * @param order GaslessCrossChainOrder to open.
     * @param signature Signature from order.user.
     * @param originFillerData Permit2 data for the origin settler.
     */
    function openFor(GaslessCrossChainOrder calldata order, bytes calldata signature, bytes calldata originFillerData)
        external
        whenNotPaused(OPEN)
        nonReentrant
    {
        address user = order.user;
        SolverNet.Order memory orderData = _validateFor(order, originFillerData);
        if (msg.sender != user) _validateSignature(order, signature);
        if (originFillerData.length > 0) _permit2(order, orderData.deposit, originFillerData);

        bytes32 id = _getOrderId(user, _gaslessUserNonce[user]++);
        _open(orderData, id, user, order.openDeadline);
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

        _upsertOrder(id, Status.Rejected, reason, msg.sender);
        _transferDeposit(id, _orderHeader[id].owner);
        _purgeState(id, Status.Rejected);

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

        if (state.status != Status.Pending) revert OrderNotPending();
        if (header.owner != msg.sender) revert Unauthorized();
        if (IOmniPortalPausable(address(omni)).isPaused(ACTION_XSUBMIT, header.destChainId)) revert PortalPaused();
        if (header.fillDeadline + CLOSE_BUFFER >= block.timestamp) revert OrderStillValid();

        _upsertOrder(id, Status.Closed, 0, msg.sender);
        _transferDeposit(id, header.owner);
        _purgeState(id, Status.Closed);

        emit Closed(id);
    }

    /**
     * @notice Fill an order.
     * @dev Only callable by the outbox.
     * @param id         ID of the order.
     * @param fillHash   Hash of fill instructions origin data.
     * @param creditedTo Address deposits are credited to, provided by the filler.
     */
    function markFilled(bytes32 id, bytes32 fillHash, address creditedTo) external xrecv {
        SolverNet.Header memory header = _orderHeader[id];
        OrderState memory state = _orderState[id];

        if (state.status != Status.Pending) revert OrderNotPending();
        if (xmsg.sourceChainId != header.destChainId) revert WrongSourceChain();
        if (xmsg.sender != _outboxes[xmsg.sourceChainId]) revert Unauthorized();

        // Ensure reported fill hash matches origin data
        if (fillHash != _fillHash(id)) {
            revert WrongFillHash();
        }

        _upsertOrder(id, Status.Filled, 0, creditedTo);
        _purgeState(id, Status.Filled);

        emit Filled(id, fillHash, creditedTo);
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

        _upsertOrder(id, Status.Claimed, 0, msg.sender);
        _transferDeposit(id, to);
        _purgeState(id, Status.Claimed);

        emit Claimed(id, msg.sender, to);
    }

    /**
     * @notice Increment the gasless nonce for the sender.
     * @param amount Amount to increment the nonce by.
     */
    function incrementGaslessNonce(uint16 amount) external {
        if (amount > MAX_NONCE_STEP) revert InvalidNonce();
        _gaslessUserNonce[msg.sender] += amount;
    }

    /**
     * @dev Return the order for the given ID.
     * @param id ID of the order.
     */
    function _getOrder(bytes32 id) internal view returns (SolverNet.Order memory) {
        return SolverNet.Order({
            header: _orderHeader[id],
            calls: _orderCalls[id],
            deposit: _orderDeposit[id],
            expenses: _orderExpenses[id]
        });
    }

    /**
     * @dev Validate and parse OnchainCrossChainOrder parameters
     * @param order OnchainCrossChainOrder to validate
     */
    function _validate(OnchainCrossChainOrder calldata order) internal view returns (SolverNet.Order memory) {
        _validateOnchainOrder(order);
        return _validateOrderData(order.orderData, order.fillDeadline);
    }

    /**
     * @dev Validate and parse GaslessCrossChainOrder parameters
     * @param order GaslessCrossChainOrder to validate
     * @param originFillerData Permit2 data for the origin settler
     */
    function _validateFor(GaslessCrossChainOrder calldata order, bytes calldata originFillerData)
        internal
        view
        returns (SolverNet.Order memory)
    {
        _validateGaslessOrder(order, originFillerData);
        return _validateOrderData(order.orderData, order.fillDeadline);
    }

    /**
     * @dev Validate OnchainCrossChainOrder parameters
     * @param order OnchainCrossChainOrder to validate
     */
    function _validateOnchainOrder(OnchainCrossChainOrder calldata order) internal view {
        if (order.fillDeadline <= block.timestamp) revert InvalidFillDeadline();
        if (order.orderDataType != ORDERDATA_TYPEHASH) revert InvalidOrderTypehash();
        if (order.orderData.length == 0) revert InvalidOrderData();
    }

    /**
     * @dev Validate GaslessCrossChainOrder parameters
     * @param order GaslessCrossChainOrder to validate
     * @param originFillerData Permit2 data for the origin settler
     */
    function _validateGaslessOrder(GaslessCrossChainOrder calldata order, bytes calldata originFillerData)
        internal
        view
    {
        uint256 gaslessUserNonce = _gaslessUserNonce[order.user];
        if (order.originSettler != address(this)) revert InvalidOriginSettler();
        if (order.user == address(0)) revert InvalidUser();
        if (order.nonce < gaslessUserNonce || order.nonce > gaslessUserNonce + MAX_NONCE_STEP) revert InvalidNonce();
        if (order.originChainId != block.chainid) revert InvalidOriginChainId();
        if (order.openDeadline < block.timestamp || order.openDeadline >= order.fillDeadline) {
            revert InvalidOpenDeadline();
        }
        if (order.fillDeadline <= block.timestamp) revert InvalidFillDeadline();
        if (order.orderDataType != ORDERDATA_TYPEHASH) revert InvalidOrderTypehash();
        if (order.orderData.length == 0) revert InvalidOrderData();
        if (originFillerData.length != 0 && originFillerData.length != 96) revert InvalidOriginFillerData();
    }

    /**
     * @dev Validate SolverNet.OrderData
     * @param orderDataBytes Undecoded SolverNet.OrderData to validate
     * @param fillDeadline Fill deadline of the order
     */
    function _validateOrderData(bytes calldata orderDataBytes, uint32 fillDeadline)
        internal
        view
        returns (SolverNet.Order memory)
    {
        SolverNet.OrderData memory orderData = abi.decode(orderDataBytes, (SolverNet.OrderData));

        // Validate SolverNet.OrderData.Header fields
        if (orderData.owner == address(0)) orderData.owner = msg.sender;
        if (orderData.destChainId == 0 || orderData.destChainId == block.chainid) revert InvalidDestinationChainId();

        SolverNet.Header memory header =
            SolverNet.Header({ owner: orderData.owner, destChainId: orderData.destChainId, fillDeadline: fillDeadline });

        // Validate SolverNet.OrderData.Call
        SolverNet.Call[] memory calls = orderData.calls;
        if (calls.length == 0) revert InvalidMissingCalls();
        if (calls.length > MAX_ARRAY_SIZE) revert InvalidArrayLength();
        for (uint256 i; i < calls.length; ++i) {
            SolverNet.Call memory call = calls[i];
            if (call.target == address(0)) revert InvalidCallTarget();
        }

        // Validate SolverNet.OrderData.Expenses
        SolverNet.TokenExpense[] memory expenses = orderData.expenses;
        if (expenses.length > MAX_ARRAY_SIZE) revert InvalidArrayLength();
        for (uint256 i; i < expenses.length; ++i) {
            if (expenses[i].token == address(0)) revert InvalidExpenseToken();
            if (expenses[i].amount == 0) revert InvalidExpenseAmount();
        }

        return SolverNet.Order({ header: header, calls: calls, deposit: orderData.deposit, expenses: expenses });
    }

    /**
     * @dev Validate the signature for the given gasless order.
     * @param order GaslessCrossChainOrder to validate.
     * @param signature Signature to validate.
     */
    function _validateSignature(GaslessCrossChainOrder calldata order, bytes calldata signature) internal view {
        bytes32 digest = getGaslessCrossChainOrderDigest(order);
        if (!SignatureCheckerLib.isValidSignatureNowCalldata(order.user, digest, signature)) {
            if (!SignatureCheckerLib.isValidERC1271SignatureNowCalldata(order.user, digest, signature)) {
                revert InvalidSignature();
            }
        }
    }

    /**
     * @dev Derive the maxSpent Output for the order.
     * @param orderData Order data to derive from.
     */
    function _deriveMaxSpent(SolverNet.Order memory orderData) internal view returns (IERC7683.Output[] memory) {
        SolverNet.Header memory header = orderData.header;
        SolverNet.Call[] memory calls = orderData.calls;
        SolverNet.TokenExpense[] memory expenses = orderData.expenses;

        uint256 totalNativeValue;
        for (uint256 i; i < calls.length; ++i) {
            if (calls[i].value > 0) totalNativeValue += calls[i].value;
        }

        IERC7683.Output[] memory maxSpent =
            new IERC7683.Output[](totalNativeValue > 0 ? expenses.length + 1 : expenses.length);
        for (uint256 i; i < expenses.length; ++i) {
            maxSpent[i] = IERC7683.Output({
                token: expenses[i].token.toBytes32(),
                amount: expenses[i].amount,
                recipient: _outboxes[header.destChainId].toBytes32(),
                chainId: header.destChainId
            });
        }
        if (totalNativeValue > 0) {
            maxSpent[expenses.length] = IERC7683.Output({
                token: bytes32(0),
                amount: totalNativeValue,
                recipient: _outboxes[header.destChainId].toBytes32(),
                chainId: header.destChainId
            });
        }

        return maxSpent;
    }

    /**
     * @dev Derive the minReceived Output for the order.
     * @param orderData Order data to derive from.
     */
    function _deriveMinReceived(SolverNet.Order memory orderData) internal view returns (IERC7683.Output[] memory) {
        SolverNet.Deposit memory deposit = orderData.deposit;

        IERC7683.Output[] memory minReceived = new IERC7683.Output[](deposit.amount > 0 ? 1 : 0);
        if (deposit.amount > 0) {
            minReceived[0] = IERC7683.Output({
                token: deposit.token.toBytes32(),
                amount: deposit.amount,
                recipient: bytes32(0),
                chainId: block.chainid
            });
        }

        return minReceived;
    }

    /**
     * @dev Derive the fillInstructions for the order.
     * @param orderData Order data to derive from.
     */
    function _deriveFillInstructions(SolverNet.Order memory orderData)
        internal
        view
        returns (IERC7683.FillInstruction[] memory)
    {
        SolverNet.Header memory header = orderData.header;
        SolverNet.Call[] memory calls = orderData.calls;
        SolverNet.TokenExpense[] memory expenses = orderData.expenses;

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
     * @param orderData Order data to resolve.
     * @param id ID of the order.
     * @param openDeadline Open deadline of the order.
     */
    function _resolve(SolverNet.Order memory orderData, bytes32 id, uint32 openDeadline)
        internal
        view
        returns (ResolvedCrossChainOrder memory)
    {
        SolverNet.Header memory header = orderData.header;

        IERC7683.Output[] memory maxSpent = _deriveMaxSpent(orderData);
        IERC7683.Output[] memory minReceived = _deriveMinReceived(orderData);
        IERC7683.FillInstruction[] memory fillInstructions = _deriveFillInstructions(orderData);

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
     * @dev Approve the origin settler to spend the deposit using Permit2.
     * @param order GaslessCrossChainOrder containing user and deadline.
     * @param deposit Deposit to approve.
     * @param originFillerData Permit2 data from the depositor.
     */
    function _permit2(
        GaslessCrossChainOrder calldata order,
        SolverNet.Deposit memory deposit,
        bytes calldata originFillerData
    ) internal {
        if (deposit.token == address(0)) return;
        (uint8 v, bytes32 r, bytes32 s) = abi.decode(originFillerData, (uint8, bytes32, bytes32));
        deposit.token.permit2(order.user, address(this), deposit.amount, order.openDeadline, v, r, s);
    }

    /**
     * @notice Validate and intake an ERC20 or native deposit.
     * @param deposit Deposit to process.
     * @param from Address to retrieve the deposit from.
     */
    function _processDeposit(SolverNet.Deposit memory deposit, address from) internal {
        if (deposit.token == address(0)) {
            if (msg.value != deposit.amount) revert InvalidNativeDeposit();
        } else {
            deposit.token.safeTransferFrom2(from, address(this), deposit.amount);
        }
    }

    /**
     * @dev Opens a new order by retrieving the deposit and initializing its state.
     * @param orderData Order data to open.
     * @param id ID of the order.
     * @param user Address of the user paying for the order.
     * @param openDeadline Open deadline of the order.
     */
    function _open(SolverNet.Order memory orderData, bytes32 id, address user, uint32 openDeadline) internal {
        _processDeposit(orderData.deposit, user);
        ResolvedCrossChainOrder memory resolved = _openOrder(orderData, id, openDeadline);

        emit Open(resolved.orderId, resolved);
        emit FillOriginData(
            resolved.orderId, abi.decode(resolved.fillInstructions[0].originData, (SolverNet.FillOriginData))
        );
    }

    /**
     * @dev Opens a new order by initializing its state.
     * @param orderData Order data to open.
     * @param id ID of the order.
     */
    function _openOrder(SolverNet.Order memory orderData, bytes32 id, uint32 openDeadline)
        internal
        returns (ResolvedCrossChainOrder memory resolved)
    {
        resolved = _resolve(orderData, id, openDeadline);
        _orderHeader[id] = orderData.header;
        _orderDeposit[id] = orderData.deposit;
        _orderOffset[id] = _incrementOffset();
        for (uint256 i; i < orderData.calls.length; ++i) {
            _orderCalls[id].push(orderData.calls[i]);
        }
        for (uint256 i; i < orderData.expenses.length; ++i) {
            _orderExpenses[id].push(orderData.expenses[i]);
        }

        _upsertOrder(id, Status.Pending, 0, msg.sender);

        return resolved;
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
     */
    function _getOrderId(address user, uint256 nonce) internal view returns (bytes32) {
        return keccak256(abi.encode(user, nonce, block.chainid));
    }

    /**
     * @dev Increment and return the next order offset.
     */
    function _incrementOffset() internal returns (uint248) {
        return ++_offset;
    }

    /**
     * @dev Returns call hash. Used to discern fulfillment.
     * @param orderId ID of the order.
     */
    function _fillHash(bytes32 orderId) internal view returns (bytes32) {
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

        return keccak256(abi.encode(orderId, abi.encode(fillOriginData)));
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
    }

    /**
     * @dev Returns the domain name and version for the EIP-712 domain separator.
     * @return name Domain name.
     * @return version Domain version.
     */
    function _domainNameAndVersion() internal pure override returns (string memory name, string memory version) {
        name = "MockSolverNetInbox";
        version = "1";
    }
}
