package e2e_test

import (
	"context"
	"crypto/ecdsa"
	"math/big"
	"strings"
	"testing"
	"time"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/ethclient"
	"github.com/omni-network/omni/test/e2e/types"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/stretchr/testify/require"
)

const (
	// pk used to deploy omniAVS contracts (anvil account 0).
	omniDeployPk = "0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80"

	// pk used to deploy EigenLayer contracts ( anvil account 9).
	elDeployPk = "0x2a871d0798f97d79848a013d4936a73bf4cc922c825d33c1cf7073dff6d409c6"

	// ETH funding account (anvil account 8)
	ethFundingPk = "0xdbda1821b80551c9d65939329250298aa3472ba22feea921c0cf5d620ea67b97"

	// URI, which stores the metadata for an operator.
	operator1MetaDataURI = "https://www.operator1.com"

	// URI, which stores the metadata for an operator.
	operator2MetaDataURI = "https://www.operator2.com"

	// URI, which stores the metadata for an operator.
	operator3MetaDataURI = "https://www.operator3.com"

	// minimum self stake for an operator to get registered in omniAVS.
	MinStateForOperatorInEigenLayer = 10

	// maximum number of operators allowed in omniAVS.
	MaxNumberOfOperators = 200

	// initial funding of WETH for operators and delegators.
	InitialWETHFunding = 1000

	// funding operator during registering to omniAVS.
	InitialOperatorStake = 100

	// initial stake that a delegator will do for an operator.
	InitialDelegatorStake = 50

	// the zero-address string.
	zeroAddr = "0x0000000000000000000000000000000000000000"
)

func TestEigenAndOmniAVS(t *testing.T) {
	t.Parallel()
	testAVS(t, func(t *testing.T, avs AVS, deployInfo map[types.ContractName]types.DeployInfo) {
		t.Helper()
		ctx := context.Background()

		delMgrAddr := deployInfo[types.ContractELDelegationManager].Address
		stratManAddr := deployInfo[types.ContractELStrategyManager].Address
		wethStratAddr := deployInfo[types.ContractELWETHStrategy].Address
		wethTokenAddr := getTokenAddr(t, avs)
		omniAVSAddr := deployInfo[types.ContractOmniAVS].Address
		//podMgrAddr := deployInfo[types.ContractELPodManager].Address
		deployInfo[types.ContractELWETH] = types.DeployInfo{
			Address: wethTokenAddr,
		}

		// create new operators, delegators and fund them with ETH
		operator1Addr, opr1Addr, opr1Pk := createAccount(t, ctx, avs.Client)
		operator2Addr, opr2Addr, opr2Pk := createAccount(t, ctx, avs.Client)
		operator3Addr, opr3Addr, opr3Pk := createAccount(t, ctx, avs.Client)
		_, del1Addr, del1Pk := createAccount(t, ctx, avs.Client)
		_, del2Addr, del2Pk := createAccount(t, ctx, avs.Client)
		_, del3Addr, del3Pk := createAccount(t, ctx, avs.Client)

		// check if contracts are deployed and configured properly
		checkIfContractsAreDeployed(t, ctx, avs, deployInfo)
		//configOmniAVS(t, ctx, avs)

		// register operators with EigenLayer
		registerOperatorWithEL(t, ctx, avs, delMgrAddr, opr1Pk, operator1Addr, operator1MetaDataURI)
		registerOperatorWithEL(t, ctx, avs, delMgrAddr, opr2Pk, operator2Addr, operator2MetaDataURI)
		registerOperatorWithEL(t, ctx, avs, delMgrAddr, opr3Pk, operator3Addr, operator3MetaDataURI)

		// fund operators and delegators with WETH
		accounts := []common.Address{
			opr1Addr,
			opr2Addr,
			opr3Addr,
			del1Addr,
			del2Addr,
			del3Addr,
		}
		fundAccountsWithWETH(t, ctx, avs, InitialWETHFunding, accounts)

		// register operators to omni AVS with a stake more than minimum stake
		whiteListStrategy(t, ctx, avs, wethStratAddr)
		registeringOperatorsWithAVS(t, ctx, avs, stratManAddr, wethStratAddr, wethTokenAddr, omniAVSAddr,
			opr1Addr, opr1Pk, InitialOperatorStake)
		registeringOperatorsWithAVS(t, ctx, avs, stratManAddr, wethStratAddr, wethTokenAddr, omniAVSAddr,
			opr2Addr, opr2Pk, InitialOperatorStake)
		registeringOperatorsWithAVS(t, ctx, avs, stratManAddr, wethStratAddr, wethTokenAddr, omniAVSAddr,
			opr3Addr, opr3Pk, InitialOperatorStake)

		// delegate stake to operators and check their balance
		delegateWETHToOperator(t, ctx, avs, del1Addr, del1Pk, opr1Addr)
		delegateWETHToStrategy(t, ctx, avs, del1Pk, stratManAddr, wethStratAddr, wethTokenAddr, InitialDelegatorStake)
		checkOperatorBalance(t, ctx, avs, opr1Addr, opr1Pk, uint64(InitialOperatorStake), uint64(InitialDelegatorStake))

		delegateWETHToOperator(t, ctx, avs, del2Addr, del2Pk, opr2Addr)
		delegateWETHToStrategy(t, ctx, avs, del2Pk, stratManAddr, wethStratAddr, wethTokenAddr, InitialDelegatorStake)
		checkOperatorBalance(t, ctx, avs, opr2Addr, opr2Pk, uint64(InitialOperatorStake), uint64(InitialDelegatorStake))

		stakeETHToDelegator(t, ctx, avs, del3Addr, del3Pk)
		//delegateETHToStrategy(t, ctx, avs, del3Pk, stratManAddr, wethStratAddr, wethTokenAddr, InitialDelegatorStake)
		checkOperatorBalance(t, ctx, avs, opr3Addr, opr3Pk, uint64(InitialOperatorStake), uint64(InitialDelegatorStake))

		// undelegate delegator 1 and check if the stake is removed from operator
		undelegateWETHForDelegattor(t, ctx, avs, del1Addr, del1Pk)
		checkOperatorBalance(t, ctx, avs, opr1Addr, opr1Pk, uint64(InitialOperatorStake), uint64(0))

		// unregister operators
		unregisterOperatorFromAVS(t, ctx, avs, opr1Addr, opr1Pk)
		unregisterOperatorFromAVS(t, ctx, avs, opr2Addr, opr2Pk)
	})
}

/*
 * internal/ helper functions.
 */

func getTokenAddr(t *testing.T, avs AVS) common.Address {
	t.Helper()

	callOpts := bind.CallOpts{}
	token, err := avs.WETHStrategyContract.UnderlyingToken(&callOpts)
	require.NoError(t, err)

	return token
}

func createAccount(t *testing.T, ctx context.Context, client ethclient.Client) (string,
	common.Address, *ecdsa.PrivateKey) {
	t.Helper()

	privateKey, err := crypto.GenerateKey()
	require.NoError(t, err)

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	require.True(t, ok)
	address := crypto.PubkeyToAddress(*publicKeyECDSA)

	// fund the account with ETH for fees
	amount := big.NewInt(100000000000000000) // 1 ETH
	transferETHTo(t, ctx, address.String(), client, amount)

	return address.String(), address, privateKey
}

func transferETHTo(t *testing.T, ctx context.Context, addr string, client ethclient.Client, amount *big.Int) {
	t.Helper()
	pk := mustHexToKey("0x4bbbf85ce3377467afe5d46f804f221813b2bb87f24d81f60f1fcdbf7cbf4356")
	adr := common.HexToAddress("0x14dC79964da2C08b23698B3D3cc7Ca32193d9955")

	nonce, err := client.PendingNonceAt(ctx, adr)
	require.NoError(t, err)
	value := amount // big.NewInt(100000000000000000) // 1 ETH
	gasLimit := uint64(21000)
	tip := big.NewInt(2000000000)
	feeCap := big.NewInt(20000000000)
	require.NoError(t, err)

	var data []byte
	chainID, err := client.ChainID(ctx)
	require.NoError(t, err)

	toAddress := common.HexToAddress(addr)
	tx := ethtypes.NewTx(&ethtypes.DynamicFeeTx{
		ChainID:   chainID,
		Nonce:     nonce,
		GasFeeCap: feeCap,
		GasTipCap: tip,
		Gas:       gasLimit,
		To:        &toAddress,
		Value:     value,
		Data:      data,
	})

	signedTx, err := ethtypes.SignTx(tx, ethtypes.LatestSignerForChainID(chainID), pk)
	require.NoError(t, err)

	err = client.SendTransaction(ctx, signedTx)
	require.NoError(t, err)
	_, err = bind.WaitMined(ctx, client, signedTx)
	require.NoError(t, err)
}

func checkIfContractsAreDeployed(
	t *testing.T,
	ctx context.Context,
	avs AVS,
	deployInfo map[types.ContractName]types.DeployInfo) {
	t.Helper()

	blockNumber, err := avs.Client.BlockNumber(ctx)
	require.NoError(t, err)

	// check if the contract has code in its respective contract addresses
	checkIfCodePresent(t, ctx, avs, deployInfo, blockNumber, types.ContractOmniAVS)
	checkIfCodePresent(t, ctx, avs, deployInfo, blockNumber, types.ContractELDelegationManager)
	checkIfCodePresent(t, ctx, avs, deployInfo, blockNumber, types.ContractELAVSDirectory)
	checkIfCodePresent(t, ctx, avs, deployInfo, blockNumber, types.ContractELStrategyManager)
	checkIfCodePresent(t, ctx, avs, deployInfo, blockNumber, types.ContractELWETHStrategy)
	checkIfCodePresent(t, ctx, avs, deployInfo, blockNumber, types.ContractELWETH)
	checkIfCodePresent(t, ctx, avs, deployInfo, blockNumber, types.ContractELPodManager)
}

//func configOmniAVS(
//	t *testing.T,
//	ctx context.Context,
//	avs AVS) {
//	t.Helper()
//
//	// set min stake
//	omniDepPk := mustHexToKey(omniDeployPk)
//	txOpts, err := bind.NewKeyedTransactorWithChainID(omniDepPk, big.NewInt(int64(avs.Chain.ID)))
//	require.NoError(t, err)
//	txOpts.Context = ctx
//	minStake := big.NewInt(MinStateForOperatorInEigenLayer)
//	tx, err := avs.AVSContract.SetMinOperatorStake(txOpts, minStake)
//	require.NoError(t, err)
//	_, err = bind.WaitMined(ctx, avs.Client, tx)
//	require.NoError(t, err)
//
//	// check if min stake is set properly
//	callOpts := bind.CallOpts{}
//	operatorStake, err := avs.AVSContract.MinOperatorStake(&callOpts)
//	require.NoError(t, err)
//	require.Equal(t, minStake.Uint64(), operatorStake.Uint64())
//
//	// set operator count
//	tx, err = avs.AVSContract.SetMaxOperatorCount(txOpts, uint32(MaxNumberOfOperators))
//	require.NoError(t, err)
//	_, err = bind.WaitMined(ctx, avs.Client, tx)
//	require.NoError(t, err)
//
//	// check if operator count is set properly
//	opCount, err := avs.AVSContract.MaxOperatorCount(&callOpts)
//	require.NoError(t, err)
//	require.Equal(t, uint32(MaxNumberOfOperators), opCount)
//}

func registerOperatorWithEL(
	t *testing.T,
	ctx context.Context,
	avs AVS,
	delMgrAddr common.Address,
	oprPk *ecdsa.PrivateKey,
	oprAddrStr string,
	oprMetaDataURI string) {
	t.Helper()

	// register operator to EL
	txOpts, err := bind.NewKeyedTransactorWithChainID(oprPk, big.NewInt(int64(avs.Chain.ID)))
	require.NoError(t, err)
	txOpts.Context = ctx
	operatorDetails := getOperatorDetails(oprAddrStr)
	tx, err := avs.DelegationManagerContract.RegisterAsOperator(txOpts, operatorDetails, oprMetaDataURI)
	require.NoError(t, err)
	receipt, err := bind.WaitMined(ctx, avs.Client, tx)
	require.NoError(t, err)

	// get block where the operator was registered
	block, err := avs.Client.BlockByNumber(ctx, receipt.BlockNumber)
	require.NoError(t, err)

	// check logs if the operator is registered properly
	checkForOperatorRegisteredToELLog(t, ctx, avs,
		common.HexToAddress(oprAddrStr),
		delMgrAddr,
		block.Header().Number.Uint64())
}

func fundAccountsWithWETH(
	t *testing.T,
	ctx context.Context,
	avs AVS,
	amount int64,
	addresses []common.Address) {
	t.Helper()

	// fund the operators/delegators and check if it funded properly
	for _, addr := range addresses {
		fundAccountWithWETH(t, ctx, avs, addr, amount)
		require.Equal(t, uint64(amount), wETHBalance(t, ctx, avs, addr))
	}
}

func registeringOperatorsWithAVS(
	t *testing.T,
	ctx context.Context,
	avs AVS,
	stratManAddr common.Address,
	wethStratAddr common.Address,
	wethTokenAddr common.Address,
	omniAVSAddr common.Address,
	oprAddr common.Address,
	oprPk *ecdsa.PrivateKey,
	selfStakeAmount int64) {
	t.Helper()

	// add the operator to omni avs allow list
	addOperatorToAllowList(t, ctx, avs, oprAddr)

	// delegate min stake from operator to strategy
	delegateWETHToStrategy(t, ctx, avs, oprPk, stratManAddr, wethStratAddr, wethTokenAddr, selfStakeAmount)

	// register the operator with omni avs
	registerOperatorToAVS(t, ctx, avs, oprAddr, omniAVSAddr, oprPk)
}

func addOperatorToAllowList(
	t *testing.T,
	ctx context.Context,
	avs AVS,
	oprAddr common.Address) {
	t.Helper()

	omniDepPk := mustHexToKey(omniDeployPk)
	txOpts, err := bind.NewKeyedTransactorWithChainID(omniDepPk, big.NewInt(int64(avs.Chain.ID)))
	require.NoError(t, err)
	txOpts.Context = ctx

	tx, err := avs.AVSContract.AddToAllowlist(txOpts, oprAddr)
	require.NoError(t, err)
	_, err = bind.WaitMined(ctx, avs.Client, tx)
	require.NoError(t, err)
}

func registerOperatorToAVS(t *testing.T,
	ctx context.Context,
	avs AVS,
	oprAddr common.Address,
	omniAVSAddr common.Address,
	oprPk *ecdsa.PrivateKey) {
	t.Helper()

	// calculate operator signature and digest hash
	blockNumber, err := avs.Client.BlockNumber(ctx)
	require.NoError(t, err)
	block, err := avs.Client.BlockByNumber(ctx, big.NewInt(int64(blockNumber)))
	require.NoError(t, err)
	operatorSignatureWithSaltAndExpiry := bindings.ISignatureUtilsSignatureWithSaltAndExpiry{
		Signature: []byte{0},
		Salt:      crypto.Keccak256Hash(oprAddr.Bytes()),
		Expiry:    big.NewInt(int64(block.Time()) + int64(24*time.Hour)),
	}
	callOpts := bind.CallOpts{}
	digestHash, err := avs.AVSDirectory.CalculateOperatorAVSRegistrationDigestHash(&callOpts,
		oprAddr, omniAVSAddr, operatorSignatureWithSaltAndExpiry.Salt, operatorSignatureWithSaltAndExpiry.Expiry)
	require.NoError(t, err)

	operatorSignatureWithSaltAndExpiry.Signature, err = crypto.Sign(digestHash[:32], oprPk)
	require.NoError(t, err)
	if len(operatorSignatureWithSaltAndExpiry.Signature) != 65 {
		require.NoError(t, errors.New("invalid signature length"))
	}
	operatorSignatureWithSaltAndExpiry.Signature[64] += 27
	txOpts, err := bind.NewKeyedTransactorWithChainID(oprPk, big.NewInt(int64(avs.Chain.ID)))
	require.NoError(t, err)
	txOpts.Context = ctx
	tx, err := avs.AVSContract.RegisterOperatorToAVS(txOpts, oprAddr, operatorSignatureWithSaltAndExpiry)
	require.NoError(t, err)
	_, err = bind.WaitMined(ctx, avs.Client, tx)
	require.NoError(t, err)
}

func getOperatorDetails(addr string) bindings.IDelegationManagerOperatorDetails {
	return bindings.IDelegationManagerOperatorDetails{
		EarningsReceiver:         common.HexToAddress(addr),
		DelegationApprover:       common.HexToAddress(zeroAddr),
		StakerOptOutWindowBlocks: uint32(1),
	}
}

func checkForOperatorRegisteredToELLog(
	t *testing.T,
	ctx context.Context,
	avs AVS,
	operatorAddr common.Address,
	contractAddr common.Address,
	height uint64) {
	t.Helper()

	filterer, err := bindings.NewDelegationManagerFilterer(contractAddr, avs.Client)
	require.NoError(t, err)

	filterOpts := bind.FilterOpts{
		Start:   height,
		End:     &height,
		Context: ctx,
	}

	operator := make([]common.Address, 0)
	operator = append(operator, operatorAddr)
	iter, err := filterer.FilterOperatorRegistered(&filterOpts, operator)
	require.NoError(t, err)

	for iter.Next() {
		e := iter.Event
		require.Equal(t, e.Operator, operatorAddr,
			"operator is not matching")
		require.Equal(t, e.OperatorDetails.DelegationApprover, common.HexToAddress(zeroAddr),
			"delegation approver is not matching")
		require.Equal(t, e.OperatorDetails.EarningsReceiver, operatorAddr,
			"earnings receiver is not matching")

		break // there should be only one log event in this block
	}
}

func fundAccountWithWETH(t *testing.T,
	ctx context.Context,
	avs AVS,
	accountToFund common.Address,
	amount int64) {
	t.Helper()

	elDepPk := mustHexToKey(elDeployPk)
	txOpts, err := bind.NewKeyedTransactorWithChainID(elDepPk, big.NewInt(int64(avs.Chain.ID)))
	require.NoError(t, err)
	txOpts.Context = ctx

	tx, err := avs.WETHTokenContract.Mint(txOpts, accountToFund, big.NewInt(amount))
	require.NoError(t, err)
	_, err = bind.WaitMined(ctx, avs.Client, tx)
	require.NoError(t, err)
}

func whiteListStrategy(
	t *testing.T,
	ctx context.Context,
	avs AVS,
	wethStratAddr common.Address) {
	t.Helper()

	elDepPk := mustHexToKey(elDeployPk)
	txOpts, err := bind.NewKeyedTransactorWithChainID(elDepPk, big.NewInt(int64(avs.Chain.ID)))
	require.NoError(t, err)
	txOpts.Context = ctx
	tx, err := avs.StrategyManagerContract.AddStrategiesToDepositWhitelist(
		txOpts, []common.Address{wethStratAddr}, []bool{false})
	require.NoError(t, err)
	_, err = bind.WaitMined(ctx, avs.Client, tx)
	require.NoError(t, err)
}

func delegateWETHToOperator(t *testing.T,
	ctx context.Context,
	avs AVS,
	delegator common.Address,
	delegatorPk *ecdsa.PrivateKey,
	operator common.Address) {
	t.Helper()

	txOpts, err := bind.NewKeyedTransactorWithChainID(delegatorPk, big.NewInt(int64(avs.Chain.ID)))
	require.NoError(t, err)
	txOpts.Context = ctx

	blockNumber, err := avs.Client.BlockNumber(ctx)
	require.NoError(t, err)
	block, err := avs.Client.BlockByNumber(ctx, big.NewInt(int64(blockNumber)))
	require.NoError(t, err)
	approverSignatureAndExpiry := bindings.ISignatureUtilsSignatureWithExpiry{
		Signature: []byte{0},
		Expiry:    big.NewInt(int64(block.Time()) + int64(24*time.Hour)),
	}
	approverSalt := crypto.Keccak256Hash(delegator.Bytes())
	tx, err := avs.DelegationManagerContract.DelegateTo(txOpts, operator, approverSignatureAndExpiry, approverSalt)
	require.NoError(t, err)
	_, err = bind.WaitMined(ctx, avs.Client, tx)
	require.NoError(t, err)
}

func stakeETHToDelegator(t *testing.T,
	ctx context.Context,
	avs AVS,
	delegator common.Address,
	delegatorPk *ecdsa.PrivateKey) {
	t.Helper()

	// first transfer some funds to delegator to stake
	amount := big.NewInt(32000000000000000000) // 32 ETH
	transferETHTo(t, ctx, delegator.String(), avs.Client, amount)

	// then delegate the ETH to operator
	txOpts, err := bind.NewKeyedTransactorWithChainID(delegatorPk, big.NewInt(int64(avs.Chain.ID)))
	require.NoError(t, err)
	txOpts.Context = ctx
	txOpts.Value = amount
	txOpts.From = delegator

	depositRoot := []byte(zeroAddr)[:32]
	signature, err := crypto.Sign(depositRoot, delegatorPk)
	require.NoError(t, err)

	pubKey, err := crypto.SigToPub(depositRoot, signature)
	require.NoError(t, err)
	pubKeyBytes := crypto.FromECDSAPub(pubKey)

	tx, err := avs.EigenPodManagerContract.Stake(txOpts, pubKeyBytes, signature, [32]byte(depositRoot))
	require.NoError(t, err)
	_, err = bind.WaitMined(ctx, avs.Client, tx)
	require.NoError(t, err)
}

func delegateWETHToStrategy(
	t *testing.T,
	ctx context.Context,
	avs AVS,
	delegatorPk *ecdsa.PrivateKey,
	strategyManager common.Address,
	strategy common.Address,
	token common.Address,
	amount int64) {
	t.Helper()

	txOpts, err := bind.NewKeyedTransactorWithChainID(delegatorPk, big.NewInt(int64(avs.Chain.ID)))
	require.NoError(t, err)
	txOpts.Context = ctx

	tx, err := avs.WETHTokenContract.Approve(txOpts, strategyManager, big.NewInt(amount))
	require.NoError(t, err)
	_, err = bind.WaitMined(ctx, avs.Client, tx)
	require.NoError(t, err)

	txOpts, err = bind.NewKeyedTransactorWithChainID(delegatorPk, big.NewInt(int64(avs.Chain.ID)))
	require.NoError(t, err)
	txOpts.Context = ctx

	tx, err = avs.StrategyManagerContract.DepositIntoStrategy(txOpts, strategy, token, big.NewInt(amount))
	require.NoError(t, err)
	_, err = bind.WaitMined(ctx, avs.Client, tx)
	require.NoError(t, err)
}

func checkOperatorBalance(
	t *testing.T,
	ctx context.Context,
	avs AVS,
	operator common.Address,
	operatorPk *ecdsa.PrivateKey,
	oprState uint64,
	delStake uint64) {
	t.Helper()

	txOpts, err := bind.NewKeyedTransactorWithChainID(operatorPk, big.NewInt(int64(avs.Chain.ID)))
	require.NoError(t, err)
	txOpts.Context = ctx

	callOpts := bind.CallOpts{
		From:    operator,
		Context: ctx,
	}
	vals, err := avs.AVSContract.Operators(&callOpts)
	require.NoError(t, err)

	found := false
	for _, val := range vals {
		if val.Addr.String() == operator.String() {
			found = true
			require.Equal(t, oprState, val.Staked.Uint64())
			require.Equal(t, delStake, val.Delegated.Uint64())

			break
		}
	}
	require.True(t, found)
}

func wETHBalance(t *testing.T,
	ctx context.Context,
	avs AVS,
	account common.Address) uint64 {
	t.Helper()

	callOpts := bind.CallOpts{
		From:    account,
		Context: ctx,
	}

	// get the balance of WETH for a given account
	balance, err := avs.WETHTokenContract.BalanceOf(&callOpts, account)
	require.NoError(t, err)

	return balance.Uint64()
}

func checkIfCodePresent(
	t *testing.T,
	ctx context.Context,
	avs AVS,
	deployInfo map[types.ContractName]types.DeployInfo,
	blockNumber uint64,
	contractName types.ContractName) {
	t.Helper()

	codeBytes, err := avs.Client.CodeAt(ctx, deployInfo[contractName].Address, big.NewInt(int64(blockNumber)))
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

func undelegateWETHForDelegattor(
	t *testing.T,
	ctx context.Context,
	avs AVS,
	delegatorAddr common.Address,
	delegatorPk *ecdsa.PrivateKey) {
	t.Helper()

	txOpts, err := bind.NewKeyedTransactorWithChainID(delegatorPk, big.NewInt(int64(avs.Chain.ID)))
	require.NoError(t, err)
	txOpts.Context = ctx
	tx, err := avs.DelegationManagerContract.Undelegate(txOpts, delegatorAddr)
	require.NoError(t, err)
	_, err = bind.WaitMined(ctx, avs.Client, tx)
	require.NoError(t, err)
}

func unregisterOperatorFromAVS(t *testing.T,
	ctx context.Context,
	avs AVS,
	oprAddr common.Address,
	oprPk *ecdsa.PrivateKey) {
	t.Helper()

	txOpts, err := bind.NewKeyedTransactorWithChainID(oprPk, big.NewInt(int64(avs.Chain.ID)))
	require.NoError(t, err)
	txOpts.Context = ctx
	tx, err := avs.AVSContract.DeregisterOperatorFromAVS(txOpts, oprAddr)
	require.NoError(t, err)
	_, err = bind.WaitMined(ctx, avs.Client, tx)
	require.NoError(t, err)
}
