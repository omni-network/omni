# Halo

Halo is a cosmos-SDK application blockchain client.

# Flow

The following describes the consensus flow:
```
   CometBFT         ABCI++        ┌──── EVMEngine ─────────
 ─────────────────────────────    │
                                  │ 1. Did we trigger an optimistic payload build earlier?
 1. Ask proposer                  │    - No:  ask geth to build a new payload (execution block)
 to build next block              │    - Yes: no need to ask geth, since we already did
       │                          │
       │                          │ 2. Get resulting built payload from geth
       └──────► ┌───────────────┐ │
                │PrepareProposal│ │ 3. Collect txs (msgs) to include in consensus block
       ┌─────── └───────────────┘ │    - Make MsgExecutionPayload with above payload
       │                          │    - Ask attest module for MsgAddVote (extracted from VE)
       │                          │    - Ask EVM tracking modules for other TXs (e.g. staking)
       ▼                          │
 2. Send proposal                 │ 4. Return txs (msgs) to include consenus block
 to all validators
       │
       │                          ┌──── App ───────────────
       │                          │
       └──────► ┌───────────────┐ │ 1. Extract and decode all txs in proposed block
                │ProcessProposal│ │
       ┌─────── └───────────────┘ │ 2. Ask EVMEngine to verify MsgExecutionPayload
       │                          │    - Ask geth if payload is a valid head
       ▼                          │
  3. Continue and                 │ 3. Ask Attest module to verify MsgAddVote is valid
  complete consensus
       │
       ▼
  4. Provide finalised
  consensus block to              ┌──── App ───────────────
  all validators                  │
  and full nodes                  │ 1. Extract and decode all txs in finalised block
       │                          │
       └────────► ┌─────────────┐ │ 2. Ask EVMEngine to process MsgExecutionPayload
                  │FinalizeBlock│ │    - Inform geth of this new payload
       ┌───────── └─────────────┘ │    - Instruct geth to update head to this new payload
       │                          │    - If this node is next proposer (and optimistic enabled)
       ▼                          │       - Ask geth to build a new payload
  5. Sleep and then               │
     start again at 1.            │ 3. Ask Attest module to process MsgAddVote
                                  │    - Merge all votes into the application DB (creating
                                  │      pending attestations for new xblocks)
                                  │    - Approve all attestations with quorum votes
```
