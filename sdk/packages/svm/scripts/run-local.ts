import { readFileSync } from 'node:fs'
import {
  TOKEN_PROGRAM_ADDRESS,
  findAssociatedTokenPda,
} from '@solana-program/token'
import {
  address,
  appendTransactionMessageInstruction,
  createKeyPairSignerFromBytes,
  createSolanaRpc,
  createTransactionMessage,
  getAddressEncoder,
  getProgramDerivedAddress,
  pipe,
  sendTransactionWithoutConfirmingFactory,
  setTransactionMessageFeePayerSigner,
  setTransactionMessageLifetimeUsingBlockhash,
  signTransactionMessageWithSigners,
} from '@solana/kit'
import bs58 from 'bs58'
import { getOrderId, inboxClient, randomU64 } from '../dist/esm/index.js'
import type { inboxClient as InboxClient } from '../dist/types/index.js'

const config = JSON.parse(readFileSync('/tmp/svm/config.json', 'utf8'))
const usdcMint = config.mints[0]

const usdcMintAccount = address(usdcMint.mint_account)
console.log('USDC mint account:', usdcMintAccount)

const signer = await createKeyPairSignerFromBytes(
  bs58.decode(config.authority_key),
)
console.log('using EOA account:', signer.address)

const rpc = createSolanaRpc(config.rpc_address)

const balance = await rpc.getBalance(signer.address).send()
console.log('EOA account balance in lamports:', balance.value)

const nonce = randomU64()
const orderId = await getOrderId(signer.address, nonce)
console.log('orderId:', orderId)

const [ownerTokenAccount] = await findAssociatedTokenPda({
  owner: signer.address,
  mint: usdcMintAccount,
  tokenProgram: TOKEN_PROGRAM_ADDRESS,
})

const programAddress = address(inboxClient.SOLVER_INBOX_PROGRAM_ADDRESS)
const addressEncoder = getAddressEncoder()
const encodedOrderId = addressEncoder.encode(orderId)
const textEncoder = new TextEncoder()

const [orderTokenAccount] = await getProgramDerivedAddress({
  programAddress,
  seeds: [textEncoder.encode('order_token'), encodedOrderId],
})

const [orderState] = await getProgramDerivedAddress({
  programAddress,
  seeds: [textEncoder.encode('order_state'), encodedOrderId],
})

const openInput: InboxClient.OpenAsyncInput = {
  owner: signer,
  orderId,
  orderState,
  mintAccount: usdcMintAccount,
  orderTokenAccount,
  ownerTokenAccount,
  nonce,
  destChainId: 1n,
  depositAmount: 1000n,
  call: {
    target: new Uint8Array(20), // EVM address
    selector: new Uint8Array(4), // EVM selector
    value: 0n,
    params: new Uint8Array(0),
  },
  expense: {
    spender: new Uint8Array(20), // EVM address
    token: new Uint8Array(20), // EVM address
    amount: 1000n,
  },
}

const openOrderInstruction =
  await inboxClient.getOpenInstructionAsync(openInput)
console.log('open order instruction:', openOrderInstruction)

const recentBlockhash = await rpc.getLatestBlockhash().send()
console.log('recent blockhash:', recentBlockhash)

const transactionMessage = pipe(
  createTransactionMessage({ version: 0 }),
  (tx) => setTransactionMessageFeePayerSigner(signer, tx),
  (tx) =>
    setTransactionMessageLifetimeUsingBlockhash(recentBlockhash.value, tx),
  (tx) => appendTransactionMessageInstruction(openOrderInstruction, tx),
)
console.log('transaction message:', transactionMessage)

const signedTransaction =
  await signTransactionMessageWithSigners(transactionMessage)
console.log('signed transaction:', signedTransaction)

const sendTransaction = sendTransactionWithoutConfirmingFactory({ rpc })
await sendTransaction(signedTransaction, { commitment: 'confirmed' })
