import { graphql } from '~/graphql'

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
