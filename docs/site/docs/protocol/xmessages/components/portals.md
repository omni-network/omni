---
sidebar_position: 2
---

import GitHubCodeBlock from '@site/src/components/GitHubCodeBlock/GitHubCodeBlock';

# Portals

Portal contracts are an integral part of Omni and are deployed to every supported rollup network EVM and the Omni EVM itself. These contracts are responsible for initiating `XMsg` requests and maintaining a record of the Omni validator set to verify inbound `XMsg` deliveries.

## Portal Send Logic

 These contracts are specifically designed for initiating cross-chain communications, acting as the gateway for emitting cross-chain messages known as `XMsg`. A notable feature is the "pay at source" fee mechanism, leveraging the native token of the source chain for transaction fees. Moreover, Portal Contracts maintain a record of the omni consensus validator set, essential for the verification of cross-chain message attestations.

### Cross-Chain Calls

To initiate a cross-chain call, the Portal Contract provides the `xcall` method. This function is accessible via the `omni.xcall()` helper, which simplifies the process of making cross-chain requests. Upon execution, an `XMsg` Event is emitted, signifying the successful forwarding of the cross-chain message. The `xcall` method is designed to facilitate seamless communication between chains, underpinning the broader objective of interoperability within the Omni protocol ecosystem.

Below is a summarized fragment for the underlying logic beneath `xcall` from the [Portal contract Solidity source](https://github.com/omni-network/omni/blob/1439d8a99f66a3bb3b7d113c63f8f073512c5377/contracts/src/protocol/OmniPortal.sol):

<GitHubCodeBlock url="https://github.com/omni-network/omni/blob/3064e8ad22557c8756547f1875d026e348846c9d/contracts/src/xchain/OmniPortal.sol#L100-L129" />

For detailed instructions on conducting cross-chain transactions, refer to the [developer section](../../../develop/introduction.md).

## Portal Receive Logic

### `xsubmit` Method

The `xsubmit` method is specifically designed for receiving calls from another chain. It is crucial for the secure and validated processing of cross-chain messages:

- Accepts calls only from valid signatures, primarily from the designated Relayer.
- Requires the destination contract `address` and `data`, which includes the `method` and method parameters (`data`).

Below is a summarized fragment for `xsubmit` from the [Portal contract Solidity source](https://github.com/omni-network/omni/blob/1439d8a99f66a3bb3b7d113c63f8f073512c5377/contracts/src/protocol/OmniPortal.sol):

```solidity
    /**
     * @notice Submit a batch of XMsgs to be executed on this chain
     * @param xsub  An xchain submisison, including an attestation root w/ validator signatures,
     *              and a block header and message batch, proven against the attestation root.
     */
    function xsubmit(XTypes.Submission calldata xsub) external {
        // Submission validations that check non-empty xmsgs, validator set validity and power,
        // matching valSetId sequence, quorum signature, and proof of blockHeader and xmsgs in attestationRoot.

        // source chain block height of this submission
        uint64 blockHeight = xsub.blockHeader.blockHeight;

        // last seen block height for this source chain
        uint64 lastBlockHeight = inXStreamBlockHeight[xsub.blockHeader.sourceChainId];

        // update in stream block height, if it's new
        if (blockHeight > lastBlockHeight) inXStreamBlockHeight[xsub.blockHeader.sourceChainId] = blockHeight;

        // update in stream validator set id, if it's new
        if (valSetId > lastValSetId) inXStreamValidatorSetId[xsub.blockHeader.sourceChainId] = valSetId;

        // execute xmsgs
        for (uint256 i = 0; i < xsub.msgs.length; i++) {
            _exec(xsub.msgs[i]); // performs the contract call
        }
    }
```

### Submission Validation

To ensure the integrity and authenticity of incoming cross-chain messages, the following validation processes are performed:

<details>
<summary>Submission Validation Code</summary>

Below is a summarized fragment for the validations in `xsubmit` from the [Portal contract Solidity source](https://github.com/omni-network/omni/blob/main/contracts/src/protocol/OmniPortal.sol):

<GitHubCodeBlock url="https://github.com/omni-network/omni/blob/3064e8ad22557c8756547f1875d026e348846c9d/contracts/src/xchain/OmniPortal.sol#L146-L183" />

</details>

#### Aggregate Attestation Data Validation

Checks include:

- **Existence of Messages:** Ensuring the message batch (`XMsgs`) is not empty.
- **Validator Set Verification:** Verification that the `ValidatorSetID` is known and the validator set has non-zero power, ensuring the availability of the corresponding validator set.
- **Validator Set Sequence:** Ensuring the submission's validator set is either the same as the last seen or the next sequential one, which helps maintain the order and prevents replay of old data.

#### XMsgBatch Data Validation

This step involves:

- **Chain ID Confirmation:** Confirming that the `TargetChainID` matches the local chain ID, ensuring that the messages are destined for the correct chain.
- **Message Index Verification:** For a new batch, verifying that `WrappedXMsg` indexes start at 0 or follow the last processed message index for ongoing submissions, maintaining the order and integrity of messages.

#### Aggregate Attestation Signatures Verification

This process includes:

- **Quorum Authentication:** Authenticating that the attestation root is signed by a quorum of validators from the validator set, ensuring that a majority agrees on the submitted batch.
- **Proof Verification:** Verifying that the block header and messages (`XMsgs`) are included in the `AttestationRoot` through a valid Merkle-multi-proof. This proves the integrity and inclusion of the message batch and block header related to the submission.

#### Update and Execution

Post-validation steps include:

- **Block Height and Validator Set Updates:** Updating the last seen block height and validator set ID for the source chain if the new data is more recent, ensuring the system remains up-to-date.
- **Message Execution:** Executing each message in the batch upon successful validation, applying the cross-chain data.

Through these meticulous validation, verification, and execution steps, the system guarantees the secure handling of cross-chain messages, facilitating reliable communication across different chains.
