import { expectTypeOf, test } from 'vitest'
import type { UseReadContractReturnType } from 'wagmi'
import { resolvedOrder } from '../../test/shared.js'
import type { outboxABI } from '../constants/abis.js'
import { useDidFill } from './useDidFill.js'
import type { useParseOpenEvent } from './useParseOpenEvent.js'

test('type: useDidFill', () => {
  const result = useDidFill({
    destChainId: 1,
    resolvedOrder,
  })

  expectTypeOf(useDidFill).parameter(0).toMatchTypeOf<{
    destChainId: number
    resolvedOrder?: ReturnType<typeof useParseOpenEvent>['resolvedOrder']
  }>()

  expectTypeOf(result).toEqualTypeOf<
    UseReadContractReturnType<typeof outboxABI>
  >()
})
