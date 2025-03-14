import React from 'react';
import { useAccount, useConnect, useDisconnect } from 'wagmi';
import { injected } from 'wagmi/connectors';

const ConnectButton: React.FC = () => {
    const { address, isConnected } = useAccount();
    const { connect } = useConnect();
    const { disconnect } = useDisconnect();

    if (isConnected && address) {
        return (
            <div className="connect-button-container">
                <div className="wallet-info">
                    Connected: {formatAddress(address)}
                    <button
                        onClick={() => disconnect()}
                        style={{ backgroundColor: '#e74c3c', marginLeft: '10px' }}
                    >
                        Disconnect
                    </button>
                </div>
            </div>
        );
    }

    return (
        <button
            onClick={() => connect({ connector: injected() })}
            style={{ marginBottom: '1rem' }}
        >
            Connect Wallet
        </button>
    );
};

// Helper to format the address (truncate the middle)
const formatAddress = (address: string): string => {
    return `${address.substring(0, 6)}...${address.substring(address.length - 4)}`;
};

export default ConnectButton;
