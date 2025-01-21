// SPDX-License-Identifier: MIT
pragma solidity 0.8.26;

interface IBridgeUpgradeable {
    /*´:°•.°+.*•´.*:˚.°*.˚•´.°:°•.°•.*•´.*:˚.°*.˚•´.°:°•.°+.*•´.*:*/
    /*                           ERRORS                           */
    /*.•°:°.´+˚.*°.˚:*.´•*.+°.•°:´*.´•*.•°.•°:°.´:•˚°.*°.˚:*.´+°.•*/

    /**
     * @dev Error thrown when an unauthorized crosschain or local mint is attempted.
     */
    error Unauthorized(uint64 chainId, address addr);

    /**
     * @dev Thrown when a destChainId matches the local chainId.
     */
    error InvalidChainId();

    /**
     * @dev Error thrown when an invalid token route is attempted.
     */
    error InvalidTokenRoute(address srcToken, uint64 destChainId);

    /**
     * @dev Error thrown when the length of the array elements do not match.
     */
    error ArrayLengthMismatch();

    /**
     * @dev Error thrown when an insufficient amount of native payment is provided to pay for a crosschain transfer.
     */
    error InsufficientFunds();

    /**
     * @dev Error thrown when an invalid address is attempted.
     */
    error ZeroAddress();

    /**
     * @dev Error thrown when an invalid amount is attempted.
     */
    error ZeroAmount();

    /*´:°•.°+.*•´.*:˚.°*.˚•´.°:°•.°•.*•´.*:˚.°*.˚•´.°:°•.°+.*•´.*:*/
    /*                           EVENTS                           */
    /*.•°:°.´+˚.*°.˚:*.´•*.+°.•°:´*.´•*.•°.•°:°.´:•˚°.*°.˚:*.´+°.•*/

    /**
     * @dev Event emitted when a bridge is configured.
     */
    event BridgeConfigured(uint64 indexed chainId, address indexed bridge);

    /**
     * @dev Event emitted when a token route and native status is configured.
     */
    event TokenConfigured(
        address indexed srcToken, bool isNative, uint64 indexed destChainId, address indexed destToken
    );

    /**
     * @dev Event emitted when a crosschain token transfer is initiated.
     */
    event TokenSent(
        uint64 indexed destChainId, address indexed srcToken, address destToken, address indexed to, uint256 value
    );

    /**
     * @dev Event emitted when a crosschain token transfer is received.
     */
    event TokenReceived(uint64 indexed srcChainId, address indexed destToken, address indexed to, uint256 value);

    /*´:°•.°+.*•´.*:˚.°*.˚•´.°:°•.°•.*•´.*:˚.°*.˚•´.°:°•.°+.*•´.*:*/
    /*                       VIEW FUNCTIONS                       */
    /*.•°:°.´+˚.*°.˚:*.´•*.+°.•°:´*.´•*.•°.•°:°.´:•˚°.*°.˚:*.´+°.•*/

    /**
     * @dev Mapping of destination chainId to bridge contract.
     * @param chainId The chainId of the destination chain.
     * @return bridge The bridge contract.
     */
    function bridgeRoutes(uint64 chainId) external view returns (address bridge);

    /**
     * @dev Mapping of token to whether it is the native representation of an ERC20 token.
     * @param token The token to check.
     * @return isNative Whether the token is the native representation of an ERC20 token.
     */
    function isNativeToken(address token) external view returns (bool isNative);

    /**
     * @dev Mapping of source token to destination chainId to destination token.
     * @param srcToken The source token.
     * @param destChainId The destination chainId.
     * @return destToken The destination token.
     */
    function tokenRoutes(address srcToken, uint64 destChainId) external view returns (address destToken);

    /**
     * @dev Returns the fee for bridging a token to a destination chain.
     * @param destChainId The chainId of the destination chain.
     * @return fee The fee paid to the `OmniPortal` contract.
     */
    function bridgeFee(uint64 destChainId) external view returns (uint256 fee);

    /*´:°•.°+.*•´.*:˚.°*.˚•´.°:°•.°•.*•´.*:˚.°*.˚•´.°:°•.°+.*•´.*:*/
    /*                      BRIDGE FUNCTIONS                      */
    /*.•°:°.´+˚.*°.˚:*.´•*.+°.•°:´*.´•*.•°.•°:°.´:•˚°.*°.˚:*.´+°.•*/

    /**
     * @dev Bridges a token to a destination chain.
     * @param destChainId The chainId of the destination chain.
     * @param token The address of the source token.
     * @param to The address of the recipient.
     * @param value The amount of tokens to bridge.
     */
    function sendToken(uint64 destChainId, address token, address to, uint256 value) external payable;

    /**
     * @dev Receives a token from a bridge contract and mints it to the recipient.
     * @param token The address of the source token.
     * @param to The address of the recipient.
     * @param value The amount of tokens to mint.
     */
    function receiveToken(address token, address to, uint256 value) external;

    /*´:°•.°+.*•´.*:˚.°*.˚•´.°:°•.°•.*•´.*:˚.°*.˚•´.°:°•.°+.*•´.*:*/
    /*                      ADMIN FUNCTIONS                       */
    /*.•°:°.´+˚.*°.˚:*.´•*.+°.•°:´*.´•*.•°.•°:°.´:•˚°.*°.˚:*.´+°.•*/

    /**
     * @dev Configures bridges for a given chainId.
     * @param chainIds The chainIds to configure.
     * @param bridges The bridges to configure.
     */
    function configureBridges(uint64[] calldata chainIds, address[] calldata bridges) external;

    /**
     * @dev Configures token routes for a given source token and destination chainId.
     * @param srcToken The source token to configure.
     * @param isNative Whether the source token is the native representation of an ERC20 token.
     * @param destChainIds The destination chainIds to configure routes for.
     * @param destTokens The destination token addresses for the source token on each destination chainId.
     */
    function configureToken(
        address srcToken,
        bool isNative,
        uint64[] calldata destChainIds,
        address[] calldata destTokens
    ) external;
}
