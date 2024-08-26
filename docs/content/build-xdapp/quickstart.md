---
sidebar_position: 1
slug: /
---

# My First XDapp

This QuickStart guide will run through how to start an Omni XDApp (cross-chain decentralized app) in 5 minutes.

In this guide you will:

- Install the Omni CLI
- Scaffold a new project
- Run a local devnet, including Omni and multiple rollups
- Deploy contracts using foundry to the local devnet and test their functionality

## Steps

### Step 1: Install the Omni CLI

First, install the Omni CLI by running the following command:

```bash
curl -sSfL https://raw.githubusercontent.com/omni-network/omni/main/scripts/install_omni_cli.sh | sh -s
```

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

You can test the contracts with Forge by running:

```bash
forge test
```

or

```bash
make test
```

</details>

### Step 3: Run a local devnet

Start a local devnet by running the following command:

```bash
omni devnet start
```

Note: this requires [Docker](https://docs.docker.com/get-docker/) to be installed on your machine.

### Step 4: Deploy contracts

First, copy `.env.example` to `.env`. You shouldn't need to modify any parameters, but this stores some important info, so you should check it out.

```bash
cp .env.example .env
```

If you'd like to know where some of these values came from

```bash
# For devnet RPC data and portal addresses
omni devnet info

# For the private key that's prefunded locally:
anvil
````

Deploy the contracts to the local devnet using foundry:

```bash
make deploy
```

Note: we know the address the `GreetingBook` will be deployed to as a new network is started and the nonce for the account used to deploy is always the same (0). The `Greeter` contract is deployed to the same address on both the mock chains, since these are also new networks and the account has no actions on that account.

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
cast send 0x8464135c8F25Da09e49BC8782676a84730C318bC 'greet(string)' 'Yay in 5 minutes!' --private-key 0x59c6995e998f97a5a0044966f0945389dc9e86dae88c7a8412f4603b6b78690d --rpc-url http://localhost:8001 --value 1ether
```

### Step 6: Check the Greet

You can check the greet has been saved on the Omni EVM `GreetingBook` by running the following command:

```bash
cast call 0x8464135c8F25Da09e49BC8782676a84730C318bC "lastGreet():(address,string,uint64,uint256)" --rpc-url http://localhost:8000
```

### 🎉 Done 🎉

You have successfully deployed and interacted with an Omni XApp in less than five minutes!

<figure align="center">
    <img src="/img/cat.png" alt="gg wp" width="350" height="350" />
    <figcaption></figcaption>
</figure>
