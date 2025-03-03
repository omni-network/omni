import { waitFor } from '@testing-library/react'
import {
  type Hex,
  type Log,
  encodeAbiParameters,
  encodeEventTopics,
  toBytes,
  toHex,
  zeroAddress,
} from 'viem'
import { expect, test } from 'vitest'
import { accounts, renderHook } from '../../test/index.js'
import { inboxABI } from '../constants/abis.js'
import { ParseOpenEventError } from '../errors/base.js'
import { useParseOpenEvent } from './useParseOpenEvent.js'

const orderId = toHex(toBytes(1n, { size: 32 }))
const originData = '0x123456' as Hex
const resolvedOrder = {
  user: accounts[0],
  originChainId: BigInt(1),
  openDeadline: 0,
  fillDeadline: 0,
  orderId,
  maxSpent: [
    {
      token: toHex(toBytes(zeroAddress, { size: 32 })),
      amount: 1000000000000000000n, // 1 eth
      recipient: toHex(toBytes(accounts[0], { size: 32 })),
      chainId: BigInt(1),
    },
  ],
  minReceived: [
    {
      token: toHex(toBytes(zeroAddress, { size: 32 })),
      amount: 1000000000000000000n, // 1eth
      recipient: toHex(toBytes(accounts[0], { size: 32 })),
      chainId: BigInt(1),
    },
  ],
  fillInstructions: [
    {
      destinationChainId: BigInt(1),
      destinationSettler: toHex(toBytes(accounts[0], { size: 32 })),
      originData,
    },
  ],
}

const encodedOpenEvent = encodeAbiParameters(
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
    data: encodedOpenEvent,
    blockHash: '0x1',
    blockNumber: 1n,
    logIndex: 1,
    transactionHash: '0x1',
    transactionIndex: 1,
    removed: false,
  },
]

test('should return undefined results when status pending', async () => {
  const { result } = renderHook(
    () =>
      useParseOpenEvent({
        status: 'pending',
        logs: [],
      }),
    { mockContractsCall: true },
  )

  expect(result.current.orderId).toBeUndefined()
  expect(result.current.originData).toBeUndefined()
})

test('should return orderId and originData when props update to expected values', async () => {
  const { result, rerender } = renderHook(
    () =>
      useParseOpenEvent({
        status: 'pending',
        logs: [],
      }),
    { mockContractsCall: true },
  )

  expect(result.current.orderId).toBeUndefined()
  expect(result.current.originData).toBeUndefined()

  rerender({
    status: 'success',
    logs: logs,
  })

  await waitFor(() => result.current.orderId === orderId)
  await waitFor(() => result.current.originData === originData)
})

test('should return instance of ParseOpenEventError if success status and empty logs', async () => {
  const { result } = renderHook(
    () =>
      useParseOpenEvent({
        status: 'success',
        logs: [],
      }),
    { mockContractsCall: true },
  )

  expect(result.current.error).toBeInstanceOf(ParseOpenEventError)
})

test('should return instance of ParseOpenEventError if success status and logs are not valid', async () => {
  const { result } = renderHook(
    () =>
      useParseOpenEvent({
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
      }),
    { mockContractsCall: true },
  )

  expect(result.current.error).toBeInstanceOf(ParseOpenEventError)
})
