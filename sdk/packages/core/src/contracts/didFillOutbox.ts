import type { Address, Chain, Client, Transport } from 'viem'
import { readContract } from 'viem/actions'
import { outboxABI } from '../constants/abis.js'
import type { ResolvedOrder } from './parseOpenEvent.js'

export type DidFillParams<chain extends Chain> = {
  client: Client<Transport, chain>
  outboxAddress: Address
  resolvedOrder: ResolvedOrder
}

export type DidFillReturn = boolean

export async function didFillOutbox<chain extends Chain>({
  client,
  outboxAddress,
  resolvedOrder,
}: DidFillParams<chain>): Promise<DidFillReturn> {
  return await readContract(client, {
    address: outboxAddress,
    abi: outboxABI,
    functionName: 'didFill',
    args: [resolvedOrder.orderId, resolvedOrder.fillInstructions[0].originData],
  })
}
