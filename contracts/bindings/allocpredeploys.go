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

// AllocPredeploysConfig is an auto generated low-level Go binding around an user-defined struct.
type AllocPredeploysConfig struct {
	Admin                  common.Address
	EnableStakingAllowlist bool
	Output                 string
}

// AllocPredeploysMetaData contains all meta data concerning the AllocPredeploys contract.
var AllocPredeploysMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"IS_SCRIPT\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"runWithCfg\",\"inputs\":[{\"name\":\"config\",\"type\":\"tuple\",\"internalType\":\"structAllocPredeploys.Config\",\"components\":[{\"name\":\"admin\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"enableStakingAllowlist\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"output\",\"type\":\"string\",\"internalType\":\"string\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setProxies\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setUp\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"}]",
	Bin: "0x6080604052600c805462ff00ff19166201000117905534801561002157600080fd5b50611bfb806100316000396000f3fe608060405234801561001057600080fd5b506004361061004c5760003560e01c80630a9254e4146100515780634a05b3f21461005b578063db5d166614610063578063f8ccbf4714610076575b600080fd5b61005961009d565b005b6100596100e8565b6100596100713660046115e3565b61015e565b600c546100899062010000900460ff1681565b604051901515815260200160405180910390f35b6100c6604051806040016040528060088152602001673232b83637bcb2b960c11b81525061049c565b600f80546001600160a01b0319166001600160a01b0392909216919091179055565b6101186040518060400160405280600f81526020016e53657474696e672070726f7869657360881b8152506104ae565b60006101226104f4565b905060005b815181101561015a5761015282828151811061014557610145611625565b602002602001015161058b565b600101610127565b5050565b80600d61016b82826117f6565b5050600f546040516303223eab60e11b81526001600160a01b039091166004820152600080516020611b60833981519152906306447d5690602401600060405180830381600087803b1580156101c057600080fd5b505af11580156101d4573d6000803e3d6000fd5b505050506101e06100e8565b6101e8610823565b6101f06109d1565b6101f8610bad565b610200610e40565b610208610f63565b610210611135565b7f885cb69240a935d632d79c317109709ecfa91a80626ff3989d68f67f5b1dd12d60001c6001600160a01b03166390c5013b6040518163ffffffff1660e01b8152600401600060405180830381600087803b15801561026e57600080fd5b505af1158015610282573d6000803e3d6000fd5b5050604051631c72346d60e01b8152336004820152600080516020611b608339815191529250631c72346d9150602401600060405180830381600087803b1580156102cc57600080fd5b505af11580156102e0573d6000803e3d6000fd5b505060405163c88a5e6d60e01b815233600482015260006024820152600080516020611b60833981519152925063c88a5e6d9150604401600060405180830381600087803b15801561033157600080fd5b505af1158015610345573d6000803e3d6000fd5b5050600f54604051631c72346d60e01b81526001600160a01b039091166004820152600080516020611b608339815191529250631c72346d9150602401600060405180830381600087803b15801561039c57600080fd5b505af11580156103b0573d6000803e3d6000fd5b5050600f5460405163c88a5e6d60e01b81526001600160a01b03909116600482015260006024820152600080516020611b60833981519152925063c88a5e6d9150604401600060405180830381600087803b15801561040e57600080fd5b505af1158015610422573d6000803e3d6000fd5b50600080516020611b60833981519152925063709ecd3f915061044a90506040840184611650565b6040518363ffffffff1660e01b81526004016104679291906118a7565b600060405180830381600087803b15801561048157600080fd5b505af1158015610495573d6000803e3d6000fd5b5050505050565b60006104a7826111de565b5092915050565b6104f1816040516024016104c29190611926565b60408051601f198184030181529190526020810180516001600160e01b031663104c13eb60e21b1790526112e8565b50565b604080516002808252606080830184529260208301908036833701905050905062048789608a1b8160008151811061052e5761052e611625565b60200260200101906001600160a01b031690816001600160a01b03168152505062333333608a1b8160018151811061056857610568611625565b60200260200101906001600160a01b031690816001600160a01b03168152505090565b63ffffffff8116156105d85760405162461bcd60e51b8152602060048201526011602482015270696e76616c6964206e616d65737061636560781b60448201526064015b60405180910390fd5b604051630fafdced60e21b815260206004820152603b60248201527f5472616e73706172656e745570677261646561626c6550726f78792e736f6c3a60448201527f5472616e73706172656e745570677261646561626c6550726f787900000000006064820152600090600080516020611b6083398151915290633ebf73b490608401600060405180830381865afa158015610678573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526106a09190810190611939565b90506106c6604051806060016040528060238152602001611b80602391396008846112f1565b60015b6008816001600160a01b03161161081e5760006106e682856119fc565b9050600362048789608a1b016001600160a01b0382160361073d5761073760405180604001604052806014815260200173536b697070696e672070726f787920617420257360601b81525082611338565b5061080c565b604051635a6b63c160e11b8152600080516020611b608339815191529063b4d6c782906107709084908790600401611a1c565b600060405180830381600087803b15801561078a57600080fd5b505af115801561079e573d6000803e3d6000fd5b505050506107c08173aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa61137d565b6107c9816113e5565b1561080a5760006107d982611472565b90506107fe604051806060016040528060238152602001611ba36023913983836114da565b6108088282611521565b505b505b8061081681611a48565b9150506106c9565b505050565b6108566040518060400160405280601281526020017129b2ba3a34b73390283937bc3ca0b236b4b760711b8152506104ae565b600d546040516370ca10bb60e01b815273aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa60048201526000602482018190526001600160a01b039092166044820152600080516020611b60833981519152906370ca10bb90606401600060405180830381600087803b1580156108cc57600080fd5b505af11580156108e0573d6000803e3d6000fd5b5050604051630fafdced60e21b815260206004820152602960248201527f6f75742f50726f787941646d696e2e736f6c2f50726f787941646d696e2e302e6044820152681c17191a173539b7b760b91b6064820152600080516020611b60833981519152925063b4d6c782915073aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa908390633ebf73b4906084015b600060405180830381865afa15801561098c573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526109b49190810190611939565b6040518363ffffffff1660e01b8152600401610467929190611a1c565b610a086040518060400160405280601681526020017553657474696e6720506f7274616c526567697374727960501b8152506104ae565b6000610a1c600162048789608a1b01611472565b604051630fafdced60e21b815260206004820152602160248201527f506f7274616c52656769737472792e736f6c3a506f7274616c526567697374726044820152607960f81b6064820152909150600080516020611b608339815191529063b4d6c7829083908390633ebf73b490608401600060405180830381865afa158015610aaa573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052610ad29190810190611939565b6040518363ffffffff1660e01b8152600401610aef929190611a1c565b600060405180830381600087803b158015610b0957600080fd5b505af1158015610b1d573d6000803e3d6000fd5b50505050806001600160a01b03166331f449006040518163ffffffff1660e01b8152600401600060405180830381600087803b158015610b5c57600080fd5b505af1158015610b70573d6000803e3d6000fd5b5050600d5460405163189acdbd60e31b81526001600160a01b039091166004820152600162048789608a1b01925063c4d66de89150602401610467565b610beb6040518060400160405280601881526020017f53657474696e67204f6d6e694272696467654e617469766500000000000000008152506104ae565b60405163c88a5e6d60e01b8152600262048789608a1b01600482015269d3c21bcecceda10000006024820181905290600080516020611b608339815191529063c88a5e6d90604401600060405180830381600087803b158015610c4d57600080fd5b505af1158015610c61573d6000803e3d6000fd5b505050506000610c79600262048789608a1b01611472565b604051630fafdced60e21b815260206004820152602560248201527f4f6d6e694272696467654e61746976652e736f6c3a4f6d6e694272696467654e604482015264617469766560d81b6064820152909150600080516020611b608339815191529063b4d6c7829083908390633ebf73b490608401600060405180830381865afa158015610d0b573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052610d339190810190611939565b6040518363ffffffff1660e01b8152600401610d50929190611a1c565b600060405180830381600087803b158015610d6a57600080fd5b505af1158015610d7e573d6000803e3d6000fd5b50505050806001600160a01b03166331f449006040518163ffffffff1660e01b8152600401600060405180830381600087803b158015610dbd57600080fd5b505af1158015610dd1573d6000803e3d6000fd5b5050600d5460405163189acdbd60e31b81526001600160a01b039091166004820152600262048789608a1b01925063c4d66de891506024015b600060405180830381600087803b158015610e2457600080fd5b505af1158015610e38573d6000803e3d6000fd5b505050505050565b610e6e6040518060400160405280600d81526020016c53657474696e6720574f6d6e6960981b8152506104ae565b604051630fafdced60e21b815260206004820152600f60248201526e574f6d6e692e736f6c3a574f6d6e6960881b6044820152600080516020611b608339815191529063b4d6c78290600362048789608a1b01908390633ebf73b490606401600060405180830381865afa158015610eea573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052610f129190810190611939565b6040518363ffffffff1660e01b8152600401610f2f929190611a1c565b600060405180830381600087803b158015610f4957600080fd5b505af1158015610f5d573d6000803e3d6000fd5b50505050565b610f936040518060400160405280600f81526020016e53657474696e67205374616b696e6760881b8152506104ae565b6000610fa7600162333333608a1b01611472565b604051630fafdced60e21b81526020600482015260136024820152725374616b696e672e736f6c3a5374616b696e6760681b6044820152909150600080516020611b608339815191529063b4d6c7829083908390633ebf73b490606401600060405180830381865afa158015611021573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526110499190810190611939565b6040518363ffffffff1660e01b8152600401611066929190611a1c565b600060405180830381600087803b15801561108057600080fd5b505af1158015611094573d6000803e3d6000fd5b50505050806001600160a01b03166331f449006040518163ffffffff1660e01b8152600401600060405180830381600087803b1580156110d357600080fd5b505af11580156110e7573d6000803e3d6000fd5b5050600d5460405163400ada7560e01b81526001600160a01b0382166004820152600160a01b90910460ff1615156024820152600162333333608a1b01925063400ada759150604401610467565b6111666040518060400160405280601081526020016f53657474696e6720536c617368696e6760801b8152506104ae565b600061117a600262333333608a1b01611472565b604051630fafdced60e21b8152602060048201526015602482015274536c617368696e672e736f6c3a536c617368696e6760581b6044820152909150600080516020611b608339815191529063b4d6c7829083908390633ebf73b49060640161096f565b600080826040516020016111f29190611a76565b60408051808303601f190181529082905280516020909101206001625e79b760e01b03198252600482018190529150600080516020611b608339815191529063ffa1864990602401602060405180830381865afa158015611257573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061127b9190611a92565b6040516318caf8e360e31b8152909250600080516020611b608339815191529063c657c718906112b19085908790600401611a1c565b600060405180830381600087803b1580156112cb57600080fd5b505af11580156112df573d6000803e3d6000fd5b50505050915091565b6104f181611589565b61081e83838360405160240161130993929190611aaf565b60408051601f198184030181529190526020810180516001600160e01b031663038fd88960e31b1790526112e8565b61015a828260405160240161134e929190611ae2565b60408051601f198184030181529190526020810180516001600160e01b031663319af33360e01b1790526112e8565b6040516370ca10bb60e01b81526001600160a01b0380841660048301527fb53127684a568b3173ae13b9f8a6016e243e63b6e8ee1178d6a717850b5d6103602483015282166044820152600080516020611b60833981519152906370ca10bb90606401610e0a565b60006001600160a01b038216600162048789608a1b01148061141857506001600160a01b038216600262048789608a1b01145b8061143457506001600160a01b038216600362048789608a1b01145b8061145057506001600160a01b038216600162333333608a1b01145b8061146c57506001600160a01b038216600262333333608a1b01145b92915050565b600061147d826115aa565b6114c95760405162461bcd60e51b815260206004820152601b60248201527f5072656465706c6f79733a206e6f742061207072656465706c6f79000000000060448201526064016105cf565b61146c826001600160a01b03611b0c565b61081e8383836040516024016114f293929190611b2c565b60408051601f198184030181529190526020810180516001600160e01b03166307e763af60e51b1790526112e8565b6040516370ca10bb60e01b81526001600160a01b0380841660048301527f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc602483015282166044820152600080516020611b60833981519152906370ca10bb90606401610e0a565b80516a636f6e736f6c652e6c6f67602083016000808483855afa5050505050565b600062048789607f1b600b83901c6001600160951b0316148061146c575062333333607f1b600b83901c6001600160951b03161461146c565b6000602082840312156115f557600080fd5b813567ffffffffffffffff81111561160c57600080fd5b82016060818503121561161e57600080fd5b9392505050565b634e487b7160e01b600052603260045260246000fd5b6001600160a01b03811681146104f157600080fd5b6000808335601e1984360301811261166757600080fd5b83018035915067ffffffffffffffff82111561168257600080fd5b60200191503681900382131561169757600080fd5b9250929050565b634e487b7160e01b600052604160045260246000fd5b600181811c908216806116c857607f821691505b6020821081036116e857634e487b7160e01b600052602260045260246000fd5b50919050565b601f82111561081e576000816000526020600020601f850160051c810160208610156117175750805b601f850160051c820191505b81811015610e3857828155600101611723565b67ffffffffffffffff83111561174e5761174e61169e565b6117628361175c83546116b4565b836116ee565b6000601f841160018114611796576000851561177e5750838201355b600019600387901b1c1916600186901b178355610495565b600083815260209020601f19861690835b828110156117c757868501358255602094850194600190920191016117a7565b50868210156117e45760001960f88860031b161c19848701351681555b505060018560011b0183555050505050565b81356118018161163b565b81546001600160a01b031981166001600160a01b03929092169182178355602084013580151580821461183357600080fd5b6001600160a81b03199290921690921760a09190911b60ff60a01b1617825550604082013536839003601e1901811261186b57600080fd5b8201803567ffffffffffffffff81111561188457600080fd5b60208201915080360382131561189957600080fd5b610f5d818360018601611736565b60208152816020820152818360408301376000818301604090810191909152601f909201601f19160101919050565b60005b838110156118f15781810151838201526020016118d9565b50506000910152565b600081518084526119128160208601602086016118d6565b601f01601f19169290920160200192915050565b60208152600061161e60208301846118fa565b60006020828403121561194b57600080fd5b815167ffffffffffffffff8082111561196357600080fd5b818401915084601f83011261197757600080fd5b8151818111156119895761198961169e565b604051601f8201601f19908116603f011681019083821181831017156119b1576119b161169e565b816040528281528760208487010111156119ca57600080fd5b6119db8360208301602088016118d6565b979650505050505050565b634e487b7160e01b600052601160045260246000fd5b6001600160a01b038181168382160190808211156104a7576104a76119e6565b6001600160a01b0383168152604060208201819052600090611a40908301846118fa565b949350505050565b60006001600160a01b038281166002600160a01b03198101611a6c57611a6c6119e6565b6001019392505050565b60008251611a888184602087016118d6565b9190910192915050565b600060208284031215611aa457600080fd5b815161161e8161163b565b606081526000611ac260608301866118fa565b6020830194909452506001600160a01b0391909116604090910152919050565b604081526000611af560408301856118fa565b905060018060a01b03831660208301529392505050565b6001600160a01b038281168282160390808211156104a7576104a76119e6565b606081526000611b3f60608301866118fa565b6001600160a01b039485166020840152929093166040909101529291505056fe0000000000000000000000007109709ecfa91a80626ff3989d68f67f5b1dd12d53657474696e672070726f7869657320257320666f72206e616d65737061636520257353657474696e672070726f787920257320696d706c656d656e746174696f6e3a202573a2646970667358221220ebda2e9fd0b280f891318ae4ec24034cdcb58f13971ecb1249a342342a316ae764736f6c63430008180033",
}

// AllocPredeploysABI is the input ABI used to generate the binding from.
// Deprecated: Use AllocPredeploysMetaData.ABI instead.
var AllocPredeploysABI = AllocPredeploysMetaData.ABI

// AllocPredeploysBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use AllocPredeploysMetaData.Bin instead.
var AllocPredeploysBin = AllocPredeploysMetaData.Bin

// DeployAllocPredeploys deploys a new Ethereum contract, binding an instance of AllocPredeploys to it.
func DeployAllocPredeploys(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *AllocPredeploys, error) {
	parsed, err := AllocPredeploysMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(AllocPredeploysBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &AllocPredeploys{AllocPredeploysCaller: AllocPredeploysCaller{contract: contract}, AllocPredeploysTransactor: AllocPredeploysTransactor{contract: contract}, AllocPredeploysFilterer: AllocPredeploysFilterer{contract: contract}}, nil
}

// AllocPredeploys is an auto generated Go binding around an Ethereum contract.
type AllocPredeploys struct {
	AllocPredeploysCaller     // Read-only binding to the contract
	AllocPredeploysTransactor // Write-only binding to the contract
	AllocPredeploysFilterer   // Log filterer for contract events
}

// AllocPredeploysCaller is an auto generated read-only Go binding around an Ethereum contract.
type AllocPredeploysCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AllocPredeploysTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AllocPredeploysTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AllocPredeploysFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AllocPredeploysFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AllocPredeploysSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AllocPredeploysSession struct {
	Contract     *AllocPredeploys  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AllocPredeploysCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AllocPredeploysCallerSession struct {
	Contract *AllocPredeploysCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// AllocPredeploysTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AllocPredeploysTransactorSession struct {
	Contract     *AllocPredeploysTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// AllocPredeploysRaw is an auto generated low-level Go binding around an Ethereum contract.
type AllocPredeploysRaw struct {
	Contract *AllocPredeploys // Generic contract binding to access the raw methods on
}

// AllocPredeploysCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AllocPredeploysCallerRaw struct {
	Contract *AllocPredeploysCaller // Generic read-only contract binding to access the raw methods on
}

// AllocPredeploysTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AllocPredeploysTransactorRaw struct {
	Contract *AllocPredeploysTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAllocPredeploys creates a new instance of AllocPredeploys, bound to a specific deployed contract.
func NewAllocPredeploys(address common.Address, backend bind.ContractBackend) (*AllocPredeploys, error) {
	contract, err := bindAllocPredeploys(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AllocPredeploys{AllocPredeploysCaller: AllocPredeploysCaller{contract: contract}, AllocPredeploysTransactor: AllocPredeploysTransactor{contract: contract}, AllocPredeploysFilterer: AllocPredeploysFilterer{contract: contract}}, nil
}

// NewAllocPredeploysCaller creates a new read-only instance of AllocPredeploys, bound to a specific deployed contract.
func NewAllocPredeploysCaller(address common.Address, caller bind.ContractCaller) (*AllocPredeploysCaller, error) {
	contract, err := bindAllocPredeploys(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AllocPredeploysCaller{contract: contract}, nil
}

// NewAllocPredeploysTransactor creates a new write-only instance of AllocPredeploys, bound to a specific deployed contract.
func NewAllocPredeploysTransactor(address common.Address, transactor bind.ContractTransactor) (*AllocPredeploysTransactor, error) {
	contract, err := bindAllocPredeploys(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AllocPredeploysTransactor{contract: contract}, nil
}

// NewAllocPredeploysFilterer creates a new log filterer instance of AllocPredeploys, bound to a specific deployed contract.
func NewAllocPredeploysFilterer(address common.Address, filterer bind.ContractFilterer) (*AllocPredeploysFilterer, error) {
	contract, err := bindAllocPredeploys(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AllocPredeploysFilterer{contract: contract}, nil
}

// bindAllocPredeploys binds a generic wrapper to an already deployed contract.
func bindAllocPredeploys(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AllocPredeploysMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AllocPredeploys *AllocPredeploysRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AllocPredeploys.Contract.AllocPredeploysCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AllocPredeploys *AllocPredeploysRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AllocPredeploys.Contract.AllocPredeploysTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AllocPredeploys *AllocPredeploysRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AllocPredeploys.Contract.AllocPredeploysTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AllocPredeploys *AllocPredeploysCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AllocPredeploys.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AllocPredeploys *AllocPredeploysTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AllocPredeploys.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AllocPredeploys *AllocPredeploysTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AllocPredeploys.Contract.contract.Transact(opts, method, params...)
}

// ISSCRIPT is a free data retrieval call binding the contract method 0xf8ccbf47.
//
// Solidity: function IS_SCRIPT() view returns(bool)
func (_AllocPredeploys *AllocPredeploysCaller) ISSCRIPT(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _AllocPredeploys.contract.Call(opts, &out, "IS_SCRIPT")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// ISSCRIPT is a free data retrieval call binding the contract method 0xf8ccbf47.
//
// Solidity: function IS_SCRIPT() view returns(bool)
func (_AllocPredeploys *AllocPredeploysSession) ISSCRIPT() (bool, error) {
	return _AllocPredeploys.Contract.ISSCRIPT(&_AllocPredeploys.CallOpts)
}

// ISSCRIPT is a free data retrieval call binding the contract method 0xf8ccbf47.
//
// Solidity: function IS_SCRIPT() view returns(bool)
func (_AllocPredeploys *AllocPredeploysCallerSession) ISSCRIPT() (bool, error) {
	return _AllocPredeploys.Contract.ISSCRIPT(&_AllocPredeploys.CallOpts)
}

// RunWithCfg is a paid mutator transaction binding the contract method 0xdb5d1666.
//
// Solidity: function runWithCfg((address,bool,string) config) returns()
func (_AllocPredeploys *AllocPredeploysTransactor) RunWithCfg(opts *bind.TransactOpts, config AllocPredeploysConfig) (*types.Transaction, error) {
	return _AllocPredeploys.contract.Transact(opts, "runWithCfg", config)
}

// RunWithCfg is a paid mutator transaction binding the contract method 0xdb5d1666.
//
// Solidity: function runWithCfg((address,bool,string) config) returns()
func (_AllocPredeploys *AllocPredeploysSession) RunWithCfg(config AllocPredeploysConfig) (*types.Transaction, error) {
	return _AllocPredeploys.Contract.RunWithCfg(&_AllocPredeploys.TransactOpts, config)
}

// RunWithCfg is a paid mutator transaction binding the contract method 0xdb5d1666.
//
// Solidity: function runWithCfg((address,bool,string) config) returns()
func (_AllocPredeploys *AllocPredeploysTransactorSession) RunWithCfg(config AllocPredeploysConfig) (*types.Transaction, error) {
	return _AllocPredeploys.Contract.RunWithCfg(&_AllocPredeploys.TransactOpts, config)
}

// SetProxies is a paid mutator transaction binding the contract method 0x4a05b3f2.
//
// Solidity: function setProxies() returns()
func (_AllocPredeploys *AllocPredeploysTransactor) SetProxies(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AllocPredeploys.contract.Transact(opts, "setProxies")
}

// SetProxies is a paid mutator transaction binding the contract method 0x4a05b3f2.
//
// Solidity: function setProxies() returns()
func (_AllocPredeploys *AllocPredeploysSession) SetProxies() (*types.Transaction, error) {
	return _AllocPredeploys.Contract.SetProxies(&_AllocPredeploys.TransactOpts)
}

// SetProxies is a paid mutator transaction binding the contract method 0x4a05b3f2.
//
// Solidity: function setProxies() returns()
func (_AllocPredeploys *AllocPredeploysTransactorSession) SetProxies() (*types.Transaction, error) {
	return _AllocPredeploys.Contract.SetProxies(&_AllocPredeploys.TransactOpts)
}

// SetUp is a paid mutator transaction binding the contract method 0x0a9254e4.
//
// Solidity: function setUp() returns()
func (_AllocPredeploys *AllocPredeploysTransactor) SetUp(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AllocPredeploys.contract.Transact(opts, "setUp")
}

// SetUp is a paid mutator transaction binding the contract method 0x0a9254e4.
//
// Solidity: function setUp() returns()
func (_AllocPredeploys *AllocPredeploysSession) SetUp() (*types.Transaction, error) {
	return _AllocPredeploys.Contract.SetUp(&_AllocPredeploys.TransactOpts)
}

// SetUp is a paid mutator transaction binding the contract method 0x0a9254e4.
//
// Solidity: function setUp() returns()
func (_AllocPredeploys *AllocPredeploysTransactorSession) SetUp() (*types.Transaction, error) {
	return _AllocPredeploys.Contract.SetUp(&_AllocPredeploys.TransactOpts)
}
