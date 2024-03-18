// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.24;

import { OwnableUpgradeable } from "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import { IFeeOracle } from "../interfaces/IFeeOracle.sol";
import { IFeeOracleV1 } from "../interfaces/IFeeOracleV1.sol";

/**
 * @title FeeOracleV1
 * @notice A simple fee oracle with a fixed fee, controlled by an admin account
 *         Used by OmniPortal to calculate xmsg fees
 */
contract FeeOracleV1 is IFeeOracle, IFeeOracleV1, OwnableUpgradeable {
    /// @notice The current fee per transaction, in wei.
    uint256 public fee;

    constructor() {
        _disableInitializers();
    }

    function initialize(address owner_, uint256 fee_) public initializer {
        __Ownable_init();
        _transferOwnership(owner_);
        _setFee(fee_);
    }

    /// @inheritdoc IFeeOracle
    function feeFor(uint64, bytes calldata, uint64) external view returns (uint256) {
        return fee;
    }

    /// @inheritdoc IFeeOracleV1
    function setFee(uint256 fee_) external onlyOwner {
        _setFee(fee_);
    }

    /// @dev Set the fee
    function _setFee(uint256 fee_) internal {
        require(fee_ > 0, "FeeOracleV1: fee must be > 0");

        uint256 oldFee = fee;
        fee = fee_;

        emit FeeChanged(oldFee, fee_);
    }
}
