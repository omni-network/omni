import { z } from 'zod/v4-mini'
import { encodeFunctionData, zeroAddress } from 'viem'
import { ValidateOrderError } from '../errors/base.js'
import { fetchJSON } from '../internal/api.js'
import type { OptionalAbis } from '../types/abi.js'
import type { Environment } from '../types/config.js'
import { type Order, isContractCall } from '../types/order.js'
import { getApiUrl } from '../utils/getApiUrl.js'
import { toJSON } from '../utils/toJSON.js'

const acceptedResponseSchema = z.object({
  accepted: z.literal(true),
  rejectCode: z.optional(z.literal(0)),
  rejected: z.optional(z.literal(false)),
  rejectReason: z.optional(z.literal('')),
  rejectDescription: z.optional(z.literal('')),
})

const rejectedResponseSchema = z.object({
  accepted: z.optional(z.literal(false)),
  rejectCode: z.optional(z.number()),
  rejected: z.literal(true),
  rejectReason: z.string(),
  rejectDescription: z.string(),
}) 

const errorResponseSchema = z.object({
  error: z.object({
    code: z.number(),
    message: z.string(),
  }),
})

export type ValidateOrderParameters<abis extends OptionalAbis> = Order<abis> & {
  environment?: Environment | string
}

export type ValidationResponse = 
  | z.infer<typeof acceptedResponseSchema>
  | z.infer<typeof rejectedResponseSchema>
  | z.infer<typeof errorResponseSchema>

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
  return (
    acceptedResponseSchema.safeParse(json).success ||
    rejectedResponseSchema.safeParse(json).success ||
    errorResponseSchema.safeParse(json).success
  )
}

export type AcceptedResult = z.infer<typeof acceptedResponseSchema>

export function assertAcceptedResult(
  res: ValidationResponse,
): asserts res is AcceptedResult {

  // if the response is accepted
  if(acceptedResponseSchema.safeParse(res).success) {
    return;
  }

  // if the response is an error
  if (errorResponseSchema.safeParse(res).success) {
    const { error } = errorResponseSchema.parse(res);
    throw new ValidateOrderError(
      error.message,
      `Code ${error.code}`,
    )
  }

  // if the response is rejected
  if (rejectedResponseSchema.safeParse(res).success) {
    const { rejectDescription, rejectReason } = rejectedResponseSchema.parse(res);
    throw new ValidateOrderError(
      rejectDescription ?? 'Server rejected order',
      rejectReason,
    )
  }

  // fallback: no matching schema
  throw new ValidateOrderError(
    'Unexpected response from server',
    'Unknown error',
  );
}
