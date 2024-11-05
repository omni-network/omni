<h1 align="center">Omni Monorepo</h1>

<p align="center">
    <strong>Monorepo for Omni node, contracts and other related tools.</strong>
</p>

<p align="center">
    <a href="https://github.com/omni-network/omni/actions?query=workflow%3Aci-release">
        <img src="https://img.shields.io/github/actions/workflow/status/omni-network/omni/ci-release.yaml?&label=Build&logo=github-actions&logoColor=white&style=flat" alt="Build">
    </a>
    <a href="https://github.com/omni-network/omni/actions?query=workflow%3Aci-main">
        <img src="https://img.shields.io/github/actions/workflow/status/omni-network/omni/ci-main.yaml?branch=main&label=Tests&logo=github-actions&logoColor=white&style=flat" alt="Tests">
    </a>
    <a href="https://docs.omni.network/">
        <img src="https://img.shields.io/badge/Docs-docs.omni.network-darkblue.svg?style=flat&logo=read-the-docs&logoColor=white" alt="Documentation">
    </a>
    <a href="https://discord.com/invite/bKNXmaX9VD">
        <img src="https://img.shields.io/discord/850218821398822922.svg?label=Discord&logo=discord&colorB=7289DA&style=flat" alt="Discord">
    </a>
    <a href="https://t.me/OmniFDN">
        <img src="https://img.shields.io/badge/Telegram-Join%20Chat-blue.svg?style=flat&logo=telegram" alt="Telegram">
    </a>
    <a href="https://x.com/OmniFDN">
        <img src="https://img.shields.io/twitter/follow/OmniFDN.svg?style=social&label=Follow" alt="Twitter Follow">
    </a>
</p>

<p align="center">
    <a href="https://omni.network/">
        <img src="static/omni_banner.png" alt="Logo">
    </a>
</p>

## About Omni

This monorepo contains all source code for the Omni protocol. Omni's goal is to make it easy for smart contract developers to source liquidity and users from anywhere. The protocol consists of various components including an EVM and cross-chain messaging.

The [Omni Docs](https://docs.omni.network/) are the best place to get started learning about Omni.

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
├── <a href="./contracts/">contracts</a>: Solidity contracts and related software.
│ ├── <a href="./contracts/core/">core/</a>: Core protocol smart contracts.
│ ├── <a href="./contracts/avs/">avs/</a>: Eigen AVS smart contracts.
│ ├── <a href="./contracts/bindings/">bindings/</a>: Go smart contract bindings.
│ └── <a href="./contracts/allocs/">allocs/</a>: Predeploy allocations.
├── <a href="./docs/">docs</a>: Documentation resources, including images and diagrams.
├── <a href="./halo/">halo</a>: The Halo instance, including application logic and attestation mechanisms.
│ ├── <a href="./halo/app/">app</a>: Application logic for Halo.
│ └── <a href="./halo/cmd/">cmd</a>: Command-line tools and utilities.
├── <a href="./lib/">lib</a>: Core libraries for various protocol functionalities.
│ ├── <a href="./lib/cchain/">cchain</a>: Consensus chain interaction utilities.
│ └── <a href="./lib/xchain/">xchain</a>: Cross-chain messaging and utilities.
├── <a href="./octane/">octane</a>: Octane is a modular framework for the EVM.
│ └── <a href="./octane/evmengine/">evmengine</a>: The EVMEngine cosmos module.
├── <a href="./relayer/">relayer</a>: Relayer service for cross-chain messages and transactions.
│ └── <a href="./relayer/app/">app</a>: Application logic for the relayer service.
├── <a href="./scripts/">scripts</a>: Utility scripts for development and operational tasks.
└── <a href="./e2e/test/">test</a>: Testing suite for end-to-end, smoke, and utility testing.
</pre>

## Contributing

We are open to contributions, but don't currently have a formal process for contributors. If you are interested, browse through [open issues](https://github.com/omni-network/omni/issues), read our [code of conduct](./CODE_OF_CONDUCT.md), or chat with the team.

Follow these steps to set up a functional development environment:

1. Install Docker Desktop.
2. [Create a PGP key pair](https://docs.github.com/en/authentication/managing-commit-signature-verification/generating-a-new-gpg-key) and [add the public key to Github](https://docs.github.com/en/authentication/managing-commit-signature-verification/adding-a-gpg-key-to-your-github-account).
3. Run `make setup` to initialize your dev environment. See `Makefile` for details.

## Security

Please refer to [SECURITY.md](./SECURITY.md).
