import type { Address as SVMAddress } from '@solana/addresses'
import type { Address as EVMAddress } from 'viem'

export type { EVMAddress, SVMAddress }

export type Address = EVMAddress | SVMAddress
