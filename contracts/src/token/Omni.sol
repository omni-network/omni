// SPDX-License-Identifier: GPL-3.0-only
pragma solidity 0.8.24;

import { ERC20 } from "@openzeppelin/contracts/token/ERC20/ERC20.sol";
import { ERC20Permit } from "@openzeppelin/contracts/token/ERC20/extensions/ERC20Permit.sol";

contract Omni is ERC20, ERC20Permit {
    /**
     * @notice Construct an OMNI ERC20 token.
     * @param initialSupply   The initial token supply, minted to `recipient`
     * @param recipient       The recipient of the initial supply
     */
    constructor(uint256 initialSupply, address recipient) ERC20("Omni Network", "OMNI") ERC20Permit("Omni Network") {
        _mint(recipient, initialSupply);
    }
}
