import { type Asset, getAssets } from '@omni-network/core'
import { type UseQueryResult, useQuery } from '@tanstack/react-query'
import { useOmniContext } from '../context/omni.js'

export type UseOmniAssetsResult = UseQueryResult<Asset[]>

export function useOmniAssets(): UseOmniAssetsResult {
  const config = useOmniContext()
  return useQuery({
    queryKey: ['assets', config.env],
    queryFn: () => getAssets({ environment: config.env }),
  })
}
