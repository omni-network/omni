import {
  TOKEN_PROGRAM_ADDRESS,
  findAssociatedTokenPda,
} from '@solana-program/token'
import type { Address } from '@solana/kit'

export type GetTokenAccountParams = {
  owner: Address
  mint: Address
  tokenProgram?: Address
}

export async function getTokenAccount(
  params: GetTokenAccountParams,
): Promise<Address> {
  const { tokenProgram = TOKEN_PROGRAM_ADDRESS, ...rest } = params
  const [account] = await findAssociatedTokenPda({ tokenProgram, ...rest })
  return account
}
