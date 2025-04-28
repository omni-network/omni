import { testAccount, testOrder } from '@omni-network/test-utils'
import { expect, test } from 'vitest'
import { encodeOrderData } from './encodeOrderData.js'

test('default: encodes the order to an hex string', () => {
  expect(
    encodeOrderData({
      ...testOrder,
      deposit: {
        token: '0x23e98253f372ee29910e22986fe75bb287b011fc',
        amount: 1n,
      },
      expense: {
        token: '0x23e98253f372ee29910e22986fe75bb287b011fc',
        amount: 1n,
      },
    }),
  ).toMatch(/^0x[0-9a-f]+$/)
})

test('behaviour: throws if the order contains an invalid address', () => {
  expect(() => {
    encodeOrderData(testOrder)
  }).toThrow('Address "0x123" is invalid.')
})

test('behaviour: throws if the order contains an invalid call', () => {
  expect(() => {
    encodeOrderData({
      ...testOrder,
      calls: [
        {
          abi: [],
          // @ts-expect-error: ABI does not contain function
          functionName: 'transfer',
          target: '0x23e98253f372ee29910e22986fe75bb287b011fc',
          value: 0n,
          args: [testAccount.address, 0n],
        },
      ],
    })
  }).toThrow('Function "transfer" not found on ABI.')
})
