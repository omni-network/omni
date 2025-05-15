import type { OmniContracts } from '@omni-network/core'
import {
  type UseQueryOptions,
  type UseQueryResult,
  useQuery,
} from '@tanstack/react-query'
import { useOmniContext } from '../context/omni.js'
import { getOmniContractsQueryOptions } from '../utils/getContracts.js'

export type UseOmniContractsParameters = {
  queryOpts?: Omit<
    UseQueryOptions<OmniContracts>,
    'queryKey' | 'queryFn' | 'enabled'
  >
}

export type UseOmniContractsResult = UseQueryResult<OmniContracts>

export function useOmniContracts({
  queryOpts,
}: UseOmniContractsParameters = {}): UseOmniContractsResult {
  const config = useOmniContext()
  return useQuery({
    ...getOmniContractsQueryOptions(config),
    ...queryOpts,
  })
}
