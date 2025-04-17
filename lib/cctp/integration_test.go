package cctp_test

import (
	"context"
	"crypto/ecdsa"
	"flag"
	"math/big"
	"os"
	"testing"

	"github.com/omni-network/omni/lib/anvil"
	"github.com/omni-network/omni/lib/bi"
	"github.com/omni-network/omni/lib/cctp"
	"github.com/omni-network/omni/lib/cctp/db"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient"
	"github.com/omni-network/omni/lib/evmchain"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/tokens"
	"github.com/omni-network/omni/lib/tutil"
	xprovider "github.com/omni-network/omni/lib/xchain/provider"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"

	cosmosdb "github.com/cosmos/cosmos-db"
	"github.com/stretchr/testify/require"
)

var (
	integration           = flag.Bool("integration", false, "run integration tests")
	messageTransmitterABI = mustGetABI(cctp.MessageTransmitterMetaData)
)

//go:generate go test . -integration -v -run=TestIntegration

// TestIntegration runs a CCTP integration test. It:
//   - runs anvil forks of mainnet chains
//   - attests and signs CCTP messages w/ a devnet attester
//   - bridges USDC from chain to chain, asserting success
func TestIntegration(t *testing.T) {
	t.Parallel()

	if !*integration {
		t.Skip("skipping integration test")
	}

	ctx := t.Context()

	rpcs := getForkRPCs(t)
	chains := getChains(t)
	network := makeNetwork(t, chains)

	clients, stop := startAnvilForks(t, rpcs, chains)
	defer stop()

	attesterPk, attesterAddr := newAccount(t) // CCTP attester
	devPk, devAddr := newAccount(t)           // Test user

	// Enable attester on all chains
	enableAttester(t, clients, attesterAddr)

	// Fund dev account with USDC & ETH
	fund(t, clients, devAddr)

	// Create xchain provider
	xprov := xprovider.New(network, clients, nil)

	// Create CCTP client, start attesting
	cctpClient := cctp.NewDevClient(attesterPk, clients)
	err := cctpClient.AttestForever(ctx, chains, xprov)
	require.NoError(t, err)

	// In-mem CCTP DB
	memDB := cosmosdb.NewMemDB()
	cctpDB, err := db.New(memDB)
	require.NoError(t, err)

	_ = cctpDB
	_ = devPk
}

// getForkRPCs returns mainnet rpcs urls from env vars.
func getForkRPCs(t *testing.T) map[uint64]string {
	t.Helper()

	notEmpty := func(env string) string {
		v := os.Getenv(env)
		require.NotEmpty(t, v, "%s must be set", env)

		return v
	}

	return map[uint64]string{
		evmchain.IDEthereum:    notEmpty("ETH_RPC"),
		evmchain.IDArbitrumOne: notEmpty("ARB_RPC"),
		evmchain.IDOptimism:    notEmpty("OP_RPC"),
		evmchain.IDBase:        notEmpty("BASE_RPC"),
	}
}

// startAnvilForks starts anvil forks fork each chain, returning a clients map and a "stop all" function.
func startAnvilForks(t *testing.T, rpcs map[uint64]string, chains []evmchain.Metadata) (map[uint64]ethclient.Client, func()) {
	t.Helper()

	ctx := t.Context()
	clients := make(map[uint64]ethclient.Client)

	var stops []func()
	for _, chain := range chains {
		ethCl, stop, err := anvil.Start(ctx, tutil.TempDir(t), chain.ChainID,
			anvil.WithFork(rpcs[chain.ChainID]),
			anvil.WithAutoImpersonate(),
		)
		require.NoError(t, err)

		log.Info(ctx, "Stated anvil fork", "chain", chain.Name, "rpc", rpcs[chain.ChainID])

		clients[chain.ChainID] = ethCl
		stops = append(stops, stop)
	}

	stopAll := func() {
		for _, stop := range stops {
			stop()
		}
	}

	return clients, stopAll
}

// fund funds accounts with USDC on each chain.
func fund(t *testing.T, clients map[uint64]ethclient.Client, account common.Address) {
	t.Helper()
	ctx := t.Context()

	// Fund USDC
	for chainID, client := range clients {
		amount := bi.Dec6(1000) // 1000 USDC
		usdc := mustUsdc(t, chainID)
		err := anvil.FundERC20(t.Context(), client, usdc.Address, amount, account)
		require.NoError(t, err)
		log.Info(ctx, "Funded USDC", "chain", chainID, "amount", amount, "account", account)
	}

	// Fund ETH
	for chainID, client := range clients {
		amount := bi.Ether(1) // 1 ETH
		err := anvil.FundAccounts(t.Context(), client, amount, account)
		require.NoError(t, err)
		log.Info(ctx, "Funded ETH", "chain", chainID, "amount", amount, "account", account)
	}
}

func getChains(t *testing.T) []evmchain.Metadata {
	t.Helper()

	return []evmchain.Metadata{
		mustMeta(t, evmchain.IDEthereum),
		mustMeta(t, evmchain.IDArbitrumOne),
		mustMeta(t, evmchain.IDOptimism),
		mustMeta(t, evmchain.IDBase),
	}
}

func newAccount(t *testing.T) (*ecdsa.PrivateKey, common.Address) {
	t.Helper()

	pk, err := crypto.GenerateKey()
	require.NoError(t, err)
	addr := crypto.PubkeyToAddress(pk.PublicKey)

	return pk, addr
}

func makeNetwork(t *testing.T, chains []evmchain.Metadata) netconf.Network {
	t.Helper()

	network := netconf.Network{ID: netconf.Devnet}
	network.Chains = make([]netconf.Chain, len(chains))
	for i, chain := range chains {
		network.Chains[i] = netconf.Chain{
			ID:   chain.ChainID,
			Name: chain.Name,
		}
	}

	return network
}

func mustMeta(t *testing.T, chainID uint64) evmchain.Metadata {
	t.Helper()

	meta, ok := evmchain.MetadataByID(chainID)
	require.True(t, ok)

	return meta
}

func mustUsdc(t *testing.T, chainID uint64) tokens.Token {
	t.Helper()
	usdc, ok := tokens.ByAsset(chainID, tokens.USDC)
	require.True(t, ok)

	return usdc
}

// enableAttester enables the attester on each chains MessageTransmitter.
func enableAttester(t *testing.T, clients map[uint64]ethclient.Client, newAttester common.Address) {
	t.Helper()

	ctx := t.Context()

	for chainID, client := range clients {
		mtAddr, ok := cctp.MessageTransmitterAddr(chainID)
		require.True(t, ok)

		mt, err := cctp.NewMessageTransmitter(mtAddr, client)
		require.NoError(t, err)

		// AttesterManager is account allowed to enable attesters
		mngr, err := mt.AttesterManager(&bind.CallOpts{Context: ctx})
		require.NoError(t, err)

		// Send unsigned MessageTransmitter.enableAttester tx from attester manager
		// This requires anvil auto impersonation
		calldata, err := messageTransmitterABI.Pack("enableAttester", newAttester)
		require.NoError(t, err)
		txHash, err := sendUnsignedTransaction(ctx, client, txArgs{
			from:  mngr,
			to:    mtAddr,
			value: nil,
			data:  calldata,
		})
		require.NoError(t, err)

		_, err = bind.WaitMinedHash(ctx, client, txHash)
		require.NoError(t, err)

		// Verify attester is enabled
		enabled, err := mt.IsEnabledAttester(&bind.CallOpts{Context: ctx}, newAttester)
		require.NoError(t, err)
		require.True(t, enabled, "attester not enabled on chain %d", chainID)

		log.Info(ctx, "Enabled attester", "chain", chainID, "attester", newAttester)
	}
}

type txArgs struct {
	from  common.Address
	to    common.Address
	value *big.Int
	data  []byte
}

// This is used to send auto impersonated txs on anvil.
func sendUnsignedTransaction(ctx context.Context, client ethclient.Client, args txArgs) (common.Hash, error) {
	value := bi.Zero()
	if args.value != nil {
		value = args.value
	}

	to := args.to
	msg := ethereum.CallMsg{
		From:  args.from,
		To:    &to,
		Value: value,
		Data:  args.data,
	}

	gas, err := client.EstimateGas(ctx, msg)
	if err != nil {
		return common.Hash{}, errors.Wrap(err, "estimate gas")
	}

	jsonArgs := map[string]any{
		"from":  args.from,
		"to":    to,
		"value": (*hexutil.Big)(value),
		"data":  hexutil.Bytes(args.data),
		"gas":   hexutil.Uint64(gas),
	}

	var result common.Hash
	err = client.CallContext(ctx, &result, "eth_sendTransaction", jsonArgs)
	if err != nil {
		return common.Hash{}, errors.Wrap(err, "send transaction")
	}

	return result, nil
}

func mustGetABI(metadata *bind.MetaData) *abi.ABI {
	abi, err := metadata.GetAbi()
	if err != nil {
		panic(err)
	}

	return abi
}
