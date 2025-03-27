import {
    arbitrum,
    arbitrumSepolia,
    base,
    baseSepolia,
    holesky,
    mainnet,
    optimism,
    optimismSepolia,
    Chain
} from 'wagmi/chains';

export type Network = 'mainnet' | 'omega' | 'staging';

export interface ChainInfo {
    name: string;
    chain: Chain;
}

export const MAINNET_CHAINS: Record<number, ChainInfo> = {
    1: { name: 'Ethereum', chain: mainnet },
    10: { name: 'Optimism', chain: optimism },
    42161: { name: 'Arbitrum One', chain: arbitrum },
    8453: { name: 'Base', chain: base }
};

export const TESTNET_CHAINS: Record<number, ChainInfo> = {
    17000: { name: 'Ethereum Holesky', chain: holesky },
    84532: { name: 'Base Sepolia', chain: baseSepolia },
    421614: { name: 'Arbitrum Sepolia', chain: arbitrumSepolia },
    11155420: { name: 'Optimism Sepolia', chain: optimismSepolia }
};

export const MAINNET_CHAIN_IDS = Object.keys(MAINNET_CHAINS).map(id => parseInt(id));
export const TESTNET_CHAIN_IDS = Object.keys(TESTNET_CHAINS).map(id => parseInt(id));

export const getChainName = (chains: Record<number, ChainInfo>, chainId: number): string => {
    return chains[chainId]?.name || `Chain ${chainId}`;
};

export const getDefaultChains = (network: Network): { source: number; destination: number } => {
    if (network === 'mainnet') {
        return { source: 1, destination: 10 };
    }
    return { source: 17000, destination: 421614 };
};
