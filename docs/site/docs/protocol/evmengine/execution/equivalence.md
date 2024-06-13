---
sidebar_position: 1
---

# Complete EVM Equivalence

Omni offers complete EVM equivalence, ensuring that developers can deploy their Ethereum smart contracts to the Omni network without any modifications. Unlike other platforms that claim "EVM compatibility," Omni provides a truly equivalent environment by running an unmodified version of the Ethereum Virtual Machine (EVM). This fidelity guarantees that opcode compatibility issues are non-existent, and all developer tooling designed for Ethereum seamlessly works with the Omni EVM.

### Omni Validators Run Unmodified EVM Code

At the heart of Omni's EVM equivalence is the fact that Omni validators execute the same EVM code as Ethereum validators. This direct execution model means that smart contracts operate under the same conditions as they would on Ethereum, offering an unparalleled level of consistency and reliability for developers.

### Zero Changes Needed for Smart Contracts

Developers can migrate their dApps to Omni without worrying about the compatibility or operational differences often encountered on other EVM-compatible chains. The Omni EVM's adherence to Ethereum's standards ensures that smart contracts execute as intended, without requiring any adjustments or special considerations.

### Full Support for All EVM Opcodes and Upgrades

Omni's commitment to maintaining an unmodified EVM client extends to full support for all EVM opcodes and the inclusion of the latest EVM upgrades. This dedication ensures that Omni remains in lockstep with Ethereum's evolution, providing developers with a stable and feature-rich environment for their applications.

Below, the [`pushPayload`](https://github.com/omni-network/omni/blob/0f09c724ac941afc45c5f7eb1ed1a773f51dac81/halo/evmengine/keeper/msg_server.go#L116) function showcases the execution of EVM transactions, which implement strictly the same behavior of Ethereum and ensures a frictionless experience for developers.

```go
// pushPayload creates a new payload from the given message and pushes it to the execution client.
// It returns the new forkchoice state.
func pushPayload(ctx context.Context, engineCl ethclient.EngineClient, msg *types.MsgExecutionPayload,
) (engine.ExecutableData, error) {
    // Transaction execution logic, forwarding payload to EVM client...
}
```

## Advantages of Omni's EVM Equivalence

- **Seamless Migration:** Developers can port their DApps to Omni without any modifications, significantly reducing the effort and complexity involved in accessing a new blockchain ecosystem.
- **Developer Tooling Compatibility:** All the tools, libraries, and frameworks designed for Ethereum development are fully compatible with the Omni EVM, streamlining the development process.
- **Future-Proof:** Omni's alignment with Ethereum's upgrade path ensures that developers can leverage the latest features and improvements without delay.
