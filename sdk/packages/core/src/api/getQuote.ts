import { type Address, type Hex, fromHex, zeroAddress } from 'viem'
import { fetchJSON } from '../internal/api.js'
import type { Quote, Quoteable } from '../types/quote.js'
import { toJSON } from '../utils/toJSON.js'

export type GetQuoteParams = {
  srcChainId?: number
  destChainId: number
  mode: 'expense' | 'deposit'
  deposit: Quoteable
  expense: Quoteable
}

// the response from /quote endpoint (amounts are hex encoded bigints)
type QuoteResponse = {
  deposit: { token: Address; amount: Hex }
  expense: { token: Address; amount: Hex }
}

// getQuote calls the /quote endpoint
export async function getQuote(
  apiBaseUrl: string,
  quote: GetQuoteParams,
): Promise<Quote> {
  const json = await fetchJSON(`${apiBaseUrl}/quote`, {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: toJSON({
      ...quote,
      deposit: toQuoteUnit(quote.deposit, quote.mode === 'deposit'),
      expense: toQuoteUnit(quote.expense, quote.mode === 'expense'),
    }),
  })

  if (!isQuoteRes(json)) {
    throw new Error(`Unexpected quote response: ${JSON.stringify(json)}`)
  }

  const { deposit, expense } = json

  return {
    deposit: { ...deposit, amount: fromHex(deposit.amount, 'bigint') },
    expense: { ...expense, amount: fromHex(expense.amount, 'bigint') },
  } satisfies Quote
}

// trim params to create obj expected by /quote endpoint
export const toQuoteUnit = (q: Quoteable, omitAmount: boolean) => ({
  amount: omitAmount ? undefined : q.amount,
  token: q.isNative ? zeroAddress : q.token,
})

// asserts a json response is QuoteResponse
function isQuoteRes(json: unknown): json is QuoteResponse {
  // TODO: schema validation
  const quote = json as QuoteResponse
  return (
    json != null &&
    quote.deposit != null &&
    quote.expense != null &&
    typeof quote.deposit.token === 'string' &&
    typeof quote.deposit.amount === 'string' &&
    typeof quote.expense.token === 'string' &&
    typeof quote.expense.amount === 'string'
  )
}
