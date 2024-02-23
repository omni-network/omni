---
sidebar_position: 1
---

# Run a Validator

Omniʼs operator needs to run the consensus client and the execution client:

  - EVM client for execution (`geth`)
  - Omniʼs consensus client (`halo`)

## Hardware Requirements

| Hardware | Requirement |
| --- | --- |
| Cores | 8 |
| Bandwidth | 1Gbps |
| RAM | 32GB |
| SSD Hard Disk (expandable) | 4 TB |

## Software Requirements

| Software | Requirement |
| --- | --- |
| Docker | 24.0.7 |
| Operating System | Linux/macOS (arm/x46) |

### Ports

Inbound ports will be enabled for cometBFT (tcp://266567) and Geth (tcp://30303, udp://30303)

## Register as a Validator

Initially, we will have a whitelist. Ultimately, operators will be ranked based on their delegated TVL. Reach out to us to get on the whitelist.
