import { useQuery } from 'urql'
import { graphql } from '~/graphql'
import { XBlock, XMsg } from '~/graphql/graphql'

export function GetBlocksInRange(from: number, to: number): XBlock[] {
  const [result] = useQuery({
    query: xblockrange,
    variables: {
      from: '0x' + from.toString(16),
      to: '0x' + to.toString(16),
    },
  })
  const { data, fetching, error } = result
  var rows: XBlock[] = []

  data?.xblockrange.map((xblock: any) => {
    let block = {
      id: xblock.BlockHeight,
      UUID: '',
      SourceChainID: xblock.SourceChainID,
      BlockHash: xblock.BlockHash,
      BlockHeight: xblock.BlockHeight,
      Messages: [],
      Timestamp: xblock.Timestamp,
      Receipts: [],
    }

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

export const GetXBlock = (sourceChainID: string, height: string): XBlock | null => {
  const [result] = useQuery({
    query: xblock,
    variables: {
      sourceChainID,
      height,
    },
  })
  const { data, fetching, error } = result
  // TODO handle error properly here
  if (!error) {
    return data as XBlock
  } else {
    return null
  }
}

export const xblock = graphql(`
  query Xblock($sourceChainID: BigInt!, $height: BigInt!) {
    xblock(sourceChainID: $sourceChainID, height: $height) {
      SourceChainID
      BlockHeight
      BlockHash
      Timestamp
      Messages {
        StreamOffset
        SourceMessageSender
        DestAddress
        DestGasLimit
        SourceChainID
        DestChainID
        TxHash
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

export const xblockrange = graphql(`
  query XBlockRange($from: BigInt!, $to: BigInt!) {
    xblockrange(from: $from, to: $to) {
      SourceChainID
      BlockHash
      BlockHeight
      Timestamp
    }
  }
`)

export const xblockcount = graphql(`
  query XblockCount {
    xblockcount
  }
`)
