import { waitFor } from '@testing-library/react'
import type { Hex, Log } from 'viem'
import { beforeEach, expect, test, vi } from 'vitest'
import { renderHook } from '../../test/index.js'
import { useWatchDidFill } from './useWatchDidFill.js'

const mockOnError = vi.fn()
const unwatch = vi.fn()
let onError: ((error: Error) => void) | undefined
const testLogs: Log[] = [
  {
    logIndex: 0,
    topics: [],
    data: '0x',
    transactionHash: '0x123',
  } as unknown as Log,
]

const { watchDidFill } = vi.hoisted(() => {
  return {
    watchDidFill: vi.fn().mockImplementation((params) => {
      onError = params.onError
      return unwatch
    }),
  }
})

vi.mock('@omni-network/core', async () => {
  const actual = await vi.importActual('@omni-network/core')
  return { ...actual, watchDidFill }
})

beforeEach(() => {
  watchDidFill.mockClear()
  mockOnError.mockClear()
  unwatch.mockClear()
  onError = undefined
})

test('default: returns destTxHash when core api triggers onLogs callback', async () => {
  const { result, rerender } = renderHook(
    (orderId?: Hex) =>
      useWatchDidFill({
        destChainId: 1,
        onError: mockOnError,
        orderId,
      }),
    { mockContractsCall: true },
  )

  expect(result.current.destTxHash).toBeUndefined()
  expect(result.current.status).toBe('idle')
  expect(result.current.unwatch).toBeTypeOf('function')
  expect(watchDidFill).not.toHaveBeenCalled()

  watchDidFill.mockImplementation((params) => {
    onError = params.onError
    params.onLogs(testLogs)
    return unwatch
  })

  rerender('0xOrderId')

  await waitFor(() => {
    expect(watchDidFill).toHaveBeenCalledTimes(1)
    expect(onError).toBeDefined()
    expect(onError).toBe(mockOnError)
    expect(result.current.destTxHash).toBe('0x123')
    expect(result.current.status).toBe('success')
  })

  result.current.unwatch()

  expect(unwatch).toHaveBeenCalledTimes(1)
})

test('params: watchDidFill is not called when orderId is undefined', async () => {
  const { result } = renderHook(() =>
    useWatchDidFill({
      destChainId: 1,
    }),
  )

  expect(result.current.destTxHash).toBeUndefined()
  expect(result.current.status).toBe('idle')
  expect(result.current.unwatch).toBeTypeOf('function')
  expect(watchDidFill).not.toHaveBeenCalled()
})
