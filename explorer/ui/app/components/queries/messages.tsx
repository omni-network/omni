import { useQuery } from 'urql'
import { XMsg } from '~/graphql/graphql'
import { xblockrange } from './block'

export function GetXMessagesInRange(from: number, to: number): XMsg[] {
  const [result] = useQuery({
    query: xblockrange,
    variables: {
      from: '0x' + from.toString(16),
      to: '0x' + to.toString(16),
    },
  })
  const { data, fetching, error } = result

  var rows: XMsg[] = []
  data?.xblockrange.map((xblock: any) => {
    if (xblock.Messages.length == 0) {
      return
    }

    xblock.Messages.map((msg: any) => {
      let xmsg = {
        DestAddress: msg.DestAddress,
        DestChainID: msg.DestChainID,
        DestGasLimit: '',
        SourceChainID: msg.SourceChainID,
        SourceMessageSender: '',
        StreamOffset: '',
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
