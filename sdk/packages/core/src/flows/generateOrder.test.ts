import {
  mockL1Client,
  testOrder,
  testResolvedOrder,
} from '@omni-network/test-utils'
import { beforeEach, expect, test, vi } from 'vitest'
import { validateOrder } from '../api/validateOrder.js'
import { sendOrder } from '../contracts/sendOrder.js'
import { waitForOrderClose } from '../contracts/waitForOrderClose.js'
import { waitForOrderOpen } from '../contracts/waitForOrderOpen.js'
import { generateOrder } from './generateOrder.js'

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
vi.mock('../contracts/waitForOrderClose.js', () => ({
  waitForOrderClose: vi.fn(),
}))

beforeEach(() => {
  vi.clearAllMocks()
})

test('default: successfully generates an order through all states', async () => {
  vi.mocked(validateOrder).mockResolvedValueOnce({ accepted: true })
  const txHash = '0xtxHash'
  vi.mocked(sendOrder).mockResolvedValueOnce(txHash)
  vi.mocked(waitForOrderOpen).mockResolvedValueOnce(testResolvedOrder)
  vi.mocked(waitForOrderClose).mockResolvedValueOnce('filled')

  const generator = generateOrder({
    client: mockL1Client,
    inboxAddress: '0xaddress',
    order: testOrder,
  })

  const validState = await generator.next()
  expect(validState.value).toEqual({ status: 'valid' })
  expect(validateOrder).toHaveBeenCalledWith(testOrder)

  const sentState = await generator.next()
  expect(sentState.value).toEqual({ status: 'sent', txHash })
  expect(sendOrder).toHaveBeenCalledWith({
    client: mockL1Client,
    inboxAddress: '0xaddress',
    order: testOrder,
  })

  const openState = await generator.next()
  expect(openState.value).toEqual({
    status: 'open',
    txHash,
    order: testResolvedOrder,
  })
  expect(waitForOrderOpen).toHaveBeenCalledWith({
    client: mockL1Client,
    txHash,
  })

  const finalState = await generator.next()
  expect(finalState.value).toEqual({
    status: 'filled',
    txHash,
    order: testResolvedOrder,
  })
  expect(waitForOrderClose).toHaveBeenCalledWith({
    client: mockL1Client,
    inboxAddress: '0xaddress',
    orderId: testResolvedOrder.orderId,
  })

  const done = await generator.next()
  expect(done.done).toBe(true)
})

test('behaviour: passes environment and pollingInterval to appropriate functions', async () => {
  const environment = 'http://localhost'
  const pollingInterval = 1000

  vi.mocked(validateOrder).mockResolvedValueOnce({ accepted: true })
  const txHash = '0xtxHash'
  vi.mocked(sendOrder).mockResolvedValueOnce(txHash)
  vi.mocked(waitForOrderOpen).mockResolvedValueOnce(testResolvedOrder)
  vi.mocked(waitForOrderClose).mockResolvedValueOnce('filled')

  const generator = generateOrder({
    client: mockL1Client,
    inboxAddress: '0xaddress',
    order: testOrder,
    environment,
    pollingInterval,
  })
  while (!(await generator.next()).done) {}

  expect(validateOrder).toHaveBeenCalledWith({
    ...testOrder,
    environment,
  })
  expect(waitForOrderOpen).toHaveBeenCalledWith({
    client: mockL1Client,
    txHash,
    pollingInterval,
  })
  expect(waitForOrderClose).toHaveBeenCalledWith({
    client: mockL1Client,
    inboxAddress: '0xaddress',
    orderId: testResolvedOrder.orderId,
    pollingInterval,
  })
})

test('behaviour: throws when order validation fails', async () => {
  const validationError = new Error('Validation failed')
  vi.mocked(validateOrder).mockRejectedValueOnce(validationError)

  const generator = generateOrder({
    client: mockL1Client,
    inboxAddress: '0xaddress',
    order: testOrder,
  })

  await expect(generator.next()).rejects.toThrow(validationError)
  expect(sendOrder).not.toHaveBeenCalled()
  expect(waitForOrderOpen).not.toHaveBeenCalled()
  expect(waitForOrderClose).not.toHaveBeenCalled()
})

test('behaviour: throws when sending order fails', async () => {
  vi.mocked(validateOrder).mockResolvedValueOnce({ accepted: true })
  const sendError = new Error('Send failed')
  vi.mocked(sendOrder).mockRejectedValueOnce(sendError)

  const generator = generateOrder({
    client: mockL1Client,
    inboxAddress: '0xaddress',
    order: testOrder,
  })
  await generator.next()

  await expect(generator.next()).rejects.toThrow(sendError)
  expect(waitForOrderOpen).not.toHaveBeenCalled()
  expect(waitForOrderClose).not.toHaveBeenCalled()
})

test('behaviour: throws when waiting for order open fails', async () => {
  vi.mocked(validateOrder).mockResolvedValueOnce({ accepted: true })
  const txHash = '0xtxHash'
  vi.mocked(sendOrder).mockResolvedValueOnce(txHash)
  const waitError = new Error('Wait failed')
  vi.mocked(waitForOrderOpen).mockRejectedValueOnce(waitError)

  const generator = generateOrder({
    client: mockL1Client,
    inboxAddress: '0xaddress',
    order: testOrder,
  })
  await generator.next()
  await generator.next()

  await expect(generator.next()).rejects.toThrow(waitError)
  expect(waitForOrderClose).not.toHaveBeenCalled()
})

test('behaviour: throws when waiting for order close fails', async () => {
  vi.mocked(validateOrder).mockResolvedValueOnce({ accepted: true })
  const txHash = '0xtxHash'
  vi.mocked(sendOrder).mockResolvedValueOnce(txHash)
  vi.mocked(waitForOrderOpen).mockResolvedValueOnce(testResolvedOrder)
  const waitError = new Error('Wait failed')
  vi.mocked(waitForOrderClose).mockRejectedValueOnce(waitError)

  const generator = generateOrder({
    client: mockL1Client,
    inboxAddress: '0xaddress',
    order: testOrder,
  })
  await generator.next()
  await generator.next()
  await generator.next()

  await expect(generator.next()).rejects.toThrow(waitError)
})

test('behaviour: forwards transaction options to sendOrder', async () => {
  vi.mocked(validateOrder).mockResolvedValueOnce({ accepted: true })
  const txHash = '0xtxHash'
  vi.mocked(sendOrder).mockResolvedValueOnce(txHash)
  vi.mocked(waitForOrderOpen).mockResolvedValueOnce(testResolvedOrder)
  vi.mocked(waitForOrderClose).mockResolvedValueOnce('filled')

  const transactionOptions = {
    gas: 100000n,
    maxFeePerGas: 100000n,
    maxPriorityFeePerGas: 100000n,
  }

  const generator = generateOrder({
    client: mockL1Client,
    inboxAddress: '0xaddress',
    order: testOrder,
    ...transactionOptions,
  })
  while (!(await generator.next()).done) {}

  expect(sendOrder).toHaveBeenCalledWith({
    client: mockL1Client,
    inboxAddress: '0xaddress',
    order: testOrder,
    ...transactionOptions,
  })
})
