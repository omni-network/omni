// SPDX-License-Identifier: GPL-3.0-only
pragma solidity 0.8.24;

// solhint-disable no-console

import { OwnableUpgradeable } from "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import { PausableUpgradeable } from "@openzeppelin/contracts-upgradeable/utils/PausableUpgradeable.sol";
import { XAppUpgradeable } from "src/pkg/XAppUpgradeable.sol";
import { IConversionRateOracle } from "src/interfaces/IConversionRateOracle.sol";
import { IOmniGasPump } from "src/interfaces/IOmniGasPump.sol";
import { ConfLevel } from "src/libraries/ConfLevel.sol";
import { OmniGasStation } from "./OmniGasStation.sol";

/**
 * @title OmniGasPump
 * @notice A unidirectional cross-chain gas exchange, allowing users to swap native ETH for native OMNI.
 */
contract OmniGasPump is IOmniGasPump, XAppUpgradeable, OwnableUpgradeable, PausableUpgradeable {
    /// @notice Emitted when the max swap is set
    event MaxSwapSet(uint256 max);

    /// @notice Emitted when the gas station is set
    event GasStationSet(address station);

    /// @notice Emitted when the fee oracle is set
    event OracleSet(address oracle);

    /// @notice Emitted when the toll is set
    event TollSet(uint256 pct);

    /// @notice Gas limit passed to OmniGasStation.settleUp xcall
    uint64 public constant SETTLE_GAS = 140_000;

    /// @notice Denominator for toll percentage calculations
    uint256 public constant TOLL_DENOM = 1000;

    /// @notice Native token conversion rate oracle
    IConversionRateOracle public oracle;

    /// @notice Address of OmniGasStation on Omni
    address public gasStation;

    /// @notice Max amt (in native token) that can be swapped in a single tx
    uint256 public maxSwap;

    /// @notice Percentage toll taken by this contract for each swap, to disincentivize spamming
    uint256 public toll;

    /// @notice Map recipient to total owed (sum of historical fillUps()), denominated in OMNI.
    mapping(address => uint256) public owed;

    constructor() {
        _disableInitializers();
    }

    struct InitParams {
        address gasStation;
        address oracle;
        address portal;
        address owner;
        uint256 maxSwap;
        uint256 toll;
    }

    function initialize(InitParams calldata p) external initializer {
        _setOracle(p.oracle);
        _setGasStation(p.gasStation);
        _setMaxSwap(p.maxSwap);
        _setToll(p.toll);

        __XApp_init(p.portal, ConfLevel.Latest);
        __Ownable_init(p.owner);
    }

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
    function fillUp(address recipient) public payable whenNotPaused returns (uint256) {
        require(recipient != address(0), "OmniGasPump: no zero addr");

        // take xcall fee
        uint256 f = xfee();
        require(msg.value >= f, "OmniGasPump: insufficient fee");
        uint256 amtETH = msg.value - f;

        // check max
        require(amtETH <= maxSwap, "OmniGasPump: over max");

        // take toll
        uint256 t = amtETH * toll / TOLL_DENOM;
        amtETH -= t;

        uint256 amtOMNI = _toOmni(amtETH);

        // update owed
        owed[recipient] += amtOMNI;

        // settle up with the gas station
        xcall({
            destChainId: omniChainId(),
            to: gasStation,
            conf: ConfLevel.Latest,
            data: abi.encodeCall(OmniGasStation.settleUp, (recipient, owed[recipient])),
            gasLimit: SETTLE_GAS
        });

        emit FilledUp(recipient, owed[recipient], msg.value, f, t, amtOMNI);

        return amtOMNI;
    }

    /**
     * @notice Simulate a fillUp()
     *      Returns the amount of OMNI that `amtETH` msg.value would buy, whether
     *      or not it would succeed, and the revert reason, if any.
     */
    function dryFillUp(uint256 amtETH) public view returns (uint256, bool, string memory) {
        // take xcall fee
        uint256 f = xfee();
        if (amtETH < f) return (0, false, "insufficient fee");
        amtETH -= f;

        // check max
        if (amtETH > maxSwap) return (0, false, "over max");

        // take toll
        amtETH -= amtETH * toll / TOLL_DENOM;

        return (_toOmni(amtETH), true, "");
    }

    /// @notice Returns the xcall fee required for fillUp(). Does not include `pctCut`.
    function xfee() public view returns (uint256) {
        // Use max addrs & amount to use no zero byte calldata to ensure max fee
        address recipient = address(type(uint160).max);
        uint256 amt = type(uint256).max;

        return feeFor({
            destChainId: omniChainId(),
            data: abi.encodeCall(OmniGasStation.settleUp, (recipient, amt)),
            gasLimit: SETTLE_GAS
        });
    }

    /// @notice Returns the amount of ETH needed to swap for `amtOMNI`
    function quote(uint256 amtOMNI) public view returns (uint256) {
        uint256 amtETH = _toEth(amtOMNI);

        // "undo" toll
        amtETH = (amtETH * TOLL_DENOM / (TOLL_DENOM - toll));

        // "undo" xcall fee
        return amtETH + xfee();
    }

    /// @notice Converts `amtOMNI` to ETH, using the current conversion rate
    function _toEth(uint256 amtOMNI) internal view returns (uint256) {
        // toNativeRate(omniChainId()) ==  ETH per OMNI
        return amtOMNI * oracle.toNativeRate(omniChainId()) / oracle.CONVERSION_RATE_DENOM();
    }

    /// @notice Converts `amtETH` to OMNI, using the current conversion rate
    function _toOmni(uint256 amtETH) internal view returns (uint256) {
        // toNativeRate(omniChainId()) ==  ETH per OMNI
        // to convert ETH to OMNI, we use 1 / toNativeRate(omniChainId())
        return amtETH * oracle.CONVERSION_RATE_DENOM() / oracle.toNativeRate(omniChainId());
    }

    //////////////////////////////////////////////////////////////////////////////
    //                                  Admin                                   //
    //////////////////////////////////////////////////////////////////////////////

    /// @notice Pause fill ups
    function pause() external onlyOwner {
        _pause();
    }

    /// @notice Unpause fill ups
    function unpause() external onlyOwner {
        _unpause();
    }

    /// @notice Withdraw collected ETH to `to`
    function withdraw(address to) external onlyOwner {
        require(to != address(0), "OmniGasPump: no zero addr");
        (bool success,) = to.call{ value: address(this).balance }("");
        require(success, "OmniGasPump: withdraw failed");
    }

    /// @notice Set the max swap, denominated in ETh
    function setMaxSwap(uint256 max) external onlyOwner {
        _setMaxSwap(max);
    }

    /// @notice Set the address of the OmniGasStation, on Omni
    function setGasStation(address station) external onlyOwner {
        _setGasStation(station);
    }

    /// @notice Set the conversion rate oracle
    function setOracle(address oracle_) external onlyOwner {
        _setOracle(oracle_);
    }

    /// @notice Set the toll (as a percentage over PCT_DENOM)
    function setToll(uint256 pct) external onlyOwner {
        _setToll(pct);
    }

    function _setToll(uint256 pct) internal {
        require(pct < TOLL_DENOM, "OmniGasPump: pct too high");
        toll = pct;
        emit TollSet(toll);
    }

    function _setMaxSwap(uint256 max) internal {
        require(max > 0, "OmniGasPump: zero max");
        maxSwap = max;
        emit MaxSwapSet(max);
    }

    function _setGasStation(address station) internal {
        require(station != address(0), "OmniGasPump: zero address");
        gasStation = station;
        emit GasStationSet(station);
    }

    function _setOracle(address oracle_) internal {
        require(oracle_ != address(0), "OmniGasPump: zero oracle");
        oracle = IConversionRateOracle(oracle_);

        emit OracleSet(oracle_);
    }
}
