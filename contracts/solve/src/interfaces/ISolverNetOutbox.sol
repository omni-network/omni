// SPDX-License-Identifier: GPL-3.0-only
pragma solidity ^0.8.24;

import { IDestinationSettler } from "../erc7683/IDestinationSettler.sol";

interface ISolverNetOutbox is IDestinationSettler {
    error NotFilled();
    error BadFillerData();
    error AlreadyFilled();
    error InvalidConfig();
    error WrongDestChain();
    error InsufficientFee();
    error InvalidSettlement();
    error FillDeadlinePassed();
    error InvalidArrayLength();

    /**
     * @notice Emitted when an inbox is set.
     * @param chainId  ID of the chain.
     * @param inbox    Address of the inbox.
     * @param provider The messaging provider used to reach the inbox.
     */
    event InboxSet(uint64 indexed chainId, address indexed inbox, Provider indexed provider);

    /**
     * @notice Emitted when a cross-chain request is filled on the destination chain
     * @param orderId  ID of the order on the source chain
     * @param fillHash Hash of the fill origin data
     * @param filledBy Address of the solver that filled the oder
     */
    event Filled(bytes32 indexed orderId, bytes32 indexed fillHash, address indexed filledBy);

    /**
     * @notice Emitted when a markFilled settlement is retried.
     * @param orderId   ID of the order on the source chain
     * @param fillHash  Hash of the fill origin data
     * @param retriedBy Address that retried the markFilled settlement
     */
    event MarkFilledRetry(bytes32 indexed orderId, bytes32 indexed fillHash, address indexed retriedBy);

    /**
     * @notice The messaging provider used to reach the inbox.
     */
    enum Provider {
        None,
        OmniCore,
        Hyperlane,
        Trusted
    }

    /**
     * @notice The configuration for an inbox.
     * @param inbox    Address of the inbox.
     * @param provider The messaging provider used to reach the inbox.
     */
    struct InboxConfig {
        address inbox;
        Provider provider;
    }

    /**
     * @notice Set the inbox addresses for the given chain IDs.
     * @param chainIds IDs of the chains.
     * @param configs  Configurations for the inboxes.
     */
    function setInboxes(uint64[] calldata chainIds, InboxConfig[] calldata configs) external;

    /**
     * @notice Returns the inbox configuration for the given chain ID.
     * @param chainId ID of the chain.
     * @return config Inbox configuration.
     */
    function getInboxConfig(uint64 chainId) external view returns (InboxConfig memory);

    /**
     * @notice Returns the address of the executor contract.
     */
    function executor() external view returns (address);

    /**
     * @notice Returns the xcall fee required to mark an order filled on the source inbox.
     * @param originData Data emitted on the origin to parameterize the fill.
     * @return           Fee amount in native currency.
     */
    function fillFee(bytes calldata originData) external view returns (uint256);

    /**
     * @notice Returns whether a call has been filled.
     * @param srcReqId   ID of the on the source inbox.
     * @param originData Data emitted on the origin to parameterize the fill
     * @return           Whether the call has been filled.
     */
    function didFill(bytes32 srcReqId, bytes calldata originData) external view returns (bool);

    /**
     * @notice Retry marking an order as filled on the source inbox.
     * @param orderId    ID of the order.
     * @param originData Data emitted on the origin to parameterize the fill.
     * @param fillerData ABI encoded address to mark as claimant for the order.
     */
    function retryMarkFilled(bytes32 orderId, bytes calldata originData, bytes calldata fillerData) external payable;
}
