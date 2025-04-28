import {
  testAccount,
  testOrderId,
  testResolvedOrder,
} from '@omni-network/test-utils'
import {
  type Hex,
  type Log,
  encodeAbiParameters,
  encodeEventTopics,
} from 'viem'
import { expect, test } from 'vitest'
import { inboxABI } from '../constants/abis.js'
import { ParseOpenEventError } from '../errors/base.js'
import { parseOpenEvent } from './parseOpenEvent.js'

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
  [testResolvedOrder],
)

const topics = encodeEventTopics({
  abi: inboxABI,
  eventName: 'Open',
  args: {
    orderId: testOrderId,
  },
}) as [Hex, ...Hex[]]

const logs: Log[] = [
  {
    address: testAccount.address,
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

test('default: parses open event', () => {
  expect(parseOpenEvent(logs)).toEqual(testResolvedOrder)
})

test('behaviour: throws ParseOpenEventError when the logs array is empty', () => {
  try {
    parseOpenEvent([])
  } catch (err) {
    expect(err).toBeInstanceOf(ParseOpenEventError)
    expect((err as ParseOpenEventError).message).toBe(
      "Expected exactly one 'Open' event but found 0.",
    )
  }
})

test("behaviour: throws ParseOpenEventError when the logs array contains more than one 'Open' event", () => {
  try {
    parseOpenEvent([...logs, ...logs])
  } catch (err) {
    expect(err).toBeInstanceOf(ParseOpenEventError)
    expect((err as ParseOpenEventError).message).toBe(
      "Expected exactly one 'Open' event but found 2.",
    )
  }
})
