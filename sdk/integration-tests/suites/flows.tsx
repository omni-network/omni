import { type Quote, getQuote, withExecAndTransfer } from '@omni-network/core'
import { useQuote, useValidateOrder } from '@omni-network/react'
import {
  createClient,
  invalidTokenAddress,
  middleman,
  mintOMNI,
  mockL1Chain,
  mockL1Id,
  mockL2Chain,
  mockL2Id,
  omniDevnetChain,
  omniDevnetId,
  testAccount,
  testAccountPrivateKeys,
  tokenAddress,
  vault,
} from '@omni-network/test-utils'
import { act, waitFor } from '@testing-library/react'
import { type PrivateKeyAccount, parseEther, zeroAddress } from 'viem'
import { privateKeyToAccount } from 'viem/accounts'
import { getBalance } from 'viem/actions'
import { beforeAll, describe, expect, test } from 'vitest'
import { useBalance } from 'wagmi'
import {
  type AnyOrder,
  createRenderHook,
  createTestConnector,
  devnetApiUrl,
  executeTestOrderUsingCore,
  executeTestOrderUsingReact,
  useOrderRef,
} from '../test-utils.js'

let nextAccountIndex = 1
function getNextAccount(): PrivateKeyAccount {
  const pk = testAccountPrivateKeys[nextAccountIndex++]
  if (pk == null) {
    throw new Error('No next private key available')
  }
  return privateKeyToAccount(pk)
}

describe.concurrent('ERC20 OMNI to native OMNI transfer orders', () => {
  describe.sequential('default: succeeds with valid expense', async () => {
    const account = getNextAccount()
    const srcClient = createClient({ account, chain: mockL1Chain })
    const destClient = createClient({ chain: omniDevnetChain })
    const amount = parseEther('10')
    const order: AnyOrder = {
      owner: account.address,
      srcChainId: mockL1Id,
      destChainId: omniDevnetId,
      expense: { token: zeroAddress, amount },
      calls: [{ target: account.address, value: amount }],
      deposit: { token: tokenAddress, amount },
    }

    beforeAll(async () => {
      await mintOMNI(srcClient)
    })

    test('using core APIs', async () => {
      await executeTestOrderUsingCore({ order, srcClient, destClient })
    })

    test('using React APIs', async () => {
      await executeTestOrderUsingReact({ account, order })
    })
  })

  describe.sequential('default: succeeds with native deposit', () => {
    const account = getNextAccount()
    const srcClient = createClient({ account, chain: mockL1Chain })
    const destClient = createClient({ chain: omniDevnetChain })
    const amount = parseEther('10')
    const order: AnyOrder = {
      owner: account.address,
      srcChainId: mockL1Id,
      destChainId: omniDevnetId,
      expense: { token: zeroAddress, amount },
      calls: [{ target: account.address, value: amount }],
      deposit: { token: zeroAddress, amount },
    }

    beforeAll(async () => {
      await mintOMNI(srcClient)
    })

    test('using core APIs', async () => {
      await executeTestOrderUsingCore({ order, srcClient, destClient })
    })

    test('using React APIs', async () => {
      await executeTestOrderUsingReact({ account, order })
    })
  })

  describe.concurrent('behaviour: fails with unsupported ERC20 deposit', () => {
    const amount = parseEther('10')
    const order: AnyOrder = {
      owner: testAccount.address,
      srcChainId: mockL1Id,
      destChainId: omniDevnetId,
      expense: { token: zeroAddress, amount },
      calls: [{ target: testAccount.address, value: amount }],
      deposit: { token: invalidTokenAddress, amount },
    }

    test('using core APIs', async () => {
      await executeTestOrderUsingCore({
        order,
        rejectReason: 'UnsupportedDeposit',
      })
    })

    test('using React APIs', async () => {
      await executeTestOrderUsingReact({
        order,
        rejectReason: 'UnsupportedDeposit',
      })
    })
  })
})

describe.concurrent('ETH transfer orders', () => {
  const destClient = createClient({ chain: mockL2Chain })

  describe.concurrent(
    'default: successfully processes order from quote to filled',
    () => {
      const quoteParams = {
        enabled: true,
        mode: 'expense',
        srcChainId: mockL1Id,
        destChainId: mockL2Id,
        deposit: {
          amount: parseEther('2'),
          isNative: true,
        },
        expense: {
          amount: parseEther('1'),
          isNative: true,
        },
      } as const

      test('using core APIs', async () => {
        const account = getNextAccount()
        const srcClient = createClient({ account, chain: mockL1Chain })
        const orderParams = {
          deposit: { token: zeroAddress, amount: parseEther('2') },
          expense: { token: zeroAddress, amount: parseEther('1') },
          calls: [{ target: account.address, value: parseEther('1') }],
          srcChainId: mockL1Id,
          destChainId: mockL2Id,
          validateEnabled: false,
        }

        const preDestBalance = await getBalance(destClient, {
          address: account.address,
        })

        const quote = await getQuote(devnetApiUrl, quoteParams)
        expect(quote).toEqual({
          deposit: { token: zeroAddress, amount: parseEther('2') },
          expense: { token: zeroAddress, amount: expect.any(BigInt) },
        })
        expect(quote.expense.amount).toBeLessThan(parseEther('2'))

        await executeTestOrderUsingCore({
          order: orderParams,
          srcClient,
          destClient,
        })

        const postDestBalance = await getBalance(destClient, {
          address: account.address,
        })
        expect(postDestBalance).toBe(
          preDestBalance + orderParams.calls[0].value,
        )
      })

      test('using React APIs', async () => {
        const account = getNextAccount()
        const orderParams = {
          deposit: { token: zeroAddress, amount: parseEther('2') },
          expense: { token: zeroAddress, amount: parseEther('1') },
          calls: [{ target: account.address, value: parseEther('1') }],
          srcChainId: mockL1Id,
          destChainId: mockL2Id,
          validateEnabled: false,
        }

        const renderHook = createRenderHook()
        const preDestBalance = renderHook(() => {
          return useBalance({
            address: account.address,
            chainId: mockL2Id,
          })
        })

        await waitFor(
          () => {
            expect(preDestBalance.result.current.data).toBeDefined()
          },
          { timeout: 5_000 },
        )

        const quoteHook = renderHook(() => {
          return useQuote(quoteParams)
        })

        await waitFor(() => {
          expect(quoteHook.result.current.query.error).toBeNull()
          expect(quoteHook.result.current.query.data).toBeDefined()
          expect(quoteHook.result.current.isSuccess).toBe(true)
          expect(quoteHook.result.current.isPending).toBe(false)
          expect(quoteHook.result.current.isError).toBe(false)
        })

        const quote = quoteHook.result.current.query.data as Quote

        expect(quote).toEqual({
          deposit: { token: zeroAddress, amount: parseEther('2') },
          expense: { token: zeroAddress, amount: expect.any(BigInt) },
        })
        expect(quote.expense.amount).toBeLessThan(parseEther('2'))

        const validateHook = renderHook(() => {
          return useValidateOrder({ enabled: true, order: orderParams })
        })

        await waitFor(() => {
          expect(validateHook.result.current.status === 'accepted').toBe(true)
        })

        const orderRef = useOrderRef(orderParams, account)
        await waitFor(() => expect(orderRef.current?.isReady).toBe(true))

        act(() => {
          orderRef.current?.open()
        })

        await waitFor(() => {
          expect(orderRef.current?.status).toBeOneOf(['open', 'opening'])
          expect(orderRef.current?.txHash).toBeDefined()
          expect(orderRef.current?.error).toBeUndefined()
          expect(orderRef.current?.isError).toBe(false)
          expect(orderRef.current?.isTxSubmitted).toBe(true)
          expect(orderRef.current?.isTxPending).toBe(false)
        })

        await waitFor(() => expect(orderRef.current?.status).toBe('filled'), {
          timeout: 10_000,
        })

        const postDestBalance = renderHook(() => {
          return useBalance({
            address: account.address,
            chainId: mockL2Id,
          })
        })

        await waitFor(
          () => {
            expect(postDestBalance.result.current.data).toBeDefined()
            expect(postDestBalance.result.current.data?.value).toBe(
              // biome-ignore lint/style/noNonNullAssertion: safe due to throwing condition above
              preDestBalance.result.current.data?.value! +
                orderParams.calls[0].value,
            )
          },
          { timeout: 5_000 },
        )
      })
    },
  )

  describe.concurrent('default: succeeds with valid expense', () => {
    test('using core APIs', async () => {
      const account = getNextAccount()
      const srcClient = createClient({ account, chain: mockL1Chain })
      const amount = parseEther('1')
      const order: AnyOrder = {
        owner: account.address,
        srcChainId: mockL1Id,
        destChainId: mockL2Id,
        expense: { token: zeroAddress, amount },
        calls: [{ target: account.address, value: amount }],
        deposit: { token: zeroAddress, amount: 2n * amount },
      }
      await executeTestOrderUsingCore({ order, srcClient, destClient })
    })

    test('using React APIs', async () => {
      const account = getNextAccount()
      const amount = parseEther('1')
      const order: AnyOrder = {
        owner: account.address,
        srcChainId: mockL1Id,
        destChainId: mockL2Id,
        expense: { token: zeroAddress, amount },
        calls: [{ target: account.address, value: amount }],
        deposit: { token: zeroAddress, amount: 2n * amount },
      }
      await executeTestOrderUsingReact({ account, order })
    })
  })

  describe.concurrent(
    'behaviour: successfully processes ETH deposit via middleman contract',
    () => {
      const amount = parseEther('1')
      const quoteParams = {
        mode: 'expense',
        srcChainId: mockL1Id,
        destChainId: mockL2Id,
        deposit: {
          amount,
          isNative: true,
        },
        expense: {
          isNative: true,
        },
      } as const

      test('using core APIs', async () => {
        const account = getNextAccount()
        const srcClient = createClient({ account, chain: mockL1Chain })
        const preDestBalance = await getBalance(destClient, {
          address: account.address,
        })

        const quote = await getQuote(devnetApiUrl, quoteParams)
        expect(quote).toEqual({
          deposit: { token: zeroAddress, amount: parseEther('1') },
          expense: { token: zeroAddress, amount: expect.any(BigInt) },
        })
        expect(quote.expense.amount).toBeLessThan(parseEther('1'))

        const call = withExecAndTransfer({
          call: {
            target: vault,
            value: quote.expense.amount,
            abi: [
              {
                type: 'function',
                name: 'depositWithoutBehalfOf',
                inputs: [],
                outputs: [],
                stateMutability: 'payable',
              },
            ],
            functionName: 'depositWithoutBehalfOf',
          },
          transfer: {
            token: zeroAddress,
            to: account.address,
          },
          middlemanAddress: middleman,
        })

        const order = {
          owner: account.address,
          srcChainId: mockL1Id,
          destChainId: mockL2Id,
          expense: { token: zeroAddress, amount: quote.expense.amount },
          calls: [call],
          deposit: { token: zeroAddress, amount: quote.deposit.amount },
          value: quote.deposit.amount,
        }
        await executeTestOrderUsingCore({
          order: order as AnyOrder,
          srcClient,
          destClient,
        })

        const postDestBalance = await getBalance(destClient, {
          address: account.address,
        })
        expect(postDestBalance).toBe(preDestBalance + quote.expense.amount)
      })

      test('using React APIs', async () => {
        const account = getNextAccount()
        const renderHook = createRenderHook()

        const preDestBalance = renderHook(() => {
          return useBalance({
            address: account.address,
            chainId: mockL2Id,
          })
        })

        const quoteHook = renderHook(() => {
          return useQuote({ enabled: true, ...quoteParams })
        })

        await waitFor(() => {
          expect(quoteHook.result.current.query.error).toBeNull()
          expect(quoteHook.result.current.query.data).toBeDefined()
          expect(quoteHook.result.current.isSuccess).toBe(true)
          expect(quoteHook.result.current.isPending).toBe(false)
          expect(quoteHook.result.current.isError).toBe(false)
        })

        const quote = quoteHook.result.current.query.data as Quote

        expect(quote).toEqual({
          deposit: { token: zeroAddress, amount: parseEther('1') },
          expense: { token: zeroAddress, amount: expect.any(BigInt) },
        })
        expect(quote.expense.amount).toBeLessThan(parseEther('1'))

        const call = withExecAndTransfer({
          call: {
            target: vault,
            value: quote.expense.amount,
            abi: [
              {
                type: 'function',
                name: 'depositWithoutBehalfOf',
                inputs: [],
                outputs: [],
                stateMutability: 'payable',
              },
            ],
            functionName: 'depositWithoutBehalfOf',
          },
          transfer: {
            token: zeroAddress,
            to: account.address,
          },
          middlemanAddress: middleman,
        })

        const order = {
          owner: account.address,
          srcChainId: mockL1Id,
          destChainId: mockL2Id,
          expense: { token: zeroAddress, amount: quote.expense.amount },
          calls: [call],
          deposit: { token: zeroAddress, amount: quote.deposit.amount },
          value: quote.deposit.amount,
        }

        const validateHook = renderHook(() => {
          return useValidateOrder({ enabled: true, order })
        })

        await waitFor(() => {
          expect(validateHook.result.current.status === 'accepted').toBe(true)
        })

        const orderRef = useOrderRef(order, account)

        await waitFor(() => expect(orderRef.current?.isReady).toBe(true))

        act(() => {
          orderRef.current?.open()
        })

        await waitFor(() => {
          expect(orderRef.current?.status).toBeOneOf(['open', 'opening'])
          expect(orderRef.current?.txHash).toBeDefined()
          expect(orderRef.current?.error).toBeUndefined()
          expect(orderRef.current?.isError).toBe(false)
          expect(orderRef.current?.isTxSubmitted).toBe(true)
          expect(orderRef.current?.isTxPending).toBe(false)
        })

        await waitFor(() => expect(orderRef.current?.status).toBe('filled'), {
          timeout: 10_000,
        })

        const postDestBalance = renderHook(() => {
          return useBalance({
            address: account.address,
            chainId: mockL2Id,
          })
        })

        await waitFor(
          () => {
            expect(postDestBalance.result.current.data).toBeDefined()
            expect(postDestBalance.result.current.data?.value).toBe(
              // biome-ignore lint/style/noNonNullAssertion: safe due to throwing condition above
              preDestBalance.result.current.data?.value! + quote.expense.amount,
            )
          },
          { timeout: 5_000 },
        )
      })
    },
  )

  describe.concurrent('behaviour: fails with expense over max amount', () => {
    const account = testAccount
    const amount = parseEther('1000')
    const order: AnyOrder = {
      owner: account.address,
      srcChainId: mockL1Id,
      destChainId: mockL2Id,
      expense: { token: zeroAddress, amount },
      calls: [{ target: account.address, value: amount }],
      deposit: { token: zeroAddress, amount: amount + parseEther('1') },
    }

    test('using core APIs', async () => {
      await executeTestOrderUsingCore({
        order,
        rejectReason: 'ExpenseOverMax',
      })
    })

    test('using React APIs', async () => {
      await executeTestOrderUsingReact({
        order,
        rejectReason: 'ExpenseOverMax',
      })
    })
  })

  describe.concurrent('behaviour: fails with expense under min amount', () => {
    const account = testAccount
    const amount = 1n
    const order: AnyOrder = {
      owner: account.address,
      srcChainId: mockL1Id,
      destChainId: mockL2Id,
      expense: { token: zeroAddress, amount },
      calls: [{ target: account.address, value: amount }],
      deposit: { token: zeroAddress, amount: 2n * amount },
    }

    test('using core APIs', async () => {
      await executeTestOrderUsingCore({
        order,
        rejectReason: 'ExpenseUnderMin',
      })
    })

    test('using React APIs', async () => {
      await executeTestOrderUsingReact({
        order,
        rejectReason: 'ExpenseUnderMin',
      })
    })
  })
})
