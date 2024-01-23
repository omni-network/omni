# End-to-End Tests

Spins up and tests Omni devnets in Docker Compose based on a testnet manifest. To run the CI testnet:

```sh
# In repo root
# Install the e2e runner
go install github.com/omni-network/omin/test/e2e/runner

# Build docker image of the code to test.
make build-docker

# Run one of the "manifests" in networks/ directory: runner -f <manifest>
runner -f test/e2e/networks/single.toml
```

This creates and runs a testnet named `single` under `test/e2e/runs/single/`.

## Conceptual Overview

Please refer to the [cometBFT E2E test framework](https://github.com/cometbft/cometbft/tree/main/test/e2e) for more details.

## Test Stages

The test runner has the following stages, which can also be executed explicitly by running `runner -f <manifest> <stage>`:

* `setup`: generates configuration files.

* `start`: starts Docker containers.

* `wait`: waits for a few blocks to be produced, and for all nodes to catch up to it.

* `stop`: stops Docker containers.

* `cleanup`: removes configuration files and Docker containers/networks.

Auxiliary commands:

* `logs`: outputs all node logs.

* `tail`: tails (follows) node logs until canceled.
