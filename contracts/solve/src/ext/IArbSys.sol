// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.24;

/**
 * @title IArbSys
 * @notice Interface for Arbitrum's ArbSys precompile (0x0000000000000000000000000000000000000064)
 * @custom:source https://github.com/OffchainLabs/nitro-contracts/blob/main/src/precompiles/ArbSys.sol
 */
interface IArbSys {
    /**
     * @notice Get Arbitrum block number (distinct from L1 block number; Arbitrum genesis block has block number 0)
     * @return block number as int
     */
    function arbBlockNumber() external view returns (uint256);
}
