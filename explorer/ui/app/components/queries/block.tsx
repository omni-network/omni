import { useQuery } from 'urql'
import { graphql } from '~/graphql'
import { XBlock, XMsg } from '~/graphql/graphql'

export function GetBlocksInRange(amount: number, offset: number): XBlock[] {
  const [result] = useQuery({
    query: xblockrange,
    variables: {
      from: '0x' + amount.toString(16),
      to: '0x' + offset.toString(16),
    },
  })
  const { data, fetching, error } = result

  var rows: XBlock[] = []

  data?.xblockrange.map((xblock: any) => {
    var msgs: XMsg[] = []
    let block = {
      id: xblock.BlockHeight,
      UUID: '',
      SourceChainID: xblock.SourceChainID,
      BlockHash: xblock.BlockHash,
      BlockHeight: xblock.BlockHeight,
      Messages: msgs,
      Timestamp: xblock.Timestamp,
      Receipts: [],
    }

    xblock.Messages.map((msg: any) => {
      let xmsg = {
        DestAddress: '',
        DestChainID: '',
        DestGasLimit: '',
        SourceChainID: '',
        SourceMessageSender: '',
        StreamOffset: '',
        TxHash: '',
        BlockHeight: '',
        BlockHash: '',
      }
      msgs.push(xmsg)
    })

    block.Messages = msgs
    rows.push(block)
  })

  return rows
}

export function GetBlockCount(): number {
  const [result] = useQuery({
    query: xblockcount,
  })
  const { data, fetching, error } = result
  let hex = data?.xblockcount
  return Number(hex)
}

export const xblock = graphql(`
  query Xblock($sourceChainID: BigInt!, $height: BigInt!) {
    xblock(sourceChainID: $sourceChainID, height: $height) {
      BlockHash
    }
  }
`)

export const xblockrange = graphql(`
  query XBlockRange($from: BigInt!, $to: BigInt!) {
    xblockrange(from: $from, to: $to) {
      SourceChainID
      BlockHash
      BlockHeight
      Messages {
        TxHash
        DestAddress
        DestChainID
        SourceChainID
      }
      Timestamp
    }
  }
`)

export const xblockcount = graphql(`
  query XblockCount {
    xblockcount
  }
`)
