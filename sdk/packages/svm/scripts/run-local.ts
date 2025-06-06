import { readFileSync } from 'node:fs'
import {
  address,
  createKeyPairSignerFromBytes,
  createSolanaRpc,
  sendTransactionWithoutConfirmingFactory,
  setTransactionMessageLifetimeUsingBlockhash,
  signTransactionMessageWithSigners,
} from '@solana/kit'
import bs58 from 'bs58'
import { createOpenOrderTransactionMessage } from '../dist/esm/index.js'

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

const orderTransactionMessage = await createOpenOrderTransactionMessage({
  owner: signer,
  mint: usdcMintAccount,
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
})
console.log('order transaction message:', orderTransactionMessage)

const recentBlockhash = await rpc.getLatestBlockhash().send()
console.log('recent blockhash:', recentBlockhash)

const transactionMessage = setTransactionMessageLifetimeUsingBlockhash(
  recentBlockhash.value,
  orderTransactionMessage,
)
console.log('transaction message:', transactionMessage)

const signedTransaction =
  await signTransactionMessageWithSigners(transactionMessage)
console.log('signed transaction:', signedTransaction)

const sendTransaction = sendTransactionWithoutConfirmingFactory({ rpc })
await sendTransaction(signedTransaction, { commitment: 'confirmed' })
