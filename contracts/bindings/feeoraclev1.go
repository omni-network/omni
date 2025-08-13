// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// IFeeOracleV1ChainFeeParams is an auto generated low-level Go binding around an user-defined struct.
type IFeeOracleV1ChainFeeParams struct {
	ChainId      uint64
	PostsTo      uint64
	GasPrice     *big.Int
	ToNativeRate *big.Int
}

// FeeOracleV1MetaData contains all meta data concerning the FeeOracleV1 contract.
var FeeOracleV1MetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"CONVERSION_RATE_DENOM\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"baseGasLimit\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"bulkSetFeeParams\",\"inputs\":[{\"name\":\"params\",\"type\":\"tuple[]\",\"internalType\":\"structIFeeOracleV1.ChainFeeParams[]\",\"components\":[{\"name\":\"chainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"postsTo\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"gasPrice\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"toNativeRate\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"feeFor\",\"inputs\":[{\"name\":\"destChainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"gasLimit\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"feeParams\",\"inputs\":[{\"name\":\"chainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIFeeOracleV1.ChainFeeParams\",\"components\":[{\"name\":\"chainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"postsTo\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"gasPrice\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"toNativeRate\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"gasPriceOn\",\"inputs\":[{\"name\":\"chainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"owner_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"manager_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"baseGasLimit_\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"protocolFee_\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"params\",\"type\":\"tuple[]\",\"internalType\":\"structIFeeOracleV1.ChainFeeParams[]\",\"components\":[{\"name\":\"chainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"postsTo\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"gasPrice\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"toNativeRate\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"manager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"postsTo\",\"inputs\":[{\"name\":\"chainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"protocolFee\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setBaseGasLimit\",\"inputs\":[{\"name\":\"gasLimit\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setGasPrice\",\"inputs\":[{\"name\":\"chainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"gasPrice\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setManager\",\"inputs\":[{\"name\":\"manager_\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setProtocolFee\",\"inputs\":[{\"name\":\"fee\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setToNativeRate\",\"inputs\":[{\"name\":\"chainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"rate\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"toNativeRate\",\"inputs\":[{\"name\":\"chainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"version\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"pure\"},{\"type\":\"event\",\"name\":\"BaseGasLimitSet\",\"inputs\":[{\"name\":\"baseGasLimit\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FeeParamsSet\",\"inputs\":[{\"name\":\"chainId\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"postsTo\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"gasPrice\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"toNativeRate\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"GasPriceSet\",\"inputs\":[{\"name\":\"chainId\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"gasPrice\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ManagerSet\",\"inputs\":[{\"name\":\"manager\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ProtocolFeeSet\",\"inputs\":[{\"name\":\"protocolFee\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ToNativeRateSet\",\"inputs\":[{\"name\":\"chainId\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"toNativeRate\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnableInvalidOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnableUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]}]",
	Bin: "0x608060405234801561000f575f80fd5b5061001861001d565b6100cf565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000900460ff161561006d5760405163f92ee8a960e01b815260040160405180910390fd5b80546001600160401b03908116146100cc5780546001600160401b0319166001600160401b0390811782556040519081527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b50565b6111b5806100dc5f395ff3fe608060405234801561000f575f80fd5b5060043610610127575f3560e01c80638f9d6ace116100a9578063b48ec8611161006e578063b48ec86114610348578063d07041571461037a578063d0ebdbe7146103a5578063ee590a53146103b8578063f2fde38b146103cb575f80fd5b80638f9d6ace146102fc57806393a871881461030657806398563b0314610319578063a34e7abb1461032c578063b0e21e8a1461033f575f80fd5b8063787dce3d116100ef578063787dce3d1461025b5780638b7bfd701461026e5780638da5cb5b146102a75780638dd9523c146102d75780638df66e34146102ea575f80fd5b80632d4634a41461012b578063361c019f146101f8578063481c6a751461020d57806354fd4d5014610238578063715018a614610253575b5f80fd5b6101ab610139366004610de6565b60408051608080820183525f808352602080840182905283850182905260609384018290526001600160401b03958616825260038152908490208451928301855280548087168452600160401b9004909516908201526001840154928101929092526002909201549181019190915290565b6040516101ef91905f6080820190506001600160401b0380845116835280602085015116602084015250604083015160408301526060830151606083015292915050565b60405180910390f35b61020b610206366004610e63565b6103de565b005b600254610220906001600160a01b031681565b6040516001600160a01b0390911681526020016101ef565b60015b6040516001600160401b0390911681526020016101ef565b61020b6104ff565b61020b610269366004610edc565b610512565b61029961027c366004610de6565b6001600160401b03165f9081526003602052604090206002015490565b6040519081526020016101ef565b7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c199300546001600160a01b0316610220565b6102996102e5366004610ef3565b610526565b5f5461023b906001600160401b031681565b610299620f424081565b61020b610314366004610f7f565b6106a9565b61020b610327366004610fbd565b6106e1565b61020b61033a366004610fbd565b610715565b61029960015481565b61023b610356366004610de6565b6001600160401b039081165f90815260036020526040902054600160401b90041690565b610299610388366004610de6565b6001600160401b03165f9081526003602052604090206001015490565b61020b6103b3366004610fe5565b610749565b61020b6103c6366004610de6565b6107b0565b61020b6103d9366004610fe5565b6107c1565b5f6103e76107fb565b805490915060ff600160401b82041615906001600160401b03165f8115801561040d5750825b90505f826001600160401b031660011480156104285750303b155b905081158015610436575080155b156104545760405163f92ee8a960e01b815260040160405180910390fd5b845467ffffffffffffffff19166001178555831561047e57845460ff60401b1916600160401b1785555b6104878b610825565b6104908a610836565b6104998961088b565b6104a2886108d9565b6104ac878761090e565b83156104f257845460ff60401b19168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b5050505050505050505050565b610507610b22565b6105105f610b7d565b565b61051a610b22565b610523816108d9565b50565b6001600160401b038085165f908152600360205260408082208054600160401b900490931682528120600283015460018401549293928491620f42409161056d9190611012565b6105779190611029565b90505f620f4240836002015484600101546105929190611012565b61059c9190611029565b90505f82116105f25760405162461bcd60e51b815260206004820152601a60248201527f4665654f7261636c6556313a206e6f2066656520706172616d7300000000000060448201526064015b60405180910390fd5b5f81116106415760405162461bcd60e51b815260206004820152601a60248201527f4665654f7261636c6556313a206e6f2066656520706172616d7300000000000060448201526064016105e9565b5f61064d886010611012565b90506106598282611012565b5f548490610671908a906001600160401b0316611048565b6001600160401b03166106849190611012565b600154610691919061106f565b61069b919061106f565b9a9950505050505050505050565b6002546001600160a01b031633146106d35760405162461bcd60e51b81526004016105e990611082565b6106dd828261090e565b5050565b6002546001600160a01b0316331461070b5760405162461bcd60e51b81526004016105e990611082565b6106dd8282610bed565b6002546001600160a01b0316331461073f5760405162461bcd60e51b81526004016105e990611082565b6106dd8282610cbf565b610751610b22565b6001600160a01b0381166107a75760405162461bcd60e51b815260206004820152601c60248201527f4665654f7261636c6556313a206e6f207a65726f206d616e616765720000000060448201526064016105e9565b61052381610836565b6107b8610b22565b6105238161088b565b6107c9610b22565b6001600160a01b0381166107f257604051631e4fbdf760e01b81525f60048201526024016105e9565b61052381610b7d565b5f807ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a005b92915050565b61082d610d85565b61052381610daa565b600280546001600160a01b0319166001600160a01b0383169081179091556040519081527f60a0f5b9f9e81e98216071b85826681c796256fe3d1354ecb675580fba64fa69906020015b60405180910390a150565b5f805467ffffffffffffffff19166001600160401b0383169081179091556040519081527f6185fbe062d94552cf644f5cb643f583db7b2e7e66fdc4b4c75ff8876a257ba690602001610880565b60018190556040518181527fdb5aafdb29539329e37d4e3ee869bc4031941fd55a5dfc92824fbe34b204e30d90602001610880565b5f5b81811015610b1d575f83838381811061092b5761092b6110b9565b90506080020180360381019061094191906110cd565b90505f8160400151116109965760405162461bcd60e51b815260206004820152601e60248201527f4665654f7261636c6556313a206e6f207a65726f20676173207072696365000060448201526064016105e9565b5f8160600151116109e55760405162461bcd60e51b81526020600482015260196024820152784665654f7261636c6556313a206e6f207a65726f207261746560381b60448201526064016105e9565b80516001600160401b03165f03610a0e5760405162461bcd60e51b81526004016105e990611148565b80602001516001600160401b03165f03610a6a5760405162461bcd60e51b815260206004820152601c60248201527f4665654f7261636c6556313a206e6f207a65726f20706f737473546f0000000060448201526064016105e9565b80516001600160401b039081165f9081526003602090815260409182902084518154838701519186166fffffffffffffffffffffffffffffffff199091168117600160401b92909616918202959095178255838601516001830181905560608088015160029094018490558551968752938601919091528484015290830152517ff378a0dd98429494eb2e26894562949c4d6e7cef5eb893b1d4c0052078d92fe59181900360800190a150600101610910565b505050565b33610b547f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c199300546001600160a01b031690565b6001600160a01b0316146105105760405163118cdaa760e01b81523360048201526024016105e9565b7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c19930080546001600160a01b031981166001600160a01b03848116918217845560405192169182907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0905f90a3505050565b5f8111610c3c5760405162461bcd60e51b815260206004820152601e60248201527f4665654f7261636c6556313a206e6f207a65726f20676173207072696365000060448201526064016105e9565b816001600160401b03165f03610c645760405162461bcd60e51b81526004016105e990611148565b6001600160401b0382165f81815260036020908152604091829020600101849055815192835282018390527f3b196e45eaa29099834d3d912ac550e4f3e13fef2e2a998100368e506a44d8ff91015b60405180910390a15050565b5f8111610d0a5760405162461bcd60e51b81526020600482015260196024820152784665654f7261636c6556313a206e6f207a65726f207261746560381b60448201526064016105e9565b816001600160401b03165f03610d325760405162461bcd60e51b81526004016105e990611148565b6001600160401b0382165f81815260036020908152604091829020600201849055815192835282018390527f4b4594c9f06af25bc504eead96f7f0eaa3f1577f8d9b075b236520ec712e13089101610cb3565b610d8d610db2565b61051057604051631afcd79f60e31b815260040160405180910390fd5b6107c9610d85565b5f610dbb6107fb565b54600160401b900460ff16919050565b80356001600160401b0381168114610de1575f80fd5b919050565b5f60208284031215610df6575f80fd5b610dff82610dcb565b9392505050565b80356001600160a01b0381168114610de1575f80fd5b5f8083601f840112610e2c575f80fd5b5081356001600160401b03811115610e42575f80fd5b6020830191508360208260071b8501011115610e5c575f80fd5b9250929050565b5f805f805f8060a08789031215610e78575f80fd5b610e8187610e06565b9550610e8f60208801610e06565b9450610e9d60408801610dcb565b93506060870135925060808701356001600160401b03811115610ebe575f80fd5b610eca89828a01610e1c565b979a9699509497509295939492505050565b5f60208284031215610eec575f80fd5b5035919050565b5f805f8060608587031215610f06575f80fd5b610f0f85610dcb565b935060208501356001600160401b0380821115610f2a575f80fd5b818701915087601f830112610f3d575f80fd5b813581811115610f4b575f80fd5b886020828501011115610f5c575f80fd5b602083019550809450505050610f7460408601610dcb565b905092959194509250565b5f8060208385031215610f90575f80fd5b82356001600160401b03811115610fa5575f80fd5b610fb185828601610e1c565b90969095509350505050565b5f8060408385031215610fce575f80fd5b610fd783610dcb565b946020939093013593505050565b5f60208284031215610ff5575f80fd5b610dff82610e06565b634e487b7160e01b5f52601160045260245ffd5b808202811582820484141761081f5761081f610ffe565b5f8261104357634e487b7160e01b5f52601260045260245ffd5b500490565b6001600160401b0381811683821601908082111561106857611068610ffe565b5092915050565b8082018082111561081f5761081f610ffe565b60208082526018908201527f4665654f7261636c6556313a206e6f74206d616e616765720000000000000000604082015260600190565b634e487b7160e01b5f52603260045260245ffd5b5f608082840312156110dd575f80fd5b604051608081018181106001600160401b038211171561110b57634e487b7160e01b5f52604160045260245ffd5b60405261111783610dcb565b815261112560208401610dcb565b602082015260408301356040820152606083013560608201528091505092915050565b6020808252601d908201527f4665654f7261636c6556313a206e6f207a65726f20636861696e20696400000060408201526060019056fea26469706673582212202fa539e859bac6181d4934021292bcce9206509a0e5510811c8f61bd84f0098564736f6c63430008180033",
}

// FeeOracleV1ABI is the input ABI used to generate the binding from.
// Deprecated: Use FeeOracleV1MetaData.ABI instead.
var FeeOracleV1ABI = FeeOracleV1MetaData.ABI

// FeeOracleV1Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use FeeOracleV1MetaData.Bin instead.
var FeeOracleV1Bin = FeeOracleV1MetaData.Bin

// DeployFeeOracleV1 deploys a new Ethereum contract, binding an instance of FeeOracleV1 to it.
func DeployFeeOracleV1(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *FeeOracleV1, error) {
	parsed, err := FeeOracleV1MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(FeeOracleV1Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &FeeOracleV1{FeeOracleV1Caller: FeeOracleV1Caller{contract: contract}, FeeOracleV1Transactor: FeeOracleV1Transactor{contract: contract}, FeeOracleV1Filterer: FeeOracleV1Filterer{contract: contract}}, nil
}

// FeeOracleV1 is an auto generated Go binding around an Ethereum contract.
type FeeOracleV1 struct {
	FeeOracleV1Caller     // Read-only binding to the contract
	FeeOracleV1Transactor // Write-only binding to the contract
	FeeOracleV1Filterer   // Log filterer for contract events
}

// FeeOracleV1Caller is an auto generated read-only Go binding around an Ethereum contract.
type FeeOracleV1Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FeeOracleV1Transactor is an auto generated write-only Go binding around an Ethereum contract.
type FeeOracleV1Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FeeOracleV1Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FeeOracleV1Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FeeOracleV1Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FeeOracleV1Session struct {
	Contract     *FeeOracleV1      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FeeOracleV1CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FeeOracleV1CallerSession struct {
	Contract *FeeOracleV1Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// FeeOracleV1TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FeeOracleV1TransactorSession struct {
	Contract     *FeeOracleV1Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// FeeOracleV1Raw is an auto generated low-level Go binding around an Ethereum contract.
type FeeOracleV1Raw struct {
	Contract *FeeOracleV1 // Generic contract binding to access the raw methods on
}

// FeeOracleV1CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FeeOracleV1CallerRaw struct {
	Contract *FeeOracleV1Caller // Generic read-only contract binding to access the raw methods on
}

// FeeOracleV1TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FeeOracleV1TransactorRaw struct {
	Contract *FeeOracleV1Transactor // Generic write-only contract binding to access the raw methods on
}

// NewFeeOracleV1 creates a new instance of FeeOracleV1, bound to a specific deployed contract.
func NewFeeOracleV1(address common.Address, backend bind.ContractBackend) (*FeeOracleV1, error) {
	contract, err := bindFeeOracleV1(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &FeeOracleV1{FeeOracleV1Caller: FeeOracleV1Caller{contract: contract}, FeeOracleV1Transactor: FeeOracleV1Transactor{contract: contract}, FeeOracleV1Filterer: FeeOracleV1Filterer{contract: contract}}, nil
}

// NewFeeOracleV1Caller creates a new read-only instance of FeeOracleV1, bound to a specific deployed contract.
func NewFeeOracleV1Caller(address common.Address, caller bind.ContractCaller) (*FeeOracleV1Caller, error) {
	contract, err := bindFeeOracleV1(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FeeOracleV1Caller{contract: contract}, nil
}

// NewFeeOracleV1Transactor creates a new write-only instance of FeeOracleV1, bound to a specific deployed contract.
func NewFeeOracleV1Transactor(address common.Address, transactor bind.ContractTransactor) (*FeeOracleV1Transactor, error) {
	contract, err := bindFeeOracleV1(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FeeOracleV1Transactor{contract: contract}, nil
}

// NewFeeOracleV1Filterer creates a new log filterer instance of FeeOracleV1, bound to a specific deployed contract.
func NewFeeOracleV1Filterer(address common.Address, filterer bind.ContractFilterer) (*FeeOracleV1Filterer, error) {
	contract, err := bindFeeOracleV1(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FeeOracleV1Filterer{contract: contract}, nil
}

// bindFeeOracleV1 binds a generic wrapper to an already deployed contract.
func bindFeeOracleV1(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := FeeOracleV1MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FeeOracleV1 *FeeOracleV1Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FeeOracleV1.Contract.FeeOracleV1Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FeeOracleV1 *FeeOracleV1Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FeeOracleV1.Contract.FeeOracleV1Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FeeOracleV1 *FeeOracleV1Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FeeOracleV1.Contract.FeeOracleV1Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FeeOracleV1 *FeeOracleV1CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FeeOracleV1.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FeeOracleV1 *FeeOracleV1TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FeeOracleV1.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FeeOracleV1 *FeeOracleV1TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FeeOracleV1.Contract.contract.Transact(opts, method, params...)
}

// CONVERSIONRATEDENOM is a free data retrieval call binding the contract method 0x8f9d6ace.
//
// Solidity: function CONVERSION_RATE_DENOM() view returns(uint256)
func (_FeeOracleV1 *FeeOracleV1Caller) CONVERSIONRATEDENOM(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FeeOracleV1.contract.Call(opts, &out, "CONVERSION_RATE_DENOM")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CONVERSIONRATEDENOM is a free data retrieval call binding the contract method 0x8f9d6ace.
//
// Solidity: function CONVERSION_RATE_DENOM() view returns(uint256)
func (_FeeOracleV1 *FeeOracleV1Session) CONVERSIONRATEDENOM() (*big.Int, error) {
	return _FeeOracleV1.Contract.CONVERSIONRATEDENOM(&_FeeOracleV1.CallOpts)
}

// CONVERSIONRATEDENOM is a free data retrieval call binding the contract method 0x8f9d6ace.
//
// Solidity: function CONVERSION_RATE_DENOM() view returns(uint256)
func (_FeeOracleV1 *FeeOracleV1CallerSession) CONVERSIONRATEDENOM() (*big.Int, error) {
	return _FeeOracleV1.Contract.CONVERSIONRATEDENOM(&_FeeOracleV1.CallOpts)
}

// BaseGasLimit is a free data retrieval call binding the contract method 0x8df66e34.
//
// Solidity: function baseGasLimit() view returns(uint64)
func (_FeeOracleV1 *FeeOracleV1Caller) BaseGasLimit(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _FeeOracleV1.contract.Call(opts, &out, "baseGasLimit")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// BaseGasLimit is a free data retrieval call binding the contract method 0x8df66e34.
//
// Solidity: function baseGasLimit() view returns(uint64)
func (_FeeOracleV1 *FeeOracleV1Session) BaseGasLimit() (uint64, error) {
	return _FeeOracleV1.Contract.BaseGasLimit(&_FeeOracleV1.CallOpts)
}

// BaseGasLimit is a free data retrieval call binding the contract method 0x8df66e34.
//
// Solidity: function baseGasLimit() view returns(uint64)
func (_FeeOracleV1 *FeeOracleV1CallerSession) BaseGasLimit() (uint64, error) {
	return _FeeOracleV1.Contract.BaseGasLimit(&_FeeOracleV1.CallOpts)
}

// FeeFor is a free data retrieval call binding the contract method 0x8dd9523c.
//
// Solidity: function feeFor(uint64 destChainId, bytes data, uint64 gasLimit) view returns(uint256)
func (_FeeOracleV1 *FeeOracleV1Caller) FeeFor(opts *bind.CallOpts, destChainId uint64, data []byte, gasLimit uint64) (*big.Int, error) {
	var out []interface{}
	err := _FeeOracleV1.contract.Call(opts, &out, "feeFor", destChainId, data, gasLimit)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FeeFor is a free data retrieval call binding the contract method 0x8dd9523c.
//
// Solidity: function feeFor(uint64 destChainId, bytes data, uint64 gasLimit) view returns(uint256)
func (_FeeOracleV1 *FeeOracleV1Session) FeeFor(destChainId uint64, data []byte, gasLimit uint64) (*big.Int, error) {
	return _FeeOracleV1.Contract.FeeFor(&_FeeOracleV1.CallOpts, destChainId, data, gasLimit)
}

// FeeFor is a free data retrieval call binding the contract method 0x8dd9523c.
//
// Solidity: function feeFor(uint64 destChainId, bytes data, uint64 gasLimit) view returns(uint256)
func (_FeeOracleV1 *FeeOracleV1CallerSession) FeeFor(destChainId uint64, data []byte, gasLimit uint64) (*big.Int, error) {
	return _FeeOracleV1.Contract.FeeFor(&_FeeOracleV1.CallOpts, destChainId, data, gasLimit)
}

// FeeParams is a free data retrieval call binding the contract method 0x2d4634a4.
//
// Solidity: function feeParams(uint64 chainId) view returns((uint64,uint64,uint256,uint256))
func (_FeeOracleV1 *FeeOracleV1Caller) FeeParams(opts *bind.CallOpts, chainId uint64) (IFeeOracleV1ChainFeeParams, error) {
	var out []interface{}
	err := _FeeOracleV1.contract.Call(opts, &out, "feeParams", chainId)

	if err != nil {
		return *new(IFeeOracleV1ChainFeeParams), err
	}

	out0 := *abi.ConvertType(out[0], new(IFeeOracleV1ChainFeeParams)).(*IFeeOracleV1ChainFeeParams)

	return out0, err

}

// FeeParams is a free data retrieval call binding the contract method 0x2d4634a4.
//
// Solidity: function feeParams(uint64 chainId) view returns((uint64,uint64,uint256,uint256))
func (_FeeOracleV1 *FeeOracleV1Session) FeeParams(chainId uint64) (IFeeOracleV1ChainFeeParams, error) {
	return _FeeOracleV1.Contract.FeeParams(&_FeeOracleV1.CallOpts, chainId)
}

// FeeParams is a free data retrieval call binding the contract method 0x2d4634a4.
//
// Solidity: function feeParams(uint64 chainId) view returns((uint64,uint64,uint256,uint256))
func (_FeeOracleV1 *FeeOracleV1CallerSession) FeeParams(chainId uint64) (IFeeOracleV1ChainFeeParams, error) {
	return _FeeOracleV1.Contract.FeeParams(&_FeeOracleV1.CallOpts, chainId)
}

// GasPriceOn is a free data retrieval call binding the contract method 0xd0704157.
//
// Solidity: function gasPriceOn(uint64 chainId) view returns(uint256)
func (_FeeOracleV1 *FeeOracleV1Caller) GasPriceOn(opts *bind.CallOpts, chainId uint64) (*big.Int, error) {
	var out []interface{}
	err := _FeeOracleV1.contract.Call(opts, &out, "gasPriceOn", chainId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GasPriceOn is a free data retrieval call binding the contract method 0xd0704157.
//
// Solidity: function gasPriceOn(uint64 chainId) view returns(uint256)
func (_FeeOracleV1 *FeeOracleV1Session) GasPriceOn(chainId uint64) (*big.Int, error) {
	return _FeeOracleV1.Contract.GasPriceOn(&_FeeOracleV1.CallOpts, chainId)
}

// GasPriceOn is a free data retrieval call binding the contract method 0xd0704157.
//
// Solidity: function gasPriceOn(uint64 chainId) view returns(uint256)
func (_FeeOracleV1 *FeeOracleV1CallerSession) GasPriceOn(chainId uint64) (*big.Int, error) {
	return _FeeOracleV1.Contract.GasPriceOn(&_FeeOracleV1.CallOpts, chainId)
}

// Manager is a free data retrieval call binding the contract method 0x481c6a75.
//
// Solidity: function manager() view returns(address)
func (_FeeOracleV1 *FeeOracleV1Caller) Manager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FeeOracleV1.contract.Call(opts, &out, "manager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Manager is a free data retrieval call binding the contract method 0x481c6a75.
//
// Solidity: function manager() view returns(address)
func (_FeeOracleV1 *FeeOracleV1Session) Manager() (common.Address, error) {
	return _FeeOracleV1.Contract.Manager(&_FeeOracleV1.CallOpts)
}

// Manager is a free data retrieval call binding the contract method 0x481c6a75.
//
// Solidity: function manager() view returns(address)
func (_FeeOracleV1 *FeeOracleV1CallerSession) Manager() (common.Address, error) {
	return _FeeOracleV1.Contract.Manager(&_FeeOracleV1.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_FeeOracleV1 *FeeOracleV1Caller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FeeOracleV1.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_FeeOracleV1 *FeeOracleV1Session) Owner() (common.Address, error) {
	return _FeeOracleV1.Contract.Owner(&_FeeOracleV1.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_FeeOracleV1 *FeeOracleV1CallerSession) Owner() (common.Address, error) {
	return _FeeOracleV1.Contract.Owner(&_FeeOracleV1.CallOpts)
}

// PostsTo is a free data retrieval call binding the contract method 0xb48ec861.
//
// Solidity: function postsTo(uint64 chainId) view returns(uint64)
func (_FeeOracleV1 *FeeOracleV1Caller) PostsTo(opts *bind.CallOpts, chainId uint64) (uint64, error) {
	var out []interface{}
	err := _FeeOracleV1.contract.Call(opts, &out, "postsTo", chainId)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// PostsTo is a free data retrieval call binding the contract method 0xb48ec861.
//
// Solidity: function postsTo(uint64 chainId) view returns(uint64)
func (_FeeOracleV1 *FeeOracleV1Session) PostsTo(chainId uint64) (uint64, error) {
	return _FeeOracleV1.Contract.PostsTo(&_FeeOracleV1.CallOpts, chainId)
}

// PostsTo is a free data retrieval call binding the contract method 0xb48ec861.
//
// Solidity: function postsTo(uint64 chainId) view returns(uint64)
func (_FeeOracleV1 *FeeOracleV1CallerSession) PostsTo(chainId uint64) (uint64, error) {
	return _FeeOracleV1.Contract.PostsTo(&_FeeOracleV1.CallOpts, chainId)
}

// ProtocolFee is a free data retrieval call binding the contract method 0xb0e21e8a.
//
// Solidity: function protocolFee() view returns(uint256)
func (_FeeOracleV1 *FeeOracleV1Caller) ProtocolFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FeeOracleV1.contract.Call(opts, &out, "protocolFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ProtocolFee is a free data retrieval call binding the contract method 0xb0e21e8a.
//
// Solidity: function protocolFee() view returns(uint256)
func (_FeeOracleV1 *FeeOracleV1Session) ProtocolFee() (*big.Int, error) {
	return _FeeOracleV1.Contract.ProtocolFee(&_FeeOracleV1.CallOpts)
}

// ProtocolFee is a free data retrieval call binding the contract method 0xb0e21e8a.
//
// Solidity: function protocolFee() view returns(uint256)
func (_FeeOracleV1 *FeeOracleV1CallerSession) ProtocolFee() (*big.Int, error) {
	return _FeeOracleV1.Contract.ProtocolFee(&_FeeOracleV1.CallOpts)
}

// ToNativeRate is a free data retrieval call binding the contract method 0x8b7bfd70.
//
// Solidity: function toNativeRate(uint64 chainId) view returns(uint256)
func (_FeeOracleV1 *FeeOracleV1Caller) ToNativeRate(opts *bind.CallOpts, chainId uint64) (*big.Int, error) {
	var out []interface{}
	err := _FeeOracleV1.contract.Call(opts, &out, "toNativeRate", chainId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ToNativeRate is a free data retrieval call binding the contract method 0x8b7bfd70.
//
// Solidity: function toNativeRate(uint64 chainId) view returns(uint256)
func (_FeeOracleV1 *FeeOracleV1Session) ToNativeRate(chainId uint64) (*big.Int, error) {
	return _FeeOracleV1.Contract.ToNativeRate(&_FeeOracleV1.CallOpts, chainId)
}

// ToNativeRate is a free data retrieval call binding the contract method 0x8b7bfd70.
//
// Solidity: function toNativeRate(uint64 chainId) view returns(uint256)
func (_FeeOracleV1 *FeeOracleV1CallerSession) ToNativeRate(chainId uint64) (*big.Int, error) {
	return _FeeOracleV1.Contract.ToNativeRate(&_FeeOracleV1.CallOpts, chainId)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(uint64)
func (_FeeOracleV1 *FeeOracleV1Caller) Version(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _FeeOracleV1.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(uint64)
func (_FeeOracleV1 *FeeOracleV1Session) Version() (uint64, error) {
	return _FeeOracleV1.Contract.Version(&_FeeOracleV1.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(uint64)
func (_FeeOracleV1 *FeeOracleV1CallerSession) Version() (uint64, error) {
	return _FeeOracleV1.Contract.Version(&_FeeOracleV1.CallOpts)
}

// BulkSetFeeParams is a paid mutator transaction binding the contract method 0x93a87188.
//
// Solidity: function bulkSetFeeParams((uint64,uint64,uint256,uint256)[] params) returns()
func (_FeeOracleV1 *FeeOracleV1Transactor) BulkSetFeeParams(opts *bind.TransactOpts, params []IFeeOracleV1ChainFeeParams) (*types.Transaction, error) {
	return _FeeOracleV1.contract.Transact(opts, "bulkSetFeeParams", params)
}

// BulkSetFeeParams is a paid mutator transaction binding the contract method 0x93a87188.
//
// Solidity: function bulkSetFeeParams((uint64,uint64,uint256,uint256)[] params) returns()
func (_FeeOracleV1 *FeeOracleV1Session) BulkSetFeeParams(params []IFeeOracleV1ChainFeeParams) (*types.Transaction, error) {
	return _FeeOracleV1.Contract.BulkSetFeeParams(&_FeeOracleV1.TransactOpts, params)
}

// BulkSetFeeParams is a paid mutator transaction binding the contract method 0x93a87188.
//
// Solidity: function bulkSetFeeParams((uint64,uint64,uint256,uint256)[] params) returns()
func (_FeeOracleV1 *FeeOracleV1TransactorSession) BulkSetFeeParams(params []IFeeOracleV1ChainFeeParams) (*types.Transaction, error) {
	return _FeeOracleV1.Contract.BulkSetFeeParams(&_FeeOracleV1.TransactOpts, params)
}

// Initialize is a paid mutator transaction binding the contract method 0x361c019f.
//
// Solidity: function initialize(address owner_, address manager_, uint64 baseGasLimit_, uint256 protocolFee_, (uint64,uint64,uint256,uint256)[] params) returns()
func (_FeeOracleV1 *FeeOracleV1Transactor) Initialize(opts *bind.TransactOpts, owner_ common.Address, manager_ common.Address, baseGasLimit_ uint64, protocolFee_ *big.Int, params []IFeeOracleV1ChainFeeParams) (*types.Transaction, error) {
	return _FeeOracleV1.contract.Transact(opts, "initialize", owner_, manager_, baseGasLimit_, protocolFee_, params)
}

// Initialize is a paid mutator transaction binding the contract method 0x361c019f.
//
// Solidity: function initialize(address owner_, address manager_, uint64 baseGasLimit_, uint256 protocolFee_, (uint64,uint64,uint256,uint256)[] params) returns()
func (_FeeOracleV1 *FeeOracleV1Session) Initialize(owner_ common.Address, manager_ common.Address, baseGasLimit_ uint64, protocolFee_ *big.Int, params []IFeeOracleV1ChainFeeParams) (*types.Transaction, error) {
	return _FeeOracleV1.Contract.Initialize(&_FeeOracleV1.TransactOpts, owner_, manager_, baseGasLimit_, protocolFee_, params)
}

// Initialize is a paid mutator transaction binding the contract method 0x361c019f.
//
// Solidity: function initialize(address owner_, address manager_, uint64 baseGasLimit_, uint256 protocolFee_, (uint64,uint64,uint256,uint256)[] params) returns()
func (_FeeOracleV1 *FeeOracleV1TransactorSession) Initialize(owner_ common.Address, manager_ common.Address, baseGasLimit_ uint64, protocolFee_ *big.Int, params []IFeeOracleV1ChainFeeParams) (*types.Transaction, error) {
	return _FeeOracleV1.Contract.Initialize(&_FeeOracleV1.TransactOpts, owner_, manager_, baseGasLimit_, protocolFee_, params)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_FeeOracleV1 *FeeOracleV1Transactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FeeOracleV1.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_FeeOracleV1 *FeeOracleV1Session) RenounceOwnership() (*types.Transaction, error) {
	return _FeeOracleV1.Contract.RenounceOwnership(&_FeeOracleV1.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_FeeOracleV1 *FeeOracleV1TransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _FeeOracleV1.Contract.RenounceOwnership(&_FeeOracleV1.TransactOpts)
}

// SetBaseGasLimit is a paid mutator transaction binding the contract method 0xee590a53.
//
// Solidity: function setBaseGasLimit(uint64 gasLimit) returns()
func (_FeeOracleV1 *FeeOracleV1Transactor) SetBaseGasLimit(opts *bind.TransactOpts, gasLimit uint64) (*types.Transaction, error) {
	return _FeeOracleV1.contract.Transact(opts, "setBaseGasLimit", gasLimit)
}

// SetBaseGasLimit is a paid mutator transaction binding the contract method 0xee590a53.
//
// Solidity: function setBaseGasLimit(uint64 gasLimit) returns()
func (_FeeOracleV1 *FeeOracleV1Session) SetBaseGasLimit(gasLimit uint64) (*types.Transaction, error) {
	return _FeeOracleV1.Contract.SetBaseGasLimit(&_FeeOracleV1.TransactOpts, gasLimit)
}

// SetBaseGasLimit is a paid mutator transaction binding the contract method 0xee590a53.
//
// Solidity: function setBaseGasLimit(uint64 gasLimit) returns()
func (_FeeOracleV1 *FeeOracleV1TransactorSession) SetBaseGasLimit(gasLimit uint64) (*types.Transaction, error) {
	return _FeeOracleV1.Contract.SetBaseGasLimit(&_FeeOracleV1.TransactOpts, gasLimit)
}

// SetGasPrice is a paid mutator transaction binding the contract method 0x98563b03.
//
// Solidity: function setGasPrice(uint64 chainId, uint256 gasPrice) returns()
func (_FeeOracleV1 *FeeOracleV1Transactor) SetGasPrice(opts *bind.TransactOpts, chainId uint64, gasPrice *big.Int) (*types.Transaction, error) {
	return _FeeOracleV1.contract.Transact(opts, "setGasPrice", chainId, gasPrice)
}

// SetGasPrice is a paid mutator transaction binding the contract method 0x98563b03.
//
// Solidity: function setGasPrice(uint64 chainId, uint256 gasPrice) returns()
func (_FeeOracleV1 *FeeOracleV1Session) SetGasPrice(chainId uint64, gasPrice *big.Int) (*types.Transaction, error) {
	return _FeeOracleV1.Contract.SetGasPrice(&_FeeOracleV1.TransactOpts, chainId, gasPrice)
}

// SetGasPrice is a paid mutator transaction binding the contract method 0x98563b03.
//
// Solidity: function setGasPrice(uint64 chainId, uint256 gasPrice) returns()
func (_FeeOracleV1 *FeeOracleV1TransactorSession) SetGasPrice(chainId uint64, gasPrice *big.Int) (*types.Transaction, error) {
	return _FeeOracleV1.Contract.SetGasPrice(&_FeeOracleV1.TransactOpts, chainId, gasPrice)
}

// SetManager is a paid mutator transaction binding the contract method 0xd0ebdbe7.
//
// Solidity: function setManager(address manager_) returns()
func (_FeeOracleV1 *FeeOracleV1Transactor) SetManager(opts *bind.TransactOpts, manager_ common.Address) (*types.Transaction, error) {
	return _FeeOracleV1.contract.Transact(opts, "setManager", manager_)
}

// SetManager is a paid mutator transaction binding the contract method 0xd0ebdbe7.
//
// Solidity: function setManager(address manager_) returns()
func (_FeeOracleV1 *FeeOracleV1Session) SetManager(manager_ common.Address) (*types.Transaction, error) {
	return _FeeOracleV1.Contract.SetManager(&_FeeOracleV1.TransactOpts, manager_)
}

// SetManager is a paid mutator transaction binding the contract method 0xd0ebdbe7.
//
// Solidity: function setManager(address manager_) returns()
func (_FeeOracleV1 *FeeOracleV1TransactorSession) SetManager(manager_ common.Address) (*types.Transaction, error) {
	return _FeeOracleV1.Contract.SetManager(&_FeeOracleV1.TransactOpts, manager_)
}

// SetProtocolFee is a paid mutator transaction binding the contract method 0x787dce3d.
//
// Solidity: function setProtocolFee(uint256 fee) returns()
func (_FeeOracleV1 *FeeOracleV1Transactor) SetProtocolFee(opts *bind.TransactOpts, fee *big.Int) (*types.Transaction, error) {
	return _FeeOracleV1.contract.Transact(opts, "setProtocolFee", fee)
}

// SetProtocolFee is a paid mutator transaction binding the contract method 0x787dce3d.
//
// Solidity: function setProtocolFee(uint256 fee) returns()
func (_FeeOracleV1 *FeeOracleV1Session) SetProtocolFee(fee *big.Int) (*types.Transaction, error) {
	return _FeeOracleV1.Contract.SetProtocolFee(&_FeeOracleV1.TransactOpts, fee)
}

// SetProtocolFee is a paid mutator transaction binding the contract method 0x787dce3d.
//
// Solidity: function setProtocolFee(uint256 fee) returns()
func (_FeeOracleV1 *FeeOracleV1TransactorSession) SetProtocolFee(fee *big.Int) (*types.Transaction, error) {
	return _FeeOracleV1.Contract.SetProtocolFee(&_FeeOracleV1.TransactOpts, fee)
}

// SetToNativeRate is a paid mutator transaction binding the contract method 0xa34e7abb.
//
// Solidity: function setToNativeRate(uint64 chainId, uint256 rate) returns()
func (_FeeOracleV1 *FeeOracleV1Transactor) SetToNativeRate(opts *bind.TransactOpts, chainId uint64, rate *big.Int) (*types.Transaction, error) {
	return _FeeOracleV1.contract.Transact(opts, "setToNativeRate", chainId, rate)
}

// SetToNativeRate is a paid mutator transaction binding the contract method 0xa34e7abb.
//
// Solidity: function setToNativeRate(uint64 chainId, uint256 rate) returns()
func (_FeeOracleV1 *FeeOracleV1Session) SetToNativeRate(chainId uint64, rate *big.Int) (*types.Transaction, error) {
	return _FeeOracleV1.Contract.SetToNativeRate(&_FeeOracleV1.TransactOpts, chainId, rate)
}

// SetToNativeRate is a paid mutator transaction binding the contract method 0xa34e7abb.
//
// Solidity: function setToNativeRate(uint64 chainId, uint256 rate) returns()
func (_FeeOracleV1 *FeeOracleV1TransactorSession) SetToNativeRate(chainId uint64, rate *big.Int) (*types.Transaction, error) {
	return _FeeOracleV1.Contract.SetToNativeRate(&_FeeOracleV1.TransactOpts, chainId, rate)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_FeeOracleV1 *FeeOracleV1Transactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _FeeOracleV1.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_FeeOracleV1 *FeeOracleV1Session) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _FeeOracleV1.Contract.TransferOwnership(&_FeeOracleV1.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_FeeOracleV1 *FeeOracleV1TransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _FeeOracleV1.Contract.TransferOwnership(&_FeeOracleV1.TransactOpts, newOwner)
}

// FeeOracleV1BaseGasLimitSetIterator is returned from FilterBaseGasLimitSet and is used to iterate over the raw logs and unpacked data for BaseGasLimitSet events raised by the FeeOracleV1 contract.
type FeeOracleV1BaseGasLimitSetIterator struct {
	Event *FeeOracleV1BaseGasLimitSet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *FeeOracleV1BaseGasLimitSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeeOracleV1BaseGasLimitSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(FeeOracleV1BaseGasLimitSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *FeeOracleV1BaseGasLimitSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FeeOracleV1BaseGasLimitSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FeeOracleV1BaseGasLimitSet represents a BaseGasLimitSet event raised by the FeeOracleV1 contract.
type FeeOracleV1BaseGasLimitSet struct {
	BaseGasLimit uint64
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterBaseGasLimitSet is a free log retrieval operation binding the contract event 0x6185fbe062d94552cf644f5cb643f583db7b2e7e66fdc4b4c75ff8876a257ba6.
//
// Solidity: event BaseGasLimitSet(uint64 baseGasLimit)
func (_FeeOracleV1 *FeeOracleV1Filterer) FilterBaseGasLimitSet(opts *bind.FilterOpts) (*FeeOracleV1BaseGasLimitSetIterator, error) {

	logs, sub, err := _FeeOracleV1.contract.FilterLogs(opts, "BaseGasLimitSet")
	if err != nil {
		return nil, err
	}
	return &FeeOracleV1BaseGasLimitSetIterator{contract: _FeeOracleV1.contract, event: "BaseGasLimitSet", logs: logs, sub: sub}, nil
}

// WatchBaseGasLimitSet is a free log subscription operation binding the contract event 0x6185fbe062d94552cf644f5cb643f583db7b2e7e66fdc4b4c75ff8876a257ba6.
//
// Solidity: event BaseGasLimitSet(uint64 baseGasLimit)
func (_FeeOracleV1 *FeeOracleV1Filterer) WatchBaseGasLimitSet(opts *bind.WatchOpts, sink chan<- *FeeOracleV1BaseGasLimitSet) (event.Subscription, error) {

	logs, sub, err := _FeeOracleV1.contract.WatchLogs(opts, "BaseGasLimitSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FeeOracleV1BaseGasLimitSet)
				if err := _FeeOracleV1.contract.UnpackLog(event, "BaseGasLimitSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseBaseGasLimitSet is a log parse operation binding the contract event 0x6185fbe062d94552cf644f5cb643f583db7b2e7e66fdc4b4c75ff8876a257ba6.
//
// Solidity: event BaseGasLimitSet(uint64 baseGasLimit)
func (_FeeOracleV1 *FeeOracleV1Filterer) ParseBaseGasLimitSet(log types.Log) (*FeeOracleV1BaseGasLimitSet, error) {
	event := new(FeeOracleV1BaseGasLimitSet)
	if err := _FeeOracleV1.contract.UnpackLog(event, "BaseGasLimitSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FeeOracleV1FeeParamsSetIterator is returned from FilterFeeParamsSet and is used to iterate over the raw logs and unpacked data for FeeParamsSet events raised by the FeeOracleV1 contract.
type FeeOracleV1FeeParamsSetIterator struct {
	Event *FeeOracleV1FeeParamsSet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *FeeOracleV1FeeParamsSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeeOracleV1FeeParamsSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(FeeOracleV1FeeParamsSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *FeeOracleV1FeeParamsSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FeeOracleV1FeeParamsSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FeeOracleV1FeeParamsSet represents a FeeParamsSet event raised by the FeeOracleV1 contract.
type FeeOracleV1FeeParamsSet struct {
	ChainId      uint64
	PostsTo      uint64
	GasPrice     *big.Int
	ToNativeRate *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterFeeParamsSet is a free log retrieval operation binding the contract event 0xf378a0dd98429494eb2e26894562949c4d6e7cef5eb893b1d4c0052078d92fe5.
//
// Solidity: event FeeParamsSet(uint64 chainId, uint64 postsTo, uint256 gasPrice, uint256 toNativeRate)
func (_FeeOracleV1 *FeeOracleV1Filterer) FilterFeeParamsSet(opts *bind.FilterOpts) (*FeeOracleV1FeeParamsSetIterator, error) {

	logs, sub, err := _FeeOracleV1.contract.FilterLogs(opts, "FeeParamsSet")
	if err != nil {
		return nil, err
	}
	return &FeeOracleV1FeeParamsSetIterator{contract: _FeeOracleV1.contract, event: "FeeParamsSet", logs: logs, sub: sub}, nil
}

// WatchFeeParamsSet is a free log subscription operation binding the contract event 0xf378a0dd98429494eb2e26894562949c4d6e7cef5eb893b1d4c0052078d92fe5.
//
// Solidity: event FeeParamsSet(uint64 chainId, uint64 postsTo, uint256 gasPrice, uint256 toNativeRate)
func (_FeeOracleV1 *FeeOracleV1Filterer) WatchFeeParamsSet(opts *bind.WatchOpts, sink chan<- *FeeOracleV1FeeParamsSet) (event.Subscription, error) {

	logs, sub, err := _FeeOracleV1.contract.WatchLogs(opts, "FeeParamsSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FeeOracleV1FeeParamsSet)
				if err := _FeeOracleV1.contract.UnpackLog(event, "FeeParamsSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseFeeParamsSet is a log parse operation binding the contract event 0xf378a0dd98429494eb2e26894562949c4d6e7cef5eb893b1d4c0052078d92fe5.
//
// Solidity: event FeeParamsSet(uint64 chainId, uint64 postsTo, uint256 gasPrice, uint256 toNativeRate)
func (_FeeOracleV1 *FeeOracleV1Filterer) ParseFeeParamsSet(log types.Log) (*FeeOracleV1FeeParamsSet, error) {
	event := new(FeeOracleV1FeeParamsSet)
	if err := _FeeOracleV1.contract.UnpackLog(event, "FeeParamsSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FeeOracleV1GasPriceSetIterator is returned from FilterGasPriceSet and is used to iterate over the raw logs and unpacked data for GasPriceSet events raised by the FeeOracleV1 contract.
type FeeOracleV1GasPriceSetIterator struct {
	Event *FeeOracleV1GasPriceSet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *FeeOracleV1GasPriceSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeeOracleV1GasPriceSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(FeeOracleV1GasPriceSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *FeeOracleV1GasPriceSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FeeOracleV1GasPriceSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FeeOracleV1GasPriceSet represents a GasPriceSet event raised by the FeeOracleV1 contract.
type FeeOracleV1GasPriceSet struct {
	ChainId  uint64
	GasPrice *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterGasPriceSet is a free log retrieval operation binding the contract event 0x3b196e45eaa29099834d3d912ac550e4f3e13fef2e2a998100368e506a44d8ff.
//
// Solidity: event GasPriceSet(uint64 chainId, uint256 gasPrice)
func (_FeeOracleV1 *FeeOracleV1Filterer) FilterGasPriceSet(opts *bind.FilterOpts) (*FeeOracleV1GasPriceSetIterator, error) {

	logs, sub, err := _FeeOracleV1.contract.FilterLogs(opts, "GasPriceSet")
	if err != nil {
		return nil, err
	}
	return &FeeOracleV1GasPriceSetIterator{contract: _FeeOracleV1.contract, event: "GasPriceSet", logs: logs, sub: sub}, nil
}

// WatchGasPriceSet is a free log subscription operation binding the contract event 0x3b196e45eaa29099834d3d912ac550e4f3e13fef2e2a998100368e506a44d8ff.
//
// Solidity: event GasPriceSet(uint64 chainId, uint256 gasPrice)
func (_FeeOracleV1 *FeeOracleV1Filterer) WatchGasPriceSet(opts *bind.WatchOpts, sink chan<- *FeeOracleV1GasPriceSet) (event.Subscription, error) {

	logs, sub, err := _FeeOracleV1.contract.WatchLogs(opts, "GasPriceSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FeeOracleV1GasPriceSet)
				if err := _FeeOracleV1.contract.UnpackLog(event, "GasPriceSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseGasPriceSet is a log parse operation binding the contract event 0x3b196e45eaa29099834d3d912ac550e4f3e13fef2e2a998100368e506a44d8ff.
//
// Solidity: event GasPriceSet(uint64 chainId, uint256 gasPrice)
func (_FeeOracleV1 *FeeOracleV1Filterer) ParseGasPriceSet(log types.Log) (*FeeOracleV1GasPriceSet, error) {
	event := new(FeeOracleV1GasPriceSet)
	if err := _FeeOracleV1.contract.UnpackLog(event, "GasPriceSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FeeOracleV1InitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the FeeOracleV1 contract.
type FeeOracleV1InitializedIterator struct {
	Event *FeeOracleV1Initialized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *FeeOracleV1InitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeeOracleV1Initialized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(FeeOracleV1Initialized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *FeeOracleV1InitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FeeOracleV1InitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FeeOracleV1Initialized represents a Initialized event raised by the FeeOracleV1 contract.
type FeeOracleV1Initialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_FeeOracleV1 *FeeOracleV1Filterer) FilterInitialized(opts *bind.FilterOpts) (*FeeOracleV1InitializedIterator, error) {

	logs, sub, err := _FeeOracleV1.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &FeeOracleV1InitializedIterator{contract: _FeeOracleV1.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_FeeOracleV1 *FeeOracleV1Filterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *FeeOracleV1Initialized) (event.Subscription, error) {

	logs, sub, err := _FeeOracleV1.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FeeOracleV1Initialized)
				if err := _FeeOracleV1.contract.UnpackLog(event, "Initialized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInitialized is a log parse operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_FeeOracleV1 *FeeOracleV1Filterer) ParseInitialized(log types.Log) (*FeeOracleV1Initialized, error) {
	event := new(FeeOracleV1Initialized)
	if err := _FeeOracleV1.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FeeOracleV1ManagerSetIterator is returned from FilterManagerSet and is used to iterate over the raw logs and unpacked data for ManagerSet events raised by the FeeOracleV1 contract.
type FeeOracleV1ManagerSetIterator struct {
	Event *FeeOracleV1ManagerSet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *FeeOracleV1ManagerSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeeOracleV1ManagerSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(FeeOracleV1ManagerSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *FeeOracleV1ManagerSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FeeOracleV1ManagerSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FeeOracleV1ManagerSet represents a ManagerSet event raised by the FeeOracleV1 contract.
type FeeOracleV1ManagerSet struct {
	Manager common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterManagerSet is a free log retrieval operation binding the contract event 0x60a0f5b9f9e81e98216071b85826681c796256fe3d1354ecb675580fba64fa69.
//
// Solidity: event ManagerSet(address manager)
func (_FeeOracleV1 *FeeOracleV1Filterer) FilterManagerSet(opts *bind.FilterOpts) (*FeeOracleV1ManagerSetIterator, error) {

	logs, sub, err := _FeeOracleV1.contract.FilterLogs(opts, "ManagerSet")
	if err != nil {
		return nil, err
	}
	return &FeeOracleV1ManagerSetIterator{contract: _FeeOracleV1.contract, event: "ManagerSet", logs: logs, sub: sub}, nil
}

// WatchManagerSet is a free log subscription operation binding the contract event 0x60a0f5b9f9e81e98216071b85826681c796256fe3d1354ecb675580fba64fa69.
//
// Solidity: event ManagerSet(address manager)
func (_FeeOracleV1 *FeeOracleV1Filterer) WatchManagerSet(opts *bind.WatchOpts, sink chan<- *FeeOracleV1ManagerSet) (event.Subscription, error) {

	logs, sub, err := _FeeOracleV1.contract.WatchLogs(opts, "ManagerSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FeeOracleV1ManagerSet)
				if err := _FeeOracleV1.contract.UnpackLog(event, "ManagerSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseManagerSet is a log parse operation binding the contract event 0x60a0f5b9f9e81e98216071b85826681c796256fe3d1354ecb675580fba64fa69.
//
// Solidity: event ManagerSet(address manager)
func (_FeeOracleV1 *FeeOracleV1Filterer) ParseManagerSet(log types.Log) (*FeeOracleV1ManagerSet, error) {
	event := new(FeeOracleV1ManagerSet)
	if err := _FeeOracleV1.contract.UnpackLog(event, "ManagerSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FeeOracleV1OwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the FeeOracleV1 contract.
type FeeOracleV1OwnershipTransferredIterator struct {
	Event *FeeOracleV1OwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *FeeOracleV1OwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeeOracleV1OwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(FeeOracleV1OwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *FeeOracleV1OwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FeeOracleV1OwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FeeOracleV1OwnershipTransferred represents a OwnershipTransferred event raised by the FeeOracleV1 contract.
type FeeOracleV1OwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_FeeOracleV1 *FeeOracleV1Filterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*FeeOracleV1OwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _FeeOracleV1.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &FeeOracleV1OwnershipTransferredIterator{contract: _FeeOracleV1.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_FeeOracleV1 *FeeOracleV1Filterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *FeeOracleV1OwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _FeeOracleV1.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FeeOracleV1OwnershipTransferred)
				if err := _FeeOracleV1.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_FeeOracleV1 *FeeOracleV1Filterer) ParseOwnershipTransferred(log types.Log) (*FeeOracleV1OwnershipTransferred, error) {
	event := new(FeeOracleV1OwnershipTransferred)
	if err := _FeeOracleV1.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FeeOracleV1ProtocolFeeSetIterator is returned from FilterProtocolFeeSet and is used to iterate over the raw logs and unpacked data for ProtocolFeeSet events raised by the FeeOracleV1 contract.
type FeeOracleV1ProtocolFeeSetIterator struct {
	Event *FeeOracleV1ProtocolFeeSet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *FeeOracleV1ProtocolFeeSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeeOracleV1ProtocolFeeSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(FeeOracleV1ProtocolFeeSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *FeeOracleV1ProtocolFeeSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FeeOracleV1ProtocolFeeSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FeeOracleV1ProtocolFeeSet represents a ProtocolFeeSet event raised by the FeeOracleV1 contract.
type FeeOracleV1ProtocolFeeSet struct {
	ProtocolFee *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterProtocolFeeSet is a free log retrieval operation binding the contract event 0xdb5aafdb29539329e37d4e3ee869bc4031941fd55a5dfc92824fbe34b204e30d.
//
// Solidity: event ProtocolFeeSet(uint256 protocolFee)
func (_FeeOracleV1 *FeeOracleV1Filterer) FilterProtocolFeeSet(opts *bind.FilterOpts) (*FeeOracleV1ProtocolFeeSetIterator, error) {

	logs, sub, err := _FeeOracleV1.contract.FilterLogs(opts, "ProtocolFeeSet")
	if err != nil {
		return nil, err
	}
	return &FeeOracleV1ProtocolFeeSetIterator{contract: _FeeOracleV1.contract, event: "ProtocolFeeSet", logs: logs, sub: sub}, nil
}

// WatchProtocolFeeSet is a free log subscription operation binding the contract event 0xdb5aafdb29539329e37d4e3ee869bc4031941fd55a5dfc92824fbe34b204e30d.
//
// Solidity: event ProtocolFeeSet(uint256 protocolFee)
func (_FeeOracleV1 *FeeOracleV1Filterer) WatchProtocolFeeSet(opts *bind.WatchOpts, sink chan<- *FeeOracleV1ProtocolFeeSet) (event.Subscription, error) {

	logs, sub, err := _FeeOracleV1.contract.WatchLogs(opts, "ProtocolFeeSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FeeOracleV1ProtocolFeeSet)
				if err := _FeeOracleV1.contract.UnpackLog(event, "ProtocolFeeSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseProtocolFeeSet is a log parse operation binding the contract event 0xdb5aafdb29539329e37d4e3ee869bc4031941fd55a5dfc92824fbe34b204e30d.
//
// Solidity: event ProtocolFeeSet(uint256 protocolFee)
func (_FeeOracleV1 *FeeOracleV1Filterer) ParseProtocolFeeSet(log types.Log) (*FeeOracleV1ProtocolFeeSet, error) {
	event := new(FeeOracleV1ProtocolFeeSet)
	if err := _FeeOracleV1.contract.UnpackLog(event, "ProtocolFeeSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FeeOracleV1ToNativeRateSetIterator is returned from FilterToNativeRateSet and is used to iterate over the raw logs and unpacked data for ToNativeRateSet events raised by the FeeOracleV1 contract.
type FeeOracleV1ToNativeRateSetIterator struct {
	Event *FeeOracleV1ToNativeRateSet // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *FeeOracleV1ToNativeRateSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeeOracleV1ToNativeRateSet)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(FeeOracleV1ToNativeRateSet)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *FeeOracleV1ToNativeRateSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FeeOracleV1ToNativeRateSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FeeOracleV1ToNativeRateSet represents a ToNativeRateSet event raised by the FeeOracleV1 contract.
type FeeOracleV1ToNativeRateSet struct {
	ChainId      uint64
	ToNativeRate *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterToNativeRateSet is a free log retrieval operation binding the contract event 0x4b4594c9f06af25bc504eead96f7f0eaa3f1577f8d9b075b236520ec712e1308.
//
// Solidity: event ToNativeRateSet(uint64 chainId, uint256 toNativeRate)
func (_FeeOracleV1 *FeeOracleV1Filterer) FilterToNativeRateSet(opts *bind.FilterOpts) (*FeeOracleV1ToNativeRateSetIterator, error) {

	logs, sub, err := _FeeOracleV1.contract.FilterLogs(opts, "ToNativeRateSet")
	if err != nil {
		return nil, err
	}
	return &FeeOracleV1ToNativeRateSetIterator{contract: _FeeOracleV1.contract, event: "ToNativeRateSet", logs: logs, sub: sub}, nil
}

// WatchToNativeRateSet is a free log subscription operation binding the contract event 0x4b4594c9f06af25bc504eead96f7f0eaa3f1577f8d9b075b236520ec712e1308.
//
// Solidity: event ToNativeRateSet(uint64 chainId, uint256 toNativeRate)
func (_FeeOracleV1 *FeeOracleV1Filterer) WatchToNativeRateSet(opts *bind.WatchOpts, sink chan<- *FeeOracleV1ToNativeRateSet) (event.Subscription, error) {

	logs, sub, err := _FeeOracleV1.contract.WatchLogs(opts, "ToNativeRateSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FeeOracleV1ToNativeRateSet)
				if err := _FeeOracleV1.contract.UnpackLog(event, "ToNativeRateSet", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseToNativeRateSet is a log parse operation binding the contract event 0x4b4594c9f06af25bc504eead96f7f0eaa3f1577f8d9b075b236520ec712e1308.
//
// Solidity: event ToNativeRateSet(uint64 chainId, uint256 toNativeRate)
func (_FeeOracleV1 *FeeOracleV1Filterer) ParseToNativeRateSet(log types.Log) (*FeeOracleV1ToNativeRateSet, error) {
	event := new(FeeOracleV1ToNativeRateSet)
	if err := _FeeOracleV1.contract.UnpackLog(event, "ToNativeRateSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
