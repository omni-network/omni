# Portals

Portal contracts are an integral part of Omni and are deployed to every supported rollup network EVM and the Omni EVM itself. These contracts are responsible for initiating `XMsg` requests and maintaining a record of the Omni validator set to verify inbound `XMsg` deliveries.

## Portal Send Logic

 These contracts are specifically designed for initiating cross-chain communications, acting as the gateway for emitting cross-chain messages known as `XMsg`. A notable feature is the "pay at source" fee mechanism, leveraging the native token of the source chain for transaction fees. Moreover, Portal Contracts maintain a record of the omni consensus validator set, essential for the verification of cross-chain message attestations.

### Cross-Chain Calls

To initiate a cross-chain call, the Portal Contract provides the `xcall` method. This function is accessible via the `omni.xcall()` helper, which simplifies the process of making cross-chain requests. Upon execution, an `XMsg` Event is emitted, signifying the successful forwarding of the cross-chain message. The `xcall` method is designed to facilitate seamless communication between chains, underpinning the broader objective of interoperability within the Omni protocol ecosystem.

Below is a summarized fragment for the underlying logic beneath `xcall` from the [Portal contract Solidity source](https://github.com/omni-network/omni/blob/main/contracts/core/src/xchain/OmniPortal.sol):


```solidity
    /**
     * @notice Call a contract on another chain.
     * @param destChainId   Destination chain ID
     * @param conf          Confirmation level
     * @param to            Address of contract to call on destination chain
     * @param data          ABI Encoded function calldata
     * @param gasLimit      Execution gas limit, enforced on destination chain
     */
    function xcall(uint64 destChainId, uint8 conf, address to, bytes calldata data, uint64 gasLimit)
        external
        payable
        whenNotPaused(ActionXCall, destChainId)
    {
        require(isSupportedDest[destChainId], "OmniPortal: unsupported dest");
        require(to != VirtualPortalAddress, "OmniPortal: no portal xcall");
        require(gasLimit <= xmsgMaxGasLimit, "OmniPortal: gasLimit too high");
        require(gasLimit >= xmsgMinGasLimit, "OmniPortal: gasLimit too low");
        require(data.length <= xmsgMaxDataSize, "OmniPortal: data too large");

        // conf level will always be first byte of shardId. for now, shardId is just conf level
        uint64 shardId = uint64(conf);
        require(isSupportedShard[shardId], "OmniPortal: unsupported shard");

        uint256 fee = feeFor(destChainId, data, gasLimit);
        require(msg.value >= fee, "OmniPortal: insufficient fee");

        outXMsgOffset[destChainId][shardId] += 1;

        emit XMsg(destChainId, shardId, outXMsgOffset[destChainId][shardId], msg.sender, to, data, gasLimit, fee);
    }
```

## Portal Receive Logic

### `xsubmit` Method

The `xsubmit` method is specifically designed for receiving calls from another chain. It is crucial for the secure and validated processing of cross-chain messages.

Below is a summarized fragment for `xsubmit` from the [Portal contract Solidity source](https://github.com/omni-network/omni/blob/main/contracts/core/src/xchain/OmniPortal.sol):

```solidity
    /**
     * @notice Submit a batch of XMsgs to be executed on this chain
     * @param xsub  An xchain submission, including an attestation root w/ validator signatures,
     *              and a block header and message batch, proven against the attestation root.
     */
    function xsubmit(XTypes.Submission calldata xsub)
        external
        whenNotPaused(ActionXSubmit, xsub.blockHeader.sourceChainId)
        nonReentrant
    {
        XTypes.Msg[] calldata xmsgs = xsub.msgs;
        XTypes.BlockHeader calldata xheader = xsub.blockHeader;
        uint64 valSetId = xsub.validatorSetId;

        // Validation logic
        ...

        // execute xmsgs
        for (uint256 i = 0; i < xmsgs.length; i++) {
            _exec(xheader, xmsgs[i]);
        }
    }
```

### Submission Validation

To ensure the integrity and authenticity of incoming cross-chain messages, the following validation processes are performed:

```solidity

        require(xheader.consensusChainId == omniCChainId, "OmniPortal: wrong cchain ID");
        require(xmsgs.length > 0, "OmniPortal: no xmsgs");
        require(valSetTotalPower[valSetId] > 0, "OmniPortal: unknown val set");
        require(valSetId >= _minValSet(), "OmniPortal: old val set");

        // check that the attestationRoot is signed by a quorum of validators in xsub.validatorsSetId
        require(
            Quorum.verify(
                xsub.attestationRoot,
                xsub.signatures,
                valSet[valSetId],
                valSetTotalPower[valSetId],
                XSubQuorumNumerator,
                XSubQuorumDenominator
            ),
            "OmniPortal: no quorum"
        );

        // check that blockHeader and xmsgs are included in attestationRoot
        require(
            XBlockMerkleProof.verify(xsub.attestationRoot, xheader, xmsgs, xsub.proof, xsub.proofFlags),
            "OmniPortal: invalid proof"
        );

```

Through these meticulous validation, verification, and execution steps, the system guarantees the secure handling of cross-chain messages, facilitating reliable communication across different chains.
