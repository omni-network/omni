import type { Block, Client, Hex } from 'viem'
import { getLogs, watchBlocks } from 'viem/actions'
import { outboxABI } from '../constants/abis.js'
import type { EVMAddress } from '../types/addresses.js'
import { computeFillHash } from '../utils/computeFillHash.js'
import { invariant } from '../utils/index.js'
import type { ResolvedOrder } from './parseOpenEvent.js'

export type WatchDidFillParams = {
  client: Client
  outboxAddress: EVMAddress
  resolvedOrder: ResolvedOrder
  onFill: (txHash: Hex) => void
  pollingInterval?: number
  onError?: (error: Error) => void
}

export function watchDidFill({
  client,
  outboxAddress,
  resolvedOrder,
  pollingInterval,
  onFill,
  onError,
}: WatchDidFillParams) {
  const fillHash = computeFillHash(resolvedOrder)

  return watchBlocks(client, {
    onBlock: async (block: Block) => {
      try {
        // default block tag is 'latest', hash is only undefined for pending blocks
        invariant(!!block.hash, 'Block hash is undefined')

        const log = await getFilledLog({
          client,
          outboxAddress,
          resolvedOrder,
          blockHash: block.hash,
          fillHash,
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
  resolvedOrder,
  blockHash,
  fillHash,
}: Pick<WatchDidFillParams, 'client' | 'outboxAddress' | 'resolvedOrder'> & {
  blockHash: Hex
  fillHash: Hex
}) {
  const logs = await getLogs(client, {
    address: outboxAddress,
    event: getFilledEventAbi(),
    args: {
      orderId: resolvedOrder.orderId,
      fillHash,
    },
    blockHash,
  })

  // it's safe to assume we always get a single Filled event per orderId and fillHash
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
