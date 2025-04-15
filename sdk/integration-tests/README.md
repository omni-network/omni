# Integration tests

## Running

Before running the tests, you need to deploy devnet.

```bash
make devnet-deploy MANIFEST=devnet1
```

Install dependencies

```bash
pnpm install
```

Run the tests

```bash
pnpm run test:integration
```
