// SPDX-License-Identifier: GPL-3.0-only
pragma solidity ^0.8.24;

import { ERC721 } from "solady/src/tokens/ERC721.sol";
import { LibString } from "solady/src/utils/LibString.sol";

contract MockERC721 is ERC721 {
    using LibString for string;
    using LibString for uint256;

    string private _NAME;
    string private _SYMBOL;
    string private _BASE_URI;

    uint256 public totalSupply;

    constructor(string memory _name, string memory _symbol, string memory _baseURI) {
        _NAME = _name;
        _SYMBOL = _symbol;
        _BASE_URI = _baseURI;
    }

    function name() public view override returns (string memory) {
        return _NAME;
    }

    function symbol() public view override returns (string memory) {
        return _SYMBOL;
    }

    function baseURI() public view returns (string memory) {
        return _BASE_URI;
    }

    function tokenURI(uint256 tokenId) public view override returns (string memory) {
        return baseURI().concat(tokenId.toString());
    }

    function mint() external {
        _mint(msg.sender, ++totalSupply);
    }

    function mintTo(address to) external {
        _mint(to, ++totalSupply);
    }
}
