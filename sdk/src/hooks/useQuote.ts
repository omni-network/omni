import { type UseMutationResult, useMutation } from '@tanstack/react-query'
import { useCallback, useMemo } from 'react'
import { type Address, zeroAddress } from 'viem'
import type { Deposit, Expense } from '../types/order.js'

type UseQuoteParams = {
  srcChainId: number
  destChainId: number
  mode: 'expense'
  deposit: Deposit
  expense: Omit<Expense, 'spender'>
}

type UseQuoteReturnType = {
  quote: () => Promise<void>
  result?: Quote
  error?: QuoteError
  isPending: boolean
  isError: boolean
  isSuccess: boolean
  mutation: UseMutationResult<Quote, Error, void, unknown>
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
  const mutation = useMutation<Quote>({
    mutationFn: async () => {
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
  })

  const quoteAsync = useCallback(async () => {
    await mutation.mutateAsync()
  }, [mutation.mutateAsync])

  const result = useMemo(() => {
    if (!mutation.data) return

    if (mutation.data.error) {
      return {
        error: {
          code: mutation.data.error.code,
          status: mutation.data.error.status,
          message: mutation.data.error.message,
        },
      }
    }

    if (!mutation.data.deposit || !mutation.data.expense) {
      return {
        error: {
          code: 500,
          status: 'Internal Server Error',
          message: 'Invalid quote response',
        },
      }
    }

    return mutation.data
  }, [mutation.data])

  return {
    quote: quoteAsync,
    result,
    isPending: mutation.isPending,
    isError: mutation.isError,
    isSuccess: mutation.isSuccess,
    mutation,
  }
}
