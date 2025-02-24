import type { Abi, Address } from 'viem'
import type { Prettify } from './utils.js'

type NativeToken = {
  readonly amount: bigint
  readonly isNative: true
}

type ERC20Token = {
  readonly token: Address
  readonly amount: bigint
  readonly isNative: false
}

export type Deposit = Prettify<NativeToken | ERC20Token>
export type Expense = Prettify<
  { readonly spender: Address } & (NativeToken | ERC20Token)
>

export type Call = {
  readonly abi: Abi
  readonly target: Address
  readonly value: bigint
  // TODO: infer selector and args from abi
  readonly functionName: string
  readonly args?: unknown[]
}

export type Order = {
  readonly owner?: Address
  readonly srcChainId?: number
  readonly destChainId: number
  readonly calls: Call[]
  readonly fillDeadline?: number
  readonly deposit: Deposit
  readonly expense: Expense
}
