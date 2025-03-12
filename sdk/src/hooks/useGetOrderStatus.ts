import type { Hex } from 'viem'
import { DidFillError } from '../errors/base.js'
import type { OrderStatus } from '../types/order.js'
import { useDidFillOutbox } from './useDidFillOutbox.js'
import { type InboxStatus, useInboxStatus } from './useInboxStatus.js'
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
  const inboxStatus = useInboxStatus({
    orderId: resolvedOrder?.orderId ?? orderId,
    chainId: srcChainId,
  })
  const didFillOutbox = useDidFillOutbox({
    destChainId,
    resolvedOrder,
  })

  const status = deriveStatus(inboxStatus, didFillOutbox)

  return {
    status,
    error: deriveError(didFillOutbox),
  }
}

function deriveError(didFillOutbox: ReturnType<typeof useDidFillOutbox>) {
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
