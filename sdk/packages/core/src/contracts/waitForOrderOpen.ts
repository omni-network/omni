import type { Client, Hex } from 'viem'
import { waitForTransactionReceipt } from 'viem/actions'
import { type ResolvedOrder, parseOpenEvent } from './parseOpenEvent.js'

export type WaitForOrderOpenParameters = {
  client: Client
  txHash: Hex
  pollingInterval?: number
}

export async function waitForOrderOpen(
  params: WaitForOrderOpenParameters,
): Promise<ResolvedOrder> {
  const receipt = await waitForTransactionReceipt(params.client, {
    hash: params.txHash,
    pollingInterval: params.pollingInterval,
  })
  return parseOpenEvent(receipt.logs)
}
