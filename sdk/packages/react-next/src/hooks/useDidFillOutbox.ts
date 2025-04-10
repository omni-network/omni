import { type ParseOpenEventReturn, didFillOutbox } from '@omni-network/core'
import { type UseQueryResult, useQuery } from '@tanstack/react-query'
import { useClient } from 'wagmi'
import { useOmniContracts } from './useOmniContracts.js'

export type UseDidFillOutboxParams = {
  destChainId: number
  resolvedOrder?: ParseOpenEventReturn
}

export type UseDidFillOutboxReturn = UseQueryResult<ParseOpenEventReturn>

export function useDidFillOutbox({
  resolvedOrder,
  destChainId,
}: UseDidFillOutboxParams): UseDidFillOutboxReturn {
  const client = useClient({ chainId: destChainId })
  const { data: contracts } = useOmniContracts()
  return useQuery({
    queryKey: ['didFill', destChainId, resolvedOrder?.orderId],
    queryFn: async () => {
      if (!client || !contracts || !resolvedOrder) {
        throw new Error('Invalid query parameters')
      }
      return await didFillOutbox({
        client,
        outboxAddress: contracts.outbox,
        resolvedOrder,
      })
    },
    enabled: !!client && !!contracts && !!resolvedOrder,
    refetchInterval: 1000,
  })
}
