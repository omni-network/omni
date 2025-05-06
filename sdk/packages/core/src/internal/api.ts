import * as z from '@zod/mini'
import { type AsyncResult, Result } from 'typescript-result'
import type { Environment } from '../types/config.js'
import { getApiUrl } from '../utils/getApiUrl.js'
import { version } from '../version.js'
import {
  type ValidationError,
  type ValidationSchema,
  safeValidate,
  safeValidateAsync,
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

export type SafeFetchResponseError<O, I = O> =
  | SafeFetchError
  | ValidationError<O, I>

export type SafeFetchResponse<O, I = O> = AsyncResult<
  O,
  SafeFetchResponseError<O, I>
>

export type SafeFetchParameters = {
  environment?: Environment | string
  request?: RequestInit
}

export function createSafeFetchResponse<O, I = O>(
  endpoint: string,
  responseSchema: ValidationSchema<O, I>,
) {
  return function safeFetchResponse(
    params: SafeFetchParameters,
  ): SafeFetchResponse<O, I> {
    const url = getApiUrl(params.environment) + endpoint
    return safeFetchJSON(url, params.request).map((data) => {
      return safeValidate(responseSchema, data)
    })
  }
}

export type SafeFetchRequestError<
  RequestInput,
  ResponseOutput,
  ResponseInput = ResponseOutput,
> =
  | ValidationError<string, RequestInput>
  | SafeFetchResponseError<ResponseOutput, ResponseInput>

export type SafeFetchRequest<
  RequestInput,
  ResponseOutput,
  ResponseInput = ResponseOutput,
> = AsyncResult<
  ResponseOutput,
  SafeFetchRequestError<RequestInput, ResponseOutput, ResponseInput>
>

export type SafeFetchRequestParameters<I> = SafeFetchParameters & { input: I }

export function encodeRequest<I>(
  schema: ValidationSchema<string, I>,
  params: SafeFetchRequestParameters<I>,
): AsyncResult<RequestInit, ValidationError<string, I>> {
  return safeValidateAsync(schema, params.input).map((body) => {
    return {
      method: 'POST',
      headers: { 'Content-Type': 'application/json' },
      ...params.request,
      body,
    }
  })
}

export function createSafeFetchRequest<
  RequestInput,
  ResponseOutput,
  ResponseInput = ResponseOutput,
>(
  endpoint: string,
  requestSchema: ValidationSchema<string, RequestInput>,
  responseSchema: ValidationSchema<ResponseOutput, ResponseInput>,
) {
  const fetchResponse = createSafeFetchResponse(endpoint, responseSchema)

  return function safeFetchRequest(
    params: SafeFetchRequestParameters<RequestInput>,
  ): SafeFetchRequest<RequestInput, ResponseOutput, ResponseInput> {
    return encodeRequest(requestSchema, params).map((request) => {
      return fetchResponse({ request })
    })
  }
}
