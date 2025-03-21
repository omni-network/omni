import { type UseQueryResult, useQuery } from '@tanstack/react-query'
import { useMemo } from 'react'
import { type Address, type Hex, fromHex, zeroAddress } from 'viem'
import { useOmniContext } from '../context/omni.js'
import { type FetchJSONError, fetchJSON } from '../internal/api.js'
import type { Quote, Quoteable } from '../types/quote.js'
import { toJSON } from './util.js'

type UseQuoteParams = {
  srcChainId?: number
  destChainId: number
  mode: 'expense' | 'deposit'
  deposit: Quoteable
  expense: Quoteable
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

// QuoteResponse is the response from the /quote endpoint, with hex encoded amounts
type QuoteResponse = {
  deposit: { token: Address; amount: Hex }
  expense: { token: Address; amount: Hex }
}

type QuoteError = FetchJSONError

// useQuote quotes an expense for deposit, or vice versa
export function useQuote(params: UseQuoteParams): UseQuoteResult {
  const { apiBaseUrl } = useOmniContext()
  const { srcChainId, destChainId, deposit, expense, mode, enabled } = params

  const request = toJSON({
    sourceChainId: srcChainId,
    destChainId: destChainId,
    deposit: toQuoteUnit(deposit, mode === 'deposit'),
    expense: toQuoteUnit(expense, mode === 'expense'),
  })

  const query = useQuery<Quote, QuoteError>({
    queryKey: ['quote', request],
    queryFn: async () => doQuote(apiBaseUrl, request),
    enabled,
  })

  return useResult(query)
}

// doQuote calls the /quote endpoint, throwing on error
async function doQuote(apiBaseUrl: string, request: string) {
  const json = await fetchJSON(`${apiBaseUrl}/quote`, {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: request,
  })

  if (!isQuoteRes(json)) {
    throw new Error(`Unexpected quote response: ${JSON.stringify(json)}`)
  }

  const { deposit, expense } = json
  return {
    deposit: { ...deposit, amount: fromHex(deposit.amount, 'bigint') },
    expense: { ...expense, amount: fromHex(expense.amount, 'bigint') },
  } as Quote
}

// toQuoteUnit translates a Quoteable to "QuoteUnit", the format expected by /quote
const toQuoteUnit = (q: Quoteable, omitAmount: boolean) => ({
  amount: omitAmount ? undefined : q.amount,
  token: q.isNative ? zeroAddress : q.token,
})

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

// isQuoteRes checks if a json is a QuoteResponse
// TODO: use zod
function isQuoteRes(json: unknown): json is QuoteResponse {
  const quote = json as QuoteResponse
  return (
    json != null &&
    quote.deposit != null &&
    quote.expense != null &&
    typeof quote.deposit.token === 'string' &&
    typeof quote.deposit.amount === 'string' &&
    typeof quote.expense.token === 'string' &&
    typeof quote.expense.amount === 'string'
  )
}
