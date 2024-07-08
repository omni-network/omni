// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.24;

/**
 * @title Events
 * @dev Contains all events that are tested against. Required until solc allows
 *      for referencing events defined in other contracts.
 */
contract Events {
    event XMsg(
        uint64 indexed destChainId,
        uint64 indexed shardId,
        uint64 indexed offset,
        address sender,
        address to,
        bytes data,
        uint64 gasLimit,
        uint256 fees
    );

    event XReceipt(
        uint64 indexed sourceChainId,
        uint64 indexed shardId,
        uint64 indexed offset,
        uint256 gasUsed,
        address relayer,
        bool success,
        bytes error
    );

    event FeeOracleChanged(address indexed oldFeeOracle, address indexed newFeeOracle);

    event FeeChanged(uint256 oldFee, uint256 newFee);
}
