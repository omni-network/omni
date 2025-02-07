import { type Hex, encodeAbiParameters } from 'viem'
import type { Order } from '../index.js'

/**
 * @description Encodes an order before sending it to the inbox contract.
 *
 * @param order - {@link Order}
 * @returns Encoded order. {@link Hex}
 *
 * @example
 *
 * const encodedOrder = encodeOrder({
 *  owner: '0x...',
 *  destChainId: 1,
 *  deposit: {
 *    token: '0x0000000000000000000000000000000000000000',
 *    amount: BigInt(1000000000000000000),
 *  },
 *  calls: [
 *    {
 *      target: '0x...',
 *       selector: '0x...',
 *       value: BigInt(1000000000000000000),
 *       params: '0x...'
 *     }
 *   ],
 *   expenses: [
 *    {
 *      spender: '0x...',
 *      token: '0x...',
 *      amount: BigInt(1000000000000000000),
 *     }
 *   ],
 * })
 */
export function encodeOrder(order: Order): Hex {
  // TODO add uint length validations
  const callsTuple = order.calls.map((call) => ({
    target: call.target,
    selector: call.selector,
    value: call.value,
    params: call.params,
  }))

  const expensesTuple = order.expenses.map((expense) => ({
    spender: expense.spender,
    token: expense.token,
    amount: expense.amount,
  }))

  const depositTuple = {
    token: order.deposit.token,
    amount: order.deposit.amount,
  }

  return encodeAbiParameters(
    [
      {
        type: 'tuple',
        components: [
          { name: 'owner', type: 'address' },
          { name: 'destChainId', type: 'uint64' },
          {
            // deposit
            type: 'tuple',
            components: [
              { name: 'token', type: 'address' },
              { name: 'amount', type: 'uint96' },
            ],
          },
          {
            // calls
            type: 'tuple[]',
            components: [
              { name: 'target', type: 'address' },
              { name: 'selector', type: 'bytes4' },
              { name: 'value', type: 'uint256' },
              { name: 'params', type: 'bytes' },
            ],
          },
          {
            // expenses array
            type: 'tuple[]',
            components: [
              { name: 'spender', type: 'address' },
              { name: 'token', type: 'address' },
              { name: 'amount', type: 'uint96' },
            ],
          },
        ],
      },
    ],
    [
      [
        order.owner,
        BigInt(order.destChainId),
        depositTuple,
        callsTuple,
        expensesTuple,
      ],
    ],
  )
}
