import { useMemo } from 'react'
import { type Log, decodeEventLog } from 'viem'
import type { UseWaitForTransactionReceiptReturnType } from 'wagmi'
import { inboxABI } from '../constants/abis.js'
import { ParseOpenEventError } from '../errors/base.js'

type UseParseOpenEventParams = {
  status: UseWaitForTransactionReceiptReturnType['status']
  logs?: Log[]
}

export function useParseOpenEvent(params: UseParseOpenEventParams) {
  const { status, logs } = params
  const eventData = useMemo(() => {
    if (!logs || status !== 'success') return
    try {
      const openEvent = decodeEventLog({
        abi: inboxABI,
        eventName: 'Open',
        data: logs[logs.length - 1].data,
        topics: logs[logs.length - 1].topics,
      })
      return {
        resolvedOrder: openEvent.args.resolvedOrder,
      }
    } catch (error) {
      return {
        error: new ParseOpenEventError(`Failed to parse open event: ${error}`),
      }
    }
  }, [status, logs])

  return {
    resolvedOrder: eventData?.resolvedOrder,
    error: eventData?.error,
  }
}
