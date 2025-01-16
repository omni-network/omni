// SPDX-License-Identifier: MIT
pragma solidity =0.8.24;

import { Initializable } from "@openzeppelin/contracts-upgradeable/proxy/utils/Initializable.sol";
import { ERC20Upgradeable } from "@openzeppelin/contracts-upgradeable/token/ERC20/ERC20Upgradeable.sol";
import { UUPSUpgradeable } from "@openzeppelin/contracts-upgradeable/proxy/utils/UUPSUpgradeable.sol";
import { AccessControlUpgradeable } from "@openzeppelin/contracts-upgradeable/access/AccessControlUpgradeable.sol";
import { ERC20PausableUpgradeable } from
    "@openzeppelin/contracts-upgradeable/token/ERC20/extensions/ERC20PausableUpgradeable.sol";
import { AccountPausableUpgradeable } from "./AccountPausableUpgradeable.sol";
import { XAppUpgradeable } from "core/src/pkg/XAppUpgradeable.sol";
import { ConfLevel } from "core/src/libraries/ConfLevel.sol";
import { TypeMax } from "core/src/libraries/TypeMax.sol";
import { SafeTransferLib } from "solady/src/utils/SafeTransferLib.sol";

interface ITokenLockbox {
    function withdraw(address from, address to, uint256 value) external;
}

contract BridgedRLUSDUpgradeable is
    Initializable,
    ERC20Upgradeable,
    UUPSUpgradeable,
    AccessControlUpgradeable,
    ERC20PausableUpgradeable,
    AccountPausableUpgradeable,
    XAppUpgradeable
{
    using SafeTransferLib for address;

    /**
     * @dev Error thrown when an unauthorized crosschain or local mint is attempted.
     */
    error Unauthorized(uint64 chainId, address addr);

    /**
     * @dev Error thrown when an invalid route is attempted.
     */
    error InvalidRoute(uint64 chainId);

    /**
     * @dev Error thrown when an insufficient amount of native payment is provided to pay for a crosschain transfer.
     */
    error InsufficientFunds();

    /**
     * @dev Error thrown when the length of the array of chainIds and token addresses do not match.
     */
    error ArrayLengthMismatch();

    /**
     * @dev Event emitted when a route is configured.
     */
    event RouteConfigured(uint64 chainId, address addr);

    /**
     * @dev Event emitted when a crosschain transfer is sent.
     */
    event CrosschainTxSent(uint64 destChainId, address from, address to, uint256 value);

    /**
     * @dev Event emitted when a crosschain transfer is received.
     */
    event CrosschainTxReceived(uint64 srcChainId, address from, address to, uint256 value);

    /**
     * @dev Gas limit for a crosschain transfer.
     */
    uint64 public constant TRANSFER_GAS_LIMIT = 50_000;

    //keccak256("CONFIGURATOR")
    bytes32 public constant CONFIGURATOR_ROLE = 0x530008d2b058137d9c475b1b7d83984f1fcf1dd0e607659d029fc1517ab89268;
    //keccak256("UPGRADER")
    bytes32 public constant UPGRADER_ROLE = 0xa615a8afb6fffcb8c6809ac0997b5c9c12b8cc97651150f14c8f6203168cff4c;

    /**
     * @dev Address of the lockbox contract on Ethereum mainnet.
     */
    address public lockbox;

    /**
     * @dev Mapping of chainId to RLUSD.e contract address.
     */
    mapping(uint64 chainId => address addr) public routes;

    /**
     * @dev Modifier to check if the caller is another known RLUSD.e contract.
     */
    modifier onlyApprovedRoutes() {
        if (msg.sender == address(omni)) {
            if (routes[xmsg.sourceChainId] != xmsg.sender) revert Unauthorized(xmsg.sourceChainId, xmsg.sender);
        } else {
            revert Unauthorized(uint64(block.chainid), msg.sender);
        }
        _;
    }

    /// @custom:oz-upgrades-unsafe-allow constructor
    constructor() {
        _disableInitializers();
    }

    /**
     * @dev This method is used to initialize the contract with values that we want to use to bootstrap and run things.
     * The modifier initializer here helps us block initialization in the constructor so that we initialize value only
     * when deploying the proxy and not the contract itself. The initializer also tracks how many times this method is
     * called and it can only be called once.
     */
    function initialize(address _admin, address _configurator, address _upgrader, address _omni, address _lockbox)
        external
        initializer
    {
        __ERC20_init("Bridged RLUSD (Omni)", "RLUSD.e");
        __UUPSUpgradeable_init();
        __AccessControl_init();
        __ERC20Pausable_init();
        __XApp_init(_omni, ConfLevel.Finalized);

        _grantRole(DEFAULT_ADMIN_ROLE, _admin);
        _grantRole(CONFIGURATOR_ROLE, _configurator);
        _grantRole(UPGRADER_ROLE, _upgrader);
        lockbox = _lockbox;
        routes[1] = _lockbox;
    }

    /**
     * @dev Sets the addresses of the token contracts on specific chainIds.
     * @param chainIds The chainIds of the chains to set the token address for.
     * @param tokenAddrs The addresses of the token contracts to set.
     */
    function setRoute(uint64[] calldata chainIds, address[] calldata tokenAddrs) external onlyRole(CONFIGURATOR_ROLE) {
        if (chainIds.length != tokenAddrs.length) revert ArrayLengthMismatch();
        for (uint256 i = 0; i < chainIds.length; i++) {
            routes[chainIds[i]] = tokenAddrs[i];
            emit RouteConfigured(chainIds[i], tokenAddrs[i]);
        }
    }

    /**
     * @dev Creates a `value` amount of tokens and assigns them to `to`, by transferring it from address(0).
     * Relies on the `_update` mechanism. Only callable by RLUSD.e contracts configured in the `routes` mapping.
     *
     * @param from  The address that is transferring the tokens.
     * @param to    The address that will be receiving the minted amount.
     * @param value The amount of tokens that are being minted to the account.
     *
     * Emits a {Transfer} event with `source` set to the zero address.
     */
    function mint(address from, address to, uint256 value) external xrecv onlyApprovedRoutes {
        _mint(to, value);
        emit CrosschainTxReceived(xmsg.sourceChainId, from, to, value);
    }

    /**
     * @dev Returns the native fee for a crosschain transfer.
     * @param destChainId The chainId of the destination chain.
     * @return fee The native fee for the crosschain transfer.
     */
    function feeForCrosschainTransfer(uint64 destChainId) public view returns (uint256 fee) {
        return feeFor(
            destChainId,
            _getCrosschainTransferData(destChainId, TypeMax.Address, TypeMax.Address, TypeMax.Uint256),
            TRANSFER_GAS_LIMIT
        );
    }

    /**
     * @dev Performs a crosschain transfer
     * @param destChainId The chainId of the destination chain.
     * @param from The address that is transferring the tokens.
     * @param to The address that will be receiving the tokens.
     * @param value The amount of tokens to transfer.
     */
    function crosschainTransfer(uint64 destChainId, address from, address to, uint256 value) external payable {
        address destContract = routes[destChainId];
        if (destContract == address(0)) revert InvalidRoute(destChainId);
        if (from == address(0)) from = msg.sender;
        if (to == address(0)) to = msg.sender;
        if (from != msg.sender) _spendAllowance(from, msg.sender, value);

        _burn(from, value);

        bytes memory data = _getCrosschainTransferData(destChainId, from, to, value);
        uint256 fee = xcall(destChainId, destContract, data, TRANSFER_GAS_LIMIT);

        if (msg.value < fee) revert InsufficientFunds();
        if (msg.value > fee) msg.sender.safeTransferETH(msg.value - fee);

        emit CrosschainTxSent(destChainId, from, to, value);
    }

    /**
     * @dev Returns the calldata to be used for a crosschain transfer.
     * @param destChainId The chainId of the destination chain.
     * @param from The address that is transferring the tokens.
     * @param to The address to transfer the tokens to.
     * @param value The amount of tokens to transfer.
     * @return data The calldata to be used for the crosschain transfer.
     */
    function _getCrosschainTransferData(uint64 destChainId, address from, address to, uint256 value)
        internal
        pure
        returns (bytes memory data)
    {
        if (destChainId == 1) return abi.encodeCall(ITokenLockbox.withdraw, (from, to, value));
        return abi.encodeCall(BridgedRLUSDUpgradeable.mint, (from, to, value));
    }

    /**
     * An overridden method from {ERC20Upgradeable} and {ERC20PausableUpgradeable} without Pausable mechanics enabled.
     */
    function _update(address from, address to, uint256 value)
        internal
        override(ERC20Upgradeable, ERC20PausableUpgradeable)
    {
        ERC20Upgradeable._update(from, to, value);
    }

    /**
     * An overridden method from {UUPSUpgradeable} which defines the permissions for authorizing an upgrade to a
     * new implementation.
     */
    function _authorizeUpgrade(address newImplementation) internal virtual override onlyRole(UPGRADER_ROLE) { }
}
