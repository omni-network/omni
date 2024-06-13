# Octane

Octane is a modular framework for the EVM. Developers can use it to run:

- any EVM execution client that implements the [Engine API](https://hackmd.io/@danielrachi/engine_api)
- any consensus client that implements [ABCI 2.0](https://github.com/cometbft/cometbft/tree/main/spec/abci)

You can read more about Octane in our [documentation](https://docs.omni.network/octane/background/introduction)

## Examples

Octane is used by [halo](../halo/) – Omni's consensus layer implementation. You can check out `halo` to learn how to use octane in your own application.

Halo also includes several EVM extensions e.g. for staking and slashing, which map directly to the staking and slashing cosmos-sdk modules. If you'd like to use cosmos modules in your own application, you can follow halo's model from for processing EVM logs.

## Status

Octane is a work in progress, and we plan to further build out features such as custom predeploys (like staking and slashing in halo), spinning up infra, connecting to various EVM clients, and more. If you'd like to contribute, please reach out to the team.
