//
// Necessary dependencies from the Omni SDK and Viem
//

import {
  type Environment,
  type OrderState,
  generateOrder,
  getContracts,
  getQuote,
} from '@omni-network/core'
import {
  http,
  type Hex,
  createWalletClient,
  formatEther,
  parseEther,
  publicActions,
} from 'viem'
import { privateKeyToAccount } from 'viem/accounts'
import { getBalance } from 'viem/actions'
import { baseSepolia, holesky } from 'viem/chains'

//
// Wallet and client setup
//

// Create an account using the private key provided as environment variable
const privateKey = process.env.WALLET_PRIVATE_KEY || ''
if (privateKey === '') {
  throw new Error('Missing WALLET_PRIVATE_KEY environment variable')
}
if (!privateKey.startsWith('0x')) {
  throw new Error(
    'WALLET_PRIVATE_KEY environment variable must be an hexadecimal string starting with 0x',
  )
}
const account = privateKeyToAccount(privateKey as Hex)
const client = createWalletClient({
  account,
  chain: baseSepolia,
  transport: http(),
}).extend(publicActions)

//
// Configuration
//

const environment: Environment = 'testnet'
const depositAmount = parseEther('0.01')

//
// Prerequisite: check account balance is high enough for wanted deposit
//

console.log(`⬇️  Loading balance for account: ${account.address}`)
const balance = await getBalance(client, { address: account.address })
if (depositAmount > balance) {
  console.log(`❌ Insufficient account balance: ${formatEther(balance)}`)
  process.exit(1)
}
console.log(`✅ Account balance: ${formatEther(balance)}`)

//
// Prerequisite: load SolverNet contracts addresses
//

console.log(
  `⬇️  Loading SolverNet contracts addresses for environment: ${environment}`,
)
const contracts = await getContracts({ environment })
console.log('✅ SolverNet contracts addresses loaded')

//
// Get a quote for the order want to make
//

console.log(
  `⬇️  Loading quote for Ether deposit amount: ${formatEther(depositAmount)}`,
)
const quote = await getQuote({
  srcChainId: baseSepolia.id,
  destChainId: holesky.id,
  deposit: { amount: depositAmount, isNative: true },
  expense: { isNative: true },
  mode: 'expense',
  environment,
})
console.log(
  `✅ Loaded quote with Ether expense amount: ${formatEther(quote.expense.amount)}`,
)

//
// Order flow
//

console.log('0️⃣  Generating order')
const generator = generateOrder({
  client,
  environment,
  inboxAddress: contracts.inbox,
  order: {
    srcChainId: baseSepolia.id,
    destChainId: holesky.id,
    // Use values provided by the quote
    deposit: quote.deposit,
    expense: quote.expense,
    calls: [{ target: account.address, value: quote.expense.amount }],
  },
})

let orderState: OrderState | undefined
for await (orderState of generator) {
  switch (orderState.status) {
    case 'valid':
      console.log('1️⃣  Order validated using SolverNet API')
      break
    case 'sent':
      console.log(`2️⃣  Order sent with transaction hash: ${orderState.txHash}`)
      break
    case 'open':
      console.log(`3️⃣  Order open with identifier: ${orderState.order.orderId}`)
      break
    default:
      console.log(`4️⃣  Order terminated with status: ${orderState.status}`)
  }
}
if (orderState == null) {
  console.log('❌ Order generation failed')
  process.exit(2)
}
if (orderState.status === 'filled') {
  console.log('✅ Order successfully filled')
} else {
  console.log(`❌ Order failed with status: ${orderState.status}`)
  process.exit(3)
}
