import { fillOriginDataAbi } from '@omni-network/core'
import { type Hex, encodeAbiParameters, parseEther, toBytes, toHex } from 'viem'
import { testAccount } from './accounts.js'

const oneEth = parseEther('1')

type Transfer = {
  token: Hex
  amount: bigint
  recipient: Hex
  chainId: bigint
}

type FillInstruction = {
  destinationChainId: bigint
  destinationSettler: Hex
  originData: Hex
}

export const testContracts = {
  inbox: '0x123',
  outbox: '0x456',
} as const

const srcChainId = 1n
const destChainId = 2n
const fillDeadline = 1748953139
export const testOrderId = toHex(toBytes(1n, { size: 32 }))
export const testBytes32Addr = toHex(toBytes(testAccount.address, { size: 32 }))
export const originData = {
  srcChainId,
  destChainId,
  fillDeadline,
  calls: [
    {
      target: testAccount.address,
      selector: '0x00000000',
      value: 9970089730807577n,
      params: '0x',
    },
  ],
  expenses: [],
} as const
export const encodedOriginData = encodeAbiParameters(
  [fillOriginDataAbi],
  [originData],
)

export const testResolvedOrder = {
  user: testAccount.address,
  originChainId: srcChainId,
  openDeadline: 0,
  fillDeadline,
  orderId: testOrderId,
  maxSpent: [
    {
      token: testBytes32Addr,
      amount: oneEth,
      recipient: testBytes32Addr,
      chainId: srcChainId,
    },
  ] as readonly Transfer[],
  minReceived: [
    {
      token: testBytes32Addr,
      amount: oneEth,
      recipient: testBytes32Addr,
      chainId: srcChainId,
    },
  ] as readonly Transfer[],
  fillInstructions: [
    {
      destinationChainId: destChainId,
      destinationSettler: testBytes32Addr,
      originData: encodedOriginData,
    },
  ] as readonly FillInstruction[],
} as const
