import type { Hex } from 'viem'
import { useReadContract } from 'wagmi'
import { inboxABI } from '../constants/abis.js'
import { useOmniContracts } from './useOmniContracts.js'

export function useGetOrder({
  chainId,
  orderId,
}: {
  chainId?: number
  orderId?: Hex
}) {
  const { data: contracts } = useOmniContracts()
  return useReadContract({
    address: contracts?.inbox,
    abi: inboxABI,
    functionName: 'getOrder',
    chainId,
    args: orderId ? [orderId] : undefined,
    query: {
      enabled: !!contracts && !!orderId && !!chainId,
      refetchInterval: 1000,
    },
  })
}
