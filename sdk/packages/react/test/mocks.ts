import * as core from '@omni-network/core'
import { vi } from 'vitest'
import type { useReadContract } from 'wagmi'
import { contracts } from './shared.js'

type UseReadContractReturn<Data> = Omit<
  ReturnType<typeof useReadContract>,
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
