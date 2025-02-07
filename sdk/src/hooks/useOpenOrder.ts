import { useCallback } from 'react'
import { useWaitForTransactionReceipt, useWriteContract } from 'wagmi'
import { type Order, encodeOrder, inbox, typeHash } from '../index.js'
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

  const orderStatus = useOrderStatus({
    txStatus: wait.status,
    txLogs: wait.data?.logs,
    originChainId,
    destChainId,
    refetchInterval: 1000,
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
