import type { UseQueryResult } from '@tanstack/react-query'
import { expectTypeOf, test } from 'vitest'
import type { FetchJSONError } from '../internal/api.js'
import type { Quote, Quoteable } from '../types/quote.js'
import { useQuote } from './useQuote.js'

test('type: useInboxStatus', () => {
  const result = useQuote({
    destChainId: 2,
    mode: 'expense',
    deposit: { isNative: true },
    expense: { isNative: true },
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
