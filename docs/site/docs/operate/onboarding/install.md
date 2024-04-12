---
sidebar_position: 2
---

# Installing the Omni CLI

The Omni CLI will allow you to interact with the Omni network and simplifies the process of performing actions.

## Install from Binary

The easiest way to install the Omni CLI is to download the latest release from the [GitHub releases page](https://github.com/omni-network/omni/releases). Once downloaded, you can extract the binary and move it to a location in your $PATH.

## Install from Script

You can also install the Omni CLI using the following script:

```bash
curl -sSfL https://raw.githubusercontent.com/omni-network/omni/main/scripts/install_omni_cli.sh | sh -s
```

## Install from Source

import Tabs from '@theme/Tabs';
import TabItem from '@theme/TabItem';

<Tabs>
  <TabItem value="source" label="src with make">
    ```bash
    git clone https://github.com/omni-network/omni.git
    cd omni
    make install-cli
    ```
  </TabItem>
  <TabItem value="go" label="src with go">
    ```bash
    git clone https://github.com/omni-network/omni.git
    cd omni
    go install ./cli/cmd/omni
    ```
  </TabItem>
</Tabs>
