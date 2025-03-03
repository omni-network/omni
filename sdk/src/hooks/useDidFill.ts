import { useReadContract } from 'wagmi'
import { outboxABI } from '../constants/abis.js'
import { useOmniContext } from '../context/omni.js'
import type { useParseOpenEvent } from './useParseOpenEvent.js'

type UseDidFillParams = {
  destChainId: number
  resolvedOrder?: ReturnType<typeof useParseOpenEvent>['resolvedOrder']
}

export function useDidFill(params: UseDidFillParams) {
  const { resolvedOrder, destChainId } = params
  const { outbox } = useOmniContext()
  const filled = useReadContract({
    chainId: destChainId,
    address: outbox,
    abi: outboxABI,
    functionName: 'didFill',
    args: resolvedOrder?.fillInstructions[0].originData
      ? [resolvedOrder.orderId, resolvedOrder.fillInstructions[0].originData]
      : undefined,
    query: {
      enabled:
        !!resolvedOrder && !!resolvedOrder.fillInstructions[0].originData,
      refetchInterval: 1000,
    },
  })

  return filled
}
