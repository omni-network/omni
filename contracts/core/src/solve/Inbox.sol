// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.24;

import { OwnableRoles } from "solady/src/auth/OwnableRoles.sol";
import { ReentrancyGuard } from "solady/src/utils/ReentrancyGuard.sol";
import { Initializable } from "solady/src/utils/Initializable.sol";
import { XAppBase } from "../pkg/XAppBase.sol";
import { SafeTransferLib } from "solady/src/utils/SafeTransferLib.sol";
import { IConversionRateOracle } from "../interfaces/IConversionRateOracle.sol";
import { Solve } from "./Solve.sol";

/**
 * @title Inbox
 * @notice Entrypoint and alt-mempoool for user solve requests.
 */
contract Inbox is OwnableRoles, ReentrancyGuard, Initializable, XAppBase {
    using SafeTransferLib for address;

    error NoDeposits();
    error InvalidCall();
    error IncorrectChain();
    error InvalidDeposit();
    error TransferFailed();
    error RequestStateInvalid();

    /**
     * @notice Emitted when a request is created.
     * @param id        ID of the request.
     * @param from      Address of the user who created the request.
     * @param call      Details of the call to be executed on another chain.
     * @param deposits  Array of deposits backing the request.
     */
    event Requested(bytes32 indexed id, address indexed from, Solve.Call call, Solve.Deposit[] deposits);

    /**
     * @notice Emitted when a request is accepted.
     * @param id  ID of the request.
     * @param by  Address of the solver who accepted the request.
     */
    event Accepted(bytes32 indexed id, address indexed by);

    /**
     * @notice Emitted when a request is rejected.
     * @param id  ID of the request.
     * @param by  Address of the solver who rejected the request.
     */
    event Rejected(bytes32 indexed id, address indexed by);

    /**
     * @notice Emitted when a request is cancelled.
     * @param id  ID of the request.
     */
    event Reverted(bytes32 indexed id);

    /**
     * @notice Emitted when a request is fulfilled.
     * @param guid        ID of the request.
     * @param callHash    Hash of the call executed on another chain.
     * @param creditedTo  Address of the recipient credited the funds by the solver.
     */
    event Fulfilled(bytes32 indexed guid, bytes32 indexed callHash, address indexed creditedTo);

    /**
     * @notice Role for solvers.
     * @dev _ROLE_0 evaluates to '1'.
     */
    uint256 internal constant SOLVER = _ROLE_0;

    /**
     * @dev uint repr of last assigned request ID.
     */
    uint256 internal _lastId;

    /**
     * @notice Address of the outbox contract.
     */
    address internal _outbox;

    /**
     * @notice Map ID to request.
     */
    mapping(bytes32 id => Solve.Request) internal _requests;

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
    function getRequest(bytes32 id) external view returns (Solve.Request memory) {
        return _requests[id];
    }

    /**
     * @notice Suggest the amount of native currency to send with a request.
     * @param call        Details of the call to be executed on another chain.
     * @param gasLimit    Maximum gas limit for the call.
     * @param gasPrice    Gas price in wei.
     * @param fulfillFee  Fee for the fulfill call, retrieved from the destination outbox.
     */
    function suggestNativePayment(Solve.Call calldata call, uint64 gasLimit, uint64 gasPrice, uint256 fulfillFee)
        external
        view
        returns (uint256)
    {
        IConversionRateOracle oracle = IConversionRateOracle(omni.feeOracle());

        uint256 nativeValue = call.value * oracle.toNativeRate(call.destChainId) / oracle.CONVERSION_RATE_DENOM();
        uint256 executionFee = omni.feeFor(call.destChainId, call.data, gasLimit);
        uint256 acceptFee = 55_000 * gasPrice;
        uint256 solveFee = 100_000 gwei; // TODO: determine solve fee

        return nativeValue + executionFee + acceptFee + solveFee + fulfillFee;
    }

    /**
     * @notice Open a request to execute a call on another chain, backed by deposits.
     *  Token deposits are transferred from msg.sender to this inbox.
     * @param call      Details of the call to be executed on another chain.
     * @param deposits  Array of deposits backing the request.
     */
    function request(Solve.Call calldata call, Solve.TokenDeposit[] calldata deposits)
        external
        payable
        nonReentrant
        returns (bytes32 id)
    {
        if (call.target == address(0)) revert InvalidCall();
        if (call.destChainId == 0) revert InvalidCall();
        if (call.data.length == 0) revert InvalidCall();
        if (deposits.length == 0 && msg.value == 0) revert NoDeposits();

        Solve.Request storage req = _openRequest(msg.sender, call, deposits);

        emit Requested(req.id, req.from, req.call, req.deposits);

        return req.id;
    }

    /**
     * @notice Accept an open request.
     * @dev Only a whitelisted solver can accept.
     * @param id  ID of the request.
     */
    function accept(bytes32 id) external onlyRoles(SOLVER) nonReentrant {
        Solve.Request storage req = _requests[id];
        if (req.status != Solve.Status.Pending) revert RequestStateInvalid();

        req.updatedAt = uint40(block.timestamp);
        req.status = Solve.Status.Accepted;
        req.acceptedBy = msg.sender;

        emit Accepted(id, msg.sender);
    }

    /**
     * @notice Reject an open request.
     * @dev Only a whitelisted solver can reject.
     * @param id  ID of the request.
     */
    function reject(bytes32 id) external onlyRoles(SOLVER) nonReentrant {
        Solve.Request storage req = _requests[id];
        if (req.status != Solve.Status.Pending) revert RequestStateInvalid();

        req.updatedAt = uint40(block.timestamp);
        req.status = Solve.Status.Rejected;

        emit Rejected(id, msg.sender);
    }

    /**
     * @notice Cancel an open or rejected request and refund deposits.
     * @dev Only request initiator can cancel.
     * @param id  ID of the request.
     */
    function cancel(bytes32 id) external nonReentrant {
        Solve.Request storage req = _requests[id];
        if (req.status != Solve.Status.Pending && req.status != Solve.Status.Rejected) revert RequestStateInvalid();
        if (req.from != msg.sender) revert Unauthorized();

        _cancelRequest(req);

        emit Reverted(id);
    }

    /**
     * @notice Fulfill a request.
     * @dev Only callable by the outbox.
     */
    function markFulfilled(bytes32 guid, bytes32 callHash, address creditTo) external xrecv nonReentrant {
        // Validate request state and fulfillment context
        Solve.Request storage req = _requests[guid];
        if (req.status != Solve.Status.Accepted) revert RequestStateInvalid();
        if (req.call.destChainId != xmsg.sourceChainId) revert IncorrectChain();
        if (xmsg.sender != _outbox) revert Unauthorized();
        // No need to check if msg.sender is OmniPortal as we use xrecv for source validation

        // Validate call hash, ensuring the correct action is being fulfilled
        bytes32 _callHash = keccak256(abi.encode(guid, req.call));
        if (_callHash != callHash) revert InvalidCall();

        // Update request state
        req.updatedAt = uint40(block.timestamp);
        req.status = Solve.Status.Fulfilled;

        // Transfer deposits to solver
        Solve.Deposit[] memory deposits = req.deposits;
        for (uint256 i; i < deposits.length; ++i) {
            if (deposits[i].isNative) {
                (bool success,) = payable(creditTo).call{ value: deposits[i].amount }("");
                if (!success) revert TransferFailed();
            } else {
                deposits[i].token.safeTransfer(creditTo, deposits[i].amount);
            }
        }

        emit Fulfilled(guid, callHash, creditTo);
    }

    /**
     * @dev Open a new request in storage at `id`.
     *      Transfer token deposits from msg.sender to this inbox.
     *      Duplicate token addresses are allowed.
     */
    function _openRequest(address from, Solve.Call calldata call, Solve.TokenDeposit[] calldata deposits)
        internal
        returns (Solve.Request storage req)
    {
        bytes32 id = _nextId();

        req = _requests[id];
        req.id = id;
        req.updatedAt = uint40(block.timestamp);
        req.status = Solve.Status.Pending;
        req.from = from;
        req.call = call;

        if (msg.value > 0) {
            req.deposits.push(Solve.Deposit({ isNative: true, token: address(0), amount: msg.value }));
        }

        for (uint256 i = 0; i < deposits.length; i++) {
            if (deposits[i].amount == 0) revert InvalidDeposit();
            if (deposits[i].token == address(0)) revert InvalidDeposit();

            req.deposits.push(Solve.Deposit({ isNative: false, token: deposits[i].token, amount: deposits[i].amount }));

            // NOTE: all external methods must be nonReentrant
            // This allows us to transfer while opening the request - saving some gas.
            deposits[i].token.safeTransferFrom(msg.sender, address(this), deposits[i].amount);
        }
    }

    function _cancelRequest(Solve.Request storage req) internal {
        // Update state
        req.updatedAt = uint40(block.timestamp);
        req.status = Solve.Status.Reverted;

        // Refund deposits
        uint256 length = req.deposits.length;
        for (uint256 i = 0; i < length; i++) {
            Solve.Deposit memory deposit = req.deposits[i];
            if (deposit.isNative) {
                (bool success,) = payable(msg.sender).call{ value: deposit.amount }("");
                if (!success) revert TransferFailed();
            } else {
                deposit.token.safeTransfer(msg.sender, deposit.amount);
            }
        }
    }

    /**
     * @dev Increment and return _lastId.
     */
    function _nextId() internal returns (bytes32) {
        _lastId++;
        return bytes32(_lastId);
    }
}
