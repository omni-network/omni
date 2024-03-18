# Omni Contracts

## Overview

This directory contains the smart contracts and related tooling for Omni. It is structured to facilitate development, testing, and deployment of the contracts.

## Contents

- `bindings/`: Go bindings for contracts, including utility scripts for generation.
- `src/`: Solidity source files, organized by functionality (deploy, examples, interfaces, libraries, protocol).
- `script/`: Scripts for deploying contracts on various networks and utility tasks.
- `test/`: Test cases for the smart contracts, using Foundry.
- `Makefile`: Utility to automate common tasks such as building contracts, generating bindings, and running tests.
- `foundry.toml`, `package.json`, `pnpm-lock.yaml`, `tsconfig.json`: Configuration files for Foundry, Node.js, and TypeScript.

## Installation

To set up your environment for development:

1. Install Foundry for smart contract compilation and testing. Follow the instructions at [Foundry's GitHub repository](https://github.com/foundry-rs/foundry).
2. Install Node.js and pnpm to handle script execution and package management:
   - Node.js: [https://nodejs.org/](https://nodejs.org/)
   - pnpm: Run `npm install -g pnpm` after installing Node.js.
3. Ensure you have `abigen` installed for generating Go bindings. It should be version 1.13.14-stable. Install with:
   ```
   go install github.com/ethereum/go-ethereum/cmd/abigen@v1.13.14
   ```

## Build

To compile the smart contracts, run:

```
make build
```

This command compiles the Solidity contracts and also prints version information of the tools used.

## Running Tests

Execute the test suite with:

```
pnpm run test
```

### AVS Tests

#### Local

Or, for AVS tests locally with a running `anvil` node:

```
make avs-fork-test-local
```

Note: checks require setting an `INFURA_KEY` and `ETHERSCAN_GOERLI_KEY`, however these may be set to dummy values for local testing.

#### Goerli

Or, for AVS tests with on a goerli fork:

```
make avs-fork-test-goerli
```

## Generating Bindings

Generate Go bindings for the contracts with:

```
make bindings
```

This command requires `abigen` and will generate bindings, including examples.

## Deployments

Deploy contracts to networks such as Goerli with scripts included in the `script/` directory. Use the following commands for deployment and verification:

- Deploy to Goerli:
  ```
  make deploy-goerli-avs
  ```
- Verify contracts on Goerli Etherscan:
  ```
  make verify-goerli-avs
  ```
