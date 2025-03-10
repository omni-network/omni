import { type UseQueryResult, useQuery } from '@tanstack/react-query'

import { useOmniContext } from '../context/omni.js'
import type { OmniContracts } from '../types/contracts.js'
import { getOmniContractsQueryOptions } from '../utils/getContracts.js'

export type UseOmniContractsResult = UseQueryResult<OmniContracts>

export function useOmniContracts(): UseOmniContractsResult {
  const config = useOmniContext()
  return useQuery(getOmniContractsQueryOptions(config))
}
