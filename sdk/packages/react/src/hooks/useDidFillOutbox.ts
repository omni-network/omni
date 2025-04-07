import { outboxABI } from '@omni-network/core'
import { useReadContract } from 'wagmi'
import { useOmniContracts } from './useOmniContracts.js'
import type { useParseOpenEvent } from './useParseOpenEvent.js'

type UseDidFillParams = {
  destChainId: number
  resolvedOrder?: ReturnType<typeof useParseOpenEvent>['resolvedOrder']
}

export function useDidFillOutbox(params: UseDidFillParams) {
  const { resolvedOrder, destChainId } = params
  const { data: contracts } = useOmniContracts()
  const filled = useReadContract({
    chainId: destChainId,
    address: contracts?.outbox,
    abi: outboxABI,
    functionName: 'didFill',
    args: resolvedOrder?.fillInstructions[0].originData
      ? [resolvedOrder.orderId, resolvedOrder.fillInstructions[0].originData]
      : undefined,
    query: {
      enabled:
        !!contracts &&
        !!resolvedOrder &&
        !!resolvedOrder.fillInstructions[0].originData,
      refetchInterval: 1000,
    },
  })

  return filled
}
