package e2e_test

import (
	"context"
	"crypto/ecdsa"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"math/big"
	"strings"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/test/e2e/types"

	"github.com/stretchr/testify/require"
)

const (
	omniDeployPk = "0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80"
	elDeployPk   = "0x2a871d0798f97d79848a013d4936a73bf4cc922c825d33c1cf7073dff6d409c6" // anvil account 9

	//// info of pre-funded anvil account 3 as delegator1
	//delegator1Addr = "0x90F79bf6EB2c4f870365E785982E1f101E93b906"
	//delegator1Pk   = "0x7c852118294e51e653712a81e05800f419141751be58f605c371e15141b007a6"
	//
	//// info of pre-funded anvil account 4 as delegator2
	//delegator2Addr = "0x15d34AAf54267DB7D7c367839AAf71A00a2C6A65"
	//delegator2Pk   = "0x47e179ec197488593b187f80a00eb0da91f1b9d0b13f8733639f19c30a34926a"
	//
	//// info of pre-funded anvil account 6 as operator1
	//operator1Addr        = "0x976EA74026E726554dB657fA54763abd0C3a0aa9"
	//operator1PK          = "0x92db14e403b83dfe3df233f83dfa3a0d7096f21ca9b0d6d6b8d88b2b4ec1564e"
	operator1MetaDataURI = "https://www.operator1.com"
	//
	//// info of pre-funded anvil account 7 as operator2
	//operator2Addr        = "0x14dC79964da2C08b23698B3D3cc7Ca32193d9955"
	//operator2PK          = "0x4bbbf85ce3377467afe5d46f804f221813b2bb87f24d81f60f1fcdbf7cbf4356"
	operator2MetaDataURI = "https://www.operator2.com"

	MinStateForOperatorInEigenLayer = 10
	MaxNumberOfOperators            = 2
)

var (
	zeroAddr = common.HexToAddress("0x0000000000000000000000000000000000000000")
	//opr1Addr = common.HexToAddress(operator1Addr)
	//opr2Addr = common.HexToAddress(operator2Addr)
	//del1Addr = common.HexToAddress(delegator1Addr)
	//del2Addr = common.HexToAddress(delegator1Addr)

	omniDepPk = mustHexToKey(omniDeployPk)
	elDepPk   = mustHexToKey(elDeployPk)
	//opr1Pk    = mustHexToKey(operator1PK)
	//opr2Pk    = mustHexToKey(operator2PK)
	//del1Addr    = mustHexToKey(delegator1Pk)
	//del2Pk    = mustHexToKey(delegator2Pk)
)

var (
	ErrStrOprAlreadyReg = "operator has already registered"
)

func TestEigenAndOmniAVS(t *testing.T) {
	testAVS(t, func(t *testing.T, avs AVS, deployInfo map[types.ContractName]types.DeployInfo) {
		t.Helper()
		ctx := context.Background()

		// create new operators, delegators and fund them with ETH
		operator1Addr, opr1Addr, opr1Pk := createAccount(t)
		operator2Addr, opr2Addr, opr2Pk := createAccount(t)
		delegator1Addr, del1Addr, del1Pk := createAccount(t)
		delegator2Addr, del2Addr, del2Pk := createAccount(t)
		transferFundTo(t, operator1Addr, avs.Client)
		transferFundTo(t, operator2Addr, avs.Client)
		transferFundTo(t, delegator1Addr, avs.Client)
		transferFundTo(t, delegator2Addr, avs.Client)

		// check if contracts are deployed and configured properly
		checkIfContractsAreDeployed(t, ctx, avs, deployInfo)
		checkOmniAVSConfigurations(t, ctx, avs)

		// register operator with EigenLayer
		checkRegisteringOperatorToEL(t, ctx, avs, deployInfo, opr1Pk, operator1Addr, operator1MetaDataURI, "")
		checkRegisteringOperatorToEL(t, ctx, avs, deployInfo, opr1Pk, operator1Addr, operator1MetaDataURI, ErrStrOprAlreadyReg)
		checkRegisteringOperatorToEL(t, ctx, avs, deployInfo, opr2Pk, operator2Addr, operator2MetaDataURI, "")

		// add funding to operators and delegators
		checkFundingWETHToAccounts(t, ctx, avs, 1000, opr1Addr, opr2Addr, del1Addr, del2Addr)

		// register operators to omni AVS
		whiteListStrategy(t, ctx, avs, deployInfo)
		checkRegisteringOperatorsWithAVS(t, ctx, avs, deployInfo, opr1Addr, opr1Pk, 100, true)
		checkRegisteringOperatorsWithAVS(t, ctx, avs, deployInfo, opr2Addr, opr2Pk, 100, true)

		// delegate stake from delegaters
		delegateWETHToOperator(t, ctx, avs, del1Addr, del1Pk, opr1Addr, 20)
		delegateWETHToOperator(t, ctx, avs, del1Addr, del1Pk, opr2Addr, 30)
		delegateWETHToOperator(t, ctx, avs, del1Addr, del1Pk, opr1Addr, 20)
		delegateWETHToOperator(t, ctx, avs, del2Addr, del2Pk, opr2Addr, 20)

	})
}

func createAccount(t *testing.T) (string, common.Address, *ecdsa.PrivateKey) {
	privateKey, err := crypto.GenerateKey()
	require.NoError(t, err)

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	require.Equal(t, true, ok)
	address := crypto.PubkeyToAddress(*publicKeyECDSA)

	return address.String(), address, privateKey
}

func transferFundTo(t *testing.T, addr string, client *ethclient.Client) {
	pk := mustHexToKey("0x4bbbf85ce3377467afe5d46f804f221813b2bb87f24d81f60f1fcdbf7cbf4356")
	adr := common.HexToAddress("0x14dC79964da2C08b23698B3D3cc7Ca32193d9955")

	nonce, err := client.PendingNonceAt(context.Background(), adr)
	require.NoError(t, err)
	value := big.NewInt(1000000000000000000) // 10 ETH
	gasLimit := uint64(21000)
	tip := big.NewInt(2000000000)
	feeCap := big.NewInt(20000000000)
	require.NoError(t, err)

	var data []byte
	chainID, err := client.NetworkID(context.Background())
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

	err = client.SendTransaction(context.Background(), signedTx)
	require.NoError(t, err)
}

/*
 * internal/ helper functions
 */
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

func checkOmniAVSConfigurations(
	t *testing.T,
	ctx context.Context,
	avs AVS) {
	t.Helper()

	txOpts, err := bind.NewKeyedTransactorWithChainID(omniDepPk, big.NewInt(int64(avs.Chain.ID)))
	require.NoError(t, err)
	txOpts.Context = ctx

	// set min stake
	minStake := big.NewInt(MinStateForOperatorInEigenLayer)
	_, err = avs.AVSContract.SetMinimumOperatorStake(txOpts, minStake)
	require.NoError(t, err)

	// check if min stake is set properly
	callOpts := bind.CallOpts{}
	operatorStake, err := avs.AVSContract.MinimumOperatorStake(&callOpts)
	require.NoError(t, err)
	require.Equal(t, minStake.Uint64(), operatorStake.Uint64())

	// set operator count
	_, err = avs.AVSContract.SetMaxOperatorCount(txOpts, uint32(MaxNumberOfOperators))
	require.NoError(t, err)

	// check if operator count is set properly
	opCount, err := avs.AVSContract.MaxOperatorCount(&callOpts)
	require.NoError(t, err)
	require.Equal(t, uint32(MaxNumberOfOperators), opCount)
}

func checkRegisteringOperatorToEL(
	t *testing.T,
	ctx context.Context,
	avs AVS,
	deployInfo map[types.ContractName]types.DeployInfo,
	oprPk *ecdsa.PrivateKey,
	oprAddrStr string,
	oprMetaDataURI string,
	regErrStr string) {
	t.Helper()

	// register operator to EL
	txOpts, err := bind.NewKeyedTransactorWithChainID(oprPk, big.NewInt(int64(avs.Chain.ID)))
	require.NoError(t, err)
	txOpts.Context = ctx
	operatorDetails := getOperatorDetails(oprAddrStr)
	tx, err := avs.DelegationManagerContract.RegisterAsOperator(txOpts, operatorDetails, oprMetaDataURI)

	if regErrStr != "" {
		require.ErrorContains(t, err, regErrStr, "")
		return
	}

	require.NoError(t, err)
	require.Equal(t, big.NewInt(int64(avs.Chain.ID)), tx.ChainId(), "chain Id not equal")

	//// get block where the operator was registered
	//hash := tx.Hash()
	//block, err := avs.Client.BlockByHash(ctx, hash)
	//require.NoError(t, err)
	//
	//// check logs if the operator is registered properly
	//checkForOperatorRegisteredToELLog(t, ctx, avs,
	//	common.HexToAddress(operator1Addr),
	//	deployInfo[types.ContractELDelegationManager].Address,
	//	block.Header().Number.Uint64())
}

func checkFundingWETHToAccounts(
	t *testing.T,
	ctx context.Context,
	avs AVS,
	amount int64,
	opr1Addr common.Address,
	opr2Addr common.Address,
	del1Addr common.Address,
	del2Addr common.Address) {
	t.Helper()

	// fund the operators and delegators
	fundAccountWithWETH(t, ctx, avs, opr1Addr, amount)
	fundAccountWithWETH(t, ctx, avs, opr2Addr, amount)
	fundAccountWithWETH(t, ctx, avs, del1Addr, amount)
	fundAccountWithWETH(t, ctx, avs, del2Addr, amount)

	require.Equal(t, uint64(amount), wETHBalance(t, ctx, avs, opr1Addr))
	require.Equal(t, uint64(amount), wETHBalance(t, ctx, avs, opr2Addr))
	require.Equal(t, uint64(amount), wETHBalance(t, ctx, avs, del1Addr))
	require.Equal(t, uint64(amount), wETHBalance(t, ctx, avs, del2Addr))
}

func checkRegisteringOperatorsWithAVS(
	t *testing.T,
	ctx context.Context,
	avs AVS,
	deployInfo map[types.ContractName]types.DeployInfo,
	oprAddr common.Address,
	oprPk *ecdsa.PrivateKey,
	selfStakeAmount int64,
	addToAllowList bool) {
	t.Helper()

	// add the operator to omni avs allow list
	if addToAllowList {
		addOperatorToAllowList(t, ctx, avs, oprAddr)
	}

	// delegate min stake from operator to strategy
	depositToStrategy(t, ctx, avs, deployInfo, oprPk, selfStakeAmount)

	// register the operator with omni avs
	registerOperatorToAVS(t, ctx, avs, deployInfo, oprAddr, oprPk)
}

func addOperatorToAllowList(
	t *testing.T,
	ctx context.Context,
	avs AVS,
	oprAddr common.Address) {
	t.Helper()

	txOpts, err := bind.NewKeyedTransactorWithChainID(omniDepPk, big.NewInt(int64(avs.Chain.ID)))
	require.NoError(t, err)
	txOpts.Context = ctx

	_, err = avs.AVSContract.AddToAllowlist(txOpts, oprAddr)
	require.NoError(t, err)
}

func registerOperatorToAVS(t *testing.T,
	ctx context.Context,
	avs AVS,
	deployInfo map[types.ContractName]types.DeployInfo,
	oprAddr common.Address,
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
		Expiry:    big.NewInt(int64(block.Time()) + int64((24 * time.Hour))),
	}
	callOpts := bind.CallOpts{}
	omniAVSAddr := deployInfo[types.ContractOmniAVS].Address
	digestHash, err := avs.AVSDirectory.CalculateOperatorAVSRegistrationDigestHash(&callOpts, oprAddr, omniAVSAddr, operatorSignatureWithSaltAndExpiry.Salt, operatorSignatureWithSaltAndExpiry.Expiry)
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
	require.Equal(t, big.NewInt(int64(avs.Chain.ID)), tx.ChainId(), "chain Id not equal")
}

func getOperatorDetails(addr string) bindings.IDelegationManagerOperatorDetails {
	return bindings.IDelegationManagerOperatorDetails{
		EarningsReceiver:         common.HexToAddress(addr),
		DelegationApprover:       zeroAddr,
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
		require.Equal(t, e.Operator, operatorAddr, "operator is not matching")
		require.Equal(t, e.OperatorDetails.DelegationApprover, zeroAddr, "delegation approver is not matching")
		require.Equal(t, e.OperatorDetails.EarningsReceiver, operatorAddr, "earnings receiver is not matching")
		break // there should be only one log event in this block
	}
}

func fundAccountWithWETH(t *testing.T,
	ctx context.Context,
	avs AVS,
	accountToFund common.Address,
	amount int64) {

	txOpts, err := bind.NewKeyedTransactorWithChainID(elDepPk, big.NewInt(int64(avs.Chain.ID)))
	require.NoError(t, err)
	txOpts.Context = ctx

	_, err = avs.WETHTokenContract.Mint(txOpts, accountToFund, big.NewInt(amount))
	require.NoError(t, err)
}

func whiteListStrategy(
	t *testing.T,
	ctx context.Context,
	avs AVS,
	deployInfo map[types.ContractName]types.DeployInfo) {
	t.Helper()

	wethStratAddr := deployInfo[types.ContractELWETHStrategy].Address
	txOpts, err := bind.NewKeyedTransactorWithChainID(elDepPk, big.NewInt(int64(avs.Chain.ID)))
	require.NoError(t, err)
	txOpts.Context = ctx
	_, err = avs.StrategyManagerContract.AddStrategiesToDepositWhitelist(txOpts, []common.Address{wethStratAddr}, []bool{false})
	require.NoError(t, err)

}

func depositToStrategy(
	t *testing.T,
	ctx context.Context,
	avs AVS,
	deployInfo map[types.ContractName]types.DeployInfo,
	depositorPk *ecdsa.PrivateKey,
	amount int64) {
	t.Helper()

	stratManAddr := deployInfo[types.ContractELStrategyManager].Address
	wethStratAddr := deployInfo[types.ContractELWETHStrategy].Address
	wethTokenAddr := deployInfo[types.ContractELWETH].Address

	// approve strategy manager to deposit on behalf of the depositor
	txOpts, err := bind.NewKeyedTransactorWithChainID(depositorPk, big.NewInt(int64(avs.Chain.ID)))
	require.NoError(t, err)
	txOpts.Context = ctx
	_, err = avs.WETHTokenContract.Approve(txOpts, stratManAddr, big.NewInt(amount))
	require.NoError(t, err)

	// deposit to strategy (assuming that this strategy is already white listed)
	txOpts, err = bind.NewKeyedTransactorWithChainID(depositorPk, big.NewInt(int64(avs.Chain.ID)))
	require.NoError(t, err)
	txOpts.Context = ctx
	_, err = avs.StrategyManagerContract.DepositIntoStrategy(txOpts, wethStratAddr, wethTokenAddr, big.NewInt(amount))
	require.NoError(t, err)
}

func delegateWETHToOperator(t *testing.T,
	ctx context.Context,
	avs AVS,
	delegator common.Address,
	delegatorPk *ecdsa.PrivateKey,
	operator common.Address,
	amount int64) {

	txOpts, err := bind.NewKeyedTransactorWithChainID(delegatorPk, big.NewInt(int64(avs.Chain.ID)))
	require.NoError(t, err)
	txOpts.Context = ctx
	txOpts.Value = big.NewInt(amount)

	blockNumber, err := avs.Client.BlockNumber(ctx)
	require.NoError(t, err)
	block, err := avs.Client.BlockByNumber(ctx, big.NewInt(int64(blockNumber)))
	require.NoError(t, err)
	approverSignatureAndExpiry := bindings.ISignatureUtilsSignatureWithExpiry{
		Signature: []byte{0},
		Expiry:    big.NewInt(int64(block.Time()) + int64((24 * time.Hour))),
	}
	approverSalt := crypto.Keccak256Hash(delegator.Bytes())
	_, err = avs.DelegationManagerContract.DelegateTo(txOpts, operator, approverSignatureAndExpiry, approverSalt)
	require.NoError(t, err)

}

func wETHBalance(t *testing.T,
	ctx context.Context,
	avs AVS,
	account common.Address) uint64 {

	callOpts := bind.CallOpts{
		From:    account,
		Context: ctx,
	}

	// get the balance of WETH for a given account
	balance, err := avs.WETHTokenContract.BalanceOf(&callOpts, account)
	require.NoError(t, err)
	return balance.Uint64()
}

func checkIfCodePresent(t *testing.T,
	ctx context.Context,
	avs AVS,
	deployInfo map[types.ContractName]types.DeployInfo,
	blockNumber uint64,
	contractName types.ContractName) {
	codeBytes, err := avs.Client.CodeAt(ctx, deployInfo[contractName].Address, big.NewInt(int64(blockNumber)))
	require.NoError(t, err)
	require.Less(t, 0, len(codeBytes))
}

func mustHexToKey(privKeyHex string) *ecdsa.PrivateKey {
	privKey, err := crypto.HexToECDSA(strings.TrimPrefix(privKeyHex, "0x"))
	if err != nil {
		panic(err)
	}

	return privKey
}
