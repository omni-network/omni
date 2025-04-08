import { waitFor } from '@testing-library/react'
import { expect, test, vi } from 'vitest'
import { renderHook } from '../../test/react.js'
import { contracts } from '../../test/shared.js'
import * as apiModule from '../internal/api.js'
import { useOmniContracts } from './useOmniContracts.js'

test('default: returns contracts when API call succeeds', async () => {
  const { result } = renderHook(() => useOmniContracts(), {
    mockContractsCall: true,
  })

  expect(result.current.isPending).toBeTruthy()

  await waitFor(() => expect(result.current.isPending).toBeFalsy())

  expect(result.current.isSuccess).toBeTruthy()
  expect(result.current.data).toEqual(contracts)
})

test('behaviour: handles API error gracefully', async () => {
  const { result } = renderHook(() => useOmniContracts(), {
    mockContractsCallFailure: true,
  })

  expect(result.current.isPending).toBeTruthy()

  await waitFor(() => expect(result.current.isPending).toBeFalsy())

  expect(result.current.isError).toBeTruthy()
  expect(result.current.error).toBeDefined()
  expect(result.current.data).toBeUndefined()
})

test('behaviour: handles invalid response format', async () => {
  const fetchJSONSpy = vi.spyOn(apiModule, 'fetchJSON')
  fetchJSONSpy.mockResolvedValueOnce({
    invalidField: 'value',
  })

  const { result } = renderHook(() => useOmniContracts())

  await waitFor(() => expect(result.current.isPending).toBeFalsy())

  expect(result.current.isError).toBeTruthy()
  expect(result.current.error).toBeDefined()
  expect(result.current.error?.message).toContain(
    'Unexpected /contracts response',
  )
  expect(result.current.data).toBeUndefined()
})
