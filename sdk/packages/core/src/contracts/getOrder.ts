import type { Address, Client, Hex } from 'viem'
import { type ReadContractReturnType, readContract } from 'viem/actions'
import { inboxABI } from '../constants/abis.js'

export type GetOrderParameters = {
  client: Client
  inboxAddress: Address
  orderId: Hex
}

export type GetOrderReturn = ReadContractReturnType<
  typeof inboxABI,
  'getOrder',
  [Hex]
>

export async function getOrder({
  client,
  inboxAddress,
  orderId,
}: GetOrderParameters): Promise<GetOrderReturn> {
  return await readContract(client, {
    address: inboxAddress,
    abi: inboxABI,
    functionName: 'getOrder',
    args: [orderId],
  })
}
