import {
  DidFillError,
  GetOrderError,
  type InboxStatus,
  type OrderStatus,
} from '@omni-network/core'
import type { Hex } from 'viem'
import { useDidFillOutbox } from './useDidFillOutbox.js'
import { useGetOrder } from './useGetOrder.js'
import { useInboxStatus } from './useInboxStatus.js'
import type { useParseOpenEvent } from './useParseOpenEvent.js'

export function useGetOrderStatus({
  srcChainId,
  destChainId,
  orderId,
  resolvedOrder,
}: {
  srcChainId?: number
  destChainId: number
  orderId?: Hex
  resolvedOrder?: ReturnType<typeof useParseOpenEvent>['resolvedOrder']
}) {
  // if resolved order is passed, we don't need to fetch the order
  const getOrder = useGetOrder({
    chainId: srcChainId,
    orderId,
    enabled: !resolvedOrder,
  })

  const resolved = resolvedOrder ?? getOrder.data?.[0]

  const inboxStatus = useInboxStatus({
    orderId,
    chainId: srcChainId,
  })

  const didFillOutbox = useDidFillOutbox({
    destChainId,
    resolvedOrder: resolved,
  })

  const status = deriveStatus(inboxStatus, didFillOutbox)
  const error = deriveError(getOrder, didFillOutbox)

  return {
    status,
    error,
  }
}

function deriveError(
  getOrder: ReturnType<typeof useGetOrder>,
  didFillOutbox: ReturnType<typeof useDidFillOutbox>,
) {
  if (getOrder.error) return new GetOrderError(getOrder.error.message)
  if (didFillOutbox.error) return new DidFillError(didFillOutbox.error.message)
  return
}

function deriveStatus(
  inboxStatus: InboxStatus,
  didFillOutbox: ReturnType<typeof useDidFillOutbox>,
): OrderStatus {
  if (didFillOutbox.error) return 'error'
  if (didFillOutbox?.data === true) return 'filled'
  if (inboxStatus === 'filled') return 'filled'
  if (inboxStatus === 'open') return 'open'
  if (inboxStatus === 'rejected') return 'rejected'
  if (inboxStatus === 'closed') return 'closed'
  return 'not-found'
}
