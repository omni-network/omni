import { encodeFunctionData, zeroAddress } from 'viem'
import { ValidateOrderError } from '../errors/base.js'
import { fetchJSON } from '../internal/api.js'
import type { OptionalAbis } from '../types/abi.js'
import { type Order, isContractCall } from '../types/order.js'
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

// validateOrder calls /check - checking if an order is valid
export async function validateOrder<abis extends OptionalAbis>(
  apiBaseUrl: string,
  order: Order<abis>,
) {
  const json = await fetchJSON(`${apiBaseUrl}/check`, {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: serialize(order),
  })

  if (!isValidateRes(json)) {
    throw new Error(`Unexpected validation response: ${JSON.stringify(json)}`)
  }

  return json satisfies ValidationResponse
}

const serialize = <abis extends OptionalAbis>(order: Order<abis>) => {
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

// asserts a json response is ValidationResponse
const isValidateRes = (json: unknown): json is ValidationResponse => {
  // TODO schema validation
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

export type AcceptedResult = {
  accepted: true
  rejected?: false
  error?: never
  rejectReason?: never
  rejectDescription?: never
}

export function assertAcceptedResult(
  res: ValidationResponse,
): asserts res is AcceptedResult {
  if (!res.accepted) {
    if (res.error != null) {
      throw new ValidateOrderError(res.error.message, `Code ${res.error.code}`)
    }
    if (res.rejected) {
      throw new ValidateOrderError(
        res.rejectDescription ?? 'Server rejected order',
        res.rejectReason,
      )
    }
  }
}
