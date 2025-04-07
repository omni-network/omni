import type { outboxABI } from '@omni-network/core'
import { expectTypeOf, test } from 'vitest'
import type { UseReadContractReturnType } from 'wagmi'
import { resolvedOrder } from '../../test/shared.js'
import { useDidFillOutbox } from './useDidFillOutbox.js'
import type { useParseOpenEvent } from './useParseOpenEvent.js'

test('type: useDidFillOutbox', () => {
  const result = useDidFillOutbox({
    destChainId: 1,
    resolvedOrder,
  })

  expectTypeOf(useDidFillOutbox).parameter(0).toMatchTypeOf<{
    destChainId: number
    resolvedOrder?: ReturnType<typeof useParseOpenEvent>['resolvedOrder']
  }>()

  expectTypeOf(result).toEqualTypeOf<
    UseReadContractReturnType<typeof outboxABI>
  >()
})
