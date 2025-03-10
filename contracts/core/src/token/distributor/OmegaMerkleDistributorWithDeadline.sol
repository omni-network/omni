// SPDX-License-Identifier: GPL-3.0-only
pragma solidity 0.8.24;

import { MerkleDistributorWithDeadline } from "./MerkleDistributorWithDeadline.sol";

contract OmegaMerkleDistributorWithDeadline is MerkleDistributorWithDeadline {
    address internal constant VALIDATOR_1 = 0xCE624ce5C5717b63CED36AfB76857183E0a8a6eb;
    address internal constant VALIDATOR_2 = 0x98Eb13371c095905985cddE937018881d4D7f229;
    address internal constant VALIDATOR_3 = 0xe96cF9Ad91cD6dc911817603Dfb3c65d5f532B95;
    address internal constant VALIDATOR_4 = 0xdBd26a685DB4475b6c58ADEC0DE06c6eE387EAa8;

    constructor(
        address token_,
        bytes32 merkleRoot_,
        uint256 endTime_,
        address omniPortal_,
        address genesisStaking_,
        address solverNetInbox_
    ) MerkleDistributorWithDeadline(token_, merkleRoot_, endTime_, omniPortal_, genesisStaking_, solverNetInbox_) { }

    function _getValidator(address addr) internal pure override returns (address) {
        uint256 selection = uint160(addr) % 4;

        if (selection == 1) return VALIDATOR_1;
        if (selection == 2) return VALIDATOR_2;
        if (selection == 3) return VALIDATOR_3;
        return VALIDATOR_4;
    }
}
