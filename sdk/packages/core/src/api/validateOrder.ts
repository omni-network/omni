import { encodeFunctionData, zeroAddress } from 'viem'
import { fetchJSON } from '../internal/api.js'
import type { OptionalAbis } from '../types/abi.js'
import type { Order } from '../types/order.js'
import { isContractCall } from '../types/order.js'
import { toJSON } from '../utils/toJSON.js'

export type ValidationResponse = {
  accepted?: boolean
  rejected?: boolean
  error?: {
    code: number
    message: string
  }
  rejectReason?: string
  rejectDescription?: string
}

export function encodeOrderForValidation<abis extends OptionalAbis>(
  order: Order<abis>,
): string {
  try {
    const calls = order.calls.map((call) => {
      if (!isContractCall(call)) {
        return {
          target: call.target,
          value: call.value,
          data: '0x',
        }
      }

      const callData = encodeFunctionData({
        abi: call.abi,
        functionName: call.functionName,
        args: call.args,
      })

      return {
        target: call.target,
        value: call.value,
        data: callData,
      }
    })

    return toJSON({
      orderId: order.owner ?? zeroAddress,
      sourceChainId: order.srcChainId,
      destChainId: order.destChainId,
      fillDeadline: order.fillDeadline ?? Math.floor(Date.now() / 1000 + 86400),
      calls,
      deposit: {
        amount: order.deposit.amount,
        token: order.deposit.token ?? zeroAddress,
      },
      expenses: [
        {
          amount: order.expense.amount,
          token: order.expense.token ?? zeroAddress,
          spender: order.expense.spender ?? zeroAddress,
        },
      ],
    })
  } catch (e) {
    const error =
      e instanceof Error
        ? e
        : new Error((e as { message: string }).message ?? 'Invalid order')
    throw error
  }
}

export async function validateOrderWithEncoded(
  apiBaseUrl: string,
  encodedOrder: string,
) {
  const json = await fetchJSON(`${apiBaseUrl}/check`, {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: encodedOrder,
  })

  if (!isValidateRes(json)) {
    throw new Error(`Unexpected validation response: ${JSON.stringify(json)}`)
  }

  return json as ValidationResponse
}

export async function validateOrder<abis extends OptionalAbis>(
  apiBaseUrl: string,
  order: Order<abis>,
) {
  return await validateOrderWithEncoded(
    apiBaseUrl,
    encodeOrderForValidation(order),
  )
}

// TODO schema validation
const isValidateRes = (json: unknown): json is ValidationResponse => {
  const res = json as ValidationResponse
  return (
    json != null &&
    (res.accepted != null ||
      (res.rejected != null &&
        res.rejectReason != null &&
        res.rejectDescription != null) ||
      res.error != null)
  )
}
