---
sidebar_position: 4
---

import GitHubCodeBlock from '@site/src/components/GitHubCodeBlock/GitHubCodeBlock';

# Contracts

A reference for Omni's user facing solidity contracts and libraries.

### [`OmniPortal`](https://github.com/omni-network/omni/blob/main/contracts/src/protocol/OmniPortal.sol)

- On-chain gateway into Omni's cross-chain messaging protocol
- Call contracts on another chain (`xcall`)
- Calculate fees for an `xcall`
- Read `xmsg` context when receiving an `xcall`

<details>
<summary>`IOmniPortal.sol` Reference Solidity Interface</summary>

<GitHubCodeBlock url="https://github.com/omni-network/omni/blob/main/contracts/src/interfaces/IOmniPortal.sol" />
</details>

### [`XApp`](https://github.com/omni-network/omni/blob/main/contracts/src/pkg/XApp.sol)

- Base contract for Omni cross-chain applications
- Simplifies sending / receiving `xcalls`

<details>
<summary>`XApp.sol` Reference Solidity Interface</summary>

<GitHubCodeBlock url="https://github.com/omni-network/omni/blob/main/contracts/src/pkg/XApp.sol" />
</details>

### [`XTypes`](https://github.com/omni-network/omni/blob/main/contracts/src/libraries/XTypes.sol)

- Defines core xchain messaging types for the Omni protocol.
- `XTypes.MsgShort` is the only type end users interact with. It provides context

<details>
<summary>`XTypes.sol` Reference Solidity Code</summary>

<GitHubCodeBlock url="https://github.com/omni-network/omni/blob/main/contracts/src/libraries/XTypes.sol" />
</details>
