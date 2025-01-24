package appv2

import (
	"context"
	"math/big"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/lib/contracts"
	"github.com/omni-network/omni/lib/ethclient/ethbackend"
	"github.com/omni-network/omni/lib/evmchain"
	"github.com/omni-network/omni/lib/netconf"
	tokenslib "github.com/omni-network/omni/lib/tokens"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

type Token struct {
	tokenslib.Token
	ChainID uint64
	Address common.Address // empty if native
}

type Expense struct {
	token  Token
	amount *big.Int
}

type Tokens []Token

func (t Token) IsNative() bool {
	return t.Address == common.Address{}
}

var tokens = Tokens{
	// Native ETH
	nativeETH(evmchain.IDHolesky),
	nativeETH(evmchain.IDArbSepolia),
	nativeETH(evmchain.IDBaseSepolia),
	nativeETH(evmchain.IDOpSepolia),
	nativeETH(evmchain.IDMockL1),
	nativeETH(evmchain.IDMockL2),

	// Native OMNI
	nativeOMNI(evmchain.IDOmniOmega),
	nativeOMNI(evmchain.IDOmniStaging),
	nativeOMNI(evmchain.IDOmniDevnet),

	// ERC20 OMNI
	omniERC20(netconf.Mainnet),
	omniERC20(netconf.Omega),
	omniERC20(netconf.Staging),
	omniERC20(netconf.Devnet),

	// wtSETH
	wstETH(evmchain.IDHolesky, common.HexToAddress("0x8d09a4502cc8cf1547ad300e066060d043f6982d")),

	// stETH
	stETH(evmchain.IDHolesky, common.HexToAddress("0x3f1c547b21f65e10480de3ad8e19faac46c95034")),
}

func (ts Tokens) find(chainID uint64, addr common.Address) (Token, bool) {
	for _, t := range ts {
		if t.ChainID == chainID && t.Address == addr {
			return t, true
		}
	}

	return Token{}, false
}

func nativeETH(chainID uint64) Token {
	return Token{
		Token:   tokenslib.ETH,
		ChainID: chainID,
	}
}

func nativeOMNI(chainID uint64) Token {
	return Token{
		Token:   tokenslib.OMNI,
		ChainID: chainID,
	}
}

func omniERC20(network netconf.ID) Token {
	return Token{
		Token:   tokenslib.OMNI,
		ChainID: netconf.EthereumChainID(network),
		Address: contracts.TokenAddr(network),
	}
}

func stETH(chainID uint64, addr common.Address) Token {
	return Token{
		Token:   tokenslib.STETH,
		ChainID: chainID,
		Address: addr,
	}
}

func wstETH(chainID uint64, addr common.Address) Token {
	return Token{
		Token:   tokenslib.WSTETH,
		ChainID: chainID,
		Address: addr,
	}
}

func balanceOf(
	ctx context.Context,
	tkn Token,
	backend *ethbackend.Backend,
	addr common.Address,
) (*big.Int, error) {
	switch {
	case tkn.IsNative():
		return backend.BalanceAt(ctx, addr, nil)
	default:
		contract, err := bindings.NewIERC20(tkn.Address, backend)
		if err != nil {
			return nil, err
		}

		return contract.BalanceOf(&bind.CallOpts{Context: ctx}, addr)
	}
}
