import { graphql } from '~/graphql'

export const search = graphql(`
  query Search($query: String!) {
    search(query: $query) {
      Type
      BlockHeight
      TxHash
      SourceChainID
    }
  }
`)
