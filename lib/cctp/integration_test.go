package cctp_test

import (
	"context"
	"crypto/ecdsa"
	"flag"
	"log/slog"
	"math/big"
	"os"
	"os/signal"
	"syscall"
	"testing"
	"time"

	"github.com/omni-network/omni/lib/anvil"
	"github.com/omni-network/omni/lib/bi"
	"github.com/omni-network/omni/lib/cctp"
	cctpdb "github.com/omni-network/omni/lib/cctp/db"
	"github.com/omni-network/omni/lib/cctp/types"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient"
	"github.com/omni-network/omni/lib/ethclient/ethbackend"
	"github.com/omni-network/omni/lib/evmchain"
	"github.com/omni-network/omni/lib/log"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/tokens"
	"github.com/omni-network/omni/lib/tokens/tokenutil"
	"github.com/omni-network/omni/lib/tutil"
	"github.com/omni-network/omni/lib/xchain"
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

	logCfg := log.DefaultConfig()
	logCfg.Level = slog.LevelDebug.String()
	logCfg.Color = log.ColorForce

	ctx, err := log.Init(ctx, logCfg)
	require.NoError(t, err)

	// Handle interrupts
	ctx, cancel := signal.NotifyContext(ctx, os.Interrupt, syscall.SIGTERM)
	defer cancel()

	rpcs := getForkRPCs(t)
	chains := getChains(t)
	network := makeNetwork(t, chains)

	clients, stop := startAnvilForks(t, ctx, rpcs, chains)
	defer stop()

	// Stop anvil on interrupt
	go func() {
		<-ctx.Done()
		stop()
	}()

	attesterPk, attesterAddr := newAccount(t) // CCTP attester
	devPk, devAddr := newAccount(t)           // Test user

	backends, err := ethbackend.BackendsFromClients(clients, devPk)
	require.NoError(t, err)

	// Enable attester on all chains
	enableAttester(t, ctx, clients, attesterAddr)

	// Fund dev account with USDC & ETH
	fund(t, ctx, clients, devAddr)

	// Create xchain provider
	xprov := xprovider.New(network, clients, nil)

	// Create CCTP client, start attesting
	cctpClient := cctp.NewDevClient(attesterPk, clients)
	err = cctpClient.AttestForever(ctx, chains, xprov)
	require.NoError(t, err)

	// In-mem CCTP DB
	memDB := cosmosdb.NewMemDB()
	db, err := cctpdb.New(memDB)
	require.NoError(t, err)

	// Mint forever
	err = cctp.MintForever(ctx, db, cctpClient, backends, chains, devAddr,
		cctp.WithMintInterval(1*time.Second),
		cctp.WithPurgeInterval(10*time.Second))
	require.NoError(t, err)

	// Audit forever
	err = cctp.AuditForever(ctx, db, netconf.Mainnet, xprov, clients, chains, devAddr)
	require.NoError(t, err)

	// Track initial balances
	initialBalances := make(map[uint64]*big.Int)
	for _, chain := range chains {
		usdc := mustUSDC(t, chain.ChainID)
		backend, err := backends.Backend(chain.ChainID)
		require.NoError(t, err)

		balance, err := tokenutil.BalanceOf(ctx, backend, usdc, devAddr)
		require.NoError(t, err)

		initialBalances[chain.ChainID] = balance
	}

	// Track expected final balances
	expectedBalances := make(map[uint64]*big.Int)
	for chainID, balance := range initialBalances {
		expectedBalances[chainID] = new(big.Int).Set(balance)
	}

	// wrongDb is used to simulate messages missed, and not stored in the db.
	// these should be caught by audit.
	wrongCosmosDB := cosmosdb.NewMemDB()
	wrongDB, err := cctpdb.New(wrongCosmosDB)
	require.NoError(t, err)

	// Define sends
	sends := []struct {
		srcChainID  uint64
		destChainID uint64
		amount      *big.Int
		wrongDB     bool
	}{
		{evmchain.IDEthereum, evmchain.IDOptimism, bi.Dec6(50), false},
		{evmchain.IDArbitrumOne, evmchain.IDBase, bi.Dec6(50), false},
		{evmchain.IDOptimism, evmchain.IDBase, bi.Dec6(25), true}, // wrong db
		{evmchain.IDBase, evmchain.IDEthereum, bi.Dec6(10), false},
		{evmchain.IDEthereum, evmchain.IDOptimism, bi.Dec6(5), false},
		{evmchain.IDOptimism, evmchain.IDBase, bi.Dec6(2), false},
		{evmchain.IDBase, evmchain.IDArbitrumOne, bi.Dec6(1), true}, // wrong db
		{evmchain.IDArbitrumOne, evmchain.IDEthereum, bi.Dec6(0.5), false},
		{evmchain.IDEthereum, evmchain.IDBase, bi.Dec6(0.1), true}, // wrong db
	}

	// Do sends
	msgs := make([]types.MsgSendUSDC, len(sends))
	for i, send := range sends {
		backend, err := backends.Backend(send.srcChainID)
		require.NoError(t, err)

		sendDB := db
		if send.wrongDB {
			sendDB = wrongDB
		}

		msg, err := cctp.SendUSDC(ctx, sendDB, backend, cctp.SendUSDCArgs{
			Sender:      devAddr,
			Recipient:   devAddr,
			SrcChainID:  send.srcChainID,
			DestChainID: send.destChainID,
			Amount:      send.amount,
		})
		require.NoError(t, err)
		msgs[i] = msg

		// Update expected balances
		expectedBalances[send.srcChainID] = bi.Sub(expectedBalances[send.srcChainID], send.amount)
		expectedBalances[send.destChainID] = bi.Add(expectedBalances[send.destChainID], send.amount)
	}

	// Wait for all sends
	tutil.RequireEventually(t, ctx, func() bool {
		for chainID, expectedBalance := range expectedBalances {
			usdc := mustUSDC(t, chainID)

			backend, err := backends.Backend(chainID)
			tutil.RequireNoError(t, err)

			balance, err := tokenutil.BalanceOf(ctx, backend, usdc, devAddr)
			tutil.RequireNoError(t, err)

			if !bi.EQ(balance, expectedBalance) {
				return false
			}
		}

		log.Info(ctx, "All sends completed")

		return true
	}, 2*time.Minute, 1*time.Second)

	// Confirm all messages received
	for _, msg := range msgs {
		received, err := cctp.DidReceive(ctx, clients[msg.DestChainID], msg, nil)
		require.NoError(t, err)
		require.True(t, received, "message not received on dest chain %d", msg.DestChainID)
	}

	// Wait for all purged (confirmed and deleted)
	tutil.RequireEventually(t, ctx, func() bool {
		msgs, err := db.GetMsgs(ctx)
		require.NoError(t, err)

		if len(msgs) > 0 {
			return false
		}

		log.Info(ctx, "All messages purged")

		return true
	}, 2*time.Minute, 1*time.Second)
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
func startAnvilForks(t *testing.T, ctx context.Context, rpcs map[uint64]string, chains []evmchain.Metadata) (map[uint64]ethclient.Client, func()) {
	t.Helper()

	clients := make(map[uint64]ethclient.Client)

	var stops []func()
	for _, chain := range chains {
		ethCl, stop, err := anvil.Start(ctx, tutil.TempDir(t), chain.ChainID,
			anvil.WithFork(rpcs[chain.ChainID]),
			anvil.WithAutoImpersonate(),
			// quick finalization for testing
			anvil.WithBlockTime(1),
			anvil.WithSlotsInEpoch(2),
		)
		require.NoError(t, err)

		log.Info(ctx, "Stated anvil fork", "chain", chain.Name, "fork_rpc", rpcs[chain.ChainID])

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
func fund(t *testing.T, ctx context.Context, clients map[uint64]ethclient.Client, account common.Address) {
	t.Helper()

	// Fund USDC
	for chainID, client := range clients {
		amount := bi.Dec6(1000) // 1000 USDC
		usdc := mustUSDC(t, chainID)
		err := anvil.FundUSDC(ctx, client, usdc.Address, amount, account)
		require.NoError(t, err)
		log.Info(ctx, "Funded USDC", "chain", chainID, "amount", amount, "account", account)
	}

	// Fund ETH
	for chainID, client := range clients {
		amount := bi.Ether(1) // 1 ETH
		err := anvil.FundAccounts(ctx, client, amount, account)
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
			ID:             chain.ChainID,
			Name:           chain.Name,
			BlockPeriod:    chain.BlockPeriod,
			Shards:         []xchain.ShardID{xchain.ShardFinalized0, xchain.ShardLatest0},
			AttestInterval: 10000,
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

func mustUSDC(t *testing.T, chainID uint64) tokens.Token {
	t.Helper()
	usdc, ok := tokens.ByAsset(chainID, tokens.USDC)
	require.True(t, ok)

	return usdc
}

// enableAttester enables the attester on each chains MessageTransmitter, and sets the signature threshold to 1.
func enableAttester(t *testing.T, ctx context.Context, clients map[uint64]ethclient.Client, newAttester common.Address) {
	t.Helper()

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

		// Reduce signature threshold to 1 (number of attesters required)
		calldata, err = messageTransmitterABI.Pack("setSignatureThreshold", bi.One())
		require.NoError(t, err)
		txHash, err = sendUnsignedTransaction(ctx, client, txArgs{
			from:  mngr,
			to:    mtAddr,
			value: nil,
			data:  calldata,
		})
		require.NoError(t, err)

		_, err = bind.WaitMinedHash(ctx, client, txHash)
		require.NoError(t, err)

		threshold, err := mt.SignatureThreshold(&bind.CallOpts{Context: ctx})
		require.NoError(t, err)
		tutil.RequireEQ(t, threshold, bi.One(), "signature threshold not set to 1 on chain %d", chainID)

		log.Info(ctx, "Signature threshold set", "chain", chainID, "threshold", threshold)
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
