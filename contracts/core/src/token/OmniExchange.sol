// SPDX-License-Identifier: GPL-3.0-only
pragma solidity 0.8.24;

import { XApp } from "src/pkg/XApp.sol";
import { FeeOracleV1 } from "src/xchain/FeeOracleV1.sol";
import { Ownable } from "@openzeppelin/contracts/access/Ownable.sol";
import { ConfLevel } from "src/libraries/ConfLevel.sol";
import { OmniFund } from "./OmniFund.sol";

contract OmniExchange is XApp, Ownable {
    /// @notice Map address to total swap requests
    mapping(address => uint256) public swapped;

    /// @notice Address of OmniFund on Omni
    address public fund;

    constructor(address portal, address fund_, address owner) XApp(portal, ConfLevel.Latest) Ownable(owner) {
        fund = fund_;
    }

    /**
     * @notice Swap native tokens for OMNI tokens, on Omni. Or, retry if last swap failed.
     */
    function swapOrRetry() external payable {
        swapped[msg.sender] += nativeToOMNI(msg.value);
        xcall({
            destChainId: omni.omniChainId(),
            to: fund,
            data: abi.encodeCall(OmniFund.tryWithdrawRemaining, (msg.sender, swapped[msg.sender])),
            gasLimit: 100_000
        });
    }

    function nativeToOMNI(uint256 amount) public view returns (uint256) {
        FeeOracleV1 oracle = FeeOracleV1(omni.feeOracle());
        FeeOracleV1.ChainFeeParams memory params = oracle.feeParams(omni.omniChainId());
        return amount * params.toNativeRate / oracle.CONVERSION_RATE_DENOM();
    }

    function withdraw(address to) external onlyOwner {
        (bool success,) = to.call{ value: address(this).balance }("");
        require(success, "OmniExchange: withdraw failed");
    }
}
