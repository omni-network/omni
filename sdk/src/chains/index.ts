import type { Chain } from "../types/chain.js"

/**
 * @summary We will want to eventually define chains in submodules, and re-export from here,
 * to allow for lazy loading for consumers. But this is fine for now while we work
 * with a small number of chains.
 */

export const mainnet: Chain = {
  name: "Ethereum Mainnet",
  id: 1,
  testnet: false,
  portalContract: "0x0000000000000000000000000000000000000000",
  nativeCurrency: { name: "Ether", symbol: "ETH", decimals: 18 },
}

////////////////////////////////////////////////////////////
/// TESTNETS
////////////////////////////////////////////////////////////
export const omniOmega: Chain = {
  name: "Omni Omega",
  id: 164,
  testnet: true,
  portalContract: "0xcB60A0451831E4865bC49f41F9C67665Fc9b75C3",
  nativeCurrency: { name: "Omni", symbol: "OMNI", decimals: 18 },
}

export const holesky: Chain = {
  name: "Ethereum Holesky",
  id: 17000,
  testnet: true,
  portalContract: "0xcB60A0451831E4865bC49f41F9C67665Fc9b75C3",
  nativeCurrency: { name: "Ether", symbol: "ETH", decimals: 18 },
}

export const baseSepolia: Chain = {
  name: "Base Sepolia",
  id: 84532,
  testnet: true,
  portalContract: "0xcB60A0451831E4865bC49f41F9C67665Fc9b75C3",
  nativeCurrency: { name: "Ether", symbol: "ETH", decimals: 18 },
}

export const arbitrumSepolia: Chain = {
  name: "Arbitrum Sepolia",
  id: 421614,
  testnet: true,
  portalContract: "0xcB60A0451831E4865bC49f41F9C67665Fc9b75C3",
  nativeCurrency: { name: "Ether", symbol: "ETH", decimals: 18 },
}

export const optimismSepolia: Chain = {
  name: "Optimism Sepolia",
  id: 11155420,
  testnet: true,
  portalContract: "0xcB60A0451831E4865bC49f41F9C67665Fc9b75C3",
  nativeCurrency: { name: "Ether", symbol: "ETH", decimals: 18 },
}
