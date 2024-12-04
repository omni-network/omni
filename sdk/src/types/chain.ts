import type { NativeCurrency } from "./nativeCurrency.js"

export type Chain = {
  name: string
  id: number
  testnet: boolean
  portalContract: string
  nativeCurrency: NativeCurrency
}
