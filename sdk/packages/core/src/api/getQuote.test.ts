import { testQuote } from '@omni-network/test-utils'
import { zeroAddress } from 'viem'
import { beforeEach, expect, test, vi } from 'vitest'
import * as api from '../internal/api.js'
import { type GetQuoteParameters, getQuote } from './getQuote.js'

const token = '0x123'

// Server response matching the testQuote object with string amounts
const testQuoteResponse = {
  deposit: { token: zeroAddress, amount: '100' },
  expense: { token: zeroAddress, amount: '99' },
} as const

beforeEach(() => {
  vi.spyOn(api, 'fetchJSON').mockResolvedValue(testQuoteResponse)
})

const params: GetQuoteParameters = {
  srcChainId: 1,
  destChainId: 2,
  mode: 'expense',
  deposit: { amount: 100n },
}

test('default: fetches a quote', async () => {
  await expect(getQuote(params)).resolves.toEqual(testQuote)
})

test('parameters: expense', async () => {
  await expect(
    getQuote({
      ...params,
      mode: 'deposit',
      expense: { token, amount: 100n },
    }),
  ).resolves.toEqual(testQuote)

  await expect(
    getQuote({
      ...params,
      mode: 'deposit',
      expense: { amount: 100n },
    }),
  ).resolves.toEqual(testQuote)
})

test('parameters: deposit', async () => {
  await expect(
    getQuote({
      ...params,
      deposit: { token, amount: 100n },
    }),
  ).resolves.toEqual(testQuote)

  await expect(
    getQuote({
      ...params,
      deposit: { amount: 100n },
    }),
  ).resolves.toEqual(testQuote)
})

test('parameters: mode', async () => {
  await expect(
    getQuote({
      ...params,
      mode: 'expense',
      deposit: { amount: 100n },
    }),
  ).resolves.toEqual(testQuote)

  await expect(
    getQuote({
      ...params,
      mode: 'deposit',
      expense: { amount: 100n },
    }),
  ).resolves.toEqual(testQuote)
})

test.each([
  'test',
  {},
  { deposit: { amount: '100' } },
  { expense: { amount: '100' } },
  { deposit: { token: zeroAddress }, expense: { token: zeroAddress } },
  { deposit: { amount: '100' }, expense: { amount: '99' } },
])('behaviour: throws if response is not a quote: %s', async (mockReturn) => {
  vi.spyOn(api, 'fetchJSON').mockResolvedValue(mockReturn)

  const expectRejection = expect(async () => {
    await getQuote(params)
  }).rejects
  await expectRejection.toBeInstanceOf(Error)
  await expectRejection.toHaveProperty(
    'message',
    `Unexpected quote response: ${JSON.stringify(mockReturn)}`,
  )
})
