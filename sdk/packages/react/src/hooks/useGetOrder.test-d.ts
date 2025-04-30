import type { GetOrderReturn } from '@omni-network/core'
import type { UseQueryResult } from '@tanstack/react-query'
import type { Hex } from 'viem'
import { expectTypeOf, test } from 'vitest'
import { orderId } from '../../test/shared.js'
import { useGetOrder } from './useGetOrder.js'

test('type: useGetOrder', () => {
  const result = useGetOrder({
    chainId: 1,
    orderId: orderId,
  })

  expectTypeOf(useGetOrder).parameter(0).toMatchObjectType<{
    chainId?: number
    orderId?: Hex
  }>()

  expectTypeOf(result).toEqualTypeOf<UseQueryResult<GetOrderReturn>>()
})
