import type { Environment, OmniConfig } from '@omni-network/core'
import { useQueryClient } from '@tanstack/react-query'
import { createContext, useContext, useEffect, useMemo } from 'react'
import { getOmniContractsQueryOptions } from '../utils/getContracts.js'
import { throwingProxy } from '../utils/throwingProxy.js'

function apiUrl(env: Environment): string {
  switch (env) {
    case 'devnet':
      return 'http://localhost:26661/api/v1'
    case 'testnet':
      return 'https://solver.omega.omni.network/api/v1'
    case 'mainnet':
      return 'https://solver.mainnet.omni.network/api/v1'
    default:
      throw new Error(`Invalid environment supplied: ${env}`)
  }
}

type OmniProviderProps = {
  env: Environment
  __apiBaseUrl?: string
}

const OmniContext = createContext<OmniConfig>(throwingProxy<OmniConfig>())

export function OmniProvider({
  env,
  children,
  __apiBaseUrl: apiOverride,
}: React.PropsWithChildren<OmniProviderProps>) {
  const apiBaseUrl = apiOverride ?? apiUrl(env)
  const config = useMemo(() => {
    return { apiBaseUrl, env }
  }, [env, apiBaseUrl])

  const queryClient = useQueryClient()
  useEffect(() => {
    queryClient.prefetchQuery(getOmniContractsQueryOptions(config))
  }, [queryClient, config])

  return <OmniContext.Provider value={config}>{children}</OmniContext.Provider>
}

export function useOmniContext() {
  const context = useContext(OmniContext)

  if (!context) {
    throw new Error('useOmniContext must be used within OmniProvider')
  }

  return context
}
