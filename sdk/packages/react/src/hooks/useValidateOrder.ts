import type {
  EVMOrder,
  FetchJSONError,
  OptionalAbis,
  ValidationResponse,
} from '@omni-network/core'
import { isAcceptedRes, isRejectedRes, validateOrder } from '@omni-network/core'
import { type UseQueryResult, useQuery } from '@tanstack/react-query'
import { useMemo } from 'react'
import { useOmniContext } from '../context/omni.js'
import { hashFn } from '../utils/query.js'
import type { QueryOpts } from './types.js'

type UseValidateOrderParams<abis extends OptionalAbis> = {
  order: EVMOrder<abis>
  enabled: boolean
  debug?: boolean
  queryOpts?: QueryOpts<ValidationResponse, FetchJSONError>
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
  trace?: Record<string, unknown>
}

type ValidationAccepted = Validation & {
  status: 'accepted'
  trace?: Record<string, unknown>
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
  debug,
  queryOpts,
}: UseValidateOrderParams<abis>): UseValidateOrderResult {
  const { apiBaseUrl } = useOmniContext()

  const query = useQuery<ValidationResponse, FetchJSONError>({
    ...queryOpts,
    queryKey: ['check', order],
    queryFn: async () => {
      return await validateOrder({ ...order, debug, environment: apiBaseUrl })
    },
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
    if (isAcceptedRes(q.data)) {
      return { status: 'accepted', trace: q.data.trace ?? undefined }
    }
    if (isRejectedRes(q.data)) {
      return {
        status: 'rejected',
        // TODO validation on rejections
        rejectReason: q.data.rejectReason ?? 'Unknown reason',
        rejectDescription:
          q.data.rejectDescription ?? 'No description provided',
        trace: q.data.trace ?? undefined,
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
