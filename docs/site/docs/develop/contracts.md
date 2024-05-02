---
sidebar_position: 4
---

import CodeSnippet from '@site/src/components/CodeSnippet/CodeSnippet';

# Contracts

A reference for Omni's user facing solidity contracts and libraries.

### [`OmniPortal`](https://github.com/omni-network/omni/blob/main/contracts/src/protocol/OmniPortal.sol)

- On-chain gateway into Omni's cross-chain messaging protocol
- Call contracts on another chain (`xcall`)
- Calculate fees for an `xcall`
- Read `xmsg` context when receiving an `xcall`

<details>
<summary>`IOmniPortal.sol` Reference Solidity Interface</summary>

<CodeSnippet repoUrl="https://github.com/omni-network/omni/blob/main/contracts/src/interfaces/IOmniPortal.sol" />
</details>

### [`XApp`](https://github.com/omni-network/omni/blob/main/contracts/src/pkg/XApp.sol)

- Base contract for Omni cross-chain applications
- Simplifies sending / receiving `xcalls`

<details>
<summary>`XApp.sol` Reference Solidity Interface</summary>

```solidity
// SPDX-License-Identifier: GPL-3.0-only
pragma solidity ^0.8.12;

import { IOmniPortal } from "../interfaces/IOmniPortal.sol";
import { XTypes } from "../libraries/XTypes.sol";

/**
 * @title XApp
 * @dev Base contract for Omni cross-chain applications
 */
contract XApp {
    /// @dev The OmniPortal contract
    IOmniPortal internal immutable omni;

    /// @dev Transient storage for the current xmsg
    XTypes.MsgShort internal xmsg;

    /// @dev Read current xmsg into storage before execution, delete it afterwards
    modifier xrecv() {
        xmsg = omni.xmsg();
        _;
        delete xmsg;
    }

    constructor(address _omni) {
        omni = IOmniPortal(_omni);
    }

    /// @dev Return true if the current call is an xcall from the OmniPortal
    function isXCall() internal view returns (bool) {
        return omni.isXCall() && msg.sender == address(omni);
    }

    /// @dev Calculate the fee for calling a contract on another chain.
    ///      Uses OmniPortal.XMSG_DEFAULT_GAS_LIMIT
    /// @return fee The fee, denominated in wei
    function feeFor(uint64 destChainId, bytes memory data) internal view returns (uint256) {
        return omni.feeFor(destChainId, data);
    }

    /// @dev Calculate the fee for calling a contract on another chain.
    ///      Uses OmniPortal.XMSG_DEFAULT_GAS_LIMIT
    /// @return fee The fee, denominated in wei
    function feeFor(uint64 destChainId, bytes memory data, uint64 gasLimit) internal view returns (uint256) {
        return omni.feeFor(destChainId, data, gasLimit);
    }

    /// @dev Call a contract on another chain. Uses OmniPortal.XMSG_DEFAULT_GAS_LIMIT
    /// @return fee The fee for the xcall
    function xcall(uint64 destChainId, address to, bytes memory data) internal returns (uint256) {
        uint256 fee = omni.feeFor(destChainId, data);
        omni.xcall{ value: fee }(destChainId, to, data);
        return fee;
    }

    /// @dev Call a contract on another chain
    /// @return fee The fee, denominated in wei
    function xcall(uint64 destChainId, address to, bytes memory data, uint64 gasLimit) internal returns (uint256) {
        uint256 fee = omni.feeFor(destChainId, data, gasLimit);
        omni.xcall{ value: fee }(destChainId, to, data, gasLimit);
        return fee;
    }
}
```
</details>

### [`XTypes`](https://github.com/omni-network/omni/blob/main/contracts/src/libraries/XTypes.sol)

- Defines core xchain messaging types for the Omni protocol.
- `XTypes.MsgShort` is the only type end users interact with. It provides context

<details>
<summary>`XTypes.sol` Reference Solidity Code</summary>

```solidity
// SPDX-License-Identifier: GPL-3.0-only
pragma solidity ^0.8.12;

/// @dev Omni xchain types (only user facing types)
library XTypes {
    /**
     * @notice Trimmed version of Msg that presents the minimum required context for consuming xapps.
     * @custom:field sourceChainId  Chain ID of the source chain
     * @custom:field sender         msg.sender of xcall on source chain
     */
    struct MsgShort {
        uint64 sourceChainId;
        address sender;
    }
}
```

</details>
