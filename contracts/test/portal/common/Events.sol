// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.24;

/**
 * @title Events
 * @dev Contains all events that are tested against. Required until solc allows
 *      for referencing events defined in other contracts.
 */
contract Events {
    event XMsg(
        uint64 indexed destChainId, uint64 indexed streamOffset, address sender, address to, bytes data, uint64 gasLimit
    );

    event XReceipt(
        uint64 indexed sourceChainId, uint64 indexed streamOffset, uint256 gasUsed, address relayer, bool success
    );

    event FeeOracleChanged(address indexed oldFeeOracle, address indexed newFeeOracle);

    event FeeChanged(uint256 oldFee, uint256 newFee);
}
