import React, { useState, useEffect, useCallback } from 'react';
import { parseEther, formatEther, zeroAddress } from 'viem';
import { useAccount, useBalance, useChainId, useSwitchChain } from 'wagmi';
// SDK INTEGRATION POINT #1 - Import the SDK components
import {
    OmniProvider,
    useQuote,
    useOrder,
    useGetOrderStatus
} from '../../src/index';
import ConnectButton from '../components/ConnectButton';
import { Network, MAINNET_CHAINS, TESTNET_CHAINS, MAINNET_CHAIN_IDS, TESTNET_CHAIN_IDS, getChainName, getDefaultChains } from './config/chains';

interface EthBridgeProps {
    selectedNetwork: Network;
    onNetworkChange: (network: Network) => void;
}

/**
 * ETH Bridge Component
 * Demonstrates Omni SDK integration for cross-chain transfers
 */
const EthBridge: React.FC<EthBridgeProps> = ({ selectedNetwork, onNetworkChange }) => {
    // Wallet connection from wagmi
    const { address, isConnected } = useAccount();
    const chainId = useChainId();
    const { switchChain } = useSwitchChain();

    // User input state
    const [amount, setAmount] = useState<string>('0.01');
    const [sourceChainId, setSourceChainId] = useState<number>(17000);
    const [destinationChainId, setDestinationChainId] = useState<number>(421614);
    const [quoteMode, setQuoteMode] = useState<'deposit' | 'expense'>('deposit');

    // Get available chains based on selected network
    const availableChains = selectedNetwork === 'mainnet' ? MAINNET_CHAINS : TESTNET_CHAINS;
    const availableChainIds = selectedNetwork === 'mainnet' ? MAINNET_CHAIN_IDS : TESTNET_CHAIN_IDS;

    // Update chain selections when network changes
    useEffect(() => {
        const { source, destination } = getDefaultChains(selectedNetwork);
        setSourceChainId(source);
        setDestinationChainId(destination);
    }, [selectedNetwork]);

    // Ensure selected chains are valid for current network
    useEffect(() => {
        const isSourceValid = availableChainIds.includes(sourceChainId);
        const isDestValid = availableChainIds.includes(destinationChainId);

        if (!isSourceValid) {
            setSourceChainId(availableChainIds[0]);
        }
        if (!isDestValid) {
            setDestinationChainId(availableChainIds[1]);
        }
    }, [availableChainIds, sourceChainId, destinationChainId]);

    // Transaction state
    const [isSubmitting, setIsSubmitting] = useState<boolean>(false);
    const [orderId, setOrderId] = useState<`0x${string}` | undefined>(undefined);
    const [errorMessage, setErrorMessage] = useState<string | null>(null);

    // User balance on source chain
    const { data: balance } = useBalance({
        address,
        chainId: sourceChainId,
    });

    // Clear error message on any input change
    const clearError = useCallback(() => {
        if (errorMessage) {
            setErrorMessage(null);
        }
    }, [errorMessage]);

    // Handle user interactions with automatic error clearing
    const handleAmountChange = (e: React.ChangeEvent<HTMLInputElement>) => {
        setAmount(e.target.value);
        clearError();
    };

    const handleSourceChainChange = (e: React.ChangeEvent<HTMLSelectElement>) => {
        const newSourceChainId = parseInt(e.target.value);
        if (newSourceChainId === destinationChainId) {
            setDestinationChainId(sourceChainId);
        }
        setSourceChainId(newSourceChainId);
        clearError();
    };

    const handleDestinationChainChange = (e: React.ChangeEvent<HTMLSelectElement>) => {
        const newDestChainId = parseInt(e.target.value);
        if (sourceChainId === newDestChainId) {
            setSourceChainId(destinationChainId);
        }
        setDestinationChainId(newDestChainId);
        clearError();
    };

    const handleModeChange = (newMode: 'deposit' | 'expense') => {
        setQuoteMode(newMode);
        clearError();
    };

    // Auto-switch to the correct network when source chain changes
    useEffect(() => {
        if (isConnected && chainId !== sourceChainId) {
            switchChain({ chainId: sourceChainId });
        }
    }, [sourceChainId, isConnected, chainId, switchChain]);

    // SDK INTEGRATION POINT #2 - Get cross-chain quote
    // Quote hook determines the deposit/expense amounts based on mode
    const quoteResult = useQuote({
        srcChainId: sourceChainId,
        destChainId: destinationChainId,
        mode: quoteMode,
        deposit: {
            isNative: true,
            amount: quoteMode === 'expense' ? parseEther(amount) : undefined // Only include amount when quoting for expense
        },
        expense: {
            isNative: true,
            amount: quoteMode === 'deposit' ? parseEther(amount) : undefined // Only include amount when quoting for deposit
        },
        enabled: true
    });

    // SDK INTEGRATION POINT #3 - Prepare order parameters
    // Order hook validates and prepares the transaction for execution
    const orderResult = useOrder({
        srcChainId: sourceChainId,
        destChainId: destinationChainId,
        owner: address,
        deposit: {
            token: zeroAddress,
            amount: quoteResult.isSuccess ? quoteResult.deposit.amount : 0n
        },
        calls: [
            {
                target: address as `0x${string}`,
                value: quoteResult.isSuccess
                    ? (quoteMode === 'deposit'
                        ? parseEther(amount)
                        : quoteResult.expense.amount)
                    : 0n,
            }
        ],
        expense: {
            token: zeroAddress,
            amount: quoteResult.isSuccess ? quoteResult.expense.amount : 0n
        },
        validateEnabled: true
    });

    // SDK INTEGRATION POINT #4 - Monitor order status
    // Status hook tracks the cross-chain transaction progress
    const orderStatus = useGetOrderStatus({
        destChainId: destinationChainId,
        orderId,
        srcChainId: sourceChainId
    });

    // Handle bridge transaction
    const handleBridge = async () => {
        if (!address || !quoteResult.isSuccess) return;

        setIsSubmitting(true);
        setErrorMessage(null);

        try {
            // Ensure connected to source chain
            if (chainId !== sourceChainId) {
                await switchChain({ chainId: sourceChainId });
            }

            // SDK INTEGRATION POINT #5 - Execute the bridge transaction
            // Submit the prepared order to initiate the cross-chain transfer
            const txHash = await orderResult.open();
            console.log('Bridge transaction hash:', txHash);

            if (orderResult.orderId) {
                setOrderId(orderResult.orderId);
            }
        } catch (error) {
            console.error('Bridge error:', error);
            setErrorMessage(error instanceof Error ? error.message : 'Failed to bridge ETH');
        } finally {
            setIsSubmitting(false);
        }
    };

    // Compute derived states
    const isWrongChain = isConnected && chainId !== sourceChainId;
    const insufficientBalance = balance && quoteResult.isSuccess &&
        balance.value < quoteResult.deposit.amount;
    const validAmount = parseFloat(amount || '0') > 0;

    return (
        <div className="eth-bridge-container">
            <div className="network-selector">
                <label htmlFor="network">Network:</label>
                <select
                    id="network"
                    value={selectedNetwork}
                    onChange={(e) => onNetworkChange(e.target.value as Network)}
                >
                    <option value="mainnet">Mainnet</option>
                    <option value="omega">Omega</option>
                    <option value="staging">Staging</option>
                </select>
            </div>

            <h2>Cross-Chain ETH Bridge</h2>

            <div className="connect-wallet">
                <ConnectButton />
            </div>

            {!isConnected ? (
                <div className="message">
                    Connect your wallet to continue
                </div>
            ) : (
                <>
                    {/* Chain Selection */}
                    <div className="chain-selection">
                        <div className="source-chain">
                            <label htmlFor="sourceChain">From:</label>
                            <select
                                id="sourceChain"
                                value={sourceChainId}
                                onChange={handleSourceChainChange}
                            >
                                {availableChainIds.map(id => (
                                    <option key={`src-${id}`} value={id}>
                                        {getChainName(availableChains, id)}
                                    </option>
                                ))}
                            </select>
                        </div>

                        <div className="destination-chain">
                            <label htmlFor="destinationChain">To:</label>
                            <select
                                id="destinationChain"
                                value={destinationChainId}
                                onChange={handleDestinationChainChange}
                            >
                                {availableChainIds.map(id => (
                                    <option key={`dest-${id}`} value={id} disabled={id === sourceChainId}>
                                        {getChainName(availableChains, id)}
                                    </option>
                                ))}
                            </select>
                        </div>
                    </div>

                    {/* Balance info */}
                    <div className="balance-info">
                        <p>Balance: {balance ? formatEther(balance.value) : '0'} ETH on {getChainName(availableChains, sourceChainId)}</p>
                    </div>

                    {/* Quote Mode Selection */}
                    <div className="quote-mode-selection">
                        <div className="radio-group">
                            <label className="mode-label">
                                <input
                                    type="radio"
                                    name="quoteMode"
                                    value="expense"
                                    checked={quoteMode === 'expense'}
                                    onChange={() => handleModeChange('expense')}
                                />
                                Send Exact
                            </label>
                            <label className="mode-label">
                                <input
                                    type="radio"
                                    name="quoteMode"
                                    value="deposit"
                                    checked={quoteMode === 'deposit'}
                                    onChange={() => handleModeChange('deposit')}
                                />
                                Receive Exact
                            </label>
                        </div>
                    </div>

                    {/* Amount Input */}
                    <div className="amount-input">
                        <label htmlFor="amount">
                            {quoteMode === 'expense' ? 'You\'ll send exactly (ETH):' : 'You\'ll receive exactly (ETH):'}
                        </label>
                        <input
                            id="amount"
                            type="number"
                            value={amount}
                            onChange={handleAmountChange}
                            min="0.001"
                            step="0.001"
                        />
                    </div>

                    {/* Quote Information */}
                    {validAmount && quoteResult.isSuccess && (
                        <div className="quote-summary">
                            <h3>Quote Summary</h3>
                            {quoteMode === 'expense' ? (
                                <>
                                    <p>You send: {formatEther(quoteResult.deposit.amount)} ETH from {getChainName(availableChains, sourceChainId)}</p>
                                    <p>You receive: {formatEther(quoteResult.expense.amount)} ETH on {getChainName(availableChains, destinationChainId)}</p>
                                    <p className="fee-info">
                                        Fee: {formatEther(quoteResult.deposit.amount - quoteResult.expense.amount)} ETH
                                        {' (deducted from what you receive)'}
                                    </p>
                                </>
                            ) : (
                                <>
                                    <p>You send: {formatEther(quoteResult.deposit.amount)} ETH from {getChainName(availableChains, sourceChainId)}</p>
                                    <p>You receive: {formatEther(quoteResult.expense.amount)} ETH on {getChainName(availableChains, destinationChainId)}</p>
                                    <p className="fee-info">
                                        Fee: {formatEther(quoteResult.deposit.amount - quoteResult.expense.amount)} ETH
                                        {' (added to what you send)'}
                                    </p>
                                </>
                            )}
                        </div>
                    )}

                    {/* Bridge Button */}
                    <button
                        className="bridge-button"
                        onClick={handleBridge}
                        disabled={
                            !isConnected ||
                            isWrongChain ||
                            insufficientBalance ||
                            !quoteResult.isSuccess ||
                            isSubmitting
                        }
                    >
                        {isSubmitting ? 'Processing...' : 'Bridge ETH'}
                    </button>

                    {/* Error Messages */}
                    {isWrongChain && (
                        <div className="error-message">
                            Please switch to {getChainName(availableChains, sourceChainId)}
                        </div>
                    )}

                    {insufficientBalance && (
                        <div className="error-message">
                            Insufficient balance. You need {formatEther(quoteResult.deposit.amount)} ETH.
                        </div>
                    )}

                    {errorMessage && (
                        <div className="error-message">
                            {errorMessage}
                        </div>
                    )}

                    {/* Transaction Status */}
                    {orderId && (
                        <div className="order-status">
                            <h3>Order Status</h3>
                            <p>Order ID: {orderId}</p>
                            <p>Status: {orderStatus.status || 'Processing...'}</p>

                            {orderStatus.status === 'filled' && (
                                <div className="success-message">
                                    Transfer complete! Your ETH has been bridged to {getChainName(availableChains, destinationChainId)}.
                                </div>
                            )}
                        </div>
                    )}
                </>
            )}
        </div>
    );
};

/**
 * Main App Component with OmniProvider wrapper
 * SDK INTEGRATION POINT #6 - Configure provider with environment and API endpoint
 */
const App: React.FC = () => {
    const [selectedNetwork, setSelectedNetwork] = useState<Network>('staging');

    const getApiUrl = (network: Network) => {
        return `https://solver.${network}.omni.network/api/v1`;
    };

    return (
        <OmniProvider
            env="testnet"
            __apiBaseUrl={getApiUrl(selectedNetwork)}
        >
            <div className="app-container">
                <h1>Omni SolverNet SDK Integration Demo</h1>
                <p>Demonstrating cross-chain ETH bridging with the Omni SolverNet SDK</p>
                <EthBridge selectedNetwork={selectedNetwork} onNetworkChange={setSelectedNetwork} />
                <div className="demo-notes">
                    <p>Note: This demo uses testnet and mainnet networks. You'll need ETH on your source network to try it.</p>
                </div>
            </div>
        </OmniProvider>
    );
};

export default App;
