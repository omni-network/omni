import {
  DidFillError,
  type GetOrderReturn,
  type InboxStatus,
  type OrderStatus,
  WatchDidFillError,
} from '@omni-network/core'
import type { Hex } from 'viem'
import type { QueryOpts } from './types.js'
import { useDidFill } from './useDidFill.js'
import { useGetOrder } from './useGetOrder.js'
import { useInboxStatus } from './useInboxStatus.js'
import type { useParseOpenEvent } from './useParseOpenEvent.js'
import { useWatchDidFill } from './useWatchDidFill.js'

type UseGetOrderStatusParams = {
  srcChainId?: number
  destChainId: number
  orderId?: Hex
  resolvedOrder?: ReturnType<typeof useParseOpenEvent>['resolvedOrder']
  getOrderQueryOpts?: QueryOpts<GetOrderReturn>
  didFillQueryOpts?: QueryOpts<boolean>
}

export function useGetOrderStatus({
  srcChainId,
  destChainId,
  orderId,
  resolvedOrder,
  getOrderQueryOpts,
  didFillQueryOpts,
}: UseGetOrderStatusParams) {
  // if resolved order is passed, we don't need to fetch the order
  const getOrder = useGetOrder({
    chainId: srcChainId,
    orderId,
    enabled: !resolvedOrder,
    queryOpts: getOrderQueryOpts,
  })

  const resolved = resolvedOrder ?? getOrder.data?.[0]
  const inboxStatus = useInboxStatus({
    orderId,
    chainId: srcChainId,
  })

  const didFill = useDidFill({
    destChainId,
    resolvedOrder: resolved,
    queryOpts: didFillQueryOpts,
  })

  const watchDidFill = useWatchDidFill({
    destChainId,
    orderId,
  })

  const status = deriveStatus(inboxStatus, didFill, watchDidFill)
  const error = deriveError(didFill, watchDidFill)

  return {
    status,
    error,
    destTxHash: watchDidFill.destTxHash,
    unwatchDestTx: watchDidFill.unwatch,
  }
}

function deriveError(
  didFill: ReturnType<typeof useDidFill>,
  watchDidFill: ReturnType<typeof useWatchDidFill>,
) {
  if (watchDidFill.status === 'error' && watchDidFill.error)
    return new WatchDidFillError(watchDidFill.error.message)
  if (didFill.status === 'error' && didFill.error)
    return new DidFillError(didFill.error.message)
  return
}

function deriveStatus(
  inboxStatus: InboxStatus,
  didFill: ReturnType<typeof useDidFill>,
  watchDidFill: ReturnType<typeof useWatchDidFill>,
): OrderStatus {
  if (didFill.status === 'error' && didFill.error) return 'error'
  // prioritise didFill since if consumers use a public RPC, event parsing
  // can sometimes fail, so at least we can inform fill status, but won't
  // propagate dest tx hash
  if (didFill.status === 'success' && didFill.data) return 'filled'
  if (watchDidFill.status === 'success' && watchDidFill.destTxHash)
    return 'filled'
  if (inboxStatus === 'open') return 'open'
  if (inboxStatus === 'filled') return 'filled'
  if (inboxStatus === 'rejected') return 'rejected'
  if (inboxStatus === 'closed') return 'closed'
  return 'not-found'
}
