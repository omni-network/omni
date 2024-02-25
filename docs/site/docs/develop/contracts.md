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

/// @dev OmniPortal interface (only user facing functions)
interface IOmniPortal {
    /**
     * @notice Default xmsg execution gas limit, enforced on destination chain
     * @return Gas limit
     */
    function XMSG_DEFAULT_GAS_LIMIT() external view returns (uint64);

    /**
     * @notice Maximum allowed xmsg gas limit
     * @return Maximum gas limit
     */
    function XMSG_MAX_GAS_LIMIT() external view returns (uint64);

    /**
     * @notice Minimum allowed xmsg gas limit
     * @return Minimum gas limit
     */
    function XMSG_MIN_GAS_LIMIT() external view returns (uint64);

    /**
     * @notice Chain ID of the chain to which this portal is deployed
     * @dev Used as sourceChainId for all outbound XMsgs
     * @return Chain ID
     */
    function chainId() external view returns (uint64);

    /**
     * @notice The current XMsg being executed via this portal
     * @dev If no XMsg is being executed, all fields will be zero
     * @return XMsg
     */
    function xmsg() external view returns (XTypes.Msg memory);

    /**
     * @notice Whether the current transaction is an xcall
     * @return True if current transaction is an xcall, false otherwise
     */
    function isXCall() external view returns (bool);

    /**
     * @notice Calculate the fee for calling a contract on another chain
     * @dev Uses OmniPortal.XMSG_DEFAULT_GAS_LIMIT
     * @dev Fees denominated in wei
     * @param destChainId Destination chain ID
     * @param data Encoded function calldata
     */
    function feeFor(uint64 destChainId, bytes calldata data) external view returns (uint256);

    /**
     * @notice Calculate the fee for calling a contract on another chain
     * @dev Fees denominated in wei
     * @param destChainId Destination chain ID
     * @param data Encoded function calldata
     * @param gasLimit Execution gas limit, enforced on destination chain
     */
    function feeFor(uint64 destChainId, bytes calldata data, uint64 gasLimit) external view returns (uint256);

    /**
     * @notice Call a contract on another chain
     * @dev Uses OmniPortal.XMSG_DEFAULT_GAS_LIMIT as execution gas limit on destination chain
     * @param destChainId Destination chain ID
     * @param to Address of contract to call on destination chain
     * @param data Encoded function calldata (use abi.encodeWithSignature
     * 	or abi.encodeWithSelector)
     */
    function xcall(uint64 destChainId, address to, bytes calldata data) external payable;

    /**
     * @notice Call a contract on another chain
     * @dev Uses provide gasLimit as execution gas limit on destination chain.
     *      Reverts if gasLimit < XMSG_MAX_GAS_LIMIT or gasLimit > XMSG_MAX_GAS_LIMIT
     * @param destChainId Destination chain ID
     * @param to Address of contract to call on destination chain
     * @param data Encoded function calldata (use abi.encodeWithSignature
     * 	or abi.encodeWithSelector)
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
 * @dev Base contract for Omni cross-chain applications
 */
contract XApp {
    /// @dev The OmniPortal contract
    IOmniPortal internal immutable omni;

    /// @dev Transient storage for the current xmsg
    XTypes.Msg internal xmsg;

    /// @dev Read current xmsg into storage before execution, delete it afterwards
    modifier xfunc() {
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
- `XTypes.Msg` is the only type end users interact with. It provides context

<details>
<summary>`XTypes.sol` Reference Solidity Code</summary>

```solidity
// SPDX-License-Identifier: GPL-3.0-only
pragma solidity ^0.8.12;

/// @dev Omni xchain types (only user facing types)
library XTypes {
    struct Msg {
        /// @dev Chain ID of the source chain
        uint64 sourceChainId;
        /// @dev Chain ID of the destination chain
        uint64 destChainId;
        /// @dev Monotonically incremented offset of Msg in source -> dest Stream
        uint64 streamOffset;
        /// @dev msg.sender of xcall on source chain
        address sender;
        /// @dev Target address to call on destination chain
        address to;
        /// @dev Data to provide to call on destination chain
        bytes data;
        /// @dev Gas limit to use for call execution on destination chain
        uint64 gasLimit;
    }
}
```

</details>
