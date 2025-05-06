import { createSafeFetchResponse } from '../internal/api.js'
import type { Environment } from '../types/config.js'
import { type OmniContracts, omniContractsSchema } from '../types/contracts.js'

export const safeFetchContracts = createSafeFetchResponse(
  '/contracts',
  omniContractsSchema,
)

export type GetContractsParameters = {
  environment?: Environment | string
}

export async function getContracts(
  params: GetContractsParameters,
): Promise<OmniContracts> {
  return await safeFetchContracts(params).getOrThrow()
}
