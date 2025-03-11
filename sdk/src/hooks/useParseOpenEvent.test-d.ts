import type { Hex, Log } from 'viem'
import { expectTypeOf, test } from 'vitest'
import type { UseWaitForTransactionReceiptReturnType } from 'wagmi'
import type { ParseOpenEventError } from '../errors/base.js'
import { useParseOpenEvent } from './useParseOpenEvent.js'

type ResolvedOrder = {
  user: Hex
  originChainId: bigint
  openDeadline: number
  fillDeadline: number
  orderId: Hex
  maxSpent: readonly {
    token: Hex
    amount: bigint
    recipient: Hex
    chainId: bigint
  }[]
  minReceived: readonly {
    token: Hex
    amount: bigint
    recipient: Hex
    chainId: bigint
  }[]
  fillInstructions: readonly {
    destinationChainId: bigint
    destinationSettler: Hex
    originData: Hex
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
