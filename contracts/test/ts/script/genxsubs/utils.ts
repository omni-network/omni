import { XMsg } from '../../xtypes'

export function groupByDestChain(msgs: readonly XMsg[]) {
  return msgs.reduce(
    (acc, msg) => {
      const group = msg.destChainId.toString()
      if (!acc[group]) acc[group] = []
      acc[group].push(msg)
      return acc
    },
    {} as Record<string, XMsg[]>,
  )
}
