package e2e_test

import (
	"context"
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/omni-network/omni/lib/errors"
	"math/big"
	"strings"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/test/e2e/types"

	"github.com/stretchr/testify/require"
)

const (
	omniDeployPk = "0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80"
	elDeployPk   = "0x2a871d0798f97d79848a013d4936a73bf4cc922c825d33c1cf7073dff6d409c6" // anvil account 9

	// info of pre-funded anvil account 3 as delegator1
	delegator1Addr = "0x90F79bf6EB2c4f870365E785982E1f101E93b906"
	delegator1Pk   = "0x7c852118294e51e653712a81e05800f419141751be58f605c371e15141b007a6"

	// info of pre-funded anvil account 4 as delegator2
	delegator2Addr = "0x15d34AAf54267DB7D7c367839AAf71A00a2C6A65"
	delegator2Pk   = "0x47e179ec197488593b187f80a00eb0da91f1b9d0b13f8733639f19c30a34926a"

	// info of pre-funded anvil account 6 as operator1
	operator1Addr        = "0x976EA74026E726554dB657fA54763abd0C3a0aa9"
	operator1PK          = "0x92db14e403b83dfe3df233f83dfa3a0d7096f21ca9b0d6d6b8d88b2b4ec1564e"
	operator1MetaDataURI = "https://www.operator1.com"

	// info of pre-funded anvil account 7 as operator2
	operator2Addr        = "0x14dC79964da2C08b23698B3D3cc7Ca32193d9955"
	operator2PK          = "0x4bbbf85ce3377467afe5d46f804f221813b2bb87f24d81f60f1fcdbf7cbf4356"
	operator2MetaDataURI = "https://www.operator2.com"

	// info of pre-funded anvil account 7 as operator2
	operator3Addr        = "0x14dC79964da2C08b23698B3D3cc7Ca32193d9955"
	operator3PK          = "0x4bbbf85ce3377467afe5d46f804f221813b2bb87f24d81f60f1fcdbf7cbf4356"
	operator3MetaDataURI = "https://www.operator2.com"

	// anvil account 8 used as token store
	tokenStoreAddr = "0x23618e81E3f5cdF7f54C3d65f7FBc0aBf5B21E8f"
	tokenStorePk   = "0xdbda1821b80551c9d65939329250298aa3472ba22feea921c0cf5d620ea67b97"

	MinStateForOperatorInEigenLayer = 10
	MaxNumberOfOperators            = 2
)

type LogEvents int

const (
	EventOperatorRegisteredToEL LogEvents = iota + 1
	EventOperatorRegisteredToAVS
	EventStakeDelegated
	EventStakeUnDelegated
	EventOperatorUnRegisteredFromAVS
	EventOperatorUnRegisteredFromEL
	EventOperatorStatus
)

var (
	zeroAddr  = common.HexToAddress("0x0000000000000000000000000000000000000000")
	omniDepPk = mustHexToKey(omniDeployPk)
	elDepPk   = mustHexToKey(elDeployPk)
	del1Pk    = mustHexToKey(delegator1Pk)
	del2Pk    = mustHexToKey(delegator2Pk)
	opr1Pk    = mustHexToKey(operator1PK)
	opr2Pk    = mustHexToKey(operator2PK)
	opr3Pk    = mustHexToKey(operator3PK)
	tokStrPk  = mustHexToKey(tokenStorePk)
)

func TestEigen_AreContractsDeployed(t *testing.T) {
	testAVS(t, func(t *testing.T, avs AVS, deployInfo map[types.ContractName]types.DeployInfo) {
		t.Helper()
		ctx := context.Background()

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
	})
}

func TestEigen_RegisterOperator(t *testing.T) {
	t.Run("register operator 1", func(t *testing.T) {
		testAVS(t, func(t *testing.T, avs AVS, deployInfo map[types.ContractName]types.DeployInfo) {
			t.Helper()
			ctx := context.Background()

			txOpts, err := bind.NewKeyedTransactorWithChainID(opr1Pk, big.NewInt(int64(avs.Chain.ID)))
			require.NoError(t, err)
			txOpts.Context = ctx

			// register operator1 to EL
			operatorDetails := getOperatorDetails(operator1Addr)
			tx, err := avs.DelegationManagerContract.RegisterAsOperator(txOpts, operatorDetails, operator1MetaDataURI)
			require.NoError(t, err)
			require.Equal(t, big.NewInt(int64(avs.Chain.ID)), tx.ChainId(), "chain Id not equal")

			// get block where the operator was registered
			hash := tx.Hash()
			block, err := avs.Client.BlockByHash(ctx, hash)
			require.NoError(t, err)

			// check logs if the operator is registered
			checkForOperatorRegisteredToELLog(t, ctx, avs,
				common.HexToAddress(operator1Addr),
				deployInfo[types.ContractELDelegationManager].Address,
				block.Header().Number.Uint64())
		})
	})

	t.Run("register operator 1 again as duplicate", func(t *testing.T) {
		testAVS(t, func(t *testing.T, avs AVS, deployInfo map[types.ContractName]types.DeployInfo) {
			t.Helper()
			ctx := context.Background()

			txOpts, err := bind.NewKeyedTransactorWithChainID(opr1Pk, big.NewInt(int64(avs.Chain.ID)))
			require.NoError(t, err)
			txOpts.Context = ctx

			// register operator1 to EL
			operatorDetails := getOperatorDetails(operator1Addr)
			_, err = avs.DelegationManagerContract.RegisterAsOperator(txOpts, operatorDetails, operator1MetaDataURI)
			require.ErrorContains(t, err, "operator has already registered", "")
		})
	})

	t.Run("register operator 2", func(t *testing.T) {
		testAVS(t, func(t *testing.T, avs AVS, deployInfo map[types.ContractName]types.DeployInfo) {
			t.Helper()
			ctx := context.Background()

			txOpts, err := bind.NewKeyedTransactorWithChainID(opr2Pk, big.NewInt(int64(avs.Chain.ID)))
			require.NoError(t, err)
			txOpts.Context = ctx

			// register operator2 to EL
			operatorDetails := getOperatorDetails(operator2Addr)
			tx, err := avs.DelegationManagerContract.RegisterAsOperator(txOpts, operatorDetails, operator2MetaDataURI)
			require.NoError(t, err)
			require.Equal(t, big.NewInt(int64(avs.Chain.ID)), tx.ChainId(), "chain Id not equal")

			// get block where the operator was registered
			hash := tx.Hash()
			block, err := avs.Client.BlockByHash(ctx, hash)
			require.NoError(t, err)

			// check logs if the operator is registered
			checkForOperatorRegisteredToELLog(t, ctx, avs,
				common.HexToAddress(operator2Addr),
				deployInfo[types.ContractELDelegationManager].Address,
				block.Header().Number.Uint64())
		})
	})
}

func TestEigen_FundOperatorsAndDelegators(t *testing.T) {
	t.Run("fund token to delegators and operators", func(t *testing.T) {
		testAVS(t, func(t *testing.T, avs AVS, deployInfo map[types.ContractName]types.DeployInfo) {
			t.Helper()
			ctx := context.Background()

			// fund the operators and delegators
			opr1Addr := common.HexToAddress(operator1Addr)
			opr2Addr := common.HexToAddress(operator2Addr)
			del1Addr := common.HexToAddress(delegator1Addr)
			del2Addr := common.HexToAddress(delegator1Addr)
			fundAccountWithWETH(t, ctx, avs, opr1Addr, 1000) // less than min stake
			fundAccountWithWETH(t, ctx, avs, opr2Addr, 1000) // more than min stake
			fundAccountWithWETH(t, ctx, avs, del1Addr, 1000)
			fundAccountWithWETH(t, ctx, avs, del2Addr, 1000)

			require.Equal(t, uint64(1000), wETHBalance(t, ctx, avs, opr1Addr))
			require.Equal(t, uint64(1000), wETHBalance(t, ctx, avs, opr2Addr))
			require.Equal(t, uint64(1000), wETHBalance(t, ctx, avs, del1Addr))
			require.Equal(t, uint64(1000), wETHBalance(t, ctx, avs, del2Addr))
		})
	})

	//t.Run("stake token to delegators and operators", func(t *testing.T) {
	//	testAVS(t, func(t *testing.T, avs AVS, deployInfo map[types.ContractName]types.DeployInfo) {
	//		t.Helper()
	//		ctx := context.Background()
	//
	//	})
	//})
}

func TestEigen_RegisterOperatorForAVS(t *testing.T) {
	t.Run("register operator 1 for AVS with less than min stake", func(t *testing.T) {
		t.Parallel()
		testAVS(t, func(t *testing.T, avs AVS, deployInfo map[types.ContractName]types.DeployInfo) {
			t.Helper()
			ctx := context.Background()
			opr1Addr := common.HexToAddress(operator1Addr)

			// add operator to avs allow list
			txOpts, err := bind.NewKeyedTransactorWithChainID(omniDepPk, big.NewInt(int64(avs.Chain.ID)))
			require.NoError(t, err)
			txOpts.Context = ctx
			//_, err = avs.AVSContract.AddToAllowlist(txOpts, opr1Addr)
			//require.NoError(t, err)

			// delegate min stake from operator to strategy
			depositToStrategy(t, ctx, avs, common.HexToAddress(operator1Addr), opr1Pk, 11, deployInfo)

			// register operator to avs
			blockNumber, err := avs.Client.BlockNumber(ctx)
			require.NoError(t, err)
			block, err := avs.Client.BlockByNumber(ctx, big.NewInt(int64(blockNumber)))
			require.NoError(t, err)
			operatorSignatureWithSaltAndExpiry := bindings.ISignatureUtilsSignatureWithSaltAndExpiry{
				Signature: []byte{0},
				Salt:      crypto.Keccak256Hash(opr1Addr.Bytes()),
				Expiry:    big.NewInt(int64(block.Time()) + int64((24 * time.Hour))),
			}

			callOpts := bind.CallOpts{}
			omniAVSADDR := deployInfo[types.ContractOmniAVS].Address
			digestHash, err := avs.AVSDirectory.CalculateOperatorAVSRegistrationDigestHash(&callOpts, opr1Addr, omniAVSADDR, operatorSignatureWithSaltAndExpiry.Salt, operatorSignatureWithSaltAndExpiry.Expiry)
			require.NoError(t, err)

			operatorSignatureWithSaltAndExpiry.Signature, err = crypto.Sign(digestHash[:32], opr1Pk)
			require.NoError(t, err)
			if len(operatorSignatureWithSaltAndExpiry.Signature) != 65 {
				require.NoError(t, errors.New("invalid signature length"))
			}
			operatorSignatureWithSaltAndExpiry.Signature[64] += 27
			txOpts, err = bind.NewKeyedTransactorWithChainID(opr1Pk, big.NewInt(int64(avs.Chain.ID)))
			require.NoError(t, err)
			txOpts.Context = ctx
			tx, err := avs.AVSContract.RegisterOperatorToAVS(txOpts, opr1Addr, operatorSignatureWithSaltAndExpiry)
			require.NoError(t, err)
			require.Equal(t, big.NewInt(int64(avs.Chain.ID)), tx.ChainId(), "chain Id not equal")

		})
	})

	t.Run("register operator 2 for AVS with no stake", func(t *testing.T) {
		t.Parallel()
		testAVS(t, func(t *testing.T, avs AVS, deployInfo map[types.ContractName]types.DeployInfo) {
			t.Helper()
			ctx := context.Background()

			txOpts, err := bind.NewKeyedTransactorWithChainID(opr2Pk, big.NewInt(int64(avs.Chain.ID)))
			require.NoError(t, err)
			txOpts.Context = ctx

			blockNumber, err := avs.Client.BlockNumber(ctx)
			require.NoError(t, err)
			block, err := avs.Client.BlockByNumber(ctx, big.NewInt(int64(blockNumber)))
			require.NoError(t, err)

			opr2Addr := common.HexToAddress(operator2Addr)
			callOpts := bind.CallOpts{
				From:    opr2Addr,
				Context: ctx,
			}
			_, err = avs.DelegationManagerContract.DelegatedTo(&callOpts, opr2Addr)
			require.NoError(t, err)

			operatorSignatureWithSaltAndExpiry := bindings.ISignatureUtilsSignatureWithSaltAndExpiry{
				Signature: []byte{0},
				Salt:      crypto.Keccak256Hash(opr2Addr.Bytes()),
				Expiry:    big.NewInt(int64(block.Time()) + int64((24 * time.Hour))),
			}

			_, err = avs.AVSContract.RegisterOperatorToAVS(txOpts, common.HexToAddress(operator2Addr), operatorSignatureWithSaltAndExpiry)
			require.ErrorContains(t, err, "minimum stake not met", "")

		})
	})
}

/*
* internal/ helper functions
 */
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

func depositToStrategy(t *testing.T,
	ctx context.Context,
	avs AVS,
	depositorAddr common.Address,
	depositorPk *ecdsa.PrivateKey,
	amount int64,
	deployInfo map[types.ContractName]types.DeployInfo) {

	stratManAddr := deployInfo[types.ContractELStrategyManager].Address
	wethStratAddr := deployInfo[types.ContractELWETHStrategy].Address
	//wethTokenAddr := deployInfo[types.ContractELWETH].Address

	// approve strategy manager to deposit on behalf of the depositor
	txOpts, err := bind.NewKeyedTransactorWithChainID(depositorPk, big.NewInt(int64(avs.Chain.ID)))
	require.NoError(t, err)
	txOpts.Context = ctx
	_, err = avs.WETHTokenContract.Approve(txOpts, stratManAddr, big.NewInt(amount))
	require.NoError(t, err)

	// whitelist strategy
	txOpts, err = bind.NewKeyedTransactorWithChainID(elDepPk, big.NewInt(int64(avs.Chain.ID)))
	require.NoError(t, err)
	txOpts.Context = ctx
	_, err = avs.StrategyManagerContract.AddStrategiesToDepositWhitelist(txOpts, []common.Address{wethStratAddr}, []bool{false})
	require.NoError(t, err)

	// deposit to strategy
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
	amount int64,
	expiryTime *big.Int) {

	txOpts, err := bind.NewKeyedTransactorWithChainID(delegatorPk, big.NewInt(int64(avs.Chain.ID)))
	require.NoError(t, err)
	txOpts.Context = ctx
	txOpts.Value = big.NewInt(amount)

	approverSignatureAndExpiry := bindings.ISignatureUtilsSignatureWithExpiry{
		Signature: []byte{0},
		Expiry:    expiryTime,
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
