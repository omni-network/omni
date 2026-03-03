// SPDX-License-Identifier: GPL-3.0-only
pragma solidity 0.8.24;

import { Nomina } from "./Nomina.sol";

/**
 * @title NominaMintLock
 * @notice Permanently locks the ability to mint NOM tokens.
 * @dev Once this contract accepts mint authority from the Nomina token, the mint authority is
 *      irrevocably held by this contract. Since this contract has no ability to set a minter or
 *      transfer the mint authority, no new NOM tokens can ever be minted.
 */
contract NominaMintLock {
    /**
     * @notice The Nomina token contract.
     */
    Nomina public immutable NOMINA;

    constructor(Nomina _nomina) {
        NOMINA = _nomina;
    }

    /**
     * @notice Accepts the mint authority from the Nomina token, permanently locking minting.
     * @dev After this call, no new NOM tokens can ever be minted, as this contract has no
     *      mechanism to set a minter or transfer the mint authority.
     */
    function acceptMintAuthority() external {
        NOMINA.acceptMintAuthority();
    }
}
