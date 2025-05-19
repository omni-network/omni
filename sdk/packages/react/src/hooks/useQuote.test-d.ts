import type { FetchJSONError, Quote } from '@omni-network/core'
import type { UseQueryResult } from '@tanstack/react-query'
import { expectTypeOf, test } from 'vitest'
import { type UseQuoteParams, useQuote } from './useQuote.js'

test('type: useQuote', () => {
  const result = useQuote({
    destChainId: 2,
    mode: 'expense',
    deposit: { amount: 0n, token: '0x00' },
    expense: { token: '0x00' },
    enabled: true,
    environment: 'devnet',
    queryOpts: {
      refetchInterval: 100,
    },
  })

  expectTypeOf(useQuote).parameter(0).toMatchTypeOf<UseQuoteParams>()

  expectTypeOf(result.isError).toBeBoolean()
  expectTypeOf(result.isPending).toBeBoolean()
  expectTypeOf(result.isSuccess).toBeBoolean()
  expectTypeOf(result.query).toEqualTypeOf<
    UseQueryResult<Quote, FetchJSONError>
  >()
  expectTypeOf(result.query.data).toEqualTypeOf<Quote | undefined>()
})
