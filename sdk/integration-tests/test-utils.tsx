import { OmniProvider, type Order, useOrder } from '@omni-network/react'
import {
  MOCK_CHAINS,
  MOCK_L1_CHAIN,
  MOCK_L1_ID,
  MOCK_L2_CHAIN,
  OMNI_DEVNET_CHAIN,
  createClient,
  mockL1Client,
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
import { expect } from 'vitest'
import {
  type Config,
  type CreateConnectorFn,
  WagmiProvider,
  createConfig,
  mock,
  useConnect,
} from 'wagmi'

const mockConnector = mock({ accounts: [testAccount.address] as const })

export function testConnector(config) {
  const connector = mockConnector(config)
  connector.getClient = async ({ chainId } = {}) => {
    if (chainId === MOCK_L1_ID) {
      return mockL1Client
    }
    const chain = chainId ? MOCK_CHAINS[chainId] : MOCK_L1_CHAIN
    if (chain == null) {
      throw new Error(`Unsupported chain: ${chainId}`)
    }
    return createClient({ chain })
  }
  return connector
}

export function createWagmiConfig() {
  return createConfig({
    chains: [MOCK_L1_CHAIN, MOCK_L2_CHAIN, OMNI_DEVNET_CHAIN],
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

export type UseOrderReturn = ReturnType<typeof useOrder>

export function useOrderRef(
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

export async function executeTestOrder(
  order: AnyOrder,
  rejectReason?: string,
): Promise<void> {
  const orderRef = useOrderRef(testConnector, order)

  await waitFor(() => expect(orderRef.current?.isReady).toBeTruthy())

  // Wait for order to be validated
  await waitFor(() =>
    expect(orderRef.current?.validation?.status).toBeOneOf([
      'accepted',
      'rejected',
    ]),
  )

  if (rejectReason) {
    expect(orderRef.current?.validation?.status).toBe('rejected')
    expect(orderRef.current?.validation?.rejectReason).toBe(rejectReason)
    return
  }

  expect(orderRef.current?.validation?.status).toBe('accepted')

  // Open the order
  act(() => {
    orderRef.current?.open()
  })

  const waitForOpts = {
    interval: 100,
    timeout: 5000,
  }

  // Assert tx submitted
  await waitFor(() => {
    expect(orderRef.current?.error).toBeUndefined()
    expect(orderRef.current?.txHash).toBeDefined()
  }, waitForOpts)

  // Assert the order was opened properly
  await waitFor(() => {
    expect(orderRef.current?.error).toBeUndefined()
    expect(orderRef.current?.orderId).toBeDefined()
    expect(orderRef.current?.status).toBeOneOf(['open', 'filled']) // allow filled, in case order was filled quickly
  }, waitForOpts)

  // Assert the order was filled
  await waitFor(() => {
    expect(orderRef.current?.error).toBeUndefined()
    expect(orderRef.current?.status).toBe('filled')
  }, waitForOpts)
}
