import type { Address, Client, Hex, Log } from 'viem'
import { watchContractEvent } from 'viem/actions'
import { outboxABI } from '../constants/abis.js'

export type WatchDidFillParams = {
  client: Client
  outboxAddress: Address
  orderId: Hex
  onLogs: (logs: Log[]) => void
  pollingInterval?: number
  onError?: (error: Error) => void
}

export type WatchDidFillReturn = () => void

export function watchDidFill({
  client,
  outboxAddress,
  orderId,
  pollingInterval,
  onLogs,
  onError,
}: WatchDidFillParams): WatchDidFillReturn {
  return watchContractEvent(client, {
    address: outboxAddress,
    eventName: 'Filled',
    abi: outboxABI,
    args: { orderId },
    onLogs,
    onError,
    pollingInterval,
  })
}
