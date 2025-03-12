import type { Log } from 'viem'
import type { UseWaitForTransactionReceiptReturnType } from 'wagmi'
import { DidFillError, type ParseOpenEventError } from '../errors/base.js'
import type { OrderStatus } from '../types/order.js'
import { useDidFillOutbox } from './useDidFillOutbox.js'
import { type InboxStatus, useInboxStatus } from './useInboxStatus.js'
import { useParseOpenEvent } from './useParseOpenEvent.js'

export function useOrderStatus({
  srcChainId,
  destChainId,
  waitTx,
  logs,
}: {
  srcChainId?: number
  destChainId: number
  waitTx: UseWaitForTransactionReceiptReturnType
  logs?: Log[]
}) {
  const { resolvedOrder, error: parseOpenEventError } = useParseOpenEvent({
    status: waitTx.status,
    logs,
  })
  const inboxStatus = useInboxStatus({
    orderId: resolvedOrder?.orderId,
    chainId: srcChainId,
  })
  const didFillOutbox = useDidFillOutbox({
    destChainId,
    resolvedOrder,
  })

  const status = deriveStatus(inboxStatus, parseOpenEventError, didFillOutbox)

  const error = deriveError(parseOpenEventError, didFillOutbox)

  return {
    orderId: resolvedOrder?.orderId,
    status,
    error,
  }
}

function deriveError(
  parseOpenEventError: ParseOpenEventError | undefined,
  didFillOutbox: ReturnType<typeof useDidFillOutbox>,
) {
  if (parseOpenEventError) return parseOpenEventError
  if (didFillOutbox.error) return new DidFillError(didFillOutbox.error.message)
  return
}

function deriveStatus(
  inboxStatus: InboxStatus,
  parseOpenEventError: ParseOpenEventError | undefined,
  didFillOutbox: ReturnType<typeof useDidFillOutbox>,
): OrderStatus {
  if (parseOpenEventError || didFillOutbox.error) return 'error'
  if (didFillOutbox?.data === true) return 'filled'
  if (inboxStatus === 'filled') return 'filled'
  if (inboxStatus === 'open') return 'open'
  if (inboxStatus === 'rejected') return 'rejected'
  if (inboxStatus === 'closed') return 'closed'
  return 'not-found'
}
