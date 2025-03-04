import { expectTypeOf, test } from 'vitest'
import type { ParseOpenEventError } from '../errors/base.js'
import { useParseOpenEvent } from './useParseOpenEvent.js'

test('type: useParseOpenEvent return', () => {
  const result = useParseOpenEvent({
    status: 'pending',
    logs: [],
  })
  expectTypeOf(result.resolvedOrder).toEqualTypeOf<
    ReturnType<typeof useParseOpenEvent>['resolvedOrder'] | undefined
  >()
  expectTypeOf(result.error).toEqualTypeOf<ParseOpenEventError | undefined>()
})
