import type { Abi, Address } from 'viem'

export type OrderStatus =
  | 'idle'
  | 'opening'
  | 'open'
  | 'closed'
  | 'rejected'
  | 'error'
  | 'filled'

type Deposit = {
  readonly token: Address
  readonly amount: bigint
}

type Expense = {
  readonly spender: Address
  readonly token: Address
  readonly amount: bigint
}

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
