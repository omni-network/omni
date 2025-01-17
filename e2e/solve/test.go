package solve

import (
	"context"
	"fmt"
	"math/big"
	"strings"
	"time"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/lib/anvil"
	"github.com/omni-network/omni/lib/contracts"
	"github.com/omni-network/omni/lib/contracts/solvernet"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient/ethbackend"
	"github.com/omni-network/omni/lib/evmchain"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/xchain"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/params"

	"cosmossdk.io/math"
	"golang.org/x/sync/errgroup"
)

type BridgeOrder struct {
	From   common.Address
	To     common.Address
	Amount *big.Int
}

func TestV2(ctx context.Context, network netconf.Network, endpoints xchain.RPCEndpoints) error {
	if network.ID != netconf.Devnet {
		return errors.New("only devnet")
	}

	log.Info(ctx, "Running solver v2 test")

	// use anvil.DevAccounts instead of eoa.DevAccounts, because eoa.DevAccounts
	// are used frequently elsewhere in e2e / e2e tests, and nonce issues get annoying
	backends, err := ethbackend.BackendsFromNetwork(network, endpoints, anvil.DevPrivateKeys()...)
	if err != nil {
		return err
	}

	orders := makeOrders()

	err = mintAndApproveAll(ctx, backends, orders)
	if err != nil {
		return errors.Wrap(err, "mint omni")
	}

	if err := bridgeToNativeAll(ctx, backends, orders); err != nil {
		return errors.Wrap(err, "bridge to native")
	}

	if err := waitNativeAll(ctx, backends, orders); err != nil {
		return errors.Wrap(err, "wait native")
	}

	log.Info(ctx, "Solver v2 test success")

	return nil
}

func makeOrders() []BridgeOrder {
	users := anvil.DevAccounts()
	amt := math.NewInt(10).MulRaw(params.Ether).BigInt()
	orders := make([]BridgeOrder, len(users))

	for i, user := range users {
		// use 0xdead0000ii... to get unique, unused To address per bridge order
		// we bridge to unused addresses to simplify balance checks in waitNativeAll
		const prefix = "0xdead0000"
		to := common.HexToAddress(fmt.Sprintf("%s%s", prefix, strings.Repeat(fmt.Sprintf("%d", i), 42-len(prefix))))

		orders[i] = BridgeOrder{
			From:   user,
			To:     to,
			Amount: amt,
		}
	}

	return orders
}

func waitNativeAll(ctx context.Context, backends ethbackend.Backends, orders []BridgeOrder) error {
	ctx, cancel := context.WithTimeout(ctx, 2*time.Minute)
	defer cancel()

	backend, err := backends.Backend(evmchain.IDOmniDevnet)
	if err != nil {
		return errors.Wrap(err, "backend")
	}

	ticker := time.NewTicker(2 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return errors.Wrap(ctx.Err(), "timeout")
		case <-ticker.C:
			bridged := 0

			for _, order := range orders {
				balance, err := backend.BalanceAt(ctx, order.To, nil)
				if err != nil {
					return errors.Wrap(err, "balance of")
				}

				if balance.Cmp(order.Amount) == 0 {
					bridged++
				}
			}

			if bridged == len(orders) {
				log.Debug(ctx, "All native bridges complete")
				return nil
			}
		}
	}
}

func bridgeToNativeAll(ctx context.Context, backends ethbackend.Backends, orders []BridgeOrder) error {
	var eg errgroup.Group
	for _, order := range orders {
		eg.Go(func() error { return bridgeToNative(ctx, backends, order) })
	}

	if err := eg.Wait(); err != nil {
		return errors.Wrap(err, "wait group")
	}

	return nil
}

func bridgeToNative(ctx context.Context, backends ethbackend.Backends, order BridgeOrder) error {
	log.Debug(ctx, "Requesting native solvernet bridge", "user", order.From.Hex(), "amt", order.Amount.Uint64())

	// bridge to native requests a user.call{value: amt} on omni, while depositing amt ERC20 omni devnet l1
	return solvernet.OpenOrder(ctx, netconf.Devnet, evmchain.IDMockL1, backends, order.From, bindings.ISolverNetOrderData{
		Call: bindings.ISolverNetCall{
			ChainId:  evmchain.IDOmniDevnet,
			Value:    order.Amount,
			Target:   toBz32(order.To),
			Data:     nil,
			Expenses: nil,
		},
		Deposits: []bindings.ISolverNetDeposit{{
			Token:  toBz32(contracts.TokenAddr(netconf.Devnet)),
			Amount: order.Amount,
		}},
	})
}

func mintAndApproveAll(ctx context.Context, backends ethbackend.Backends, orders []BridgeOrder) error {
	backend, err := backends.Backend(evmchain.IDMockL1)
	if err != nil {
		return errors.Wrap(err, "backend")
	}

	var eg errgroup.Group
	for _, order := range orders {
		eg.Go(func() error { return mintAndApprove(ctx, backend, order) })
	}

	if err := eg.Wait(); err != nil {
		return errors.Wrap(err, "wait group")
	}

	return nil
}

func mintAndApprove(ctx context.Context, backend *ethbackend.Backend, order BridgeOrder) error {
	addrs, err := contracts.GetAddresses(ctx, netconf.Devnet)
	if err != nil {
		return errors.Wrap(err, "get addrs")
	}

	txOpts, err := backend.BindOpts(ctx, order.From)
	if err != nil {
		return errors.Wrap(err, "bind opts")
	}

	contract, err := bindings.NewMockERC20(addrs.Token, backend)
	if err != nil {
		return errors.Wrap(err, "bind contract")
	}

	tx, err := contract.Mint(txOpts, order.From, order.Amount)
	if err != nil {
		return errors.Wrap(err, "mint tx")
	}

	_, err = backend.WaitMined(ctx, tx)
	if err != nil {
		return errors.Wrap(err, "wait mined")
	}

	tx, err = contract.Approve(txOpts, addrs.SolverNetInbox, order.Amount)
	if err != nil {
		return errors.Wrap(err, "mint tx")
	}

	_, err = backend.WaitMined(ctx, tx)
	if err != nil {
		return errors.Wrap(err, "wait mined")
	}

	return nil
}

func toBz32(addr common.Address) [32]byte {
	var bz [32]byte
	copy(bz[12:], addr.Bytes())

	return bz
}
