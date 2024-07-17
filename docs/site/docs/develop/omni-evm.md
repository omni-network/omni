---
sidebar_position: 3
---

# Omni EVM

Omni's dual client architecture means that the EVM execution clients are exactly the same – unmodified – as the clients running on Ethereum L1. This means that your solidity code will work on Omni out of the box.

## Precompiles, Preinstalls, and Predeploys

These are a set of contracts that exist on Omni's EVM at genesis.

### Precompiles

Precompiles are part of every EVM execution client, and give developers access to sophisticated logic that may not be possible or feasible through opcodes / solidity directly. Omni's pre-compiles match Ethereum L1's directly.

An example of a precompile is `ecrecover`, which can be used to recover an address associated with the public key from elliptic curve signatures.

A list of precompiles is available at [evm.codes](https://www.evm.codes/precompiled)

### Preinstalls

Preinstalls are solidity contracts that are frequently used by developers that Omni deploys at genesis. In some cases, these contracts can be deployed through standard means. However, in other cases, some of the contracts are deployed to Ethereum and other chains as singletons at a standard address through [Nick's method](https://yamenmerhi.medium.com/nicks-method-ethereum-keyless-execution-168a6659479c). If Nick's method was used before EIP-150, the contract can't typically be deployed to the expected address. Thus, Omni deploys them at genesis to the expected address.

A full list of preinstalls can be found in the Omni code base [here](https://github.com/omni-network/omni/blob/main/contracts/core/src/octane/Preinstalls.sol).

A few examples include:

- [multicall3](https://www.multicall3.com/)
- [DeterministicDeploymentProxy](https://github.com/Zoltu/deterministic-deployment-proxy)
- [ERC1820Registry](https://eips.ethereum.org/EIPS/eip-1820)
- and more...

### Predeploys

These are contracts deployed at genesis that are frequently used by the Omni protocol, but aren't relevant to smart contract developers, except in very specific cases (like if you are building a liquid staking protocol).

Examples include:

- [Staking.sol](https://github.com/omni-network/omni/blob/main/contracts/core/src/octane/Staking.sol): for validators to register, stake their $OMNI tokens, etc.
- [Slashing.sol](https://github.com/omni-network/omni/blob/main/contracts/core/src/octane/Slashing.sol): for validators to unjail themselves (if jailed for inactivity)

## EVM Differences

Please note that the Omni EVM does not currently support [RANDAO](https://eth2book.info/capella/part2/building_blocks/randomness/).
