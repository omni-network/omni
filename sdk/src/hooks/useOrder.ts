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
  DidFillError,
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
import { useDidFill } from './useDidFill.js'
import { type InboxStatus, useInboxStatus } from './useInboxStatus.js'
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
  status: OrderStatus
  isTxPending: boolean
  isTxSubmitted: boolean
  isValidated: boolean
  isOpen: boolean
  isError: boolean
  isReady: boolean
  txMutation: UseWriteContractReturnType<Config, unknown>
  waitForTx: UseWaitForTransactionReceiptReturnType<Config, number>
}

const defaultFillDeadline = () => Math.floor(Date.now() / 1000 + 86400)

export function useOrder<abis extends OptionalAbis>(
  params: UseOrderParams<abis>,
): UseOrderReturnType {
  const { validateEnabled, ...order } = params
  const txMutation = useWriteContract()
  const wait = useWaitForTransactionReceipt({ hash: txMutation.data })
  const { resolvedOrder, error: parseOpenEventError } = useParseOpenEvent({
    status: wait.status,
    logs: wait.data?.logs,
  })

  const connected = useChainId()
  const srcChainId = order.srcChainId ?? connected
  const destChainId = order.destChainId
  const inboxStatus = useInboxStatus({
    orderId: resolvedOrder?.orderId,
    chainId: srcChainId,
  })
  const didFill = useDidFill({ destChainId, resolvedOrder })

  const contractsResult = useOmniContracts()
  const inboxAddress = contractsResult.data?.inbox

  const status = deriveStatus(
    contractsResult,
    inboxStatus,
    didFill.data ?? false,
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
    didFill,
    validation,
    inboxStatus,
    parseOpenEventError,
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
  didFill: ReturnType<typeof useDidFill>
  validation: ReturnType<typeof useValidateOrder>
  inboxStatus: InboxStatus
  parseOpenEventError?: ParseOpenEventError
}

function deriveError(params: DeriveErrorParams): UseOrderError {
  const { contracts, txMutation, wait, didFill, validation, inboxStatus } =
    params

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

  if (params.parseOpenEventError) {
    return params.parseOpenEventError
  }

  if (didFill.error) {
    return new DidFillError(didFill.error.message)
  }

  if (wait.isSuccess && inboxStatus === 'not-found') {
    return new GetOrderError(inboxStatus)
  }

  return
}

// deriveStatus returns a status derived from open tx, inbox and outbox statuses
function deriveStatus(
  contracts: UseOmniContractsResult,
  inboxStatus: InboxStatus,
  didFill: boolean,
  txStatus: UseWriteContractReturnType['status'],
  receiptStatus: UseWaitForTransactionReceiptReturnType['status'],
  receiptFetchStatus: UseWaitForTransactionReceiptReturnType['fetchStatus'],
): OrderStatus {
  // inbox contract address needs to be loaded
  if (contracts.isError) return 'error'
  if (!contracts.data) return 'initializing'

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
  if (receiptFetchStatus === 'idle') return 'ready' // pending is true when !txHash, so we prioritise fetchStatus to check if query is executing
  if (receiptStatus === 'pending') return 'opening'

  // fallback to tx status
  if (txStatus === 'error') return 'error'
  if (txStatus === 'pending') return 'opening'
  if (txStatus === 'success') return 'opening' // still need to wait for receipt

  return 'ready'
}
