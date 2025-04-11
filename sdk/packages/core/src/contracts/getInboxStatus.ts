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

type ValidStatusCode = keyof typeof strs

export type InboxStatus = (typeof strs)[ValidStatusCode] | 'unknown'

export function getInboxStatus(statusCode?: number): InboxStatus {
  return strs[statusCode as ValidStatusCode] || 'unknown'
}
