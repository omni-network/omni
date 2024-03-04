---
sidebar_position: 2
---

# Validator Requirements

Omniʼs operator needs to run the consensus client and the execution client:

- EVM client for execution (`geth`, `erigon`, etc)
- Omniʼs consensus client (`halo`)

## Hardware Requirements

| Hardware | Requirement (minimum) |
| --- | --- |
| Cores | 4 |
| Bandwidth | 100 Mbps |
| RAM | 32GB |
| SSD Hard Disk | 500 GB |

## Software Requirements

| Software | Requirement |
| --- | --- |
| Docker | 24.0.7 |
| Operating System | Linux/macOS (arm/64) |

### Ports

Inbound ports will be enabled for cometBFT (tcp://266567) and Geth (tcp://30303, udp://30303)
