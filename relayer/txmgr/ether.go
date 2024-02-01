package txmgr

import (
	"math"
	"math/big"

	"github.com/omni-network/omni/lib/errors"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/params"
)

// GweiToWei converts a float64 GWei value into a big.Int Wei value.
func GweiToWei(gwei float64) (*big.Int, error) {
	if math.IsNaN(gwei) || math.IsInf(gwei, 0) {
		return nil, errors.New("invalid gwei value", gwei)
	}

	// convert float GWei value into integer Wei value
	wei, _ := new(big.Float).Mul(
		big.NewFloat(gwei),
		big.NewFloat(params.GWei)).
		Int(nil)

	if wei.Cmp(abi.MaxUint256) == 1 {
		return nil, errors.New("gwei value larger than max uint256")
	}

	return wei, nil
}
