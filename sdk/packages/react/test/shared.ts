import { type Hex, parseEther, toBytes, toHex } from 'viem'
import { arbitrum, base, optimism } from 'viem/chains'
import { http, createConfig, mock } from 'wagmi'
import { mainnet } from 'wagmi/chains'

////////////////////////////////////////
//// TEST DATA
////////////////////////////////////////
export const MOCK_L1_ID = 1652
export const MOCK_L2_ID = 1654
export const ZERO_ADDRESS = '0x0000000000000000000000000000000000000000'

export const accounts = ['0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266'] as const

// 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266
export const privateKey =
  '0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80'

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
export const contracts = {
  inbox: '0x123',
  outbox: '0x456',
  middleman: '0x789',
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
