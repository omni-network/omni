---
sidebar_position: 6
---

# Client Diversity

Omni's EVM Engine champions the principle of client diversity, allowing the integration and use of any Ethereum Virtual Machine (EVM) client within its ecosystem. This approach underscores Omni's commitment to flexibility, scalability, and interoperability, promoting a robust and resilient blockchain network.

### Integration of EVM Clients

Omni supports the seamless integration of various EVM clients, such as Geth, Besu, Erigon, and others, without requiring specialized modifications. This inclusivity enables the Omni network to benefit from the diverse set of features and optimizations offered by different clients, enhancing the overall ecosystem.

```go
// Example of client diversity support
func NewKeeper(...) *Keeper {
    // Keeper initialization with flexible client integration...
}
```

Omni's Keeper module can be initialized with any EVM client of choice, ensuring that the network remains open to a wide range of execution clients.

### Advantages of Client Diversity

- **Innovation and Improvement:** By supporting a variety of EVM clients, Omni encourages innovation and continuous improvement within the ecosystem.
- **Security and Robustness:** Diversity in execution clients can enhance network security, mitigating the risk of widespread issues stemming from client-specific vulnerabilities.
- **Customization and Optimization:** Developers and node operators have the freedom to choose an EVM client that best fits their needs, optimizing for performance, tooling, or other criteria.
