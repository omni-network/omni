import type {
  FetchJSONError,
  OptionalAbis,
  Order,
  ValidationResponse,
} from '@omni-network/core'
import {
  encodeOrderForValidation,
  validateOrderWithEncoded,
} from '@omni-network/core'
import { type UseQueryResult, useQuery } from '@tanstack/react-query'
import { useMemo } from 'react'
import { useOmniContext } from '../context/omni.js'

type UseValidateOrderParams<abis extends OptionalAbis> = {
  order: Order<abis>
  enabled: boolean
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
      return await validateOrderWithEncoded(apiBaseUrl, encoded.value)
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
    return { ok: true, value: encodeOrderForValidation<abis>(order) }
  } catch (error) {
    return { ok: false, error: error as Error }
  }
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
