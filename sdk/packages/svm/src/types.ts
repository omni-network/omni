export type EvmCall = {
  target: Uint8Array // 20 bytes (EVM address)
  selector: Uint8Array // 4 bytes
  value: bigint // u128
  params: Uint8Array // variable length
}

export type EvmTokenExpense = {
  spender: Uint8Array // 20 bytes (EVM address)
  token: Uint8Array // 20 bytes (EVM address)
  amount: bigint // u128
}
