import {
  type Hex,
  encodeAbiParameters,
  encodeFunctionData,
  slice,
  zeroAddress,
} from 'viem'
import type { Order } from '../types/order.js'

/**
 * @description Encodes order params in preparation for sending to the inbox contract
 *
 * @param order - {@link Order}
 * @returns Encoded order - {@link Hex}
 *
 * @example
 *
 * const encodedOrder = encodeOrder({
 *  owner: '0x...',
 *  destChainId: 1,
 *  deposit: {
 *    token: '0x...',
 *    amount: BigInt(1),
 *  },
 *  calls: [
 *    {
 *      target: '0x...',
 *       selector: 'deposit',
 *       value: BigInt(1),
 *       args: [
 *         '0x...',
 *        BigInt(1),
 *       ]
 *     }
 *   ],
 *   expense {
 *      spender: '0x...',
 *      token: '0x...',
 *      amount: BigInt(1),
 *     }
 *   ],
 * })
 */
export function encodeOrder(order: Order) {
  // TODO custom error
  if (!order.deposit || !order.expense || !order.owner) {
    throw new Error('Invalid order')
  }

  const callsTuple = order.calls.map((call) => {
    const callData = encodeFunctionData({
      abi: call.abi,
      functionName: call.functionName,
      args: call.args,
    })

    const selector = slice(callData, 0, 4)

    const params = callData.length > 10 ? slice(callData, 4) : '0x'

    return {
      target: call.target,
      selector,
      value: call.value,
      params,
    }
  })

  const expenseTuple = {
    spender: order.expense.spender,
    token: order.expense.isNative ? zeroAddress : order.expense.token,
    amount: order.expense.amount,
  }

  const depositTuple = {
    token: order.deposit.isNative ? zeroAddress : order.deposit.token,
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
            // expense
            type: 'tuple',
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
        expenseTuple,
      ],
    ],
  )
}
