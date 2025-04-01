package tokens

import (
	"math/big"

	e2e "github.com/omni-network/omni/e2e/solve"
	"github.com/omni-network/omni/lib/bi"
	"github.com/omni-network/omni/lib/contracts"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/evmchain"
	"github.com/omni-network/omni/lib/netconf"
	tokenslib "github.com/omni-network/omni/lib/tokens"

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

// Token represents a token (erc20 or native) on a specific chain.
type Token struct {
	tokenslib.Token
	ChainID    uint64
	ChainClass ChainClass
	Address    common.Address // empty if native
	MaxSpend   *big.Int
	MinSpend   *big.Int
	IsMock     bool
}

func (t Token) IsNative() bool {
	return t.Address == NativeAddr
}

func (t Token) IsOMNI() bool {
	return t.Token == tokenslib.OMNI
}

type SpendBounds struct {
	MinSpend *big.Int // minimum spend amount
	MaxSpend *big.Int // maximum spend amount
}

var (
	spendBounds = map[tokenslib.Token]map[ChainClass]SpendBounds{
		tokenslib.ETH: {
			ClassMainnet: {
				MinSpend: bi.Ether(0.001), // 0.001 ETH
				MaxSpend: bi.Ether(6),     // 6 ETH
			},
			ClassTestnet: {
				MinSpend: bi.Ether(0.001), // 0.001 ETH
				MaxSpend: bi.Ether(3),     // 3 ETH
			},
			ClassDevent: {
				MinSpend: bi.Ether(0.001), // 0.001 ETH
				MaxSpend: bi.Ether(3),     // 3 ETH
			},
		},
		tokenslib.OMNI: {
			ClassMainnet: {
				MinSpend: bi.Ether(0.1),     // 0.1 OMNI
				MaxSpend: bi.Ether(120_000), // 120k OMNI
			},
			ClassTestnet: {
				MinSpend: bi.Ether(0.1),   // 0.1 OMNI
				MaxSpend: bi.Ether(1_000), // 1k OMNI
			},
			ClassDevent: {
				MinSpend: bi.Ether(0.1),   // 0.1 OMNI
				MaxSpend: bi.Ether(1_000), // 1k OMNI
			},
		},
		tokenslib.WSTETH: {
			ClassMainnet: {
				MinSpend: bi.Ether(0.001), // 0.001 wstETH
				MaxSpend: bi.Ether(6),     // 6 wstETH
			},
			ClassTestnet: {
				MinSpend: bi.Ether(0.001), // 0.001 wstETH
				MaxSpend: bi.Ether(1),     // 1 wstETH
			},
			ClassDevent: {
				MinSpend: bi.Ether(0.001), // 0.001 wstETH
				MaxSpend: bi.Ether(1),     // 1 wstETH
			},
		},
		tokenslib.STETH: {
			ClassMainnet: {
				MinSpend: bi.Ether(0.001), // 0.001 stETH
				MaxSpend: bi.Ether(6),     // 6 stETH
			},
			ClassTestnet: {
				MinSpend: bi.Ether(0.001), // 0.001 stETH
				MaxSpend: bi.Ether(1),     // 1 stETH
			},
			ClassDevent: {
				MinSpend: bi.Ether(0.001), // 0.001 stETH
				MaxSpend: bi.Ether(1),     // 1 stETH
			},
		},
		tokenslib.USDC: {
			ClassMainnet: {
				MinSpend: bi.Dec6(0.1),    // 0.1 USDC
				MaxSpend: bi.Dec6(10_000), // 10k USDC
			},
			ClassTestnet: {
				MinSpend: bi.Dec6(0.1), // 0.1 USDC
				MaxSpend: bi.Dec6(10),  // 10 USDC
			},
			ClassDevent: {
				MinSpend: bi.Dec6(0.1), // 0.1 USDC
				MaxSpend: bi.Dec6(10),  // 10 USDC
			},
		},
	}
)

var tokens = append([]Token{
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

	// USDC (mainnet)
	usdc(evmchain.IDEthereum, common.HexToAddress("0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48")),
	usdc(evmchain.IDArbitrumOne, common.HexToAddress("0xaf88d065e77c8cC2239327C5EDb3A432268e5831")),
	usdc(evmchain.IDOptimism, common.HexToAddress("0x0b2c639c533813f4aa9d7837caf62653d097ff85")),
	usdc(evmchain.IDBase, common.HexToAddress("0x833589fCD6eDb6E08f4c7C32D4f71b54bdA02913")),

	// USDC (testnet)
	usdc(evmchain.IDSepolia, common.HexToAddress("0x1c7D4B196Cb0C7B01d743Fbc6116a902379C7238")),
	usdc(evmchain.IDArbSepolia, common.HexToAddress("0x75faf114eafb1BDbe2F0316DF893fd58CE46AA4d")),
	usdc(evmchain.IDOpSepolia, common.HexToAddress("0x5fd84259d66Cd46123540766Be93DFE6D43130D7")),
	usdc(evmchain.IDBaseSepolia, common.HexToAddress("0x036CbD53842c5426634e7929541eC2318f3dCF7e")),

	// ERC20 OMNI
	omniERC20(netconf.Mainnet),
	omniERC20(netconf.Omega),
	omniERC20(netconf.Staging),
	omniERC20(netconf.Devnet),

	// wstETH
	wstETH(evmchain.IDBase, common.HexToAddress("0xc1cba3fcea344f92d9239c08c0568f6f2f0ee452")),
	wstETH(evmchain.IDEthereum, common.HexToAddress("0x7f39c581f595b53c5cb19bd0b3f8da6c935e2ca0")),
	wstETH(evmchain.IDHolesky, common.HexToAddress("0x8d09a4502cc8cf1547ad300e066060d043f6982d")),
	wstETH(evmchain.IDSepolia, common.HexToAddress("0xB82381A3fBD3FaFA77B3a7bE693342618240067b")),
	// Mocks contain wstETH on IDBaseSepolia for omega and staging

	// stETH
	stETH(evmchain.IDEthereum, common.HexToAddress("0xae7ab96520de3a18e5e111b5eaab095312d7fe84")),
	stETH(evmchain.IDHolesky, common.HexToAddress("0x3f1c547b21f65e10480de3ad8e19faac46c95034")),
}, mocks()...)

func All() []Token {
	return tokens
}

// UniqueSymbols returns the unique set of tokenslib.Tokens in the tokens list.
func UniqueSymbols() []tokenslib.Token {
	uniq := make(map[tokenslib.Token]bool)

	for _, t := range tokens {
		uniq[t.Token] = true
	}

	var resp []tokenslib.Token
	for t := range uniq {
		resp = append(resp, t)
	}

	return resp
}

func BySymbol(chainID uint64, symbol string) (Token, bool) {
	for _, t := range tokens {
		if t.ChainID == chainID && t.Symbol == symbol {
			return t, true
		}
	}

	return Token{}, false
}

// Native is an alias for ByAddress with the native token address.
func Native(chainID uint64) (Token, bool) {
	return ByAddress(chainID, NativeAddr)
}

func ByAddress(chainID uint64, addr common.Address) (Token, bool) {
	for _, t := range tokens {
		if t.ChainID == chainID && t.Address == addr {
			return t, true
		}
	}

	return Token{}, false
}

func ByChain(chainID uint64) []Token {
	var tkns []Token

	for _, t := range tokens {
		if t.ChainID == chainID {
			tkns = append(tkns, t)
		}
	}

	return tkns
}

func nativeETH(chainID uint64) Token {
	chainClass := mustChainClass(chainID)
	bounds := mustSpendBounds(tokenslib.ETH, chainClass)

	return Token{
		Token:      tokenslib.ETH,
		ChainID:    chainID,
		ChainClass: mustChainClass(chainID),
		MaxSpend:   bounds.MaxSpend,
		MinSpend:   bounds.MinSpend,
		Address:    NativeAddr,
	}
}

func nativeOMNI(chainID uint64) Token {
	chainClass := mustChainClass(chainID)
	bounds := mustSpendBounds(tokenslib.OMNI, chainClass)

	return Token{
		Token:      tokenslib.OMNI,
		ChainID:    chainID,
		ChainClass: mustChainClass(chainID),
		MaxSpend:   bounds.MaxSpend,
		MinSpend:   bounds.MinSpend,
		Address:    NativeAddr,
	}
}

func omniERC20(network netconf.ID) Token {
	chainID := netconf.EthereumChainID(network)
	chainClass := mustChainClass(chainID)
	bounds := mustSpendBounds(tokenslib.OMNI, chainClass)

	return Token{
		Token:      tokenslib.OMNI,
		ChainID:    chainID,
		ChainClass: mustChainClass(chainID),
		Address:    contracts.TokenAddr(network),
		MaxSpend:   bounds.MaxSpend,
		MinSpend:   bounds.MinSpend,
	}
}

// mockOMNI returns a manually deployed OMNI token on a given chain for testing purposes.
func mockOMNI(chainID uint64, addr common.Address) Token {
	chainClass := mustChainClass(chainID)
	bounds := mustSpendBounds(tokenslib.OMNI, chainClass)

	return Token{
		Token:      tokenslib.OMNI,
		ChainID:    chainID,
		ChainClass: mustChainClass(chainID),
		Address:    addr,
		MaxSpend:   bounds.MaxSpend,
		MinSpend:   bounds.MinSpend,
		IsMock:     true,
	}
}

func stETH(chainID uint64, addr common.Address) Token {
	chainClass := mustChainClass(chainID)
	bounds := mustSpendBounds(tokenslib.STETH, chainClass)

	return Token{
		Token:      tokenslib.STETH,
		ChainID:    chainID,
		ChainClass: mustChainClass(chainID),
		MaxSpend:   bounds.MaxSpend,
		MinSpend:   bounds.MinSpend,
		Address:    addr,
	}
}

func wstETH(chainID uint64, addr common.Address) Token {
	chainClass := mustChainClass(chainID)
	bounds := mustSpendBounds(tokenslib.WSTETH, chainClass)

	return Token{
		Token:      tokenslib.WSTETH,
		ChainID:    chainID,
		ChainClass: mustChainClass(chainID),
		Address:    addr,
		MaxSpend:   bounds.MaxSpend,
		MinSpend:   bounds.MinSpend,
	}
}

func usdc(chainID uint64, addr common.Address) Token {
	chainClass := mustChainClass(chainID)
	bounds := mustSpendBounds(tokenslib.USDC, chainClass)

	return Token{
		Token:      tokenslib.USDC,
		ChainID:    chainID,
		ChainClass: mustChainClass(chainID),
		Address:    addr,
		MaxSpend:   bounds.MaxSpend,
		MinSpend:   bounds.MinSpend,
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

func mustChainClass(chainID uint64) ChainClass {
	class, err := chainClass(chainID)
	if err != nil {
		panic(err)
	}

	return class
}

func mustSpendBounds(tkn tokenslib.Token, chainClass ChainClass) SpendBounds {
	bounds, ok := spendBounds[tkn][chainClass]
	if !ok {
		panic(errors.New("spend bounds not found", "token", tkn, "chain_class", chainClass))
	}

	return bounds
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
