import { waitFor } from '@testing-library/react'
import { expect, test, vi } from 'vitest'
import { renderHook } from '../../test/react.js'
import { contracts } from '../../test/shared.js'
import * as apiModule from '../internal/api.js'
import { useOmniContracts } from './useOmniContracts.js'

const renderOmniContractsHook = (
  options?: Parameters<typeof renderHook>[1],
) => {
  return renderHook(() => useOmniContracts(), options)
}

test('default: returns contracts when API call succeeds', async () => {
  const { result } = renderOmniContractsHook({
    mockContractsCall: true,
  })

  expect(result.current.isPending).toBe(true)

  await waitFor(() => expect(result.current.isPending).toBe(false))

  expect(result.current.isSuccess).toBe(true)
  expect(result.current.data).toEqual(contracts)
})

test('behaviour: handles API error gracefully', async () => {
  const { result } = renderOmniContractsHook({
    mockContractsCallFailure: true,
  })

  expect(result.current.isPending).toBe(true)

  await waitFor(() => expect(result.current.isPending).toBe(false))

  expect(result.current.isError).toBe(true)
  expect(result.current.error).toBeDefined()
  expect(result.current.data).toBeUndefined()
})

test('behaviour: handles invalid response format', async () => {
  const fetchJSONSpy = vi.spyOn(apiModule, 'fetchJSON')
  fetchJSONSpy.mockResolvedValueOnce({
    invalidField: 'value',
  })

  const { result } = renderOmniContractsHook()

  await waitFor(() => expect(result.current.isPending).toBe(false))

  expect(result.current.isError).toBe(true)
  expect(result.current.error).toBeDefined()
  expect(result.current.error?.message).toContain(
    'Unexpected /contracts response',
  )
  expect(result.current.data).toBeUndefined()
})
