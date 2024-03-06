---
sidebar_position: 3
---

# Omni AVS Contract Registration

This tells EigenLayer that you'd like to be an operator specifically for the Omni AVS. Thus, the **\$ETH** that you, and your delegators restaked, will be used to secure Omni.

You will need to have an `operator.yml` file to perform this registration. This file is created as part of [registering as an operator with the EigenLayer CLI](https://docs.eigenlayer.xyz/eigenlayer/operator-guides/operator-installation).

<details>
<summary>`operator.yml` Reference</summary>

For further information on this reference, please refer to the [EigenLayer reference example](https://github.com/Layr-Labs/eigenlayer-cli/blob/master/pkg/operator/config/operator-config-example.yaml).

```yaml
operator:
    address: 0xfd23f7f705344bce1582fcf9bc6a0dc8e33b3b61 # Your operator address
    earnings_receiver_address: 0xfd23f7f705344bce1582fcf9bc6a0dc8e33b3b61 # Your operator payout address, may be the same as above
    delegation_approver_address: "0x0000000000000000000000000000000000000000" # Your delegation approver address, may be left as shown
    staker_opt_out_window_blocks: 0 # may be left as shown, and can be updated later using EigenLayer CLI
    metadata_url: "https://raw.githubusercontent.com/idea404/resources/main/eigenlayer/metadata.json" # Your metadata URL
el_delegation_manager_address: 0x8ce361602B935680E8DeC218b820ff5056BeB7af # The address of the EigenLayer delegation manager on the Omni Network
eth_rpc_url: http://127.0.0.1:8002 # Your node Ethereum RPC URL
private_key_store_path: /Users/idea404/.eigenlayer/operator_keys/OpKeys1.ecdsa.key.json # Your private key store path generated or imported by EigenLayer CLI
signer_type: local_keystore # Your signer type, may be left as shown
chain_id: 100 # The chain ID of the Omni Network
```

</details>

## Register as an Operator

1. Ensure that your node address has been added to the allowed list of operators.
2. Run the following command to register as an operator:

```bash
omni operator register --config-file ~/path/to/operator.yml
```

:::info

The AVS address (`0x848BE3DBcd054c17EbC712E0d29D15C2e638aBCe`) is the address of the Omni AVS contract [deployed on Goerli testnet](https://goerli.etherscan.io/address/0x848BE3DBcd054c17EbC712E0d29D15C2e638aBCe) that will be called by the `register` command.

:::
