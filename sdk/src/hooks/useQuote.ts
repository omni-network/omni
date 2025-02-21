import { type UseQueryResult, useQuery } from '@tanstack/react-query'
import { useMemo } from 'react'
import { type Address, zeroAddress } from 'viem'
import type { Deposit, Expense } from '../types/order.js'

type UseQuoteParams = {
  srcChainId: number
  destChainId: number
  mode: 'expense'
  deposit: Deposit
  expense: Omit<Expense, 'spender'>
  enabled?: boolean
}

type UseQuoteReturnType = {
  result?: Quote
  error?: QuoteError
  isPending: boolean
  isError: boolean
  isSuccess: boolean
  query: UseQueryResult<Quote, Error>
}

type Quote = {
  deposit?: {
    token: Address
    amount: bigint
  }
  expense?: {
    token: Address
    amount: bigint
  }
  error?: QuoteError
}

type QuoteError = {
  code: number
  status: string
  message: string
}

export function useQuote(params: UseQuoteParams): UseQuoteReturnType {
  const query = useQuery<Quote>({
    queryKey: ['quote'],
    queryFn: async () => {
      const deposit = {
        amount: params.deposit.amount ?? 0,
        token: params.deposit.isNative ? zeroAddress : params.deposit.token,
      }
      const expense = {
        amount: params.deposit.amount ?? 0,
        token: params.deposit.isNative ? zeroAddress : params.deposit.token,
      }

      const request = JSON.stringify({
        sourceChainId: params.srcChainId,
        destChainId: params.destChainId,
        deposit,
        expense,
      })

      // TODO remove hardcoded api url
      const response = await fetch(
        'https://solver.staging.omni.network/api/v1/quote',
        {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json',
          },
          body: request,
        },
      )
      return await response.json()
    },
    enabled: params.enabled ?? true,
  })

  const result = useMemo(() => {
    if (!query.data) return

    if (query.data.error) {
      return {
        error: {
          code: query.data.error.code,
          status: query.data.error.status,
          message: query.data.error.message,
        },
      }
    }

    if (!query.data.deposit || !query.data.expense) {
      return {
        error: {
          code: 500,
          status: 'Internal Error',
          message: 'Invalid quote response',
        },
      }
    }

    return query.data
  }, [query.data])

  return {
    result,
    isPending: query.isPending,
    isError: query.isError,
    isSuccess: query.isSuccess,
    query,
  }
}
