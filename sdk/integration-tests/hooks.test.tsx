import { renderHook, waitFor } from '@testing-library/react'
import { describe, expect, test } from 'vitest'

import { useQuote } from '../src/index.js'

import {
  ContextProvider,
  MOCK_L1_ID,
  MOCK_L2_ID,
  ZERO_ADDRESS,
} from './test-utils.js'

describe('useQuote()', () => {
  test('successfully gets a quote', async () => {
    const { result } = renderHook(
      () => {
        return useQuote({
          enabled: true,
          mode: 'expense',
          srcChainId: MOCK_L1_ID,
          destChainId: MOCK_L2_ID,
          deposit: {
            amount: 1n,
            isNative: true,
          },
          expense: {
            amount: 1n,
            isNative: true,
          },
        })
      },
      { wrapper: ContextProvider },
    )

    await waitFor(() => expect(result.current.isSuccess).toBe(true))

    expect(result.current.query.data).toEqual({
      deposit: { token: ZERO_ADDRESS, amount: 1n },
      expense: { token: ZERO_ADDRESS, amount: 0n },
    })
  })
})
