import type { ReadonlyUint8Array } from '@solana/kit'
import { decodeU64 } from './codecs.js'

export async function digestSHA256(
  ...inputs: Array<Uint8Array | ReadonlyUint8Array>
): Promise<Uint8Array> {
  const length = inputs.reduce((acc, input) => acc + input.length, 0)
  const data = new Uint8Array(length)
  let offset = 0
  for (const input of inputs) {
    data.set(input, offset)
    offset += input.length
  }
  const hash = await crypto.subtle.digest('SHA-256', data)
  return new Uint8Array(hash)
}

export function randomBytes(length: number): Uint8Array {
  const bytes = new Uint8Array(length)
  crypto.getRandomValues(bytes)
  return bytes
}

export function randomU64(): bigint {
  return decodeU64(randomBytes(8))
}
