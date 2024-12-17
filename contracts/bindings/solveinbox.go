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

// SolveCall is an auto generated low-level Go binding around an user-defined struct.
type SolveCall struct {
	DestChainId uint64
	Target      common.Address
	Value       *big.Int
	Data        []byte
}

// SolveDeposit is an auto generated low-level Go binding around an user-defined struct.
type SolveDeposit struct {
	IsNative bool
	Token    common.Address
	Amount   *big.Int
}

// SolveRequest is an auto generated low-level Go binding around an user-defined struct.
type SolveRequest struct {
	Id         [32]byte
	Status     uint8
	From       common.Address
	AcceptedBy common.Address
	Call       SolveCall
	Deposits   []SolveDeposit
	History    []SolveStatusUpdate
}

// SolveStatusUpdate is an auto generated low-level Go binding around an user-defined struct.
type SolveStatusUpdate struct {
	Status    uint8
	Timestamp *big.Int
}

// SolveTokenDeposit is an auto generated low-level Go binding around an user-defined struct.
type SolveTokenDeposit struct {
	Token  common.Address
	Amount *big.Int
}

// SolveInboxMetaData contains all meta data concerning the SolveInbox contract.
var SolveInboxMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"accept\",\"inputs\":[{\"name\":\"id\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"cancel\",\"inputs\":[{\"name\":\"id\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"cancelOwnershipHandover\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"claim\",\"inputs\":[{\"name\":\"id\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"completeOwnershipHandover\",\"inputs\":[{\"name\":\"pendingOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"defaultConfLevel\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"deployedAt\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getLatestRequestByStatus\",\"inputs\":[{\"name\":\"status\",\"type\":\"uint8\",\"internalType\":\"enumSolve.Status\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structSolve.Request\",\"components\":[{\"name\":\"id\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"updatedAt\",\"type\":\"uint40\",\"internalType\":\"uint40\"},{\"name\":\"status\",\"type\":\"uint8\",\"internalType\":\"enumSolve.Status\"},{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"acceptedBy\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"call\",\"type\":\"tuple\",\"internalType\":\"structSolve.Call\",\"components\":[{\"name\":\"destChainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"deposits\",\"type\":\"tuple[]\",\"internalType\":\"structSolve.Deposit[]\",\"components\":[{\"name\":\"isNative\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRequest\",\"inputs\":[{\"name\":\"id\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structSolve.Request\",\"components\":[{\"name\":\"id\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"updatedAt\",\"type\":\"uint40\",\"internalType\":\"uint40\"},{\"name\":\"status\",\"type\":\"uint8\",\"internalType\":\"enumSolve.Status\"},{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"acceptedBy\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"call\",\"type\":\"tuple\",\"internalType\":\"structSolve.Call\",\"components\":[{\"name\":\"destChainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"deposits\",\"type\":\"tuple[]\",\"internalType\":\"structSolve.Deposit[]\",\"components\":[{\"name\":\"isNative\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"grantRoles\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"roles\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"hasAllRoles\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"roles\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"hasAnyRole\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"roles\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"owner_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"solver_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"omni_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"outbox_\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"markFulfilled\",\"inputs\":[{\"name\":\"id\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"callHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"omni\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIOmniPortal\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"result\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"ownershipHandoverExpiresAt\",\"inputs\":[{\"name\":\"pendingOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"result\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"reject\",\"inputs\":[{\"name\":\"id\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"reason\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"renounceRoles\",\"inputs\":[{\"name\":\"roles\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"request\",\"inputs\":[{\"name\":\"call\",\"type\":\"tuple\",\"internalType\":\"structSolve.Call\",\"components\":[{\"name\":\"destChainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"deposits\",\"type\":\"tuple[]\",\"internalType\":\"structSolve.TokenDeposit[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[{\"name\":\"id\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"requestOwnershipHandover\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"revokeRoles\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"roles\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"rolesOf\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"roles\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"event\",\"name\":\"Accepted\",\"inputs\":[{\"name\":\"id\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"by\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Claimed\",\"inputs\":[{\"name\":\"id\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"by\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"deposits\",\"type\":\"tuple[]\",\"indexed\":false,\"internalType\":\"structSolve.Deposit[]\",\"components\":[{\"name\":\"isNative\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DefaultConfLevelSet\",\"inputs\":[{\"name\":\"conf\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Fulfilled\",\"inputs\":[{\"name\":\"id\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"callHash\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"creditedTo\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OmniPortalSet\",\"inputs\":[{\"name\":\"omni\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipHandoverCanceled\",\"inputs\":[{\"name\":\"pendingOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipHandoverRequested\",\"inputs\":[{\"name\":\"pendingOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"oldOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Rejected\",\"inputs\":[{\"name\":\"id\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"by\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"reason\",\"type\":\"uint8\",\"indexed\":true,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Requested\",\"inputs\":[{\"name\":\"id\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"call\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structSolve.Call\",\"components\":[{\"name\":\"destChainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"deposits\",\"type\":\"tuple[]\",\"indexed\":false,\"internalType\":\"structSolve.Deposit[]\",\"components\":[{\"name\":\"isNative\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Reverted\",\"inputs\":[{\"name\":\"id\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RolesUpdated\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"roles\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AlreadyInitialized\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidDeposit\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidRecipient\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NewOwnerIsZeroAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NoDeposits\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NoHandoverRequest\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotAccepted\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotFulfilled\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotOutbox\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotPending\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotPendingOrRejected\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"Reentrancy\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"TransferFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"Unauthorized\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"WrongCallHash\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"WrongSourceChain\",\"inputs\":[]}]",
	Bin: "0x60a06040523480156200001157600080fd5b5063ffffffff60643b1615620000975760646001600160a01b031663a3b1b31d6040518163ffffffff1660e01b8152600401602060405180830381865afa9250505080156200007f575060408051601f3d908101601f191682019092526200007c9181019062000110565b60015b6200008e57436080526200009c565b6080526200009c565b436080525b620000a6620000ac565b6200012a565b63409feecd1980546001811615620000cc5763f92ee8a96000526004601cfd5b8160c01c808260011c146200010b578060011b8355806020527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2602080a15b505050565b6000602082840312156200012357600080fd5b5051919050565b60805161236e6200014660003960006103f8015261236e6000f3fe6080604052600436106101665760003560e01c80637667fd69116100d1578063e4725ba11161008a578063f2fde38b11610064578063f2fde38b1461042d578063f8c8765e14610440578063fb1e61ca14610460578063fee81cf41461048057600080fd5b8063e4725ba1146103c6578063eae4c19f146103e6578063f04e283e1461041a57600080fd5b80637667fd691461030d5780638da5cb5b1461033a57806396c144f014610353578063c4d252f514610373578063db3ea55314610393578063e17771c8146103b357600080fd5b806339acf9f11161012357806339acf9f1146102485780634a4ee7b114610280578063514e62fc1461029357806354d1f13d146102ca578063715018a6146102d257806374eeb847146102da57600080fd5b8063183a4f6e1461016b5780631c10893f146101805780631cd64df41461019357806325692962146101df5780632de94807146101e7578063337ffe2014610228575b600080fd5b61017e610179366004611ac3565b6104b3565b005b61017e61018e366004611af1565b6104c0565b34801561019f57600080fd5b506101ca6101ae366004611af1565b638b78c6d8600c90815260009290925260209091205481161490565b60405190151581526020015b60405180910390f35b61017e6104d6565b3480156101f357600080fd5b5061021a610202366004611b1d565b638b78c6d8600c908152600091909152602090205490565b6040519081526020016101d6565b34801561023457600080fd5b5061017e610243366004611b41565b610525565b34801561025457600080fd5b50600054610268906001600160a01b031681565b6040516001600160a01b0390911681526020016101d6565b61017e61028e366004611af1565b61078f565b34801561029f57600080fd5b506101ca6102ae366004611af1565b638b78c6d8600c90815260009290925260209091205416151590565b61017e6107a1565b61017e6107dd565b3480156102e657600080fd5b506000546102fb90600160a01b900460ff1681565b60405160ff90911681526020016101d6565b34801561031957600080fd5b5061032d610328366004611b63565b6107f1565b6040516101d69190611c79565b34801561034657600080fd5b50638b78c6d81954610268565b34801561035f57600080fd5b5061017e61036e366004611d31565b610a30565b34801561037f57600080fd5b5061017e61038e366004611ac3565b610c12565b34801561039f57600080fd5b5061017e6103ae366004611d61565b610e07565b61021a6103c1366004611d8c565b610f1f565b3480156103d257600080fd5b5061017e6103e1366004611ac3565b611093565b3480156103f257600080fd5b5061021a7f000000000000000000000000000000000000000000000000000000000000000081565b61017e610428366004611b1d565b6111bc565b61017e61043b366004611b1d565b6111f9565b34801561044c57600080fd5b5061017e61045b366004611e2d565b611220565b34801561046c57600080fd5b5061032d61047b366004611ac3565b6112cc565b34801561048c57600080fd5b5061021a61049b366004611b1d565b63389a75e1600c908152600091909152602090205490565b6104bd3382611328565b50565b6104c8611334565b6104d2828261134f565b5050565b60006202a3006001600160401b03164201905063389a75e1600c5233600052806020600c2055337fdbf36a107da19e49527a7176a1babf963b4b0ff8cde35ee35d6cd8f1f9ac7e1d600080a250565b60005460408051631799380760e11b815281516001600160a01b0390931692632f32700e926004808401939192918290030181865afa15801561056c573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906105909190611eb4565b8051600180546020909301516001600160a01b0316600160401b026001600160e01b03199093166001600160401b039092169190911791909117905568929eee149b4bd21268543090036105ec5763ab143c066000526004601cfd5b3068929eee149b4bd2126855600082815260046020526040902060026001820154600160281b900460ff16600681111561062857610628611b84565b146106465760405163029d79a560e41b815260040160405180910390fd5b600354600154600160401b90046001600160a01b0390811691161461067e5760405163bda8fc9560e01b815260040160405180910390fd5b60038101546001546001600160401b039081169116146106b157604051633687f39960e21b815260040160405180910390fd5b6106bf83468360030161135b565b82146106de57604051631c6060ab60e11b815260040160405180910390fd5b60018101805464ffffffffff421665ffffffffffff199091161765050000000000179055600560008181526020919091527f458b30c2d72bfd2c6317304a4594ecbafe5f729d3111b65fdc3a33bd48e5432d84905560028201546040516001600160a01b0390911691849186917f7898a125e0970666c80e00bbf2e7041d84dfe5bbe6bcf562ce53d540fd6cd89191a450503868929eee149b4bd212685550600180546001600160e01b0319169055565b610797611334565b6104d28282611328565b63389a75e1600c523360005260006020600c2055337ffa7b8eab7da67f412cc9575ed43464468f9bfbae89d1675917346ca6d8fe3c92600080a2565b6107e5611334565b6107ef6000611391565b565b6107f9611a39565b600460006005600085600681111561081357610813611b84565b600681111561082457610824611b84565b81526020808201929092526040908101600090812054845283830194909452918201909220815160e08101835281548152600182015464ffffffffff811694820194909452929091830190600160281b900460ff16600681111561088a5761088a611b84565b600681111561089b5761089b611b84565b815260018201546001600160a01b03600160301b90910481166020808401919091526002840154821660408085019190915280516080810182526003860180546001600160401b0381168352600160401b9004909416928101929092526004850154908201526005840180546060948501949293928401919061091d90611f12565b80601f016020809104026020016040519081016040528092919081815260200182805461094990611f12565b80156109965780601f1061096b57610100808354040283529160200191610996565b820191906000526020600020905b81548152906001019060200180831161097957829003601f168201915b505050505081525050815260200160068201805480602002602001604051908101604052809291908181526020016000905b82821015610a225760008481526020908190206040805160608101825260028602909201805460ff81161515845261010090046001600160a01b0316838501526001908101549183019190915290835290920191016109c8565b505050915250909392505050565b3068929eee149b4bd212685403610a4f5763ab143c066000526004601cfd5b3068929eee149b4bd2126855600082815260046020526040902060056001820154600160281b900460ff166006811115610a8b57610a8b611b84565b14610aa9576040516303de361f60e11b815260040160405180910390fd5b60028101546001600160a01b03163314610ad5576040516282b42960e81b815260040160405180910390fd5b60018101805464ffffffffff421665ffffffffffff19909116176506000000000017905560066000818152600560209081527f069400f22b28c6c362558d92f66163cec5671cba50b61abd2eecfcd0eaeac51886905591830180546040805182860281018601909152818152610bb2948794919392919084015b82821015610ba95760008481526020908190206040805160608101825260028602909201805460ff81161515845261010090046001600160a01b031683850152600190810154918301919091529083529092019101610b4f565b505050506113cf565b816001600160a01b0316336001600160a01b0316847f1291cec4ea55732527f1ae8ff3e53def0dbe1feb99c094983449e59cdce4674b84600601604051610bf99190611fa5565b60405180910390a4503868929eee149b4bd21268555050565b3068929eee149b4bd212685403610c315763ab143c066000526004601cfd5b3068929eee149b4bd21268556000818152600460205260409020600180820154600160281b900460ff166006811115610c6c57610c6c611b84565b14158015610c9a575060036001820154600160281b900460ff166006811115610c9757610c97611b84565b14155b15610cb857604051631fa4148760e21b815260040160405180910390fd5b6001810154600160301b90046001600160a01b03163314610ceb576040516282b42960e81b815260040160405180910390fd5b6001810180546504000000000065ffffffffffff199091164264ffffffffff161717815560046000908152600560209081527f3eec716f11ba9e820c81ca75eb978ffb45831ef8b7a53e5e422c26008e1ca6d585905591546006840180546040805182870281018701909152818152610dcc95600160301b9094046001600160a01b031694909384018215610ba95760008481526020908190206040805160608101825260028602909201805460ff81161515845261010090046001600160a01b031683850152600190810154918301919091529083529092019101610b4f565b60405182907fb66b13449e4bb2c30749a37f3081f1988fcee5ff5d98ce740b354d4e2d94409590600090a2503868929eee149b4bd212685550565b6001610e1281611521565b3068929eee149b4bd212685403610e315763ab143c066000526004601cfd5b3068929eee149b4bd21268556000838152600460205260409020600180820154600160281b900460ff166006811115610e6c57610e6c611b84565b14610e8a57604051633ee3282d60e11b815260040160405180910390fd5b60018101805464ffffffffff421665ffffffffffff1990911617650300000000001790556003600090815260056020527fa9bc9a3a348c357ba16b37005d7e6b3236198c0e939f4af8c5f19b8deeb8ebc085905560405160ff851691339187917f21f84ee3a6e9bc7c10f855f8c9829e22c613861cef10add09eccdbc88df9f59f91a4503868929eee149b4bd2126855505050565b60003068929eee149b4bd212685403610f405763ab143c066000526004601cfd5b3068929eee149b4bd21268556000610f5e6040860160208701611b1d565b6001600160a01b031603610f855760405163574b16a760e11b815260040160405180910390fd5b610f926020850185611fb8565b6001600160401b0316600003610fbb5760405163574b16a760e11b815260040160405180910390fd5b610fc86060850185611fd5565b9050600003610fea5760405163574b16a760e11b815260040160405180910390fd5b81158015610ff6575034155b1561101457604051630558800760e21b815260040160405180910390fd5b600061102233868686611547565b60018101548154604051929350600160301b9091046001600160a01b0316917f3622a563ce1f96b477fa827bf0b60be8aee346b515754f3645f4d7a2ef5d4b299061107690600386019060068701906120e7565b60405180910390a3543868929eee149b4bd2126855949350505050565b600161109e81611521565b3068929eee149b4bd2126854036110bd5763ab143c066000526004601cfd5b3068929eee149b4bd21268556000828152600460205260409020600180820154600160281b900460ff1660068111156110f8576110f8611b84565b1461111657604051633ee3282d60e11b815260040160405180910390fd5b60018101805464ffffffffff421665ffffffffffff199091161765020000000000179055600280820180546001600160a01b03191633908117909155600091825260056020527f89832631fb3c3307a103ba2c84ab569c64d6182a18893dcd163f0f1c2090733a859055604051909185917f9deed34441ca75bb2dbbe101d2201930f40e18a9ce521c77fbdca6690a8999679190a3503868929eee149b4bd21268555050565b6111c4611334565b63389a75e1600c52806000526020600c2080544211156111ec57636f5e88186000526004601cfd5b600090556104bd81611391565b611201611334565b8060601b61121757637448fbae6000526004601cfd5b6104bd81611391565b63409feecd1980546003825580156112575760018160011c14303b1061124e5763f92ee8a96000526004601cfd5b818160ff1b1b91505b5061126185611835565b61126c84600161134f565b61127583611871565b600380546001600160a01b0319166001600160a01b03841617905580156112c5576002815560016020527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2602080a15b5050505050565b6112d4611a39565b600082815260046020908152604091829020825160e08101845281548152600182015464ffffffffff81169382019390935292909190830190600160281b900460ff16600681111561088a5761088a611b84565b6104d282826000611914565b638b78c6d8195433146107ef576382b429006000526004601cfd5b6104d282826001611914565b60008383836040516020016113729392919061210c565b6040516020818303038152906040528051906020012090509392505050565b638b78c6d81980546001600160a01b039092169182907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0600080a355565b6001600160a01b0382166113f657604051634e46966960e11b815260040160405180910390fd5b60005b815181101561151c5781818151811061141457611414612134565b602002602001015160000151156114bc576000836001600160a01b031683838151811061144357611443612134565b60200260200101516040015160405160006040518083038185875af1925050503d806000811461148f576040519150601f19603f3d011682016040523d82523d6000602084013e611494565b606091505b50509050806114b6576040516312171d8360e31b815260040160405180910390fd5b50611514565b611514838383815181106114d2576114d2612134565b6020026020010151604001518484815181106114f0576114f0612134565b6020026020010151602001516001600160a01b031661196d9092919063ffffffff16565b6001016113f9565b505050565b638b78c6d8600c5233600052806020600c2054166104bd576382b429006000526004601cfd5b6000806115526119bd565b600081815260046020526040902081815560018101805464ffffffffff421665ffffffffffff1990911617600160281b176601000000000000600160d01b031916600160301b6001600160a01b038b16021790559250905084600383016115b98282612259565b5050600160005260056020527f1471eb6eb2c5e789fc3de43f8ce62938c7d1836ec861730447e2ada8fd81017b819055341561166857604080516060810182526001808252600060208084018281523495850195865260068801805480860182559084529190922093516002909102909301805491516001600160a81b0319909216931515610100600160a81b031916939093176101006001600160a01b039092169190910217825591519101555b60005b8381101561182b5784848281811061168557611685612134565b905060400201602001356000036116af57604051635972996f60e11b815260040160405180910390fd5b60008585838181106116c3576116c3612134565b6116d99260206040909202019081019150611b1d565b6001600160a01b03160361170057604051635972996f60e11b815260040160405180910390fd5b82600601604051806060016040528060001515815260200187878581811061172a5761172a612134565b6117409260206040909202019081019150611b1d565b6001600160a01b0316815260200187878581811061176057611760612134565b60206040918202939093018301359093525083546001818101865560009586529482902084516002909202018054928501516001600160a01b031661010002610100600160a81b0319921515929092166001600160a81b03199093169290921717815591015191015561182333308787858181106117e0576117e0612134565b905060400201602001358888868181106117fc576117fc612134565b6118129260206040909202019081019150611b1d565b6001600160a01b03169291906119db565b60010161166b565b5050949350505050565b6001600160a01b0316638b78c6d8198190558060007f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e08180a350565b6001600160a01b0381166118c05760405162461bcd60e51b8152602060048201526012602482015271584170703a206e6f207a65726f206f6d6e6960701b604482015260640160405180910390fd5b600080546001600160a01b0319166001600160a01b0383169081179091556040519081527f79162c8d053a07e70cdc1ccc536f0440b571f8508377d2bef51094fadab98f479060200160405180910390a150565b638b78c6d8600c52826000526020600c20805483811783611936575080841681185b80835580600c5160601c7f715ad5ce61fc9595c7b415289d59cf203f23a94fa06f04af7e489a0a76e1fe26600080a3505050505050565b816014528060345263a9059cbb60601b60005260206000604460106000875af180600160005114166119b257803d853b1517106119b2576390b8ec186000526004601cfd5b506000603452505050565b60028054600091826119ce83612311565b9091555050600254919050565b60405181606052826040528360601b602c526323b872dd60601b600c52602060006064601c6000895af18060016000511416611a2a57803d873b151710611a2a57637939f4246000526004601cfd5b50600060605260405250505050565b6040805160e0810182526000808252602082018190529091820190815260200160006001600160a01b0316815260200160006001600160a01b03168152602001611ab6604051806080016040528060006001600160401b0316815260200160006001600160a01b0316815260200160008152602001606081525090565b8152602001606081525090565b600060208284031215611ad557600080fd5b5035919050565b6001600160a01b03811681146104bd57600080fd5b60008060408385031215611b0457600080fd5b8235611b0f81611adc565b946020939093013593505050565b600060208284031215611b2f57600080fd5b8135611b3a81611adc565b9392505050565b60008060408385031215611b5457600080fd5b50508035926020909101359150565b600060208284031215611b7557600080fd5b813560078110611b3a57600080fd5b634e487b7160e01b600052602160045260246000fd5b6001600160401b0381511682526000602060018060a01b03602084015116602085015260408301516040850152606083015160806060860152805180608087015260005b81811015611bfa5782810184015187820160a001528301611bde565b50600060a0828801015260a0601f19601f830116870101935050505092915050565b60008151808452602080850194506020840160005b83811015611c6e578151805115158852838101516001600160a01b0316848901526040908101519088015260609096019590820190600101611c31565b509495945050505050565b602081528151602082015264ffffffffff60208301511660408201526000604083015160078110611cba57634e487b7160e01b600052602160045260246000fd5b806060840152506060830151611cdb60808401826001600160a01b03169052565b5060808301516001600160a01b03811660a08401525060a083015160e060c0840152611d0b610100840182611b9a565b905060c0840151601f198483030160e0850152611d288282611c1c565b95945050505050565b60008060408385031215611d4457600080fd5b823591506020830135611d5681611adc565b809150509250929050565b60008060408385031215611d7457600080fd5b82359150602083013560ff81168114611d5657600080fd5b600080600060408486031215611da157600080fd5b83356001600160401b0380821115611db857600080fd5b9085019060808288031215611dcc57600080fd5b90935060208501359080821115611de257600080fd5b818601915086601f830112611df657600080fd5b813581811115611e0557600080fd5b8760208260061b8501011115611e1a57600080fd5b6020830194508093505050509250925092565b60008060008060808587031215611e4357600080fd5b8435611e4e81611adc565b93506020850135611e5e81611adc565b92506040850135611e6e81611adc565b91506060850135611e7e81611adc565b939692955090935050565b634e487b7160e01b600052604160045260246000fd5b6001600160401b03811681146104bd57600080fd5b600060408284031215611ec657600080fd5b604051604081018181106001600160401b0382111715611ee857611ee8611e89565b6040528251611ef681611e9f565b81526020830151611f0681611adc565b60208201529392505050565b600181811c90821680611f2657607f821691505b602082108103611f4657634e487b7160e01b600052602260045260246000fd5b50919050565b600081548084526020808501945083600052602060002060005b83811015611c6e57815460ff81161515885260081c6001600160a01b031683880152600180830154604089015260609097019660029092019101611f66565b602081526000611b3a6020830184611f4c565b600060208284031215611fca57600080fd5b8135611b3a81611e9f565b6000808335601e19843603018112611fec57600080fd5b8301803591506001600160401b0382111561200657600080fd5b60200191503681900382131561201b57600080fd5b9250929050565b600081546001600160401b0381168452602060018060a01b038260401c166020860152600191506001840154604086015260028401608060608701526000815461206b81611f12565b8060808a015260a0600183166000811461208c57600181146120a8576120d8565b60ff19841660a08c015260a083151560051b8c010194506120d8565b85600052602060002060005b848110156120cf5781548d82018501529089019088016120b4565b8c0160a0019550505b50929998505050505050505050565b6040815260006120fa6040830185612022565b8281036020840152611d288185611f4c565b8381526001600160401b0383166020820152606060408201526000611d286060830184612022565b634e487b7160e01b600052603260045260246000fd5b601f82111561151c576000816000526020600020601f850160051c810160208610156121735750805b601f850160051c820191505b818110156121925782815560010161217f565b505050505050565b6001600160401b038311156121b1576121b1611e89565b6121c5836121bf8354611f12565b8361214a565b6000601f8411600181146121f957600085156121e15750838201355b600019600387901b1c1916600186901b1783556112c5565b600083815260209020601f19861690835b8281101561222a578685013582556020948501946001909201910161220a565b50868210156122475760001960f88860031b161c19848701351681555b505060018560011b0183555050505050565b813561226481611e9f565b6001600160401b0380821691508254826001600160401b03198216178455602085013561229081611adc565b6001600160e01b031991909116909217604092831b68010000000000000000600160e01b031617835590830135600183015560608301359036849003601e190182126122db57600080fd5b908301908135818111156122ee57600080fd5b60208301925080360383131561230357600080fd5b6112c581846002870161219a565b60006001820161233157634e487b7160e01b600052601160045260246000fd5b506001019056fea264697066735822122007ec7e45f3d811cb159f4a7b9a03656f5aae204b984cf866160cd0c3c83748a564736f6c63430008180033",
}

// SolveInboxABI is the input ABI used to generate the binding from.
// Deprecated: Use SolveInboxMetaData.ABI instead.
var SolveInboxABI = SolveInboxMetaData.ABI

// SolveInboxBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SolveInboxMetaData.Bin instead.
var SolveInboxBin = SolveInboxMetaData.Bin

// DeploySolveInbox deploys a new Ethereum contract, binding an instance of SolveInbox to it.
func DeploySolveInbox(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SolveInbox, error) {
	parsed, err := SolveInboxMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SolveInboxBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SolveInbox{SolveInboxCaller: SolveInboxCaller{contract: contract}, SolveInboxTransactor: SolveInboxTransactor{contract: contract}, SolveInboxFilterer: SolveInboxFilterer{contract: contract}}, nil
}

// SolveInbox is an auto generated Go binding around an Ethereum contract.
type SolveInbox struct {
	SolveInboxCaller     // Read-only binding to the contract
	SolveInboxTransactor // Write-only binding to the contract
	SolveInboxFilterer   // Log filterer for contract events
}

// SolveInboxCaller is an auto generated read-only Go binding around an Ethereum contract.
type SolveInboxCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SolveInboxTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SolveInboxTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SolveInboxFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SolveInboxFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SolveInboxSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SolveInboxSession struct {
	Contract     *SolveInbox       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SolveInboxCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SolveInboxCallerSession struct {
	Contract *SolveInboxCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// SolveInboxTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SolveInboxTransactorSession struct {
	Contract     *SolveInboxTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// SolveInboxRaw is an auto generated low-level Go binding around an Ethereum contract.
type SolveInboxRaw struct {
	Contract *SolveInbox // Generic contract binding to access the raw methods on
}

// SolveInboxCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SolveInboxCallerRaw struct {
	Contract *SolveInboxCaller // Generic read-only contract binding to access the raw methods on
}

// SolveInboxTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SolveInboxTransactorRaw struct {
	Contract *SolveInboxTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSolveInbox creates a new instance of SolveInbox, bound to a specific deployed contract.
func NewSolveInbox(address common.Address, backend bind.ContractBackend) (*SolveInbox, error) {
	contract, err := bindSolveInbox(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SolveInbox{SolveInboxCaller: SolveInboxCaller{contract: contract}, SolveInboxTransactor: SolveInboxTransactor{contract: contract}, SolveInboxFilterer: SolveInboxFilterer{contract: contract}}, nil
}

// NewSolveInboxCaller creates a new read-only instance of SolveInbox, bound to a specific deployed contract.
func NewSolveInboxCaller(address common.Address, caller bind.ContractCaller) (*SolveInboxCaller, error) {
	contract, err := bindSolveInbox(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SolveInboxCaller{contract: contract}, nil
}

// NewSolveInboxTransactor creates a new write-only instance of SolveInbox, bound to a specific deployed contract.
func NewSolveInboxTransactor(address common.Address, transactor bind.ContractTransactor) (*SolveInboxTransactor, error) {
	contract, err := bindSolveInbox(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SolveInboxTransactor{contract: contract}, nil
}

// NewSolveInboxFilterer creates a new log filterer instance of SolveInbox, bound to a specific deployed contract.
func NewSolveInboxFilterer(address common.Address, filterer bind.ContractFilterer) (*SolveInboxFilterer, error) {
	contract, err := bindSolveInbox(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SolveInboxFilterer{contract: contract}, nil
}

// bindSolveInbox binds a generic wrapper to an already deployed contract.
func bindSolveInbox(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SolveInboxMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SolveInbox *SolveInboxRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SolveInbox.Contract.SolveInboxCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SolveInbox *SolveInboxRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SolveInbox.Contract.SolveInboxTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SolveInbox *SolveInboxRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SolveInbox.Contract.SolveInboxTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SolveInbox *SolveInboxCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SolveInbox.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SolveInbox *SolveInboxTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SolveInbox.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SolveInbox *SolveInboxTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SolveInbox.Contract.contract.Transact(opts, method, params...)
}

// DefaultConfLevel is a free data retrieval call binding the contract method 0x74eeb847.
//
// Solidity: function defaultConfLevel() view returns(uint8)
func (_SolveInbox *SolveInboxCaller) DefaultConfLevel(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _SolveInbox.contract.Call(opts, &out, "defaultConfLevel")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// DefaultConfLevel is a free data retrieval call binding the contract method 0x74eeb847.
//
// Solidity: function defaultConfLevel() view returns(uint8)
func (_SolveInbox *SolveInboxSession) DefaultConfLevel() (uint8, error) {
	return _SolveInbox.Contract.DefaultConfLevel(&_SolveInbox.CallOpts)
}

// DefaultConfLevel is a free data retrieval call binding the contract method 0x74eeb847.
//
// Solidity: function defaultConfLevel() view returns(uint8)
func (_SolveInbox *SolveInboxCallerSession) DefaultConfLevel() (uint8, error) {
	return _SolveInbox.Contract.DefaultConfLevel(&_SolveInbox.CallOpts)
}

// DeployedAt is a free data retrieval call binding the contract method 0xeae4c19f.
//
// Solidity: function deployedAt() view returns(uint256)
func (_SolveInbox *SolveInboxCaller) DeployedAt(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _SolveInbox.contract.Call(opts, &out, "deployedAt")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DeployedAt is a free data retrieval call binding the contract method 0xeae4c19f.
//
// Solidity: function deployedAt() view returns(uint256)
func (_SolveInbox *SolveInboxSession) DeployedAt() (*big.Int, error) {
	return _SolveInbox.Contract.DeployedAt(&_SolveInbox.CallOpts)
}

// DeployedAt is a free data retrieval call binding the contract method 0xeae4c19f.
//
// Solidity: function deployedAt() view returns(uint256)
func (_SolveInbox *SolveInboxCallerSession) DeployedAt() (*big.Int, error) {
	return _SolveInbox.Contract.DeployedAt(&_SolveInbox.CallOpts)
}

// GetLatestRequestByStatus is a free data retrieval call binding the contract method 0x7667fd69.
//
// Solidity: function getLatestRequestByStatus(uint8 status) view returns((bytes32,uint8,address,address,(uint64,address,uint256,bytes),(bool,address,uint256)[],(uint8,uint40)[]))
func (_SolveInbox *SolveInboxCaller) GetLatestRequestByStatus(opts *bind.CallOpts, status uint8) (SolveRequest, error) {
	var out []interface{}
	err := _SolveInbox.contract.Call(opts, &out, "getLatestRequestByStatus", status)

	if err != nil {
		return *new(SolveRequest), err
	}

	out0 := *abi.ConvertType(out[0], new(SolveRequest)).(*SolveRequest)

	return out0, err

}

// GetLatestRequestByStatus is a free data retrieval call binding the contract method 0x7667fd69.
//
// Solidity: function getLatestRequestByStatus(uint8 status) view returns((bytes32,uint8,address,address,(uint64,address,uint256,bytes),(bool,address,uint256)[],(uint8,uint40)[]))
func (_SolveInbox *SolveInboxSession) GetLatestRequestByStatus(status uint8) (SolveRequest, error) {
	return _SolveInbox.Contract.GetLatestRequestByStatus(&_SolveInbox.CallOpts, status)
}

// GetLatestRequestByStatus is a free data retrieval call binding the contract method 0x7667fd69.
//
// Solidity: function getLatestRequestByStatus(uint8 status) view returns((bytes32,uint8,address,address,(uint64,address,uint256,bytes),(bool,address,uint256)[],(uint8,uint40)[]))
func (_SolveInbox *SolveInboxCallerSession) GetLatestRequestByStatus(status uint8) (SolveRequest, error) {
	return _SolveInbox.Contract.GetLatestRequestByStatus(&_SolveInbox.CallOpts, status)
}

// GetRequest is a free data retrieval call binding the contract method 0xfb1e61ca.
//
// Solidity: function getRequest(bytes32 id) view returns((bytes32,uint8,address,address,(uint64,address,uint256,bytes),(bool,address,uint256)[],(uint8,uint40)[]))
func (_SolveInbox *SolveInboxCaller) GetRequest(opts *bind.CallOpts, id [32]byte) (SolveRequest, error) {
	var out []interface{}
	err := _SolveInbox.contract.Call(opts, &out, "getRequest", id)

	if err != nil {
		return *new(SolveRequest), err
	}

	out0 := *abi.ConvertType(out[0], new(SolveRequest)).(*SolveRequest)

	return out0, err

}

// GetRequest is a free data retrieval call binding the contract method 0xfb1e61ca.
//
// Solidity: function getRequest(bytes32 id) view returns((bytes32,uint8,address,address,(uint64,address,uint256,bytes),(bool,address,uint256)[],(uint8,uint40)[]))
func (_SolveInbox *SolveInboxSession) GetRequest(id [32]byte) (SolveRequest, error) {
	return _SolveInbox.Contract.GetRequest(&_SolveInbox.CallOpts, id)
}

// GetRequest is a free data retrieval call binding the contract method 0xfb1e61ca.
//
// Solidity: function getRequest(bytes32 id) view returns((bytes32,uint8,address,address,(uint64,address,uint256,bytes),(bool,address,uint256)[],(uint8,uint40)[]))
func (_SolveInbox *SolveInboxCallerSession) GetRequest(id [32]byte) (SolveRequest, error) {
	return _SolveInbox.Contract.GetRequest(&_SolveInbox.CallOpts, id)
}

// GetUpdateHistory is a free data retrieval call binding the contract method 0x2cea64e8.
//
// Solidity: function getUpdateHistory(bytes32 id) view returns((uint8,uint40)[])
func (_SolveInbox *SolveInboxCaller) GetUpdateHistory(opts *bind.CallOpts, id [32]byte) ([]SolveStatusUpdate, error) {
	var out []interface{}
	err := _SolveInbox.contract.Call(opts, &out, "getUpdateHistory", id)

	if err != nil {
		return *new([]SolveStatusUpdate), err
	}

	out0 := *abi.ConvertType(out[0], new([]SolveStatusUpdate)).(*[]SolveStatusUpdate)

	return out0, err

}

// GetUpdateHistory is a free data retrieval call binding the contract method 0x2cea64e8.
//
// Solidity: function getUpdateHistory(bytes32 id) view returns((uint8,uint40)[])
func (_SolveInbox *SolveInboxSession) GetUpdateHistory(id [32]byte) ([]SolveStatusUpdate, error) {
	return _SolveInbox.Contract.GetUpdateHistory(&_SolveInbox.CallOpts, id)
}

// GetUpdateHistory is a free data retrieval call binding the contract method 0x2cea64e8.
//
// Solidity: function getUpdateHistory(bytes32 id) view returns((uint8,uint40)[])
func (_SolveInbox *SolveInboxCallerSession) GetUpdateHistory(id [32]byte) ([]SolveStatusUpdate, error) {
	return _SolveInbox.Contract.GetUpdateHistory(&_SolveInbox.CallOpts, id)
}

// HasAllRoles is a free data retrieval call binding the contract method 0x1cd64df4.
//
// Solidity: function hasAllRoles(address user, uint256 roles) view returns(bool)
func (_SolveInbox *SolveInboxCaller) HasAllRoles(opts *bind.CallOpts, user common.Address, roles *big.Int) (bool, error) {
	var out []interface{}
	err := _SolveInbox.contract.Call(opts, &out, "hasAllRoles", user, roles)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasAllRoles is a free data retrieval call binding the contract method 0x1cd64df4.
//
// Solidity: function hasAllRoles(address user, uint256 roles) view returns(bool)
func (_SolveInbox *SolveInboxSession) HasAllRoles(user common.Address, roles *big.Int) (bool, error) {
	return _SolveInbox.Contract.HasAllRoles(&_SolveInbox.CallOpts, user, roles)
}

// HasAllRoles is a free data retrieval call binding the contract method 0x1cd64df4.
//
// Solidity: function hasAllRoles(address user, uint256 roles) view returns(bool)
func (_SolveInbox *SolveInboxCallerSession) HasAllRoles(user common.Address, roles *big.Int) (bool, error) {
	return _SolveInbox.Contract.HasAllRoles(&_SolveInbox.CallOpts, user, roles)
}

// HasAnyRole is a free data retrieval call binding the contract method 0x514e62fc.
//
// Solidity: function hasAnyRole(address user, uint256 roles) view returns(bool)
func (_SolveInbox *SolveInboxCaller) HasAnyRole(opts *bind.CallOpts, user common.Address, roles *big.Int) (bool, error) {
	var out []interface{}
	err := _SolveInbox.contract.Call(opts, &out, "hasAnyRole", user, roles)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasAnyRole is a free data retrieval call binding the contract method 0x514e62fc.
//
// Solidity: function hasAnyRole(address user, uint256 roles) view returns(bool)
func (_SolveInbox *SolveInboxSession) HasAnyRole(user common.Address, roles *big.Int) (bool, error) {
	return _SolveInbox.Contract.HasAnyRole(&_SolveInbox.CallOpts, user, roles)
}

// HasAnyRole is a free data retrieval call binding the contract method 0x514e62fc.
//
// Solidity: function hasAnyRole(address user, uint256 roles) view returns(bool)
func (_SolveInbox *SolveInboxCallerSession) HasAnyRole(user common.Address, roles *big.Int) (bool, error) {
	return _SolveInbox.Contract.HasAnyRole(&_SolveInbox.CallOpts, user, roles)
}

// Omni is a free data retrieval call binding the contract method 0x39acf9f1.
//
// Solidity: function omni() view returns(address)
func (_SolveInbox *SolveInboxCaller) Omni(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SolveInbox.contract.Call(opts, &out, "omni")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Omni is a free data retrieval call binding the contract method 0x39acf9f1.
//
// Solidity: function omni() view returns(address)
func (_SolveInbox *SolveInboxSession) Omni() (common.Address, error) {
	return _SolveInbox.Contract.Omni(&_SolveInbox.CallOpts)
}

// Omni is a free data retrieval call binding the contract method 0x39acf9f1.
//
// Solidity: function omni() view returns(address)
func (_SolveInbox *SolveInboxCallerSession) Omni() (common.Address, error) {
	return _SolveInbox.Contract.Omni(&_SolveInbox.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address result)
func (_SolveInbox *SolveInboxCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SolveInbox.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address result)
func (_SolveInbox *SolveInboxSession) Owner() (common.Address, error) {
	return _SolveInbox.Contract.Owner(&_SolveInbox.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address result)
func (_SolveInbox *SolveInboxCallerSession) Owner() (common.Address, error) {
	return _SolveInbox.Contract.Owner(&_SolveInbox.CallOpts)
}

// OwnershipHandoverExpiresAt is a free data retrieval call binding the contract method 0xfee81cf4.
//
// Solidity: function ownershipHandoverExpiresAt(address pendingOwner) view returns(uint256 result)
func (_SolveInbox *SolveInboxCaller) OwnershipHandoverExpiresAt(opts *bind.CallOpts, pendingOwner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _SolveInbox.contract.Call(opts, &out, "ownershipHandoverExpiresAt", pendingOwner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OwnershipHandoverExpiresAt is a free data retrieval call binding the contract method 0xfee81cf4.
//
// Solidity: function ownershipHandoverExpiresAt(address pendingOwner) view returns(uint256 result)
func (_SolveInbox *SolveInboxSession) OwnershipHandoverExpiresAt(pendingOwner common.Address) (*big.Int, error) {
	return _SolveInbox.Contract.OwnershipHandoverExpiresAt(&_SolveInbox.CallOpts, pendingOwner)
}

// OwnershipHandoverExpiresAt is a free data retrieval call binding the contract method 0xfee81cf4.
//
// Solidity: function ownershipHandoverExpiresAt(address pendingOwner) view returns(uint256 result)
func (_SolveInbox *SolveInboxCallerSession) OwnershipHandoverExpiresAt(pendingOwner common.Address) (*big.Int, error) {
	return _SolveInbox.Contract.OwnershipHandoverExpiresAt(&_SolveInbox.CallOpts, pendingOwner)
}

// RolesOf is a free data retrieval call binding the contract method 0x2de94807.
//
// Solidity: function rolesOf(address user) view returns(uint256 roles)
func (_SolveInbox *SolveInboxCaller) RolesOf(opts *bind.CallOpts, user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _SolveInbox.contract.Call(opts, &out, "rolesOf", user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RolesOf is a free data retrieval call binding the contract method 0x2de94807.
//
// Solidity: function rolesOf(address user) view returns(uint256 roles)
func (_SolveInbox *SolveInboxSession) RolesOf(user common.Address) (*big.Int, error) {
	return _SolveInbox.Contract.RolesOf(&_SolveInbox.CallOpts, user)
}

// RolesOf is a free data retrieval call binding the contract method 0x2de94807.
//
// Solidity: function rolesOf(address user) view returns(uint256 roles)
func (_SolveInbox *SolveInboxCallerSession) RolesOf(user common.Address) (*big.Int, error) {
	return _SolveInbox.Contract.RolesOf(&_SolveInbox.CallOpts, user)
}

// Accept is a paid mutator transaction binding the contract method 0xe4725ba1.
//
// Solidity: function accept(bytes32 id) returns()
func (_SolveInbox *SolveInboxTransactor) Accept(opts *bind.TransactOpts, id [32]byte) (*types.Transaction, error) {
	return _SolveInbox.contract.Transact(opts, "accept", id)
}

// Accept is a paid mutator transaction binding the contract method 0xe4725ba1.
//
// Solidity: function accept(bytes32 id) returns()
func (_SolveInbox *SolveInboxSession) Accept(id [32]byte) (*types.Transaction, error) {
	return _SolveInbox.Contract.Accept(&_SolveInbox.TransactOpts, id)
}

// Accept is a paid mutator transaction binding the contract method 0xe4725ba1.
//
// Solidity: function accept(bytes32 id) returns()
func (_SolveInbox *SolveInboxTransactorSession) Accept(id [32]byte) (*types.Transaction, error) {
	return _SolveInbox.Contract.Accept(&_SolveInbox.TransactOpts, id)
}

// Cancel is a paid mutator transaction binding the contract method 0xc4d252f5.
//
// Solidity: function cancel(bytes32 id) returns()
func (_SolveInbox *SolveInboxTransactor) Cancel(opts *bind.TransactOpts, id [32]byte) (*types.Transaction, error) {
	return _SolveInbox.contract.Transact(opts, "cancel", id)
}

// Cancel is a paid mutator transaction binding the contract method 0xc4d252f5.
//
// Solidity: function cancel(bytes32 id) returns()
func (_SolveInbox *SolveInboxSession) Cancel(id [32]byte) (*types.Transaction, error) {
	return _SolveInbox.Contract.Cancel(&_SolveInbox.TransactOpts, id)
}

// Cancel is a paid mutator transaction binding the contract method 0xc4d252f5.
//
// Solidity: function cancel(bytes32 id) returns()
func (_SolveInbox *SolveInboxTransactorSession) Cancel(id [32]byte) (*types.Transaction, error) {
	return _SolveInbox.Contract.Cancel(&_SolveInbox.TransactOpts, id)
}

// CancelOwnershipHandover is a paid mutator transaction binding the contract method 0x54d1f13d.
//
// Solidity: function cancelOwnershipHandover() payable returns()
func (_SolveInbox *SolveInboxTransactor) CancelOwnershipHandover(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SolveInbox.contract.Transact(opts, "cancelOwnershipHandover")
}

// CancelOwnershipHandover is a paid mutator transaction binding the contract method 0x54d1f13d.
//
// Solidity: function cancelOwnershipHandover() payable returns()
func (_SolveInbox *SolveInboxSession) CancelOwnershipHandover() (*types.Transaction, error) {
	return _SolveInbox.Contract.CancelOwnershipHandover(&_SolveInbox.TransactOpts)
}

// CancelOwnershipHandover is a paid mutator transaction binding the contract method 0x54d1f13d.
//
// Solidity: function cancelOwnershipHandover() payable returns()
func (_SolveInbox *SolveInboxTransactorSession) CancelOwnershipHandover() (*types.Transaction, error) {
	return _SolveInbox.Contract.CancelOwnershipHandover(&_SolveInbox.TransactOpts)
}

// Claim is a paid mutator transaction binding the contract method 0x96c144f0.
//
// Solidity: function claim(bytes32 id, address to) returns()
func (_SolveInbox *SolveInboxTransactor) Claim(opts *bind.TransactOpts, id [32]byte, to common.Address) (*types.Transaction, error) {
	return _SolveInbox.contract.Transact(opts, "claim", id, to)
}

// Claim is a paid mutator transaction binding the contract method 0x96c144f0.
//
// Solidity: function claim(bytes32 id, address to) returns()
func (_SolveInbox *SolveInboxSession) Claim(id [32]byte, to common.Address) (*types.Transaction, error) {
	return _SolveInbox.Contract.Claim(&_SolveInbox.TransactOpts, id, to)
}

// Claim is a paid mutator transaction binding the contract method 0x96c144f0.
//
// Solidity: function claim(bytes32 id, address to) returns()
func (_SolveInbox *SolveInboxTransactorSession) Claim(id [32]byte, to common.Address) (*types.Transaction, error) {
	return _SolveInbox.Contract.Claim(&_SolveInbox.TransactOpts, id, to)
}

// CompleteOwnershipHandover is a paid mutator transaction binding the contract method 0xf04e283e.
//
// Solidity: function completeOwnershipHandover(address pendingOwner) payable returns()
func (_SolveInbox *SolveInboxTransactor) CompleteOwnershipHandover(opts *bind.TransactOpts, pendingOwner common.Address) (*types.Transaction, error) {
	return _SolveInbox.contract.Transact(opts, "completeOwnershipHandover", pendingOwner)
}

// CompleteOwnershipHandover is a paid mutator transaction binding the contract method 0xf04e283e.
//
// Solidity: function completeOwnershipHandover(address pendingOwner) payable returns()
func (_SolveInbox *SolveInboxSession) CompleteOwnershipHandover(pendingOwner common.Address) (*types.Transaction, error) {
	return _SolveInbox.Contract.CompleteOwnershipHandover(&_SolveInbox.TransactOpts, pendingOwner)
}

// CompleteOwnershipHandover is a paid mutator transaction binding the contract method 0xf04e283e.
//
// Solidity: function completeOwnershipHandover(address pendingOwner) payable returns()
func (_SolveInbox *SolveInboxTransactorSession) CompleteOwnershipHandover(pendingOwner common.Address) (*types.Transaction, error) {
	return _SolveInbox.Contract.CompleteOwnershipHandover(&_SolveInbox.TransactOpts, pendingOwner)
}

// GrantRoles is a paid mutator transaction binding the contract method 0x1c10893f.
//
// Solidity: function grantRoles(address user, uint256 roles) payable returns()
func (_SolveInbox *SolveInboxTransactor) GrantRoles(opts *bind.TransactOpts, user common.Address, roles *big.Int) (*types.Transaction, error) {
	return _SolveInbox.contract.Transact(opts, "grantRoles", user, roles)
}

// GrantRoles is a paid mutator transaction binding the contract method 0x1c10893f.
//
// Solidity: function grantRoles(address user, uint256 roles) payable returns()
func (_SolveInbox *SolveInboxSession) GrantRoles(user common.Address, roles *big.Int) (*types.Transaction, error) {
	return _SolveInbox.Contract.GrantRoles(&_SolveInbox.TransactOpts, user, roles)
}

// GrantRoles is a paid mutator transaction binding the contract method 0x1c10893f.
//
// Solidity: function grantRoles(address user, uint256 roles) payable returns()
func (_SolveInbox *SolveInboxTransactorSession) GrantRoles(user common.Address, roles *big.Int) (*types.Transaction, error) {
	return _SolveInbox.Contract.GrantRoles(&_SolveInbox.TransactOpts, user, roles)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address owner_, address solver_, address omni_, address outbox_) returns()
func (_SolveInbox *SolveInboxTransactor) Initialize(opts *bind.TransactOpts, owner_ common.Address, solver_ common.Address, omni_ common.Address, outbox_ common.Address) (*types.Transaction, error) {
	return _SolveInbox.contract.Transact(opts, "initialize", owner_, solver_, omni_, outbox_)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address owner_, address solver_, address omni_, address outbox_) returns()
func (_SolveInbox *SolveInboxSession) Initialize(owner_ common.Address, solver_ common.Address, omni_ common.Address, outbox_ common.Address) (*types.Transaction, error) {
	return _SolveInbox.Contract.Initialize(&_SolveInbox.TransactOpts, owner_, solver_, omni_, outbox_)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address owner_, address solver_, address omni_, address outbox_) returns()
func (_SolveInbox *SolveInboxTransactorSession) Initialize(owner_ common.Address, solver_ common.Address, omni_ common.Address, outbox_ common.Address) (*types.Transaction, error) {
	return _SolveInbox.Contract.Initialize(&_SolveInbox.TransactOpts, owner_, solver_, omni_, outbox_)
}

// MarkFulfilled is a paid mutator transaction binding the contract method 0x337ffe20.
//
// Solidity: function markFulfilled(bytes32 id, bytes32 callHash) returns()
func (_SolveInbox *SolveInboxTransactor) MarkFulfilled(opts *bind.TransactOpts, id [32]byte, callHash [32]byte) (*types.Transaction, error) {
	return _SolveInbox.contract.Transact(opts, "markFulfilled", id, callHash)
}

// MarkFulfilled is a paid mutator transaction binding the contract method 0x337ffe20.
//
// Solidity: function markFulfilled(bytes32 id, bytes32 callHash) returns()
func (_SolveInbox *SolveInboxSession) MarkFulfilled(id [32]byte, callHash [32]byte) (*types.Transaction, error) {
	return _SolveInbox.Contract.MarkFulfilled(&_SolveInbox.TransactOpts, id, callHash)
}

// MarkFulfilled is a paid mutator transaction binding the contract method 0x337ffe20.
//
// Solidity: function markFulfilled(bytes32 id, bytes32 callHash) returns()
func (_SolveInbox *SolveInboxTransactorSession) MarkFulfilled(id [32]byte, callHash [32]byte) (*types.Transaction, error) {
	return _SolveInbox.Contract.MarkFulfilled(&_SolveInbox.TransactOpts, id, callHash)
}

// Reject is a paid mutator transaction binding the contract method 0xdb3ea553.
//
// Solidity: function reject(bytes32 id, uint8 reason) returns()
func (_SolveInbox *SolveInboxTransactor) Reject(opts *bind.TransactOpts, id [32]byte, reason uint8) (*types.Transaction, error) {
	return _SolveInbox.contract.Transact(opts, "reject", id, reason)
}

// Reject is a paid mutator transaction binding the contract method 0xdb3ea553.
//
// Solidity: function reject(bytes32 id, uint8 reason) returns()
func (_SolveInbox *SolveInboxSession) Reject(id [32]byte, reason uint8) (*types.Transaction, error) {
	return _SolveInbox.Contract.Reject(&_SolveInbox.TransactOpts, id, reason)
}

// Reject is a paid mutator transaction binding the contract method 0xdb3ea553.
//
// Solidity: function reject(bytes32 id, uint8 reason) returns()
func (_SolveInbox *SolveInboxTransactorSession) Reject(id [32]byte, reason uint8) (*types.Transaction, error) {
	return _SolveInbox.Contract.Reject(&_SolveInbox.TransactOpts, id, reason)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() payable returns()
func (_SolveInbox *SolveInboxTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SolveInbox.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() payable returns()
func (_SolveInbox *SolveInboxSession) RenounceOwnership() (*types.Transaction, error) {
	return _SolveInbox.Contract.RenounceOwnership(&_SolveInbox.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() payable returns()
func (_SolveInbox *SolveInboxTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _SolveInbox.Contract.RenounceOwnership(&_SolveInbox.TransactOpts)
}

// RenounceRoles is a paid mutator transaction binding the contract method 0x183a4f6e.
//
// Solidity: function renounceRoles(uint256 roles) payable returns()
func (_SolveInbox *SolveInboxTransactor) RenounceRoles(opts *bind.TransactOpts, roles *big.Int) (*types.Transaction, error) {
	return _SolveInbox.contract.Transact(opts, "renounceRoles", roles)
}

// RenounceRoles is a paid mutator transaction binding the contract method 0x183a4f6e.
//
// Solidity: function renounceRoles(uint256 roles) payable returns()
func (_SolveInbox *SolveInboxSession) RenounceRoles(roles *big.Int) (*types.Transaction, error) {
	return _SolveInbox.Contract.RenounceRoles(&_SolveInbox.TransactOpts, roles)
}

// RenounceRoles is a paid mutator transaction binding the contract method 0x183a4f6e.
//
// Solidity: function renounceRoles(uint256 roles) payable returns()
func (_SolveInbox *SolveInboxTransactorSession) RenounceRoles(roles *big.Int) (*types.Transaction, error) {
	return _SolveInbox.Contract.RenounceRoles(&_SolveInbox.TransactOpts, roles)
}

// Request is a paid mutator transaction binding the contract method 0xe17771c8.
//
// Solidity: function request((uint64,address,uint256,bytes) call, (address,uint256)[] deposits) payable returns(bytes32 id)
func (_SolveInbox *SolveInboxTransactor) Request(opts *bind.TransactOpts, call SolveCall, deposits []SolveTokenDeposit) (*types.Transaction, error) {
	return _SolveInbox.contract.Transact(opts, "request", call, deposits)
}

// Request is a paid mutator transaction binding the contract method 0xe17771c8.
//
// Solidity: function request((uint64,address,uint256,bytes) call, (address,uint256)[] deposits) payable returns(bytes32 id)
func (_SolveInbox *SolveInboxSession) Request(call SolveCall, deposits []SolveTokenDeposit) (*types.Transaction, error) {
	return _SolveInbox.Contract.Request(&_SolveInbox.TransactOpts, call, deposits)
}

// Request is a paid mutator transaction binding the contract method 0xe17771c8.
//
// Solidity: function request((uint64,address,uint256,bytes) call, (address,uint256)[] deposits) payable returns(bytes32 id)
func (_SolveInbox *SolveInboxTransactorSession) Request(call SolveCall, deposits []SolveTokenDeposit) (*types.Transaction, error) {
	return _SolveInbox.Contract.Request(&_SolveInbox.TransactOpts, call, deposits)
}

// RequestOwnershipHandover is a paid mutator transaction binding the contract method 0x25692962.
//
// Solidity: function requestOwnershipHandover() payable returns()
func (_SolveInbox *SolveInboxTransactor) RequestOwnershipHandover(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SolveInbox.contract.Transact(opts, "requestOwnershipHandover")
}

// RequestOwnershipHandover is a paid mutator transaction binding the contract method 0x25692962.
//
// Solidity: function requestOwnershipHandover() payable returns()
func (_SolveInbox *SolveInboxSession) RequestOwnershipHandover() (*types.Transaction, error) {
	return _SolveInbox.Contract.RequestOwnershipHandover(&_SolveInbox.TransactOpts)
}

// RequestOwnershipHandover is a paid mutator transaction binding the contract method 0x25692962.
//
// Solidity: function requestOwnershipHandover() payable returns()
func (_SolveInbox *SolveInboxTransactorSession) RequestOwnershipHandover() (*types.Transaction, error) {
	return _SolveInbox.Contract.RequestOwnershipHandover(&_SolveInbox.TransactOpts)
}

// RevokeRoles is a paid mutator transaction binding the contract method 0x4a4ee7b1.
//
// Solidity: function revokeRoles(address user, uint256 roles) payable returns()
func (_SolveInbox *SolveInboxTransactor) RevokeRoles(opts *bind.TransactOpts, user common.Address, roles *big.Int) (*types.Transaction, error) {
	return _SolveInbox.contract.Transact(opts, "revokeRoles", user, roles)
}

// RevokeRoles is a paid mutator transaction binding the contract method 0x4a4ee7b1.
//
// Solidity: function revokeRoles(address user, uint256 roles) payable returns()
func (_SolveInbox *SolveInboxSession) RevokeRoles(user common.Address, roles *big.Int) (*types.Transaction, error) {
	return _SolveInbox.Contract.RevokeRoles(&_SolveInbox.TransactOpts, user, roles)
}

// RevokeRoles is a paid mutator transaction binding the contract method 0x4a4ee7b1.
//
// Solidity: function revokeRoles(address user, uint256 roles) payable returns()
func (_SolveInbox *SolveInboxTransactorSession) RevokeRoles(user common.Address, roles *big.Int) (*types.Transaction, error) {
	return _SolveInbox.Contract.RevokeRoles(&_SolveInbox.TransactOpts, user, roles)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) payable returns()
func (_SolveInbox *SolveInboxTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _SolveInbox.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) payable returns()
func (_SolveInbox *SolveInboxSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _SolveInbox.Contract.TransferOwnership(&_SolveInbox.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) payable returns()
func (_SolveInbox *SolveInboxTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _SolveInbox.Contract.TransferOwnership(&_SolveInbox.TransactOpts, newOwner)
}

// SolveInboxAcceptedIterator is returned from FilterAccepted and is used to iterate over the raw logs and unpacked data for Accepted events raised by the SolveInbox contract.
type SolveInboxAcceptedIterator struct {
	Event *SolveInboxAccepted // Event containing the contract specifics and raw log

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
func (it *SolveInboxAcceptedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SolveInboxAccepted)
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
		it.Event = new(SolveInboxAccepted)
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
func (it *SolveInboxAcceptedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SolveInboxAcceptedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SolveInboxAccepted represents a Accepted event raised by the SolveInbox contract.
type SolveInboxAccepted struct {
	Id  [32]byte
	By  common.Address
	Raw types.Log // Blockchain specific contextual infos
}

// FilterAccepted is a free log retrieval operation binding the contract event 0x9deed34441ca75bb2dbbe101d2201930f40e18a9ce521c77fbdca6690a899967.
//
// Solidity: event Accepted(bytes32 indexed id, address indexed by)
func (_SolveInbox *SolveInboxFilterer) FilterAccepted(opts *bind.FilterOpts, id [][32]byte, by []common.Address) (*SolveInboxAcceptedIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var byRule []interface{}
	for _, byItem := range by {
		byRule = append(byRule, byItem)
	}

	logs, sub, err := _SolveInbox.contract.FilterLogs(opts, "Accepted", idRule, byRule)
	if err != nil {
		return nil, err
	}
	return &SolveInboxAcceptedIterator{contract: _SolveInbox.contract, event: "Accepted", logs: logs, sub: sub}, nil
}

// WatchAccepted is a free log subscription operation binding the contract event 0x9deed34441ca75bb2dbbe101d2201930f40e18a9ce521c77fbdca6690a899967.
//
// Solidity: event Accepted(bytes32 indexed id, address indexed by)
func (_SolveInbox *SolveInboxFilterer) WatchAccepted(opts *bind.WatchOpts, sink chan<- *SolveInboxAccepted, id [][32]byte, by []common.Address) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var byRule []interface{}
	for _, byItem := range by {
		byRule = append(byRule, byItem)
	}

	logs, sub, err := _SolveInbox.contract.WatchLogs(opts, "Accepted", idRule, byRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SolveInboxAccepted)
				if err := _SolveInbox.contract.UnpackLog(event, "Accepted", log); err != nil {
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

// ParseAccepted is a log parse operation binding the contract event 0x9deed34441ca75bb2dbbe101d2201930f40e18a9ce521c77fbdca6690a899967.
//
// Solidity: event Accepted(bytes32 indexed id, address indexed by)
func (_SolveInbox *SolveInboxFilterer) ParseAccepted(log types.Log) (*SolveInboxAccepted, error) {
	event := new(SolveInboxAccepted)
	if err := _SolveInbox.contract.UnpackLog(event, "Accepted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SolveInboxClaimedIterator is returned from FilterClaimed and is used to iterate over the raw logs and unpacked data for Claimed events raised by the SolveInbox contract.
type SolveInboxClaimedIterator struct {
	Event *SolveInboxClaimed // Event containing the contract specifics and raw log

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
func (it *SolveInboxClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SolveInboxClaimed)
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
		it.Event = new(SolveInboxClaimed)
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
func (it *SolveInboxClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SolveInboxClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SolveInboxClaimed represents a Claimed event raised by the SolveInbox contract.
type SolveInboxClaimed struct {
	Id       [32]byte
	By       common.Address
	To       common.Address
	Deposits []SolveDeposit
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterClaimed is a free log retrieval operation binding the contract event 0x1291cec4ea55732527f1ae8ff3e53def0dbe1feb99c094983449e59cdce4674b.
//
// Solidity: event Claimed(bytes32 indexed id, address indexed by, address indexed to, (bool,address,uint256)[] deposits)
func (_SolveInbox *SolveInboxFilterer) FilterClaimed(opts *bind.FilterOpts, id [][32]byte, by []common.Address, to []common.Address) (*SolveInboxClaimedIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var byRule []interface{}
	for _, byItem := range by {
		byRule = append(byRule, byItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _SolveInbox.contract.FilterLogs(opts, "Claimed", idRule, byRule, toRule)
	if err != nil {
		return nil, err
	}
	return &SolveInboxClaimedIterator{contract: _SolveInbox.contract, event: "Claimed", logs: logs, sub: sub}, nil
}

// WatchClaimed is a free log subscription operation binding the contract event 0x1291cec4ea55732527f1ae8ff3e53def0dbe1feb99c094983449e59cdce4674b.
//
// Solidity: event Claimed(bytes32 indexed id, address indexed by, address indexed to, (bool,address,uint256)[] deposits)
func (_SolveInbox *SolveInboxFilterer) WatchClaimed(opts *bind.WatchOpts, sink chan<- *SolveInboxClaimed, id [][32]byte, by []common.Address, to []common.Address) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var byRule []interface{}
	for _, byItem := range by {
		byRule = append(byRule, byItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _SolveInbox.contract.WatchLogs(opts, "Claimed", idRule, byRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SolveInboxClaimed)
				if err := _SolveInbox.contract.UnpackLog(event, "Claimed", log); err != nil {
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

// ParseClaimed is a log parse operation binding the contract event 0x1291cec4ea55732527f1ae8ff3e53def0dbe1feb99c094983449e59cdce4674b.
//
// Solidity: event Claimed(bytes32 indexed id, address indexed by, address indexed to, (bool,address,uint256)[] deposits)
func (_SolveInbox *SolveInboxFilterer) ParseClaimed(log types.Log) (*SolveInboxClaimed, error) {
	event := new(SolveInboxClaimed)
	if err := _SolveInbox.contract.UnpackLog(event, "Claimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SolveInboxDefaultConfLevelSetIterator is returned from FilterDefaultConfLevelSet and is used to iterate over the raw logs and unpacked data for DefaultConfLevelSet events raised by the SolveInbox contract.
type SolveInboxDefaultConfLevelSetIterator struct {
	Event *SolveInboxDefaultConfLevelSet // Event containing the contract specifics and raw log

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
func (it *SolveInboxDefaultConfLevelSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SolveInboxDefaultConfLevelSet)
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
		it.Event = new(SolveInboxDefaultConfLevelSet)
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
func (it *SolveInboxDefaultConfLevelSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SolveInboxDefaultConfLevelSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SolveInboxDefaultConfLevelSet represents a DefaultConfLevelSet event raised by the SolveInbox contract.
type SolveInboxDefaultConfLevelSet struct {
	Conf uint8
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterDefaultConfLevelSet is a free log retrieval operation binding the contract event 0x8de08a798b4e50b4f351c1eaa91a11530043802be3ffac2df87db0c45a2e8483.
//
// Solidity: event DefaultConfLevelSet(uint8 conf)
func (_SolveInbox *SolveInboxFilterer) FilterDefaultConfLevelSet(opts *bind.FilterOpts) (*SolveInboxDefaultConfLevelSetIterator, error) {

	logs, sub, err := _SolveInbox.contract.FilterLogs(opts, "DefaultConfLevelSet")
	if err != nil {
		return nil, err
	}
	return &SolveInboxDefaultConfLevelSetIterator{contract: _SolveInbox.contract, event: "DefaultConfLevelSet", logs: logs, sub: sub}, nil
}

// WatchDefaultConfLevelSet is a free log subscription operation binding the contract event 0x8de08a798b4e50b4f351c1eaa91a11530043802be3ffac2df87db0c45a2e8483.
//
// Solidity: event DefaultConfLevelSet(uint8 conf)
func (_SolveInbox *SolveInboxFilterer) WatchDefaultConfLevelSet(opts *bind.WatchOpts, sink chan<- *SolveInboxDefaultConfLevelSet) (event.Subscription, error) {

	logs, sub, err := _SolveInbox.contract.WatchLogs(opts, "DefaultConfLevelSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SolveInboxDefaultConfLevelSet)
				if err := _SolveInbox.contract.UnpackLog(event, "DefaultConfLevelSet", log); err != nil {
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

// ParseDefaultConfLevelSet is a log parse operation binding the contract event 0x8de08a798b4e50b4f351c1eaa91a11530043802be3ffac2df87db0c45a2e8483.
//
// Solidity: event DefaultConfLevelSet(uint8 conf)
func (_SolveInbox *SolveInboxFilterer) ParseDefaultConfLevelSet(log types.Log) (*SolveInboxDefaultConfLevelSet, error) {
	event := new(SolveInboxDefaultConfLevelSet)
	if err := _SolveInbox.contract.UnpackLog(event, "DefaultConfLevelSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SolveInboxFulfilledIterator is returned from FilterFulfilled and is used to iterate over the raw logs and unpacked data for Fulfilled events raised by the SolveInbox contract.
type SolveInboxFulfilledIterator struct {
	Event *SolveInboxFulfilled // Event containing the contract specifics and raw log

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
func (it *SolveInboxFulfilledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SolveInboxFulfilled)
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
		it.Event = new(SolveInboxFulfilled)
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
func (it *SolveInboxFulfilledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SolveInboxFulfilledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SolveInboxFulfilled represents a Fulfilled event raised by the SolveInbox contract.
type SolveInboxFulfilled struct {
	Id         [32]byte
	CallHash   [32]byte
	CreditedTo common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterFulfilled is a free log retrieval operation binding the contract event 0x7898a125e0970666c80e00bbf2e7041d84dfe5bbe6bcf562ce53d540fd6cd891.
//
// Solidity: event Fulfilled(bytes32 indexed id, bytes32 indexed callHash, address indexed creditedTo)
func (_SolveInbox *SolveInboxFilterer) FilterFulfilled(opts *bind.FilterOpts, id [][32]byte, callHash [][32]byte, creditedTo []common.Address) (*SolveInboxFulfilledIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var callHashRule []interface{}
	for _, callHashItem := range callHash {
		callHashRule = append(callHashRule, callHashItem)
	}
	var creditedToRule []interface{}
	for _, creditedToItem := range creditedTo {
		creditedToRule = append(creditedToRule, creditedToItem)
	}

	logs, sub, err := _SolveInbox.contract.FilterLogs(opts, "Fulfilled", idRule, callHashRule, creditedToRule)
	if err != nil {
		return nil, err
	}
	return &SolveInboxFulfilledIterator{contract: _SolveInbox.contract, event: "Fulfilled", logs: logs, sub: sub}, nil
}

// WatchFulfilled is a free log subscription operation binding the contract event 0x7898a125e0970666c80e00bbf2e7041d84dfe5bbe6bcf562ce53d540fd6cd891.
//
// Solidity: event Fulfilled(bytes32 indexed id, bytes32 indexed callHash, address indexed creditedTo)
func (_SolveInbox *SolveInboxFilterer) WatchFulfilled(opts *bind.WatchOpts, sink chan<- *SolveInboxFulfilled, id [][32]byte, callHash [][32]byte, creditedTo []common.Address) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var callHashRule []interface{}
	for _, callHashItem := range callHash {
		callHashRule = append(callHashRule, callHashItem)
	}
	var creditedToRule []interface{}
	for _, creditedToItem := range creditedTo {
		creditedToRule = append(creditedToRule, creditedToItem)
	}

	logs, sub, err := _SolveInbox.contract.WatchLogs(opts, "Fulfilled", idRule, callHashRule, creditedToRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SolveInboxFulfilled)
				if err := _SolveInbox.contract.UnpackLog(event, "Fulfilled", log); err != nil {
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

// ParseFulfilled is a log parse operation binding the contract event 0x7898a125e0970666c80e00bbf2e7041d84dfe5bbe6bcf562ce53d540fd6cd891.
//
// Solidity: event Fulfilled(bytes32 indexed id, bytes32 indexed callHash, address indexed creditedTo)
func (_SolveInbox *SolveInboxFilterer) ParseFulfilled(log types.Log) (*SolveInboxFulfilled, error) {
	event := new(SolveInboxFulfilled)
	if err := _SolveInbox.contract.UnpackLog(event, "Fulfilled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SolveInboxInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the SolveInbox contract.
type SolveInboxInitializedIterator struct {
	Event *SolveInboxInitialized // Event containing the contract specifics and raw log

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
func (it *SolveInboxInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SolveInboxInitialized)
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
		it.Event = new(SolveInboxInitialized)
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
func (it *SolveInboxInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SolveInboxInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SolveInboxInitialized represents a Initialized event raised by the SolveInbox contract.
type SolveInboxInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_SolveInbox *SolveInboxFilterer) FilterInitialized(opts *bind.FilterOpts) (*SolveInboxInitializedIterator, error) {

	logs, sub, err := _SolveInbox.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &SolveInboxInitializedIterator{contract: _SolveInbox.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_SolveInbox *SolveInboxFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *SolveInboxInitialized) (event.Subscription, error) {

	logs, sub, err := _SolveInbox.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SolveInboxInitialized)
				if err := _SolveInbox.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_SolveInbox *SolveInboxFilterer) ParseInitialized(log types.Log) (*SolveInboxInitialized, error) {
	event := new(SolveInboxInitialized)
	if err := _SolveInbox.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SolveInboxOmniPortalSetIterator is returned from FilterOmniPortalSet and is used to iterate over the raw logs and unpacked data for OmniPortalSet events raised by the SolveInbox contract.
type SolveInboxOmniPortalSetIterator struct {
	Event *SolveInboxOmniPortalSet // Event containing the contract specifics and raw log

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
func (it *SolveInboxOmniPortalSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SolveInboxOmniPortalSet)
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
		it.Event = new(SolveInboxOmniPortalSet)
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
func (it *SolveInboxOmniPortalSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SolveInboxOmniPortalSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SolveInboxOmniPortalSet represents a OmniPortalSet event raised by the SolveInbox contract.
type SolveInboxOmniPortalSet struct {
	Omni common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterOmniPortalSet is a free log retrieval operation binding the contract event 0x79162c8d053a07e70cdc1ccc536f0440b571f8508377d2bef51094fadab98f47.
//
// Solidity: event OmniPortalSet(address omni)
func (_SolveInbox *SolveInboxFilterer) FilterOmniPortalSet(opts *bind.FilterOpts) (*SolveInboxOmniPortalSetIterator, error) {

	logs, sub, err := _SolveInbox.contract.FilterLogs(opts, "OmniPortalSet")
	if err != nil {
		return nil, err
	}
	return &SolveInboxOmniPortalSetIterator{contract: _SolveInbox.contract, event: "OmniPortalSet", logs: logs, sub: sub}, nil
}

// WatchOmniPortalSet is a free log subscription operation binding the contract event 0x79162c8d053a07e70cdc1ccc536f0440b571f8508377d2bef51094fadab98f47.
//
// Solidity: event OmniPortalSet(address omni)
func (_SolveInbox *SolveInboxFilterer) WatchOmniPortalSet(opts *bind.WatchOpts, sink chan<- *SolveInboxOmniPortalSet) (event.Subscription, error) {

	logs, sub, err := _SolveInbox.contract.WatchLogs(opts, "OmniPortalSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SolveInboxOmniPortalSet)
				if err := _SolveInbox.contract.UnpackLog(event, "OmniPortalSet", log); err != nil {
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

// ParseOmniPortalSet is a log parse operation binding the contract event 0x79162c8d053a07e70cdc1ccc536f0440b571f8508377d2bef51094fadab98f47.
//
// Solidity: event OmniPortalSet(address omni)
func (_SolveInbox *SolveInboxFilterer) ParseOmniPortalSet(log types.Log) (*SolveInboxOmniPortalSet, error) {
	event := new(SolveInboxOmniPortalSet)
	if err := _SolveInbox.contract.UnpackLog(event, "OmniPortalSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SolveInboxOwnershipHandoverCanceledIterator is returned from FilterOwnershipHandoverCanceled and is used to iterate over the raw logs and unpacked data for OwnershipHandoverCanceled events raised by the SolveInbox contract.
type SolveInboxOwnershipHandoverCanceledIterator struct {
	Event *SolveInboxOwnershipHandoverCanceled // Event containing the contract specifics and raw log

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
func (it *SolveInboxOwnershipHandoverCanceledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SolveInboxOwnershipHandoverCanceled)
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
		it.Event = new(SolveInboxOwnershipHandoverCanceled)
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
func (it *SolveInboxOwnershipHandoverCanceledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SolveInboxOwnershipHandoverCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SolveInboxOwnershipHandoverCanceled represents a OwnershipHandoverCanceled event raised by the SolveInbox contract.
type SolveInboxOwnershipHandoverCanceled struct {
	PendingOwner common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterOwnershipHandoverCanceled is a free log retrieval operation binding the contract event 0xfa7b8eab7da67f412cc9575ed43464468f9bfbae89d1675917346ca6d8fe3c92.
//
// Solidity: event OwnershipHandoverCanceled(address indexed pendingOwner)
func (_SolveInbox *SolveInboxFilterer) FilterOwnershipHandoverCanceled(opts *bind.FilterOpts, pendingOwner []common.Address) (*SolveInboxOwnershipHandoverCanceledIterator, error) {

	var pendingOwnerRule []interface{}
	for _, pendingOwnerItem := range pendingOwner {
		pendingOwnerRule = append(pendingOwnerRule, pendingOwnerItem)
	}

	logs, sub, err := _SolveInbox.contract.FilterLogs(opts, "OwnershipHandoverCanceled", pendingOwnerRule)
	if err != nil {
		return nil, err
	}
	return &SolveInboxOwnershipHandoverCanceledIterator{contract: _SolveInbox.contract, event: "OwnershipHandoverCanceled", logs: logs, sub: sub}, nil
}

// WatchOwnershipHandoverCanceled is a free log subscription operation binding the contract event 0xfa7b8eab7da67f412cc9575ed43464468f9bfbae89d1675917346ca6d8fe3c92.
//
// Solidity: event OwnershipHandoverCanceled(address indexed pendingOwner)
func (_SolveInbox *SolveInboxFilterer) WatchOwnershipHandoverCanceled(opts *bind.WatchOpts, sink chan<- *SolveInboxOwnershipHandoverCanceled, pendingOwner []common.Address) (event.Subscription, error) {

	var pendingOwnerRule []interface{}
	for _, pendingOwnerItem := range pendingOwner {
		pendingOwnerRule = append(pendingOwnerRule, pendingOwnerItem)
	}

	logs, sub, err := _SolveInbox.contract.WatchLogs(opts, "OwnershipHandoverCanceled", pendingOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SolveInboxOwnershipHandoverCanceled)
				if err := _SolveInbox.contract.UnpackLog(event, "OwnershipHandoverCanceled", log); err != nil {
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

// ParseOwnershipHandoverCanceled is a log parse operation binding the contract event 0xfa7b8eab7da67f412cc9575ed43464468f9bfbae89d1675917346ca6d8fe3c92.
//
// Solidity: event OwnershipHandoverCanceled(address indexed pendingOwner)
func (_SolveInbox *SolveInboxFilterer) ParseOwnershipHandoverCanceled(log types.Log) (*SolveInboxOwnershipHandoverCanceled, error) {
	event := new(SolveInboxOwnershipHandoverCanceled)
	if err := _SolveInbox.contract.UnpackLog(event, "OwnershipHandoverCanceled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SolveInboxOwnershipHandoverRequestedIterator is returned from FilterOwnershipHandoverRequested and is used to iterate over the raw logs and unpacked data for OwnershipHandoverRequested events raised by the SolveInbox contract.
type SolveInboxOwnershipHandoverRequestedIterator struct {
	Event *SolveInboxOwnershipHandoverRequested // Event containing the contract specifics and raw log

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
func (it *SolveInboxOwnershipHandoverRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SolveInboxOwnershipHandoverRequested)
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
		it.Event = new(SolveInboxOwnershipHandoverRequested)
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
func (it *SolveInboxOwnershipHandoverRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SolveInboxOwnershipHandoverRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SolveInboxOwnershipHandoverRequested represents a OwnershipHandoverRequested event raised by the SolveInbox contract.
type SolveInboxOwnershipHandoverRequested struct {
	PendingOwner common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterOwnershipHandoverRequested is a free log retrieval operation binding the contract event 0xdbf36a107da19e49527a7176a1babf963b4b0ff8cde35ee35d6cd8f1f9ac7e1d.
//
// Solidity: event OwnershipHandoverRequested(address indexed pendingOwner)
func (_SolveInbox *SolveInboxFilterer) FilterOwnershipHandoverRequested(opts *bind.FilterOpts, pendingOwner []common.Address) (*SolveInboxOwnershipHandoverRequestedIterator, error) {

	var pendingOwnerRule []interface{}
	for _, pendingOwnerItem := range pendingOwner {
		pendingOwnerRule = append(pendingOwnerRule, pendingOwnerItem)
	}

	logs, sub, err := _SolveInbox.contract.FilterLogs(opts, "OwnershipHandoverRequested", pendingOwnerRule)
	if err != nil {
		return nil, err
	}
	return &SolveInboxOwnershipHandoverRequestedIterator{contract: _SolveInbox.contract, event: "OwnershipHandoverRequested", logs: logs, sub: sub}, nil
}

// WatchOwnershipHandoverRequested is a free log subscription operation binding the contract event 0xdbf36a107da19e49527a7176a1babf963b4b0ff8cde35ee35d6cd8f1f9ac7e1d.
//
// Solidity: event OwnershipHandoverRequested(address indexed pendingOwner)
func (_SolveInbox *SolveInboxFilterer) WatchOwnershipHandoverRequested(opts *bind.WatchOpts, sink chan<- *SolveInboxOwnershipHandoverRequested, pendingOwner []common.Address) (event.Subscription, error) {

	var pendingOwnerRule []interface{}
	for _, pendingOwnerItem := range pendingOwner {
		pendingOwnerRule = append(pendingOwnerRule, pendingOwnerItem)
	}

	logs, sub, err := _SolveInbox.contract.WatchLogs(opts, "OwnershipHandoverRequested", pendingOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SolveInboxOwnershipHandoverRequested)
				if err := _SolveInbox.contract.UnpackLog(event, "OwnershipHandoverRequested", log); err != nil {
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

// ParseOwnershipHandoverRequested is a log parse operation binding the contract event 0xdbf36a107da19e49527a7176a1babf963b4b0ff8cde35ee35d6cd8f1f9ac7e1d.
//
// Solidity: event OwnershipHandoverRequested(address indexed pendingOwner)
func (_SolveInbox *SolveInboxFilterer) ParseOwnershipHandoverRequested(log types.Log) (*SolveInboxOwnershipHandoverRequested, error) {
	event := new(SolveInboxOwnershipHandoverRequested)
	if err := _SolveInbox.contract.UnpackLog(event, "OwnershipHandoverRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SolveInboxOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the SolveInbox contract.
type SolveInboxOwnershipTransferredIterator struct {
	Event *SolveInboxOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *SolveInboxOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SolveInboxOwnershipTransferred)
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
		it.Event = new(SolveInboxOwnershipTransferred)
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
func (it *SolveInboxOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SolveInboxOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SolveInboxOwnershipTransferred represents a OwnershipTransferred event raised by the SolveInbox contract.
type SolveInboxOwnershipTransferred struct {
	OldOwner common.Address
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed oldOwner, address indexed newOwner)
func (_SolveInbox *SolveInboxFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, oldOwner []common.Address, newOwner []common.Address) (*SolveInboxOwnershipTransferredIterator, error) {

	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _SolveInbox.contract.FilterLogs(opts, "OwnershipTransferred", oldOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &SolveInboxOwnershipTransferredIterator{contract: _SolveInbox.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed oldOwner, address indexed newOwner)
func (_SolveInbox *SolveInboxFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *SolveInboxOwnershipTransferred, oldOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _SolveInbox.contract.WatchLogs(opts, "OwnershipTransferred", oldOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SolveInboxOwnershipTransferred)
				if err := _SolveInbox.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
// Solidity: event OwnershipTransferred(address indexed oldOwner, address indexed newOwner)
func (_SolveInbox *SolveInboxFilterer) ParseOwnershipTransferred(log types.Log) (*SolveInboxOwnershipTransferred, error) {
	event := new(SolveInboxOwnershipTransferred)
	if err := _SolveInbox.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SolveInboxRejectedIterator is returned from FilterRejected and is used to iterate over the raw logs and unpacked data for Rejected events raised by the SolveInbox contract.
type SolveInboxRejectedIterator struct {
	Event *SolveInboxRejected // Event containing the contract specifics and raw log

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
func (it *SolveInboxRejectedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SolveInboxRejected)
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
		it.Event = new(SolveInboxRejected)
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
func (it *SolveInboxRejectedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SolveInboxRejectedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SolveInboxRejected represents a Rejected event raised by the SolveInbox contract.
type SolveInboxRejected struct {
	Id     [32]byte
	By     common.Address
	Reason uint8
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRejected is a free log retrieval operation binding the contract event 0x21f84ee3a6e9bc7c10f855f8c9829e22c613861cef10add09eccdbc88df9f59f.
//
// Solidity: event Rejected(bytes32 indexed id, address indexed by, uint8 indexed reason)
func (_SolveInbox *SolveInboxFilterer) FilterRejected(opts *bind.FilterOpts, id [][32]byte, by []common.Address, reason []uint8) (*SolveInboxRejectedIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var byRule []interface{}
	for _, byItem := range by {
		byRule = append(byRule, byItem)
	}
	var reasonRule []interface{}
	for _, reasonItem := range reason {
		reasonRule = append(reasonRule, reasonItem)
	}

	logs, sub, err := _SolveInbox.contract.FilterLogs(opts, "Rejected", idRule, byRule, reasonRule)
	if err != nil {
		return nil, err
	}
	return &SolveInboxRejectedIterator{contract: _SolveInbox.contract, event: "Rejected", logs: logs, sub: sub}, nil
}

// WatchRejected is a free log subscription operation binding the contract event 0x21f84ee3a6e9bc7c10f855f8c9829e22c613861cef10add09eccdbc88df9f59f.
//
// Solidity: event Rejected(bytes32 indexed id, address indexed by, uint8 indexed reason)
func (_SolveInbox *SolveInboxFilterer) WatchRejected(opts *bind.WatchOpts, sink chan<- *SolveInboxRejected, id [][32]byte, by []common.Address, reason []uint8) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var byRule []interface{}
	for _, byItem := range by {
		byRule = append(byRule, byItem)
	}
	var reasonRule []interface{}
	for _, reasonItem := range reason {
		reasonRule = append(reasonRule, reasonItem)
	}

	logs, sub, err := _SolveInbox.contract.WatchLogs(opts, "Rejected", idRule, byRule, reasonRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SolveInboxRejected)
				if err := _SolveInbox.contract.UnpackLog(event, "Rejected", log); err != nil {
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

// ParseRejected is a log parse operation binding the contract event 0x21f84ee3a6e9bc7c10f855f8c9829e22c613861cef10add09eccdbc88df9f59f.
//
// Solidity: event Rejected(bytes32 indexed id, address indexed by, uint8 indexed reason)
func (_SolveInbox *SolveInboxFilterer) ParseRejected(log types.Log) (*SolveInboxRejected, error) {
	event := new(SolveInboxRejected)
	if err := _SolveInbox.contract.UnpackLog(event, "Rejected", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SolveInboxRequestedIterator is returned from FilterRequested and is used to iterate over the raw logs and unpacked data for Requested events raised by the SolveInbox contract.
type SolveInboxRequestedIterator struct {
	Event *SolveInboxRequested // Event containing the contract specifics and raw log

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
func (it *SolveInboxRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SolveInboxRequested)
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
		it.Event = new(SolveInboxRequested)
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
func (it *SolveInboxRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SolveInboxRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SolveInboxRequested represents a Requested event raised by the SolveInbox contract.
type SolveInboxRequested struct {
	Id       [32]byte
	From     common.Address
	Call     SolveCall
	Deposits []SolveDeposit
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterRequested is a free log retrieval operation binding the contract event 0x3622a563ce1f96b477fa827bf0b60be8aee346b515754f3645f4d7a2ef5d4b29.
//
// Solidity: event Requested(bytes32 indexed id, address indexed from, (uint64,address,uint256,bytes) call, (bool,address,uint256)[] deposits)
func (_SolveInbox *SolveInboxFilterer) FilterRequested(opts *bind.FilterOpts, id [][32]byte, from []common.Address) (*SolveInboxRequestedIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _SolveInbox.contract.FilterLogs(opts, "Requested", idRule, fromRule)
	if err != nil {
		return nil, err
	}
	return &SolveInboxRequestedIterator{contract: _SolveInbox.contract, event: "Requested", logs: logs, sub: sub}, nil
}

// WatchRequested is a free log subscription operation binding the contract event 0x3622a563ce1f96b477fa827bf0b60be8aee346b515754f3645f4d7a2ef5d4b29.
//
// Solidity: event Requested(bytes32 indexed id, address indexed from, (uint64,address,uint256,bytes) call, (bool,address,uint256)[] deposits)
func (_SolveInbox *SolveInboxFilterer) WatchRequested(opts *bind.WatchOpts, sink chan<- *SolveInboxRequested, id [][32]byte, from []common.Address) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _SolveInbox.contract.WatchLogs(opts, "Requested", idRule, fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SolveInboxRequested)
				if err := _SolveInbox.contract.UnpackLog(event, "Requested", log); err != nil {
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

// ParseRequested is a log parse operation binding the contract event 0x3622a563ce1f96b477fa827bf0b60be8aee346b515754f3645f4d7a2ef5d4b29.
//
// Solidity: event Requested(bytes32 indexed id, address indexed from, (uint64,address,uint256,bytes) call, (bool,address,uint256)[] deposits)
func (_SolveInbox *SolveInboxFilterer) ParseRequested(log types.Log) (*SolveInboxRequested, error) {
	event := new(SolveInboxRequested)
	if err := _SolveInbox.contract.UnpackLog(event, "Requested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SolveInboxRevertedIterator is returned from FilterReverted and is used to iterate over the raw logs and unpacked data for Reverted events raised by the SolveInbox contract.
type SolveInboxRevertedIterator struct {
	Event *SolveInboxReverted // Event containing the contract specifics and raw log

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
func (it *SolveInboxRevertedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SolveInboxReverted)
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
		it.Event = new(SolveInboxReverted)
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
func (it *SolveInboxRevertedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SolveInboxRevertedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SolveInboxReverted represents a Reverted event raised by the SolveInbox contract.
type SolveInboxReverted struct {
	Id  [32]byte
	Raw types.Log // Blockchain specific contextual infos
}

// FilterReverted is a free log retrieval operation binding the contract event 0xb66b13449e4bb2c30749a37f3081f1988fcee5ff5d98ce740b354d4e2d944095.
//
// Solidity: event Reverted(bytes32 indexed id)
func (_SolveInbox *SolveInboxFilterer) FilterReverted(opts *bind.FilterOpts, id [][32]byte) (*SolveInboxRevertedIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _SolveInbox.contract.FilterLogs(opts, "Reverted", idRule)
	if err != nil {
		return nil, err
	}
	return &SolveInboxRevertedIterator{contract: _SolveInbox.contract, event: "Reverted", logs: logs, sub: sub}, nil
}

// WatchReverted is a free log subscription operation binding the contract event 0xb66b13449e4bb2c30749a37f3081f1988fcee5ff5d98ce740b354d4e2d944095.
//
// Solidity: event Reverted(bytes32 indexed id)
func (_SolveInbox *SolveInboxFilterer) WatchReverted(opts *bind.WatchOpts, sink chan<- *SolveInboxReverted, id [][32]byte) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _SolveInbox.contract.WatchLogs(opts, "Reverted", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SolveInboxReverted)
				if err := _SolveInbox.contract.UnpackLog(event, "Reverted", log); err != nil {
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

// ParseReverted is a log parse operation binding the contract event 0xb66b13449e4bb2c30749a37f3081f1988fcee5ff5d98ce740b354d4e2d944095.
//
// Solidity: event Reverted(bytes32 indexed id)
func (_SolveInbox *SolveInboxFilterer) ParseReverted(log types.Log) (*SolveInboxReverted, error) {
	event := new(SolveInboxReverted)
	if err := _SolveInbox.contract.UnpackLog(event, "Reverted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SolveInboxRolesUpdatedIterator is returned from FilterRolesUpdated and is used to iterate over the raw logs and unpacked data for RolesUpdated events raised by the SolveInbox contract.
type SolveInboxRolesUpdatedIterator struct {
	Event *SolveInboxRolesUpdated // Event containing the contract specifics and raw log

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
func (it *SolveInboxRolesUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SolveInboxRolesUpdated)
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
		it.Event = new(SolveInboxRolesUpdated)
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
func (it *SolveInboxRolesUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SolveInboxRolesUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SolveInboxRolesUpdated represents a RolesUpdated event raised by the SolveInbox contract.
type SolveInboxRolesUpdated struct {
	User  common.Address
	Roles *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterRolesUpdated is a free log retrieval operation binding the contract event 0x715ad5ce61fc9595c7b415289d59cf203f23a94fa06f04af7e489a0a76e1fe26.
//
// Solidity: event RolesUpdated(address indexed user, uint256 indexed roles)
func (_SolveInbox *SolveInboxFilterer) FilterRolesUpdated(opts *bind.FilterOpts, user []common.Address, roles []*big.Int) (*SolveInboxRolesUpdatedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var rolesRule []interface{}
	for _, rolesItem := range roles {
		rolesRule = append(rolesRule, rolesItem)
	}

	logs, sub, err := _SolveInbox.contract.FilterLogs(opts, "RolesUpdated", userRule, rolesRule)
	if err != nil {
		return nil, err
	}
	return &SolveInboxRolesUpdatedIterator{contract: _SolveInbox.contract, event: "RolesUpdated", logs: logs, sub: sub}, nil
}

// WatchRolesUpdated is a free log subscription operation binding the contract event 0x715ad5ce61fc9595c7b415289d59cf203f23a94fa06f04af7e489a0a76e1fe26.
//
// Solidity: event RolesUpdated(address indexed user, uint256 indexed roles)
func (_SolveInbox *SolveInboxFilterer) WatchRolesUpdated(opts *bind.WatchOpts, sink chan<- *SolveInboxRolesUpdated, user []common.Address, roles []*big.Int) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var rolesRule []interface{}
	for _, rolesItem := range roles {
		rolesRule = append(rolesRule, rolesItem)
	}

	logs, sub, err := _SolveInbox.contract.WatchLogs(opts, "RolesUpdated", userRule, rolesRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SolveInboxRolesUpdated)
				if err := _SolveInbox.contract.UnpackLog(event, "RolesUpdated", log); err != nil {
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

// ParseRolesUpdated is a log parse operation binding the contract event 0x715ad5ce61fc9595c7b415289d59cf203f23a94fa06f04af7e489a0a76e1fe26.
//
// Solidity: event RolesUpdated(address indexed user, uint256 indexed roles)
func (_SolveInbox *SolveInboxFilterer) ParseRolesUpdated(log types.Log) (*SolveInboxRolesUpdated, error) {
	event := new(SolveInboxRolesUpdated)
	if err := _SolveInbox.contract.UnpackLog(event, "RolesUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
