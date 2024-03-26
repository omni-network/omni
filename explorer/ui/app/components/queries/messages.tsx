import { useQuery } from 'urql'
import { XMsg } from '~/graphql/graphql'
import { xblockrange } from './block'

export function GetXMessagesInRange(amount: number, offset: number): XMsg[] {
  let amt = '0x' + amount.toString(16)
  let off = '0x' + offset.toString(16)

  const [result] = useQuery({
    query: xblockrange,
    variables: {
      amount: amt,
      offset: off,
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
      }
      rows.push(xmsg)
    })
  })
  return rows
}
