import { expect, test } from 'vitest'
import type { GetOrderReturn } from './getOrder.js'
import { parseInboxStatus, status } from './parseInboxStatus.js'

test.each(Object.entries(status))('code %i returns "%s"', (status, text) => {
  expect(
    parseInboxStatus({
      order: [{}, { status }, {}] as unknown as GetOrderReturn,
    }),
  ).toBe(text)
})

test('unknown codes return "unknown"', () => {
  expect(
    parseInboxStatus({
      order: [{}, { status: 100 }, {}] as unknown as GetOrderReturn,
    }),
  ).toBe('unknown')
  expect(
    parseInboxStatus({
      order: [{}, { status: 111 }, {}] as unknown as GetOrderReturn,
    }),
  ).toBe('unknown')
})
