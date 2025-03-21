import { toHex } from 'viem'

export const toJSON = (value: unknown) =>
  JSON.stringify(value, (_, value) =>
    typeof value === 'bigint' ? toHex(value) : value,
  )
