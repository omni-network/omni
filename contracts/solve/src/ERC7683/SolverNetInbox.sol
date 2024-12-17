// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.24;

import { OwnableRoles } from "solady/src/auth/OwnableRoles.sol";
import { ReentrancyGuard } from "solady/src/utils/ReentrancyGuard.sol";
import { Initializable } from "solady/src/utils/Initializable.sol";
import { XAppBase } from "core/src/pkg/XAppBase.sol";
import { SafeTransferLib } from "solady/src/utils/SafeTransferLib.sol";
import { ISolverNetInbox } from "./interfaces/ISolverNetInbox.sol";
import { IArbSys } from "../interfaces/IArbSys.sol";

/**
 * @title SolverNetInbox
 * @notice Entrypoint and alt-mempoool for user solve orders.
 */
contract SolverNetInbox is OwnableRoles, ReentrancyGuard, Initializable, XAppBase, ISolverNetInbox {
    using SafeTransferLib for address;

    /**
     * @notice Block number at which the contract was deployed.
     * @dev Uses L2 block number on Arbitrum, L1 block number elsewhere
     */
    uint256 public immutable deployedAt;

    /**
     * @notice Role for solvers.
     * @dev _ROLE_0 evaluates to '1'.
     */
    uint256 internal constant SOLVER = _ROLE_0;

    /**
     * @notice Arbitrum's ArbSys precompile (0x0000000000000000000000000000000000000064)
     * @dev Used to get Arbitrum block number.
     */
    address internal constant ARB_SYS = 0x0000000000000000000000000000000000000064;

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
    mapping(bytes32 id => SolverNetOrderParams) internal _orderParams;

    /**
     * @notice Map order ID to order history.
     */
    mapping(bytes32 id => StatusUpdate[]) internal _orderHistory;

    /**
     * @notice Map status to latest order ID.
     */
    mapping(Status => bytes32 id) internal _latestOrderIdByStatus;

    constructor() {
        // Must get Arbitrum block number from ArbSys precompile, block.number returns L1 block number on Arbitrum.
        if (_isContract(ARB_SYS)) {
            try IArbSys(ARB_SYS).arbBlockNumber() returns (uint256 arbBlockNumber) {
                deployedAt = arbBlockNumber;
            } catch {
                deployedAt = block.number;
            }
        } else {
            deployedAt = block.number;
        }

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
    function getOrder(bytes32 id) external view returns (SolverNetOrder memory) {
        return SolverNetOrder({ order: _orders[id], params: _orderParams[id], history: _orderHistory[id] });
    }

    /**
     * @notice Returns the order parameters for the given order ID.
     * @param id  ID of the order.
     */
    function getOrderParams(bytes32 id) external view returns (SolverNetOrderParams memory) {
        return _orderParams[id];
    }

    /**
     * @notice Returns the order history for the given order ID.
     * @param id  ID of the order.
     */
    function getOrderHistory(bytes32 id) external view returns (StatusUpdate[] memory) {
        return _orderHistory[id];
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

        ResolvedCrossChainOrder memory resolvedOrder = ResolvedCrossChainOrder({
            user: msg.sender,
            originChainId: block.chainid,
            openDeadline: uint32(block.timestamp),
            fillDeadline: order.fillDeadline,
            orderId: _nextId(),
            maxSpent: maxSpent,
            minReceived: minReceived,
            fillInstructions: fillInstructions
        });

        return resolvedOrder;
    }

    /**
     * @notice Open an order to execute a call on another chain, backed by deposits.
     * @dev Token deposits are transferred from msg.sender to this inbox.
     * @param order OnchainCrossChainOrder to open.
     */
    function open(OnchainCrossChainOrder calldata order) external payable nonReentrant {
        SolverNetOrderData memory orderData = _validateOrder(order);

        _processDeposits(orderData.deposits);

        ResolvedCrossChainOrder storage resolvedOrder = _openOrder(order);

        emit Open(resolvedOrder.orderId, resolvedOrder);
    }

    /**
     * @notice Accept an open order.
     * @dev Only a whitelisted solver can accept.
     * @param id  ID of the order.
     */
    function accept(bytes32 id) external onlyRoles(SOLVER) nonReentrant {
        SolverNetOrderParams memory orderParams = _orderParams[id];
        if (orderParams.status != Status.Pending) revert NotPending();

        orderParams.status = Status.Accepted;
        orderParams.updatedAt = uint40(block.timestamp);
        orderParams.acceptedBy = msg.sender;
        StatusUpdate memory statusUpdate = StatusUpdate({ status: Status.Accepted, timestamp: uint40(block.timestamp) });

        _orderParams[id] = orderParams;
        _orderHistory[id].push(statusUpdate);
        _latestOrderIdByStatus[Status.Accepted] = id;

        emit Accepted(id, msg.sender);
    }

    /**
     * @notice Reject an open order.
     * @dev Only a whitelisted solver can reject.
     * @param id      ID of the order.
     * @param reason  Reason code for rejection.
     */
    function reject(bytes32 id, uint8 reason) external onlyRoles(SOLVER) nonReentrant {
        SolverNetOrderParams memory orderParams = _orderParams[id];
        if (orderParams.status != Status.Pending) revert NotPending();

        orderParams.status = Status.Rejected;
        orderParams.updatedAt = uint40(block.timestamp);
        StatusUpdate memory statusUpdate = StatusUpdate({ status: Status.Rejected, timestamp: uint40(block.timestamp) });

        _orderParams[id] = orderParams;
        _orderHistory[id].push(statusUpdate);
        _latestOrderIdByStatus[Status.Rejected] = id;

        emit Rejected(id, msg.sender, reason);
    }

    /**
     * @notice Cancel an open or rejected order and refund deposits.
     * @dev Only order initiator can cancel.
     * @param id  ID of the order.
     */
    function cancel(bytes32 id) external nonReentrant {
        ResolvedCrossChainOrder memory order = _orders[id];
        SolverNetOrderParams memory orderParams = _orderParams[id];
        if (orderParams.status != Status.Pending && orderParams.status != Status.Rejected) {
            revert NotPendingOrRejected();
        }
        if (order.user != msg.sender) revert Unauthorized();

        orderParams.status = Status.Reverted;
        orderParams.updatedAt = uint40(block.timestamp);
        StatusUpdate memory statusUpdate = StatusUpdate({ status: Status.Reverted, timestamp: uint40(block.timestamp) });

        _orderParams[id] = orderParams;
        _orderHistory[id].push(statusUpdate);
        _latestOrderIdByStatus[Status.Reverted] = id;

        _transferDeposits(order.user, order.minReceived);

        emit Reverted(id);
    }

    /**
     * @notice Fulfill an order.
     * @dev Only callable by the outbox.
     * @param id        ID of the order.
     * @param callHash  Hash of the calls for this order executed on another chain.
     */
    function markFulfilled(bytes32 id, bytes32 callHash) external xrecv nonReentrant {
        ResolvedCrossChainOrder memory order = _orders[id];
        SolverNetOrderParams memory orderParams = _orderParams[id];
        if (orderParams.status != Status.Accepted) revert NotAccepted();
        if (xmsg.sender != _outbox) revert NotOutbox();
        if (xmsg.sourceChainId != order.fillInstructions[0].destinationChainId) revert WrongSourceChain();

        // Ensure reported call hash matches requested call hash
        if (callHash != _callHash(id, uint64(block.chainid), order.fillInstructions[0].originData)) {
            revert WrongCallHash();
        }

        orderParams.status = Status.Fulfilled;
        orderParams.updatedAt = uint40(block.timestamp);
        StatusUpdate memory statusUpdate =
            StatusUpdate({ status: Status.Fulfilled, timestamp: uint40(block.timestamp) });

        _orderParams[id] = orderParams;
        _orderHistory[id].push(statusUpdate);
        _latestOrderIdByStatus[Status.Fulfilled] = id;

        emit Fulfilled(id, callHash, orderParams.acceptedBy);
    }

    /**
     * @notice Claim a fulfilled order.
     * @param id  ID of the order.
     * @param to  Address to send deposits to.
     */
    function claim(bytes32 id, address to) external nonReentrant {
        ResolvedCrossChainOrder memory order = _orders[id];
        SolverNetOrderParams memory orderParams = _orderParams[id];
        if (orderParams.status != Status.Fulfilled) revert NotFulfilled();
        if (orderParams.acceptedBy != msg.sender) revert Unauthorized();

        orderParams.status = Status.Claimed;
        orderParams.updatedAt = uint40(block.timestamp);
        StatusUpdate memory statusUpdate = StatusUpdate({ status: Status.Claimed, timestamp: uint40(block.timestamp) });

        _orderParams[id] = orderParams;
        _orderHistory[id].push(statusUpdate);
        _latestOrderIdByStatus[Status.Claimed] = id;

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
     * @return resolvedOrder The storage reference to the newly created order.
     */
    function _openOrder(OnchainCrossChainOrder calldata order)
        internal
        returns (ResolvedCrossChainOrder storage resolvedOrder)
    {
        ResolvedCrossChainOrder memory _resolvedOrder = resolve(order);
        bytes32 orderId = _incrementId();
        resolvedOrder = _orders[orderId];

        resolvedOrder.user = _resolvedOrder.user;
        resolvedOrder.originChainId = _resolvedOrder.originChainId;
        resolvedOrder.openDeadline = _resolvedOrder.openDeadline;
        resolvedOrder.fillDeadline = _resolvedOrder.fillDeadline;
        resolvedOrder.orderId = orderId;

        for (uint256 i; i < resolvedOrder.maxSpent.length; ++i) {
            resolvedOrder.maxSpent.push(_resolvedOrder.maxSpent[i]);
        }

        for (uint256 i; i < resolvedOrder.minReceived.length; ++i) {
            resolvedOrder.minReceived.push(_resolvedOrder.minReceived[i]);
        }

        for (uint256 i; i < resolvedOrder.fillInstructions.length; ++i) {
            resolvedOrder.fillInstructions.push(_resolvedOrder.fillInstructions[i]);
        }

        _orderParams[orderId] =
            SolverNetOrderParams({ status: Status.Pending, updatedAt: uint40(block.timestamp), acceptedBy: address(0) });

        _orderHistory[orderId].push(StatusUpdate({ status: Status.Pending, timestamp: uint40(block.timestamp) }));

        _latestOrderIdByStatus[Status.Pending] = orderId;

        return resolvedOrder;
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
     * @param id          ID of the order.
     * @param srcChainId  Chain ID of the source chain.
     * @param orderData   Encoded order data.
     */
    function _callHash(bytes32 id, uint64 srcChainId, bytes memory orderData) internal pure returns (bytes32) {
        return keccak256(abi.encode(id, srcChainId, orderData));
    }

    /**
     * @dev Returns true if the address is a contract.
     * @param addr  Address to check.
     */
    function _isContract(address addr) internal view returns (bool) {
        uint32 size;
        assembly {
            size := extcodesize(addr)
        }
        return (size > 0);
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
}
