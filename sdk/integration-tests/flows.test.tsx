import { act, render, waitFor } from '@testing-library/react'
import { createRef } from 'react'
import { describe, expect, test } from 'vitest'
import { type CreateConnectorFn, useConnect } from 'wagmi'

import {
  type Order,
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
  OMNI_DEVNET_ID,
  TOKEN_ADDRESS,
  ZERO_ADDRESS,
  createRenderHook,
  omniMintedPromise,
  testAccount,
  testConnector,
} from './test-utils.js'

type AnyOrder = Order<Array<unknown>>

type UseOrderReturn = ReturnType<typeof useOrder>

function useOrderRef(
  connector: CreateConnectorFn,
  order: AnyOrder,
): React.RefObject<UseOrderReturn | null> {
  const connectRef = createRef()
  const orderRef = createRef<UseOrderReturn>()

  // useOrder() can only be used with a connected account, so we need to render it conditionally
  function TestOrder() {
    orderRef.current = useOrder({ validateEnabled: true, ...order })
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
    connectRef.current?.connect({ connector })
  })

  return orderRef
}

async function executeTestOrder(
  order: AnyOrder,
  rejectReason?: string,
): Promise<void> {
  const orderRef = useOrderRef(testConnector, order)
  await waitFor(() => expect(orderRef.current?.isReady).toBe(true))

  if (rejectReason) {
    expect(orderRef.current?.validation?.status).toBe('rejected')
    expect(orderRef.current?.validation?.rejectReason).toBe(rejectReason)
  } else {
    expect(orderRef.current?.validation?.status).toBe('accepted')
    act(() => {
      orderRef.current?.open()
    })
    await waitFor(() => expect(orderRef.current?.txHash).toBeDefined())
  }
}

describe('ERC20 OMNI to native OMNI transfer orders', () => {
  test('default: succeeds with valid expense', async () => {
    await omniMintedPromise
    const amount = 10n * ETHER
    const order: AnyOrder = {
      owner: testAccount.address,
      srcChainId: MOCK_L1_ID,
      destChainId: OMNI_DEVNET_ID,
      expense: { token: ZERO_ADDRESS, amount },
      calls: [{ target: testAccount.address, value: amount }],
      deposit: { token: TOKEN_ADDRESS, amount },
    }
    await executeTestOrder(order)
  }, 15_000)

  test('behaviour: fails with native deposit', async () => {
    const amount = 10n * ETHER
    const order: AnyOrder = {
      owner: testAccount.address,
      srcChainId: MOCK_L1_ID,
      destChainId: OMNI_DEVNET_ID,
      expense: { token: ZERO_ADDRESS, amount },
      calls: [{ target: testAccount.address, value: amount }],
      deposit: { token: ZERO_ADDRESS, amount },
    }
    await executeTestOrder(order, 'InvalidDeposit')
  })

  test('behaviour: fails with unsupported ERC20 deposit', async () => {
    const amount = 10n * ETHER
    const order: AnyOrder = {
      owner: testAccount.address,
      srcChainId: MOCK_L1_ID,
      destChainId: OMNI_DEVNET_ID,
      expense: { token: ZERO_ADDRESS, amount },
      calls: [{ target: testAccount.address, value: amount }],
      deposit: { token: '0x1234000000000000000000000000000000000000', amount },
    }
    await executeTestOrder(order, 'UnsupportedDeposit')
  })
})

describe('ETH transfer orders', () => {
  test('default: succeeds with valid expense', async () => {
    const account = testAccount
    const amount = ETHER
    const order: AnyOrder = {
      owner: account.address,
      srcChainId: MOCK_L1_ID,
      destChainId: MOCK_L2_ID,
      expense: { token: ZERO_ADDRESS, amount },
      calls: [{ target: account.address, value: amount }],
      deposit: { token: ZERO_ADDRESS, amount: amount + ETHER },
    }
    await executeTestOrder(order)
  })

  test('behaviour: fails with expense over max amount', async () => {
    const account = testAccount
    const amount = 2n * ETHER
    const order: AnyOrder = {
      owner: account.address,
      srcChainId: MOCK_L1_ID,
      destChainId: MOCK_L2_ID,
      expense: { token: ZERO_ADDRESS, amount },
      calls: [{ target: account.address, value: amount }],
      deposit: { token: ZERO_ADDRESS, amount: amount + ETHER },
    }
    await executeTestOrder(order, 'ExpenseOverMax')
  })

  test('behaviour: fails with expense under min amount', async () => {
    const account = testAccount
    const amount = 1n
    const order: AnyOrder = {
      owner: account.address,
      srcChainId: MOCK_L1_ID,
      destChainId: MOCK_L2_ID,
      expense: { token: ZERO_ADDRESS, amount },
      calls: [{ target: account.address, value: amount }],
      deposit: { token: ZERO_ADDRESS, amount: amount + ETHER },
    }
    await executeTestOrder(order, 'ExpenseUnderMin')
  })
})

test('default: successfully processes order from quote to filled', async () => {
  const renderHook = createRenderHook()

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
    calls: [{ target: testAccount.address, value: 1n * ETHER }],
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

  const orderRef = useOrderRef(testConnector, orderParams)
  await waitFor(() => expect(orderRef.current?.isReady).toBe(true))

  act(() => {
    orderRef.current?.open()
  })
  await waitFor(() => expect(orderRef.current?.txHash).toBeDefined())
})
