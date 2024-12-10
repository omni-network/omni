// SPDX-License-Identifier: GPL-3.0-only
pragma solidity 0.8.24;

import { PausableUpgradeable } from "@openzeppelin/contracts-upgradeable/utils/PausableUpgradeable.sol";
import { OwnableUpgradeable } from "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";
import { XAppUpgradeable } from "src/pkg/XAppUpgradeable.sol";
import { ConfLevel } from "src/libraries/ConfLevel.sol";

/**
 * @title OmniGasStation
 * @notice Pays out all gas owed by way of OmniGasPumps on other chains
 */
contract OmniGasStation is XAppUpgradeable, OwnableUpgradeable, PausableUpgradeable {
    /**
     * @notice Emitted on settleUp
     * @param recipient Address
     * @param chainID   ChainID of the pump
     * @param owed      Total amount owed to `recipient` by way of `chainID`'s OmniGasPump
     * @param fueled    Total amount sent to `recipient` by this contract.
     * @param success   True if the transfer was successful (if true, owed == fueled)
     */
    event SettledUp(address indexed recipient, uint64 indexed chainID, uint256 owed, uint256 fueled, bool success);

    /// @notice Emitted when a OmniGasPump is set for a chain
    event GasPumpAdded(uint64 indexed chainID, bytes32 addr);

    //// @notice Map chainID to addr to true, if authorized to withdraw
    mapping(uint64 => bytes32) public pumps;

    /// @notice Map recipient to chainID to total fueled
    mapping(address => mapping(uint64 => uint256)) public fueled;

    constructor() {
        _disableInitializers();
    }

    /// @dev GasPump struct, just used in initialize params
    struct GasPump {
        uint64 chainID;
        bytes32 addr;
    }

    function initialize(address portal, address owner, GasPump[] calldata pumps_) external initializer {
        __XApp_init(portal, ConfLevel.Finalized);
        __Ownable_init(owner);

        for (uint256 i = 0; i < pumps_.length; i++) {
            _setPump(pumps_[i].chainID, pumps_[i].addr);
        }
    }

    /**
     * @notice Settle up with a recipient. If `owed` is more than they've been fueled, send the difference.
     * @param recipient Address to receive the funds
     * @param owed      Total amount owed to `recipient`, by way of xmsg.sourceChainId's OmniGasPump
     */
    function settleUp(address recipient, uint256 owed) external xrecv whenNotPaused {
        require(isXCall() && isPump(xmsg.sourceChainId, xmsg.sender), "GasStation: unauthorized");

        uint256 settled = fueled[recipient][xmsg.sourceChainId];

        // If already settled, revert
        require(owed > settled, "GasStation: already funded");

        // Transfer the difference
        (bool success,) = recipient.call{ value: owed - settled }("");

        // Update books. We do not bother doing so pre-transfer, because isXCall() prevents reentrancy
        if (success) fueled[recipient][xmsg.sourceChainId] = owed;

        emit SettledUp(recipient, xmsg.sourceChainId, owed, fueled[recipient][xmsg.sourceChainId], success);
    }

    /// @notice Set the pump addr for a chain
    function setPump(uint64 chainId, bytes32 addr) external onlyOwner {
        _setPump(chainId, addr);
    }

    /// @notice Return true if `chainID` has a registered pump at `addr`
    function isPump(uint64 chainID, bytes32 addr) public view returns (bool) {
        return addr != bytes32(0) && addr == pumps[chainID];
    }

    /// @notice Pause withdrawals
    function pause() external onlyOwner {
        _pause();
    }

    /// @notice Unpause withdrawals
    function unpause() external onlyOwner {
        _unpause();
    }

    function _setPump(uint64 chainId, bytes32 addr) internal {
        require(addr != bytes32(0), "GasStation: zero addr");
        require(chainId != 0, "GasStation: zero chainId");

        pumps[chainId] = addr;
        emit GasPumpAdded(chainId, addr);
    }

    receive() external payable { }
}
