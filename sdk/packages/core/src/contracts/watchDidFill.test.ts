import { testResolvedOrder } from '@omni-network/test-utils'
import type { Block, Client } from 'viem'
import { beforeEach, expect, test, vi } from 'vitest'
import type { EVMAddress } from '../types/addresses.js'
import { watchDidFill } from './watchDidFill.js'

const client = {} as Client
const outboxAddrs = '0x0' as EVMAddress
const resolvedOrder = testResolvedOrder
const txHash = '0xTxHash'
const blockHash = '0xBlockHash'
const mockOnError = vi.fn()
const mockOnFill = vi.fn()
const unwatch = vi.fn()
let onBlock: ((block: Block) => Promise<void>) | undefined
let onError: ((error: Error) => void) | undefined

const block = {
  hash: blockHash,
} as unknown as Block

const { watchBlocks, getLogs } = vi.hoisted(() => ({
  watchBlocks: vi.fn().mockImplementation((_client, params) => {
    onError = params.onError
    onBlock = params.onBlock
    return unwatch
  }),
  getLogs: vi.fn().mockResolvedValue([
    {
      transactionHash: '0xTxHash',
    },
  ]),
}))

vi.mock('viem/actions', () => ({
  watchBlocks,
  getLogs,
}))

beforeEach(() => {
  watchBlocks.mockClear()
  mockOnError.mockClear()
  mockOnFill.mockClear()
  onBlock = undefined
  onError = undefined
  unwatch.mockClear()
})

test('default: should trigger onFill callback when filled function in block and return an unwatch func', async () => {
  const unwatch = watchDidFill({
    client,
    outboxAddress: outboxAddrs,
    resolvedOrder,
    onFill: mockOnFill,
    onError: mockOnError,
    pollingInterval: 5000,
  })

  expect(unwatch).toBeTypeOf('function')
  expect(onBlock).toBeDefined()
  expect(onError).toBe(mockOnError)

  await onBlock?.(block)
  expect(mockOnFill).toHaveBeenCalledTimes(1)
  expect(mockOnFill).toHaveBeenCalledWith(txHash)

  expect(getLogs).toHaveBeenCalledTimes(1)

  expect(watchBlocks).toHaveBeenCalledTimes(1)
  expect(watchBlocks).toHaveBeenCalledWith(client, {
    onBlock,
    onError: mockOnError,
    pollingInterval: 5000,
    emitMissed: true,
  })
  unwatch()
  expect(unwatch).toHaveBeenCalledTimes(1)
})

test('behaviour: should trigger onError callback when watchBlocks emits an error', () => {
  watchDidFill({
    client,
    outboxAddress: outboxAddrs,
    resolvedOrder,
    onFill: mockOnFill,
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
    resolvedOrder,
    onFill: mockOnFill,
  })

  expect(watchBlocks).toHaveBeenCalledTimes(1)
  expect(watchBlocks).toHaveBeenCalledWith(client, {
    onBlock,
    onError: undefined,
    pollingInterval: undefined,
    emitMissed: true,
  })
})
