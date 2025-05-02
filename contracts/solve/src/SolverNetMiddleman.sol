// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.24;

import { ReentrancyGuard } from "solady/src/utils/ReentrancyGuard.sol";
import { SafeTransferLib } from "solady/src/utils/SafeTransferLib.sol";
import { IERC721 } from "@openzeppelin/contracts/interfaces/IERC721.sol";

/// @dev This contract is being deprecated in favor of moving this logic into the SolverNetExecutor contract.
contract SolverNetMiddleman is ReentrancyGuard {
    using SafeTransferLib for address;

    error CallFailed();
    error InvalidToken();

    /**
     * @notice Execute a call and transfer any received ERC20 tokens back to the recipient
     * @dev Intended to be used when interacting with contracts that don't allow us to specify a recipient
     * @param token  Token to transfer
     * @param to     Recipient address
     * @param target Call target address
     * @param data   Calldata for the call
     */
    function executeAndTransfer(address token, address to, address target, bytes calldata data)
        external
        payable
        nonReentrant
    {
        (bool success,) = target.call{ value: msg.value }(data);
        if (!success) revert CallFailed();

        if (token == address(0)) SafeTransferLib.safeTransferAllETH(to);
        else token.safeTransferAll(to);
    }

    /**
     * @notice Execute a call and transfer a received ERC721 token back to the recipient
     * @dev Intended to be used when interacting with contracts that don't allow us to specify a recipient
     * @param token     Token to transfer
     * @param tokenId   Token ID to transfer
     * @param to        Recipient address
     * @param target    Call target address
     * @param data      Calldata for the call
     */
    function executeAndTransfer721(address token, uint256 tokenId, address to, address target, bytes calldata data)
        external
        payable
        nonReentrant
    {
        (bool success,) = target.call{ value: msg.value }(data);
        if (!success) revert CallFailed();

        if (token == address(0)) revert InvalidToken();
        IERC721(token).transferFrom(address(this), to, tokenId);
    }

    /**
     * @dev Allows contract to receive ETH as a result of call execution
     */
    receive() external payable { }

    /**
     * @dev Prevents callbacks into this contract from triggering a revert
     */
    fallback() external payable { }
}
