// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.24;

import { OwnableUpgradeable } from "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";

/**
 * @title Upgrade
 * @notice The EVM interface to the consensus chain's x/upgrade module.
 *         Calls are proxied, and not executed synchronously. Their execution is left to
 *         the consensus chain, and they may fail.
 * @dev This contract is predeployed, and requires storage slots to be set in genesis.
 *      initialize(...) is called pre-deployment, in script/genesis/AllocPredeploys.s.sol
 *      Initializers on the implementation are disabled via manual storage updates, rather than in a constructor.
 *      If a new implementation is required, a constructor should be added.
 */
contract Upgrade is OwnableUpgradeable {
    /**
     * @notice Emitted when a software upgrade is planned
     * @param name     (MsgSoftwareUpgrade.plan.name) The name for the upgrade
     * @param height   (MsgSoftwareUpgrade.plan.height) The height at which the upgrade must be performed
     * @param info     (MsgSoftwareUpgrade.plan.info) Any application specific upgrade info to be included on-chain such as a git commit that validators could automatically upgrade to
     */
    event PlanUpgrade(string name, uint64 height, string info);

    /**
     * @notice Emitted when the current planned (non-executed) upgrade plan should be cancelled.
     */
    event CancelUpgrade();

    /**
     * @notice Plan specifies information about a planned upgrade and when it should occur..
     * @custom:field name      (MsgSoftwareUpgrade.plan.name) The name for the upgrade
     * @custom:field height    (MsgSoftwareUpgrade.plan.height) The height at which the upgrade must be performed
     * @custom:field info      (MsgSoftwareUpgrade.plan.info) Any application specific upgrade info to be included on-chain such as a git commit that validators could automatically upgrade to
     */
    struct Plan {
        string name;
        uint64 height;
        string info;
    }

    function initialize(address owner_) public initializer {
        __Ownable_init(owner_);
    }

    //////////////////////////////////////////////////////////////////////////////
    //                                  Admin                                   //
    //////////////////////////////////////////////////////////////////////////////

    /**
     * @notice Plan a new software upgrade
     */
    function planUpgrade(Plan calldata plan) external onlyOwner {
        emit PlanUpgrade(plan.name, plan.height, plan.info);
    }

    /**
     * @notice Cancels the current planned (non-executed) upgrade.
     */
    function cancelUpgrade() external onlyOwner {
        emit CancelUpgrade();
    }
}
