export function bytesFromU64(n: bigint): Uint8Array {
  const buffer = new ArrayBuffer(8)
  const view = new DataView(buffer)
  view.setBigUint64(0, n, true) // little endian
  return new Uint8Array(buffer)
}

export function bytesToU64(bytes: Uint8Array): bigint {
  return new DataView(bytes.buffer).getBigUint64(0, true) // little endian
}

export async function digestSHA256(
  ...inputs: Array<Uint8Array>
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
  const buffer = new Uint8Array(length)
  crypto.getRandomValues(buffer)
  return buffer
}

export function randomU64(): bigint {
  return bytesToU64(randomBytes(8))
}
