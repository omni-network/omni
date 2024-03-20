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
	"strings"
	"testing"
	"time"

	clicmd "github.com/omni-network/omni/cli/cmd"
	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/e2e/app/static"
	"github.com/omni-network/omni/e2e/tutil"
	"github.com/omni-network/omni/lib/avs"
	"github.com/omni-network/omni/lib/avs/anvil"
	"github.com/omni-network/omni/lib/ethclient"
	"github.com/omni-network/omni/lib/ethclient/ethbackend"
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
	omniChainID = 10
	blockPeriod = time.Second

	// pk used to deploy omniAVS contracts.
	omniDeployPk = "0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80"

	// pk used to deploy EigenLayer contracts, anvil account 9.
	eigenDeployPk = "0x2a871d0798f97d79848a013d4936a73bf4cc922c825d33c1cf7073dff6d409c6"

	// the zero-address string.
	zeroAddr = "0x0000000000000000000000000000000000000000"
)

func setup(t *testing.T) (context.Context, ethclient.Client, *avs.Deployer) {
	t.Helper()

	ctx := context.Background()

	ethCl, stop, err := anvil.Start(ctx, t.TempDir(), chainID)
	require.NoError(t, err)
	t.Cleanup(stop)

	// Create a backend for deploying AVS contract using the omniDeployPk.
	var deployments avs.EigenDeployments
	err = json.Unmarshal(static.GetElDeployments(), &deployments)
	require.NoError(t, err)

	// Deploy the OmniAVS contract using omniDeployKey
	ad := avs.NewDeployer(
		avs.DefaultTestAVSConfig(deployments),
		deployments,
		common.Address{},
		omniChainID,
	)

	return ctx, ethCl, ad
}

//nolint:paralleltest // Parallel tests not supported since we start docker containers.
func TestEigenAndOmniAVS(t *testing.T) {
	initialWETHFunding := toWei(1000)  // initial funding of WETH for operators and delegators.
	initialOperatorStake := toWei(100) // funding operator during registering to omniAVS.
	initialDelegatorStake := toWei(50) // initial stake that a delegator will do for an operator.

	ctx, ethCl, avsDeploy := setup(t)

	backend, err := ethbackend.NewBackend(chainName, chainID, blockPeriod, ethCl)
	require.NoError(t, err)

	// Add the two owner accounts to the backend
	ownerAVS, err := backend.AddAccount(mustHexToKey(omniDeployPk))
	require.NoError(t, err)
	ownerEigen, err := backend.AddAccount(mustHexToKey(eigenDeployPk))
	require.NoError(t, err)

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

	// Deploy the AVS contract and get the contracts struct to interact with it.
	require.NoError(t, avsDeploy.Deploy(ctx, backend, ownerAVS))
	contracts, err := avsDeploy.Contracts(backend)
	require.NoError(t, err)

	// Combine 2 operators and 2 delegators
	operatorKeys := []*ecdsa.PrivateKey{operator1Key, operator2Key}
	operators := []common.Address{operator1, operator2}
	delegators := []common.Address{delegator1, delegator2}

	// Check if contracts are deployed and configured properly
	checkIfContractsAreDeployed(t, ctx, ethCl, contracts)

	// Fund operators and delegators with ETH and WETH (using the pre-funded eigen owner account)
	for _, account := range slices.Concat(operators, delegators) {
		fundAccount(t, ctx, backend, ownerEigen, account)
		mintWETHToAddresses(t, ctx, backend, contracts, ownerEigen, initialWETHFunding, account)
	}

	// Register operators with EigenLayer
	for i, operator := range operators {
		rec, err := avs.RegisterOperatorWithEigen(ctx, contracts, backend, operator, fmt.Sprintf("https://operator%d.com", i))
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
	err = avs.Undelegate(ctx, contracts, backend, delegators[0])
	require.NoError(t, err)
	assertOperatorBalance(t, ctx, contracts, operators[0], initialOperatorStake, big.NewInt(0))

	// Deregister operators
	for _, operator := range operators {
		err := avs.DeregisterOperatorFromAVS(ctx, contracts, backend, operator)
		require.NoError(t, err)
	}
}

func assertOperatorRegistered(t *testing.T, ctx context.Context, contracts avs.Contracts, operator common.Address) {
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
	contracts avs.Contracts) {
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
	contracts avs.Contracts,
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
	contracts avs.Contracts,
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
	contracts avs.Contracts,
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
	contracts avs.Contracts,
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
		require.Equal(t, e.OperatorDetails.DelegationApprover, common.HexToAddress(zeroAddr),
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
	contracts avs.Contracts,
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
	contracts avs.Contracts,
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
	contracts avs.Contracts,
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
	contracts avs.Contracts,
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

func mustHexToKey(privKeyHex string) *ecdsa.PrivateKey {
	privKey, err := crypto.HexToECDSA(strings.TrimPrefix(privKeyHex, "0x"))
	if err != nil {
		panic(err)
	}

	return privKey
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

func registerOperatorCLI(t *testing.T, ctx context.Context, contracts avs.Contracts, b *ethbackend.Backend, key *ecdsa.PrivateKey) {
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
