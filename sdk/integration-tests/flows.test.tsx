import { act, render, waitFor } from '@testing-library/react'
import { createRef } from 'react'
import { expect, test } from 'vitest'
import { useConnect } from 'wagmi'

import {
  type Quote,
  useOrder,
  useQuote,
  useValidateOrder,
} from '../src/index.js'

import {
  ContextProvider,
  ETHER,
  MOCK_L1_ID,
  MOCK_L2_ID,
  ZERO_ADDRESS,
  accounts,
  createRenderHook,
  createWagmiConfig,
  testConnector,
} from './test-utils.js'

test.skip('successfully processes order from quote to filled', async () => {
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
