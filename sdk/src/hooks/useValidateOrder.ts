import { type UseQueryResult, useQuery } from '@tanstack/react-query'
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

  function _encodeCalls() {
    return order.calls.map((call) => {
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
  }

  const request = toJSON({
    orderId: order.owner ?? zeroAddress,
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

  const query = useQuery<ValidationResponse, FetchJSONError>({
    queryKey: ['check', request],
    queryFn: async () => doValidate(apiBaseUrl, request),
    enabled,
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
  query: UseQueryResult<ValidationResponse, FetchJSONError>,
): UseValidateOrderResult =>
  useMemo(() => {
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
  }, [query])
