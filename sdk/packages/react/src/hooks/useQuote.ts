import type {
  FetchJSONError,
  GetQuoteParameters,
  Quote,
} from '@omni-network/core'
import { getQuote } from '@omni-network/core'
import {
  type UseQueryOptions,
  type UseQueryResult,
  useQuery,
} from '@tanstack/react-query'
import { useMemo } from 'react'
import { useOmniContext } from '../context/omni.js'
import { hashFn } from '../utils/query.js'

type UseQuoteParams = GetQuoteParameters & {
  enabled: boolean
  queryOpts?: Omit<
    UseQueryOptions<Quote, QuoteError>,
    'queryKey' | 'queryFn' | 'enabled'
  >
}

type UseQuoteSuccess = Quote & {
  isPending: false
  isError: false
  isSuccess: true
}

type UseQuoteError = {
  error: QuoteError
  isPending: false
  isError: true
  isSuccess: false
}

type UseQuotePending = {
  isPending: true
  isError: false
  isSuccess: false
}

type UseQuoteResult = (UseQuoteSuccess | UseQuoteError | UseQuotePending) & {
  query: UseQueryResult<Quote, QuoteError>
}

type QuoteError = FetchJSONError

// quotes expense amount for a given deposit, or vice versa
export function useQuote(params: UseQuoteParams): UseQuoteResult {
  const { apiBaseUrl } = useOmniContext()
  const { enabled, ...quoteParams } = params
  const query = useQuery<Quote, QuoteError>({
    ...params.queryOpts,
    queryKey: ['quote', quoteParams],
    queryFn: async () => getQuote({ ...quoteParams, environment: apiBaseUrl }),
    queryKeyHashFn: hashFn,
    enabled,
  })

  return useResult(query)
}

// memoizes query result as UseQuoteResult
const useResult = (q: UseQueryResult<Quote, QuoteError>): UseQuoteResult =>
  useMemo(() => {
    if (q.isError) {
      return {
        error: q.error,
        isPending: false,
        isError: true,
        isSuccess: false,
        query: q,
      } as const
    }

    if (q.isSuccess) {
      return {
        ...q.data,
        isPending: false,
        isError: false,
        isSuccess: true,
        query: q,
      } as const
    }

    return {
      isPending: true,
      isError: false,
      isSuccess: false,
      query: q,
    } as const
  }, [q])
