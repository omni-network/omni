import type { Address, Hex } from 'viem'
import { z } from 'zod/v4-mini'
export const hex = () =>
  z.string().check(
    z.regex(/^0x[0-9a-fA-F]+$/, {
      message:
        'Invalid hex string: must start with 0x and contain only hex characters',
    }),
  ) as z.ZodMiniType<Hex>

export const address = () =>
  z.string().check(
    z.regex(/^0x[0-9a-fA-F]{40}$/, {
      message: 'Invalid address: must start with 0x and be 20 bytes long',
    }),
  ) as z.ZodMiniType<Address>
