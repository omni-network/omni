---
sidebar_position: 1
---

# `XMsg` Lifecycle

Cross-rollup messages are referred to as `XMsg` in the Omni protocol. Omni uses CometBFT to process `XMsg`s according to the following sequence.

## Flow Diagram

The following steps provide a comprehensive overview of how an XMsg travels from a source rollup VM to a destination rollup VM. This process is visualized in Figure 4.

<figure>
  <img src="/img/xmsg.png" alt="xmsg" />
  <figcaption>*Relaying of `XMsg` through the network*</figcaption>
</figure>

## Steps

### 1. User Call

- A user calls an xdapp smart contract function on a rollup VM.

### 2. Smart Contract Logic

- The smart contract converts the user’s logic into an `xcall` that is made on the rollup VM’s Portal contract. An `xcall` is defined in Solidity below:

    ```solidity
    omni.xcall(
      destChainId,  // destination chain id, e.g. 1 for Ethereum mainnet
      to,           // contract address on the destination chain
      data,         // abi encoded calldata, ex abi.encodeWithSignature("foo()")
      gasLimit      // (optional) gas limit for the call on the destination chain
    )
    ```

### 3. Portal Contract Event Emission

- The Portal contract converts the `xcall` into an `XMsg` and emits the corresponding `XMsg` event. `XMsg` events are defined in Solidity as:

    ```solidity
    event XMsg(
      uint64 indexed destChainId,   // Destination chain ID
      uint64 indexed streamOffset,  // Offset of XMsg in the XStream
      address sender,               // Address of the sender
      address to,                   // Address of the recipient
      bytes data,                   // ABI encoded calldata
      uint64 gasLimit               // Gas limit for the call on the destination chain
    );
    ```

### 4. Validator `XMsg` Packaging

- For each rollup VM block, validators package all `XMsg`s into a corresponding `XBlock`  using a deterministic 1:1 mapping. In `halo`, `XBlock` are typed as:

    ```go
    // Block is a deterministic representation of the omni cross-chain properties of a source chain EVM block.
    type Block struct {
      BlockHeader
      Msgs      []Msg     // All cross-chain messages sent/emittted in the block
      Receipts  []Receipt // Receipts of all submitted cross-chain messages applied in the block
      Timestamp time.Time // Timestamp of the source chain block
    }
    ```

    `XBlock` structure provides the following properties to the Omni Network:

    - Succinctly verifiable merkle-multi-proofs for sub-ranges of `XMsg` per source-target pair allowing relayers to manage submission costs at single `XMsg` granularity.
    - Omni Consensus attestations are not required for source network blocks without any `XMsg` requests (aka empty `XBlock`).
    - Relayer submissions are not required on destination networks for batches without `XMsg`s.
    - The logic to create a `XBlock` is deterministic for any finalized source rollup block height.

    <br />
    <details>
    <summary><code>XBlock</code> Storage and Calculation</summary>

    `XBlock` is not stored as they are deterministically calculated from a source network. So in effect, the source rollup stores them.
    Any component that depends on `XBlock`, calculates it from a source rollup.

    $XBlock = f(chain_A)$ where $f(x)$ is a deterministic `pure` function that takes a finalized network as input and produces `XBlock` as output.
    In practice, source blocks can be streamed and transformed using a simple translation function backed by an in-memory cache.

    </details>

    `XMsg`s are associated with an `XStream`. An `XStream`  is a logical connection between a source and destination network. It contains many `XMsg`s, each with a monotonically incrementing `XStreamOffset` (the offset is like a EOA nonce, it is incremented for each subsequent message sent from a source network to a destination network). `XMsg`s  are therefore uniquely identified and strictly ordered by their associated `XStream` and `XStreamOffset`.

    `XStreamOffset` allows for exactly-once delivery guarantees with strict ordering per source-destination network pair.

### 5. Validator Attestation

- Validators attest to `XBlock` hashes during CometBFT consensus. Each Omni consensus layer validator monitors every finalized block for all source networks in `halo`. By default, validators wait for block finalization, or some other agreed-upon finality mechanism, to ensure consistent and secure `XMsg` processing.

    All validators use halo to attest to `XBlock`s via CometBFT vote extensions. All validators in the CometBFT validator set should attest to all `XBlock`. An attestation is defined by the following `Attestation` type:

    ```go
    // Attestation containing quorum votes by the validator set of a cross-chain Block.
    type Attestation struct {
      BlockHeader                 // BlockHeader identifies the cross-chain Block
      ValidatorSetID  uint64      // Validator set that approved this attestation.
      AttestationRoot common.Hash // Attestation merkle root of the cross-chain Block
      Signatures      []SigTuple  // Validator signatures and public keys
    }
    ```

    Validators return an array of `Attestations` during the ABCI++ `ExtendVote` method.

### 6. Relayer Collects `XBlocks`

- Relayers monitoring the Omni consensus layer gather finalized `XBlock` hashes and their corresponding `XBlock` and `XMsg`s. Finalized rollup blocks without any `XMsg` events are ignored. The Relayer submits `XMsg`s to destination networks along with the set of validator signatures.

    The relayer determines how many `XMsg`s to package within each submission based on the “cost” of transactions submitted to the destination network. This is primarily defined by the data size and gas limit of the messages, the portal contract verification costs, and processing overhead.

    A merkle-multi-proof is generated for the set of identified `XMsg` that match the quorum `XBlock` attestations root.

    ```go
    // Msg is a cross-chain message.
    type Msg struct {
      MsgID                          // Unique ID of the message
      SourceMsgSender common.Address // Sender on source chain, set to msg.Sender
      DestAddress     common.Address // Target/To address to "call" on destination chain
      Data            []byte         // Data to provide to "call" on destination chain
      DestGasLimit    uint64         // Gas limit to use for "call" on destination chain
      TxHash          common.Hash    // Hash of the source chain transaction that emitted the message
    }
    ```

### 7. Relayer Submits `XBlocks`

- For every finalized `XBlock` hash, relayers construct an `XBlock` submission containing the `XBlock` and the validator signatures for the `XBlock` hashes. Merkle proofs are generated to prove the inclusion of each `XMsg` in the `XBlock` hash. This ensures the decoupling of attestations from execution, thereby allowing the relayer to split the execution of an `XBlock` into many transactions so that it adheres to the constraints of the destination rollup VM.

### 8. Relayer Submits `XMsg`

- The relayer delivers its submissions to the Portal contract on the destination rollup VM. Relayer submissions are defined in Go as:

    ```go
    // Submission is a cross-chain submission of a set of messages and their proofs.
    type Submission struct {
      AttestationRoot common.Hash // Attestation merkle root of the cross-chain Block
      ValidatorSetID  uint64      // Validator set that approved the attestation.
      BlockHeader     BlockHeader // BlockHeader identifies the cross-chain Block
      Msgs            []Msg       // Messages to be submitted
      Proof           [][32]byte  // Merkle multi proofs of the messages
      ProofFlags      []bool      // Flags indicating whether the proof is a left or right proof
      Signatures      []SigTuple  // Validator signatures and public keys
      DestChainID     uint64      // Destination chain ID, for internal use only
    }
    ```

### 9. Portal Verification and Forwarding

- The receiving network’s Portal contract verifies the `XBlock`’s validator signatures and `XMsg` merkle proofs before passing all verified `XMsg`s to their destination smart contracts. After verifying each submitted `XMsg`, the portal contract emits an `XReceipt` event. This marks the `XMsg` as “successful” or “reverted” by `halo`. `XMsg`s can revert if the gas limit was exceeded or if target address smart contract logic reverted for other reasons. `XReceipt`s are included in `XBlock`s (same as `XMsg`).

    ```go
    // Receipt is a cross-chain message receipt, the result of applying the Msg on the destination chain.
    type Receipt struct {
      MsgID                         // Unique ID of the cross chain message that was applied.
      GasUsed        uint64         // Gas used during message "call"
      Success        bool           // Result, true for success, false for revert
      RelayerAddress common.Address // Address of relayer that submitted the message
      TxHash         common.Hash    // Hash of the relayer submission transaction
    }
    ```

### 10. Smart Contract Execution

- The receiving smart contracts execute the logic for their users.
