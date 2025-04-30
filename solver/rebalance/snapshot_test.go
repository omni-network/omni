package rebalance_test

import (
	"flag"
	"log/slog"
	"math/big"
	"testing"

	"github.com/omni-network/omni/e2e/app/eoa"
	"github.com/omni-network/omni/lib/bi"
	"github.com/omni-network/omni/lib/ethclient"
	"github.com/omni-network/omni/lib/evmchain"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/tokens"
	"github.com/omni-network/omni/lib/tokens/tokenutil"
	"github.com/omni-network/omni/lib/tutil"
	solver "github.com/omni-network/omni/solver/app"
	"github.com/omni-network/omni/solver/rebalance"
)

var (
	snapshot = flag.Bool("snapshot", false, "run snapshot test")
)

type tokenAmt struct {
	token tokens.Token
	amt   *big.Int
}

// Usage: go test . -snapshot -v -run=TestSnapshot.
func TestSnapshot(t *testing.T) {
	t.Parallel()

	if !*snapshot {
		t.Skip("skipping snapshot test")
	}

	ctx := t.Context()

	logCfg := log.DefaultConfig()
	logCfg.Level = slog.LevelDebug.String()
	logCfg.Color = log.ColorForce

	ctx, err := log.Init(ctx, logCfg)
	tutil.RequireNoError(t, err)

	rpcs := getRPCs(t)

	clients := make(map[uint64]ethclient.Client)
	for chainID, rpc := range rpcs {
		client, err := ethclient.Dial(evmchain.Name(chainID), rpc)
		tutil.RequireNoError(t, err)
		clients[chainID] = client
	}

	solverAddr := eoa.MustAddress(netconf.Mainnet, eoa.RoleSolver)

	var deficits []tokenAmt
	var surpluses []tokenAmt

	// Log all token balances w/ surplus/deficit
	for _, token := range tokens.All() {
		if !solver.IsSupportedToken(token) {
			continue
		}

		if token.ChainClass != tokens.ClassMainnet {
			continue
		}

		if !rebalance.CanRebalance(token.ChainID) {
			continue
		}

		balance, err := tokenutil.BalanceOf(ctx, clients[token.ChainID], token, solverAddr)
		tutil.RequireNoError(t, err)

		surplus, err := rebalance.GetSurplus(ctx, clients[token.ChainID], token, solverAddr)
		tutil.RequireNoError(t, err)

		deficit, err := rebalance.GetDeficit(ctx, clients[token.ChainID], token, solverAddr)
		tutil.RequireNoError(t, err)

		log.Info(ctx, "Token balance",
			"chain", evmchain.Name(token.ChainID),
			"token", token.Asset,
			"balance", token.FormatAmt(balance),
			"surplus", token.FormatAmt(surplus),
			"deficit", token.FormatAmt(deficit))

		if !bi.IsZero(surplus) {
			surpluses = append(surpluses, tokenAmt{token: token, amt: surplus})
		}

		if !bi.IsZero(deficit) {
			deficits = append(deficits, tokenAmt{token: token, amt: deficit})
		}
	}

	// Log deficits
	log.Info(ctx, "------------------")
	for _, d := range deficits {
		log.Info(ctx, "Deficit",
			"chain", evmchain.Name(d.token.ChainID),
			"token", d.token.Asset,
			"deficit", d.token.FormatAmt(d.amt),
			"target", d.token.FormatAmt(rebalance.GetFundThreshold(d.token).Target()))
	}

	// Log surpluses
	log.Info(ctx, "------------------")
	for _, s := range surpluses {
		log.Info(ctx, "Surplus",
			"chain", evmchain.Name(s.token.ChainID),
			"token", s.token.Asset,
			"surplus", s.token.FormatAmt(s.amt),
			"surplus_threshold", s.token.FormatAmt(rebalance.GetFundThreshold(s.token).Surplus()))
	}
}
