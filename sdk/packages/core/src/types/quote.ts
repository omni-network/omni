import type { Address, EVMAddress } from './addresses.js'

export type Quoteable = { token?: Address; amount?: bigint }

export type Quote = {
  // EVM and SVM deposits are supported
  deposit: { token: Address; amount: bigint }
  // only EVM expenses are supported
  expense: { token: EVMAddress; amount: bigint }
}
