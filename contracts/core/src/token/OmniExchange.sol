// SPDX-License-Identifier: GPL-3.0-only
pragma solidity 0.8.24;

import { IOmniPortal } from "src/interfaces/IOmniPortal.sol";
import { OwnableUpgradeable } from "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import { FeeOracleV1 } from "src/xchain/FeeOracleV1.sol";
import { ConfLevel } from "src/libraries/ConfLevel.sol";
import { OmniFund } from "./OmniFund.sol";
import { IOmniExchange } from "src/interfaces/IOmniExchange.sol";

contract OmniExchange is IOmniExchange, OwnableUpgradeable {
    /// @notice Max amtETH (in OMNI) that can be swapped in a single tx
    uint256 internal _maxOmniPerSwap = 1e18;

    /// @notice Map address to total swap requests
    mapping(address => uint256) public swapped;

    /// @notice Address of OmniFund on Omni
    address public omnifund;

    IOmniPortal public omni;

    constructor() {
        _disableInitializers();
    }

    function initialize(address portal, address fund_, address owner) external initializer {
        omnifund = fund_;
        omni = IOmniPortal(portal);
        __Ownable_init(owner);
    }

    function retry(address recipient, uint64 gasLimit) external payable {
        _fund(recipient, 0, gasLimit);
    }

    function fund(address recipient, uint256 amtETH, uint64 gasLimit) external payable {
        _fund(recipient, amtETH, gasLimit);
    }

    function _fund(address recipient, uint256 amtETH, uint64 gasLimit) internal {
        uint256 fee = fundFee(recipient, amtETH, gasLimit);
        require(msg.value == amtETH + fee, "OmniExchange: incorrect amount");

        uint256 amtOMNI = nativeToOMNI(amtETH);
        require(amtOMNI <= _maxOmniPerSwap, "OmniExchange: amtETH too large");

        swapped[recipient] += amtOMNI;

        omni.xcall{ value: fee }({
            destChainId: omni.omniChainId(),
            to: omnifund,
            conf: ConfLevel.Latest,
            data: abi.encodeCall(OmniFund.tryWithdrawRemaining, (recipient, swapped[recipient])),
            gasLimit: gasLimit
        });
    }

    function fundFee(address recipient, uint256 amtETH, uint64 gasLimit) public view returns (uint256) {
        return omni.feeFor({
            destChainId: omni.omniChainId(),
            data: abi.encodeCall(OmniFund.tryWithdrawRemaining, (recipient, nativeToOMNI(amtETH))),
            gasLimit: gasLimit
        });
    }

    function maxSwap() external view returns (uint256) {
        return _omniToNative(_maxOmniPerSwap);
    }

    function nativeToOMNI(uint256 amtETH) public view returns (uint256) {
        FeeOracleV1 oracle = FeeOracleV1(omni.feeOracle());
        // TODO: check
        return amtETH * oracle.toNativeRate(omni.omniChainId()) / oracle.CONVERSION_RATE_DENOM();
    }

    function _omniToNative(uint256 amtETH) public view returns (uint256) {
        FeeOracleV1 oracle = FeeOracleV1(omni.feeOracle());
        // TODO: check
        return amtETH * oracle.CONVERSION_RATE_DENOM() / oracle.toNativeRate(omni.omniChainId());
    }

    //////////////////////////////////////////////////////////////////////////////
    //                                  Admin                                   //
    //////////////////////////////////////////////////////////////////////////////

    function withdraw(address to) external onlyOwner {
        (bool success,) = to.call{ value: address(this).balance }("");
        require(success, "OmniExchange: withdraw failed");
    }

    function setMaxSwap(uint256 max) external onlyOwner {
        _maxOmniPerSwap = max;
    }
}
