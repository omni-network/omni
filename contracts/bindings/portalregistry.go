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

// PortalRegistryDeployment is an auto generated low-level Go binding around an user-defined struct.
type PortalRegistryDeployment struct {
	Addr           common.Address
	ChainId        uint64
	DeployHeight   uint64
	AttestInterval uint64
	BlockPeriod    uint64
	Shards         []uint64
	Name           string
}

// PortalRegistryMetaData contains all meta data concerning the PortalRegistry contract.
var PortalRegistryMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"bulkRegister\",\"inputs\":[{\"name\":\"deps\",\"type\":\"tuple[]\",\"internalType\":\"structPortalRegistry.Deployment[]\",\"components\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"chainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"deployHeight\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"attestInterval\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"blockPeriod\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"shards\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"},{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"}]}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"chainIds\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"deployments\",\"inputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"chainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"deployHeight\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"attestInterval\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"blockPeriod\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"get\",\"inputs\":[{\"name\":\"chainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structPortalRegistry.Deployment\",\"components\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"chainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"deployHeight\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"attestInterval\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"blockPeriod\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"shards\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"},{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"owner_\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"list\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structPortalRegistry.Deployment[]\",\"components\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"chainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"deployHeight\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"attestInterval\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"blockPeriod\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"shards\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"},{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"register\",\"inputs\":[{\"name\":\"dep\",\"type\":\"tuple\",\"internalType\":\"structPortalRegistry.Deployment\",\"components\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"chainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"deployHeight\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"attestInterval\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"blockPeriod\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"shards\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"},{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PortalRegistered\",\"inputs\":[{\"name\":\"chainId\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"addr\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"deployHeight\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"attestInterval\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"blockPeriod\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"shards\",\"type\":\"uint64[]\",\"indexed\":false,\"internalType\":\"uint64[]\"},{\"name\":\"name\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnableInvalidOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnableUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]}]",
	Bin: "0x608060405234801561001057600080fd5b5061190b806100206000396000f3fe6080604052600436106100915760003560e01c8063715018a611610059578063715018a6146101605780638da5cb5b14610175578063ada86798146101bc578063c4d66de8146101e9578063f2fde38b1461020957600080fd5b80630f560cd71461009657806321d93090146100c157806347153cbf146100f9578063473d04521461011b57806352d482e21461014d575b600080fd5b3480156100a257600080fd5b506100ab610229565b6040516100b8919061114e565b60405180910390f35b3480156100cd57600080fd5b506100e16100dc3660046111b2565b6104ca565b6040516001600160401b0390911681526020016100b8565b34801561010557600080fd5b506101196101143660046111cb565b610507565b005b34801561012757600080fd5b5061013b610136366004611221565b61051b565b6040516100b89695949392919061123e565b61011961015b366004611289565b6105ff565b34801561016c57600080fd5b50610119610667565b34801561018157600080fd5b507f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c199300546040516001600160a01b0390911681526020016100b8565b3480156101c857600080fd5b506101dc6101d7366004611221565b61067b565b6040516100b891906112fd565b3480156101f557600080fd5b50610119610204366004611325565b610840565b34801561021557600080fd5b50610119610224366004611325565b61094e565b60008054606091906001600160401b0381111561024857610248611342565b6040519080825280602002602001820160405280156102b157816020015b6040805160e081018252600080825260208083018290529282018190526060808301829052608083019190915260a0820181905260c082015282526000199092019101816102665790505b50905060005b6000546001600160401b03821610156104c4576001600080836001600160401b0316815481106102e9576102e9611358565b6000918252602080832060048304015460039092166008026101000a9091046001600160401b039081168452838201949094526040928301909120825160e08101845281546001600160a01b0381168252600160a01b9004851681840152600182015480861682860152600160401b810486166060830152600160801b90049094166080850152600281018054845181850281018501909552808552919360a08601939092908301828280156103f057602002820191906000526020600020906000905b82829054906101000a90046001600160401b03166001600160401b0316815260200190600801906020826007010492830192600103820291508084116103ad5790505b505050505081526020016003820180546104099061136e565b80601f01602080910402602001604051908101604052809291908181526020018280546104359061136e565b80156104825780601f1061045757610100808354040283529160200191610482565b820191906000526020600020905b81548152906001019060200180831161046557829003601f168201915b50505050508152505082826001600160401b0316815181106104a6576104a6611358565b602002602001018190525080806104bc906113a2565b9150506102b7565b50919050565b600081815481106104da57600080fd5b9060005260206000209060049182820401919006600802915054906101000a90046001600160401b031681565b61050f61098e565b610518816109e9565b50565b60016020819052600091825260409091208054918101546003820180546001600160a01b038516946001600160401b03600160a01b90910481169484821694600160401b8104831694600160801b9091049092169290919061057c9061136e565b80601f01602080910402602001604051908101604052809291908181526020018280546105a89061136e565b80156105f55780601f106105ca576101008083540402835291602001916105f5565b820191906000526020600020905b8154815290600101906020018083116105d857829003601f168201915b5050505050905086565b61060761098e565b60005b6001600160401b038116821115610662576106508383836001600160401b031681811061063957610639611358565b905060200281019061064b91906113d6565b6109e9565b8061065a816113a2565b91505061060a565b505050565b61066f61098e565b6106796000610f5e565b565b6040805160e0810182526000808252602082018190529181018290526060808201839052608082019290925260a0810182905260c08101919091526001600160401b03808316600090815260016020818152604092839020835160e08101855281546001600160a01b0381168252600160a01b90048616818401529281015480861684860152600160401b810486166060850152600160801b9004909416608083015260028401805484518184028101840190955280855292949360a0860193909283018282801561079e57602002820191906000526020600020906000905b82829054906101000a90046001600160401b03166001600160401b03168152602001906008019060208260070104928301926001038202915080841161075b5790505b505050505081526020016003820180546107b79061136e565b80601f01602080910402602001604051908101604052809291908181526020018280546107e39061136e565b80156108305780601f1061080557610100808354040283529160200191610830565b820191906000526020600020905b81548152906001019060200180831161081357829003601f168201915b5050505050815250509050919050565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a008054600160401b810460ff1615906001600160401b03166000811580156108855750825b90506000826001600160401b031660011480156108a15750303b155b9050811580156108af575080155b156108cd5760405163f92ee8a960e01b815260040160405180910390fd5b845467ffffffffffffffff1916600117855583156108f757845460ff60401b1916600160401b1785555b61090086610fcf565b831561094657845460ff60401b19168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b505050505050565b61095661098e565b6001600160a01b03811661098557604051631e4fbdf760e01b8152600060048201526024015b60405180910390fd5b61051881610f5e565b336109c07f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c199300546001600160a01b031690565b6001600160a01b0316146106795760405163118cdaa760e01b815233600482015260240161097c565b60006109f86020830183611325565b6001600160a01b031603610a4e5760405162461bcd60e51b815260206004820152601960248201527f506f7274616c52656769737472793a207a65726f206164647200000000000000604482015260640161097c565b6000610a606040830160208401611221565b6001600160401b031611610ab65760405162461bcd60e51b815260206004820152601d60248201527f506f7274616c52656769737472793a207a65726f20636861696e204944000000604482015260640161097c565b6000610ac86080830160608401611221565b6001600160401b031611610b1e5760405162461bcd60e51b815260206004820152601d60248201527f506f7274616c52656769737472793a207a65726f20696e74657276616c000000604482015260640161097c565b677fffffffffffffff610b3760a0830160808401611221565b6001600160401b03161115610b8e5760405162461bcd60e51b815260206004820181905260248201527f506f7274616c52656769737472793a20706572696f6420746f6f206c61726765604482015260640161097c565b6000610ba060a0830160808401611221565b6001600160401b031611610bf65760405162461bcd60e51b815260206004820152601b60248201527f506f7274616c52656769737472793a207a65726f20706572696f640000000000604482015260640161097c565b6000610c0560c08301836113f6565b905011610c545760405162461bcd60e51b815260206004820152601760248201527f506f7274616c52656769737472793a206e6f206e616d65000000000000000000604482015260640161097c565b6000610c6360a0830183611443565b905011610cb25760405162461bcd60e51b815260206004820152601960248201527f506f7274616c52656769737472793a206e6f2073686172647300000000000000604482015260640161097c565b6000600181610cc76040850160208601611221565b6001600160401b031681526020810191909152604001600020546001600160a01b031614610d375760405162461bcd60e51b815260206004820152601b60248201527f506f7274616c52656769737472793a20616c7265616479207365740000000000604482015260640161097c565b60005b610d4760a0830183611443565b9050816001600160401b03161015610e16576000610d6860a0840184611443565b836001600160401b0316818110610d8157610d81611358565b9050602002016020810190610d969190611221565b90508060ff16816001600160401b0316148015610db75750610db781610fe0565b610e035760405162461bcd60e51b815260206004820152601d60248201527f506f7274616c52656769737472793a20696e76616c6964207368617264000000604482015260640161097c565b5080610e0e816113a2565b915050610d3a565b508060016000610e2c6040840160208501611221565b6001600160401b031681526020810191909152604001600020610e4f82826116f3565b5060009050610e646040830160208401611221565b815460018101835560009283526020928390206004820401805460039092166008026101000a6001600160401b03818102199093169390921691909102919091179055610eb390820182611325565b6001600160a01b0316610ecc6040830160208401611221565b6001600160401b03167fb08d1911b978b0c040fa5e01711aa326770a97c5f00039d45e7ae8dec7409e73610f066060850160408601611221565b610f166080860160608701611221565b610f2660a0870160808801611221565b610f3360a0880188611443565b610f4060c08a018a6113f6565b604051610f53979695949392919061183c565b60405180910390a350565b7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c19930080546001600160a01b031981166001600160a01b03848116918217845560405192169182907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a3505050565b610fd7610ffd565b61051881611046565b600060ff821660011480610ff7575060ff82166004145b92915050565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a0054600160401b900460ff1661067957604051631afcd79f60e31b815260040160405180910390fd5b610956610ffd565b6000815180845260005b8181101561107457602081850181015186830182015201611058565b506000602082860101526020601f19601f83011685010191505092915050565b600060e0830160018060a01b0383511684526020808401516001600160401b03808216602088015280604087015116604088015280606087015116606088015280608087015116608088015260a0860151915060e060a088015283825180865261010089019150602084019550600093505b8084101561112857855183168252948401946001939093019290840190611106565b5060c0870151945087810360c0890152611142818661104e565b98975050505050505050565b600060208083016020845280855180835260408601915060408160051b87010192506020870160005b828110156111a557603f19888603018452611193858351611094565b94509285019290850190600101611177565b5092979650505050505050565b6000602082840312156111c457600080fd5b5035919050565b6000602082840312156111dd57600080fd5b81356001600160401b038111156111f357600080fd5b820160e0818503121561120557600080fd5b9392505050565b6001600160401b038116811461051857600080fd5b60006020828403121561123357600080fd5b81356112058161120c565b6001600160a01b03871681526001600160401b038681166020830152858116604083015284811660608301528316608082015260c060a082018190526000906111429083018461104e565b6000806020838503121561129c57600080fd5b82356001600160401b03808211156112b357600080fd5b818501915085601f8301126112c757600080fd5b8135818111156112d657600080fd5b8660208260051b85010111156112eb57600080fd5b60209290920196919550909350505050565b6020815260006112056020830184611094565b6001600160a01b038116811461051857600080fd5b60006020828403121561133757600080fd5b813561120581611310565b634e487b7160e01b600052604160045260246000fd5b634e487b7160e01b600052603260045260246000fd5b600181811c9082168061138257607f821691505b6020821081036104c457634e487b7160e01b600052602260045260246000fd5b60006001600160401b038083168181036113cc57634e487b7160e01b600052601160045260246000fd5b6001019392505050565b6000823560de198336030181126113ec57600080fd5b9190910192915050565b6000808335601e1984360301811261140d57600080fd5b8301803591506001600160401b0382111561142757600080fd5b60200191503681900382131561143c57600080fd5b9250929050565b6000808335601e1984360301811261145a57600080fd5b8301803591506001600160401b0382111561147457600080fd5b6020019150600581901b360382131561143c57600080fd5b60008135610ff78161120c565b5b818110156114ae576000815560010161149a565b5050565b600160401b8211156114c6576114c6611342565b805482825580831015610662578160005260206000206003840160021c810160188560031b168015611509576000198083018054828460200360031b1c16815550505b5061151c6003840160021c830182611499565b5050505050565b6001600160401b0383111561153a5761153a611342565b61154483826114b2565b60008181526020902082908460021c60005b818110156115b2576000805b60048110156115a5576115946115778761148c565b6001600160401b03908116600684901b90811b91901b1984161790565b602096909601959150600101611562565b5083820155600101611556565b5060031986168087038188146115f1576000805b828110156115eb576115da6115778861148c565b6020979097019691506001016115c6565b50848401555b5050505050505050565b601f82111561066257806000526020600020601f840160051c810160208510156116225750805b61151c601f850160051c830182611499565b6001600160401b0383111561164b5761164b611342565b61165f83611659835461136e565b836115fb565b6000601f841160018114611693576000851561167b5750838201355b600019600387901b1c1916600186901b17835561151c565b600083815260209020601f19861690835b828110156116c457868501358255602094850194600190920191016116a4565b50868210156116e15760001960f88860031b161c19848701351681555b505060018560011b0183555050505050565b81356116fe81611310565b81546001600160a01b031981166001600160a01b03929092169182178355602084013561172a8161120c565b6001600160e01b03199190911690911760a09190911b67ffffffffffffffff60a01b16178155600181016117816117636040850161148c565b825467ffffffffffffffff19166001600160401b0391909116178255565b6117c26117906060850161148c565b82546fffffffffffffffff0000000000000000191660409190911b6fffffffffffffffff000000000000000016178255565b6117fd6117d16080850161148c565b82805467ffffffffffffffff60801b191660809290921b67ffffffffffffffff60801b16919091179055565b5061180b60a0830183611443565b611819818360028601611523565b505061182860c08301836113f6565b611836818360038601611634565b50505050565b600060a082016001600160401b03808b1684526020818b1681860152818a16604086015260a060608601528288845260c08601905089935060005b8981101561189e57843561188a8161120c565b841682529382019390820190600101611877565b5085810360808701528681528688838301376000818801830152601f909601601f19169095019094019a995050505050505050505056fea2646970667358221220a001b8c39d3ce4e766e6e03e79eab4645f3ecdfbe7d17a38d87ff7a73148936c64736f6c63430008180033",
}

// PortalRegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use PortalRegistryMetaData.ABI instead.
var PortalRegistryABI = PortalRegistryMetaData.ABI

// PortalRegistryBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use PortalRegistryMetaData.Bin instead.
var PortalRegistryBin = PortalRegistryMetaData.Bin

// DeployPortalRegistry deploys a new Ethereum contract, binding an instance of PortalRegistry to it.
func DeployPortalRegistry(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *PortalRegistry, error) {
	parsed, err := PortalRegistryMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(PortalRegistryBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &PortalRegistry{PortalRegistryCaller: PortalRegistryCaller{contract: contract}, PortalRegistryTransactor: PortalRegistryTransactor{contract: contract}, PortalRegistryFilterer: PortalRegistryFilterer{contract: contract}}, nil
}

// PortalRegistry is an auto generated Go binding around an Ethereum contract.
type PortalRegistry struct {
	PortalRegistryCaller     // Read-only binding to the contract
	PortalRegistryTransactor // Write-only binding to the contract
	PortalRegistryFilterer   // Log filterer for contract events
}

// PortalRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type PortalRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PortalRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PortalRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PortalRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PortalRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PortalRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PortalRegistrySession struct {
	Contract     *PortalRegistry   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PortalRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PortalRegistryCallerSession struct {
	Contract *PortalRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// PortalRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PortalRegistryTransactorSession struct {
	Contract     *PortalRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// PortalRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type PortalRegistryRaw struct {
	Contract *PortalRegistry // Generic contract binding to access the raw methods on
}

// PortalRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PortalRegistryCallerRaw struct {
	Contract *PortalRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// PortalRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PortalRegistryTransactorRaw struct {
	Contract *PortalRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPortalRegistry creates a new instance of PortalRegistry, bound to a specific deployed contract.
func NewPortalRegistry(address common.Address, backend bind.ContractBackend) (*PortalRegistry, error) {
	contract, err := bindPortalRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PortalRegistry{PortalRegistryCaller: PortalRegistryCaller{contract: contract}, PortalRegistryTransactor: PortalRegistryTransactor{contract: contract}, PortalRegistryFilterer: PortalRegistryFilterer{contract: contract}}, nil
}

// NewPortalRegistryCaller creates a new read-only instance of PortalRegistry, bound to a specific deployed contract.
func NewPortalRegistryCaller(address common.Address, caller bind.ContractCaller) (*PortalRegistryCaller, error) {
	contract, err := bindPortalRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PortalRegistryCaller{contract: contract}, nil
}

// NewPortalRegistryTransactor creates a new write-only instance of PortalRegistry, bound to a specific deployed contract.
func NewPortalRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*PortalRegistryTransactor, error) {
	contract, err := bindPortalRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PortalRegistryTransactor{contract: contract}, nil
}

// NewPortalRegistryFilterer creates a new log filterer instance of PortalRegistry, bound to a specific deployed contract.
func NewPortalRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*PortalRegistryFilterer, error) {
	contract, err := bindPortalRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PortalRegistryFilterer{contract: contract}, nil
}

// bindPortalRegistry binds a generic wrapper to an already deployed contract.
func bindPortalRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := PortalRegistryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PortalRegistry *PortalRegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PortalRegistry.Contract.PortalRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PortalRegistry *PortalRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PortalRegistry.Contract.PortalRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PortalRegistry *PortalRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PortalRegistry.Contract.PortalRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PortalRegistry *PortalRegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _PortalRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PortalRegistry *PortalRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PortalRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PortalRegistry *PortalRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PortalRegistry.Contract.contract.Transact(opts, method, params...)
}

// ChainIds is a free data retrieval call binding the contract method 0x21d93090.
//
// Solidity: function chainIds(uint256 ) view returns(uint64)
func (_PortalRegistry *PortalRegistryCaller) ChainIds(opts *bind.CallOpts, arg0 *big.Int) (uint64, error) {
	var out []interface{}
	err := _PortalRegistry.contract.Call(opts, &out, "chainIds", arg0)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// ChainIds is a free data retrieval call binding the contract method 0x21d93090.
//
// Solidity: function chainIds(uint256 ) view returns(uint64)
func (_PortalRegistry *PortalRegistrySession) ChainIds(arg0 *big.Int) (uint64, error) {
	return _PortalRegistry.Contract.ChainIds(&_PortalRegistry.CallOpts, arg0)
}

// ChainIds is a free data retrieval call binding the contract method 0x21d93090.
//
// Solidity: function chainIds(uint256 ) view returns(uint64)
func (_PortalRegistry *PortalRegistryCallerSession) ChainIds(arg0 *big.Int) (uint64, error) {
	return _PortalRegistry.Contract.ChainIds(&_PortalRegistry.CallOpts, arg0)
}

// Deployments is a free data retrieval call binding the contract method 0x473d0452.
//
// Solidity: function deployments(uint64 ) view returns(address addr, uint64 chainId, uint64 deployHeight, uint64 attestInterval, uint64 blockPeriod, string name)
func (_PortalRegistry *PortalRegistryCaller) Deployments(opts *bind.CallOpts, arg0 uint64) (struct {
	Addr           common.Address
	ChainId        uint64
	DeployHeight   uint64
	AttestInterval uint64
	BlockPeriod    uint64
	Name           string
}, error) {
	var out []interface{}
	err := _PortalRegistry.contract.Call(opts, &out, "deployments", arg0)

	outstruct := new(struct {
		Addr           common.Address
		ChainId        uint64
		DeployHeight   uint64
		AttestInterval uint64
		BlockPeriod    uint64
		Name           string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Addr = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.ChainId = *abi.ConvertType(out[1], new(uint64)).(*uint64)
	outstruct.DeployHeight = *abi.ConvertType(out[2], new(uint64)).(*uint64)
	outstruct.AttestInterval = *abi.ConvertType(out[3], new(uint64)).(*uint64)
	outstruct.BlockPeriod = *abi.ConvertType(out[4], new(uint64)).(*uint64)
	outstruct.Name = *abi.ConvertType(out[5], new(string)).(*string)

	return *outstruct, err

}

// Deployments is a free data retrieval call binding the contract method 0x473d0452.
//
// Solidity: function deployments(uint64 ) view returns(address addr, uint64 chainId, uint64 deployHeight, uint64 attestInterval, uint64 blockPeriod, string name)
func (_PortalRegistry *PortalRegistrySession) Deployments(arg0 uint64) (struct {
	Addr           common.Address
	ChainId        uint64
	DeployHeight   uint64
	AttestInterval uint64
	BlockPeriod    uint64
	Name           string
}, error) {
	return _PortalRegistry.Contract.Deployments(&_PortalRegistry.CallOpts, arg0)
}

// Deployments is a free data retrieval call binding the contract method 0x473d0452.
//
// Solidity: function deployments(uint64 ) view returns(address addr, uint64 chainId, uint64 deployHeight, uint64 attestInterval, uint64 blockPeriod, string name)
func (_PortalRegistry *PortalRegistryCallerSession) Deployments(arg0 uint64) (struct {
	Addr           common.Address
	ChainId        uint64
	DeployHeight   uint64
	AttestInterval uint64
	BlockPeriod    uint64
	Name           string
}, error) {
	return _PortalRegistry.Contract.Deployments(&_PortalRegistry.CallOpts, arg0)
}

// Get is a free data retrieval call binding the contract method 0xada86798.
//
// Solidity: function get(uint64 chainId) view returns((address,uint64,uint64,uint64,uint64,uint64[],string))
func (_PortalRegistry *PortalRegistryCaller) Get(opts *bind.CallOpts, chainId uint64) (PortalRegistryDeployment, error) {
	var out []interface{}
	err := _PortalRegistry.contract.Call(opts, &out, "get", chainId)

	if err != nil {
		return *new(PortalRegistryDeployment), err
	}

	out0 := *abi.ConvertType(out[0], new(PortalRegistryDeployment)).(*PortalRegistryDeployment)

	return out0, err

}

// Get is a free data retrieval call binding the contract method 0xada86798.
//
// Solidity: function get(uint64 chainId) view returns((address,uint64,uint64,uint64,uint64,uint64[],string))
func (_PortalRegistry *PortalRegistrySession) Get(chainId uint64) (PortalRegistryDeployment, error) {
	return _PortalRegistry.Contract.Get(&_PortalRegistry.CallOpts, chainId)
}

// Get is a free data retrieval call binding the contract method 0xada86798.
//
// Solidity: function get(uint64 chainId) view returns((address,uint64,uint64,uint64,uint64,uint64[],string))
func (_PortalRegistry *PortalRegistryCallerSession) Get(chainId uint64) (PortalRegistryDeployment, error) {
	return _PortalRegistry.Contract.Get(&_PortalRegistry.CallOpts, chainId)
}

// List is a free data retrieval call binding the contract method 0x0f560cd7.
//
// Solidity: function list() view returns((address,uint64,uint64,uint64,uint64,uint64[],string)[])
func (_PortalRegistry *PortalRegistryCaller) List(opts *bind.CallOpts) ([]PortalRegistryDeployment, error) {
	var out []interface{}
	err := _PortalRegistry.contract.Call(opts, &out, "list")

	if err != nil {
		return *new([]PortalRegistryDeployment), err
	}

	out0 := *abi.ConvertType(out[0], new([]PortalRegistryDeployment)).(*[]PortalRegistryDeployment)

	return out0, err

}

// List is a free data retrieval call binding the contract method 0x0f560cd7.
//
// Solidity: function list() view returns((address,uint64,uint64,uint64,uint64,uint64[],string)[])
func (_PortalRegistry *PortalRegistrySession) List() ([]PortalRegistryDeployment, error) {
	return _PortalRegistry.Contract.List(&_PortalRegistry.CallOpts)
}

// List is a free data retrieval call binding the contract method 0x0f560cd7.
//
// Solidity: function list() view returns((address,uint64,uint64,uint64,uint64,uint64[],string)[])
func (_PortalRegistry *PortalRegistryCallerSession) List() ([]PortalRegistryDeployment, error) {
	return _PortalRegistry.Contract.List(&_PortalRegistry.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PortalRegistry *PortalRegistryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _PortalRegistry.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PortalRegistry *PortalRegistrySession) Owner() (common.Address, error) {
	return _PortalRegistry.Contract.Owner(&_PortalRegistry.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_PortalRegistry *PortalRegistryCallerSession) Owner() (common.Address, error) {
	return _PortalRegistry.Contract.Owner(&_PortalRegistry.CallOpts)
}

// BulkRegister is a paid mutator transaction binding the contract method 0x52d482e2.
//
// Solidity: function bulkRegister((address,uint64,uint64,uint64,uint64,uint64[],string)[] deps) payable returns()
func (_PortalRegistry *PortalRegistryTransactor) BulkRegister(opts *bind.TransactOpts, deps []PortalRegistryDeployment) (*types.Transaction, error) {
	return _PortalRegistry.contract.Transact(opts, "bulkRegister", deps)
}

// BulkRegister is a paid mutator transaction binding the contract method 0x52d482e2.
//
// Solidity: function bulkRegister((address,uint64,uint64,uint64,uint64,uint64[],string)[] deps) payable returns()
func (_PortalRegistry *PortalRegistrySession) BulkRegister(deps []PortalRegistryDeployment) (*types.Transaction, error) {
	return _PortalRegistry.Contract.BulkRegister(&_PortalRegistry.TransactOpts, deps)
}

// BulkRegister is a paid mutator transaction binding the contract method 0x52d482e2.
//
// Solidity: function bulkRegister((address,uint64,uint64,uint64,uint64,uint64[],string)[] deps) payable returns()
func (_PortalRegistry *PortalRegistryTransactorSession) BulkRegister(deps []PortalRegistryDeployment) (*types.Transaction, error) {
	return _PortalRegistry.Contract.BulkRegister(&_PortalRegistry.TransactOpts, deps)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address owner_) returns()
func (_PortalRegistry *PortalRegistryTransactor) Initialize(opts *bind.TransactOpts, owner_ common.Address) (*types.Transaction, error) {
	return _PortalRegistry.contract.Transact(opts, "initialize", owner_)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address owner_) returns()
func (_PortalRegistry *PortalRegistrySession) Initialize(owner_ common.Address) (*types.Transaction, error) {
	return _PortalRegistry.Contract.Initialize(&_PortalRegistry.TransactOpts, owner_)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address owner_) returns()
func (_PortalRegistry *PortalRegistryTransactorSession) Initialize(owner_ common.Address) (*types.Transaction, error) {
	return _PortalRegistry.Contract.Initialize(&_PortalRegistry.TransactOpts, owner_)
}

// Register is a paid mutator transaction binding the contract method 0x47153cbf.
//
// Solidity: function register((address,uint64,uint64,uint64,uint64,uint64[],string) dep) returns()
func (_PortalRegistry *PortalRegistryTransactor) Register(opts *bind.TransactOpts, dep PortalRegistryDeployment) (*types.Transaction, error) {
	return _PortalRegistry.contract.Transact(opts, "register", dep)
}

// Register is a paid mutator transaction binding the contract method 0x47153cbf.
//
// Solidity: function register((address,uint64,uint64,uint64,uint64,uint64[],string) dep) returns()
func (_PortalRegistry *PortalRegistrySession) Register(dep PortalRegistryDeployment) (*types.Transaction, error) {
	return _PortalRegistry.Contract.Register(&_PortalRegistry.TransactOpts, dep)
}

// Register is a paid mutator transaction binding the contract method 0x47153cbf.
//
// Solidity: function register((address,uint64,uint64,uint64,uint64,uint64[],string) dep) returns()
func (_PortalRegistry *PortalRegistryTransactorSession) Register(dep PortalRegistryDeployment) (*types.Transaction, error) {
	return _PortalRegistry.Contract.Register(&_PortalRegistry.TransactOpts, dep)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PortalRegistry *PortalRegistryTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PortalRegistry.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PortalRegistry *PortalRegistrySession) RenounceOwnership() (*types.Transaction, error) {
	return _PortalRegistry.Contract.RenounceOwnership(&_PortalRegistry.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_PortalRegistry *PortalRegistryTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _PortalRegistry.Contract.RenounceOwnership(&_PortalRegistry.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_PortalRegistry *PortalRegistryTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _PortalRegistry.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_PortalRegistry *PortalRegistrySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _PortalRegistry.Contract.TransferOwnership(&_PortalRegistry.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_PortalRegistry *PortalRegistryTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _PortalRegistry.Contract.TransferOwnership(&_PortalRegistry.TransactOpts, newOwner)
}

// PortalRegistryInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the PortalRegistry contract.
type PortalRegistryInitializedIterator struct {
	Event *PortalRegistryInitialized // Event containing the contract specifics and raw log

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
func (it *PortalRegistryInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PortalRegistryInitialized)
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
		it.Event = new(PortalRegistryInitialized)
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
func (it *PortalRegistryInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PortalRegistryInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PortalRegistryInitialized represents a Initialized event raised by the PortalRegistry contract.
type PortalRegistryInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_PortalRegistry *PortalRegistryFilterer) FilterInitialized(opts *bind.FilterOpts) (*PortalRegistryInitializedIterator, error) {

	logs, sub, err := _PortalRegistry.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &PortalRegistryInitializedIterator{contract: _PortalRegistry.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_PortalRegistry *PortalRegistryFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *PortalRegistryInitialized) (event.Subscription, error) {

	logs, sub, err := _PortalRegistry.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PortalRegistryInitialized)
				if err := _PortalRegistry.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_PortalRegistry *PortalRegistryFilterer) ParseInitialized(log types.Log) (*PortalRegistryInitialized, error) {
	event := new(PortalRegistryInitialized)
	if err := _PortalRegistry.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PortalRegistryOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the PortalRegistry contract.
type PortalRegistryOwnershipTransferredIterator struct {
	Event *PortalRegistryOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *PortalRegistryOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PortalRegistryOwnershipTransferred)
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
		it.Event = new(PortalRegistryOwnershipTransferred)
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
func (it *PortalRegistryOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PortalRegistryOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PortalRegistryOwnershipTransferred represents a OwnershipTransferred event raised by the PortalRegistry contract.
type PortalRegistryOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_PortalRegistry *PortalRegistryFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*PortalRegistryOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _PortalRegistry.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &PortalRegistryOwnershipTransferredIterator{contract: _PortalRegistry.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_PortalRegistry *PortalRegistryFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *PortalRegistryOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _PortalRegistry.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PortalRegistryOwnershipTransferred)
				if err := _PortalRegistry.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_PortalRegistry *PortalRegistryFilterer) ParseOwnershipTransferred(log types.Log) (*PortalRegistryOwnershipTransferred, error) {
	event := new(PortalRegistryOwnershipTransferred)
	if err := _PortalRegistry.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// PortalRegistryPortalRegisteredIterator is returned from FilterPortalRegistered and is used to iterate over the raw logs and unpacked data for PortalRegistered events raised by the PortalRegistry contract.
type PortalRegistryPortalRegisteredIterator struct {
	Event *PortalRegistryPortalRegistered // Event containing the contract specifics and raw log

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
func (it *PortalRegistryPortalRegisteredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PortalRegistryPortalRegistered)
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
		it.Event = new(PortalRegistryPortalRegistered)
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
func (it *PortalRegistryPortalRegisteredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PortalRegistryPortalRegisteredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PortalRegistryPortalRegistered represents a PortalRegistered event raised by the PortalRegistry contract.
type PortalRegistryPortalRegistered struct {
	ChainId        uint64
	Addr           common.Address
	DeployHeight   uint64
	AttestInterval uint64
	BlockPeriod    uint64
	Shards         []uint64
	Name           string
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterPortalRegistered is a free log retrieval operation binding the contract event 0xb08d1911b978b0c040fa5e01711aa326770a97c5f00039d45e7ae8dec7409e73.
//
// Solidity: event PortalRegistered(uint64 indexed chainId, address indexed addr, uint64 deployHeight, uint64 attestInterval, uint64 blockPeriod, uint64[] shards, string name)
func (_PortalRegistry *PortalRegistryFilterer) FilterPortalRegistered(opts *bind.FilterOpts, chainId []uint64, addr []common.Address) (*PortalRegistryPortalRegisteredIterator, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}
	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}

	logs, sub, err := _PortalRegistry.contract.FilterLogs(opts, "PortalRegistered", chainIdRule, addrRule)
	if err != nil {
		return nil, err
	}
	return &PortalRegistryPortalRegisteredIterator{contract: _PortalRegistry.contract, event: "PortalRegistered", logs: logs, sub: sub}, nil
}

// WatchPortalRegistered is a free log subscription operation binding the contract event 0xb08d1911b978b0c040fa5e01711aa326770a97c5f00039d45e7ae8dec7409e73.
//
// Solidity: event PortalRegistered(uint64 indexed chainId, address indexed addr, uint64 deployHeight, uint64 attestInterval, uint64 blockPeriod, uint64[] shards, string name)
func (_PortalRegistry *PortalRegistryFilterer) WatchPortalRegistered(opts *bind.WatchOpts, sink chan<- *PortalRegistryPortalRegistered, chainId []uint64, addr []common.Address) (event.Subscription, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}
	var addrRule []interface{}
	for _, addrItem := range addr {
		addrRule = append(addrRule, addrItem)
	}

	logs, sub, err := _PortalRegistry.contract.WatchLogs(opts, "PortalRegistered", chainIdRule, addrRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PortalRegistryPortalRegistered)
				if err := _PortalRegistry.contract.UnpackLog(event, "PortalRegistered", log); err != nil {
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

// ParsePortalRegistered is a log parse operation binding the contract event 0xb08d1911b978b0c040fa5e01711aa326770a97c5f00039d45e7ae8dec7409e73.
//
// Solidity: event PortalRegistered(uint64 indexed chainId, address indexed addr, uint64 deployHeight, uint64 attestInterval, uint64 blockPeriod, uint64[] shards, string name)
func (_PortalRegistry *PortalRegistryFilterer) ParsePortalRegistered(log types.Log) (*PortalRegistryPortalRegistered, error) {
	event := new(PortalRegistryPortalRegistered)
	if err := _PortalRegistry.contract.UnpackLog(event, "PortalRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
