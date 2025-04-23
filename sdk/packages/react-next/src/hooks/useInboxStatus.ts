import { type InboxStatus, getInboxStatus } from '@omni-network/core'
import { type UseQueryResult, useQuery } from '@tanstack/react-query'
import type { Hex } from 'viem'
import { useClient } from 'wagmi'
import { invariant } from '../utils/invariant.js'
import { useOmniContracts } from './useOmniContracts.js'

export type UseInboxStatusParams = {
  chainId: number
  orderId?: Hex
}

export type UseInboxStatusReturn = UseQueryResult<InboxStatus>

export function useInboxStatus({
  chainId,
  orderId,
}: UseInboxStatusParams): UseInboxStatusReturn {
  const { data: contracts } = useOmniContracts()
  const client = useClient({ chainId })

  return useQuery({
    queryKey: ['inboxStatus', chainId, orderId],
    queryFn: async () => {
      invariant(!!client && !!contracts && !!orderId)
      return await getInboxStatus({
        client,
        inboxAddress: contracts.inbox,
        orderId,
      })
    },
    enabled: !!client && !!contracts && !!orderId,
    refetchInterval: 1000,
  })
}
