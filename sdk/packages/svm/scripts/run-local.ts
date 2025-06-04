import { readFileSync } from 'node:fs'
import { Program } from '@coral-xyz/anchor'
import {
  Connection,
  Keypair,
  LAMPORTS_PER_SOL,
  PublicKey,
} from '@solana/web3.js'
import BN from 'bn.js'
import bs58 from 'bs58'
import { createOpenParams, inboxIDL } from '../dist/esm/index.js'
import type { SolverInbox } from '../dist/types/index.js'

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
  publicKey: keypair.publicKey,
})

const openParams = await createOpenParams({
  owner: keypair.publicKey,
  depositAmount: new BN(1000),
  destChainId: new BN(1),
  call: {
    target: new Array(20).fill(0),
    selector: new Array(4).fill(0),
    value: new BN(0),
    params: Buffer.from([]),
  },
  expense: {
    spender: new Array(20).fill(0),
    token: new Array(20).fill(0),
    amount: new BN(0),
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
console.log('open instruction', openInstruction)
