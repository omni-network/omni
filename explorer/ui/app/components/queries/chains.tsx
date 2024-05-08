import { gql } from 'urql'

export const supportedChains = gql(`
  query supportedChains {
      id
      chainID
      name
    }
`)
