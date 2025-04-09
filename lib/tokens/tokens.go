package tokens

import (
	_ "embed"
	"encoding/json"
	"github.com/ethereum/go-ethereum/common"
	"github.com/omni-network/omni/lib/contracts"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/evmchain"
	"github.com/omni-network/omni/lib/netconf"
)

// NativeAddr is the "address" of the native token; the zero address.
var NativeAddr common.Address

type ChainClass string

const (
	ClassDevent  ChainClass = "devnet"
	ClassTestnet ChainClass = "testnet"
	ClassMainnet ChainClass = "mainnet"
)

// Token represents a deployed instance of an Asset on a specific blockchain.
// It includes both ERC20 and native tokens.
type Token struct {
	Asset
	ChainID    uint64
	ChainClass ChainClass
	Address    common.Address // empty if native
	IsMock     bool
}

func (t Token) IsNative() bool {
	return t.Address == NativeAddr
}

func (t Token) Is(asset Asset) bool {
	return t.Asset == asset
}

func (t Token) IsOMNI() bool {
	return t.Is(OMNI)
}

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
	usdc(evmchain.IDEthereum, addr("0xa0b86991c6218b36c1d19d4a2e9eb0ce3606eb48")),
	usdc(evmchain.IDArbitrumOne, addr("0xaf88d065e77c8cC2239327C5EDb3A432268e5831")),
	usdc(evmchain.IDOptimism, addr("0x0b2c639c533813f4aa9d7837caf62653d097ff85")),
	usdc(evmchain.IDBase, addr("0x833589fCD6eDb6E08f4c7C32D4f71b54bdA02913")),

	// USDC (testnet)
	usdc(evmchain.IDSepolia, addr("0x1c7D4B196Cb0C7B01d743Fbc6116a902379C7238")),
	usdc(evmchain.IDArbSepolia, addr("0x75faf114eafb1BDbe2F0316DF893fd58CE46AA4d")),
	usdc(evmchain.IDOpSepolia, addr("0x5fd84259d66Cd46123540766Be93DFE6D43130D7")),
	usdc(evmchain.IDBaseSepolia, addr("0x036CbD53842c5426634e7929541eC2318f3dCF7e")),

	// ERC20 OMNI
	omniERC20(netconf.Mainnet),
	omniERC20(netconf.Omega),
	omniERC20(netconf.Staging),
	omniERC20(netconf.Devnet),

	// wstETH
	wstETH(evmchain.IDBase, addr("0xc1cba3fcea344f92d9239c08c0568f6f2f0ee452")),
	wstETH(evmchain.IDEthereum, addr("0x7f39c581f595b53c5cb19bd0b3f8da6c935e2ca0")),
	wstETH(evmchain.IDHolesky, addr("0x8d09a4502cc8cf1547ad300e066060d043f6982d")),
	wstETH(evmchain.IDSepolia, addr("0xB82381A3fBD3FaFA77B3a7bE693342618240067b")),
	// Mocks contain wstETH on IDBaseSepolia for omega and staging

	// stETH
	stETH(evmchain.IDEthereum, addr("0xae7ab96520de3a18e5e111b5eaab095312d7fe84")),
	stETH(evmchain.IDHolesky, addr("0x3f1c547b21f65e10480de3ad8e19faac46c95034")),

	// WETH
	weth(evmchain.IDEthereum, addr("0xc02aaa39b223fe8d0a0e5c4f27ead9083c756cc2")),
	weth(evmchain.IDArbitrumOne, addr("0x82af49447d8a07e3bd95bd0d56f35241523fbab1")),
	weth(evmchain.IDBase, addr("0x4200000000000000000000000000000000000006")),
	weth(evmchain.IDOptimism, addr("0x4200000000000000000000000000000000000006")),
}, mocks()...)

func All() []Token {
	return tokens
}

// BySymbol returns the token with the given symbol and chain ID.
func BySymbol(chainID uint64, symbol string) (Token, bool) {
	for _, t := range tokens {
		if t.ChainID == chainID && t.Symbol == symbol {
			return t, true
		}
	}

	return Token{}, false
}

// ByAsset returns the token with the given Asset and chain ID.
func ByAsset(chainID uint64, asset Asset) (Token, bool) {
	for _, t := range tokens {
		if t.ChainID == chainID && t.Is(asset) {
			return t, true
		}
	}

	return Token{}, false
}

// Native is an alias for ByAddress with the native token address.
func Native(chainID uint64) (Token, bool) {
	return ByAddress(chainID, NativeAddr)
}

// ByAddress returns the token with the given address and chain ID.
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
	return Token{
		Asset:      ETH,
		ChainID:    chainID,
		ChainClass: mustChainClass(chainID),
		Address:    NativeAddr,
	}
}

func weth(chainID uint64, addr common.Address) Token {
	return Token{
		Asset:      WETH,
		ChainID:    chainID,
		ChainClass: mustChainClass(chainID),
		Address:    addr,
	}
}

func nativeOMNI(chainID uint64) Token {
	return Token{
		Asset:      OMNI,
		ChainID:    chainID,
		ChainClass: mustChainClass(chainID),
		Address:    NativeAddr,
	}
}

func omniERC20(network netconf.ID) Token {
	chainID := netconf.EthereumChainID(network)

	return Token{
		Asset:      OMNI,
		ChainID:    chainID,
		ChainClass: mustChainClass(chainID),
		Address:    contracts.TokenAddr(network),
	}
}

// mockOMNI returns a manually deployed OMNI token on a given chain for testing purposes.
func mockOMNI(chainID uint64, addr common.Address) Token {
	return Token{
		Asset:      OMNI,
		ChainID:    chainID,
		ChainClass: mustChainClass(chainID),
		Address:    addr,
		IsMock:     true,
	}
}

func stETH(chainID uint64, addr common.Address) Token {
	return Token{
		Asset:      STETH,
		ChainID:    chainID,
		ChainClass: mustChainClass(chainID),
		Address:    addr,
	}
}

func wstETH(chainID uint64, addr common.Address) Token {
	return Token{
		Asset:      WSTETH,
		ChainID:    chainID,
		ChainClass: mustChainClass(chainID),
		Address:    addr,
	}
}

func usdc(chainID uint64, addr common.Address) Token {
	return Token{
		Asset:      USDC,
		ChainID:    chainID,
		ChainClass: mustChainClass(chainID),
		Address:    addr,
	}
}

//go:embed mock_tokens.json
var mockTokenJSON []byte

// mocks returns MockTokens deployed in e2e for testing purposes.
func mocks() []Token {
	// e2e/solve#MockTokens are embedded as json instead of imported.
	// This prevents /lib/tokens from depending on e2e/solve.
	// /lib/* packages should be proper libraries so only depend on other lib/* packages.
	var tkns []Token
	if err := json.Unmarshal(mockTokenJSON, &tkns); err != nil {
		panic(errors.Wrap(err, "unmarshal mock tokens"))
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

func addr(hex string) common.Address {
	return common.HexToAddress(hex)
}
