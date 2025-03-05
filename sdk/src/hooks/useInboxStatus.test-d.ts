import type { Hex } from 'viem'
import { expectTypeOf, test } from 'vitest'
import { orderId } from '../../test/shared.js'
import { useInboxStatus } from './useInboxStatus.js'

test('type: useInboxStatus', () => {
  const result = useInboxStatus({
    chainId: 1,
    orderId,
  })

  expectTypeOf(useInboxStatus).parameter(0).toMatchTypeOf<{
    chainId: number
    orderId?: Hex
  }>()

  expectTypeOf(result).toBeString()
})
