import {
  mockL1Client,
  testOrder,
  testResolvedOrder,
} from '@omni-network/test-utils'
import { beforeEach, expect, test, vi } from 'vitest'
import { validateOrder } from '../api/validateOrder.js'
import { sendOrder } from '../contracts/sendOrder.js'
import { waitForOrderOpen } from '../contracts/waitForOrderOpen.js'
import { openOrder } from './openOrder.js'

vi.mock('../api/validateOrder.js', () => ({
  validateOrder: vi.fn(),
  assertAcceptedResult: vi.fn(),
}))
vi.mock('../contracts/sendOrder.js', () => ({
  sendOrder: vi.fn(),
}))
vi.mock('../contracts/waitForOrderOpen.js', () => ({
  waitForOrderOpen: vi.fn(),
}))

beforeEach(() => {
  vi.clearAllMocks()
})

test('default: successfully opens an order', async () => {
  vi.mocked(validateOrder).mockResolvedValueOnce({ accepted: true })
  const txHash = '0xtxHash'
  vi.mocked(sendOrder).mockResolvedValueOnce(txHash)
  vi.mocked(waitForOrderOpen).mockResolvedValueOnce(testResolvedOrder)

  const result = await openOrder({
    client: mockL1Client,
    inboxAddress: '0xaddress',
    order: testOrder,
  })

  expect(result).toEqual(testResolvedOrder)
  expect(validateOrder).toHaveBeenCalledWith({
    ...testOrder,
    environment: undefined,
  })
  expect(sendOrder).toHaveBeenCalledWith({
    client: mockL1Client,
    inboxAddress: '0xaddress',
    order: testOrder,
  })
  expect(waitForOrderOpen).toHaveBeenCalledWith({
    client: mockL1Client,
    txHash,
    pollingInterval: undefined,
  })
})

test('behaviour: passes environment and pollingInterval to appropriate functions', async () => {
  const environment = 'http://localhost'
  const pollingInterval = 1000

  vi.mocked(validateOrder).mockResolvedValueOnce({ accepted: true })
  const txHash = '0xtxHash'
  vi.mocked(sendOrder).mockResolvedValueOnce(txHash)
  vi.mocked(waitForOrderOpen).mockResolvedValueOnce(testResolvedOrder)

  await openOrder({
    client: mockL1Client,
    inboxAddress: '0xaddress',
    order: testOrder,
    environment,
    pollingInterval,
  })

  expect(validateOrder).toHaveBeenCalledWith({
    ...testOrder,
    environment,
  })
  expect(waitForOrderOpen).toHaveBeenCalledWith({
    client: mockL1Client,
    txHash,
    pollingInterval,
  })
})

test('behaviour: throws when order validation fails', async () => {
  const validationError = new Error('Validation failed')
  vi.mocked(validateOrder).mockRejectedValueOnce(validationError)

  await expect(
    openOrder({
      client: mockL1Client,
      inboxAddress: '0xaddress',
      order: testOrder,
    }),
  ).rejects.toThrow(validationError)

  expect(sendOrder).not.toHaveBeenCalled()
  expect(waitForOrderOpen).not.toHaveBeenCalled()
})

test('behaviour: throws when sending order fails', async () => {
  vi.mocked(validateOrder).mockResolvedValueOnce({ accepted: true })
  const sendError = new Error('Send failed')
  vi.mocked(sendOrder).mockRejectedValueOnce(sendError)

  await expect(
    openOrder({
      client: mockL1Client,
      inboxAddress: '0xaddress',
      order: testOrder,
    }),
  ).rejects.toThrow(sendError)

  expect(waitForOrderOpen).not.toHaveBeenCalled()
})

test('behaviour: throws when waiting for order open fails', async () => {
  vi.mocked(validateOrder).mockResolvedValueOnce({ accepted: true })
  const txHash = '0xtxHash'
  vi.mocked(sendOrder).mockResolvedValueOnce(txHash)
  const waitError = new Error('Wait failed')
  vi.mocked(waitForOrderOpen).mockRejectedValueOnce(waitError)

  await expect(
    openOrder({
      client: mockL1Client,
      inboxAddress: '0xaddress',
      order: testOrder,
    }),
  ).rejects.toThrow(waitError)
})

test('behaviour: forwards transaction options to sendOrder', async () => {
  vi.mocked(validateOrder).mockResolvedValueOnce({ accepted: true })
  const txHash = '0xtxHash'
  vi.mocked(sendOrder).mockResolvedValueOnce(txHash)
  vi.mocked(waitForOrderOpen).mockResolvedValueOnce(testResolvedOrder)

  const transactionOptions = {
    gas: 100000n,
    maxFeePerGas: 100000n,
    maxPriorityFeePerGas: 100000n,
  }
  await openOrder({
    client: mockL1Client,
    inboxAddress: '0xaddress',
    order: testOrder,
    ...transactionOptions,
  })
  expect(sendOrder).toHaveBeenCalledWith({
    client: mockL1Client,
    inboxAddress: '0xaddress',
    order: testOrder,
    ...transactionOptions,
  })
})
