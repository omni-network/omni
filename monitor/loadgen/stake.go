package loadgen

import (
	"context"
	"crypto/ecdsa"
	"math/big"
	"time"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient/ethbackend"
	"github.com/omni-network/omni/lib/k1util"
	"github.com/omni-network/omni/lib/log"

	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/params"

	"math/rand/v2"
)

const selfDelegateJitter = 0.2 // 20% jitter

func selfDelegateForever(ctx context.Context, contract *bindings.OmniStake, backend *ethbackend.Backend, validator *ecdsa.PublicKey, period time.Duration) {
	addr := crypto.PubkeyToAddress(*validator)

	log.Info(ctx, "Starting periodic self-delegation", "validator", addr.Hex(), "period", period)

	nextPeriod := func() time.Duration {
		jitter := time.Duration(float64(period) * rand.Float64() * selfDelegateJitter)
		return period + jitter
	}

	timer := time.NewTimer(nextPeriod())
	defer timer.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-timer.C:
			if err := selfDelegateOnce(ctx, contract, backend, validator); err != nil {
				log.Error(ctx, "Failed to self-delegate (will retry)", err)
			}
			timer.Reset(nextPeriod())
		}
	}
}

func selfDelegateOnce(ctx context.Context, contract *bindings.OmniStake, backend *ethbackend.Backend, validator *ecdsa.PublicKey) error {
	addr := crypto.PubkeyToAddress(*validator)

	ethBalance, err := backend.EtherBalanceAt(ctx, addr)
	if err != nil {
		return err
	} else if ethBalance < 1 {
		return errors.New("insufficient balance to self-delegate",
			"balance", ethBalance,
			"validator", addr.Hex(),
		)
	}

	txOpts, err := backend.BindOpts(ctx, addr)
	if err != nil {
		return err
	}
	txOpts.Value = big.NewInt(params.Ether) // 1 ETH (in wei)

	tx, err := contract.Deposit(txOpts, k1util.PubKeyToBytes64(validator))
	if err != nil {
		return errors.Wrap(err, "deposit")
	}

	rec, err := backend.WaitMined(ctx, tx)
	if err != nil {
		return err
	}

	log.Info(ctx, "Deposited validator self-delegation",
		"height", rec.BlockNumber,
		"validator", addr.Hex(),
	)

	return nil
}
