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
  const res = await fetch(url, init)
  const json = await res.json()

  if (!res.ok) {
    if (!isJSONError(json.error)) new Error(`${res.status} ${res.statusText}`)
    throw json.error
  }

  return json
}

// TODO: use zod
function isJSONError(error: unknown): error is JSONError {
  return (
    error != null &&
    typeof (error as any).code === 'number' &&
    typeof (error as any).status === 'string' &&
    typeof (error as any).message === 'string'
  )
}
