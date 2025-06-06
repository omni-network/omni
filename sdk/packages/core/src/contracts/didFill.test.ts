import {
  encodedOriginData,
  testOrderId,
  testResolvedOrder,
} from '@omni-network/test-utils'
import type { Client } from 'viem'
import { readContract } from 'viem/actions'
import { expect, test, vi } from 'vitest'
import { outboxABI } from '../constants/abis.js'
import { didFill } from './didFill.js'

vi.mock('viem/actions', () => {
  return {
    readContract: vi.fn().mockResolvedValue(true),
  }
})

test('default: returns true when outbox read is truthy', async () => {
  const client = {} as Client
  const outboxAddress = '0xaddress'

  await expect(
    didFill({ client, outboxAddress, resolvedOrder: testResolvedOrder }),
  ).resolves.toBe(true)

  expect(readContract).toHaveBeenCalledWith(client, {
    address: outboxAddress,
    abi: outboxABI,
    functionName: 'didFill',
    args: [testOrderId, encodedOriginData],
  })
})
