import * as core from '@omni-network/core'
import { testAssets } from '@omni-network/test-utils'
import type { useQuery } from '@tanstack/react-query'
import { waitFor } from '@testing-library/react'
import { expect, test, vi } from 'vitest'
import { renderHook } from '../../test/index.js'
import { useOmniAssets } from './useOmniAssets.js'

const { useQueryMock } = vi.hoisted(() => {
  return { useQueryMock: vi.fn() }
})

vi.mock('@tanstack/react-query', async () => {
  const actual = await vi.importActual('@tanstack/react-query')
  const actualUseQuery = actual.useQuery as typeof useQuery
  return {
    ...actual,
    useQuery: useQueryMock.mockImplementation((params) => {
      return actualUseQuery(params)
    }),
  }
})

test('default: returns assets when API call succeeds', async () => {
  vi.spyOn(core, 'getAssets').mockResolvedValueOnce(
    testAssets as unknown as core.Asset[],
  )

  const { result } = renderHook(() => useOmniAssets())

  expect(result.current.isPending).toBe(true)

  await waitFor(() => expect(result.current.isPending).toBe(false))

  expect(result.current.isSuccess).toBe(true)
  expect(result.current.data).toEqual(testAssets)
})

test('parameters: passes through queryOpts to useQuery', async () => {
  const queryOpts = {
    refetchInterval: 5000,
    staleTime: 10000,
  }
  renderHook(() => useOmniAssets({ queryOpts }), { mockContractsCall: true })
  expect(useQueryMock).toHaveBeenCalledWith(expect.objectContaining(queryOpts))
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
