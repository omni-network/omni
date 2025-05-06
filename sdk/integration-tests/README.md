# Integration tests

## Running

Before running the tests, you need to deploy devnet. You can do this by first setting up e2e (see [README](../../e2e/README.md)) and then running from the monorepo root:

```bash
make devnet-deploy MANIFEST=devnet1
```

Install dependencies in this directory

```bash
pnpm install
```

Run the tests from this directory

```bash
pnpm run test:integration
```

To run a specific test you can supply the test file and line number

```bash
pnpm run test suites/react-hooks.tsx:123
```
