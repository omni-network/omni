import type { Hex } from 'viem'
import { expectTypeOf, test } from 'vitest'
import type { UseReadContractReturnType } from 'wagmi'
import { orderId } from '../../test/shared.js'
import type { inboxABI } from '../constants/abis.js'
import { useGetOrder } from './useGetOrder.js'

test('type: useGetOrder', () => {
  const result = useGetOrder({
    chainId: 1,
    orderId: orderId,
  })

  expectTypeOf(useGetOrder).parameter(0).toMatchTypeOf<{
    chainId?: number
    orderId?: Hex
  }>()

  expectTypeOf(result).toEqualTypeOf<
    UseReadContractReturnType<typeof inboxABI>
  >()
})
