---
sidebar_position: 2
---

# Client Diversity

Omni's EVM Engine champions the principle of client diversity, allowing the integration and use of any Ethereum Virtual Machine (EVM) client within its ecosystem. This approach underscores Omni's commitment to flexibility, scalability, and interoperability, promoting a robust and resilient blockchain network.

### Integration of EVM Clients

Omni adheres to the Engine API, a standard that all EVM clients also comply with. This adherence ensures that any EVM client, such as Geth, Besu, Erigon, and others, can be seamlessly integrated into the Omni network without the need for specialized modifications. This approach allows the Omni ecosystem to leverage the unique features and optimizations that different clients provide.

<details>
<summary>Engine API Integration</summary>

The Engine API is a set of authenticated JSON-RPC endpoints that extend the normal JSON-RPC interface. It includes methods for creating new EVM blocks, updating the fork choice, and retrieving cached payloads. The Engine API is designed to be compatible with any EVM client, allowing for seamless integration of the EVM into the Omni network.

The [Engine API Client interface](https://github.com/omni-network/omni/blob/0f09c724ac941afc45c5f7eb1ed1a773f51dac81/lib/ethclient/engineclient.go#L31) is defined as follows:

```go
// EngineClient defines the Engine API authenticated JSON-RPC endpoints.
// It extends the normal Client interface with the Engine API.
type EngineClient interface {
	Client

	// NewPayloadV2 creates an Eth1 block, inserts it in the chain, and returns the status of the chain.
	NewPayloadV2(ctx context.Context, params engine.ExecutableData) (engine.PayloadStatusV1, error)
	// NewPayloadV3 creates an Eth1 block, inserts it in the chain, and returns the status of the chain.
	NewPayloadV3(ctx context.Context, params engine.ExecutableData, versionedHashes []common.Hash,
		beaconRoot *common.Hash) (engine.PayloadStatusV1, error)

	// ForkchoiceUpdatedV2 has several responsibilities:
	//  - It sets the chain the head.
	//  - And/or it sets the chain's finalized block hash.
	//  - And/or it starts assembling (async) a block with the payload attributes.
	ForkchoiceUpdatedV2(ctx context.Context, update engine.ForkchoiceStateV1,
		payloadAttributes *engine.PayloadAttributes) (engine.ForkChoiceResponse, error)

	// ForkchoiceUpdatedV3 is equivalent to V2 with the addition of parent beacon block root in the payload attributes.
	ForkchoiceUpdatedV3(ctx context.Context, update engine.ForkchoiceStateV1,
		payloadAttributes *engine.PayloadAttributes) (engine.ForkChoiceResponse, error)

	// GetPayloadV2 returns a cached payload by id.
	GetPayloadV2(ctx context.Context, payloadID engine.PayloadID) (*engine.ExecutionPayloadEnvelope, error)
	// GetPayloadV3 returns a cached payload by id.
	GetPayloadV3(ctx context.Context, payloadID engine.PayloadID) (*engine.ExecutionPayloadEnvelope, error)
}
```

</details>

### Advantages of Client Diversity

- **Innovation and Improvement:** By supporting a variety of EVM clients, Omni encourages innovation and continuous improvement within the ecosystem.
- **Security and Robustness:** Diversity in execution clients can enhance network security, mitigating the risk of widespread issues stemming from client-specific vulnerabilities.
- **Customization and Optimization:** Developers and node operators have the freedom to choose an EVM client that best fits their needs, optimizing for performance, tooling, or other criteria.
