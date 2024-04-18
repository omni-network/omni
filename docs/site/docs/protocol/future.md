---
sidebar_position: 4
---

# Future Work

Omni was intentionally designed not only to establish a new standard in the interoperability industry, but also to expand in both functionality and coverage in the future. Specifically, Omni will introduce novel finality mechanisms and support for alternative DA and consensus layers.

## Fast Finality Mechanisms

Omni’s `XMsg` verification process is designed with a modular approach. By default, Omni validators wait for `XMsg` requests to finalize on Ethereum L1 (2 epochs, ∼12 minutes) before attesting to their validity in consensus. However, low latency finality mechanisms can be substituted so that users can enjoy a UX that mirrors modern cloud applications. Specifically, Omni is exploring implementing transaction insurance mechanisms on source rollups or pre-confirmations from technologies like shared sequencers.

## Expanded DA and Consensus Support

Ethereum is still in the early phases of rollup design development. In the coming years, we anticipate the rollup ecosystem will proliferate in diversity. Projects will develop more customized rollup solutions, each tailored for specific functionalities and performance needs, incorporating unique virtual machines, programming languages, and data availability architectures. Understanding this progression, Omni is intentionally designed with minimal integration requirements to ensure compatibility with any rollup architecture, reflecting the distinct layer separations within the Internet Protocol stack. Thus, Omni promotes innovation at the rollup level while functioning as a global hub that integrates the entire rollup ecosystem.
