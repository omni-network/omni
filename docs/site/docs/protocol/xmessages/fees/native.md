---
sidebar_position: 2
---

# Native Token Roles & Utility

**\$OMNI** is the native token powering the Omni protocol. This section covers the roles and utility aspects of the **\$OMNI** tokenomics. The token serves the following functions within the network:

## Cross-Rollup Gas Fees

**\$OMNI** is used as a payment mechanism to compensate relayers for submitting cross-rollup messages to destination networks. Relayers maintain an inventory of gas assets for all supported rollups and accept payments in the form of **\$OMNI**. This allows Omni to establish a universal gas marketplace that simplifies the process of gas payments across all rollups for end users.

<figure>
  <img src="/img/gas-marketplace.png" alt="Global Gas Marketplace" />
  <figcaption>*Universal gas marketplace simplifying gas payments across all rollups*</figcaption>
</figure>

## Omni EVM Gas Fees

**\$OMNI** also serves as the native gas asset that powers the Omni EVM. The Omni EVM functions as a global orchestration layer for application instances across multiple rollups, allowing users and developers to initiate transactions and manage applications on any rollup from a single source. **\$OMNI** provides an anti-sybil mechanism for transactions submitted to the Omni EVM, deterring spam and malicious activities such as denial-of-service attacks. **\$OMNI** also serves as compensation for Omni validators investing computational power for transaction processing and network security. Users can choose to pay higher **\$OMNI** fees to validators for priority transactions, thereby establishing a fee market based on **\$OMNI**.

## Network Governance

At launch, the initial Omni protocol will provide basic cross-rollup messaging functionality across a specified set of rollups. As Omni matures, **\$OMNI** stakeholders will be responsible for various governance decisions such as protocol upgrades and additional developer features.

## Reinforced Security

Omni achieves stronger and more stable security guarantees than existing interoperability protocols by deriving its cryptoeconomic security from restaked **\$ETH**. Omni extends its security model further by incorporating staked **\$OMNI** using a dual staking model. Effectively, the total cryptoeconomic security of Omni is determined by the combined value of restaked **\$ETH** and staked **\$OMNI**.

By implementing this dual staking model, Omni’s security scales across two dimensions. Restaked **\$ETH** anchors Omni’s security to Ethereum L1, enabling it to grow in line with Ethereum’s own security budget. The addition of staked **\$OMNI** builds upon this base, expanding Omni’s security alongside its own network activity. Collectively, these two complementary mechanisms provide robust and dynamic security guarantees for Omni, setting a new standard for secure interoperability for the Ethereum ecosystem.
