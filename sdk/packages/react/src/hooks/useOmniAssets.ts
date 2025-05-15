import { type Asset, getAssets } from '@omni-network/core'
import {
  type UseQueryOptions,
  type UseQueryResult,
  useQuery,
} from '@tanstack/react-query'
import { useOmniContext } from '../context/omni.js'

export type UseOmniAssetsParameters = {
  queryOpts?: Omit<UseQueryOptions<Asset[]>, 'queryKey' | 'queryFn' | 'enabled'>
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
