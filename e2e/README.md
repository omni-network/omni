# End-to-End Tests

Spins up and tests Omni devnets in Docker Compose based on a testnet manifest. To run the CI testnet:

```sh
# In repo root
# Install the e2e app
go install github.com/omni-network/omni/e2e

# Build docker image of the code to test.
make build-docker

# Run one of the "manifests" in manifests/ directory: e2e -f <manifest>
e2e -f e2e/manifests/devnet0.toml
```

This creates and runs a testnet named `single` under `e2e/runs/single/`.

## Conceptual Overview

Please refer to the [cometBFT E2E test framework](https://github.com/cometbft/cometbft/tree/main/test/e2e) for more details.

In order to perform any action on a network (deploy/test/show logs), the following process is followed to create a network `Definition`:
1. A network is initially declared in a `manifest` file, see [manifests/](./manifests) folder. It defines the desired network topology. See the `e2e/types#Manifest` type for details.
2. Then the infrastructure provider (only `docker compose` supported at the moment) subsequently generates the `e2e/types#InfrastructureData` from the manifest. This defines the instance IPs and ports of everything we will deploy.
3. Subsequently, we generate a `Testnet` struct which is basically contains all the configuration/keys/peers/images/files/folders required to deploy a network. See `e2e/types#Testnet` for details.
4. We then instantiate a `netman.Manager` which is responsible for deploying portals. It takes a `Testnet` struct as input.
5. Finally, we instantiate new `InfrastructureProvider` which can deploy the network. It takes a `Testnet` struct and `InfrastructureData` as input.

These objects are then wrapped in a `e2e/app#Definition` that can be used to perform any action on a network.

## Test Stages

The e2e test has the following stages, which can also be executed explicitly by running `e2e -f <manifest> <stage>`:

* `setup`: generates configuration files.

* `start`: starts Docker containers.

* `wait`: waits for a few blocks to be produced, and for all nodes to catch up to it.

* `stop`: stops Docker containers.

* `cleanup`: removes configuration files and Docker containers/networks.

Auxiliary commands:

* `logs`: outputs all node logs.

* `tail`: tails (follows) node logs until canceled.


## Troubleshooting
**MacBook E2E test fails to start docker container**
If you are experiencing an issue running the e2e tests and the error output looks like this:
```
Error response from daemon: no match for platform in manifest
```
This error is known to happen on a freshly installed MacBook, we are investigating the underlying issue, but meanwhile, there is a workaround:
1. Clean all the build images `docker system prune -a -f --volumes`
2. Set the platform env: `export DOCKER_DEFAULT_PLATFORM=linux/amd64`
3. Rerun docker build: `make build-docker`
4. Run tests again: `make e2e-ci` or any other tests using the e2e command.

Please let the team know if you experienced the above issue.

**Failure to start SVM container**

If the e2e tests fail to run with an error such as the following:

```sh
ERRO !! Fatal error occurred, app died !!     err="svm init: request airdrop for role account: rpc call requestAirdrop() on http://localhost:8899: Post \"http://localhost:8899\": dial tcp [::1]:8899: connect: connection refused"
```

Check the SVM container logs for an error such as: `Incompatible CPU detected: missing AVX support`.
This issue typically happens when using an unsupported virtual machine. To solve it:

1. Go to Docker Desktop's settings > General > Virtual Machine Options
2. Select the Docker VMM option and restart Docker
3. Rerun docker build: `make build-docker`
