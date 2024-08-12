---
sidebar_position: 4
---

import Tabs from '@theme/Tabs';
import TabItem from '@theme/TabItem';

# Updating

This guide will walk you through the process of updating your Omni CLI to the latest version.

## Update from Binary

Similar to installation, locate the old binary in your `$PATH` and replace it with the new binary. You can download the latest release from the [GitHub releases page](https://github.com/omni-network/omni/releases/).

## Update from Script

You can also update the Omni CLI using the installation script, which will overwrite the binary with the latest version:

```bash
curl -sSfL https://raw.githubusercontent.com/omni-network/omni/main/scripts/install_omni_cli.sh | sh -s
```

## Update from Source

If you have the Omni monorepo, you can update it by installing it from the `main` branch:

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
