// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.24;

import { OwnableRoles } from "solady/src/auth/OwnableRoles.sol";
import { ReentrancyGuard } from "solady/src/utils/ReentrancyGuard.sol";
import { Initializable } from "solady/src/utils/Initializable.sol";
import { SafeTransferLib } from "solady/src/utils/SafeTransferLib.sol";
import { XAppBase } from "core/src/pkg/XAppBase.sol";
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
     * @dev uint repr of last assigned order ID.
     */
    uint256 internal _lastId;

    /**
     * @notice Address of the outbox contract.
     */
    address internal _outbox;

    /**
     * @notice Map order ID to request.
     */
    mapping(bytes32 id => Request) internal _requests;

    /**
     * @notice Map status to latest order ID.
     */
    mapping(Status => bytes32 id) internal _latestRequestByStatus;

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
     * @notice Returns the request with the given ID.
     */
    function getRequest(bytes32 id) external view returns (Request memory) {
        return _requests[id];
    }

    /**
     * @notice Returns the latest request with the given status.
     */
    function getLatestRequestByStatus(Status status) external view returns (Request memory) {
        return _requests[_latestRequestByStatus[status]];
    }

    /**
     * @notice Returns the update history for a request.
     */
    function getUpdateHistory(bytes32 id) external view returns (StatusUpdate[] memory) {
        return _requests[id].history;
    }

    /**
     * @dev Validate the onchain order.
     */
    function validateOnchainOrder(OnchainCrossChainOrder calldata order) external view returns (bool) {
        _validateOnchainOrder(order);
        return true;
    }

    /**
     * @dev Resolve the onchain order.
     */
    function resolve(OnchainCrossChainOrder calldata order) public view returns (ResolvedCrossChainOrder memory) {
        OrderData memory orderData = abi.decode(order.orderData, (OrderData));

        Output[] memory maxSpent = new Output[](orderData.prereqs.length);
        for (uint256 i; i < orderData.prereqs.length; ++i) {
            maxSpent[i] = Output({
                token: orderData.prereqs[i].token,
                amount: orderData.prereqs[i].amount,
                recipient: _addressToBytes32(_outbox),
                chainId: orderData.destChainId
            });
        }

        Output[] memory minReceived = new Output[](orderData.deposits.length);
        for (uint256 i; i < orderData.deposits.length; ++i) {
            minReceived[i] = Output({
                token: _addressToBytes32(orderData.deposits[i].token),
                amount: orderData.deposits[i].amount,
                recipient: bytes32(0),
                chainId: block.chainid
            });
        }

        FillOriginData memory fillOriginData = FillOriginData({
            srcChainId: uint64(block.chainid),
            destChainId: orderData.destChainId,
            calls: orderData.calls,
            prereqs: orderData.prereqs
        });
        FillInstruction[] memory fillInstructions = new FillInstruction[](1);
        fillInstructions[0] = FillInstruction({
            destinationChainId: orderData.destChainId,
            destinationSettler: _addressToBytes32(_outbox),
            originData: abi.encode(fillOriginData)
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
     * @notice Open a request to execute a call on another chain, backed by deposits.
     * @dev Token deposits are transferred from msg.sender to this inbox.
     * @param order OnchainCrossChainOrder to open.
     */
    function open(OnchainCrossChainOrder calldata order) external payable nonReentrant {
        OrderData memory orderData = _validateOnchainOrder(order);
        _validateTokenPrereqs(orderData.prereqs);
        _validateCalls(orderData.calls);

        _processDeposits(orderData.deposits);

        Request storage request = _openRequest(order);

        emit Open(request.order.orderId, request.order);
    }

    /**
     * @notice Accept an open request.
     * @dev Only a whitelisted solver can accept.
     * @param id  ID of the request.
     */
    function accept(bytes32 id) external onlyRoles(SOLVER) nonReentrant {
        Request storage request = _requests[id];
        if (request.status != Status.Pending) revert NotPending();

        request.status = Status.Accepted;
        request.updatedAt = uint40(block.timestamp);
        request.acceptedBy = msg.sender;
        request.history.push(StatusUpdate({ status: Status.Accepted, timestamp: uint40(block.timestamp) }));

        _latestRequestByStatus[Status.Accepted] = id;

        emit Accepted(id, msg.sender);
    }

    /**
     * @notice Reject an open request.
     * @dev Only a whitelisted solver can reject.
     * @param id  ID of the request.
     */
    function reject(bytes32 id, RejectReason reason) external onlyRoles(SOLVER) nonReentrant {
        Request storage request = _requests[id];
        if (request.status != Status.Pending) revert NotPending();

        request.status = Status.Rejected;
        request.updatedAt = uint40(block.timestamp);
        request.history.push(StatusUpdate({ status: Status.Rejected, timestamp: uint40(block.timestamp) }));

        _latestRequestByStatus[Status.Rejected] = id;

        emit Rejected(id, msg.sender, reason);
    }

    /**
     * @notice Cancel an open or rejected request and refund deposits.
     * @dev Only request initiator can cancel.
     * @param id  ID of the request.
     */
    function cancel(bytes32 id) external nonReentrant {
        Request storage request = _requests[id];
        if (request.status != Status.Pending && request.status != Status.Rejected) revert NotPendingOrRejected();
        if (request.order.user != msg.sender) revert Unauthorized();

        request.status = Status.Reverted;
        request.updatedAt = uint40(block.timestamp);
        request.history.push(StatusUpdate({ status: Status.Reverted, timestamp: uint40(block.timestamp) }));

        _latestRequestByStatus[Status.Reverted] = id;

        _transferDeposits(request.order.user, request.order.minReceived);

        emit Reverted(id);
    }

    /**
     * @notice Fulfill a request.
     * @dev Only callable by the outbox.
     * @param id        ID of the request.
     * @param callHash  Hash of the calls for this request executed on another chain.
     */
    function markFulfilled(bytes32 id, bytes32 callHash) external xrecv nonReentrant {
        Request storage request = _requests[id];
        if (request.status != Status.Accepted) revert NotAccepted();
        if (xmsg.sender != _outbox) revert NotOutbox();
        if (xmsg.sourceChainId != request.order.fillInstructions[0].destinationChainId) revert WrongSourceChain();

        // Ensure reported call hash matches requested call hash
        if (callHash != _callHash(id, uint64(block.chainid), request.order.fillInstructions[0].originData)) {
            revert WrongCallHash();
        }

        request.status = Status.Fulfilled;
        request.updatedAt = uint40(block.timestamp);
        request.history.push(StatusUpdate({ status: Status.Fulfilled, timestamp: uint40(block.timestamp) }));

        _latestRequestByStatus[Status.Fulfilled] = id;

        emit Fulfilled(id, callHash, request.acceptedBy);
    }

    /**
     * @notice Claim a fulfilled request.
     * @param id  ID of the request.
     * @param to  Address to send deposits to.
     */
    function claim(bytes32 id, address to) external nonReentrant {
        Request storage request = _requests[id];
        if (request.status != Status.Fulfilled) revert NotFulfilled();
        if (request.acceptedBy != msg.sender) revert Unauthorized();

        request.status = Status.Claimed;
        request.updatedAt = uint40(block.timestamp);
        request.history.push(StatusUpdate({ status: Status.Claimed, timestamp: uint40(block.timestamp) }));

        _latestRequestByStatus[Status.Claimed] = id;

        _transferDeposits(to, request.order.minReceived);

        emit Claimed(id, msg.sender, to, request.order.minReceived);
    }

    /**
     * @dev Validate all order fields.
     */
    function _validateOnchainOrder(OnchainCrossChainOrder calldata order) internal view returns (OrderData memory) {
        if (order.fillDeadline < block.timestamp) revert InvalidFillDeadline();
        // TODO: validate orderDataType
        if (order.orderData.length == 0) revert InvalidOrderData();

        OrderData memory orderData = abi.decode(order.orderData, (OrderData));
        if (orderData.destChainId == 0 || orderData.destChainId == block.chainid) revert InvalidChain();
        if (orderData.deposits.length == 0) revert NoDeposits(); // Do we need to enforce this?
        if (orderData.prereqs.length == 0) revert NoTokenPrereqs(); // Do we need to enforce this?
        if (orderData.calls.length == 0) revert NoCalls();

        return orderData;
    }

    /**
     * @dev Validate all token pre-requisites.
     * @dev `Output.recipient` is populated with the address the Outbox will sign an approval to.
     */
    function _validateTokenPrereqs(TokenPrereq[] memory tokenPrereqs) internal pure {
        for (uint256 i; i < tokenPrereqs.length; ++i) {
            TokenPrereq memory tokenPrereq = tokenPrereqs[i];
            // No zero address tokens, natives are handled directly by the outbox
            if (tokenPrereq.token == bytes32(0)) revert ZeroAddress();
            // No zero amount prereqs
            if (tokenPrereq.amount == 0) revert ZeroAmount();
            // Ensure ERC20 spender is specified
            if (tokenPrereq.spender == bytes32(0)) revert NoSpender();
        }
    }

    /**
     * @dev Validate all calls.
     */
    function _validateCalls(Call[] memory calls) internal pure {
        for (uint256 i; i < calls.length; ++i) {
            Call memory call = calls[i];
            // Only prevent calls to zero address
            // We support native payments, so no calldata check, amount check isn't needed either
            if (call.target == bytes32(0)) revert ZeroAddress();
        }
    }

    /**
     * @dev Process and validate all deposits.
     * @dev Native deposit is validated by checking msg.value against deposit amount, and must be included in array.
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
     * @dev Open a request to execute a call on another chain, backed by deposits.
     */
    function _openRequest(OnchainCrossChainOrder calldata order) internal returns (Request storage request) {
        ResolvedCrossChainOrder memory resolvedOrder = resolve(order);
        request = _requests[_incrementId()];

        request.order.user = resolvedOrder.user;
        request.order.originChainId = resolvedOrder.originChainId;
        request.order.openDeadline = resolvedOrder.openDeadline;
        request.order.fillDeadline = resolvedOrder.fillDeadline;
        request.order.orderId = resolvedOrder.orderId;

        for (uint256 i; i < resolvedOrder.maxSpent.length; ++i) {
            request.order.maxSpent.push(resolvedOrder.maxSpent[i]);
        }

        for (uint256 i; i < resolvedOrder.minReceived.length; ++i) {
            request.order.minReceived.push(resolvedOrder.minReceived[i]);
        }

        for (uint256 i; i < resolvedOrder.fillInstructions.length; ++i) {
            request.order.fillInstructions.push(resolvedOrder.fillInstructions[i]);
        }

        request.status = Status.Pending;
        request.updatedAt = uint40(block.timestamp);
        request.acceptedBy = address(0);
        request.history.push(StatusUpdate({ status: Status.Pending, timestamp: uint40(block.timestamp) }));

        _latestRequestByStatus[Status.Pending] = resolvedOrder.orderId;

        return request;
    }

    /**
     * @dev Transfer deposits to recipient. Used for both refunds and claims.
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
        _lastId++;
        return bytes32(_lastId);
    }

    /**
     * @dev Returns call hash. Used to discern fullfilment.
     */
    function _callHash(bytes32 id, uint64 srcChainId, bytes memory fillOriginData) internal pure returns (bytes32) {
        return keccak256(abi.encode(id, srcChainId, fillOriginData));
    }

    /**
     * @dev Returns true if the address is a contract.
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
     */
    function _addressToBytes32(address addr) internal pure returns (bytes32) {
        return bytes32(uint256(uint160(addr)));
    }

    /**
     * @dev Convert bytes32 to address.
     */
    function _bytes32ToAddress(bytes32 b) internal pure returns (address) {
        return address(uint160(uint256(b)));
    }
}
