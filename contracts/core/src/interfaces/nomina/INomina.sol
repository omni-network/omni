// SPDX-License-Identifier: GPL-3.0-only
pragma solidity ^0.8.24;

interface INomina {
    function convert(address to, uint256 amount) external;
}
