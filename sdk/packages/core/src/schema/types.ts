import type { Address, Hex } from 'viem'
import { z, ZodType } from 'zod'

export const hex = (): ZodType<Hex> =>
  z
  .string()
  .refine((val): val is Hex => /^0x[0-9a-fA-F]+$/.test(val), {
    message: 'Invalid hex string: must start with 0x and contain only hex characters',
  }) as unknown as ZodType<Hex>

export const address = (): ZodType<Address> =>
  z
  .string()
  .refine((val): val is Address => /^0x[0-9a-fA-F]{40}$/.test(val), {
    message: 'Invalid address: must start with 0x and be 20 bytes long',
  }) as unknown as ZodType<Address>