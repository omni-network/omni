// SPDX-License-Identifier: GPL-3.0-only
pragma solidity ^0.8.24;

import { OwnableRoles } from "solady/src/auth/OwnableRoles.sol";
import { ReentrancyGuard } from "solady/src/utils/ReentrancyGuard.sol";
import { Initializable } from "solady/src/utils/Initializable.sol";
import { DeployedAt } from "./util/DeployedAt.sol";
import { XAppBase } from "core/src/pkg/XAppBase.sol";
import { OAppCore, OAppSender, MessagingFee } from "./ext/layerzero/oapp/OAppSender.sol";
import { OptionsBuilder } from "./ext/layerzero/libs/OptionsBuilder.sol";
import { ISolverNetOutboxLZ } from "./interfaces/ISolverNetOutboxLZ.sol";
import { SolverNetExecutor } from "./SolverNetExecutor.sol";
import { SafeTransferLib } from "solady/src/utils/SafeTransferLib.sol";
import { FixedPointMathLib } from "solady/src/utils/FixedPointMathLib.sol";
import { ConfLevel } from "core/src/libraries/ConfLevel.sol";
import { TypeMax } from "core/src/libraries/TypeMax.sol";
import { SolverNet } from "./lib/SolverNet.sol";
import { AddrUtils } from "./lib/AddrUtils.sol";
import { ISolverNetInbox } from "./interfaces/ISolverNetInbox.sol";

/**
 * @title SolverNetOutboxLZ
 * @dev This version is capable of sending messages to non-Core chains via LayerZero in addition to Omni Core.
 * @notice Entrypoint for fulfillments of user solve requests.
 */
contract SolverNetOutboxLZ is
    OwnableRoles,
    ReentrancyGuard,
    Initializable,
    DeployedAt,
    XAppBase,
    OAppSender,
    ISolverNetOutboxLZ
{
    using SafeTransferLib for address;
    using OptionsBuilder for bytes;
    using AddrUtils for address;
    using AddrUtils for bytes32;

    /**
     * @notice Role for solvers.
     * @dev _ROLE_0 evaluates to '1'.
     */
    uint256 internal constant SOLVER = _ROLE_0;

    /**
     * @notice Stubbed calldata for SolveInbox.markFilled. Used to estimate the gas cost.
     * @dev Type maxes used to ensure no non-zero bytes in fee estimation.
     */
    bytes internal constant MARK_FILLED_STUB_CDATA =
        abi.encodeCall(ISolverNetInbox.markFilled, (TypeMax.Bytes32, TypeMax.Bytes32, TypeMax.Address));

    /**
     * @notice Addresses of the inbox contracts.
     */
    mapping(uint64 chainId => InboxConfig) internal _inboxes;

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

    constructor(address _endpoint) OAppCore(_endpoint) {
        _disableInitializers();
    }

    /**
     * @notice Initialize the contract's owner and solver.
     * @dev Used instead of constructor as we want to use the transparent upgradeable proxy pattern.
     * @param owner_    Address of the owner.
     * @param solver_   Address of the solver.
     * @param omni_     Address of the OmniPortal.
     * @param executor_ Address of the executor.
     */
    function initialize(address owner_, address solver_, address omni_, address executor_) external initializer {
        _initializeOwner(owner_);
        _grantRoles(solver_, SOLVER);
        _setOmniPortal(omni_);
        _initializeOAppCore(owner_);
        _executor = SolverNetExecutor(payable(executor_));
    }

    /**
     * @notice Initialize the contract's LayerZero delegate.
     */
    function initializeV2() external reinitializer(2) {
        _initializeOAppCore(owner());
        // NOTE: `setInboxes` MUST be called for ALL chain IDs after initialization due to storage layout changes.
    }

    /**
     * @notice Set the inbox addresses for the given chain IDs.
     * @param chainIds IDs of the chains.
     * @param configs  Configurations for the inboxes.
     */
    function setInboxes(uint64[] calldata chainIds, InboxConfig[] calldata configs) external onlyOwner {
        if (chainIds.length != configs.length) revert InvalidArrayLength();
        for (uint256 i; i < chainIds.length; ++i) {
            _inboxes[chainIds[i]] = configs[i];
            emit InboxSet(chainIds[i], configs[i].inbox, configs[i].provider);
        }
    }

    /**
     * @notice Returns the address of the executor contract.
     */
    function executor() external view returns (address) {
        return address(_executor);
    }

    /**
     * @notice Returns the xcall fee required to mark an order filled on the source inbox.
     * @param originData Data emitted on the origin to parameterize the fill.
     * @return           Fee amount in native currency.
     */
    function fillFee(bytes calldata originData) public view returns (uint256) {
        SolverNet.FillOriginData memory fillData = abi.decode(originData, (SolverNet.FillOriginData));
        InboxConfig memory inboxConfig = _inboxes[fillData.srcChainId];

        if (inboxConfig.provider == Provider.OmniCore) {
            return feeFor(fillData.srcChainId, MARK_FILLED_STUB_CDATA, uint64(_fillGasLimit(fillData)));
        } else if (inboxConfig.provider == Provider.LayerZero) {
            bytes memory message = abi.encode(TypeMax.Bytes32, TypeMax.Bytes32, TypeMax.Address);
            bytes memory options =
                OptionsBuilder.newOptions().addExecutorLzReceiveOption(uint128(_fillGasLimit(fillData)), 0);

            MessagingFee memory _fee = _quote(uint32(fillData.srcChainId), message, options, false);
            return _fee.nativeFee;
        } else {
            revert InvalidConfig();
        }
    }

    /**
     * @notice Returns true if the order has been filled.
     * @param orderId    ID of the order the source inbox.
     * @param originData Data emitted on the origin to parameterize the fill
     */
    function didFill(bytes32 orderId, bytes calldata originData) external view returns (bool) {
        return _filled[_fillHash(orderId, originData)];
    }

    /**
     * @notice Fills a particular order on the destination chain.
     * @param orderId    Unique order identifier for this order.
     * @param originData Data emitted on the origin to parameterize the fill.
     * @param fillerData ABI encoded address to mark as claimant for the order.
     */
    function fill(bytes32 orderId, bytes calldata originData, bytes calldata fillerData)
        external
        payable
        onlyRoles(SOLVER)
        nonReentrant
    {
        SolverNet.FillOriginData memory fillData = abi.decode(originData, (SolverNet.FillOriginData));
        address creditTo = msg.sender;

        if (fillData.destChainId != block.chainid) revert WrongDestChain();
        if (fillData.fillDeadline < block.timestamp) revert FillDeadlinePassed();
        if (fillerData.length != 0 && fillerData.length != 32) revert BadFillerData();
        if (fillerData.length == 32) creditTo = abi.decode(fillerData, (address));

        uint256 totalNativeValue = _executeCalls(fillData);
        _markFilled(orderId, fillData, creditTo, totalNativeValue);
    }

    /**
     * @notice Wrap a call with approved / enforced expenses.
     * Approve spenders. Verify post-call balances match pre-call.
     * @dev Expenses doesn't contain native tokens sent alongside the call.
     */
    modifier withExpenses(SolverNet.TokenExpense[] memory expenses) {
        // transfer from solver, approve spenders
        for (uint256 i; i < expenses.length; ++i) {
            SolverNet.TokenExpense memory expense = expenses[i];
            address spender = expense.spender;
            address token = expense.token;
            uint256 amount = expense.amount;

            token.safeTransferFrom(msg.sender, address(_executor), amount);
            // We remotely set token approvals on executor so we don't need to reprocess Call expenses there.
            if (spender != address(0)) _executor.approve(token, spender, amount);
        }

        _;

        // refund excess, revoke approvals
        //
        // NOTE: If anyone transfers this token to this outbox outside
        // SolverNetOutbox.fill(...), the next solver to fill a call with
        // that token as an expense will get the balance.
        // This includes the call target.
        for (uint256 i; i < expenses.length; ++i) {
            SolverNet.TokenExpense memory expense = expenses[i];
            address token = expense.token;
            uint256 tokenBalance = token.balanceOf(address(_executor));

            if (tokenBalance > 0) {
                address spender = expense.spender;
                if (spender != address(0)) _executor.tryRevokeApproval(token, spender);
                _executor.transfer(token, msg.sender, tokenBalance);
            }
        }

        // send any potential native refund sent to executor back to solver
        uint256 nativeBalance = address(_executor).balance;
        if (nativeBalance > 0) _executor.transferNative(msg.sender, nativeBalance);
    }

    /**
     * @notice Verify and execute a call. Expenses are processed and enforced.
     * @param fillData ABI decoded fill originData.
     * @return totalNativeValue total native value of the calls.
     */
    function _executeCalls(SolverNet.FillOriginData memory fillData)
        internal
        withExpenses(fillData.expenses)
        returns (uint256)
    {
        uint256 totalNativeValue;

        for (uint256 i; i < fillData.calls.length; ++i) {
            SolverNet.Call memory call = fillData.calls[i];
            _executor.execute{ value: call.value }(
                call.target, call.value, abi.encodePacked(call.selector, call.params)
            );
            unchecked {
                totalNativeValue += call.value;
            }
        }

        return totalNativeValue;
    }

    /**
     * @notice Mark an order as filled. Require sufficient native payment, refund excess.
     * @param orderId          ID of the order.
     * @param fillData         ABI decoded fill originData.
     * @param claimant         Address specified by the filler to claim the order (msg.sender if none specified).
     * @param totalNativeValue Total native value of the calls.
     */
    function _markFilled(
        bytes32 orderId,
        SolverNet.FillOriginData memory fillData,
        address claimant,
        uint256 totalNativeValue
    ) internal {
        // mark filled on outbox (here)
        bytes32 fillHash = _fillHash(orderId, abi.encode(fillData));
        if (_filled[fillHash]) revert AlreadyFilled();
        _filled[fillHash] = true;

        uint256 fee = _routeMsg(orderId, fillHash, claimant, fillData);
        uint256 totalSpent = totalNativeValue + fee;
        if (msg.value < totalSpent) revert InsufficientFee();

        // refund any overpayment in native currency
        uint256 refund = msg.value - totalSpent;
        if (refund > 0) msg.sender.safeTransferETH(refund);

        emit Filled(orderId, fillHash, msg.sender);
    }

    /**
     * @notice Route a message to the inbox.
     * @param orderId  ID of the order.
     * @param fillHash Hash of the fill instructions origin data.
     * @param claimant Address specified by the filler to claim the order (msg.sender if none specified).
     * @param fillData ABI decoded fill originData.
     * @return fee     Fee amount in native currency.
     */
    function _routeMsg(bytes32 orderId, bytes32 fillHash, address claimant, SolverNet.FillOriginData memory fillData)
        internal
        returns (uint256)
    {
        InboxConfig memory inboxConfig = _inboxes[fillData.srcChainId];
        uint256 fee;

        if (inboxConfig.provider == Provider.OmniCore) {
            // mark filled on inbox
            fee = xcall({
                destChainId: fillData.srcChainId,
                conf: ConfLevel.Finalized,
                to: inboxConfig.inbox,
                data: abi.encodeCall(ISolverNetInbox.markFilled, (orderId, fillHash, claimant)),
                gasLimit: uint64(_fillGasLimit(fillData))
            });
        } else if (inboxConfig.provider == Provider.LayerZero) {
            bytes memory message = abi.encode(orderId, fillHash, claimant);
            bytes memory options =
                OptionsBuilder.newOptions().addExecutorLzReceiveOption(uint128(_fillGasLimit(fillData)), 0);

            MessagingFee memory _fee = _quote(uint32(fillData.srcChainId), message, options, false);
            fee = _fee.nativeFee;

            _lzSend(uint32(fillData.srcChainId), message, options, _fee, msg.sender);
        } else {
            revert InvalidConfig();
        }

        return fee;
    }

    /**
     * @dev Returns call hash. Used to discern fulfillment.
     */
    function _fillHash(bytes32 srcReqId, bytes memory originData) internal pure returns (bytes32) {
        return keccak256(abi.encode(srcReqId, originData));
    }

    /**
     * @notice Returns the gas limit required to mark an order as filled on the source inbox.
     * @param fillData ABI decoded fill originData.
     * @return gasLimit Gas limit for the fill.
     */
    function _fillGasLimit(SolverNet.FillOriginData memory fillData) internal pure returns (uint256) {
        // 2500 gas for the Metadata struct SLOAD.
        uint256 metadataGas = 2500;

        // 2500 gas for Call array length SLOAD + dynamic cost of reading each call.
        uint256 callsGas = 2500;
        for (uint256 i; i < fillData.calls.length; ++i) {
            SolverNet.Call memory call = fillData.calls[i];
            uint256 paramsLength = call.params.length;
            unchecked {
                // 5000 gas for the two slots that hold target, selector, and value.
                // 2500 gas per params slot (1 per function argument) used (minimum of 1 slot).
                callsGas += 5000 + (FixedPointMathLib.divUp(call.params.length + 32, 32) * 2500);
                // Plus memory expansion costs: 3 gas per 32 bytes + bytes^2 / 524288
                callsGas += (3 * FixedPointMathLib.divUp(paramsLength, 32))
                    + FixedPointMathLib.mulDivUp(paramsLength, paramsLength, 524_288);
            }
        }

        // 2500 gas for TokenExpense array length SLOAD + cost of reading each expense.
        uint256 expensesGas = 2500;
        unchecked {
            expensesGas += fillData.expenses.length * 5000;
        }

        return metadataGas + callsGas + expensesGas + 100_000; // 100k base gas limit
    }

    /**
     * @notice Internal function to get the peer address associated with a specific endpoint; never reverts.
     * @param _eid The endpoint ID.
     * @return peer The address of the peer associated with the specified endpoint.
     */
    function _getPeer(uint32 _eid) internal view override returns (bytes32) {
        return _inboxes[_eid].inbox.toBytes32();
    }

    /**
     * @notice Internal function to get the peer address associated with a specific endpoint; reverts if NOT set.
     * ie. the peer is set to bytes32(0).
     * @param _eid The endpoint ID.
     * @return peer The address of the peer associated with the specified endpoint.
     */
    function _getPeerOrRevert(uint32 _eid) internal view override returns (bytes32) {
        bytes32 peer = _getPeer(_eid);
        if (peer == bytes32(0)) revert InvalidInbox();
        return peer;
    }
}
