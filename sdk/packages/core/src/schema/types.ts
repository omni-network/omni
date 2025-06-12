import type { Hex } from 'viem'
import { z } from 'zod/v4-mini'
import type { Address, EVMAddress, SVMAddress } from '../types/addresses.js'

export const hex = () =>
  z.string().check(
    z.regex(/^0x[0-9a-fA-F]+$/, {
      message:
        'Invalid hex string: must start with 0x and contain only hex characters',
    }),
  ) as z.ZodMiniType<Hex>

export const evmAddress = () =>
  z.string().check(
    z.regex(/^0x[0-9a-fA-F]{40}$/, {
      message: 'Invalid EVM address: must start with 0x and be 20 bytes long',
    }),
  ) as z.ZodMiniType<EVMAddress>

export const svmAddress = () =>
  z.string().check(
    z.regex(/^[1-9A-HJ-NP-Za-km-z]{32,44}$/, {
      message:
        'Invalid SVM address: must be base58 encoded and 32 to 44 characters long',
    }),
  ) as unknown as z.ZodMiniType<SVMAddress>

export const address = () =>
  z.union([evmAddress(), svmAddress()]) as z.ZodMiniType<Address>
