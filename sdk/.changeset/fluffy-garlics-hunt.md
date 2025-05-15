---
"@omni-network/react": patch
"@omni-network/core": patch
---

Use the new useWatchDidFill hook in useOrder to propagate the destTxHash, but fallback to didFill in case consumers rely on public RPCs
