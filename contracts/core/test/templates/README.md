# Omni Test Templates

Each test template is setup with necessary fixtures to test a single contract. Available templates include

- [OmniPortal.t.sol](./OmniPortal.t.sol)
- [OmniBridgeNative.t.sol](./OmniBridgeNative.t.sol)
- [OmniBridgeL1.t.sol](./OmniBridgeL1.t.sol)
- [PortalRegistry.t.sol](./PortalRegistry.t.sol)
- [Slashing.t.sol](./Slashing.t.sol)
- [Staking.t.sol](./Staking.t.sol)
- [Upgrade.t.sol](./Upgrade.t.sol)

Each template includes list of available fixtures and their purpose. Some include example test cases. Use these templates as a base from which to build out test suites.

## Running Tests

To run tests, run `make test` from this directory. Or, `forge test --match-path 'test/templates/**'`
