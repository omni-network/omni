// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.24;

import { IFeeOracle } from "src/interfaces/IFeeOracle.sol";
import { Base } from "./common/Base.sol";

/**
 * @title OmniPortal_feeFor_Test
 * @dev Test of OmniPortal feeFor functions
 */
contract OmniPortal_feeFor_Test is Base {
    /// @dev Test feeFor matches oracle
    function test_feeFor_succeeds() public {
        uint64 destChainId = 1;
        uint64 gasLimit;
        bytes memory data = abi.encodeWithSignature("test()");

        assertEq(
            IFeeOracle(portal.feeOracle()).feeFor(destChainId, data, gasLimit),
            portal.feeFor(destChainId, data, gasLimit)
        );
    }

    /// @dev Test feeFor with default gasLimit matches oracle
    function test_feeFor_defaultGasLimit_succeeds() public {
        uint64 destChainId = 1;
        bytes memory data = abi.encodeWithSignature("test()");

        assertEq(
            IFeeOracle(portal.feeOracle()).feeFor(destChainId, data, portal.XMSG_DEFAULT_GAS_LIMIT()),
            portal.feeFor(destChainId, data)
        );
    }
}
