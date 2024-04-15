import { useQuery } from 'urql'
import { graphql } from '~/graphql'
import { XMsg } from '~/graphql/graphql'

export function GetXMessagesInRange(from: number, to: number): XMsg[] {
  const [result] = useQuery({
    query: xmsgrange,
    variables: {
      from: '0x' + from.toString(16),
      to: '0x' + to.toString(16),
    },
  })
  const { data, fetching, error } = result

  var rows: XMsg[] = []
  data?.xmsgrange.map((xblock: any) => {
    if (xblock.Messages.length == 0) {
      return
    }

    xblock.Messages.map((msg: any) => {
      let xmsg = {
        DestAddress: msg.DestAddress,
        DestChainID: msg.DestChainID,
        DestGasLimit: msg.DestGasLimit,
        SourceChainID: msg.SourceChainID,
        SourceMessageSender: msg.SourceMessageSender,
        StreamOffset: msg.StreamOffset,
        TxHash: msg.TxHash,
        BlockHeight: msg.BlockHeight,
        BlockHash: msg.BlockHash,
        Receipts: msg.Receipts,
        Block: msg.Block,
      }
      rows.push(xmsg)
    })
  })
  return rows
}

export function GetMessage(sourceChainID: number, destChainID: number, streamOffset: number): XMsg {
  const [result] = useQuery({
    query: xmsg,
    variables: {
      sourceChainID: '0x' + sourceChainID.toString(16),
      destChainID: '0x' + destChainID.toString(16),
      streamOffset: '0x' + streamOffset.toString(16),
    },
  })
  const { data, fetching, error } = result

  if (data?.xmsg === undefined || data?.xmsg === null) {
    return {
      DestAddress: '',
      DestChainID: '',
      DestGasLimit: '',
      SourceChainID: '',
      SourceMessageSender: '',
      StreamOffset: '',
      TxHash: '',
      BlockHeight: '',
      BlockHash: '',
      Receipts: [],
      Block: {
        UUID: '',
        SourceChainID: '',
        BlockHeight: '',
        BlockHash: '',
        Timestamp: '',
        Messages: [],
        Receipts: [],
      },
    }
  }

  return data?.xmsg as XMsg
}

export function GetBlockCount(): number {
  const [result] = useQuery({
    query: xmsgcount,
  })
  const { data, fetching, error } = result
  let hex = data?.xmsgcount
  return Number(hex)
}

export const xmsg = graphql(`
  query XMsg($sourceChainID: BigInt!, $destChainID: BigInt!, $streamOffset: BigInt!) {
    xmsg(sourceChainID: $sourceChainID, destChainID: $destChainID, streamOffset: $streamOffset) {
      StreamOffset
      SourceMessageSender
      DestAddress
      DestGasLimit
      SourceChainID
      DestChainID
      TxHash
      BlockHeight
      BlockHash
      Block {
        SourceChainID
        BlockHeight
        BlockHash
        Timestamp
      }
      Receipts {
        GasUsed
        Success
        RelayerAddress
        SourceChainID
        DestChainID
        StreamOffset
        TxHash
        Timestamp
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
