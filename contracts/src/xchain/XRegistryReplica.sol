// SPDX-License-Identifier: GPL-3.0-only
pragma solidity =0.8.24;

import { Predeploys } from "../libraries/Predeploys.sol";
import { IOmniPortal } from "../interfaces/IOmniPortal.sol";
import { IOmniPortalSys } from "../interfaces/IOmniPortalSys.sol";
import { XRegistryBase } from "./XRegistryBase.sol";
import { XTypes } from "../libraries/XTypes.sol";

/**
 * @title XRegistryReplica
 * @notice XRegistry deployed, deployed alongside each portal. Only tracks registrations forwarded
 *         from the XRegistry predeploy on Omni's EVM.
 */
contract XRegistryReplica is XRegistryBase {
    /// @notice The OmniPortal contract
    address internal immutable portal;

    constructor(address portal_) {
        require(portal_ != address(0), "XRegistryReplica: no zero portal");
        portal = portal_;
    }

    modifier onlyXRegistry() {
        IOmniPortal omni = IOmniPortal(portal);

        // require the portal is the sender
        require(msg.sender == address(omni), "XReplica: not xcall");

        // require the xmsg sender is the XRegistry predeployed on omni evm
        XTypes.MsgShort memory xmsg = omni.xmsg();
        require(xmsg.sourceChainId == omni.omniChainId(), "XReplica: not from omni");
        require(xmsg.sender == Predeploys.XRegistry, "XReplica: not from XRegistry");

        _;
    }

    function set(uint64 chainId, string memory name, address registrant, Deployment calldata dep)
        public
        onlyXRegistry
    {
        _set(chainId, name, registrant, dep);

        // if OmniPortal registration for, intialize the new source chain on this chain's portal deployment
        if (_isPortalRegistration(name, registrant)) {
            uint64[] memory shards = abi.decode(dep.metadata, (uint64[]));
            IOmniPortalSys(portal).initSourceChain(chainId, shards);
        }
    }
}
