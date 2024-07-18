package cmd_test

import (
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"math/big"
	"os"
	"path/filepath"
	"testing"
	"time"

	clicmd "github.com/omni-network/omni/cli/cmd"
	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/e2e/app/eoa"
	"github.com/omni-network/omni/e2e/app/static"
	"github.com/omni-network/omni/lib/anvil"
	"github.com/omni-network/omni/lib/contracts"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient"
	"github.com/omni-network/omni/lib/ethclient/ethbackend"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/tutil"
	"github.com/omni-network/omni/lib/txmgr"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/params"

	eigentypes "github.com/Layr-Labs/eigenlayer-cli/pkg/types"
	eigenutils "github.com/Layr-Labs/eigenlayer-cli/pkg/utils"
	eigenecdsa "github.com/Layr-Labs/eigensdk-go/crypto/ecdsa"
	eigensdktypes "github.com/Layr-Labs/eigensdk-go/types"
	"github.com/stretchr/testify/require"
	"gopkg.in/yaml.v3"
)

const (
	chainName   = "test"
	chainID     = 99
	blockPeriod = time.Second
)

//nolint:gochecknoglobals // These are test constants.
var (
	avsABI = mustGetABI(bindings.OmniAVSMetaData)

	// devnet eigen strategy manger, and test weth strategy.
	stratMngrAddr = common.HexToAddress("0xe1DA8919f262Ee86f9BE05059C9280142CF23f48")
	wethStratAddr = common.HexToAddress("0xdBD296711eC8eF9Aacb623ee3F1C0922dce0D7b2")
)

//nolint:paralleltest // Parallel tests not supported since we start docker containers.
func TestRegister(t *testing.T) {
	ctx, backend, contracts, eoas := setup(t)

	// Register operators to omni AVS with a stake more than minimum stake
	for _, operator := range eoas.operators() {
		delegateWETH(t, ctx, contracts, backend, operator, toWei(100))
		registerOperator(t, ctx, contracts, backend, eoas.operatorKey(operator))
		assertOperatorRegistered(t, ctx, contracts, operator)
	}
}

// setup deploys the omni avs contract, using an anvil instance with pre-loaded eigen deployments.
func setup(t *testing.T) (context.Context, *ethbackend.Backend, Contracts, EOAS) {
	t.Helper()

	ctx := context.Background()

	ethCl, _, stop, err := anvil.Start(ctx, tutil.TempDir(t), chainID)
	require.NoError(t, err)
	t.Cleanup(stop)

	backend, err := ethbackend.NewAnvilBackend(chainName, chainID, blockPeriod, ethCl)
	require.NoError(t, err)

	eoas := makeEOAS(t, backend)
	addr := deployAVS(t, ctx, backend)
	contracts := makeContracts(t, addr, devnetEigenDeployments(t), backend)
	whiteListStrategy(t, ctx, backend, contracts, eoas.EigenOwner)

	// Fund operators with ETH and WETH (using the pre-funded eigen owner account)
	// So they can transact, and delegate to operators.
	for _, account := range eoas.operators() {
		fundAccount(t, ctx, backend, eoas.EigenOwner, account)
		mintWETHToAddresses(t, ctx, backend, contracts, eoas.EigenOwner, toWei(1000), account)
	}

	// Register operators with EigenLayer
	for i, operator := range eoas.operators() {
		err := registerOperatorWithEigen(ctx, contracts, backend, operator, fmt.Sprintf("https://operator%d.com", i))
		tutil.RequireNoError(t, err)
	}

	return ctx, backend, contracts, eoas
}

// eigenDeployments is a struct that holds the addresses of the EigenLayer core contracts.
type eigenDeployments struct {
	AVSDirectory      common.Address `json:"AVSDirectory"`
	DelegationManager common.Address `json:"DelegationManager"`
}

// deployConfig is a struct that holds the configuration for deploying the OmniAVS contract.
type deployConfig struct {
	eigen            eigenDeployments
	deployer         common.Address
	owner            common.Address
	portal           common.Address
	omniChainID      uint64
	metadataURI      string
	stratParams      []bindings.IOmniAVSStrategyParam
	ethStakeInbox    common.Address
	minOperatorStake *big.Int
	maxOperatorCount uint32
	allowlistEnabled bool
}

// testStratParams returns a list of strategy parameters for testing.
func testStratParams() []bindings.IOmniAVSStrategyParam {
	return []bindings.IOmniAVSStrategyParam{
		{
			// devnet WETH
			Strategy:   common.HexToAddress("0xdBD296711eC8eF9Aacb623ee3F1C0922dce0D7b2"),
			Multiplier: big.NewInt(1e18), // OmniAVS.STRATEGY_WEIGHTING_DIVISOR
		},
	}
}

// testDeployCfg returns a test deployment configuration.
func testDeployCfg(t *testing.T) deployConfig {
	t.Helper()

	return deployConfig{
		deployer:         eoa.MustAddress(netconf.Devnet, eoa.RoleDeployer),
		owner:            eoa.MustAddress(netconf.Devnet, eoa.RoleAdmin),
		eigen:            devnetEigenDeployments(t),
		metadataURI:      "https://test-operator.com",
		omniChainID:      netconf.Devnet.Static().OmniExecutionChainID,
		stratParams:      testStratParams(),
		portal:           contracts.Portal(netconf.Devnet),
		ethStakeInbox:    common.HexToAddress("0x1234"), // stub
		minOperatorStake: big.NewInt(1e18),              // 1 ETH
		maxOperatorCount: 10,
		allowlistEnabled: false,
	}
}

// deployAVS deploys the OmniAVS contract without a proxy, and initialize it in a second transaction.
func deployAVS(t *testing.T, ctx context.Context, backend *ethbackend.Backend) common.Address {
	t.Helper()

	cfg := testDeployCfg(t)

	txOpts, err := backend.BindOpts(ctx, cfg.deployer)
	require.NoError(t, err)

	impl, tx, _, err := bindings.DeployOmniAVS(txOpts, backend, cfg.eigen.DelegationManager, cfg.eigen.AVSDirectory)
	require.NoError(t, err)
	_, err = backend.WaitMined(ctx, tx)
	require.NoError(t, err)

	addr, tx, _, err := bindings.DeployTransparentUpgradeableProxy(txOpts, backend, impl, cfg.owner, packInitialzer(t, cfg))
	require.NoError(t, err)
	_, err = backend.WaitMined(ctx, tx)
	require.NoError(t, err)

	require.NoError(t, err)
	_, err = backend.WaitMined(ctx, tx)
	require.NoError(t, err)

	return addr
}

// packInitializer encodes the initializer parameters for the AVS contract.
func packInitialzer(t *testing.T, cfg deployConfig) []byte {
	t.Helper()

	enc, err := avsABI.Pack("initialize",
		cfg.owner, cfg.portal, cfg.omniChainID, cfg.ethStakeInbox,
		cfg.minOperatorStake, cfg.maxOperatorCount, cfg.stratParams,
		cfg.metadataURI, cfg.allowlistEnabled)
	require.NoError(t, err)

	return enc
}

func devnetEigenDeployments(t *testing.T) eigenDeployments {
	t.Helper()

	var el eigenDeployments
	err := json.Unmarshal(static.GetDevnetElDeployments(), &el)
	require.NoError(t, err)

	return el
}

type EOAS struct {
	AVSOwner     common.Address
	EigenOwner   common.Address
	Operator1    common.Address
	Operator1Key *ecdsa.PrivateKey
	Operator2    common.Address
	Operator2Key *ecdsa.PrivateKey
	Delgator1    common.Address
	Delgator2    common.Address
}

func (e EOAS) operatorKey(addr common.Address) *ecdsa.PrivateKey {
	switch addr {
	case e.Operator1:
		return e.Operator1Key
	case e.Operator2:
		return e.Operator2Key
	default:
		panic("unknown operator")
	}
}

func (e EOAS) operators() []common.Address {
	return []common.Address{e.Operator1, e.Operator2}
}

func makeEOAS(t *testing.T, backend *ethbackend.Backend) EOAS {
	t.Helper()

	// setup accounts
	operator1Key := genPrivKey(t)
	operator1, err := backend.AddAccount(operator1Key)
	require.NoError(t, err)
	operator2Key := genPrivKey(t)
	operator2, err := backend.AddAccount(operator2Key)
	require.NoError(t, err)
	delegator1, err := backend.AddAccount(genPrivKey(t))
	require.NoError(t, err)
	delegator2, err := backend.AddAccount(genPrivKey(t))
	require.NoError(t, err)

	return EOAS{
		AVSOwner:     eoa.MustAddress(netconf.Devnet, eoa.RoleAdmin),
		EigenOwner:   anvil.DevAccount9(), // account used to deploy eigen contracts
		Operator1:    operator1,
		Operator1Key: operator1Key,
		Operator2:    operator2,
		Operator2Key: operator2Key,
		Delgator1:    delegator1,
		Delgator2:    delegator2,
	}
}

type Contracts struct {
	OmniAVS           *bindings.OmniAVS
	DelegationManager *bindings.DelegationManager
	StrategyManager   *bindings.StrategyManager
	WETHStrategy      *bindings.StrategyBase
	WETHToken         *bindings.MockERC20
	AVSDirectory      *bindings.AVSDirectory

	OmniAVSAddr           common.Address
	DelegationManagerAddr common.Address
	StrategyManagerAddr   common.Address
	WETHStrategyAddr      common.Address
	WETHTokenAddr         common.Address
	AVSDirectoryAddr      common.Address
}

func makeContracts(t *testing.T, avsAddr common.Address, eigen eigenDeployments, backend *ethbackend.Backend) Contracts {
	t.Helper()

	delMan, err := bindings.NewDelegationManager(eigen.DelegationManager, backend)
	require.NoError(t, err)

	stratMan, err := bindings.NewStrategyManager(stratMngrAddr, backend)
	require.NoError(t, err)

	wethStrategy, err := bindings.NewStrategyBase(wethStratAddr, backend)
	require.NoError(t, err)

	wethTokenAddr, err := wethStrategy.UnderlyingToken(&bind.CallOpts{})
	require.NoError(t, err)

	wethToken, err := bindings.NewMockERC20(wethTokenAddr, backend)
	require.NoError(t, err)

	avsDir, err := bindings.NewAVSDirectory(eigen.AVSDirectory, backend)
	require.NoError(t, err)

	avs, err := bindings.NewOmniAVS(avsAddr, backend)
	require.NoError(t, err)

	return Contracts{
		OmniAVS:           avs,
		DelegationManager: delMan,
		StrategyManager:   stratMan,
		WETHStrategy:      wethStrategy,
		WETHToken:         wethToken,
		AVSDirectory:      avsDir,

		OmniAVSAddr:           avsAddr,
		DelegationManagerAddr: eigen.DelegationManager,
		StrategyManagerAddr:   stratMngrAddr,
		WETHStrategyAddr:      wethStratAddr,
		WETHTokenAddr:         wethTokenAddr,
		AVSDirectoryAddr:      eigen.AVSDirectory,
	}
}

func registerOperatorWithEigen(ctx context.Context, contracts Contracts, backend *ethbackend.Backend, operator common.Address, metadataURI string) error {
	operatorDetails := bindings.IDelegationManagerOperatorDetails{
		EarningsReceiver:         operator,
		DelegationApprover:       common.Address{},
		StakerOptOutWindowBlocks: uint32(0), // Currently unused by Eigen
	}

	txOpts, err := backend.BindOpts(ctx, operator)
	if err != nil {
		return err
	}

	tx, err := contracts.DelegationManager.RegisterAsOperator(txOpts, operatorDetails, metadataURI)
	if err != nil {
		return errors.Wrap(err, "register as operator")
	}

	_, err = backend.WaitMined(ctx, tx)
	if err != nil {
		return errors.Wrap(err, "wait mined")
	}

	return nil
}

func assertOperatorRegistered(t *testing.T, ctx context.Context, contracts Contracts, operator common.Address) {
	t.Helper()

	ops, err := contracts.OmniAVS.Operators(&bind.CallOpts{Context: ctx})
	require.NoError(t, err)

	for _, op := range ops {
		if op.Addr == operator {
			return
		}
	}

	require.Fail(t, "operator not found in omni avs")
}

func mintWETHToAddresses(
	t *testing.T,
	ctx context.Context,
	backend *ethbackend.Backend,
	contracts Contracts,
	funder common.Address,
	amount *big.Int,
	addrs ...common.Address) {
	t.Helper()

	for _, addr := range addrs {
		txOpts, err := backend.BindOpts(ctx, funder)
		require.NoError(t, err)

		tx, err := contracts.WETHToken.Mint(txOpts, addr, amount)
		require.NoError(t, err)
		_, err = backend.WaitMined(ctx, tx)
		require.NoError(t, err)
	}
}

func delegateWETH(
	t *testing.T,
	ctx context.Context,
	contracts Contracts,
	backend *ethbackend.Backend,
	delegator common.Address,
	amount *big.Int) {
	t.Helper()

	txOpts, err := backend.BindOpts(ctx, delegator)
	require.NoError(t, err)

	// First approve the strategy manager to "assign" the WETH to itself.
	tx, err := contracts.WETHToken.Approve(txOpts, contracts.StrategyManagerAddr, amount)
	require.NoError(t, err)

	receipt, err := backend.WaitMined(ctx, tx)
	require.NoError(t, err)
	require.Equal(t, ethtypes.ReceiptStatusSuccessful, receipt.Status)

	// Then deposit the WETH into the strategy (it will assign it to itself).
	tx, err = contracts.StrategyManager.DepositIntoStrategy(txOpts, contracts.WETHStrategyAddr, contracts.WETHTokenAddr, amount)
	require.NoError(t, err)

	receipt, err = backend.WaitMined(ctx, tx)
	require.NoError(t, err)
	require.Equal(t, ethtypes.ReceiptStatusSuccessful, receipt.Status)
}

func whiteListStrategy(
	t *testing.T,
	ctx context.Context,
	backend *ethbackend.Backend,
	contracts Contracts,
	ownerEigen common.Address) {
	t.Helper()

	txOpts, err := backend.BindOpts(ctx, ownerEigen)
	require.NoError(t, err)

	tx, err := contracts.StrategyManager.AddStrategiesToDepositWhitelist(
		txOpts, []common.Address{contracts.WETHStrategyAddr}, []bool{false})
	require.NoError(t, err)

	_, err = backend.WaitMined(ctx, tx)
	tutil.RequireNoError(t, err)
}

func fundAccount(t *testing.T, ctx context.Context, backend *ethbackend.Backend, funder, account common.Address) {
	t.Helper()
	tx, _, err := backend.Send(ctx, funder, txmgr.TxCandidate{
		To:    &account,
		Value: new(big.Int).Mul(big.NewInt(params.Ether), big.NewInt(10)), // 10 ETH
	})
	require.NoError(t, err)

	_, err = backend.WaitMined(ctx, tx)
	require.NoError(t, err)
}

func genPrivKey(t *testing.T) *ecdsa.PrivateKey {
	t.Helper()

	privKey, err := crypto.GenerateKey()
	require.NoError(t, err)

	return privKey
}

func registerOperator(t *testing.T, ctx context.Context, contracts Contracts, b *ethbackend.Backend, key *ecdsa.PrivateKey) {
	t.Helper()

	addr := crypto.PubkeyToAddress(key.PublicKey)
	dir := filepath.Join(t.TempDir(), addr.Hex())
	keystoreFile := filepath.Join(dir, "keystore.json")
	configFile := filepath.Join(dir, "config.yaml")

	_, chainID := b.Chain()

	const password = "12345678"

	err := eigenecdsa.WriteKey(keystoreFile, key, password)
	require.NoError(t, err)

	cfg := eigentypes.OperatorConfig{
		Operator: eigensdktypes.Operator{
			Address:                   addr.Hex(),
			EarningsReceiverAddress:   addr.Hex(),
			DelegationApproverAddress: eigensdktypes.ZeroAddress,
			StakerOptOutWindowBlocks:  0,
		},
		ELDelegationManagerAddress: contracts.DelegationManagerAddr.Hex(),
		EthRPCUrl:                  b.Address(),
		SignerConfig: eigentypes.SignerConfig{
			PrivateKeyStorePath: keystoreFile,
			SignerType:          eigentypes.LocalKeystoreSigner,
		},
		ChainId: *big.NewInt(int64(chainID)),
	}

	cfgYAML, err := cfg.MarshalYAML() // Convert into custom yaml struct first
	require.NoError(t, err)
	bz, err := yaml.Marshal(cfgYAML)
	require.NoError(t, err)
	require.NoError(t, os.WriteFile(configFile, bz, 0644))

	require.NoError(t, yaml.Unmarshal(bz, new(eigentypes.OperatorConfig))) // Ensure unmarshalling works

	// Override register options for testing.
	testOpts := func(deps *clicmd.RegDeps) {
		deps.Prompter = stubPrompter{password: password}
		deps.NewBackendFunc = func(_ string, _ uint64, _ time.Duration, _ ethclient.Client, _ ...*ecdsa.PrivateKey) (*ethbackend.Backend, error) {
			return b, nil // Have to provide the test backend for nonce management
		}
		deps.VerifyFunc = func(eigensdktypes.Operator) error {
			return nil // Skip operator verification since it requires non-localhost urls.
		}
	}

	regCfg := clicmd.RegConfig{
		ConfigFile: configFile,
		AVSAddr:    contracts.OmniAVSAddr.Hex(),
	}

	err = clicmd.Register(ctx, regCfg, testOpts)
	tutil.RequireNoError(t, err)
}

var _ eigenutils.Prompter = stubPrompter{}

type stubPrompter struct {
	eigenutils.Prompter
	password string
}

func (s stubPrompter) InputHiddenString(_, _ string, _ func(string) error) (string, error) {
	return s.password, nil
}

func toWei(amount int64) *big.Int {
	return new(big.Int).Mul(big.NewInt(amount), big.NewInt(params.Ether))
}

// mustGetABI returns the metadata's ABI as an abi.ABI type.
// It panics on error.
func mustGetABI(metadata *bind.MetaData) *abi.ABI {
	abi, err := metadata.GetAbi()
	if err != nil {
		panic(err)
	}

	return abi
}
