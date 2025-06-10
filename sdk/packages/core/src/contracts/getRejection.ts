import type { BlockNumber, Client, Hex } from 'viem'
import { getLogs } from 'viem/actions'
import { inboxABI } from '../constants/abis.js'
import { GetRejectionError } from '../errors/base.js'
import type { EVMAddress } from '../types/addresses.js'
import type { Rejection } from '../types/rejection.js'
import { invariant } from '../utils/index.js'

// these rejections are defined in the solver
// https://github.com/omni-network/omni/blob/main/solver/types/reject.go
export const rejectReasons = {
  0: 'None',
  1: 'Destination call reverts',
  2: 'Invalid deposit',
  3: 'Invalid expense',
  4: 'Insufficient deposit',
  5: 'Solver has insufficient inventory',
  6: 'Unsupported deposit',
  7: 'Unsupported expense',
  8: 'Unsupported destination chain',
  9: 'Unsupported source chain',
  11: 'Expense over max',
  12: 'Expense under min',
  13: 'Call not allowed',
}

export type GetRejectionParams = {
  client: Client
  orderId: Hex
  inboxAddress: EVMAddress
  fromBlock: BlockNumber
}

export async function getRejection({
  client,
  orderId,
  inboxAddress,
  fromBlock,
}: GetRejectionParams): Promise<Rejection> {
  const logs = await getLogs(client, {
    address: inboxAddress,
    fromBlock,
    toBlock: 'latest',
    event: getRejectEventAbi(),
    args: {
      id: orderId,
    },
    // only returns logs that conform to the args, aka event args will always be present
    strict: true,
  })

  if (logs.length !== 1) {
    throw new GetRejectionError(
      `Expected exactly one 'Rejected' event but found ${logs.length}.`,
    )
  }

  invariant(
    !!logs[0].transactionHash,
    'Tx hash is always present for "latest" blocks',
  )

  return {
    // it's safe to assume we expect a single Rejected event per orderId
    txHash: logs[0].transactionHash,
    rejectReason: parseRejectReason(logs[0].args.reason),
  }
}

const getRejectEventAbi = () => {
  const abi = inboxABI.find(
    (item) => item.type === 'event' && item.name === 'Rejected',
  )
  invariant(!!abi, 'Rejected event not found')
  return abi
}

const parseRejectReason = (reason: number): string => {
  invariant(reason in rejectReasons, 'Invalid reject reason key')
  return rejectReasons[reason as keyof typeof rejectReasons]
}
