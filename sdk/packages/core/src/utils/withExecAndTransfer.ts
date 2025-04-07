import type { Abi, Address } from 'viem'
import { encodeFunctionData } from 'viem'
import { middlemanABI } from '../constants/abis.js'
import type { Call } from '../types/order.js'

type WithExecAndTransferParams = {
  readonly middlemanAddress: Address
  readonly call: Call<Abi>
  readonly transfer: {
    readonly to: Address
    readonly token: Address
  }
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
 *  transfer: {
 *    to: '0x...',
 *    token: '0x...',
 *  }
 * })
 */
export function withExecAndTransfer(
  params: WithExecAndTransferParams,
): Call<typeof middlemanABI> {
  const { call, transfer } = params
  const _callData = encodeFunctionData({ ...call })

  return {
    ...call,
    abi: middlemanABI,
    target: params.middlemanAddress,
    functionName: 'executeAndTransfer',
    args: [transfer.token, transfer.to, call.target, _callData],
  }
}
