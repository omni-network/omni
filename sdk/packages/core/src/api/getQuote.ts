import * as z from '@zod/mini'
import { zeroAddress } from 'viem'
import { createSafeFetchRequest } from '../internal/api.js'
import type { Environment } from '../types/config.js'
import { addressSchema } from '../types/primitives.js'
import { type Quote, type Quoteable, quoteableSchema } from '../types/quote.js'
import { toJSON } from '../utils/toJSON.js'

// trim params to create obj expected by /quote endpoint
export const toQuoteUnit = (q: Quoteable, omitAmount: boolean) => ({
  amount: omitAmount ? undefined : q.amount,
  token: q.isNative ? zeroAddress : q.token,
})

export const quoteParametersSchema = z.pipe(
  z.object({
    srcChainId: z.optional(z.number()),
    destChainId: z.number(),
    mode: z.enum(['expense', 'deposit']),
    deposit: quoteableSchema,
    expense: quoteableSchema,
  }),
  z.transform((params) => {
    const {
      srcChainId,
      destChainId,
      deposit: depositInput,
      expense: expenseInput,
      mode,
    } = params
    return toJSON({
      sourceChainId: srcChainId,
      destChainId: destChainId,
      deposit: toQuoteUnit(depositInput, mode === 'deposit'),
      expense: toQuoteUnit(expenseInput, mode === 'expense'),
    })
  }),
)
export type GetQuoteParams = z.input<typeof quoteParametersSchema>

const quoteResponseEntrySchema = z.object({
  token: addressSchema,
  amount: z.coerce.bigint(), // string-encoded number we parse to bigint
})
export const quoteResponseSchema = z.object({
  deposit: quoteResponseEntrySchema,
  expense: quoteResponseEntrySchema,
})

export const safeFetchQuote = createSafeFetchRequest(
  '/quote',
  quoteParametersSchema,
  quoteResponseSchema,
)

// getQuote calls the /quote endpoint
export async function getQuote(
  quote: GetQuoteParams,
  environment?: Environment | string,
): Promise<Quote> {
  return await safeFetchQuote({ input: quote, environment }).getOrThrow()
}
