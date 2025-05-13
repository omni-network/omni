import * as core from '@omni-network/core'
import { testAssets } from '@omni-network/test-utils'
import { waitFor } from '@testing-library/react'
import { expect, test, vi } from 'vitest'
import { renderHook } from '../../test/index.js'
import { useOmniAssets } from './useOmniAssets.js'

test('default: returns assets when API call succeeds', async () => {
  vi.spyOn(core, 'getAssets').mockResolvedValueOnce(
    testAssets as unknown as core.Asset[],
  )

  const { result } = renderHook(() => useOmniAssets())

  expect(result.current.isPending).toBe(true)

  await waitFor(() => expect(result.current.isPending).toBe(false))

  expect(result.current.isSuccess).toBe(true)
  expect(result.current.data).toEqual(testAssets) // 'assets' should be your mock asset data
})

test('behaviour: handles API error gracefully', async () => {
  vi.spyOn(core, 'getAssets').mockRejectedValueOnce(new Error('mock error'))

  const { result } = renderHook(() => useOmniAssets())

  expect(result.current.isPending).toBe(true)

  await waitFor(() => expect(result.current.isPending).toBe(false))

  expect(result.current.isError).toBe(true)
  expect(result.current.error).toBeDefined()
  expect(result.current.data).toBeUndefined()
})
