import type { Hex } from 'viem'

export type Rejection = {
  txHash: Hex
  rejectReason: string
}
