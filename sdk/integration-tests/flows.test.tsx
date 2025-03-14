import { act, render, waitFor } from '@testing-library/react'
import { createRef } from 'react'
import type { PrivateKeyAccount } from 'viem/accounts'
import { describe, expect, should, test } from 'vitest'
import { useConnect } from 'wagmi'

import {
  type Order,
  type Quote,
  useOrder,
  useQuote,
  useValidateOrder,
} from '../src/index.js'

import {
  ACCOUNTS_RECORD,
  ContextProvider,
  ETHER,
  MOCK_L1_ID,
  MOCK_L2_ID,
  ZERO_ADDRESS,
  accounts,
  createClientFactory,
  createRenderHook,
  createTestConnector,
  createWagmiConfig,
  testConnector,
} from './test-utils.js'

type AnyOrder = Order<Array<unknown>>

describe('solver test orders', () => {
  type TestOrder = {
    label: string
    account: PrivateKeyAccount
    order: AnyOrder
    shouldReject: boolean
    rejectReason: string
  }

  const testAccounts = Object.values(ACCOUNTS_RECORD)

  const nativeOrders: Array<TestOrder> = testAccounts.map((account, i) => {
    const isOverMax = i < 3
    const isUnderMin = i > 6

    const [label, amount, shouldReject, rejectReason] = isOverMax
      ? [
          'Native order failing with expense over max amount',
          2n * ETHER,
          true,
          'ExpenseOverMax',
        ]
      : isUnderMin
        ? [
            'Native order failing with expense under min amount',
            1n,
            true,
            'ExpenseUnderMin',
          ]
        : ['Native order succeeding', ETHER, false, '']

    const order: AnyOrder = {
      owner: account.address,
      srcChainId: MOCK_L1_ID,
      destChainId: MOCK_L2_ID,
      expense: { token: ZERO_ADDRESS, amount },
      calls: [{ target: account.address, value: amount }],
      deposit: { token: ZERO_ADDRESS, amount: amount + ETHER },
    }

    return { label, account, order, shouldReject, rejectReason }
  })

  test.each(nativeOrders)(
    '$label with account $account.address',
    async ({ account, order, shouldReject, rejectReason }) => {
      const createClient = createClientFactory(account)
      const testConnector = createTestConnector(createClient)

      const connectRef = createRef()
      const orderRef = createRef()

      // useOrder() can only be used with a connected account, so we need to render it conditionally
      function TestOrder() {
        orderRef.current = useOrder({ ...order, validateEnabled: true })
        return null
      }

      // Wrap TestOrder to only render if connected
      function TestConnectAndOrder() {
        const connectReturn = useConnect()
        connectRef.current = connectReturn
        return connectReturn.data ? <TestOrder /> : null
      }

      render(<TestConnectAndOrder />, { wrapper: ContextProvider })
      act(() => {
        connectRef.current?.connect({ connector: testConnector })
      })

      await waitFor(() => expect(orderRef.current?.isReady).toBe(true))

      if (shouldReject) {
        expect(orderRef.current?.validation?.status).toBe('rejected')
        expect(orderRef.current?.validation?.rejectReason).toBe(rejectReason)
      } else {
        expect(orderRef.current?.validation?.status).toBe('accepted')
        act(() => {
          orderRef.current?.open()
        })
        await waitFor(() => expect(orderRef.current?.txHash).toBeDefined())
      }
    },
  )
})

test('successfully processes order from quote to filled', async () => {
  const wagmiConfig = createWagmiConfig()
  const renderHook = createRenderHook({ wagmiConfig })

  const quoteHook = renderHook(() => {
    return useQuote({
      enabled: true,
      mode: 'expense',
      srcChainId: MOCK_L1_ID,
      destChainId: MOCK_L2_ID,
      deposit: {
        amount: 2n * ETHER,
        isNative: true,
      },
      expense: {
        amount: 1n * ETHER,
        isNative: true,
      },
    })
  })
  await waitFor(() => expect(quoteHook.result.current.isSuccess).toBe(true))

  const quote = quoteHook.result.current.query.data as Quote
  expect(quote).toEqual({
    deposit: { token: ZERO_ADDRESS, amount: 2n * ETHER },
    expense: { token: ZERO_ADDRESS, amount: expect.any(BigInt) },
  })
  expect(quote.expense.amount).toBeLessThan(2n * ETHER)

  const orderParams = {
    deposit: { token: ZERO_ADDRESS, amount: 2n * ETHER },
    expense: { token: ZERO_ADDRESS, amount: 1n * ETHER },
    calls: [{ target: accounts[0], value: 1n * ETHER }],
    srcChainId: MOCK_L1_ID,
    destChainId: MOCK_L2_ID,
    validateEnabled: false,
  }

  const validateHook = renderHook(() => {
    return useValidateOrder({ enabled: true, order: orderParams })
  })
  await waitFor(() => {
    return expect(validateHook.result.current.status === 'accepted').toBe(true)
  })

  const connectRef = createRef()
  const orderRef = createRef()

  // useOrder() can only be used with a connected account, so we need to render it conditionally
  function TestOrder() {
    orderRef.current = useOrder(orderParams)
    return null
  }

  // Wrap TestOrder to only render if connected
  function TestConnectAndOrder() {
    const connectReturn = useConnect()
    connectRef.current = connectReturn
    return connectReturn.data ? <TestOrder /> : null
  }

  render(<TestConnectAndOrder />, { wrapper: ContextProvider })
  act(() => {
    connectRef.current?.connect({ connector: testConnector })
  })

  await waitFor(() => expect(orderRef.current?.isReady).toBe(true))
  act(() => {
    orderRef.current?.open()
  })
  await waitFor(() => expect(orderRef.current?.txHash).toBeDefined())
})
