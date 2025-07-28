// SPDX-License-Identifier: GPL-3.0-only
pragma solidity ^0.8.30;

import { ERC20 } from "solady/src/tokens/ERC20.sol";
import { SafeTransferLib } from "solady/src/utils/SafeTransferLib.sol";

contract Nomina is ERC20 {
    using SafeTransferLib for address;

    error ZeroAmount();
    error ZeroAddress();
    error ConversionDisabled();

    string private constant _NAME = "Nomina";
    string private constant _SYMBOL = "NOM";

    bytes32 private constant _NAME_HASH = keccak256(bytes(_NAME));
    address private constant _DEAD_ADDRESS = address(0xdead);
    uint8 private constant _CONVERSION_RATE = 75;

    address public immutable omni;

    constructor(address _omni) {
        omni = _omni;
    }

    function name() public pure override returns (string memory) {
        return _NAME;
    }

    function symbol() public pure override returns (string memory) {
        return _SYMBOL;
    }

    function mint(address to, uint256 amount) public { }

    function burn(uint256 amount) public {
        if (amount == 0) revert ZeroAmount();
        _burn(msg.sender, amount);
    }

    function convert(address to, uint256 amount) public {
        address _omni = omni;
        if (to == address(0)) revert ZeroAddress();
        if (amount == 0) revert ZeroAmount();
        if (_omni == address(0)) revert ConversionDisabled();

        _omni.safeTransferFrom(msg.sender, _DEAD_ADDRESS, amount);
        _mint(to, amount * _CONVERSION_RATE);
    }

    function _constantNameHash() internal pure override returns (bytes32) {
        return _NAME_HASH;
    }
}
