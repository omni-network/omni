---
sidebar_position: 3
---

# Finality

After verifying the contents of a proposed consensus block and appending `XBlock` hash attestations, `halo` clients use CometBFT consensus to vote on the validity of the block.

If the proposal is determined valid by ⅔ of the active validator set, the block is finalized and becomes the latest block committed to the network. These confirmed blocks contain all transactions from the Omni EVM along with the hashes of attested `XBlock`s, allowing anyone with a validator node for a given rollup VM to reconstruct any `XBlock` and verify its contents.

Omni’s [Integrated Consensus](../xmessages/components/validator/cometbft.md) processes state transitions for both the Omni EVM and external VMs, preventing both complementary sub-processes from interfering with one another prior to block finalization. ABCI++ allows validators to keep `XBlock` attestations separate from the single CometBFT transaction within each block.
