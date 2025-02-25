import { useQuery } from '@tanstack/react-query'
import { useCallback, useMemo } from 'react'
import { encodeFunctionData } from 'viem'
import type { Hex, WriteContractErrorType } from 'viem'
import {
  type Config,
  type UseWaitForTransactionReceiptReturnType,
  type UseWriteContractReturnType,
  useChainId,
  useWaitForTransactionReceipt,
  useWriteContract,
} from 'wagmi'
import { inboxABI } from '../constants/abis.js'
import { typeHash } from '../constants/typehash.js'
import { useOmniContext } from '../context/omni.js'
import type { Order } from '../types/order.js'
import type { OrderStatus } from '../types/order.js'
import { encodeOrder } from '../utils/encodeOrder.js'
import { useDidFill } from './useDidFill.js'
import { useGetOpenOrder } from './useGetOpenOrder.js'
import { type InboxStatus, useInboxStatus } from './useInboxStatus.js'
import { toJSON } from './util.js'

type UseOrderParams = {
  order: Order
  validateEnabled?: boolean
}

type UseOrderReturnType = {
  open: () => Promise<Hex>
  orderId?: Hex
  validation?: Validation
  txHash?: Hex
  error?: WriteContractErrorType
  validationError?: ValidationError
  status: OrderStatus
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

export function useOrder(params: UseOrderParams): UseOrderReturnType {
  const txMutation = useWriteContract()
  const wait = useWaitForTransactionReceipt({ hash: txMutation.data })

  const { orderId, originData } = useGetOpenOrder({
    status: wait.status,
    logs: wait.data?.logs,
  })

  const connected = useChainId()
  const srcChainId = params.order.srcChainId ?? connected
  const destChainId = params.order.destChainId
  const inboxStatus = useInboxStatus({ orderId, chainId: srcChainId })
  const didFill = useDidFill({ destChainId, orderId, originData })

  const status = deriveStatus(
    inboxStatus,
    didFill,
    txMutation.status,
    wait.status,
    wait.fetchStatus,
  )

  const validate = useValidateOrder(params)

  const validation = useMemo(() => {
    if (validate?.data?.error)
      return {
        error: {
          code: validate.data.error.code,
          message: validate.data.error.message,
        },
      }
    if (validate?.data?.rejected)
      return {
        rejected: true,
        rejectReason: validate.data.rejectReason,
        rejectDescription: validate.data.rejectDescription,
      } as const
    if (validate?.data?.accepted) return { accepted: true } as const
    return
  }, [validate?.data])

  const { inbox } = useOmniContext()

  const open = useCallback(async () => {
    const order = params.order
    if (
      !order.deposit.token ||
      !order.expense.token ||
      !order.deposit.amount ||
      !order.expense.amount
    ) {
      throw new Error('Tokens and amounts are required')
    }

    // we check for overriden deposit and expense params above
    const encoded = encodeOrder(order)
    return await txMutation.writeContractAsync({
      abi: inboxABI,
      address: inbox,
      functionName: 'open',
      chainId: params.order.srcChainId,
      value: params.order.calls.reduce(
        (acc, call) => acc + (call.value ?? 0n),
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
  }, [params, txMutation.writeContractAsync, inbox])

  return {
    open,
    orderId,
    validation,
    txHash: txMutation.data,
    status,
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

// TODO: runtime assertions
function useValidateOrder({ order, validateEnabled = true }: UseOrderParams) {
  // biome-ignore lint/correctness/useExhaustiveDependencies: deep compare on obj properties
  const request = useMemo(() => {
    if (!order.owner) return ''

    function _encodeCalls() {
      return order.calls.map((call) => {
        const callData = encodeFunctionData({
          abi: call.abi,
          functionName: call.functionName,
          args: call.args,
        })
        return {
          target: call.target,
          value: call.value,
          data: callData,
        }
      })
    }

    return toJSON({
      sourceChainId: order.srcChainId,
      destChainId: order.destChainId,
      fillDeadline: order.fillDeadline ?? Math.floor(Date.now() / 1000 + 86400),
      calls: _encodeCalls(),
      deposit: {
        amount: order.deposit.amount,
        token: order.deposit.token,
      },
      expenses: [
        {
          amount: order.expense.amount,
          token: order.expense.token,
          spender: order.expense.spender,
        },
      ],
    })
  }, [
    order.srcChainId,
    order.destChainId,
    order.deposit.amount?.toString(),
    order.deposit.token,
    order.expense.amount?.toString(),
    order.expense.token,
    order.expense.spender,
    order.owner,
    order.fillDeadline,
  ])

  return useQuery<ValidationResponse>({
    queryKey: ['check', request],
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
    enabled:
      !!order.owner &&
      !!order.srcChainId &&
      validateEnabled !== false &&
      !(order.deposit.amount === 0n && order.expense.amount === 0n),
  })
}

// deriveStatus returns a status derived from open tx, inbox and outbox statuses
function deriveStatus(
  inboxStatus: InboxStatus,
  didFill: boolean,
  txStatus: UseWriteContractReturnType['status'],
  receiptStatus: UseWaitForTransactionReceiptReturnType['status'],
  receiptFetchStatus: UseWaitForTransactionReceiptReturnType['fetchStatus'],
): OrderStatus {
  // if outbox says filled, it's filled
  if (didFill) return 'filled'

  // prioritize inbox status over tx status
  if (inboxStatus === 'filled') return 'filled'
  if (inboxStatus === 'open') return 'open'
  if (inboxStatus === 'rejected') return 'rejected'
  if (inboxStatus === 'closed') return 'closed'

  // prioritize receipt status over tx status
  if (receiptStatus === 'error') return 'error'
  if (receiptStatus === 'success') return 'open' // receipt success == open (may be seen before inboxStatus is updated)
  if (receiptFetchStatus === 'idle') return 'idle' // pending is true when !txHash, so we prioritise fetchStatus to check if query is executing
  if (receiptStatus === 'pending') return 'opening'

  // fallback to tx status
  if (txStatus === 'error') return 'error'
  if (txStatus === 'pending') return 'opening'
  if (txStatus === 'success') return 'opening' // still need to wait for receipt

  return 'idle'
}
