---
sidebar_position: 1
---

# Background

Ethereum’s commitment to the rollup-centric roadmap has forced the network to scale via isolated execution environments. While this allows rollups to support various execution environments and programming languages, it creates negative externalities that degrade Ethereum’s network effects. Specifically, liquidity, users, and developers are fragmented between disparate ecosystems. The expanding variety of rollup designs and their growing adoption will only exacerbate these issues. Consequently, Ethereum requires a native interoperability protocol that realigns the network with its original vision of being a single, unified operating system for decentralized applications.

### The Ideal Interoperability Solution

The ideal Ethereum interoperability solution must be **secure**, **performant**, and **globally compatible with the Ethereum ecosystem**. To meet our security standards, the solution should derive its security from the same source as Ethereum rollups: Ethereum L1. To be considered performant, the interoperability protocol must verify and process cross-rollup messages with minimal latency. For the protocol to be globally compatible, it must enable applications to be Turing-complete across all rollup environments, ensuring that applications are not limited by the resource constraints of any single rollup.

### Interoperability Verification Approaches

Existing interoperability protocol designs can be broadly classified according to their verification approach.

<figure>
  <img src="/img/verification-table.png" alt="Interoperability Verification Approaches" />
  <figcaption>Interoperability verification approaches</figcaption>
</figure>

Natively verified systems ensure trustless interoperability but require the underlying network security providers to validate state changes. Since rollups derive security from Ethereum L1, Ethereum L1 is the only natively verified solution for rollup interoperability. However, validating rollup state changes on Ethereum L1 via optimistic or zero-knowledge mechanisms is both costly and time consuming, making Ethereum L1 an unfit interoperability solution for rollups.

Similar to a strictly native approach, other approaches also fail to meet our three criteria of security, performance, and global compatibility. Solutions that rely on local verification provide strong security guarantees and can be applied
across various rollups, yet their inability to handle arbitrary messages limits their application compatibility. Optimistic approaches, which incorporate a latency period for potential challenges, inherently do not satisfy our performance requirement.

Externally verified systems meet our performance and global compatibility needs but require trust in an external verifier set, compromising security. To enhance security, these systems can implement cryptoeconomic security measures that require verifiers to stake capital that is at risk of being slashed for misconduct. However, existing solutions that use this approach derive their cryptoeconomic security from a native asset rather than Ethereum L1, making them unfit for our security requirements. By devising a method to extend Ethereum L1’s cryptoeconomic security to an external set of verifiers, we can fulfill our security criteria and create a novel interoperability solution that offers all three of our desired properties.
