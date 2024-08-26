# helpverify

A set of tools to help verify Omni contracts on EVM explorers.

## Usage

(Commands run from foundry root)

Get the creation tx hash

```bash
go run ./script/helpverify get-creation-tx-hash \
  --etherscan-api-key <key> \ # OR --arbscan-api-key <key> \
  <chain-name> \
  <contract-address>
```

With the tx hash, parse the creation tx. Note, this only works for transparent proxies deployed via Create3. This will print out the construct args and implmention address.

```bash
go run ./script/helpverify parse-proxy-create3-tx \
  <chain-name> \
  <tx-hash>
```

Then, use the constructor args to verify the proxy contract.

```bash
forge verify-contract
    --verifier-url <url> \ # see static.go for verifier urls
    --chain-id <foundry-chain-id> \ # optimism-sepolia, arbitrum-sepolia, etc
    --num-of-optimizations 200 \
    --compiler-version 0.8.24 \
    <proxy-address> \
    TransparentUpgradeableProxy
```

Verify the implementation contract by manually encoding constructor args, if there are any.
