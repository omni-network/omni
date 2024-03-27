import type { CodegenConfig } from '@graphql-codegen/cli'

const config: CodegenConfig = {
  overwrite: true,
  schema: '../graphql/app/schema.graphql',
  documents: ['./app/components/**/*.ts?(x)'],
  generates: {
    './app/graphql/': {
      preset: 'client',
      plugins: [],
    },
  },
}

export default config
