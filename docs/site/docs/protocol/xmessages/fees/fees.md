---
sidebar_position: 1
---

# Fee Model

Omni implements a Solidity-based fee payment interface that accepts **\$ETH** as the payment currency. In the future, Omni will add support for fees payable in **\$OMNI** and other tokens to streamline the developer experience.

Fees are paid in **\$ETH** and calculated during transactions via the payable `xcall` function on the portal contracts, ensuring simplicity for developers and compatibility with existing Ethereum tooling. This setup allows for easy off-chain fee estimations and the possibility for developers to pass the cost on to users, with a straightforward upgrade path to a more dynamic fee structure that can adapt to the network's evolving needs without necessitating changes to developer contracts.

### Collection

Each portal will be configured with a `feeTo` address. All collected fees will be sent to this address. This address is set to the relayer address.

### Pricing

Portal contracts need to know how much to charge for each transaction, implemented in the `feeFor` method. The parameters to fee calculation are:

- Destination chain id
- Calldata
- Gas limit


## Fee Handling

Omni fees are charged synchronously on `xcall` call to a portal contract. This function is therefore `payable`. It allows specification of a custom gas limit, enforced at the destination chain.

```solidity
interface IOmniPortal {

  /**
   * @notice Call a contract on another chain
   * @dev Uses OmniPortal.xmsgDefaultGasLimit as execution gas limit on destination chain
   *      Fees are denomninated in wei, and paid via msg.value. Call reverts if fees
   *      are insufficient. Calculate fees with feeFor(...)
   * @param destChainId Destination chain ID
   * @param to Address of contract to call on destination chain
   * @param data Encoded function calldata
   */
  function xcall(uint64 destChainId, address to, bytes calldata data)
    external
    payable;

 /**
   * @notice Call a contract on another chain
   * @dev Uses provide gasLimit as execution gas limit on destination chain.
   *      Reverts if gasLimit < xmsgMinGasLimit or gasLimit > xmsgMaxGasLimit.
   *      Fees are denomninated in wei, and paid via msg.value. Call reverts
   * 	  if fees are insufficient. Calculate fees with feeFor(...)
   * @param destChainId Destination chain ID
   * @param to Address of contract to call on destination chain
   * @param data Encoded function calldata
   */
  function xcall(uint64 destChainId, address to, bytes calldata data, uint64 gasLimit)
    external
    payable;

    // ...
}
```

Native fees must be provided explicitly in the call. An interface must therefore be exposed to allow synchronous fee calculation. This interface is exposed via the portal.

```solidity
interface IOmniPortal {
  // ...

  /**
   * @notice Calculate the fee for calling a contract on another chain
   * @dev Uses OmniPortal.xmsgDefaultGasLimit. Fees denominated in wei.
   * @param destChainId Destination chain ID
   * @param data Encoded function calldata
   */
  function feeFor(uint64 destChainId, bytes calldata data)
    external
    view
    returns (uint256);

  /**
   * @notice Calculate the fee for calling a contract on another chain
   * @dev Fees denominated in wei
   * @param destChainId Destination chain ID
   * @param data Encoded function calldata
   * @param gasLimit Custom gas limit, enforced on destination chain
   */
  function feeFor(uint64 destChainId, bytes calldata data, uint64 gasLimit)
    external
    view
    returns (uint256);

    // ...
}
```
