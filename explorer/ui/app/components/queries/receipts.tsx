import { useQuery } from 'urql'
import { graphql } from '~/graphql'
import { XReceipt } from '~/graphql/graphql'

export function GetReceipt(
  sourceChainID: number,
  destChainID: number,
  streamOffset: number,
): XReceipt {
  const [result] = useQuery({
    query: xmsg,
    variables: {
      sourceChainID: '0x' + sourceChainID.toString(16),
      destChainID: '0x' + destChainID.toString(16),
      streamOffset: '0x' + streamOffset.toString(16),
    },
  })
  const { data, fetching, error } = result

  if (data?.xreceipt === undefined || data?.xreceipt === null) {
    return {
      UUID: '',
      GasUsed: '',
      Success: false,
      RelayerAddress: '',
      SourceChainID: '',
      DestChainID: '',
      StreamOffset: '',
      TxHash: '',
      Timestamp: '',
      Block: {
        SourceChainID: '',
        BlockHeight: '',
        BlockHash: '',
        Timestamp: '',
        Messages: [],
        Receipts: [],
        UUID: '',
      },
      Messages: [],
    }
  }

  return data?.xreceipt as XReceipt
}

export const xmsg = graphql(`
  query XReceipt($sourceChainID: BigInt!, $destChainID: BigInt!, $streamOffset: BigInt!) {
    xreceipt(
      sourceChainID: $sourceChainID
      destChainID: $destChainID
      streamOffset: $streamOffset
    ) {
      GasUsed
      Success
      RelayerAddress
      SourceChainID
      DestChainID
      StreamOffset
      TxHash
      Timestamp
      Block {
        SourceChainID
        BlockHeight
        BlockHash
        Timestamp
      }
      Messages {
        StreamOffset
        SourceMessageSender
        DestAddress
        DestGasLimit
        SourceChainID
        DestChainID
        TxHash
      }
    }
  }
`)
