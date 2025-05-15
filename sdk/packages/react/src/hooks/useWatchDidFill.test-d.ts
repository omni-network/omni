import { expectTypeOf, test } from 'vitest'
import {
  type UseWatchDidFillParams,
  type UseWatchDidFillReturn,
  useWatchDidFill,
} from './useWatchDidFill.js'

test('type: useWatchDidFill', () => {
  const result = useWatchDidFill({
    destChainId: 1,
    orderId: '0x123',
    pollingInterval: 1000,
  })

  expectTypeOf(useWatchDidFill)
    .parameter(0)
    .toMatchTypeOf<UseWatchDidFillParams>()

  expectTypeOf(result).toEqualTypeOf<UseWatchDidFillReturn>()
})
