import type { Hex } from 'viem'

export type Asset = {
  enabled: boolean
  name: string
  symbol: string
  chainId: number
  address: Hex
  decimals: number
  expenseMin: bigint
  expenseMax: bigint
}
