package app

import (
	"context"
	"math/big"

	"github.com/omni-network/omni/contracts/bindings"
	e2e "github.com/omni-network/omni/e2e/solve"
	"github.com/omni-network/omni/lib/contracts"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient/ethbackend"
	"github.com/omni-network/omni/lib/evmchain"
	"github.com/omni-network/omni/lib/netconf"
	tokenslib "github.com/omni-network/omni/lib/tokens"
	"github.com/omni-network/omni/lib/umath"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

// NativeAddr is the "address" of the native token; the zero address.
var NativeAddr common.Address

type ChainClass string

const (
	ClassDevent  ChainClass = "devnet"
	ClassTestnet ChainClass = "testnet"
	ClassMainnet ChainClass = "mainnet"
)

type Token struct {
	tokenslib.Token
	ChainID    uint64
	ChainClass ChainClass
	Address    common.Address // empty if native
	MaxSpend   *big.Int
	MinSpend   *big.Int
	IsMock     bool
}

// TokenAmt represents a token and an amount.
// It differs from types.AddrAmt in that it contains fully resolved token type, not just the token address.
type TokenAmt struct {
	Token  Token
	Amount *big.Int
}

type Tokens []Token

func (t Token) IsNative() bool {
	return t.Address == NativeAddr
}

func (t Token) IsOMNI() bool {
	return t.Token == tokenslib.OMNI
}

var (
	// TODO: increase max spend on mainnet, keep low for staging / omega.
	maxETHSpend    = umath.Ether(1)     // 1 ETH
	minETHSpend    = umath.Ether(0.001) // 0.001 ETH
	maxWSTETHSpend = umath.Ether(1)     // 1 wstETH
	minWSTETHSpend = umath.Ether(0.001) // 0.001 wstETH
	maxOMNISpend   = umath.Ether(1_000) // 1000 OMNI
	minOMNISpend   = umath.Ether(0.1)   // 0.1 OMNI
)

var tokens = append(Tokens{
	// Native ETH (mainnet)
	nativeETH(evmchain.IDEthereum),
	nativeETH(evmchain.IDArbitrumOne),
	nativeETH(evmchain.IDBase),
	nativeETH(evmchain.IDOptimism),

	// Native ETH (testnet)
	nativeETH(evmchain.IDHolesky),
	nativeETH(evmchain.IDArbSepolia),
	nativeETH(evmchain.IDBaseSepolia),
	nativeETH(evmchain.IDOpSepolia),

	// Native ETH (devnet)
	nativeETH(evmchain.IDMockL1),
	nativeETH(evmchain.IDMockL2),

	// Native OMNI
	nativeOMNI(evmchain.IDOmniMainnet),
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
	wstETH(evmchain.IDSepolia, common.HexToAddress("0xB82381A3fBD3FaFA77B3a7bE693342618240067b")),

	// stETH
	stETH(evmchain.IDHolesky, common.HexToAddress("0x3f1c547b21f65e10480de3ad8e19faac46c95034")),

	// mock l1 copies (for e2e fork testing)
	wstETH(evmchain.IDMockL1, common.HexToAddress("0x8d09a4502cc8cf1547ad300e066060d043f6982d")), // holesky wstETH
	wstETH(evmchain.IDMockL1, common.HexToAddress("0xB82381A3fBD3FaFA77B3a7bE693342618240067b")), // sepolia wstETH
	stETH(evmchain.IDMockL1, common.HexToAddress("0x3f1c547b21f65e10480de3ad8e19faac46c95034")),  // holesky stETH
}, mocks()...)

func AllTokens() Tokens {
	return tokens
}

func (ts Tokens) Find(chainID uint64, addr common.Address) (Token, bool) {
	for _, t := range ts {
		if t.ChainID == chainID && t.Address == addr {
			return t, true
		}
	}

	return Token{}, false
}

func (ts Tokens) ForChain(chainID uint64) Tokens {
	var tkns Tokens

	for _, t := range ts {
		if t.ChainID == chainID {
			tkns = append(tkns, t)
		}
	}

	return tkns
}

func nativeETH(chainID uint64) Token {
	return Token{
		Token:      tokenslib.ETH,
		ChainID:    chainID,
		ChainClass: mustChainClass(chainID),
		MaxSpend:   maxETHSpend,
		MinSpend:   minETHSpend,
		Address:    NativeAddr,
	}
}

func nativeOMNI(chainID uint64) Token {
	return Token{
		Token:      tokenslib.OMNI,
		ChainID:    chainID,
		ChainClass: mustChainClass(chainID),
		MaxSpend:   maxOMNISpend,
		MinSpend:   minOMNISpend,
		Address:    NativeAddr,
	}
}

func omniERC20(network netconf.ID) Token {
	chainID := netconf.EthereumChainID(network)

	return Token{
		Token:      tokenslib.OMNI,
		ChainID:    chainID,
		ChainClass: mustChainClass(chainID),
		Address:    contracts.TokenAddr(network),
		MaxSpend:   maxOMNISpend,
		MinSpend:   minOMNISpend,
	}
}

// mockOMNI returns a manually deployed OMNI token on a given chain for testing purposes.
func mockOMNI(chainID uint64, addr common.Address) Token {
	return Token{
		Token:      tokenslib.OMNI,
		ChainID:    chainID,
		ChainClass: mustChainClass(chainID),
		Address:    addr,
		MaxSpend:   maxOMNISpend,
		MinSpend:   minOMNISpend,
		IsMock:     true,
	}
}

func stETH(chainID uint64, addr common.Address) Token {
	return Token{
		Token:      tokenslib.STETH,
		ChainID:    chainID,
		ChainClass: mustChainClass(chainID),
		Address:    addr,
	}
}

func wstETH(chainID uint64, addr common.Address) Token {
	return Token{
		Token:      tokenslib.WSTETH,
		ChainID:    chainID,
		ChainClass: mustChainClass(chainID),
		Address:    addr,
		MaxSpend:   maxWSTETHSpend,
		MinSpend:   minWSTETHSpend,
	}
}

// mocks returns MockTokens deployed in e2e for testing purposes.
func mocks() []Token {
	var tkns []Token

	for _, mock := range e2e.MockTokens() {
		tkns = append(tkns, Token{
			Token:      mock.Token,
			Address:    mock.Address(),
			ChainID:    mock.ChainID,
			ChainClass: mustChainClass(mock.ChainID),
			IsMock:     true,
		})
	}

	// Add manually deployed tokens that aren't part of the automatic mock deployment
	tkns = append(tkns,
		mockOMNI(evmchain.IDBaseSepolia, common.HexToAddress("0xe4075366F03C286846dECE8AAF104cF2a542294d")),
		mockOMNI(evmchain.IDOpSepolia, common.HexToAddress("0x0b3AED256a51919b660fF79a280A309EecA9d688")),
		mockOMNI(evmchain.IDArbSepolia, common.HexToAddress("0xd859f9Ff3C9700fB623Dc8e76217ba2a9f8613F0")),
	)

	return tkns
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

func mustChainClass(chainID uint64) ChainClass {
	class, err := chainClass(chainID)
	if err != nil {
		panic(err)
	}

	return class
}

func chainClass(chainID uint64) (ChainClass, error) {
	switch chainID {
	case
		evmchain.IDOmniMainnet,
		evmchain.IDEthereum,
		evmchain.IDArbitrumOne,
		evmchain.IDBase,
		evmchain.IDOptimism:
		return ClassMainnet, nil
	case
		evmchain.IDOmniOmega,
		evmchain.IDOmniStaging, // classify omni staging as testnet, because it interops with other testnets
		evmchain.IDHolesky,
		evmchain.IDSepolia,
		evmchain.IDArbSepolia,
		evmchain.IDBaseSepolia,
		evmchain.IDOpSepolia:
		return ClassTestnet, nil
	case
		evmchain.IDOmniDevnet,
		evmchain.IDMockL1,
		evmchain.IDMockL2,
		evmchain.IDMockArb,
		evmchain.IDMockOp:
		return ClassDevent, nil
	default:
		return "", errors.New("unsupported chain ID", "chain_id", chainID)
	}
}
