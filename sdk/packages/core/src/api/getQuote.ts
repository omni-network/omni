import { type Address, type Hex, fromHex, zeroAddress } from 'viem'
import { fetchJSON } from '../internal/api.js'
import type { Environment } from '../types/config.js'
import type { Quote, Quoteable } from '../types/quote.js'
import { getApiUrl } from '../utils/getApiUrl.js'
import { toJSON } from '../utils/toJSON.js'

export type GetQuoteParameters = {
  srcChainId?: number
  destChainId: number
  environment?: Environment | string
} & (
  | {
      mode: 'deposit'
      deposit: Omit<Quoteable, 'amount'>
      expense: Omit<Quoteable, 'amount'> & { amount: bigint }
    }
  | {
      mode: 'expense'
      deposit: Omit<Quoteable, 'amount'> & { amount: bigint }
      expense: Omit<Quoteable, 'amount'>
    }
)

// the response from /quote endpoint (amounts are hex encoded bigints)
type QuoteResponse = {
  deposit: { token: Address; amount: Hex }
  expense: { token: Address; amount: Hex }
}

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
      deposit: toQuoteUnit(depositInput, mode === 'deposit'),
      expense: toQuoteUnit(expenseInput, mode === 'expense'),
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
  token: q.token ?? zeroAddress,
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
