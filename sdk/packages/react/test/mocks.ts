import { beforeEach, vi } from 'vitest'
import type {
  useReadContract,
  useWaitForTransactionReceipt,
  useWriteContract,
} from 'wagmi'
import * as apiModule from '../src/internal/api.js'
import { contracts } from './shared.js'

type UseReadContractReturn<Data> = Omit<
  ReturnType<typeof useReadContract>,
  'error' | 'data'
> & {
  error: Error | null
  data: Data
}

type UseWriteContractReturn<Data> = Omit<
  ReturnType<typeof useWriteContract>,
  'error' | 'data'
> & {
  error: Error | null
  data: Data
}

type UseWaitForTransactionReceiptReturn<Data> = Omit<
  ReturnType<typeof useWaitForTransactionReceipt>,
  'error' | 'data'
> & {
  error: Error | null
  data: Data
}

export function mockContractsQuery(failure = false) {
  vi.spyOn(apiModule, 'fetchJSON').mockImplementation((url: string) => {
    if (failure) {
      return Promise.reject(new Error('mock error'))
    }

    if (url.includes('/contracts')) {
      return Promise.resolve(contracts)
    }
    return apiModule.fetchJSON(url)
  })
}

export function createMockReadContractResult<
  TResult extends ReturnType<typeof useReadContract> = never,
>(
  overrides?: Partial<UseReadContractReturn<TResult['data']>>,
): UseReadContractReturn<TResult['data']> {
  const result = {
    data: undefined as TResult['data'],
    error: null,
    isError: false,
    isPending: true,
    isLoading: false,
    isLoadingError: false,
    isRefetchError: false,
    isSuccess: false,
    isPlaceholderData: false,
    status: 'pending' as const,
    dataUpdatedAt: 0,
    errorUpdatedAt: 0,
    failureCount: 0,
    failureReason: null,
    errorUpdateCount: 0,
    isFetched: false,
    isFetchedAfterMount: false,
    isFetching: false,
    isInitialLoading: false,
    isPaused: false,
    isRefetching: false,
    isStale: false,
    refetch: vi.fn(),
    fetchStatus: 'idle' as const,
    queryKey: [],
    promise: Promise.resolve(),
    ...overrides,
  }

  return result
}

export function createMockWriteContractResult<
  TResult extends ReturnType<typeof useWriteContract> = never,
>(
  overrides?: Partial<UseWriteContractReturn<TResult['data']>>,
): UseWriteContractReturn<TResult['data']> {
  return {
    isError: false,
    isPending: false,
    isSuccess: true,
    status: 'success',
    data: '0xTxHash',
    error: null,
    failureCount: 0,
    failureReason: null,
    isPaused: false,
    variables: undefined,
    isIdle: false,
    reset: vi.fn(),
    context: undefined,
    submittedAt: 0,
    writeContract: vi.fn().mockReturnValue('0xTxHash'),
    writeContractAsync: vi.fn().mockResolvedValue('0xTxHash'),
    ...overrides,
  }
}

export function createMockWaitForTransactionReceiptResult<
  TResult extends ReturnType<typeof useWaitForTransactionReceipt> = never,
>(
  overrides?: Partial<UseWaitForTransactionReceiptReturn<TResult['data']>>,
): UseWaitForTransactionReceiptReturn<TResult['data']> {
  return {
    isError: false,
    isPending: false,
    isSuccess: true,
    isLoading: false,
    isStale: false,
    isLoadingError: false,
    isRefetchError: false,
    isPlaceholderData: false,
    dataUpdatedAt: 0,
    errorUpdatedAt: 0,
    failureCount: 0,
    failureReason: null,
    errorUpdateCount: 0,
    isFetched: true,
    isFetchedAfterMount: true,
    isFetching: false,
    isInitialLoading: false,
    isRefetching: false,
    status: 'success',
    data: '0xTxHash',
    isPaused: false,
    refetch: vi.fn(),
    fetchStatus: 'idle' as const,
    queryKey: [],
    promise: Promise.resolve(),
    error: null,
    ...overrides,
  }
}

export function mockWagmiHooks() {
  const { useReadContract, useWriteContract, useWaitForTransactionReceipt } =
    vi.hoisted(() => {
      return {
        useReadContract: vi.fn(),
        useWriteContract: vi.fn(),
        useWaitForTransactionReceipt: vi.fn(),
      }
    })

  vi.mock('wagmi', async () => {
    const actual = await vi.importActual('wagmi')
    return {
      ...actual,
      useReadContract,
      useWriteContract,
      useWaitForTransactionReceipt,
    }
  })

  beforeEach(() => {
    useReadContract.mockReturnValue(createMockReadContractResult())
    useWriteContract.mockReturnValue(createMockWriteContractResult())
    useWaitForTransactionReceipt.mockReturnValue(
      createMockWaitForTransactionReceiptResult(),
    )
  })

  return { useReadContract, useWriteContract, useWaitForTransactionReceipt }
}
