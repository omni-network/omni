import { type ResolvedOrder, didFill } from '@omni-network/core'
import { type UseQueryResult, useQuery } from '@tanstack/react-query'
import { useClient } from 'wagmi'
import { invariant } from '../utils/invariant.js'
import { useOmniContracts } from './useOmniContracts.js'

export type UseDidFillParams = {
  destChainId: number
  resolvedOrder?: ResolvedOrder
}

export type UseDidFillReturn = UseQueryResult<boolean>

export function useDidFill({
  resolvedOrder,
  destChainId,
}: UseDidFillParams): UseDidFillReturn {
  const client = useClient({ chainId: destChainId })
  const { data: contracts } = useOmniContracts()
  const canQuery = !!client && !!contracts && !!resolvedOrder

  return useQuery({
    queryKey: ['didFill', destChainId, resolvedOrder?.orderId],
    queryFn: async () => {
      invariant(canQuery)
      return await didFill({
        client,
        outboxAddress: contracts.outbox,
        resolvedOrder,
      })
    },
    enabled: canQuery,
    refetchInterval: 1000,
  })
}
