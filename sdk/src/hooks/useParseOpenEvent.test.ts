import { waitFor } from '@testing-library/react'
import {
  type Hex,
  type Log,
  encodeAbiParameters,
  encodeEventTopics,
} from 'viem'
import { expect, test } from 'vitest'
import {
  accounts,
  orderId,
  renderHook,
  resolvedOrder,
} from '../../test/index.js'
import { inboxABI } from '../constants/abis.js'
import { ParseOpenEventError } from '../errors/base.js'
import { useParseOpenEvent } from './useParseOpenEvent.js'

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

test('default', async () => {
  const { result, rerender } = renderHook(
    () =>
      useParseOpenEvent({
        status: 'pending',
        logs: [],
      }),
    { mockContractsCall: true },
  )

  expect(result.current.resolvedOrder).toBeUndefined()

  rerender({
    status: 'success',
    logs,
  })

  await waitFor(() => result.current.resolvedOrder?.orderId === orderId)
})

// TODO
test('parameters: status', () => {})

// TODO
test('parameters: logs', () => {})

// TODO: assert decodeEventLog not called via spy/mock
// vi.mock('./src/calculator.ts', { spy: true })
test('behaviour: no parsing when status is pending', () => {
  const { result } = renderHook(
    () =>
      useParseOpenEvent({
        status: 'pending',
        logs,
      }),
    { mockContractsCall: true },
  )

  expect(result.current.resolvedOrder).toBeUndefined()
})

// TODO: assert decodeEventLog not called via spy/mock
test('behaviour: no parsing when logs is undefined', () => {
  const { result } = renderHook(
    () =>
      useParseOpenEvent({
        status: 'success',
      }),
    { mockContractsCall: true },
  )

  expect(result.current.resolvedOrder).toBeUndefined()
})

test('behaviour: error when status is success and logs is []', () => {
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

test('behaviour: error when status is success and logs not valid', () => {
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
