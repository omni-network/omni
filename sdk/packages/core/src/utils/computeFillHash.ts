import { decodeAbiParameters, encodeAbiParameters, keccak256 } from 'viem'
import { fillOriginDataAbi } from '../constants/abis.js'
import type { ResolvedOrder } from '../contracts/parseOpenEvent.js'

export function computeFillHash(resolvedOrder: ResolvedOrder) {
  const [decoded] = decodeAbiParameters(
    [fillOriginDataAbi],
    resolvedOrder.fillInstructions[0].originData,
  )

  return keccak256(
    encodeAbiParameters(
      [{ type: 'bytes32', name: 'srcReqId' }, fillOriginDataAbi],
      [resolvedOrder.orderId, decoded],
    ),
  )
}
