import { gql } from 'urql'

export const supportedChains = gql(`
  query supportedChains {
    supportedChains {
    id
    chainID
    name
    logoUrl
    }
  }
`)
export const chainStats = gql(`
query chainStats {
  stats {
    totalMsgs
    topStreams {
      sourceChain{
        id
        chainID
        displayID
        name
        logoUrl
      }
      destChain {
        id
        chainID
        displayID
        name
        logoUrl
      }
      msgCount
    }
  }
}
`)
