---
"@omni-network/react": minor
"@omni-network/core": minor
---

Removed isNative flag from the quote inputs, deposit.token and expense.token are always optional now. By default we assume native when token is not provided, but we also support tokens in which case you need to provide the address.

This cleans up the interface for consumers as the isNative flag was proving awkward.
