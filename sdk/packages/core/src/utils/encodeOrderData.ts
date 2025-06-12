import type { Hex } from 'viem'
import {
  encodeAbiParameters,
  encodeFunctionData,
  slice,
  zeroAddress,
} from 'viem'
import type { OptionalAbis } from '../types/abi.js'
import { type EVMOrder, isContractCall } from '../types/order.js'

/**
 * @description Encodes order params in preparation for sending to the inbox contract
 *
 * @param order - {@link EVMOrder}
 * @returns Encoded order - {@link Hex}
 *
 * @example
 *
 * const encodedOrderData = encodeOrderData({
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
export function encodeOrderData(order: EVMOrder<OptionalAbis>): Hex {
  const callsTuple = order.calls.map((call) => {
    if (!isContractCall(call)) {
      return {
        target: call.target,
        selector: '0x00000000',
        value: call.value,
        params: '0x',
      } as const
    }

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
      value: call.value ?? 0n,
      params,
    }
  })

  const expenseTuple = [
    {
      spender: order.expense.spender ?? zeroAddress,
      token: order.expense.token ?? zeroAddress,
      amount: order.expense.amount,
    },
    // native expenses not included in order data
  ].filter((e) => e.token !== zeroAddress)

  const depositTuple = {
    token: order.deposit.token ?? zeroAddress,
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
        order.owner ?? zeroAddress,
        BigInt(order.destChainId),
        depositTuple,
        callsTuple,
        expenseTuple,
      ],
    ],
  )
}
