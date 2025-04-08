import { beforeEach, expect, test, vi } from 'vitest'
import { resolvedOrder } from '../../test/index.js'
import { createMockReadContractResult } from '../../test/mocks.js'
import { renderHook } from '../../test/react.js'
import { useGetOrderStatus } from './useGetOrderStatus.js'

const { useGetOrder, useInboxStatus, useDidFillOutbox } = vi.hoisted(() => {
  return {
    useGetOrder: vi.fn(),
    useInboxStatus: vi.fn(),
    useDidFillOutbox: vi.fn(),
  }
})

vi.mock('./useGetOrder.js', async () => {
  return {
    useGetOrder: useGetOrder,
  }
})

vi.mock('./useInboxStatus.js', async () => {
  return {
    useInboxStatus: useInboxStatus,
  }
})

vi.mock('./useDidFillOutbox.js', async () => {
  return {
    useDidFillOutbox: useDidFillOutbox,
  }
})

beforeEach(() => {
  useGetOrder.mockReturnValue(createMockReadContractResult())
  useInboxStatus.mockRestore()
  useDidFillOutbox.mockImplementation(() => ({
    data: false,
    error: null,
  }))
})

const renderGetOrderStatusHook = () => {
  return renderHook(
    () =>
      useGetOrderStatus({
        destChainId: 1,
        srcChainId: 2,
        orderId: '0x123',
        resolvedOrder,
      }),
    {
      mockContractsCall: true,
    },
  )
}

test('default: if inbox status is open, returned status is open', async () => {
  useInboxStatus.mockReturnValue('open')

  const { result } = renderGetOrderStatusHook()

  // once on mount only
  expect(useInboxStatus).toHaveBeenCalledOnce()
  expect(useDidFillOutbox).toHaveBeenCalledOnce()
  expect(result.current.status).toBe('open')
  expect(result.current.error).toBeUndefined()
})

test('behaviour: error defined if didFillOutbox error', async () => {
  useDidFillOutbox.mockReturnValue(
    createMockReadContractResult({
      error: new Error('test error'),
    }),
  )

  const { result } = renderGetOrderStatusHook()

  expect(result.current.status).toBe('error')
  expect(result.current.error).toBeDefined()
})

test('behaviour: error defined if getOrder error', async () => {
  useGetOrder.mockReturnValue(
    createMockReadContractResult({
      error: new Error('test error'),
    }),
  )

  const { result } = renderGetOrderStatusHook()

  expect(result.current.status).toBe('not-found')
  expect(result.current.error).toBeDefined()
})

test('behaviour: status filled if didFillOutbox is true', async () => {
  useDidFillOutbox.mockReturnValue({ data: true })

  const { result } = renderGetOrderStatusHook()

  expect(result.current.status).toBe('filled')
  expect(result.current.error).toBeUndefined()
})

test('behaviour: status filled if inboxStatus is filled', async () => {
  useInboxStatus.mockReturnValue('filled')

  const { result } = renderGetOrderStatusHook()

  expect(result.current.status).toBe('filled')
  expect(result.current.error).toBeUndefined()
})

test('behaviour: status rejected if inboxStatus is rejected', async () => {
  useInboxStatus.mockReturnValue('rejected')

  const { result } = renderGetOrderStatusHook()

  expect(result.current.status).toBe('rejected')
  expect(result.current.error).toBeUndefined()
})

test('behaviour: status closed if inboxStatus is closed', async () => {
  useInboxStatus.mockReturnValue('closed')

  const { result } = renderGetOrderStatusHook()

  expect(result.current.status).toBe('closed')
  expect(result.current.error).toBeUndefined()
})

test('behaviour: status not found if inboxStatus is unknown', async () => {
  useInboxStatus.mockReturnValue('unknown')

  const { result } = renderGetOrderStatusHook()

  expect(result.current.status).toBe('not-found')
  expect(result.current.error).toBeUndefined()
})
