import { z } from 'zod'
import { encodeFunctionData, zeroAddress } from 'viem'
import { ValidateOrderError } from '../errors/base.js'
import { fetchJSON } from '../internal/api.js'
import type { OptionalAbis } from '../types/abi.js'
import type { Environment } from '../types/config.js'
import { type Order, isContractCall } from '../types/order.js'
import { getApiUrl } from '../utils/getApiUrl.js'
import { toJSON } from '../utils/toJSON.js'

const validationResponseSchema = z.object({
  accepted: z.boolean().optional(),
  rejected: z.boolean().optional(),
  error: z
    .object({
      code: z.number(),
      message: z.string(),
    })
    .optional(),
  rejectReason: z.string().optional(),
  rejectDescription: z.string().optional(),
})

export type ValidateOrderParameters<abis extends OptionalAbis> = Order<abis> & {
  environment?: Environment | string
}

export type ValidationResponse = z.infer<typeof validationResponseSchema>

// validateOrder calls /check - checking if an order is valid
export async function validateOrder<abis extends OptionalAbis>(
  params: ValidateOrderParameters<abis>,
) {
  const { environment, ...order } = params
  const apiUrl = getApiUrl(environment)
  const json = await fetchJSON(`${apiUrl}/check`, {
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
  return validationResponseSchema.safeParse(json).success
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
