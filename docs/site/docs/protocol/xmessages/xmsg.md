---
sidebar_position: 1
---

# `XMsg` Lifecycle

## Flow Diagram

The following steps provide a comprehensive overview of how an XMsg travels from a source rollup VM to a destination rollup VM. This process is visualized in Figure 4.

<figure>
  <img src="/img/xmsg.png" alt="xmsg" />
  <figcaption>*Relaying of `XMsg` through the network*</figcaption>
</figure>

## Steps

### 1. User Call

1. A user calls an xapp smart contract function on a rollup VM.

### 2. Smart Contract Logic

2. The smart contract converts the user’s logic into an `xcall` that is made on the rollup VM’s Portal contract. `xcall` are defined in Solidity as below (read more in the [develop section](../../develop/introduction.md)):

    ```solidity
    omni.xcall(
      uint64 destChainId,  // desintation chain id
      address to,           // contract address on the dest rollup to execute the call on
      bytes memory data     // calldata for the transaction, abi encoded
    )
    ```

### 3. Portal Contract Event Emission

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

### 4. Validator `XMsg` Packaging

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

### 5. Validator Attestation

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

### 6. Relayer Collects `XBlocks`

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

### 7. Relayer Submits `XBlocks`

7. For every finalized `XBlock` hash, relayers construct an `XBlock` submission containing the `XBlock` and the validator signatures for the `XBlock` hashes. Merkle proofs are generated to prove the inclusion of each `XMsg` in the `XBlock` hash. This ensures the decoupling of attestations from execution, thereby allowing the relayer to split the execution of an `XBlock` into many transactions so that it adheres to the constraints of the destination rollup VM.

### 8. Relayer Submits `XMsg`

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

### 9. Portal Verification and Forwarding

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

### 10. Smart Contract Execution

10. The receiving smart contracts execute the logic for their users.
