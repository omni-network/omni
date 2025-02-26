// SPDX-License-Identifier: GPL-3.0-only
pragma solidity 0.8.24;

import { ERC20 } from "solady/src/tokens/ERC20.sol";

contract DebugMockOmni is ERC20 {
    constructor() {
        _mint(msg.sender, 10_000 ether);
        _mint(address(0xF6CDB1E733EA00D0eEa1A32F218B0ec76ABF1517), 5000 ether);
        _mint(address(0xBeD17aa3E1c99ea86e19e7B38356C54007BB6CDe), 5000 ether);
        _mint(address(0x2D61bE547b365BD5CdCc02920818492Fb7bdb765), 5000 ether);
        _mint(address(0xA6C9c842dc0C9C16338444e8bB77b885986Ef38b), 5000 ether);
    }

    function name() public pure override returns (string memory) {
        return "Mock Omni";
    }

    function symbol() public pure override returns (string memory) {
        return "OMNI";
    }

    function _constantNameHash() internal pure override returns (bytes32) {
        return keccak256(bytes("Mock Omni"));
    }
}
