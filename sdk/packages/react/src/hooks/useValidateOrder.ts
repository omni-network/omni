import {
  type UseQueryResult,
  useQuery,
} from '@tanstack/react-query'
import { useMemo } from 'react'
import { encodeFunctionData, zeroAddress } from 'viem'
import { useOmniContext } from '../context/omni.js'
import { type FetchJSONError, fetchJSON } from '../internal/api.js'
import type { OptionalAbis } from '../types/abi.js'
import type { Order } from '../types/order.js'
import { isContractCall } from '../types/order.js'
import { toJSON } from './util.js'

type UseValidateOrderParams<abis extends OptionalAbis> = {
  order: Order<abis>
  enabled: boolean
}

type ValidationResponse = {
  accepted?: boolean
  rejected?: boolean
  error?: {
    code: number
    message: string
  }
  rejectReason?: string
  rejectDescription?: string
}

type Validation = {
  status: 'pending' | 'rejected' | 'accepted' | 'error'
}

type ValidationPending = Validation & {
  status: 'pending'
}

type ValidationRejected = Validation & {
  status: 'rejected'
  rejectReason: string
  rejectDescription: string
}

type ValidationAccepted = Validation & {
  status: 'accepted'
}

type ValidationError = Validation & {
  status: 'error'
  error:
    | {
        code: number
        message: string
      }
    | FetchJSONError
}

export type UseValidateOrderResult =
  | ValidationPending
  | ValidationRejected
  | ValidationAccepted
  | ValidationError

export function useValidateOrder<abis extends OptionalAbis>({
  order,
  enabled,
}: UseValidateOrderParams<abis>): UseValidateOrderResult {
  const { apiBaseUrl } = useOmniContext()
  const encoded = encodeOrder(order)

  const query = useQuery<ValidationResponse, FetchJSONError>({
    queryKey: ['check', encoded.ok ? encoded.value : 'error'],
    queryFn: async () => {
      if (!encoded.ok) {
        throw encoded.error
      }
      return await doValidate(apiBaseUrl, encoded.value)
    },
    enabled: enabled && encoded.ok,
  })

  return useResult(encoded, query)
}

type EncodeOrderResult =
  | { ok: true; value: string }
  | { ok: false; error: Error }

function encodeOrder<abis extends OptionalAbis>(
  order: Order<abis>,
): EncodeOrderResult {
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

    const value = toJSON({
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
    return { ok: true, value }
  } catch (e) {
    const error =
      e instanceof Error
        ? e
        : new Error((e as { message: string }).message ?? 'Invalid order')
    return { ok: false, error }
  }
}

async function doValidate(apiBaseUrl: string, request: string) {
  const json = await fetchJSON(`${apiBaseUrl}/check`, {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: request,
  })

  if (!isValidateRes(json)) {
    throw new Error(`Unexpected validation response: ${JSON.stringify(json)}`)
  }

  const validation = json

  return validation as ValidationResponse
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

const useResult = (
  encoded: EncodeOrderResult,
  query: UseQueryResult<ValidationResponse, FetchJSONError>,
): UseValidateOrderResult =>
  useMemo(() => {
    if (!encoded.ok) return { status: 'error', error: encoded.error }

    if (query.isError) return { status: 'error', error: query.error }
    if (query.isPending) return { status: 'pending' }
    if (query.data.accepted) return { status: 'accepted' }

    if (query.data.rejected) {
      return {
        status: 'rejected',
        // TODO validation on rejections
        rejectReason: query.data.rejectReason ?? 'Unknown reason',
        rejectDescription:
          query.data.rejectDescription ?? 'No description provided',
      }
    }

    return {
      status: 'error',
      error: query.data?.error ?? {
        code: 0,
        message: 'Unknown validation error',
      },
      query,
    }
  }, [encoded, query])
