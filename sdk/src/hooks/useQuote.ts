import { type UseQueryResult, useQuery } from '@tanstack/react-query'
import { useMemo } from 'react'
import { type Address, fromHex, zeroAddress } from 'viem'
import { toJSON } from './util.js'

type Quoteable =
  | { isNative: true; token?: never; amount: bigint }
  | { isNative: false; token: Address; amount: bigint }

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

type QuoteError = { code: number; status: string; message: string }

// useQuote quotes an expense for deposit, or vice versa[
export function useQuote(p: UseQuoteParams): UseQuoteResult {
  // TODO: move to context
  const apiBaseUrl = 'https://solver.staging.omni.network/api/v1/check'

  const request = useMemo(() => {
    return toJSON({
      sourceChainId: p.srcChainId,
      destChainId: p.destChainId,
      deposit: toQuoteUnit(p.deposit, p.mode == 'deposit'),
      expense: toQuoteUnit(p.expense, p.mode == 'expense'),
    })
  }, [
    p.srcChainId,
    p.destChainId,
    p.deposit.amount,
    p.deposit.isNative,
    p.deposit.token,
    p.expense.amount,
    p.expense.isNative,
    p.expense.token,
  ])

  const enabled = !p.srcChainId || p.enabled
  const query = useQuery<Quote, QuoteError>({
    queryKey: ['quote', request],
    queryFn: async () => doQuote(apiBaseUrl, request),
    enabled,
  })

  return useResult(query)
}

// doQuote calls the /quote endpoint, throwing on error
async function doQuote(apiBaseUrl: string, request: string) {
  const response = await fetch(apiBaseUrl + '/quote', {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: request,
  })

  if (!response.ok) {
    const { error } = await response.json()
    throw error as QuoteError
  }

  const { deposit, expense } = await response.json()

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
