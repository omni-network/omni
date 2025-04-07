import { getOrder } from '@omni-network/core'
import { useQuery } from '@tanstack/react-query'
import type { Hex } from 'viem'
import { useClient } from 'wagmi'
import { useOmniContracts } from './useOmniContracts.js'

export function useGetOrder({
  chainId,
  orderId,
  enabled,
}: {
  chainId?: number
  orderId?: Hex
  enabled?: boolean
}) {
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
