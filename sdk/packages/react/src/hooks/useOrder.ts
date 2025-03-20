import type { Hex } from 'viem'
import { zeroAddress } from 'viem'
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
import {
  type DidFillError,
  GetOrderError,
  LoadContractsError,
  OpenError,
  type ParseOpenEventError,
  TxReceiptError,
  ValidateOrderError,
} from '../errors/base.js'
import type { OptionalAbis } from '../types/abi.js'
import type { Order, OrderStatus } from '../types/order.js'
import { encodeOrder } from '../utils/encodeOrder.js'
import { useGetOrderStatus } from './useGetOrderStatus.js'
import {
  type UseOmniContractsResult,
  useOmniContracts,
} from './useOmniContracts.js'
import { useParseOpenEvent } from './useParseOpenEvent.js'
import {
  type UseValidateOrderResult,
  useValidateOrder,
} from './useValidateOrder.js'
type UseOrderParams<abis extends OptionalAbis> = Order<abis> & {
  validateEnabled: boolean
}

type UseOrderError =
  | OpenError
  | TxReceiptError
  | GetOrderError
  | DidFillError
  | ValidateOrderError
  | ParseOpenEventError
  | LoadContractsError
  | undefined

type UseOrderReturnType = {
  open: () => Promise<Hex>
  orderId?: Hex
  validation?: UseValidateOrderResult
  txHash?: Hex
  error?: UseOrderError
  status: UseOrderStatus
  isTxPending: boolean
  isTxSubmitted: boolean
  isValidated: boolean
  isOpen: boolean
  isError: boolean
  isReady: boolean
  txMutation: UseWriteContractReturnType<Config, unknown>
  waitForTx: UseWaitForTransactionReceiptReturnType<Config, number>
}

type UseOrderStatus =
  | 'initializing'
  | 'ready'
  | 'opening'
  | 'open'
  | 'closed'
  | 'rejected'
  | 'error'
  | 'filled'

const defaultFillDeadline = () => Math.floor(Date.now() / 1000 + 86400)

export function useOrder<abis extends OptionalAbis>(
  params: UseOrderParams<abis>,
): UseOrderReturnType {
  const { validateEnabled, ...order } = params
  const connected = useChainId()
  const txMutation = useWriteContract()

  const wait = useWaitForTransactionReceipt({
    hash: txMutation.data,
    chainId: order.srcChainId,
  })

  const { resolvedOrder, error: parseOpenEventError } = useParseOpenEvent({
    status: wait.status,
    logs: wait.data?.logs,
  })

  const orderStatus = useGetOrderStatus({
    srcChainId: order.srcChainId ?? connected,
    destChainId: order.destChainId,
    orderId: resolvedOrder?.orderId,
    resolvedOrder,
  })
  const contractsResult = useOmniContracts()
  const inboxAddress = contractsResult.data?.inbox

  const status = deriveStatus(
    contractsResult,
    orderStatus.status,
    txMutation.status,
    wait.status,
    wait.fetchStatus,
  )

  const validation = useValidateOrder({ order, enabled: validateEnabled })

  const open = inboxAddress
    ? async () => {
        const encoded = encodeOrder(order)

        const isNativeDeposit =
          order.deposit.token == null || order.deposit.token === zeroAddress

        return await txMutation.writeContractAsync({
          abi: inboxABI,
          address: inboxAddress,
          functionName: 'open',
          chainId: order.srcChainId,
          value: isNativeDeposit ? order.deposit.amount : 0n,
          args: [
            {
              fillDeadline: order.fillDeadline ?? defaultFillDeadline(),
              orderDataType: typeHash,
              orderData: encoded,
            },
          ],
        })
      }
    : () => {
        return Promise.reject(
          new LoadContractsError('Inbox contract address needs to be loaded'),
        )
      }

  const error = deriveError({
    contracts: contractsResult,
    txMutation,
    wait,
    validation,
    parseOpenEventError,
    orderStatus,
  })

  return {
    open,
    orderId: resolvedOrder?.orderId,
    validation,
    txHash: txMutation.data,
    error,
    status,
    isError: !!error,
    isValidated: validation?.status === 'accepted',
    isTxPending: txMutation.isPending,
    isTxSubmitted: txMutation.isSuccess,
    isOpen: !!wait.data,
    isReady: !!inboxAddress,
    txMutation,
    waitForTx: wait,
  }
}

type DeriveErrorParams = {
  contracts: UseOmniContractsResult
  txMutation: UseWriteContractReturnType<Config, unknown>
  wait: UseWaitForTransactionReceiptReturnType
  validation: ReturnType<typeof useValidateOrder>
  orderStatus: ReturnType<typeof useGetOrderStatus>
  parseOpenEventError?: ParseOpenEventError
}

function deriveError(params: DeriveErrorParams): UseOrderError {
  const {
    contracts,
    txMutation,
    wait,
    validation,
    orderStatus,
    parseOpenEventError,
  } = params

  if (contracts.error) {
    return new LoadContractsError(contracts.error.message)
  }

  if (validation.status === 'error') {
    return new ValidateOrderError(validation.error.message)
  }

  if (txMutation.error) {
    return new OpenError(txMutation.error.message)
  }

  if (wait.error) {
    return new TxReceiptError(wait.error.message)
  }

  if (parseOpenEventError) {
    return parseOpenEventError
  }

  if (orderStatus.error) {
    return orderStatus.error
  }

  if (wait.isSuccess && orderStatus.status === 'not-found') {
    return new GetOrderError(orderStatus.status)
  }

  return
}

// deriveStatus returns a status derived from open tx, inbox and outbox statuses
function deriveStatus(
  contracts: UseOmniContractsResult,
  orderStatus: OrderStatus, // TODO rename
  txStatus: UseWriteContractReturnType['status'],
  receiptStatus: UseWaitForTransactionReceiptReturnType['status'],
  receiptFetchStatus: UseWaitForTransactionReceiptReturnType['fetchStatus'],
): UseOrderStatus {
  // inbox contract address needs to be loaded
  if (contracts.isError) return 'error'
  if (!contracts.data) return 'initializing'

  // prioritize on chain status over tx status
  if (orderStatus === 'filled') return 'filled'
  if (orderStatus === 'open') return 'open'
  if (orderStatus === 'rejected') return 'rejected'
  if (orderStatus === 'closed') return 'closed'
  if (orderStatus === 'error') return 'error'

  // prioritize receipt status over tx status
  if (receiptStatus === 'error') return 'error'
  if (receiptStatus === 'success') return 'open' // receipt success == open (may be seen before inboxStatus is updated)
  if (receiptFetchStatus === 'idle') return 'ready' // pending is true when !txHash, so we prioritise fetchStatus to check if query is executing
  if (receiptStatus === 'pending') return 'opening'

  // fallback to tx status
  if (txStatus === 'error') return 'error'
  if (txStatus === 'pending') return 'opening'
  if (txStatus === 'success') return 'opening' // still need to wait for receipt

  return 'ready'
}
