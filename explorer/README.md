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



### Devnet
You can run this all with the following command to build all the backend go binaries and run them. The components in the devnet include the omni network dev net, indexer, db and the graphql server.

### Fist Time Setup
You need to have docker installed and running to execute the following command:

You can either build all of the component with the following:
```bash
make build
```

then run:
```bash
make run-devnet
```

or run:
```bash
make run-clean
```

### Stopping the network
If you want to stop the local devnet you can run:

```bash
make stop
```

or from the `omni` root folder:

```bash
make devnet-clean
```

#### GraphQL

Locally visit: http://localhost:21335 for the GraphiQL interface

We currently have 3 chains in the devnet currently, those chainids are:
- 100
- 200
- 16561

So if you were to go and query an xblock, you would need to set one of those as the `sourceChainID` field.

ex:
```graphql
query{
  xblock(sourceChainID: 200, height: 35){
		SourceChainID
    BlockHeight
    Messages{
      StreamOffset
    }
    Receipts{
      TxHash
    }
  }
}
```

### Explorer UI
Currently we aren't deploying the UI (yet) as part of our devnet. You can run it from the `explorer/ui` folder with the following:

```bash
make run-explorer
```

For local development in the `explorer/ui` folder:
```bash
pnpm run dev
```
