import type {
  FetchJSONError,
  OptionalAbis,
  Order,
  ValidationResponse,
} from '@omni-network/core'
import { validateOrder } from '@omni-network/core'
import { type UseQueryResult, useQuery } from '@tanstack/react-query'
import { useMemo } from 'react'
import { useOmniContext } from '../context/omni.js'
import { hashFn } from '../utils/query.js'

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

  const query = useQuery<ValidationResponse, FetchJSONError>({
    queryKey: ['check', order],
    queryFn: async () => validateOrder({ ...order, environment: apiBaseUrl }),
    queryKeyHashFn: hashFn,
    enabled,
  })

  return useResult(query)
}

const useResult = (
  q: UseQueryResult<ValidationResponse, FetchJSONError>,
): UseValidateOrderResult =>
  useMemo(() => {
    if (q.isError) return { status: 'error', error: q.error }
    if (q.isPending) return { status: 'pending' }
    if (q.data.accepted) return { status: 'accepted' }

    if (q.data.rejected) {
      return {
        status: 'rejected',
        // TODO validation on rejections
        rejectReason: q.data.rejectReason ?? 'Unknown reason',
        rejectDescription:
          q.data.rejectDescription ?? 'No description provided',
      }
    }

    return {
      status: 'error',
      error: q.data?.error ?? {
        code: 0,
        message: 'Unknown validation error',
      },
      q,
    }
  }, [q])
