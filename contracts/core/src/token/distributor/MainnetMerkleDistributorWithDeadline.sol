// SPDX-License-Identifier: GPL-3.0-only
pragma solidity 0.8.24;

import { MerkleDistributorWithDeadline } from "./MerkleDistributorWithDeadline.sol";

contract OmegaMerkleDistributorWithDeadline is MerkleDistributorWithDeadline {
    address internal constant VALIDATOR_1 = 0x58D2A4e3880635B7682A1BB7Ed8a43F5ac6cFD3d;
    address internal constant VALIDATOR_2 = 0x19a4Cb685af95A96BEd67C764b6dB137978a5B17;
    address internal constant VALIDATOR_3 = 0xD5f9e687c1EA2b0Da7C06bEbe80ddAb03B33C075;
    address internal constant VALIDATOR_4 = 0x8be1aBb26435fc1AF39Fc88DF9499f626094f9AF;

    constructor(
        address token_,
        bytes32 merkleRoot_,
        uint256 endTime_,
        address omniPortal_,
        address genesisStaking_,
        address solverNetInbox_
    ) MerkleDistributorWithDeadline(token_, merkleRoot_, endTime_, omniPortal_, genesisStaking_, solverNetInbox_) { }

    function _getValidator() internal override returns (address) {
        uint256 selection = ++_delegationCount % 4;

        if (selection == 1) return VALIDATOR_1;
        if (selection == 2) return VALIDATOR_2;
        if (selection == 3) return VALIDATOR_3;
        return VALIDATOR_4;
    }
}
