import type { FetchJSONError, GetQuoteParams, Quote } from '@omni-network/core'
import { encodeQuoteRequest, getQuoteEncoded } from '@omni-network/core'
import { type UseQueryResult, useQuery } from '@tanstack/react-query'
import { useMemo } from 'react'
import { useOmniContext } from '../context/omni.js'

type UseQuoteParams = GetQuoteParams & {
  enabled: boolean
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

// useQuote quotes an expense for deposit, or vice versa
export function useQuote(params: UseQuoteParams): UseQuoteResult {
  const { apiBaseUrl } = useOmniContext()
  const { enabled, ...quoteParams } = params

  const request = encodeQuoteRequest(quoteParams)
  const query = useQuery<Quote, QuoteError>({
    queryKey: ['quote', request],
    queryFn: async () => await getQuoteEncoded(apiBaseUrl, request),
    enabled,
  })

  return useResult(query)
}

// useResult memoizes a query into a UseQuoteResult
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
