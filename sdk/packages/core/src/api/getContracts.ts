import {
  type SafeFetchTypeResult,
  createSafeFetchType,
} from '../internal/api.js'
import type { Environment } from '../types/config.js'
import { type OmniContracts, omniContractsSchema } from '../types/contracts.js'
import { getApiUrl } from '../utils/getApiUrl.js'

export const safeFetchContracts = createSafeFetchType(omniContractsSchema)

export function safeGetContracts(
  envOrApiBaseUrl?: Environment | string,
): SafeFetchTypeResult<OmniContracts> {
  const apiUrl = getApiUrl(envOrApiBaseUrl)
  return safeFetchContracts(`${apiUrl}/contracts`)
}

export async function getContracts(
  envOrApiBaseUrl?: Environment | string,
): Promise<OmniContracts> {
  return await safeGetContracts(envOrApiBaseUrl).getOrThrow()
}
