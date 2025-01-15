package appv2

import (
	"context"
	"math/big"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/lib/contracts"
	"github.com/omni-network/omni/lib/ethclient/ethbackend"
	"github.com/omni-network/omni/lib/evmchain"
	"github.com/omni-network/omni/lib/netconf"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

// TODO: consider moving this to lib/tokens

type Token struct {
	symbol      string
	name        string
	chainID     uint64
	address     common.Address // empty if native
	coingeckoID string
}

type Expense struct {
	token  Token
	amount *big.Int
}

type Tokens []Token

func (t Token) isNative() bool {
	return t.address == common.Address{}
}

var tokens = Tokens{
	nativeETH(evmchain.IDHolesky),
	nativeETH(evmchain.IDArbSepolia),
	nativeETH(evmchain.IDBaseSepolia),
	nativeETH(evmchain.IDOpSepolia),
	nativeETH(evmchain.IDMockL1),
	nativeETH(evmchain.IDMockL2),

	nativeOMNI(evmchain.IDOmniOmega),
	nativeOMNI(evmchain.IDOmniStaging),
	nativeOMNI(evmchain.IDOmniDevnet),

	omniERC20(netconf.Mainnet),
	omniERC20(netconf.Omega),
	omniERC20(netconf.Staging),
	omniERC20(netconf.Devnet),
}

func (ts Tokens) find(chainID uint64, addr common.Address) (Token, bool) {
	for _, t := range ts {
		if t.chainID == chainID && t.address == addr {
			return t, true
		}
	}

	return Token{}, false
}

func nativeETH(chainID uint64) Token {
	return Token{
		symbol:      "ETH",
		name:        "Ether",
		chainID:     chainID,
		coingeckoID: "ethereum",
	}
}

func nativeOMNI(chainID uint64) Token {
	return Token{
		symbol:      "OMNI",
		name:        "Omni Network",
		chainID:     chainID,
		coingeckoID: "omni-network",
	}
}

func omniERC20(network netconf.ID) Token {
	return Token{
		symbol:      "OMNI",
		name:        "Omni Network",
		chainID:     netconf.EthereumChainID(network),
		address:     contracts.TokenAddr(network),
		coingeckoID: "omni-network",
	}
}

func (t Token) balanceOf(
	ctx context.Context,
	backend *ethbackend.Backend,
	addr common.Address,
) (*big.Int, error) {
	switch {
	case t.isNative():
		return backend.BalanceAt(ctx, addr, nil)
	default:
		contract, err := bindings.NewIERC20(t.address, backend)
		if err != nil {
			return nil, err
		}

		return contract.BalanceOf(&bind.CallOpts{Context: ctx}, addr)
	}
}
