import { gql } from 'urql'
import { graphql } from '~/graphql'

export const xmsg = gql(`
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
      sourceChainID
      destChainID
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

export const xmsgrange = graphql(`
  query XMsgRange($from: BigInt!, $to: BigInt!) {
    xmsgrange(from: $from, to: $to) {
      StreamOffset
      SourceMessageSender
      DestAddress
      DestGasLimit
      SourceChainID
      DestChainID
      TxHash
      BlockHeight
      BlockHash
    }
  }
`)

export const xmsgcount = graphql(`
  query XMsgCount {
    xmsgcount
  }
`)

// query xmsgs($first: Int, $last: Int, $after: Id, $before: Id) {
//   xmsgs(first: $first, last: $last, after: $after, before: $before) {

export const xmsgs = gql(`
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
        sourceChainID
        sender
        senderUrl
        to
        toUrl
        destChainID
        gasLimit
        status
        txHash
        txUrl
        block {
          id
          chainID
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
          sourceChainID
          destChainID
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
