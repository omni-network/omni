// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.12;

/**
 * @title MockPortal
 * @dev Stub portal that lets use us vm.expectCall to test the OmniAVS contract.
 */
contract MockPortal {
    function feeFor(uint64, /*destChainId*/ bytes calldata, /*data*/ uint64 /*gasLimit*/ )
        external
        pure
        returns (uint256)
    {
        return 1 gwei;
    }

    function xcall(uint64, /*destChainId*/ address, /*to*/ bytes calldata, /*data*/ uint64 /*gasLimit*/ )
        external
        payable
    { }
}
