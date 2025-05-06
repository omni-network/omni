import * as z from '@zod/mini'
import { addressSchema } from './primitives.js'

export const omniContractsSchema = z.looseObject({
  inbox: addressSchema,
  outbox: addressSchema,
  middleman: addressSchema,
})
export type OmniContracts = z.infer<typeof omniContractsSchema>
