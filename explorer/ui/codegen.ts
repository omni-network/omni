
import type { CodegenConfig } from '@graphql-codegen/cli';

const config: CodegenConfig = {
  overwrite: true,
  schema: "../graphql/app/schema.graphql",
  documents: "src/**/*.tsx",
  ignoreNoDocuments: true, // for better experience with the watcher
  generates: {
    "./app/graphql/": {
      preset: "client",
      plugins: []
    }
  }
};

export default config;
