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

// QuoteResponse is the response from the /quote endpoint, with hex encoded amounts
export type QuoteResponse = {
  deposit: { token: Address; amount: Hex }
  expense: { token: Address; amount: Hex }
}

export function encodeQuote(params: GetQuoteParams): string {
  const { srcChainId, destChainId, deposit, expense, mode } = params
  return toJSON({
    sourceChainId: srcChainId,
    destChainId: destChainId,
    deposit: toQuoteUnit(deposit, mode === 'deposit'),
    expense: toQuoteUnit(expense, mode === 'expense'),
  })
}

// getQuoteWithEncoded calls the /quote endpoint, throwing on error
export async function getQuoteWithEncoded(
  apiBaseUrl: string,
  encodedQuote: string,
): Promise<Quote> {
  const json = await fetchJSON(`${apiBaseUrl}/quote`, {
    method: 'POST',
    headers: { 'Content-Type': 'application/json' },
    body: encodedQuote,
  })

  if (!isQuoteRes(json)) {
    throw new Error(`Unexpected quote response: ${JSON.stringify(json)}`)
  }

  const { deposit, expense } = json
  return {
    deposit: { ...deposit, amount: fromHex(deposit.amount, 'bigint') },
    expense: { ...expense, amount: fromHex(expense.amount, 'bigint') },
  } as Quote
}

export async function getQuote(
  apiBaseUrl: string,
  params: GetQuoteParams,
): Promise<Quote> {
  const encoded = encodeQuote(params)
  return await getQuoteWithEncoded(apiBaseUrl, encoded)
}

// toQuoteUnit translates a Quoteable to "QuoteUnit", the format expected by /quote
export const toQuoteUnit = (q: Quoteable, omitAmount: boolean) => ({
  amount: omitAmount ? undefined : q.amount,
  token: q.isNative ? zeroAddress : q.token,
})

// isQuoteRes checks if a json is a QuoteResponse
// TODO: use zod
function isQuoteRes(json: unknown): json is QuoteResponse {
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
