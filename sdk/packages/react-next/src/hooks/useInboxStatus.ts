import { type InboxStatus, getInboxStatus } from '@omni-network/core'
import type { Hex } from 'viem'
import { useGetOrder } from './useGetOrder.js'

export function useInboxStatus({
  chainId,
  orderId,
}: {
  chainId?: number
  orderId?: Hex
}): InboxStatus {
  const order = useGetOrder({ chainId, orderId })
  // TODO propagate error if getOrder fails / data not found
  return getInboxStatus(order?.data?.[1].status)
}
