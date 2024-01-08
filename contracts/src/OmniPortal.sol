// SPDX-License-Identifier: GPL-3.0-only
pragma solidity 0.8.23;

import {IOmniPortal} from "./interfaces/IOmniPortal.sol";

contract OmniPortal is IOmniPortal {
    /// @inheritdoc IOmniPortal
    uint64 public constant XMSG_DEFAULT_GAS_LIMIT = 200_000;

    /// @inheritdoc IOmniPortal
    uint64 public chainId;

    /// @inheritdoc IOmniPortal
    mapping(uint64 => uint64) public outXStreamOffset;

    constructor(uint64 _chainId) {
        chainId = _chainId;
    }

    /// @inheritdoc IOmniPortal
    function xcall(uint64 destChainId, address to, bytes calldata data) external payable {
        _xcall(destChainId, msg.sender, to, data, XMSG_DEFAULT_GAS_LIMIT);
    }

    /// @dev Emit an XMsg event, increment dest chain outXStreamOffset
    function _xcall(uint64 destChainId, address sender, address to, bytes calldata data, uint64 gasLimit) private {
        emit XMsg(destChainId, outXStreamOffset[destChainId], sender, to, data, gasLimit);
        outXStreamOffset[destChainId] += 1;
    }
}
