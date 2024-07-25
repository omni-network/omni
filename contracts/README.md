# Omni Contracts

Omni smart contracts and related software.

## Contents

<pre>
├── <a href="./core/">core/</a>: Core protocol smart contracts.
├── <a href="./avs/">avs/</a>: Eigen AVS smart contracts.
├── <a href="./bindings/">bindings/</a>: Go smart contract bindings.
├── <a href="./allocs/">allocs/</a>: Predeploy allocations.
</pre>


## Build

```bash
make build      # compile the smart contracts (avs & core)
make bindings   # generate go bindings
make allocs     # generate predeploy allocations
make all        # all of the above
```
