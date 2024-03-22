// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.24;

import { Test } from "forge-std/Test.sol";
import { Events } from "./Events.sol";
import { Fixtures } from "./Fixtures.sol";
import { Utils } from "./Utils.sol";

/**
 * @title Base
 * @dev An extension of forge Test that includes Omni specifc setup, fixtures, utils, and events.
 */
contract Base is Test, Events, Fixtures, Utils { }
