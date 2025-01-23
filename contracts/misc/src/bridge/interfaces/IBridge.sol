// SPDX-License-Identifier: MIT
pragma solidity 0.8.26;

interface IBridge {
    /*´:°•.°+.*•´.*:˚.°*.˚•´.°:°•.°•.*•´.*:˚.°*.˚•´.°:°•.°+.*•´.*:*/
    /*                           ERRORS                           */
    /*.•°:°.´+˚.*°.˚:*.´•*.+°.•°:´*.´•*.•°.•°:°.´:•˚°.*°.˚:*.´+°.•*/

    /**
     * @dev Error thrown when no token and lockbox are set.
     */
    error CannotWrap();

    /**
     * @dev Error thrown when the amount is zero.
     */
    error ZeroAmount();

    /**
     * @dev Error thrown when the address is zero.
     */
    error ZeroAddress();

    /**
     * @dev Error thrown when the destination chainId is invalid.
     */
    error InvalidRoute(uint64 chainId);

    /**
     * @dev Error thrown when the caller is not authorized.
     */
    error Unauthorized(uint64 chainId, address addr);

    /**
     * @dev Error thrown when the funds are insufficient.
     */
    error InsufficientPayment();

    /**
     * @dev Error thrown when the array lengths mismatch.
     */
    error ArrayLengthMismatch();

    /*´:°•.°+.*•´.*:˚.°*.˚•´.°:°•.°•.*•´.*:˚.°*.˚•´.°:°•.°+.*•´.*:*/
    /*                           EVENTS                           */
    /*.•°:°.´+˚.*°.˚:*.´•*.+°.•°:´*.´•*.•°.•°:°.´:•˚°.*°.˚:*.´+°.•*/

    /**
     * @dev Event emitted when a bridge route is configured.
     */
    event BridgeConfigured(uint64 indexed destChainId, address indexed bridge);

    /**
     * @dev Event emitted when a crosschain token transfer is initiated.
     */
    event TokenSent(uint64 indexed destChainId, address indexed from, address indexed to, uint256 value);

    /**
     * @dev Event emitted when a crosschain token transfer is received.
     */
    event TokenReceived(uint64 indexed srcChainId, address indexed to, uint256 value);

    /**
     * @dev Event emitted when a lockbox withdrawal fails.
     */
    event LockboxWithdrawalFailed(address indexed badLockbox, address indexed to, uint256 value);

    /*´:°•.°+.*•´.*:˚.°*.˚•´.°:°•.°•.*•´.*:˚.°*.˚•´.°:°•.°+.*•´.*:*/
    /*                          STORAGE                           */
    /*.•°:°.´+˚.*°.˚:*.´•*.+°.•°:´*.´•*.•°.•°:°.´:•˚°.*°.˚:*.´+°.•*/

    /**
     * @dev Address of the token being bridged.
     */
    function token() external view returns (address);

    /**
     * @dev Lockbox (if assigned) indicating where the token is wrapped.
     */
    function lockbox() external view returns (address);

    /*´:°•.°+.*•´.*:˚.°*.˚•´.°:°•.°•.*•´.*:˚.°*.˚•´.°:°•.°+.*•´.*:*/
    /*                       VIEW FUNCTIONS                       */
    /*.•°:°.´+˚.*°.˚:*.´•*.+°.•°:´*.´•*.•°.•°:°.´:•˚°.*°.˚:*.´+°.•*/

    /**
     * @dev Returns the bridge address for a given destination chainId.
     * @param destChainId The chainId of the destination chain.
     * @return bridge     The bridge address.
     */
    function routes(uint64 destChainId) external view returns (address bridge);

    /**
     * @dev Returns the fee for bridging a token to a destination chain.
     * @param destChainId The chainId of the destination chain.
     * @return fee        The fee paid to the `OmniPortal` contract.
     */
    function bridgeFee(uint64 destChainId) external view returns (uint256 fee);

    /*´:°•.°+.*•´.*:˚.°*.˚•´.°:°•.°•.*•´.*:˚.°*.˚•´.°:°•.°+.*•´.*:*/
    /*                      BRIDGE FUNCTIONS                      */
    /*.•°:°.´+˚.*°.˚:*.´•*.+°.•°:´*.´•*.•°.•°:°.´:•˚°.*°.˚:*.´+°.•*/

    /**
     * @dev Bridges a token to a destination chain.
     * @param destChainId The chainId of the destination chain.
     * @param to          The address of the recipient.
     * @param value       The amount of tokens to bridge.
     * @param wrap        Whether to wrap the token first.
     */
    function sendToken(uint64 destChainId, address to, uint256 value, bool wrap) external payable;

    /**
     * @dev Receives a token from a bridge contract and mints/unwraps it to the recipient.
     * @param to    The address of the recipient.
     * @param value The amount of tokens to mint/unwrap.
     */
    function receiveToken(address to, uint256 value) external;

    /*´:°•.°+.*•´.*:˚.°*.˚•´.°:°•.°•.*•´.*:˚.°*.˚•´.°:°•.°+.*•´.*:*/
    /*                      ADMIN FUNCTIONS                       */
    /*.•°:°.´+˚.*°.˚:*.´•*.+°.•°:´*.´•*.•°.•°:°.´:•˚°.*°.˚:*.´+°.•*/

    /**
     * @dev Sets bridge routes for given chainIds.
     * @param chainIds    The chainIds to configure.
     * @param bridgeAddrs The bridges addresses to configure.
     */
    function setRoutes(uint64[] calldata chainIds, address[] calldata bridgeAddrs) external;
}
