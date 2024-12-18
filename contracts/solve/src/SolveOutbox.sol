// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.24;

import { OwnableRoles } from "solady/src/auth/OwnableRoles.sol";
import { ReentrancyGuard } from "solady/src/utils/ReentrancyGuard.sol";
import { Initializable } from "solady/src/utils/Initializable.sol";
import { XAppBase } from "core/src/pkg/XAppBase.sol";

import { SafeTransferLib } from "solady/src/utils/SafeTransferLib.sol";
import { ConfLevel } from "core/src/libraries/ConfLevel.sol";
import { TypeMax } from "core/src/libraries/TypeMax.sol";
import { Solve } from "./Solve.sol";

import { ISolveInbox } from "./interfaces/ISolveInbox.sol";
import { IArbSys } from "./interfaces/IArbSys.sol";

/**
 * @title SolveOutbox
 * @notice Entrypoint for fulfillments of user solve requests.
 */
contract SolveOutbox is OwnableRoles, ReentrancyGuard, Initializable, XAppBase {
    using SafeTransferLib for address;

    error CallFailed();
    error CallNotAllowed();
    error AreadyFulfilled();
    error WrongDestChain();
    error IncorrectPrereqs();
    error InsufficientFee();

    event AllowedCallSet(address indexed target, bytes4 indexed selector, bool allowed);

    /**
     * @notice Emitted when a request is fulfilled.
     * @param reqId       ID of the request.
     * @param callHash    Hash of the call executed.
     * @param solvedBy    Address of the solver.
     */
    event Fulfilled(bytes32 indexed reqId, bytes32 indexed callHash, address indexed solvedBy);

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
     * @notice Gas limit for SolveInbox.markFulfilled callback.
     */
    uint64 internal constant MARK_FULFILLED_GAS_LIMIT = 125_000;

    /**
     * @notice Stubbed calldata for SolveInbox.markFulfilled. Used to estimate the gas cost.
     * @dev Type maxes used to ensure no non-zero bytes in fee estimation.
     */
    bytes internal constant MARK_FULFILLED_STUB_CDATA =
        abi.encodeCall(ISolveInbox.markFulfilled, (TypeMax.Bytes32, TypeMax.Bytes32));

    /**
     * @notice Address of the inbox contract.
     */
    address internal _inbox;

    /**
     * @notice Mapping of allowed calls per contract.
     */
    mapping(address target => mapping(bytes4 selector => bool)) public allowedCalls;

    /**
     * @notice Mapping of fulfilled calls.
     * @dev callHash used to prevent duplicate fulfillment.
     */
    mapping(bytes32 callHash => bool fulfilled) public fulfilledCalls;

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
    function initialize(address owner_, address solver_, address omni_, address inbox_) external initializer {
        _initializeOwner(owner_);
        _grantRoles(solver_, SOLVER);
        _setOmniPortal(omni_);
        _inbox = inbox_;
    }

    /**
     * @notice Calculate the message passing fee for a fulfill call.
     * @param srcChainId ID of the source chain.
     */
    function fulfillFee(uint64 srcChainId) public view returns (uint256) {
        return feeFor(srcChainId, MARK_FULFILLED_STUB_CDATA, MARK_FULFILLED_GAS_LIMIT);
    }

    /**
     * @notice Check if a call has been fulfilled.
     * @param srcReqId      ID of the on the source inbox.
     * @param srcChainId    ID of the source chain.
     * @param call          Details of the call executed.
     */
    function didFulfill(bytes32 srcReqId, uint64 srcChainId, Solve.Call calldata call) external view returns (bool) {
        return fulfilledCalls[_callHash(srcReqId, srcChainId, call)];
    }

    /**
     * @notice Set an allowed call for a target contract.
     * @param target    Address of the target contract.
     * @param selector  4-byte selector of the function to allow.
     * @param allowed   Whether the call is allowed.
     */
    function setAllowedCall(address target, bytes4 selector, bool allowed) external onlyOwner {
        allowedCalls[target][selector] = allowed;
        emit AllowedCallSet(target, selector, allowed);
    }

    /**
     * @notice Fulfill a request.
     * @param srcReqId      ID of the request on the source inbox.
     * @param srcChainId    ID of the source chain.
     * @param call          Details of the call to execute.
     * @param prereqs       Pre-requisite token deposits required by the call.
     */
    function fulfill(
        bytes32 srcReqId,
        uint64 srcChainId,
        Solve.Call calldata call,
        Solve.TokenPrereq[] calldata prereqs
    ) external payable onlyRoles(SOLVER) nonReentrant {
        if (call.destChainId != block.chainid) revert WrongDestChain();
        if (!allowedCalls[call.target][bytes4(call.data)]) revert CallNotAllowed();

        // If the call has already been fulfilled, revert. Else, mark fulfilled
        bytes32 callHash = _callHash(srcReqId, srcChainId, call);
        if (fulfilledCalls[callHash]) revert AreadyFulfilled();
        fulfilledCalls[callHash] = true;

        // Process token pre-requisites. Record pre-call balances (checked against post-call balances)
        uint256[] memory prereqBalances = new uint256[](prereqs.length);
        for (uint256 i; i < prereqs.length; ++i) {
            prereqBalances[i] = prereqs[i].token.balanceOf(address(this));
            prereqs[i].token.safeTransferFrom(msg.sender, address(this), prereqs[i].amount);
            prereqs[i].token.safeApprove(prereqs[i].spender, prereqs[i].amount);
        }

        // Execute the call
        (bool success,) = payable(call.target).call{ value: call.value }(call.data);
        if (!success) revert CallFailed();

        // Require post-call balances matches pre-call. Ensures prequisites match call transfers.
        for (uint256 i; i < prereqs.length; ++i) {
            if (prereqs[i].token.balanceOf(address(this)) != prereqBalances[i]) revert IncorrectPrereqs();
        }

        // Mark the call as fulfilled on inbox
        bytes memory xcalldata = abi.encodeCall(ISolveInbox.markFulfilled, (srcReqId, callHash));
        uint256 fee = xcall(srcChainId, ConfLevel.Finalized, _inbox, xcalldata, MARK_FULFILLED_GAS_LIMIT);
        if (msg.value - call.value < fee) revert InsufficientFee();

        emit Fulfilled(srcReqId, callHash, msg.sender);
    }

    /**
     * @dev Returns call hash. Used to discern fullfilment.
     */
    function _callHash(bytes32 srcReqId, uint64 srcChainId, Solve.Call calldata call) internal pure returns (bytes32) {
        return keccak256(abi.encode(srcReqId, srcChainId, call));
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
