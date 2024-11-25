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

// IFeeOracleV2DataCostParams is an auto generated low-level Go binding around an user-defined struct.
type IFeeOracleV2DataCostParams struct {
	GasToken       uint8
	BaseDataBuffer uint32
	DataCostId     uint64
	GasPrice       uint64
	GasPerByte     uint64
}

// IFeeOracleV2FeeParams is an auto generated low-level Go binding around an user-defined struct.
type IFeeOracleV2FeeParams struct {
	GasToken     uint8
	BaseGasLimit uint32
	ChainId      uint64
	GasPrice     uint64
	DataCostId   uint64
}

// IFeeOracleV2NativeRateParams is an auto generated low-level Go binding around an user-defined struct.
type IFeeOracleV2NativeRateParams struct {
	GasToken   uint8
	NativeRate *big.Int
}

// FeeOracleV2MetaData contains all meta data concerning the FeeOracleV2 contract.
var FeeOracleV2MetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"CONVERSION_RATE_DENOM\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"baseDataBuffer\",\"inputs\":[{\"name\":\"dataCostId\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"baseGasLimit\",\"inputs\":[{\"name\":\"chainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"bulkSetDataCostParams\",\"inputs\":[{\"name\":\"params\",\"type\":\"tuple[]\",\"internalType\":\"structIFeeOracleV2.DataCostParams[]\",\"components\":[{\"name\":\"gasToken\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"baseDataBuffer\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"dataCostId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"gasPrice\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"gasPerByte\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"bulkSetFeeParams\",\"inputs\":[{\"name\":\"params\",\"type\":\"tuple[]\",\"internalType\":\"structIFeeOracleV2.FeeParams[]\",\"components\":[{\"name\":\"gasToken\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"baseGasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"chainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"gasPrice\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"dataCostId\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"bulkSetToNativeRate\",\"inputs\":[{\"name\":\"params\",\"type\":\"tuple[]\",\"internalType\":\"structIFeeOracleV2.NativeRateParams[]\",\"components\":[{\"name\":\"gasToken\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"nativeRate\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"dataCostParams\",\"inputs\":[{\"name\":\"dataCostId\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIFeeOracleV2.DataCostParams\",\"components\":[{\"name\":\"gasToken\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"baseDataBuffer\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"dataCostId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"gasPrice\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"gasPerByte\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"dataGasPerByte\",\"inputs\":[{\"name\":\"dataCostId\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"dataGasPrice\",\"inputs\":[{\"name\":\"dataCostId\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"dataGasToken\",\"inputs\":[{\"name\":\"dataCostId\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"execDataCostId\",\"inputs\":[{\"name\":\"chainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"execGasPrice\",\"inputs\":[{\"name\":\"chainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"execGasToken\",\"inputs\":[{\"name\":\"chainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"feeFor\",\"inputs\":[{\"name\":\"destChainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"gasLimit\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"feeParams\",\"inputs\":[{\"name\":\"chainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIFeeOracleV2.FeeParams\",\"components\":[{\"name\":\"gasToken\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"baseGasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"chainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"gasPrice\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"dataCostId\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"owner_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"manager_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"protocolFee_\",\"type\":\"uint96\",\"internalType\":\"uint96\"},{\"name\":\"feeParams_\",\"type\":\"tuple[]\",\"internalType\":\"structIFeeOracleV2.FeeParams[]\",\"components\":[{\"name\":\"gasToken\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"baseGasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"chainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"gasPrice\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"dataCostId\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"dataCostParams_\",\"type\":\"tuple[]\",\"internalType\":\"structIFeeOracleV2.DataCostParams[]\",\"components\":[{\"name\":\"gasToken\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"baseDataBuffer\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"dataCostId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"gasPrice\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"gasPerByte\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"nativeRateParams_\",\"type\":\"tuple[]\",\"internalType\":\"structIFeeOracleV2.NativeRateParams[]\",\"components\":[{\"name\":\"gasToken\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"nativeRate\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"manager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"protocolFee\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint96\",\"internalType\":\"uint96\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setBaseDataBuffer\",\"inputs\":[{\"name\":\"dataCostId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"newBaseDataBuffer\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setBaseGasLimit\",\"inputs\":[{\"name\":\"chainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"newBaseGasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setDataCostId\",\"inputs\":[{\"name\":\"chainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"dataCostId\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setDataGasPrice\",\"inputs\":[{\"name\":\"dataCostId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"gasPrice\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setExecGasPrice\",\"inputs\":[{\"name\":\"chainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"gasPrice\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setGasPerByte\",\"inputs\":[{\"name\":\"dataCostId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"gasPerByte\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setManager\",\"inputs\":[{\"name\":\"manager_\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setProtocolFee\",\"inputs\":[{\"name\":\"fee\",\"type\":\"uint96\",\"internalType\":\"uint96\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setToNativeRate\",\"inputs\":[{\"name\":\"gasToken\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"nativeRate\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"toNativeRate\",\"inputs\":[{\"name\":\"chainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"toNativeRateData\",\"inputs\":[{\"name\":\"dataCostId\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"tokenToNativeRate\",\"inputs\":[{\"name\":\"gasToken\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"version\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"pure\"},{\"type\":\"event\",\"name\":\"BaseDataBufferSet\",\"inputs\":[{\"name\":\"dataCostId\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"baseDataBuffer\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"BaseGasLimitSet\",\"inputs\":[{\"name\":\"chainId\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"baseGasLimit\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DataCostIdSet\",\"inputs\":[{\"name\":\"chainId\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"dataCostId\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DataCostParamsSet\",\"inputs\":[{\"name\":\"gasToken\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"},{\"name\":\"dataCostId\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"gasPrice\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"gasPerByte\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DataGasPriceSet\",\"inputs\":[{\"name\":\"dataCostId\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"gasPrice\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ExecGasPriceSet\",\"inputs\":[{\"name\":\"chainId\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"gasPrice\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FeeParamsSet\",\"inputs\":[{\"name\":\"gasToken\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"},{\"name\":\"chainId\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"gasPrice\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"dataCostId\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"GasPerByteSet\",\"inputs\":[{\"name\":\"dataCostId\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"gasPerByte\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ManagerSet\",\"inputs\":[{\"name\":\"manager\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ProtocolFeeSet\",\"inputs\":[{\"name\":\"protocolFee\",\"type\":\"uint96\",\"indexed\":false,\"internalType\":\"uint96\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ToNativeRateSet\",\"inputs\":[{\"name\":\"gasToken\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"},{\"name\":\"nativeRate\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NoFeeParams\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotManager\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnableInvalidOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnableUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ZeroAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroChainId\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroDataCostId\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroGasPerByte\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroGasPrice\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroGasToken\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroNativeRate\",\"inputs\":[]}]",
	Bin: "0x608060405234801561001057600080fd5b5061001961001e565b6100d0565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000900460ff161561006e5760405163f92ee8a960e01b815260040160405180910390fd5b80546001600160401b03908116146100cd5780546001600160401b0319166001600160401b0390811782556040519081527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b50565b611eef806100df6000396000f3fe608060405234801561001057600080fd5b50600436106102065760003560e01c80638dd9523c1161011a578063bc51bf37116100ad578063db0018e81161007c578063db0018e8146106f3578063e21497b914610726578063e5737f6214610759578063f2fde38b1461076c578063fee86f691461077f57600080fd5b8063bc51bf37146105f8578063bfc714161461060b578063d0ebdbe7146106b4578063d32b68ad146106c757600080fd5b8063b0e21e8a116100e9578063b0e21e8a14610569578063b15268b014610594578063b984cc0b146105a7578063b9923e1c146105e557600080fd5b80638dd9523c146105065780638f9d6ace146105195780638ff83049146105235780639a5551c31461053657600080fd5b806350b815391161019d578063671402321161016c578063671402321461043b5780636b6dccfe14610482578063715018a6146104955780638b7bfd701461049d5780638da5cb5b146104d657600080fd5b806350b81539146103da57806354fd4d501461040e57806356bce459146104155780635d3acee21461042857600080fd5b8063176c73e3116101d9578063176c73e3146102945780632d4634a4146102a7578063415070af1461035d578063481c6a75146103a857600080fd5b806304e53a151461020b5780630baaa6aa146102205780630f0b435b146102335780631101885314610246575b600080fd5b61021e610219366004611920565b61079f565b005b61021e61022e366004611977565b6107df565b61021e6102413660046119cc565b61081b565b61027a6102543660046119ff565b6001600160401b0316600090815260036020526040902054610100900463ffffffff1690565b60405163ffffffff90911681526020015b60405180910390f35b61021e6102a2366004611a93565b610857565b6103506102b53660046119ff565b6040805160a081018252600080825260208201819052918101829052606081018290526080810191909152506001600160401b03908116600090815260026020908152604091829020825160a081018452905460ff8116825263ffffffff61010082041692820192909252600160281b8204841692810192909252600160681b810483166060830152600160a81b9004909116608082015290565b60405161028b9190611b60565b61039061036b3660046119ff565b6001600160401b03908116600090815260036020526040902054600160681b90041690565b6040516001600160401b03909116815260200161028b565b6000546103c290600160601b90046001600160a01b031681565b6040516001600160a01b03909116815260200161028b565b61027a6103e83660046119ff565b6001600160401b0316600090815260026020526040902054610100900463ffffffff1690565b6002610390565b61021e610423366004611bb9565b61099d565b61021e610436366004611bb9565b6109d9565b6104746104493660046119ff565b6001600160401b031660009081526003602090815260408083205460ff168352600190915290205490565b60405190815260200161028b565b61021e610490366004611be3565b610a15565b61021e610a29565b6104746104ab3660046119ff565b6001600160401b031660009081526002602090815260408083205460ff168352600190915290205490565b7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c199300546001600160a01b03166103c2565b610474610514366004611bfe565b610a3d565b610474620f424081565b61021e610531366004611c91565b610bc6565b6103906105443660046119ff565b6001600160401b03908116600090815260026020526040902054600160a81b90041690565b60005461057c906001600160601b031681565b6040516001600160601b03909116815260200161028b565b61021e6105a2366004611920565b610c02565b6105d36105b53660046119ff565b6001600160401b031660009081526002602052604090205460ff1690565b60405160ff909116815260200161028b565b61021e6105f3366004611bb9565b610c3e565b61021e610606366004611bb9565b610c7a565b6103506106193660046119ff565b6040805160a081018252600080825260208201819052918101829052606081018290526080810191909152506001600160401b03908116600090815260036020908152604091829020825160a081018452905460ff8116825263ffffffff61010082041692820192909252600160281b8204841692810192909252600160681b810483166060830152600160a81b9004909116608082015290565b61021e6106c2366004611cc6565b610cb6565b6105d36106d53660046119ff565b6001600160401b031660009081526003602052604090205460ff1690565b6103906107013660046119ff565b6001600160401b03908116600090815260036020526040902054600160a81b90041690565b6103906107343660046119ff565b6001600160401b03908116600090815260026020526040902054600160681b90041690565b61021e6107673660046119cc565b610cee565b61021e61077a366004611cc6565b610d2a565b61047461078d366004611ce1565b60016020526000908152604090205481565b600054600160601b90046001600160a01b031633146107d15760405163607e454560e11b815260040160405180910390fd5b6107db8282610d6a565b5050565b600054600160601b90046001600160a01b031633146108115760405163607e454560e11b815260040160405180910390fd5b6107db8282610f56565b600054600160601b90046001600160a01b0316331461084d5760405163607e454560e11b815260040160405180910390fd5b6107db8282610fee565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a008054600160401b810460ff1615906001600160401b031660008115801561089c5750825b90506000826001600160401b031660011480156108b85750303b155b9050811580156108c6575080155b156108e45760405163f92ee8a960e01b815260040160405180910390fd5b845467ffffffffffffffff19166001178555831561090e57845460ff60401b1916600160401b1785555b6109178e611083565b6109208d611094565b6109298c6110f1565b6109338b8b611144565b61093d8989610d6a565b610947878761132b565b831561098d57845460ff60401b19168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b5050505050505050505050505050565b600054600160601b90046001600160a01b031633146109cf5760405163607e454560e11b815260040160405180910390fd5b6107db828261141b565b600054600160601b90046001600160a01b03163314610a0b5760405163607e454560e11b815260040160405180910390fd5b6107db82826114dd565b610a1d61159f565b610a26816110f1565b50565b610a3161159f565b610a3b60006115fa565b565b6001600160401b0380851660009081526002602090815260408083208054600160a81b8104861685526003845282852060ff82168652600190945291842054939490938592620f424092610a999291600160681b900416611d12565b610aa39190611d29565b825460ff811660009081526001602052604081205492935091620f424091610adb9190600160681b90046001600160401b0316611d12565b610ae59190611d29565b905081600003610b0857604051633532119760e11b815260040160405180910390fd5b80600003610b2957604051633532119760e11b815260040160405180910390fd5b8254600090600160a81b81046001600160401b031690610b55908a90610100900463ffffffff16611d4b565b610b5f9190611d12565b9050610b6b8282611d12565b85548490610b85908a90610100900463ffffffff16611d5e565b6001600160401b0316610b989190611d12565b600054610bae91906001600160601b0316611d4b565b610bb89190611d4b565b9a9950505050505050505050565b600054600160601b90046001600160a01b03163314610bf85760405163607e454560e11b815260040160405180910390fd5b6107db828261132b565b600054600160601b90046001600160a01b03163314610c345760405163607e454560e11b815260040160405180910390fd5b6107db8282611144565b600054600160601b90046001600160a01b03163314610c705760405163607e454560e11b815260040160405180910390fd5b6107db828261166b565b600054600160601b90046001600160a01b03163314610cac5760405163607e454560e11b815260040160405180910390fd5b6107db828261172d565b610cbe61159f565b6001600160a01b038116610ce55760405163d92e233d60e01b815260040160405180910390fd5b610a2681611094565b600054600160601b90046001600160a01b03163314610d205760405163607e454560e11b815260040160405180910390fd5b6107db82826117ef565b610d3261159f565b6001600160a01b038116610d6157604051631e4fbdf760e01b8152600060048201526024015b60405180910390fd5b610a26816115fa565b60005b81811015610f51576000838383818110610d8957610d89611d85565b905060a00201803603810190610d9f9190611e39565b805190915060ff16600003610dc7576040516350614df960e01b815260040160405180910390fd5b80604001516001600160401b0316600003610df557604051630c26851b60e11b815260040160405180910390fd5b80606001516001600160401b0316600003610e2357604051630e661aed60e41b815260040160405180910390fd5b80608001516001600160401b0316600003610e51576040516348cfc33560e11b815260040160405180910390fd5b604081810180516001600160401b0390811660009081526003602090815290849020855181548388015195516060808a01516080808c015160ff90961664ffffffffff19909516851761010063ffffffff909b169a909a029990991765010000000000600160a81b031916600160281b93891693840267ffffffffffffffff60681b191617600160681b9189169182021767ffffffffffffffff60a81b1916600160a81b959098169485029790971790945587519182529381019390935294820192909252908101929092527fdc963752c4cbfdae934dd4c5abf689a81f0e8f2eb66558af02d252abc9cec7a9910160405180910390a150600101610d6d565b505050565b80600003610f775760405163fa90419960e01b815260040160405180910390fd5b8160ff16600003610f9b576040516350614df960e01b815260040160405180910390fd5b60ff8216600081815260016020908152604091829020849055815192835282018390527fd665ed89605ad8a805fa330621c3a15c9e6de67d51dbc52a478a868e1ed616e491015b60405180910390a15050565b816001600160401b031660000361101857604051630c26851b60e11b815260040160405180910390fd5b6001600160401b038216600081815260036020908152604091829020805464ffffffff00191661010063ffffffff8716908102919091179091558251938452908301527f9f1d6f43f208ea6217bdb6274ed9ade43967929063eaa2f16274565d3af986e29101610fe2565b61108b611884565b610a26816118cd565b600080546001600160601b0316600160601b6001600160a01b038416908102919091179091556040519081527f60a0f5b9f9e81e98216071b85826681c796256fe3d1354ecb675580fba64fa69906020015b60405180910390a150565b600080546bffffffffffffffffffffffff19166001600160601b0383169081179091556040519081527fd91752439e358587fc0828ed743df5939f16a918de501834bd954d03be15c959906020016110e6565b60005b81811015610f5157600083838381811061116357611163611d85565b905060a002018036038101906111799190611e39565b805190915060ff166000036111a1576040516350614df960e01b815260040160405180910390fd5b80604001516001600160401b03166000036111cf57604051633212217560e21b815260040160405180910390fd5b80606001516001600160401b03166000036111fd57604051630e661aed60e41b815260040160405180910390fd5b80608001516001600160401b031660000361122b57604051630c26851b60e11b815260040160405180910390fd5b604081810180516001600160401b0390811660009081526002602090815290849020855181548388015195516060808a01516080808c015160ff90961664ffffffffff19909516851761010063ffffffff909b169a909a029990991765010000000000600160a81b031916600160281b93891693840267ffffffffffffffff60681b191617600160681b9189169182021767ffffffffffffffff60a81b1916600160a81b959098169485029790971790945587519182529381019390935294820192909252908101929092527f4928444fd66b441b3ace1058b261e66af8a0193f1e0533d23ed48d0270075551910160405180910390a150600101611147565b60005b81811015610f5157600083838381811061134a5761134a611d85565b9050604002018036038101906113609190611e55565b805190915060ff16600003611388576040516350614df960e01b815260040160405180910390fd5b80602001516000036113ad5760405163fa90419960e01b815260040160405180910390fd5b60208082018051835160ff1660009081526001909352604092839020558251905191517fd665ed89605ad8a805fa330621c3a15c9e6de67d51dbc52a478a868e1ed616e49261140a929160ff929092168252602082015260400190565b60405180910390a15060010161132e565b816001600160401b031660000361144557604051630c26851b60e11b815260040160405180910390fd5b806001600160401b031660000361146f576040516348cfc33560e11b815260040160405180910390fd5b6001600160401b03828116600081815260036020908152604091829020805467ffffffffffffffff60a81b1916600160a81b95871695860217905581519283528201929092527f9e8c8606adb2b50f48cb69ccb1c3e349e6046ed63765e5931c01f4fe6aacd5e99101610fe2565b806001600160401b031660000361150757604051630e661aed60e41b815260040160405180910390fd5b816001600160401b031660000361153157604051630c26851b60e11b815260040160405180910390fd5b6001600160401b03828116600081815260036020908152604091829020805467ffffffffffffffff60681b1916600160681b95871695860217905581519283528201929092527fd7d8dd5a956a8bd500e02d52d0a9dd8a0e2955ec48771a8c9da485e6706c66fb9101610fe2565b336115d17f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c199300546001600160a01b031690565b6001600160a01b031614610a3b5760405163118cdaa760e01b8152336004820152602401610d58565b7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c19930080546001600160a01b031981166001600160a01b03848116918217845560405192169182907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a3505050565b806001600160401b031660000361169557604051630e661aed60e41b815260040160405180910390fd5b816001600160401b03166000036116bf57604051633212217560e21b815260040160405180910390fd5b6001600160401b03828116600081815260026020908152604091829020805467ffffffffffffffff60681b1916600160681b95871695860217905581519283528201929092527fe0e5abb8929e27a69d77f47a4e3f9575411a5be1fa596e5b55078d7850f358db9101610fe2565b816001600160401b031660000361175757604051633212217560e21b815260040160405180910390fd5b806001600160401b031660000361178157604051630c26851b60e11b815260040160405180910390fd5b6001600160401b03828116600081815260026020908152604091829020805467ffffffffffffffff60a81b1916600160a81b95871695860217905581519283528201929092527f0a5853014cbdb5103840fd3b7fcd886e7a93ef446d8c8707a5269d25ed32b4fe9101610fe2565b816001600160401b031660000361181957604051633212217560e21b815260040160405180910390fd5b6001600160401b038216600081815260026020908152604091829020805464ffffffff00191661010063ffffffff8716908102919091179091558251938452908301527f525a2cd9c1093178959cb9c72fe00c6be026fc953e93c9bc789d2176da98c4059101610fe2565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a0054600160401b900460ff16610a3b57604051631afcd79f60e31b815260040160405180910390fd5b610d32611884565b60008083601f8401126118e757600080fd5b5081356001600160401b038111156118fe57600080fd5b60208301915083602060a08302850101111561191957600080fd5b9250929050565b6000806020838503121561193357600080fd5b82356001600160401b0381111561194957600080fd5b611955858286016118d5565b90969095509350505050565b803560ff8116811461197257600080fd5b919050565b6000806040838503121561198a57600080fd5b61199383611961565b946020939093013593505050565b80356001600160401b038116811461197257600080fd5b803563ffffffff8116811461197257600080fd5b600080604083850312156119df57600080fd5b6119e8836119a1565b91506119f6602084016119b8565b90509250929050565b600060208284031215611a1157600080fd5b611a1a826119a1565b9392505050565b80356001600160a01b038116811461197257600080fd5b80356001600160601b038116811461197257600080fd5b60008083601f840112611a6157600080fd5b5081356001600160401b03811115611a7857600080fd5b6020830191508360208260061b850101111561191957600080fd5b600080600080600080600080600060c08a8c031215611ab157600080fd5b611aba8a611a21565b9850611ac860208b01611a21565b9750611ad660408b01611a38565b965060608a01356001600160401b0380821115611af257600080fd5b611afe8d838e016118d5565b909850965060808c0135915080821115611b1757600080fd5b611b238d838e016118d5565b909650945060a08c0135915080821115611b3c57600080fd5b50611b498c828d01611a4f565b915080935050809150509295985092959850929598565b60a08101611bb3828460ff815116825263ffffffff602082015116602083015260408101516001600160401b03808216604085015280606084015116606085015280608084015116608085015250505050565b92915050565b60008060408385031215611bcc57600080fd5b611bd5836119a1565b91506119f6602084016119a1565b600060208284031215611bf557600080fd5b611a1a82611a38565b60008060008060608587031215611c1457600080fd5b611c1d856119a1565b935060208501356001600160401b0380821115611c3957600080fd5b818701915087601f830112611c4d57600080fd5b813581811115611c5c57600080fd5b886020828501011115611c6e57600080fd5b602083019550809450505050611c86604086016119a1565b905092959194509250565b60008060208385031215611ca457600080fd5b82356001600160401b03811115611cba57600080fd5b61195585828601611a4f565b600060208284031215611cd857600080fd5b611a1a82611a21565b600060208284031215611cf357600080fd5b611a1a82611961565b634e487b7160e01b600052601160045260246000fd5b8082028115828204841417611bb357611bb3611cfc565b600082611d4657634e487b7160e01b600052601260045260246000fd5b500490565b80820180821115611bb357611bb3611cfc565b6001600160401b03818116838216019080821115611d7e57611d7e611cfc565b5092915050565b634e487b7160e01b600052603260045260246000fd5b600060a08284031215611dad57600080fd5b60405160a081018181106001600160401b0382111715611ddd57634e487b7160e01b600052604160045260246000fd5b604052905080611dec83611961565b8152611dfa602084016119b8565b6020820152611e0b604084016119a1565b6040820152611e1c606084016119a1565b6060820152611e2d608084016119a1565b60808201525092915050565b600060a08284031215611e4b57600080fd5b611a1a8383611d9b565b600060408284031215611e6757600080fd5b604051604081018181106001600160401b0382111715611e9757634e487b7160e01b600052604160045260246000fd5b604052611ea383611961565b815260208301356020820152809150509291505056fea26469706673582212208f71023df533b8fb316ad48c577646337d057f519e52734bb9811fa5ecffc21064736f6c63430008180033",
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

// BaseDataBuffer is a free data retrieval call binding the contract method 0x11018853.
//
// Solidity: function baseDataBuffer(uint64 dataCostId) view returns(uint32)
func (_FeeOracleV2 *FeeOracleV2Caller) BaseDataBuffer(opts *bind.CallOpts, dataCostId uint64) (uint32, error) {
	var out []interface{}
	err := _FeeOracleV2.contract.Call(opts, &out, "baseDataBuffer", dataCostId)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// BaseDataBuffer is a free data retrieval call binding the contract method 0x11018853.
//
// Solidity: function baseDataBuffer(uint64 dataCostId) view returns(uint32)
func (_FeeOracleV2 *FeeOracleV2Session) BaseDataBuffer(dataCostId uint64) (uint32, error) {
	return _FeeOracleV2.Contract.BaseDataBuffer(&_FeeOracleV2.CallOpts, dataCostId)
}

// BaseDataBuffer is a free data retrieval call binding the contract method 0x11018853.
//
// Solidity: function baseDataBuffer(uint64 dataCostId) view returns(uint32)
func (_FeeOracleV2 *FeeOracleV2CallerSession) BaseDataBuffer(dataCostId uint64) (uint32, error) {
	return _FeeOracleV2.Contract.BaseDataBuffer(&_FeeOracleV2.CallOpts, dataCostId)
}

// BaseGasLimit is a free data retrieval call binding the contract method 0x50b81539.
//
// Solidity: function baseGasLimit(uint64 chainId) view returns(uint32)
func (_FeeOracleV2 *FeeOracleV2Caller) BaseGasLimit(opts *bind.CallOpts, chainId uint64) (uint32, error) {
	var out []interface{}
	err := _FeeOracleV2.contract.Call(opts, &out, "baseGasLimit", chainId)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// BaseGasLimit is a free data retrieval call binding the contract method 0x50b81539.
//
// Solidity: function baseGasLimit(uint64 chainId) view returns(uint32)
func (_FeeOracleV2 *FeeOracleV2Session) BaseGasLimit(chainId uint64) (uint32, error) {
	return _FeeOracleV2.Contract.BaseGasLimit(&_FeeOracleV2.CallOpts, chainId)
}

// BaseGasLimit is a free data retrieval call binding the contract method 0x50b81539.
//
// Solidity: function baseGasLimit(uint64 chainId) view returns(uint32)
func (_FeeOracleV2 *FeeOracleV2CallerSession) BaseGasLimit(chainId uint64) (uint32, error) {
	return _FeeOracleV2.Contract.BaseGasLimit(&_FeeOracleV2.CallOpts, chainId)
}

// DataCostParams is a free data retrieval call binding the contract method 0xbfc71416.
//
// Solidity: function dataCostParams(uint64 dataCostId) view returns((uint8,uint32,uint64,uint64,uint64))
func (_FeeOracleV2 *FeeOracleV2Caller) DataCostParams(opts *bind.CallOpts, dataCostId uint64) (IFeeOracleV2DataCostParams, error) {
	var out []interface{}
	err := _FeeOracleV2.contract.Call(opts, &out, "dataCostParams", dataCostId)

	if err != nil {
		return *new(IFeeOracleV2DataCostParams), err
	}

	out0 := *abi.ConvertType(out[0], new(IFeeOracleV2DataCostParams)).(*IFeeOracleV2DataCostParams)

	return out0, err

}

// DataCostParams is a free data retrieval call binding the contract method 0xbfc71416.
//
// Solidity: function dataCostParams(uint64 dataCostId) view returns((uint8,uint32,uint64,uint64,uint64))
func (_FeeOracleV2 *FeeOracleV2Session) DataCostParams(dataCostId uint64) (IFeeOracleV2DataCostParams, error) {
	return _FeeOracleV2.Contract.DataCostParams(&_FeeOracleV2.CallOpts, dataCostId)
}

// DataCostParams is a free data retrieval call binding the contract method 0xbfc71416.
//
// Solidity: function dataCostParams(uint64 dataCostId) view returns((uint8,uint32,uint64,uint64,uint64))
func (_FeeOracleV2 *FeeOracleV2CallerSession) DataCostParams(dataCostId uint64) (IFeeOracleV2DataCostParams, error) {
	return _FeeOracleV2.Contract.DataCostParams(&_FeeOracleV2.CallOpts, dataCostId)
}

// DataGasPerByte is a free data retrieval call binding the contract method 0xdb0018e8.
//
// Solidity: function dataGasPerByte(uint64 dataCostId) view returns(uint64)
func (_FeeOracleV2 *FeeOracleV2Caller) DataGasPerByte(opts *bind.CallOpts, dataCostId uint64) (uint64, error) {
	var out []interface{}
	err := _FeeOracleV2.contract.Call(opts, &out, "dataGasPerByte", dataCostId)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// DataGasPerByte is a free data retrieval call binding the contract method 0xdb0018e8.
//
// Solidity: function dataGasPerByte(uint64 dataCostId) view returns(uint64)
func (_FeeOracleV2 *FeeOracleV2Session) DataGasPerByte(dataCostId uint64) (uint64, error) {
	return _FeeOracleV2.Contract.DataGasPerByte(&_FeeOracleV2.CallOpts, dataCostId)
}

// DataGasPerByte is a free data retrieval call binding the contract method 0xdb0018e8.
//
// Solidity: function dataGasPerByte(uint64 dataCostId) view returns(uint64)
func (_FeeOracleV2 *FeeOracleV2CallerSession) DataGasPerByte(dataCostId uint64) (uint64, error) {
	return _FeeOracleV2.Contract.DataGasPerByte(&_FeeOracleV2.CallOpts, dataCostId)
}

// DataGasPrice is a free data retrieval call binding the contract method 0x415070af.
//
// Solidity: function dataGasPrice(uint64 dataCostId) view returns(uint64)
func (_FeeOracleV2 *FeeOracleV2Caller) DataGasPrice(opts *bind.CallOpts, dataCostId uint64) (uint64, error) {
	var out []interface{}
	err := _FeeOracleV2.contract.Call(opts, &out, "dataGasPrice", dataCostId)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// DataGasPrice is a free data retrieval call binding the contract method 0x415070af.
//
// Solidity: function dataGasPrice(uint64 dataCostId) view returns(uint64)
func (_FeeOracleV2 *FeeOracleV2Session) DataGasPrice(dataCostId uint64) (uint64, error) {
	return _FeeOracleV2.Contract.DataGasPrice(&_FeeOracleV2.CallOpts, dataCostId)
}

// DataGasPrice is a free data retrieval call binding the contract method 0x415070af.
//
// Solidity: function dataGasPrice(uint64 dataCostId) view returns(uint64)
func (_FeeOracleV2 *FeeOracleV2CallerSession) DataGasPrice(dataCostId uint64) (uint64, error) {
	return _FeeOracleV2.Contract.DataGasPrice(&_FeeOracleV2.CallOpts, dataCostId)
}

// DataGasToken is a free data retrieval call binding the contract method 0xd32b68ad.
//
// Solidity: function dataGasToken(uint64 dataCostId) view returns(uint8)
func (_FeeOracleV2 *FeeOracleV2Caller) DataGasToken(opts *bind.CallOpts, dataCostId uint64) (uint8, error) {
	var out []interface{}
	err := _FeeOracleV2.contract.Call(opts, &out, "dataGasToken", dataCostId)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// DataGasToken is a free data retrieval call binding the contract method 0xd32b68ad.
//
// Solidity: function dataGasToken(uint64 dataCostId) view returns(uint8)
func (_FeeOracleV2 *FeeOracleV2Session) DataGasToken(dataCostId uint64) (uint8, error) {
	return _FeeOracleV2.Contract.DataGasToken(&_FeeOracleV2.CallOpts, dataCostId)
}

// DataGasToken is a free data retrieval call binding the contract method 0xd32b68ad.
//
// Solidity: function dataGasToken(uint64 dataCostId) view returns(uint8)
func (_FeeOracleV2 *FeeOracleV2CallerSession) DataGasToken(dataCostId uint64) (uint8, error) {
	return _FeeOracleV2.Contract.DataGasToken(&_FeeOracleV2.CallOpts, dataCostId)
}

// ExecDataCostId is a free data retrieval call binding the contract method 0x9a5551c3.
//
// Solidity: function execDataCostId(uint64 chainId) view returns(uint64)
func (_FeeOracleV2 *FeeOracleV2Caller) ExecDataCostId(opts *bind.CallOpts, chainId uint64) (uint64, error) {
	var out []interface{}
	err := _FeeOracleV2.contract.Call(opts, &out, "execDataCostId", chainId)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// ExecDataCostId is a free data retrieval call binding the contract method 0x9a5551c3.
//
// Solidity: function execDataCostId(uint64 chainId) view returns(uint64)
func (_FeeOracleV2 *FeeOracleV2Session) ExecDataCostId(chainId uint64) (uint64, error) {
	return _FeeOracleV2.Contract.ExecDataCostId(&_FeeOracleV2.CallOpts, chainId)
}

// ExecDataCostId is a free data retrieval call binding the contract method 0x9a5551c3.
//
// Solidity: function execDataCostId(uint64 chainId) view returns(uint64)
func (_FeeOracleV2 *FeeOracleV2CallerSession) ExecDataCostId(chainId uint64) (uint64, error) {
	return _FeeOracleV2.Contract.ExecDataCostId(&_FeeOracleV2.CallOpts, chainId)
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

// ExecGasToken is a free data retrieval call binding the contract method 0xb984cc0b.
//
// Solidity: function execGasToken(uint64 chainId) view returns(uint8)
func (_FeeOracleV2 *FeeOracleV2Caller) ExecGasToken(opts *bind.CallOpts, chainId uint64) (uint8, error) {
	var out []interface{}
	err := _FeeOracleV2.contract.Call(opts, &out, "execGasToken", chainId)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// ExecGasToken is a free data retrieval call binding the contract method 0xb984cc0b.
//
// Solidity: function execGasToken(uint64 chainId) view returns(uint8)
func (_FeeOracleV2 *FeeOracleV2Session) ExecGasToken(chainId uint64) (uint8, error) {
	return _FeeOracleV2.Contract.ExecGasToken(&_FeeOracleV2.CallOpts, chainId)
}

// ExecGasToken is a free data retrieval call binding the contract method 0xb984cc0b.
//
// Solidity: function execGasToken(uint64 chainId) view returns(uint8)
func (_FeeOracleV2 *FeeOracleV2CallerSession) ExecGasToken(chainId uint64) (uint8, error) {
	return _FeeOracleV2.Contract.ExecGasToken(&_FeeOracleV2.CallOpts, chainId)
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
// Solidity: function feeParams(uint64 chainId) view returns((uint8,uint32,uint64,uint64,uint64))
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
// Solidity: function feeParams(uint64 chainId) view returns((uint8,uint32,uint64,uint64,uint64))
func (_FeeOracleV2 *FeeOracleV2Session) FeeParams(chainId uint64) (IFeeOracleV2FeeParams, error) {
	return _FeeOracleV2.Contract.FeeParams(&_FeeOracleV2.CallOpts, chainId)
}

// FeeParams is a free data retrieval call binding the contract method 0x2d4634a4.
//
// Solidity: function feeParams(uint64 chainId) view returns((uint8,uint32,uint64,uint64,uint64))
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
// Solidity: function protocolFee() view returns(uint96)
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
// Solidity: function protocolFee() view returns(uint96)
func (_FeeOracleV2 *FeeOracleV2Session) ProtocolFee() (*big.Int, error) {
	return _FeeOracleV2.Contract.ProtocolFee(&_FeeOracleV2.CallOpts)
}

// ProtocolFee is a free data retrieval call binding the contract method 0xb0e21e8a.
//
// Solidity: function protocolFee() view returns(uint96)
func (_FeeOracleV2 *FeeOracleV2CallerSession) ProtocolFee() (*big.Int, error) {
	return _FeeOracleV2.Contract.ProtocolFee(&_FeeOracleV2.CallOpts)
}

// ToNativeRate is a free data retrieval call binding the contract method 0x8b7bfd70.
//
// Solidity: function toNativeRate(uint64 chainId) view returns(uint256)
func (_FeeOracleV2 *FeeOracleV2Caller) ToNativeRate(opts *bind.CallOpts, chainId uint64) (*big.Int, error) {
	var out []interface{}
	err := _FeeOracleV2.contract.Call(opts, &out, "toNativeRate", chainId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ToNativeRate is a free data retrieval call binding the contract method 0x8b7bfd70.
//
// Solidity: function toNativeRate(uint64 chainId) view returns(uint256)
func (_FeeOracleV2 *FeeOracleV2Session) ToNativeRate(chainId uint64) (*big.Int, error) {
	return _FeeOracleV2.Contract.ToNativeRate(&_FeeOracleV2.CallOpts, chainId)
}

// ToNativeRate is a free data retrieval call binding the contract method 0x8b7bfd70.
//
// Solidity: function toNativeRate(uint64 chainId) view returns(uint256)
func (_FeeOracleV2 *FeeOracleV2CallerSession) ToNativeRate(chainId uint64) (*big.Int, error) {
	return _FeeOracleV2.Contract.ToNativeRate(&_FeeOracleV2.CallOpts, chainId)
}

// ToNativeRateData is a free data retrieval call binding the contract method 0x67140232.
//
// Solidity: function toNativeRateData(uint64 dataCostId) view returns(uint256)
func (_FeeOracleV2 *FeeOracleV2Caller) ToNativeRateData(opts *bind.CallOpts, dataCostId uint64) (*big.Int, error) {
	var out []interface{}
	err := _FeeOracleV2.contract.Call(opts, &out, "toNativeRateData", dataCostId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ToNativeRateData is a free data retrieval call binding the contract method 0x67140232.
//
// Solidity: function toNativeRateData(uint64 dataCostId) view returns(uint256)
func (_FeeOracleV2 *FeeOracleV2Session) ToNativeRateData(dataCostId uint64) (*big.Int, error) {
	return _FeeOracleV2.Contract.ToNativeRateData(&_FeeOracleV2.CallOpts, dataCostId)
}

// ToNativeRateData is a free data retrieval call binding the contract method 0x67140232.
//
// Solidity: function toNativeRateData(uint64 dataCostId) view returns(uint256)
func (_FeeOracleV2 *FeeOracleV2CallerSession) ToNativeRateData(dataCostId uint64) (*big.Int, error) {
	return _FeeOracleV2.Contract.ToNativeRateData(&_FeeOracleV2.CallOpts, dataCostId)
}

// TokenToNativeRate is a free data retrieval call binding the contract method 0xfee86f69.
//
// Solidity: function tokenToNativeRate(uint8 gasToken) view returns(uint256)
func (_FeeOracleV2 *FeeOracleV2Caller) TokenToNativeRate(opts *bind.CallOpts, gasToken uint8) (*big.Int, error) {
	var out []interface{}
	err := _FeeOracleV2.contract.Call(opts, &out, "tokenToNativeRate", gasToken)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenToNativeRate is a free data retrieval call binding the contract method 0xfee86f69.
//
// Solidity: function tokenToNativeRate(uint8 gasToken) view returns(uint256)
func (_FeeOracleV2 *FeeOracleV2Session) TokenToNativeRate(gasToken uint8) (*big.Int, error) {
	return _FeeOracleV2.Contract.TokenToNativeRate(&_FeeOracleV2.CallOpts, gasToken)
}

// TokenToNativeRate is a free data retrieval call binding the contract method 0xfee86f69.
//
// Solidity: function tokenToNativeRate(uint8 gasToken) view returns(uint256)
func (_FeeOracleV2 *FeeOracleV2CallerSession) TokenToNativeRate(gasToken uint8) (*big.Int, error) {
	return _FeeOracleV2.Contract.TokenToNativeRate(&_FeeOracleV2.CallOpts, gasToken)
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

// BulkSetDataCostParams is a paid mutator transaction binding the contract method 0x04e53a15.
//
// Solidity: function bulkSetDataCostParams((uint8,uint32,uint64,uint64,uint64)[] params) returns()
func (_FeeOracleV2 *FeeOracleV2Transactor) BulkSetDataCostParams(opts *bind.TransactOpts, params []IFeeOracleV2DataCostParams) (*types.Transaction, error) {
	return _FeeOracleV2.contract.Transact(opts, "bulkSetDataCostParams", params)
}

// BulkSetDataCostParams is a paid mutator transaction binding the contract method 0x04e53a15.
//
// Solidity: function bulkSetDataCostParams((uint8,uint32,uint64,uint64,uint64)[] params) returns()
func (_FeeOracleV2 *FeeOracleV2Session) BulkSetDataCostParams(params []IFeeOracleV2DataCostParams) (*types.Transaction, error) {
	return _FeeOracleV2.Contract.BulkSetDataCostParams(&_FeeOracleV2.TransactOpts, params)
}

// BulkSetDataCostParams is a paid mutator transaction binding the contract method 0x04e53a15.
//
// Solidity: function bulkSetDataCostParams((uint8,uint32,uint64,uint64,uint64)[] params) returns()
func (_FeeOracleV2 *FeeOracleV2TransactorSession) BulkSetDataCostParams(params []IFeeOracleV2DataCostParams) (*types.Transaction, error) {
	return _FeeOracleV2.Contract.BulkSetDataCostParams(&_FeeOracleV2.TransactOpts, params)
}

// BulkSetFeeParams is a paid mutator transaction binding the contract method 0xb15268b0.
//
// Solidity: function bulkSetFeeParams((uint8,uint32,uint64,uint64,uint64)[] params) returns()
func (_FeeOracleV2 *FeeOracleV2Transactor) BulkSetFeeParams(opts *bind.TransactOpts, params []IFeeOracleV2FeeParams) (*types.Transaction, error) {
	return _FeeOracleV2.contract.Transact(opts, "bulkSetFeeParams", params)
}

// BulkSetFeeParams is a paid mutator transaction binding the contract method 0xb15268b0.
//
// Solidity: function bulkSetFeeParams((uint8,uint32,uint64,uint64,uint64)[] params) returns()
func (_FeeOracleV2 *FeeOracleV2Session) BulkSetFeeParams(params []IFeeOracleV2FeeParams) (*types.Transaction, error) {
	return _FeeOracleV2.Contract.BulkSetFeeParams(&_FeeOracleV2.TransactOpts, params)
}

// BulkSetFeeParams is a paid mutator transaction binding the contract method 0xb15268b0.
//
// Solidity: function bulkSetFeeParams((uint8,uint32,uint64,uint64,uint64)[] params) returns()
func (_FeeOracleV2 *FeeOracleV2TransactorSession) BulkSetFeeParams(params []IFeeOracleV2FeeParams) (*types.Transaction, error) {
	return _FeeOracleV2.Contract.BulkSetFeeParams(&_FeeOracleV2.TransactOpts, params)
}

// BulkSetToNativeRate is a paid mutator transaction binding the contract method 0x8ff83049.
//
// Solidity: function bulkSetToNativeRate((uint8,uint256)[] params) returns()
func (_FeeOracleV2 *FeeOracleV2Transactor) BulkSetToNativeRate(opts *bind.TransactOpts, params []IFeeOracleV2NativeRateParams) (*types.Transaction, error) {
	return _FeeOracleV2.contract.Transact(opts, "bulkSetToNativeRate", params)
}

// BulkSetToNativeRate is a paid mutator transaction binding the contract method 0x8ff83049.
//
// Solidity: function bulkSetToNativeRate((uint8,uint256)[] params) returns()
func (_FeeOracleV2 *FeeOracleV2Session) BulkSetToNativeRate(params []IFeeOracleV2NativeRateParams) (*types.Transaction, error) {
	return _FeeOracleV2.Contract.BulkSetToNativeRate(&_FeeOracleV2.TransactOpts, params)
}

// BulkSetToNativeRate is a paid mutator transaction binding the contract method 0x8ff83049.
//
// Solidity: function bulkSetToNativeRate((uint8,uint256)[] params) returns()
func (_FeeOracleV2 *FeeOracleV2TransactorSession) BulkSetToNativeRate(params []IFeeOracleV2NativeRateParams) (*types.Transaction, error) {
	return _FeeOracleV2.Contract.BulkSetToNativeRate(&_FeeOracleV2.TransactOpts, params)
}

// Initialize is a paid mutator transaction binding the contract method 0x176c73e3.
//
// Solidity: function initialize(address owner_, address manager_, uint96 protocolFee_, (uint8,uint32,uint64,uint64,uint64)[] feeParams_, (uint8,uint32,uint64,uint64,uint64)[] dataCostParams_, (uint8,uint256)[] nativeRateParams_) returns()
func (_FeeOracleV2 *FeeOracleV2Transactor) Initialize(opts *bind.TransactOpts, owner_ common.Address, manager_ common.Address, protocolFee_ *big.Int, feeParams_ []IFeeOracleV2FeeParams, dataCostParams_ []IFeeOracleV2DataCostParams, nativeRateParams_ []IFeeOracleV2NativeRateParams) (*types.Transaction, error) {
	return _FeeOracleV2.contract.Transact(opts, "initialize", owner_, manager_, protocolFee_, feeParams_, dataCostParams_, nativeRateParams_)
}

// Initialize is a paid mutator transaction binding the contract method 0x176c73e3.
//
// Solidity: function initialize(address owner_, address manager_, uint96 protocolFee_, (uint8,uint32,uint64,uint64,uint64)[] feeParams_, (uint8,uint32,uint64,uint64,uint64)[] dataCostParams_, (uint8,uint256)[] nativeRateParams_) returns()
func (_FeeOracleV2 *FeeOracleV2Session) Initialize(owner_ common.Address, manager_ common.Address, protocolFee_ *big.Int, feeParams_ []IFeeOracleV2FeeParams, dataCostParams_ []IFeeOracleV2DataCostParams, nativeRateParams_ []IFeeOracleV2NativeRateParams) (*types.Transaction, error) {
	return _FeeOracleV2.Contract.Initialize(&_FeeOracleV2.TransactOpts, owner_, manager_, protocolFee_, feeParams_, dataCostParams_, nativeRateParams_)
}

// Initialize is a paid mutator transaction binding the contract method 0x176c73e3.
//
// Solidity: function initialize(address owner_, address manager_, uint96 protocolFee_, (uint8,uint32,uint64,uint64,uint64)[] feeParams_, (uint8,uint32,uint64,uint64,uint64)[] dataCostParams_, (uint8,uint256)[] nativeRateParams_) returns()
func (_FeeOracleV2 *FeeOracleV2TransactorSession) Initialize(owner_ common.Address, manager_ common.Address, protocolFee_ *big.Int, feeParams_ []IFeeOracleV2FeeParams, dataCostParams_ []IFeeOracleV2DataCostParams, nativeRateParams_ []IFeeOracleV2NativeRateParams) (*types.Transaction, error) {
	return _FeeOracleV2.Contract.Initialize(&_FeeOracleV2.TransactOpts, owner_, manager_, protocolFee_, feeParams_, dataCostParams_, nativeRateParams_)
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

// SetBaseDataBuffer is a paid mutator transaction binding the contract method 0x0f0b435b.
//
// Solidity: function setBaseDataBuffer(uint64 dataCostId, uint32 newBaseDataBuffer) returns()
func (_FeeOracleV2 *FeeOracleV2Transactor) SetBaseDataBuffer(opts *bind.TransactOpts, dataCostId uint64, newBaseDataBuffer uint32) (*types.Transaction, error) {
	return _FeeOracleV2.contract.Transact(opts, "setBaseDataBuffer", dataCostId, newBaseDataBuffer)
}

// SetBaseDataBuffer is a paid mutator transaction binding the contract method 0x0f0b435b.
//
// Solidity: function setBaseDataBuffer(uint64 dataCostId, uint32 newBaseDataBuffer) returns()
func (_FeeOracleV2 *FeeOracleV2Session) SetBaseDataBuffer(dataCostId uint64, newBaseDataBuffer uint32) (*types.Transaction, error) {
	return _FeeOracleV2.Contract.SetBaseDataBuffer(&_FeeOracleV2.TransactOpts, dataCostId, newBaseDataBuffer)
}

// SetBaseDataBuffer is a paid mutator transaction binding the contract method 0x0f0b435b.
//
// Solidity: function setBaseDataBuffer(uint64 dataCostId, uint32 newBaseDataBuffer) returns()
func (_FeeOracleV2 *FeeOracleV2TransactorSession) SetBaseDataBuffer(dataCostId uint64, newBaseDataBuffer uint32) (*types.Transaction, error) {
	return _FeeOracleV2.Contract.SetBaseDataBuffer(&_FeeOracleV2.TransactOpts, dataCostId, newBaseDataBuffer)
}

// SetBaseGasLimit is a paid mutator transaction binding the contract method 0xe5737f62.
//
// Solidity: function setBaseGasLimit(uint64 chainId, uint32 newBaseGasLimit) returns()
func (_FeeOracleV2 *FeeOracleV2Transactor) SetBaseGasLimit(opts *bind.TransactOpts, chainId uint64, newBaseGasLimit uint32) (*types.Transaction, error) {
	return _FeeOracleV2.contract.Transact(opts, "setBaseGasLimit", chainId, newBaseGasLimit)
}

// SetBaseGasLimit is a paid mutator transaction binding the contract method 0xe5737f62.
//
// Solidity: function setBaseGasLimit(uint64 chainId, uint32 newBaseGasLimit) returns()
func (_FeeOracleV2 *FeeOracleV2Session) SetBaseGasLimit(chainId uint64, newBaseGasLimit uint32) (*types.Transaction, error) {
	return _FeeOracleV2.Contract.SetBaseGasLimit(&_FeeOracleV2.TransactOpts, chainId, newBaseGasLimit)
}

// SetBaseGasLimit is a paid mutator transaction binding the contract method 0xe5737f62.
//
// Solidity: function setBaseGasLimit(uint64 chainId, uint32 newBaseGasLimit) returns()
func (_FeeOracleV2 *FeeOracleV2TransactorSession) SetBaseGasLimit(chainId uint64, newBaseGasLimit uint32) (*types.Transaction, error) {
	return _FeeOracleV2.Contract.SetBaseGasLimit(&_FeeOracleV2.TransactOpts, chainId, newBaseGasLimit)
}

// SetDataCostId is a paid mutator transaction binding the contract method 0xbc51bf37.
//
// Solidity: function setDataCostId(uint64 chainId, uint64 dataCostId) returns()
func (_FeeOracleV2 *FeeOracleV2Transactor) SetDataCostId(opts *bind.TransactOpts, chainId uint64, dataCostId uint64) (*types.Transaction, error) {
	return _FeeOracleV2.contract.Transact(opts, "setDataCostId", chainId, dataCostId)
}

// SetDataCostId is a paid mutator transaction binding the contract method 0xbc51bf37.
//
// Solidity: function setDataCostId(uint64 chainId, uint64 dataCostId) returns()
func (_FeeOracleV2 *FeeOracleV2Session) SetDataCostId(chainId uint64, dataCostId uint64) (*types.Transaction, error) {
	return _FeeOracleV2.Contract.SetDataCostId(&_FeeOracleV2.TransactOpts, chainId, dataCostId)
}

// SetDataCostId is a paid mutator transaction binding the contract method 0xbc51bf37.
//
// Solidity: function setDataCostId(uint64 chainId, uint64 dataCostId) returns()
func (_FeeOracleV2 *FeeOracleV2TransactorSession) SetDataCostId(chainId uint64, dataCostId uint64) (*types.Transaction, error) {
	return _FeeOracleV2.Contract.SetDataCostId(&_FeeOracleV2.TransactOpts, chainId, dataCostId)
}

// SetDataGasPrice is a paid mutator transaction binding the contract method 0x5d3acee2.
//
// Solidity: function setDataGasPrice(uint64 dataCostId, uint64 gasPrice) returns()
func (_FeeOracleV2 *FeeOracleV2Transactor) SetDataGasPrice(opts *bind.TransactOpts, dataCostId uint64, gasPrice uint64) (*types.Transaction, error) {
	return _FeeOracleV2.contract.Transact(opts, "setDataGasPrice", dataCostId, gasPrice)
}

// SetDataGasPrice is a paid mutator transaction binding the contract method 0x5d3acee2.
//
// Solidity: function setDataGasPrice(uint64 dataCostId, uint64 gasPrice) returns()
func (_FeeOracleV2 *FeeOracleV2Session) SetDataGasPrice(dataCostId uint64, gasPrice uint64) (*types.Transaction, error) {
	return _FeeOracleV2.Contract.SetDataGasPrice(&_FeeOracleV2.TransactOpts, dataCostId, gasPrice)
}

// SetDataGasPrice is a paid mutator transaction binding the contract method 0x5d3acee2.
//
// Solidity: function setDataGasPrice(uint64 dataCostId, uint64 gasPrice) returns()
func (_FeeOracleV2 *FeeOracleV2TransactorSession) SetDataGasPrice(dataCostId uint64, gasPrice uint64) (*types.Transaction, error) {
	return _FeeOracleV2.Contract.SetDataGasPrice(&_FeeOracleV2.TransactOpts, dataCostId, gasPrice)
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

// SetGasPerByte is a paid mutator transaction binding the contract method 0x56bce459.
//
// Solidity: function setGasPerByte(uint64 dataCostId, uint64 gasPerByte) returns()
func (_FeeOracleV2 *FeeOracleV2Transactor) SetGasPerByte(opts *bind.TransactOpts, dataCostId uint64, gasPerByte uint64) (*types.Transaction, error) {
	return _FeeOracleV2.contract.Transact(opts, "setGasPerByte", dataCostId, gasPerByte)
}

// SetGasPerByte is a paid mutator transaction binding the contract method 0x56bce459.
//
// Solidity: function setGasPerByte(uint64 dataCostId, uint64 gasPerByte) returns()
func (_FeeOracleV2 *FeeOracleV2Session) SetGasPerByte(dataCostId uint64, gasPerByte uint64) (*types.Transaction, error) {
	return _FeeOracleV2.Contract.SetGasPerByte(&_FeeOracleV2.TransactOpts, dataCostId, gasPerByte)
}

// SetGasPerByte is a paid mutator transaction binding the contract method 0x56bce459.
//
// Solidity: function setGasPerByte(uint64 dataCostId, uint64 gasPerByte) returns()
func (_FeeOracleV2 *FeeOracleV2TransactorSession) SetGasPerByte(dataCostId uint64, gasPerByte uint64) (*types.Transaction, error) {
	return _FeeOracleV2.Contract.SetGasPerByte(&_FeeOracleV2.TransactOpts, dataCostId, gasPerByte)
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

// SetProtocolFee is a paid mutator transaction binding the contract method 0x6b6dccfe.
//
// Solidity: function setProtocolFee(uint96 fee) returns()
func (_FeeOracleV2 *FeeOracleV2Transactor) SetProtocolFee(opts *bind.TransactOpts, fee *big.Int) (*types.Transaction, error) {
	return _FeeOracleV2.contract.Transact(opts, "setProtocolFee", fee)
}

// SetProtocolFee is a paid mutator transaction binding the contract method 0x6b6dccfe.
//
// Solidity: function setProtocolFee(uint96 fee) returns()
func (_FeeOracleV2 *FeeOracleV2Session) SetProtocolFee(fee *big.Int) (*types.Transaction, error) {
	return _FeeOracleV2.Contract.SetProtocolFee(&_FeeOracleV2.TransactOpts, fee)
}

// SetProtocolFee is a paid mutator transaction binding the contract method 0x6b6dccfe.
//
// Solidity: function setProtocolFee(uint96 fee) returns()
func (_FeeOracleV2 *FeeOracleV2TransactorSession) SetProtocolFee(fee *big.Int) (*types.Transaction, error) {
	return _FeeOracleV2.Contract.SetProtocolFee(&_FeeOracleV2.TransactOpts, fee)
}

// SetToNativeRate is a paid mutator transaction binding the contract method 0x0baaa6aa.
//
// Solidity: function setToNativeRate(uint8 gasToken, uint256 nativeRate) returns()
func (_FeeOracleV2 *FeeOracleV2Transactor) SetToNativeRate(opts *bind.TransactOpts, gasToken uint8, nativeRate *big.Int) (*types.Transaction, error) {
	return _FeeOracleV2.contract.Transact(opts, "setToNativeRate", gasToken, nativeRate)
}

// SetToNativeRate is a paid mutator transaction binding the contract method 0x0baaa6aa.
//
// Solidity: function setToNativeRate(uint8 gasToken, uint256 nativeRate) returns()
func (_FeeOracleV2 *FeeOracleV2Session) SetToNativeRate(gasToken uint8, nativeRate *big.Int) (*types.Transaction, error) {
	return _FeeOracleV2.Contract.SetToNativeRate(&_FeeOracleV2.TransactOpts, gasToken, nativeRate)
}

// SetToNativeRate is a paid mutator transaction binding the contract method 0x0baaa6aa.
//
// Solidity: function setToNativeRate(uint8 gasToken, uint256 nativeRate) returns()
func (_FeeOracleV2 *FeeOracleV2TransactorSession) SetToNativeRate(gasToken uint8, nativeRate *big.Int) (*types.Transaction, error) {
	return _FeeOracleV2.Contract.SetToNativeRate(&_FeeOracleV2.TransactOpts, gasToken, nativeRate)
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

// FeeOracleV2BaseDataBufferSetIterator is returned from FilterBaseDataBufferSet and is used to iterate over the raw logs and unpacked data for BaseDataBufferSet events raised by the FeeOracleV2 contract.
type FeeOracleV2BaseDataBufferSetIterator struct {
	Event *FeeOracleV2BaseDataBufferSet // Event containing the contract specifics and raw log

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
func (it *FeeOracleV2BaseDataBufferSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeeOracleV2BaseDataBufferSet)
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
		it.Event = new(FeeOracleV2BaseDataBufferSet)
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
func (it *FeeOracleV2BaseDataBufferSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FeeOracleV2BaseDataBufferSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FeeOracleV2BaseDataBufferSet represents a BaseDataBufferSet event raised by the FeeOracleV2 contract.
type FeeOracleV2BaseDataBufferSet struct {
	DataCostId     uint64
	BaseDataBuffer uint32
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterBaseDataBufferSet is a free log retrieval operation binding the contract event 0x9f1d6f43f208ea6217bdb6274ed9ade43967929063eaa2f16274565d3af986e2.
//
// Solidity: event BaseDataBufferSet(uint64 dataCostId, uint32 baseDataBuffer)
func (_FeeOracleV2 *FeeOracleV2Filterer) FilterBaseDataBufferSet(opts *bind.FilterOpts) (*FeeOracleV2BaseDataBufferSetIterator, error) {

	logs, sub, err := _FeeOracleV2.contract.FilterLogs(opts, "BaseDataBufferSet")
	if err != nil {
		return nil, err
	}
	return &FeeOracleV2BaseDataBufferSetIterator{contract: _FeeOracleV2.contract, event: "BaseDataBufferSet", logs: logs, sub: sub}, nil
}

// WatchBaseDataBufferSet is a free log subscription operation binding the contract event 0x9f1d6f43f208ea6217bdb6274ed9ade43967929063eaa2f16274565d3af986e2.
//
// Solidity: event BaseDataBufferSet(uint64 dataCostId, uint32 baseDataBuffer)
func (_FeeOracleV2 *FeeOracleV2Filterer) WatchBaseDataBufferSet(opts *bind.WatchOpts, sink chan<- *FeeOracleV2BaseDataBufferSet) (event.Subscription, error) {

	logs, sub, err := _FeeOracleV2.contract.WatchLogs(opts, "BaseDataBufferSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FeeOracleV2BaseDataBufferSet)
				if err := _FeeOracleV2.contract.UnpackLog(event, "BaseDataBufferSet", log); err != nil {
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

// ParseBaseDataBufferSet is a log parse operation binding the contract event 0x9f1d6f43f208ea6217bdb6274ed9ade43967929063eaa2f16274565d3af986e2.
//
// Solidity: event BaseDataBufferSet(uint64 dataCostId, uint32 baseDataBuffer)
func (_FeeOracleV2 *FeeOracleV2Filterer) ParseBaseDataBufferSet(log types.Log) (*FeeOracleV2BaseDataBufferSet, error) {
	event := new(FeeOracleV2BaseDataBufferSet)
	if err := _FeeOracleV2.contract.UnpackLog(event, "BaseDataBufferSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
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
	ChainId      uint64
	BaseGasLimit uint32
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterBaseGasLimitSet is a free log retrieval operation binding the contract event 0x525a2cd9c1093178959cb9c72fe00c6be026fc953e93c9bc789d2176da98c405.
//
// Solidity: event BaseGasLimitSet(uint64 chainId, uint32 baseGasLimit)
func (_FeeOracleV2 *FeeOracleV2Filterer) FilterBaseGasLimitSet(opts *bind.FilterOpts) (*FeeOracleV2BaseGasLimitSetIterator, error) {

	logs, sub, err := _FeeOracleV2.contract.FilterLogs(opts, "BaseGasLimitSet")
	if err != nil {
		return nil, err
	}
	return &FeeOracleV2BaseGasLimitSetIterator{contract: _FeeOracleV2.contract, event: "BaseGasLimitSet", logs: logs, sub: sub}, nil
}

// WatchBaseGasLimitSet is a free log subscription operation binding the contract event 0x525a2cd9c1093178959cb9c72fe00c6be026fc953e93c9bc789d2176da98c405.
//
// Solidity: event BaseGasLimitSet(uint64 chainId, uint32 baseGasLimit)
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

// ParseBaseGasLimitSet is a log parse operation binding the contract event 0x525a2cd9c1093178959cb9c72fe00c6be026fc953e93c9bc789d2176da98c405.
//
// Solidity: event BaseGasLimitSet(uint64 chainId, uint32 baseGasLimit)
func (_FeeOracleV2 *FeeOracleV2Filterer) ParseBaseGasLimitSet(log types.Log) (*FeeOracleV2BaseGasLimitSet, error) {
	event := new(FeeOracleV2BaseGasLimitSet)
	if err := _FeeOracleV2.contract.UnpackLog(event, "BaseGasLimitSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FeeOracleV2DataCostIdSetIterator is returned from FilterDataCostIdSet and is used to iterate over the raw logs and unpacked data for DataCostIdSet events raised by the FeeOracleV2 contract.
type FeeOracleV2DataCostIdSetIterator struct {
	Event *FeeOracleV2DataCostIdSet // Event containing the contract specifics and raw log

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
func (it *FeeOracleV2DataCostIdSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeeOracleV2DataCostIdSet)
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
		it.Event = new(FeeOracleV2DataCostIdSet)
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
func (it *FeeOracleV2DataCostIdSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FeeOracleV2DataCostIdSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FeeOracleV2DataCostIdSet represents a DataCostIdSet event raised by the FeeOracleV2 contract.
type FeeOracleV2DataCostIdSet struct {
	ChainId    uint64
	DataCostId uint64
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterDataCostIdSet is a free log retrieval operation binding the contract event 0x0a5853014cbdb5103840fd3b7fcd886e7a93ef446d8c8707a5269d25ed32b4fe.
//
// Solidity: event DataCostIdSet(uint64 chainId, uint64 dataCostId)
func (_FeeOracleV2 *FeeOracleV2Filterer) FilterDataCostIdSet(opts *bind.FilterOpts) (*FeeOracleV2DataCostIdSetIterator, error) {

	logs, sub, err := _FeeOracleV2.contract.FilterLogs(opts, "DataCostIdSet")
	if err != nil {
		return nil, err
	}
	return &FeeOracleV2DataCostIdSetIterator{contract: _FeeOracleV2.contract, event: "DataCostIdSet", logs: logs, sub: sub}, nil
}

// WatchDataCostIdSet is a free log subscription operation binding the contract event 0x0a5853014cbdb5103840fd3b7fcd886e7a93ef446d8c8707a5269d25ed32b4fe.
//
// Solidity: event DataCostIdSet(uint64 chainId, uint64 dataCostId)
func (_FeeOracleV2 *FeeOracleV2Filterer) WatchDataCostIdSet(opts *bind.WatchOpts, sink chan<- *FeeOracleV2DataCostIdSet) (event.Subscription, error) {

	logs, sub, err := _FeeOracleV2.contract.WatchLogs(opts, "DataCostIdSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FeeOracleV2DataCostIdSet)
				if err := _FeeOracleV2.contract.UnpackLog(event, "DataCostIdSet", log); err != nil {
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

// ParseDataCostIdSet is a log parse operation binding the contract event 0x0a5853014cbdb5103840fd3b7fcd886e7a93ef446d8c8707a5269d25ed32b4fe.
//
// Solidity: event DataCostIdSet(uint64 chainId, uint64 dataCostId)
func (_FeeOracleV2 *FeeOracleV2Filterer) ParseDataCostIdSet(log types.Log) (*FeeOracleV2DataCostIdSet, error) {
	event := new(FeeOracleV2DataCostIdSet)
	if err := _FeeOracleV2.contract.UnpackLog(event, "DataCostIdSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FeeOracleV2DataCostParamsSetIterator is returned from FilterDataCostParamsSet and is used to iterate over the raw logs and unpacked data for DataCostParamsSet events raised by the FeeOracleV2 contract.
type FeeOracleV2DataCostParamsSetIterator struct {
	Event *FeeOracleV2DataCostParamsSet // Event containing the contract specifics and raw log

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
func (it *FeeOracleV2DataCostParamsSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeeOracleV2DataCostParamsSet)
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
		it.Event = new(FeeOracleV2DataCostParamsSet)
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
func (it *FeeOracleV2DataCostParamsSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FeeOracleV2DataCostParamsSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FeeOracleV2DataCostParamsSet represents a DataCostParamsSet event raised by the FeeOracleV2 contract.
type FeeOracleV2DataCostParamsSet struct {
	GasToken   uint8
	DataCostId uint64
	GasPrice   uint64
	GasPerByte uint64
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterDataCostParamsSet is a free log retrieval operation binding the contract event 0xdc963752c4cbfdae934dd4c5abf689a81f0e8f2eb66558af02d252abc9cec7a9.
//
// Solidity: event DataCostParamsSet(uint8 gasToken, uint64 dataCostId, uint64 gasPrice, uint64 gasPerByte)
func (_FeeOracleV2 *FeeOracleV2Filterer) FilterDataCostParamsSet(opts *bind.FilterOpts) (*FeeOracleV2DataCostParamsSetIterator, error) {

	logs, sub, err := _FeeOracleV2.contract.FilterLogs(opts, "DataCostParamsSet")
	if err != nil {
		return nil, err
	}
	return &FeeOracleV2DataCostParamsSetIterator{contract: _FeeOracleV2.contract, event: "DataCostParamsSet", logs: logs, sub: sub}, nil
}

// WatchDataCostParamsSet is a free log subscription operation binding the contract event 0xdc963752c4cbfdae934dd4c5abf689a81f0e8f2eb66558af02d252abc9cec7a9.
//
// Solidity: event DataCostParamsSet(uint8 gasToken, uint64 dataCostId, uint64 gasPrice, uint64 gasPerByte)
func (_FeeOracleV2 *FeeOracleV2Filterer) WatchDataCostParamsSet(opts *bind.WatchOpts, sink chan<- *FeeOracleV2DataCostParamsSet) (event.Subscription, error) {

	logs, sub, err := _FeeOracleV2.contract.WatchLogs(opts, "DataCostParamsSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FeeOracleV2DataCostParamsSet)
				if err := _FeeOracleV2.contract.UnpackLog(event, "DataCostParamsSet", log); err != nil {
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

// ParseDataCostParamsSet is a log parse operation binding the contract event 0xdc963752c4cbfdae934dd4c5abf689a81f0e8f2eb66558af02d252abc9cec7a9.
//
// Solidity: event DataCostParamsSet(uint8 gasToken, uint64 dataCostId, uint64 gasPrice, uint64 gasPerByte)
func (_FeeOracleV2 *FeeOracleV2Filterer) ParseDataCostParamsSet(log types.Log) (*FeeOracleV2DataCostParamsSet, error) {
	event := new(FeeOracleV2DataCostParamsSet)
	if err := _FeeOracleV2.contract.UnpackLog(event, "DataCostParamsSet", log); err != nil {
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
	DataCostId uint64
	GasPrice   uint64
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterDataGasPriceSet is a free log retrieval operation binding the contract event 0xd7d8dd5a956a8bd500e02d52d0a9dd8a0e2955ec48771a8c9da485e6706c66fb.
//
// Solidity: event DataGasPriceSet(uint64 dataCostId, uint64 gasPrice)
func (_FeeOracleV2 *FeeOracleV2Filterer) FilterDataGasPriceSet(opts *bind.FilterOpts) (*FeeOracleV2DataGasPriceSetIterator, error) {

	logs, sub, err := _FeeOracleV2.contract.FilterLogs(opts, "DataGasPriceSet")
	if err != nil {
		return nil, err
	}
	return &FeeOracleV2DataGasPriceSetIterator{contract: _FeeOracleV2.contract, event: "DataGasPriceSet", logs: logs, sub: sub}, nil
}

// WatchDataGasPriceSet is a free log subscription operation binding the contract event 0xd7d8dd5a956a8bd500e02d52d0a9dd8a0e2955ec48771a8c9da485e6706c66fb.
//
// Solidity: event DataGasPriceSet(uint64 dataCostId, uint64 gasPrice)
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
// Solidity: event DataGasPriceSet(uint64 dataCostId, uint64 gasPrice)
func (_FeeOracleV2 *FeeOracleV2Filterer) ParseDataGasPriceSet(log types.Log) (*FeeOracleV2DataGasPriceSet, error) {
	event := new(FeeOracleV2DataGasPriceSet)
	if err := _FeeOracleV2.contract.UnpackLog(event, "DataGasPriceSet", log); err != nil {
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

// FeeOracleV2FeeParamsSetIterator is returned from FilterFeeParamsSet and is used to iterate over the raw logs and unpacked data for FeeParamsSet events raised by the FeeOracleV2 contract.
type FeeOracleV2FeeParamsSetIterator struct {
	Event *FeeOracleV2FeeParamsSet // Event containing the contract specifics and raw log

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
func (it *FeeOracleV2FeeParamsSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeeOracleV2FeeParamsSet)
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
		it.Event = new(FeeOracleV2FeeParamsSet)
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
func (it *FeeOracleV2FeeParamsSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FeeOracleV2FeeParamsSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FeeOracleV2FeeParamsSet represents a FeeParamsSet event raised by the FeeOracleV2 contract.
type FeeOracleV2FeeParamsSet struct {
	GasToken   uint8
	ChainId    uint64
	GasPrice   uint64
	DataCostId uint64
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterFeeParamsSet is a free log retrieval operation binding the contract event 0x4928444fd66b441b3ace1058b261e66af8a0193f1e0533d23ed48d0270075551.
//
// Solidity: event FeeParamsSet(uint8 gasToken, uint64 chainId, uint64 gasPrice, uint64 dataCostId)
func (_FeeOracleV2 *FeeOracleV2Filterer) FilterFeeParamsSet(opts *bind.FilterOpts) (*FeeOracleV2FeeParamsSetIterator, error) {

	logs, sub, err := _FeeOracleV2.contract.FilterLogs(opts, "FeeParamsSet")
	if err != nil {
		return nil, err
	}
	return &FeeOracleV2FeeParamsSetIterator{contract: _FeeOracleV2.contract, event: "FeeParamsSet", logs: logs, sub: sub}, nil
}

// WatchFeeParamsSet is a free log subscription operation binding the contract event 0x4928444fd66b441b3ace1058b261e66af8a0193f1e0533d23ed48d0270075551.
//
// Solidity: event FeeParamsSet(uint8 gasToken, uint64 chainId, uint64 gasPrice, uint64 dataCostId)
func (_FeeOracleV2 *FeeOracleV2Filterer) WatchFeeParamsSet(opts *bind.WatchOpts, sink chan<- *FeeOracleV2FeeParamsSet) (event.Subscription, error) {

	logs, sub, err := _FeeOracleV2.contract.WatchLogs(opts, "FeeParamsSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FeeOracleV2FeeParamsSet)
				if err := _FeeOracleV2.contract.UnpackLog(event, "FeeParamsSet", log); err != nil {
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

// ParseFeeParamsSet is a log parse operation binding the contract event 0x4928444fd66b441b3ace1058b261e66af8a0193f1e0533d23ed48d0270075551.
//
// Solidity: event FeeParamsSet(uint8 gasToken, uint64 chainId, uint64 gasPrice, uint64 dataCostId)
func (_FeeOracleV2 *FeeOracleV2Filterer) ParseFeeParamsSet(log types.Log) (*FeeOracleV2FeeParamsSet, error) {
	event := new(FeeOracleV2FeeParamsSet)
	if err := _FeeOracleV2.contract.UnpackLog(event, "FeeParamsSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FeeOracleV2GasPerByteSetIterator is returned from FilterGasPerByteSet and is used to iterate over the raw logs and unpacked data for GasPerByteSet events raised by the FeeOracleV2 contract.
type FeeOracleV2GasPerByteSetIterator struct {
	Event *FeeOracleV2GasPerByteSet // Event containing the contract specifics and raw log

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
func (it *FeeOracleV2GasPerByteSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeeOracleV2GasPerByteSet)
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
		it.Event = new(FeeOracleV2GasPerByteSet)
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
func (it *FeeOracleV2GasPerByteSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FeeOracleV2GasPerByteSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FeeOracleV2GasPerByteSet represents a GasPerByteSet event raised by the FeeOracleV2 contract.
type FeeOracleV2GasPerByteSet struct {
	DataCostId uint64
	GasPerByte uint64
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterGasPerByteSet is a free log retrieval operation binding the contract event 0x9e8c8606adb2b50f48cb69ccb1c3e349e6046ed63765e5931c01f4fe6aacd5e9.
//
// Solidity: event GasPerByteSet(uint64 dataCostId, uint64 gasPerByte)
func (_FeeOracleV2 *FeeOracleV2Filterer) FilterGasPerByteSet(opts *bind.FilterOpts) (*FeeOracleV2GasPerByteSetIterator, error) {

	logs, sub, err := _FeeOracleV2.contract.FilterLogs(opts, "GasPerByteSet")
	if err != nil {
		return nil, err
	}
	return &FeeOracleV2GasPerByteSetIterator{contract: _FeeOracleV2.contract, event: "GasPerByteSet", logs: logs, sub: sub}, nil
}

// WatchGasPerByteSet is a free log subscription operation binding the contract event 0x9e8c8606adb2b50f48cb69ccb1c3e349e6046ed63765e5931c01f4fe6aacd5e9.
//
// Solidity: event GasPerByteSet(uint64 dataCostId, uint64 gasPerByte)
func (_FeeOracleV2 *FeeOracleV2Filterer) WatchGasPerByteSet(opts *bind.WatchOpts, sink chan<- *FeeOracleV2GasPerByteSet) (event.Subscription, error) {

	logs, sub, err := _FeeOracleV2.contract.WatchLogs(opts, "GasPerByteSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FeeOracleV2GasPerByteSet)
				if err := _FeeOracleV2.contract.UnpackLog(event, "GasPerByteSet", log); err != nil {
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

// ParseGasPerByteSet is a log parse operation binding the contract event 0x9e8c8606adb2b50f48cb69ccb1c3e349e6046ed63765e5931c01f4fe6aacd5e9.
//
// Solidity: event GasPerByteSet(uint64 dataCostId, uint64 gasPerByte)
func (_FeeOracleV2 *FeeOracleV2Filterer) ParseGasPerByteSet(log types.Log) (*FeeOracleV2GasPerByteSet, error) {
	event := new(FeeOracleV2GasPerByteSet)
	if err := _FeeOracleV2.contract.UnpackLog(event, "GasPerByteSet", log); err != nil {
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

// FilterProtocolFeeSet is a free log retrieval operation binding the contract event 0xd91752439e358587fc0828ed743df5939f16a918de501834bd954d03be15c959.
//
// Solidity: event ProtocolFeeSet(uint96 protocolFee)
func (_FeeOracleV2 *FeeOracleV2Filterer) FilterProtocolFeeSet(opts *bind.FilterOpts) (*FeeOracleV2ProtocolFeeSetIterator, error) {

	logs, sub, err := _FeeOracleV2.contract.FilterLogs(opts, "ProtocolFeeSet")
	if err != nil {
		return nil, err
	}
	return &FeeOracleV2ProtocolFeeSetIterator{contract: _FeeOracleV2.contract, event: "ProtocolFeeSet", logs: logs, sub: sub}, nil
}

// WatchProtocolFeeSet is a free log subscription operation binding the contract event 0xd91752439e358587fc0828ed743df5939f16a918de501834bd954d03be15c959.
//
// Solidity: event ProtocolFeeSet(uint96 protocolFee)
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

// ParseProtocolFeeSet is a log parse operation binding the contract event 0xd91752439e358587fc0828ed743df5939f16a918de501834bd954d03be15c959.
//
// Solidity: event ProtocolFeeSet(uint96 protocolFee)
func (_FeeOracleV2 *FeeOracleV2Filterer) ParseProtocolFeeSet(log types.Log) (*FeeOracleV2ProtocolFeeSet, error) {
	event := new(FeeOracleV2ProtocolFeeSet)
	if err := _FeeOracleV2.contract.UnpackLog(event, "ProtocolFeeSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// FeeOracleV2ToNativeRateSetIterator is returned from FilterToNativeRateSet and is used to iterate over the raw logs and unpacked data for ToNativeRateSet events raised by the FeeOracleV2 contract.
type FeeOracleV2ToNativeRateSetIterator struct {
	Event *FeeOracleV2ToNativeRateSet // Event containing the contract specifics and raw log

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
func (it *FeeOracleV2ToNativeRateSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeeOracleV2ToNativeRateSet)
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
		it.Event = new(FeeOracleV2ToNativeRateSet)
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
func (it *FeeOracleV2ToNativeRateSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FeeOracleV2ToNativeRateSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FeeOracleV2ToNativeRateSet represents a ToNativeRateSet event raised by the FeeOracleV2 contract.
type FeeOracleV2ToNativeRateSet struct {
	GasToken   uint8
	NativeRate *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterToNativeRateSet is a free log retrieval operation binding the contract event 0xd665ed89605ad8a805fa330621c3a15c9e6de67d51dbc52a478a868e1ed616e4.
//
// Solidity: event ToNativeRateSet(uint8 gasToken, uint256 nativeRate)
func (_FeeOracleV2 *FeeOracleV2Filterer) FilterToNativeRateSet(opts *bind.FilterOpts) (*FeeOracleV2ToNativeRateSetIterator, error) {

	logs, sub, err := _FeeOracleV2.contract.FilterLogs(opts, "ToNativeRateSet")
	if err != nil {
		return nil, err
	}
	return &FeeOracleV2ToNativeRateSetIterator{contract: _FeeOracleV2.contract, event: "ToNativeRateSet", logs: logs, sub: sub}, nil
}

// WatchToNativeRateSet is a free log subscription operation binding the contract event 0xd665ed89605ad8a805fa330621c3a15c9e6de67d51dbc52a478a868e1ed616e4.
//
// Solidity: event ToNativeRateSet(uint8 gasToken, uint256 nativeRate)
func (_FeeOracleV2 *FeeOracleV2Filterer) WatchToNativeRateSet(opts *bind.WatchOpts, sink chan<- *FeeOracleV2ToNativeRateSet) (event.Subscription, error) {

	logs, sub, err := _FeeOracleV2.contract.WatchLogs(opts, "ToNativeRateSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FeeOracleV2ToNativeRateSet)
				if err := _FeeOracleV2.contract.UnpackLog(event, "ToNativeRateSet", log); err != nil {
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

// ParseToNativeRateSet is a log parse operation binding the contract event 0xd665ed89605ad8a805fa330621c3a15c9e6de67d51dbc52a478a868e1ed616e4.
//
// Solidity: event ToNativeRateSet(uint8 gasToken, uint256 nativeRate)
func (_FeeOracleV2 *FeeOracleV2Filterer) ParseToNativeRateSet(log types.Log) (*FeeOracleV2ToNativeRateSet, error) {
	event := new(FeeOracleV2ToNativeRateSet)
	if err := _FeeOracleV2.contract.UnpackLog(event, "ToNativeRateSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
