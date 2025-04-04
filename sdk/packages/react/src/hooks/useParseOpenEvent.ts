import { useMemo } from 'react'
import { type Log, decodeEventLog, parseEventLogs } from 'viem'
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
      const parsed = parseEventLogs({
        abi: inboxABI,
        logs,
        eventName: 'Open',
      })

      if (parsed.length !== 1) {
        return {
          error: new ParseOpenEventError(
            `Expected exactly one 'Open' event but found ${parsed.length}.`,
          ),
        }
      }

      const openLog = parsed[0]

      const openEvent = decodeEventLog({
        abi: inboxABI,
        eventName: 'Open',
        data: openLog.data,
        topics: openLog.topics,
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
