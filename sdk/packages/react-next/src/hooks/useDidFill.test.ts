import { waitFor } from '@testing-library/react'
import { expect, test, vi } from 'vitest'
import { renderHook, resolvedOrder } from '../../test/index.js'
import { type UseDidFillParams, useDidFill } from './useDidFill.js'

const { didFill } = vi.hoisted(() => {
  return {
    didFill: vi.fn().mockImplementation(() => {
      return Promise.reject(new Error('No mock'))
    }),
  }
})

vi.mock('@omni-network/core', async () => {
  const actual = await vi.importActual('@omni-network/core')
  return { ...actual, didFill }
})

const renderDidFillHook = (withResolvedOrder = false) => {
  return renderHook(
    (props: Partial<UseDidFillParams>) =>
      useDidFill({
        destChainId: 1,
        resolvedOrder: withResolvedOrder ? resolvedOrder : undefined,
        ...props,
      }),
    { mockContractsCall: true },
  )
}

test('default: returns true when core api returns truthy', async () => {
  const { result, rerender } = renderDidFillHook()

  expect(result.current.data).toBeUndefined()
  expect(didFill).not.toHaveBeenCalled()

  didFill.mockResolvedValue(true)

  rerender({
    destChainId: 1,
    resolvedOrder,
  })

  await waitFor(() => expect(result.current.data).toBe(true))
  expect(didFill).toHaveBeenCalled()
})

test('behaviour: no exception if core api throws', async () => {
  didFill.mockRejectedValue(new Error('Contract read failed'))

  const { result } = renderDidFillHook(true)

  await waitFor(() => expect(result.current.status).toBe('error'))
  expect(result.current.isError).toBe(true)
  expect(result.current.data).toBeUndefined()
  expect(didFill).toHaveBeenCalled()
})

test('behaviour: no core api call when resolvedOrder is undefined', async () => {
  const { result } = renderDidFillHook()

  expect(result.current.data).toBeUndefined()
  expect(result.current.status).toBe('pending')
  expect(result.current.isFetched).toBe(false)
  expect(didFill).not.toHaveBeenCalledOnce()
})
