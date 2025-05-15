import type {
  DidFillError,
  OrderStatus,
  WatchDidFillError,
} from '@omni-network/core'
import type { Hex } from 'viem'
import { expectTypeOf, test } from 'vitest'
import { useGetOrderStatus } from './useGetOrderStatus.js'

test('type: useGetOrderStatus', () => {
  const result1 = useGetOrderStatus({
    destChainId: 1,
  })

  expectTypeOf(result1).toEqualTypeOf<{
    status: OrderStatus
    error: WatchDidFillError | DidFillError | undefined
    destTxHash: Hex | undefined
    unwatchDestTx: () => void
  }>()

  const result2 = useGetOrderStatus({
    srcChainId: 1,
    destChainId: 2,
    orderId: '0x123' as const,
  })

  expectTypeOf(result2).toEqualTypeOf<{
    status: OrderStatus
    error: WatchDidFillError | DidFillError | undefined
    destTxHash: Hex | undefined
    unwatchDestTx: () => void
  }>()

  const { status, error, destTxHash, unwatchDestTx } = useGetOrderStatus({
    destChainId: 1,
  })

  expectTypeOf(status).toEqualTypeOf<OrderStatus>()
  expectTypeOf(error).toEqualTypeOf<
    WatchDidFillError | DidFillError | undefined
  >()
  expectTypeOf(destTxHash).toEqualTypeOf<Hex | undefined>()
  expectTypeOf(unwatchDestTx).toEqualTypeOf<() => void>()
})
