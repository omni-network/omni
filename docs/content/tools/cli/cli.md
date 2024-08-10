---
sidebar_position: 1
---

import Tabs from '@theme/Tabs';
import TabItem from '@theme/TabItem';

# Installation

The Omni CLI is a command-line interface for managing your Omni projects as a developer and interacting with the network as an operator. This guide will walk you through the installation process.

## Install from Binary

The easiest way to install the Omni CLI is to download the latest release from the [GitHub releases page](https://github.com/omni-network/omni/releases). Once downloaded, you can extract the binary and move it to a location in your `$PATH`.

## Install from Script

You can also install the Omni CLI using the following script:

```bash
curl -sSfL https://raw.githubusercontent.com/omni-network/omni/main/scripts/install_omni_cli.sh | sh -s
```

## Install from Source

<Tabs>
  <TabItem value="source" label="src with make">
    ```bash
    git clone https://github.com/omni-network/omni.git
    cd omni
    latest_tag=$(curl -s https://api.github.com/repos/omni-network/omni/releases/latest | jq -r .tag_name)
    git checkout $latest_tag
    make install-cli
    ```
  </TabItem>
  <TabItem value="go" label="src with go">
    ```bash
    git clone https://github.com/omni-network/omni.git
    cd omni
    latest_tag=$(curl -s https://api.github.com/repos/omni-network/omni/releases/latest | jq -r .tag_name)
    git checkout $latest_tag
    go install ./cli/cmd/omni
    ```
  </TabItem>
</Tabs>
