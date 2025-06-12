import { type Hex, fromHex, zeroAddress } from 'viem'
import { string, z } from 'zod/v4-mini'
import { fetchJSON } from '../internal/api.js'
import { address, evmAddress, hex } from '../schema/types.js'
import type { Environment } from '../types/config.js'
import type { Quote, Quoteable } from '../types/quote.js'
import type { Prettify } from '../types/utils.js'
import { getApiUrl, toJSON } from '../utils/index.js'

export const quoteResponseSchema = z.object({
  deposit: z.object({
    token: address(),
    amount: z.union([hex(), string()]),
  }),
  expense: z.object({
    token: evmAddress(),
    amount: z.union([hex(), string()]),
  }),
})

export type GetQuoteParameters = Prettify<
  {
    srcChainId?: number
    destChainId: number
    environment?: Environment | string
  } & (
    | {
        mode: 'deposit'
        deposit?: Prettify<Omit<Quoteable, 'amount'>>
        expense: Prettify<Omit<Quoteable, 'amount'> & { amount: bigint }>
      }
    | {
        mode: 'expense'
        deposit: Prettify<Omit<Quoteable, 'amount'> & { amount: bigint }>
        expense?: Prettify<Omit<Quoteable, 'amount'>>
      }
  )
>

// the response from /quote endpoint (amounts are hex encoded bigints)
type QuoteResponse = z.infer<typeof quoteResponseSchema>

// getQuote calls the /quote endpoint
export async function getQuote(quote: GetQuoteParameters): Promise<Quote> {
  const {
    srcChainId,
    destChainId,
    deposit: depositInput,
    expense: expenseInput,
    mode,
    environment,
  } = quote
  const apiUrl = getApiUrl(environment)
  const json = await fetchJSON(`${apiUrl}/quote`, {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: toJSON({
      sourceChainId: srcChainId,
      destChainId: destChainId,
      deposit: depositInput
        ? toQuoteUnit(depositInput, mode === 'deposit')
        : {},
      expense: expenseInput
        ? toQuoteUnit(expenseInput, mode === 'expense')
        : {},
    }),
  })

  if (!isQuoteRes(json)) {
    throw new Error(`Unexpected quote response: ${JSON.stringify(json)}`)
  }

  const { deposit, expense } = json

  return {
    deposit: { ...deposit, amount: fromHex(deposit.amount as Hex, 'bigint') },
    expense: { ...expense, amount: fromHex(expense.amount as Hex, 'bigint') },
  } satisfies Quote
}

// trim params to create obj expected by /quote endpoint
export const toQuoteUnit = (q: Quoteable, omitAmount: boolean) => ({
  amount: omitAmount ? undefined : q.amount,
  token: q.token ?? zeroAddress,
})

// asserts a json response is QuoteResponse
function isQuoteRes(json: unknown): json is QuoteResponse {
  return quoteResponseSchema.safeParse(json).success
}
