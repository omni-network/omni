import type { ParseOpenEventReturn } from '@omni-network/core'
import { expectTypeOf, test } from 'vitest'
import { resolvedOrder } from '../../test/shared.js'
import {
  type UseDidFillOutboxReturn,
  useDidFillOutbox,
} from './useDidFillOutbox.js'

test('type: useDidFillOutbox', () => {
  const result = useDidFillOutbox({
    destChainId: 1,
    resolvedOrder,
  })

  expectTypeOf(useDidFillOutbox).parameter(0).toMatchTypeOf<{
    destChainId: number
    resolvedOrder?: ParseOpenEventReturn
  }>()

  expectTypeOf(result).toEqualTypeOf<UseDidFillOutboxReturn>()
})
