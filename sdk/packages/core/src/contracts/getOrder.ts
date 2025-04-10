import type { Address, Chain, Client, Hex, Transport } from 'viem'
import { type ReadContractReturnType, readContract } from 'viem/actions'
import { inboxABI } from '../constants/abis.js'

export type GetOrderParameters<chain extends Chain> = {
  client: Client<Transport, chain>
  inboxAddress: Address
  orderId: Hex
}

export type GetOrderReturn = ReadContractReturnType<
  typeof inboxABI,
  'getOrder',
  [Hex]
>

export async function getOrder<chain extends Chain>({
  client,
  inboxAddress,
  orderId,
}: GetOrderParameters<chain>): Promise<GetOrderReturn> {
  return readContract(client, {
    address: inboxAddress,
    abi: inboxABI,
    functionName: 'getOrder',
    args: [orderId],
  })
}
