import { testOrderId, testResolvedOrder } from '@omni-network/test-utils'
import type { Client } from 'viem'
import { expect, test, vi } from 'vitest'
import { inboxABI } from '../constants/abis.js'
import { getOrder } from './getOrder.js'

const { readContract } = vi.hoisted(() => ({
  readContract: vi.fn(),
}))
vi.mock('viem/actions', () => ({ readContract }))

test('default: returns order when contract read returns an order', async () => {
  const client = {} as Client
  const inboxAddress = '0xaddress'

  const orderResult = [
    testResolvedOrder,
    {
      status: 1,
      updatedBy: '0x123',
      timestamp: 1,
      rejectReason: 0,
    } as const,
    0n,
  ]
  readContract.mockResolvedValueOnce(orderResult)

  await expect(
    getOrder({ client, inboxAddress, orderId: testOrderId }),
  ).resolves.toEqual(orderResult)

  expect(readContract).toHaveBeenCalledWith(client, {
    address: inboxAddress,
    abi: inboxABI,
    functionName: 'getOrder',
    args: [testOrderId],
  })
})
