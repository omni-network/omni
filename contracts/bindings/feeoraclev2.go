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

// IFeeOracleV2DataFeeParams is an auto generated low-level Go binding around an user-defined struct.
type IFeeOracleV2DataFeeParams struct {
	ChainId      uint64
	SizeBuffer   uint64
	DataGasPrice uint64
	ToNativeRate uint64
}

// IFeeOracleV2ExecFeeParams is an auto generated low-level Go binding around an user-defined struct.
type IFeeOracleV2ExecFeeParams struct {
	ChainId      uint64
	PostsTo      uint64
	ExecGasPrice uint64
	ToNativeRate uint64
}

// IFeeOracleV2FeeParams is an auto generated low-level Go binding around an user-defined struct.
type IFeeOracleV2FeeParams struct {
	ExecChainId      uint64
	ExecPostsTo      uint64
	ExecGasPrice     uint64
	ExecToNativeRate uint64
	DataChainId      uint64
	DataSizeBuffer   uint64
	DataGasPrice     uint64
	DataToNativeRate uint64
}

// FeeOracleV2MetaData contains all meta data concerning the FeeOracleV2 contract.
var FeeOracleV2MetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"CONVERSION_RATE_DENOM\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"baseGasLimit\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint24\",\"internalType\":\"uint24\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"bulkSetDataFeeParams\",\"inputs\":[{\"name\":\"params\",\"type\":\"tuple[]\",\"internalType\":\"structIFeeOracleV2.DataFeeParams[]\",\"components\":[{\"name\":\"chainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"sizeBuffer\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"dataGasPrice\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"toNativeRate\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"bulkSetExecFeeParams\",\"inputs\":[{\"name\":\"params\",\"type\":\"tuple[]\",\"internalType\":\"structIFeeOracleV2.ExecFeeParams[]\",\"components\":[{\"name\":\"chainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"postsTo\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"execGasPrice\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"toNativeRate\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"dataGasPrice\",\"inputs\":[{\"name\":\"chainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"dataSizeBuffer\",\"inputs\":[{\"name\":\"chainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"dataToNativeRate\",\"inputs\":[{\"name\":\"chainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"execGasPrice\",\"inputs\":[{\"name\":\"chainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"execPostsTo\",\"inputs\":[{\"name\":\"chainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"execToNativeRate\",\"inputs\":[{\"name\":\"chainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"feeFor\",\"inputs\":[{\"name\":\"destChainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"gasLimit\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"feeParams\",\"inputs\":[{\"name\":\"chainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIFeeOracleV2.FeeParams\",\"components\":[{\"name\":\"execChainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"execPostsTo\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"execGasPrice\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"execToNativeRate\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"dataChainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"dataSizeBuffer\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"dataGasPrice\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"dataToNativeRate\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"owner_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"manager_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"baseGasLimit_\",\"type\":\"uint24\",\"internalType\":\"uint24\"},{\"name\":\"protocolFee_\",\"type\":\"uint72\",\"internalType\":\"uint72\"},{\"name\":\"execParams\",\"type\":\"tuple[]\",\"internalType\":\"structIFeeOracleV2.ExecFeeParams[]\",\"components\":[{\"name\":\"chainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"postsTo\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"execGasPrice\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"toNativeRate\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"dataParams\",\"type\":\"tuple[]\",\"internalType\":\"structIFeeOracleV2.DataFeeParams[]\",\"components\":[{\"name\":\"chainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"sizeBuffer\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"dataGasPrice\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"toNativeRate\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"manager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"protocolFee\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint72\",\"internalType\":\"uint72\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setBaseGasLimit\",\"inputs\":[{\"name\":\"gasLimit\",\"type\":\"uint24\",\"internalType\":\"uint24\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setDataGasPrice\",\"inputs\":[{\"name\":\"chainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"gasPrice\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setDataSizeBuffer\",\"inputs\":[{\"name\":\"chainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"sizeBuffer\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setDataToNativeRate\",\"inputs\":[{\"name\":\"chainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"rate\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setExecGasPrice\",\"inputs\":[{\"name\":\"chainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"gasPrice\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setExecPostsTo\",\"inputs\":[{\"name\":\"chainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"postsTo\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setExecToNativeRate\",\"inputs\":[{\"name\":\"chainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"rate\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setManager\",\"inputs\":[{\"name\":\"manager_\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setProtocolFee\",\"inputs\":[{\"name\":\"fee\",\"type\":\"uint72\",\"internalType\":\"uint72\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"version\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"pure\"},{\"type\":\"event\",\"name\":\"BaseGasLimitSet\",\"inputs\":[{\"name\":\"baseGasLimit\",\"type\":\"uint24\",\"indexed\":false,\"internalType\":\"uint24\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DataFeeParamsSet\",\"inputs\":[{\"name\":\"chainId\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"sizeBuffer\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"dataGasPrice\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"toNativeRate\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DataGasPriceSet\",\"inputs\":[{\"name\":\"chainId\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"gasPrice\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DataSizeBufferSet\",\"inputs\":[{\"name\":\"chainId\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"sizeBuffer\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DataToNativeRateSet\",\"inputs\":[{\"name\":\"chainId\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"toNativeRate\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ExecFeeParamsSet\",\"inputs\":[{\"name\":\"chainId\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"postsTo\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"execGasPrice\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"toNativeRate\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ExecGasPriceSet\",\"inputs\":[{\"name\":\"chainId\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"gasPrice\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ExecPostsToSet\",\"inputs\":[{\"name\":\"chainId\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"postsTo\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ExecToNativeRateSet\",\"inputs\":[{\"name\":\"chainId\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"toNativeRate\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ManagerSet\",\"inputs\":[{\"name\":\"manager\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ProtocolFeeSet\",\"inputs\":[{\"name\":\"protocolFee\",\"type\":\"uint72\",\"indexed\":false,\"internalType\":\"uint72\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnableInvalidOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnableUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]}]",
	Bin: "0x608060405234801561001057600080fd5b5061001961001e565b6100d0565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000900460ff161561006e5760405163f92ee8a960e01b815260040160405180910390fd5b80546001600160401b03908116146100cd5780546001600160401b0319166001600160401b0390811782556040519081527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b50565b611bf2806100df6000396000f3fe608060405234801561001057600080fd5b50600436106101c45760003560e01c8063944de210116100f9578063e21497b911610097578063f2fde38b11610071578063f2fde38b14610608578063f35eecbd1461061b578063f3ccd5001461062e578063fcb1deb11461064157600080fd5b8063e21497b91461056f578063e476d475146105a2578063f263980e146105d557600080fd5b8063b0e21e8a116100d3578063b0e21e8a14610507578063b9923e1c14610536578063d0ebdbe714610549578063dc8d7b2c1461055c57600080fd5b8063944de210146104ae5780639d1d335d146104c1578063a8d8216f146104f457600080fd5b806365c6850d116101665780638da5cb5b116101405780638da5cb5b146104295780638dd9523c146104595780638df66e341461047a5780638f9d6ace146104a457600080fd5b806365c6850d146103fb578063715018a61461040e57806379b1f0e91461041657600080fd5b8063481c6a75116101a2578063481c6a751461037c57806354fd4d50146103ae57806355bd7c10146103b55780635d3acee2146103e857600080fd5b80630a8041f4146101c95780632d4634a4146101de578063415070af14610331575b600080fd5b6101dc6101d736600461169c565b610654565b005b61031b6101ec366004611764565b6040805161010081018252600080825260208201819052918101829052606081018290526080810182905260a0810182905260c0810182905260e0810191909152506001600160401b03908116600081815260016020908152604080832081516080808201845291548088168252600160401b8082048916838701818152600160801b8085048c16868901908152600160c01b958690048d166060978801908152938b5260028a529988902088518089018a529054808e1682529485048d16818b019081529185048d16818a01908152959094048c169386019384528751610100810189529a8b5281518c16988b019890985297518a1695890195909552935188169187019190915293518616908501529051841660a08401529051831660c08301525190911660e082015290565b6040516103289190611786565b60405180910390f35b61036461033f366004611764565b6001600160401b03908116600090815260026020526040902054600160801b90041690565b6040516001600160401b039091168152602001610328565b60005461039690600160601b90046001600160a01b031681565b6040516001600160a01b039091168152602001610328565b6002610364565b6103646103c3366004611764565b6001600160401b03908116600090815260016020526040902054600160c01b90041690565b6101dc6103f6366004611816565b610798565b6101dc610409366004611816565b6107e0565b6101dc61081b565b6101dc610424366004611816565b61082f565b7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c199300546001600160a01b0316610396565b61046c610467366004611849565b61086a565b604051908152602001610328565b60005461049090600160481b900462ffffff1681565b60405162ffffff9091168152602001610328565b61046c620f424081565b6101dc6104bc3660046118dc565b610a52565b6103646104cf366004611764565b6001600160401b03908116600090815260016020526040902054600160401b90041690565b6101dc61050236600461191d565b610a8d565b60005461051c9068ffffffffffffffffff1681565b60405168ffffffffffffffffff9091168152602001610328565b6101dc610544366004611816565b610aa1565b6101dc610557366004611938565b610adc565b6101dc61056a3660046118dc565b610b43565b61036461057d366004611764565b6001600160401b03908116600090815260016020526040902054600160801b90041690565b6103646105b0366004611764565b6001600160401b03908116600090815260026020526040902054600160401b90041690565b6103646105e3366004611764565b6001600160401b03908116600090815260026020526040902054600160c01b90041690565b6101dc610616366004611938565b610b7e565b6101dc610629366004611816565b610bb9565b6101dc61063c366004611816565b610bf4565b6101dc61064f366004611953565b610c2f565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a008054600160401b810460ff1615906001600160401b03166000811580156106995750825b90506000826001600160401b031660011480156106b55750303b155b9050811580156106c3575080155b156106e15760405163f92ee8a960e01b815260040160405180910390fd5b845467ffffffffffffffff19166001178555831561070b57845460ff60401b1916600160401b1785555b6107148d610c40565b61071d8c610c51565b6107268b610cb3565b61072f8a610d0b565b6107398989610d5d565b6107438787610f13565b831561078957845460ff60401b19168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b50505050505050505050505050565b600054600160601b90046001600160a01b031633146107d25760405162461bcd60e51b81526004016107c99061196e565b60405180910390fd5b6107dc8282611097565b5050565b600054600160601b90046001600160a01b031633146108115760405162461bcd60e51b81526004016107c99061196e565b6107dc828261115f565b61082361121a565b61082d6000611275565b565b600054600160601b90046001600160a01b031633146108605760405162461bcd60e51b81526004016107c99061196e565b6107dc82826112e6565b6001600160401b0380851660009081526001602090815260408083208054600160401b810486168552600290935290832092939092918491620f4240916108c291600160c01b8204811691600160801b9004166119bb565b6001600160401b03166108d591906119e6565b8254909150600090620f424090610905906001600160401b03600160c01b8204811691600160801b9004166119bb565b6001600160401b031661091891906119e6565b90506000821161096a5760405162461bcd60e51b815260206004820152601f60248201527f4665654f7261636c6556323a206e6f20657865632066656520706172616d730060448201526064016107c9565b600081116109ba5760405162461bcd60e51b815260206004820152601f60248201527f4665654f7261636c6556323a206e6f20646174612066656520706172616d730060448201526064016107c9565b60006109c7886010611a08565b845490915082906109e9908390600160401b90046001600160401b0316611a25565b6109f39190611a08565b6000548490610a0f908a90600160481b900462ffffff16611a38565b6001600160401b0316610a229190611a08565b600054610a3a919068ffffffffffffffffff16611a25565b610a449190611a25565b9a9950505050505050505050565b600054600160601b90046001600160a01b03163314610a835760405162461bcd60e51b81526004016107c99061196e565b6107dc8282610f13565b610a9561121a565b610a9e81610cb3565b50565b600054600160601b90046001600160a01b03163314610ad25760405162461bcd60e51b81526004016107c99061196e565b6107dc828261137d565b610ae461121a565b6001600160a01b038116610b3a5760405162461bcd60e51b815260206004820152601c60248201527f4665654f7261636c6556323a206e6f207a65726f206d616e616765720000000060448201526064016107c9565b610a9e81610c51565b600054600160601b90046001600160a01b03163314610b745760405162461bcd60e51b81526004016107c99061196e565b6107dc8282610d5d565b610b8661121a565b6001600160a01b038116610bb057604051631e4fbdf760e01b8152600060048201526024016107c9565b610a9e81611275565b600054600160601b90046001600160a01b03163314610bea5760405162461bcd60e51b81526004016107c99061196e565b6107dc828261143d565b600054600160601b90046001600160a01b03163314610c255760405162461bcd60e51b81526004016107c99061196e565b6107dc82826114fd565b610c3761121a565b610a9e81610d0b565b610c486115b8565b610a9e81611601565b600080546bffffffffffffffffffffffff16600160601b6001600160a01b038416908102919091179091556040519081527f60a0f5b9f9e81e98216071b85826681c796256fe3d1354ecb675580fba64fa69906020015b60405180910390a150565b600080546bffffff0000000000000000001916600160481b62ffffff8416908102919091179091556040519081527fb2497d19702e6a33eb2a5487a9ad5977a0284bfe1ccad332e31b8a81a8c5ebaf90602001610ca8565b6000805468ffffffffffffffffff191668ffffffffffffffffff83169081179091556040519081527f92c98513056b0d0d263b0f7153c99e1651791a983e75fc558b424b3aa0b623f790602001610ca8565b60005b81811015610f0e576000838383818110610d7c57610d7c611a58565b905060800201803603810190610d929190611afb565b9050600081604001516001600160401b031611610dc15760405162461bcd60e51b81526004016107c990611b17565b600081602001516001600160401b031611610dee5760405162461bcd60e51b81526004016107c990611b4e565b600081606001516001600160401b031611610e1b5760405162461bcd60e51b81526004016107c990611b85565b80516001600160401b0316600003610e455760405162461bcd60e51b81526004016107c990611b4e565b80516001600160401b039081166000908152600160209081526040918290208451815483870151858801516060808a01519489166001600160801b03199094168417600160401b938a16938402176001600160801b0316600160801b928a169283026001600160c01b031617600160c01b959099169485029890981790945585519182529381019390935292820152918201527f0a05333d850372b39fcfc14c6eed42ebc766b01f516add8c20240696552c44649060800160405180910390a150600101610d60565b505050565b60005b81811015610f0e576000838383818110610f3257610f32611a58565b905060800201803603810190610f489190611afb565b9050600081604001516001600160401b031611610f775760405162461bcd60e51b81526004016107c990611b17565b600081606001516001600160401b031611610fa45760405162461bcd60e51b81526004016107c990611b85565b80516001600160401b0316600003610fce5760405162461bcd60e51b81526004016107c990611b4e565b80516001600160401b039081166000908152600260209081526040918290208451815483870151858801516060808a01519489166001600160801b03199094168417600160401b938a16938402176001600160801b0316600160801b928a169283026001600160c01b031617600160c01b959099169485029890981790945585519182529381019390935292820152918201527fd0b18837916cfd80f0f320486a107794cf007ff14ddafb4483a62c3fb5ec05a39060800160405180910390a150600101610f16565b6000816001600160401b0316116110c05760405162461bcd60e51b81526004016107c990611b17565b816001600160401b03166000036110e95760405162461bcd60e51b81526004016107c990611b4e565b6001600160401b03828116600081815260026020908152604091829020805467ffffffffffffffff60801b1916600160801b95871695860217905581519283528201929092527fd7d8dd5a956a8bd500e02d52d0a9dd8a0e2955ec48771a8c9da485e6706c66fb91015b60405180910390a15050565b6000816001600160401b0316116111885760405162461bcd60e51b81526004016107c990611b85565b816001600160401b03166000036111b15760405162461bcd60e51b81526004016107c990611b4e565b6001600160401b0382811660008181526001602090815260409182902080546001600160c01b0316600160c01b95871695860217905581519283528201929092527f5728e411b1606c4a6fd53e10049ebe42400b3617c09e478e10239a913ab3493b9101611153565b3361124c7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c199300546001600160a01b031690565b6001600160a01b03161461082d5760405163118cdaa760e01b81523360048201526024016107c9565b7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c19930080546001600160a01b031981166001600160a01b03848116918217845560405192169182907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a3505050565b816001600160401b031660000361130f5760405162461bcd60e51b81526004016107c990611b4e565b6001600160401b03828116600081815260026020908152604091829020805467ffffffffffffffff60401b1916600160401b95871695860217905581519283528201929092527f0b1d18b6fe29df95e0bd90abbbf45e65e4a14f3ef2126bfcfb7f2a08112669989101611153565b6000816001600160401b0316116113a65760405162461bcd60e51b81526004016107c990611b17565b816001600160401b03166000036113cf5760405162461bcd60e51b81526004016107c990611b4e565b6001600160401b03828116600081815260016020908152604091829020805467ffffffffffffffff60801b1916600160801b95871695860217905581519283528201929092527fe0e5abb8929e27a69d77f47a4e3f9575411a5be1fa596e5b55078d7850f358db9101611153565b806001600160401b03166000036114665760405162461bcd60e51b81526004016107c990611b4e565b816001600160401b031660000361148f5760405162461bcd60e51b81526004016107c990611b4e565b6001600160401b03828116600081815260016020908152604091829020805467ffffffffffffffff60401b1916600160401b95871695860217905581519283528201929092527f62f13d8522c451be7798aa5f1e258a7780ef5296115067b8e632436d4548a11d9101611153565b6000816001600160401b0316116115265760405162461bcd60e51b81526004016107c990611b85565b816001600160401b031660000361154f5760405162461bcd60e51b81526004016107c990611b4e565b6001600160401b0382811660008181526002602090815260409182902080546001600160c01b0316600160c01b95871695860217905581519283528201929092527f1e87d1bd7c559f1a6fcee0e50d1f7d7df76cee0bafc691051b174dc2a83139219101611153565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a0054600160401b900460ff1661082d57604051631afcd79f60e31b815260040160405180910390fd5b610b866115b8565b80356001600160a01b038116811461162057600080fd5b919050565b803562ffffff8116811461162057600080fd5b803568ffffffffffffffffff8116811461162057600080fd5b60008083601f84011261166357600080fd5b5081356001600160401b0381111561167a57600080fd5b6020830191508360208260071b850101111561169557600080fd5b9250929050565b60008060008060008060008060c0898b0312156116b857600080fd5b6116c189611609565b97506116cf60208a01611609565b96506116dd60408a01611625565b95506116eb60608a01611638565b945060808901356001600160401b038082111561170757600080fd5b6117138c838d01611651565b909650945060a08b013591508082111561172c57600080fd5b506117398b828c01611651565b999c989b5096995094979396929594505050565b80356001600160401b038116811461162057600080fd5b60006020828403121561177657600080fd5b61177f8261174d565b9392505050565b6000610100820190506001600160401b038084511683528060208501511660208401528060408501511660408401528060608501511660608401528060808501511660808401528060a08501511660a08401525060c08301516117f460c08401826001600160401b03169052565b5060e083015161180f60e08401826001600160401b03169052565b5092915050565b6000806040838503121561182957600080fd5b6118328361174d565b91506118406020840161174d565b90509250929050565b6000806000806060858703121561185f57600080fd5b6118688561174d565b935060208501356001600160401b038082111561188457600080fd5b818701915087601f83011261189857600080fd5b8135818111156118a757600080fd5b8860208285010111156118b957600080fd5b6020830195508094505050506118d16040860161174d565b905092959194509250565b600080602083850312156118ef57600080fd5b82356001600160401b0381111561190557600080fd5b61191185828601611651565b90969095509350505050565b60006020828403121561192f57600080fd5b61177f82611625565b60006020828403121561194a57600080fd5b61177f82611609565b60006020828403121561196557600080fd5b61177f82611638565b60208082526018908201527f4665654f7261636c6556323a206e6f74206d616e616765720000000000000000604082015260600190565b634e487b7160e01b600052601160045260246000fd5b6001600160401b038181168382160280821691908281146119de576119de6119a5565b505092915050565b600082611a0357634e487b7160e01b600052601260045260246000fd5b500490565b8082028115828204841417611a1f57611a1f6119a5565b92915050565b80820180821115611a1f57611a1f6119a5565b6001600160401b0381811683821601908082111561180f5761180f6119a5565b634e487b7160e01b600052603260045260246000fd5b600060808284031215611a8057600080fd5b604051608081018181106001600160401b0382111715611ab057634e487b7160e01b600052604160045260246000fd5b604052905080611abf8361174d565b8152611acd6020840161174d565b6020820152611ade6040840161174d565b6040820152611aef6060840161174d565b60608201525092915050565b600060808284031215611b0d57600080fd5b61177f8383611a6e565b6020808252601e908201527f4665654f7261636c6556323a206e6f207a65726f206761732070726963650000604082015260600190565b6020808252601d908201527f4665654f7261636c6556323a206e6f207a65726f20636861696e206964000000604082015260600190565b60208082526019908201527f4665654f7261636c6556323a206e6f207a65726f20726174650000000000000060408201526060019056fea2646970667358221220be7739c13e913ff724cc43977f5aedf48b723bf4869c8136354ada3b3ab37f2764736f6c63430008180033",
}

// FeeOracleV2ABI is the input ABI used to generate the binding from.
// Deprecated: Use FeeOracleV2MetaData.ABI instead.
var FeeOracleV2ABI = FeeOracleV2MetaData.ABI

// FeeOracleV2Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use FeeOracleV2MetaData.Bin instead.
var FeeOracleV2Bin = FeeOracleV2MetaData.Bin

// DeployFeeOracleV2 deploys a new Ethereum contract, binding an instance of FeeOracleV2 to it.
func DeployFeeOracleV2(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *FeeOracleV2, error) {
	parsed, err := FeeOracleV2MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(FeeOracleV2Bin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &FeeOracleV2{FeeOracleV2Caller: FeeOracleV2Caller{contract: contract}, FeeOracleV2Transactor: FeeOracleV2Transactor{contract: contract}, FeeOracleV2Filterer: FeeOracleV2Filterer{contract: contract}}, nil
}

// FeeOracleV2 is an auto generated Go binding around an Ethereum contract.
type FeeOracleV2 struct {
	FeeOracleV2Caller     // Read-only binding to the contract
	FeeOracleV2Transactor // Write-only binding to the contract
	FeeOracleV2Filterer   // Log filterer for contract events
}

// FeeOracleV2Caller is an auto generated read-only Go binding around an Ethereum contract.
type FeeOracleV2Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FeeOracleV2Transactor is an auto generated write-only Go binding around an Ethereum contract.
type FeeOracleV2Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FeeOracleV2Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type FeeOracleV2Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// FeeOracleV2Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type FeeOracleV2Session struct {
	Contract     *FeeOracleV2      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// FeeOracleV2CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type FeeOracleV2CallerSession struct {
	Contract *FeeOracleV2Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// FeeOracleV2TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type FeeOracleV2TransactorSession struct {
	Contract     *FeeOracleV2Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// FeeOracleV2Raw is an auto generated low-level Go binding around an Ethereum contract.
type FeeOracleV2Raw struct {
	Contract *FeeOracleV2 // Generic contract binding to access the raw methods on
}

// FeeOracleV2CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type FeeOracleV2CallerRaw struct {
	Contract *FeeOracleV2Caller // Generic read-only contract binding to access the raw methods on
}

// FeeOracleV2TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type FeeOracleV2TransactorRaw struct {
	Contract *FeeOracleV2Transactor // Generic write-only contract binding to access the raw methods on
}

// NewFeeOracleV2 creates a new instance of FeeOracleV2, bound to a specific deployed contract.
func NewFeeOracleV2(address common.Address, backend bind.ContractBackend) (*FeeOracleV2, error) {
	contract, err := bindFeeOracleV2(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &FeeOracleV2{FeeOracleV2Caller: FeeOracleV2Caller{contract: contract}, FeeOracleV2Transactor: FeeOracleV2Transactor{contract: contract}, FeeOracleV2Filterer: FeeOracleV2Filterer{contract: contract}}, nil
}

// NewFeeOracleV2Caller creates a new read-only instance of FeeOracleV2, bound to a specific deployed contract.
func NewFeeOracleV2Caller(address common.Address, caller bind.ContractCaller) (*FeeOracleV2Caller, error) {
	contract, err := bindFeeOracleV2(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &FeeOracleV2Caller{contract: contract}, nil
}

// NewFeeOracleV2Transactor creates a new write-only instance of FeeOracleV2, bound to a specific deployed contract.
func NewFeeOracleV2Transactor(address common.Address, transactor bind.ContractTransactor) (*FeeOracleV2Transactor, error) {
	contract, err := bindFeeOracleV2(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &FeeOracleV2Transactor{contract: contract}, nil
}

// NewFeeOracleV2Filterer creates a new log filterer instance of FeeOracleV2, bound to a specific deployed contract.
func NewFeeOracleV2Filterer(address common.Address, filterer bind.ContractFilterer) (*FeeOracleV2Filterer, error) {
	contract, err := bindFeeOracleV2(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &FeeOracleV2Filterer{contract: contract}, nil
}

// bindFeeOracleV2 binds a generic wrapper to an already deployed contract.
func bindFeeOracleV2(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := FeeOracleV2MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FeeOracleV2 *FeeOracleV2Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FeeOracleV2.Contract.FeeOracleV2Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FeeOracleV2 *FeeOracleV2Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FeeOracleV2.Contract.FeeOracleV2Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FeeOracleV2 *FeeOracleV2Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FeeOracleV2.Contract.FeeOracleV2Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_FeeOracleV2 *FeeOracleV2CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _FeeOracleV2.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_FeeOracleV2 *FeeOracleV2TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FeeOracleV2.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_FeeOracleV2 *FeeOracleV2TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _FeeOracleV2.Contract.contract.Transact(opts, method, params...)
}

// CONVERSIONRATEDENOM is a free data retrieval call binding the contract method 0x8f9d6ace.
//
// Solidity: function CONVERSION_RATE_DENOM() view returns(uint256)
func (_FeeOracleV2 *FeeOracleV2Caller) CONVERSIONRATEDENOM(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FeeOracleV2.contract.Call(opts, &out, "CONVERSION_RATE_DENOM")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CONVERSIONRATEDENOM is a free data retrieval call binding the contract method 0x8f9d6ace.
//
// Solidity: function CONVERSION_RATE_DENOM() view returns(uint256)
func (_FeeOracleV2 *FeeOracleV2Session) CONVERSIONRATEDENOM() (*big.Int, error) {
	return _FeeOracleV2.Contract.CONVERSIONRATEDENOM(&_FeeOracleV2.CallOpts)
}

// CONVERSIONRATEDENOM is a free data retrieval call binding the contract method 0x8f9d6ace.
//
// Solidity: function CONVERSION_RATE_DENOM() view returns(uint256)
func (_FeeOracleV2 *FeeOracleV2CallerSession) CONVERSIONRATEDENOM() (*big.Int, error) {
	return _FeeOracleV2.Contract.CONVERSIONRATEDENOM(&_FeeOracleV2.CallOpts)
}

// BaseGasLimit is a free data retrieval call binding the contract method 0x8df66e34.
//
// Solidity: function baseGasLimit() view returns(uint24)
func (_FeeOracleV2 *FeeOracleV2Caller) BaseGasLimit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FeeOracleV2.contract.Call(opts, &out, "baseGasLimit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BaseGasLimit is a free data retrieval call binding the contract method 0x8df66e34.
//
// Solidity: function baseGasLimit() view returns(uint24)
func (_FeeOracleV2 *FeeOracleV2Session) BaseGasLimit() (*big.Int, error) {
	return _FeeOracleV2.Contract.BaseGasLimit(&_FeeOracleV2.CallOpts)
}

// BaseGasLimit is a free data retrieval call binding the contract method 0x8df66e34.
//
// Solidity: function baseGasLimit() view returns(uint24)
func (_FeeOracleV2 *FeeOracleV2CallerSession) BaseGasLimit() (*big.Int, error) {
	return _FeeOracleV2.Contract.BaseGasLimit(&_FeeOracleV2.CallOpts)
}

// DataGasPrice is a free data retrieval call binding the contract method 0x415070af.
//
// Solidity: function dataGasPrice(uint64 chainId) view returns(uint64)
func (_FeeOracleV2 *FeeOracleV2Caller) DataGasPrice(opts *bind.CallOpts, chainId uint64) (uint64, error) {
	var out []interface{}
	err := _FeeOracleV2.contract.Call(opts, &out, "dataGasPrice", chainId)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// DataGasPrice is a free data retrieval call binding the contract method 0x415070af.
//
// Solidity: function dataGasPrice(uint64 chainId) view returns(uint64)
func (_FeeOracleV2 *FeeOracleV2Session) DataGasPrice(chainId uint64) (uint64, error) {
	return _FeeOracleV2.Contract.DataGasPrice(&_FeeOracleV2.CallOpts, chainId)
}

// DataGasPrice is a free data retrieval call binding the contract method 0x415070af.
//
// Solidity: function dataGasPrice(uint64 chainId) view returns(uint64)
func (_FeeOracleV2 *FeeOracleV2CallerSession) DataGasPrice(chainId uint64) (uint64, error) {
	return _FeeOracleV2.Contract.DataGasPrice(&_FeeOracleV2.CallOpts, chainId)
}

// DataSizeBuffer is a free data retrieval call binding the contract method 0xe476d475.
//
// Solidity: function dataSizeBuffer(uint64 chainId) view returns(uint64)
func (_FeeOracleV2 *FeeOracleV2Caller) DataSizeBuffer(opts *bind.CallOpts, chainId uint64) (uint64, error) {
	var out []interface{}
	err := _FeeOracleV2.contract.Call(opts, &out, "dataSizeBuffer", chainId)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// DataSizeBuffer is a free data retrieval call binding the contract method 0xe476d475.
//
// Solidity: function dataSizeBuffer(uint64 chainId) view returns(uint64)
func (_FeeOracleV2 *FeeOracleV2Session) DataSizeBuffer(chainId uint64) (uint64, error) {
	return _FeeOracleV2.Contract.DataSizeBuffer(&_FeeOracleV2.CallOpts, chainId)
}

// DataSizeBuffer is a free data retrieval call binding the contract method 0xe476d475.
//
// Solidity: function dataSizeBuffer(uint64 chainId) view returns(uint64)
func (_FeeOracleV2 *FeeOracleV2CallerSession) DataSizeBuffer(chainId uint64) (uint64, error) {
	return _FeeOracleV2.Contract.DataSizeBuffer(&_FeeOracleV2.CallOpts, chainId)
}

// DataToNativeRate is a free data retrieval call binding the contract method 0xf263980e.
//
// Solidity: function dataToNativeRate(uint64 chainId) view returns(uint64)
func (_FeeOracleV2 *FeeOracleV2Caller) DataToNativeRate(opts *bind.CallOpts, chainId uint64) (uint64, error) {
	var out []interface{}
	err := _FeeOracleV2.contract.Call(opts, &out, "dataToNativeRate", chainId)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// DataToNativeRate is a free data retrieval call binding the contract method 0xf263980e.
//
// Solidity: function dataToNativeRate(uint64 chainId) view returns(uint64)
func (_FeeOracleV2 *FeeOracleV2Session) DataToNativeRate(chainId uint64) (uint64, error) {
	return _FeeOracleV2.Contract.DataToNativeRate(&_FeeOracleV2.CallOpts, chainId)
}

// DataToNativeRate is a free data retrieval call binding the contract method 0xf263980e.
//
// Solidity: function dataToNativeRate(uint64 chainId) view returns(uint64)
func (_FeeOracleV2 *FeeOracleV2CallerSession) DataToNativeRate(chainId uint64) (uint64, error) {
	return _FeeOracleV2.Contract.DataToNativeRate(&_FeeOracleV2.CallOpts, chainId)
}

// ExecGasPrice is a free data retrieval call binding the contract method 0xe21497b9.
//
// Solidity: function execGasPrice(uint64 chainId) view returns(uint64)
func (_FeeOracleV2 *FeeOracleV2Caller) ExecGasPrice(opts *bind.CallOpts, chainId uint64) (uint64, error) {
	var out []interface{}
	err := _FeeOracleV2.contract.Call(opts, &out, "execGasPrice", chainId)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// ExecGasPrice is a free data retrieval call binding the contract method 0xe21497b9.
//
// Solidity: function execGasPrice(uint64 chainId) view returns(uint64)
func (_FeeOracleV2 *FeeOracleV2Session) ExecGasPrice(chainId uint64) (uint64, error) {
	return _FeeOracleV2.Contract.ExecGasPrice(&_FeeOracleV2.CallOpts, chainId)
}

// ExecGasPrice is a free data retrieval call binding the contract method 0xe21497b9.
//
// Solidity: function execGasPrice(uint64 chainId) view returns(uint64)
func (_FeeOracleV2 *FeeOracleV2CallerSession) ExecGasPrice(chainId uint64) (uint64, error) {
	return _FeeOracleV2.Contract.ExecGasPrice(&_FeeOracleV2.CallOpts, chainId)
}

// ExecPostsTo is a free data retrieval call binding the contract method 0x9d1d335d.
//
// Solidity: function execPostsTo(uint64 chainId) view returns(uint64)
func (_FeeOracleV2 *FeeOracleV2Caller) ExecPostsTo(opts *bind.CallOpts, chainId uint64) (uint64, error) {
	var out []interface{}
	err := _FeeOracleV2.contract.Call(opts, &out, "execPostsTo", chainId)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// ExecPostsTo is a free data retrieval call binding the contract method 0x9d1d335d.
//
// Solidity: function execPostsTo(uint64 chainId) view returns(uint64)
func (_FeeOracleV2 *FeeOracleV2Session) ExecPostsTo(chainId uint64) (uint64, error) {
	return _FeeOracleV2.Contract.ExecPostsTo(&_FeeOracleV2.CallOpts, chainId)
}

// ExecPostsTo is a free data retrieval call binding the contract method 0x9d1d335d.
//
// Solidity: function execPostsTo(uint64 chainId) view returns(uint64)
func (_FeeOracleV2 *FeeOracleV2CallerSession) ExecPostsTo(chainId uint64) (uint64, error) {
	return _FeeOracleV2.Contract.ExecPostsTo(&_FeeOracleV2.CallOpts, chainId)
}

// ExecToNativeRate is a free data retrieval call binding the contract method 0x55bd7c10.
//
// Solidity: function execToNativeRate(uint64 chainId) view returns(uint64)
func (_FeeOracleV2 *FeeOracleV2Caller) ExecToNativeRate(opts *bind.CallOpts, chainId uint64) (uint64, error) {
	var out []interface{}
	err := _FeeOracleV2.contract.Call(opts, &out, "execToNativeRate", chainId)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// ExecToNativeRate is a free data retrieval call binding the contract method 0x55bd7c10.
//
// Solidity: function execToNativeRate(uint64 chainId) view returns(uint64)
func (_FeeOracleV2 *FeeOracleV2Session) ExecToNativeRate(chainId uint64) (uint64, error) {
	return _FeeOracleV2.Contract.ExecToNativeRate(&_FeeOracleV2.CallOpts, chainId)
}

// ExecToNativeRate is a free data retrieval call binding the contract method 0x55bd7c10.
//
// Solidity: function execToNativeRate(uint64 chainId) view returns(uint64)
func (_FeeOracleV2 *FeeOracleV2CallerSession) ExecToNativeRate(chainId uint64) (uint64, error) {
	return _FeeOracleV2.Contract.ExecToNativeRate(&_FeeOracleV2.CallOpts, chainId)
}

// FeeFor is a free data retrieval call binding the contract method 0x8dd9523c.
//
// Solidity: function feeFor(uint64 destChainId, bytes data, uint64 gasLimit) view returns(uint256)
func (_FeeOracleV2 *FeeOracleV2Caller) FeeFor(opts *bind.CallOpts, destChainId uint64, data []byte, gasLimit uint64) (*big.Int, error) {
	var out []interface{}
	err := _FeeOracleV2.contract.Call(opts, &out, "feeFor", destChainId, data, gasLimit)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FeeFor is a free data retrieval call binding the contract method 0x8dd9523c.
//
// Solidity: function feeFor(uint64 destChainId, bytes data, uint64 gasLimit) view returns(uint256)
func (_FeeOracleV2 *FeeOracleV2Session) FeeFor(destChainId uint64, data []byte, gasLimit uint64) (*big.Int, error) {
	return _FeeOracleV2.Contract.FeeFor(&_FeeOracleV2.CallOpts, destChainId, data, gasLimit)
}

// FeeFor is a free data retrieval call binding the contract method 0x8dd9523c.
//
// Solidity: function feeFor(uint64 destChainId, bytes data, uint64 gasLimit) view returns(uint256)
func (_FeeOracleV2 *FeeOracleV2CallerSession) FeeFor(destChainId uint64, data []byte, gasLimit uint64) (*big.Int, error) {
	return _FeeOracleV2.Contract.FeeFor(&_FeeOracleV2.CallOpts, destChainId, data, gasLimit)
}

// FeeParams is a free data retrieval call binding the contract method 0x2d4634a4.
//
// Solidity: function feeParams(uint64 chainId) view returns((uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64))
func (_FeeOracleV2 *FeeOracleV2Caller) FeeParams(opts *bind.CallOpts, chainId uint64) (IFeeOracleV2FeeParams, error) {
	var out []interface{}
	err := _FeeOracleV2.contract.Call(opts, &out, "feeParams", chainId)

	if err != nil {
		return *new(IFeeOracleV2FeeParams), err
	}

	out0 := *abi.ConvertType(out[0], new(IFeeOracleV2FeeParams)).(*IFeeOracleV2FeeParams)

	return out0, err

}

// FeeParams is a free data retrieval call binding the contract method 0x2d4634a4.
//
// Solidity: function feeParams(uint64 chainId) view returns((uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64))
func (_FeeOracleV2 *FeeOracleV2Session) FeeParams(chainId uint64) (IFeeOracleV2FeeParams, error) {
	return _FeeOracleV2.Contract.FeeParams(&_FeeOracleV2.CallOpts, chainId)
}

// FeeParams is a free data retrieval call binding the contract method 0x2d4634a4.
//
// Solidity: function feeParams(uint64 chainId) view returns((uint64,uint64,uint64,uint64,uint64,uint64,uint64,uint64))
func (_FeeOracleV2 *FeeOracleV2CallerSession) FeeParams(chainId uint64) (IFeeOracleV2FeeParams, error) {
	return _FeeOracleV2.Contract.FeeParams(&_FeeOracleV2.CallOpts, chainId)
}

// Manager is a free data retrieval call binding the contract method 0x481c6a75.
//
// Solidity: function manager() view returns(address)
func (_FeeOracleV2 *FeeOracleV2Caller) Manager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FeeOracleV2.contract.Call(opts, &out, "manager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Manager is a free data retrieval call binding the contract method 0x481c6a75.
//
// Solidity: function manager() view returns(address)
func (_FeeOracleV2 *FeeOracleV2Session) Manager() (common.Address, error) {
	return _FeeOracleV2.Contract.Manager(&_FeeOracleV2.CallOpts)
}

// Manager is a free data retrieval call binding the contract method 0x481c6a75.
//
// Solidity: function manager() view returns(address)
func (_FeeOracleV2 *FeeOracleV2CallerSession) Manager() (common.Address, error) {
	return _FeeOracleV2.Contract.Manager(&_FeeOracleV2.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_FeeOracleV2 *FeeOracleV2Caller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _FeeOracleV2.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_FeeOracleV2 *FeeOracleV2Session) Owner() (common.Address, error) {
	return _FeeOracleV2.Contract.Owner(&_FeeOracleV2.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_FeeOracleV2 *FeeOracleV2CallerSession) Owner() (common.Address, error) {
	return _FeeOracleV2.Contract.Owner(&_FeeOracleV2.CallOpts)
}

// ProtocolFee is a free data retrieval call binding the contract method 0xb0e21e8a.
//
// Solidity: function protocolFee() view returns(uint72)
func (_FeeOracleV2 *FeeOracleV2Caller) ProtocolFee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _FeeOracleV2.contract.Call(opts, &out, "protocolFee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ProtocolFee is a free data retrieval call binding the contract method 0xb0e21e8a.
//
// Solidity: function protocolFee() view returns(uint72)
func (_FeeOracleV2 *FeeOracleV2Session) ProtocolFee() (*big.Int, error) {
	return _FeeOracleV2.Contract.ProtocolFee(&_FeeOracleV2.CallOpts)
}

// ProtocolFee is a free data retrieval call binding the contract method 0xb0e21e8a.
//
// Solidity: function protocolFee() view returns(uint72)
func (_FeeOracleV2 *FeeOracleV2CallerSession) ProtocolFee() (*big.Int, error) {
	return _FeeOracleV2.Contract.ProtocolFee(&_FeeOracleV2.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(uint64)
func (_FeeOracleV2 *FeeOracleV2Caller) Version(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _FeeOracleV2.contract.Call(opts, &out, "version")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(uint64)
func (_FeeOracleV2 *FeeOracleV2Session) Version() (uint64, error) {
	return _FeeOracleV2.Contract.Version(&_FeeOracleV2.CallOpts)
}

// Version is a free data retrieval call binding the contract method 0x54fd4d50.
//
// Solidity: function version() pure returns(uint64)
func (_FeeOracleV2 *FeeOracleV2CallerSession) Version() (uint64, error) {
	return _FeeOracleV2.Contract.Version(&_FeeOracleV2.CallOpts)
}

// BulkSetDataFeeParams is a paid mutator transaction binding the contract method 0x944de210.
//
// Solidity: function bulkSetDataFeeParams((uint64,uint64,uint64,uint64)[] params) returns()
func (_FeeOracleV2 *FeeOracleV2Transactor) BulkSetDataFeeParams(opts *bind.TransactOpts, params []IFeeOracleV2DataFeeParams) (*types.Transaction, error) {
	return _FeeOracleV2.contract.Transact(opts, "bulkSetDataFeeParams", params)
}

// BulkSetDataFeeParams is a paid mutator transaction binding the contract method 0x944de210.
//
// Solidity: function bulkSetDataFeeParams((uint64,uint64,uint64,uint64)[] params) returns()
func (_FeeOracleV2 *FeeOracleV2Session) BulkSetDataFeeParams(params []IFeeOracleV2DataFeeParams) (*types.Transaction, error) {
	return _FeeOracleV2.Contract.BulkSetDataFeeParams(&_FeeOracleV2.TransactOpts, params)
}

// BulkSetDataFeeParams is a paid mutator transaction binding the contract method 0x944de210.
//
// Solidity: function bulkSetDataFeeParams((uint64,uint64,uint64,uint64)[] params) returns()
func (_FeeOracleV2 *FeeOracleV2TransactorSession) BulkSetDataFeeParams(params []IFeeOracleV2DataFeeParams) (*types.Transaction, error) {
	return _FeeOracleV2.Contract.BulkSetDataFeeParams(&_FeeOracleV2.TransactOpts, params)
}

// BulkSetExecFeeParams is a paid mutator transaction binding the contract method 0xdc8d7b2c.
//
// Solidity: function bulkSetExecFeeParams((uint64,uint64,uint64,uint64)[] params) returns()
func (_FeeOracleV2 *FeeOracleV2Transactor) BulkSetExecFeeParams(opts *bind.TransactOpts, params []IFeeOracleV2ExecFeeParams) (*types.Transaction, error) {
	return _FeeOracleV2.contract.Transact(opts, "bulkSetExecFeeParams", params)
}

// BulkSetExecFeeParams is a paid mutator transaction binding the contract method 0xdc8d7b2c.
//
// Solidity: function bulkSetExecFeeParams((uint64,uint64,uint64,uint64)[] params) returns()
func (_FeeOracleV2 *FeeOracleV2Session) BulkSetExecFeeParams(params []IFeeOracleV2ExecFeeParams) (*types.Transaction, error) {
	return _FeeOracleV2.Contract.BulkSetExecFeeParams(&_FeeOracleV2.TransactOpts, params)
}

// BulkSetExecFeeParams is a paid mutator transaction binding the contract method 0xdc8d7b2c.
//
// Solidity: function bulkSetExecFeeParams((uint64,uint64,uint64,uint64)[] params) returns()
func (_FeeOracleV2 *FeeOracleV2TransactorSession) BulkSetExecFeeParams(params []IFeeOracleV2ExecFeeParams) (*types.Transaction, error) {
	return _FeeOracleV2.Contract.BulkSetExecFeeParams(&_FeeOracleV2.TransactOpts, params)
}

// Initialize is a paid mutator transaction binding the contract method 0x0a8041f4.
//
// Solidity: function initialize(address owner_, address manager_, uint24 baseGasLimit_, uint72 protocolFee_, (uint64,uint64,uint64,uint64)[] execParams, (uint64,uint64,uint64,uint64)[] dataParams) returns()
func (_FeeOracleV2 *FeeOracleV2Transactor) Initialize(opts *bind.TransactOpts, owner_ common.Address, manager_ common.Address, baseGasLimit_ *big.Int, protocolFee_ *big.Int, execParams []IFeeOracleV2ExecFeeParams, dataParams []IFeeOracleV2DataFeeParams) (*types.Transaction, error) {
	return _FeeOracleV2.contract.Transact(opts, "initialize", owner_, manager_, baseGasLimit_, protocolFee_, execParams, dataParams)
}

// Initialize is a paid mutator transaction binding the contract method 0x0a8041f4.
//
// Solidity: function initialize(address owner_, address manager_, uint24 baseGasLimit_, uint72 protocolFee_, (uint64,uint64,uint64,uint64)[] execParams, (uint64,uint64,uint64,uint64)[] dataParams) returns()
func (_FeeOracleV2 *FeeOracleV2Session) Initialize(owner_ common.Address, manager_ common.Address, baseGasLimit_ *big.Int, protocolFee_ *big.Int, execParams []IFeeOracleV2ExecFeeParams, dataParams []IFeeOracleV2DataFeeParams) (*types.Transaction, error) {
	return _FeeOracleV2.Contract.Initialize(&_FeeOracleV2.TransactOpts, owner_, manager_, baseGasLimit_, protocolFee_, execParams, dataParams)
}

// Initialize is a paid mutator transaction binding the contract method 0x0a8041f4.
//
// Solidity: function initialize(address owner_, address manager_, uint24 baseGasLimit_, uint72 protocolFee_, (uint64,uint64,uint64,uint64)[] execParams, (uint64,uint64,uint64,uint64)[] dataParams) returns()
func (_FeeOracleV2 *FeeOracleV2TransactorSession) Initialize(owner_ common.Address, manager_ common.Address, baseGasLimit_ *big.Int, protocolFee_ *big.Int, execParams []IFeeOracleV2ExecFeeParams, dataParams []IFeeOracleV2DataFeeParams) (*types.Transaction, error) {
	return _FeeOracleV2.Contract.Initialize(&_FeeOracleV2.TransactOpts, owner_, manager_, baseGasLimit_, protocolFee_, execParams, dataParams)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_FeeOracleV2 *FeeOracleV2Transactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _FeeOracleV2.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_FeeOracleV2 *FeeOracleV2Session) RenounceOwnership() (*types.Transaction, error) {
	return _FeeOracleV2.Contract.RenounceOwnership(&_FeeOracleV2.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_FeeOracleV2 *FeeOracleV2TransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _FeeOracleV2.Contract.RenounceOwnership(&_FeeOracleV2.TransactOpts)
}

// SetBaseGasLimit is a paid mutator transaction binding the contract method 0xa8d8216f.
//
// Solidity: function setBaseGasLimit(uint24 gasLimit) returns()
func (_FeeOracleV2 *FeeOracleV2Transactor) SetBaseGasLimit(opts *bind.TransactOpts, gasLimit *big.Int) (*types.Transaction, error) {
	return _FeeOracleV2.contract.Transact(opts, "setBaseGasLimit", gasLimit)
}

// SetBaseGasLimit is a paid mutator transaction binding the contract method 0xa8d8216f.
//
// Solidity: function setBaseGasLimit(uint24 gasLimit) returns()
func (_FeeOracleV2 *FeeOracleV2Session) SetBaseGasLimit(gasLimit *big.Int) (*types.Transaction, error) {
	return _FeeOracleV2.Contract.SetBaseGasLimit(&_FeeOracleV2.TransactOpts, gasLimit)
}

// SetBaseGasLimit is a paid mutator transaction binding the contract method 0xa8d8216f.
//
// Solidity: function setBaseGasLimit(uint24 gasLimit) returns()
func (_FeeOracleV2 *FeeOracleV2TransactorSession) SetBaseGasLimit(gasLimit *big.Int) (*types.Transaction, error) {
	return _FeeOracleV2.Contract.SetBaseGasLimit(&_FeeOracleV2.TransactOpts, gasLimit)
}

// SetDataGasPrice is a paid mutator transaction binding the contract method 0x5d3acee2.
//
// Solidity: function setDataGasPrice(uint64 chainId, uint64 gasPrice) returns()
func (_FeeOracleV2 *FeeOracleV2Transactor) SetDataGasPrice(opts *bind.TransactOpts, chainId uint64, gasPrice uint64) (*types.Transaction, error) {
	return _FeeOracleV2.contract.Transact(opts, "setDataGasPrice", chainId, gasPrice)
}

// SetDataGasPrice is a paid mutator transaction binding the contract method 0x5d3acee2.
//
// Solidity: function setDataGasPrice(uint64 chainId, uint64 gasPrice) returns()
func (_FeeOracleV2 *FeeOracleV2Session) SetDataGasPrice(chainId uint64, gasPrice uint64) (*types.Transaction, error) {
	return _FeeOracleV2.Contract.SetDataGasPrice(&_FeeOracleV2.TransactOpts, chainId, gasPrice)
}

// SetDataGasPrice is a paid mutator transaction binding the contract method 0x5d3acee2.
//
// Solidity: function setDataGasPrice(uint64 chainId, uint64 gasPrice) returns()
func (_FeeOracleV2 *FeeOracleV2TransactorSession) SetDataGasPrice(chainId uint64, gasPrice uint64) (*types.Transaction, error) {
	return _FeeOracleV2.Contract.SetDataGasPrice(&_FeeOracleV2.TransactOpts, chainId, gasPrice)
}

// SetDataSizeBuffer is a paid mutator transaction binding the contract method 0x79b1f0e9.
//
// Solidity: function setDataSizeBuffer(uint64 chainId, uint64 sizeBuffer) returns()
func (_FeeOracleV2 *FeeOracleV2Transactor) SetDataSizeBuffer(opts *bind.TransactOpts, chainId uint64, sizeBuffer uint64) (*types.Transaction, error) {
	return _FeeOracleV2.contract.Transact(opts, "setDataSizeBuffer", chainId, sizeBuffer)
}

// SetDataSizeBuffer is a paid mutator transaction binding the contract method 0x79b1f0e9.
//
// Solidity: function setDataSizeBuffer(uint64 chainId, uint64 sizeBuffer) returns()
func (_FeeOracleV2 *FeeOracleV2Session) SetDataSizeBuffer(chainId uint64, sizeBuffer uint64) (*types.Transaction, error) {
	return _FeeOracleV2.Contract.SetDataSizeBuffer(&_FeeOracleV2.TransactOpts, chainId, sizeBuffer)
}

// SetDataSizeBuffer is a paid mutator transaction binding the contract method 0x79b1f0e9.
//
// Solidity: function setDataSizeBuffer(uint64 chainId, uint64 sizeBuffer) returns()
func (_FeeOracleV2 *FeeOracleV2TransactorSession) SetDataSizeBuffer(chainId uint64, sizeBuffer uint64) (*types.Transaction, error) {
	return _FeeOracleV2.Contract.SetDataSizeBuffer(&_FeeOracleV2.TransactOpts, chainId, sizeBuffer)
}

// SetDataToNativeRate is a paid mutator transaction binding the contract method 0xf3ccd500.
//
// Solidity: function setDataToNativeRate(uint64 chainId, uint64 rate) returns()
func (_FeeOracleV2 *FeeOracleV2Transactor) SetDataToNativeRate(opts *bind.TransactOpts, chainId uint64, rate uint64) (*types.Transaction, error) {
	return _FeeOracleV2.contract.Transact(opts, "setDataToNativeRate", chainId, rate)
}

// SetDataToNativeRate is a paid mutator transaction binding the contract method 0xf3ccd500.
//
// Solidity: function setDataToNativeRate(uint64 chainId, uint64 rate) returns()
func (_FeeOracleV2 *FeeOracleV2Session) SetDataToNativeRate(chainId uint64, rate uint64) (*types.Transaction, error) {
	return _FeeOracleV2.Contract.SetDataToNativeRate(&_FeeOracleV2.TransactOpts, chainId, rate)
}

// SetDataToNativeRate is a paid mutator transaction binding the contract method 0xf3ccd500.
//
// Solidity: function setDataToNativeRate(uint64 chainId, uint64 rate) returns()
func (_FeeOracleV2 *FeeOracleV2TransactorSession) SetDataToNativeRate(chainId uint64, rate uint64) (*types.Transaction, error) {
	return _FeeOracleV2.Contract.SetDataToNativeRate(&_FeeOracleV2.TransactOpts, chainId, rate)
}

// SetExecGasPrice is a paid mutator transaction binding the contract method 0xb9923e1c.
//
// Solidity: function setExecGasPrice(uint64 chainId, uint64 gasPrice) returns()
func (_FeeOracleV2 *FeeOracleV2Transactor) SetExecGasPrice(opts *bind.TransactOpts, chainId uint64, gasPrice uint64) (*types.Transaction, error) {
	return _FeeOracleV2.contract.Transact(opts, "setExecGasPrice", chainId, gasPrice)
}

// SetExecGasPrice is a paid mutator transaction binding the contract method 0xb9923e1c.
//
// Solidity: function setExecGasPrice(uint64 chainId, uint64 gasPrice) returns()
func (_FeeOracleV2 *FeeOracleV2Session) SetExecGasPrice(chainId uint64, gasPrice uint64) (*types.Transaction, error) {
	return _FeeOracleV2.Contract.SetExecGasPrice(&_FeeOracleV2.TransactOpts, chainId, gasPrice)
}

// SetExecGasPrice is a paid mutator transaction binding the contract method 0xb9923e1c.
//
// Solidity: function setExecGasPrice(uint64 chainId, uint64 gasPrice) returns()
func (_FeeOracleV2 *FeeOracleV2TransactorSession) SetExecGasPrice(chainId uint64, gasPrice uint64) (*types.Transaction, error) {
	return _FeeOracleV2.Contract.SetExecGasPrice(&_FeeOracleV2.TransactOpts, chainId, gasPrice)
}

// SetExecPostsTo is a paid mutator transaction binding the contract method 0xf35eecbd.
//
// Solidity: function setExecPostsTo(uint64 chainId, uint64 postsTo) returns()
func (_FeeOracleV2 *FeeOracleV2Transactor) SetExecPostsTo(opts *bind.TransactOpts, chainId uint64, postsTo uint64) (*types.Transaction, error) {
	return _FeeOracleV2.contract.Transact(opts, "setExecPostsTo", chainId, postsTo)
}

// SetExecPostsTo is a paid mutator transaction binding the contract method 0xf35eecbd.
//
// Solidity: function setExecPostsTo(uint64 chainId, uint64 postsTo) returns()
func (_FeeOracleV2 *FeeOracleV2Session) SetExecPostsTo(chainId uint64, postsTo uint64) (*types.Transaction, error) {
	return _FeeOracleV2.Contract.SetExecPostsTo(&_FeeOracleV2.TransactOpts, chainId, postsTo)
}

// SetExecPostsTo is a paid mutator transaction binding the contract method 0xf35eecbd.
//
// Solidity: function setExecPostsTo(uint64 chainId, uint64 postsTo) returns()
func (_FeeOracleV2 *FeeOracleV2TransactorSession) SetExecPostsTo(chainId uint64, postsTo uint64) (*types.Transaction, error) {
	return _FeeOracleV2.Contract.SetExecPostsTo(&_FeeOracleV2.TransactOpts, chainId, postsTo)
}

// SetExecToNativeRate is a paid mutator transaction binding the contract method 0x65c6850d.
//
// Solidity: function setExecToNativeRate(uint64 chainId, uint64 rate) returns()
func (_FeeOracleV2 *FeeOracleV2Transactor) SetExecToNativeRate(opts *bind.TransactOpts, chainId uint64, rate uint64) (*types.Transaction, error) {
	return _FeeOracleV2.contract.Transact(opts, "setExecToNativeRate", chainId, rate)
}

// SetExecToNativeRate is a paid mutator transaction binding the contract method 0x65c6850d.
//
// Solidity: function setExecToNativeRate(uint64 chainId, uint64 rate) returns()
func (_FeeOracleV2 *FeeOracleV2Session) SetExecToNativeRate(chainId uint64, rate uint64) (*types.Transaction, error) {
	return _FeeOracleV2.Contract.SetExecToNativeRate(&_FeeOracleV2.TransactOpts, chainId, rate)
}

// SetExecToNativeRate is a paid mutator transaction binding the contract method 0x65c6850d.
//
// Solidity: function setExecToNativeRate(uint64 chainId, uint64 rate) returns()
func (_FeeOracleV2 *FeeOracleV2TransactorSession) SetExecToNativeRate(chainId uint64, rate uint64) (*types.Transaction, error) {
	return _FeeOracleV2.Contract.SetExecToNativeRate(&_FeeOracleV2.TransactOpts, chainId, rate)
}

// SetManager is a paid mutator transaction binding the contract method 0xd0ebdbe7.
//
// Solidity: function setManager(address manager_) returns()
func (_FeeOracleV2 *FeeOracleV2Transactor) SetManager(opts *bind.TransactOpts, manager_ common.Address) (*types.Transaction, error) {
	return _FeeOracleV2.contract.Transact(opts, "setManager", manager_)
}

// SetManager is a paid mutator transaction binding the contract method 0xd0ebdbe7.
//
// Solidity: function setManager(address manager_) returns()
func (_FeeOracleV2 *FeeOracleV2Session) SetManager(manager_ common.Address) (*types.Transaction, error) {
	return _FeeOracleV2.Contract.SetManager(&_FeeOracleV2.TransactOpts, manager_)
}

// SetManager is a paid mutator transaction binding the contract method 0xd0ebdbe7.
//
// Solidity: function setManager(address manager_) returns()
func (_FeeOracleV2 *FeeOracleV2TransactorSession) SetManager(manager_ common.Address) (*types.Transaction, error) {
	return _FeeOracleV2.Contract.SetManager(&_FeeOracleV2.TransactOpts, manager_)
}

// SetProtocolFee is a paid mutator transaction binding the contract method 0xfcb1deb1.
//
// Solidity: function setProtocolFee(uint72 fee) returns()
func (_FeeOracleV2 *FeeOracleV2Transactor) SetProtocolFee(opts *bind.TransactOpts, fee *big.Int) (*types.Transaction, error) {
	return _FeeOracleV2.contract.Transact(opts, "setProtocolFee", fee)
}

// SetProtocolFee is a paid mutator transaction binding the contract method 0xfcb1deb1.
//
// Solidity: function setProtocolFee(uint72 fee) returns()
func (_FeeOracleV2 *FeeOracleV2Session) SetProtocolFee(fee *big.Int) (*types.Transaction, error) {
	return _FeeOracleV2.Contract.SetProtocolFee(&_FeeOracleV2.TransactOpts, fee)
}

// SetProtocolFee is a paid mutator transaction binding the contract method 0xfcb1deb1.
//
// Solidity: function setProtocolFee(uint72 fee) returns()
func (_FeeOracleV2 *FeeOracleV2TransactorSession) SetProtocolFee(fee *big.Int) (*types.Transaction, error) {
	return _FeeOracleV2.Contract.SetProtocolFee(&_FeeOracleV2.TransactOpts, fee)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_FeeOracleV2 *FeeOracleV2Transactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _FeeOracleV2.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_FeeOracleV2 *FeeOracleV2Session) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _FeeOracleV2.Contract.TransferOwnership(&_FeeOracleV2.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_FeeOracleV2 *FeeOracleV2TransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _FeeOracleV2.Contract.TransferOwnership(&_FeeOracleV2.TransactOpts, newOwner)
}

// FeeOracleV2BaseGasLimitSetIterator is returned from FilterBaseGasLimitSet and is used to iterate over the raw logs and unpacked data for BaseGasLimitSet events raised by the FeeOracleV2 contract.
type FeeOracleV2BaseGasLimitSetIterator struct {
	Event *FeeOracleV2BaseGasLimitSet // Event containing the contract specifics and raw log

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
func (it *FeeOracleV2BaseGasLimitSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeeOracleV2BaseGasLimitSet)
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
		it.Event = new(FeeOracleV2BaseGasLimitSet)
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
func (it *FeeOracleV2BaseGasLimitSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FeeOracleV2BaseGasLimitSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FeeOracleV2BaseGasLimitSet represents a BaseGasLimitSet event raised by the FeeOracleV2 contract.
type FeeOracleV2BaseGasLimitSet struct {
	BaseGasLimit *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterBaseGasLimitSet is a free log retrieval operation binding the contract event 0xb2497d19702e6a33eb2a5487a9ad5977a0284bfe1ccad332e31b8a81a8c5ebaf.
//
// Solidity: event BaseGasLimitSet(uint24 baseGasLimit)
func (_FeeOracleV2 *FeeOracleV2Filterer) FilterBaseGasLimitSet(opts *bind.FilterOpts) (*FeeOracleV2BaseGasLimitSetIterator, error) {

	logs, sub, err := _FeeOracleV2.contract.FilterLogs(opts, "BaseGasLimitSet")
	if err != nil {
		return nil, err
	}
	return &FeeOracleV2BaseGasLimitSetIterator{contract: _FeeOracleV2.contract, event: "BaseGasLimitSet", logs: logs, sub: sub}, nil
}

// WatchBaseGasLimitSet is a free log subscription operation binding the contract event 0xb2497d19702e6a33eb2a5487a9ad5977a0284bfe1ccad332e31b8a81a8c5ebaf.
//
// Solidity: event BaseGasLimitSet(uint24 baseGasLimit)
func (_FeeOracleV2 *FeeOracleV2Filterer) WatchBaseGasLimitSet(opts *bind.WatchOpts, sink chan<- *FeeOracleV2BaseGasLimitSet) (event.Subscription, error) {

	logs, sub, err := _FeeOracleV2.contract.WatchLogs(opts, "BaseGasLimitSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FeeOracleV2BaseGasLimitSet)
				if err := _FeeOracleV2.contract.UnpackLog(event, "BaseGasLimitSet", log); err != nil {
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

// ParseBaseGasLimitSet is a log parse operation binding the contract event 0xb2497d19702e6a33eb2a5487a9ad5977a0284bfe1ccad332e31b8a81a8c5ebaf.
//
// Solidity: event BaseGasLimitSet(uint24 baseGasLimit)
func (_FeeOracleV2 *FeeOracleV2Filterer) ParseBaseGasLimitSet(log types.Log) (*FeeOracleV2BaseGasLimitSet, error) {
	event := new(FeeOracleV2BaseGasLimitSet)
	if err := _FeeOracleV2.contract.UnpackLog(event, "BaseGasLimitSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FeeOracleV2DataFeeParamsSetIterator is returned from FilterDataFeeParamsSet and is used to iterate over the raw logs and unpacked data for DataFeeParamsSet events raised by the FeeOracleV2 contract.
type FeeOracleV2DataFeeParamsSetIterator struct {
	Event *FeeOracleV2DataFeeParamsSet // Event containing the contract specifics and raw log

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
func (it *FeeOracleV2DataFeeParamsSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeeOracleV2DataFeeParamsSet)
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
		it.Event = new(FeeOracleV2DataFeeParamsSet)
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
func (it *FeeOracleV2DataFeeParamsSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FeeOracleV2DataFeeParamsSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FeeOracleV2DataFeeParamsSet represents a DataFeeParamsSet event raised by the FeeOracleV2 contract.
type FeeOracleV2DataFeeParamsSet struct {
	ChainId      uint64
	SizeBuffer   uint64
	DataGasPrice uint64
	ToNativeRate uint64
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterDataFeeParamsSet is a free log retrieval operation binding the contract event 0xd0b18837916cfd80f0f320486a107794cf007ff14ddafb4483a62c3fb5ec05a3.
//
// Solidity: event DataFeeParamsSet(uint64 chainId, uint64 sizeBuffer, uint64 dataGasPrice, uint64 toNativeRate)
func (_FeeOracleV2 *FeeOracleV2Filterer) FilterDataFeeParamsSet(opts *bind.FilterOpts) (*FeeOracleV2DataFeeParamsSetIterator, error) {

	logs, sub, err := _FeeOracleV2.contract.FilterLogs(opts, "DataFeeParamsSet")
	if err != nil {
		return nil, err
	}
	return &FeeOracleV2DataFeeParamsSetIterator{contract: _FeeOracleV2.contract, event: "DataFeeParamsSet", logs: logs, sub: sub}, nil
}

// WatchDataFeeParamsSet is a free log subscription operation binding the contract event 0xd0b18837916cfd80f0f320486a107794cf007ff14ddafb4483a62c3fb5ec05a3.
//
// Solidity: event DataFeeParamsSet(uint64 chainId, uint64 sizeBuffer, uint64 dataGasPrice, uint64 toNativeRate)
func (_FeeOracleV2 *FeeOracleV2Filterer) WatchDataFeeParamsSet(opts *bind.WatchOpts, sink chan<- *FeeOracleV2DataFeeParamsSet) (event.Subscription, error) {

	logs, sub, err := _FeeOracleV2.contract.WatchLogs(opts, "DataFeeParamsSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FeeOracleV2DataFeeParamsSet)
				if err := _FeeOracleV2.contract.UnpackLog(event, "DataFeeParamsSet", log); err != nil {
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

// ParseDataFeeParamsSet is a log parse operation binding the contract event 0xd0b18837916cfd80f0f320486a107794cf007ff14ddafb4483a62c3fb5ec05a3.
//
// Solidity: event DataFeeParamsSet(uint64 chainId, uint64 sizeBuffer, uint64 dataGasPrice, uint64 toNativeRate)
func (_FeeOracleV2 *FeeOracleV2Filterer) ParseDataFeeParamsSet(log types.Log) (*FeeOracleV2DataFeeParamsSet, error) {
	event := new(FeeOracleV2DataFeeParamsSet)
	if err := _FeeOracleV2.contract.UnpackLog(event, "DataFeeParamsSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FeeOracleV2DataGasPriceSetIterator is returned from FilterDataGasPriceSet and is used to iterate over the raw logs and unpacked data for DataGasPriceSet events raised by the FeeOracleV2 contract.
type FeeOracleV2DataGasPriceSetIterator struct {
	Event *FeeOracleV2DataGasPriceSet // Event containing the contract specifics and raw log

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
func (it *FeeOracleV2DataGasPriceSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeeOracleV2DataGasPriceSet)
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
		it.Event = new(FeeOracleV2DataGasPriceSet)
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
func (it *FeeOracleV2DataGasPriceSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FeeOracleV2DataGasPriceSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FeeOracleV2DataGasPriceSet represents a DataGasPriceSet event raised by the FeeOracleV2 contract.
type FeeOracleV2DataGasPriceSet struct {
	ChainId  uint64
	GasPrice uint64
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterDataGasPriceSet is a free log retrieval operation binding the contract event 0xd7d8dd5a956a8bd500e02d52d0a9dd8a0e2955ec48771a8c9da485e6706c66fb.
//
// Solidity: event DataGasPriceSet(uint64 chainId, uint64 gasPrice)
func (_FeeOracleV2 *FeeOracleV2Filterer) FilterDataGasPriceSet(opts *bind.FilterOpts) (*FeeOracleV2DataGasPriceSetIterator, error) {

	logs, sub, err := _FeeOracleV2.contract.FilterLogs(opts, "DataGasPriceSet")
	if err != nil {
		return nil, err
	}
	return &FeeOracleV2DataGasPriceSetIterator{contract: _FeeOracleV2.contract, event: "DataGasPriceSet", logs: logs, sub: sub}, nil
}

// WatchDataGasPriceSet is a free log subscription operation binding the contract event 0xd7d8dd5a956a8bd500e02d52d0a9dd8a0e2955ec48771a8c9da485e6706c66fb.
//
// Solidity: event DataGasPriceSet(uint64 chainId, uint64 gasPrice)
func (_FeeOracleV2 *FeeOracleV2Filterer) WatchDataGasPriceSet(opts *bind.WatchOpts, sink chan<- *FeeOracleV2DataGasPriceSet) (event.Subscription, error) {

	logs, sub, err := _FeeOracleV2.contract.WatchLogs(opts, "DataGasPriceSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FeeOracleV2DataGasPriceSet)
				if err := _FeeOracleV2.contract.UnpackLog(event, "DataGasPriceSet", log); err != nil {
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

// ParseDataGasPriceSet is a log parse operation binding the contract event 0xd7d8dd5a956a8bd500e02d52d0a9dd8a0e2955ec48771a8c9da485e6706c66fb.
//
// Solidity: event DataGasPriceSet(uint64 chainId, uint64 gasPrice)
func (_FeeOracleV2 *FeeOracleV2Filterer) ParseDataGasPriceSet(log types.Log) (*FeeOracleV2DataGasPriceSet, error) {
	event := new(FeeOracleV2DataGasPriceSet)
	if err := _FeeOracleV2.contract.UnpackLog(event, "DataGasPriceSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FeeOracleV2DataSizeBufferSetIterator is returned from FilterDataSizeBufferSet and is used to iterate over the raw logs and unpacked data for DataSizeBufferSet events raised by the FeeOracleV2 contract.
type FeeOracleV2DataSizeBufferSetIterator struct {
	Event *FeeOracleV2DataSizeBufferSet // Event containing the contract specifics and raw log

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
func (it *FeeOracleV2DataSizeBufferSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeeOracleV2DataSizeBufferSet)
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
		it.Event = new(FeeOracleV2DataSizeBufferSet)
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
func (it *FeeOracleV2DataSizeBufferSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FeeOracleV2DataSizeBufferSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FeeOracleV2DataSizeBufferSet represents a DataSizeBufferSet event raised by the FeeOracleV2 contract.
type FeeOracleV2DataSizeBufferSet struct {
	ChainId    uint64
	SizeBuffer uint64
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterDataSizeBufferSet is a free log retrieval operation binding the contract event 0x0b1d18b6fe29df95e0bd90abbbf45e65e4a14f3ef2126bfcfb7f2a0811266998.
//
// Solidity: event DataSizeBufferSet(uint64 chainId, uint64 sizeBuffer)
func (_FeeOracleV2 *FeeOracleV2Filterer) FilterDataSizeBufferSet(opts *bind.FilterOpts) (*FeeOracleV2DataSizeBufferSetIterator, error) {

	logs, sub, err := _FeeOracleV2.contract.FilterLogs(opts, "DataSizeBufferSet")
	if err != nil {
		return nil, err
	}
	return &FeeOracleV2DataSizeBufferSetIterator{contract: _FeeOracleV2.contract, event: "DataSizeBufferSet", logs: logs, sub: sub}, nil
}

// WatchDataSizeBufferSet is a free log subscription operation binding the contract event 0x0b1d18b6fe29df95e0bd90abbbf45e65e4a14f3ef2126bfcfb7f2a0811266998.
//
// Solidity: event DataSizeBufferSet(uint64 chainId, uint64 sizeBuffer)
func (_FeeOracleV2 *FeeOracleV2Filterer) WatchDataSizeBufferSet(opts *bind.WatchOpts, sink chan<- *FeeOracleV2DataSizeBufferSet) (event.Subscription, error) {

	logs, sub, err := _FeeOracleV2.contract.WatchLogs(opts, "DataSizeBufferSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FeeOracleV2DataSizeBufferSet)
				if err := _FeeOracleV2.contract.UnpackLog(event, "DataSizeBufferSet", log); err != nil {
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

// ParseDataSizeBufferSet is a log parse operation binding the contract event 0x0b1d18b6fe29df95e0bd90abbbf45e65e4a14f3ef2126bfcfb7f2a0811266998.
//
// Solidity: event DataSizeBufferSet(uint64 chainId, uint64 sizeBuffer)
func (_FeeOracleV2 *FeeOracleV2Filterer) ParseDataSizeBufferSet(log types.Log) (*FeeOracleV2DataSizeBufferSet, error) {
	event := new(FeeOracleV2DataSizeBufferSet)
	if err := _FeeOracleV2.contract.UnpackLog(event, "DataSizeBufferSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FeeOracleV2DataToNativeRateSetIterator is returned from FilterDataToNativeRateSet and is used to iterate over the raw logs and unpacked data for DataToNativeRateSet events raised by the FeeOracleV2 contract.
type FeeOracleV2DataToNativeRateSetIterator struct {
	Event *FeeOracleV2DataToNativeRateSet // Event containing the contract specifics and raw log

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
func (it *FeeOracleV2DataToNativeRateSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeeOracleV2DataToNativeRateSet)
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
		it.Event = new(FeeOracleV2DataToNativeRateSet)
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
func (it *FeeOracleV2DataToNativeRateSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FeeOracleV2DataToNativeRateSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FeeOracleV2DataToNativeRateSet represents a DataToNativeRateSet event raised by the FeeOracleV2 contract.
type FeeOracleV2DataToNativeRateSet struct {
	ChainId      uint64
	ToNativeRate uint64
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterDataToNativeRateSet is a free log retrieval operation binding the contract event 0x1e87d1bd7c559f1a6fcee0e50d1f7d7df76cee0bafc691051b174dc2a8313921.
//
// Solidity: event DataToNativeRateSet(uint64 chainId, uint64 toNativeRate)
func (_FeeOracleV2 *FeeOracleV2Filterer) FilterDataToNativeRateSet(opts *bind.FilterOpts) (*FeeOracleV2DataToNativeRateSetIterator, error) {

	logs, sub, err := _FeeOracleV2.contract.FilterLogs(opts, "DataToNativeRateSet")
	if err != nil {
		return nil, err
	}
	return &FeeOracleV2DataToNativeRateSetIterator{contract: _FeeOracleV2.contract, event: "DataToNativeRateSet", logs: logs, sub: sub}, nil
}

// WatchDataToNativeRateSet is a free log subscription operation binding the contract event 0x1e87d1bd7c559f1a6fcee0e50d1f7d7df76cee0bafc691051b174dc2a8313921.
//
// Solidity: event DataToNativeRateSet(uint64 chainId, uint64 toNativeRate)
func (_FeeOracleV2 *FeeOracleV2Filterer) WatchDataToNativeRateSet(opts *bind.WatchOpts, sink chan<- *FeeOracleV2DataToNativeRateSet) (event.Subscription, error) {

	logs, sub, err := _FeeOracleV2.contract.WatchLogs(opts, "DataToNativeRateSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FeeOracleV2DataToNativeRateSet)
				if err := _FeeOracleV2.contract.UnpackLog(event, "DataToNativeRateSet", log); err != nil {
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

// ParseDataToNativeRateSet is a log parse operation binding the contract event 0x1e87d1bd7c559f1a6fcee0e50d1f7d7df76cee0bafc691051b174dc2a8313921.
//
// Solidity: event DataToNativeRateSet(uint64 chainId, uint64 toNativeRate)
func (_FeeOracleV2 *FeeOracleV2Filterer) ParseDataToNativeRateSet(log types.Log) (*FeeOracleV2DataToNativeRateSet, error) {
	event := new(FeeOracleV2DataToNativeRateSet)
	if err := _FeeOracleV2.contract.UnpackLog(event, "DataToNativeRateSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FeeOracleV2ExecFeeParamsSetIterator is returned from FilterExecFeeParamsSet and is used to iterate over the raw logs and unpacked data for ExecFeeParamsSet events raised by the FeeOracleV2 contract.
type FeeOracleV2ExecFeeParamsSetIterator struct {
	Event *FeeOracleV2ExecFeeParamsSet // Event containing the contract specifics and raw log

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
func (it *FeeOracleV2ExecFeeParamsSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeeOracleV2ExecFeeParamsSet)
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
		it.Event = new(FeeOracleV2ExecFeeParamsSet)
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
func (it *FeeOracleV2ExecFeeParamsSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FeeOracleV2ExecFeeParamsSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FeeOracleV2ExecFeeParamsSet represents a ExecFeeParamsSet event raised by the FeeOracleV2 contract.
type FeeOracleV2ExecFeeParamsSet struct {
	ChainId      uint64
	PostsTo      uint64
	ExecGasPrice uint64
	ToNativeRate uint64
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterExecFeeParamsSet is a free log retrieval operation binding the contract event 0x0a05333d850372b39fcfc14c6eed42ebc766b01f516add8c20240696552c4464.
//
// Solidity: event ExecFeeParamsSet(uint64 chainId, uint64 postsTo, uint64 execGasPrice, uint64 toNativeRate)
func (_FeeOracleV2 *FeeOracleV2Filterer) FilterExecFeeParamsSet(opts *bind.FilterOpts) (*FeeOracleV2ExecFeeParamsSetIterator, error) {

	logs, sub, err := _FeeOracleV2.contract.FilterLogs(opts, "ExecFeeParamsSet")
	if err != nil {
		return nil, err
	}
	return &FeeOracleV2ExecFeeParamsSetIterator{contract: _FeeOracleV2.contract, event: "ExecFeeParamsSet", logs: logs, sub: sub}, nil
}

// WatchExecFeeParamsSet is a free log subscription operation binding the contract event 0x0a05333d850372b39fcfc14c6eed42ebc766b01f516add8c20240696552c4464.
//
// Solidity: event ExecFeeParamsSet(uint64 chainId, uint64 postsTo, uint64 execGasPrice, uint64 toNativeRate)
func (_FeeOracleV2 *FeeOracleV2Filterer) WatchExecFeeParamsSet(opts *bind.WatchOpts, sink chan<- *FeeOracleV2ExecFeeParamsSet) (event.Subscription, error) {

	logs, sub, err := _FeeOracleV2.contract.WatchLogs(opts, "ExecFeeParamsSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FeeOracleV2ExecFeeParamsSet)
				if err := _FeeOracleV2.contract.UnpackLog(event, "ExecFeeParamsSet", log); err != nil {
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

// ParseExecFeeParamsSet is a log parse operation binding the contract event 0x0a05333d850372b39fcfc14c6eed42ebc766b01f516add8c20240696552c4464.
//
// Solidity: event ExecFeeParamsSet(uint64 chainId, uint64 postsTo, uint64 execGasPrice, uint64 toNativeRate)
func (_FeeOracleV2 *FeeOracleV2Filterer) ParseExecFeeParamsSet(log types.Log) (*FeeOracleV2ExecFeeParamsSet, error) {
	event := new(FeeOracleV2ExecFeeParamsSet)
	if err := _FeeOracleV2.contract.UnpackLog(event, "ExecFeeParamsSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FeeOracleV2ExecGasPriceSetIterator is returned from FilterExecGasPriceSet and is used to iterate over the raw logs and unpacked data for ExecGasPriceSet events raised by the FeeOracleV2 contract.
type FeeOracleV2ExecGasPriceSetIterator struct {
	Event *FeeOracleV2ExecGasPriceSet // Event containing the contract specifics and raw log

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
func (it *FeeOracleV2ExecGasPriceSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeeOracleV2ExecGasPriceSet)
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
		it.Event = new(FeeOracleV2ExecGasPriceSet)
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
func (it *FeeOracleV2ExecGasPriceSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FeeOracleV2ExecGasPriceSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FeeOracleV2ExecGasPriceSet represents a ExecGasPriceSet event raised by the FeeOracleV2 contract.
type FeeOracleV2ExecGasPriceSet struct {
	ChainId  uint64
	GasPrice uint64
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterExecGasPriceSet is a free log retrieval operation binding the contract event 0xe0e5abb8929e27a69d77f47a4e3f9575411a5be1fa596e5b55078d7850f358db.
//
// Solidity: event ExecGasPriceSet(uint64 chainId, uint64 gasPrice)
func (_FeeOracleV2 *FeeOracleV2Filterer) FilterExecGasPriceSet(opts *bind.FilterOpts) (*FeeOracleV2ExecGasPriceSetIterator, error) {

	logs, sub, err := _FeeOracleV2.contract.FilterLogs(opts, "ExecGasPriceSet")
	if err != nil {
		return nil, err
	}
	return &FeeOracleV2ExecGasPriceSetIterator{contract: _FeeOracleV2.contract, event: "ExecGasPriceSet", logs: logs, sub: sub}, nil
}

// WatchExecGasPriceSet is a free log subscription operation binding the contract event 0xe0e5abb8929e27a69d77f47a4e3f9575411a5be1fa596e5b55078d7850f358db.
//
// Solidity: event ExecGasPriceSet(uint64 chainId, uint64 gasPrice)
func (_FeeOracleV2 *FeeOracleV2Filterer) WatchExecGasPriceSet(opts *bind.WatchOpts, sink chan<- *FeeOracleV2ExecGasPriceSet) (event.Subscription, error) {

	logs, sub, err := _FeeOracleV2.contract.WatchLogs(opts, "ExecGasPriceSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FeeOracleV2ExecGasPriceSet)
				if err := _FeeOracleV2.contract.UnpackLog(event, "ExecGasPriceSet", log); err != nil {
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

// ParseExecGasPriceSet is a log parse operation binding the contract event 0xe0e5abb8929e27a69d77f47a4e3f9575411a5be1fa596e5b55078d7850f358db.
//
// Solidity: event ExecGasPriceSet(uint64 chainId, uint64 gasPrice)
func (_FeeOracleV2 *FeeOracleV2Filterer) ParseExecGasPriceSet(log types.Log) (*FeeOracleV2ExecGasPriceSet, error) {
	event := new(FeeOracleV2ExecGasPriceSet)
	if err := _FeeOracleV2.contract.UnpackLog(event, "ExecGasPriceSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FeeOracleV2ExecPostsToSetIterator is returned from FilterExecPostsToSet and is used to iterate over the raw logs and unpacked data for ExecPostsToSet events raised by the FeeOracleV2 contract.
type FeeOracleV2ExecPostsToSetIterator struct {
	Event *FeeOracleV2ExecPostsToSet // Event containing the contract specifics and raw log

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
func (it *FeeOracleV2ExecPostsToSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeeOracleV2ExecPostsToSet)
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
		it.Event = new(FeeOracleV2ExecPostsToSet)
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
func (it *FeeOracleV2ExecPostsToSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FeeOracleV2ExecPostsToSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FeeOracleV2ExecPostsToSet represents a ExecPostsToSet event raised by the FeeOracleV2 contract.
type FeeOracleV2ExecPostsToSet struct {
	ChainId uint64
	PostsTo uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterExecPostsToSet is a free log retrieval operation binding the contract event 0x62f13d8522c451be7798aa5f1e258a7780ef5296115067b8e632436d4548a11d.
//
// Solidity: event ExecPostsToSet(uint64 chainId, uint64 postsTo)
func (_FeeOracleV2 *FeeOracleV2Filterer) FilterExecPostsToSet(opts *bind.FilterOpts) (*FeeOracleV2ExecPostsToSetIterator, error) {

	logs, sub, err := _FeeOracleV2.contract.FilterLogs(opts, "ExecPostsToSet")
	if err != nil {
		return nil, err
	}
	return &FeeOracleV2ExecPostsToSetIterator{contract: _FeeOracleV2.contract, event: "ExecPostsToSet", logs: logs, sub: sub}, nil
}

// WatchExecPostsToSet is a free log subscription operation binding the contract event 0x62f13d8522c451be7798aa5f1e258a7780ef5296115067b8e632436d4548a11d.
//
// Solidity: event ExecPostsToSet(uint64 chainId, uint64 postsTo)
func (_FeeOracleV2 *FeeOracleV2Filterer) WatchExecPostsToSet(opts *bind.WatchOpts, sink chan<- *FeeOracleV2ExecPostsToSet) (event.Subscription, error) {

	logs, sub, err := _FeeOracleV2.contract.WatchLogs(opts, "ExecPostsToSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FeeOracleV2ExecPostsToSet)
				if err := _FeeOracleV2.contract.UnpackLog(event, "ExecPostsToSet", log); err != nil {
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

// ParseExecPostsToSet is a log parse operation binding the contract event 0x62f13d8522c451be7798aa5f1e258a7780ef5296115067b8e632436d4548a11d.
//
// Solidity: event ExecPostsToSet(uint64 chainId, uint64 postsTo)
func (_FeeOracleV2 *FeeOracleV2Filterer) ParseExecPostsToSet(log types.Log) (*FeeOracleV2ExecPostsToSet, error) {
	event := new(FeeOracleV2ExecPostsToSet)
	if err := _FeeOracleV2.contract.UnpackLog(event, "ExecPostsToSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FeeOracleV2ExecToNativeRateSetIterator is returned from FilterExecToNativeRateSet and is used to iterate over the raw logs and unpacked data for ExecToNativeRateSet events raised by the FeeOracleV2 contract.
type FeeOracleV2ExecToNativeRateSetIterator struct {
	Event *FeeOracleV2ExecToNativeRateSet // Event containing the contract specifics and raw log

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
func (it *FeeOracleV2ExecToNativeRateSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeeOracleV2ExecToNativeRateSet)
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
		it.Event = new(FeeOracleV2ExecToNativeRateSet)
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
func (it *FeeOracleV2ExecToNativeRateSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FeeOracleV2ExecToNativeRateSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FeeOracleV2ExecToNativeRateSet represents a ExecToNativeRateSet event raised by the FeeOracleV2 contract.
type FeeOracleV2ExecToNativeRateSet struct {
	ChainId      uint64
	ToNativeRate uint64
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterExecToNativeRateSet is a free log retrieval operation binding the contract event 0x5728e411b1606c4a6fd53e10049ebe42400b3617c09e478e10239a913ab3493b.
//
// Solidity: event ExecToNativeRateSet(uint64 chainId, uint64 toNativeRate)
func (_FeeOracleV2 *FeeOracleV2Filterer) FilterExecToNativeRateSet(opts *bind.FilterOpts) (*FeeOracleV2ExecToNativeRateSetIterator, error) {

	logs, sub, err := _FeeOracleV2.contract.FilterLogs(opts, "ExecToNativeRateSet")
	if err != nil {
		return nil, err
	}
	return &FeeOracleV2ExecToNativeRateSetIterator{contract: _FeeOracleV2.contract, event: "ExecToNativeRateSet", logs: logs, sub: sub}, nil
}

// WatchExecToNativeRateSet is a free log subscription operation binding the contract event 0x5728e411b1606c4a6fd53e10049ebe42400b3617c09e478e10239a913ab3493b.
//
// Solidity: event ExecToNativeRateSet(uint64 chainId, uint64 toNativeRate)
func (_FeeOracleV2 *FeeOracleV2Filterer) WatchExecToNativeRateSet(opts *bind.WatchOpts, sink chan<- *FeeOracleV2ExecToNativeRateSet) (event.Subscription, error) {

	logs, sub, err := _FeeOracleV2.contract.WatchLogs(opts, "ExecToNativeRateSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FeeOracleV2ExecToNativeRateSet)
				if err := _FeeOracleV2.contract.UnpackLog(event, "ExecToNativeRateSet", log); err != nil {
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

// ParseExecToNativeRateSet is a log parse operation binding the contract event 0x5728e411b1606c4a6fd53e10049ebe42400b3617c09e478e10239a913ab3493b.
//
// Solidity: event ExecToNativeRateSet(uint64 chainId, uint64 toNativeRate)
func (_FeeOracleV2 *FeeOracleV2Filterer) ParseExecToNativeRateSet(log types.Log) (*FeeOracleV2ExecToNativeRateSet, error) {
	event := new(FeeOracleV2ExecToNativeRateSet)
	if err := _FeeOracleV2.contract.UnpackLog(event, "ExecToNativeRateSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FeeOracleV2InitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the FeeOracleV2 contract.
type FeeOracleV2InitializedIterator struct {
	Event *FeeOracleV2Initialized // Event containing the contract specifics and raw log

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
func (it *FeeOracleV2InitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeeOracleV2Initialized)
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
		it.Event = new(FeeOracleV2Initialized)
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
func (it *FeeOracleV2InitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FeeOracleV2InitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FeeOracleV2Initialized represents a Initialized event raised by the FeeOracleV2 contract.
type FeeOracleV2Initialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_FeeOracleV2 *FeeOracleV2Filterer) FilterInitialized(opts *bind.FilterOpts) (*FeeOracleV2InitializedIterator, error) {

	logs, sub, err := _FeeOracleV2.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &FeeOracleV2InitializedIterator{contract: _FeeOracleV2.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_FeeOracleV2 *FeeOracleV2Filterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *FeeOracleV2Initialized) (event.Subscription, error) {

	logs, sub, err := _FeeOracleV2.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FeeOracleV2Initialized)
				if err := _FeeOracleV2.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_FeeOracleV2 *FeeOracleV2Filterer) ParseInitialized(log types.Log) (*FeeOracleV2Initialized, error) {
	event := new(FeeOracleV2Initialized)
	if err := _FeeOracleV2.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FeeOracleV2ManagerSetIterator is returned from FilterManagerSet and is used to iterate over the raw logs and unpacked data for ManagerSet events raised by the FeeOracleV2 contract.
type FeeOracleV2ManagerSetIterator struct {
	Event *FeeOracleV2ManagerSet // Event containing the contract specifics and raw log

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
func (it *FeeOracleV2ManagerSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeeOracleV2ManagerSet)
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
		it.Event = new(FeeOracleV2ManagerSet)
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
func (it *FeeOracleV2ManagerSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FeeOracleV2ManagerSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FeeOracleV2ManagerSet represents a ManagerSet event raised by the FeeOracleV2 contract.
type FeeOracleV2ManagerSet struct {
	Manager common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterManagerSet is a free log retrieval operation binding the contract event 0x60a0f5b9f9e81e98216071b85826681c796256fe3d1354ecb675580fba64fa69.
//
// Solidity: event ManagerSet(address manager)
func (_FeeOracleV2 *FeeOracleV2Filterer) FilterManagerSet(opts *bind.FilterOpts) (*FeeOracleV2ManagerSetIterator, error) {

	logs, sub, err := _FeeOracleV2.contract.FilterLogs(opts, "ManagerSet")
	if err != nil {
		return nil, err
	}
	return &FeeOracleV2ManagerSetIterator{contract: _FeeOracleV2.contract, event: "ManagerSet", logs: logs, sub: sub}, nil
}

// WatchManagerSet is a free log subscription operation binding the contract event 0x60a0f5b9f9e81e98216071b85826681c796256fe3d1354ecb675580fba64fa69.
//
// Solidity: event ManagerSet(address manager)
func (_FeeOracleV2 *FeeOracleV2Filterer) WatchManagerSet(opts *bind.WatchOpts, sink chan<- *FeeOracleV2ManagerSet) (event.Subscription, error) {

	logs, sub, err := _FeeOracleV2.contract.WatchLogs(opts, "ManagerSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FeeOracleV2ManagerSet)
				if err := _FeeOracleV2.contract.UnpackLog(event, "ManagerSet", log); err != nil {
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
func (_FeeOracleV2 *FeeOracleV2Filterer) ParseManagerSet(log types.Log) (*FeeOracleV2ManagerSet, error) {
	event := new(FeeOracleV2ManagerSet)
	if err := _FeeOracleV2.contract.UnpackLog(event, "ManagerSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FeeOracleV2OwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the FeeOracleV2 contract.
type FeeOracleV2OwnershipTransferredIterator struct {
	Event *FeeOracleV2OwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *FeeOracleV2OwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeeOracleV2OwnershipTransferred)
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
		it.Event = new(FeeOracleV2OwnershipTransferred)
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
func (it *FeeOracleV2OwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FeeOracleV2OwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FeeOracleV2OwnershipTransferred represents a OwnershipTransferred event raised by the FeeOracleV2 contract.
type FeeOracleV2OwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_FeeOracleV2 *FeeOracleV2Filterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*FeeOracleV2OwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _FeeOracleV2.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &FeeOracleV2OwnershipTransferredIterator{contract: _FeeOracleV2.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_FeeOracleV2 *FeeOracleV2Filterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *FeeOracleV2OwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _FeeOracleV2.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FeeOracleV2OwnershipTransferred)
				if err := _FeeOracleV2.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_FeeOracleV2 *FeeOracleV2Filterer) ParseOwnershipTransferred(log types.Log) (*FeeOracleV2OwnershipTransferred, error) {
	event := new(FeeOracleV2OwnershipTransferred)
	if err := _FeeOracleV2.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FeeOracleV2ProtocolFeeSetIterator is returned from FilterProtocolFeeSet and is used to iterate over the raw logs and unpacked data for ProtocolFeeSet events raised by the FeeOracleV2 contract.
type FeeOracleV2ProtocolFeeSetIterator struct {
	Event *FeeOracleV2ProtocolFeeSet // Event containing the contract specifics and raw log

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
func (it *FeeOracleV2ProtocolFeeSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeeOracleV2ProtocolFeeSet)
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
		it.Event = new(FeeOracleV2ProtocolFeeSet)
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
func (it *FeeOracleV2ProtocolFeeSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FeeOracleV2ProtocolFeeSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FeeOracleV2ProtocolFeeSet represents a ProtocolFeeSet event raised by the FeeOracleV2 contract.
type FeeOracleV2ProtocolFeeSet struct {
	ProtocolFee *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterProtocolFeeSet is a free log retrieval operation binding the contract event 0x92c98513056b0d0d263b0f7153c99e1651791a983e75fc558b424b3aa0b623f7.
//
// Solidity: event ProtocolFeeSet(uint72 protocolFee)
func (_FeeOracleV2 *FeeOracleV2Filterer) FilterProtocolFeeSet(opts *bind.FilterOpts) (*FeeOracleV2ProtocolFeeSetIterator, error) {

	logs, sub, err := _FeeOracleV2.contract.FilterLogs(opts, "ProtocolFeeSet")
	if err != nil {
		return nil, err
	}
	return &FeeOracleV2ProtocolFeeSetIterator{contract: _FeeOracleV2.contract, event: "ProtocolFeeSet", logs: logs, sub: sub}, nil
}

// WatchProtocolFeeSet is a free log subscription operation binding the contract event 0x92c98513056b0d0d263b0f7153c99e1651791a983e75fc558b424b3aa0b623f7.
//
// Solidity: event ProtocolFeeSet(uint72 protocolFee)
func (_FeeOracleV2 *FeeOracleV2Filterer) WatchProtocolFeeSet(opts *bind.WatchOpts, sink chan<- *FeeOracleV2ProtocolFeeSet) (event.Subscription, error) {

	logs, sub, err := _FeeOracleV2.contract.WatchLogs(opts, "ProtocolFeeSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FeeOracleV2ProtocolFeeSet)
				if err := _FeeOracleV2.contract.UnpackLog(event, "ProtocolFeeSet", log); err != nil {
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

// ParseProtocolFeeSet is a log parse operation binding the contract event 0x92c98513056b0d0d263b0f7153c99e1651791a983e75fc558b424b3aa0b623f7.
//
// Solidity: event ProtocolFeeSet(uint72 protocolFee)
func (_FeeOracleV2 *FeeOracleV2Filterer) ParseProtocolFeeSet(log types.Log) (*FeeOracleV2ProtocolFeeSet, error) {
	event := new(FeeOracleV2ProtocolFeeSet)
	if err := _FeeOracleV2.contract.UnpackLog(event, "ProtocolFeeSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
