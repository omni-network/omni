// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.24;

import { IDestinationSettler } from "../erc7683/IDestinationSettler.sol";

interface ISolverNetOutbox is IDestinationSettler {
    error BadFillerData();
    error AlreadyFilled();
    error WrongDestChain();
    error InsufficientFee();
    error FillDeadlinePassed();

    /**
     * @notice Emitted when an inbox is set.
     * @param chainId ID of the chain.
     * @param inbox   Address of the inbox.
     */
    event InboxSet(uint64 indexed chainId, address indexed inbox);

    /**
     * @notice Emitted when a cross-chain request is filled on the destination chain
     * @param orderId  ID of the order on the source chain
     * @param fillHash Hash of the fill origin data
     * @param filledBy Address of the solver that filled the oder
     */
    event Filled(bytes32 indexed orderId, bytes32 indexed fillHash, address indexed filledBy);

    /**
     * @notice Set the inbox addresses for the given chain IDs.
     * @param chainIds IDs of the chains.
     * @param inboxes  Addresses of the inboxes.
     */
    function setInboxes(uint64[] calldata chainIds, address[] calldata inboxes) external;

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
}
