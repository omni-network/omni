import { type ResolvedOrder, watchDidFill } from '@omni-network/core'
import { useCallback, useEffect, useRef, useState } from 'react'
import type { Hex } from 'viem'
import { useClient } from 'wagmi'
import { useOmniContracts } from './useOmniContracts.js'

export type UseWatchDidFillParams = {
  destChainId: number
  resolvedOrder?: ResolvedOrder
  pollingInterval?: number
}

export type UseWatchDidFillReturn = {
  unwatch: () => void
  status: 'idle' | 'pending' | 'success' | 'error'
  destTxHash?: Hex
  error?: Error
}

export function useWatchDidFill({
  destChainId,
  resolvedOrder,
  pollingInterval,
}: UseWatchDidFillParams): UseWatchDidFillReturn {
  const unwatchRef = useRef<(() => void) | undefined>()
  const [destTxHash, setDestTxHash] = useState<Hex | undefined>()
  const [error, setError] = useState<Error | undefined>()
  const client = useClient({ chainId: destChainId })
  const outboxAddress = useOmniContracts().data?.outbox
  const status = destTxHash
    ? 'success'
    : error
      ? 'error'
      : unwatchRef.current
        ? 'pending'
        : 'idle'

  const unwatch = useCallback(() => {
    unwatchRef.current?.()
    unwatchRef.current = undefined
  }, [])

  useEffect(() => {
    unwatch()
    setError(undefined)
    setDestTxHash(undefined)

    if (!client || !resolvedOrder || !outboxAddress) return

    unwatchRef.current = watchDidFill({
      client,
      outboxAddress,
      resolvedOrder,
      onFill: (txHash) => {
        setDestTxHash(txHash)
        unwatch()
      },
      onError: (error) => {
        setError(error)
        unwatch()
      },
      pollingInterval,
    })

    return unwatch
  }, [client, outboxAddress, resolvedOrder, pollingInterval, unwatch])

  return { destTxHash, unwatch, status, error }
}
