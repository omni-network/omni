import { vi } from 'vitest'
import type { useReadContract } from 'wagmi'
import * as apiModule from '../src/internal/api.js'
import { contracts } from './shared.js'

type UseReadContractReturn<Data> = Omit<
  ReturnType<typeof useReadContract>,
  'error' | 'data'
> & {
  error: Error | null
  data: Data
}

export function mockContractsQuery() {
  vi.spyOn(apiModule, 'fetchJSON').mockImplementation((url: string) => {
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
