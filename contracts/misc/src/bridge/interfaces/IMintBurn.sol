// SPDX-License-Identifier: MIT
pragma solidity 0.8.26;

interface IMintBurn {
    function mint(address to, uint256 value) external;
    function burn(uint256 value) external;
}
