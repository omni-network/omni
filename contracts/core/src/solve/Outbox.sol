// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.24;

import { OwnableRoles } from "solady/src/auth/OwnableRoles.sol";
import { ReentrancyGuard } from "solady/src/utils/ReentrancyGuard.sol";
import { Initializable } from "solady/src/utils/Initializable.sol";
import { XAppBase } from "../pkg/XAppBase.sol";
import { SafeTransferLib } from "solady/src/utils/SafeTransferLib.sol";
import { ConfLevel } from "../libraries/ConfLevel.sol";
import { Solve } from "./Solve.sol";

interface ISymbioticVault {
    function collateral() external view returns (address);
    function deposit(address onBehalfOf, uint256 amount) external;
}

/**
 * @title Outbox
 * @notice Entrypoint for fulfillments of user solve requests.
 */
contract Outbox is OwnableRoles, ReentrancyGuard, Initializable, XAppBase {
    using SafeTransferLib for address;

    error CallFailed();
    error IncorrectChain();
    error InsufficientFee();
    error IncorrectBalance();
    error UnauthorizedCall();

    event AllowedCallSet(address indexed target, bytes4 indexed selector, bool allowed);

    /**
     * @notice Emitted when a request is fulfilled.
     * @param guid        ID of the request.
     * @param callHash    Hash of the call executed.
     * @param creditedTo  Origin address credited the funds by the solver.
     */
    event Fulfilled(bytes32 indexed guid, bytes32 indexed callHash, address indexed creditedTo);

    /**
     * @notice Role for solvers.
     * @dev _ROLE_0 evaluates to '1'.
     */
    uint256 internal constant SOLVER = _ROLE_0;

    /**
     * @notice Address of the inbox contract.
     */
    address internal _inbox;

    /**
     * @notice Mapping of allowed calls per contract.
     */
    mapping(address target => mapping(bytes4 selector => bool)) public allowedCalls;

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

    function setAllowedCall(address target, bytes4 selector, bool allowed) external onlyOwner {
        allowedCalls[target][selector] = allowed;
        emit AllowedCallSet(target, selector, allowed);
    }

    /**
     * @notice Calculate the message passing fee for a fulfill call.
     * @param sourceChainId ID of the source chain.
     */
    function fulfillFee(uint64 sourceChainId) public view returns (uint256) {
        bytes memory data = abi.encodeWithSignature(
            "fulfill(bytes32,bytes32,address)",
            keccak256(abi.encode("GUID")),
            keccak256(abi.encode("CALLHASH")),
            address(uint160(uint256(keccak256(abi.encode("SOLVER")))))
        );
        return feeFor(sourceChainId, data, 200_000);
    }

    /**
     * @notice Fulfill a request.
     * @param guid           ID of the request.
     * @param sourceChainId  ID of the source chain.
     * @param creditTo       Address to credit funds to on the origin chain.
     * @param call           Details of the call to execute.
     * @param prereqs        Pre-requisite token deposits required by the call.
     */
    function fulfill(
        bytes32 guid,
        uint64 sourceChainId,
        address creditTo,
        Solve.Call calldata call,
        Solve.TokenPrereq[] calldata prereqs
    ) external payable onlyRoles(SOLVER) nonReentrant {
        // Verify the call is intended for this chain and is authorized
        if (call.destChainId != block.chainid) revert IncorrectChain();
        if (!allowedCalls[call.target][bytes4(call.data)]) revert UnauthorizedCall();

        // Process pre-requisite deposits
        uint256[] memory prereqBalances = new uint256[](prereqs.length);
        for (uint256 i; i < prereqs.length; ++i) {
            prereqBalances[i] = prereqs[i].token.balanceOf(address(this));
            prereqs[i].token.safeTransferFrom(msg.sender, address(this), prereqs[i].amount);
            prereqs[i].token.safeApprove(prereqs[i].spender, prereqs[i].amount);
        }

        // Execute the call
        (bool success,) = payable(call.target).call{ value: call.value }(call.data);
        if (!success) revert CallFailed();

        // Post check balances to ensure token prereqs were properly used
        for (uint256 i; i < prereqs.length; ++i) {
            if (prereqs[i].token.balanceOf(address(this)) != prereqBalances[i]) revert IncorrectBalance();
        }

        // Calculate the call hash and prepare fulfillment call
        bytes32 callHash = keccak256(abi.encode(guid, call));
        bytes memory data = abi.encodeWithSignature("markFulfilled(bytes32,bytes32,address)", guid, callHash, creditTo);

        // Send the fulfillment call to the inbox
        uint256 fee = xcall(sourceChainId, ConfLevel.Finalized, _inbox, data, 200_000);
        if (msg.value - call.value < fee) revert InsufficientFee();

        emit Fulfilled(guid, callHash, creditTo);
    }
}
