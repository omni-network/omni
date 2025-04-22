import { toHex } from 'viem'

const convertBigIntsToHex = (value: unknown): unknown => {
  if (typeof value === 'bigint') {
    return toHex(value)
  }

  if (Array.isArray(value)) {
    return value.map(convertBigIntsToHex)
  }

  if (typeof value === 'object' && value !== null) {
    const newValue: Record<string, unknown> = {}
    for (const key in value) {
      if (Object.prototype.hasOwnProperty.call(value, key)) {
        newValue[key] = convertBigIntsToHex(
          (value as Record<string, unknown>)[key],
        )
      }
    }
    return newValue
  }

  return value
}

export const toJSON = (value: unknown) => {
  // stringify w/out the replacer callback since we've seen browser
  // extensions bypass the replacer by attaching a toJSON method to BigInt
  // https://developer.mozilla.org/en-US/docs/Web/JavaScript/Reference/Global_Objects/JSON/stringify#:~:text=Attempting%20to%20serialize,by%20the%20user
  return JSON.stringify(convertBigIntsToHex(value))
}
