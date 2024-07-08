// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.24;

import { XTypes } from "src/libraries/XTypes.sol";

/**
 * @title TestXTypes
 * @dev Defines xtypes not needed on chain, and therefore not included in src/interfaces/XTypes.sol,
 *      but part of Omni's xchain messaging protocol (see lib/xchain/types.go), and useful in tests.
 */
library TestXTypes {
    struct Receipt {
        uint64 sourceChainId;
        uint64 shardId;
        uint64 offset;
        uint256 gasUsed;
        address relayer;
        bool success;
        bytes error;
    }

    /// @dev receipts omitted, as they are not needed to construct XSubmissions
    struct Block {
        XTypes.BlockHeader blockHeader;
        XTypes.Msg[] msgs;
    }
}
