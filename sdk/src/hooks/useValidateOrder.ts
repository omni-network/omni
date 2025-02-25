import { type UseQueryResult, useQuery } from '@tanstack/react-query'
import { useMemo } from 'react'
import { encodeFunctionData } from 'viem'
import { type FetchJSONError, fetchJSON } from '../internal/api.js'
import type { Order } from '../types/order.js'
import { toJSON } from './util.js'

type UseValidateOrderParams = {
  order: Order
  enabled?: boolean
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

export function useValidateOrder({
  order,
  enabled = true,
}: UseValidateOrderParams): UseValidateOrderResult {
  const apiBaseUrl = 'https://solver.staging.omni.network/api/v1'

  // biome-ignore lint/correctness/useExhaustiveDependencies: deep compare on obj properties
  const request = useMemo(() => {
    if (!order.owner) return ''

    function _encodeCalls() {
      return order.calls.map((call) => {
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
    }

    return toJSON({
      sourceChainId: order.srcChainId,
      destChainId: order.destChainId,
      fillDeadline: order.fillDeadline ?? Math.floor(Date.now() / 1000 + 86400),
      calls: _encodeCalls(),
      deposit: {
        amount: order.deposit.amount,
        token: order.deposit.token,
      },
      expenses: [
        {
          amount: order.expense.amount,
          token: order.expense.token,
          spender: order.expense.spender,
        },
      ],
    })
  }, [
    order.srcChainId,
    order.destChainId,
    order.deposit.amount?.toString(),
    order.deposit.token,
    order.expense.amount?.toString(),
    order.expense.token,
    order.expense.spender,
    order.owner,
    order.fillDeadline,
  ])

  const query = useQuery<ValidationResponse, FetchJSONError>({
    queryKey: ['check', request],
    queryFn: async () => doValidate(apiBaseUrl, request),
    enabled:
      !!order.owner &&
      !!order.srcChainId &&
      enabled !== false &&
      !(order.deposit.amount === 0n && order.expense.amount === 0n),
  })

  return useResult(query)
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
  q: UseQueryResult<ValidationResponse, FetchJSONError>,
): UseValidateOrderResult =>
  useMemo(() => {
    if (q.isError) {
      return {
        status: 'error',
        error: q.error,
      }
    }

    if (q.isPending) {
      return {
        status: 'pending',
      }
    }

    if (q.data.rejected) {
      return {
        status: 'rejected',
        // TODO validation on rejections
        rejectReason: q.data.rejectReason ?? 'Unknown reason',
        rejectDescription:
          q.data.rejectDescription ?? 'No description provided',
      }
    }

    if (q.data.accepted) {
      return {
        status: 'accepted',
      }
    }

    return {
      status: 'error',
      error: q.data?.error ?? {
        code: 0,
        message: 'Unknown validation error',
      },
    }
  }, [q])
