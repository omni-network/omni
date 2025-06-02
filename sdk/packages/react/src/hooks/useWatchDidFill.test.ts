import type { ResolvedOrder } from '@omni-network/core'
import { testResolvedOrder } from '@omni-network/test-utils'
import { waitFor } from '@testing-library/react'
import { beforeEach, expect, test, vi } from 'vitest'
import { renderHook } from '../../test/index.js'
import { useWatchDidFill } from './useWatchDidFill.js'

const unwatch = vi.fn()
const txHash = '0x123'

const { watchDidFill } = vi.hoisted(() => {
  return {
    watchDidFill: vi.fn().mockImplementation(() => {
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
  unwatch.mockClear()
})

test('default: returns destTxHash when core api triggers onFill callback', async () => {
  const { result, rerender } = renderHook(
    (resolvedOrder?: ResolvedOrder) =>
      useWatchDidFill({
        destChainId: 1,
        resolvedOrder,
      }),
    { mockContractsCall: true },
  )

  expect(result.current.destTxHash).toBeUndefined()
  expect(result.current.status).toBe('idle')
  expect(result.current.unwatch).toBeTypeOf('function')
  expect(watchDidFill).not.toHaveBeenCalled()

  watchDidFill.mockImplementation((params) => {
    params.onFill(txHash)
    return unwatch
  })

  rerender(testResolvedOrder)

  await waitFor(() => {
    expect(watchDidFill).toHaveBeenCalledTimes(1)
    expect(result.current.destTxHash).toBe(txHash)
    expect(result.current.status).toBe('success')
  })

  result.current.unwatch()

  expect(unwatch).toHaveBeenCalledTimes(1)
})

test('behaviour: error and status are set when core api triggers onError callback', async () => {
  const error = new Error('Test error')
  watchDidFill.mockImplementation((params) => {
    params.onError?.(error)
    return unwatch
  })

  const { result } = renderHook(() =>
    useWatchDidFill({
      destChainId: 1,
      resolvedOrder: testResolvedOrder,
    }),
  )

  await waitFor(() => {
    expect(result.current.error).toBe(error)
    expect(result.current.status).toBe('error')
    expect(result.current.destTxHash).toBeUndefined()
  })
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
