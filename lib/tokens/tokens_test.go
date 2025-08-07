package tokens_test

import (
	"encoding/json"
	"flag"
	"io"
	"net/http"
	"strings"
	"testing"

	e2e "github.com/omni-network/omni/e2e/solve"
	"github.com/omni-network/omni/lib/contracts"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/evmchain"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/svmutil"
	"github.com/omni-network/omni/lib/tokens"
	"github.com/omni-network/omni/lib/tutil"

	"github.com/ethereum/go-ethereum/common"

	"github.com/gagliardetto/solana-go"
	"github.com/stretchr/testify/require"
)

var integration = flag.Bool("integration", false, "run integration tests")

//go:generate go test . -golden -integration

// TestGenTokens generates tokens.json
//
// This ensures that this lib/tokens package is self-contained
// and doesn't depend on external libraries which result in
// cyclic dependencies.
func TestGenTokens(t *testing.T) {
	t.Parallel()

	tkns := append([]tokens.Token{},
		// Native ETH (mainnet)
		nativeETH(evmchain.IDEthereum),
		nativeETH(evmchain.IDArbitrumOne),
		nativeETH(evmchain.IDBase),
		nativeETH(evmchain.IDOptimism),

		// Native ETH (testnet)
		nativeETH(evmchain.IDHolesky),
		nativeETH(evmchain.IDSepolia),
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
		usdc(evmchain.IDMantle, addr("0x09Bc4E0D864854c6aFB6eB9A9cdF58aC190D0dF9")),

		// USDC (testnet)
		usdc(evmchain.IDSepolia, addr("0x1c7D4B196Cb0C7B01d743Fbc6116a902379C7238")),
		usdc(evmchain.IDArbSepolia, addr("0x75faf114eafb1BDbe2F0316DF893fd58CE46AA4d")),
		usdc(evmchain.IDOpSepolia, addr("0x5fd84259d66Cd46123540766Be93DFE6D43130D7")),
		usdc(evmchain.IDBaseSepolia, addr("0x036CbD53842c5426634e7929541eC2318f3dCF7e")),

		// USDT (mainnet)
		usdt(evmchain.IDEthereum, addr("0xdac17f958d2ee523a2206206994597c13d831ec7")),
		usdt(evmchain.IDArbitrumOne, addr("0xfd086bc7cd5c481dcc9c85ebe478a1c0b69fcbb9")),
		usdt(evmchain.IDOptimism, addr("0x94b008aA00579c1307B0EF2c499aD98a8ce58e58")),

		// USDT0 (mainnet)
		usdt0(evmchain.IDArbitrumOne, addr("0xFd086bC7CD5C481DCC9C85ebE478A1C0b69FCbb9")),
		usdt0(evmchain.IDOptimism, addr("0x01bFF41798a0BcF287b996046Ca68b395DbC1071")),
		usdt0(evmchain.IDHyperEVM, addr("0xB8CE59FC3717ada4C02eaDF9682A9e934F625ebb")),

		// ERC20 OMNI
		omniERC20(netconf.Mainnet),
		omniERC20(netconf.Omega),
		omniERC20(netconf.Staging),
		omniERC20(netconf.Devnet),

		// ERC20 NOM
		nomERC20(netconf.Mainnet),
		nomERC20(netconf.Omega),
		nomERC20(netconf.Staging),
		nomERC20(netconf.Devnet),

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
		weth(evmchain.IDMantle, addr("0xdEAddEaDdeadDEadDEADDEAddEADDEAddead1111")),

		// MNT
		mntNative(evmchain.IDMantle),
		mntERC20(evmchain.IDEthereum, addr("0x3c3a81e81dc49a522a592e7622a7e711c06bf354")),

		// mETH
		meth(evmchain.IDEthereum, addr("0xd5f7838f5c461feff7fe49ea5ebaf7728bb0adfa")),
		meth(evmchain.IDMantle, addr("0xcDA86A272531e8640cD7F1a92c01839911B90bb0")),

		// HYPE
		tokens.Token{
			Asset:      tokens.HYPE,
			ChainID:    evmchain.IDHyperEVM,
			ChainClass: tokens.ClassMainnet,
		},

		// SVM mints/tokens
		svm(evmchain.IDSolanaLocal, tokens.USDC, svmutil.DevnetUSDCMint.PublicKey()),
	)

	for _, mock := range e2e.MockTokens() {
		chainClass := tokens.ClassDevent
		if mock.NetworkID == netconf.Staging {
			chainClass = tokens.ClassTestnet
		}

		tkns = append(tkns, tokens.Token{
			Asset:      mock.Asset,
			Address:    mock.Address(),
			ChainID:    mock.ChainID,
			ChainClass: chainClass,
			IsMock:     true,
		})
	}

	// Add manually deployed tokens that aren't part of the automatic mock deployment
	tkns = append(tkns,
		mockOMNI(evmchain.IDBaseSepolia, common.HexToAddress("0xe4075366F03C286846dECE8AAF104cF2a542294d")),
		mockOMNI(evmchain.IDOpSepolia, common.HexToAddress("0x0b3AED256a51919b660fF79a280A309EecA9d688")),
		mockOMNI(evmchain.IDArbSepolia, common.HexToAddress("0xd859f9Ff3C9700fB623Dc8e76217ba2a9f8613F0")),
	)

	tutil.RequireGoldenJSON(t, tkns, tutil.WithFilename("../tokens.json"))
}

// TestAssetCGIDs tests that all assets have a valid CoingeckoIDs.
func TestAssetCGIDs(t *testing.T) {
	t.Parallel()
	if !*integration {
		t.Skip("integration test")
	}

	coins, err := listCGCoins()
	require.NoError(t, err)

	for _, asset := range tokens.UniqueAssets() {
		// TODO(zodomo): remove this once NOM is listed on CoinGecko
		if asset.Symbol == "NOM" {
			continue
		}

		symbol, ok := coins[asset.CoingeckoID]
		require.True(t, ok, "missing asset %s", asset.CoingeckoID)
		require.Equal(t, strings.ToLower(asset.Symbol), symbol, "asset %s has different symbol", asset.CoingeckoID)
	}
}

func nativeETH(chainID uint64) tokens.Token {
	return tokens.Token{
		Asset:      tokens.ETH,
		ChainID:    chainID,
		ChainClass: mustChainClass(chainID),
		Address:    tokens.NativeAddr,
	}
}

func weth(chainID uint64, addr common.Address) tokens.Token {
	return tokens.Token{
		Asset:      tokens.WETH,
		ChainID:    chainID,
		ChainClass: mustChainClass(chainID),
		Address:    addr,
	}
}

func nativeOMNI(chainID uint64) tokens.Token {
	return tokens.Token{
		Asset:      tokens.OMNI,
		ChainID:    chainID,
		ChainClass: mustChainClass(chainID),
		Address:    tokens.NativeAddr,
	}
}

func mntNative(chainID uint64) tokens.Token {
	return tokens.Token{
		Asset:      tokens.MNT,
		ChainID:    chainID,
		ChainClass: mustChainClass(chainID),
		Address:    tokens.NativeAddr,
	}
}

func mntERC20(chainID uint64, addr common.Address) tokens.Token {
	return tokens.Token{
		Asset:      tokens.MNT,
		ChainID:    chainID,
		ChainClass: mustChainClass(chainID),
		Address:    addr,
	}
}

func meth(chainID uint64, addr common.Address) tokens.Token {
	return tokens.Token{
		Asset:      tokens.METH,
		ChainID:    chainID,
		ChainClass: mustChainClass(chainID),
		Address:    addr,
	}
}

func omniERC20(network netconf.ID) tokens.Token {
	chainID := netconf.EthereumChainID(network)

	return tokens.Token{
		Asset:      tokens.OMNI,
		ChainID:    chainID,
		ChainClass: mustChainClass(chainID),
		Address:    contracts.TokenAddr(network),
	}
}

func nomERC20(network netconf.ID) tokens.Token {
	chainID := netconf.EthereumChainID(network)

	return tokens.Token{
		Asset:      tokens.NOM,
		ChainID:    chainID,
		ChainClass: mustChainClass(chainID),
		Address:    contracts.NomAddr(network),
	}
}

func svm(chainID uint64, asset tokens.Asset, address solana.PublicKey) tokens.Token {
	return tokens.Token{
		Asset:   asset,
		ChainID: chainID,
		ChainClass: map[uint64]tokens.ChainClass{
			evmchain.IDSolanaLocal:  tokens.ClassDevent,
			evmchain.IDSolanaDevnet: tokens.ClassTestnet,
			evmchain.IDSolana:       tokens.ClassMainnet,
		}[chainID],
		SVMAddress: address,
		IsMock:     chainID == evmchain.IDSolanaLocal,
	}
}

// mockOMNI returns a manually deployed OMNI tokens.Token on a given chain for testing purposes.
func mockOMNI(chainID uint64, addr common.Address) tokens.Token {
	return tokens.Token{
		Asset:      tokens.OMNI,
		ChainID:    chainID,
		ChainClass: mustChainClass(chainID),
		Address:    addr,
		IsMock:     true,
	}
}

func stETH(chainID uint64, addr common.Address) tokens.Token {
	return tokens.Token{
		Asset:      tokens.STETH,
		ChainID:    chainID,
		ChainClass: mustChainClass(chainID),
		Address:    addr,
	}
}

func wstETH(chainID uint64, addr common.Address) tokens.Token {
	return tokens.Token{
		Asset:      tokens.WSTETH,
		ChainID:    chainID,
		ChainClass: mustChainClass(chainID),
		Address:    addr,
	}
}

func usdc(chainID uint64, addr common.Address) tokens.Token {
	return tokens.Token{
		Asset:      tokens.USDC,
		ChainID:    chainID,
		ChainClass: mustChainClass(chainID),
		Address:    addr,
	}
}

func usdt(chainID uint64, addr common.Address) tokens.Token {
	return tokens.Token{
		Asset:      tokens.USDT,
		ChainID:    chainID,
		ChainClass: mustChainClass(chainID),
		Address:    addr,
	}
}

func usdt0(chainID uint64, addr common.Address) tokens.Token {
	return tokens.Token{
		Asset:      tokens.USDT0,
		ChainID:    chainID,
		ChainClass: mustChainClass(chainID),
		Address:    addr,
	}
}

func mustChainClass(chainID uint64) tokens.ChainClass {
	class, err := chainClass(chainID)
	if err != nil {
		panic(err)
	}

	return class
}

func chainClass(chainID uint64) (tokens.ChainClass, error) {
	switch chainID {
	case
		evmchain.IDOmniMainnet,
		evmchain.IDEthereum,
		evmchain.IDArbitrumOne,
		evmchain.IDBase,
		evmchain.IDOptimism,
		evmchain.IDMantle,
		evmchain.IDHyperEVM:
		return tokens.ClassMainnet, nil
	case
		evmchain.IDOmniOmega,
		evmchain.IDOmniStaging, // classify omni staging as testnet, because it interops with other testnets
		evmchain.IDHolesky,
		evmchain.IDSepolia,
		evmchain.IDArbSepolia,
		evmchain.IDBaseSepolia,
		evmchain.IDOpSepolia:
		return tokens.ClassTestnet, nil
	case
		evmchain.IDOmniDevnet,
		evmchain.IDMockL1,
		evmchain.IDMockL2,
		evmchain.IDMockArb,
		evmchain.IDMockOp:
		return tokens.ClassDevent, nil
	default:
		return "", errors.New("unsupported chain ID", "chain_id", chainID)
	}
}

func addr(hex string) common.Address {
	return common.HexToAddress(hex)
}

// listCGCount returns all supported coins from CoinGecko; map[CoinGeckoID]Symbol.
func listCGCoins() (map[string]string, error) {
	var bodyJSON []struct {
		ID     string `json:"id"`
		Symbol string `json:"symbol"`
		Name   string `json:"name"`
	}

	resp, err := http.Get("https://api.coingecko.com/api/v3/coins/list")
	if err != nil {
		return nil, errors.Wrap(err, "get")
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "read")
	}

	if err := json.Unmarshal(body, &bodyJSON); err != nil {
		return nil, errors.Wrap(err, "unmarshal")
	}

	coins := make(map[string]string)
	for _, coin := range bodyJSON {
		if coin.ID == "" {
			continue
		}
		coins[coin.ID] = coin.Symbol
	}

	return coins, nil
}
