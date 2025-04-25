import type { Client } from 'viem'
import { readContract } from 'viem/actions'
import { expect, test, vi } from 'vitest'
import { outboxABI } from '../constants/abis.js'
import { didFillOutbox } from './didFillOutbox.js'
import type { ResolvedOrder } from './parseOpenEvent.js'

vi.mock('viem/actions', () => {
  return {
    readContract: vi.fn().mockResolvedValue(true),
  }
})

test('default: returns true when outbox read is truthy', async () => {
  const client = {} as Client
  const outboxAddress = '0xaddress'
  const resolvedOrder = {
    orderId: '0xorderId',
    fillInstructions: [{ originData: '0xoriginData' }],
  } as unknown as ResolvedOrder

  await expect(
    didFillOutbox({ client, outboxAddress, resolvedOrder }),
  ).resolves.toBe(true)

  expect(readContract).toHaveBeenCalledWith(client, {
    address: outboxAddress,
    abi: outboxABI,
    functionName: 'didFill',
    args: ['0xorderId', '0xoriginData'],
  })
})
