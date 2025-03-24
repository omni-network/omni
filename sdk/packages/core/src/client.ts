import { up } from 'up-fetch'

import {
  type ContractAddresses,
  type Hex,
  type Order,
  type Quote,
  type ValidationResult,
  contractAddressesSchema,
  quoteSchema,
  validationResultSchema,
} from './schemas.js'

const DEFAULT_USER_AGENT = '@omni-network/core:0.1.0'

type UpFetch = ReturnType<typeof up>

export type SolverClientParams = {
  baseUrl: string
  userAgent?: string
}

export type Quoteable = {
  token: Hex
  amount?: Hex
}

export type QuoteRequest = {
  sourceChainId?: number
  destChainId: number
  deposit: Quoteable
  expense: Quoteable
}

export type GetQuoteParams = {
  cacheKey?: string
  request: QuoteRequest
}

export type ValidateOrderParams = {
  cacheKey?: string
  order: Order
}

export class SolverClient {
  #requests: Map<string, Promise<unknown>> = new Map()
  #upfetch: UpFetch

  constructor(params: SolverClientParams) {
    const userAgent = params.userAgent ?? DEFAULT_USER_AGENT
    this.#upfetch = up(fetch, () => ({
      baseUrl: params.baseUrl,
      headers: { 'user-agent': userAgent },
    }))
  }

  #queryCache<T>(key: string, runQuery: () => Promise<T>): Promise<T> {
    let request = this.#requests.get(key)
    if (request == null) {
      request = runQuery()
      this.#requests.set(key, request)
    }
    return request as Promise<T>
  }

  getContracts(): Promise<ContractAddresses> {
    return this.#queryCache('contracts', () => {
      return this.#upfetch('contracts', { schema: contractAddressesSchema })
    })
  }

  async #getQuote(request: QuoteRequest): Promise<Quote> {
    return await this.#upfetch('quote', {
      method: 'POST',
      body: request,
      schema: quoteSchema,
    })
  }

  getQuote(params: GetQuoteParams): Promise<Quote> {
    return params.cacheKey == null
      ? this.#getQuote(params.request)
      : this.#queryCache(params.cacheKey, () => this.#getQuote(params.request))
  }

  async #checkOrder(order: Order): Promise<ValidationResult> {
    return await this.#upfetch('check', {
      method: 'POST',
      body: order,
      schema: validationResultSchema,
    })
  }

  validateOrder(params: ValidateOrderParams): Promise<ValidationResult> {
    return params.cacheKey == null
      ? this.#checkOrder(params.order)
      : this.#queryCache(params.cacheKey, () => this.#checkOrder(params.order))
  }
}
