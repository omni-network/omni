package avs

import (
	"context"
	"math/big"
	"time"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/log"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/params"
)

// monitorAVSOperatorsForever blocks and periodically monitors AVS operators.
func monitorAVSOperatorsForever(ctx context.Context, avs *bindings.OmniAVS) {
	ticker := time.NewTicker(time.Second * 30)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			err := monitorAVSOnce(avs)
			if ctx.Err() != nil {
				return
			} else if err != nil {
				log.Error(ctx, "Monitoring AVS failed (will retry)", err)

				continue
			}
		}
	}
}

// monitorAVSOperatorsOnce monitors the AVS operators once, tracking operators
// registered, total delegations, and delegations per operator.
func monitorAVSOnce(avs *bindings.OmniAVS) error {
	operators, err := avs.Operators(&bind.CallOpts{})
	if err != nil {
		return errors.Wrap(err, "get operators")
	}

	// Reset all operator labeled metrics since some operators may have been removed.
	operatorStake.Reset()
	operatorDelegations.Reset()

	var total float64
	for _, operator := range operators {
		addr := operator.Addr.Hex()
		staked := weiToEth(operator.Staked)
		delegated := weiToEth(operator.Delegated)

		operatorStake.WithLabelValues(addr).Set(staked)
		operatorDelegations.WithLabelValues(addr).Set(delegated)

		total += delegated + staked
	}

	numOperators.Set(float64(len(operators)))
	totalDelegations.Set(total)

	return nil
}

// weiToEth converts a wei amount to an ether amount.
func weiToEth(wei *big.Int) float64 {
	f, _ := wei.Float64()
	return f / params.Ether
}
