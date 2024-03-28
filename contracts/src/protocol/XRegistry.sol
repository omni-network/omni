// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.24;

import { OwnableUpgradeable } from "@openzeppelin/contracts-upgradeable/access/OwnableUpgradeable.sol";

import { OmniPortal } from "./OmniPortal.sol";
import { XRegistryBase } from "./XRegistryBase.sol";
import { ContractNames } from "../libraries/ContractNames.sol";
import { IXRegistry } from "../interfaces/IXRegistry.sol";

/**
 * @title XRegistry
 * @notice Local contract registry, deployed along each portal. Controlled by the GlobalXRegistry
 *         predeployed on Omni.
 */
contract XRegistry is IXRegistry, OwnableUpgradeable { }
