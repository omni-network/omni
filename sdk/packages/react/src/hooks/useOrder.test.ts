import { waitFor } from '@testing-library/react'
import { expect, test, vi } from 'vitest'
import {
  createMockWaitForTransactionReceiptResult,
  createMockWriteContractResult,
} from '../../test/mocks.js'
import { renderHook } from '../../test/react.js'
import { order } from '../../test/shared.js'
import { useOrder } from './useOrder.js'

const {
  mockUseValidateOrder,
  mockUseGetOrderStatus,
  mockUseWriteContract,
  mockUseWaitForTransactionReceipt,
} = vi.hoisted(() => {
  return {
    mockUseValidateOrder: vi.fn(),
    mockUseGetOrderStatus: vi.fn(),
    mockUseWaitForTransactionReceipt: vi.fn().mockImplementation(() => {
      return createMockWaitForTransactionReceiptResult({
        isPending: true,
        isSuccess: false,
        data: undefined,
        status: 'pending',
      })
    }),
    mockUseWriteContract: vi.fn().mockImplementation(() => {
      return createMockWriteContractResult({
        isPending: true,
        isSuccess: false,
        data: undefined,
        isIdle: true,
        status: 'pending',
      })
    }),
  }
})

vi.mock('wagmi', async () => {
  const actual = await vi.importActual('wagmi')
  return {
    ...actual,
    useWriteContract: mockUseWriteContract,
    useWaitForTransactionReceipt: mockUseWaitForTransactionReceipt,
  }
})

vi.mock('./useValidateOrder.js', async () => {
  return {
    useValidateOrder: mockUseValidateOrder,
  }
})

vi.mock('./useGetOrderStatus.js', async () => {
  return {
    useGetOrderStatus: mockUseGetOrderStatus,
  }
})

test(`default: validates, opens, and transitions order through it's lifecycle`, async () => {
  mockUseValidateOrder.mockReturnValue({
    status: 'pending',
  })

  mockUseGetOrderStatus.mockReturnValue({
    status: 'not-found',
  })

  const { result, rerender } = renderHook(
    ({ validateEnabled }: { validateEnabled: boolean }) =>
      useOrder({ ...order, validateEnabled }),
    { mockContractsCall: true, initialProps: { validateEnabled: false } },
  )

  await waitFor(() => {
    expect(result.current.isReady).toBeFalsy()
    expect(result.current.isError).toBeFalsy()
    expect(result.current.isTxPending).toBeTruthy()
    expect(result.current.isTxSubmitted).toBeFalsy()
    expect(result.current.txMutation.data).toBeUndefined()
    expect(result.current.isOpen).toBeFalsy()
    expect(result.current.orderId).toBeUndefined()
    expect(result.current.txHash).toBeUndefined()
    expect(result.current.error).toBeUndefined()
  })

  mockUseValidateOrder.mockReturnValue({
    status: 'accepted',
  })

  mockUseGetOrderStatus.mockReturnValue({
    status: 'open',
  })

  rerender({ validateEnabled: true })

  await waitFor(() => expect(result.current.isValidated).toBeTruthy())

  mockUseWriteContract
    .mockReset()
    .mockImplementation(() => createMockWriteContractResult())

  mockUseWaitForTransactionReceipt.mockImplementation(() =>
    createMockWaitForTransactionReceiptResult(),
  )

  rerender({ validateEnabled: true })

  const res = await result.current.open()

  await waitFor(() => {
    expect(result.current.isOpen).toBeTruthy()
    expect(result.current.isTxPending).toBeFalsy()
    expect(result.current.isTxSubmitted).toBeTruthy()
    expect(result.current.txMutation.data).toBe('0xTxHash')
    expect(result.current.txMutation.isSuccess).toBeTruthy()
    expect(res).toBe('0xTxHash')
  })

  mockUseGetOrderStatus.mockReturnValue({
    status: 'filled',
  })

  rerender({ validateEnabled: true })

  await waitFor(() => expect(result.current.status).toBe('filled'))
})
