import * as z from '@zod/mini'
import { type AsyncResult, Result } from 'typescript-result'
import { version } from '../version.js'
import {
  type ValidationError,
  type ValidationSchema,
  safeValidate,
} from './validation.js'

// TODO: remove these types, only kept there while they are used by the React package

// JSONError is a json error response from the solver api
export type JSONError = {
  code: number
  status: string
  message: string
}
// FetchJSONError is a JSONError or a generic Error (e.g. network error)
export type FetchJSONError = JSONError | Error

const APIErrorSchema = z.object({
  code: z.number(),
  status: z.string(),
  message: z.string(),
})
type APIErrorObject = z.infer<typeof APIErrorSchema>

class APIError extends Error implements APIErrorObject {
  readonly type = 'APIError'
  readonly code: number
  readonly status: string

  constructor(data: APIErrorObject) {
    super(data.message)
    this.code = data.code
    this.status = data.status
  }
}

class FetchError extends Error {
  readonly type = 'FetchError'
}

type RequestErrorParams = ErrorOptions & {
  response: Response
  message?: string
}

class RequestError extends Error {
  readonly type = 'RequestError'
  readonly response: Response

  constructor(params: RequestErrorParams) {
    const { response, message, ...options } = params
    super(
      message ??
        `Request failed with status ${response.status}: ${response.statusText}`,
      options,
    )
    this.response = response
  }
}

export type SafeFetchError = FetchError | RequestError | APIError

export function safeFetchJSON<Value = Record<string, unknown>>(
  url: string,
  init?: RequestInit,
): AsyncResult<Value, SafeFetchError> {
  const headers = new Headers(init?.headers)
  headers.set('user-agent', `@omni-network/core:${version}`)
  return Result.fromAsyncCatching(fetch(url, { ...init, headers }))
    .mapError((cause) => {
      // Request failed before returning a response
      return new FetchError(`Failed to fetch ${url}`, { cause })
    })
    .map((response) => {
      return Result.fromAsyncCatching(response.json())
        .map((data) => {
          if (response.ok) {
            return Result.ok(data as Value)
          }
          const parsedAPIError = APIErrorSchema.safeParse(
            (data as Record<string, unknown>).error,
          )
          return Result.error(
            parsedAPIError.success
              ? new APIError(parsedAPIError.data)
              : new RequestError({ response }),
          )
        })
        .mapError((cause) => {
          return new RequestError({
            response,
            message: 'Failed to parse JSON response',
            cause,
          })
        })
    })
}

export type SafeFetchTypeError<T> = SafeFetchError | ValidationError<T>

export type SafeFetchTypeResult<T> = AsyncResult<T, SafeFetchTypeError<T>>

export function createSafeFetchType<T>(schema: ValidationSchema<T>) {
  return function safeFetchType(
    url: string,
    init?: RequestInit,
  ): AsyncResult<T, SafeFetchTypeError<T>> {
    return safeFetchJSON(url, init).map((data) => safeValidate(schema, data))
  }
}
