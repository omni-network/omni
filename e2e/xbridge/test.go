package xbridge

import (
	"context"
	"math/big"
	"math/rand"
	"time"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/e2e/app/eoa"
	"github.com/omni-network/omni/e2e/xbridge/rlusd"
	"github.com/omni-network/omni/e2e/xbridge/types"
	"github.com/omni-network/omni/lib/anvil"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient/ethbackend"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/umath"
	"github.com/omni-network/omni/lib/xchain"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

func Test(ctx context.Context, network netconf.Network, endpoints xchain.RPCEndpoints) error {
	if network.ID != netconf.Devnet {
		return errors.New("only devnet")
	}

	pks := append(anvil.DevPrivateKeys(), eoa.DevPrivateKeys()...)
	backends, err := ethbackend.BackendsFromNetwork(network, endpoints, pks...)
	if err != nil {
		return err
	}

	if err := testRLUSD(ctx, network, backends); err != nil {
		return errors.Wrap(err, "test tokens")
	}

	return nil
}

func testRLUSD(ctx context.Context, network netconf.Network, backends ethbackend.Backends) error {
	user := anvil.DevAccount5()
	amt := umath.Ether(1_000_000) // 1M

	if err := rlusd.MintCanonical(ctx, network, backends, user, amt); err != nil {
		return errors.Wrap(err, "mint wrapped")
	}

	if err := testBridge(ctx, network, backends, rlusd.XToken(), user, amt); err != nil {
		return errors.Wrap(err, "test bridge")
	}

	return nil
}

// testBridge tests the bridging for a given xtoken.
func testBridge(
	ctx context.Context,
	network netconf.Network,
	backends ethbackend.Backends,
	tkn types.XToken,
	user common.Address,
	amt *big.Int,
) error {
	canon, err := tkn.Canonical(ctx, network.ID)
	if err != nil {
		return errors.Wrap(err, "canonical", "xtoken", tkn.Symbol())
	}

	chains := network.EVMChains()

	bridge, err := newBridger(ctx, network, backends, tkn)
	if err != nil {
		return errors.Wrap(err, "bridger")
	}

	// returns a random chain, not equal to src
	nextDst := func(src uint64) uint64 {
		for {
			//nolint:gosec // does not need to be secure
			c := rand.Intn(len(chains))
			if chains[c].ID == src {
				continue
			}

			return chains[c].ID
		}
	}

	// starting with wrapped, hop 3 times
	src := canon.ChainID
	for i := 0; i < 3; i++ {
		dst := nextDst(src)

		log.Info(ctx, "Bridging", "src", src, "dst", dst, "xtoken", tkn.Symbol())

		if err := bridge(src, dst, user, amt); err != nil {
			return errors.Wrap(err, "bridge", "src", src, "dst", dst)
		}

		log.Info(ctx, "Bridged", "src", src, "dst", dst, "xtoken", tkn.Symbol())

		src = dst
	}

	return nil
}

func newBridger(
	ctx context.Context,
	network netconf.Network,
	backends ethbackend.Backends,
	xtoken types.XToken,
) (func(src, dst uint64, user common.Address, amt *big.Int) error, error) {
	canon, err := xtoken.Canonical(ctx, network.ID)
	if err != nil {
		return nil, errors.Wrap(err, "canonical")
	}

	bridgeAddr, err := BridgeAddr(ctx, network.ID, xtoken)
	if err != nil {
		return nil, errors.Wrap(err, "bridge addr")
	}

	xtokenAddr, err := xtoken.Address(ctx, network.ID)
	if err != nil {
		return nil, errors.Wrap(err, "token addr")
	}

	tknAddr := func(chainID uint64) common.Address {
		if chainID == canon.ChainID {
			return canon.Address
		}

		return xtokenAddr
	}

	return func(src, dst uint64, user common.Address, amt *big.Int) error {
		srcBackend, err := backends.Backend(src)
		if err != nil {
			return errors.Wrap(err, "src backend")
		}

		dstBackend, err := backends.Backend(dst)
		if err != nil {
			return errors.Wrap(err, "dst backend")
		}

		dstTkn, err := bindings.NewIERC20(tknAddr(dst), dstBackend)
		if err != nil {
			return errors.Wrap(err, "dst token")
		}

		srcTkn, err := bindings.NewIERC20(tknAddr(src), srcBackend)
		if err != nil {
			return errors.Wrap(err, "src token")
		}

		bridge, err := bindings.NewBridge(bridgeAddr, srcBackend)
		if err != nil {
			return errors.Wrap(err, "src bridge")
		}

		//
		// approve bridge
		//

		txOpts, err := srcBackend.BindOpts(ctx, user)
		if err != nil {
			return errors.Wrap(err, "tx opts")
		}

		wrap := src == canon.ChainID

		if wrap {
			// approval only needed when wrapping
			tx, err := srcTkn.Approve(txOpts, bridgeAddr, amt)
			if err != nil {
				return errors.Wrap(err, "approve")
			}

			_, err = srcBackend.WaitMined(ctx, tx)
			if err != nil {
				return errors.Wrap(err, "wait mined")
			}
		}

		//
		// bridge to dst
		//

		fee, err := bridge.BridgeFee(&bind.CallOpts{Context: ctx}, dst)
		if err != nil {
			return errors.Wrap(err, "bridge fee")
		}

		txOpts.Value = fee

		tx, err := bridge.SendToken(txOpts, dst, user, amt, wrap, user)
		if err != nil {
			log.Error(ctx, "Send token", err, "custom", detectCustomError(err), "wrap", wrap, "src", src, "dst", dst)
			return errors.Wrap(err, "send token", "custom", detectCustomError(err))
		}

		_, err = srcBackend.WaitMined(ctx, tx)
		if err != nil {
			return errors.Wrap(err, "wait mined")
		}

		return waitForBalance(ctx, dstTkn, user, amt)
	}, nil
}

func waitForBalance(ctx context.Context, tkn *bindings.IERC20, addr common.Address, amount *big.Int) error {
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()

	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return errors.Wrap(ctx.Err(), "timeout")
		case <-ticker.C:
			balance, err := tkn.BalanceOf(&bind.CallOpts{Context: ctx}, addr)
			if err != nil {
				return errors.Wrap(err, "balance")
			}

			if umath.GTE(balance, amount) {
				return nil
			}
		}
	}
}
