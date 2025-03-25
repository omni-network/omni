import {
  http,
  type Account,
  type Chain,
  type HttpTransport,
  type PublicActions,
  type WalletClient,
  createWalletClient,
  publicActions,
} from 'viem'
import { type PrivateKeyAccount, privateKeyToAccount } from 'viem/accounts'

import {
  ETHER,
  MOCK_L1_CHAIN,
  OMNI_TOKEN_ABI,
  SOLVERNET_INBOX_ADDRESS,
  TOKEN_ADDRESS,
} from './constants.js'

export const testAccount: PrivateKeyAccount = privateKeyToAccount(
  '0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80',
)

export type Client<ChainType extends Chain> = WalletClient<
  HttpTransport,
  ChainType
> &
  PublicActions<HttpTransport, ChainType, Account>

export function createClient<ChainType extends Chain>({
  chain,
}: { chain: ChainType }): Client<ChainType> {
  return createWalletClient({
    account: testAccount,
    chain,
    transport: http(),
  }).extend(publicActions)
}

export const mockL1Client: Client<typeof MOCK_L1_CHAIN> = createClient({
  chain: MOCK_L1_CHAIN,
})

export async function mintOMNI(): Promise<void> {
  // mint token
  const mintHash = await mockL1Client.writeContract({
    account: testAccount.address,
    address: TOKEN_ADDRESS,
    abi: OMNI_TOKEN_ABI,
    functionName: 'mint',
    args: [testAccount.address, 100n * ETHER],
  })
  // wait for transaction to be mined
  await mockL1Client.waitForTransactionReceipt({ hash: mintHash })
  // approve transfers to inbox contract
  const approveHash = await mockL1Client.writeContract({
    account: testAccount.address,
    address: TOKEN_ADDRESS,
    abi: OMNI_TOKEN_ABI,
    functionName: 'approve',
    args: [SOLVERNET_INBOX_ADDRESS, 100n * ETHER],
  })
  // wait for transaction to be mined
  await mockL1Client.waitForTransactionReceipt({ hash: approveHash })
}
