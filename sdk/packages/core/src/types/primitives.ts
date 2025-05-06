import * as z from '@zod/mini'
import { fromHex } from 'viem'
import type { ValidationSchema } from '../internal/validation.js'

export const hexStringSchema = z
  .string({ error: 'Value must be an hexadecimal string starting with "0x"' })
  .check(z.regex(/0x[0-9A-Fa-f]+/)) as ValidationSchema<`0x${string}`>

export const addressSchema = hexStringSchema.check(z.length(42))

export const hexToBigIntSchema = z.pipe(
  hexStringSchema,
  z.transform((input) => fromHex(input, 'bigint')),
)
