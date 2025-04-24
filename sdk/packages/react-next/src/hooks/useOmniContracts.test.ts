import { waitFor } from '@testing-library/react'
import { expect, test } from 'vitest'
import { contracts, renderHook } from '../../test/index.js'
import { useOmniContracts } from './useOmniContracts.js'

test('default: returns contracts when API call succeeds', async () => {
  const { result } = renderHook(() => useOmniContracts(), {
    mockContractsCall: true,
  })

  expect(result.current.isPending).toBe(true)

  await waitFor(() => expect(result.current.isPending).toBe(false))

  expect(result.current.isSuccess).toBe(true)
  expect(result.current.data).toEqual(contracts)
})

test('behaviour: handles API error gracefully', async () => {
  const { result } = renderHook(() => useOmniContracts(), {
    mockContractsCallFailure: true,
  })

  expect(result.current.isPending).toBe(true)

  await waitFor(() => expect(result.current.isPending).toBe(false))

  expect(result.current.isError).toBe(true)
  expect(result.current.error).toBeDefined()
  expect(result.current.data).toBeUndefined()
})
