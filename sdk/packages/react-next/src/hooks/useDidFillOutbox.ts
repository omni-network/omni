import { type ResolvedOrder, didFillOutbox } from '@omni-network/core'
import { type UseQueryResult, useQuery } from '@tanstack/react-query'
import { useClient } from 'wagmi'
import { invariant } from '../utils/invariant.js'
import { useOmniContracts } from './useOmniContracts.js'

export type UseDidFillOutboxParams = {
  destChainId: number
  resolvedOrder?: ResolvedOrder
}

export type UseDidFillOutboxReturn = UseQueryResult<boolean>

export function useDidFillOutbox({
  resolvedOrder,
  destChainId,
}: UseDidFillOutboxParams): UseDidFillOutboxReturn {
  const client = useClient({ chainId: destChainId })
  const { data: contracts } = useOmniContracts()
  const canQuery = !!client && !!contracts && !!resolvedOrder

  return useQuery({
    queryKey: ['didFill', destChainId, resolvedOrder?.orderId],
    queryFn: async () => {
      invariant(canQuery)
      return await didFillOutbox({
        client,
        outboxAddress: contracts.outbox,
        resolvedOrder,
      })
    },
    enabled: canQuery,
    refetchInterval: 1000,
  })
}
