// SPDX-License-Identifier: GPL-3.0-only
pragma solidity 0.8.23;

import { IOmniPortal } from "./interfaces/IOmniPortal.sol";
import { XChain } from "./libraries/XChain.sol";

contract OmniPortal is IOmniPortal {
    /// @inheritdoc IOmniPortal
    uint64 public constant XMSG_DEFAULT_GAS_LIMIT = 200_000;

    /// @inheritdoc IOmniPortal
    uint64 public constant XMSG_MAX_GAS_LIMIT = 5_000_000;

    /// @inheritdoc IOmniPortal
    uint64 public constant XMSG_MIN_GAS_LIMIT = 21_000;

    /// @inheritdoc IOmniPortal
    uint64 public immutable chainId;

    /// @inheritdoc IOmniPortal
    mapping(uint64 => uint64) public outXStreamOffset;

    constructor() {
        chainId = uint64(block.chainid);
    }

    /// @inheritdoc IOmniPortal
    function xcall(uint64 destChainId, address to, bytes calldata data) external payable {
        _xcall(destChainId, msg.sender, to, data, XMSG_DEFAULT_GAS_LIMIT);
    }

    /// @inheritdoc IOmniPortal
    function xcall(uint64 destChainId, address to, bytes calldata data, uint64 gasLimit)
        external
        payable
    {
        require(gasLimit <= XMSG_MAX_GAS_LIMIT, "OmniPortal: gasLimit too high");
        require(gasLimit >= XMSG_MIN_GAS_LIMIT, "OmniPortal: gasLimit too low");
        _xcall(destChainId, msg.sender, to, data, gasLimit);
    }

    /// @inheritdoc IOmniPortal
    function xsubmit(XChain.Submission calldata xsub) external {
        // TODO: verify a quorum of validators have signed off on the attestation root.

        // TODO: verify block header and msgs are included in the attestation merkle root

        // TODO: verify msgs are intended for this chain, and are next messages in stream

        // TODO: execute messages, emit receipts, update stream offsets
    }

    /// @dev Emit an XMsg event, increment dest chain outXStreamOffset
    function _xcall(
        uint64 destChainId,
        address sender,
        address to,
        bytes calldata data,
        uint64 gasLimit
    ) private {
        emit XMsg(destChainId, outXStreamOffset[destChainId], sender, to, data, gasLimit);
        outXStreamOffset[destChainId] += 1;
    }
}
