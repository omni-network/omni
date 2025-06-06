import { getAddressDecoder, getAddressEncoder } from '@solana/kit'

export const addressDecoder = getAddressDecoder()
export const addressEncoder = getAddressEncoder()

export const textEncoder = new TextEncoder()

export function encodeU64(n: bigint): Uint8Array {
  const bytes = new Uint8Array(8)
  new DataView(bytes.buffer).setBigUint64(0, n, true) // little endian
  return bytes
}

export function decodeU64(bytes: Uint8Array): bigint {
  return new DataView(bytes.buffer).getBigUint64(0, true) // little endian
}
