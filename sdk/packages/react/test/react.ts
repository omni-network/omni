import { QueryClient, QueryClientProvider } from '@tanstack/react-query'
import {
  type RenderHookOptions,
  type RenderHookResult,
  renderHook as _renderHook,
} from '@testing-library/react'
import { createElement } from 'react'
import { WagmiProvider, type createConfig } from 'wagmi'
import { OmniProvider } from '../src/context/omni.js'
import { mockContractsQuery } from './mocks.js'
import { web3Config as defaultWeb3Config } from './shared.js'

const queryClient = new QueryClient({
  defaultOptions: {
    queries: {
      // turns retries off to trigger failures immediately
      retry: false,
    },
  },
})

// biome-ignore lint/suspicious/noExplicitAny: test module
export function createWrapper<TComponent extends React.FunctionComponent<any>>(
  Wrapper: TComponent,
  props: Parameters<TComponent>[0],
  web3Config: ReturnType<typeof createConfig> = defaultWeb3Config,
) {
  return function CreatedWrapper({
    children,
  }: { children?: React.ReactNode | undefined }) {
    return createElement(
      QueryClientProvider,
      { client: queryClient },
      createElement(
        WagmiProvider,
        {
          config: web3Config,
          reconnectOnMount: false,
        },
        createElement(Wrapper, props, children),
      ),
    )
  }
}

export function renderHook<Result, Props>(
  render: (props: Props) => Result,
  options?: RenderHookOptions<Props> & {
    mockContractsCall?: boolean
    mockContractsCallFailure?: boolean
    env?: 'devnet' | 'testnet'
  },
): RenderHookResult<Result, Props> {
  queryClient.clear()

  if (options?.mockContractsCall) {
    mockContractsQuery()

    const { mockContractsCall, ...restOptions } = options
    // biome-ignore lint/style/noParameterAssign: test module
    options = restOptions
  }

  if (options?.mockContractsCallFailure) {
    mockContractsQuery(true)

    const { mockContractsCall, ...restOptions } = options
    // biome-ignore lint/style/noParameterAssign: test module
    options = restOptions
  }

  return _renderHook(render, {
    wrapper: createWrapper(OmniProvider, {
      env: options?.env ?? 'devnet',
    }),
    ...options,
  })
}
