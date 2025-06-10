import type { AnyAddress, EVMAddress } from './addresses.js'

export type Quoteable = { token?: AnyAddress; amount?: bigint }

export type Quote = {
  // EVM and SVM deposits are supported
  deposit: { token: AnyAddress; amount: bigint }
  // only EVM expenses are supported
  expense: { token: EVMAddress; amount: bigint }
}
