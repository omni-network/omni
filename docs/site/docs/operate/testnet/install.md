---
sidebar_position: 2
---

# Installing the Omni CLI

The Omni CLI will allow you to interact with the Omni network and simplifies the process of performing actions as an operator.

import Tabs from '@theme/Tabs';
import TabItem from '@theme/TabItem';

<Tabs>
  <TabItem value="source" label="Source with `make`">
    ```bash
    git clone https://github.com/omni-network/omni.git
    cd omni
    make install-cli
    ```
  </TabItem>
  <TabItem value="go" label="Source with `go`">
    ```bash
    git clone https://github.com/omni-network/omni.git
    cd omni
    go install ./cli/cmd/omni
    ```
  </TabItem>
</Tabs>
