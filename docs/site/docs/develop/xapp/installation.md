---
sidebar_position: 2
---

# Installation

`XApp` is a contract maintained in Omni's [monorepo](https://github.com/omni-network/omni/tree/main/contracts), alongside the rest of Omni's smart contracts. To use it, install Omni's smart contracts in your project.

:::info Note

Note that Omni's contracts are under active development, and are subject to change.

:::

import Tabs from '@theme/Tabs';
import TabItem from '@theme/TabItem';

### Installing the XApp contract

<Tabs groupId="sol-manager">
  <TabItem value="npm" label="npm" default>
    ```bash
    npm install @omni-network/contracts
    ```
  </TabItem>
  <TabItem value="forge" label="forge">
    ```bash
    forge install omni-network/omni
    ```
  </TabItem>
</Tabs>

### Importing the XApp contract

<Tabs groupId="sol-manager">
  <TabItem value="npm" label="npm" default>
    ```solidity
    import { XApp } from "@omni-network/contracts/src/pkg/XApp.sol"
    ```
  </TabItem>
  <TabItem value="forge" label="forge">
    ```solidity
    import { XApp } from "omni/contracts/src/pkg/XApp.sol"
    ```
  </TabItem>
</Tabs>
