# Omni Contract Admin Commands

Defines contract admin commands for live omni networks.

## Contract Specs

We define a number of `ensure-<contract>-spec` commands. These commands
apply local contract specifications defined in here to live contracts, if
necessary. This includes pausing / unpausing, setting parameters, etc.

### `ensure-bridge-spec`

- specs defined in `bridgespec.go`
- manages `OmniBridgeL1` and `OmniBridgeNative` contracts

### `ensure-portal-spec`

- specs defined in `portalspec.go`
- manages `OmniPortal` deployments
