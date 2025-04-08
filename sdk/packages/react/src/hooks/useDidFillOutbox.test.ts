import { waitFor } from '@testing-library/react'
import { beforeEach, expect, test, vi } from 'vitest'
import { resolvedOrder } from '../../test/index.js'
import { createMockReadContractResult } from '../../test/mocks.js'
import { renderHook } from '../../test/react.js'
import { useDidFillOutbox } from './useDidFillOutbox.js'

const { useReadContract } = vi.hoisted(() => {
  return {
    useReadContract: vi.fn().mockImplementation(() => {
      return createMockReadContractResult()
    }),
  }
})

vi.mock('wagmi', async () => {
  const actual = await vi.importActual('wagmi')
  return {
    ...actual,
    useReadContract,
  }
})

beforeEach(() => {
  useReadContract.mockReturnValue(createMockReadContractResult())
})

test('default: returns true when outbox read is truthy', async () => {
  const { result, rerender } = renderHook(
    () =>
      useDidFillOutbox({
        destChainId: 1,
      }),
    { mockContractsCall: true },
  )

  expect(result.current.data).toBeUndefined()
  expect(useReadContract).toHaveBeenCalled()

  useReadContract.mockReturnValue(
    createMockReadContractResult<ReturnType<typeof useDidFillOutbox>>({
      data: true,
      isSuccess: true,
      status: 'success',
    }),
  )

  rerender({
    destChainId: 1,
    resolvedOrder,
  })

  await waitFor(() => expect(result.current.data).toBe(true))
})

test('behaviour: no exception if contract read fails', () => {
  useReadContract.mockReturnValue(
    createMockReadContractResult<ReturnType<typeof useDidFillOutbox>>({
      isSuccess: false,
      isError: true,
      status: 'error',
    }),
  )

  const { result } = renderHook(
    () =>
      useDidFillOutbox({
        destChainId: 1,
        resolvedOrder,
      }),
    { mockContractsCall: true },
  )

  expect(result.current.status).toBe('error')
  expect(result.current.isError).toBe(true)
  expect(result.current.data).toBeUndefined()
  expect(useReadContract).toHaveBeenCalled()
})

test('behaviour: no contract read when resolvedOrder is undefined', async () => {
  const { result } = renderHook(
    () =>
      useDidFillOutbox({
        destChainId: 1,
        resolvedOrder: undefined,
      }),
    { mockContractsCall: true },
  )

  expect(result.current.data).toBeUndefined()
  expect(result.current.status).toBe('pending')
  expect(result.current.isFetched).toBe(false)
  // once on mount
  expect(useReadContract).toHaveBeenCalledOnce()
})
