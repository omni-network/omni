import {
  type Environment,
  type OmniConfig,
  getApiUrl,
} from '@omni-network/core'
import { useQueryClient } from '@tanstack/react-query'
import { createContext, useContext, useEffect, useMemo } from 'react'
import { getOmniContractsQueryOptions } from '../utils/getContracts.js'
import { throwingProxy } from '../utils/throwingProxy.js'

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
  const apiBaseUrl = apiOverride ?? getApiUrl(env)
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
