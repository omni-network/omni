# Omni Core Contracts

Omni's core protocol smart contracts.

## Contents

<pre>
└── <a href="./src/">src/</a>
  ├── <a href="./xchain/">xchain/</a>: Cross-chain protocol smart contracts (incl OmniPortal)
  ├── <a href="./octane/">octane/</a>: Octane protocol smart contracts
  ├── <a href="./interfaces/">interfaces/</a>: All interfaces
  ├── <a href="./libraries/">libraries/</a>: All libraries
  ├── <a href="./pkg/">pkg/</a>: Exported utility contracts
  ├── <a href="./utils/">utils/</a>: Internal utility contracts
  ├── <a href="./token/">token/</a>: Token & bridge contracts
  ├── <a href="./deploy/">deploy/</a>: Deployment utilities
  └── <a href="./examples/">examples/</a>: Example cross-chain contracts
</pre>

## Development

```bash
pnpm i          # install dependencies
pnpm test:gen   # generate test fixtures
forge test      # run forge tests
pnpm test       # alias for (pnpm test:gen && forge test)
```
