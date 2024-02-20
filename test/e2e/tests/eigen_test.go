package e2e_test

import (
	"context"
	"github.com/ethereum/go-ethereum/crypto"
	"math/big"
	"testing"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/test/e2e/netman"
	"github.com/omni-network/omni/test/e2e/types"

	"github.com/stretchr/testify/require"
)

const (
	deployPk = "0xac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80"

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
	zeroAddr = common.HexToAddress("0x0000000000000000000000000000000000000000")
	depPk    = netman.MustHexToKey(deployPk)
	del1Pk   = netman.MustHexToKey(delegator1Pk)
	del2Pk   = netman.MustHexToKey(delegator2Pk)
	opr1Pk   = netman.MustHexToKey(operator1PK)
	opr2Pk   = netman.MustHexToKey(operator2PK)
	opr3Pk   = netman.MustHexToKey(operator3PK)
)

func TestEigen_AreContractsDeployed(t *testing.T) {
	testAVS(t, func(t *testing.T, avs AVS, deployInfo map[types.ContractName]types.DeployInfo) {
		t.Helper()
		ctx := context.Background()

		blockNumber, err := avs.Client.BlockNumber(ctx)
		require.NoError(t, err)

		// check iof the avs contract has code in its respective contract address
		codeBytes, err := avs.Client.CodeAt(ctx, deployInfo[types.ContractOmniAVS].Address, big.NewInt(int64(blockNumber)))
		require.NoError(t, err)
		require.Less(t, 0, len(codeBytes))

		// check if the delegation manager contract has code in its respective contract address
		codeBytes, err = avs.Client.CodeAt(ctx, deployInfo[types.ContractELDelegationManager].Address, big.NewInt(int64(blockNumber)))
		require.NoError(t, err)
		require.Less(t, 0, len(codeBytes))

		txOpts, err := bind.NewKeyedTransactorWithChainID(depPk, big.NewInt(int64(avs.Chain.ID)))
		require.NoError(t, err)
		txOpts.Context = ctx

		// set min stake
		stake := big.NewInt(10)
		_, err = avs.AVSContract.SetMinimumOperatorStake(txOpts, stake)
		require.NoError(t, err)

		// check if min stake is set
		callOpts := bind.CallOpts{}
		operatorStake, err := avs.AVSContract.MinimumOperatorStake(&callOpts)
		require.NoError(t, err)
		require.Equal(t, stake.Uint64(), operatorStake.Uint64())

		// set operator count
		_, err = avs.AVSContract.SetMaxOperatorCount(txOpts, uint32(2))
		require.NoError(t, err)

		// check if operator count is set
		opCount, err := avs.AVSContract.MaxOperatorCount(&callOpts)
		require.NoError(t, err)
		require.Equal(t, uint32(2), opCount)
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

			// check logs
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

			// check logs
			checkForOperatorRegisteredToELLog(t, ctx, avs,
				common.HexToAddress(operator2Addr),
				deployInfo[types.ContractELDelegationManager].Address,
				block.Header().Number.Uint64())
		})
	})

}

func TestEigen_RegisterOperatorForAVS(t *testing.T) {
	t.Run("register operator 1 for AVS", func(t *testing.T) {
		t.Parallel()
		testAVS(t, func(t *testing.T, avs AVS, deployInfo map[types.ContractName]types.DeployInfo) {
			t.Helper()
			ctx := context.Background()

			txOpts, err := bind.NewKeyedTransactorWithChainID(opr1Pk, big.NewInt(int64(avs.Chain.ID)))
			require.NoError(t, err)
			txOpts.Context = ctx

			blockNumber, err := avs.Client.BlockNumber(ctx)
			require.NoError(t, err)
			block, err := avs.Client.BlockByNumber(ctx, big.NewInt(int64(blockNumber)))
			require.NoError(t, err)

			del1Addr := common.HexToAddress(delegator1Addr)
			opr1Addr := common.HexToAddress(operator1Addr)
			callOpts := bind.CallOpts{
				From:    del1Addr,
				Context: ctx,
			}
			_, err = avs.DelegationManagerContract.DelegatedTo(&callOpts, opr1Addr)
			require.NoError(t, err)

			operatorSignatureWithSaltAndExpiry := bindings.ISignatureUtilsSignatureWithSaltAndExpiry{
				Signature: []byte{0},
				Salt:      crypto.Keccak256Hash(opr1Addr.Bytes()),
				Expiry:    big.NewInt(int64(block.Time()) + int64((24 * time.Hour))),
			}

			tx, err := avs.AVSContract.RegisterOperatorToAVS(txOpts, common.HexToAddress(operator1Addr), operatorSignatureWithSaltAndExpiry)
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

func registerOperatorToEL(
	t *testing.T,
	ctx context.Context,
	avs AVS,
	txOpts *bind.TransactOpts,
	operatorAddr string,
	metadataURI string) *ethtypes.Block {
	operatorDetails := getOperatorDetails(operatorAddr)
	tx, err := avs.DelegationManagerContract.RegisterAsOperator(txOpts, operatorDetails, metadataURI)
	require.NoError(t, err)
	require.Equal(t, big.NewInt(int64(avs.Chain.ID)), tx.ChainId(), "chain Id not equal")

	// get block where the operator was registered
	hash := tx.Hash()
	block, err := avs.Client.BlockByHash(ctx, hash)
	require.NoError(t, err)
	return block
}

func registerOperatorToAVSL(
	t *testing.T,
	avs AVS,
	txOpts *bind.TransactOpts,
	operatorAddr common.Address,
	operatorSig bindings.ISignatureUtilsSignatureWithSaltAndExpiry) {
	tx, err := avs.AVSContract.RegisterOperatorToAVS(txOpts, operatorAddr, operatorSig)
	require.NoError(t, err)
	require.Equal(t, big.NewInt(int64(avs.Chain.ID)), tx.ChainId(), "chain Id not equal")
}

func checkOperatorDetails(t *testing.T, operatorAddr common.Address,
	operatorDetails *bindings.IDelegationManagerOperatorDetails) {
	require.Equal(t, operatorDetails.DelegationApprover, zeroAddr, "delegation approver is not matching")
	require.Equal(t, operatorDetails.EarningsReceiver, operatorAddr, "earnings receiver is not matching")
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
