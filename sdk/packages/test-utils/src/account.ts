import {
  http,
  type Account,
  type Chain,
  type HttpTransport,
  type PublicActions,
  type WalletClient,
  createWalletClient,
  parseEther,
  publicActions,
} from 'viem'
import { type PrivateKeyAccount, privateKeyToAccount } from 'viem/accounts'
import { inbox, mockL1Chain, omniTokenAbi, tokenAddress } from './constants.js'

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

export const mockL1Client: Client<typeof mockL1Chain> = createClient({
  chain: mockL1Chain,
})

export async function mintOMNI(): Promise<void> {
  const amount = parseEther('100')
  const mintHash = await mockL1Client.writeContract({
    account: testAccount.address,
    address: tokenAddress,
    abi: omniTokenAbi,
    functionName: 'mint',
    args: [testAccount.address, amount],
  })
  await mockL1Client.waitForTransactionReceipt({ hash: mintHash })
  const approveHash = await mockL1Client.writeContract({
    account: testAccount.address,
    address: tokenAddress,
    abi: omniTokenAbi,
    functionName: 'approve',
    args: [inbox, amount],
  })
  await mockL1Client.waitForTransactionReceipt({ hash: approveHash })
}
