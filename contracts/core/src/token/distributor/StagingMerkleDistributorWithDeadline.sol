// SPDX-License-Identifier: GPL-3.0-only
pragma solidity 0.8.24;

import { MerkleDistributorWithDeadline } from "./MerkleDistributorWithDeadline.sol";

contract StagingMerkleDistributorWithDeadline is MerkleDistributorWithDeadline {
    address internal constant VALIDATOR_1 = 0xD6CD71dF91a6886f69761826A9C4D123178A8d9D;
    address internal constant VALIDATOR_2 = 0x9C7bf21f72CA34af89F620D27E0F18C4366b88c6;

    constructor(
        address token_,
        bytes32 merkleRoot_,
        uint256 endTime_,
        address omniPortal_,
        address genesisStaking_,
        address solverNetInbox_
    ) MerkleDistributorWithDeadline(token_, merkleRoot_, endTime_, omniPortal_, genesisStaking_, solverNetInbox_) { }

    function _getValidator(address addr) internal pure override returns (address) {
        uint256 selection = uint160(addr) % 2;

        if (selection == 1) return VALIDATOR_1;
        return VALIDATOR_2;
    }
}
