import { watchBlockNumber } from 'viem/actions'
import { AbortError } from '../errors/base.js'
import { type GetOrderParameters, getOrder } from './getOrder.js'
import { type InboxStatus, parseInboxStatus } from './parseInboxStatus.js'

const terminalStatus = [
  'closed',
  'filled',
  'rejected',
] as const satisfies InboxStatus[]

export type TerminalStatus = (typeof terminalStatus)[number]

export type WaitForOrderCloseParameters = GetOrderParameters & {
  pollingInterval?: number
  signal?: AbortSignal
}

export function waitForOrderClose(
  params: WaitForOrderCloseParameters,
): Promise<TerminalStatus> {
  const { pollingInterval, signal, ...getOrderParams } = params
  if (signal?.aborted) {
    return Promise.reject(new AbortError('Aborted'))
  }

  return new Promise((resolve, reject) => {
    const onAbort = () => {
      stop()
      reject(new AbortError('Aborted'))
    }
    signal?.addEventListener('abort', onAbort, { once: true })

    const stop = watchBlockNumber(params.client, {
      onBlockNumber: async () => {
        const order = await getOrder(getOrderParams)
        const status = parseInboxStatus({ order }) as TerminalStatus
        if (terminalStatus.includes(status)) {
          stop()
          signal?.removeEventListener('abort', onAbort)
          resolve(status)
        }
      },
      onError: (err) => {
        stop()
        signal?.removeEventListener('abort', onAbort)
        reject(err)
      },
    })
  })
}
