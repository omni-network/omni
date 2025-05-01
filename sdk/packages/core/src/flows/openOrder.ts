import { assertAcceptedResult, validateOrder } from '../api/validateOrder.js'
import type { ResolvedOrder } from '../contracts/parseOpenEvent.js'
import { type SendOrderParameters, sendOrder } from '../contracts/sendOrder.js'
import { waitForOrderOpen } from '../contracts/waitForOrderOpen.js'
import type { OptionalAbis } from '../types/abi.js'
import type { Environment } from '../types/config.js'

export type OpenOrderParameters<abis extends OptionalAbis> =
  SendOrderParameters<abis> & {
    environment?: Environment | string
    pollingInterval?: number
  }

export async function openOrder<abis extends OptionalAbis>(
  params: OpenOrderParameters<abis>,
): Promise<ResolvedOrder> {
  const { environment, pollingInterval, ...sendOrderParams } = params
  const validationResult = await validateOrder(params.order, environment)
  assertAcceptedResult(validationResult)
  const txHash = await sendOrder(sendOrderParams)
  return await waitForOrderOpen({
    client: params.client,
    txHash,
    pollingInterval,
  })
}
