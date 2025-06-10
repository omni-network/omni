import { type Hex, fromHex } from 'viem'
import { fetchJSON } from '../internal/api.js'
import type { Asset } from '../types/asset.js'
import type { Environment } from '../types/config.js'
import { getApiUrl } from '../utils/getApiUrl.js'

export function isAssets(json: unknown): json is Record<string, Asset[]> {
  function _isAssetObj(item: unknown): boolean {
    if (typeof item !== 'object' || item === null) {
      return false
    }

    const asset = item as unknown as Asset

    if (
      typeof asset.enabled === 'boolean' &&
      typeof asset.name === 'string' &&
      typeof asset.symbol === 'string' &&
      typeof asset.expenseMin === 'string' &&
      typeof asset.expenseMax === 'string' &&
      typeof asset.chainId === 'number' &&
      typeof asset.address === 'string' &&
      typeof asset.decimals === 'number'
    ) {
      const parsedExpenseMin = fromHex(asset.expenseMin as Hex, 'bigint')
      const parsedExpenseMax = fromHex(asset.expenseMax as Hex, 'bigint')

      // mutate in-place
      asset.expenseMin = parsedExpenseMin
      asset.expenseMax = parsedExpenseMax

      return true
    }

    return false
  }

  return (
    json !== null &&
    typeof json === 'object' &&
    'tokens' in json &&
    Array.isArray(json.tokens) &&
    json.tokens.every(_isAssetObj)
  )
}

export type GetAssetsParameters = {
  environment?: Environment | string
}

export async function getAssets(
  params: GetAssetsParameters = {},
): Promise<Asset[]> {
  const apiUrl = getApiUrl(params?.environment)
  const json = await fetchJSON(`${apiUrl}/tokens`)

  if (!isAssets(json)) throw new Error('Unexpected /tokens response')

  return json.tokens
}
