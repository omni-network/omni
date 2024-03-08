---
sidebar_position: 5
---

# EVM Equivalence

The Omni architecture addresses two main challenges faced by existing blockchain frameworks: scalability and EVM compatibility. By moving the transaction mempool to the execution layer and employing an ABCI++ wrapper for state translation, Omni ensures that the consensus process remains lightweight and efficient, even under high network activity.

### Engine API and State Translation

The integration of the Ethereum Engine API and the use of an ABCI++ wrapper around the CometBFT engine facilitate seamless state translation between the Omni EVM and the CometBFT consensus layer. This architecture allows Omni to convert EVM blocks into single transactions that can be easily processed by CometBFT, significantly enhancing throughput and reducing consensus overhead.

```go
// ForkchoiceUpdatedV2 integration for state translation and consensus
forkchoiceResp, err := k.engineCl.ForkchoiceUpdatedV2(ctx, forkchoiceState, &payloadAttrs)
if err != nil {
    // Handle error...
} else if forkchoiceResp.PayloadStatus.Status != engine.VALID {
    // Handle invalid status...
}
```

The Engine API is employed to manage the fork choice and payload attributes, ensuring efficient consensus and state synchronization.

### Flexibility and Compatibility

Omni's design philosophy prioritizes flexibility and compatibility, supporting the use of any existing Ethereum execution client without the need for specialized modifications. This approach not only reduces the risk of introducing new bugs but also ensures ongoing compatibility with Ethereum's evolving ecosystem.

Omni's commitment to EVM equivalence extends to adopting recent upgrades such as dynamic transaction fees from EIP-1559 and future enhancements like proposer-builder separation. This ensures that Omni remains at the forefront of blockchain technology, offering a platform that is both cutting-edge and deeply integrated with the Ethereum ecosystem.
