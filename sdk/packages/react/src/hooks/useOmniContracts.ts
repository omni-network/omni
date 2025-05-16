import type { OmniContracts } from '@omni-network/core'
import { type UseQueryResult, useQuery } from '@tanstack/react-query'
import { useOmniContext } from '../context/omni.js'
import { getOmniContractsQueryOptions } from '../utils/getContracts.js'
import type { QueryOpts } from './types.js'

export type UseOmniContractsParameters = {
  queryOpts?: QueryOpts<OmniContracts>
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
