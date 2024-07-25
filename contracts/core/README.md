# Omni Core Contracts

Omni's core protocol smart contracts.

## Contents

- `src/`
    - `xchain/`     : Cross-chain protocol smart contracts (incl OmniPortal)
    - `octane/`     : Octane protocol smart contracts
    - `interfaces/` : All interfaces
    - `libraries/`  : All libraries
    - `pkg/`        : Exported utility contracts
    - `utils/`      : Internal utility contracts
    - `token/`      : Token & bridge contracts
    - `deploy/`     : Deployment utilities
    - `examples/`   : Example cross-chain contracts

## Development

```bash
pnpm i          # install dependencies
pnpm test:gen   # generate test fixtures
forge test      # run forge tests
pnpm test       # alias for (pnpm test:gen && forge test)
```
