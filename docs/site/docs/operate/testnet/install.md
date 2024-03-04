---
sidebar_position: 2
---

# Installing the Omni CLI

The Omni CLI will allow you to interact with the Omni network and simplifies the process of performing actions as an operator.

import Tabs from '@theme/Tabs';
import TabItem from '@theme/TabItem';

<Tabs>
  <TabItem value="binary" label="Binary" default>
    ```bash
    curl -sSfL https://raw.githubusercontent.com/omni-network/omni/cli/install.sh | sh -s
    ```
  </TabItem>
  <TabItem value="go" label="Go">
    ```bash
    go install github.com/omni-network/omni/cmd/omni@latest
    ```
  </TabItem>
  <TabItem value="source" label="Source">
    ```bash
    git clone https://github.com/omni-network/omni.git
    cd omni
    make build-cli
    ```
  </TabItem>
</Tabs>
