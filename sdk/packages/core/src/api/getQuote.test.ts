import { testQuote } from '@omni-network/test-utils'
import { type AsyncResult, Result } from 'typescript-result'
import { zeroAddress } from 'viem'
import { beforeEach, expect, test, vi } from 'vitest'
import * as api from '../internal/api.js'
import type { Quoteable } from '../types/quote.js'
import { type GetQuoteParams, getQuote } from './getQuote.js'

const token = '0x123'
const deposit = { token, isNative: false } satisfies Quoteable
const nativeExpense = { isNative: true } satisfies Quoteable

function asyncResult<T>(data: T): AsyncResult<T, never> {
  return Result.fromAsync(Promise.resolve(Result.ok(data)))
}

// Server response matching the testQuote object with string amounts
const testQuoteResponse = {
  deposit: { token: zeroAddress, amount: '100' },
  expense: { token: zeroAddress, amount: '99' },
} as const

beforeEach(() => {
  vi.spyOn(api, 'safeFetchJSON').mockReturnValue(asyncResult(testQuoteResponse))
})

const params: GetQuoteParams = {
  srcChainId: 1,
  destChainId: 2,
  mode: 'expense',
  deposit: deposit,
  expense: nativeExpense,
}

test('default: fetches a quote', async () => {
  await expect(getQuote(params)).resolves.toEqual(testQuote)
})

test('parameters: expense', async () => {
  await expect(
    getQuote({
      ...params,
      expense: { token, isNative: false },
    }),
  ).resolves.toEqual(testQuote)

  await expect(
    getQuote({
      ...params,
      expense: { isNative: true },
    }),
  ).resolves.toEqual(testQuote)
})

test('parameters: deposit', async () => {
  await expect(
    getQuote({
      ...params,
      deposit: { token, isNative: false },
    }),
  ).resolves.toEqual(testQuote)

  await expect(
    getQuote({
      ...params,
      deposit: { isNative: true },
    }),
  ).resolves.toEqual(testQuote)
})

test('parameters: mode', async () => {
  await expect(
    getQuote({
      ...params,
      mode: 'expense',
      deposit: { isNative: true, amount: 100n },
      // TODO expense amount shouldn't be allowed if mode === 'expense'
      expense: { isNative: true, amount: 100n },
    }),
  ).resolves.toEqual(testQuote)

  await expect(
    getQuote({
      ...params,
      mode: 'deposit',
      deposit: { isNative: true, amount: 100n },
      expense: { isNative: true, amount: 100n },
    }),
  ).resolves.toEqual(testQuote)
})

test.each([
  'test',
  {},
  { deposit: { token, amount: '100' } },
  { expense: { token, amount: '100' } },
  { deposit: { token }, expense: { token } },
  { deposit: { amount: '100' }, expense: { amount: '99' } },
])('behaviour: throws if response is not a quote: %s', async (mockReturn) => {
  vi.spyOn(api, 'safeFetchJSON').mockReturnValue(asyncResult(mockReturn))

  const expectRejection = expect(async () => {
    await getQuote(params)
  }).rejects
  await expectRejection.toBeInstanceOf(Error)
  await expectRejection.toHaveProperty(
    'message',
    `Unexpected quote response: ${JSON.stringify(mockReturn)}`,
  )
})
