import { testAccount, testOrder } from '@omni-network/test-utils'
import { expect, test, vi } from 'vitest'
import * as api from '../internal/api.js'
import { validateOrder } from './validateOrder.js'

// TODO calls as empty array should not be allowed // throw error

test('default: native transfer order', async () => {
  vi.spyOn(api, 'fetchJSON').mockResolvedValue({
    accepted: true,
  })

  const resultPromise = validateOrder(
    {
      owner: testAccount.address,
      srcChainId: 1,
      destChainId: 2,
      calls: [
        {
          target: testAccount.address,
          value: 0n,
        },
      ],
      deposit: {
        amount: 0n,
      },
      expense: {
        amount: 0n,
      },
    },
    'http://localhost',
  )
  await expect(resultPromise).resolves.toEqual({ accepted: true })
  expect(api.fetchJSON).toHaveBeenCalledWith(
    'http://localhost/check',
    expect.any(Object),
  )
})

test('default: order', async () => {
  vi.spyOn(api, 'fetchJSON').mockResolvedValue({
    accepted: true,
  })
  await expect(validateOrder(testOrder, 'http://localhost')).resolves.toEqual({
    accepted: true,
  })
  expect(api.fetchJSON).toHaveBeenCalledWith(
    'http://localhost/check',
    expect.any(Object),
  )
})

test('behaviour: resolves if response is supported error object', async () => {
  const response = {
    error: {
      code: 1,
      message: 'an error',
    },
  }
  vi.spyOn(api, 'fetchJSON').mockResolvedValue(response)
  await expect(validateOrder(testOrder, 'http://localhost')).resolves.toBe(
    response,
  )
})

test('behaviour: resolves if response is supported rejection object', async () => {
  const response = {
    rejected: true,
    rejectReason: 'a reason',
    rejectDescription: 'a description',
  }
  vi.spyOn(api, 'fetchJSON').mockResolvedValue(response)
  await expect(validateOrder(testOrder)).resolves.toBe(response)
})

test.each([
  'test',
  {},
  { rejected: true },
  { rejected: true, rejectReason: 'a reason' },
  { rejecetd: true, rejectDescription: 'a description' },
])('behaviour: throws if response is not valid: %s', async (mockReturn) => {
  vi.spyOn(api, 'fetchJSON').mockResolvedValue(mockReturn)

  const expectRejection = expect(async () => {
    await validateOrder(testOrder)
  }).rejects
  await expectRejection.toBeInstanceOf(Error)
  await expectRejection.toHaveProperty(
    'message',
    `Unexpected validation response: ${JSON.stringify(mockReturn)}`,
  )
})

test('behaviour: throws when the order encoding is invalid', async () => {
  const invalidOrder = {
    ...testOrder,
    calls: [{ ...testOrder.calls[0], args: ['0xinvalid', 0n] }],
  }
  const expectRejection = expect(async () => {
    // @ts-expect-error: invalid order
    await validateOrder(invalidOrder)
  }).rejects
  await expectRejection.toBeInstanceOf(Error)
  await expectRejection.toMatchObject({
    message: expect.stringMatching(`Address "0xinvalid" is invalid`),
  })
})
