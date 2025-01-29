// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.24;

import { OwnableRoles } from "solady/src/auth/OwnableRoles.sol";
import { ReentrancyGuard } from "solady/src/utils/ReentrancyGuard.sol";
import { Initializable } from "solady/src/utils/Initializable.sol";
import { XAppBase } from "core/src/pkg/XAppBase.sol";
import { SolverNetExecutor } from "./SolverNetExecutor.sol";
import { SafeTransferLib } from "solady/src/utils/SafeTransferLib.sol";
import { ConfLevel } from "core/src/libraries/ConfLevel.sol";
import { TypeMax } from "core/src/libraries/TypeMax.sol";
import { DeployedAt } from "./util/DeployedAt.sol";
import { AddrUtils } from "./lib/AddrUtils.sol";
import { ISolverNetInbox } from "./interfaces/ISolverNetInbox.sol";
import { ISolverNetOutbox } from "./interfaces/ISolverNetOutbox.sol";

/**
 * @title SolverNetOutbox
 * @notice Entrypoint for fulfillments of user solve requests.
 */
contract SolverNetOutbox is OwnableRoles, ReentrancyGuard, Initializable, DeployedAt, XAppBase, ISolverNetOutbox {
    using SafeTransferLib for address;
    using AddrUtils for bytes32;

    /**
     * @notice Role for solvers.
     * @dev _ROLE_0 evaluates to '1'.
     */
    uint256 internal constant SOLVER = _ROLE_0;

    /**
     * @notice Gas limit for SolveInbox.markFilled callback.
     */
    uint64 internal constant MARK_FILLED_GAS_LIMIT = 100_000;

    /**
     * @notice Stubbed calldata for SolveInbox.markFilled. Used to estimate the gas cost.
     * @dev Type maxes used to ensure no non-zero bytes in fee estimation.
     */
    bytes internal constant MARK_FILLED_STUB_CDATA =
        abi.encodeCall(ISolverNetInbox.markFilled, (TypeMax.Bytes32, TypeMax.Bytes32));

    /**
     * @notice Address of the inbox contract.
     */
    address internal _inbox;

    /**
     * @notice Executor contract handling calls.
     * @dev An executor is used so infinite approvals from solvers cannot be abused.
     */
    SolverNetExecutor internal _executor;

    /**
     * @notice Maps fillHash (hash of fill instruction origin data) to true, if filled.
     * @dev Used to prevent duplicate fulfillment.
     */
    mapping(bytes32 fillHash => bool filled) internal _filled;

    constructor() {
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
        _executor = new SolverNetExecutor(address(this));
    }

    /**
     * @notice Returns the address of the executor contract.
     */
    function executor() external view returns (address) {
        return address(_executor);
    }

    /**
     * @notice Returns the xcall fee required to mark an order filled on the source inbox
     * @param srcChainId  Chain ID on which the order was opened.
     * @return            Fee amount in native currency.
     */
    function fillFee(uint64 srcChainId) public view returns (uint256) {
        return feeFor(srcChainId, MARK_FILLED_STUB_CDATA, MARK_FILLED_GAS_LIMIT);
    }

    /**
     * @notice Returns true if the order has been filled.
     * @param orderId     ID of the order the source inbox.
     * @param originData  Data emitted on the origin to parameterize the fill
     */
    function didFill(bytes32 orderId, bytes calldata originData) external view returns (bool) {
        return _filled[_fillHash(orderId, originData)];
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
        FillOriginData memory fillData = abi.decode(originData, (FillOriginData));
        Call memory call = fillData.call;

        if (fillData.fillDeadline < block.timestamp && fillData.fillDeadline != 0) revert FillDeadlinePassed();

        _executeCall(call);
        _markFilled(orderId, fillData.srcChainId, call, _fillHash(orderId, originData));
    }

    /**
     * @notice Wrap a call with approved / enforced expenses.
     *  Approve spenders. Verify post-call balances match pre-call.
     */
    modifier withExpenses(TokenExpense[] memory expenses) {
        // transfer from solver, approve spenders
        for (uint256 i; i < expenses.length; ++i) {
            TokenExpense memory expense = expenses[i];
            address token = expense.token.toAddress();
            address spender = expense.spender.toAddress();

            token.safeTransferFrom(msg.sender, address(_executor), expense.amount);
            // We remotely set token approvals on executor so we don't need to reprocess Call expenses there.
            if (spender != address(0)) _executor.approve(token, spender, expense.amount);
        }

        _;

        // refund excess, revoke approvals
        //
        // NOTE: If anyone transfers this token to this outbox outside
        // SolverNetOutbox.fill(...), the next solver to fill a call with
        // that token as an expense will get the balance.
        // This includes the call target.
        for (uint256 i; i < expenses.length; ++i) {
            TokenExpense memory expense = expenses[i];
            address token = expense.token.toAddress();
            uint256 tokenBalance = token.balanceOf(address(_executor));

            if (tokenBalance > 0) {
                if (expense.spender != bytes32(0)) _executor.approve(token, expense.spender.toAddress(), 0);
                _executor.transfer(token, msg.sender, tokenBalance);
            }
        }

        // send any potential native refund sent to executor back to solver
        uint256 nativeBalance = address(_executor).balance;
        if (nativeBalance > 0) _executor.transferNative(msg.sender, nativeBalance);
    }

    /**
     * @notice Verify and execute a call. Expenses are processed and enforced.
     * @param call  Call to execute.
     */
    function _executeCall(Call memory call) internal withExpenses(call.expenses) {
        if (call.chainId != block.chainid) revert WrongDestChain();

        _executor.execute{ value: call.value }(call);
    }

    /**
     * @notice Mark an order as filled. Require sufficient native payment, refund excess.
     * @param orderId     ID of the order.
     * @param srcChainId  Chain ID on which the order was opened.
     * @param call        Call executed.
     * @param fillHash    Hash of fill data, verifies fill matches order.
     */
    function _markFilled(bytes32 orderId, uint64 srcChainId, Call memory call, bytes32 fillHash) internal {
        // mark filled on outbox (here)
        if (_filled[fillHash]) revert AlreadyFilled();
        _filled[fillHash] = true;

        // mark filled on inbox
        uint256 fee = xcall({
            destChainId: srcChainId,
            conf: ConfLevel.Finalized,
            to: _inbox,
            data: abi.encodeCall(ISolverNetInbox.markFilled, (orderId, fillHash)),
            gasLimit: MARK_FILLED_GAS_LIMIT
        });
        if (msg.value - call.value < fee) revert InsufficientFee();

        // refund any overpayment in native currency
        uint256 refund = msg.value - call.value - fee;
        if (refund > 0) msg.sender.safeTransferETH(refund);

        emit Filled(orderId, fillHash, msg.sender);
    }

    /**
     * @dev Returns call hash. Used to discern fulfillment.
     */
    function _fillHash(bytes32 srcReqId, bytes memory originData) internal pure returns (bytes32) {
        return keccak256(abi.encode(srcReqId, originData));
    }
}
