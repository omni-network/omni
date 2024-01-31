// SPDX-License-Identifier: GPL-3.0-only
pragma solidity 0.8.23;

import { IFeeOracle } from "./interfaces/IFeeOracle.sol";
import { IFeeOracleV1 } from "./interfaces/IFeeOracleV1.sol";

/**
 * @title FeeOracleV1
 * @notice A simple fee oracle with a fixed fee, controlled by an admin account
 * @dev Used by OmniPortal to calculate xmsg fees
 */
contract FeeOracleV1 is IFeeOracle, IFeeOracleV1 {
    /// @inheritdoc IFeeOracleV1
    address public admin;

    /// @inheritdoc IFeeOracleV1
    uint256 public fee;

    modifier onlyAdmin() {
        require(msg.sender == admin, "FeeOracleV1: only admin");
        _;
    }

    constructor(address admin_, uint256 fee_) {
        _setAdmin(admin_);
        _setFee(fee_);
    }

    /// @inheritdoc IFeeOracle
    function feeFor(uint64, bytes calldata, uint64) external view returns (uint256) {
        return fee;
    }

    /// @inheritdoc IFeeOracleV1
    function setAdmin(address admin_) external onlyAdmin {
        _setAdmin(admin_);
    }

    /// @inheritdoc IFeeOracleV1
    function setFee(uint256 fee_) external onlyAdmin {
        _setFee(fee_);
    }

    /// @dev Set the fee
    function _setFee(uint256 fee_) internal {
        require(fee_ > 0, "FeeOracleV1: fee must be > 0");

        uint256 oldFee = fee;
        fee = fee_;

        emit FeeChanged(oldFee, fee_);
    }

    /// @dev Set the admin account
    function _setAdmin(address admin_) internal {
        require(admin_ != address(0), "FeeOracleV1: no zero admin");

        address oldAdmin = admin;
        admin = admin_;

        emit AdminChanged(oldAdmin, admin_);
    }
}
