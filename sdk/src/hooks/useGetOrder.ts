import type { Hex, Address } from 'viem'
import { useReadContract } from 'wagmi'
import { inbox } from '../constants/contracts.js'

export function useGetOrder({
  chainId,
  orderId,
  inbox: address,
}: {
  chainId: number
  inbox?: Address
  orderId?: Hex
}) {
  return useReadContract({
    address,
    abi: inbox.abi,
    functionName: 'getOrder',
    chainId,
    args: orderId ? [orderId] : undefined,
    query: {
      enabled: !!orderId && !!address,
      refetchInterval: 1000,
    },
  })
}
