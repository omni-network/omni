import type { ResolvedOrder } from '@omni-network/core'
import { expectTypeOf, test } from 'vitest'
import { resolvedOrder } from '../../test/shared.js'
import { type UseDidFillReturn, useDidFill } from './useDidFill.js'

test('type: useDidFill', () => {
  const result = useDidFill({
    destChainId: 1,
    resolvedOrder,
  })

  expectTypeOf(useDidFill).parameter(0).toMatchTypeOf<{
    destChainId: number
    resolvedOrder?: ResolvedOrder
  }>()

  expectTypeOf(result).toEqualTypeOf<UseDidFillReturn>()
})
