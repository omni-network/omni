// import { readFileSync } from 'node:fs'
// import { AnchorProvider, Wallet } from '@coral-xyz/anchor'
// import { getOrCreateAssociatedTokenAccount } from '@solana/spl-token'
// import {
//   Connection,
//   Keypair,
//   LAMPORTS_PER_SOL,
//   PublicKey,
//   VersionedTransaction,
// } from '@solana/web3.js'
// import BN from 'bn.js'
// import bs58 from 'bs58'
// import { createOpenMessage } from '../dist/esm/index.js'

// const config = JSON.parse(readFileSync('/tmp/svm/config.json', 'utf8'))

// const usdcMintAccount = new PublicKey(config.mints[0].mint_account)
// console.log('USDC mint account', usdcMintAccount.toBase58())

// const connection = new Connection(config.rpc_address, 'confirmed')
// const keypair = Keypair.fromSecretKey(bs58.decode(config.authority_key))
// console.log('using account', keypair.publicKey.toBase58())

// const balance = await connection.getBalance(keypair.publicKey)
// console.log('account balance in SOL:', balance / LAMPORTS_PER_SOL)

// const wallet = new Wallet(keypair)
// const provider = new AnchorProvider(connection, wallet, {
//   commitment: 'confirmed',
// })

// const tokenAccount = await getOrCreateAssociatedTokenAccount(
//   connection,
//   keypair,
//   usdcMintAccount,
//   keypair.publicKey,
// )
// console.log('token account address', tokenAccount.address.toBase58())

// const message = await createOpenMessage({
//   provider,
//   order: {
//     owner: keypair.publicKey,
//     depositAmount: new BN(1000),
//     destChainId: new BN(1),
//     call: {
//       target: new Array(20).fill(0),
//       selector: new Array(4).fill(0),
//       value: new BN(0),
//       params: Buffer.from([]),
//     },
//     expense: {
//       spender: new Array(20).fill(0),
//       token: new Array(20).fill(0),
//       amount: new BN(0),
//     },
//   },
//   ownerTokenAccount: tokenAccount.address,
//   mintAccount: usdcMintAccount,
// })

// const transaction = new VersionedTransaction(message)
// transaction.sign([keypair])

// const txId = await connection.sendTransaction(transaction)
// console.log('open transaction txId', txId)

import { inboxClient } from '../dist/esm/index.js'

console.log(inboxClient)
