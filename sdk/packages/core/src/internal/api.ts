import { version } from '../version.js'
import { z } from 'zod'

const jsonErrorSchema = z.object({
  code: z.number(),
  status: z.string(),
  message: z.string(),
})

// JSONError is a json error response from the solver api
export type JSONError = z.infer<typeof jsonErrorSchema>

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
    if (!isJSONError(json.error)) {
      throw new Error(`${res.status} ${res.statusText}`)
    }
    throw json.error
  }

  return json
}

function isJSONError(error: unknown): error is JSONError {
  return jsonErrorSchema.safeParse(error).success
}
