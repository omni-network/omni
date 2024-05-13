import { gql } from 'urql'
import { XmsgDocument, XmsgsDocument } from '~/graphql/graphql'

export const xmsg: typeof XmsgDocument = gql(`
  query xmsg($sourceChainID: BigInt!, $destChainID: BigInt!, $offset: BigInt!) {
    xmsg(sourceChainID: $sourceChainID, destChainID: $destChainID, offset: $offset) {
      id
      displayID
      offset
      sender
      senderUrl
      to
      toUrl
      gasLimit
      sourceChain{
        chainID
        logoUrl
        name
      }
      destChain{
        chainID
        logoUrl
        name
      }
      txHash
      txUrl
      status
      block {
        height
        hash
        timestamp
      }
      receipt {
        revertReason
        txHash
        txUrl
        relayer
        timestamp
        gasUsed
      }
    }
  }
`)

export const xmsgs: typeof XmsgsDocument = gql(`
query xmsgs($first: Int, $last: Int, $after: ID, $before: ID, $filters: [FilterInput!]) {
  xmsgs(first: $first, last: $last, after: $after, before: $before, filters: $filters) {
    totalCount
    edges {
      cursor
      node {
        id
        txHash
        offset
        displayID
        sender
        senderUrl
        to
        toUrl
        sourceChain{
          chainID
          logoUrl
          name
        }
        destChain{
          chainID
          logoUrl
          name
        }
        gasLimit
        status
        txHash
        txUrl
        block {
          id
          chain {
            chainID
          }
          hash
          height
          timestamp
        }
        receipt {
          txHash
          txUrl
          timestamp
          success
          offset
          sourceChain{
            chainID
          }
          destChain{
            chainID
          }
          relayer
          revertReason
        }
      }
    }
    pageInfo {
      currentPage
      totalPages
      hasNextPage
      hasPrevPage
    }
  }
}
`)
