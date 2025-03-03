import { arbitrum, base, optimism } from 'viem/chains'
import { http, createConfig } from 'wagmi'
import { mainnet } from 'wagmi/chains'

export const MOCK_L1_ID = 1652
export const MOCK_L2_ID = 1654
export const ZERO_ADDRESS = '0x0000000000000000000000000000000000000000'

export const accounts = ['0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266'] as const

// 0xf39Fd6e51aad88F6F4ce6aB8827279cffFb92266
export const privateKey =
  '0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80'

export const web3Config = createConfig({
  chains: [mainnet, base, optimism, arbitrum],
  pollingInterval: 100,
  storage: null,
  transports: {
    [mainnet.id]: http(),
    [optimism.id]: http(),
    [base.id]: http(),
    [arbitrum.id]: http(),
  },
})
