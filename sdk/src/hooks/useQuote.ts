import { useMemo } from 'react'
import { type Hex } from 'viem'
import { type UseQueryResult, useQuery } from '@tanstack/react-query'
import { type Address, fromHex, zeroAddress } from 'viem'
import { type FetchJSONError, fetchJSON } from '../internal/api.js'
import { toJSON } from './util.js'

// TODO add complex type to enforce one of the amounts is defined
type Quoteable =
  | { isNative: true; token?: never; amount?: bigint }
  | { isNative: false; token: Address; amount?: bigint }

type UseQuoteParams = {
  srcChainId?: number
  destChainId: number
  mode: 'expense' | 'deposit'
  deposit: Quoteable
  expense: Quoteable
  enabled?: boolean
}

type UseQuoteSuccess = {
  quote: Quote
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

type Quote = {
  deposit: { token: Address; amount: bigint }
  expense: { token: Address; amount: bigint }
}

// QuoteResponse is the response from the /quote endpoint, with hex encoded amounts
type QuoteResponse = {
  deposit: { token: Address; amount: Hex }
  expense: { token: Address; amount: Hex }
}

type QuoteError = FetchJSONError

// useQuote quotes an expense for deposit, or vice versa
export function useQuote(params: UseQuoteParams): UseQuoteResult {
  // TODO: move to context
  const apiBaseUrl = 'https://solver.staging.omni.network/api/v1'

  // biome-ignore lint/correctness/useExhaustiveDependencies: deep compare on obj properties
  const request = useMemo(() => {
    return toJSON({
      sourceChainId: params.srcChainId,
      destChainId: params.destChainId,
      deposit: toQuoteUnit(params.deposit, params.mode === 'deposit'),
      expense: toQuoteUnit(params.expense, params.mode === 'expense'),
    })
  }, [
    params.srcChainId,
    params.destChainId,
    params.deposit.isNative,
    params.deposit.amount,
    params.deposit.token,
    params.expense.isNative,
    params.expense.amount,
    params.expense.token,
    params.mode,
  ])

  const query = useQuery<Quote, QuoteError>({
    queryKey: ['quote', request],
    queryFn: async () => doQuote(apiBaseUrl, request),
    enabled:
      !!params.srcChainId &&
      params.enabled !== false &&
      !(
        (params.deposit.amount ?? 0n) > 0n || (params.expense.amount ?? 0n) > 0n
      ),
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
      }
    }

    if (q.isSuccess) {
      return {
        quote: q.data,
        isPending: false,
        isError: false,
        isSuccess: true,
        query: q,
      }
    }

    return {
      isPending: true,
      isError: false,
      isSuccess: false,
      query: q,
    }
  }, [q])

// isQuoteRes checks if a json is a QuoteResponse
// TODO: use zod
function isQuoteRes(json: unknown): json is QuoteResponse {
  return (
    json != null &&
    (json as any).deposit != null &&
    (json as any).expense != null &&
    typeof (json as any).deposit.token === 'string' &&
    typeof (json as any).deposit.amount === 'string' &&
    typeof (json as any).expense.token === 'string' &&
    typeof (json as any).expense.amount === 'string'
  )
}
