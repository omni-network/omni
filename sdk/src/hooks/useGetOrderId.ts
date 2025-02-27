import { useMemo } from 'react'
import { type Log, decodeEventLog } from 'viem'
import type { UseWaitForTransactionReceiptReturnType } from 'wagmi'
import { inboxABI } from '../constants/abis.js'

type UseGetOrderDataParams = {
  status: UseWaitForTransactionReceiptReturnType['status']
  logs?: Log[]
}

export function useGetOrderData(params: UseGetOrderDataParams) {
  const { status, logs } = params
  const eventData = useMemo(() => {
    if (!logs || status !== 'success') return
    const openEvent = decodeEventLog({
      abi: inboxABI,
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
