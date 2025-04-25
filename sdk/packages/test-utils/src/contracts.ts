import { type Hex, parseEther, toBytes, toHex } from 'viem'
import { testAccount } from './account.js'

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
  middleman: '0x789',
} as const

export const testOrderId = toHex(toBytes(1n, { size: 32 }))

export const testOriginData = '0x123456' as const

export const testBytes32Addr = toHex(toBytes(testAccount.address, { size: 32 }))

export const testResolvedOrder = {
  user: testAccount.address,
  originChainId: 1n,
  openDeadline: 0,
  fillDeadline: 0,
  orderId: testOrderId,
  maxSpent: [
    {
      token: testBytes32Addr,
      amount: oneEth,
      recipient: testBytes32Addr,
      chainId: 1n,
    },
  ] as readonly Transfer[],
  minReceived: [
    {
      token: testBytes32Addr,
      amount: oneEth,
      recipient: testBytes32Addr,
      chainId: 1n,
    },
  ] as readonly Transfer[],
  fillInstructions: [
    {
      destinationChainId: 1n,
      destinationSettler: testBytes32Addr,
      originData: testOriginData,
    },
  ] as readonly FillInstruction[],
} as const
