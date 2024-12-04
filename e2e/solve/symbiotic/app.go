package symbiotic

import (
	"context"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/e2e/app/eoa"
	"github.com/omni-network/omni/e2e/solve/devapp"
	"github.com/omni-network/omni/lib/cast"
	"github.com/omni-network/omni/lib/contracts"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient/ethbackend"
	"github.com/omni-network/omni/lib/evmchain"
	"github.com/omni-network/omni/lib/netconf"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"

	_ "embed"
)

type App struct {
	L1wstETHCollateral common.Address
	L1wstETH           common.Address
	L2wstETH           common.Address
	L1                 evmchain.Metadata
	L2                 evmchain.Metadata
}

var (
	//go:embed default-collateral-abi.json
	collateralABIJSON []byte

	collateralABI = mustParseABI(collateralABIJSON)
	depositABI    = mustGetMethod(collateralABI, "deposit")

	mockL1      = mustChainMeta(evmchain.IDMockL1) // should be forked from holesky
	mockL2      = mustChainMeta(evmchain.IDMockL2)
	holesky     = mustChainMeta(evmchain.IDHolesky)
	baseSepolia = mustChainMeta(evmchain.IDBaseSepolia)
)

func GetApp(network netconf.ID) (App, error) {
	app := App{
		// holesky wsETH collateral
		L1wstETHCollateral: common.HexToAddress("0x23e98253f372ee29910e22986fe75bb287b011fc"),
		// holesky wsETH
		L1wstETH: common.HexToAddress("0x8d09a4502cc8cf1547ad300e066060d043f6982d"),
		// use mintable devapp mintable mock token, base sepolia has no canonical wsETH
		L2wstETH: devapp.MustGetApp(network).L2Token,
	}

	if network == netconf.Devnet {
		app.L1 = mockL1
		app.L2 = mockL2

		return app, nil
	}

	if network == netconf.Staging {
		app.L1 = holesky
		app.L2 = baseSepolia

		return app, nil
	}

	return App{}, errors.New("unsupported network", "network", network)
}

func MustGetApp(network netconf.ID) App {
	app, err := GetApp(network)
	if err != nil {
		panic(err)
	}

	return app
}

// AllowOutboxCalls allows the outbox to call the L1 wstETH collateral contract.
func AllowOutboxCalls(ctx context.Context, network netconf.ID, backends ethbackend.Backends) error {
	if !network.IsEphemeral() {
		return errors.New("only ephemeral networks")
	}

	app, err := GetApp(network)
	if err != nil {
		return errors.Wrap(err, "get app")
	}

	addrs, err := contracts.GetAddresses(ctx, network)
	if err != nil {
		return errors.Wrap(err, "get addresses")
	}

	l1Backend, err := backends.Backend(app.L1.ChainID)
	if err != nil {
		return errors.Wrap(err, "backend mock l1")
	}

	if err := allowCalls(ctx, app, l1Backend, addrs.SolveOutbox); err != nil {
		return errors.Wrap(err, "allow calls")
	}

	return nil
}

// allowCalls allows the outbox to call the L1 wstETH collateral deposit method.
func allowCalls(ctx context.Context, app App, backend *ethbackend.Backend, outboxAddr common.Address) error {
	outbox, err := bindings.NewSolveOutbox(outboxAddr, backend)
	if err != nil {
		return errors.Wrap(err, "new solve outbox")
	}

	manager := eoa.MustAddress(netconf.Devnet, eoa.RoleManager)

	txOpts, err := backend.BindOpts(ctx, manager)
	if err != nil {
		return errors.Wrap(err, "bind opts")
	}

	depositID, err := cast.Array4(depositABI.ID[:4])
	if err != nil {
		return err
	}

	tx, err := outbox.SetAllowedCall(txOpts, app.L1wstETHCollateral, depositID, true)
	if err != nil {
		return errors.Wrap(err, "set allowed call")
	} else if _, err := backend.WaitMined(ctx, tx); err != nil {
		return errors.Wrap(err, "wait mined")
	}

	return nil
}

func mustParseABI(json []byte) *abi.ABI {
	var abi abi.ABI
	if err := abi.UnmarshalJSON(json); err != nil {
		panic(err)
	}

	return &abi
}

func mustGetMethod(abi *abi.ABI, name string) abi.Method {
	method, ok := abi.Methods[name]
	if !ok {
		panic(errors.New("missing method", "name", name))
	}

	return method
}

func mustChainMeta(chainID uint64) evmchain.Metadata {
	meta, ok := evmchain.MetadataByID(chainID)
	if !ok {
		panic(errors.New("missing chain meta", "chain_id", chainID))
	}

	return meta
}
