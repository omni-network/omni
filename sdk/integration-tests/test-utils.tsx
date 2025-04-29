import {
  type DidFillParams,
  type GetOrderParameters,
  didFillOutbox,
  getOrder,
  openOrder,
  parseInboxStatus,
  parseOpenEvent,
  validateOrder,
} from '@omni-network/core'
import {
  OmniProvider,
  type Order,
  useOrder,
  type useParseOpenEvent,
} from '@omni-network/react'
import {
  createClient,
  inbox,
  mockChains,
  mockL1Chain,
  mockL1Id,
  mockL2Chain,
  mockL2Id,
  omniDevnetChain,
  outbox,
  testAccount,
} from '@omni-network/test-utils'
import { QueryClient, QueryClientProvider } from '@tanstack/react-query'
import {
  type RenderHookResult,
  act,
  render,
  renderHook,
  waitFor,
} from '@testing-library/react'
import { createRef } from 'react'
import {
  type Abi,
  type Account,
  type Client,
  type WalletClient,
  pad,
  parseEther,
  zeroAddress,
} from 'viem'
import { waitForTransactionReceipt, watchBlocks } from 'viem/actions'
import { expect } from 'vitest'
import {
  type Config,
  type CreateConnectorFn,
  WagmiProvider,
  createConfig,
  mock,
  useConnect,
} from 'wagmi'

export const devnetApiUrl = 'http://localhost:26661/api/v1'

const txHashRegexp = /^0x[0-9a-f]{64}$/

type WaitForInboxOrderFilledParams = GetOrderParameters & {
  pollingInterval?: number
  timeout?: number
}

function waitForInboxOrderFilled(
  params: WaitForInboxOrderFilledParams,
): Promise<void> {
  const { pollingInterval, timeout, ...getOrderParams } = params
  return new Promise<void>((resolve, reject) => {
    const timeoutId = setTimeout(() => {
      stopWatching()
      reject(new Error('Timeout waiting for order to be filled on inbox'))
    }, timeout ?? 60_000)
    const stopWatching = watchBlocks(params.client, {
      onBlock: async () => {
        const order = await getOrder(getOrderParams)
        const status = parseInboxStatus({ order })
        if (status === 'filled') {
          stopWatching
          clearTimeout(timeoutId)
          resolve()
        }
      },
      pollingInterval,
    })
  })
}

type WaitForOutboxOrderFilledParams = DidFillParams & {
  pollingInterval?: number
  timeout?: number
}

function waitForOutboxOrderFilled(
  params: WaitForOutboxOrderFilledParams,
): Promise<void> {
  const { pollingInterval, timeout, ...didFillParams } = params
  return new Promise<void>((resolve, reject) => {
    const timeoutId = setTimeout(() => {
      stopWatching()
      reject(new Error('Timeout waiting for order to be filled on outbox'))
    }, timeout ?? 60_000)
    const stopWatching = watchBlocks(params.client, {
      onBlock: async () => {
        const isFilled = await didFillOutbox(didFillParams)
        if (isFilled) {
          stopWatching()
          clearTimeout(timeoutId)
          resolve()
        }
      },
      pollingInterval,
    })
  })
}

type ExecuteTestOrderUsingCoreParams = {
  order: AnyOrder
} & (
  | { rejectReason: string }
  | { rejectReason?: never; srcClient: WalletClient; destClient: Client }
)

export async function executeTestOrderUsingCore(
  params: ExecuteTestOrderUsingCoreParams,
) {
  const { order, rejectReason } = params
  if (rejectReason != null) {
    await expect(validateOrder(devnetApiUrl, order)).resolves.toMatchObject({
      rejected: true,
      rejectReason,
    })
    return
  }

  await expect(validateOrder(devnetApiUrl, order)).resolves.toMatchObject({
    accepted: true,
  })

  const txHash = await openOrder({
    client: params.srcClient,
    inboxAddress: inbox,
    order,
  })
  expect(txHash).toMatch(txHashRegexp)

  const receipt = await waitForTransactionReceipt(params.srcClient, {
    hash: txHash,
    pollingInterval: 100,
  })
  const resolvedOrder = parseOpenEvent(receipt.logs)

  await Promise.all([
    waitForInboxOrderFilled({
      client: params.srcClient,
      inboxAddress: inbox,
      orderId: resolvedOrder.orderId,
      pollingInterval: 100,
      timeout: 20_000,
    }),
    waitForOutboxOrderFilled({
      client: params.destClient,
      outboxAddress: outbox,
      resolvedOrder,
      pollingInterval: 100,
      timeout: 20_000,
    }),
  ])
}

export function createTestConnector(account: Account) {
  // biome-ignore lint/suspicious/noExplicitAny: test file
  return function testConnector(config: any) {
    const connector = mock({ accounts: [account.address] })(config)
    connector.getClient = async ({ chainId } = {}) => {
      const chain = chainId ? mockChains[chainId] : mockL1Chain
      if (chain == null) {
        throw new Error(`Unsupported chain: ${chainId}`)
      }
      return createClient({ account, chain })
    }
    return connector
  }
}

export const testConnector = createTestConnector(testAccount)

export function createWagmiConfig() {
  return createConfig({
    chains: [mockL1Chain, mockL2Chain, omniDevnetChain],
    client: createClient,
  })
}

export function createQueryClient() {
  return new QueryClient({
    defaultOptions: {
      queries: {
        retry: false,
      },
    },
  })
}

export type TestConfig = {
  queryClient?: QueryClient
  wagmiConfig?: Config
}

export type ContextProviderProps = TestConfig & {
  children: React.ReactNode
}

export function ContextProvider(props: ContextProviderProps) {
  const client = props.queryClient ?? createQueryClient()
  const config = props.wagmiConfig ?? createWagmiConfig()

  return (
    <WagmiProvider config={config}>
      <QueryClientProvider client={client}>
        <OmniProvider env="devnet">{props.children}</OmniProvider>
      </QueryClientProvider>
    </WagmiProvider>
  )
}

export function createRenderHook(config: TestConfig = {}) {
  return function customRenderHook<Result>(
    render: () => Result,
  ): RenderHookResult<Result, ContextProviderProps> {
    return renderHook(render, {
      initialProps: config,
      wrapper: ContextProvider,
    })
  }
}

export type AnyOrder = Order<Array<unknown>>
export type ContractOrder = Order<readonly Abi[]>

function isContractOrder(
  order: AnyOrder | ContractOrder,
): order is ContractOrder {
  return order.calls.every((call) => call.abi != null)
}

export type UseOrderReturn = ReturnType<typeof useOrder>

export function useOrderRef(
  connector: CreateConnectorFn,
  order: AnyOrder | ContractOrder,
): React.RefObject<UseOrderReturn | null> {
  const connectRef = createRef<ReturnType<typeof useConnect>>()
  const orderRef = createRef<UseOrderReturn>()

  // useOrder() can only be used with a connected account, so we need to render it conditionally
  function TestOrder() {
    if (isContractOrder(order)) {
      orderRef.current = useOrder<readonly Abi[]>({
        validateEnabled: true,
        ...order,
      })
    } else {
      orderRef.current = useOrder<Array<unknown>>({
        validateEnabled: true,
        ...order,
      })
    }
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

type ExecuteTestOrderUsingReactParams = {
  connector?: CreateConnectorFn
  order: AnyOrder
  rejectReason?: string
}

export async function executeTestOrderUsingReact(
  params: ExecuteTestOrderUsingReactParams,
): Promise<void> {
  const { connector, order, rejectReason } = params
  const orderRef = useOrderRef(connector ?? testConnector, order)

  await waitFor(() => expect(orderRef.current?.isReady).toBe(true))

  await waitFor(() =>
    expect(orderRef.current?.validation?.status).toBeOneOf([
      'accepted',
      'rejected',
    ]),
  )

  if (rejectReason) {
    if (orderRef.current?.validation?.status !== 'rejected')
      throw new Error('Rejection expected')
    expect(orderRef.current?.validation?.status).toBe('rejected')
    expect(orderRef.current?.validation?.rejectReason).toBe(rejectReason)
    return
  }

  expect(orderRef.current?.validation?.status).toBe('accepted')

  act(() => {
    orderRef.current?.open()
  })

  const waitForOpts = {
    interval: 100,
    timeout: 20_000,
  }

  await waitFor(() => {
    expect(orderRef.current?.error).toBeUndefined()
    expect(orderRef.current?.txHash).toBeDefined()
  }, waitForOpts)

  await waitFor(() => {
    // allow filled, in case order was filled quickly
    expect(orderRef.current?.status).toBeOneOf(['open', 'filled'])
    expect(orderRef.current?.txHash).toBeDefined()
    expect(orderRef.current?.error).toBeUndefined()
    expect(orderRef.current?.orderId).toBeDefined()
  }, waitForOpts)

  await waitFor(() => {
    expect(orderRef.current?.txHash).toBeDefined()
    expect(orderRef.current?.error).toBeUndefined()
    expect(orderRef.current?.isError).toBe(false)
    expect(orderRef.current?.isTxSubmitted).toBe(true)
    expect(orderRef.current?.isTxPending).toBe(false)
    expect(orderRef.current?.status).toBe('filled')
  }, waitForOpts)
}

export function assertResolvedOrder(
  resolvedOrder: ReturnType<typeof useParseOpenEvent>['resolvedOrder'],
) {
  if (!resolvedOrder) throw new Error('Resolved order must be defined')

  expect(resolvedOrder.user).toEqual(testAccount.address)
  expect(resolvedOrder.originChainId).toEqual(BigInt(mockL2Id))
  expect(resolvedOrder.openDeadline).toEqual(0)
  expect(resolvedOrder.fillDeadline).toBeTypeOf('number')
  expect(resolvedOrder.orderId).toBeTypeOf('string')
  expect(resolvedOrder.orderId).toContain('0x')

  // maxSpent
  expect(resolvedOrder.maxSpent).toBeInstanceOf(Array)
  expect(resolvedOrder.maxSpent[0].token).toEqual(
    pad(zeroAddress, { size: 32, dir: 'left' }),
  )
  expect(resolvedOrder.maxSpent[0].amount).toEqual(parseEther('1'))
  expect(resolvedOrder.maxSpent[0].chainId).toEqual(BigInt(mockL1Id))
  expect(resolvedOrder.maxSpent[0].recipient).toEqual(
    pad(outbox, { size: 32, dir: 'left' }),
  )

  // minReceived
  expect(resolvedOrder.minReceived).toBeInstanceOf(Array)
  expect(resolvedOrder.minReceived[0].token).toEqual(
    pad(zeroAddress, { size: 32, dir: 'left' }),
  )
  expect(resolvedOrder.minReceived[0].amount).toEqual(parseEther('2'))
  expect(resolvedOrder.minReceived[0].chainId).toEqual(BigInt(mockL2Id))
  expect(resolvedOrder.minReceived[0].recipient).toEqual(
    pad(zeroAddress, { size: 32, dir: 'left' }),
  )

  // fillInstructions
  expect(resolvedOrder.fillInstructions).toBeInstanceOf(Array)
  expect(resolvedOrder.fillInstructions[0]).toBeTypeOf('object')
  expect(resolvedOrder.fillInstructions[0].destinationChainId).toEqual(
    BigInt(mockL1Id),
  )
  expect(resolvedOrder.fillInstructions[0].destinationSettler).toEqual(
    pad(outbox, { size: 32, dir: 'left' }),
  )
  expect(resolvedOrder.fillInstructions[0].originData).toBeTypeOf('string')
}
