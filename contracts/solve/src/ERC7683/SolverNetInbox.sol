// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.24;

import { OwnableRoles } from "solady/src/auth/OwnableRoles.sol";
import { ReentrancyGuard } from "solady/src/utils/ReentrancyGuard.sol";
import { Initializable } from "solady/src/utils/Initializable.sol";
import { XAppBase } from "core/src/pkg/XAppBase.sol";
import { SafeTransferLib } from "solady/src/utils/SafeTransferLib.sol";
import { DeployedAt } from "src/util/DeployedAt.sol";
import { ISolverNetInbox } from "./interfaces/ISolverNetInbox.sol";

/**
 * @title SolverNetInbox
 * @notice Entrypoint and alt-mempoool for user solve orders.
 */
contract SolverNetInbox is OwnableRoles, ReentrancyGuard, Initializable, DeployedAt, XAppBase, ISolverNetInbox {
    using SafeTransferLib for address;

    /**
     * @notice Role for solvers.
     * @dev _ROLE_0 evaluates to '1'.
     */
    uint256 internal constant SOLVER = _ROLE_0;

    /**
     * @notice Typehash for the order data.
     */
    bytes32 internal constant ORDER_DATA_TYPEHASH = keccak256(
        "SolverNetIntent(uint64 srcChainId,uint64 destChainId,TokenPrereq[] tokenPrereqs,Call call)TokenPrereq(bytes32 token,bytes32 spender,uint256 amount)Call(bytes32 target,uint256 value,bytes data)"
    ); // Not really needed until we support more than one order type or gasless orders

    /**
     * @dev Counter for generating unique order IDs. Incremented each time a new order is created.
     */
    uint256 internal _lastId;

    /**
     * @notice Address of the outbox contract.
     */
    address internal _outbox;

    /**
     * @notice Map order ID to resolved onchain order.
     */
    mapping(bytes32 id => ResolvedCrossChainOrder) internal _orders;

    /**
     * @notice Map order ID to order parameters.
     */
    mapping(bytes32 id => OrderState) internal _orderState;

    /**
     * @notice Map order ID to order history.
     */
    mapping(bytes32 id => StatusUpdate[]) internal _orderHistory;

    /**
     * @notice Map status to latest order ID.
     */
    mapping(Status => bytes32 id) internal _latestOrderIdByStatus;

    constructor() {
        _disableInitializers();
    }

    /**
     * @notice Initialize the contract's owner and solver.
     * @dev Used instead of constructor as we want to use the transparent upgradeable proxy pattern.
     * @param owner_  Address of the owner.
     * @param solver_ Address of the solver.
     */
    function initialize(address owner_, address solver_, address omni_, address outbox_) external initializer {
        _initializeOwner(owner_);
        _grantRoles(solver_, SOLVER);
        _setOmniPortal(omni_);
        _outbox = outbox_;
    }

    /**
     * @notice Returns the order with the given ID.
     * @param id  ID of the order.
     */
    function getOrder(bytes32 id)
        external
        view
        returns (ResolvedCrossChainOrder memory order, OrderState memory state, StatusUpdate[] memory history)
    {
        return (_orders[id], _orderState[id], _orderHistory[id]);
    }

    /**
     * @notice Returns the latest order with the given status.
     * @param status  Order status to query.
     */
    function getLatestOrderIdByStatus(Status status) external view returns (bytes32) {
        return _latestOrderIdByStatus[status];
    }

    /**
     * @dev Validate the onchain order.
     * @param order  OnchainCrossChainOrder to validate.
     */
    function validateOrder(OnchainCrossChainOrder calldata order) external view returns (bool) {
        _validateOrder(order);
        return true;
    }

    /**
     * @dev Resolve the onchain order.
     * @param order  OnchainCrossChainOrder to resolve.
     */
    function resolve(OnchainCrossChainOrder calldata order) public view returns (ResolvedCrossChainOrder memory) {
        SolverNetOrderData memory orderData = abi.decode(order.orderData, (SolverNetOrderData));
        SolverNetIntent memory intent = orderData.intent;
        TokenPrereq[] memory prereqs = intent.tokenPrereqs;
        TokenDeposit[] memory deposits = orderData.deposits;

        Output[] memory maxSpent = new Output[](prereqs.length);
        for (uint256 i; i < prereqs.length; ++i) {
            maxSpent[i] = Output({
                token: prereqs[i].token,
                amount: prereqs[i].amount,
                recipient: _addressToBytes32(_outbox),
                chainId: intent.destChainId
            });
        }

        Output[] memory minReceived = new Output[](deposits.length);
        for (uint256 i; i < deposits.length; ++i) {
            minReceived[i] = Output({
                token: _addressToBytes32(deposits[i].token),
                amount: deposits[i].amount,
                recipient: bytes32(0),
                chainId: block.chainid
            });
        }

        FillInstruction[] memory fillInstructions = new FillInstruction[](1);
        fillInstructions[0] = FillInstruction({
            destinationChainId: intent.destChainId,
            destinationSettler: _addressToBytes32(_outbox),
            originData: abi.encode(intent)
        });

        return ResolvedCrossChainOrder({
            user: msg.sender,
            originChainId: block.chainid,
            openDeadline: uint32(block.timestamp),
            fillDeadline: order.fillDeadline,
            orderId: _nextId(),
            maxSpent: maxSpent,
            minReceived: minReceived,
            fillInstructions: fillInstructions
        });
    }

    /**
     * @notice Open an order to execute a call on another chain, backed by deposits.
     * @dev Token deposits are transferred from msg.sender to this inbox.
     * @param order OnchainCrossChainOrder to open.
     */
    function open(OnchainCrossChainOrder calldata order) external payable nonReentrant {
        SolverNetOrderData memory orderData = _validateOrder(order);

        _processDeposits(orderData.deposits);

        ResolvedCrossChainOrder storage resolved = _openOrder(order);

        emit Open(resolved.orderId, resolved);
    }

    /**
     * @notice Accept an open order.
     * @dev Only a whitelisted solver can accept.
     * @param id  ID of the order.
     */
    function accept(bytes32 id) external onlyRoles(SOLVER) nonReentrant {
        OrderState memory state = _orderState[id];
        if (state.status != Status.Pending) revert NotPending();

        state.status = Status.Accepted;
        state.acceptedBy = msg.sender;
        _upsertOrder(id, state);

        emit Accepted(id, msg.sender);
    }

    /**
     * @notice Reject an open order.
     * @dev Only a whitelisted solver can reject.
     * @param id      ID of the order.
     * @param reason  Reason code for rejection.
     */
    function reject(bytes32 id, uint8 reason) external onlyRoles(SOLVER) nonReentrant {
        OrderState memory state = _orderState[id];
        if (state.status != Status.Pending) revert NotPending();

        state.status = Status.Rejected;
        _upsertOrder(id, state);

        emit Rejected(id, msg.sender, reason);
    }

    /**
     * @notice Cancel an open or rejected order and refund deposits.
     * @dev Only order initiator can cancel.
     * @param id  ID of the order.
     */
    function cancel(bytes32 id) external nonReentrant {
        ResolvedCrossChainOrder memory order = _orders[id];
        OrderState memory state = _orderState[id];
        if (state.status != Status.Pending && state.status != Status.Rejected) {
            revert NotPendingOrRejected();
        }
        if (order.user != msg.sender) revert Unauthorized();

        state.status = Status.Reverted;
        _upsertOrder(id, state);
        _transferDeposits(order.user, order.minReceived);

        emit Reverted(id);
    }

    /**
     * @notice Fill an order.
     * @dev Only callable by the outbox.
     * @param id        ID of the order.
     * @param fillHash  Hash of fill instructions origin data.
     */
    function markFilled(bytes32 id, bytes32 fillHash) external xrecv nonReentrant {
        ResolvedCrossChainOrder memory order = _orders[id];
        OrderState memory state = _orderState[id];
        if (state.status != Status.Accepted) revert NotAccepted();
        if (xmsg.sender != _outbox) revert NotOutbox();
        if (xmsg.sourceChainId != order.fillInstructions[0].destinationChainId) revert WrongSourceChain();

        // Ensure reported fill hash matches origin data
        if (fillHash != _fillHash(id, order.fillInstructions[0].originData)) {
            revert WrongCallHash();
        }

        state.status = Status.Filled;
        _upsertOrder(id, state);

        emit Filled(id, fillHash, state.acceptedBy);
    }

    /**
     * @notice Claim a filled order.
     * @param id  ID of the order.
     * @param to  Address to send deposits to.
     */
    function claim(bytes32 id, address to) external nonReentrant {
        ResolvedCrossChainOrder memory order = _orders[id];
        OrderState memory state = _orderState[id];
        if (state.status != Status.Filled) revert NotFilled();
        if (state.acceptedBy != msg.sender) revert Unauthorized();

        state.status = Status.Claimed;

        _upsertOrder(id, state);
        _transferDeposits(to, order.minReceived);

        emit Claimed(id, msg.sender, to, order.minReceived);
    }

    /**
     * @dev Validate all order fields.
     * @param order  OnchainCrossChainOrder to validate.
     */
    function _validateOrder(OnchainCrossChainOrder calldata order) internal view returns (SolverNetOrderData memory) {
        if (order.fillDeadline < block.timestamp) revert InvalidFillDeadline();
        if (order.orderDataType != ORDER_DATA_TYPEHASH) revert InvalidOrderDataTypehash();
        if (order.orderData.length == 0) revert InvalidOrderData();

        SolverNetOrderData memory orderData = abi.decode(order.orderData, (SolverNetOrderData));
        SolverNetIntent memory intent = orderData.intent;

        if (intent.srcChainId != block.chainid) revert InvalidSrcChain();
        // We should perform a chainId => outbox address lookup here in the future to validate the route
        if (intent.destChainId == 0 || intent.destChainId == block.chainid) revert InvalidDestChain();
        if (intent.call.target == bytes32(0)) revert ZeroAddress();
        if (intent.call.callData.length == 0) revert NoCalldata();
        if (orderData.deposits.length == 0) revert NoDeposits(); // Should we prevent requests without deposits?
        for (uint256 i; i < intent.tokenPrereqs.length; ++i) {
            TokenPrereq memory prereq = intent.tokenPrereqs[i];
            if (prereq.token != bytes32(0) && prereq.spender == bytes32(0)) revert NoSpender();
            if (prereq.amount == 0) revert ZeroAmount();
        }

        return orderData;
    }

    /**
     * @dev Process and validate all deposits.
     * @dev Native deposit is validated by checking msg.value against deposit amount, and must be included in array.
     * @param deposits  Array of TokenDeposit to process.
     */
    function _processDeposits(TokenDeposit[] memory deposits) internal {
        bool nativeDepositValidated = msg.value > 0 ? false : true;
        for (uint256 i; i < deposits.length; ++i) {
            TokenDeposit memory deposit = deposits[i];
            // Handle native deposit
            if (deposit.token == address(0)) {
                if (nativeDepositValidated) revert InvalidNativeDeposit();
                if (deposit.amount != msg.value) revert InvalidNativeDeposit();
                nativeDepositValidated = true;
            }
            // Handle ERC20 deposit
            if (deposit.token != address(0)) {
                if (deposit.amount == 0) revert ZeroAmount();
                deposit.token.safeTransferFrom(msg.sender, address(this), deposit.amount);
            }
        }
        // Validate frontend properly processed native deposit
        if (!nativeDepositValidated) revert InvalidNativeDeposit();
    }

    /**
     * @dev Opens a new order and initializes its state.
     * @param order The cross-chain order to open.
     * @return resolved The storage reference to the newly created order.
     */
    function _openOrder(OnchainCrossChainOrder calldata order)
        internal
        returns (ResolvedCrossChainOrder storage resolved)
    {
        bytes32 id = _incrementId();

        ResolvedCrossChainOrder memory _resolved = resolve(order);
        resolved = _orders[id];
        resolved.user = _resolved.user;
        resolved.originChainId = _resolved.originChainId;
        resolved.openDeadline = _resolved.openDeadline;
        resolved.fillDeadline = _resolved.fillDeadline;
        resolved.orderId = id;

        for (uint256 i; i < _resolved.maxSpent.length; ++i) {
            resolved.maxSpent.push(_resolved.maxSpent[i]);
        }

        for (uint256 i; i < _resolved.minReceived.length; ++i) {
            resolved.minReceived.push(_resolved.minReceived[i]);
        }

        for (uint256 i; i < _resolved.fillInstructions.length; ++i) {
            resolved.fillInstructions.push(_resolved.fillInstructions[i]);
        }

        _upsertOrder(id, OrderState({ status: Status.Pending, acceptedBy: address(0) }));

        return resolved;
    }

    /**
     * @dev Transfer deposits to recipient. Used for both refunds and claims.
     * @param to  Address to send deposits to.
     * @param deposits  Array of Output to transfer.
     */
    function _transferDeposits(address to, Output[] memory deposits) internal {
        if (to == address(0)) revert InvalidRecipient();

        for (uint256 i; i < deposits.length; ++i) {
            Output memory deposit = deposits[i];
            if (deposit.token == bytes32(0)) {
                to.safeTransferETH(deposit.amount);
            } else {
                address token = _bytes32ToAddress(deposit.token);
                token.safeTransferFrom(address(this), to, deposit.amount);
            }
        }
    }

    /**
     * @dev Return the next order ID.
     */
    function _nextId() internal view returns (bytes32) {
        return bytes32(_lastId + 1);
    }

    /**
     * @dev Increment and return the next order ID.
     */
    function _incrementId() internal returns (bytes32) {
        return bytes32(++_lastId);
    }

    /**
     * @dev Returns call hash. Used to discern fullfilment.
     * @param orderId      ID of the order.
     * @param originData   Encoded fill instruction origin data.
     */
    function _fillHash(bytes32 orderId, bytes memory originData) internal pure returns (bytes32) {
        return keccak256(abi.encode(orderId, originData));
    }

    /**
     * @dev Convert address to bytes32.
     * @param addr  Address to convert.
     */
    function _addressToBytes32(address addr) internal pure returns (bytes32) {
        return bytes32(uint256(uint160(addr)));
    }

    /**
     * @dev Convert bytes32 to address.
     * @param b  Bytes32 to convert.
     */
    function _bytes32ToAddress(bytes32 b) internal pure returns (address) {
        return address(uint160(uint256(b)));
    }

    /**
     * @dev Update or insert order state by id.
     */
    function _upsertOrder(bytes32 id, OrderState memory state) internal {
        _orderState[id] = state;
        _orderHistory[id].push(StatusUpdate({ status: state.status, timestamp: uint40(block.timestamp) }));
        _latestOrderIdByStatus[state.status] = id;
    }
}
