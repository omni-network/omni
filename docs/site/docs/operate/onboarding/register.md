---
sidebar_position: 3
---

# Omni AVS Contract Registration

This tells EigenLayer that you'd like to be an operator specifically for the Omni AVS. Thus, the **\$ETH** that you, and your delegators restaked, will be used to secure Omni.

You will need to have an `operator.yml` file to perform this registration. This file is created as part of [registering as an operator with the EigenLayer CLI](https://docs.eigenlayer.xyz/eigenlayer/operator-guides/operator-installation). You should use the same file created then, and don't need to modify it at all.

<details>
<summary>`operator.yml` Reference</summary>

For further information on this reference, please refer to the [EigenLayer reference example](https://github.com/Layr-Labs/eigenlayer-cli/blob/master/pkg/operator/config/operator-config-example.yaml). Note that this example was generated for Holesky testnet. You should use your pre-existing `operator.yaml`, configured for the network you're registering on.

**Testnet**

```yaml
operator:
    address: 0xfd23f7f705344bce1582fcf9bc6a0dc8e33b3b61 # Your operator address
    earnings_receiver_address: 0xfd23f7f705344bce1582fcf9bc6a0dc8e33b3b61 # Your operator payout address, may be the same as above
    delegation_approver_address: "0x0000000000000000000000000000000000000000" # Your delegation approver address, may be left as shown
    staker_opt_out_window_blocks: 0 # may be left as shown, and can be updated later using EigenLayer CLI
    metadata_url: "https://raw.githubusercontent.com/idea404/resources/main/eigenlayer/metadata.json" # Your metadata URL
el_delegation_manager_address: 0xA44151489861Fe9e3055d95adC98FbD462B948e7 # The address of the EigenLayer delegation manager on Holesky
eth_rpc_url: https://ethereum-holesky-rpc.publicnode.com # Holesky RPC URL
private_key_store_path: /Users/idea404/.eigenlayer/operator_keys/OpKeys1.ecdsa.key.json # Your private key store path generated or imported by EigenLayer CLI
signer_type: local_keystore # Your signer type, may be left as shown
chain_id: 17000 # The chain ID of Holesky
```

**Mainnet**

```yaml
operator:
    address: 0xfd23f7f705344bce1582fcf9bc6a0dc8e33b3b61 # Your operator address
    earnings_receiver_address: 0xfd23f7f705344bce1582fcf9bc6a0dc8e33b3b61 # Your operator payout address, may be the same as above
    delegation_approver_address: "0x0000000000000000000000000000000000000000" # Your delegation approver address, may be left as shown
    staker_opt_out_window_blocks: 0 # may be left as shown, and can be updated later using EigenLayer CLI
    metadata_url: "https://raw.githubusercontent.com/idea404/resources/main/eigenlayer/metadata.json" # Your metadata URL
el_delegation_manager_address: 0x39053D51B77DC0d36036Fc1fCc8Cb819df8Ef37A # Ethereum mainnet address for the EigenLayer delegation manager
eth_rpc_url: https://ethereum-rpc.publicnode.com # Ethereum mainnet RPC URL
private_key_store_path: /Users/idea404/.eigenlayer/operator_keys/OpKeys1.ecdsa.key.json # Your private key store path generated or imported by EigenLayer CLI
signer_type: local_keystore # Your signer type, may be left as shown
chain_id: 1 # The chain ID of Ethereum mainnet
```

</details>

## Register as an Operator

Run the following command to register as an operator:

```bash
omni operator register --config-file ~/path/to/operator.yaml
```

:::info

The Testnet AVS address (`0xa7b2e7830C51728832D33421670DbBE30299fD92`) is the address of the Omni AVS contract [deployed on Holesky](https://holesky.etherscan.io/address/0xa7b2e7830C51728832D33421670DbBE30299fD92) that will be called by the `register` command.

The Mainnet AVS address is [`0xed2f4d90b073128ae6769a9A8D51547B1Df766C8`](https://etherscan.io/address/0xed2f4d90b073128ae6769a9A8D51547B1Df766C8).

:::
