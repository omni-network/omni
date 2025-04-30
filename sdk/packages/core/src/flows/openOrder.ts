import { assertAcceptedResult, validateOrder } from '../api/validateOrder.js'
import type { ResolvedOrder } from '../contracts/parseOpenEvent.js'
import {
  type SendOrderTransactionParameters,
  sendOrderTransaction,
} from '../contracts/sendOrderTransaction.js'
import { waitForOrderOpen } from '../contracts/waitForOrderOpen.js'
import type { OptionalAbis } from '../types/abi.js'

export type OpenOrderParameters<abis extends OptionalAbis> =
  SendOrderTransactionParameters<abis> & {
    apiBaseUrl: string
    pollingInterval?: number
  }

export async function openOrder<abis extends OptionalAbis>(
  params: OpenOrderParameters<abis>,
): Promise<ResolvedOrder> {
  const { apiBaseUrl, pollingInterval, ...sendOrderParams } = params
  const validationResult = await validateOrder(apiBaseUrl, params.order)
  assertAcceptedResult(validationResult)
  const txHash = await sendOrderTransaction(sendOrderParams)
  return await waitForOrderOpen({
    client: params.client,
    txHash,
    pollingInterval,
  })
}
