import type { Address, Client } from 'viem'
import { readContract } from 'viem/actions'
import { outboxABI } from '../constants/abis.js'
import type { ResolvedOrder } from './parseOpenEvent.js'

export type DidFillParams = {
  client: Client
  outboxAddress: Address
  resolvedOrder: ResolvedOrder
}

export type DidFillReturn = boolean

export function didFillOutbox({
  client,
  outboxAddress,
  resolvedOrder,
}: DidFillParams) {
  return readContract(client, {
    address: outboxAddress,
    abi: outboxABI,
    functionName: 'didFill',
    args: [resolvedOrder.orderId, resolvedOrder.fillInstructions[0].originData],
  })
}
