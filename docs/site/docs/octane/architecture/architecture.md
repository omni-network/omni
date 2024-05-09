---
sidebar_position: 1
---

# Architecture

You can read more about the architecture in our documentation portal – check out [this page](https://docs.omni.network/protocol/evmengine/dual) for an overview, and [this one](https://docs.omni.network/protocol/evmengine/lifecycle) for details on the EVM payload lifecycle. Continue reading below for a high level overview.

## Overview

Octane is built on a sophisticated modern architecture that includes the following components:

### Ethereum Engine API

Ethereum introduced the [Engine API](https://github.com/ethereum/execution-apis/tree/main/src/engine) to decouple the execution layer from the consensus layer. This is almost like a `consensus.EngineV2` interface, except that this is an HTTP RPC interface and therefore introduces more decoupling.

Here is a [visual guide](https://hackmd.io/@danielrachi/engine_api) and this is the original [design space doc](https://hackmd.io/@n0ble/consensus_api_design_space).

The benefit to using the Engine API for decoupling execution vs consensus is that multiple execution layer implementations are available with new ones being added constantly. This means that the latest and greatest EVM implementation is always available for use (standing on the shoulders of giants).

Noteworthy implementations:

- [geth](https://github.com/ethereum/go-ethereum): the original and most-used go implementation.
- [erigon](https://github.com/ledgerwatch/erigon): GA from [2022](https://erigon.substack.com/p/post-merge-release-of-erigon-dropping); performant go implementation, almost like geth v2.
- [reth](https://github.com/paradigmxyz/reth): soon to be available rust implementation, almost like Erigon v2.

Octane drives its execution layer via the EngineAPI for the block building engine.
