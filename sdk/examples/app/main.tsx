import React from 'react'
import ReactDOM from 'react-dom/client'
import { createConfig, http, WagmiProvider } from 'wagmi'
import { QueryClient, QueryClientProvider } from '@tanstack/react-query'
import {
    arbitrumSepolia,
    baseSepolia,
    holesky,
    optimismSepolia
} from 'wagmi/chains'
import App from './bridge-eth'

// Create a Wagmi config with the chains we want to support
const config = createConfig({
    chains: [holesky, arbitrumSepolia, baseSepolia, optimismSepolia],
    transports: {
        [holesky.id]: http(),
        [arbitrumSepolia.id]: http(),
        [baseSepolia.id]: http(),
        [optimismSepolia.id]: http(),
    },
})

// Create a react-query client
const queryClient = new QueryClient()

// Render our app wrapped with the required providers
ReactDOM.createRoot(document.getElementById('root')!).render(
    <React.StrictMode>
        <WagmiProvider config={config}>
            <QueryClientProvider client={queryClient}>
                <App />
            </QueryClientProvider>
        </WagmiProvider>
    </React.StrictMode>,
)
