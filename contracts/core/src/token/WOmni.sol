// SPDX-License-Identifier: GPL-3.0
// Copyright (C) 2015, 2016, 2017 Dapphub

// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.

// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.

// You should have received a copy of the GNU General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

// Based on WETH9 by Dapphub, and WETH98 by OP Labs.
// Modified  by Omni Network

pragma solidity 0.8.24;

import { IWOmni } from "src/interfaces/IWOmni.sol";

/**
 * @title WOmni
 * @notice WOmni is a modifed version of WETH98, is a version of WETH9 upgraded for Solidity 0.8.x.
 * @dev Changes: Renamed "WETH" to "WOmni", compiler version bumped to 0.8.24
 * @custom:attribution https://github.com/ethereum-optimism/optimism/blob/develop/packages/contracts-bedrock/src/dispute/weth/WETH98.sol
 */
contract WOmni is IWOmni {
    uint8 public constant decimals = 18;

    mapping(address => uint256) public balanceOf;
    mapping(address => mapping(address => uint256)) public allowance;

    /// @notice Pipes to deposit.
    receive() external payable {
        deposit();
    }

    /// @notice Pipes to deposit.
    fallback() external payable {
        deposit();
    }

    /// @inheritdoc IWOmni
    function name() external view virtual override returns (string memory) {
        return "Wrapped Omni";
    }

    /// @inheritdoc IWOmni
    function symbol() external view virtual override returns (string memory) {
        return "WOMNI";
    }

    /// @inheritdoc IWOmni
    function deposit() public payable {
        balanceOf[msg.sender] += msg.value;
        emit Deposit(msg.sender, msg.value);
    }

    /// @inheritdoc IWOmni
    function withdraw(uint256 wad) public virtual {
        require(balanceOf[msg.sender] >= wad);
        balanceOf[msg.sender] -= wad;
        payable(msg.sender).transfer(wad);
        emit Withdrawal(msg.sender, wad);
    }

    /// @inheritdoc IWOmni
    function totalSupply() external view returns (uint256) {
        return address(this).balance;
    }

    /// @inheritdoc IWOmni
    function approve(address guy, uint256 wad) external returns (bool) {
        allowance[msg.sender][guy] = wad;
        emit Approval(msg.sender, guy, wad);
        return true;
    }

    /// @inheritdoc IWOmni
    function transfer(address dst, uint256 wad) external returns (bool) {
        return transferFrom(msg.sender, dst, wad);
    }

    /// @inheritdoc IWOmni
    function transferFrom(address src, address dst, uint256 wad) public returns (bool) {
        require(balanceOf[src] >= wad);

        if (src != msg.sender && allowance[src][msg.sender] != type(uint256).max) {
            require(allowance[src][msg.sender] >= wad);
            allowance[src][msg.sender] -= wad;
        }

        balanceOf[src] -= wad;
        balanceOf[dst] += wad;

        emit Transfer(src, dst, wad);

        return true;
    }
}
