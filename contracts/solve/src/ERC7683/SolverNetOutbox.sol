// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.24;

import { OwnableRoles } from "solady/src/auth/OwnableRoles.sol";
import { ReentrancyGuard } from "solady/src/utils/ReentrancyGuard.sol";
import { Initializable } from "solady/src/utils/Initializable.sol";
import { XAppBase } from "core/src/pkg/XAppBase.sol";

import { SafeTransferLib } from "solady/src/utils/SafeTransferLib.sol";
import { ConfLevel } from "core/src/libraries/ConfLevel.sol";
import { TypeMax } from "core/src/libraries/TypeMax.sol";

import { ISolverNetInbox } from "./interfaces/ISolverNetInbox.sol";
import { ISolverNetOutbox } from "./interfaces/ISolverNetOutbox.sol";
import { IArbSys } from "../interfaces/IArbSys.sol";

/**
 * @title SolverNetOutbox
 * @notice Entrypoint for fulfillments of user solve requests.
 */
contract SolverNetOutbox is OwnableRoles, ReentrancyGuard, Initializable, XAppBase, ISolverNetOutbox {
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
     * @notice Gas limit for SolveInbox.markFulfilled callback.
     */
    uint64 internal constant MARK_FULFILLED_GAS_LIMIT = 100_000;

    /**
     * @notice Stubbed calldata for SolveInbox.markFulfilled. Used to estimate the gas cost.
     * @dev Type maxes used to ensure no non-zero bytes in fee estimation.
     */
    bytes internal constant MARK_FULFILLED_STUB_CDATA =
        abi.encodeCall(ISolverNetInbox.markFulfilled, (TypeMax.Bytes32, TypeMax.Bytes32));

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
     * @param srcReqId          ID of the on the source inbox.
     * @param srcChainId        ID of the source chain.
     * @param fillOriginData    Data emitted on the origin to parameterize the fill
     */
    function didFulfill(bytes32 srcReqId, uint64 srcChainId, bytes calldata fillOriginData) external view returns (bool) {
        return fulfilledCalls[_callHash(srcReqId, srcChainId, fillOriginData)];
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
     * @notice Fills a single leg of a particular order on the destination chain
     * @param orderId     Unique order identifier for this order
     * @param originData  Data emitted on the origin to parameterize the fill
     * @dev fillerData (currently unused): Data provided by the filler to inform the fill or express their preferences
     */
    function fill(bytes32 orderId, bytes calldata originData, bytes calldata) external onlyRoles(SOLVER) nonReentrant {
        FillOriginData memory requestData = abi.decode(originData, (FillOriginData));

        if (requestData.destChainId != block.chainid) revert WrongDestChain();
        for (uint256 i; i < requestData.calls.length; ++i) {
            Call memory call = requestData.calls[i];
            if (!allowedCalls[_bytes32ToAddress(call.target)][bytes4(call.callData)]) revert CallNotAllowed();
        }

        // If the call has already been fulfilled, revert. Else, mark fulfilled
        bytes32 callHash = _callHash(orderId, requestData.srcChainId, originData);
        if (fulfilledCalls[callHash]) revert AlreadyFulfilled();
        fulfilledCalls[callHash] = true;

        // Process token pre-requisites. Record pre-call balances (checked against post-call balances)
        uint256[] memory prereqBalances = new uint256[](requestData.prereqs.length);
        for (uint256 i; i < requestData.prereqs.length; ++i) {
            address token = _bytes32ToAddress(requestData.prereqs[i].token);
            address spender = _bytes32ToAddress(requestData.prereqs[i].recipient);

            prereqBalances[i] = token.balanceOf(address(this));
            token.safeTransferFrom(msg.sender, address(this), requestData.prereqs[i].amount);
            token.safeApprove(spender, requestData.prereqs[i].amount);
        }

        // Execute the calls
        uint256 nativeAmount = 0;
        for (uint256 i; i < requestData.calls.length; ++i) {
            Call memory call = requestData.calls[i];
            address target = _bytes32ToAddress(call.target);
            (bool success,) = payable(target).call{ value: call.value }(call.callData);
            if (!success) revert CallFailed();
            nativeAmount += call.value;
        }

        // Require post-call balances matches pre-call. Ensures prequisites match call transfers.
        for (uint256 i; i < requestData.prereqs.length; ++i) {
            address token = _bytes32ToAddress(requestData.prereqs[i].token);
            if (token.balanceOf(address(this)) != prereqBalances[i]) revert IncorrectPrereqs();
        }

        // Mark the call as fulfilled on inbox
        bytes memory xcalldata = abi.encodeCall(ISolverNetInbox.markFulfilled, (orderId, callHash));
        uint256 fee = xcall(uint64(requestData.srcChainId), ConfLevel.Finalized, _inbox, xcalldata, MARK_FULFILLED_GAS_LIMIT);
        if (msg.value - nativeAmount < fee) revert InsufficientFee();

        emit Fulfilled(orderId, callHash, msg.sender);
    }

    /**
     * @dev Returns call hash. Used to discern fullfilment.
     */
    function _callHash(bytes32 srcReqId, uint64 srcChainId, bytes memory fillOriginData) internal pure returns (bytes32) {
        return keccak256(abi.encode(srcReqId, srcChainId, fillOriginData));
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
     * @dev Convert bytes32 to address.
     */
    function _bytes32ToAddress(bytes32 b) internal pure returns (address) {
        return address(uint160(uint256(b)));
    }
}
