import { beforeEach, expect, test, vi } from 'vitest'
import { resolvedOrder } from '../../test/index.js'
import { createMockReadContractResult } from '../../test/mocks.js'
import { renderHook } from '../../test/react.js'
import { useGetOrderStatus } from './useGetOrderStatus.js'

const { mockUseGetOrder, mockUseInboxStatus, mockUseDidFillOutbox } =
  vi.hoisted(() => {
    return {
      mockUseGetOrder: vi.fn(),
      mockUseInboxStatus: vi.fn(),
      mockUseDidFillOutbox: vi.fn(),
    }
  })

vi.mock('./useGetOrder.js', async () => {
  return {
    useGetOrder: mockUseGetOrder,
  }
})

vi.mock('./useInboxStatus.js', async () => {
  return {
    useInboxStatus: mockUseInboxStatus,
  }
})

vi.mock('./useDidFillOutbox.js', async () => {
  return {
    useDidFillOutbox: mockUseDidFillOutbox,
  }
})

beforeEach(() => {
  mockUseGetOrder.mockReturnValue(createMockReadContractResult())
  mockUseInboxStatus.mockRestore()
  mockUseDidFillOutbox.mockImplementation(() => ({
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
  mockUseInboxStatus.mockReturnValue('open')

  const { result } = renderGetOrderStatusHook()

  // once on mount only
  expect(mockUseInboxStatus).toHaveBeenCalledOnce()
  expect(mockUseDidFillOutbox).toHaveBeenCalledOnce()
  expect(result.current.status).toBe('open')
  expect(result.current.error).toBeUndefined()
})

test('behaviour: error defined if didFillOutbox error', async () => {
  mockUseDidFillOutbox.mockReturnValue(
    createMockReadContractResult({
      error: new Error('test error'),
    }),
  )

  const { result } = renderGetOrderStatusHook()

  expect(result.current.status).toBe('error')
  expect(result.current.error).toBeDefined()
})

test('behaviour: error defined if getOrder error', async () => {
  mockUseGetOrder.mockReturnValue(
    createMockReadContractResult({
      error: new Error('test error'),
    }),
  )

  const { result } = renderGetOrderStatusHook()

  expect(result.current.status).toBe('not-found')
  expect(result.current.error).toBeDefined()
})

test('behaviour: status filled if didFillOutbox is true', async () => {
  mockUseDidFillOutbox.mockReturnValue({ data: true })

  const { result } = renderGetOrderStatusHook()

  expect(result.current.status).toBe('filled')
  expect(result.current.error).toBeUndefined()
})

test('behaviour: status filled if inboxStatus is filled', async () => {
  mockUseInboxStatus.mockReturnValue('filled')

  const { result } = renderGetOrderStatusHook()

  expect(result.current.status).toBe('filled')
  expect(result.current.error).toBeUndefined()
})

test('behaviour: status rejected if inboxStatus is rejected', async () => {
  mockUseInboxStatus.mockReturnValue('rejected')

  const { result } = renderGetOrderStatusHook()

  expect(result.current.status).toBe('rejected')
  expect(result.current.error).toBeUndefined()
})

test('behaviour: status closed if inboxStatus is closed', async () => {
  mockUseInboxStatus.mockReturnValue('closed')

  const { result } = renderGetOrderStatusHook()

  expect(result.current.status).toBe('closed')
  expect(result.current.error).toBeUndefined()
})

test('behaviour: status not found if inboxStatus is unknown', async () => {
  mockUseInboxStatus.mockReturnValue('unknown')

  const { result } = renderGetOrderStatusHook()

  expect(result.current.status).toBe('not-found')
  expect(result.current.error).toBeUndefined()
})
