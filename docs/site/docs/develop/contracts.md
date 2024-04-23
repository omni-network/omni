---
sidebar_position: 4
---

# Contracts

A reference for Omni's user facing solidity contracts and libraries.

### [`OmniPortal`](https://github.com/omni-network/omni/blob/main/contracts/src/protocol/OmniPortal.sol)

- On-chain gateway into Omni's cross-chain messaging protocol
- Call contracts on another chain (`xcall`)
- Calculate fees for an `xcall`
- Read `xmsg` context when receiving an `xcall`

<details>
<summary>`IOmniPortal.sol` Reference Solidity Interface</summary>

```solidity
// SPDX-License-Identifier: GPL-3.0-only
pragma solidity ^0.8.12;

/**
 * @title IOmniPortal
 * @notice The OmniPortal is the on-chain interface to Omni's cross-chain
 *         messaging protocol. It is used to initiate and execute cross-chain calls.
 * @dev This snippet only shows functions meant to be called by users.
 */
interface IOmniPortal {
    /**
     * @notice Default xmsg execution gas limit, enforced on destination chain
     */
    function xmsgDefaultGasLimit() external view returns (uint64);

    /**
     * @notice Maximum allowed xmsg gas limit
     */
    function xmsgMaxGasLimit() external view returns (uint64);

    /**
     * @notice Minimum allowed xmsg gas limit
     */
    function xmsgMinGasLimit() external view returns (uint64);

    /**
     * @notice Maxium number of bytes allowed in xreceipt result
     */
    function xreceiptMaxErrorBytes() external view returns (uint64);


    /**
     * @notice Returns Chain ID of the chain to which this portal is deployed
     */
    function chainId() external view returns (uint64);

    /**
     * @notice Returns the current XMsg being executed via this portal.
     *          - xmsg().sourceChainId  Chain ID of the source xcall
     *          - xmsg().sender         msg.sender of the source xcall
     *         If no XMsg is being executed, all fields will be zero.
     *          - xmsg().sourceChainId  == 0
     *          - xmsg().sender         == address(0)
     */
    function xmsg() external view returns (XTypes.MsgShort memory);

    /**
     * @notice Returns true the current transaction is an xcall, false otherwise
     */
    function isXCall() external view returns (bool);

    /**
     * @notice Calculate the fee for calling a contract on another chain. Uses xmsgDefaultGasLimit.
     *         Fees denominated in wei.
     * @param destChainId   Destination chain ID
     * @param data          Encoded function calldata
     */
    function feeFor(uint64 destChainId, bytes calldata data) external view returns (uint256);

    /**
     * @notice Calculate the fee for calling a contract on another chain
     *         Fees denominated in wei.
     * @param destChainId   Destination chain ID
     * @param data          Encoded function calldata
     * @param gasLimit      Execution gas limit, enforced on destination chain
     */
    function feeFor(uint64 destChainId, bytes calldata data, uint64 gasLimit) external view returns (uint256);

    /**
     * @notice Call a contract on another chain Uses xmsgDefaultGasLimit as execution
     *         gas limit on destination chain
     * @param destChainId   Destination chain ID
     * @param to            Address of contract to call on destination chain
     * @param data          ABI Encoded function calldata
     */
    function xcall(uint64 destChainId, address to, bytes calldata data) external payable;

    /**
     * @notice Call a contract on another chain Uses provide gasLimit as execution gas limit on
     *          destination chain. Reverts if gasLimit < xmsgMinGasLimit or gasLimit > xmsgMaxGasLimit.
     * @param destChainId   Destination chain ID
     * @param to            Address of contract to call on destination chain
     * @param data          ABI Encoded function calldata
     * @param gasLimit      Execution gas limit, enforced on destination chain
     */
    function xcall(uint64 destChainId, address to, bytes calldata data, uint64 gasLimit) external payable;
}
```
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
 * @notice Base contract for Omni cross-chain applications
 */
contract XApp {
    /**
     * @notice The OmniPortal contract
     */
    IOmniPortal internal immutable omni;

    /**
     * @notice Transient storage for the current xmsg
     */
    XTypes.MsgShort internal xmsg;

    /**
     * @notice Read current xmsg into storage before execution, delete it afterwards
     */
    modifier xrecv() {
        xmsg = omni.xmsg();
        _;
        delete xmsg;
    }

    constructor(address _omni) {
        omni = IOmniPortal(_omni);
    }

    /**
     * @notice Return true if the current call is an xcall from the OmniPortal
     */
    function isXCall() internal view returns (bool) {
        return omni.isXCall() && msg.sender == address(omni);
    }

    /**
     * @notice Returns the fee for calling a contract on another chain. Uses OmniPortal.xmsgDefaultGasLimit
     */
    function feeFor(uint64 destChainId, bytes memory data) internal view returns (uint256) {
        return omni.feeFor(destChainId, data);
    }

    /**
     * @notice Retruns the fee for calling a contract on another chain, with the specified gas limit
     */
    function feeFor(uint64 destChainId, bytes memory data, uint64 gasLimit) internal view returns (uint256) {
        return omni.feeFor(destChainId, data, gasLimit);
    }

    /**
     * @notice Call a contract on another chain. Uses OmniPortal.xmsgDefaultGasLimit
     * @return fee The fee for the xcall
     */
    function xcall(uint64 destChainId, address to, bytes memory data) internal returns (uint256) {
        uint256 fee = omni.feeFor(destChainId, data);
        require(address(this).balance >= fee || msg.value >= fee, "XApp: insufficient funds");
        omni.xcall{ value: fee }(destChainId, to, data);
        return fee;
    }

    /**
     * @notice Call a contract on another chain, with the specified gas limit
     * @return fee The fee, denominated in wei
     */
    function xcall(uint64 destChainId, address to, bytes memory data, uint64 gasLimit) internal returns (uint256) {
        uint256 fee = omni.feeFor(destChainId, data, gasLimit);
        require(address(this).balance >= fee || msg.value >= fee, "XApp: insufficient funds");
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
