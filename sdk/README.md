# Omni SDK

## Overview

The Omni SDK is a TypeScript library for interfacing with the Omni network.

## Code Architecture

### Named Exports

We should always use named exports, as they're more friendly to tree-shaking, thus reducing bundle sizes. The alternative, `export *` may make it difficult for a bundler to identify what can be discarded.

### Barrel Files

Where possible, we use barrel files to re-export functionality from subdirectories. For example:

- `src/utils/index.ts` - re-exports utilOne and utilTwo
- `src/utils/utilOne.ts`
- `src/utils/utilTwo.ts`

Reasons for this are:

1. This allows us to simplify import statements for consumers:
   Instead of: `import { mainnet } from "@package/chains/mainnet"`
   We get: `import { mainnet } from "@package/chains"`
2. We can rename exports without changing consumer code
3. Internal file structures can evolve over time without affecting consumers
4. Reduces the number of import statements required by consumers

The top level barrel file (`src/index.ts`) re-exports functionality, that can be imported from the root:

`import { util } from "@package/sdk"`

Where logic is not exported here, it would be exported from a barrel file in the relevant subdirectory, imported as:

`import { util } from "@package/sdk/utils"`

The only exception to this is types, which we export directly from their modules, since:

1. Typescript handles `type` imports differently from value imports
2. Types are erased at runtime
3. Exporting types from an `index.ts` will avoid potential circular dependency issues
