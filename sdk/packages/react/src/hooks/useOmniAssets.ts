import { type Asset, getAssets } from '@omni-network/core'
import { type UseQueryResult, useQuery } from '@tanstack/react-query'
import { useOmniContext } from '../context/omni.js'
import type { QueryOpts } from './types.js'

export type UseOmniAssetsParameters = {
  queryOpts?: QueryOpts<Asset[]>
}

export type UseOmniAssetsResult = UseQueryResult<Asset[]>

export function useOmniAssets({
  queryOpts,
}: UseOmniAssetsParameters = {}): UseOmniAssetsResult {
  const config = useOmniContext()
  return useQuery({
    queryKey: ['assets', config.env],
    queryFn: () => getAssets({ environment: config.env }),
    ...queryOpts,
  })
}
