import type { Hex } from 'viem'
import { assertAcceptedResult, validateOrder } from '../api/validateOrder.js'
import type { ResolvedOrder } from '../contracts/parseOpenEvent.js'
import { type SendOrderParameters, sendOrder } from '../contracts/sendOrder.js'
import {
  type TerminalStatus,
  waitForOrderClose,
} from '../contracts/waitForOrderClose.js'
import { waitForOrderOpen } from '../contracts/waitForOrderOpen.js'
import { watchDidFill } from '../contracts/watchDidFill.js'
import type { OptionalAbis } from '../types/abi.js'
import type { EVMAddress } from '../types/addresses.js'
import type { Environment } from '../types/config.js'

export type GenerateOrderParameters<abis extends OptionalAbis> =
  SendOrderParameters<abis> & {
    environment?: Environment | string
    pollingInterval?: number
    outboxAddress: EVMAddress
  }

export type OrderState =
  | { status: 'valid'; txHash?: never; order?: never }
  | { status: 'sent'; txHash: Hex; order?: never }
  | { status: 'open'; txHash: Hex; order: ResolvedOrder }
  | {
      status: TerminalStatus
      txHash: Hex
      order: ResolvedOrder
      destTxHash?: Hex
    }

export async function* generateOrder<abis extends OptionalAbis>(
  params: GenerateOrderParameters<abis>,
): AsyncGenerator<OrderState> {
  const { environment, pollingInterval, ...sendOrderParams } = params

  const validationResult = await validateOrder({
    ...params.order,
    environment,
  })
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

  let destTxHash: Hex | undefined
  const unwatch = watchDidFill({
    client: params.client,
    outboxAddress: params.outboxAddress,
    resolvedOrder: order,
    pollingInterval,
    onFill: (txHash) => {
      destTxHash = txHash
      unwatch()
    },
  })

  const status = await waitForOrderClose({
    client: params.client,
    inboxAddress: params.inboxAddress,
    orderId: order.orderId,
    pollingInterval,
  })

  yield { status, txHash, order, destTxHash }
}
