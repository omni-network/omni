import { erc20Abi } from 'viem'
import { testAccount } from './account.js'

export const testOrder = {
  owner: testAccount.address,
  srcChainId: 1,
  destChainId: 2,
  calls: [
    {
      abi: erc20Abi,
      functionName: 'transfer',
      target: '0x23e98253f372ee29910e22986fe75bb287b011fc',
      value: 0n,
      args: [testAccount.address, 0n],
    },
  ],
  deposit: {
    token: '0x123',
    amount: 0n,
  },
  expense: {
    token: '0x123',
    amount: 0n,
  },
} as const
