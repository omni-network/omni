import type { ParseOpenEventError, ResolvedOrder } from '@omni-network/core'
import type { Log } from 'viem'
import { expectTypeOf, test } from 'vitest'
import type { UseWaitForTransactionReceiptReturnType } from 'wagmi'
import { useParseOpenEvent } from './useParseOpenEvent.js'

test('type: useParseOpenEvent return', () => {
  const result = useParseOpenEvent({
    status: 'pending',
    logs: [],
  })

  expectTypeOf(useParseOpenEvent).parameter(0).toMatchTypeOf<{
    status: UseWaitForTransactionReceiptReturnType['status']
    logs?: Log[]
  }>()

  expectTypeOf(result.resolvedOrder).toMatchTypeOf<ResolvedOrder | undefined>()
  expectTypeOf(result.error).toEqualTypeOf<ParseOpenEventError | undefined>()
})
