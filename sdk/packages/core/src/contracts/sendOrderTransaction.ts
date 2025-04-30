import type { Address, Client } from 'viem'
import { zeroAddress } from 'viem'
import { type WriteContractReturnType, writeContract } from 'viem/actions'
import { inboxABI } from '../constants/abis.js'
import { typeHash } from '../constants/typehash.js'
import { AccountRequiredError } from '../errors/base.js'
import type { OptionalAbis } from '../types/abi.js'
import type { Order } from '../types/order.js'
import { encodeOrderData } from '../utils/encodeOrderData.js'

const defaultFillDeadline = () => Math.floor(Date.now() / 1000 + 86400)

export type SendOrderTransactionParameters<abis extends OptionalAbis> = {
  client: Client
  inboxAddress: Address
  order: Order<abis>
}

export type SendOrderTransactionReturn = WriteContractReturnType

export async function sendOrderTransaction<abis extends OptionalAbis>(
  params: SendOrderTransactionParameters<abis>,
): Promise<SendOrderTransactionReturn> {
  const { client, inboxAddress, order } = params

  if (client.account == null) {
    throw new AccountRequiredError(
      'Client needs to have an associated account to open an order',
    )
  }

  const isNativeDeposit =
    order.deposit.token == null || order.deposit.token === zeroAddress

  return await writeContract(client, {
    abi: inboxABI,
    address: inboxAddress,
    functionName: 'open',
    account: client.account,
    chain: client.chain,
    value: isNativeDeposit ? order.deposit.amount : 0n,
    args: [
      {
        fillDeadline: order.fillDeadline ?? defaultFillDeadline(),
        orderDataType: typeHash,
        orderData: encodeOrderData(order),
      },
    ],
  })
}
