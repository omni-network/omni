import type { Address } from 'viem'

export type Quoteable = { token?: Address; amount?: bigint }

export type Quote = {
  deposit: { token: Address; amount: bigint }
  expense: { token: Address; amount: bigint }
}
