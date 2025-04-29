import { beforeEach, expect, test, vi } from 'vitest'
import {
  createMockQueryResult,
  renderHook,
  resolvedOrder,
} from '../../test/index.js'
import { useGetOrderStatus } from './useGetOrderStatus.js'

const { useGetOrder, useInboxStatus, useDidFill } = vi.hoisted(() => {
  return {
    useGetOrder: vi.fn(),
    useInboxStatus: vi.fn(),
    useDidFill: vi.fn(),
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

vi.mock('./useDidFill.js', async () => {
  return {
    useDidFill: useDidFill,
  }
})

beforeEach(() => {
  useGetOrder.mockReturnValue(createMockQueryResult())
  useInboxStatus.mockRestore()
  useDidFill.mockImplementation(() => ({
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
  expect(useDidFill).toHaveBeenCalledOnce()
  expect(result.current.status).toBe('open')
  expect(result.current.error).toBeUndefined()
})

test('behaviour: error defined if didFill error', async () => {
  useDidFill.mockReturnValue(
    createMockQueryResult({
      error: new Error('test error'),
    }),
  )

  const { result } = renderGetOrderStatusHook()

  expect(result.current.status).toBe('error')
  expect(result.current.error).toBeDefined()
})

test('behaviour: error defined if getOrder error', async () => {
  useGetOrder.mockReturnValue(
    createMockQueryResult({
      error: new Error('test error'),
    }),
  )

  const { result } = renderGetOrderStatusHook()

  expect(result.current.status).toBe('not-found')
  expect(result.current.error).toBeDefined()
})

test('behaviour: status filled if didFill is true', async () => {
  useDidFill.mockReturnValue({ data: true })

  const { result } = renderGetOrderStatusHook()

  expect(result.current.status).toBe('filled')
  expect(result.current.error).toBeUndefined()
})

test.each(['filled', 'rejected', 'closed', 'unknown'])(
  'behaviour: status %s if inboxStatus is %s',
  async (status) => {
    useInboxStatus.mockReturnValue(status)

    const { result } = renderGetOrderStatusHook()

    expect(result.current.status).toBe(
      status === 'unknown' ? 'not-found' : status,
    )
    expect(result.current.error).toBeUndefined()
  },
)
