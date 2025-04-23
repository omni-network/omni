import type { Client, Hex } from 'viem'
import type { Address } from 'viem'
import { readContract } from 'viem/actions'
import { inboxABI } from '../constants/abis.js'

export type InboxStatus = (typeof strs)[keyof typeof strs] | 'unknown'

export type GetInboxStatusParameters = {
  client: Client
  inboxAddress: Address
  orderId: Hex
}

export type GetInboxStatusReturn = InboxStatus

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

export async function getInboxStatus({
  client,
  inboxAddress,
  orderId,
}: GetInboxStatusParameters): Promise<GetInboxStatusReturn> {
  const order = await readContract(client, {
    address: inboxAddress,
    abi: inboxABI,
    functionName: 'getOrder',
    args: [orderId],
  })

  const status = order?.[1].status

  if (!isKnown(status)) return 'unknown'

  return strs[status]
}
