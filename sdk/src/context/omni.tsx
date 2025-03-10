import { createContext, useContext, useMemo } from 'react'
import { throwingProxy } from '../internal/util.js'

type Environment = 'devnet' | 'testnet'

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

type OmniProviderProps = {
  env: Environment
  __apiBaseUrl?: string
}

type OmniContextValue = {
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

  const value = useMemo(() => {
    return { apiBaseUrl, env }
  }, [env, apiBaseUrl])

  return <OmniContext.Provider value={value}>{children}</OmniContext.Provider>
}

export function useOmniContext() {
  const context = useContext(OmniContext)

  if (!context) {
    throw new Error('useOmniContext must be used within OmniProvider')
  }

  return context
}
