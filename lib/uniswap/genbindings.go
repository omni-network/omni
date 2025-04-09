package uniswap

// Generates bindings for the Uniswap's SwapRouter02 contract.
//go:generate abigen --abi uniswaprouter02.json --type UniSwapRouter02 --pkg uniswap --out uniswaprouter02_bindings.go
//go:generate abigen --abi uniquoterv2.json --type UniQuoterV2 --pkg uniswap --out uniquoterv2_bindings.go
