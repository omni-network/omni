import { vi } from 'vitest'
import * as apiModule from '../src/internal/api.js'
import { contracts } from './shared.js'

export function mockContractsQuery() {
  vi.spyOn(apiModule, 'fetchJSON').mockImplementation((url: string) => {
    if (url.includes('/contracts')) {
      return Promise.resolve(contracts)
    }
    return apiModule.fetchJSON(url)
  })
}

export function createMockReadContractResult(overrides = {}) {
  return {
    data: undefined,
    error: null,
    isError: false,
    isPending: true,
    isLoading: false,
    isLoadingError: false,
    isRefetchError: false,
    isSuccess: false,
    isPlaceholderData: false,
    status: 'pending',
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
    remove: vi.fn(),
    fetchStatus: 'idle',
    queryKey: [],
    queryHash: '',
    promise: Promise.resolve(),
    ...overrides,
  } as const
}
