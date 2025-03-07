import type { Log } from 'viem'
import { expectTypeOf, test } from 'vitest'
import type { UseWaitForTransactionReceiptReturnType } from 'wagmi'
import type { ParseOpenEventError } from '../errors/base.js'
import { useParseOpenEvent } from './useParseOpenEvent.js'

type ResolvedOrder = {
  user: `0x${string}`
  originChainId: bigint
  openDeadline: number
  fillDeadline: number
  orderId: `0x${string}`
  maxSpent: readonly {
    token: `0x${string}`
    amount: bigint
    recipient: `0x${string}`
    chainId: bigint
  }[]
  minReceived: readonly {
    token: `0x${string}`
    amount: bigint
    recipient: `0x${string}`
    chainId: bigint
  }[]
  fillInstructions: readonly {
    destinationChainId: bigint
    destinationSettler: `0x${string}`
    originData: `0x${string}`
  }[]
}

test('type: useParseOpenEvent return', () => {
  const result = useParseOpenEvent({
    status: 'pending',
    logs: [],
  })

  expectTypeOf(useParseOpenEvent).parameter(0).toMatchTypeOf<{
    status: UseWaitForTransactionReceiptReturnType['status']
    logs?: Log[]
  }>()

  expectTypeOf(result.resolvedOrder).toMatchTypeOf<ResolvedOrder | undefined>()
  expectTypeOf(result.error).toEqualTypeOf<ParseOpenEventError | undefined>()
})
