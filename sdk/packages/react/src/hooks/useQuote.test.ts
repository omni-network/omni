import * as core from '@omni-network/core'
import { QueryClient } from '@tanstack/react-query'
import { waitFor } from '@testing-library/react'
import { zeroAddress } from 'viem'
import { beforeEach, expect, test, vi } from 'vitest'
import { quote, renderHook } from '../../test/index.js'
import { type UseQuoteParams, useQuote } from './useQuote.js'

beforeEach(() => {
  vi.spyOn(core, 'getQuote').mockResolvedValue(quote)
})

const params: UseQuoteParams = {
  srcChainId: 1,
  destChainId: 2,
  mode: 'expense',
  deposit: { amount: 100n },
  enabled: true,
} as const

const renderQuoteHook = (
  params: Parameters<typeof useQuote>[0],
  queryClient?: QueryClient,
) => {
  return renderHook(() => useQuote(params), {
    mockContractsCall: true,
    queryClient,
  })
}

test('default: fetches a quote once enabled', async () => {
  const { result, rerender } = renderHook(
    ({ enabled }: { enabled: boolean }) => useQuote({ ...params, enabled }),
    { initialProps: { enabled: false } },
  )

  expect(result.current.isPending).toBe(true)
  expect(result.current.query.data).toBeUndefined()
  expect(result.current.query.isFetched).toBe(false)

  rerender({ enabled: true })

  await waitFor(() => {
    expect(result.current.query.isFetched).toBe(true)
    expect(result.current.isPending).toBe(false)
    expect(result.current.query.data).toBeDefined()
  })
})

test('parameters: expense', () => {
  const { result, rerender } = renderQuoteHook({
    ...params,
    mode: 'deposit',
    expense: { amount: 100n },
  })

  expect(result.current).toBeDefined()

  rerender({ ...params, expense: { token: '0x123', amount: 100n } })

  expect(result.current).toBeDefined()

  rerender({ ...params, expense: { token: zeroAddress, amount: 100n } })

  expect(result.current).toBeDefined()
})

test('parameters: deposit', () => {
  const { result, rerender } = renderQuoteHook({
    ...params,
    mode: 'expense',
    deposit: { amount: 100n },
  })

  expect(result.current).toBeDefined()

  rerender({ ...params, deposit: { token: '0x123', amount: 100n } })

  expect(result.current).toBeDefined()

  rerender({ ...params, deposit: { token: zeroAddress, amount: 100n } })

  expect(result.current).toBeDefined()
})

test('parameters: mode', () => {
  const { result, rerender } = renderQuoteHook({
    ...params,
    mode: 'expense',
    deposit: { amount: 100n },
  })

  expect(result.current).toBeDefined()

  rerender({
    ...params,
    mode: 'deposit',
    expense: { token: zeroAddress, amount: 100n },
  })

  expect(result.current).toBeDefined()
})

test('behaviour: quote does not fire when enabled is false', () => {
  const { result } = renderQuoteHook({ ...params, enabled: false })

  expect(result.current.isPending).toBe(true)
  expect(result.current.query.data).toBeUndefined()
  expect(result.current.query.isFetched).toBe(false)
})

test('behaviour: quote does not fire when both deposit and expense are zero', () => {
  const { result } = renderQuoteHook({
    ...params,
    enabled: true,
    deposit: { token: zeroAddress, amount: 0n },
    expense: { token: zeroAddress },
  })

  expect(result.current.isPending).toBe(true)
  expect(result.current.query.data).toBeUndefined()
  expect(result.current.query.isFetched).toBe(false)
})

test('behaviour: quote is error if call throws', async () => {
  // Use custom query client that will retry queries
  const queryClient = new QueryClient({
    defaultOptions: {
      queries: {
        retry: 3,
      },
    },
  })

  const error = new Error('Unexpected quote response')
  const spy = vi.spyOn(core, 'getQuote').mockRejectedValue(error)

  const { result } = renderQuoteHook({ ...params, enabled: true }, queryClient)
  expect(spy).toHaveBeenCalledTimes(1)

  await waitFor(() => {
    expect(result.current.query.isLoading).toBe(false)
    expect(result.current.isError).toBe(true)
    if (result.current.isError) {
      expect(result.current.error).toBe(error)
    }
  })
})
