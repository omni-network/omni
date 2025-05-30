import type { Address, Block, Client, Hex } from 'viem'
import { getLogs, watchBlocks } from 'viem/actions'
import { outboxABI } from '../constants/abis.js'
import { invariant } from '../utils/index.js'

export type WatchDidFillParams = {
  client: Client
  outboxAddress: Address
  orderId: Hex
  onFill: (txHash: Hex) => void
  pollingInterval?: number
  onError?: (error: Error) => void
}

export type WatchDidFillReturn = () => void

export function watchDidFill({
  client,
  outboxAddress,
  orderId,
  pollingInterval,
  onFill,
  onError,
}: WatchDidFillParams) {
  return watchBlocks(client, {
    onBlock: async (block: Block) => {
      try {
        // default block tag is 'latest', hash is only undefined for pending blocks
        invariant(!!block.hash, 'Block hash is undefined')

        const log = await getFilledLog({
          client,
          outboxAddress,
          orderId,
          blockHash: block.hash,
        })

        if (log) onFill(log.transactionHash)
      } catch (error) {
        onError?.(error as Error)
      }
    },
    pollingInterval,
    onError,
    emitMissed: true,
  })
}

async function getFilledLog({
  client,
  outboxAddress,
  orderId,
  blockHash,
}: Pick<WatchDidFillParams, 'client' | 'outboxAddress' | 'orderId'> & {
  blockHash: Hex
}) {
  const logs = await getLogs(client, {
    address: outboxAddress,
    event: getFilledEventAbi(),
    args: {
      orderId,
    },
    blockHash,
  })

  // it's safe to assume we always get a single Filled event per orderId
  if (logs.length > 0) return logs[0]

  return
}

const getFilledEventAbi = () => {
  const abi = outboxABI.find(
    (item) => item.type === 'event' && item.name === 'Filled',
  )
  invariant(!!abi, 'Filled event not found')
  return abi
}
