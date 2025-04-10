import { waitFor } from '@testing-library/react'
import { expect, test, vi } from 'vitest'
import { resolvedOrder } from '../../test/index.js'
import { renderHook } from '../../test/react.js'
import {
  type UseDidFillOutboxParams,
  useDidFillOutbox,
} from './useDidFillOutbox.js'

const { didFillOutbox } = vi.hoisted(() => {
  return {
    didFillOutbox: vi.fn().mockImplementation(() => {
      return Promise.reject(new Error('No mock'))
    }),
  }
})

vi.mock('@omni-network/core', async () => {
  const actual = await vi.importActual('@omni-network/core')
  return { ...actual, didFillOutbox }
})

const renderDidFillOutboxHook = (withResolvedOrder = false) => {
  return renderHook(
    (props: Partial<UseDidFillOutboxParams>) =>
      useDidFillOutbox({
        destChainId: 1,
        resolvedOrder: withResolvedOrder ? resolvedOrder : undefined,
        ...props,
      }),
    { mockContractsCall: true },
  )
}

test('default: returns true when outbox read is truthy', async () => {
  const { result, rerender } = renderDidFillOutboxHook()

  expect(result.current.data).toBeUndefined()
  expect(didFillOutbox).not.toHaveBeenCalled()

  didFillOutbox.mockResolvedValue(true)

  rerender({
    destChainId: 1,
    resolvedOrder,
  })

  await waitFor(() => expect(result.current.data).toBe(true))
  expect(didFillOutbox).toHaveBeenCalled()
})

test('behaviour: no exception if contract read fails', async () => {
  didFillOutbox.mockRejectedValue(new Error('Contract read failed'))

  const { result } = renderDidFillOutboxHook(true)

  await waitFor(() => expect(result.current.status).toBe('error'))
  expect(result.current.isError).toBe(true)
  expect(result.current.data).toBeUndefined()
  expect(didFillOutbox).toHaveBeenCalled()
})

test('behaviour: no contract read when resolvedOrder is undefined', async () => {
  const { result } = renderDidFillOutboxHook()

  expect(result.current.data).toBeUndefined()
  expect(result.current.status).toBe('pending')
  expect(result.current.isFetched).toBe(false)
  expect(didFillOutbox).not.toHaveBeenCalledOnce()
})
