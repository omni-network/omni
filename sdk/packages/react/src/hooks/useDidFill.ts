import { type ResolvedOrder, didFill } from '@omni-network/core'
import { invariant } from '@omni-network/core'
import { type UseQueryResult, useQuery } from '@tanstack/react-query'
import { useClient } from 'wagmi'
import type { QueryOpts } from './types.js'
import { useOmniContracts } from './useOmniContracts.js'

export type UseDidFillParams = {
  destChainId: number
  resolvedOrder?: ResolvedOrder
  queryOpts?: QueryOpts<boolean>
}

export type UseDidFillReturn = UseQueryResult<boolean>

export function useDidFill({
  resolvedOrder,
  destChainId,
  queryOpts,
}: UseDidFillParams): UseDidFillReturn {
  const client = useClient({ chainId: destChainId })
  const { data: contracts } = useOmniContracts()
  const canQuery = !!client && !!contracts && !!resolvedOrder

  return useQuery({
    refetchInterval: 1000,
    ...queryOpts,
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
  })
}
