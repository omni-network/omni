package umath

import (
	"math/big"

	"github.com/ethereum/go-ethereum/params"
)

// WeiToEtherF64 converts big.Int wei to float64 ether (wei/1e18).
func WeiToEtherF64(wei *big.Int) float64 {
	resp, _ := new(big.Int).Quo(wei, big.NewInt(params.Ether)).Float64()
	return resp
}

// WeiToGweiF64 converts big.Int wei to float64 gwei (wei/1e9).
func WeiToGweiF64(wei *big.Int) float64 {
	resp, _ := new(big.Int).Quo(wei, big.NewInt(params.GWei)).Float64()
	return resp
}
