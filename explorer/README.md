# Explorer

The explorer has four components broken down into their own respective directory

```
/db
/graphql
/indexer
/ui
```

## DB
The db is driven by the [ent framework](entgo.io). It is an ORM where we generate our schemas as code and automatically create migrations and update our DB.

## GraphQL
We have a graphql server that serves data from our `db` component. We use `gqlgen` to generate our graphql schema and resolvers.

## Indexer
The Indexer listens to events published by the `xprovider` subscription that we generate for each rollup.

## UI
The UI is a [remix](https://remix.run/) application using `react`, `typescript` and hosted on the `remix app server`. We use `tailwind` to drive all of our CSS as well.

## Running the Explorer

In order to run the explorer locally you need a few things.
1. A devnet running on your local machine
2. Indexer
3. DB
4. GraphQL
5. UI

### Fist Time Setup

You can run this all with the following command to build all the images and run them:

```bash
make run-clean
```

This will create a fresh build of all the components, run a dev net and start the components in the background. You can access the UI at `http://localhost:3000`, you can access the GraphQL at `http://localhost:8080`.

If you want to stop the components you can run:

```bash
make stop
```

### Local Advanced Configurations

The explorer needs a copy of the network.json file to be able to know which network to connect to. You can find the network.json file in the `explorer` directory. If you want to run the explorer against a different network you can set the `NETWORK` environment variable to the network you want to run against. For example:

The following command assumes you have some network running locally, and you want to run the explorer against it. This will grab the `network.json` file via `@cp ../e2e/runs/$(NETWORK)/relayer/network.json` and copy it to the explorer directory. Then it will run the explorer.

```bash
make copy-network NETWORK=devnet-1
make run-explorer
```
