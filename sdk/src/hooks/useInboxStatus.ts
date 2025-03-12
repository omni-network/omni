import type { Hex } from 'viem'
import { useGetOrder } from './useGetOrder.js'

export type InboxStatus = ReturnType<typeof useInboxStatus>

export function useInboxStatus({
  chainId,
  orderId,
}: {
  chainId?: number
  orderId?: Hex
}) {
  const order = useGetOrder({ chainId, orderId })
  // TODO propagate error if getOrder fails / data not found
  const status = order?.data?.[1].status
  if (!isKnown(status)) return 'unknown'
  return strs[status]
}

// strs maps inbox enum uint status to user-friendly strings
// 0: ISolvernetInbox.Status.Invalid
// 1: ISolvernetInbox.Status.Open
// 2: ISolvernetInbox.Status.Rejected
// 3: ISolvernetInbox.Status.Closed
// 4: ISolvernetInbox.Status.Filled
// 5: ISolvernetInbox.Status.Claimed (reported as 'filled' to user)
const strs = {
  0: 'not-found',
  1: 'open',
  2: 'rejected',
  3: 'closed',
  4: 'filled',
  5: 'filled',
} as const

const isKnown = (s?: number): s is keyof typeof strs => s != null && s in strs
