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

### Step 4: Deploy contracts

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

These values are kept in `./script/bash/.env.example` and are used to deploy the contracts. You can rename the file to `.env` and fill in the values for other networks.

</details>

```bash
bash script/bash/deploy.sh
```

### Step 5: Perform a Cross-Chain Greet

You can now perform a cross-rollup greet by running the following command:

<details>
<summary>Obtaining Deployment Addresses</summary>

You can obtain the XGreeter deployment addresses from the output of the previous forge script deployment.

Because the devnet has just been started, the addresses will be the same as the ones shown below:

```bash
omni_evm: 0x8464135c8F25Da09e49BC8782676a84730C318bC
mock_op: 0x8464135c8F25Da09e49BC8782676a84730C318bC
mock_arb: 0x8464135c8F25Da09e49BC8782676a84730C318bC
```

</details>

```bash
cast send 0x8464135c8F25Da09e49BC8782676a84730C318bC 'greet(string)' 'Yay in 3 minutes!' --private-key 0x59c6995e998f97a5a0044966f0945389dc9e86dae88c7a8412f4603b6b78690d --rpc-url http://localhost:8001 --value 1ether
```

### Step 6: Check the Greet

You can check the greet has been saved on the Omni EVM global state by running the following command:

```bash
cast call 0x8464135c8F25Da09e49BC8782676a84730C318bC "lastGreet():(uint64,uint256,uint256,address,address,string)" --rpc-url http://localhost:8000
```

### 🎉 Done 🎉

You have successfully deployed and interacted with an Omni cross-chain dApp in less than three minutes!

<figure align="center">
    <img src="/img/cat.png" alt="gg wp" width="350" height="350" />
    <figcaption></figcaption>
</figure>
