package avs_test

import (
	"context"
	"crypto/ecdsa"
	"encoding/json"
	"fmt"
	"math/big"
	"os"
	"path/filepath"
	"slices"
	"testing"
	"time"

	clicmd "github.com/omni-network/omni/cli/cmd"
	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/e2e/app/eoa"
	"github.com/omni-network/omni/e2e/app/static"
	"github.com/omni-network/omni/lib/anvil"
	"github.com/omni-network/omni/lib/contracts/avs"
	"github.com/omni-network/omni/lib/contracts/create3"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient"
	"github.com/omni-network/omni/lib/ethclient/ethbackend"
	"github.com/omni-network/omni/lib/netconf"
	"github.com/omni-network/omni/lib/tutil"
	"github.com/omni-network/omni/lib/txmgr"

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
	// eigen strategy manger, and test weth strategy.
	stratMngrAddr = common.HexToAddress("0xe1DA8919f262Ee86f9BE05059C9280142CF23f48")
	wethStratAddr = common.HexToAddress("0xdBD296711eC8eF9Aacb623ee3F1C0922dce0D7b2")

	zeroAddr = common.HexToAddress("0x0000000000000000000000000000000000000000")
)

func setup(t *testing.T) (context.Context, *ethbackend.Backend, Contracts, EOAS) {
	t.Helper()

	ctx := context.Background()

	ethCl, _, stop, err := anvil.Start(ctx, tutil.TempDir(t), chainID)
	require.NoError(t, err)
	t.Cleanup(stop)

	backend, err := ethbackend.NewAnvilBackend(chainName, chainID, blockPeriod, ethCl)
	require.NoError(t, err)

	eoas := makeEOAS(t, backend)

	_, _, err = create3.Deploy(ctx, netconf.Devnet, backend)
	require.NoError(t, err)

	addr, _, err := avs.Deploy(ctx, netconf.Devnet, backend)
	require.NoError(t, err)

	contracts, err := makeContracts(addr, devnetEigenDeployments(t), backend)
	require.NoError(t, err)

	checkIfContractsAreDeployed(t, ctx, ethCl, contracts)

	return ctx, backend, contracts, eoas
}

//nolint:paralleltest // Parallel tests not supported since we start docker containers.
func TestEigenAndOmniAVS(t *testing.T) {
	initialWETHFunding := toWei(1000)  // initial funding of WETH for operators and delegators.
	initialOperatorStake := toWei(100) // funding operator during registering to omniAVS.
	initialDelegatorStake := toWei(50) // initial stake that a delegator will do for an operator.

	ctx, backend, contracts, eoas := setup(t)

	ownerAVS := eoas.AVSOwner
	ownerEigen := eoas.EigenOwner
	operator1 := eoas.Operator1
	operator1Key := eoas.Operator1Key
	operator2 := eoas.Operator2
	operator2Key := eoas.Operator2Key
	delegator1 := eoas.Delgator1
	delegator2 := eoas.Delgator2

	// Combine 2 operators and 2 delegators
	operatorKeys := []*ecdsa.PrivateKey{operator1Key, operator2Key}
	operators := []common.Address{operator1, operator2}
	delegators := []common.Address{delegator1, delegator2}

	// Fund operators and delegators with ETH and WETH (using the pre-funded eigen owner account)
	for _, account := range slices.Concat(operators, delegators) {
		fundAccount(t, ctx, backend, ownerEigen, account)
		mintWETHToAddresses(t, ctx, backend, contracts, ownerEigen, initialWETHFunding, account)
	}

	// Register operators with EigenLayer
	for i, operator := range operators {
		rec, err := registerOperatorWithEigen(ctx, contracts, backend, operator, fmt.Sprintf("https://operator%d.com", i))
		tutil.RequireNoError(t, err)

		checkForOperatorRegisteredToELLog(t, ctx, backend, contracts, operator, rec.BlockNumber.Uint64())
	}

	// Whitelist the WETH strategy on eigen strategy manager (using the eigen owner account)
	whiteListStrategy(t, ctx, backend, contracts, ownerEigen)

	// Register operators to omni AVS with a stake more than minimum stake
	for i, operator := range operators {
		// Add the operator to omni avs allow list
		addOperatorToAllowList(t, ctx, contracts, ownerAVS, backend, operator)

		// Delegate some WETH to Eigen strategy manager
		delegateWETH(t, ctx, contracts, backend, operator, initialOperatorStake)

		registerOperatorCLI(t, ctx, contracts, backend, operatorKeys[i])

		assertOperatorRegistered(t, ctx, contracts, operator)
	}

	// delegate stake to operators and check their balance
	for i, delegator := range delegators {
		operator := operators[i]

		// Delegation is all-or-nothing, so delegate everything to the operator
		delegateToOperator(t, ctx, contracts, backend, delegator, operator)

		// Delegate actual tokens
		delegateWETH(t, ctx, contracts, backend, delegator, initialDelegatorStake)

		assertOperatorBalance(t, ctx, contracts, operator, initialOperatorStake, initialDelegatorStake)
	}

	// Undelegate delegator 1 and check if the stake is removed from operator
	err := undelegate(ctx, contracts, backend, delegators[0])
	require.NoError(t, err)
	assertOperatorBalance(t, ctx, contracts, operators[0], initialOperatorStake, big.NewInt(0))
}

func devnetEigenDeployments(t *testing.T) avs.EigenDeployments {
	t.Helper()

	var el avs.EigenDeployments
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
		AVSOwner:     eoa.MustAddress(netconf.Devnet, eoa.RoleAVSAdmin),
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

func makeContracts(
	avsAddr common.Address,
	eigen avs.EigenDeployments,
	backend *ethbackend.Backend,
) (Contracts, error) {
	delMan, err := bindings.NewDelegationManager(eigen.DelegationManager, backend)
	if err != nil {
		return Contracts{}, errors.Wrap(err, "delegation manager")
	}

	stratMan, err := bindings.NewStrategyManager(stratMngrAddr, backend)
	if err != nil {
		return Contracts{}, errors.Wrap(err, "strategy manager")
	}

	wethStrategy, err := bindings.NewStrategyBase(wethStratAddr, backend)
	if err != nil {
		return Contracts{}, errors.Wrap(err, "weth strategy")
	}

	wethTokenAddr, err := wethStrategy.UnderlyingToken(&bind.CallOpts{})
	if err != nil {
		return Contracts{}, errors.Wrap(err, "underlying token")
	}

	wethToken, err := bindings.NewMockERC20(wethTokenAddr, backend)
	if err != nil {
		return Contracts{}, errors.Wrap(err, "weth token")
	}

	avsDir, err := bindings.NewAVSDirectory(eigen.AVSDirectory, backend)
	if err != nil {
		return Contracts{}, errors.Wrap(err, "avs directory")
	}

	avs, err := bindings.NewOmniAVS(avsAddr, backend)
	if err != nil {
		return Contracts{}, errors.Wrap(err, "omni avs")
	}

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
	}, nil
}

func registerOperatorWithEigen(ctx context.Context, contracts Contracts, backend *ethbackend.Backend, operator common.Address, metadataURI string) (*ethtypes.Receipt, error) {
	operatorDetails := bindings.IDelegationManagerOperatorDetails{
		EarningsReceiver:         operator,
		DelegationApprover:       common.Address{},
		StakerOptOutWindowBlocks: uint32(0), // Currently unused by Eigen
	}

	txOpts, err := backend.BindOpts(ctx, operator)
	if err != nil {
		return nil, err
	}

	tx, err := contracts.DelegationManager.RegisterAsOperator(txOpts, operatorDetails, metadataURI)
	if err != nil {
		return nil, errors.Wrap(err, "register as operator")
	}

	receipt, err := backend.WaitMined(ctx, tx)
	if err != nil {
		return nil, errors.Wrap(err, "wait mined")
	}

	return receipt, nil
}

func undelegate(ctx context.Context, contracts Contracts, backend *ethbackend.Backend, delegator common.Address) error {
	txOpts, err := backend.BindOpts(ctx, delegator)
	if err != nil {
		return err
	}

	tx, err := contracts.DelegationManager.Undelegate(txOpts, delegator)
	if err != nil {
		return errors.Wrap(err, "deposit into strategy")
	}

	if _, err = backend.WaitMined(ctx, tx); err != nil {
		return err
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

func checkIfContractsAreDeployed(
	t *testing.T,
	ctx context.Context,
	ethCl ethclient.Client,
	contracts Contracts) {
	t.Helper()

	// check if the contract has code in its respective contract addresses
	checkIfCodePresent(t, ctx, ethCl, contracts.OmniAVSAddr)
	checkIfCodePresent(t, ctx, ethCl, contracts.DelegationManagerAddr)
	checkIfCodePresent(t, ctx, ethCl, contracts.StrategyManagerAddr)
	checkIfCodePresent(t, ctx, ethCl, contracts.WETHStrategyAddr)
	checkIfCodePresent(t, ctx, ethCl, contracts.WETHTokenAddr)
	checkIfCodePresent(t, ctx, ethCl, contracts.AVSDirectoryAddr)
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

		require.Equal(t, amount, wETHBalance(t, ctx, contracts, addr))
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

func addOperatorToAllowList(
	t *testing.T,
	ctx context.Context,
	contracts Contracts,
	ownerAVS common.Address,
	backend *ethbackend.Backend,
	operator common.Address) {
	t.Helper()

	txOpts, err := backend.BindOpts(ctx, ownerAVS)
	require.NoError(t, err)
	tx, err := contracts.OmniAVS.AddToAllowlist(txOpts, operator)
	require.NoError(t, err)
	_, err = backend.WaitMined(ctx, tx)
	require.NoError(t, err)
}

func checkForOperatorRegisteredToELLog(
	t *testing.T,
	ctx context.Context,
	backend *ethbackend.Backend,
	contracts Contracts,
	operator common.Address,
	height uint64) {
	t.Helper()

	filterer, err := bindings.NewDelegationManagerFilterer(contracts.DelegationManagerAddr, backend)
	require.NoError(t, err)

	filterOpts := bind.FilterOpts{
		Start:   height,
		End:     &height,
		Context: ctx,
	}

	iter, err := filterer.FilterOperatorRegistered(&filterOpts, []common.Address{operator})
	require.NoError(t, err)

	for iter.Next() {
		e := iter.Event
		require.Equal(t, e.Operator, operator,
			"operator is not matching")
		require.Equal(t, e.OperatorDetails.DelegationApprover, zeroAddr,
			"delegation approver is not matching")
		require.Equal(t, e.OperatorDetails.EarningsReceiver, operator,
			"earnings receiver is not matching")

		break // there should be only one log event in this block
	}
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

func delegateToOperator(t *testing.T,
	ctx context.Context,
	contracts Contracts,
	backend *ethbackend.Backend,
	delegator common.Address,
	operator common.Address) {
	t.Helper()

	txOpts, err := backend.BindOpts(ctx, delegator)
	require.NoError(t, err)

	// TODO(corver): I do not think this is required since operator doesn't have a delegation approver.
	approverSignatureAndExpiry := bindings.ISignatureUtilsSignatureWithExpiry{
		Signature: []byte{0},
		Expiry:    big.NewInt(time.Now().Add(time.Hour).Unix()),
	}
	approverSalt := crypto.Keccak256Hash(delegator.Bytes()) // Not sure whether any salt is ok?

	tx, err := contracts.DelegationManager.DelegateTo(txOpts, operator, approverSignatureAndExpiry, approverSalt)
	require.NoError(t, err)

	_, err = backend.WaitMined(ctx, tx)
	require.NoError(t, err)
}

func assertOperatorBalance(
	t *testing.T,
	ctx context.Context,
	contracts Contracts,
	operator common.Address,
	oprStake *big.Int,
	delStake *big.Int) {
	t.Helper()

	callOpts := bind.CallOpts{
		From:    operator,
		Context: ctx,
	}
	operators, err := contracts.OmniAVS.Operators(&callOpts)
	require.NoError(t, err)

	found := false
	for _, op := range operators {
		if op.Addr == operator {
			found = true

			// when one of staked/delegated is zero, require.Equal(*big.Int, *big.Int) fails
			// so we compare strings
			require.Equal(t, oprStake.String(), op.Staked.String())
			require.Equal(t, delStake.String(), op.Delegated.String())

			break
		}
	}
	require.True(t, found)
}

func wETHBalance(t *testing.T,
	ctx context.Context,
	contracts Contracts,
	account common.Address) *big.Int {
	t.Helper()

	callOpts := bind.CallOpts{
		From:    account,
		Context: ctx,
	}

	// get the balance of WETH for a given account
	balance, err := contracts.WETHToken.BalanceOf(&callOpts, account)
	require.NoError(t, err)

	return balance
}

func checkIfCodePresent(
	t *testing.T,
	ctx context.Context,
	ethCl ethclient.Client,
	addr common.Address) {
	t.Helper()

	codeBytes, err := ethCl.CodeAt(ctx, addr, nil)
	require.NoError(t, err)
	require.NotEmpty(t, len(codeBytes))
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

func registerOperatorCLI(t *testing.T, ctx context.Context, contracts Contracts, b *ethbackend.Backend, key *ecdsa.PrivateKey) {
	t.Helper()

	addr := crypto.PubkeyToAddress(key.PublicKey)
	dir := filepath.Join(t.TempDir(), addr.Hex())
	keystoreFile := filepath.Join(dir, "keystore.json")
	configFile := filepath.Join(dir, "config.yaml")

	_, chainID := b.Chain()

	const password = "12345678"

	err := eigenecdsa.WriteKey(keystoreFile, key, password)
	require.NoError(t, err)

	cfg := eigentypes.OperatorConfigNew{
		Operator: eigensdktypes.Operator{
			Address:                   addr.Hex(),
			EarningsReceiverAddress:   addr.Hex(),
			DelegationApproverAddress: eigensdktypes.ZeroAddress,
			StakerOptOutWindowBlocks:  0,
		},
		ELDelegationManagerAddress: contracts.DelegationManagerAddr.Hex(),
		EthRPCUrl:                  b.Address(),
		PrivateKeyStorePath:        keystoreFile,
		SignerType:                 eigentypes.LocalKeystoreSigner,
		ChainId:                    *big.NewInt(int64(chainID)),
	}

	bz, err := yaml.Marshal(cfg)
	require.NoError(t, err)
	require.NoError(t, os.WriteFile(configFile, bz, 0644))

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
