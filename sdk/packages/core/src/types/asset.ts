import type { Address } from './addresses.js'

export type Asset = {
  enabled: boolean
  name: string
  symbol: string
  chainId: number
  address: Address
  decimals: number
  expenseMin: bigint
  expenseMax: bigint
}
