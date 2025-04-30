import { watchBlockNumber } from 'viem/actions'
import { GetOrderError } from '../errors/base.js'
import {
  type GetOrderParameters,
  type GetOrderReturn,
  getOrder,
} from './getOrder.js'

export type WatchOrderParameters = GetOrderParameters & {
  onOrder: (order: GetOrderReturn) => void
  onError?: (error: Error) => void
  pollingInterval?: number
}

export type WatchOrderReturn = () => void

export function watchOrder(params: WatchOrderParameters): WatchOrderReturn {
  const { onError, onOrder, pollingInterval, ...getOrderParams } = params
  return watchBlockNumber(params.client, {
    onBlockNumber: async () => {
      try {
        const order = await getOrder(getOrderParams)
        onOrder(order)
      } catch (err) {
        if (onError != null) {
          const error =
            err instanceof Error
              ? err
              : new GetOrderError('Failed to get order', String(err))
          onError(error)
        }
      }
    },
    onError,
    pollingInterval,
  })
}
