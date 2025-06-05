import type { Address, AnchorProvider } from '@coral-xyz/anchor'
import {
  type MessageV0,
  PublicKey,
  type TransactionInstruction,
  TransactionMessage,
} from '@solana/web3.js'
import type BN from 'bn.js'
import { type InboxProgram, createInboxProgram } from './inbox.js'
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

export type CreateOpenInstructionParams = {
  inboxProgram: InboxProgram
  order: CreateParams
  ownerTokenAccount: Address
  mintAccount: Address
}

export async function createOpenInstruction(
  params: CreateOpenInstructionParams,
): Promise<TransactionInstruction> {
  const { inboxProgram, order, ...accounts } = params
  const openParams = await createOpenParams(order)
  return await inboxProgram.methods
    .open(openParams)
    .accounts(accounts)
    .instruction()
}

export type CreateOpenMessageParams = {
  provider: AnchorProvider
  inboxProgram?: InboxProgram
  order: CreateParams
  ownerTokenAccount: Address
  mintAccount: Address
  recentBlockhash?: string
}

export async function createOpenMessage(
  params: CreateOpenMessageParams,
): Promise<MessageV0> {
  const {
    provider,
    recentBlockhash: maybeRecentBlockhash,
    ...instructionParams
  } = params
  const inboxProgram =
    params.inboxProgram ?? (await createInboxProgram(provider))
  const [openInstruction, recentBlockhash] = await Promise.all([
    createOpenInstruction({ ...instructionParams, inboxProgram }),
    maybeRecentBlockhash ??
      provider.connection
        .getLatestBlockhash()
        .then(({ blockhash }) => blockhash),
  ])
  return new TransactionMessage({
    payerKey: provider.wallet.publicKey,
    recentBlockhash,
    instructions: [openInstruction],
  }).compileToV0Message()
}
