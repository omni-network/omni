import React, { useState, useEffect, useCallback } from 'react';
import { parseEther, formatEther, zeroAddress } from 'viem';
import { useAccount, useBalance, useChainId, useSwitchChain } from 'wagmi';
import {
    OmniProvider,
    useQuote,
    useOrder,
    useGetOrderStatus
} from '../../src/index';
import ConnectButton from '../components/ConnectButton';

// Network configuration
const MAINNET_CHAINS: Record<number, { name: string; }> = {
    1: { name: 'Ethereum' },
    10: { name: 'Optimism' },
    42161: { name: 'Arbitrum One' },
    8453: { name: 'Base' }
};

const TESTNET_CHAINS: Record<number, { name: string; }> = {
    17000: { name: 'Ethereum Holesky' },
    84532: { name: 'Base Sepolia' },
    421614: { name: 'Arbitrum Sepolia' },
    11155420: { name: 'Optimism Sepolia' }
};

const MAINNET_CHAIN_IDS = Object.keys(MAINNET_CHAINS).map(id => parseInt(id));
const TESTNET_CHAIN_IDS = Object.keys(TESTNET_CHAINS).map(id => parseInt(id));

// Helper function to get chain name safely
const getChainName = (chains: Record<number, { name: string }>, chainId: number): string => {
    return chains[chainId]?.name || `Chain ${chainId}`;
};

/**
 * ETH Bridge Component
 * Demonstrates Omni SDK integration for cross-chain transfers
 */
const EthBridge: React.FC = () => {
    // Wallet connection from wagmi
    const { address, isConnected } = useAccount();
    const chainId = useChainId();
    const { switchChain } = useSwitchChain();

    // User input state
    const [amount, setAmount] = useState<string>('0.01');
    const [sourceChainId, setSourceChainId] = useState<number>(17000);
    const [destinationChainId, setDestinationChainId] = useState<number>(421614);
    const [quoteMode, setQuoteMode] = useState<'deposit' | 'expense'>('deposit');
    const [selectedNetwork, setSelectedNetwork] = useState<'mainnet' | 'omega' | 'staging'>('staging');

    // Get available chains based on selected network
    const availableChains = selectedNetwork === 'mainnet' ? MAINNET_CHAINS : TESTNET_CHAINS;
    const availableChainIds = selectedNetwork === 'mainnet' ? MAINNET_CHAIN_IDS : TESTNET_CHAIN_IDS;

    // Update chain selections when network changes
    useEffect(() => {
        const defaultSourceId = selectedNetwork === 'mainnet' ? 1 : 17000;
        const defaultDestId = selectedNetwork === 'mainnet' ? 10 : 421614;

        setSourceChainId(defaultSourceId);
        setDestinationChainId(defaultDestId);
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
    const quoteResult = useQuote({
        // The SDK uses inverted mode terminology compared to our UI:
        // UI "Send Exact" (deposit) -> SDK uses "expense" mode to calculate what user receives
        // UI "Receive Exact" (expense) -> SDK uses "deposit" mode to calculate what user sends
        srcChainId: sourceChainId,
        destChainId: destinationChainId,
        mode: quoteMode === 'deposit' ? 'expense' : 'deposit',
        deposit: {
            isNative: true,
            amount: parseEther(amount || '0')
        },
        expense: {
            isNative: true,
            amount: parseEther(amount || '0')
        },
        enabled: !!address && parseFloat(amount || '0') > 0
    });

    // SDK INTEGRATION POINT #3 - Setup order parameters
    const orderResult = useOrder({
        srcChainId: sourceChainId,
        destChainId: destinationChainId,
        owner: address,
        deposit: {
            token: zeroAddress, // For native ETH
            amount: quoteResult.isSuccess ? quoteResult.deposit.amount : 0n
        },
        calls: [
            {
                target: (address || '0x') as `0x${string}`, // Send to same address by default
                value: quoteResult.isSuccess
                    ? (quoteMode === 'deposit'
                        ? parseEther(amount)  // "Send Exact" mode: use user's input amount
                        : quoteResult.expense.amount  // "Receive Exact" mode: use quote result
                    )
                    : 0n,
            }
        ],
        expense: {
            token: undefined,
            amount: 0n
        },
        validateEnabled: true
    });

    // SDK INTEGRATION POINT #4 - Track transaction status
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
                    onChange={(e) => setSelectedNetwork(e.target.value as 'mainnet' | 'omega' | 'staging')}
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
                                    value="deposit"
                                    checked={quoteMode === 'deposit'}
                                    onChange={() => handleModeChange('deposit')}
                                />
                                Send Exact
                            </label>
                            <label className="mode-label">
                                <input
                                    type="radio"
                                    name="quoteMode"
                                    value="expense"
                                    checked={quoteMode === 'expense'}
                                    onChange={() => handleModeChange('expense')}
                                />
                                Receive Exact
                            </label>
                        </div>
                    </div>

                    {/* Amount Input */}
                    <div className="amount-input">
                        <label htmlFor="amount">
                            {quoteMode === 'deposit' ? 'You\'ll send exactly (ETH):' : 'You\'ll receive exactly (ETH):'}
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
                            {quoteMode === 'deposit' ? (
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
 * SDK INTEGRATION POINT #6 - Root provider setup
 */
const App: React.FC = () => {
    return (
        <OmniProvider env="testnet">
            <div className="app-container">
                <h1>Omni SDK Integration Demo</h1>
                <p>Demonstrating cross-chain ETH bridging with the Omni SDK</p>
                <EthBridge />
                <div className="demo-notes">
                    <p>Note: This demo uses testnet networks. You'll need testnet ETH to try it.</p>
                </div>
            </div>
        </OmniProvider>
    );
};

export default App;
