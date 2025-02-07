import { useMemo } from 'react'
import { type Hex, type Log, decodeEventLog } from 'viem'
import {
  type UseWaitForTransactionReceiptReturnType,
  useReadContract,
} from 'wagmi'
import { inbox, outbox } from '../index.js'
import type { OrderStatus } from '../types/orderStatus.js'

type UseOrderStatusParams = {
  destChainId: number
  txStatus?: UseWaitForTransactionReceiptReturnType['status']
  txLogs?: Log[]
  originChainId?: number
  refetchInterval?: number
}

type UseDidFillParams = {
  destChainId: number
  id?: Hex
  originData?: Hex
  refetchInterval?: number
}

function useDidFill(params: UseDidFillParams) {
  const { id, originData, destChainId, refetchInterval } = params
  const filled = useReadContract({
    chainId: destChainId,
    address: outbox.address,
    abi: outbox.abi,
    functionName: 'didFill',
    args: id && originData ? [id, originData] : undefined,
    query: {
      enabled: !!id && !!originData,
      refetchInterval: refetchInterval ?? 1000,
    },
  })

  return filled.data
}

export function useOrderStatus(params: UseOrderStatusParams) {
  const { txStatus, txLogs, refetchInterval, originChainId } = params
  const eventData = useMemo(() => {
    if (!txStatus || !txLogs || txStatus !== 'success') return
    const openEvent = decodeEventLog({
      abi: inbox.abi,
      eventName: 'Open',
      data: txLogs[txLogs.length - 1].data,
      topics: txLogs[txLogs.length - 1].topics,
    })
    return {
      id: openEvent.args.resolvedOrder.orderId,
      originData: openEvent.args.resolvedOrder.fillInstructions[0].originData,
    }
  }, [txStatus, txLogs])

  const filled = useDidFill({
    id: eventData?.id,
    originData: eventData?.originData,
    refetchInterval,
    ...params,
  })

  const order = useReadContract({
    address: inbox.address,
    abi: inbox.abi,
    functionName: 'getOrder',
    chainId: originChainId,
    args: eventData?.id ? [eventData.id] : undefined,
    query: {
      enabled: !!eventData?.id || !filled,
      refetchInterval: refetchInterval ?? 1000,
    },
  })

  const status: OrderStatus | undefined = useMemo(() => {
    return (
      order?.data &&
      (filled ? 'filled' : parseOrderStatus(order.data[1].status))
    )
  }, [order, filled])

  return status
}

const ORDER_STATUS: Record<number, OrderStatus> = {
  0: 'invalid',
  1: 'pending',
  2: 'accepted',
  3: 'rejected',
  4: 'reverted',
  5: 'filled',
  6: 'claimed',
} as const

function parseOrderStatus(status: number): OrderStatus {
  const orderStatus = ORDER_STATUS[status]
  if (!orderStatus) {
    throw new Error(`Invalid order status: ${status}`)
  }
  return orderStatus
}
