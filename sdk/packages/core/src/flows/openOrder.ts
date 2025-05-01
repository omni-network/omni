import { assertAcceptedResult, validateOrder } from '../api/validateOrder.js'
import type { ResolvedOrder } from '../contracts/parseOpenEvent.js'
import { type SendOrderParameters, sendOrder } from '../contracts/sendOrder.js'
import { waitForOrderOpen } from '../contracts/waitForOrderOpen.js'
import type { OptionalAbis } from '../types/abi.js'

export type OpenOrderParameters<abis extends OptionalAbis> =
  SendOrderParameters<abis> & {
    apiBaseUrl: string
    pollingInterval?: number
  }

export async function openOrder<abis extends OptionalAbis>(
  params: OpenOrderParameters<abis>,
): Promise<ResolvedOrder> {
  const { apiBaseUrl, pollingInterval, ...sendOrderParams } = params
  const validationResult = await validateOrder(apiBaseUrl, params.order)
  assertAcceptedResult(validationResult)
  const txHash = await sendOrder(sendOrderParams)
  return await waitForOrderOpen({
    client: params.client,
    txHash,
    pollingInterval,
  })
}
