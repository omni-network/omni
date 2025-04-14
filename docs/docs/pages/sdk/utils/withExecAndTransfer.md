---
sidebar_position: 1
title: withExecAndTransfer
---

# `withExecAndTransfer`

Use the `withExecAndTransfer` utility when interacting with destination contracts whose functions credit assets directly to `msg.sender` instead of allowing a recipient address to be specified.

Many smart contracts operate this way (e.g., some vault deposits, staking rewards, NFT mints). When called via SolverNet, the `msg.sender` on the destination chain is an Omni Network contract, not the end user. `withExecAndTransfer` ensures assets sent to `msg.sender` are automatically forwarded to the intended user address after the primary function call completes.

## How it Works

This utility wraps your intended destination contract call (`call`) within another call targeting the `SolverNetMiddleman` contract. When the solver fulfills the intent:

1.  The `SolverNetMiddleman` contract is called on the destination chain.
2.  The `SolverNetMiddleman` executes the original `call` you specified (e.g., `Vault.deposit()` which credits vault shares to `msg.sender`).
3.  After the original `call` completes, the `SolverNetMiddleman` (which *was* the `msg.sender` and received the assets) checks its own balance of a specified `token`.
4.  It transfers the *entire balance* of that `token` it holds to the specified final recipient address (`to`).

Here's the relevant part of the `SolverNetMiddleman` contract:

```solidity
contract SolverNetMiddleman {
    // ...

    /**
     * @notice Execute a call and transfer any received tokens back to the recipient
     * @dev Intended to be used when interacting with contracts that don't allow us to specify a recipient
     * @param token  Token to transfer (address(0) for ETH)
     * @param to     Recipient address
     * @param target Call target address
     * @param data   Calldata for the call
     */
    function executeAndTransfer(
        address token,
        address to,
        address target,
        bytes calldata data
    ) external payable nonReentrant {
        (bool success, ) = target.call{value: msg.value}(data); // Executes the original call
        if (!success) revert CallFailed();

        // Transfer received assets to the intended recipient
        if (token == address(0)) {
            SafeTransferLib.safeTransferAllETH(to);
        } else {
            IERC20(token).safeTransferAll(to); // safeTransferAll checks balance and transfers if > 0
        }
    }

    // ...
}
```

## Example

### High Level Usage

Before diving into the full React hook example, here's the basic idea:

1.  **Define your original call:** This is the call to the target contract function (e.g., `vault.deposit()`) that sends assets to `msg.sender`.
2.  **Wrap it:** Use `withExecAndTransfer`, providing the Omni `middlemanAddress`, your original call, the `token` you expect the middleman to receive, and the final recipient (`to`).
3.  **Use the wrapped call:** Pass the result of `withExecAndTransfer` (which is a `Call` object itself) into the `calls` array for `useOrder`.

```typescript
import { withExecAndTransfer } from '@omni-network/react';
import { type Call, type Address } from 'viem';

// Assume these are defined elsewhere
declare const middlemanAddress: Address;
declare const vaultAddress: Address;
declare const vaultTokenAddress: Address; // The token the vault sends to msg.sender
declare const userAddress: Address;
declare const vaultABI: any;
declare const depositAmount: bigint;

// 1. Define the original call (e.g., vault.deposit() that sends tokens to msg.sender)
const originalCall: Call = {
    target: vaultAddress,
    abi: vaultABI,
    functionName: 'deposit',
    value: depositAmount, // If the deposit function is payable
};

// 2. Wrap the call
const wrappedCall: Call = withExecAndTransfer({
    middlemanAddress: middlemanAddress,
    call: originalCall,
    transfer: {
        token: vaultTokenAddress, // Token the middleman should forward
        to: userAddress,          // Final recipient
    }
});

// 3. Use the wrapped call in useOrder's calls array
// const order = useOrder({
//    ...
//    calls: [wrappedCall],
//    expense: {
//        ...,
//        spender: middlemanAddress // Middleman needs to spend to execute wrapped call
//    }
// });
```

### React Example

Let's say you have a `TokenizedVault` that mints vault shares (ERC20 tokens) to `msg.sender` upon deposit. This example shows how to configure `useOrder` with `withExecAndTransfer` to handle this, including fetching the necessary middleman address.

```tsx
import {
  useOrder,
  useQuote,
  useOmniContracts,
  withExecAndTransfer
} from '@omni-network/react'
import { parseEther, zeroAddress, type Abi, type Address } from 'viem'
import { baseSepolia, holesky } from 'viem/chains'
import { useAccount } from 'wagmi' // Added useAccount for userAddress
import React, { useMemo } from 'react' // Added React/useMemo for context

// ABI for TokenizedVault.deposit{ value: amount }()
const tokenizedVaultABI = [
  {
    inputs: [],
    name: 'deposit',
    outputs: [],
    stateMutability: 'payable',
    type: 'function',
  },
] as const

// Addresses (replace with actual values)
const vaultAddress = '0x...' as const      // Your tokenized vault address
const vaultTokenAddress = '0x...' as const // The ERC20 token minted by the vault to msg.sender

function DepositTokenizedVault() {
    const { address: userAddress } = useAccount();

    // --- Step 1: Get Quote ---
    // Use `useQuote` in EITHER "deposit" OR "expense" mode
    // to get the required amounts. Example assumes a successful quote.
    const quote = useQuote({ /* ... quote config (mode: 'deposit' or 'expense') ... */ });

    // Get the appropriate amounts from the single successful quote result:
    const depositAmt = quote.data?.deposit.amount ?? 0n; // Amount user deposits on source
    const expenseAmt = quote.data?.expense.amount ?? 0n; // Amount spent on destination (value for vault call)

    // --- Step 2: Get Middleman Address ---
    // Fetch the Middleman contract address for the destination chain
    const destChainId = holesky.id; // Example destination
    const contracts = useOmniContracts(destChainId) // Pass destChainId
    // Need the middleman to wrap the call
    const middlemanAddress = contracts.data?.middleman ?? zeroAddress

    // --- Step 3: Define and Wrap the Call ---
    // 1. Define the original call to the vault (this is the one sending to msg.sender)
    const originalVaultCall = useMemo(() => ({
        target: vaultAddress,
        abi: tokenizedVaultABI,
        functionName: 'deposit',
        value: expenseAmt, // Pass ETH value if required by the deposit function
    }), [expenseAmt]); // Recalculate if expenseAmt changes

    // 2. Wrap the call using withExecAndTransfer
    const middlemanCall = useMemo(() => {
        // Ensure we have the user's address and the middleman address before creating the call
        if (!userAddress || middlemanAddress === zeroAddress) return undefined;

        return withExecAndTransfer({
            middlemanAddress: middlemanAddress, // The Omni contract that will execute the call
            call: originalVaultCall,             // The actual vault interaction
            transfer: {
                // Specify the token the middleman should transfer *after* calling the vault.
                // This is the token the vault sends to msg.sender (the middleman).
                token: vaultTokenAddress,
                // Specify the final recipient of the vault tokens (the user).
                to: userAddress,
            }
        });
        // Recalculate if dependencies change
    }, [userAddress, middlemanAddress, originalVaultCall, vaultTokenAddress]);

    // --- Step 4: Configure useOrder ---
    // 3. Pass the *wrapped* call to useOrder
    const order = useOrder({
        srcChainId: baseSepolia.id, // Example source
        destChainId: destChainId,
        deposit: { isNative: true, amount: depositAmt }, // User deposits this on source
        // Expense: funds the middleman needs to execute the *wrapped* call
        expense: {
            isNative: true, // Assuming vault needs ETH
            amount: expenseAmt,
            // Spender must be the middleman, as it's the target of the
            // effective call from the solver and needs to spend the expenseAmt
            spender: middlemanAddress
        },
        // Pass the call generated by withExecAndTransfer
        // Ensure middlemanCall is defined before including it
        calls: middlemanCall ? [middlemanCall] : [],
        // Enable validation only when quote, contracts, and the wrapped call are ready
        validateEnabled: quote.isSuccess && contracts.isSuccess && middlemanCall != null,
    })

    // ... rest of the component (open button, status display, etc.)
    // const { open, status, validation } = order;
    return <div>{/* UI using order object */}</div>;
}
```

**Key Points:**
*   The `spender` in the `expense` configuration for `useOrder` **must** be the `middlemanAddress` when using `withExecAndTransfer`. The solver sends the `expense` funds to the middleman, which then uses them to execute your `call` (including any `value`).

*   The `token` in the `transfer` configuration of `withExecAndTransfer` is the address of the asset you expect the *middleman* to receive *as a result of* executing your `call`. This is the asset the middleman will forward to the final `to` address. Use `zeroAddress` if the `call` results in native ETH being sent to the middleman.

*   If the target function requires sending ETH (like a payable function), include the `value` field in the original `call` object passed to `withExecAndTransfer`.
