import type { Environment } from '../types/config.js'

export const apiUrls: Record<Environment, string> = {
  devnet: 'http://localhost:26661/api/v1',
  testnet: 'https://solver.omega.omni.network/api/v1',
  mainnet: 'https://solver.mainnet.omni.network/api/v1',
}
