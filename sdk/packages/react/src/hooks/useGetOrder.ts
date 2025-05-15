import { type GetOrderReturn, getOrder } from '@omni-network/core'
import { type UseQueryResult, useQuery } from '@tanstack/react-query'
import type { Hex } from 'viem'
import { useClient } from 'wagmi'
import { invariant } from '../utils/invariant.js'
import type { QueryOpts } from './types.js'
import { useOmniContracts } from './useOmniContracts.js'

export type UseGetOrderParameters = {
  chainId?: number
  orderId?: Hex
  enabled?: boolean
  queryOpts?: QueryOpts<GetOrderReturn>
}

export type UseGetOrderReturn = UseQueryResult<GetOrderReturn>

export function useGetOrder({
  chainId,
  orderId,
  enabled,
  queryOpts,
}: UseGetOrderParameters): UseGetOrderReturn {
  const client = useClient({ chainId })
  const { data: contracts } = useOmniContracts()
  const canQuery =
    !!client && !!contracts && !!orderId && !!chainId && (enabled ?? true)

  return useQuery({
    refetchInterval: 1000,
    ...queryOpts,
    queryKey: ['getOrder', chainId, orderId],
    queryFn: async () => {
      invariant(canQuery)
      return await getOrder({ client, inboxAddress: contracts.inbox, orderId })
    },
    enabled: canQuery,
  })
}
