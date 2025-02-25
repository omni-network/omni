package bridging

import (
	"context"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/e2e/app/eoa"
	"github.com/omni-network/omni/lib/contracts/solvernet"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient/ethbackend"
	"github.com/omni-network/omni/lib/evmchain"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/tokens"
	"github.com/omni-network/omni/monitor/flowgen/types"
	"github.com/omni-network/omni/monitor/flowgen/util"
	"github.com/omni-network/omni/solver/app"
)

type job struct {
	name    string
	cadence time.Duration
}

func NewJob() types.Job {
	return job{
		name:    "Bridging",
		cadence: 1 * time.Second,
	}
}

func (j job) Name() string {
	return j.name
}

func (j job) Cadence() time.Duration {
	return j.cadence
}

func (j job) Run(ctx context.Context, backends ethbackend.Backends) error {
	token, found := app.AllTokens().Find(evmchain.IDBaseSepolia, common.HexToAddress("0x6319df7c227e34B967C1903A08a698A3cC43492B"))
	if !found {
		return errors.New("unknown token", "token", tokens.ETH)
	}
	deposit := app.Payment{Token: token, Amount: util.MilGwei}

	expense, err := app.QuoteExpense(deposit.Token, deposit)
	if err != nil {
		return errors.Wrap(err, "quote expense")
	}
	user := eoa.MustAddress(netconf.Omega, eoa.RoleTester)

	log.Info(ctx, "Starting symbiotic flow", "deposit", deposit, "expense", expense, "user", user)

	symbioticContractOnHolesky := common.HexToAddress("0x23E98253F372Ee29910e22986fe75Bb287b011fC")

	orderData := bindings.SolverNetOrderData{
		Owner:       user,
		DestChainId: evmchain.IDHolesky,
		Deposit: solvernet.Deposit{
			Token:  deposit.Token.Address,
			Amount: deposit.Amount,
		},
		Expenses: []bindings.SolverNetTokenExpense{{
			Spender: symbioticContractOnHolesky,
			Token:   expense.Token.Address,
			Amount:  expense.Amount,
		}},
		Calls: solvernet.Calls{
			{
				Target: symbioticContractOnHolesky,
				Value:  nil,
				Data:   nil, // abi.encode("deposit(address, uint256)", user, deposit),
			},
		}.ToBindings(),
	}

	solvernet.OpenOrder(ctx, netconf.Staging, evmchain.IDHolesky, backends, user, orderData)

	return nil
}
