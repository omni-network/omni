package uniswap

// Generates bindings for the Uniswap's SwapRouter02 and QuoterV2 contracts.
//
//go:generate abigen --abi uniswaprouter02.json --type UniSwapRouter02 --pkg uniswap --out uniswaprouter02_bindings.go
//go:generate abigen --abi uniquoterv2.json --type UniQuoterV2 --pkg uniswap --out uniquoterv2_bindings.go
//go:generate abigen --abi weth9.json --type WETH9 --pkg uniswap --out weth9_bindings.go
