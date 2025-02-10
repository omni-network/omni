import type { Abi, Address } from 'viem'

/**
 * @description Defines the structure of an order.
 *
 * @param params Order object:
 * - owner: owner of the order
 * - destChainId: destination chainID
 * - deposit: token and amount to be sent on source chain
 * - calls: array of calls to be executed on the destination chain
 * - expenses: array of expenses required to fulfill the order
 *
 * @example
 *
 * const order: Order = {
 *  owner: '0x...',
 *  destChainId: 1,
 *  deposit: {
 *    token: '0x0000000000000000000000000000000000000000',
 *    amount: BigInt(1000000000000000000),
 *  },
 *  calls: [
 *    {
 *      target: '0x...',
 *      selector: '0x...',
 *      value: BigInt(1000000000000000000),
 *      params: '0x...'
 *     }
 *   ],
 *   expenses: [
 *    {
 *      spender: '0x...',
 *      token: '0x...',
 *      amount: BigInt(1000000000000000000),
 *     }
 *   ],
 * }
 */
export type Order = {
  readonly owner: Address
  readonly destChainId: number
  readonly calls: Call[]
  readonly deposit: Deposit
  readonly expenses: Expense[]
}

type Deposit = {
  readonly token: Address
  readonly amount: bigint
}

type Expense = {
  readonly spender: Address
  readonly token: Address
  readonly amount: bigint
}

type Call = {
  readonly abi: Abi
  readonly target: Address
  readonly value: bigint
  // TODO: infer selector and args from abi
  readonly functionName: string
  readonly args?: unknown[]
}
