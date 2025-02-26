import { useQuery } from '@tanstack/react-query'
import { createContext, useContext, useMemo } from 'react'
import type { Address } from 'viem'
import { fetchJSON } from '../internal/api.js'
import { throwingProxy } from '../internal/util.js'

type Environment = 'devnet' | 'testnet'

type OmniProviderProps = {
  env: Environment
  __apiBaseUrl?: string
}

type OmniContracts = {
  inbox: Address
  outbox: Address
  middleman: Address
}

type OmniContextValue = OmniContracts & {
  apiBaseUrl: string
  env: Environment
}

const OmniContext = createContext<OmniContextValue>(
  throwingProxy<OmniContextValue>(),
)

export function OmniProvider({
  env,
  children,
  __apiBaseUrl: apiOverride,
}: React.PropsWithChildren<OmniProviderProps>) {
  const apiBaseUrl = apiOverride ?? apiUrl(env)

  const contracts = useQuery({
    queryKey: ['contracts', env],
    queryFn: () => getContracts(apiBaseUrl),
  })

  const value = useMemo(() => {
    // if api url overrided, provide no defaults, must fetch from api
    const defaults = apiOverride
      ? ({
          inbox: '0x',
          outbox: '0x',
          middleman: '0x',
        } as const)
      : defaultContracts(env)

    // Currently we override defaults with API response when available
    //
    // Reasoning: defaults will almost always match api response. so for most
    // cases, explicit api loading state would be unnecessarily cumbersome.
    // Additionally, api response will generally be available before
    // an address is used.
    //
    // This does mean that if the api is unreachable, we will use defaults.
    // If we are changing addresses in prod, this is a problem.
    //
    // TODO: track ready / error states in context, and log warnings
    // when opening orders, or block order operations entirely.
    return {
      ...defaults,
      ...contracts.data,
      apiBaseUrl,
      env,
    }
  }, [contracts.data, env, apiBaseUrl, apiOverride])

  return <OmniContext.Provider value={value}>{children}</OmniContext.Provider>
}

export function useOmniContext() {
  const context = useContext(OmniContext)

  if (!context) {
    throw new Error('useOmniContext must be used within OmniProvider')
  }

  return context
}

async function getContracts(apiBaseUrl: string) {
  const json = await fetchJSON(`${apiBaseUrl}/contracts`)

  if (!isContracts(json)) throw new Error('Unexpected /contracts response')

  return json
}

function isContracts(json: unknown): json is OmniContracts {
  const contracts = json as OmniContracts
  return (
    contracts != null &&
    typeof contracts.inbox === 'string' &&
    typeof contracts.outbox === 'string' &&
    typeof contracts.middleman === 'string'
  )
}

function apiUrl(env: Environment): string {
  switch (env) {
    case 'devnet':
      return 'http://localhost:26661/api/v1'
    case 'testnet':
      return 'https://solver.omega.omni.network/api/v1'
    default:
      throw new Error(`Invalid environment supplied: ${env}`)
  }
}

function defaultContracts(env: Environment): OmniContracts {
  switch (env) {
    case 'devnet':
      return {
        inbox: '0x7c7759b801078ecb2c41c9caecc2db13c3079c76',
        outbox: '0x29d9e8faa760841aacbe79a8632c1f42e0a858e6',
        middleman: '0x1b99e432d5f9e8110102b8d3dce2d0b462a37942',
      }
    case 'testnet':
      return {
        inbox: '0x80b6ed465241a17080dc4a68be42e80fea1214dd',
        outbox: '0x020b76746606c6ddb4708b6996cad9adb821604e',
        middleman: '0x191e0a0aab2b21e946a0ff0ecbd36218b90801c9',
      }
  }
}
