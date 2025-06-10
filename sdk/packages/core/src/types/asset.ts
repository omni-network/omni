import type { AnyAddress } from './addresses.js'

export type Asset = {
  enabled: boolean
  name: string
  symbol: string
  chainId: number
  address: AnyAddress
  decimals: number
  expenseMin: bigint
  expenseMax: bigint
}
