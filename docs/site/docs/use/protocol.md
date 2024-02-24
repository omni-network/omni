---
sidebar_position: 2
---

# Protocol

Omni is designed to enhance the Ethereum ecosystem's scalability and interoperability. It integrates smart contracts across the Omni Chain EVM and Ethereum rollups for cross-chain interactions, with security ensured by a dPoS validator set, supported by a native token and re-staked ETH through Eigenlayer. This framework enables efficient cross-network storage and contract calls.

## `XMsg` Lifecycle

The following steps provide a comprehensive overview of how an XMsg travels from a source rollup VM to a destination rollup VM. This process is visualized in Figure 4.

<figure>
  <img src="/img/xmsg.png" alt="xmsg" />
  <figcaption>*Relaying of `XMsg` through the network*</figcaption>
</figure>

1. A user calls an xdApp smart contract function on a rollup VM.

2. The smart contract converts the user’s logic into an `xcall` that is made on the rollup VM’s Portal contract. `xcall` are defined in Solidity as below ([also shown in the overview](./overview.md#xcall-parameters)):

    ```solidity
    omni.xcall(
      uint64 destChainId,  // desintation chain id
      address to,           // contract address on the dest rollup to execute the call on
      bytes memory data     // calldata for the transaction, abi encoded
    )
    ```

3. The Portal contract converts the `xcall` into an `XMsg` and emits the corresponding `XMsg` event. `XMsg` events are defined in Solidity as:

    ```solidity
    event XMsg (
      address SourceMsgSender     // Sender on source chain, set to msg.Sender
      uint256 XStreamOffset       // Monotonically incremented offset of XMsg in the XStream
      uint256 DestChainID         // Target chain ID as per https://chainlist.org/
      address DestAddress         // Target/To address to "call" on destination chain
      bytes   Data                // Data to provide to "call" on destination chain
      uint256 DestGasLimit        // Gas limit to use for "call" on destination chain
    )
    ```

4. Validators package multiple `XMsg` into an `XBlock`. It is a deterministic one-to-one mapping. In `halo`, `XBlock` are typed as:

    ```go
    // XBlock represents the cross-chain properties of a source chain finalised block.
    type XBlock (
      uint         SourceChainID  // Source chain ID as per https://chainlist.org
      uint         BlockHeight    // Height of the source chain block
      bytes32      BlockHash      // Hash of the source chain block
      XMsg[]       Msgs           // XMsg events included in the source block.
      XReceipt[]   Receipts       // XReceipt events included in the source block.
    )
    ```

    `XBlock` structure provides the following properties to the Omni Network:

    - Succinctly verifiable merkle-multi-proofs for sub-ranges of `XMsg` per source-target pair allowing relayers to manage submission costs at single `XMsg` granularity.
    - Omni Consensus attestations are not required for source chain blocks without any cross chain messages (aka empty `XBlock`).
    - Relayer submissions are not required on destination chains for batches without cross chain messages.
    - The logic to create a `XBlock` is deterministic for any finalized source chain block height.

    <br />
    <details>
    <summary><code>XBlock</code> Storage and Calculation</summary>

    `XBlock` is not stored as they are deterministically calculated from a source blockchain. So in effect, the source blockchain stores them.
    Any component that depends on `XBlock`, calculates it from a source chain.

    $XBlock = f(chain_A)$ where $f(x)$ is a deterministic `pure` function that takes a finalized blockchain as input and produces `XBlock` as output.
    In practice, source blocks can be streamed and transformed using a simple translation function backed by an in-memory cache.

    </details>

    `XMsg` are associated with an `XStream`. An `XStream`  is a logical connection between a source and destination chain. It contains many `XMsg`, each with a monotonically incrementing `XStreamOffset` (the offset is like a EOA nonce, it is incremented for each subsequent message sent from a source chain to a destination chain). `XMsg`  are therefore uniquely identified and strictly ordered by their associated `XStream` and `XStreamOffset`.

    `XStreamOffset` allows for exactly-once delivery guarantees with strict ordering per source-destination chain pair.

5. Validators attest to `XBlock` hashes during consensus. Each Omni consensus layer validator monitors every finalized block for all source chains in `halo`. Validators need to wait for block finalization, or some other agreed-upon threshold, to ensure consistent and secure cross-chain messaging.

    `halo` attests via CometBFT vote extensions, all validators in the CometBFT validator set should attest to all `XBlock` (in addition to their normal validator duties). An attestation is defined by the following `Attestation` type:

    ```go
    type Attestation (
      // Composite primary key of XBlock + ValSet
      uint    SourceChainID    // SourceChainID of the XBlock
      uint    BlockHeight      // Source chain block height
      bytes32 BlockHash        // Source chain block hash
      uint    ValidatorSetID   // Unique identifier of current validator set

      bytes32 XBlockRoot       // Merkle root of the XBlock
      bytes32 ValidatorPubkey  // Validator identifier
      bytes32 Signature        // Validator signature of XBlockHash
    )
    ```

    Validators return an array of `Attestation` during the ABCI++ `ExtendVote` method.

6. Relayers monitoring the Omni consensus layer gather finalized `XBlock` hashes and their corresponding `XBlock` and `XMsg`. Finalized blocks without any `XMsg` events are ignored. The Relayer submits `XMsg`s to destination chain that have ⅔ of validator signatures.

    For each destination chain, the relayer decides how many `XMsg` to submit, which defines the “cost” of transactions being submitted to the destination chain. This is primarily defined by the data size and gas limit of the messages and the portal contract verification and processing overhead.

    A merkle-multi-proof is generated for the set of identified `XMsg` that match the quorum `XBlock` attestations root. The relayer submits an EVM transaction to the destination chain, ensuring it gets included on-chain as soon as possible. The transaction contains the following data:

    <!-- TODO: rename after refactoring of Attestation and AggregateAttestation -->

    ```go
    type Submission (
      AggregateAttestation att    // Aggregate attestation containing quorum signatures for a specific validator set.
      XMsg[]               msgs   // Subset of XMsgs in a XBlock (could also be all)
      bytes                Proofs // Merkle-multi-proof for XMsgBatch
    )
    ```

7. For every finalized `XBlock` hash, relayers construct an `XBlock` submission containing the `XBlock` and the validator signatures for the `XBlock` hashes. Merkle proofs are generated to prove the inclusion of each `XMsg` in the `XBlock` hash. This ensures the decoupling of attestations from execution, thereby allowing the relayer to split the execution of an `XBlock` into many transactions so that it adheres to the constraints of the destination rollup VM.

8. The relayer delivers its submissions to the Portal contract on the destination rollup VM. Relayer submissions are defined in Go as:

    ```go
    // Submission is a cross-chain submission of a set of messages and their proofs.
    type Submission struct {
      AttestationRoot common.Hash // Merkle root of the attestations
      ValidatorSetID  uint64      // Unique identified of the validator set included in this aggregate.
      BlockHeader     BlockHeader // BlockHeader identifies the cross-chain Block
      Msgs            []Msg       // Messages to be submitted
      Proof           [][32]byte  // Merkle multi proofs of the messages
      ProofFlags      []bool      // Flags indicating whether the proof is a left or right proof
      Signatures      []SigTuple  // Validator signatures and public keys
      DestChainID     uint64      // Destination chain ID, for internal use only
    }
    ```

9. The receiving Portal contract verifies the `XBlock`’s validator signatures and `XMsg` merkle proofs before passing all verified `XMsg`s to their destination smart contracts on the rollup VM. After validating and processing the submitted `XMsg`, the portal contract emits an `XReceipt` event. This marks the `XMsg` as “successful” or “reverted” by `halo`. `XMsg` can revert if the gas limit was exceeded or if target address smart contract logic reverted for other reasons. `XReceipt` are included in `XBlock` (same as `XMsg`).

    ```go
    type XReceipt (
      uint    SourceChainID         // The cross-chain message's source chain
      uint    XStreamOffset         // Offset of XMsg in the XStream
      uint    GasUsed               // Gas used dueing message "call"
      uint    Result                // 0 for success, 1 for revert
      address RelayerAddress        // Address of relayer that submitted the message
    )
    ```

10. The receiving smart contracts execute the logic for their users.

## Validator Components & Communication

Omni nodes are configured with a new modular framework based on [Ethereum’s Engine API](https://github.com/ethereum/execution-apis/tree/4b225e0d273e92982b2c539d63eaaa756c5285a4/src/engine) to combine `halo` with the EVM execution client.

<figure>
  <img src="/img/validator.png" alt="Validator" />
  <figcaption>*Overview of the Components Run by Validators and Interactivity*</figcaption>
</figure>

Using the Engine API, Omni nodes pair existing high performance Ethereum execution clients with a new consensus client, referred to as `halo`, that implements CometBFT consensus.The Engine API allows clients to be substituted or upgraded without breaking the system. This allows the protocol to maintain flexibility on Ethereum and Cosmos technology while promoting client diversity within its execution layer and consensus layer. We consider this new network framework to be a public good that future projects may leverage for their own network designs.

### Halo

It implements the server side of the ABCI++ interface and drives the Omni Execution Layer via the Engine API. CometBFT validators attest to source chain blocks containing cross chain messages using CometBFT Vote Extensions.

Omni's consensus layer is established through the Halo consensus client which supports a Delegated Proof of Stake (DPoS) consensus mechanism implemented by the [CometBFT consensus engine](https://docs.cometbft.com/v0.38/) (formerly known as Tendermint). CometBFT stands out as the optimal consensus engine for Omni nodes for three reasons:

1. CometBFT achieves immediate transaction finalization upon block inclusion, eliminating the need for additional confirmations. This feature allows Omni validators to agree on a single global state for all connected networks, updated every second.
2. CometBFT provides support for DPoS consensus mechanisms, allowing re-staked ETH to be delegated directly to individual Omni validators as part of the network’s security model.
3. CometBFT is one of the most robust and widely tested PoS consensus models and is already used in many production blockchain networks securing billions of dollars.

By default, Omni validators ensure the integrity of rollup transactions by awaiting their posting and finalization on the Ethereum L1 before attesting to them. This proactive approach mitigates the risk of reorganizations on both Ethereum L1 and the rollup, enhancing the overall security and reliability of the system.

#### Consensus Process

Omni Consensus clients process the consensus blocks and maintain consensus state that tracks the “status” of each `XBlock` (by the validator set) from `Pending` to `Approved` (including an `AggregateAttestation`). Then only the “latest approved” `XBlock` for each source chain needs to be maintained; earlier `XBlock`s can be trimmed from the state. Validators in the current validator set must attest to all subsequent (after the “last approved”) `XBlock`.

#### Validator Set Changes

When the validator set changes, all `XBlock`s marked as `Pending` need to be updated by:

- Updating the associated validator set ID to the current.
- Deleting all attestations by validators not in the current set.
- Updating the weights of each remaining attestation according to the new validator set.

Validators that already attest to the `Pending` marked `XBlock` during the previous validator set, do not need to re-attest. Only the new set validators must attest (ie. to all `XBlock`s after the latest approved).

#### ABCI++

Cosmos’ [ABCI++](https://docs.cometbft.com/v0.37/spec/abci/) is a supplementary tool that allows `halo` clients to adapt the CometBFT consensus engine to the node’s architecture. An ABCI++ adapter is wrapped around the CometBFT consensus engine to convert messages from the Engine API into a format that can be used in CometBFT consensus. These messages are inserted into CometBFT blocks as single transactions – this makes Omni consensus lightweight and enables Omni’s sub-second finality time.

During consensus, validators also use ABCI++ to attest to the state of external Rollup VMs. Omni validators run the state transition function, $f(i, S_n)$, for each Rollup VM and compute the output, $S_{n+1}$.

#### Parallelized Consensus & CometBFT

Omni introduces Parallelized Consensus, a consensus framework that allows validators to run consensus for the Omni EVM in parallel with consensus for cross-network messages without compromising on performance.

Omni’s Parallelized Consensus contains two sub-processes:

- validating state changes within the Omni EVM and
- attesting to `XBlock` hashes originating from external rollup VMs.

The aggregate Parallelized Consensus process is visualized below.

<figure>
  <img src="/img/parallel-consensus.png" alt="Parallel Consensus" />
  <figcaption>*Parallelized Consensus: Validating Omni EVM State Changes and Rollup `XBlock` Hash Attestation*</figcaption>
</figure>

1. Every `halo` client runs a node for each rollup VM to check for `XMsg` events emitted from Portal contracts.
2. For every rollup VM block that contains `XMsg`s, `halo` clients build `XBlock`s that contain the corresponding `XMsg`s.
3. Once the calldata for a rollup VM block has been posted and finalized on Ethereum L1, Omni validators use ABCI++ vote extensions to attest to the hash of the corresponding `XBlock`. These attestations are appended to the current consensus layer block.

### EngineAPI

The Engine API allows the Omni execution layer (Omni EVM) to mirror the functionality and design of Ethereum L1’s execution layer (EVM). Users send transactions to the Omni EVM mempool and execution clients share those transactions through the execution layer’s peer-to-peer (P2P) node network. For each block of transactions, a node’s execution client computes the state transition function for the Omni EVM and shares the output with its `halo` client through the Engine API. Since the transaction mempool resides on the execution layer, Omni can scale activity without congesting the CometBFT mempool used by validators on the consensus layer. Previously, using the CometBFT mempool to handle transaction requests caused the network to become overloaded and resulted in liveness disruptions. Similar challenges have been observed in other projects that adopted a comparable approach.

### Execution Consensus

1. When it is a validator's turn to propose a block, its `halo` client requests the latest Omni EVM block from its execution client using the Engine API.
2. The execution client builds a block from the transactions in its mempool and returns the block header to the `halo` client through the Engine API.
3. The `halo` client packages the new block proposal as a single CometBFT transaction and includes it in the consensus layer block.
4. The block is proposed to the rest of the validator network through the consensus layer’s P2P network.
5. Non-proposing validators use the Engine API and their execution clients to run the state transition function on the proposed block header to verify the block’s validity.

### Omni EVM

Post-merge ethereum decoupled the execution layer from the consensus layer introducing a more modular approach to building blockchains. This modular approach allows the EVM to scale (somewhat) independently from consensus, by simply adopting the latest performant execution client like `erigon` or `geth`.

The Omni consensus layer needs smart contracts to manage native staking and delegated re-staking from ETH L1. The Omni EVM is a natural fit as fees would be much lower and syncing with the consensus layer is already built-in. Providing an EVM purposely built for cross-chain dapps that has both low fees and short block times allows for a simple adoption path and hub-and-spoke mental model to onboard projects into Omni Protocol.

The Omni execution layer also benefits from using existing Ethereum execution clients without specialized modifications. This approach eliminates the risk of introducing new bugs and increases release velocity for features strictly related to cross-rollup interoperability. Furthermore, Omni nodes can seamlessly adopt upgrades from any EVM client, ensuring ongoing compatibility with the Omni consensus layer. For example, the Omni EVM natively supports dynamic transaction fees and partial fee burning through its support for [EIP-1559](https://eips.ethereum.org/EIPS/eip-1559). In contrast, frameworks like Ethermint have faced delays spanning multiple years due to challenges in adapting EVM upgrades.

## Relayer

The Relayer is a permissionless actor that submits cross chain messages to destination chains and monitors the Omni Consensus Layer until ⅔ (>66%) of the validator set attested to the next block on each source chain. It then submits the applicable cross-chain messages to each destination chain providing the quorum validator signatures and a multi-merkle-proof. Will be incentivized in future versions of the network.

The Relayer is also responsible for monitoring attestations. Similar to validators, relayers should maintain an `XBlock` cache. I.e., track all source chain blocks, convert them to `XBlock`s, cache them, and make them available for internal indexed querying. Relayers monitor the Omni Consensus Chain state for attested `XBlock`.

For each destination chain, the relayer has to decide how many `XMsg`s to submit, which defines the [cost](./protocol.md#fee-handling) of transactions being submitted to the destination chain. This is primarily defined by the data size and gas limit of the messages and the portal contract verification and processing overhead.

A merkle-multi-proof is generated for the set of identified `XMsg`s that match the quorum `XBlock` attestations root. The relayer submits a EVM transaction to the destination chain, ensuring it gets included on-chain as soon as possible. The transaction contains the following data:

```go
type Submission (
  AggregateAttestation att    // Aggregate attestation containing quorum signatures for a specific validator set.
  XMsg[]               msgs   // Subset of XMsgs in a XBlock (could also be all)
  bytes                Proofs // Merkle-multi-proof for XMsgBatch
)
```

## Portal Contract

Portal Contracts are a set of smart contracts that implement the on-chain logic of the Omni protocol and are deployed to all supported Rollup EVMs as well as the Omni EVM. They provide the main interface to call a cross-chain smart contract which results in a cross-chain message being emitted as a `XMsg`. They also provides a "pay at source" fee mechanism using the source chain’s native token. They track the omni consensus validator set, used to verify submitted cross chain message attestations ([more below](#submission-validation)).

### Relaying Methods

The Portal Contract exposes both an `xcall` method for cross-chain calls to another chain, and an `xsubmit` method for receiving calls from another chain.

- `xcall` is called by the `omni.xcall()` helper and requires the parameters listed above, with the Portal Contract emitting an `XMsg` Event for every call to it.
- `xsubmit` accepts only calls from valid signatures (currently the Relayer) and requires the destination contract `address` and `data` (containing `method` and method params `data`).

See more on performing cross chain transactions in the [developer section](../develop/contracts.md).

### `Submission` Validation

Portal Contracts keep a cursor for each source chain that:

- Tracks the latest valid `Submission`’s `XBlockHash` that contained valid `XMsgBatch` to the local destination chain.
- The total messages in that batch.
- The index of the last message that was submitted.
- And implicitly, whether the latest `XBlock` is partially or completely submitted.

They validate the `AggregateAttestation` data:

- Ensuring the `SourceChainID` is known
- Ensuring the `ValidateSetID` is known and the validator set if available.
- If the cursor is partial, ensuring the `XBlockHash` matches that of the cursor.

Validate the `XMsgBatch` data:

- Ensuring the `TargetChainID` matches the local chain ID.
- If the cursor is complete, ensure the `ParentBlockHash` matches that of the cursor.
- Ensure the included `WrappedXMsg` indexes start from 0 if the cursor is complete or follow on the cursor index if partial.

Verify the `AggregateAttestation` signatures:

- Verify all validator signatures over the root `XBlockHash`
- Ensure that quorum is reached; more than 66% validators in the set signed.
- Verify a merke-multi-proof against the `XBlockHash` that proves the following fields of the `XBlock`:
  - All fields used in above validator.
  - All included `XMsgBatch` hashes.

## Fee Handling

Omni fees are charged synchronously on `xcall` call to a portal contract. This function is therefore `payable`. It allows specification of a custom gas limit, enforced at the destination chain.

```solidity
interface IOmniPortal {

  /**
   * @notice Call a contract on another chain
   * @dev Uses OmniPortal.XMSG_DEFAULT_GAS_LIMIT as execution gas limit
   *      on destination chain
   * @dev Fees are denomninated in wei, and paid via msg.value. Call reverts
   * 	   if fees are insufficient. Calculate fees with feeFor(...)
   * @param destChainId Destination chain ID
   * @param to Address of contract to call on destination chain
   * @param data Encoded function calldata (use abi.encodeWithSignature
   * 	or abi.encodeWithSelector)
   */
  function xcall(uint64 destChainId, address to, bytes calldata data)
external
payable;

 /**
   * @notice Call a contract on another chain
   * @dev Uses provide gasLimit as execution gas limit on destination chain.
   *      Reverts if gasLimit < XMSG_MAX_GAS_LIMIT or gasLimit > XMSG_MAX_GAS_LIMIT
   * @dev Fees are denomninated in wei, and paid via msg.value. Call reverts
   * 	   if fees are insufficient. Calculate fees with feeFor(...)
   * @param destChainId Destination chain ID
   * @param to Address of contract to call on destination chain
   * @param data Encoded function calldata (use abi.encodeWithSignature
   * 	or abi.encodeWithSelector)
   */
  function xcall(uint64 destChainId, address to, bytes calldata data, uint64 gasLimit)
external
payable;

}
```

Native fees must be provided explicitly in the call. An interface must therefore be exposed to allow synchronous fee calculation. This interface is exposed via the portal.

```solidity
interface IOmniPortal {
  // ...

  /**
   * @notice Calculate the fee for calling a contract on another chain
   * @dev Uses OmniPortal.XMSG_DEFAULT_GAS_LIMIT
   * @dev Fees denominated in wei
   * @param destChainId Destination chain ID
   * @param data Encoded function calldata (use abi.encodeWithSignature
   * 	or abi.encodeWithSelector)
   */
  function feeFor(uint64 destChainId, bytes calldata data)
external
view
returns (uint256);

  /**
   * @notice Calculate the fee for calling a contract on another chain
   * @dev Fees denominated in wei
   * @param destChainId Destination chain ID
   * @param data Encoded function calldata (use abi.encodeWithSignature
   * 	or abi.encodeWithSelector)
   * @param gasLimit Custom gas limit, enforced on destination chain
   */
  function feeFor(uint64 destChainId, bytes calldata data, uint64 gasLimit)
external
 view
returns (uint256);

    // ...
}
```

### Collection

Each portal will be configured with a `feeTo` address. All collected fees will be sent to this address. This address is set to the relayer address.

### Pricing

Portal contracts need to know how much to charge for each transaction, implemented in the `feeFor` method. The parameters to fee calculation are:

- Destination chain id
- Calldata
- Gas limit
