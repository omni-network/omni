import type { EVMOrder, OmniContracts, OptionalAbis } from '@omni-network/core'
import { testAccount } from '@omni-network/test-utils'
import {
  type Hex,
  erc20Abi,
  parseEther,
  toBytes,
  toHex,
  zeroAddress,
} from 'viem'
import { arbitrum, base, optimism } from 'viem/chains'
import { http, createConfig, mock } from 'wagmi'
import { mainnet } from 'wagmi/chains'

////////////////////////////////////////
//// TEST DATA
////////////////////////////////////////

// 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266
export const accounts = [testAccount.address] as const

export const web3Config = createConfig({
  chains: [mainnet, base, optimism, arbitrum],
  connectors: [mock({ accounts }), mock({ accounts })],
  pollingInterval: 100,
  storage: null,
  transports: {
    [mainnet.id]: http(),
    [optimism.id]: http(),
    [base.id]: http(),
    [arbitrum.id]: http(),
  },
})

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

export const oneEth = parseEther('1')
export const contracts: OmniContracts = {
  inbox: '0x123',
  outbox: '0x456',
}
export const orderId = toHex(toBytes(1n, { size: 32 }))
export const originData = '0x123456' as const
export const bytes32Addr = toHex(toBytes(accounts[0], { size: 32 }))
export const resolvedOrder = {
  user: accounts[0],
  originChainId: 1n,
  openDeadline: 0,
  fillDeadline: 0,
  orderId,
  maxSpent: [
    {
      token: bytes32Addr,
      amount: oneEth,
      recipient: bytes32Addr,
      chainId: 1n,
    },
  ] as readonly Transfer[],
  minReceived: [
    {
      token: bytes32Addr,
      amount: oneEth,
      recipient: bytes32Addr,
      chainId: 1n,
    },
  ] as readonly Transfer[],
  fillInstructions: [
    {
      destinationChainId: 1n,
      destinationSettler: bytes32Addr,
      originData,
    },
  ] as readonly FillInstruction[],
} as const

export const orderRequest: EVMOrder<OptionalAbis> = {
  srcChainId: 1,
  destChainId: 2,
  deposit: {
    token: zeroAddress,
    amount: 100n,
  },
  expense: {
    token: zeroAddress,
    amount: 90n,
  },
  calls: [],
}

export const quote = {
  deposit: { token: zeroAddress, amount: 100n },
  expense: { token: zeroAddress, amount: 99n },
}

export const order = {
  owner: accounts[0],
  srcChainId: 1,
  destChainId: 2,
  calls: [
    {
      abi: erc20Abi,
      functionName: 'transfer',
      target: '0x23e98253f372ee29910e22986fe75bb287b011fc',
      value: BigInt(0),
      args: [accounts[0], 0n],
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

export const orderStatusData = {
  status: 1,
  updatedBy: '0x123',
  timestamp: 1,
  rejectReason: 0,
} as const
