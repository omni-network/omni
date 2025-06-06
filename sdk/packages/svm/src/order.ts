import {
  type Address,
  type ReadonlyUint8Array,
  type TransactionSigner,
  appendTransactionMessageInstruction,
  createTransactionMessage,
  getProgramDerivedAddress,
  pipe,
  setTransactionMessageFeePayerSigner,
} from '@solana/kit'
import {
  type EVMCallArgs,
  type EVMTokenExpenseArgs,
  getOpenInstructionAsync,
} from './__generated__/inbox/index.js'
import {
  addressDecoder,
  addressEncoder,
  encodeU64,
  textEncoder,
} from './codecs.js'
import { digestSHA256, randomU64 } from './crypto.js'
import { inboxProgramAddress } from './inbox.js'
import { getTokenAccount } from './token.js'

const orderTokenSeed = textEncoder.encode('order_token')
const orderStateSeed = textEncoder.encode('order_state')

export async function getOrderId(
  owner: Address,
  nonce: bigint,
): Promise<Address> {
  const bytes = await digestSHA256(
    addressEncoder.encode(owner),
    encodeU64(nonce),
  )
  return addressDecoder.decode(bytes)
}

export async function getInboxDerivedAddress(
  seeds: Array<Uint8Array | ReadonlyUint8Array>,
): Promise<Address> {
  const [address] = await getProgramDerivedAddress({
    programAddress: inboxProgramAddress,
    seeds,
  })
  return address
}

export type OrderAccounts = {
  orderState: Address
  orderTokenAccount: Address
}

export async function getOrderAccounts(
  orderId: Address,
): Promise<OrderAccounts> {
  const orderIdSeed = addressEncoder.encode(orderId)
  const [orderState, orderTokenAccount] = await Promise.all([
    getInboxDerivedAddress([orderStateSeed, orderIdSeed]),
    getInboxDerivedAddress([orderTokenSeed, orderIdSeed]),
  ])
  return { orderState, orderTokenAccount }
}

export type OpenOrderInstructionParams = {
  owner: TransactionSigner
  mint: Address
  ownerTokenAccount?: Address
  depositAmount: number | bigint
  destChainId: number | bigint
  call: EVMCallArgs
  expense: EVMTokenExpenseArgs
  nonce?: bigint
}

export async function getOpenOrderInstruction(
  params: OpenOrderInstructionParams,
) {
  const {
    owner,
    mint,
    nonce: maybeNonce,
    ownerTokenAccount: maybeOwnerTokenAccount,
    ...rest
  } = params
  const nonce = maybeNonce ?? randomU64()
  const orderId = await getOrderId(owner.address, nonce)
  const [orderAccounts, ownerTokenAccount] = await Promise.all([
    getOrderAccounts(orderId),
    maybeOwnerTokenAccount ?? getTokenAccount({ owner: owner.address, mint }),
  ])
  return await getOpenInstructionAsync({
    owner,
    orderId,
    nonce,
    mintAccount: mint,
    ownerTokenAccount,
    ...orderAccounts,
    ...rest,
  })
}

export async function createOpenOrderTransactionMessage(
  params: OpenOrderInstructionParams,
) {
  const instruction = await getOpenOrderInstruction(params)
  return pipe(
    createTransactionMessage({ version: 0 }),
    (tx) => setTransactionMessageFeePayerSigner(params.owner, tx),
    (tx) => appendTransactionMessageInstruction(instruction, tx),
  )
}
