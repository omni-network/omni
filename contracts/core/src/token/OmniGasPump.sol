// SPDX-License-Identifier: GPL-3.0-only
pragma solidity 0.8.24;

import { OwnableUpgradeable } from "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import { PausableUpgradeable } from "@openzeppelin/contracts-upgradeable/utils/PausableUpgradeable.sol";
import { XAppUpgradeable } from "src/pkg/XAppUpgradeable.sol";
import { FeeOracleV1 } from "src/xchain/FeeOracleV1.sol";
import { ConfLevel } from "src/libraries/ConfLevel.sol";
import { OmniXFund } from "./OmniXFund.sol";

/**
 * @title OmniGasPump
 * @notice A unidirectional cross-chain gas exchange, allowing users to swap native ETH for native OMNI.[]
 */
contract OmniGasPump is XAppUpgradeable, OwnableUpgradeable, PausableUpgradeable {
    event MaxSwapSet(uint256 max);
    event FundSet(address fund);
    event PctCutSet(uint256 pct);
    event FillUp(address indexed recipient, uint256 total, uint256 paid, uint256 fee, uint256 cut, uint256 amtOMNI);

    /// @notice Gas limit passed to OmniXFund.settleUp xcall
    uint64 internal constant SETTLE_GAS = 100_000;

    /// @notice Denominator for pct cut
    uint256 internal constant PCT_DENOM = 1000;

    /// @notice Address of OmniXFund on Omni
    address public fund;

    /// @notice Max amt (in native token) that can be swapped in a single tx
    uint256 public maxSwap;

    /// @notice Percentage cut (over PCT_DENOM) taken by this contract for each swap
    ///     Used to disencentivize spamming
    uint256 public pctCut;

    /// @notice Map recipient to total owed (sum of historical fillUps()), denominated in OMNI.
    mapping(address => uint256) public owed;

    constructor() {
        _disableInitializers();
    }

    struct InitParams {
        address fund;
        address portal;
        address owner;
        uint256 maxSwap;
        uint256 pctCut;
    }

    function initialize(InitParams calldata p) external initializer {
        _setFund(p.fund);
        _setMaxSwap(p.maxSwap);
        _setPctCut(p.pctCut);
        __XApp_init(p.portal, ConfLevel.Latest);
        __Ownable_init(p.owner);
    }

    /**
     * @notice Swaps msg.value ETH for OMNI and sends it to `recipient` on Omni.
     *
     *      Takes an xcall fee and a pct cut. Cut taken to disencentivize spamming.
     *      Returns the amount of OMNI swapped for.
     *
     *      To retry (if OmniXFund transfer fails), call swap() again with the
     *      same `recipient`, and msg.value == swapFee().
     *
     * @param recipient Address on Omni to send OMNI to
     */
    function fillUp(address recipient) public payable whenNotPaused returns (uint256) {
        uint256 f = fee();
        require(msg.value >= f, "OmniExchange: insufficient fee");

        uint256 amtETH = msg.value - f;
        require(amtETH <= maxSwap, "OmniExchange: over max");

        uint256 cut = amtETH * pctCut / PCT_DENOM;
        amtETH -= cut;

        uint256 amtOMNI = _toOmni(amtETH);
        owed[recipient] += amtOMNI;

        xcall({
            destChainId: omniChainId(),
            to: fund,
            conf: ConfLevel.Latest,
            data: abi.encodeCall(OmniXFund.settleUp, (recipient, owed[recipient])),
            gasLimit: SETTLE_GAS
        });

        emit FillUp(recipient, owed[recipient], msg.value, f, cut, amtOMNI);

        return amtOMNI;
    }

    /**
     * @notice Simulate a fillUp()
     *      Returns the amount of OMNI that `amtETH` msg.value would buy, whether
     *      or not it would succeed, and the revert reason, if any.
     */
    function dryFillUp(uint256 amtETH) public view returns (uint256, bool, string memory) {
        uint256 f = fee();
        if (amtETH < f) return (0, false, "insufficient fee");

        amtETH = amtETH - f;
        if (amtETH > maxSwap) return (0, false, "over max");

        return (_toOmni(_takeCut(amtETH)), true, "");
    }

    /// @notice Returns the xcall fee required for fillUp(). Does not include `pctCut`.
    function fee() public view returns (uint256) {
        // Use max addrs & amount to use no zero byte calldata to ensure max fee
        address recipient = address(type(uint160).max);
        uint256 amt = type(uint256).max;

        return feeFor({
            destChainId: omniChainId(),
            data: abi.encodeCall(OmniXFund.settleUp, (recipient, amt)),
            gasLimit: SETTLE_GAS
        });
    }

    /// @notice Returns the amount of ETH needed to swap for `amtOMNI`
    function quote(uint256 amtOMNI) public view returns (uint256) {
        return _undoCut(_toEth(amtOMNI)) + fee();
    }

    /// @notice Converts `amtOMNI` to ETH, using the current conversion rate
    function _toEth(uint256 amtOMNI) internal view returns (uint256) {
        FeeOracleV1 o = FeeOracleV1(omni.feeOracle());
        return amtOMNI * o.CONVERSION_RATE_DENOM() / o.toNativeRate(omniChainId());
    }

    /// @notice Converts `amtETH` to OMNI, using the current conversion rate
    function _toOmni(uint256 amtETH) internal view returns (uint256) {
        FeeOracleV1 o = FeeOracleV1(omni.feeOracle());
        return amtETH * o.toNativeRate(omniChainId()) / o.CONVERSION_RATE_DENOM();
    }

    /// @notice Returns `amtETH` cut by `pctCut`.
    function _takeCut(uint256 amtETH) internal view returns (uint256) {
        return amtETH - (amtETH * pctCut / PCT_DENOM);
    }

    /// @notice Returns n such `n = _takeCut(amtETH)`
    function _undoCut(uint256 amtETH) internal view returns (uint256) {
        return amtETH + (amtETH * PCT_DENOM / (PCT_DENOM - pctCut));
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
        (bool success,) = to.call{ value: address(this).balance }("");
        require(success, "OmniExchange: withdraw failed");
    }

    /// @notice Set the max swap, denominated in ETh
    function setMaxSwap(uint256 max) external onlyOwner {
        _setMaxSwap(max);
    }

    /// @notice Set the address of the OmniXFund, on Omni
    function setOmniXFund(address omnifund_) external onlyOwner {
        _setFund(omnifund_);
    }

    /// @notice Set the pct cut taken on each swap
    function setPctCut(uint256 pct) external onlyOwner {
        _setPctCut(pct);
    }

    function _setPctCut(uint256 pct) internal {
        require(pct < PCT_DENOM, "OmniExchange: over pct cut");
        pctCut = pct;
        emit PctCutSet(pct);
    }

    function _setMaxSwap(uint256 max) internal {
        require(max > 0, "OmniExchange: zero max");
        maxSwap = max;
        emit MaxSwapSet(max);
    }

    function _setFund(address fund_) internal {
        require(fund_ != address(0), "OmniExchange: zero address");
        fund = fund_;
        emit FundSet(fund_);
    }
}
