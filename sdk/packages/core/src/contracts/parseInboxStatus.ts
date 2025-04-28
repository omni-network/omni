import type { GetOrderReturn } from './getOrder.js'

export type InboxStatus = (typeof status)[keyof typeof status] | 'unknown'

export type ParseInboxStatusParameters = {
  order?: GetOrderReturn
}

export type ParseInboxStatusReturn = InboxStatus

// strs maps inbox enum uint status to user-friendly strings
// 0: ISolvernetInbox.Status.Invalid
// 1: ISolvernetInbox.Status.Open
// 2: ISolvernetInbox.Status.Rejected
// 3: ISolvernetInbox.Status.Closed
// 4: ISolvernetInbox.Status.Filled
// 5: ISolvernetInbox.Status.Claimed (reported as 'filled' to user)
export const status = {
  0: 'not-found',
  1: 'open',
  2: 'rejected',
  3: 'closed',
  4: 'filled',
  5: 'filled',
} as const

const isKnown = (s?: number): s is keyof typeof status =>
  s != null && s in status

export function parseInboxStatus({
  order,
}: ParseInboxStatusParameters): ParseInboxStatusReturn {
  const res = order?.[1].status

  if (!isKnown(res)) return 'unknown'

  return status[res]
}
