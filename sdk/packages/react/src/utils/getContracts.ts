import {
  type OmniConfig,
  type OmniContracts,
  getContracts,
} from '@omni-network/core'
import type { FetchQueryOptions } from '@tanstack/react-query'

export function getOmniContractsQueryOptions(
  config: OmniConfig,
): FetchQueryOptions<OmniContracts> {
  return {
    queryKey: ['contracts', config.env],
    queryFn: () => getContracts({ environment: config.apiBaseUrl }),
  }
}
