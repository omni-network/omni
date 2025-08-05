package testupgrade

import (
	"context"
	"math/big"
	"time"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/e2e/app/eoa"
	"github.com/omni-network/omni/e2e/types"
	"github.com/omni-network/omni/halo/evmredenom"
	"github.com/omni-network/omni/halo/genutil/evm/predeploys"
	"github.com/omni-network/omni/lib/anvil"
	"github.com/omni-network/omni/lib/bi"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient/ethbackend"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/txmgr"

	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"

	"golang.org/x/sync/errgroup"
)

var (
	preBalances = map[string]*big.Int{
		"0x0000000000000000111111000000000000000000": bi.Wei(99), // Less than 1 gwei, no ignored.
		"0x0000000000000000222222000000000000000000": bi.Gwei(123.999_999_999),
		"0x0000000000000000333333000000000000000000": bi.Zero(),
		"0xFFFFFFFFFFFFFFFF444444000000000000000000": bi.Ether(456),
	}

	delegator  = eoa.RoleTester
	delegation = bi.Ether(2)
)

// prepForEarhart funds predefined addresses with specific balances.
func prepForEarhart(ctx context.Context, testnet types.Testnet, omniEVM *ethbackend.Backend) error {
	from := anvil.DevAccount9()

	if err := delegate(ctx, testnet, omniEVM); err != nil {
		return err
	}

	eg, ctx := errgroup.WithContext(ctx)
	for addrHex, bal := range preBalances {
		eg.Go(func() error {
			addr := common.HexToAddress(addrHex)
			if b, err := omniEVM.BalanceAt(ctx, addr, nil); err != nil {
				return err
			} else if !bi.IsZero(b) {
				return errors.Wrap(err, "non zero balance")
			}

			_, rec, err := omniEVM.Send(ctx, from, txmgr.TxCandidate{
				To:    &addr,
				Value: bal,
			})
			if err != nil {
				return err
			} else if rec.Status != ethtypes.ReceiptStatusSuccessful {
				return errors.New("send failed")
			}

			if b, err := omniEVM.BalanceAt(ctx, addr, nil); err != nil {
				return err
			} else if !bi.EQ(b, bal) {
				return errors.New("balance mismatch", "addr", addrHex, "expected", bal, "got", b)
			}

			return nil
		})
	}

	if err := eg.Wait(); err != nil {
		return errors.Wrap(err, "funding addresses failed")
	}

	return nil
}

func undelegate(ctx context.Context, testnet types.Testnet, evm *ethbackend.Backend) error {
	var valAddr common.Address
	for n := range testnet.Validators {
		pk, err := crypto.ToECDSA(n.PrivvalKey.Bytes())
		if err != nil {
			return errors.Wrap(err, "to ecdsa")
		}
		valAddr = crypto.PubkeyToAddress(pk.PublicKey)
	}

	addr := eoa.MustAddress(testnet.Network, delegator)

	txOpts, err := evm.BindOpts(ctx, addr)
	if err != nil {
		return errors.Wrap(err, "bind opts")
	}
	txOpts.Value = bi.Ether(0.1) // Fee

	contract, err := bindings.NewStaking(common.HexToAddress(predeploys.Staking), evm)
	if err != nil {
		return errors.Wrap(err, "new staking contract")
	}

	preBal, err := evm.BalanceAt(ctx, addr, nil)
	if err != nil {
		return err
	}

	// undelegation amount is 2 $STAKE, which is 150 $NATIVE_EVM.
	amount := bi.MulRaw(delegation, evmredenom.Factor)

	_, err = contract.Undelegate(txOpts, valAddr, amount)
	if err != nil {
		return errors.Wrap(err, "delegate for")
	}

	// Wait for balance to be increased by more than evmredenom.Factor (incl rewards).
	ctx, cancel := context.WithTimeout(ctx, time.Minute)
	defer cancel()
	for i := 0; ; i++ {
		b, err := evm.BalanceAt(ctx, eoa.MustAddress(testnet.Network, delegator), nil)
		if err != nil {
			return errors.Wrap(err, "balance check failed")
		}

		increase := bi.Sub(b, preBal)
		if bi.GT(increase, amount) {
			break // Done
		}

		if i > 5 {
			log.Warn(ctx, "Waiting for undelegation withdraw", nil, "increase", increase, "expect", amount)
		}

		select {
		case <-ctx.Done():
			return errors.New("timeout waiting for undelegation withdraw", "increase", increase, "expect", amount)
		case <-time.After(time.Second):
		}
	}

	return nil
}
func delegate(ctx context.Context, testnet types.Testnet, evm *ethbackend.Backend) error {
	var valAddr common.Address
	for n := range testnet.Validators {
		pk, err := crypto.ToECDSA(n.PrivvalKey.Bytes())
		if err != nil {
			return errors.Wrap(err, "to ecdsa")
		}
		valAddr = crypto.PubkeyToAddress(pk.PublicKey)
	}

	txOpts, err := evm.BindOpts(ctx, eoa.MustAddress(testnet.Network, delegator))
	if err != nil {
		return errors.Wrap(err, "bind opts")
	}
	txOpts.Value = delegation

	contract, err := bindings.NewStaking(common.HexToAddress(predeploys.Staking), evm)
	if err != nil {
		return errors.Wrap(err, "new staking contract")
	}

	_, err = contract.Delegate(txOpts, valAddr)
	if err != nil {
		return errors.Wrap(err, "delegate for")
	}

	return nil
}

// ensureEarhart ensure the balances of predefined addresses after the Earhart upgrade are increased.
func ensureEarhart(ctx context.Context, testnet types.Testnet, omniEVM *ethbackend.Backend) error {
	ctx, cancel := context.WithTimeout(ctx, time.Minute)
	defer cancel()

	log.Debug(ctx, "Ensuring 4_earhart redenomination")

	for addrHex, bal := range preBalances {
		// Calculate expected balance after upgrade.
		add := bi.MulRaw(bal, evmredenom.Factor-1)
		addRem := new(big.Int).Rem(add, bi.Gwei(1)) // Gwei truncation
		exp := bi.Add(bal, bi.Sub(add, addRem))

		for {
			addr := common.HexToAddress(addrHex)
			b, err := omniEVM.BalanceAt(ctx, addr, nil)
			if err != nil {
				return errors.Wrap(err, "balance check failed", "addr", addrHex)
			} else if bi.IsZero(b) {
				// Account not created, so prepEarhart did not run, skip balance check.
				break
			} else if bi.EQ(b, exp) {
				break
			}

			log.Debug(ctx, "Still waiting for redenominated balance (will retry)", "addr", addrHex, "expected", exp, "got", b)

			select {
			case <-ctx.Done():
				return errors.New("timeout waiting for redenominated balance")
			case <-time.After(time.Second):
			}
		}
	}

	log.Info(ctx, "All 4_earhart redenomination withdrawals complete")

	if err := undelegate(ctx, testnet, omniEVM); err != nil {
		return errors.Wrap(err, "undelegate failed")
	}

	return nil
}
