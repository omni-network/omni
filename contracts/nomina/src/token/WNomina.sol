// SPDX-License-Identifier: GPL-3.0-only
pragma solidity 0.8.30;

import { ERC20 } from "solady/src/tokens/ERC20.sol";

contract WNomina is ERC20 {
    error NOMTransferFailed();

    /**
     * @notice Returns the name of the token.
     */
    function name() public pure override returns (string memory) {
        return "Wrapped Nomina";
    }

    /**
     * @notice Returns the symbol of the token.
     */
    function symbol() public pure override returns (string memory) {
        return "WNOM";
    }

    /**
     * @notice Deposits NOM in exchange for minting WNOM to the caller.
     */
    function deposit() public payable {
        _mint(msg.sender, msg.value);
    }

    /**
     * @notice Deposits NOM in exchange for minting WNOM to the specified address.
     * @param to The address to mint the WNOM to.
     */
    function depositTo(address to) public payable {
        _mint(to, msg.value);
    }

    /**
     * @notice Burns WNOM in exchange for withdrawing NOM to the caller.
     * @param amount The amount of WNOM to burn and withdraw.
     */
    function withdraw(uint256 amount) public {
        _withdraw(msg.sender, amount);
    }

    /**
     * @notice Burns WNOM in exchange for withdrawing NOM to the specified address.
     * @param to The address to withdraw the NOM to.
     * @param amount The amount of WNOM to burn and withdraw.
     */
    function withdrawTo(address to, uint256 amount) public {
        _withdraw(to, amount);
    }

    /**
     * @notice Burns WNOM in exchange for withdrawing NOM to the specified address.
     * @param to The address to withdraw the NOM to.
     * @param amount The amount of WNOM to burn and withdraw.
     */
    function _withdraw(address to, uint256 amount) internal {
        _burn(msg.sender, amount);
        /// @solidity memory-safe-assembly
        assembly {
            // Transfer the NOM to the specified address and confirm if it succeeded.
            if iszero(call(gas(), to, amount, codesize(), 0x00, codesize(), 0x00)) {
                mstore(0x00, 0x252db595) // `bytes4(keccak256("NOMTransferFailed()"))`.
                revert(0x1c, 0x04)
            }
        }
    }

    /**
     * @notice Returns a constant name hash to optimize permit gas costs in Solady ERC20.
     */
    function _constantNameHash() internal pure override returns (bytes32) {
        return 0x84fe092f00820265d8abbb39159a80f1dc751d5463fcd172fce598140f7b360c;
    }

    /**
     * @notice All native transfers trigger `deposit()`.
     */
    receive() external payable {
        deposit();
    }

    /**
     * @notice All native transfers trigger `deposit()`.
     */
    fallback() external payable {
        deposit();
    }
}
