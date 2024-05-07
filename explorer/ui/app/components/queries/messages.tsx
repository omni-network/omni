import { DocumentNode } from 'graphql'
import { gql, useQuery } from 'urql'
import { gqlClient } from '~/entry.server'
import { graphql } from '~/graphql'
import { XMsg, XMsgsDocument } from '~/graphql/graphql'

export const GetXMsg = async (sourceChainID: string, destChainID: string, streamOffset: string) => {
  console.log("======");
  return await gqlClient.query(xmsg, {sourceChainID, destChainID , streamOffset})
  // const x = useQuery({query: xmsg, variables: {sourceChainID, destChainID , streamOffset}})
  // console.log(x);
  // return null
  // const [result] = useQuery({
  //   query: xmsg,
  //   variables: {
  //     sourceChainID,
  //     destChainID,
  //     streamOffset,
  //   },
  // })
  // const { data, fetching, error } = result
  // // TODO handle error properly here
  // if (!error) {
  //   return data as XMsg
  // } else {
  //   return null
  // }
}



export const xmsg = gql(`
  query xmsg($sourceChainID: BigInt!, $destChainID: BigInt!, $offset: BigInt!) {
    xmsg(sourceChainID: $sourceChainID, destChainID: $destChainID, offset: $offset) {
      streamOffset
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

export const xmsgs = gql(`
query {
  xmsgs(first: 10, after: "eyJJRCI6MTksIlBhZ2VOdW0iOjJ9") {
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
  txHashUrl
  block {
    id
    sourceChainID
    hash
    height
    timestamp
  }
  receipt {
    txHash
    txHashUrl
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
