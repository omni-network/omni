import { type GetOrderReturn, getOrder } from '@omni-network/core'
import { type UseQueryResult, useQuery } from '@tanstack/react-query'
import type { Hex } from 'viem'
import { useClient } from 'wagmi'
import { useOmniContracts } from './useOmniContracts.js'

export type UseGetOrderParameters = {
  chainId?: number
  orderId?: Hex
  enabled?: boolean
}

export type UseGetOrderReturn = UseQueryResult<GetOrderReturn>

export function useGetOrder({
  chainId,
  orderId,
  enabled,
}: UseGetOrderParameters): UseGetOrderReturn {
  const client = useClient({ chainId })
  const { data: contracts } = useOmniContracts()
  return useQuery({
    queryKey: ['getOrder', chainId, orderId],
    queryFn: async () => {
      if (!client || !contracts || !orderId) {
        throw new Error('Invalid query parameters')
      }
      return await getOrder({ client, inboxAddress: contracts.inbox, orderId })
    },
    enabled:
      !!client && !!contracts && !!orderId && !!chainId && (enabled ?? true),
    refetchInterval: 1000,
  })
}
