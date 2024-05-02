// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.24;

import { Predeploys } from "../libraries/Predeploys.sol";
import { IOmniPortal } from "../interfaces/IOmniPortal.sol";
import { XRegistryBase } from "./XRegistryBase.sol";
import { XTypes } from "../libraries/XTypes.sol";

/**
 * @title XRegistryReplica
 * @notice XRegistry deployed, deployed alongside each portal. Only tracks registrations forwarded
 *         from the XRegistry predeploy on Omni's EVM.
 */
contract XRegistryReplica is XRegistryBase {
    /// @notice The OmniPortal contract
    IOmniPortal internal immutable omni;

    constructor(address omni_) {
        omni = IOmniPortal(omni_);
    }

    modifier onlyXRegistry() {
        // require the portal is the sender
        require(msg.sender == address(omni), "XReplica: not xcall");

        // require the xmsg sender is the XRegistry predeployed on omni evm
        XTypes.MsgShort memory xmsg = omni.xmsg();
        require(xmsg.sourceChainId == omni.omniChainId(), "XReplica: not from omni");
        require(xmsg.sender == Predeploys.XRegistry, "XReplica: not from XRegistry");

        _;
    }

    function set(uint64 chainId, string memory name, address registrant, address addr) public onlyXRegistry {
        _set(chainId, name, registrant, addr);
    }
}
