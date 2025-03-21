package app

import (
	"context"
	"math/big"
	"time"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/lib/bi"
	"github.com/omni-network/omni/lib/contracts"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient"
	"github.com/omni-network/omni/lib/ethclient/ethbackend"
	"github.com/omni-network/omni/lib/expbackoff"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/umath"
	stokens "github.com/omni-network/omni/solver/tokens"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"

	"golang.org/x/sync/errgroup"
)

// approveOutboxes gives each outbox max allowance for all supported tokens.
// Most tokens will not decrement allowance when set to max, though some do.
// TODO: monitor allowances, alert or reset when "low" (half type(uint256).max).
func approveOutboxes(ctx context.Context, network netconf.Network, backends ethbackend.Backends, solverAddr common.Address) error {
	addrs, err := contracts.GetAddresses(ctx, network.ID)
	if err != nil {
		return errors.Wrap(err, "get addresses")
	}

	for _, chain := range network.EVMChains() {
		backend, err := backends.Backend(chain.ID)
		if err != nil {
			return errors.Wrap(err, "backend", "chain", chain.Name)
		}

		if err := approveOutbox(ctx, chain, backend, solverAddr, addrs.SolverNetOutbox); err != nil {
			return err
		}
	}

	return nil
}

// approveOutbox gives `outboxAddr` max allowance for all supported tokens on `chain`.
func approveOutbox(ctx context.Context, chain netconf.Chain, backend *ethbackend.Backend, solverAddr, outboxAddr common.Address) error {
	var eg errgroup.Group

	for _, token := range stokens.ByChain(chain.ID) {
		if token.IsNative() {
			continue
		}

		eg.Go(func() error {
			// use backoff incase approval fails. top level loop is too slow for retries
			// timeout single attempt after 2 minutes

			ctx, cancel := context.WithTimeout(ctx, time.Minute*2)
			defer cancel()

			// skip tokens that are not deployed. this ignores devnet tokens only "deployed" in fork mode
			isDeployed, err := contracts.IsDeployed(ctx, backend, token.Address)
			if err != nil {
				return errors.Wrap(err, "is deployed")
			} else if !isDeployed {
				return nil
			}

			backoff := expbackoff.New(ctx, expbackoff.WithPeriodicConfig(time.Second*5))

			for ctx.Err() == nil {
				if err := approveToken(ctx, backend, token, solverAddr, outboxAddr); err != nil {
					log.Warn(ctx, "Failed approve outbox, will retry", err, "chain", chain.Name, "token", token.Symbol)
					backoff()

					continue
				}

				log.Debug(ctx, "Approved token spend", "chain", chain.Name, "token", token.Symbol, "address", token.Address)

				return nil
			}

			return ctx.Err()
		})
	}

	if err := eg.Wait(); err != nil {
		return errors.Wrap(err, "approve outbox", "chain", chain.Name)
	}

	return nil
}

// approveToken gives `outboxAddr` max allowance for `token`.
func approveToken(ctx context.Context, backend *ethbackend.Backend, token stokens.Token, solverAddr, outboxAddr common.Address) error {
	erc20, err := bindings.NewIERC20(token.Address, backend)
	if err != nil {
		return errors.Wrap(err, "new token")
	}

	isApproved := func() (bool, error) {
		return isAppproved(ctx, token.Address, backend, solverAddr, outboxAddr, umath.MaxUint256)
	}

	if approved, err := isApproved(); err != nil {
		return err
	} else if approved {
		return nil
	}

	txOpts, err := backend.BindOpts(ctx, solverAddr)
	if err != nil {
		return err
	}

	tx, err := erc20.Approve(txOpts, outboxAddr, umath.MaxUint256)
	if err != nil {
		return errors.Wrap(err, "approve token")
	} else if _, err := backend.WaitMined(ctx, tx); err != nil {
		return errors.Wrap(err, "wait mined")
	}

	if approved, err := isApproved(); err != nil {
		return err
	} else if !approved {
		return errors.New("approve failed")
	}

	return nil
}

func isAppproved(
	ctx context.Context,
	token common.Address,
	client ethclient.Client,
	solverAddr, outboxAddr common.Address,
	spend *big.Int,
) (bool, error) {
	tkn, err := bindings.NewIERC20(token, client)
	if err != nil {
		return false, errors.Wrap(err, "new token")
	}

	allowance, err := tkn.Allowance(&bind.CallOpts{Context: ctx}, solverAddr, outboxAddr)
	if err != nil {
		return false, errors.Wrap(err, "get allowance")
	}

	return bi.LTE(spend, allowance), nil
}
