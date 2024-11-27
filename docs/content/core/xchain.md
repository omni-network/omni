# XChain Messaging

Omni Core’s cross chain messaging layer enables developers to send function calls across chains using the Portal contract or the XApp framework.

## OmniPortal

The `OmniPortal` contracts are the backbone of Omni’s messaging layer, allowing developers to send and receive cross-chain function calls (`xcall`). These contracts are the endpoints for interacting with Omni’s cross chain messaging layer.

### **Key Functionality:**

**Outbound Calls (`xcall`)**:
  Developers can trigger cross-chain function calls by interacting with the `xcall` method in the Portal contract.

  Example:

  ```solidity
  omni.xcall(
      destChainId,  // Target chain ID
      conf,         // Confirmation level
      to,           // Destination contract address
      data,         // Encoded calldata
      gasLimit      // Gas limit for execution
  );
  ```

**Fee Calculation (`feeFor`)**:

  Use the `feeFor` method to calculate the cost of a cross-chain call based on calldata size and gas requirements.

 ```solidity
  uint256 fee = omni.feeFor(destChainId, data, gasLimit);
  ```
You can find the source code for `OmniPortal` [here](https://github.com/omni-network/omni/blob/main/contracts/core/src/xchain/OmniPortal.sol).

## Building with XApp

For developers looking to simplify cross-chain interactions, Omni provides the `XApp` contract, which handles the boilerplate logic for interacting with the Portal contracts.

Features include:

**Pre-Built Utilities**:

- `xcall`: Simplifies sending cross-chain calls, automatically handles fee payments.
- `feeFor`: Calculates the required fee for a call.
- `isXCall`: Detects whether a function is being executed via an inbound cross-chain call.

**Default Confirmation Level**:

- Don't worry about confirmation levels.

See more about `XApp` in the code base: [XApp](https://github.com/omni-network/omni/blob/main/contracts/core/src/pkg/XApp.sol)
