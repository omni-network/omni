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
     * @notice Returns the message passing fee required to mark a request as fulfilled on the source chain
     * @param srcChainId  ID of the source chain.
     * @return            Fee amount in native currency.
     */
    function fulfillFee(uint64 srcChainId) public view returns (uint256) {
        return feeFor(srcChainId, MARK_FULFILLED_STUB_CDATA, MARK_FULFILLED_GAS_LIMIT);
    }

    /**
     * @notice Check if a call has been fulfilled.
     * @param srcReqId    ID of the on the source inbox.
     * @param originData  Data emitted on the origin to parameterize the fill
     */
    function didFulfill(bytes32 srcReqId, bytes calldata originData) external view returns (bool) {
        return fulfilledCalls[_callHash(srcReqId, originData)];
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
     * @notice Fills a particular order on the destination chain
     * @param orderId     Unique order identifier for this order
     * @param originData  Data emitted on the origin to parameterize the fill
     * @dev fillerData (currently unused): Data provided by the filler to inform the fill or express their preferences
     */
    function fill(bytes32 orderId, bytes calldata originData, bytes calldata)
        external
        payable
        onlyRoles(SOLVER)
        nonReentrant
    {
        SolverNetIntent memory intent = abi.decode(originData, (SolverNetIntent));

        // Check that the destination chain is the current chain
        if (intent.destChainId != block.chainid) revert WrongDestChain();

        // If the call has already been fulfilled, revert. Else, mark fulfilled
        bytes32 callHash = _callHash(orderId, originData);
        if (fulfilledCalls[callHash]) revert AlreadyFulfilled();
        fulfilledCalls[callHash] = true;

        // Determine tokens required, record pre-call balances, retrieve tokens from solver, and sign approvals
        (address[] memory tokens, uint256[] memory preBalances) = _prepareIntent(intent);

        // Execute the calls
        uint256 nativeAmountRequired = _executeIntent(intent);

        // Require post-call balance matches pre-call. Ensures prerequisites match call transfers.
        // Native balance is validated after xcall
        for (uint256 i; i < tokens.length; ++i) {
            if (tokens[i] != address(0)) {
                if (tokens[i].balanceOf(address(this)) != preBalances[i]) revert InvalidPrereq();
            }
        }

        // Mark the call as fulfilled on inbox
        bytes memory xcalldata = abi.encodeCall(ISolverNetInbox.markFulfilled, (orderId, callHash));
        uint256 fee = xcall(uint64(intent.srcChainId), ConfLevel.Finalized, _inbox, xcalldata, MARK_FULFILLED_GAS_LIMIT);
        if (msg.value - nativeAmountRequired < fee) revert InsufficientFee();

        // Refund any overpayment in native currency
        uint256 refund = msg.value - nativeAmountRequired - fee;
        if (refund > 0) msg.sender.safeTransferETH(refund);

        emit Fulfilled(orderId, callHash, msg.sender);
    }

    function _prepareIntent(SolverNetIntent memory intent)
        internal
        returns (address[] memory tokens, uint256[] memory preBalances)
    {
        TokenPrereq[] memory prereqs = intent.tokenPrereqs;
        tokens = new address[](prereqs.length);
        preBalances = new uint256[](prereqs.length);

        for (uint256 i; i < prereqs.length; ++i) {
            TokenPrereq memory prereq = prereqs[i];
            address token = _bytes32ToAddress(prereq.token);
            tokens[i] = token;

            if (token == address(0)) {
                if (prereq.amount >= msg.value || prereq.amount != intent.call.value) revert InvalidPrereq();
                preBalances[i] = address(this).balance - msg.value;
            } else {
                preBalances[i] = token.balanceOf(address(this));
                address spender = _bytes32ToAddress(prereq.spender);

                token.safeTransferFrom(msg.sender, address(this), prereq.amount);
                token.safeApprove(spender, prereq.amount);
            }
        }

        return (tokens, preBalances);
    }

    function _executeIntent(SolverNetIntent memory intent) internal returns (uint256 nativeAmountRequired) {
        Call memory call = intent.call;
        address target = _bytes32ToAddress(call.target);

        if (!allowedCalls[target][bytes4(call.callData)]) revert CallNotAllowed();

        (bool success,) = payable(target).call{ value: call.value }(call.callData);
        if (!success) revert CallFailed();

        return call.value;
    }

    /**
     * @dev Returns call hash. Used to discern fullfilment.
     */
    function _callHash(bytes32 srcReqId, bytes memory originData) internal pure returns (bytes32) {
        return keccak256(abi.encode(srcReqId, originData));
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
