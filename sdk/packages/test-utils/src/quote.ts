import { zeroAddress } from 'viem'

export const testQuote = {
  deposit: { token: zeroAddress, amount: 100n },
  expense: { token: zeroAddress, amount: 99n },
} as const
