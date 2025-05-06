import type { Address, Client, Log } from 'viem'
import { beforeEach, expect, test, vi } from 'vitest'
import { outboxABI } from '../constants/abis.js'
import { watchDidFill } from './watchDidFill.js'

const client = {} as Client
const outboxAddrs = '0x0' as Address
const orderId = '0x123'
const testLogs: Log[] = [
  { logIndex: 0, topics: [], data: '0x' } as unknown as Log,
]
const mockOnLogs = vi.fn()
const mockOnError = vi.fn()
const unwatch = vi.fn()
let onLogs: ((logs: Log[]) => void) | undefined
let onError: ((error: Error) => void) | undefined

const { watchContractEvent } = vi.hoisted(() => ({
  watchContractEvent: vi.fn().mockImplementation((_client, params) => {
    onLogs = params.onLogs
    onError = params.onError
    return unwatch
  }),
}))

vi.mock('viem/actions', () => ({
  watchContractEvent,
}))

beforeEach(() => {
  watchContractEvent.mockClear()
  unwatch.mockClear()
  mockOnLogs.mockClear()
  mockOnError.mockClear()
  onLogs = undefined
  onError = undefined
})

test('default: should trigger onLogs callback on Filled event and return an unwatch func', () => {
  const unwatch = watchDidFill({
    client,
    outboxAddress: outboxAddrs,
    orderId,
    onLogs: mockOnLogs,
    onError: mockOnError,
    pollingInterval: 5000,
  })

  expect(unwatch).toBeTypeOf('function')

  expect(onLogs).toBeDefined()
  expect(onLogs).toBe(mockOnLogs)

  if (onLogs) {
    onLogs(testLogs)
  }

  expect(mockOnLogs).toHaveBeenCalledTimes(1)
  expect(mockOnLogs).toHaveBeenCalledWith(testLogs)

  expect(watchContractEvent).toHaveBeenCalledTimes(1)
  expect(watchContractEvent).toHaveBeenCalledWith(client, {
    address: outboxAddrs,
    eventName: 'Filled',
    abi: outboxABI,
    args: { orderId },
    onLogs: mockOnLogs,
    onError: mockOnError,
    pollingInterval: 5000,
  })

  unwatch()

  expect(unwatch).toHaveBeenCalledTimes(1)
})

test('behaviour: should trigger onError callback when watchContractEvent emits an error', () => {
  watchDidFill({
    client,
    outboxAddress: outboxAddrs,
    orderId,
    onLogs: mockOnLogs,
    onError: mockOnError,
  })

  const error = new Error('Simulated watch error')
  expect(onError).toBeDefined()
  expect(onError).toBe(mockOnError)

  if (onError) {
    onError(error)
  }

  expect(mockOnError).toHaveBeenCalledTimes(1)
  expect(mockOnError).toHaveBeenCalledWith(error)
})

test('params: should pass undefined for optional params if not provided', () => {
  watchDidFill({
    client,
    outboxAddress: outboxAddrs,
    orderId,
    onLogs: mockOnLogs,
  })

  expect(watchContractEvent).toHaveBeenCalledTimes(1)
  expect(watchContractEvent).toHaveBeenCalledWith(client, {
    address: outboxAddrs,
    eventName: 'Filled',
    abi: outboxABI,
    args: { orderId },
    onLogs: mockOnLogs,
    onError: undefined,
    pollingInterval: undefined,
  })
})
