import { readFileSync } from 'node:fs'
import { Program } from '@coral-xyz/anchor'
import {
  Connection,
  Keypair,
  LAMPORTS_PER_SOL,
  PublicKey,
} from '@solana/web3.js'
import bs58 from 'bs58'
import type { SolverInbox } from './idl/solver_inbox.js'
import inboxIDL from './idl/solver_inbox.json' with { type: 'json' }
import { createOpenParams } from './order.js'

const config = JSON.parse(readFileSync('/tmp/svm/config.json', 'utf8'))

const usdcMintAccount = new PublicKey(config.mints[0].mint_account)
console.log('USDC mint account', usdcMintAccount.toBase58())

const connection = new Connection(config.rpc_address, 'confirmed')
const keypair = Keypair.fromSecretKey(bs58.decode(config.authority_key))
console.log('using account', keypair.publicKey.toBase58())

const balance = await connection.getBalance(keypair.publicKey)
console.log('account balance in SOL:', balance / LAMPORTS_PER_SOL)

export const inboxProgram = new Program<SolverInbox>(inboxIDL as SolverInbox, {
  connection,
})

const openParams = await createOpenParams({
  owner: keypair.publicKey,
  depositAmount: 1000n,
  destChainId: 1n,
  call: {
    target: new Uint8Array(20),
    selector: new Uint8Array(4),
    value: 0n,
    params: new Uint8Array(),
  },
  expense: {
    spender: new Uint8Array(20),
    token: new Uint8Array(20),
    amount: 0n,
  },
})

const openInstruction = await inboxProgram.methods
  .open(openParams)
  .accounts({
    ownerTokenAccount: keypair.publicKey,
    mintAccount: usdcMintAccount,
  })
  .signers([keypair])
  .instruction()
console.log('instruction', openInstruction)

// const inboxAccounts = await inboxProgram.account.inboxState.all();
// console.log(inboxAccounts);

// const orderAccounts = await inboxProgram.account.orderState.all();
// console.log(orderAccounts);
