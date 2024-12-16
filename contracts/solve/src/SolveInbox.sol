// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.24;

import { OwnableRoles } from "solady/src/auth/OwnableRoles.sol";
import { ReentrancyGuard } from "solady/src/utils/ReentrancyGuard.sol";
import { Initializable } from "solady/src/utils/Initializable.sol";
import { SafeTransferLib } from "solady/src/utils/SafeTransferLib.sol";
import { XAppBase } from "core/src/pkg/XAppBase.sol";
import { ISolveInbox } from "./interfaces/ISolveInbox.sol";
import { IArbSys } from "./interfaces/IArbSys.sol";
import { Solve } from "./Solve.sol";

/**
 * @title SolveInbox
 * @notice Entrypoint and alt-mempoool for user solve requests.
 */
contract SolveInbox is OwnableRoles, ReentrancyGuard, Initializable, XAppBase, ISolveInbox {
    using SafeTransferLib for address;

    // Request creation errors
    error NoDeposits();
    error InvalidCall();
    error InvalidDeposit();

    // Request state transition errors
    error NotPending();
    error NotPendingOrRejected();
    error NotAccepted();
    error NotFulfilled();

    // Request fulfillment errors
    error NotOutbox();
    error WrongCallHash();
    error WrongSourceChain();

    // Transfer errors
    error TransferFailed();
    error InvalidRecipient();

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
     * @notice Map status to latest request ID.
     */
    mapping(Solve.Status => bytes32 id) internal _latestReqByStatus;

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
    function getRequest(bytes32 id) external view returns (Solve.Request memory) {
        return _requests[id];
    }

    /**
     * @notice Returns the latest request with the given status.
     */
    function getLatestRequestByStatus(Solve.Status status) external view returns (Solve.Request memory) {
        return _requests[_latestReqByStatus[status]];
    }

    /**
     * @notice Returns the update history for a request.
     */
    function getUpdateHistory(bytes32 id) external view returns (Solve.StatusUpdate[] memory) {
        return _requests[id].updateHistory;
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
        if (req.status != Solve.Status.Pending) revert NotPending();

        req.updatedAt = uint40(block.timestamp);
        req.status = Solve.Status.Accepted;
        req.acceptedBy = msg.sender;
        req.updateHistory.push(
            Solve.StatusUpdate({ status: Solve.Status.Accepted, timestamp: uint40(block.timestamp) })
        );

        _latestReqByStatus[Solve.Status.Accepted] = id;

        emit Accepted(id, msg.sender);
    }

    /**
     * @notice Reject an open request.
     * @dev Only a whitelisted solver can reject.
     * @param id  ID of the request.
     * @param reason  Reason code for rejecting the request (see solver/app/reject.go for current codes)
     */
    function reject(bytes32 id, uint8 reason) external onlyRoles(SOLVER) nonReentrant {
        Solve.Request storage req = _requests[id];
        if (req.status != Solve.Status.Pending) revert NotPending();

        req.updatedAt = uint40(block.timestamp);
        req.status = Solve.Status.Rejected;
        req.updateHistory.push(
            Solve.StatusUpdate({ status: Solve.Status.Rejected, timestamp: uint40(block.timestamp) })
        );

        _latestReqByStatus[Solve.Status.Rejected] = id;

        emit Rejected(id, msg.sender, reason);
    }

    /**
     * @notice Cancel an open or rejected request and refund deposits.
     * @dev Only request initiator can cancel.
     * @param id  ID of the request.
     */
    function cancel(bytes32 id) external nonReentrant {
        Solve.Request storage req = _requests[id];
        if (req.status != Solve.Status.Pending && req.status != Solve.Status.Rejected) revert NotPendingOrRejected();
        if (req.from != msg.sender) revert Unauthorized();

        req.updatedAt = uint40(block.timestamp);
        req.status = Solve.Status.Reverted;
        req.updateHistory.push(
            Solve.StatusUpdate({ status: Solve.Status.Reverted, timestamp: uint40(block.timestamp) })
        );

        _latestReqByStatus[Solve.Status.Reverted] = id;

        _transferDeposits(req.from, req.deposits);

        emit Reverted(id);
    }

    /**
     * @notice Fulfill a request.
     * @dev Only callable by the outbox.
     */
    function markFulfilled(bytes32 id, bytes32 callHash) external xrecv nonReentrant {
        Solve.Request storage req = _requests[id];
        if (req.status != Solve.Status.Accepted) revert NotAccepted();
        if (xmsg.sender != _outbox) revert NotOutbox();
        if (xmsg.sourceChainId != req.call.destChainId) revert WrongSourceChain();

        // Ensure reported call hash matches requested call hash
        if (callHash != _callHash(id, uint64(block.chainid), req.call)) revert WrongCallHash();

        req.updatedAt = uint40(block.timestamp);
        req.status = Solve.Status.Fulfilled;
        req.updateHistory.push(
            Solve.StatusUpdate({ status: Solve.Status.Fulfilled, timestamp: uint40(block.timestamp) })
        );

        _latestReqByStatus[Solve.Status.Fulfilled] = id;

        emit Fulfilled(id, callHash, req.acceptedBy);
    }

    /**
     * @notice Claim a fulfilled request.
     * @param id  ID of the request.
     * @param to  Address to send deposits to.
     */
    function claim(bytes32 id, address to) external nonReentrant {
        Solve.Request storage req = _requests[id];
        if (req.status != Solve.Status.Fulfilled) revert NotFulfilled();
        if (req.acceptedBy != msg.sender) revert Unauthorized();

        req.updatedAt = uint40(block.timestamp);
        req.status = Solve.Status.Claimed;
        req.updateHistory.push(Solve.StatusUpdate({ status: Solve.Status.Claimed, timestamp: uint40(block.timestamp) }));

        _latestReqByStatus[Solve.Status.Claimed] = id;

        _transferDeposits(to, req.deposits);

        emit Claimed(id, msg.sender, to, req.deposits);
    }

    /**
     * @dev Transfer deposits to recipient. Used regardless of refund or claim.
     */
    function _transferDeposits(address recipient, Solve.Deposit[] memory deposits) internal {
        if (recipient == address(0)) revert InvalidRecipient();

        for (uint256 i; i < deposits.length; ++i) {
            if (deposits[i].isNative) {
                (bool success,) = payable(recipient).call{ value: deposits[i].amount }("");
                if (!success) revert TransferFailed();
            } else {
                deposits[i].token.safeTransfer(recipient, deposits[i].amount);
            }
        }
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
        req.updateHistory.push(Solve.StatusUpdate({ status: Solve.Status.Pending, timestamp: uint40(block.timestamp) }));

        _latestReqByStatus[Solve.Status.Pending] = id;

        if (msg.value > 0) {
            req.deposits.push(Solve.Deposit({ isNative: true, token: address(0), amount: msg.value }));
        }

        for (uint256 i = 0; i < deposits.length; i++) {
            if (deposits[i].amount == 0) revert InvalidDeposit();
            if (deposits[i].token == address(0)) revert InvalidDeposit();

            req.deposits.push(Solve.Deposit({ isNative: false, token: deposits[i].token, amount: deposits[i].amount }));

            // NOTE: all external methods must be nonReentrant
            // This allows us to transfer while opening the request
            deposits[i].token.safeTransferFrom(msg.sender, address(this), deposits[i].amount);
        }
    }

    /**
     * @dev Increment and return _lastId.
     */
    function _nextId() internal returns (bytes32) {
        _lastId++;
        return bytes32(_lastId);
    }

    /**
     * @dev Returns call hash. Used to discern fullfilment.
     */
    function _callHash(bytes32 id, uint64 sourceChainId, Solve.Call storage call) internal pure returns (bytes32) {
        return keccak256(abi.encode(id, sourceChainId, call));
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
}
