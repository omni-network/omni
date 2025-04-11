import {
  type ParseOpenEventError,
  type ResolvedOrder,
  parseOpenEvent,
} from '@omni-network/core'
import { useMemo } from 'react'
import type { Log } from 'viem'
import type { UseWaitForTransactionReceiptReturnType } from 'wagmi'

type UseParseOpenEventParams = {
  status: UseWaitForTransactionReceiptReturnType['status']
  logs?: Log[]
}

type UseParseOpenEventReturn = {
  resolvedOrder: ResolvedOrder | undefined
  error: ParseOpenEventError | undefined
}

export function useParseOpenEvent(
  params: UseParseOpenEventParams,
): UseParseOpenEventReturn {
  const { status, logs } = params
  const eventData = useMemo(() => {
    if (!logs || status !== 'success') return
    try {
      return { resolvedOrder: parseOpenEvent(logs) }
    } catch (error) {
      return { error: error as ParseOpenEventError }
    }
  }, [status, logs])

  return {
    resolvedOrder: eventData?.resolvedOrder,
    error: eventData?.error,
  }
}
