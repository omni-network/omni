# `getContracts`

The `getContracts` function fetches the addresses of the contracts used by SolverNet, that need to be provided for some interactions.

## Usage

`import { getContracts } from '@omni-network/core'`

```tsx
import { getContracts } from '@omni-network/core'

const contracts = await getContracts();
```

## Parameters

| Prop                | Type                                 | Required | Description                                                                                                                         |
| ------------------- | ------------------------------------ | -------- | ----------------------------------------------------------------------------------------------------------------------------------- |
| `environment`           | `Environment | string`                         | No      | SolverNet environment to use, either `mainnet` (default) or `testnet`. |

## Return

`getContracts` returns the [Promise](https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/Promise) of an `OmniContracts` object

### `OmniContracts`

```ts
type OmniContracts = {
  inbox: Address
  outbox: Address
}
```

## Example

### Get testnet contract addresses

```ts
const contracts = await getContracts('testnet')
```
