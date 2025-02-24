import type { Hex, Address } from 'viem'
import { useGetOrder } from './useGetOrder.js'

export function useInboxStatus({
  chainId,
  orderId,
  inbox,
}: {
  chainId: number
  inbox?: Address
  orderId?: Hex
}) {
  const order = useGetOrder({ chainId, orderId, inbox })
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
