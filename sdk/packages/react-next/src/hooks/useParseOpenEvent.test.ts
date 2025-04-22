import { ParseOpenEventError, inboxABI } from '@omni-network/core'
import { waitFor } from '@testing-library/react'
import {
  type Hex,
  type Log,
  encodeAbiParameters,
  encodeEventTopics,
} from 'viem'
import { expect, test } from 'vitest'
import type { UseWaitForTransactionReceiptReturnType } from 'wagmi'
import {
  accounts,
  orderId,
  renderHook,
  resolvedOrder,
} from '../../test/index.js'
import { useParseOpenEvent } from './useParseOpenEvent.js'

const eventData = encodeAbiParameters(
  [
    {
      type: 'tuple',
      components: [
        { type: 'address', name: 'user' },
        { type: 'uint256', name: 'originChainId' },
        { type: 'uint32', name: 'openDeadline' },
        { type: 'uint32', name: 'fillDeadline' },
        { type: 'bytes32', name: 'orderId' },
        {
          type: 'tuple[]',
          name: 'maxSpent',
          components: [
            { type: 'bytes32', name: 'token' },
            { type: 'uint256', name: 'amount' },
            { type: 'bytes32', name: 'recipient' },
            { type: 'uint256', name: 'chainId' },
          ],
        },
        {
          type: 'tuple[]',
          name: 'minReceived',
          components: [
            { type: 'bytes32', name: 'token' },
            { type: 'uint256', name: 'amount' },
            { type: 'bytes32', name: 'recipient' },
            { type: 'uint256', name: 'chainId' },
          ],
        },
        {
          type: 'tuple[]',
          name: 'fillInstructions',
          components: [
            { type: 'uint64', name: 'destinationChainId' },
            { type: 'bytes32', name: 'destinationSettler' },
            { type: 'bytes', name: 'originData' },
          ],
        },
      ],
    },
  ],
  [resolvedOrder],
)

const topics = encodeEventTopics({
  abi: inboxABI,
  eventName: 'Open',
  args: {
    orderId,
  },
}) as [Hex, ...Hex[]]

const logs: Log[] = [
  {
    address: accounts[0],
    topics,
    data: eventData,
    blockHash: '0x1',
    blockNumber: 1n,
    logIndex: 1,
    transactionHash: '0x1',
    transactionIndex: 1,
    removed: false,
  },
]

const renderParseOpenEventHook = (
  params: Parameters<typeof useParseOpenEvent>[0],
  options?: Parameters<typeof renderHook>[1],
) => {
  return renderHook(() => useParseOpenEvent(params), {
    mockContractsCall: true,
    ...options,
  })
}

test('default: parses open event', async () => {
  const { result, rerender } = renderHook(
    ({
      status,
      logs,
    }: {
      status: UseWaitForTransactionReceiptReturnType['status']
      logs: Log[]
    }) =>
      useParseOpenEvent({
        status,
        logs,
      }),
    { mockContractsCall: true, initialProps: { status: 'pending', logs: [] } },
  )

  expect(result.current.resolvedOrder).toBeUndefined()

  rerender({
    status: 'success',
    logs,
  })

  await waitFor(() =>
    expect(result.current.resolvedOrder?.orderId).toBe(orderId),
  )
})

test('behaviour: no parsing when status is pending', () => {
  const { result } = renderParseOpenEventHook({
    status: 'pending',
    logs: [],
  })

  expect(result.current.resolvedOrder).toBeUndefined()
})

test('behaviour: no parsing when logs is undefined', () => {
  const { result } = renderParseOpenEventHook({
    status: 'success',
    logs: [],
  })

  expect(result.current.resolvedOrder).toBeUndefined()
})

test('behaviour: error when status is success and logs is empty array', () => {
  const { result } = renderParseOpenEventHook({
    status: 'success',
    logs: [],
  })

  expect(result.current.error).toBeInstanceOf(ParseOpenEventError)
})

test('behaviour: error when status is success and logs not valid', () => {
  const { result } = renderParseOpenEventHook({
    status: 'success',
    logs: [
      {
        address: accounts[0],
        topics: ['0x1'],
        data: '0x1',
        blockHash: '0x1',
        blockNumber: 1n,
        logIndex: 1,
        transactionHash: '0x1',
        transactionIndex: 1,
        removed: false,
      },
    ],
  })

  expect(result.current.error).toBeInstanceOf(ParseOpenEventError)
})
