import { PublicKey } from '@solana/web3.js'
import type BN from 'bn.js'
import type { EvmCall, EvmTokenExpense } from './types.js'
import { bytesFromU64, digestSHA256, randomU64 } from './utils.js'

export async function createOrderId(
  owner: PublicKey,
  nonce: BN,
): Promise<PublicKey> {
  const bytes = await digestSHA256(owner.toBuffer(), bytesFromU64(nonce))
  return new PublicKey(bytes)
}

export type OpenParams = {
  orderId: PublicKey
  nonce: BN // u64
  depositAmount: BN // u64
  destChainId: BN // u64
  call: EvmCall
  expense: EvmTokenExpense
}

export type CreateParams = {
  owner: PublicKey
  nonce?: BN
  depositAmount: BN
  destChainId: BN
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
