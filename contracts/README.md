# Omni Contracts

Omni smart contracts and related software.

## Contents

- `core/`: Core protocol smart contracts
- `avs/`: Eigen AVS smart contracts
- `bindings/`: Go smark contract bindings
- `allocs/`: Predeploy allocations

## Build

```bash
make build      # compile the smart contracts (avs & core)
make bindings   # generate go bindings
make allocs     # generate predeploy allocations
make all        # all of the above
```
