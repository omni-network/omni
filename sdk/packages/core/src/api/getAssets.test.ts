import { testAssets } from '@omni-network/test-utils'
import { fromHex } from 'viem'
import { expect, vi } from 'vitest'
import { test } from 'vitest'
import * as api from '../internal/api.js'
import { getAssets } from './getAssets.js'

const expected = testAssets.map((asset) => ({
  ...asset,
  expenseMin: fromHex(asset.expenseMin, 'bigint'),
  expenseMax: fromHex(asset.expenseMax, 'bigint'),
}))

test('default: returns assets with parsed hex values', async () => {
  const fetchJSONSpy = vi.spyOn(api, 'fetchJSON')
  fetchJSONSpy.mockResolvedValueOnce({
    tokens: testAssets,
  })

  await expect(getAssets({ environment: 'http://localhost' })).resolves.toEqual(
    expected,
  )
  expect(fetchJSONSpy).toHaveBeenCalledWith('http://localhost/tokens')
})

test('behaviour: supports SVM assets', async () => {
  const tokens = [
    {
      enabled: true,
      name: 'USD Coin',
      symbol: 'UDSC',
      chainId: 352,
      address: '6AfTP38VbvM3gxMRy8ubvb7NbhzhjVB8VGMy2uBby1pY',
      decimals: 18,
      expenseMin: '0x1000000000000000000',
      expenseMax: '0x1000000000000000000',
    } as const,
  ]

  const fetchJSONSpy = vi.spyOn(api, 'fetchJSON')
  fetchJSONSpy.mockResolvedValueOnce({ tokens })

  await expect(getAssets({ environment: 'http://localhost' })).resolves.toEqual(
    tokens.map((token) => ({
      ...token,
      expenseMin: fromHex(token.expenseMin, 'bigint'),
      expenseMax: fromHex(token.expenseMax, 'bigint'),
    })),
  )
})

test('behaviour: handles invalid response format (incorrect type)', async () => {
  vi.spyOn(api, 'fetchJSON').mockResolvedValueOnce({
    invalidField: 'value',
  })

  const rejection = expect(async () => {
    await getAssets({ environment: 'http://localhost' })
  }).rejects
  await rejection.toBeInstanceOf(Error)
  await rejection.toHaveProperty('message', 'Unexpected /tokens response')
})

test('behaviour: handles invalid response format (missing fields)', async () => {
  const fetchJSONSpy = vi.spyOn(api, 'fetchJSON')
  fetchJSONSpy.mockResolvedValueOnce([
    {
      enabled: true,
      name: 'Ether',
      symbol: 'ETH',
      chainId: 1,
      address: '0x123',
      decimals: 18,
    },
  ])

  const rejection = expect(async () => {
    await getAssets({ environment: 'http://localhost' })
  }).rejects
  await rejection.toBeInstanceOf(Error)
  await rejection.toHaveProperty('message', 'Unexpected /tokens response')
})

test('behaviour: handles invalid response format (invalid hex)', async () => {
  const fetchJSONSpy = vi.spyOn(api, 'fetchJSON')
  fetchJSONSpy.mockResolvedValueOnce([
    {
      ...testAssets[0],
      expenseMin: 'abcd',
    },
  ])

  const rejection = expect(async () => {
    await getAssets({ environment: 'http://localhost' })
  }).rejects
  await rejection.toBeInstanceOf(Error)
})
