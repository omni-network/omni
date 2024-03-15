# Fireblocks
https://developers.fireblocks.com/reference/api-overview

## Transactions

This is going to be our primary API call. We use raw signing for all of our create transaction requests. This is because we cannot guranteee fireblocks is integrated with the chains we are deploying on.


## Request Signing:

https://developers.fireblocks.com/reference/signing-a-request-jwt-structure

### Note:
A deployment transaction is no different than a normal transaction, we build a transaction payload (the payload just specifies to deploy), we send the txn payload to fireblocks to be signed, we get it back and then send it to the chain.
