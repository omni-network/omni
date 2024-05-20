# Explorer

The explorer has four components broken down into their own respective directory

```bash
/db
/graphql
/indexer
/ui
```

## DB

The db is driven by the [ent framework](entgo.io) - an ORM which generates the schemas as code and automatically create migrations and update the DB.

## GraphQL

We have a graphql server that serves data from our `db` component. We use `gqlgen` to generate our graphql schema and resolvers.

## Indexer

The Indexer listens to events published by the `xprovider` subscription that we generate for each rollup.

## UI

The UI is a [remix](https://remix.run/) application using `react`, `typescript` and hosted on the `remix app server`. We use `tailwind` to drive all of our CSS as well.

## Running the Explorer

In order to run the explorer locally you need a few things:

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

#### GraphQL API

Locally visit: `http://localhost:21335` for the GraphiQL interface

We currently have 3 chains in the devnet currently, those chainids are:

- 100
- 200
- 16561

The documentation for the graphql interfact is intuative and self explanitory following the local url the documentation for all queries can be found in the top left icon.

### Explorer UI

Currently we aren't deploying the UI (yet) as part of our devnet. You can run it from the `explorer/ui` folder with the following:

```bash
make run-explorer
```

For local development in the `explorer/ui` folder:

```bash
pnpm run dev
```

## Understanding the explorer UI

Here is some quick infomation on what has been developed for the explorer UI to familurize other devs who wish to expand on the UI's functionality.

### Pages

This project currently consists of two main pages:

1. **xmsgs Page**: Located under `_index.tsx`, this serves as the landing page where users can view all xmsgs.

2. **xmsg Page**: Located under `route.tsx`, this is the view for a single xmsg. Users can navigate to this page by clicking on a relevant ID within the xmsgs table, such as "123-456-7890".

Feel free to explore these pages to get a better understanding of the project.

### Query Handling in the Project

In this project, GraphQL queries are managed in the `ui/queries` directory. If there are any changes to the schema, updates must be made in both the `ui/queries` file and in the relevant page to remap the data.

To ensure that queries remain up-to-date and strongly typed, we utilize code generation. Simply run `pnpm run codegen` to update the data types.

Queries are invoked in the loader function and accessed using the `useRevalidator` hook. For search functionality, we extract the query from the URL and apply the search parameters within the loader function.

This setup ensures efficient handling of queries and facilitates maintenance and updates as needed.
