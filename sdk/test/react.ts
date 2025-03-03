import { QueryClient, QueryClientProvider } from '@tanstack/react-query'
import {
  type RenderHookOptions,
  type RenderHookResult,
  renderHook as _renderHook,
} from '@testing-library/react'
import { createElement } from 'react'
import { WagmiProvider, type createConfig } from 'wagmi'
import { OmniProvider } from '../src/context/omni.js'
import { web3Config as defaultWeb3Config } from './shared.js'

const queryClient = new QueryClient()

// biome-ignore lint/suspicious/noExplicitAny: test
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
  env: 'devnet' | 'testnet' = 'devnet',
  options?: RenderHookOptions<Props> | undefined,
): RenderHookResult<Result, Props> {
  queryClient.clear()
  return _renderHook(render, {
    wrapper: createWrapper(OmniProvider, {
      env,
    }),
    ...options,
  })
}
