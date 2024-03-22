---
sidebar_position: 2
---

# Portal Send Logic

Portal Contracts are an integral part of the Omni protocol's architecture, deployed across various Rollup EVMs and the Omni EVM itself. These contracts are specifically designed for initiating cross-chain communications, acting as the gateway for emitting cross-chain messages known as `XMsg`. A notable feature is the "pay at source" fee mechanism, leveraging the native token of the source chain for transaction fees. Moreover, Portal Contracts maintain a record of the omni consensus validator set, essential for the verification of cross-chain message attestations.

## Cross-Chain Calls

To initiate a cross-chain call, the Portal Contract provides the `xcall` method. This function is accessible via the `omni.xcall()` helper, which simplifies the process of making cross-chain requests. Upon execution, an `XMsg` Event is emitted, signifying the successful forwarding of the cross-chain message. The `xcall` method is designed to facilitate seamless communication between chains, underpinning the broader objective of interoperability within the Omni protocol ecosystem.

For detailed instructions on conducting cross-chain transactions, refer to the [developer section](../../../develop/introduction.md).
