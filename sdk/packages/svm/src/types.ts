import type BN from 'bn.js'

export type EvmCall = {
  target: Array<number> // 20 bytes (EVM address)
  selector: Array<number> // 4 bytes
  value: BN // u128
  params: Buffer // variable length
}

export type EvmTokenExpense = {
  spender: Array<number> // 20 bytes (EVM address)
  token: Array<number> // 20 bytes (EVM address)
  amount: BN // u128
}
