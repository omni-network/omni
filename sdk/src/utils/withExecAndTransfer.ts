import { type Address, encodeFunctionData } from 'viem'
import { middleman } from '../constants/contracts.js'
import type { Call } from '../types/order.js'

type WithExecAndTransferParams = {
  readonly call: Call
  readonly to: Address
  readonly token: Address
}

/**
 * @description Constructs a call to be executed via a middleware contract, that executes the input
 * call and transfers the specified token to the recipient.
 *
 * @param params - Parameters for the middleware call.
 * @returns Call
 *
 * @example
 *
 * const callWithExecAndTransfer = withExecAndTransfer({
 *  call: {
 *    target: '0x...',
 *    abi,
 *    value: BigInt(1),
 *    functionName: '0x...',
 *    value: BigInt(1),
 *    args: ['0x...', BigInt(1)],
 *  },
 *  to: '0x...',
 *  token: '0x...',
 * })
 */
export function withExecAndTransfer(params: WithExecAndTransferParams): Call {
  const { call, to, token } = params
  const _callData = encodeFunctionData({
    abi: call.abi,
    functionName: call.functionName,
    args: call.args,
  })

  return {
    ...call,
    abi: middleman.abi,
    target: middleman.address,
    functionName: 'executeAndTransfer',
    args: [token, to, call.target, _callData],
  }
}
