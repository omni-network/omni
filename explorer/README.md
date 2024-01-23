# Explorer

The explorer has four components broken down into their own respective directory

```
/api
/db
/indexer
/ui
```

## API
The api is a simple go application where we build endpoints that serve data from our `db` component. The api is driven by its `openapi.yaml` file. Where we generate all of our requests, we also copy this file into our `ui` component such that we have no discrepancies between what we are publishing and what we are fetching

## DB
The db is driven by the [ent framework](entgo.io). It is an ORM where we generate our schemas as code and automatically create migrations and update our DB.

## Indexer
The Indexer listens to events published by the `xprovider` subscription that we generate for each rollup.

## UI
The UI is a [remix](https://remix.run/) application using `react`, `typescript` and hosted on the `remix app server`. We use `tailwind` to drive all of our CSS as well.
