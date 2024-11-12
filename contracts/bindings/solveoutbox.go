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
// autocommented by commenttypes.go
// type SolveCall struct {
// 	DestChainId uint64
// 	Target      common.Address
// 	Value       *big.Int
// 	Data        []byte
// }

// SolveTokenPrereq is an auto generated low-level Go binding around an user-defined struct.
type SolveTokenPrereq struct {
	Token   common.Address
	Spender common.Address
	Amount  *big.Int
}

// SolveOutboxMetaData contains all meta data concerning the SolveOutbox contract.
var SolveOutboxMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"allowedCalls\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"selector\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"cancelOwnershipHandover\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"completeOwnershipHandover\",\"inputs\":[{\"name\":\"pendingOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"defaultConfLevel\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"didFulfill\",\"inputs\":[{\"name\":\"reqId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"sourceChainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"call\",\"type\":\"tuple\",\"internalType\":\"structSolve.Call\",\"components\":[{\"name\":\"destChainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"fulfill\",\"inputs\":[{\"name\":\"reqId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"sourceChainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"call\",\"type\":\"tuple\",\"internalType\":\"structSolve.Call\",\"components\":[{\"name\":\"destChainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]},{\"name\":\"prereqs\",\"type\":\"tuple[]\",\"internalType\":\"structSolve.TokenPrereq[]\",\"components\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"fulfillFee\",\"inputs\":[{\"name\":\"sourceChainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"fulfilledCalls\",\"inputs\":[{\"name\":\"callHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"fulfilled\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"grantRoles\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"roles\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"hasAllRoles\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"roles\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"hasAnyRole\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"roles\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"owner_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"solver_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"omni_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"inbox_\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"omni\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIOmniPortal\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"result\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"ownershipHandoverExpiresAt\",\"inputs\":[{\"name\":\"pendingOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"result\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"renounceRoles\",\"inputs\":[{\"name\":\"roles\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"requestOwnershipHandover\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"revokeRoles\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"roles\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"rolesOf\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"roles\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setAllowedCall\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"selector\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"},{\"name\":\"allowed\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"event\",\"name\":\"AllowedCallSet\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"selector\",\"type\":\"bytes4\",\"indexed\":true,\"internalType\":\"bytes4\"},{\"name\":\"allowed\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DefaultConfLevelSet\",\"inputs\":[{\"name\":\"conf\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Fulfilled\",\"inputs\":[{\"name\":\"reqId\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"callHash\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"solvedBy\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OmniPortalSet\",\"inputs\":[{\"name\":\"omni\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipHandoverCanceled\",\"inputs\":[{\"name\":\"pendingOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipHandoverRequested\",\"inputs\":[{\"name\":\"pendingOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"oldOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RolesUpdated\",\"inputs\":[{\"name\":\"user\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"roles\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AlreadyInitialized\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AreadyFulfilled\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CallFailed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CallNotAllowed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"IncorrectPreReqs\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientFee\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NewOwnerIsZeroAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NoHandoverRequest\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"Reentrancy\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"Unauthorized\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"WrongDestChain\",\"inputs\":[]}]",
	Bin: "0x608060405234801561001057600080fd5b5061001961001e565b610080565b63409feecd198054600181161561003d5763f92ee8a96000526004601cfd5b8160c01c808260011c1461007b578060011b8355806020527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2602080a15b505050565b6115d88061008f6000396000f3fe6080604052600436106101355760003560e01c80635ba15647116100ab578063b23ade801161006f578063b23ade8014610365578063cb01a09a146103a0578063f04e283e146103b3578063f2fde38b146103c6578063f8c8765e146103d9578063fee81cf4146103f957600080fd5b80635ba15647146102c15780635db9cbe4146102e1578063715018a61461031157806374eeb847146103195780638da5cb5b1461034c57600080fd5b80632b370b67116100fd5780632b370b67146101e45780632de948071461020457806339acf9f1146102375780634a4ee7b11461026f578063514e62fc1461028257806354d1f13d146102b957600080fd5b8063183a4f6e1461013a578063188a97aa1461014f5780631c10893f146101825780631cd64df41461019557806325692962146101dc575b600080fd5b61014d61014836600461104b565b61042c565b005b34801561015b57600080fd5b5061016f61016a366004611081565b610439565b6040519081526020015b60405180910390f35b61014d6101903660046110ba565b610490565b3480156101a157600080fd5b506101cc6101b03660046110ba565b638b78c6d8600c90815260009290925260209091205481161490565b6040519015158152602001610179565b61014d6104a6565b3480156101f057600080fd5b5061014d6101ff3660046110fc565b6104f6565b34801561021057600080fd5b5061016f61021f366004611148565b638b78c6d8600c908152600091909152602090205490565b34801561024357600080fd5b50600054610257906001600160a01b031681565b6040516001600160a01b039091168152602001610179565b61014d61027d3660046110ba565b610575565b34801561028e57600080fd5b506101cc61029d3660046110ba565b638b78c6d8600c90815260009290925260209091205416151590565b61014d610587565b3480156102cd57600080fd5b506101cc6102dc36600461117b565b6105c3565b3480156102ed57600080fd5b506101cc6102fc36600461104b565b60046020526000908152604090205460ff1681565b61014d6105f0565b34801561032557600080fd5b5060005461033a90600160a01b900460ff1681565b60405160ff9091168152602001610179565b34801561035857600080fd5b50638b78c6d81954610257565b34801561037157600080fd5b506101cc6103803660046111d2565b600360209081526000928352604080842090915290825290205460ff1681565b61014d6103ae366004611205565b610604565b61014d6103c1366004611148565b610a9e565b61014d6103d4366004611148565b610adb565b3480156103e557600080fd5b5061014d6103f43660046112b7565b610b02565b34801561040557600080fd5b5061016f610414366004611148565b63389a75e1600c908152600091909152602090205490565b6104363382610bae565b50565b60405160001960248201819052604482015260009061048a90839060640160408051601f198184030181529190526020810180516001600160e01b031663019bfff160e51b179052620186a0610bba565b92915050565b610498610c38565b6104a28282610c53565b5050565b60006202a30067ffffffffffffffff164201905063389a75e1600c5233600052806020600c2055337fdbf36a107da19e49527a7176a1babf963b4b0ff8cde35ee35d6cd8f1f9ac7e1d600080a250565b6104fe610c38565b6001600160a01b03831660008181526003602090815260408083206001600160e01b0319871680855290835292819020805460ff191686151590811790915590519081529192917f4a2dc3dabd793cd88cb7b56ba4aa70196892e5b996fc72f4f3d45e20343d305b910160405180910390a3505050565b61057d610c38565b6104a28282610bae565b63389a75e1600c523360005260006020600c2055337ffa7b8eab7da67f412cc9575ed43464468f9bfbae89d1675917346ca6d8fe3c92600080a2565b6000600460006105d4868686610c5f565b815260208101919091526040016000205460ff16949350505050565b6105f8610c38565b6106026000610c95565b565b600161060f81610cd3565b3068929eee149b4bd21268540361062e5763ab143c066000526004601cfd5b3068929eee149b4bd2126855466106486020860186611081565b67ffffffffffffffff16146106705760405163fd24301760e01b815260040160405180910390fd5b600360006106846040870160208801611148565b6001600160a01b0316815260208101919091526040016000908120906106ad606087018761130b565b6106b691611359565b6001600160e01b031916815260208101919091526040016000205460ff166106f1576040516315dace2d60e21b815260040160405180910390fd5b60006106fe878787610c5f565b60008181526004602052604090205490915060ff161561073157604051630c4a31a760e01b815260040160405180910390fd5b60008367ffffffffffffffff81111561074c5761074c611389565b604051908082528060200260200182016040528015610775578160200160208202803683370190505b50905060005b848110156108bf576107be308787848181106107995761079961139f565b6107af9260206060909202019081019150611148565b6001600160a01b031690610cf9565b8282815181106107d0576107d061139f565b60200260200101818152505061083633308888858181106107f3576107f361139f565b9050606002016040013589898681811061080f5761080f61139f565b6108259260206060909202019081019150611148565b6001600160a01b0316929190610d25565b6108b786868381811061084b5761084b61139f565b90506060020160200160208101906108639190611148565b8787848181106108755761087561139f565b905060600201604001358888858181106108915761089161139f565b6108a79260206060909202019081019150611148565b6001600160a01b03169190610d83565b60010161077b565b5060006108d26040880160208901611148565b6001600160a01b031660408801356108ed60608a018a61130b565b6040516108fb9291906113b5565b60006040518083038185875af1925050503d8060008114610938576040519150601f19603f3d011682016040523d82523d6000602084013e61093d565b606091505b505090508061095f57604051633204506f60e01b815260040160405180910390fd5b60005b858110156109c05782818151811061097c5761097c61139f565b602002602001015161099a308989858181106107995761079961139f565b146109b857604051630979361760e01b815260040160405180910390fd5b600101610962565b50604051602481018a90526044810184905260009060640160408051601f198184030181529190526020810180516001600160e01b031663019bfff160e51b179052600254909150600090610a27908b906004906001600160a01b031685620186a0610dd3565b905080610a3860408b0135346113c5565b1015610a565760405162976f7560e21b815260040160405180910390fd5b604051339086908d907f7898a125e0970666c80e00bbf2e7041d84dfe5bbe6bcf562ce53d540fd6cd89190600090a450505050503868929eee149b4bd2126855505050505050565b610aa6610c38565b63389a75e1600c52806000526020600c208054421115610ace57636f5e88186000526004601cfd5b6000905561043681610c95565b610ae3610c38565b8060601b610af957637448fbae6000526004601cfd5b61043681610c95565b63409feecd198054600382558015610b395760018160011c14303b10610b305763f92ee8a96000526004601cfd5b818160ff1b1b91505b50610b4385610f17565b610b4e846001610c53565b610b5783610f53565b600280546001600160a01b0319166001600160a01b0384161790558015610ba7576002815560016020527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2602080a15b5050505050565b6104a282826000610ff2565b60008054604051632376548f60e21b81526001600160a01b0390911690638dd9523c90610bef9087908790879060040161142c565b602060405180830381865afa158015610c0c573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610c309190611463565b949350505050565b638b78c6d819543314610602576382b429006000526004601cfd5b6104a282826001610ff2565b6000838383604051602001610c76939291906114a5565b6040516020818303038152906040528051906020012090509392505050565b638b78c6d81980546001600160a01b039092169182907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0600080a355565b638b78c6d8600c5233600052806020600c205416610436576382b429006000526004601cfd5b6000816014526370a0823160601b60005260208060246010865afa601f3d111660205102905092915050565b60405181606052826040528360601b602c526323b872dd60601b600c52602060006064601c6000895af18060016000511416610d7457803d873b151710610d7457637939f4246000526004601cfd5b50600060605260405250505050565b816014528060345263095ea7b360601b60005260206000604460106000875af18060016000511416610dc857803d853b151710610dc857633e3f8f736000526004601cfd5b506000603452505050565b60008054604051632376548f60e21b815282916001600160a01b031690638dd9523c90610e08908a908890889060040161142c565b602060405180830381865afa158015610e25573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610e499190611463565b905080471015610ea05760405162461bcd60e51b815260206004820152601860248201527f584170703a20696e73756666696369656e742066756e6473000000000000000060448201526064015b60405180910390fd5b60005460405163c21dda4f60e01b81526001600160a01b039091169063c21dda4f908390610eda908b908b908b908b908b90600401611552565b6000604051808303818588803b158015610ef357600080fd5b505af1158015610f07573d6000803e3d6000fd5b50939a9950505050505050505050565b6001600160a01b0316638b78c6d8198190558060007f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e08180a350565b6001600160a01b038116610f9e5760405162461bcd60e51b8152602060048201526012602482015271584170703a206e6f207a65726f206f6d6e6960701b6044820152606401610e97565b600080546001600160a01b0319166001600160a01b0383169081179091556040519081527f79162c8d053a07e70cdc1ccc536f0440b571f8508377d2bef51094fadab98f479060200160405180910390a150565b638b78c6d8600c52826000526020600c20805483811783611014575080841681185b80835580600c5160601c7f715ad5ce61fc9595c7b415289d59cf203f23a94fa06f04af7e489a0a76e1fe26600080a3505050505050565b60006020828403121561105d57600080fd5b5035919050565b803567ffffffffffffffff8116811461107c57600080fd5b919050565b60006020828403121561109357600080fd5b61109c82611064565b9392505050565b80356001600160a01b038116811461107c57600080fd5b600080604083850312156110cd57600080fd5b6110d6836110a3565b946020939093013593505050565b80356001600160e01b03198116811461107c57600080fd5b60008060006060848603121561111157600080fd5b61111a846110a3565b9250611128602085016110e4565b91506040840135801515811461113d57600080fd5b809150509250925092565b60006020828403121561115a57600080fd5b61109c826110a3565b60006080828403121561117557600080fd5b50919050565b60008060006060848603121561119057600080fd5b833592506111a060208501611064565b9150604084013567ffffffffffffffff8111156111bc57600080fd5b6111c886828701611163565b9150509250925092565b600080604083850312156111e557600080fd5b6111ee836110a3565b91506111fc602084016110e4565b90509250929050565b60008060008060006080868803121561121d57600080fd5b8535945061122d60208701611064565b9350604086013567ffffffffffffffff8082111561124a57600080fd5b61125689838a01611163565b9450606088013591508082111561126c57600080fd5b818801915088601f83011261128057600080fd5b81358181111561128f57600080fd5b8960206060830285010111156112a457600080fd5b9699959850939650602001949392505050565b600080600080608085870312156112cd57600080fd5b6112d6856110a3565b93506112e4602086016110a3565b92506112f2604086016110a3565b9150611300606086016110a3565b905092959194509250565b6000808335601e1984360301811261132257600080fd5b83018035915067ffffffffffffffff82111561133d57600080fd5b60200191503681900382131561135257600080fd5b9250929050565b6001600160e01b031981358181169160048510156113815780818660040360031b1b83161692505b505092915050565b634e487b7160e01b600052604160045260246000fd5b634e487b7160e01b600052603260045260246000fd5b8183823760009101908152919050565b8181038181111561048a57634e487b7160e01b600052601160045260246000fd5b6000815180845260005b8181101561140c576020818501810151868301820152016113f0565b506000602082860101526020601f19601f83011685010191505092915050565b600067ffffffffffffffff80861683526060602084015261145060608401866113e6565b9150808416604084015250949350505050565b60006020828403121561147557600080fd5b5051919050565b81835281816020850137506000828201602090810191909152601f909101601f19169091010190565b838152600067ffffffffffffffff808516602084015260606040840152806114cc85611064565b1660608401526001600160a01b036114e6602086016110a3565b166080840152604084013560a08401526060840135601e1985360301811261150d57600080fd5b84016020810190358281111561152257600080fd5b80360382131561153157600080fd5b608060c086015261154660e08601828461147c565b98975050505050505050565b600067ffffffffffffffff808816835260ff8716602084015260018060a01b038616604084015260a0606084015261158d60a08401866113e6565b9150808416608084015250969550505050505056fea264697066735822122011c1d969aadfaac04426294c7ad04678461f98502e0a88851843393e77f081ec64736f6c63430008180033",
}

// SolveOutboxABI is the input ABI used to generate the binding from.
// Deprecated: Use SolveOutboxMetaData.ABI instead.
var SolveOutboxABI = SolveOutboxMetaData.ABI

// SolveOutboxBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use SolveOutboxMetaData.Bin instead.
var SolveOutboxBin = SolveOutboxMetaData.Bin

// DeploySolveOutbox deploys a new Ethereum contract, binding an instance of SolveOutbox to it.
func DeploySolveOutbox(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *SolveOutbox, error) {
	parsed, err := SolveOutboxMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(SolveOutboxBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &SolveOutbox{SolveOutboxCaller: SolveOutboxCaller{contract: contract}, SolveOutboxTransactor: SolveOutboxTransactor{contract: contract}, SolveOutboxFilterer: SolveOutboxFilterer{contract: contract}}, nil
}

// SolveOutbox is an auto generated Go binding around an Ethereum contract.
type SolveOutbox struct {
	SolveOutboxCaller     // Read-only binding to the contract
	SolveOutboxTransactor // Write-only binding to the contract
	SolveOutboxFilterer   // Log filterer for contract events
}

// SolveOutboxCaller is an auto generated read-only Go binding around an Ethereum contract.
type SolveOutboxCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SolveOutboxTransactor is an auto generated write-only Go binding around an Ethereum contract.
type SolveOutboxTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SolveOutboxFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type SolveOutboxFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// SolveOutboxSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type SolveOutboxSession struct {
	Contract     *SolveOutbox      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// SolveOutboxCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type SolveOutboxCallerSession struct {
	Contract *SolveOutboxCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// SolveOutboxTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type SolveOutboxTransactorSession struct {
	Contract     *SolveOutboxTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// SolveOutboxRaw is an auto generated low-level Go binding around an Ethereum contract.
type SolveOutboxRaw struct {
	Contract *SolveOutbox // Generic contract binding to access the raw methods on
}

// SolveOutboxCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type SolveOutboxCallerRaw struct {
	Contract *SolveOutboxCaller // Generic read-only contract binding to access the raw methods on
}

// SolveOutboxTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type SolveOutboxTransactorRaw struct {
	Contract *SolveOutboxTransactor // Generic write-only contract binding to access the raw methods on
}

// NewSolveOutbox creates a new instance of SolveOutbox, bound to a specific deployed contract.
func NewSolveOutbox(address common.Address, backend bind.ContractBackend) (*SolveOutbox, error) {
	contract, err := bindSolveOutbox(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &SolveOutbox{SolveOutboxCaller: SolveOutboxCaller{contract: contract}, SolveOutboxTransactor: SolveOutboxTransactor{contract: contract}, SolveOutboxFilterer: SolveOutboxFilterer{contract: contract}}, nil
}

// NewSolveOutboxCaller creates a new read-only instance of SolveOutbox, bound to a specific deployed contract.
func NewSolveOutboxCaller(address common.Address, caller bind.ContractCaller) (*SolveOutboxCaller, error) {
	contract, err := bindSolveOutbox(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &SolveOutboxCaller{contract: contract}, nil
}

// NewSolveOutboxTransactor creates a new write-only instance of SolveOutbox, bound to a specific deployed contract.
func NewSolveOutboxTransactor(address common.Address, transactor bind.ContractTransactor) (*SolveOutboxTransactor, error) {
	contract, err := bindSolveOutbox(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &SolveOutboxTransactor{contract: contract}, nil
}

// NewSolveOutboxFilterer creates a new log filterer instance of SolveOutbox, bound to a specific deployed contract.
func NewSolveOutboxFilterer(address common.Address, filterer bind.ContractFilterer) (*SolveOutboxFilterer, error) {
	contract, err := bindSolveOutbox(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &SolveOutboxFilterer{contract: contract}, nil
}

// bindSolveOutbox binds a generic wrapper to an already deployed contract.
func bindSolveOutbox(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := SolveOutboxMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SolveOutbox *SolveOutboxRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SolveOutbox.Contract.SolveOutboxCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SolveOutbox *SolveOutboxRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SolveOutbox.Contract.SolveOutboxTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SolveOutbox *SolveOutboxRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SolveOutbox.Contract.SolveOutboxTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_SolveOutbox *SolveOutboxCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _SolveOutbox.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_SolveOutbox *SolveOutboxTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SolveOutbox.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_SolveOutbox *SolveOutboxTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _SolveOutbox.Contract.contract.Transact(opts, method, params...)
}

// AllowedCalls is a free data retrieval call binding the contract method 0xb23ade80.
//
// Solidity: function allowedCalls(address target, bytes4 selector) view returns(bool)
func (_SolveOutbox *SolveOutboxCaller) AllowedCalls(opts *bind.CallOpts, target common.Address, selector [4]byte) (bool, error) {
	var out []interface{}
	err := _SolveOutbox.contract.Call(opts, &out, "allowedCalls", target, selector)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AllowedCalls is a free data retrieval call binding the contract method 0xb23ade80.
//
// Solidity: function allowedCalls(address target, bytes4 selector) view returns(bool)
func (_SolveOutbox *SolveOutboxSession) AllowedCalls(target common.Address, selector [4]byte) (bool, error) {
	return _SolveOutbox.Contract.AllowedCalls(&_SolveOutbox.CallOpts, target, selector)
}

// AllowedCalls is a free data retrieval call binding the contract method 0xb23ade80.
//
// Solidity: function allowedCalls(address target, bytes4 selector) view returns(bool)
func (_SolveOutbox *SolveOutboxCallerSession) AllowedCalls(target common.Address, selector [4]byte) (bool, error) {
	return _SolveOutbox.Contract.AllowedCalls(&_SolveOutbox.CallOpts, target, selector)
}

// DefaultConfLevel is a free data retrieval call binding the contract method 0x74eeb847.
//
// Solidity: function defaultConfLevel() view returns(uint8)
func (_SolveOutbox *SolveOutboxCaller) DefaultConfLevel(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _SolveOutbox.contract.Call(opts, &out, "defaultConfLevel")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// DefaultConfLevel is a free data retrieval call binding the contract method 0x74eeb847.
//
// Solidity: function defaultConfLevel() view returns(uint8)
func (_SolveOutbox *SolveOutboxSession) DefaultConfLevel() (uint8, error) {
	return _SolveOutbox.Contract.DefaultConfLevel(&_SolveOutbox.CallOpts)
}

// DefaultConfLevel is a free data retrieval call binding the contract method 0x74eeb847.
//
// Solidity: function defaultConfLevel() view returns(uint8)
func (_SolveOutbox *SolveOutboxCallerSession) DefaultConfLevel() (uint8, error) {
	return _SolveOutbox.Contract.DefaultConfLevel(&_SolveOutbox.CallOpts)
}

// DidFulfill is a free data retrieval call binding the contract method 0x5ba15647.
//
// Solidity: function didFulfill(bytes32 reqId, uint64 sourceChainId, (uint64,address,uint256,bytes) call) view returns(bool)
func (_SolveOutbox *SolveOutboxCaller) DidFulfill(opts *bind.CallOpts, reqId [32]byte, sourceChainId uint64, call SolveCall) (bool, error) {
	var out []interface{}
	err := _SolveOutbox.contract.Call(opts, &out, "didFulfill", reqId, sourceChainId, call)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// DidFulfill is a free data retrieval call binding the contract method 0x5ba15647.
//
// Solidity: function didFulfill(bytes32 reqId, uint64 sourceChainId, (uint64,address,uint256,bytes) call) view returns(bool)
func (_SolveOutbox *SolveOutboxSession) DidFulfill(reqId [32]byte, sourceChainId uint64, call SolveCall) (bool, error) {
	return _SolveOutbox.Contract.DidFulfill(&_SolveOutbox.CallOpts, reqId, sourceChainId, call)
}

// DidFulfill is a free data retrieval call binding the contract method 0x5ba15647.
//
// Solidity: function didFulfill(bytes32 reqId, uint64 sourceChainId, (uint64,address,uint256,bytes) call) view returns(bool)
func (_SolveOutbox *SolveOutboxCallerSession) DidFulfill(reqId [32]byte, sourceChainId uint64, call SolveCall) (bool, error) {
	return _SolveOutbox.Contract.DidFulfill(&_SolveOutbox.CallOpts, reqId, sourceChainId, call)
}

// FulfillFee is a free data retrieval call binding the contract method 0x188a97aa.
//
// Solidity: function fulfillFee(uint64 sourceChainId) view returns(uint256)
func (_SolveOutbox *SolveOutboxCaller) FulfillFee(opts *bind.CallOpts, sourceChainId uint64) (*big.Int, error) {
	var out []interface{}
	err := _SolveOutbox.contract.Call(opts, &out, "fulfillFee", sourceChainId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FulfillFee is a free data retrieval call binding the contract method 0x188a97aa.
//
// Solidity: function fulfillFee(uint64 sourceChainId) view returns(uint256)
func (_SolveOutbox *SolveOutboxSession) FulfillFee(sourceChainId uint64) (*big.Int, error) {
	return _SolveOutbox.Contract.FulfillFee(&_SolveOutbox.CallOpts, sourceChainId)
}

// FulfillFee is a free data retrieval call binding the contract method 0x188a97aa.
//
// Solidity: function fulfillFee(uint64 sourceChainId) view returns(uint256)
func (_SolveOutbox *SolveOutboxCallerSession) FulfillFee(sourceChainId uint64) (*big.Int, error) {
	return _SolveOutbox.Contract.FulfillFee(&_SolveOutbox.CallOpts, sourceChainId)
}

// FulfilledCalls is a free data retrieval call binding the contract method 0x5db9cbe4.
//
// Solidity: function fulfilledCalls(bytes32 callHash) view returns(bool fulfilled)
func (_SolveOutbox *SolveOutboxCaller) FulfilledCalls(opts *bind.CallOpts, callHash [32]byte) (bool, error) {
	var out []interface{}
	err := _SolveOutbox.contract.Call(opts, &out, "fulfilledCalls", callHash)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// FulfilledCalls is a free data retrieval call binding the contract method 0x5db9cbe4.
//
// Solidity: function fulfilledCalls(bytes32 callHash) view returns(bool fulfilled)
func (_SolveOutbox *SolveOutboxSession) FulfilledCalls(callHash [32]byte) (bool, error) {
	return _SolveOutbox.Contract.FulfilledCalls(&_SolveOutbox.CallOpts, callHash)
}

// FulfilledCalls is a free data retrieval call binding the contract method 0x5db9cbe4.
//
// Solidity: function fulfilledCalls(bytes32 callHash) view returns(bool fulfilled)
func (_SolveOutbox *SolveOutboxCallerSession) FulfilledCalls(callHash [32]byte) (bool, error) {
	return _SolveOutbox.Contract.FulfilledCalls(&_SolveOutbox.CallOpts, callHash)
}

// HasAllRoles is a free data retrieval call binding the contract method 0x1cd64df4.
//
// Solidity: function hasAllRoles(address user, uint256 roles) view returns(bool)
func (_SolveOutbox *SolveOutboxCaller) HasAllRoles(opts *bind.CallOpts, user common.Address, roles *big.Int) (bool, error) {
	var out []interface{}
	err := _SolveOutbox.contract.Call(opts, &out, "hasAllRoles", user, roles)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasAllRoles is a free data retrieval call binding the contract method 0x1cd64df4.
//
// Solidity: function hasAllRoles(address user, uint256 roles) view returns(bool)
func (_SolveOutbox *SolveOutboxSession) HasAllRoles(user common.Address, roles *big.Int) (bool, error) {
	return _SolveOutbox.Contract.HasAllRoles(&_SolveOutbox.CallOpts, user, roles)
}

// HasAllRoles is a free data retrieval call binding the contract method 0x1cd64df4.
//
// Solidity: function hasAllRoles(address user, uint256 roles) view returns(bool)
func (_SolveOutbox *SolveOutboxCallerSession) HasAllRoles(user common.Address, roles *big.Int) (bool, error) {
	return _SolveOutbox.Contract.HasAllRoles(&_SolveOutbox.CallOpts, user, roles)
}

// HasAnyRole is a free data retrieval call binding the contract method 0x514e62fc.
//
// Solidity: function hasAnyRole(address user, uint256 roles) view returns(bool)
func (_SolveOutbox *SolveOutboxCaller) HasAnyRole(opts *bind.CallOpts, user common.Address, roles *big.Int) (bool, error) {
	var out []interface{}
	err := _SolveOutbox.contract.Call(opts, &out, "hasAnyRole", user, roles)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasAnyRole is a free data retrieval call binding the contract method 0x514e62fc.
//
// Solidity: function hasAnyRole(address user, uint256 roles) view returns(bool)
func (_SolveOutbox *SolveOutboxSession) HasAnyRole(user common.Address, roles *big.Int) (bool, error) {
	return _SolveOutbox.Contract.HasAnyRole(&_SolveOutbox.CallOpts, user, roles)
}

// HasAnyRole is a free data retrieval call binding the contract method 0x514e62fc.
//
// Solidity: function hasAnyRole(address user, uint256 roles) view returns(bool)
func (_SolveOutbox *SolveOutboxCallerSession) HasAnyRole(user common.Address, roles *big.Int) (bool, error) {
	return _SolveOutbox.Contract.HasAnyRole(&_SolveOutbox.CallOpts, user, roles)
}

// Omni is a free data retrieval call binding the contract method 0x39acf9f1.
//
// Solidity: function omni() view returns(address)
func (_SolveOutbox *SolveOutboxCaller) Omni(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SolveOutbox.contract.Call(opts, &out, "omni")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Omni is a free data retrieval call binding the contract method 0x39acf9f1.
//
// Solidity: function omni() view returns(address)
func (_SolveOutbox *SolveOutboxSession) Omni() (common.Address, error) {
	return _SolveOutbox.Contract.Omni(&_SolveOutbox.CallOpts)
}

// Omni is a free data retrieval call binding the contract method 0x39acf9f1.
//
// Solidity: function omni() view returns(address)
func (_SolveOutbox *SolveOutboxCallerSession) Omni() (common.Address, error) {
	return _SolveOutbox.Contract.Omni(&_SolveOutbox.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address result)
func (_SolveOutbox *SolveOutboxCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _SolveOutbox.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address result)
func (_SolveOutbox *SolveOutboxSession) Owner() (common.Address, error) {
	return _SolveOutbox.Contract.Owner(&_SolveOutbox.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address result)
func (_SolveOutbox *SolveOutboxCallerSession) Owner() (common.Address, error) {
	return _SolveOutbox.Contract.Owner(&_SolveOutbox.CallOpts)
}

// OwnershipHandoverExpiresAt is a free data retrieval call binding the contract method 0xfee81cf4.
//
// Solidity: function ownershipHandoverExpiresAt(address pendingOwner) view returns(uint256 result)
func (_SolveOutbox *SolveOutboxCaller) OwnershipHandoverExpiresAt(opts *bind.CallOpts, pendingOwner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _SolveOutbox.contract.Call(opts, &out, "ownershipHandoverExpiresAt", pendingOwner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// OwnershipHandoverExpiresAt is a free data retrieval call binding the contract method 0xfee81cf4.
//
// Solidity: function ownershipHandoverExpiresAt(address pendingOwner) view returns(uint256 result)
func (_SolveOutbox *SolveOutboxSession) OwnershipHandoverExpiresAt(pendingOwner common.Address) (*big.Int, error) {
	return _SolveOutbox.Contract.OwnershipHandoverExpiresAt(&_SolveOutbox.CallOpts, pendingOwner)
}

// OwnershipHandoverExpiresAt is a free data retrieval call binding the contract method 0xfee81cf4.
//
// Solidity: function ownershipHandoverExpiresAt(address pendingOwner) view returns(uint256 result)
func (_SolveOutbox *SolveOutboxCallerSession) OwnershipHandoverExpiresAt(pendingOwner common.Address) (*big.Int, error) {
	return _SolveOutbox.Contract.OwnershipHandoverExpiresAt(&_SolveOutbox.CallOpts, pendingOwner)
}

// RolesOf is a free data retrieval call binding the contract method 0x2de94807.
//
// Solidity: function rolesOf(address user) view returns(uint256 roles)
func (_SolveOutbox *SolveOutboxCaller) RolesOf(opts *bind.CallOpts, user common.Address) (*big.Int, error) {
	var out []interface{}
	err := _SolveOutbox.contract.Call(opts, &out, "rolesOf", user)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// RolesOf is a free data retrieval call binding the contract method 0x2de94807.
//
// Solidity: function rolesOf(address user) view returns(uint256 roles)
func (_SolveOutbox *SolveOutboxSession) RolesOf(user common.Address) (*big.Int, error) {
	return _SolveOutbox.Contract.RolesOf(&_SolveOutbox.CallOpts, user)
}

// RolesOf is a free data retrieval call binding the contract method 0x2de94807.
//
// Solidity: function rolesOf(address user) view returns(uint256 roles)
func (_SolveOutbox *SolveOutboxCallerSession) RolesOf(user common.Address) (*big.Int, error) {
	return _SolveOutbox.Contract.RolesOf(&_SolveOutbox.CallOpts, user)
}

// CancelOwnershipHandover is a paid mutator transaction binding the contract method 0x54d1f13d.
//
// Solidity: function cancelOwnershipHandover() payable returns()
func (_SolveOutbox *SolveOutboxTransactor) CancelOwnershipHandover(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SolveOutbox.contract.Transact(opts, "cancelOwnershipHandover")
}

// CancelOwnershipHandover is a paid mutator transaction binding the contract method 0x54d1f13d.
//
// Solidity: function cancelOwnershipHandover() payable returns()
func (_SolveOutbox *SolveOutboxSession) CancelOwnershipHandover() (*types.Transaction, error) {
	return _SolveOutbox.Contract.CancelOwnershipHandover(&_SolveOutbox.TransactOpts)
}

// CancelOwnershipHandover is a paid mutator transaction binding the contract method 0x54d1f13d.
//
// Solidity: function cancelOwnershipHandover() payable returns()
func (_SolveOutbox *SolveOutboxTransactorSession) CancelOwnershipHandover() (*types.Transaction, error) {
	return _SolveOutbox.Contract.CancelOwnershipHandover(&_SolveOutbox.TransactOpts)
}

// CompleteOwnershipHandover is a paid mutator transaction binding the contract method 0xf04e283e.
//
// Solidity: function completeOwnershipHandover(address pendingOwner) payable returns()
func (_SolveOutbox *SolveOutboxTransactor) CompleteOwnershipHandover(opts *bind.TransactOpts, pendingOwner common.Address) (*types.Transaction, error) {
	return _SolveOutbox.contract.Transact(opts, "completeOwnershipHandover", pendingOwner)
}

// CompleteOwnershipHandover is a paid mutator transaction binding the contract method 0xf04e283e.
//
// Solidity: function completeOwnershipHandover(address pendingOwner) payable returns()
func (_SolveOutbox *SolveOutboxSession) CompleteOwnershipHandover(pendingOwner common.Address) (*types.Transaction, error) {
	return _SolveOutbox.Contract.CompleteOwnershipHandover(&_SolveOutbox.TransactOpts, pendingOwner)
}

// CompleteOwnershipHandover is a paid mutator transaction binding the contract method 0xf04e283e.
//
// Solidity: function completeOwnershipHandover(address pendingOwner) payable returns()
func (_SolveOutbox *SolveOutboxTransactorSession) CompleteOwnershipHandover(pendingOwner common.Address) (*types.Transaction, error) {
	return _SolveOutbox.Contract.CompleteOwnershipHandover(&_SolveOutbox.TransactOpts, pendingOwner)
}

// Fulfill is a paid mutator transaction binding the contract method 0xcb01a09a.
//
// Solidity: function fulfill(bytes32 reqId, uint64 sourceChainId, (uint64,address,uint256,bytes) call, (address,address,uint256)[] prereqs) payable returns()
func (_SolveOutbox *SolveOutboxTransactor) Fulfill(opts *bind.TransactOpts, reqId [32]byte, sourceChainId uint64, call SolveCall, prereqs []SolveTokenPrereq) (*types.Transaction, error) {
	return _SolveOutbox.contract.Transact(opts, "fulfill", reqId, sourceChainId, call, prereqs)
}

// Fulfill is a paid mutator transaction binding the contract method 0xcb01a09a.
//
// Solidity: function fulfill(bytes32 reqId, uint64 sourceChainId, (uint64,address,uint256,bytes) call, (address,address,uint256)[] prereqs) payable returns()
func (_SolveOutbox *SolveOutboxSession) Fulfill(reqId [32]byte, sourceChainId uint64, call SolveCall, prereqs []SolveTokenPrereq) (*types.Transaction, error) {
	return _SolveOutbox.Contract.Fulfill(&_SolveOutbox.TransactOpts, reqId, sourceChainId, call, prereqs)
}

// Fulfill is a paid mutator transaction binding the contract method 0xcb01a09a.
//
// Solidity: function fulfill(bytes32 reqId, uint64 sourceChainId, (uint64,address,uint256,bytes) call, (address,address,uint256)[] prereqs) payable returns()
func (_SolveOutbox *SolveOutboxTransactorSession) Fulfill(reqId [32]byte, sourceChainId uint64, call SolveCall, prereqs []SolveTokenPrereq) (*types.Transaction, error) {
	return _SolveOutbox.Contract.Fulfill(&_SolveOutbox.TransactOpts, reqId, sourceChainId, call, prereqs)
}

// GrantRoles is a paid mutator transaction binding the contract method 0x1c10893f.
//
// Solidity: function grantRoles(address user, uint256 roles) payable returns()
func (_SolveOutbox *SolveOutboxTransactor) GrantRoles(opts *bind.TransactOpts, user common.Address, roles *big.Int) (*types.Transaction, error) {
	return _SolveOutbox.contract.Transact(opts, "grantRoles", user, roles)
}

// GrantRoles is a paid mutator transaction binding the contract method 0x1c10893f.
//
// Solidity: function grantRoles(address user, uint256 roles) payable returns()
func (_SolveOutbox *SolveOutboxSession) GrantRoles(user common.Address, roles *big.Int) (*types.Transaction, error) {
	return _SolveOutbox.Contract.GrantRoles(&_SolveOutbox.TransactOpts, user, roles)
}

// GrantRoles is a paid mutator transaction binding the contract method 0x1c10893f.
//
// Solidity: function grantRoles(address user, uint256 roles) payable returns()
func (_SolveOutbox *SolveOutboxTransactorSession) GrantRoles(user common.Address, roles *big.Int) (*types.Transaction, error) {
	return _SolveOutbox.Contract.GrantRoles(&_SolveOutbox.TransactOpts, user, roles)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address owner_, address solver_, address omni_, address inbox_) returns()
func (_SolveOutbox *SolveOutboxTransactor) Initialize(opts *bind.TransactOpts, owner_ common.Address, solver_ common.Address, omni_ common.Address, inbox_ common.Address) (*types.Transaction, error) {
	return _SolveOutbox.contract.Transact(opts, "initialize", owner_, solver_, omni_, inbox_)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address owner_, address solver_, address omni_, address inbox_) returns()
func (_SolveOutbox *SolveOutboxSession) Initialize(owner_ common.Address, solver_ common.Address, omni_ common.Address, inbox_ common.Address) (*types.Transaction, error) {
	return _SolveOutbox.Contract.Initialize(&_SolveOutbox.TransactOpts, owner_, solver_, omni_, inbox_)
}

// Initialize is a paid mutator transaction binding the contract method 0xf8c8765e.
//
// Solidity: function initialize(address owner_, address solver_, address omni_, address inbox_) returns()
func (_SolveOutbox *SolveOutboxTransactorSession) Initialize(owner_ common.Address, solver_ common.Address, omni_ common.Address, inbox_ common.Address) (*types.Transaction, error) {
	return _SolveOutbox.Contract.Initialize(&_SolveOutbox.TransactOpts, owner_, solver_, omni_, inbox_)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() payable returns()
func (_SolveOutbox *SolveOutboxTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SolveOutbox.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() payable returns()
func (_SolveOutbox *SolveOutboxSession) RenounceOwnership() (*types.Transaction, error) {
	return _SolveOutbox.Contract.RenounceOwnership(&_SolveOutbox.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() payable returns()
func (_SolveOutbox *SolveOutboxTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _SolveOutbox.Contract.RenounceOwnership(&_SolveOutbox.TransactOpts)
}

// RenounceRoles is a paid mutator transaction binding the contract method 0x183a4f6e.
//
// Solidity: function renounceRoles(uint256 roles) payable returns()
func (_SolveOutbox *SolveOutboxTransactor) RenounceRoles(opts *bind.TransactOpts, roles *big.Int) (*types.Transaction, error) {
	return _SolveOutbox.contract.Transact(opts, "renounceRoles", roles)
}

// RenounceRoles is a paid mutator transaction binding the contract method 0x183a4f6e.
//
// Solidity: function renounceRoles(uint256 roles) payable returns()
func (_SolveOutbox *SolveOutboxSession) RenounceRoles(roles *big.Int) (*types.Transaction, error) {
	return _SolveOutbox.Contract.RenounceRoles(&_SolveOutbox.TransactOpts, roles)
}

// RenounceRoles is a paid mutator transaction binding the contract method 0x183a4f6e.
//
// Solidity: function renounceRoles(uint256 roles) payable returns()
func (_SolveOutbox *SolveOutboxTransactorSession) RenounceRoles(roles *big.Int) (*types.Transaction, error) {
	return _SolveOutbox.Contract.RenounceRoles(&_SolveOutbox.TransactOpts, roles)
}

// RequestOwnershipHandover is a paid mutator transaction binding the contract method 0x25692962.
//
// Solidity: function requestOwnershipHandover() payable returns()
func (_SolveOutbox *SolveOutboxTransactor) RequestOwnershipHandover(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _SolveOutbox.contract.Transact(opts, "requestOwnershipHandover")
}

// RequestOwnershipHandover is a paid mutator transaction binding the contract method 0x25692962.
//
// Solidity: function requestOwnershipHandover() payable returns()
func (_SolveOutbox *SolveOutboxSession) RequestOwnershipHandover() (*types.Transaction, error) {
	return _SolveOutbox.Contract.RequestOwnershipHandover(&_SolveOutbox.TransactOpts)
}

// RequestOwnershipHandover is a paid mutator transaction binding the contract method 0x25692962.
//
// Solidity: function requestOwnershipHandover() payable returns()
func (_SolveOutbox *SolveOutboxTransactorSession) RequestOwnershipHandover() (*types.Transaction, error) {
	return _SolveOutbox.Contract.RequestOwnershipHandover(&_SolveOutbox.TransactOpts)
}

// RevokeRoles is a paid mutator transaction binding the contract method 0x4a4ee7b1.
//
// Solidity: function revokeRoles(address user, uint256 roles) payable returns()
func (_SolveOutbox *SolveOutboxTransactor) RevokeRoles(opts *bind.TransactOpts, user common.Address, roles *big.Int) (*types.Transaction, error) {
	return _SolveOutbox.contract.Transact(opts, "revokeRoles", user, roles)
}

// RevokeRoles is a paid mutator transaction binding the contract method 0x4a4ee7b1.
//
// Solidity: function revokeRoles(address user, uint256 roles) payable returns()
func (_SolveOutbox *SolveOutboxSession) RevokeRoles(user common.Address, roles *big.Int) (*types.Transaction, error) {
	return _SolveOutbox.Contract.RevokeRoles(&_SolveOutbox.TransactOpts, user, roles)
}

// RevokeRoles is a paid mutator transaction binding the contract method 0x4a4ee7b1.
//
// Solidity: function revokeRoles(address user, uint256 roles) payable returns()
func (_SolveOutbox *SolveOutboxTransactorSession) RevokeRoles(user common.Address, roles *big.Int) (*types.Transaction, error) {
	return _SolveOutbox.Contract.RevokeRoles(&_SolveOutbox.TransactOpts, user, roles)
}

// SetAllowedCall is a paid mutator transaction binding the contract method 0x2b370b67.
//
// Solidity: function setAllowedCall(address target, bytes4 selector, bool allowed) returns()
func (_SolveOutbox *SolveOutboxTransactor) SetAllowedCall(opts *bind.TransactOpts, target common.Address, selector [4]byte, allowed bool) (*types.Transaction, error) {
	return _SolveOutbox.contract.Transact(opts, "setAllowedCall", target, selector, allowed)
}

// SetAllowedCall is a paid mutator transaction binding the contract method 0x2b370b67.
//
// Solidity: function setAllowedCall(address target, bytes4 selector, bool allowed) returns()
func (_SolveOutbox *SolveOutboxSession) SetAllowedCall(target common.Address, selector [4]byte, allowed bool) (*types.Transaction, error) {
	return _SolveOutbox.Contract.SetAllowedCall(&_SolveOutbox.TransactOpts, target, selector, allowed)
}

// SetAllowedCall is a paid mutator transaction binding the contract method 0x2b370b67.
//
// Solidity: function setAllowedCall(address target, bytes4 selector, bool allowed) returns()
func (_SolveOutbox *SolveOutboxTransactorSession) SetAllowedCall(target common.Address, selector [4]byte, allowed bool) (*types.Transaction, error) {
	return _SolveOutbox.Contract.SetAllowedCall(&_SolveOutbox.TransactOpts, target, selector, allowed)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) payable returns()
func (_SolveOutbox *SolveOutboxTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _SolveOutbox.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) payable returns()
func (_SolveOutbox *SolveOutboxSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _SolveOutbox.Contract.TransferOwnership(&_SolveOutbox.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) payable returns()
func (_SolveOutbox *SolveOutboxTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _SolveOutbox.Contract.TransferOwnership(&_SolveOutbox.TransactOpts, newOwner)
}

// SolveOutboxAllowedCallSetIterator is returned from FilterAllowedCallSet and is used to iterate over the raw logs and unpacked data for AllowedCallSet events raised by the SolveOutbox contract.
type SolveOutboxAllowedCallSetIterator struct {
	Event *SolveOutboxAllowedCallSet // Event containing the contract specifics and raw log

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
func (it *SolveOutboxAllowedCallSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SolveOutboxAllowedCallSet)
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
		it.Event = new(SolveOutboxAllowedCallSet)
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
func (it *SolveOutboxAllowedCallSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SolveOutboxAllowedCallSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SolveOutboxAllowedCallSet represents a AllowedCallSet event raised by the SolveOutbox contract.
type SolveOutboxAllowedCallSet struct {
	Target   common.Address
	Selector [4]byte
	Allowed  bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterAllowedCallSet is a free log retrieval operation binding the contract event 0x4a2dc3dabd793cd88cb7b56ba4aa70196892e5b996fc72f4f3d45e20343d305b.
//
// Solidity: event AllowedCallSet(address indexed target, bytes4 indexed selector, bool allowed)
func (_SolveOutbox *SolveOutboxFilterer) FilterAllowedCallSet(opts *bind.FilterOpts, target []common.Address, selector [][4]byte) (*SolveOutboxAllowedCallSetIterator, error) {

	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}
	var selectorRule []interface{}
	for _, selectorItem := range selector {
		selectorRule = append(selectorRule, selectorItem)
	}

	logs, sub, err := _SolveOutbox.contract.FilterLogs(opts, "AllowedCallSet", targetRule, selectorRule)
	if err != nil {
		return nil, err
	}
	return &SolveOutboxAllowedCallSetIterator{contract: _SolveOutbox.contract, event: "AllowedCallSet", logs: logs, sub: sub}, nil
}

// WatchAllowedCallSet is a free log subscription operation binding the contract event 0x4a2dc3dabd793cd88cb7b56ba4aa70196892e5b996fc72f4f3d45e20343d305b.
//
// Solidity: event AllowedCallSet(address indexed target, bytes4 indexed selector, bool allowed)
func (_SolveOutbox *SolveOutboxFilterer) WatchAllowedCallSet(opts *bind.WatchOpts, sink chan<- *SolveOutboxAllowedCallSet, target []common.Address, selector [][4]byte) (event.Subscription, error) {

	var targetRule []interface{}
	for _, targetItem := range target {
		targetRule = append(targetRule, targetItem)
	}
	var selectorRule []interface{}
	for _, selectorItem := range selector {
		selectorRule = append(selectorRule, selectorItem)
	}

	logs, sub, err := _SolveOutbox.contract.WatchLogs(opts, "AllowedCallSet", targetRule, selectorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SolveOutboxAllowedCallSet)
				if err := _SolveOutbox.contract.UnpackLog(event, "AllowedCallSet", log); err != nil {
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

// ParseAllowedCallSet is a log parse operation binding the contract event 0x4a2dc3dabd793cd88cb7b56ba4aa70196892e5b996fc72f4f3d45e20343d305b.
//
// Solidity: event AllowedCallSet(address indexed target, bytes4 indexed selector, bool allowed)
func (_SolveOutbox *SolveOutboxFilterer) ParseAllowedCallSet(log types.Log) (*SolveOutboxAllowedCallSet, error) {
	event := new(SolveOutboxAllowedCallSet)
	if err := _SolveOutbox.contract.UnpackLog(event, "AllowedCallSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SolveOutboxDefaultConfLevelSetIterator is returned from FilterDefaultConfLevelSet and is used to iterate over the raw logs and unpacked data for DefaultConfLevelSet events raised by the SolveOutbox contract.
type SolveOutboxDefaultConfLevelSetIterator struct {
	Event *SolveOutboxDefaultConfLevelSet // Event containing the contract specifics and raw log

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
func (it *SolveOutboxDefaultConfLevelSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SolveOutboxDefaultConfLevelSet)
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
		it.Event = new(SolveOutboxDefaultConfLevelSet)
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
func (it *SolveOutboxDefaultConfLevelSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SolveOutboxDefaultConfLevelSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SolveOutboxDefaultConfLevelSet represents a DefaultConfLevelSet event raised by the SolveOutbox contract.
type SolveOutboxDefaultConfLevelSet struct {
	Conf uint8
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterDefaultConfLevelSet is a free log retrieval operation binding the contract event 0x8de08a798b4e50b4f351c1eaa91a11530043802be3ffac2df87db0c45a2e8483.
//
// Solidity: event DefaultConfLevelSet(uint8 conf)
func (_SolveOutbox *SolveOutboxFilterer) FilterDefaultConfLevelSet(opts *bind.FilterOpts) (*SolveOutboxDefaultConfLevelSetIterator, error) {

	logs, sub, err := _SolveOutbox.contract.FilterLogs(opts, "DefaultConfLevelSet")
	if err != nil {
		return nil, err
	}
	return &SolveOutboxDefaultConfLevelSetIterator{contract: _SolveOutbox.contract, event: "DefaultConfLevelSet", logs: logs, sub: sub}, nil
}

// WatchDefaultConfLevelSet is a free log subscription operation binding the contract event 0x8de08a798b4e50b4f351c1eaa91a11530043802be3ffac2df87db0c45a2e8483.
//
// Solidity: event DefaultConfLevelSet(uint8 conf)
func (_SolveOutbox *SolveOutboxFilterer) WatchDefaultConfLevelSet(opts *bind.WatchOpts, sink chan<- *SolveOutboxDefaultConfLevelSet) (event.Subscription, error) {

	logs, sub, err := _SolveOutbox.contract.WatchLogs(opts, "DefaultConfLevelSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SolveOutboxDefaultConfLevelSet)
				if err := _SolveOutbox.contract.UnpackLog(event, "DefaultConfLevelSet", log); err != nil {
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
func (_SolveOutbox *SolveOutboxFilterer) ParseDefaultConfLevelSet(log types.Log) (*SolveOutboxDefaultConfLevelSet, error) {
	event := new(SolveOutboxDefaultConfLevelSet)
	if err := _SolveOutbox.contract.UnpackLog(event, "DefaultConfLevelSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SolveOutboxFulfilledIterator is returned from FilterFulfilled and is used to iterate over the raw logs and unpacked data for Fulfilled events raised by the SolveOutbox contract.
type SolveOutboxFulfilledIterator struct {
	Event *SolveOutboxFulfilled // Event containing the contract specifics and raw log

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
func (it *SolveOutboxFulfilledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SolveOutboxFulfilled)
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
		it.Event = new(SolveOutboxFulfilled)
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
func (it *SolveOutboxFulfilledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SolveOutboxFulfilledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SolveOutboxFulfilled represents a Fulfilled event raised by the SolveOutbox contract.
type SolveOutboxFulfilled struct {
	ReqId    [32]byte
	CallHash [32]byte
	SolvedBy common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterFulfilled is a free log retrieval operation binding the contract event 0x7898a125e0970666c80e00bbf2e7041d84dfe5bbe6bcf562ce53d540fd6cd891.
//
// Solidity: event Fulfilled(bytes32 indexed reqId, bytes32 indexed callHash, address indexed solvedBy)
func (_SolveOutbox *SolveOutboxFilterer) FilterFulfilled(opts *bind.FilterOpts, reqId [][32]byte, callHash [][32]byte, solvedBy []common.Address) (*SolveOutboxFulfilledIterator, error) {

	var reqIdRule []interface{}
	for _, reqIdItem := range reqId {
		reqIdRule = append(reqIdRule, reqIdItem)
	}
	var callHashRule []interface{}
	for _, callHashItem := range callHash {
		callHashRule = append(callHashRule, callHashItem)
	}
	var solvedByRule []interface{}
	for _, solvedByItem := range solvedBy {
		solvedByRule = append(solvedByRule, solvedByItem)
	}

	logs, sub, err := _SolveOutbox.contract.FilterLogs(opts, "Fulfilled", reqIdRule, callHashRule, solvedByRule)
	if err != nil {
		return nil, err
	}
	return &SolveOutboxFulfilledIterator{contract: _SolveOutbox.contract, event: "Fulfilled", logs: logs, sub: sub}, nil
}

// WatchFulfilled is a free log subscription operation binding the contract event 0x7898a125e0970666c80e00bbf2e7041d84dfe5bbe6bcf562ce53d540fd6cd891.
//
// Solidity: event Fulfilled(bytes32 indexed reqId, bytes32 indexed callHash, address indexed solvedBy)
func (_SolveOutbox *SolveOutboxFilterer) WatchFulfilled(opts *bind.WatchOpts, sink chan<- *SolveOutboxFulfilled, reqId [][32]byte, callHash [][32]byte, solvedBy []common.Address) (event.Subscription, error) {

	var reqIdRule []interface{}
	for _, reqIdItem := range reqId {
		reqIdRule = append(reqIdRule, reqIdItem)
	}
	var callHashRule []interface{}
	for _, callHashItem := range callHash {
		callHashRule = append(callHashRule, callHashItem)
	}
	var solvedByRule []interface{}
	for _, solvedByItem := range solvedBy {
		solvedByRule = append(solvedByRule, solvedByItem)
	}

	logs, sub, err := _SolveOutbox.contract.WatchLogs(opts, "Fulfilled", reqIdRule, callHashRule, solvedByRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SolveOutboxFulfilled)
				if err := _SolveOutbox.contract.UnpackLog(event, "Fulfilled", log); err != nil {
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
// Solidity: event Fulfilled(bytes32 indexed reqId, bytes32 indexed callHash, address indexed solvedBy)
func (_SolveOutbox *SolveOutboxFilterer) ParseFulfilled(log types.Log) (*SolveOutboxFulfilled, error) {
	event := new(SolveOutboxFulfilled)
	if err := _SolveOutbox.contract.UnpackLog(event, "Fulfilled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SolveOutboxInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the SolveOutbox contract.
type SolveOutboxInitializedIterator struct {
	Event *SolveOutboxInitialized // Event containing the contract specifics and raw log

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
func (it *SolveOutboxInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SolveOutboxInitialized)
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
		it.Event = new(SolveOutboxInitialized)
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
func (it *SolveOutboxInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SolveOutboxInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SolveOutboxInitialized represents a Initialized event raised by the SolveOutbox contract.
type SolveOutboxInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_SolveOutbox *SolveOutboxFilterer) FilterInitialized(opts *bind.FilterOpts) (*SolveOutboxInitializedIterator, error) {

	logs, sub, err := _SolveOutbox.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &SolveOutboxInitializedIterator{contract: _SolveOutbox.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_SolveOutbox *SolveOutboxFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *SolveOutboxInitialized) (event.Subscription, error) {

	logs, sub, err := _SolveOutbox.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SolveOutboxInitialized)
				if err := _SolveOutbox.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_SolveOutbox *SolveOutboxFilterer) ParseInitialized(log types.Log) (*SolveOutboxInitialized, error) {
	event := new(SolveOutboxInitialized)
	if err := _SolveOutbox.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SolveOutboxOmniPortalSetIterator is returned from FilterOmniPortalSet and is used to iterate over the raw logs and unpacked data for OmniPortalSet events raised by the SolveOutbox contract.
type SolveOutboxOmniPortalSetIterator struct {
	Event *SolveOutboxOmniPortalSet // Event containing the contract specifics and raw log

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
func (it *SolveOutboxOmniPortalSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SolveOutboxOmniPortalSet)
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
		it.Event = new(SolveOutboxOmniPortalSet)
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
func (it *SolveOutboxOmniPortalSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SolveOutboxOmniPortalSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SolveOutboxOmniPortalSet represents a OmniPortalSet event raised by the SolveOutbox contract.
type SolveOutboxOmniPortalSet struct {
	Omni common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterOmniPortalSet is a free log retrieval operation binding the contract event 0x79162c8d053a07e70cdc1ccc536f0440b571f8508377d2bef51094fadab98f47.
//
// Solidity: event OmniPortalSet(address omni)
func (_SolveOutbox *SolveOutboxFilterer) FilterOmniPortalSet(opts *bind.FilterOpts) (*SolveOutboxOmniPortalSetIterator, error) {

	logs, sub, err := _SolveOutbox.contract.FilterLogs(opts, "OmniPortalSet")
	if err != nil {
		return nil, err
	}
	return &SolveOutboxOmniPortalSetIterator{contract: _SolveOutbox.contract, event: "OmniPortalSet", logs: logs, sub: sub}, nil
}

// WatchOmniPortalSet is a free log subscription operation binding the contract event 0x79162c8d053a07e70cdc1ccc536f0440b571f8508377d2bef51094fadab98f47.
//
// Solidity: event OmniPortalSet(address omni)
func (_SolveOutbox *SolveOutboxFilterer) WatchOmniPortalSet(opts *bind.WatchOpts, sink chan<- *SolveOutboxOmniPortalSet) (event.Subscription, error) {

	logs, sub, err := _SolveOutbox.contract.WatchLogs(opts, "OmniPortalSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SolveOutboxOmniPortalSet)
				if err := _SolveOutbox.contract.UnpackLog(event, "OmniPortalSet", log); err != nil {
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
func (_SolveOutbox *SolveOutboxFilterer) ParseOmniPortalSet(log types.Log) (*SolveOutboxOmniPortalSet, error) {
	event := new(SolveOutboxOmniPortalSet)
	if err := _SolveOutbox.contract.UnpackLog(event, "OmniPortalSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SolveOutboxOwnershipHandoverCanceledIterator is returned from FilterOwnershipHandoverCanceled and is used to iterate over the raw logs and unpacked data for OwnershipHandoverCanceled events raised by the SolveOutbox contract.
type SolveOutboxOwnershipHandoverCanceledIterator struct {
	Event *SolveOutboxOwnershipHandoverCanceled // Event containing the contract specifics and raw log

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
func (it *SolveOutboxOwnershipHandoverCanceledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SolveOutboxOwnershipHandoverCanceled)
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
		it.Event = new(SolveOutboxOwnershipHandoverCanceled)
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
func (it *SolveOutboxOwnershipHandoverCanceledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SolveOutboxOwnershipHandoverCanceledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SolveOutboxOwnershipHandoverCanceled represents a OwnershipHandoverCanceled event raised by the SolveOutbox contract.
type SolveOutboxOwnershipHandoverCanceled struct {
	PendingOwner common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterOwnershipHandoverCanceled is a free log retrieval operation binding the contract event 0xfa7b8eab7da67f412cc9575ed43464468f9bfbae89d1675917346ca6d8fe3c92.
//
// Solidity: event OwnershipHandoverCanceled(address indexed pendingOwner)
func (_SolveOutbox *SolveOutboxFilterer) FilterOwnershipHandoverCanceled(opts *bind.FilterOpts, pendingOwner []common.Address) (*SolveOutboxOwnershipHandoverCanceledIterator, error) {

	var pendingOwnerRule []interface{}
	for _, pendingOwnerItem := range pendingOwner {
		pendingOwnerRule = append(pendingOwnerRule, pendingOwnerItem)
	}

	logs, sub, err := _SolveOutbox.contract.FilterLogs(opts, "OwnershipHandoverCanceled", pendingOwnerRule)
	if err != nil {
		return nil, err
	}
	return &SolveOutboxOwnershipHandoverCanceledIterator{contract: _SolveOutbox.contract, event: "OwnershipHandoverCanceled", logs: logs, sub: sub}, nil
}

// WatchOwnershipHandoverCanceled is a free log subscription operation binding the contract event 0xfa7b8eab7da67f412cc9575ed43464468f9bfbae89d1675917346ca6d8fe3c92.
//
// Solidity: event OwnershipHandoverCanceled(address indexed pendingOwner)
func (_SolveOutbox *SolveOutboxFilterer) WatchOwnershipHandoverCanceled(opts *bind.WatchOpts, sink chan<- *SolveOutboxOwnershipHandoverCanceled, pendingOwner []common.Address) (event.Subscription, error) {

	var pendingOwnerRule []interface{}
	for _, pendingOwnerItem := range pendingOwner {
		pendingOwnerRule = append(pendingOwnerRule, pendingOwnerItem)
	}

	logs, sub, err := _SolveOutbox.contract.WatchLogs(opts, "OwnershipHandoverCanceled", pendingOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SolveOutboxOwnershipHandoverCanceled)
				if err := _SolveOutbox.contract.UnpackLog(event, "OwnershipHandoverCanceled", log); err != nil {
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
func (_SolveOutbox *SolveOutboxFilterer) ParseOwnershipHandoverCanceled(log types.Log) (*SolveOutboxOwnershipHandoverCanceled, error) {
	event := new(SolveOutboxOwnershipHandoverCanceled)
	if err := _SolveOutbox.contract.UnpackLog(event, "OwnershipHandoverCanceled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SolveOutboxOwnershipHandoverRequestedIterator is returned from FilterOwnershipHandoverRequested and is used to iterate over the raw logs and unpacked data for OwnershipHandoverRequested events raised by the SolveOutbox contract.
type SolveOutboxOwnershipHandoverRequestedIterator struct {
	Event *SolveOutboxOwnershipHandoverRequested // Event containing the contract specifics and raw log

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
func (it *SolveOutboxOwnershipHandoverRequestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SolveOutboxOwnershipHandoverRequested)
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
		it.Event = new(SolveOutboxOwnershipHandoverRequested)
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
func (it *SolveOutboxOwnershipHandoverRequestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SolveOutboxOwnershipHandoverRequestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SolveOutboxOwnershipHandoverRequested represents a OwnershipHandoverRequested event raised by the SolveOutbox contract.
type SolveOutboxOwnershipHandoverRequested struct {
	PendingOwner common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterOwnershipHandoverRequested is a free log retrieval operation binding the contract event 0xdbf36a107da19e49527a7176a1babf963b4b0ff8cde35ee35d6cd8f1f9ac7e1d.
//
// Solidity: event OwnershipHandoverRequested(address indexed pendingOwner)
func (_SolveOutbox *SolveOutboxFilterer) FilterOwnershipHandoverRequested(opts *bind.FilterOpts, pendingOwner []common.Address) (*SolveOutboxOwnershipHandoverRequestedIterator, error) {

	var pendingOwnerRule []interface{}
	for _, pendingOwnerItem := range pendingOwner {
		pendingOwnerRule = append(pendingOwnerRule, pendingOwnerItem)
	}

	logs, sub, err := _SolveOutbox.contract.FilterLogs(opts, "OwnershipHandoverRequested", pendingOwnerRule)
	if err != nil {
		return nil, err
	}
	return &SolveOutboxOwnershipHandoverRequestedIterator{contract: _SolveOutbox.contract, event: "OwnershipHandoverRequested", logs: logs, sub: sub}, nil
}

// WatchOwnershipHandoverRequested is a free log subscription operation binding the contract event 0xdbf36a107da19e49527a7176a1babf963b4b0ff8cde35ee35d6cd8f1f9ac7e1d.
//
// Solidity: event OwnershipHandoverRequested(address indexed pendingOwner)
func (_SolveOutbox *SolveOutboxFilterer) WatchOwnershipHandoverRequested(opts *bind.WatchOpts, sink chan<- *SolveOutboxOwnershipHandoverRequested, pendingOwner []common.Address) (event.Subscription, error) {

	var pendingOwnerRule []interface{}
	for _, pendingOwnerItem := range pendingOwner {
		pendingOwnerRule = append(pendingOwnerRule, pendingOwnerItem)
	}

	logs, sub, err := _SolveOutbox.contract.WatchLogs(opts, "OwnershipHandoverRequested", pendingOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SolveOutboxOwnershipHandoverRequested)
				if err := _SolveOutbox.contract.UnpackLog(event, "OwnershipHandoverRequested", log); err != nil {
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
func (_SolveOutbox *SolveOutboxFilterer) ParseOwnershipHandoverRequested(log types.Log) (*SolveOutboxOwnershipHandoverRequested, error) {
	event := new(SolveOutboxOwnershipHandoverRequested)
	if err := _SolveOutbox.contract.UnpackLog(event, "OwnershipHandoverRequested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SolveOutboxOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the SolveOutbox contract.
type SolveOutboxOwnershipTransferredIterator struct {
	Event *SolveOutboxOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *SolveOutboxOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SolveOutboxOwnershipTransferred)
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
		it.Event = new(SolveOutboxOwnershipTransferred)
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
func (it *SolveOutboxOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SolveOutboxOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SolveOutboxOwnershipTransferred represents a OwnershipTransferred event raised by the SolveOutbox contract.
type SolveOutboxOwnershipTransferred struct {
	OldOwner common.Address
	NewOwner common.Address
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed oldOwner, address indexed newOwner)
func (_SolveOutbox *SolveOutboxFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, oldOwner []common.Address, newOwner []common.Address) (*SolveOutboxOwnershipTransferredIterator, error) {

	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _SolveOutbox.contract.FilterLogs(opts, "OwnershipTransferred", oldOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &SolveOutboxOwnershipTransferredIterator{contract: _SolveOutbox.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed oldOwner, address indexed newOwner)
func (_SolveOutbox *SolveOutboxFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *SolveOutboxOwnershipTransferred, oldOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var oldOwnerRule []interface{}
	for _, oldOwnerItem := range oldOwner {
		oldOwnerRule = append(oldOwnerRule, oldOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _SolveOutbox.contract.WatchLogs(opts, "OwnershipTransferred", oldOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SolveOutboxOwnershipTransferred)
				if err := _SolveOutbox.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_SolveOutbox *SolveOutboxFilterer) ParseOwnershipTransferred(log types.Log) (*SolveOutboxOwnershipTransferred, error) {
	event := new(SolveOutboxOwnershipTransferred)
	if err := _SolveOutbox.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// SolveOutboxRolesUpdatedIterator is returned from FilterRolesUpdated and is used to iterate over the raw logs and unpacked data for RolesUpdated events raised by the SolveOutbox contract.
type SolveOutboxRolesUpdatedIterator struct {
	Event *SolveOutboxRolesUpdated // Event containing the contract specifics and raw log

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
func (it *SolveOutboxRolesUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(SolveOutboxRolesUpdated)
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
		it.Event = new(SolveOutboxRolesUpdated)
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
func (it *SolveOutboxRolesUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *SolveOutboxRolesUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// SolveOutboxRolesUpdated represents a RolesUpdated event raised by the SolveOutbox contract.
type SolveOutboxRolesUpdated struct {
	User  common.Address
	Roles *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterRolesUpdated is a free log retrieval operation binding the contract event 0x715ad5ce61fc9595c7b415289d59cf203f23a94fa06f04af7e489a0a76e1fe26.
//
// Solidity: event RolesUpdated(address indexed user, uint256 indexed roles)
func (_SolveOutbox *SolveOutboxFilterer) FilterRolesUpdated(opts *bind.FilterOpts, user []common.Address, roles []*big.Int) (*SolveOutboxRolesUpdatedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var rolesRule []interface{}
	for _, rolesItem := range roles {
		rolesRule = append(rolesRule, rolesItem)
	}

	logs, sub, err := _SolveOutbox.contract.FilterLogs(opts, "RolesUpdated", userRule, rolesRule)
	if err != nil {
		return nil, err
	}
	return &SolveOutboxRolesUpdatedIterator{contract: _SolveOutbox.contract, event: "RolesUpdated", logs: logs, sub: sub}, nil
}

// WatchRolesUpdated is a free log subscription operation binding the contract event 0x715ad5ce61fc9595c7b415289d59cf203f23a94fa06f04af7e489a0a76e1fe26.
//
// Solidity: event RolesUpdated(address indexed user, uint256 indexed roles)
func (_SolveOutbox *SolveOutboxFilterer) WatchRolesUpdated(opts *bind.WatchOpts, sink chan<- *SolveOutboxRolesUpdated, user []common.Address, roles []*big.Int) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}
	var rolesRule []interface{}
	for _, rolesItem := range roles {
		rolesRule = append(rolesRule, rolesItem)
	}

	logs, sub, err := _SolveOutbox.contract.WatchLogs(opts, "RolesUpdated", userRule, rolesRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(SolveOutboxRolesUpdated)
				if err := _SolveOutbox.contract.UnpackLog(event, "RolesUpdated", log); err != nil {
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
func (_SolveOutbox *SolveOutboxFilterer) ParseRolesUpdated(log types.Log) (*SolveOutboxRolesUpdated, error) {
	event := new(SolveOutboxRolesUpdated)
	if err := _SolveOutbox.contract.UnpackLog(event, "RolesUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
