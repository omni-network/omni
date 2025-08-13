// SPDX-License-Identifier: GPL-3.0-only
pragma solidity 0.8.24;

import { IFeeOracle } from "src/interfaces/IFeeOracle.sol";
import { Base } from "test/xchain/nomina/common/Base.sol";

/**
 * @title NominaPortal_feeFor_Test
 * @dev Test of NominaPortal feeFor functions
 */
contract NominaPortal_feeFor_Test is Base {
    /// @dev Test feeFor matches oracle
    function test_feeFor_succeeds() public view {
        uint64 destChainId = chainBId;
        uint64 gasLimit;
        bytes memory data = abi.encodeWithSignature("test()");

        assertEq(
            IFeeOracle(portal.feeOracle()).feeFor(destChainId, data, gasLimit),
            portal.feeFor(destChainId, data, gasLimit)
        );
    }
}
