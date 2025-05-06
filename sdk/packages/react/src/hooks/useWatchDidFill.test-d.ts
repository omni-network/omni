import { expectTypeOf, test } from 'vitest'
import {
  type UseWatchDidFillParams,
  type UseWatchDidFillReturn,
  useWatchDidFill,
} from './useWatchDidFill.js'

test('type: useWatchDidFill', () => {
  // Mock onError function for the test
  const mockOnError = (error: Error) => {
    console.error('Error received:', error)
  }

  const result = useWatchDidFill({
    destChainId: 1,
    orderId: '0x123',
    pollingInterval: 1000,
    onError: mockOnError,
  })

  expectTypeOf(useWatchDidFill)
    .parameter(0)
    .toMatchTypeOf<UseWatchDidFillParams>()

  expectTypeOf(result).toEqualTypeOf<UseWatchDidFillReturn>()
})
