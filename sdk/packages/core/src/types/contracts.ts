import * as z from '@zod/mini'
import { hexStringSchema } from '../internal/validation.js'

export const omniContractsSchema = z.looseObject({
  inbox: hexStringSchema,
  outbox: hexStringSchema,
  middleman: hexStringSchema,
})

export type OmniContracts = z.infer<typeof omniContractsSchema>
