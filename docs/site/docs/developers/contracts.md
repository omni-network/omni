---
sidebar_position: 1
id: contracts
---

# Contracts Overview

This section is an overview of the system contracts that can be used in an Omni system. See below for greater details on method signatures and how to use them.

## `OmniScient.sol`

- Inherited by your global contracts on the Omni EVM.
- Exposes methods to interact with the sources of transactions and send transactions to other rollups.
  - `omni.txSourceChain()` - the source of the current transaction.
  - `omni.sendTx(...)` - send a transaction to another rollup.
- Exposes virtual methods that can be overwritten to receive callback functions upon success or failure of cross-rollup transactions. These methods are not required, but may be useful for richer expression of results in your smart contracts.
  - `function onXChainTxSuccess(...)` - called upon successful execution of a cross-rollup transaction.
  - `function onXChainTxReverted(...)` - called upon a reverted execution of a cross-rollup transaction.

## `OmniPortal.sol`

- An Omni system contract deployed to all integrated rollups.
- Exposes methods to interact with the Omni EVM and other rollups.
  - `function sendOmniTx(...)` - send a transaction to Omni.
  - `function sendXChainTx(...)` - send a transaction to another rollup.
  - `function verifyOmniState(...)` - verify the state of a smart contract on Omni.

## How We Recommend Using the System Contracts

### On the Omni EVM

#### Basics

- You should store your application’s global logic within Omni’s EVM. To interact with interoperability primitives, import the `OmniScient.sol` contract and extend its class in your contracts.
- You may want some functions to only be called through the Omni EVM, and some to only be called from other rollups. For this, check the `omni.txSourceChain()`.
  - For a contract you only want to be executed from omni, ensure: `require(omni.isOmniTx())`
  - For a contract you only want to be executed from other rollups, ensure: `require(omni.isExternalTx())`
- You can also require that some functions are only called from specific domains. For example, you could use:
  ```solidity
  require(isTxFrom("arbitrum-goerli"))

  // OR

  require(isTxFromOneOf("arbitrum-goerli", "optimism-goerli"))
  // this is an overloaded function, you can provide up to 5 arguments
  ```

You’ll likely want to track where transactions are coming from in your global state. You can use `omni.txSourceChain()` in a mapping or other data structures.

#### Cross Rollup Transactions

To send a transaction to another rollup, use:

```solidity
omni.sendTx(
    string memory _chain,  // rollup identifier, ex 'optimism-goerli'
    address _to,           // contract address on the dest rollup to execute the call on
    bytes memory _data     // calldata for the transaction, abi encoded, 10kb limit
)
```

#### Callbacks

When you use the `omni.sendTx(...)` method to send a cross rollup transaction, you can opt to receive a callback function based on the successful or unsuccessful execution of that transaction on the destination chain. You can override these virtual functions with your own logic.

When you override these functions, make sure to keep the `onlyOmni` modifier in your logic. This will ensure that only the Omni system can execute these callback functions.

The method signatures are:

```solidity
function onXChainTxSuccess(
    OmniCodec.Tx memory _xtx,
    address _sender,
    bytes memory _returnValue,
    uint256 _gasSpent
) virtual external onlyOmni {}

function onXChainTxReverted(
    OmniCodec.Tx memory _xtx,
    address _sender,
    uint256 _gasSpent
) virtual external onlyOmni {}
```

If you’d like to use these callbacks, you should import OmniCodec library because the first argument is `OmniCodec.Tx`:

```solidity
struct Tx {
    bytes32 sourceTxHash;
    string sourceChain;
    string destChain;
    uint64 nonce;
    address from;
    address to;
    uint256 value;
    uint256 paid;
    uint64 gasLimit;
    bytes data;
}
```

and you can use these values however you'd like.

### On Rollups

#### Sending Transactions

To send transactions _from_ another rollup, your smart contracts on those rollups must interact with an `OmniPortal` contract, a singleton contract deployed on each rollup. Check the docs for relevant addresses.

You can either execute a transaction on Omni, or on another rollup. We recommend sending a transaction to Omni if you’d like to update global state, and sending a transaction directly to other rollups for other function types that don’t require global state updates.

To send a transaction to Omni, use:

```solidity
function sendOmniTx(
    address _to,
    bytes memory _data
) public;
```

To send transactions from another rollup, your smart contracts on those rollups must interact with an `OmniPortal` contract, a singleton contract deployed on each rollup. Check the docs for relevant addresses.

You can either execute a transaction on Omni, or on another rollup. We recommend sending a transaction to Omni if you’d like to update global state, and sending a transaction directly to other rollups for other function types that don’t require global state updates.

```solidity
function sendOmniTx(
    address _to,
    bytes memory _data
) public;
```

```solidity
function sendXChainTx(
    string memory _chain,
    address _to,
    bytes memory _data
) public;

```

#### Verifying Omni Contract State

From each rollup, you can also check the state of your global contracts on Omni. For example, you may want to check the value of some state variable in your Omni EVM contract, and then execute certain logic based on that value.

:::note

This functionality only works on rollups with the sha256 opcode enabled (Scroll does not currently support this opcode).

:::

```solidity
function verifyOmniState(
        uint64 _blockNumber,
        bytes memory _storageProof,
        bytes memory _storageKey,
        bytes memory _storageValue
) public view returns (bool)
```

The block number is the block you’re interested in checking the state against. The `_storageProof` can be queried from an Omni RPC with the `eth_getProof` method, and the `_storageKey`, and `_storageValue` inputs are based on your implemented contracts.
