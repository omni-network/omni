import { encodeFunctionData, zeroAddress } from 'viem'
import { z } from 'zod/v4-mini'
import { ValidateOrderError } from '../errors/base.js'
import { fetchJSON } from '../internal/api.js'
import type { OptionalAbis } from '../types/abi.js'
import type { Environment } from '../types/config.js'
import { type EVMOrder, isContractCall } from '../types/order.js'
import { getApiUrl, toJSON } from '../utils/index.js'

const traceSchema = z.union([z.record(z.string(), z.unknown()), z.null()])

const acceptedResponseSchema = z.object({
  accepted: z.literal(true),
  rejectCode: z.optional(z.literal(0)),
  rejected: z.optional(z.literal(false)),
  rejectReason: z.optional(z.literal('')),
  rejectDescription: z.optional(z.literal('')),
  trace: z.optional(traceSchema),
})

const rejectedResponseSchema = z.object({
  accepted: z.optional(z.literal(false)),
  rejectCode: z.optional(z.number()),
  rejected: z.literal(true),
  rejectReason: z.string(),
  rejectDescription: z.string(),
  trace: z.optional(traceSchema),
})

const errorResponseSchema = z.object({
  error: z.object({
    code: z.number(),
    message: z.string(),
  }),
})

export type ValidateOrderParameters<abis extends OptionalAbis> =
  EVMOrder<abis> & {
    debug?: boolean
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
  const { environment, ...orderWithDebug } = params
  const apiUrl = getApiUrl(environment)
  const json = await fetchJSON(`${apiUrl}/check`, {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: serialize(orderWithDebug),
  })

  if (!isValidateRes(json)) {
    throw new Error(`Unexpected validation response: ${JSON.stringify(json)}`)
  }

  return json satisfies ValidationResponse
}

const serialize = <abis extends OptionalAbis>(
  order: EVMOrder<abis> & { debug?: boolean },
) => {
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
      debug: order.debug ?? false,
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

// asserts a json response is AcceptedResult
export function isAcceptedRes(json: unknown): json is AcceptedResult {
  return acceptedResponseSchema.safeParse(json).success
}

// asserts a json response is RejectedResult
export function isRejectedRes(
  json: unknown,
): json is z.infer<typeof rejectedResponseSchema> {
  return rejectedResponseSchema.safeParse(json).success
}

// asserts a json response is ErrorResult
export function isErrorRes(
  json: unknown,
): json is z.infer<typeof errorResponseSchema> {
  return errorResponseSchema.safeParse(json).success
}

export type AcceptedResult = z.infer<typeof acceptedResponseSchema>

export function assertAcceptedResult(
  res: ValidationResponse,
): asserts res is AcceptedResult {
  // if the response is accepted
  if (isAcceptedRes(res)) {
    return
  }

  // if the response is an error
  if (isErrorRes(res)) {
    const { error } = errorResponseSchema.parse(res)
    throw new ValidateOrderError(error.message, `Code ${error.code}`)
  }

  // if the response is rejected
  if (isRejectedRes(res)) {
    const { rejectDescription, rejectReason } =
      rejectedResponseSchema.parse(res)
    throw new ValidateOrderError(
      rejectDescription ?? 'Server rejected order',
      rejectReason,
    )
  }

  // fallback: no matching schema
  throw new ValidateOrderError(
    'Unexpected response from server',
    'Unknown error',
  )
}
