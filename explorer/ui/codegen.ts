
import type { CodegenConfig } from '@graphql-codegen/cli';

const config: CodegenConfig = {
  overwrite: true,
  schema: "./app/graphql/schema.graphql",
  documents: "src/**/*.tsx",
  ignoreNoDocuments: true, // for better experience with the watcher
  generates: {
    "./app/graphql/gql/": {
      preset: "client",
      plugins: []
    }
  }
};

export default config;
