package app

import (
	"context"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/lib/bi"
	"github.com/omni-network/omni/lib/contracts"
	"github.com/omni-network/omni/lib/contracts/omnitoken"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/netconf"

	"github.com/ethereum/go-ethereum/common"
)

type RunERC20FaucetConfig struct {
	AddrToFund string // Hex address to fund.
	Amount     uint64 // Amount to fund.
}

func DefaultRunERC20FaucetConfig() RunERC20FaucetConfig {
	return RunERC20FaucetConfig{
		AddrToFund: "",
		Amount:     100,
	}
}

// RunERC20Faucet runs the ERC20 faucet, funding an address from the initial supply recipient.
func RunERC20Faucet(ctx context.Context, def Definition, cfg RunERC20FaucetConfig) error {
	if def.Testnet.Network == netconf.Mainnet {
		return errors.New("no mainnet faucet")
	}

	if !common.IsHexAddress(cfg.AddrToFund) {
		return errors.New("not a hex address", "addr", cfg.AddrToFund)
	}

	networkID := def.Testnet.Network
	addrs, err := contracts.GetAddresses(ctx, networkID)
	if err != nil {
		return errors.Wrap(err, "get addrs")
	}

	account := common.HexToAddress(cfg.AddrToFund)
	amt := bi.Ether(cfg.Amount)

	funder, err := omnitoken.InitialSupplyRecipient(networkID)
	if err != nil {
		return errors.Wrap(err, "initial supply recipient")
	}

	chain, ok := def.Testnet.EthereumChain()
	if !ok {
		return errors.New("no ethereum chain")
	}

	backend, err := def.Backends().Backend(chain.ChainID)
	if err != nil {
		return errors.Wrap(err, "backend")
	}

	token, err := bindings.NewNomina(addrs.NomToken, backend)
	if err != nil {
		return errors.Wrap(err, "new nomina")
	}

	txOpts, err := backend.BindOpts(ctx, funder)
	if err != nil {
		return errors.Wrap(err, "bind opts")
	}

	tx, err := token.Transfer(txOpts, account, amt)
	if err != nil {
		return errors.Wrap(err, "transfer")
	}

	rec, err := backend.WaitMined(ctx, tx)
	if err != nil {
		return errors.Wrap(err, "wait mined")
	}

	log.Info(ctx, "Funded", "addr", account.Hex(), "token", addrs.NomToken, "amount", cfg.Amount, "tx", rec.TxHash.Hex())

	return nil
}
