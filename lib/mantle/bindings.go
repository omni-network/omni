package mantle

// Generates bindings for the Mantle smart contracts.
//
//go:generate abigen --abi l1-bridge.json --type L1Bridge --pkg mantle --out l1bridge_bindings.go
