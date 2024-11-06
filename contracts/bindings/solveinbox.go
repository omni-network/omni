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
	UpdatedAt  *big.Int
	Status     uint8
	From       common.Address
	AcceptedBy common.Address
	Call       SolveCall
	Deposits   []SolveDeposit
}

// SolveTokenDeposit is an auto generated low-level Go binding around an user-defined struct.
type SolveTokenDeposit struct {
	Token  common.Address
	Amount *big.Int
}

// SolveInboxMetaData contains all meta data concerning the SolveInbox contract.
var SolveInboxMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"accept\",\"inputs\":[{\"name\":\"id\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"cancel\",\"inputs\":[{\"name\":\"id\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"cancelOwnershipHandover\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"claim\",\"inputs\":[{\"name\":\"id\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"completeOwnershipHandover\",\"inputs\":[{\"name\":\"pendingOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"defaultConfLevel\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRequest\",\"inputs\":[{\"name\":\"id\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structSolve.Request\",\"components\":[{\"name\":\"id\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"updatedAt\",\"type\":\"uint40\",\"internalType\":\"uint40\"},{\"name\":\"status\",\"type\":\"uint8\",\"internalType\":\"enumSolve.Status\"},{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"acceptedBy\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"call\",\"type\":\"tuple\",\"internalType\":\"structSolve.Call\",\"components\":[{\"name\":\"destChainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"deposits\",\"type\":\"tuple[]\",\"internalType\":\"structSolve.Deposit[]\",\"components\":[{\"name\":\"isNative\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"grantRoles\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"roles\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"hasAllRoles\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"roles\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"hasAnyRole\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"roles\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"owner_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"solver_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"omni_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"outbox_\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"markFulfilled\",\"inputs\":[{\"name\":\"id\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"callHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"creditTo\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"omni\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIOmniPortal\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"result\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"ownershipHandoverExpiresAt\",\"inputs\":[{\"name\":\"pendingOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"result\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"reject\",\"inputs\":[{\"name\":\"id\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"reason\",\"type\":\"uint8\",\"internalType\":\"enumSolve.RejectReason\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"renounceRoles\",\"inputs\":[{\"name\":\"roles\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"request\",\"inputs\":[{\"name\":\"call\",\"type\":\"tuple\",\"internalType\":\"structSolve.Call\",\"components\":[{\"name\":\"destChainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"deposits\",\"type\":\"tuple[]\",\"internalType\":\"structSolve.TokenDeposit[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[{\"name\":\"id\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"requestOwnershipHandover\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"revokeRoles\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"roles\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"rolesOf\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"roles\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"suggestNativePayment\",\"inputs\":[{\"name\":\"call\",\"type\":\"tuple\",\"internalType\":\"structSolve.Call\",\"components\":[{\"name\":\"destChainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"gasLimit\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"gasPrice\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"fulfillFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"event\",\"name\":\"Accepted\",\"inputs\":[{\"name\":\"id\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"by\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Claimed\",\"inputs\":[{\"name\":\"id\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DefaultConfLevelSet\",\"inputs\":[{\"name\":\"conf\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Fulfilled\",\"inputs\":[{\"name\":\"id\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"callHash\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"creditedTo\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OmniPortalSet\",\"inputs\":[{\"name\":\"omni\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipHandoverCanceled\",\"inputs\":[{\"name\":\"pendingOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipHandoverRequested\",\"inputs\":[{\"name\":\"pendingOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"oldOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Rejected\",\"inputs\":[{\"name\":\"id\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"by\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"reason\",\"type\":\"uint8\",\"indexed\":true,\"internalType\":\"enumSolve.RejectReason\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Requested\",\"inputs\":[{\"name\":\"id\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"call\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structSolve.Call\",\"components\":[{\"name\":\"destChainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"deposits\",\"type\":\"tuple[]\",\"indexed\":false,\"internalType\":\"structSolve.Deposit[]\",\"components\":[{\"name\":\"isNative\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Reverted\",\"inputs\":[{\"name\":\"id\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RolesUpdated\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"roles\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AlreadyInitialized\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"IncorrectChain\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidDeposit\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NewOwnerIsZeroAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NoDeposits\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NoHandoverRequest\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"Reentrancy\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"RequestStateInvalid\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"TransferFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"Unauthorized\",\"inputs\":[]}]",
	Bin: "0x608060405234801561001057600080fd5b5061001961001e565b610080565b63409feecd198054600181161561003d5763f92ee8a96000526004601cfd5b8160c01c808260011c1461007b578060011b8355806020527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2602080a15b505050565b6124a480620000906000396000f3fe60806040526004361061014b5760003560e01c806374eeb847116100b6578063e4725ba11161006f578063e4725ba114610387578063f04e283e146103a7578063f2fde38b146103ba578063f8c8765e146103cd578063fb1e61ca146103ed578063fee81cf41461041a57600080fd5b806374eeb847146102c85780638da5cb5b146102fb578063bd66528a14610314578063c4d252f514610334578063db3ea55314610354578063e17771c81461037457600080fd5b806339acf9f11161010857806339acf9f1146102165780634a4ee7b11461024e5780634aa59afd14610261578063514e62fc1461028157806354d1f13d146102b8578063715018a6146102c057600080fd5b8063183a4f6e146101505780631c10893f146101655780631cd64df41461017857806325692962146101ad5780632de94807146101b55780632e88da3f146101f6575b600080fd5b61016361015e366004611acf565b61044d565b005b610163610173366004611afd565b61045a565b34801561018457600080fd5b50610198610193366004611afd565b610470565b60405190151581526020015b60405180910390f35b61016361048f565b3480156101c157600080fd5b506101e86101d0366004611b29565b638b78c6d8600c908152600091909152602090205490565b6040519081526020016101a4565b34801561020257600080fd5b50610163610211366004611b4d565b6104de565b34801561022257600080fd5b50600054610236906001600160a01b031681565b6040516001600160a01b0390911681526020016101a4565b61016361025c366004611afd565b61074f565b34801561026d57600080fd5b506101e861027c366004611bb3565b610761565b34801561028d57600080fd5b5061019861029c366004611afd565b638b78c6d8600c90815260009290925260209091205416151590565b6101636109bd565b6101636109f9565b3480156102d457600080fd5b506000546102e990600160a01b900460ff1681565b60405160ff90911681526020016101a4565b34801561030757600080fd5b50638b78c6d81954610236565b34801561032057600080fd5b5061016361032f366004611acf565b610a0d565b34801561034057600080fd5b5061016361034f366004611acf565b610b7e565b34801561036057600080fd5b5061016361036f366004611c1c565b610d47565b6101e8610382366004611c50565b610e3f565b34801561039357600080fd5b506101636103a2366004611acf565b610fb3565b6101636103b5366004611b29565b6110ad565b6101636103c8366004611b29565b6110ea565b3480156103d957600080fd5b506101636103e8366004611ce9565b611111565b3480156103f957600080fd5b5061040d610408366004611acf565b6111bd565b6040516101a49190611e3a565b34801561042657600080fd5b506101e8610435366004611b29565b63389a75e1600c908152600091909152602090205490565b61045733826113bf565b50565b6104626113cb565b61046c82826113e6565b5050565b638b78c6d8600c90815260008390526020902054811681145b92915050565b60006202a3006001600160401b03164201905063389a75e1600c5233600052806020600c2055337fdbf36a107da19e49527a7176a1babf963b4b0ff8cde35ee35d6cd8f1f9ac7e1d600080a250565b60005460408051631799380760e11b815281516001600160a01b0390931692632f32700e926004808401939192918290030181865afa158015610525573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906105499190611f08565b8051600180546020909301516001600160a01b0316600160401b026001600160e01b03199093166001600160401b039092169190911791909117905568929eee149b4bd21268543090036105a55763ab143c066000526004601cfd5b3068929eee149b4bd2126855600083815260046020526040902060026001820154600160281b900460ff1660068111156105e1576105e1611d45565b146105ff57604051632d47b60360e11b815260040160405180910390fd5b60015460038201546001600160401b03908116911614610632576040516377d30c8b60e11b815260040160405180910390fd5b600354600154600160401b90046001600160a01b03908116911614610669576040516282b42960e81b815260040160405180910390fd5b60008446836003016040516020016106839392919061205f565b6040516020818303038152906040528051906020012090508381146106bb5760405163574b16a760e11b815260040160405180910390fd5b60018201805464ffffffffff421665ffffffffffff1990911617650500000000001790556002820180546001600160a01b0319166001600160a01b038516908117909155604051859087907f7898a125e0970666c80e00bbf2e7041d84dfe5bbe6bcf562ce53d540fd6cd89190600090a450503868929eee149b4bd21268555050600180546001600160e01b031916905550565b6107576113cb565b61046c82826113bf565b600080546040805163500b19e760e01b8152905183926001600160a01b03169163500b19e79160048083019260209291908290030181865afa1580156107ab573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906107cf919061207e565b90506000816001600160a01b0316638f9d6ace6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610811573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610835919061209b565b6001600160a01b038316638b7bfd7061085160208b018b6120b4565b6040516001600160e01b031960e084901b1681526001600160401b039091166004820152602401602060405180830381865afa158015610895573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906108b9919061209b565b6108c79060408a01356120e7565b6108d191906120fe565b60008054919250906001600160a01b0316638dd9523c6108f460208b018b6120b4565b61090160608c018c612120565b8b6040518563ffffffff1660e01b8152600401610921949392919061216d565b602060405180830381865afa15801561093e573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610962919061209b565b905060006109728761d6d86121b9565b6001600160401b03169050655af3107a400086818361099186886121e4565b61099b91906121e4565b6109a591906121e4565b6109af91906121e4565b9a9950505050505050505050565b63389a75e1600c523360005260006020600c2055337ffa7b8eab7da67f412cc9575ed43464468f9bfbae89d1675917346ca6d8fe3c92600080a2565b610a016113cb565b610a0b60006113f2565b565b3068929eee149b4bd212685403610a2c5763ab143c066000526004601cfd5b3068929eee149b4bd2126855600081815260046020526040902060056001820154600160281b900460ff166006811115610a6857610a68611d45565b14610a8657604051632d47b60360e11b815260040160405180910390fd5b6001810180546506000000000065ffffffffffff199091164264ffffffffff1617179055600281015460068201805460408051602080840282018101909252828152610b43946001600160a01b0316939092909160009084015b82821015610b3a5760008481526020908190206040805160608101825260028602909201805460ff81161515845261010090046001600160a01b031683850152600190810154918301919091529083529092019101610ae0565b50505050611430565b60405182907feddf608ef698454af2fb41c1df7b7e5154ff0d46969f895e0f39c7dfe7e6380a90600090a2503868929eee149b4bd212685550565b3068929eee149b4bd212685403610b9d5763ab143c066000526004601cfd5b3068929eee149b4bd21268556000818152600460205260409020600180820154600160281b900460ff166006811115610bd857610bd8611d45565b14158015610c06575060036001820154600160281b900460ff166006811115610c0357610c03611d45565b14155b15610c2457604051632d47b60360e11b815260040160405180910390fd5b6001810154600160301b90046001600160a01b03163314610c57576040516282b42960e81b815260040160405180910390fd5b6001810180546504000000000065ffffffffffff199091164264ffffffffff1617179081905560068201805460408051602080840282018101909252828152610d0c94600160301b90046001600160a01b0316939092909160009084018215610b3a5760008481526020908190206040805160608101825260028602909201805460ff81161515845261010090046001600160a01b031683850152600190810154918301919091529083529092019101610ae0565b60405182907fb66b13449e4bb2c30749a37f3081f1988fcee5ff5d98ce740b354d4e2d94409590600090a2503868929eee149b4bd212685550565b6001610d528161155b565b3068929eee149b4bd212685403610d715763ab143c066000526004601cfd5b3068929eee149b4bd21268556000838152600460205260409020600180820154600160281b900460ff166006811115610dac57610dac611d45565b14610dca57604051632d47b60360e11b815260040160405180910390fd5b6001810180546503000000000065ffffffffffff199091164264ffffffffff1617179055826003811115610e0057610e00611d45565b604051339086907f21f84ee3a6e9bc7c10f855f8c9829e22c613861cef10add09eccdbc88df9f59f90600090a4503868929eee149b4bd2126855505050565b60003068929eee149b4bd212685403610e605763ab143c066000526004601cfd5b3068929eee149b4bd21268556000610e7e6040860160208701611b29565b6001600160a01b031603610ea55760405163574b16a760e11b815260040160405180910390fd5b610eb260208501856120b4565b6001600160401b0316600003610edb5760405163574b16a760e11b815260040160405180910390fd5b610ee86060850185612120565b9050600003610f0a5760405163574b16a760e11b815260040160405180910390fd5b81158015610f16575034155b15610f3457604051630558800760e21b815260040160405180910390fd5b6000610f4233868686611581565b60018101548154604051929350600160301b9091046001600160a01b0316917f3622a563ce1f96b477fa827bf0b60be8aee346b515754f3645f4d7a2ef5d4b2990610f9690600386019060068701906121f7565b60405180910390a3543868929eee149b4bd2126855949350505050565b6001610fbe8161155b565b3068929eee149b4bd212685403610fdd5763ab143c066000526004601cfd5b3068929eee149b4bd21268556000828152600460205260409020600180820154600160281b900460ff16600681111561101857611018611d45565b1461103657604051632d47b60360e11b815260040160405180910390fd5b60018101805464ffffffffff421665ffffffffffff1990911617650200000000001790556002810180546001600160a01b0319163390811790915560405184907f9deed34441ca75bb2dbbe101d2201930f40e18a9ce521c77fbdca6690a89996790600090a3503868929eee149b4bd21268555050565b6110b56113cb565b63389a75e1600c52806000526020600c2080544211156110dd57636f5e88186000526004601cfd5b60009055610457816113f2565b6110f26113cb565b8060601b61110857637448fbae6000526004601cfd5b610457816113f2565b63409feecd1980546003825580156111485760018160011c14303b1061113f5763f92ee8a96000526004601cfd5b818160ff1b1b91505b5061115285611841565b61115d8460016113e6565b6111668361187d565b600380546001600160a01b0319166001600160a01b03841617905580156111b6576002815560016020527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2602080a15b5050505050565b6111c5611a45565b600082815260046020908152604091829020825160e08101845281548152600182015464ffffffffff81169382019390935292909190830190600160281b900460ff16600681111561121957611219611d45565b600681111561122a5761122a611d45565b815260018201546001600160a01b03600160301b90910481166020808401919091526002840154821660408085019190915280516080810182526003860180546001600160401b0381168352600160401b900490941692810192909252600485015490820152600584018054606094850194929392840191906112ac90611f66565b80601f01602080910402602001604051908101604052809291908181526020018280546112d890611f66565b80156113255780601f106112fa57610100808354040283529160200191611325565b820191906000526020600020905b81548152906001019060200180831161130857829003601f168201915b505050505081525050815260200160068201805480602002602001604051908101604052809291908181526020016000905b828210156113b15760008481526020908190206040805160608101825260028602909201805460ff81161515845261010090046001600160a01b031683850152600190810154918301919091529083529092019101611357565b505050915250909392505050565b61046c82826000611920565b638b78c6d819543314610a0b576382b429006000526004601cfd5b61046c82826001611920565b638b78c6d81980546001600160a01b039092169182907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0600080a355565b60005b81518110156115565781818151811061144e5761144e61227b565b602002602001015160000151156114f6576000836001600160a01b031683838151811061147d5761147d61227b565b60200260200101516040015160405160006040518083038185875af1925050503d80600081146114c9576040519150601f19603f3d011682016040523d82523d6000602084013e6114ce565b606091505b50509050806114f0576040516312171d8360e31b815260040160405180910390fd5b5061154e565b61154e8383838151811061150c5761150c61227b565b60200260200101516040015184848151811061152a5761152a61227b565b6020026020010151602001516001600160a01b03166119799092919063ffffffff16565b600101611433565b505050565b638b78c6d8600c5233600052806020600c205416610457576382b429006000526004601cfd5b60008061158c6119c9565b600081815260046020526040902081815560018101805464ffffffffff421665ffffffffffff1990911617600160281b176601000000000000600160d01b031916600160301b6001600160a01b038b16021790559250905084600383016115f382826123a0565b5050341561167457604080516060810182526001808252600060208084018281523495850195865260068801805480860182559084529190922093516002909102909301805491516001600160a81b0319909216931515610100600160a81b031916939093176101006001600160a01b039092169190910217825591519101555b60005b83811015611837578484828181106116915761169161227b565b905060400201602001356000036116bb57604051635972996f60e11b815260040160405180910390fd5b60008585838181106116cf576116cf61227b565b6116e59260206040909202019081019150611b29565b6001600160a01b03160361170c57604051635972996f60e11b815260040160405180910390fd5b8260060160405180606001604052806000151581526020018787858181106117365761173661227b565b61174c9260206040909202019081019150611b29565b6001600160a01b0316815260200187878581811061176c5761176c61227b565b60206040918202939093018301359093525083546001818101865560009586529482902084516002909202018054928501516001600160a01b031661010002610100600160a81b0319921515929092166001600160a81b03199093169290921717815591015191015561182f33308787858181106117ec576117ec61227b565b905060400201602001358888868181106118085761180861227b565b61181e9260206040909202019081019150611b29565b6001600160a01b03169291906119e7565b600101611677565b5050949350505050565b6001600160a01b0316638b78c6d8198190558060007f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e08180a350565b6001600160a01b0381166118cc5760405162461bcd60e51b8152602060048201526012602482015271584170703a206e6f207a65726f206f6d6e6960701b604482015260640160405180910390fd5b600080546001600160a01b0319166001600160a01b0383169081179091556040519081527f79162c8d053a07e70cdc1ccc536f0440b571f8508377d2bef51094fadab98f479060200160405180910390a150565b638b78c6d8600c52826000526020600c20805483811783611942575080841681185b80835580600c5160601c7f715ad5ce61fc9595c7b415289d59cf203f23a94fa06f04af7e489a0a76e1fe26600080a3505050505050565b816014528060345263a9059cbb60601b60005260206000604460106000875af180600160005114166119be57803d853b1517106119be576390b8ec186000526004601cfd5b506000603452505050565b60028054600091826119da83612455565b9091555050600254919050565b60405181606052826040528360601b602c526323b872dd60601b600c52602060006064601c6000895af18060016000511416611a3657803d873b151710611a3657637939f4246000526004601cfd5b50600060605260405250505050565b6040805160e0810182526000808252602082018190529091820190815260200160006001600160a01b0316815260200160006001600160a01b03168152602001611ac2604051806080016040528060006001600160401b0316815260200160006001600160a01b0316815260200160008152602001606081525090565b8152602001606081525090565b600060208284031215611ae157600080fd5b5035919050565b6001600160a01b038116811461045757600080fd5b60008060408385031215611b1057600080fd5b8235611b1b81611ae8565b946020939093013593505050565b600060208284031215611b3b57600080fd5b8135611b4681611ae8565b9392505050565b600080600060608486031215611b6257600080fd5b83359250602084013591506040840135611b7b81611ae8565b809150509250925092565b600060808284031215611b9857600080fd5b50919050565b6001600160401b038116811461045757600080fd5b60008060008060808587031215611bc957600080fd5b84356001600160401b03811115611bdf57600080fd5b611beb87828801611b86565b9450506020850135611bfc81611b9e565b92506040850135611c0c81611b9e565b9396929550929360600135925050565b60008060408385031215611c2f57600080fd5b82359150602083013560048110611c4557600080fd5b809150509250929050565b600080600060408486031215611c6557600080fd5b83356001600160401b0380821115611c7c57600080fd5b611c8887838801611b86565b94506020860135915080821115611c9e57600080fd5b818601915086601f830112611cb257600080fd5b813581811115611cc157600080fd5b8760208260061b8501011115611cd657600080fd5b6020830194508093505050509250925092565b60008060008060808587031215611cff57600080fd5b8435611d0a81611ae8565b93506020850135611d1a81611ae8565b92506040850135611d2a81611ae8565b91506060850135611d3a81611ae8565b939692955090935050565b634e487b7160e01b600052602160045260246000fd5b6001600160401b0381511682526000602060018060a01b03602084015116602085015260408301516040850152606083015160806060860152805180608087015260005b81811015611dbb5782810184015187820160a001528301611d9f565b50600060a0828801015260a0601f19601f830116870101935050505092915050565b60008151808452602080850194506020840160005b83811015611e2f578151805115158852838101516001600160a01b0316848901526040908101519088015260609096019590820190600101611df2565b509495945050505050565b602081528151602082015264ffffffffff60208301511660408201526000604083015160078110611e7b57634e487b7160e01b600052602160045260246000fd5b806060840152506060830151611e9c60808401826001600160a01b03169052565b5060808301516001600160a01b03811660a08401525060a083015160e060c0840152611ecc610100840182611d5b565b905060c0840151601f198483030160e0850152611ee98282611ddd565b95945050505050565b634e487b7160e01b600052604160045260246000fd5b600060408284031215611f1a57600080fd5b604051604081018181106001600160401b0382111715611f3c57611f3c611ef2565b6040528251611f4a81611b9e565b81526020830151611f5a81611ae8565b60208201529392505050565b600181811c90821680611f7a57607f821691505b602082108103611b9857634e487b7160e01b600052602260045260246000fd5b600081546001600160401b0381168452602060018060a01b038260401c1660208601526001915060018401546040860152600284016080606087015260008154611fe381611f66565b8060808a015260a06001831660008114612004576001811461202057612050565b60ff19841660a08c015260a083151560051b8c01019450612050565b85600052602060002060005b848110156120475781548d820185015290890190880161202c565b8c0160a0019550505b50929998505050505050505050565b838152826020820152606060408201526000611ee96060830184611f9a565b60006020828403121561209057600080fd5b8151611b4681611ae8565b6000602082840312156120ad57600080fd5b5051919050565b6000602082840312156120c657600080fd5b8135611b4681611b9e565b634e487b7160e01b600052601160045260246000fd5b8082028115828204841417610489576104896120d1565b60008261211b57634e487b7160e01b600052601260045260246000fd5b500490565b6000808335601e1984360301811261213757600080fd5b8301803591506001600160401b0382111561215157600080fd5b60200191503681900382131561216657600080fd5b9250929050565b60006001600160401b038087168352606060208401528460608401528486608085013760008386016080908101919091529316604083015250601f909201601f19169091010192915050565b6001600160401b038181168382160280821691908281146121dc576121dc6120d1565b505092915050565b80820180821115610489576104896120d1565b6000604080835261220b6040840186611f9a565b60208482036020860152818654808452602084019150876000526020600020935060005b8181101561226d57845460ff81161515845260081c6001600160a01b031684840152600180860154878501526002909501946060909301920161222f565b509098975050505050505050565b634e487b7160e01b600052603260045260246000fd5b601f821115611556576000816000526020600020601f850160051c810160208610156122ba5750805b601f850160051c820191505b818110156122d9578281556001016122c6565b505050505050565b6001600160401b038311156122f8576122f8611ef2565b61230c836123068354611f66565b83612291565b6000601f84116001811461234057600085156123285750838201355b600019600387901b1c1916600186901b1783556111b6565b600083815260209020601f19861690835b828110156123715786850135825560209485019460019092019101612351565b508682101561238e5760001960f88860031b161c19848701351681555b505060018560011b0183555050505050565b81356123ab81611b9e565b6001600160401b0380821691508254826001600160401b0319821617845560208501356123d781611ae8565b6001600160e01b031991909116909217604092831b68010000000000000000600160e01b0316178355908301356001830155600090606084013536859003601e19018112612423578283fd5b8401803582811115612433578384fd5b602082019250803603831315612447578384fd5b6122d98184600288016122e1565b600060018201612467576124676120d1565b506001019056fea264697066735822122071e974b8b2ee71bad0138e945badd54211f8b7458f589516893c63a9c2628a8a64736f6c63430008180033",
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

// GetRequest is a free data retrieval call binding the contract method 0xfb1e61ca.
//
// Solidity: function getRequest(bytes32 id) view returns((bytes32,uint40,uint8,address,address,(uint64,address,uint256,bytes),(bool,address,uint256)[]))
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
// Solidity: function getRequest(bytes32 id) view returns((bytes32,uint40,uint8,address,address,(uint64,address,uint256,bytes),(bool,address,uint256)[]))
func (_SolveInbox *SolveInboxSession) GetRequest(id [32]byte) (SolveRequest, error) {
	return _SolveInbox.Contract.GetRequest(&_SolveInbox.CallOpts, id)
}

// GetRequest is a free data retrieval call binding the contract method 0xfb1e61ca.
//
// Solidity: function getRequest(bytes32 id) view returns((bytes32,uint40,uint8,address,address,(uint64,address,uint256,bytes),(bool,address,uint256)[]))
func (_SolveInbox *SolveInboxCallerSession) GetRequest(id [32]byte) (SolveRequest, error) {
	return _SolveInbox.Contract.GetRequest(&_SolveInbox.CallOpts, id)
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

// SuggestNativePayment is a free data retrieval call binding the contract method 0x4aa59afd.
//
// Solidity: function suggestNativePayment((uint64,address,uint256,bytes) call, uint64 gasLimit, uint64 gasPrice, uint256 fulfillFee) view returns(uint256)
func (_SolveInbox *SolveInboxCaller) SuggestNativePayment(opts *bind.CallOpts, call SolveCall, gasLimit uint64, gasPrice uint64, fulfillFee *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _SolveInbox.contract.Call(opts, &out, "suggestNativePayment", call, gasLimit, gasPrice, fulfillFee)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// SuggestNativePayment is a free data retrieval call binding the contract method 0x4aa59afd.
//
// Solidity: function suggestNativePayment((uint64,address,uint256,bytes) call, uint64 gasLimit, uint64 gasPrice, uint256 fulfillFee) view returns(uint256)
func (_SolveInbox *SolveInboxSession) SuggestNativePayment(call SolveCall, gasLimit uint64, gasPrice uint64, fulfillFee *big.Int) (*big.Int, error) {
	return _SolveInbox.Contract.SuggestNativePayment(&_SolveInbox.CallOpts, call, gasLimit, gasPrice, fulfillFee)
}

// SuggestNativePayment is a free data retrieval call binding the contract method 0x4aa59afd.
//
// Solidity: function suggestNativePayment((uint64,address,uint256,bytes) call, uint64 gasLimit, uint64 gasPrice, uint256 fulfillFee) view returns(uint256)
func (_SolveInbox *SolveInboxCallerSession) SuggestNativePayment(call SolveCall, gasLimit uint64, gasPrice uint64, fulfillFee *big.Int) (*big.Int, error) {
	return _SolveInbox.Contract.SuggestNativePayment(&_SolveInbox.CallOpts, call, gasLimit, gasPrice, fulfillFee)
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

// Claim is a paid mutator transaction binding the contract method 0xbd66528a.
//
// Solidity: function claim(bytes32 id) returns()
func (_SolveInbox *SolveInboxTransactor) Claim(opts *bind.TransactOpts, id [32]byte) (*types.Transaction, error) {
	return _SolveInbox.contract.Transact(opts, "claim", id)
}

// Claim is a paid mutator transaction binding the contract method 0xbd66528a.
//
// Solidity: function claim(bytes32 id) returns()
func (_SolveInbox *SolveInboxSession) Claim(id [32]byte) (*types.Transaction, error) {
	return _SolveInbox.Contract.Claim(&_SolveInbox.TransactOpts, id)
}

// Claim is a paid mutator transaction binding the contract method 0xbd66528a.
//
// Solidity: function claim(bytes32 id) returns()
func (_SolveInbox *SolveInboxTransactorSession) Claim(id [32]byte) (*types.Transaction, error) {
	return _SolveInbox.Contract.Claim(&_SolveInbox.TransactOpts, id)
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

// MarkFulfilled is a paid mutator transaction binding the contract method 0x2e88da3f.
//
// Solidity: function markFulfilled(bytes32 id, bytes32 callHash, address creditTo) returns()
func (_SolveInbox *SolveInboxTransactor) MarkFulfilled(opts *bind.TransactOpts, id [32]byte, callHash [32]byte, creditTo common.Address) (*types.Transaction, error) {
	return _SolveInbox.contract.Transact(opts, "markFulfilled", id, callHash, creditTo)
}

// MarkFulfilled is a paid mutator transaction binding the contract method 0x2e88da3f.
//
// Solidity: function markFulfilled(bytes32 id, bytes32 callHash, address creditTo) returns()
func (_SolveInbox *SolveInboxSession) MarkFulfilled(id [32]byte, callHash [32]byte, creditTo common.Address) (*types.Transaction, error) {
	return _SolveInbox.Contract.MarkFulfilled(&_SolveInbox.TransactOpts, id, callHash, creditTo)
}

// MarkFulfilled is a paid mutator transaction binding the contract method 0x2e88da3f.
//
// Solidity: function markFulfilled(bytes32 id, bytes32 callHash, address creditTo) returns()
func (_SolveInbox *SolveInboxTransactorSession) MarkFulfilled(id [32]byte, callHash [32]byte, creditTo common.Address) (*types.Transaction, error) {
	return _SolveInbox.Contract.MarkFulfilled(&_SolveInbox.TransactOpts, id, callHash, creditTo)
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
	Id  [32]byte
	Raw types.Log // Blockchain specific contextual infos
}

// FilterClaimed is a free log retrieval operation binding the contract event 0xeddf608ef698454af2fb41c1df7b7e5154ff0d46969f895e0f39c7dfe7e6380a.
//
// Solidity: event Claimed(bytes32 indexed id)
func (_SolveInbox *SolveInboxFilterer) FilterClaimed(opts *bind.FilterOpts, id [][32]byte) (*SolveInboxClaimedIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _SolveInbox.contract.FilterLogs(opts, "Claimed", idRule)
	if err != nil {
		return nil, err
	}
	return &SolveInboxClaimedIterator{contract: _SolveInbox.contract, event: "Claimed", logs: logs, sub: sub}, nil
}

// WatchClaimed is a free log subscription operation binding the contract event 0xeddf608ef698454af2fb41c1df7b7e5154ff0d46969f895e0f39c7dfe7e6380a.
//
// Solidity: event Claimed(bytes32 indexed id)
func (_SolveInbox *SolveInboxFilterer) WatchClaimed(opts *bind.WatchOpts, sink chan<- *SolveInboxClaimed, id [][32]byte) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _SolveInbox.contract.WatchLogs(opts, "Claimed", idRule)
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

// ParseClaimed is a log parse operation binding the contract event 0xeddf608ef698454af2fb41c1df7b7e5154ff0d46969f895e0f39c7dfe7e6380a.
//
// Solidity: event Claimed(bytes32 indexed id)
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
