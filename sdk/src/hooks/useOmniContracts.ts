import { type UseQueryResult, useQuery } from '@tanstack/react-query'
import type { Address } from 'viem'
import { useOmniContext } from '../context/omni.js'
import { fetchJSON } from '../internal/api.js'

export type OmniContracts = {
  inbox: Address
  outbox: Address
  middleman: Address
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

export type UseOmniContractsResult = UseQueryResult<OmniContracts>

export function useOmniContracts(): UseOmniContractsResult {
  const { apiBaseUrl, env } = useOmniContext()
  return useQuery({
    queryKey: ['contracts', env],
    queryFn: () => getContracts(apiBaseUrl),
  })
}
