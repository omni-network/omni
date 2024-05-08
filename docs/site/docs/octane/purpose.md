# Purpose and Benefits

- Rather than wrapping EVM transactions as Cosmos transactions (monolithic approach), EVM transactions are handled by the EVM, while consensus chain transactions are handled by the consensus module. This solves the bottleneck of forcing the Cosmos mempool to handle EVM load.
- This modular approach allows the EVM to scale independently from consensus, by simply adopting the latest performant execution client.
- Transaction throughput and gas throughput is limited only by the execution client, which can be highly optimized.
- Supports blocktimes ~1 second without optimizations.
- Is EVM equivalent: the code running the EVM is an unmodified EVM client, so it is exactly the same code that is running for Ethereum L1.
- Supports client diversity, as any EVM client that implements the EngineAPI can be used.
- Consensus chain transactions can be proxied through the EVM (via predeploys), if desired. Omni’s consensus chain implementation uses predeploys for all consensus chain transactions.
- Can support non-EVM execution clients if they implement the EngineAPI.
