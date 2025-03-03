import { useReadContract } from 'wagmi'
import { outboxABI } from '../constants/abis.js'
import { useOmniContext } from '../context/omni.js'
import type { ResolvedOrder } from '../types/order.js'

type UseDidFillParams = {
  destChainId: number
  resolvedOrder?: ResolvedOrder
}

export function useDidFill(params: UseDidFillParams) {
  const { resolvedOrder, destChainId } = params
  const { outbox } = useOmniContext()
  const filled = useReadContract({
    chainId: destChainId,
    address: outbox,
    abi: outboxABI,
    functionName: 'didFill',
    args: resolvedOrder
      ? [resolvedOrder.orderId, resolvedOrder.fillInstructions[0].originData]
      : undefined,
    query: {
      enabled: !!resolvedOrder,
      refetchInterval: 1000,
    },
  })

  return filled
}
