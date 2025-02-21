import { useQuery } from '@tanstack/react-query'
import { useCallback, useMemo } from 'react'
import { encodeFunctionData, slice, zeroAddress } from 'viem'
import type { Hex, WriteContractErrorType } from 'viem'
import {
  type Config,
  type UseWaitForTransactionReceiptReturnType,
  type UseWriteContractReturnType,
  useWaitForTransactionReceipt,
  useWriteContract,
} from 'wagmi'
import { inbox } from '../constants/contracts.js'
import { typeHash } from '../constants/typehash.js'
import type { Order, OrderStatus } from '../types/order.js'
import { encodeOrder } from '../utils/encodeOrder.js'
import { useGetOpenOrder } from './useGetOpenOrder.js'
import { useOrderStatus } from './useOrderStatus.js'

type UseOrderParams = {
  order: Order
  validateEnabled?: boolean
}

type UseOrderReturnType = {
  open: () => Promise<Hex>
  validation?: Validation
  txHash?: Hex
  error?: WriteContractErrorType
  validationError?: ValidationError
  orderStatus?: OrderStatus
  canOpen: boolean
  isTxPending: boolean
  isTxSubmitted: boolean
  isValidated: boolean
  isOpen: boolean
  isRejected: boolean
  isError: boolean
  txMutation: UseWriteContractReturnType<Config, unknown>
  waitForTx: UseWaitForTransactionReceiptReturnType<Config, number>
}

type ValidationRejected = {
  rejected: true
  rejectReason?: string
  rejectDescription?: string
}

type ValidationAccepted = {
  accepted: true
}

type ValidationError = {
  error: {
    code: number
    message: string
  }
}

type Validation = ValidationRejected | ValidationAccepted | ValidationError

export function useOrder(params: UseOrderParams): UseOrderReturnType {
  const txMutation = useWriteContract()
  const wait = useWaitForTransactionReceipt({
    hash: txMutation.data,
  })
  const { orderId, originData } = useGetOpenOrder({
    status: wait.status,
    logs: wait.data?.logs,
  })
  const orderStatus = useOrderStatus({
    orderId: orderId,
    originData: originData,
    ...params.order,
  })
  const validate = useValidateOrder(
    params.order,
    params.validateEnabled ?? true,
  )

  const validation = useMemo(() => {
    if (validate.data?.error)
      return {
        error: {
          code: validate.data.error.code,
          message: validate.data.error.message,
        },
      }
    if (validate.data?.rejected)
      return {
        rejected: true,
        rejectReason: validate.data.rejectReason,
        rejectDescription: validate.data.rejectDescription,
      } as const
    if (validate.data?.accepted) return { accepted: true } as const
    return
  }, [validate.data])

  const open = useCallback(async () => {
    const encoded = encodeOrder(params.order)
    return await txMutation.writeContractAsync({
      ...inbox,
      functionName: 'open',
      chainId: params.order.srcChainId,
      value: params.order.calls.reduce(
        (acc, call) => acc + call.value,
        BigInt(0),
      ),
      args: [
        {
          fillDeadline:
            params.order.fillDeadline ?? Math.floor(Date.now() / 1000 + 86400),
          orderDataType: typeHash,
          orderData: encoded,
        },
      ],
    })
  }, [params, txMutation.writeContractAsync])

  return {
    open,
    validation,
    txHash: txMutation.data,
    orderStatus,
    error: txMutation.error as WriteContractErrorType | undefined,
    canOpen: validation?.accepted ?? false,
    isTxPending: txMutation.isPending,
    isTxSubmitted: txMutation.isSuccess,
    isValidated: validation?.accepted ?? false,
    isRejected: validation?.rejected ?? false,
    isError: !!(validation?.error || txMutation.error),
    isOpen: !!wait.data,
    txMutation,
    waitForTx: wait,
  }
}

////////////////////////
//// order validation
////////////////////////
type ValidationResponse = {
  accepted?: boolean
  rejected?: boolean
  error?: {
    code: number
    message: string
  }
  rejectReason?: string
  rejectDescription?: string
}

// TODO: runtime assertions?
function useValidateOrder(order: Order, enabled: boolean) {
  const calls = order.calls.map((call) => {
    const callData = encodeFunctionData({
      abi: call.abi,
      functionName: call.functionName,
      args: call.args,
    })
    return {
      target: call.target,
      selector: slice(callData, 0, 4),
      value: call.value,
      params: callData.length > 10 ? slice(callData, 4) : '0x',
    }
  })

  const expense = {
    amount: order.expense.amount,
    spender: order.expense.spender,
    token: order.expense.isNative ? zeroAddress : order.expense.token,
  }
  const deposit = {
    amount: order.deposit.amount,
    token: order.deposit.isNative ? zeroAddress : order.deposit.token,
  }

  const request = JSON.stringify({
    sourceChainId: order.srcChainId,
    destChainId: order.destChainId,
    fillDeadline: order.fillDeadline ?? Math.floor(Date.now() / 1000 + 86400),
    calls: calls,
    expenses: [expense],
    deposit,
  })

  return useQuery<ValidationResponse>({
    queryKey: ['check'],
    queryFn: async () => {
      // TODO remove hardcoded api url
      const response = await fetch(
        'https://solver.staging.omni.network/api/v1/check',
        {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json',
          },
          body: request,
        },
      )
      return await response.json()
    },
    enabled,
  })
}
