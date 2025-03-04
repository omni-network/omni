import { expectTypeOf, test } from 'vitest'
import { resolvedOrder } from '../../test/shared.js'
import { useDidFill } from './useDidFill.js'

test('type: useDidFill return', () => {
  const result = useDidFill({
    destChainId: 1,
    resolvedOrder,
  })
  // TODO replace return type as it doesnt assert exact type
  expectTypeOf(result).toEqualTypeOf<ReturnType<typeof useDidFill>>()
})
