// SPDX-License-Identifier: GPL-3.0-only
pragma solidity ^0.8.12;

interface IOmniGasPump {
    /**
     * @notice Emitted on a fillUp
     * @param recipient Address on Omni to send OMNI to
     * @param owed      Total amount owed to `recipient` (sum of historical fillUps()), denominated in OMNI.
     * @param amtETH    Amount of ETH paid
     * @param fee       Xcall fee, denominated in ETH
     * @param toll      Toll taken by this contract, denominated in ETH
     * @param amtOMNI   Amount of OMNI swapped for
     */
    event FilledUp(address indexed recipient, uint256 owed, uint256 amtETH, uint256 fee, uint256 toll, uint256 amtOMNI);

    /// @notice Address of OmniGasStation on Omni
    function gasStation() external view returns (address);

    /// @notice Max amt (in native token) that can be swapped in a single tx
    function maxSwap() external view returns (uint256);

    /// @notice Percentage toll taken by this contract for each swap, to disincentivize spamming
    function toll() external view returns (uint256);

    /// @notice Map recipient to total owed (sum of historical fillUps()), denominated in OMNI.
    function owed(address) external view returns (uint256);

    /**
     * @notice Swaps msg.value ETH for OMNI and sends it to `recipient` on Omni.
     *
     *      Takes an xcall fee and a pct cut. Cut taken to disincentivize spamming.
     *      Returns the amount of OMNI swapped for.
     *
     *      To retry (if OmniGasStation transfer fails), call swap() again with the
     *      same `recipient`, and msg.value == swapFee().
     *
     * @param recipient Address on Omni to send OMNI to
     */
    function fillUp(address recipient) external payable returns (uint256);

    /**
     * @notice Simulate a fillUp()
     *      Returns the amount of OMNI that `amtETH` msg.value would buy, whether
     *      or not it would succeed, and the revert reason, if any.
     */
    function dryFillUp(uint256 amtETH) external view returns (uint256, bool, string memory);

    /// @notice Returns the xcall fee required for fillUp(). Does not include `pctCut`.
    function xfee() external view returns (uint256);

    /// @notice Returns the amount of ETH needed to swap for `amtOMNI`
    function quote(uint256 amtOMNI) external view returns (uint256);
}
