import type {
  Account,
  Chain,
  HttpTransport,
  PublicActions,
  TestClient,
  WalletClient,
} from 'viem'
import {
  http,
  createTestClient,
  createWalletClient,
  parseEther,
  publicActions,
} from 'viem'
import { testAccount } from './accounts.js'
import { inbox, mockL1Chain, omniTokenAbi, tokenAddress } from './constants.js'

export function createAnvilClient(chain: Chain): TestClient {
  return createTestClient({ chain, mode: 'anvil', transport: http() })
}

export type Client<ChainType extends Chain> = WalletClient<
  HttpTransport,
  ChainType
> &
  PublicActions<HttpTransport, ChainType, Account>

export function createClient<ChainType extends Chain>({
  account,
  chain,
}: { account?: Account; chain: ChainType }): Client<ChainType> {
  return createWalletClient({
    account: account ?? testAccount,
    chain,
    transport: http(),
  }).extend(publicActions)
}

export const mockL1Client: Client<typeof mockL1Chain> = createClient({
  chain: mockL1Chain,
})

export async function mintOMNI(client: Client<Chain>): Promise<void> {
  const account = client.account?.address
  if (account == null) {
    throw new Error('Missing account on client')
  }
  const amount = parseEther('100')
  const mintHash = await client.writeContract({
    account,
    address: tokenAddress,
    abi: omniTokenAbi,
    functionName: 'mint',
    args: [account, amount],
  })

  await client.waitForTransactionReceipt({
    hash: mintHash,
    pollingInterval: 500,
  })
  const approveHash = await client.writeContract({
    account,
    address: tokenAddress,
    abi: omniTokenAbi,
    functionName: 'approve',
    args: [inbox, amount],
  })
  await client.waitForTransactionReceipt({
    hash: approveHash,
    pollingInterval: 500,
  })
}
