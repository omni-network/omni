import type { Hex } from 'viem'
import { assertAcceptedResult, validateOrder } from '../api/validateOrder.js'
import type { ResolvedOrder } from '../contracts/parseOpenEvent.js'
import { type SendOrderParameters, sendOrder } from '../contracts/sendOrder.js'
import { waitForOrderOpen } from '../contracts/waitForOrderOpen.js'
import { waitForOrderStatus } from '../contracts/waitForOrderStatus.js'
import type { OptionalAbis } from '../types/abi.js'
import type { Environment } from '../types/config.js'

export type GenerateOrderParameters<abis extends OptionalAbis> =
  SendOrderParameters<abis> & {
    environment?: Environment | string
    pollingInterval?: number
  }

export type TerminalStatus = 'rejected' | 'closed' | 'filled'

export type OrderState =
  | { status: 'valid' }
  | { status: 'sent'; txHash: Hex }
  | { status: 'open'; txHash: Hex; order: ResolvedOrder }
  | { status: TerminalStatus; txHash: Hex; order: ResolvedOrder }

export async function* generateOrder<abis extends OptionalAbis>(
  params: GenerateOrderParameters<abis>,
): AsyncGenerator<OrderState> {
  const { environment, pollingInterval, ...sendOrderParams } = params

  const validationResult = await validateOrder(params.order, environment)
  assertAcceptedResult(validationResult)
  yield { status: 'valid' }

  const txHash = await sendOrder(sendOrderParams)
  yield { status: 'sent', txHash }

  const order = await waitForOrderOpen({
    client: params.client,
    txHash,
    pollingInterval,
  })
  yield { status: 'open', txHash, order }

  const status = await waitForOrderStatus({
    client: params.client,
    inboxAddress: params.inboxAddress,
    orderId: order.orderId,
    status: ['closed', 'filled', 'rejected'],
    pollingInterval,
  })
  yield { status, txHash, order }
}
