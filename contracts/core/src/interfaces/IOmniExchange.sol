// SPDX-License-Identifier: GPL-3.0-only
pragma solidity 0.8.24;

interface IOmniExchange {
    function retry(address recipient, uint64 gasLimit) external payable;
    function fund(address recipient, uint256 amtETH, uint64 gasLimit) external payable;
    function fundFee(address recipient, uint256 amtETH, uint64 gasLimit) external view returns (uint256);
    function maxSwap() external view returns (uint256);
    function nativeToOMNI(uint256 amtETH) external view returns (uint256);
}
