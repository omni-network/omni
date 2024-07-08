// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.24;

import { OmniPortal } from "src/xchain/OmniPortal.sol";
import { XTypes } from "src/libraries/XTypes.sol";

/**
 * @title PortalHarness
 * @dev A test contract that exposes OmniPortal internal functions, and allows state manipulation.
 */
contract PortalHarness is OmniPortal {
    function exec(XTypes.BlockHeader calldata xheader, XTypes.Msg calldata xmsg) external {
        _exec(xheader, xmsg);
    }

    function call(address to, uint64 gasLimit, bytes calldata data) external {
        _call(to, gasLimit, data);
    }

    function syscall(bytes calldata data) external {
        _syscall(data);
    }

    function setLatestValSetId(uint64 valSetId) external {
        latestValSetId = valSetId;
    }

    function setNetworkNoAuth(XTypes.Chain[] calldata network) external {
        _setNetwork(network);
    }
}
