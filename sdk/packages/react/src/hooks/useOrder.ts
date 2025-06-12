import {
  type DidFillError,
  type EVMOrder,
  GetOrderError,
  type GetRejectionError,
  LoadContractsError,
  type OmniContracts,
  OpenError,
  type OptionalAbis,
  type OrderStatus,
  type ParseOpenEventError,
  type Rejection,
  type SendOrderReturn,
  TxReceiptError,
  ValidateOrderError,
  type WatchDidFillError,
  sendOrder,
} from '@omni-network/core'
import {
  type UseMutateFunction,
  type UseMutationResult,
  useMutation,
} from '@tanstack/react-query'
import type { Hex, WriteContractErrorType } from 'viem'
import {
  type Config,
  type UseWaitForTransactionReceiptReturnType,
  type UseWriteContractReturnType,
  useConfig,
  useWaitForTransactionReceipt,
} from 'wagmi'
import { getConnectorClient } from 'wagmi/actions'
import type { NoClientError } from '../errors/index.js'
import type { QueryOpts } from './types.js'
import { useGetOrderStatus } from './useGetOrderStatus.js'
import {
  type UseOmniContractsResult,
  useOmniContracts,
} from './useOmniContracts.js'
import { useParseOpenEvent } from './useParseOpenEvent.js'
import { useRejection } from './useRejection.js'
import {
  type UseValidateOrderResult,
  useValidateOrder,
} from './useValidateOrder.js'

type UseOrderParams<abis extends OptionalAbis> = EVMOrder<abis> & {
  validateEnabled: boolean
  debugValidation?: boolean
  omniContractsQueryOpts?: QueryOpts<OmniContracts>
  didFillQueryOpts?: QueryOpts<boolean>
  rejectionQueryOpts?: QueryOpts<Rejection, GetRejectionError>
}

type MutationError = LoadContractsError | NoClientError | WriteContractErrorType

export type MutationResult = UseMutationResult<
  SendOrderReturn,
  MutationError,
  void
>

type UseOrderError =
  | OpenError
  | TxReceiptError
  | GetOrderError
  | WatchDidFillError
  | DidFillError
  | ValidateOrderError
  | ParseOpenEventError
  | MutationError
  | undefined

type UseOrderReturnType = {
  open: UseMutateFunction<`0x${string}`, MutationError, void, unknown>
  orderId?: Hex
  validation?: UseValidateOrderResult
  txHash?: Hex
  destTxHash?: Hex
  unwatchDestTx?: () => void
  error?: UseOrderError
  status: UseOrderStatus
  isTxPending: boolean
  isTxSubmitted: boolean
  isValidated: boolean
  isOpen: boolean
  isError: boolean
  isReady: boolean
  txMutation: MutationResult
  waitForTx: UseWaitForTransactionReceiptReturnType<Config, number>
  rejection?: Rejection
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

export function useOrder<abis extends OptionalAbis>(
  params: UseOrderParams<abis>,
): UseOrderReturnType {
  const {
    validateEnabled,
    debugValidation,
    omniContractsQueryOpts,
    didFillQueryOpts,
    rejectionQueryOpts,
    ...order
  } = params
  const config = useConfig()
  const contractsResult = useOmniContracts({
    queryOpts: omniContractsQueryOpts,
  })
  const inboxAddress = contractsResult.data?.inbox

  const txMutation = useMutation<SendOrderReturn, MutationError>({
    mutationFn: async () => {
      const client = await getConnectorClient(config, {
        chainId: order.srcChainId,
      })
      if (inboxAddress == null) {
        throw new LoadContractsError(
          'Inbox contract address needs to be loaded',
        )
      }
      return await sendOrder({ client, inboxAddress, order })
    },
  })

  const wait = useWaitForTransactionReceipt({
    hash: txMutation.data,
    chainId: order.srcChainId,
  })

  const { resolvedOrder, error: parseOpenEventError } = useParseOpenEvent({
    status: wait.status,
    logs: wait.data?.logs,
  })

  const orderStatus = useGetOrderStatus({
    srcChainId: order.srcChainId,
    destChainId: order.destChainId,
    resolvedOrder,
    didFillQueryOpts,
  })

  const rejection = useRejection({
    orderId: resolvedOrder?.orderId,
    fromBlock: wait.data?.blockNumber,
    enabled: wait.status === 'success' && orderStatus.status === 'rejected',
    srcChainId: order.srcChainId,
    queryOpts: rejectionQueryOpts,
  })

  const status = deriveStatus(
    contractsResult,
    orderStatus.status,
    txMutation.status,
    wait.status,
    wait.fetchStatus,
  )

  const validation = useValidateOrder({
    order,
    enabled: validateEnabled,
    debug: debugValidation,
  })

  const error = deriveError({
    contracts: contractsResult,
    txMutation,
    wait,
    validation,
    parseOpenEventError,
    orderStatus,
  })

  return {
    open: txMutation.mutate,
    orderId: resolvedOrder?.orderId,
    validation,
    txHash: txMutation.data,
    destTxHash: orderStatus.destTxHash,
    unwatchDestTx: orderStatus.unwatchDestTx,
    error,
    status,
    isError: !!error,
    isValidated: validation?.status === 'accepted',
    isTxPending: txMutation.isPending,
    isTxSubmitted: txMutation.isSuccess,
    isOpen: !!wait.data && status !== 'closed' && status !== 'rejected',
    isReady: !!inboxAddress,
    txMutation,
    waitForTx: wait,
    rejection: rejection.data ?? undefined,
  }
}

type DeriveErrorParams = {
  contracts: UseOmniContractsResult
  txMutation: MutationResult
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
  orderStatus: OrderStatus,
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
