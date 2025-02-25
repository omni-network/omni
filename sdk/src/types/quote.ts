import type { Address } from 'viem'

// TODO add complex type to enforce one of the amounts is defined
export type Quoteable =
  | { isNative: true; token?: never; amount?: bigint }
  | { isNative: false; token: Address; amount?: bigint }

export type Quote = {
  deposit: { token: Address; amount: bigint }
  expense: { token: Address; amount: bigint }
}
