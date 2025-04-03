import type { Hex } from 'viem'
import { expectTypeOf, test } from 'vitest'
import type { DidFillError, GetOrderError } from '../errors/base.js'
import type { OrderStatus } from '../types/order.js'
import { useGetOrderStatus } from './useGetOrderStatus.js'
import type { useParseOpenEvent } from './useParseOpenEvent.js'

test('type: useGetOrderStatus', () => {
  const result1 = useGetOrderStatus({
    destChainId: 1,
  })

  expectTypeOf(result1).toEqualTypeOf<{
    status: OrderStatus
    error: GetOrderError | DidFillError | undefined
  }>()

  const mockOrderId = '0x123' as Hex
  const mockResolvedOrder = {} as ReturnType<
    typeof useParseOpenEvent
  >['resolvedOrder']

  const result2 = useGetOrderStatus({
    srcChainId: 1,
    destChainId: 2,
    orderId: mockOrderId,
    resolvedOrder: mockResolvedOrder,
  })

  expectTypeOf(result2).toEqualTypeOf<{
    status: OrderStatus
    error: GetOrderError | DidFillError | undefined
  }>()

  const { status, error } = useGetOrderStatus({ destChainId: 1 })

  expectTypeOf(status).toEqualTypeOf<OrderStatus>()
  expectTypeOf(error).toEqualTypeOf<GetOrderError | DidFillError | undefined>()
})
