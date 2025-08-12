# evmredenom

## Overview

The `evmredenom` module implements a secure token redenomination process for the Omni->Nomina rebrand, multiplying all EVM account balances by a factor of 75 after the `4_earhart` network upgrade. This is a critical one-time operation that must maintain state integrity while handling the conversion from `OMNI` to `NOM` tokens.

## Problem Statement

The `OMNI` token is being migrated to `NOM` as part of the Omni->Nomina rebrand.
This includes a redenomination of the token supply, where 1 `OMNI` will be equivalent to 75 `NOM`.
This implies that EVM account balances must be adjusted accordingly to preserve user balances in the new denomination.

It also implies changes on Ethereum L1 related to the token contract and the token bridge, which are out of scope
for this `evmredenom` module.

## Solution Architecture

Introduced as part of the `4_earhart` network upgrade, this `evmredenom` module
is responsible for increasing the balances of all EVM accounts by a factor of 75.

**Key Design Decisions:**
- Redenominating EVM account balances over a period of time after the `4_earhart` network upgrade is sufficient.
- All EVM smart contract state is assumed to be compatible with balances increases.
- The consensus chain state/balances isn't redenominated to avoid complex migration of staking state
- Instead, the conversion rate from EVM native token to consensus chain staking module bond token is adjusted from 1:1 to 75:1
- This preserves existing consensus layer functionality while achieving the required EVM balance adjustments

The admin-controlled `Redenom.sol` smart contract proxies submitted account data in batches
to the `evmredenom` module, which then creates withdrawals to increase account balances.
The ethereum snapsync protocol's account sync logic is used to secure this process through cryptographic proofs.

## Implementation Details

- During the `4_earhart` network upgrade, the `evmredenom` module state is initialized with the EVM state root of the latest execution head.
- After the upgrade, an off-chain process scans the EVM state at that height, splits it into batches, and calls the `Redenom` contract to submit the data.
- Each batch contains a monotonic sequence of account addresses, account bodies, and a merkle range proof.
- The `evmredenom` verifies this is the expected next batch, then creates a withdrawal for each account
  to increase the balance by the expected amount.
- The EVM snapsync protocol account sync logic is used to query the EVM state (via P2P) and generate the proofs.
- The same snapsync logic is used in the `evmredenom` module to verify the proofs and apply the changes.
- This ensures that all accounts are submitted in the correct order.
- Mapping of account hash to account address is done by EVM genesis alloc and debug_preimage queries.
- For operational efficiency, a single archive can be configured to run the submission process at the upgrade height.
- Since that node stops processing the chain, this ensures that geth has the required state.
- Once all batches are submitted, the `evmredenom` module marks the process as complete.
