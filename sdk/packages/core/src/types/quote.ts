import * as z from '@zod/mini'
import { addressSchema } from './primitives.js'

export const quoteableSchema = z.union([
  z.object({ isNative: z.literal(true), amount: z.optional(z.bigint()) }),
  z.object({
    isNative: z.literal(false),
    token: addressSchema,
    amount: z.optional(z.bigint()),
  }),
])
export type Quoteable = z.infer<typeof quoteableSchema>

export const quoteEntrySchema = z.object({
  token: addressSchema,
  amount: z.bigint(),
})
export type QuoteEntry = z.infer<typeof quoteEntrySchema>

export const quoteSchema = z.object({
  deposit: quoteEntrySchema,
  expense: quoteEntrySchema,
})
export type Quote = z.infer<typeof quoteSchema>
