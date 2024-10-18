---
slug: /
---

# Getting Started

Omni makes it easy to create chain-abstracted applications.


This guide will walk you through deploying and interacting with
**[xstake](https://github.com/omni-network/xstake) - a simple chain-abstracted staking
app**. It's an example application that demonstrates how to accept deposits on multiple chains, while maintaining global accounting on Omni.

In this guide you will:

- Run an Omni devnet
- Deploy and use a chain-abstracted application

#### Quick Overview

The app has two contracts:

- [`XStaker`](https://github.com/omni-network/xstake/blob/ad4cbb/src/XStaker.sol) - deployed on multiple chains, accepts ERC20 deposits.
- [`XStakeController`](https://github.com/omni-network/xstake/blob/ad4cbb/src/XStakeController.sol) - deployed on Omni, tracks stake across all chains.

This is all the context you need for this guide. For a full explanation of how the contracts work, see [Walkthrough](/build/walkthrough).

### 1. Requirements

You'll need the following installed:

- [docker](https://docs.docker.com/get-started/get-docker/)
- [foundry](https://book.getfoundry.sh/getting-started/installation)


And the Omni CLI:


```bash
# install
curl -sSfL https://raw.githubusercontent.com/omni-network/omni/main/scripts/install_omni_cli.sh | bash -s

# check installation
omni version
```


### 2. Setup


Create a new directory, and scaffold a new xstake project.

```bash
mkdir my-xstake
forge init --template https://github.com/omni-network/xstake
```


### 3. Run a devnet

Start a local devnet by running the following command:

```bash
omni devnet start

```

You should have a devnet running with three chains: `omni_evm`, `mock_op`, and `mock_arb`. Run `omni devnet info` to check.

```bash
omni devnet info

[
  {
    "chain_id": 1655,
    "chain_name": "mock_op",
    "portal_address": "0xb835dc695c6bfc8373c0d56973b5d9e9b083e97b",
    "rpc_url": "http://127.0.0.1:8001"
  },
  {
    "chain_id": 1656,
    "chain_name": "mock_arb",
    "portal_address": "0xb835dc695c6bfc8373c0d56973b5d9e9b083e97b",
    "rpc_url": "http://127.0.0.1:8002"
  },
  {
    "chain_id": 1651,
    "chain_name": "omni_evm",
    "portal_address": "0xb835dc695c6bfc8373c0d56973b5d9e9b083e97b",
    "rpc_url": "http://127.0.0.1:8000"
  }
]

```


### 4. Deploy contracts


```bash
make devnet-deploy
```

This outputs `deployments.sh` with addresses, RPCs and chain IDs - for your convenience :handshake:.

```bash
source deployments.sh

# rpcs
echo $OMNI_RPC      # Omni EVM RPC
echo $OP_RPC        # Mock OP RPC
echo $ARB_RPC       # Mock Arb RPC

# chain ids
echo $OMNI_CHAINID  # Omni EVM chain ID
echo $OP_CHAINID    # Mock OP chain ID
echo $ARB_CHAINID   # Mock Arb chain ID

# addresses
echo $OP_TOKEN      # ERC20 token address on Mock OP
echo $ARB_TOKEN     # ERC20 token address on Mock Arb
echo $OP_XSTAKER    # XStaker contract address on Mock OP
echo $ARB_XSTAKER   # XStaker contract address on Mock Arb
echo $CONTROLLER    # XStakeController contract address on Omni EVM
```

### 5. Try it out

Setup environment.
```bash
source deployments.sh

# prefunded dev account
DEV_ACCOUNT=0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266
DEV_PRIVKEY=0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80
```

Mint and approve tokens.
```bash
cast send $OP_TOKEN "mint(address,uint256)"     $DEV_ACCOUNT 100 --rpc-url $OP_RPC --private-key $DEV_PRIVKEY
cast send $OP_TOKEN "approve(address,uint256)"   $OP_XSTAKER 100 --rpc-url $OP_RPC --private-key $DEV_PRIVKEY
```

Calculate xcall fee.
```bash
cast call $OP_XSTAKER "stakeFee(uint256)(uint256)" 100 --rpc-url $OP_RPC
```


Stake!
```bash
# using --value from step above
cast send $OP_XSTAKER "stake(uint256)" 100 --rpc-url $OP_RPC --private-key $DEV_PRIVKEY --value 500000000
```

Verify stake:
```bash
cast call $CONTROLLER "stakeOn(address,uint64)(uint256)" $DEV_ACCOUNT $OP_CHAINID --rpc-url $OMNI_RPC
```

Note that we made our deposit on `mock_op`, but our stake is tracked on Omni. You can follow the same steps to deposit on `mock_arb`, substituting rpcs / addresses as needed.


### 6. 🎉 Done 🎉

That's it! You've deployed and used your first chain-abstacted application on Omni. Continue to [Walkthrough](/build/walkthrough) to learn more about how the contracts work.

<figure align="center">
    <img src="/img/cat.png" alt="gg wp" width="350" height="350" />
    <figcaption></figcaption>
</figure>
