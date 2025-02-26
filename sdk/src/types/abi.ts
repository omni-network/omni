import type { Abi } from 'viem'

export type AbiWriteMutability = 'nonpayable' | 'payable'

export type OptionalAbi = Abi | undefined

// all unknown to let abi narrow type before reporting errors
export type OptionalAbis = readonly OptionalAbi[] | readonly unknown[]
