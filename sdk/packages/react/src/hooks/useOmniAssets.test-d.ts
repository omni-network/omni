import type { Asset } from '@omni-network/core'
import type { UseQueryResult } from '@tanstack/react-query'
import { expectTypeOf, test } from 'vitest'
import { type UseOmniAssetsResult, useOmniAssets } from './useOmniAssets.js'

test('type: useOmniAssets', () => {
  const result = useOmniAssets()
  expectTypeOf(result).toEqualTypeOf<UseOmniAssetsResult>()
  expectTypeOf(result).toEqualTypeOf<UseQueryResult<Asset[]>>()

  expectTypeOf(result.data).toEqualTypeOf<Asset[] | undefined>()

  if (result.data) {
    expectTypeOf(result.data[0].enabled).toBeBoolean()
    expectTypeOf(result.data[0].name).toBeString()
    expectTypeOf(result.data[0].symbol).toBeString()
    expectTypeOf(result.data[0].chainId).toBeNumber()
    expectTypeOf(result.data[0].address).toBeString()
    expectTypeOf(result.data[0].decimals).toBeNumber()
    expectTypeOf(result.data[0].expenseMin).toBeBigInt()
    expectTypeOf(result.data[0].expenseMax).toBeBigInt()
  }
})
