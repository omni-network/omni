---
sidebar_position: 1
---

# Omega Testnet

This page will walk through how to scaffold, test, deploy and interact with a Solidity program added to the Omni Omega testnet.

## Scaffolding a New Project

You can scaffold a new [Foundry](https://github.com/foundry-rs/foundry) template project with the [Omni CLI](./install.md) or with `forge`.

To scaffold a new project, run the following from your new project directory:

import Tabs from '@theme/Tabs';
import TabItem from '@theme/TabItem';

<Tabs>
  <TabItem value="cli" label="Omni CLI">
    ```bash
    omni developer new
    ```
  </TabItem>
  <TabItem value="forge" label="Forge">
    ```bash
    forge init --template https://github.com/omni-network/omni-forge-template.git
    ```
  </TabItem>
</Tabs>

## Testing Your Project

You can test your project by running:

```forge
forge test
```

## Deploying Your Project

You can deploy your project to the Omni Omega testnet by

<!-- TODO: include how to scaffold, test, deploy, interact -->
