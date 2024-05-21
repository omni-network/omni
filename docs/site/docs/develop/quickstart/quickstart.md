---
sidebar_position: 1
---

# 3-Minute Cross-Chain dApp

This QuickStart guide will run through how to start an Omni cross chain dApp in less than three minutes.

In this guide you will:

- Install the Omni CLI, scaffold a new project and run a local devnet
- Deploy contracts using foundry to the local devnet and test their functionality

## Steps

### Step 1: Install the Omni CLI

First, install the Omni CLI by running the following command:

```bash
curl -sSfL https://raw.githubusercontent.com/omni-network/omni/main/scripts/install_omni_cli.sh | sh -s
```

You may otherwise install from source by following the steps shown in the [Omni CLI Tools section](../../tools/cli/cli.md).

### Step 2: Scaffold a new project

Next, create a new directory for your project and scaffold a new project using the Omni CLI:

```bash
mkdir my-omni-dapp
cd my-omni-dapp
omni developer new
```

Note: this requires [foundry](https://github.com/foundry-rs/foundry) to be installed on your machine.

<details>
<summary>Test the Contracts with Forge</summary>

You can test the contracts with Forge by running the following command:

```bash
forge test
```

</details>

### Step 3: Run a local devnet

Start a local devnet by running the following command:

```bash
omni devnet start
```

Note: this requires [Docker](https://docs.docker.com/get-docker/) to be installed on your machine.

### Step 4: Deploy contract

Deploy the contracts to the local devnet using foundry:

<details>
<summary>Obtaining Parameter Values</summary>

You can obtain RPC URL values and portal addresses for the running devnet chains by running the following command:

```bash
omni devnet info
```

And you the private key value is the first anvil private key, found by running:

```bash
anvil
```

</details>

```bash
forge script DeployXGreeter --broadcast --rpc-url http://localhost:8000 --private-key 0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80
forge script DeployXGreeter --broadcast --rpc-url http://localhost:8001 --private-key 0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80
forge script DeployXGreeter --broadcast --rpc-url http://localhost:8002 --private-key 0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80
```

### Step 5: Perform an XGreet

You can now perform an XGreet by running the following command:

<details>
<summary>Obtaining Deployment Addresses</summary>

You can obtain the XGreeter deployment addresses from the output of the previous forge script deployment.

Because the devnet has just been started, the addresses will be the same as the ones shown below:

```bash
omni_evm: 0x0DCd1Bf9A1b36cE34237eEaFef220932846BCD82
mock_op: 0xa513E6E4b8f2a923D98304ec87F64353C4D5C853
mock_arb: 0xa513E6E4b8f2a923D98304ec87F64353C4D5C853
```

</details>

```bash
cast send 0x0DCd1Bf9A1b36cE34237eEaFef220932846BCD82 'xgreet(uint64,address,string)' 1655 0xa513E6E4b8f2a923D98304ec87F64353C4D5C853 '3 minutes go by too fast!' --value 0.1ether --private-key 0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80 --rpc-url http://localhost:8000
```

### 🎉 Done 🎉

You have successfully deployed and interacted with an Omni cross-chain dApp in less than three minutes!
