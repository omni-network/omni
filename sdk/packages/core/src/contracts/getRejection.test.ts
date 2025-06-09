import type { Client } from 'viem'
import { expect, test, vi } from 'vitest'
import { GetRejectionError } from '../errors/base.js'
import { getRejection, rejectReasons } from './getRejection.js'

const { getLogs } = vi.hoisted(() => ({
  getLogs: vi.fn(),
}))

vi.mock('viem/actions', () => ({ getLogs }))

const mockClient = {} as Client
const orderId = '0x123'
const mockInboxAddress = '0x1234567890123456789012345678901234567890'
const mockTxHash =
  '0xabcdef1234567890abcdef1234567890abcdef1234567890abcdef1234567890'

test('default: returns rejection when reject event logs are found', async () => {
  const logs = [
    {
      transactionHash: mockTxHash,
      args: {
        id: orderId,
        by: '0x123',
        reason: 4,
      },
    },
  ]

  getLogs.mockResolvedValueOnce(logs)

  const result = await getRejection({
    client: mockClient,
    orderId,
    inboxAddress: mockInboxAddress,
    fromBlock: 1n,
  })

  expect(result).toEqual({
    txHash: mockTxHash,
    rejectReason: 'Insufficient deposit',
  })

  expect(getLogs).toHaveBeenCalledWith(mockClient, {
    address: mockInboxAddress,
    fromBlock: 1n,
    toBlock: 'latest',
    event: expect.objectContaining({
      type: 'event',
      name: 'Rejected',
    }),
    args: {
      id: orderId,
    },
    strict: true,
  })
})

test('behaviour: throws GetRejectionError when no reject event logs are found', async () => {
  getLogs.mockResolvedValueOnce([])

  await expect(
    getRejection({
      client: mockClient,
      orderId,
      inboxAddress: mockInboxAddress,
      fromBlock: 100n,
    }),
  ).rejects.toBeInstanceOf(GetRejectionError)
})

test.each(
  Object.entries(rejectReasons).map(([code, reason]) => [Number(code), reason]),
)(
  'behaviour: correctly parses reject reason %i as "%s"',
  async (reasonCode, expectedReason) => {
    const logs = [
      {
        transactionHash: mockTxHash,
        args: {
          id: orderId,
          by: '0x123',
          reason: reasonCode,
        },
      },
    ]

    getLogs.mockResolvedValueOnce(logs)

    const result = await getRejection({
      client: mockClient,
      orderId,
      inboxAddress: mockInboxAddress,
      fromBlock: 1n,
    })

    expect(result?.rejectReason).toBe(expectedReason)
  },
)

test('behaviour: throws error when invalid reject reason is found', async () => {
  const logs = [
    {
      transactionHash: mockTxHash,
      args: {
        id: orderId,
        by: '0x123',
        reason: 999, // invalid reason code
      },
    },
  ]

  getLogs.mockResolvedValueOnce(logs)

  await expect(
    getRejection({
      client: mockClient,
      orderId,
      inboxAddress: mockInboxAddress,
      fromBlock: 1n,
    }),
  ).rejects.toThrow('Invalid reject reason key')
})

test('behaviour: throws error when transaction hash is missing', async () => {
  const logs = [
    {
      transactionHash: null, // missing tx hash
      args: {
        id: orderId,
        by: '0x123',
        reason: 4,
      },
    },
  ]

  getLogs.mockResolvedValueOnce(logs)

  await expect(
    getRejection({
      client: mockClient,
      orderId,
      inboxAddress: mockInboxAddress,
      fromBlock: 1n,
    }),
  ).rejects.toThrow('Tx hash is always present for "latest" blocks')
})

test('behaviour: uses correct block range parameters', async () => {
  const logs = [
    {
      transactionHash: mockTxHash,
      args: {
        id: orderId,
        by: '0x123',
        reason: 4,
      },
    },
  ]

  getLogs.mockResolvedValueOnce(logs)

  await getRejection({
    client: mockClient,
    orderId,
    inboxAddress: mockInboxAddress,
    fromBlock: 12345n,
  })

  expect(getLogs).toHaveBeenCalledWith(mockClient, {
    address: mockInboxAddress,
    fromBlock: 12345n,
    toBlock: 'latest',
    event: expect.objectContaining({
      type: 'event',
      name: 'Rejected',
    }),
    args: {
      id: orderId,
    },
    strict: true,
  })
})
