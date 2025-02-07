import { useCallback } from 'react'
import { useWaitForTransactionReceipt, useWriteContract } from 'wagmi'
import { type Order, encodeOrder, inbox, typeHash } from '../index.js'
import { useGetOpenOrder } from './useGetOpenOrder.js'
import { useOrderStatus } from './useOrderStatus.js'

type UseOpenOrderParams = {
  destChainId: number
  originChainId?: number
}

export function useOpenOrder(params: UseOpenOrderParams) {
  const { destChainId, originChainId } = params
  const mutation = useWriteContract()
  const wait = useWaitForTransactionReceipt({
    hash: mutation.data,
  })
  const { orderId, originData } = useGetOpenOrder({
    status: wait.status,
    logs: wait.data?.logs,
  })
  const orderStatus = useOrderStatus({
    originChainId,
    destChainId,
    orderId,
    originData,
  })

  const openOrderAsync = useCallback(
    async (order: Order, fillDeadline: number) => {
      const encodedOrder = encodeOrder(order)
      const value = order.calls.reduce(
        (acc, call) => acc + call.value,
        BigInt(0),
      )
      return await mutation.writeContractAsync({
        abi: inbox.abi,
        address: inbox.address,
        functionName: 'open',
        chainId: originChainId,
        value,
        args: [
          {
            fillDeadline: fillDeadline,
            orderDataType: typeHash,
            orderData: encodedOrder,
          },
        ],
      })
    },
    [mutation.writeContractAsync, originChainId],
  )

  return {
    wait,
    mutation,
    orderStatus,
    openOrderAsync,
  }
}
