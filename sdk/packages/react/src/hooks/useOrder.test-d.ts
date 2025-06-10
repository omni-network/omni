import type { EVMOrder, OptionalAbis } from '@omni-network/core'
import type { UseMutateFunction } from '@tanstack/react-query'
import type { Hex } from 'viem'
import { expectTypeOf, test } from 'vitest'
import type { Config, UseWaitForTransactionReceiptReturnType } from 'wagmi'
import { type MutationResult, useOrder } from './useOrder.js'
import type { UseValidateOrderResult } from './useValidateOrder.js'

test('type: useOrder', () => {
  const order: EVMOrder<OptionalAbis> & { validateEnabled: boolean } = {
    srcChainId: 1,
    destChainId: 2,
    deposit: { token: '0x123', amount: 100n },
    expense: { token: '0x456', amount: 200n },
    validateEnabled: true,
    calls: [],
  }
  const result = useOrder(order)

  expectTypeOf(result).toMatchTypeOf<{
    open: UseMutateFunction<Hex, Error, void, unknown>
    orderId?: Hex
    validation?: UseValidateOrderResult
    txHash?: Hex
    error?: unknown
    status: string
    isTxPending: boolean
    isTxSubmitted: boolean
    isValidated: boolean
    isOpen: boolean
    isError: boolean
    isReady: boolean
    txMutation: MutationResult
    waitForTx: UseWaitForTransactionReceiptReturnType<Config, number>
  }>()

  expectTypeOf(result.open).toMatchTypeOf<
    UseMutateFunction<`0x${string}`, Error, void, unknown>
  >()

  expectTypeOf(result.orderId).toEqualTypeOf<Hex | undefined>()
  expectTypeOf(result.txHash).toEqualTypeOf<Hex | undefined>()

  expectTypeOf(result.status).toBeString()

  expectTypeOf(result.isTxPending).toBeBoolean()
  expectTypeOf(result.isTxSubmitted).toBeBoolean()
  expectTypeOf(result.isValidated).toBeBoolean()
  expectTypeOf(result.isOpen).toBeBoolean()
  expectTypeOf(result.isError).toBeBoolean()
  expectTypeOf(result.isReady).toBeBoolean()

  expectTypeOf(result.txMutation).toEqualTypeOf<MutationResult>()
  expectTypeOf(result.waitForTx).toEqualTypeOf<
    UseWaitForTransactionReceiptReturnType<Config, number>
  >()
})
