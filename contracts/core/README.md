# Omni Core Contracts

Omni's core protocol smart contracts.

## Contents

<pre>
└── <a href="./src/">src/</a>
  ├── <a href="./src/xchain/">xchain/</a>: Cross-chain protocol smart contracts (incl OmniPortal)
  ├── <a href="./src/octane/">octane/</a>: Octane protocol smart contracts
  ├── <a href="./src/interfaces/">interfaces/</a>: All interfaces
  ├── <a href="./src/libraries/">libraries/</a>: All libraries
  ├── <a href="./src/pkg/">pkg/</a>: Exported utility contracts
  ├── <a href="./src/utils/">utils/</a>: Internal utility contracts
  ├── <a href="./src/token/">token/</a>: Token & bridge contracts
  ├── <a href="./src/deploy/">deploy/</a>: Deployment utilities
  └── <a href="./src/examples/">examples/</a>: Example cross-chain contracts
</pre>

## Development

```bash
pnpm i          # install dependencies
pnpm test:gen   # generate test fixtures
forge test      # run forge tests
pnpm test       # alias for (pnpm test:gen && forge test)
```
