// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.24;

import { OwnableUpgradeable } from "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";

/**
 * @title Redenom
 * @notice The EVM interface to the consensus chain's halo/evmredenom module.
 *         Calls are proxied, and not executed synchronously. Their execution is left to
 *         the consensus chain, and they may fail.
 */
contract Redenom is OwnableUpgradeable {
    /**
     * @notice Emitted when an AccountRange is submitted.
     */
    event Submitted(address[] addresses, bytes[] accounts, bytes[] proof);

    /**
     * @notice AccountRange specifies a range of accounts; see go-ethereum/eth/protocols/snap#AccountRangePacket
     * @custom:field addresses        List of consecutive account hash pre-images (address) from the trie
     * @custom:field accountBodies    RLP encoded bodies of above accounts
     * @custom:field proof            List of trie nodes proving the account range
     */
    struct AccountRange {
        address[] addresses;
        bytes[] accounts;
        bytes[] proof;
    }

    function initialize(address owner_) public initializer {
        __Ownable_init(owner_);
    }

    //////////////////////////////////////////////////////////////////////////////
    //                                  Admin                                   //
    //////////////////////////////////////////////////////////////////////////////

    /**
     * @notice Submit a new account range to redenominate.
     */
    function submit(AccountRange calldata range) external onlyOwner {
        emit Submitted(range.addresses, range.accounts, range.proof);
    }
}
