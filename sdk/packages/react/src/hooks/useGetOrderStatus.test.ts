import { WatchDidFillError } from '@omni-network/core'
import { beforeEach, expect, test, vi } from 'vitest'
import { renderHook } from '../../test/index.js'
import { useGetOrderStatus } from './useGetOrderStatus.js'

const { useInboxStatus, useWatchDidFill, useDidFill } = vi.hoisted(() => {
  return {
    useInboxStatus: vi.fn(),
    useWatchDidFill: vi.fn(),
    useDidFill: vi.fn(),
  }
})

vi.mock('./useInboxStatus.js', async () => {
  return {
    useInboxStatus,
  }
})

vi.mock('./useWatchDidFill.js', async () => {
  return {
    useWatchDidFill,
  }
})

vi.mock('./useDidFill.js', async () => {
  return {
    useDidFill,
  }
})

beforeEach(() => {
  useInboxStatus.mockRestore()
  useDidFill.mockImplementation(() => ({
    data: false,
    error: null,
  }))
  useWatchDidFill.mockImplementation(() => ({
    status: 'idle',
    error: undefined,
    destTxHash: undefined,
    unwatchDestTx: vi.fn(),
  }))
})

const renderGetOrderStatusHook = () => {
  return renderHook(
    () =>
      useGetOrderStatus({
        destChainId: 1,
        srcChainId: 2,
        orderId: '0x123',
      }),
    {
      mockContractsCall: true,
    },
  )
}

test('default: transitions status through order lifecycle', async () => {
  useInboxStatus.mockReturnValue('unknown')

  const { result, rerender } = renderGetOrderStatusHook()

  expect(useInboxStatus).toHaveBeenCalledOnce()
  expect(useWatchDidFill).toHaveBeenCalledOnce()
  expect(result.current.status).toBe('not-found')
  expect(result.current.error).toBeUndefined()
  expect(result.current.destTxHash).toBeUndefined()

  useInboxStatus.mockReturnValue('open')
  rerender()

  expect(result.current.status).toBe('open')
  expect(result.current.error).toBeUndefined()
  expect(result.current.destTxHash).toBeUndefined()

  useWatchDidFill.mockReturnValue({
    status: 'success',
    destTxHash: '0x123',
    error: undefined,
    unwatchDestTx: vi.fn(),
  })

  rerender()

  expect(result.current.status).toBe('filled')
  expect(result.current.error).toBeUndefined()
  expect(result.current.destTxHash).toBe('0x123')
})

test('behaviour: error defined if watchDidFill error', async () => {
  useWatchDidFill.mockReturnValue({
    status: 'error',
    error: new Error('test error'),
    destTxHash: undefined,
    unwatchDestTx: vi.fn(),
  })

  const { result } = renderGetOrderStatusHook()

  expect(result.current.error).toBeDefined()
})

test('behaviour: status filled if watchDidFill is true', async () => {
  useWatchDidFill.mockReturnValue({ status: 'success', destTxHash: '0x123' })

  const { result } = renderGetOrderStatusHook()

  expect(result.current.status).toBe('filled')
  expect(result.current.destTxHash).toBe('0x123')
  expect(result.current.error).toBeUndefined()
})

test('behaviour: status filled if didFill is true, and watchDidFill fails', async () => {
  useDidFill.mockReturnValue({ status: 'success', data: true })
  useWatchDidFill.mockReturnValue({
    status: 'error',
    error: new Error('test error'),
    destTxHash: undefined,
    unwatchDestTx: vi.fn(),
  })

  const { result } = renderGetOrderStatusHook()

  expect(result.current.status).toBe('filled')
  expect(result.current.destTxHash).toBeUndefined()
  expect(result.current.error).toBeInstanceOf(WatchDidFillError)
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
