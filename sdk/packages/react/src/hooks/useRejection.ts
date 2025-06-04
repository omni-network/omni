import {
  type GetRejectionError,
  type GetRejectionParams,
  type Rejection,
  getRejection,
  invariant,
} from '@omni-network/core'
import { useQuery } from '@tanstack/react-query'
import { useClient } from 'wagmi'
import { useOmniContracts } from '../index.js'
import { hashFn } from '../utils/query.js'
import type { QueryOpts } from './types.js'

export type UseRejectionParams = Partial<
  Pick<GetRejectionParams, 'orderId' | 'fromBlock'>
> & {
  enabled?: boolean
  srcChainId?: number
  queryOpts?: QueryOpts<Rejection, GetRejectionError>
}

export function useRejection({
  orderId,
  fromBlock,
  srcChainId,
  enabled = true,
  queryOpts,
}: UseRejectionParams) {
  const inbox = useOmniContracts().data?.inbox
  const client = useClient({
    chainId: srcChainId,
  })

  const canQuery = !!inbox && !!client && !!orderId && !!fromBlock && enabled

  const query = useQuery<Rejection, GetRejectionError>({
    ...queryOpts,
    retry: false,
    queryKey: ['reject', { orderId, fromBlock, srcChainId }],
    queryFn: async () => {
      invariant(canQuery)
      return await getRejection({
        client: client,
        orderId: orderId,
        inboxAddress: inbox,
        fromBlock: fromBlock,
      })
    },
    queryKeyHashFn: hashFn,
    enabled: canQuery,
  })

  return query
}
