import * as core from '@omni-network/core'
import type { UseQueryResult } from '@tanstack/react-query'
import { vi } from 'vitest'
import type { useWaitForTransactionReceipt } from 'wagmi'
import { contracts } from './shared.js'

type UseWaitForTransactionReceiptReturn<Data> = Omit<
  ReturnType<typeof useWaitForTransactionReceipt>,
  'error' | 'data'
> & {
  error: Error | null
  data: Data
}

export function mockContractsQuery(failure = false) {
  vi.spyOn(core, 'getContracts').mockImplementation(() => {
    if (failure) {
      return Promise.reject(new Error('mock error'))
    }
    return Promise.resolve(contracts)
  })
}

export function createMockQueryResult<TData = never>(
  overrides?: Partial<UseQueryResult<TData>>,
): UseQueryResult<TData> {
  const result = {
    data: undefined as TData,
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

  return result as UseQueryResult<TData>
}

export function createMockWaitForTransactionReceiptResult<TData = unknown>(
  overrides?: Partial<UseWaitForTransactionReceiptReturn<TData>>,
): UseWaitForTransactionReceiptReturn<TData> {
  return {
    data: undefined as TData,
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
    isPaused: false,
    refetch: vi.fn(),
    fetchStatus: 'idle' as const,
    queryKey: [],
    promise: Promise.resolve(),
    error: null,
    ...overrides,
  } as UseWaitForTransactionReceiptReturn<TData>
}
