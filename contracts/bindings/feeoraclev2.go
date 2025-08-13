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
	GasToken   uint16
	BaseBytes  uint32
	Id         uint64
	GasPrice   uint64
	GasPerByte uint64
}

// IFeeOracleV2FeeParams is an auto generated low-level Go binding around an user-defined struct.
type IFeeOracleV2FeeParams struct {
	GasToken     uint16
	BaseGasLimit uint32
	ChainId      uint64
	GasPrice     uint64
	DataCostId   uint64
}

// IFeeOracleV2ToNativeRateParams is an auto generated low-level Go binding around an user-defined struct.
type IFeeOracleV2ToNativeRateParams struct {
	GasToken   uint16
	NativeRate *big.Int
}

// FeeOracleV2MetaData contains all meta data concerning the FeeOracleV2 contract.
var FeeOracleV2MetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"CONVERSION_RATE_DENOM\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"baseBytes\",\"inputs\":[{\"name\":\"dataCostId\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"baseGasLimit\",\"inputs\":[{\"name\":\"chainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"bulkSetDataCostParams\",\"inputs\":[{\"name\":\"params\",\"type\":\"tuple[]\",\"internalType\":\"structIFeeOracleV2.DataCostParams[]\",\"components\":[{\"name\":\"gasToken\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"baseBytes\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"id\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"gasPrice\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"gasPerByte\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"bulkSetFeeParams\",\"inputs\":[{\"name\":\"params\",\"type\":\"tuple[]\",\"internalType\":\"structIFeeOracleV2.FeeParams[]\",\"components\":[{\"name\":\"gasToken\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"baseGasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"chainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"gasPrice\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"dataCostId\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"bulkSetToNativeRate\",\"inputs\":[{\"name\":\"params\",\"type\":\"tuple[]\",\"internalType\":\"structIFeeOracleV2.ToNativeRateParams[]\",\"components\":[{\"name\":\"gasToken\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"nativeRate\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"dataCostParams\",\"inputs\":[{\"name\":\"dataCostId\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIFeeOracleV2.DataCostParams\",\"components\":[{\"name\":\"gasToken\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"baseBytes\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"id\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"gasPrice\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"gasPerByte\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"dataGasPerByte\",\"inputs\":[{\"name\":\"dataCostId\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"dataGasPrice\",\"inputs\":[{\"name\":\"dataCostId\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"dataGasToken\",\"inputs\":[{\"name\":\"dataCostId\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"execDataCostId\",\"inputs\":[{\"name\":\"chainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"execGasPrice\",\"inputs\":[{\"name\":\"chainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"execGasToken\",\"inputs\":[{\"name\":\"chainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"feeFor\",\"inputs\":[{\"name\":\"destChainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"gasLimit\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"feeParams\",\"inputs\":[{\"name\":\"chainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structIFeeOracleV2.FeeParams\",\"components\":[{\"name\":\"gasToken\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"baseGasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"chainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"gasPrice\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"dataCostId\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"owner_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"manager_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"protocolFee_\",\"type\":\"uint96\",\"internalType\":\"uint96\"},{\"name\":\"feeParams_\",\"type\":\"tuple[]\",\"internalType\":\"structIFeeOracleV2.FeeParams[]\",\"components\":[{\"name\":\"gasToken\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"baseGasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"chainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"gasPrice\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"dataCostId\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"dataCostParams_\",\"type\":\"tuple[]\",\"internalType\":\"structIFeeOracleV2.DataCostParams[]\",\"components\":[{\"name\":\"gasToken\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"baseBytes\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"id\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"gasPrice\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"gasPerByte\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"toNativeRateParams_\",\"type\":\"tuple[]\",\"internalType\":\"structIFeeOracleV2.ToNativeRateParams[]\",\"components\":[{\"name\":\"gasToken\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"nativeRate\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"manager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"protocolFee\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint96\",\"internalType\":\"uint96\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setBaseBytes\",\"inputs\":[{\"name\":\"dataCostId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"newBaseBytes\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setBaseGasLimit\",\"inputs\":[{\"name\":\"chainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"newBaseGasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setDataCostId\",\"inputs\":[{\"name\":\"chainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"dataCostId\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setDataGasPrice\",\"inputs\":[{\"name\":\"dataCostId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"gasPrice\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setExecGasPrice\",\"inputs\":[{\"name\":\"chainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"gasPrice\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setGasPerByte\",\"inputs\":[{\"name\":\"dataCostId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"gasPerByte\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setManager\",\"inputs\":[{\"name\":\"manager_\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setProtocolFee\",\"inputs\":[{\"name\":\"fee\",\"type\":\"uint96\",\"internalType\":\"uint96\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setToNativeRate\",\"inputs\":[{\"name\":\"gasToken\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"nativeRate\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"toNativeRate\",\"inputs\":[{\"name\":\"chainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"tokenToNativeRate\",\"inputs\":[{\"name\":\"gasToken\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"version\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"pure\"},{\"type\":\"event\",\"name\":\"BaseBytesSet\",\"inputs\":[{\"name\":\"dataCostId\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"baseBytes\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"BaseGasLimitSet\",\"inputs\":[{\"name\":\"chainId\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"baseGasLimit\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DataCostIdSet\",\"inputs\":[{\"name\":\"chainId\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"dataCostId\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DataCostParamsSet\",\"inputs\":[{\"name\":\"gasToken\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"},{\"name\":\"baseBytes\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"},{\"name\":\"id\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"gasPrice\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"gasPerByte\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DataGasPriceSet\",\"inputs\":[{\"name\":\"dataCostId\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"gasPrice\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ExecGasPriceSet\",\"inputs\":[{\"name\":\"chainId\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"gasPrice\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FeeParamsSet\",\"inputs\":[{\"name\":\"gasToken\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"},{\"name\":\"baseGasLimit\",\"type\":\"uint32\",\"indexed\":false,\"internalType\":\"uint32\"},{\"name\":\"chainId\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"gasPrice\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"dataCostId\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"GasPerByteSet\",\"inputs\":[{\"name\":\"dataCostId\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"gasPerByte\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ManagerSet\",\"inputs\":[{\"name\":\"manager\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ProtocolFeeSet\",\"inputs\":[{\"name\":\"protocolFee\",\"type\":\"uint96\",\"indexed\":false,\"internalType\":\"uint96\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ToNativeRateSet\",\"inputs\":[{\"name\":\"gasToken\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"},{\"name\":\"nativeRate\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NoFeeParams\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotManager\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnableInvalidOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnableUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ZeroAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroChainId\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroDataCostId\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroGasPerByte\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroGasPrice\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroGasToken\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroNativeRate\",\"inputs\":[]}]",
	Bin: "0x608060405234801561000f575f80fd5b5061001861001d565b6100cf565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000900460ff161561006d5760405163f92ee8a960e01b815260040160405180910390fd5b80546001600160401b03908116146100cc5780546001600160401b0319166001600160401b0390811782556040519081527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b50565b611d9a806100dc5f395ff3fe608060405234801561000f575f80fd5b50600436106101f2575f3560e01c80639a5551c311610114578063bfc71416116100a9578063db847a5911610079578063db847a59146106d1578063e21497b9146106e4578063e5737f6214610716578063f2fde38b14610729578063f364f6871461073c575f80fd5b8063bfc71416146105b7578063d0ebdbe714610660578063d32b68ad14610673578063db0018e81461069f575f80fd5b8063b17db68a116100e4578063b17db68a14610533578063b984cc0b14610552578063b9923e1c14610591578063bc51bf37146105a4575f80fd5b80639a5551c3146104b15780639c742ced146104e3578063a12f2c58146104f6578063b0e21e8a14610509575f80fd5b80635d3acee21161018a5780638b7bfd701161015a5780638b7bfd701461041d5780638da5cb5b146104645780638dd9523c146104945780638f9d6ace146104a7575f80fd5b80635d3acee2146103bb578063653c356e146103ce5780636b6dccfe14610402578063715018a614610415575f80fd5b8063481c6a75116101c5578063481c6a751461032757806350b815391461035857806354fd4d50146103a157806356bce459146103a8575f80fd5b80632105b755146101f6578063223aacf81461020b5780632d4634a41461021e578063415070af146102dd575b5f80fd5b61020961020436600461180d565b61074f565b005b6102096102193660046118bc565b61078e565b6102c761022c366004611996565b6040805160a0810182525f80825260208201819052918101829052606081018290526080810191909152506001600160401b039081165f90815260026020908152604091829020825160a081018452905461ffff8116825263ffffffff6201000082041692820192909252600160301b8204841692810192909252600160701b810483166060830152600160b01b9004909116608082015290565b6040516102d491906119b6565b60405180910390f35b61030f6102eb366004611996565b6001600160401b039081165f90815260036020526040902054600160701b90041690565b6040516001600160401b0390911681526020016102d4565b5f5461034090600160601b90046001600160a01b031681565b6040516001600160a01b0390911681526020016102d4565b61038c610366366004611996565b6001600160401b03165f9081526002602052604090205462010000900463ffffffff1690565b60405163ffffffff90911681526020016102d4565b600261030f565b6102096103b6366004611a0a565b6108bd565b6102096103c9366004611a0a565b6108f8565b61038c6103dc366004611996565b6001600160401b03165f9081526003602052604090205462010000900463ffffffff1690565b610209610410366004611a3b565b610933565b610209610947565b61045661042b366004611996565b6001600160401b03165f9081526002602090815260408083205461ffff168352600190915290205490565b6040519081526020016102d4565b7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c199300546001600160a01b0316610340565b6104566104a2366004611a54565b61095a565b610456620f424081565b61030f6104bf366004611996565b6001600160401b039081165f90815260026020526040902054600160b01b90041690565b6102096104f1366004611af3565b610ae1565b610209610504366004611b2c565b610b1c565b5f5461051b906001600160601b031681565b6040516001600160601b0390911681526020016102d4565b610456610541366004611b54565b60016020525f908152604090205481565b61057e610560366004611996565b6001600160401b03165f9081526002602052604090205461ffff1690565b60405161ffff90911681526020016102d4565b61020961059f366004611a0a565b610b57565b6102096105b2366004611a0a565b610b92565b6102c76105c5366004611996565b6040805160a0810182525f80825260208201819052918101829052606081018290526080810191909152506001600160401b039081165f90815260036020908152604091829020825160a081018452905461ffff8116825263ffffffff6201000082041692820192909252600160301b8204841692810192909252600160701b810483166060830152600160b01b9004909116608082015290565b61020961066e366004611b6d565b610bcd565b61057e610681366004611996565b6001600160401b03165f9081526003602052604090205461ffff1690565b61030f6106ad366004611996565b6001600160401b039081165f90815260036020526040902054600160b01b90041690565b6102096106df36600461180d565b610c05565b61030f6106f2366004611996565b6001600160401b039081165f90815260026020526040902054600160701b90041690565b610209610724366004611af3565b610c40565b610209610737366004611b6d565b610c7b565b61020961074a366004611b86565b610cba565b5f54600160601b90046001600160a01b031633146107805760405163607e454560e11b815260040160405180910390fd5b61078a8282610cf5565b5050565b5f610797610edd565b805490915060ff600160401b82041615906001600160401b03165f811580156107bd5750825b90505f826001600160401b031660011480156107d85750303b155b9050811580156107e6575080155b156108045760405163f92ee8a960e01b815260040160405180910390fd5b845467ffffffffffffffff19166001178555831561082e57845460ff60401b1916600160401b1785555b6108378e610f07565b6108408d610f18565b6108498c610f74565b6108538b8b610fc6565b61085d8989610cf5565b61086787876111a9565b83156108ad57845460ff60401b19168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b5050505050505050505050505050565b5f54600160601b90046001600160a01b031633146108ee5760405163607e454560e11b815260040160405180910390fd5b61078a82826111f8565b5f54600160601b90046001600160a01b031633146109295760405163607e454560e11b815260040160405180910390fd5b61078a82826112bf565b61093b61137e565b61094481610f74565b50565b61094f61137e565b6109585f6113d9565b565b6001600160401b038085165f9081526002602090815260408083208054600160b01b8104861685526003845282852061ffff82168652600190945291842054939490938592620f4240926109b69291600160701b900416611bcc565b6109c09190611be3565b825461ffff81165f9081526001602052604081205492935091620f4240916109f89190600160701b90046001600160401b0316611bcc565b610a029190611be3565b9050815f03610a2457604051633532119760e11b815260040160405180910390fd5b805f03610a4457604051633532119760e11b815260040160405180910390fd5b82545f90600160b01b81046001600160401b031690610a70908a9062010000900463ffffffff16611c02565b610a7a9190611bcc565b9050610a868282611bcc565b85548490610aa1908a9062010000900463ffffffff16611c15565b6001600160401b0316610ab49190611bcc565b5f54610ac991906001600160601b0316611c02565b610ad39190611c02565b9a9950505050505050505050565b5f54600160601b90046001600160a01b03163314610b125760405163607e454560e11b815260040160405180910390fd5b61078a8282611449565b5f54600160601b90046001600160a01b03163314610b4d5760405163607e454560e11b815260040160405180910390fd5b61078a82826114de565b5f54600160601b90046001600160a01b03163314610b885760405163607e454560e11b815260040160405180910390fd5b61078a828261156d565b5f54600160601b90046001600160a01b03163314610bc35760405163607e454560e11b815260040160405180910390fd5b61078a828261162c565b610bd561137e565b6001600160a01b038116610bfc5760405163d92e233d60e01b815260040160405180910390fd5b61094481610f18565b5f54600160601b90046001600160a01b03163314610c365760405163607e454560e11b815260040160405180910390fd5b61078a8282610fc6565b5f54600160601b90046001600160a01b03163314610c715760405163607e454560e11b815260040160405180910390fd5b61078a82826116eb565b610c8361137e565b6001600160a01b038116610cb157604051631e4fbdf760e01b81525f60048201526024015b60405180910390fd5b610944816113d9565b5f54600160601b90046001600160a01b03163314610ceb5760405163607e454560e11b815260040160405180910390fd5b61078a82826111a9565b5f5b81811015610ed8575f838383818110610d1257610d12611c3c565b905060a00201803603810190610d289190611cea565b805190915061ffff165f03610d50576040516350614df960e01b815260040160405180910390fd5b80604001516001600160401b03165f03610d7d57604051630c26851b60e11b815260040160405180910390fd5b80606001516001600160401b03165f03610daa57604051630e661aed60e41b815260040160405180910390fd5b80608001516001600160401b03165f03610dd7576040516348cfc33560e11b815260040160405180910390fd5b604081810180516001600160401b039081165f9081526003602090815290849020855181548388015195516060808a01516080808c015161ffff90961665ffffffffffff1990951685176201000063ffffffff909b169a8b02176601000000000000600160b01b031916600160301b948a1694850267ffffffffffffffff60701b191617600160701b928a169283021767ffffffffffffffff60b01b1916600160b01b969099169586029890981790955588519283529482019690965295860194909452908401528201527fd143a0934cc5e5337dca3eb0afa1e7f86680796ca6f132bfe7e0828b7155bd409060a00160405180910390a150600101610cf7565b505050565b5f807ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a005b92915050565b610f0f611780565b610944816117a5565b5f80546001600160601b0316600160601b6001600160a01b038416908102919091179091556040519081527f60a0f5b9f9e81e98216071b85826681c796256fe3d1354ecb675580fba64fa69906020015b60405180910390a150565b5f80546bffffffffffffffffffffffff19166001600160601b0383169081179091556040519081527fd91752439e358587fc0828ed743df5939f16a918de501834bd954d03be15c95990602001610f69565b5f5b81811015610ed8575f838383818110610fe357610fe3611c3c565b905060a00201803603810190610ff99190611cea565b805190915061ffff165f03611021576040516350614df960e01b815260040160405180910390fd5b80604001516001600160401b03165f0361104e57604051633212217560e21b815260040160405180910390fd5b80606001516001600160401b03165f0361107b57604051630e661aed60e41b815260040160405180910390fd5b80608001516001600160401b03165f036110a857604051630c26851b60e11b815260040160405180910390fd5b604081810180516001600160401b039081165f9081526002602090815290849020855181548388015195516060808a01516080808c015161ffff90961665ffffffffffff1990951685176201000063ffffffff909b169a8b02176601000000000000600160b01b031916600160301b948a1694850267ffffffffffffffff60701b191617600160701b928a169283021767ffffffffffffffff60b01b1916600160b01b969099169586029890981790955588519283529482019690965295860194909452908401528201527f600e7ff14e74285e17debda1fee2df93741c6518e12ba908ba4417c0610974a99060a00160405180910390a150600101610fc8565b5f5b81811015610ed8575f8383838181106111c6576111c6611c3c565b9050604002018036038101906111dc9190611d04565b90506111ef815f015182602001516114de565b506001016111ab565b816001600160401b03165f0361122157604051630c26851b60e11b815260040160405180910390fd5b806001600160401b03165f0361124a576040516348cfc33560e11b815260040160405180910390fd5b6001600160401b038281165f81815260036020908152604091829020805467ffffffffffffffff60b01b1916600160b01b95871695860217905581519283528201929092527f9e8c8606adb2b50f48cb69ccb1c3e349e6046ed63765e5931c01f4fe6aacd5e991015b60405180910390a15050565b806001600160401b03165f036112e857604051630e661aed60e41b815260040160405180910390fd5b816001600160401b03165f0361131157604051630c26851b60e11b815260040160405180910390fd5b6001600160401b038281165f81815260036020908152604091829020805467ffffffffffffffff60701b1916600160701b95871695860217905581519283528201929092527fd7d8dd5a956a8bd500e02d52d0a9dd8a0e2955ec48771a8c9da485e6706c66fb91016112b3565b336113b07f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c199300546001600160a01b031690565b6001600160a01b0316146109585760405163118cdaa760e01b8152336004820152602401610ca8565b7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c19930080546001600160a01b031981166001600160a01b03848116918217845560405192169182907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0905f90a3505050565b816001600160401b03165f0361147257604051630c26851b60e11b815260040160405180910390fd5b6001600160401b0382165f81815260036020908152604091829020805465ffffffff000019166201000063ffffffff8716908102919091179091558251938452908301527f406ba3ab8c23e58f883620734592e662ac0018177590c1f9f27afc58a4b7b58391016112b3565b805f036114fe5760405163fa90419960e01b815260040160405180910390fd5b8161ffff165f03611522576040516350614df960e01b815260040160405180910390fd5b61ffff82165f81815260016020908152604091829020849055815192835282018390527f770404c81b5a366795e9a06ff3969e3fb2eb0515b7282e986f261c180b97623d91016112b3565b806001600160401b03165f0361159657604051630e661aed60e41b815260040160405180910390fd5b816001600160401b03165f036115bf57604051633212217560e21b815260040160405180910390fd5b6001600160401b038281165f81815260026020908152604091829020805467ffffffffffffffff60701b1916600160701b95871695860217905581519283528201929092527fe0e5abb8929e27a69d77f47a4e3f9575411a5be1fa596e5b55078d7850f358db91016112b3565b816001600160401b03165f0361165557604051633212217560e21b815260040160405180910390fd5b806001600160401b03165f0361167e57604051630c26851b60e11b815260040160405180910390fd5b6001600160401b038281165f81815260026020908152604091829020805467ffffffffffffffff60b01b1916600160b01b95871695860217905581519283528201929092527f0a5853014cbdb5103840fd3b7fcd886e7a93ef446d8c8707a5269d25ed32b4fe91016112b3565b816001600160401b03165f0361171457604051633212217560e21b815260040160405180910390fd5b6001600160401b0382165f81815260026020908152604091829020805465ffffffff000019166201000063ffffffff8716908102919091179091558251938452908301527f525a2cd9c1093178959cb9c72fe00c6be026fc953e93c9bc789d2176da98c40591016112b3565b6117886117ad565b61095857604051631afcd79f60e31b815260040160405180910390fd5b610c83611780565b5f6117b6610edd565b54600160401b900460ff16919050565b5f8083601f8401126117d6575f80fd5b5081356001600160401b038111156117ec575f80fd5b60208301915083602060a083028501011115611806575f80fd5b9250929050565b5f806020838503121561181e575f80fd5b82356001600160401b03811115611833575f80fd5b61183f858286016117c6565b90969095509350505050565b80356001600160a01b0381168114611861575f80fd5b919050565b80356001600160601b0381168114611861575f80fd5b5f8083601f84011261188c575f80fd5b5081356001600160401b038111156118a2575f80fd5b6020830191508360208260061b8501011115611806575f80fd5b5f805f805f805f805f60c08a8c0312156118d4575f80fd5b6118dd8a61184b565b98506118eb60208b0161184b565b97506118f960408b01611866565b965060608a01356001600160401b0380821115611914575f80fd5b6119208d838e016117c6565b909850965060808c0135915080821115611938575f80fd5b6119448d838e016117c6565b909650945060a08c013591508082111561195c575f80fd5b506119698c828d0161187c565b915080935050809150509295985092959850929598565b80356001600160401b0381168114611861575f80fd5b5f602082840312156119a6575f80fd5b6119af82611980565b9392505050565b60a08101610f01828461ffff815116825263ffffffff602082015116602083015260408101516001600160401b03808216604085015280606084015116606085015280608084015116608085015250505050565b5f8060408385031215611a1b575f80fd5b611a2483611980565b9150611a3260208401611980565b90509250929050565b5f60208284031215611a4b575f80fd5b6119af82611866565b5f805f8060608587031215611a67575f80fd5b611a7085611980565b935060208501356001600160401b0380821115611a8b575f80fd5b818701915087601f830112611a9e575f80fd5b813581811115611aac575f80fd5b886020828501011115611abd575f80fd5b602083019550809450505050611ad560408601611980565b905092959194509250565b803563ffffffff81168114611861575f80fd5b5f8060408385031215611b04575f80fd5b611b0d83611980565b9150611a3260208401611ae0565b803561ffff81168114611861575f80fd5b5f8060408385031215611b3d575f80fd5b611b4683611b1b565b946020939093013593505050565b5f60208284031215611b64575f80fd5b6119af82611b1b565b5f60208284031215611b7d575f80fd5b6119af8261184b565b5f8060208385031215611b97575f80fd5b82356001600160401b03811115611bac575f80fd5b61183f8582860161187c565b634e487b7160e01b5f52601160045260245ffd5b8082028115828204841417610f0157610f01611bb8565b5f82611bfd57634e487b7160e01b5f52601260045260245ffd5b500490565b80820180821115610f0157610f01611bb8565b6001600160401b03818116838216019080821115611c3557611c35611bb8565b5092915050565b634e487b7160e01b5f52603260045260245ffd5b5f60a08284031215611c60575f80fd5b60405160a081018181106001600160401b0382111715611c8e57634e487b7160e01b5f52604160045260245ffd5b604052905080611c9d83611b1b565b8152611cab60208401611ae0565b6020820152611cbc60408401611980565b6040820152611ccd60608401611980565b6060820152611cde60808401611980565b60808201525092915050565b5f60a08284031215611cfa575f80fd5b6119af8383611c50565b5f60408284031215611d14575f80fd5b604051604081018181106001600160401b0382111715611d4257634e487b7160e01b5f52604160045260245ffd5b604052611d4e83611b1b565b815260208301356020820152809150509291505056fea2646970667358221220d5ef050ed7952abd0e509693a529139613b7ee33fd90bd66cb90d7ad14eb248164736f6c63430008180033",
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

// BaseBytes is a free data retrieval call binding the contract method 0x653c356e.
//
// Solidity: function baseBytes(uint64 dataCostId) view returns(uint32)
func (_FeeOracleV2 *FeeOracleV2Caller) BaseBytes(opts *bind.CallOpts, dataCostId uint64) (uint32, error) {
	var out []interface{}
	err := _FeeOracleV2.contract.Call(opts, &out, "baseBytes", dataCostId)

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// BaseBytes is a free data retrieval call binding the contract method 0x653c356e.
//
// Solidity: function baseBytes(uint64 dataCostId) view returns(uint32)
func (_FeeOracleV2 *FeeOracleV2Session) BaseBytes(dataCostId uint64) (uint32, error) {
	return _FeeOracleV2.Contract.BaseBytes(&_FeeOracleV2.CallOpts, dataCostId)
}

// BaseBytes is a free data retrieval call binding the contract method 0x653c356e.
//
// Solidity: function baseBytes(uint64 dataCostId) view returns(uint32)
func (_FeeOracleV2 *FeeOracleV2CallerSession) BaseBytes(dataCostId uint64) (uint32, error) {
	return _FeeOracleV2.Contract.BaseBytes(&_FeeOracleV2.CallOpts, dataCostId)
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
// Solidity: function dataCostParams(uint64 dataCostId) view returns((uint16,uint32,uint64,uint64,uint64))
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
// Solidity: function dataCostParams(uint64 dataCostId) view returns((uint16,uint32,uint64,uint64,uint64))
func (_FeeOracleV2 *FeeOracleV2Session) DataCostParams(dataCostId uint64) (IFeeOracleV2DataCostParams, error) {
	return _FeeOracleV2.Contract.DataCostParams(&_FeeOracleV2.CallOpts, dataCostId)
}

// DataCostParams is a free data retrieval call binding the contract method 0xbfc71416.
//
// Solidity: function dataCostParams(uint64 dataCostId) view returns((uint16,uint32,uint64,uint64,uint64))
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
// Solidity: function dataGasToken(uint64 dataCostId) view returns(uint16)
func (_FeeOracleV2 *FeeOracleV2Caller) DataGasToken(opts *bind.CallOpts, dataCostId uint64) (uint16, error) {
	var out []interface{}
	err := _FeeOracleV2.contract.Call(opts, &out, "dataGasToken", dataCostId)

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// DataGasToken is a free data retrieval call binding the contract method 0xd32b68ad.
//
// Solidity: function dataGasToken(uint64 dataCostId) view returns(uint16)
func (_FeeOracleV2 *FeeOracleV2Session) DataGasToken(dataCostId uint64) (uint16, error) {
	return _FeeOracleV2.Contract.DataGasToken(&_FeeOracleV2.CallOpts, dataCostId)
}

// DataGasToken is a free data retrieval call binding the contract method 0xd32b68ad.
//
// Solidity: function dataGasToken(uint64 dataCostId) view returns(uint16)
func (_FeeOracleV2 *FeeOracleV2CallerSession) DataGasToken(dataCostId uint64) (uint16, error) {
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
// Solidity: function execGasToken(uint64 chainId) view returns(uint16)
func (_FeeOracleV2 *FeeOracleV2Caller) ExecGasToken(opts *bind.CallOpts, chainId uint64) (uint16, error) {
	var out []interface{}
	err := _FeeOracleV2.contract.Call(opts, &out, "execGasToken", chainId)

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// ExecGasToken is a free data retrieval call binding the contract method 0xb984cc0b.
//
// Solidity: function execGasToken(uint64 chainId) view returns(uint16)
func (_FeeOracleV2 *FeeOracleV2Session) ExecGasToken(chainId uint64) (uint16, error) {
	return _FeeOracleV2.Contract.ExecGasToken(&_FeeOracleV2.CallOpts, chainId)
}

// ExecGasToken is a free data retrieval call binding the contract method 0xb984cc0b.
//
// Solidity: function execGasToken(uint64 chainId) view returns(uint16)
func (_FeeOracleV2 *FeeOracleV2CallerSession) ExecGasToken(chainId uint64) (uint16, error) {
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
// Solidity: function feeParams(uint64 chainId) view returns((uint16,uint32,uint64,uint64,uint64))
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
// Solidity: function feeParams(uint64 chainId) view returns((uint16,uint32,uint64,uint64,uint64))
func (_FeeOracleV2 *FeeOracleV2Session) FeeParams(chainId uint64) (IFeeOracleV2FeeParams, error) {
	return _FeeOracleV2.Contract.FeeParams(&_FeeOracleV2.CallOpts, chainId)
}

// FeeParams is a free data retrieval call binding the contract method 0x2d4634a4.
//
// Solidity: function feeParams(uint64 chainId) view returns((uint16,uint32,uint64,uint64,uint64))
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

// TokenToNativeRate is a free data retrieval call binding the contract method 0xb17db68a.
//
// Solidity: function tokenToNativeRate(uint16 gasToken) view returns(uint256)
func (_FeeOracleV2 *FeeOracleV2Caller) TokenToNativeRate(opts *bind.CallOpts, gasToken uint16) (*big.Int, error) {
	var out []interface{}
	err := _FeeOracleV2.contract.Call(opts, &out, "tokenToNativeRate", gasToken)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenToNativeRate is a free data retrieval call binding the contract method 0xb17db68a.
//
// Solidity: function tokenToNativeRate(uint16 gasToken) view returns(uint256)
func (_FeeOracleV2 *FeeOracleV2Session) TokenToNativeRate(gasToken uint16) (*big.Int, error) {
	return _FeeOracleV2.Contract.TokenToNativeRate(&_FeeOracleV2.CallOpts, gasToken)
}

// TokenToNativeRate is a free data retrieval call binding the contract method 0xb17db68a.
//
// Solidity: function tokenToNativeRate(uint16 gasToken) view returns(uint256)
func (_FeeOracleV2 *FeeOracleV2CallerSession) TokenToNativeRate(gasToken uint16) (*big.Int, error) {
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

// BulkSetDataCostParams is a paid mutator transaction binding the contract method 0x2105b755.
//
// Solidity: function bulkSetDataCostParams((uint16,uint32,uint64,uint64,uint64)[] params) returns()
func (_FeeOracleV2 *FeeOracleV2Transactor) BulkSetDataCostParams(opts *bind.TransactOpts, params []IFeeOracleV2DataCostParams) (*types.Transaction, error) {
	return _FeeOracleV2.contract.Transact(opts, "bulkSetDataCostParams", params)
}

// BulkSetDataCostParams is a paid mutator transaction binding the contract method 0x2105b755.
//
// Solidity: function bulkSetDataCostParams((uint16,uint32,uint64,uint64,uint64)[] params) returns()
func (_FeeOracleV2 *FeeOracleV2Session) BulkSetDataCostParams(params []IFeeOracleV2DataCostParams) (*types.Transaction, error) {
	return _FeeOracleV2.Contract.BulkSetDataCostParams(&_FeeOracleV2.TransactOpts, params)
}

// BulkSetDataCostParams is a paid mutator transaction binding the contract method 0x2105b755.
//
// Solidity: function bulkSetDataCostParams((uint16,uint32,uint64,uint64,uint64)[] params) returns()
func (_FeeOracleV2 *FeeOracleV2TransactorSession) BulkSetDataCostParams(params []IFeeOracleV2DataCostParams) (*types.Transaction, error) {
	return _FeeOracleV2.Contract.BulkSetDataCostParams(&_FeeOracleV2.TransactOpts, params)
}

// BulkSetFeeParams is a paid mutator transaction binding the contract method 0xdb847a59.
//
// Solidity: function bulkSetFeeParams((uint16,uint32,uint64,uint64,uint64)[] params) returns()
func (_FeeOracleV2 *FeeOracleV2Transactor) BulkSetFeeParams(opts *bind.TransactOpts, params []IFeeOracleV2FeeParams) (*types.Transaction, error) {
	return _FeeOracleV2.contract.Transact(opts, "bulkSetFeeParams", params)
}

// BulkSetFeeParams is a paid mutator transaction binding the contract method 0xdb847a59.
//
// Solidity: function bulkSetFeeParams((uint16,uint32,uint64,uint64,uint64)[] params) returns()
func (_FeeOracleV2 *FeeOracleV2Session) BulkSetFeeParams(params []IFeeOracleV2FeeParams) (*types.Transaction, error) {
	return _FeeOracleV2.Contract.BulkSetFeeParams(&_FeeOracleV2.TransactOpts, params)
}

// BulkSetFeeParams is a paid mutator transaction binding the contract method 0xdb847a59.
//
// Solidity: function bulkSetFeeParams((uint16,uint32,uint64,uint64,uint64)[] params) returns()
func (_FeeOracleV2 *FeeOracleV2TransactorSession) BulkSetFeeParams(params []IFeeOracleV2FeeParams) (*types.Transaction, error) {
	return _FeeOracleV2.Contract.BulkSetFeeParams(&_FeeOracleV2.TransactOpts, params)
}

// BulkSetToNativeRate is a paid mutator transaction binding the contract method 0xf364f687.
//
// Solidity: function bulkSetToNativeRate((uint16,uint256)[] params) returns()
func (_FeeOracleV2 *FeeOracleV2Transactor) BulkSetToNativeRate(opts *bind.TransactOpts, params []IFeeOracleV2ToNativeRateParams) (*types.Transaction, error) {
	return _FeeOracleV2.contract.Transact(opts, "bulkSetToNativeRate", params)
}

// BulkSetToNativeRate is a paid mutator transaction binding the contract method 0xf364f687.
//
// Solidity: function bulkSetToNativeRate((uint16,uint256)[] params) returns()
func (_FeeOracleV2 *FeeOracleV2Session) BulkSetToNativeRate(params []IFeeOracleV2ToNativeRateParams) (*types.Transaction, error) {
	return _FeeOracleV2.Contract.BulkSetToNativeRate(&_FeeOracleV2.TransactOpts, params)
}

// BulkSetToNativeRate is a paid mutator transaction binding the contract method 0xf364f687.
//
// Solidity: function bulkSetToNativeRate((uint16,uint256)[] params) returns()
func (_FeeOracleV2 *FeeOracleV2TransactorSession) BulkSetToNativeRate(params []IFeeOracleV2ToNativeRateParams) (*types.Transaction, error) {
	return _FeeOracleV2.Contract.BulkSetToNativeRate(&_FeeOracleV2.TransactOpts, params)
}

// Initialize is a paid mutator transaction binding the contract method 0x223aacf8.
//
// Solidity: function initialize(address owner_, address manager_, uint96 protocolFee_, (uint16,uint32,uint64,uint64,uint64)[] feeParams_, (uint16,uint32,uint64,uint64,uint64)[] dataCostParams_, (uint16,uint256)[] toNativeRateParams_) returns()
func (_FeeOracleV2 *FeeOracleV2Transactor) Initialize(opts *bind.TransactOpts, owner_ common.Address, manager_ common.Address, protocolFee_ *big.Int, feeParams_ []IFeeOracleV2FeeParams, dataCostParams_ []IFeeOracleV2DataCostParams, toNativeRateParams_ []IFeeOracleV2ToNativeRateParams) (*types.Transaction, error) {
	return _FeeOracleV2.contract.Transact(opts, "initialize", owner_, manager_, protocolFee_, feeParams_, dataCostParams_, toNativeRateParams_)
}

// Initialize is a paid mutator transaction binding the contract method 0x223aacf8.
//
// Solidity: function initialize(address owner_, address manager_, uint96 protocolFee_, (uint16,uint32,uint64,uint64,uint64)[] feeParams_, (uint16,uint32,uint64,uint64,uint64)[] dataCostParams_, (uint16,uint256)[] toNativeRateParams_) returns()
func (_FeeOracleV2 *FeeOracleV2Session) Initialize(owner_ common.Address, manager_ common.Address, protocolFee_ *big.Int, feeParams_ []IFeeOracleV2FeeParams, dataCostParams_ []IFeeOracleV2DataCostParams, toNativeRateParams_ []IFeeOracleV2ToNativeRateParams) (*types.Transaction, error) {
	return _FeeOracleV2.Contract.Initialize(&_FeeOracleV2.TransactOpts, owner_, manager_, protocolFee_, feeParams_, dataCostParams_, toNativeRateParams_)
}

// Initialize is a paid mutator transaction binding the contract method 0x223aacf8.
//
// Solidity: function initialize(address owner_, address manager_, uint96 protocolFee_, (uint16,uint32,uint64,uint64,uint64)[] feeParams_, (uint16,uint32,uint64,uint64,uint64)[] dataCostParams_, (uint16,uint256)[] toNativeRateParams_) returns()
func (_FeeOracleV2 *FeeOracleV2TransactorSession) Initialize(owner_ common.Address, manager_ common.Address, protocolFee_ *big.Int, feeParams_ []IFeeOracleV2FeeParams, dataCostParams_ []IFeeOracleV2DataCostParams, toNativeRateParams_ []IFeeOracleV2ToNativeRateParams) (*types.Transaction, error) {
	return _FeeOracleV2.Contract.Initialize(&_FeeOracleV2.TransactOpts, owner_, manager_, protocolFee_, feeParams_, dataCostParams_, toNativeRateParams_)
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

// SetBaseBytes is a paid mutator transaction binding the contract method 0x9c742ced.
//
// Solidity: function setBaseBytes(uint64 dataCostId, uint32 newBaseBytes) returns()
func (_FeeOracleV2 *FeeOracleV2Transactor) SetBaseBytes(opts *bind.TransactOpts, dataCostId uint64, newBaseBytes uint32) (*types.Transaction, error) {
	return _FeeOracleV2.contract.Transact(opts, "setBaseBytes", dataCostId, newBaseBytes)
}

// SetBaseBytes is a paid mutator transaction binding the contract method 0x9c742ced.
//
// Solidity: function setBaseBytes(uint64 dataCostId, uint32 newBaseBytes) returns()
func (_FeeOracleV2 *FeeOracleV2Session) SetBaseBytes(dataCostId uint64, newBaseBytes uint32) (*types.Transaction, error) {
	return _FeeOracleV2.Contract.SetBaseBytes(&_FeeOracleV2.TransactOpts, dataCostId, newBaseBytes)
}

// SetBaseBytes is a paid mutator transaction binding the contract method 0x9c742ced.
//
// Solidity: function setBaseBytes(uint64 dataCostId, uint32 newBaseBytes) returns()
func (_FeeOracleV2 *FeeOracleV2TransactorSession) SetBaseBytes(dataCostId uint64, newBaseBytes uint32) (*types.Transaction, error) {
	return _FeeOracleV2.Contract.SetBaseBytes(&_FeeOracleV2.TransactOpts, dataCostId, newBaseBytes)
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

// SetToNativeRate is a paid mutator transaction binding the contract method 0xa12f2c58.
//
// Solidity: function setToNativeRate(uint16 gasToken, uint256 nativeRate) returns()
func (_FeeOracleV2 *FeeOracleV2Transactor) SetToNativeRate(opts *bind.TransactOpts, gasToken uint16, nativeRate *big.Int) (*types.Transaction, error) {
	return _FeeOracleV2.contract.Transact(opts, "setToNativeRate", gasToken, nativeRate)
}

// SetToNativeRate is a paid mutator transaction binding the contract method 0xa12f2c58.
//
// Solidity: function setToNativeRate(uint16 gasToken, uint256 nativeRate) returns()
func (_FeeOracleV2 *FeeOracleV2Session) SetToNativeRate(gasToken uint16, nativeRate *big.Int) (*types.Transaction, error) {
	return _FeeOracleV2.Contract.SetToNativeRate(&_FeeOracleV2.TransactOpts, gasToken, nativeRate)
}

// SetToNativeRate is a paid mutator transaction binding the contract method 0xa12f2c58.
//
// Solidity: function setToNativeRate(uint16 gasToken, uint256 nativeRate) returns()
func (_FeeOracleV2 *FeeOracleV2TransactorSession) SetToNativeRate(gasToken uint16, nativeRate *big.Int) (*types.Transaction, error) {
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

// FeeOracleV2BaseBytesSetIterator is returned from FilterBaseBytesSet and is used to iterate over the raw logs and unpacked data for BaseBytesSet events raised by the FeeOracleV2 contract.
type FeeOracleV2BaseBytesSetIterator struct {
	Event *FeeOracleV2BaseBytesSet // Event containing the contract specifics and raw log

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
func (it *FeeOracleV2BaseBytesSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(FeeOracleV2BaseBytesSet)
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
		it.Event = new(FeeOracleV2BaseBytesSet)
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
func (it *FeeOracleV2BaseBytesSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *FeeOracleV2BaseBytesSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// FeeOracleV2BaseBytesSet represents a BaseBytesSet event raised by the FeeOracleV2 contract.
type FeeOracleV2BaseBytesSet struct {
	DataCostId uint64
	BaseBytes  uint32
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterBaseBytesSet is a free log retrieval operation binding the contract event 0x406ba3ab8c23e58f883620734592e662ac0018177590c1f9f27afc58a4b7b583.
//
// Solidity: event BaseBytesSet(uint64 dataCostId, uint32 baseBytes)
func (_FeeOracleV2 *FeeOracleV2Filterer) FilterBaseBytesSet(opts *bind.FilterOpts) (*FeeOracleV2BaseBytesSetIterator, error) {

	logs, sub, err := _FeeOracleV2.contract.FilterLogs(opts, "BaseBytesSet")
	if err != nil {
		return nil, err
	}
	return &FeeOracleV2BaseBytesSetIterator{contract: _FeeOracleV2.contract, event: "BaseBytesSet", logs: logs, sub: sub}, nil
}

// WatchBaseBytesSet is a free log subscription operation binding the contract event 0x406ba3ab8c23e58f883620734592e662ac0018177590c1f9f27afc58a4b7b583.
//
// Solidity: event BaseBytesSet(uint64 dataCostId, uint32 baseBytes)
func (_FeeOracleV2 *FeeOracleV2Filterer) WatchBaseBytesSet(opts *bind.WatchOpts, sink chan<- *FeeOracleV2BaseBytesSet) (event.Subscription, error) {

	logs, sub, err := _FeeOracleV2.contract.WatchLogs(opts, "BaseBytesSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(FeeOracleV2BaseBytesSet)
				if err := _FeeOracleV2.contract.UnpackLog(event, "BaseBytesSet", log); err != nil {
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

// ParseBaseBytesSet is a log parse operation binding the contract event 0x406ba3ab8c23e58f883620734592e662ac0018177590c1f9f27afc58a4b7b583.
//
// Solidity: event BaseBytesSet(uint64 dataCostId, uint32 baseBytes)
func (_FeeOracleV2 *FeeOracleV2Filterer) ParseBaseBytesSet(log types.Log) (*FeeOracleV2BaseBytesSet, error) {
	event := new(FeeOracleV2BaseBytesSet)
	if err := _FeeOracleV2.contract.UnpackLog(event, "BaseBytesSet", log); err != nil {
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
	GasToken   uint16
	BaseBytes  uint32
	Id         uint64
	GasPrice   uint64
	GasPerByte uint64
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterDataCostParamsSet is a free log retrieval operation binding the contract event 0xd143a0934cc5e5337dca3eb0afa1e7f86680796ca6f132bfe7e0828b7155bd40.
//
// Solidity: event DataCostParamsSet(uint16 gasToken, uint32 baseBytes, uint64 id, uint64 gasPrice, uint64 gasPerByte)
func (_FeeOracleV2 *FeeOracleV2Filterer) FilterDataCostParamsSet(opts *bind.FilterOpts) (*FeeOracleV2DataCostParamsSetIterator, error) {

	logs, sub, err := _FeeOracleV2.contract.FilterLogs(opts, "DataCostParamsSet")
	if err != nil {
		return nil, err
	}
	return &FeeOracleV2DataCostParamsSetIterator{contract: _FeeOracleV2.contract, event: "DataCostParamsSet", logs: logs, sub: sub}, nil
}

// WatchDataCostParamsSet is a free log subscription operation binding the contract event 0xd143a0934cc5e5337dca3eb0afa1e7f86680796ca6f132bfe7e0828b7155bd40.
//
// Solidity: event DataCostParamsSet(uint16 gasToken, uint32 baseBytes, uint64 id, uint64 gasPrice, uint64 gasPerByte)
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

// ParseDataCostParamsSet is a log parse operation binding the contract event 0xd143a0934cc5e5337dca3eb0afa1e7f86680796ca6f132bfe7e0828b7155bd40.
//
// Solidity: event DataCostParamsSet(uint16 gasToken, uint32 baseBytes, uint64 id, uint64 gasPrice, uint64 gasPerByte)
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
	GasToken     uint16
	BaseGasLimit uint32
	ChainId      uint64
	GasPrice     uint64
	DataCostId   uint64
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterFeeParamsSet is a free log retrieval operation binding the contract event 0x600e7ff14e74285e17debda1fee2df93741c6518e12ba908ba4417c0610974a9.
//
// Solidity: event FeeParamsSet(uint16 gasToken, uint32 baseGasLimit, uint64 chainId, uint64 gasPrice, uint64 dataCostId)
func (_FeeOracleV2 *FeeOracleV2Filterer) FilterFeeParamsSet(opts *bind.FilterOpts) (*FeeOracleV2FeeParamsSetIterator, error) {

	logs, sub, err := _FeeOracleV2.contract.FilterLogs(opts, "FeeParamsSet")
	if err != nil {
		return nil, err
	}
	return &FeeOracleV2FeeParamsSetIterator{contract: _FeeOracleV2.contract, event: "FeeParamsSet", logs: logs, sub: sub}, nil
}

// WatchFeeParamsSet is a free log subscription operation binding the contract event 0x600e7ff14e74285e17debda1fee2df93741c6518e12ba908ba4417c0610974a9.
//
// Solidity: event FeeParamsSet(uint16 gasToken, uint32 baseGasLimit, uint64 chainId, uint64 gasPrice, uint64 dataCostId)
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

// ParseFeeParamsSet is a log parse operation binding the contract event 0x600e7ff14e74285e17debda1fee2df93741c6518e12ba908ba4417c0610974a9.
//
// Solidity: event FeeParamsSet(uint16 gasToken, uint32 baseGasLimit, uint64 chainId, uint64 gasPrice, uint64 dataCostId)
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
	GasToken   uint16
	NativeRate *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterToNativeRateSet is a free log retrieval operation binding the contract event 0x770404c81b5a366795e9a06ff3969e3fb2eb0515b7282e986f261c180b97623d.
//
// Solidity: event ToNativeRateSet(uint16 gasToken, uint256 nativeRate)
func (_FeeOracleV2 *FeeOracleV2Filterer) FilterToNativeRateSet(opts *bind.FilterOpts) (*FeeOracleV2ToNativeRateSetIterator, error) {

	logs, sub, err := _FeeOracleV2.contract.FilterLogs(opts, "ToNativeRateSet")
	if err != nil {
		return nil, err
	}
	return &FeeOracleV2ToNativeRateSetIterator{contract: _FeeOracleV2.contract, event: "ToNativeRateSet", logs: logs, sub: sub}, nil
}

// WatchToNativeRateSet is a free log subscription operation binding the contract event 0x770404c81b5a366795e9a06ff3969e3fb2eb0515b7282e986f261c180b97623d.
//
// Solidity: event ToNativeRateSet(uint16 gasToken, uint256 nativeRate)
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

// ParseToNativeRateSet is a log parse operation binding the contract event 0x770404c81b5a366795e9a06ff3969e3fb2eb0515b7282e986f261c180b97623d.
//
// Solidity: event ToNativeRateSet(uint16 gasToken, uint256 nativeRate)
func (_FeeOracleV2 *FeeOracleV2Filterer) ParseToNativeRateSet(log types.Log) (*FeeOracleV2ToNativeRateSet, error) {
	event := new(FeeOracleV2ToNativeRateSet)
	if err := _FeeOracleV2.contract.UnpackLog(event, "ToNativeRateSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
