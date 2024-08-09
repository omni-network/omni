// SPDX-License-Identifier: GPL-3.0-only
pragma solidity 0.8.24;

import { XApp } from "src/pkg/XApp.sol";
import { FeeOracleV1 } from "src/xchain/FeeOracleV1.sol";
import { Ownable } from "@openzeppelin/contracts/access/Ownable.sol";
import { ConfLevel } from "src/libraries/ConfLevel.sol";
import { OmniFund } from "./OmniFund.sol";

contract OmniExchange is XApp, Ownable {
    /// @notice Max amount (in OMNI) that can be swapped in a single tx
    uint256 internal _maxOmniPerSwap = 1e18;

    /// @notice Map address to total swap requests
    mapping(address => uint256) public swapped;

    /// @notice Address of OmniFund on Omni
    address public fund;

    constructor(address portal, address fund_, address owner) XApp(portal, ConfLevel.Latest) Ownable(owner) {
        fund = fund_;
    }

    /**
     * @notice Swap native tokens for OMNI tokens, on Omni. Or, retry if last swap failed.
     * @param recipient Address to receive the OMNI tokens
     */
    function swapOrRetry(address recipient) external payable {
        uint256 amt = nativeToOMNI(msg.value);
        require(amt <= _maxOmniPerSwap, "OmniExchange: swap t");

        swapped[recipient] += amt;

        xcall({
            destChainId: omni.omniChainId(),
            to: fund,
            data: abi.encodeCall(OmniFund.tryWithdrawRemaining, (recipient, swapped[msg.sender])),
            gasLimit: 100_000
        });
    }

    /**
     * @notice Returns the max amount of native tokens that can be exchanged
     *         for OMNI in a single swap.
     */
    function maxSwap() external view returns (uint256) {
        return _omniToNative(_maxOmniPerSwap);
    }

    function nativeToOMNI(uint256 amount) public view returns (uint256) {
        FeeOracleV1 oracle = FeeOracleV1(omni.feeOracle());
        FeeOracleV1.ChainFeeParams memory params = oracle.feeParams(omni.omniChainId());
        return amount * params.toNativeRate / oracle.CONVERSION_RATE_DENOM();
    }

    function _omniToNative(uint256 amount) public view returns (uint256) {
        FeeOracleV1 oracle = FeeOracleV1(omni.feeOracle());
        FeeOracleV1.ChainFeeParams memory params = oracle.feeParams(omni.omniChainId());
        return amount * oracle.CONVERSION_RATE_DENOM() / params.toNativeRate;
    }

    function withdraw(address to) external onlyOwner {
        (bool success,) = to.call{ value: address(this).balance }("");
        require(success, "OmniExchange: withdraw failed");
    }

    function setMaxSwap(uint256 max) external onlyOwner {
        _maxOmniPerSwap = max;
    }
}
