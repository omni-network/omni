import { fetchJSON } from '../internal/api.js'
import type { Environment } from '../types/config.js'
import type { OmniContracts } from '../types/contracts.js'
import { getApiUrl } from '../utils/getApiUrl.js'

function isContracts(json: unknown): json is OmniContracts {
  const contracts = json as OmniContracts
  return (
    contracts != null &&
    typeof contracts.inbox === 'string' &&
    typeof contracts.outbox === 'string' &&
    typeof contracts.middleman === 'string'
  )
}

export async function getContracts(
  envOrApiBaseUrl?: Environment | string,
): Promise<OmniContracts> {
  const apiUrl = getApiUrl(envOrApiBaseUrl)
  const json = await fetchJSON(`${apiUrl}/contracts`)

  if (!isContracts(json)) throw new Error('Unexpected /contracts response')

  return json
}
