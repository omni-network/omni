import type { Hex } from 'viem'
import { expectTypeOf, test } from 'vitest'
import type { ParseOpenEventError } from '../errors/base.js'
import { useParseOpenEvent } from './useParseOpenEvent.js'

test('select data', () => {
  const result = useParseOpenEvent({
    status: 'pending',
    logs: [],
  })
  expectTypeOf(result.orderId).toEqualTypeOf<Hex | undefined>()
  expectTypeOf(result.originData).toEqualTypeOf<Hex | undefined>()
  expectTypeOf(result.error).toEqualTypeOf<ParseOpenEventError | undefined>()
})
