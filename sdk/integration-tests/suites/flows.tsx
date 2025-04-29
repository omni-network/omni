import { type Quote, getQuote, withExecAndTransfer } from '@omni-network/core'
import { useQuote, useValidateOrder } from '@omni-network/react'
import {
  createClient,
  invalidTokenAddress,
  middleman,
  mintOMNI,
  mockL1Id,
  mockL2Chain,
  mockL2Id,
  omniDevnetChain,
  omniDevnetId,
  testAccount,
  tokenAddress,
  vault,
} from '@omni-network/test-utils'
import { act, waitFor } from '@testing-library/react'
import { parseEther, zeroAddress } from 'viem'
import { getBalance } from 'viem/actions'
import { beforeAll, describe, expect, test } from 'vitest'
import { useBalance } from 'wagmi'
import {
  type AnyOrder,
  createRenderHook,
  devnetApiUrl,
  executeTestOrderUsingCore,
  executeTestOrderUsingReact,
  testConnector,
  useOrderRef,
} from '../test-utils.js'

describe.concurrent('ERC20 OMNI to native OMNI transfer orders', () => {
  const outboxClient = createClient({ chain: omniDevnetChain })

  beforeAll(async () => {
    await mintOMNI()
  })

  describe.sequential('default: succeeds with valid expense', async () => {
    const amount = parseEther('10')
    const order: AnyOrder = {
      owner: testAccount.address,
      srcChainId: mockL1Id,
      destChainId: omniDevnetId,
      expense: { token: zeroAddress, amount },
      calls: [{ target: testAccount.address, value: amount }],
      deposit: { token: tokenAddress, amount },
    }

    test('using core APIs', async () => {
      await executeTestOrderUsingCore({ order, outboxClient })
    })

    test('using React APIs', async () => {
      await executeTestOrderUsingReact(order)
    })
  })

  describe.sequential('default: succeeds with native deposit', () => {
    const amount = parseEther('10')
    const order: AnyOrder = {
      owner: testAccount.address,
      srcChainId: mockL1Id,
      destChainId: omniDevnetId,
      expense: { token: zeroAddress, amount },
      calls: [{ target: testAccount.address, value: amount }],
      deposit: { token: zeroAddress, amount },
    }

    test('using core APIs', async () => {
      await executeTestOrderUsingCore({ order, outboxClient })
    })

    test('using React APIs', async () => {
      await executeTestOrderUsingReact(order)
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
      await executeTestOrderUsingReact(order, 'UnsupportedDeposit')
    })
  })
})

describe.sequential('ETH transfer orders', () => {
  const outboxClient = createClient({ chain: mockL2Chain })

  describe.sequential(
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

      const orderParams = {
        deposit: { token: zeroAddress, amount: parseEther('2') },
        expense: { token: zeroAddress, amount: parseEther('1') },
        calls: [{ target: testAccount.address, value: parseEther('1') }],
        srcChainId: mockL1Id,
        destChainId: mockL2Id,
        validateEnabled: false,
      }

      test('using core APIs', async () => {
        const preDestBalance = await getBalance(outboxClient, {
          address: testAccount.address,
        })

        const quote = await getQuote(devnetApiUrl, quoteParams)
        expect(quote).toEqual({
          deposit: { token: zeroAddress, amount: parseEther('2') },
          expense: { token: zeroAddress, amount: expect.any(BigInt) },
        })
        expect(quote.expense.amount).toBeLessThan(parseEther('2'))

        await executeTestOrderUsingCore({ order: orderParams, outboxClient })

        const postDestBalance = await getBalance(outboxClient, {
          address: testAccount.address,
        })
        expect(postDestBalance).toBe(
          preDestBalance + orderParams.calls[0].value,
        )
      })

      test('using React APIs', async () => {
        const renderHook = createRenderHook()
        const preDestBalance = renderHook(() => {
          return useBalance({
            address: testAccount.address,
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

        const orderRef = useOrderRef(testConnector, orderParams)
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
            address: testAccount.address,
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

  describe.sequential('default: succeeds with valid expense', () => {
    const account = testAccount
    const amount = parseEther('1')
    const order: AnyOrder = {
      owner: account.address,
      srcChainId: mockL1Id,
      destChainId: mockL2Id,
      expense: { token: zeroAddress, amount },
      calls: [{ target: account.address, value: amount }],
      deposit: { token: zeroAddress, amount: 2n * amount },
    }

    test('using core APIs', async () => {
      await executeTestOrderUsingCore({ order, outboxClient })
    })

    test('using React APIs', async () => {
      await executeTestOrderUsingReact(order)
    })
  })

  describe.sequential(
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
        const preDestBalance = await getBalance(outboxClient, {
          address: testAccount.address,
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
            to: testAccount.address,
          },
          middlemanAddress: middleman,
        })

        const order = {
          owner: testAccount.address,
          srcChainId: mockL1Id,
          destChainId: mockL2Id,
          expense: { token: zeroAddress, amount: quote.expense.amount },
          calls: [call],
          deposit: { token: zeroAddress, amount: quote.deposit.amount },
          value: quote.deposit.amount,
        }
        await executeTestOrderUsingCore({
          order: order as AnyOrder,
          outboxClient,
        })

        const postDestBalance = await getBalance(outboxClient, {
          address: testAccount.address,
        })
        expect(postDestBalance).toBe(preDestBalance + quote.expense.amount)
      })

      test('using React APIs', async () => {
        const renderHook = createRenderHook()

        const preDestBalance = renderHook(() => {
          return useBalance({
            address: testAccount.address,
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
            to: testAccount.address,
          },
          middlemanAddress: middleman,
        })

        const order = {
          owner: testAccount.address,
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

        const orderRef = useOrderRef(testConnector, order)

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
            address: testAccount.address,
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
      await executeTestOrderUsingReact(order, 'ExpenseOverMax')
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
      await executeTestOrderUsingReact(order, 'ExpenseUnderMin')
    })
  })
})
