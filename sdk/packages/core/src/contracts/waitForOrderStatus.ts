import type { GetOrderParameters } from './getOrder.js'
import { type InboxStatus, parseInboxStatus } from './parseInboxStatus.js'
import { watchOrder } from './watchOrder.js'

export type WaitForOrderStatusParameters<Status extends InboxStatus> =
  GetOrderParameters & {
    status: Status | Status[]
    pollingInterval?: number
  }

export function waitForOrderStatus<Status extends InboxStatus>(
  params: WaitForOrderStatusParameters<Status>,
): Promise<Status> {
  const { status, ...watchParams } = params
  const expectedStatus = (
    Array.isArray(status) ? status : [status]
  ) as Array<InboxStatus>
  return new Promise((resolve, reject) => {
    const stop = watchOrder({
      ...watchParams,
      onOrder: (order) => {
        const status = parseInboxStatus({ order })
        if (expectedStatus.includes(status)) {
          stop()
          resolve(status as Status)
        }
      },
      onError: (err) => {
        stop()
        reject(err)
      },
    })
  })
}
