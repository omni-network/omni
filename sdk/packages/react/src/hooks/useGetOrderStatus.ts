import {
  DidFillError,
  type InboxStatus,
  type OrderStatus,
  WatchDidFillError,
} from '@omni-network/core'
import type { QueryOpts } from './types.js'
import { useDidFill } from './useDidFill.js'
import { useInboxStatus } from './useInboxStatus.js'
import type { useParseOpenEvent } from './useParseOpenEvent.js'
import { useWatchDidFill } from './useWatchDidFill.js'

type UseGetOrderStatusParams = {
  srcChainId?: number
  destChainId: number
  resolvedOrder?: ReturnType<typeof useParseOpenEvent>['resolvedOrder']
  didFillQueryOpts?: QueryOpts<boolean>
}

export function useGetOrderStatus({
  srcChainId,
  destChainId,
  resolvedOrder,
  didFillQueryOpts,
}: UseGetOrderStatusParams) {
  const inboxStatus = useInboxStatus({
    orderId: resolvedOrder?.orderId,
    chainId: srcChainId,
  })

  const didFill = useDidFill({
    destChainId,
    resolvedOrder,
    queryOpts: didFillQueryOpts,
  })

  const watchDidFill = useWatchDidFill({
    destChainId,
    resolvedOrder,
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
