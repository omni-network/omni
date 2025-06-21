<h1 align="center">Omni Monorepo</h1>

<p align="center"><a href="https://docs.omni.network/"><img src="https://img.shields.io/badge/Docs-docs.omni.network-176CFF.svg"></a>
<a href="https://deepwiki.com/omni-network/omni"><img src="https://img.shields.io/badge/DeepWiki-View%20on%20DeepWiki-8A2BE2.svg"></a>
<a href="https://github.com/omni-network/omni/releases/"><img src="https://img.shields.io/github/release/omni-network/omni.svg?color=176CFF"></a>
<a href="https://goreportcard.com/report/github.com/omni-network/omni"><img src="https://goreportcard.com/badge/github.com/omni-network/omni"></a>
<a href="https://github.com/omni-network/omni/actions?query=workflow%3Aci-main"><img src="https://img.shields.io/github/actions/workflow/status/omni-network/omni/ci-main.yaml?label=Tests&logoColor=white" alt="Tests"></a>
<a href="https://x.com/OmniFDN"><img src="https://img.shields.io/twitter/follow/OmniFDN.svg?label=Follow"></a></p>

<div align="center"><img src="https://docs.omni.network/img/omni-banner.png" alt="Logo"></div>

## About Omni

This monorepo contains all source code for the Omni protocol. Omni's goal is to make it easy for smart contract developers to source liquidity and users from anywhere. The protocol consists of various components including an EVM and cross-chain messaging.

The [Omni Docs](https://docs.omni.network/) are the best place to get started learning about Omni. Our [DeepWiki page](https://deepwiki.com/omni-network/omni) is a great place to learn about the structure of this monorepo and Omni's architecture.

## Quickstart

Ensure [go](https://go.dev/doc/install), [goreleaser](https://goreleaser.com/install/) and [docker](https://docs.docker.com/engine/install/) are installed.

```bash
# Build local docker containers
make build-docker

# Run the end-to-end tests on a local devnet
MANIFEST=devnet1 make e2e-run

# Start a local devnet
make devnet-deploy

# Stop the local devnet
make devnet-clean
```
_If any of above commands fail, see the [troubleshooting section](./e2e/README.md#troubleshooting)._

## Directory Structure

<pre>
├── <a href="./cli/">cli</a>: Omni command line interface.
├── <a href="./contracts/">contracts</a>: Solidity contracts and related software.
├── <a href="./docs/">docs</a>: Resources and website for https://docs.omni.network.
├── <a href="./e2e/">e2e</a>: Deployments, testing, and live network management.
├── <a href="./halo/">halo</a>: Omni consensus client.
├── <a href="./lib/">lib</a>: Core libraries used across the Omni stack.
├── <a href="./monitor/">monitor</a>: Monitoring service for the network.
├── <a href="./octane/">octane</a>: Cosmos SDK module for the EVM.
├── <a href="./relayer/">relayer</a>: Relayer service for cross-chain messages and transactions.
├── <a href="./scripts/">scripts</a>: Utility scripts for development and operational tasks.
├── <a href="./sdk/">sdk</a>: SDK for building applications with Omni.
├── <a href="./solver/">solver</a>: Reference implementation for SolverNet.
</pre>

## Contributing

We are open to contributions, but don't currently have a formal process for contributors. If you are interested, browse through [open issues](https://github.com/omni-network/omni/issues) or reach out to chat with the team.

Follow these steps to set up a functional development environment:

1. Install Docker Desktop.
2. Setup commit signing:
  2a. [Create a PGP key pair](https://docs.github.com/en/authentication/managing-commit-signature-verification/generating-a-new-gpg-key)
  2b. [Add the public key to Github](https://docs.github.com/en/authentication/managing-commit-signature-verification/adding-a-gpg-key-to-your-github-account).
  2c. [Enabled commit signing](https://docs.github.com/en/authentication/managing-commit-signature-verification/signing-commits)
  2d. [Troubleshoot any issues](https://gist.github.com/paolocarrasco/18ca8fe6e63490ae1be23e84a7039374)
3. Run `make setup` to initialize your dev environment. See `Makefile` for details.

## Security

Omni has a bug bounty program via Immunefi. You can find more information [here](https://immunefi.com/bug-bounty/omni-network/information/).
