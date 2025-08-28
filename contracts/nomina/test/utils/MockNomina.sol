// SPDX-License-Identifier: GPL-3.0-only
pragma solidity ^0.8.24;

import { ERC20 } from "solady/src/tokens/ERC20.sol";
import { SafeTransferLib } from "solady/src/utils/SafeTransferLib.sol";

/**
 * @title MockNomina
 * @notice Nomina with public mints.
 * @dev Because mint is public, mint authority mechanisms are not includued.
 */
contract MockNomina is ERC20 {
    using SafeTransferLib for address;

    error ZeroAddress();

    address private constant _DEAD_ADDRESS = address(0xdead);

    uint8 public constant CONVERSION_RATE = 75;

    address public immutable OMNI;

    constructor(address _omni) {
        if (_omni == address(0)) revert ZeroAddress();
        OMNI = _omni;
    }

    function name() public pure override returns (string memory) {
        return "Nomina";
    }

    function symbol() public pure override returns (string memory) {
        return "NOM";
    }

    function mint(address to, uint256 amount) external {
        _mint(to, amount);
    }

    function burn(uint256 amount) public {
        if (amount == 0) return;
        _burn(msg.sender, amount);
    }

    function convert(address to, uint256 amount) public {
        if (amount == 0) return;
        if (to == address(0)) revert ZeroAddress();

        OMNI.safeTransferFrom(msg.sender, _DEAD_ADDRESS, amount);
        _mint(to, amount * CONVERSION_RATE);
    }

    function _constantNameHash() internal pure override returns (bytes32) {
        return 0xc72733118dabad3698b4044c2dc83c8c688bd907b50ed9d09d93a263878bf518;
    }
}
