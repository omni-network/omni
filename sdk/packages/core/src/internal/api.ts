import { version } from '../version.js'

// JSONError is a json error response from the solver api
export type JSONError = {
  code: number
  status: string
  message: string
}

// FetchJSONError is a JSONError or a generic Error (e.g. network error)
export type FetchJSONError = JSONError | Error

export async function fetchJSON(
  url: string,
  init?: RequestInit,
): Promise<unknown> {
  const headers = new Headers(init?.headers)
  headers.set('user-agent', `@omni-network/core:${version}`)
  const _init = { ...init, headers }

  const res = await fetch(url, _init)
  const json = await res.json()

  if (!res.ok) {
    if (!isJSONError(json.error)) new Error(`${res.status} ${res.statusText}`)
    throw json.error
  }

  return json
}

// TODO: use zod
function isJSONError(error: unknown): error is JSONError {
  const err = error as JSONError
  return (
    err != null &&
    typeof err.code === 'number' &&
    typeof err.status === 'string' &&
    typeof err.message === 'string'
  )
}
