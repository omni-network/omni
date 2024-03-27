# Explorer

The explorer has four components broken down into their own respective directory

```
/api
/db
/indexer
/ui
```

## GraphQL
We have a graphql server that serves data from our `db` component. We use `gqlgen` to generate our graphql schema and resolvers.

## DB
The db is driven by the [ent framework](entgo.io). It is an ORM where we generate our schemas as code and automatically create migrations and update our DB.

## Indexer
The Indexer listens to events published by the `xprovider` subscription that we generate for each rollup.

## UI
The UI is a [remix](https://remix.run/) application using `react`, `typescript` and hosted on the `remix app server`. We use `tailwind` to drive all of our CSS as well.

### How to run the explorer

In order to run the explorer locally you need a few things.
1. A devnet running on your local machine
2. Indexer
3. DB
3. GraphQL
4. UI

You can run this all with the following command:

```bash
make run-clean
```

This will create a fresh build of all the components, run a dev net and start the components in the background. You can access the UI at `http://localhost:3000`, you can access the GraphQL at `http://localhost:8080`.
