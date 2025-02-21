import { useMemo } from 'react'
import { type Log, decodeEventLog } from 'viem'
import type { UseWaitForTransactionReceiptReturnType } from 'wagmi'
import { inbox } from '../constants/contracts.js'

type UseGetOpenOrderParams = {
  status: UseWaitForTransactionReceiptReturnType['status']
  logs?: Log[]
}

export function useGetOpenOrder(params: UseGetOpenOrderParams) {
  const { status, logs } = params
  const eventData = useMemo(() => {
    if (!logs || status !== 'success') return
    const openEvent = decodeEventLog({
      abi: inbox.abi,
      eventName: 'Open',
      data: logs[logs.length - 1].data,
      topics: logs[logs.length - 1].topics,
    })
    return {
      id: openEvent.args.resolvedOrder.orderId,
      originData: openEvent.args.resolvedOrder.fillInstructions[0].originData,
    }
  }, [status, logs])

  return {
    orderId: eventData?.id,
    originData: eventData?.originData,
  }
}
