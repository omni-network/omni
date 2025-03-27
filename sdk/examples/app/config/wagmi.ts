import { createConfig, http } from 'wagmi';
import { MAINNET_CHAINS, TESTNET_CHAINS } from './chains';
import { mainnet } from 'wagmi/chains';

// Create a Wagmi config with all supported chains
const config = createConfig({
    chains: [
        mainnet, // Ensure at least one chain is present
        // Mainnet chains
        ...Object.values(MAINNET_CHAINS).map(info => info.chain),
        // Testnet chains
        ...Object.values(TESTNET_CHAINS).map(info => info.chain),
    ],
    transports: {
        // Mainnet chains
        ...Object.entries(MAINNET_CHAINS).reduce((acc, [id, info]) => ({
            ...acc,
            [parseInt(id)]: http()
        }), {}),
        // Testnet chains
        ...Object.entries(TESTNET_CHAINS).reduce((acc, [id, info]) => ({
            ...acc,
            [parseInt(id)]: http()
        }), {})
    },
});

export default config;
