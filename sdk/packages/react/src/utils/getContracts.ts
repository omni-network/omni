import type { FetchQueryOptions } from '@tanstack/react-query'
import { fetchJSON } from '../internal/api.js'
import type { OmniConfig } from '../types/config.js'
import type { OmniContracts } from '../types/contracts.js'

function isContracts(json: unknown): json is OmniContracts {
  const contracts = json as OmniContracts
  return (
    contracts != null &&
    typeof contracts.inbox === 'string' &&
    typeof contracts.outbox === 'string' &&
    typeof contracts.middleman === 'string'
  )
}

export async function getContracts(apiBaseUrl: string): Promise<OmniContracts> {
  const json = await fetchJSON(`${apiBaseUrl}/contracts`)

  if (!isContracts(json)) throw new Error('Unexpected /contracts response')

  return json
}

export function getOmniContractsQueryOptions(
  config: OmniConfig,
): FetchQueryOptions<OmniContracts> {
  return {
    queryKey: ['contracts', config.env],
    queryFn: () => getContracts(config.apiBaseUrl),
  }
}
