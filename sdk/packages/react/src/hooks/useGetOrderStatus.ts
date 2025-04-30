import {
  DidFillError,
  GetOrderError,
  type InboxStatus,
  type OrderStatus,
} from '@omni-network/core'
import type { Hex } from 'viem'
import { useDidFill } from './useDidFill.js'
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

  const didFill = useDidFill({
    destChainId,
    resolvedOrder: resolved,
  })

  const status = deriveStatus(inboxStatus, didFill)
  const error = deriveError(getOrder, didFill)

  return {
    status,
    error,
  }
}

function deriveError(
  getOrder: ReturnType<typeof useGetOrder>,
  didFill: ReturnType<typeof useDidFill>,
) {
  if (getOrder.error) return new GetOrderError(getOrder.error.message)
  if (didFill.error) return new DidFillError(didFill.error.message)
  return
}

function deriveStatus(
  inboxStatus: InboxStatus,
  didFill: ReturnType<typeof useDidFill>,
): OrderStatus {
  if (didFill.error) return 'error'
  if (didFill?.data === true) return 'filled'
  if (inboxStatus === 'filled') return 'filled'
  if (inboxStatus === 'open') return 'open'
  if (inboxStatus === 'rejected') return 'rejected'
  if (inboxStatus === 'closed') return 'closed'
  return 'not-found'
}
