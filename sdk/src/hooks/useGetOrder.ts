import type { Hex } from 'viem'
import { useReadContract } from 'wagmi'
import { inboxABI } from '../constants/abis.js'
import { useOmniContext } from '../context/omni.js'

export function useGetOrder({
  chainId,
  orderId,
}: {
  chainId: number
  orderId?: Hex
}) {
  const { inbox } = useOmniContext()
  return useReadContract({
    address: inbox,
    abi: inboxABI,
    functionName: 'getOrder',
    chainId,
    args: orderId ? [orderId] : undefined,
    query: {
      enabled: !!orderId,
      refetchInterval: 1000,
    },
  })
}
