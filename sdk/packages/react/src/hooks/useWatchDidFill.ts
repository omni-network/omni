import { watchDidFill } from '@omni-network/core'
import { useCallback, useEffect, useRef, useState } from 'react'
import type { Hex } from 'viem'
import { useClient } from 'wagmi'
import { useOmniContracts } from './useOmniContracts.js'

export type UseWatchDidFillProps = {
  destChainId: number
  orderId?: Hex
  pollingInterval?: number
  onError?: (error: Error) => void
}

export type UseWatchDidFillReturn = {
  unwatch: () => void
  status: 'idle' | 'pending' | 'success'
  destTxHash?: Hex
}

export function useWatchDidFill({
  destChainId,
  orderId,
  pollingInterval,
  onError,
}: UseWatchDidFillProps): UseWatchDidFillReturn {
  const unwatchRef = useRef<(() => void) | undefined>()
  const [destTxHash, setDestTxHash] = useState<Hex | undefined>()
  const client = useClient({ chainId: destChainId })
  const outboxAddress = useOmniContracts().data?.outbox
  const status = destTxHash
    ? 'success'
    : unwatchRef.current
      ? 'pending'
      : 'idle'

  const unwatch = useCallback(() => {
    unwatchRef.current?.()
    unwatchRef.current = undefined
  }, [])

  useEffect(() => {
    unwatch()

    if (!client || !orderId || !outboxAddress) return
    setDestTxHash(undefined)

    unwatchRef.current = watchDidFill({
      client,
      outboxAddress,
      orderId,
      onLogs: (logs) => {
        setDestTxHash(logs[0].transactionHash ?? undefined)
        unwatch()
      },
      pollingInterval,
      onError,
    })

    return unwatch
  }, [client, outboxAddress, orderId, pollingInterval, onError, unwatch])

  return { destTxHash, unwatch, status }
}
