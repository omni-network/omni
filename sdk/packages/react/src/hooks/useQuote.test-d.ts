import type { FetchJSONError, Quote, Quoteable } from '@omni-network/core'
import type { UseQueryResult } from '@tanstack/react-query'
import { expectTypeOf, test } from 'vitest'
import { useQuote } from './useQuote.js'

test('type: useQuote', () => {
  const result = useQuote({
    destChainId: 2,
    mode: 'expense',
    deposit: { amount: 0n, token: '0x00' },
    expense: { token: '0x00' },
    enabled: true,
  })

  expectTypeOf(useQuote).parameter(0).toMatchTypeOf<{
    srcChainId?: number
    destChainId: number
    mode: 'expense' | 'deposit'
    deposit: Quoteable
    expense: Quoteable
    enabled: boolean
  }>()

  expectTypeOf(result.isError).toBeBoolean()
  expectTypeOf(result.isPending).toBeBoolean()
  expectTypeOf(result.isSuccess).toBeBoolean()
  expectTypeOf(result.query).toEqualTypeOf<
    UseQueryResult<Quote, FetchJSONError>
  >()
  expectTypeOf(result.query.data).toEqualTypeOf<Quote | undefined>()
})
