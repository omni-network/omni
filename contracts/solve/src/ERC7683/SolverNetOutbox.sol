// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.24;

import { OwnableRoles } from "solady/src/auth/OwnableRoles.sol";
import { ReentrancyGuard } from "solady/src/utils/ReentrancyGuard.sol";
import { Initializable } from "solady/src/utils/Initializable.sol";
import { XAppBase } from "core/src/pkg/XAppBase.sol";
import { SafeTransferLib } from "solady/src/utils/SafeTransferLib.sol";
import { ConfLevel } from "core/src/libraries/ConfLevel.sol";
import { TypeMax } from "core/src/libraries/TypeMax.sol";
import { DeployedAt } from "src/util/DeployedAt.sol";
import { ISolverNetInbox } from "./interfaces/ISolverNetInbox.sol";
import { ISolverNetOutbox } from "./interfaces/ISolverNetOutbox.sol";

/**
 * @title SolverNetOutbox
 * @notice Entrypoint for fillments of user solve requests.
 */
contract SolverNetOutbox is OwnableRoles, ReentrancyGuard, Initializable, DeployedAt, XAppBase, ISolverNetOutbox {
    using SafeTransferLib for address;

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
     * @notice Maps fillHash (hash of fill instruction origin data) to true, if filled.
     * @dev Used to prevent duplicate fillment.
     */
    mapping(bytes32 fillHash => bool filled) internal _filled;

    /**
     * @notice Mapping of allowed calls per contract.
     */
    mapping(address target => mapping(bytes4 selector => bool)) public allowedCalls;

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
        FillOriginData memory fillData = abi.decode(originData, (FillOriginData));
        Call memory call = fillData.call;

        _executeCall(call);
        _markFilled(orderId, fillData.srcChainId, call, _fillHash(orderId, originData));
    }

    /**
     * @notice Wrap a call with approved / enforced expenses.
     *  Approve spenders. Verify post-call balances match pre-call.
     */
    modifier withExpenses(TokenExpense[] memory expenses) {
        address[] memory tokens = new address[](expenses.length);
        uint256[] memory preBalances = new uint256[](expenses.length);

        for (uint256 i; i < expenses.length; ++i) {
            TokenExpense memory expense = expenses[i];
            address token = _bytes32ToAddress(expense.token);

            tokens[i] = token;
            preBalances[i] = token.balanceOf(address(this));

            address spender = _bytes32ToAddress(expense.spender);
            token.safeTransferFrom(msg.sender, address(this), expense.amount);
            token.safeApprove(spender, expense.amount);
        }

        _;

        for (uint256 i; i < tokens.length; ++i) {
            if (tokens[i].balanceOf(address(this)) != preBalances[i]) revert InvalidExpenses();
        }
    }

    /**
     * @notice Verifiy and execute a call. Expenses are processed and enforced.
     * @param call  Call to execute.
     */
    function _executeCall(Call memory call) internal withExpenses(call.expenses) {
        if (call.destChainId != block.chainid) revert WrongDestChain();

        address target = _bytes32ToAddress(call.target);

        if (!allowedCalls[target][bytes4(call.data)]) revert CallNotAllowed();

        (bool success,) = payable(target).call{ value: call.value }(call.data);
        if (!success) revert CallFailed();
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
     * @dev Returns call hash. Used to discern fullfilment.
     */
    function _fillHash(bytes32 srcReqId, bytes memory originData) internal pure returns (bytes32) {
        return keccak256(abi.encode(srcReqId, originData));
    }

    /**
     * @dev Convert bytes32 to address.
     */
    function _bytes32ToAddress(bytes32 b) internal pure returns (address) {
        return address(uint160(uint256(b)));
    }
}
