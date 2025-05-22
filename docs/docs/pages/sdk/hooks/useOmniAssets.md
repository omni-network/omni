---
sidebar_position: 4
title: useOmniAssets
---

# `useOmniAssets`

The `useOmniAssets` hook can be used to obtain the supported assets from the solver API. The environment is controlled based on what you set in your `OmniProvider` (testnet or mainnet).


## Usage

`import { useOmniAssets } from '@omni-network/react'`

```tsx
import { useOmniAssets } from '@omni-network/react'

function Component() {
    const assets = useOmniAssets()
}
```

## Parameters

### queryOpts

`Omit<UseQueryOptions, 'queryKey' | 'queryFn' | 'enabled'>`

Optional query options for the react query fetching assets - consult [`@tanstack/react-query`](https://tanstack.com/query/latest/docs/react/reference/useQuery) for more information.

## Return

### query

`UseQueryResult<OmniAssets, FetchJSONError>`

The return is a query object from [`@tanstack/react-query`](https://tanstack.com/query/latest/docs/react/reference/useQuery) - consult their documentation for all available properties.

### data

`OmniAssets | undefined`

Contains the addresses of the Omni contracts.
