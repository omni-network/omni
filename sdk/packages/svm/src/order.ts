import { PublicKey } from '@solana/web3.js'
import type { EvmCall, EvmTokenExpense } from './types.js'
import { bytesFromU64, digestSHA256, randomU64 } from './utils.js'

export async function createOrderId(
  owner: PublicKey,
  nonce: bigint,
): Promise<PublicKey> {
  const bytes = await digestSHA256(owner.toBuffer(), bytesFromU64(nonce))
  return new PublicKey(bytes)
}

export type OpenParams = {
  orderId: PublicKey
  nonce: bigint // u64
  depositAmount: bigint // u64
  destChainId: bigint // u64
  call: EvmCall
  expense: EvmTokenExpense
}

export type CreateParams = {
  owner: PublicKey
  nonce?: bigint
  depositAmount: bigint
  destChainId: bigint
  call: EvmCall
  expense: EvmTokenExpense
}

export async function createOpenParams(
  params: CreateParams,
): Promise<OpenParams> {
  const { owner, nonce: maybeNonce, ...rest } = params
  const nonce = maybeNonce ?? randomU64()
  const orderId = await createOrderId(owner, nonce)
  return { orderId, nonce, ...rest }
}
