import { testContracts } from '@omni-network/test-utils'
import { expect, test, vi } from 'vitest'
import * as api from '../internal/api.js'
import { getContracts } from './getContracts.js'

test('default: returns contracts addresses', async () => {
  const fetchJSONSpy = vi.spyOn(api, 'fetchJSON')
  fetchJSONSpy.mockResolvedValueOnce(testContracts)

  await expect(
    getContracts({ environment: 'http://localhost' }),
  ).resolves.toEqual(testContracts)
  expect(fetchJSONSpy).toHaveBeenCalledWith('http://localhost/contracts')
})

test('behaviour: handles invalid response format', async () => {
  vi.spyOn(api, 'fetchJSON').mockResolvedValueOnce({
    invalidField: 'value',
  })

  const expectRejection = expect(async () => {
    await getContracts({ environment: 'http://localhost' })
  }).rejects
  await expectRejection.toBeInstanceOf(Error)
  await expectRejection.toHaveProperty(
    'message',
    'Unexpected /contracts response',
  )
})

test('behaviour: handles invalid response format', async () => {
  const fetchJSONSpy = vi.spyOn(api, 'fetchJSON')
  fetchJSONSpy.mockResolvedValueOnce({
    invalidField: 'value',
  })

  const expectRejection = expect(async () => {
    await getContracts({ environment: 'http://localhost' })
  }).rejects
  await expectRejection.toBeInstanceOf(Error)
  await expectRejection.toHaveProperty(
    'message',
    'Unexpected /contracts response',
  )
})
