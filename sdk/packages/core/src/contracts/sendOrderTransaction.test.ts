import { mockL1Client, testOrder } from '@omni-network/test-utils'
import { type Client, zeroAddress } from 'viem'
import { expect, test, vi } from 'vitest'
import { inboxABI } from '../constants/abis.js'
import { typeHash } from '../constants/typehash.js'
import { AccountRequiredError } from '../errors/base.js'
import * as encode from '../utils/encodeOrderData.js'
import { sendOrderTransaction } from './sendOrderTransaction.js'

const { writeContract } = vi.hoisted(() => ({
  writeContract: vi.fn(),
}))
vi.mock('viem/actions', () => ({ writeContract }))

test('default: opens order and returns the transaction hash', async () => {
  writeContract.mockResolvedValueOnce('0xtxHash')
  vi.spyOn(encode, 'encodeOrderData').mockReturnValueOnce('0xencodedOrder')

  await expect(
    sendOrderTransaction({
      client: mockL1Client,
      inboxAddress: '0xaddress',
      order: testOrder,
    }),
  ).resolves.toEqual('0xtxHash')
  expect(writeContract).toHaveBeenLastCalledWith(mockL1Client, {
    abi: inboxABI,
    address: '0xaddress',
    functionName: 'open',
    account: mockL1Client.account,
    chain: mockL1Client.chain,
    value: 0n,
    args: [
      {
        fillDeadline: expect.any(Number),
        orderDataType: typeHash,
        orderData: '0xencodedOrder',
      },
    ],
  })
})

test('behaviour: throws an AccountRequiredError if the client does not have an associated account', async () => {
  await expect(
    sendOrderTransaction({
      client: {} as Client,
      inboxAddress: '0xaddress',
      order: testOrder,
    }),
  ).rejects.toThrowError(AccountRequiredError)
})

test('behaviour: sets the order value when the deposit does not set a token address', async () => {
  writeContract.mockResolvedValueOnce('0xtxHash')
  vi.spyOn(encode, 'encodeOrderData').mockReturnValueOnce('0xencodedOrder')

  await expect(
    sendOrderTransaction({
      client: mockL1Client,
      inboxAddress: '0xaddress',
      order: { ...testOrder, deposit: { amount: 2n } },
    }),
  ).resolves.toEqual('0xtxHash')
  expect(writeContract).toHaveBeenLastCalledWith(mockL1Client, {
    abi: inboxABI,
    address: '0xaddress',
    functionName: 'open',
    account: mockL1Client.account,
    chain: mockL1Client.chain,
    value: 2n,
    args: [
      {
        fillDeadline: expect.any(Number),
        orderDataType: typeHash,
        orderData: '0xencodedOrder',
      },
    ],
  })
})

test('behaviour: sets the order value when the deposit token address is zero', async () => {
  writeContract.mockResolvedValueOnce('0xtxHash')
  vi.spyOn(encode, 'encodeOrderData').mockReturnValueOnce('0xencodedOrder')

  await expect(
    sendOrderTransaction({
      client: mockL1Client,
      inboxAddress: '0xaddress',
      order: { ...testOrder, deposit: { token: zeroAddress, amount: 3n } },
    }),
  ).resolves.toEqual('0xtxHash')
  expect(writeContract).toHaveBeenLastCalledWith(mockL1Client, {
    abi: inboxABI,
    address: '0xaddress',
    functionName: 'open',
    account: mockL1Client.account,
    chain: mockL1Client.chain,
    value: 3n,
    args: [
      {
        fillDeadline: expect.any(Number),
        orderDataType: typeHash,
        orderData: '0xencodedOrder',
      },
    ],
  })
})

test('behaviour: sets the order value to zero when the deposit token address is not zero', async () => {
  writeContract.mockResolvedValueOnce('0xtxHash')
  vi.spyOn(encode, 'encodeOrderData').mockReturnValueOnce('0xencodedOrder')

  await expect(
    sendOrderTransaction({
      client: mockL1Client,
      inboxAddress: '0xaddress',
      order: { ...testOrder, deposit: { token: '0x123', amount: 4n } },
    }),
  ).resolves.toEqual('0xtxHash')
  expect(writeContract).toHaveBeenLastCalledWith(mockL1Client, {
    abi: inboxABI,
    address: '0xaddress',
    functionName: 'open',
    account: mockL1Client.account,
    chain: mockL1Client.chain,
    value: 0n,
    args: [
      {
        fillDeadline: expect.any(Number),
        orderDataType: typeHash,
        orderData: '0xencodedOrder',
      },
    ],
  })
})
