import { type UseQueryResult, useQuery } from '@tanstack/react-query'
import { useMemo } from 'react'
import { type Address, fromHex, toHex, zeroAddress } from 'viem'
import type { Deposit, Expense } from '../types/order.js'

type UseQuoteParams = {
  srcChainId?: number
  destChainId: number
  mode: 'expense' | 'deposit'
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
  const queryEnabled = !!params.srcChainId && !!params.enabled

  const query = useQuery<Quote>({
    queryKey: ['quote'],
    queryFn: async () => {
      const deposit = {
        amount: params.deposit.amount ?? BigInt(0),
        token: params.deposit.isNative ? zeroAddress : params.deposit.token,
      }
      const expense = {
        amount: params.deposit.amount ?? BigInt(0),
        token: params.deposit.isNative ? zeroAddress : params.deposit.token,
      }

      const request = JSON.stringify(
        {
          sourceChainId: params.srcChainId,
          destChainId: params.destChainId,
          deposit,
          expense,
        },
        (_, value) => {
          if (typeof value === 'bigint') {
            return toHex(value)
          }
          return value
        },
      )

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
      const parsed = await response.json()

      return {
        deposit: {
          ...parsed.deposit,
          amount: fromHex(parsed.deposit.amount, 'bigint'),
        },
        expense: {
          ...parsed.expense,
          amount: fromHex(parsed.expense.amount, 'bigint'),
        },
      }
    },
    enabled: queryEnabled,
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
