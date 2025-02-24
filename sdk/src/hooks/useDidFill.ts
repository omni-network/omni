import type { Hex } from 'viem'
import { useReadContract } from 'wagmi'
import { outboxABI } from '../constants/abis.js'
import { useOmniContext } from '../context/omni.js'

type UseDidFillParams = {
  destChainId: number
  orderId?: Hex
  originData?: Hex
}

export function useDidFill(params: UseDidFillParams) {
  const { orderId, originData, destChainId } = params
  const { outbox } = useOmniContext()
  const filled = useReadContract({
    chainId: destChainId,
    address: outbox,
    abi: outboxABI,
    functionName: 'didFill',
    args: orderId && originData ? [orderId, originData] : undefined,
    query: {
      enabled: !!orderId && !!originData,
      refetchInterval: 1000,
    },
  })

  return filled.data ?? false
}
