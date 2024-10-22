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
	BlockPeriodNs  uint64
	Shards         []uint64
	Name           string
}

// PortalRegistryMetaData contains all meta data concerning the PortalRegistry contract.
var PortalRegistryMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"bulkRegister\",\"inputs\":[{\"name\":\"deps\",\"type\":\"tuple[]\",\"internalType\":\"structPortalRegistry.Deployment[]\",\"components\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"chainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"deployHeight\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"attestInterval\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"blockPeriodNs\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"shards\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"},{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"chainIds\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"deployments\",\"inputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"chainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"deployHeight\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"attestInterval\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"blockPeriodNs\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"get\",\"inputs\":[{\"name\":\"chainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structPortalRegistry.Deployment\",\"components\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"chainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"deployHeight\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"attestInterval\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"blockPeriodNs\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"shards\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"},{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"owner_\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"list\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structPortalRegistry.Deployment[]\",\"components\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"chainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"deployHeight\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"attestInterval\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"blockPeriodNs\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"shards\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"},{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"register\",\"inputs\":[{\"name\":\"dep\",\"type\":\"tuple\",\"internalType\":\"structPortalRegistry.Deployment\",\"components\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"chainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"deployHeight\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"attestInterval\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"blockPeriodNs\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"shards\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"},{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PortalRegistered\",\"inputs\":[{\"name\":\"chainId\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"addr\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"deployHeight\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"attestInterval\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"blockPeriodNs\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"shards\",\"type\":\"uint64[]\",\"indexed\":false,\"internalType\":\"uint64[]\"},{\"name\":\"name\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnableInvalidOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnableUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]}]",
	Bin: "0x608060405234801561000f575f80fd5b5061001861001d565b6100cf565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000900460ff161561006d5760405163f92ee8a960e01b815260040160405180910390fd5b80546001600160401b03908116146100cc5780546001600160401b0319166001600160401b0390811782556040519081527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b50565b61182e806100dc5f395ff3fe608060405234801561000f575f80fd5b506004361061009b575f3560e01c8063715018a611610063578063715018a6146101355780638da5cb5b1461013d578063ada8679814610177578063c4d66de814610197578063f2fde38b146101aa575f80fd5b80630f560cd71461009f57806321d93090146100bd57806347153cbf146100e8578063473d0452146100fd57806352d482e214610122575b5f80fd5b6100a76101bd565b6040516100b491906110ae565b60405180910390f35b6100d06100cb366004611110565b610452565b6040516001600160401b0390911681526020016100b4565b6100fb6100f6366004611127565b61048b565b005b61011061010b366004611178565b61049f565b6040516100b496959493929190611193565b6100fb6101303660046111dd565b610580565b6100fb6105e7565b7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c199300546040516001600160a01b0390911681526020016100b4565b61018a610185366004611178565b6105fa565b6040516100b4919061124b565b6100fb6101a5366004611271565b6107b8565b6100fb6101b8366004611271565b6108c4565b5f8054606091906001600160401b038111156101db576101db61128c565b60405190808252806020026020018201604052801561024257816020015b6040805160e0810182525f80825260208083018290529282018190526060808301829052608083019190915260a0820181905260c082015282525f199092019101816101f95790505b5090505f5b5f546001600160401b038216101561044c5760015f80836001600160401b031681548110610277576102776112a0565b5f918252602080832060048304015460039092166008026101000a9091046001600160401b039081168452838201949094526040928301909120825160e08101845281546001600160a01b0381168252600160a01b9004851681840152600182015480861682860152600160401b810486166060830152600160801b90049094166080850152600281018054845181850281018501909552808552919360a086019390929083018282801561037a57602002820191905f5260205f20905f905b82829054906101000a90046001600160401b03166001600160401b0316815260200190600801906020826007010492830192600103820291508084116103375790505b50505050508152602001600382018054610393906112b4565b80601f01602080910402602001604051908101604052809291908181526020018280546103bf906112b4565b801561040a5780601f106103e15761010080835404028352916020019161040a565b820191905f5260205f20905b8154815290600101906020018083116103ed57829003601f168201915b50505050508152505082826001600160401b03168151811061042e5761042e6112a0565b60200260200101819052508080610444906112e6565b915050610247565b50919050565b5f8181548110610460575f80fd5b905f5260205f209060049182820401919006600802915054906101000a90046001600160401b031681565b610493610903565b61049c8161095e565b50565b600160208190525f91825260409091208054918101546003820180546001600160a01b038516946001600160401b03600160a01b90910481169484821694600160401b8104831694600160801b909104909216929091906104ff906112b4565b80601f016020809104026020016040519081016040528092919081815260200182805461052b906112b4565b80156105765780601f1061054d57610100808354040283529160200191610576565b820191905f5260205f20905b81548152906001019060200180831161055957829003601f168201915b5050505050905086565b610588610903565b5f5b6001600160401b0381168211156105e2576105d08383836001600160401b03168181106105b9576105b96112a0565b90506020028101906105cb9190611317565b61095e565b806105da816112e6565b91505061058a565b505050565b6105ef610903565b6105f85f610ec5565b565b6040805160e0810182525f808252602082018190529181018290526060808201839052608082019290925260a0810182905260c08101919091526001600160401b038083165f90815260016020818152604092839020835160e08101855281546001600160a01b0381168252600160a01b90048616818401529281015480861684860152600160401b810486166060850152600160801b9004909416608083015260028401805484518184028101840190955280855292949360a0860193909283018282801561071857602002820191905f5260205f20905f905b82829054906101000a90046001600160401b03166001600160401b0316815260200190600801906020826007010492830192600103820291508084116106d55790505b50505050508152602001600382018054610731906112b4565b80601f016020809104026020016040519081016040528092919081815260200182805461075d906112b4565b80156107a85780601f1061077f576101008083540402835291602001916107a8565b820191905f5260205f20905b81548152906001019060200180831161078b57829003601f168201915b5050505050815250509050919050565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a008054600160401b810460ff1615906001600160401b03165f811580156107fc5750825b90505f826001600160401b031660011480156108175750303b155b905081158015610825575080155b156108435760405163f92ee8a960e01b815260040160405180910390fd5b845467ffffffffffffffff19166001178555831561086d57845460ff60401b1916600160401b1785555b61087686610f35565b83156108bc57845460ff60401b19168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b505050505050565b6108cc610903565b6001600160a01b0381166108fa57604051631e4fbdf760e01b81525f60048201526024015b60405180910390fd5b61049c81610ec5565b336109357f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c199300546001600160a01b031690565b6001600160a01b0316146105f85760405163118cdaa760e01b81523360048201526024016108f1565b5f61096c6020830183611271565b6001600160a01b0316036109c25760405162461bcd60e51b815260206004820152601960248201527f506f7274616c52656769737472793a207a65726f20616464720000000000000060448201526064016108f1565b5f6109d36040830160208401611178565b6001600160401b031611610a295760405162461bcd60e51b815260206004820152601d60248201527f506f7274616c52656769737472793a207a65726f20636861696e20494400000060448201526064016108f1565b5f610a3a6080830160608401611178565b6001600160401b031611610a905760405162461bcd60e51b815260206004820152601d60248201527f506f7274616c52656769737472793a207a65726f20696e74657276616c00000060448201526064016108f1565b677fffffffffffffff610aa960a0830160808401611178565b6001600160401b03161115610b005760405162461bcd60e51b815260206004820181905260248201527f506f7274616c52656769737472793a20706572696f6420746f6f206c6172676560448201526064016108f1565b5f610b1160a0830160808401611178565b6001600160401b031611610b675760405162461bcd60e51b815260206004820152601b60248201527f506f7274616c52656769737472793a207a65726f20706572696f64000000000060448201526064016108f1565b5f610b7560c0830183611335565b905011610bc45760405162461bcd60e51b815260206004820152601760248201527f506f7274616c52656769737472793a206e6f206e616d6500000000000000000060448201526064016108f1565b5f610bd260a083018361137e565b905011610c215760405162461bcd60e51b815260206004820152601960248201527f506f7274616c52656769737472793a206e6f207368617264730000000000000060448201526064016108f1565b5f600181610c356040850160208601611178565b6001600160401b0316815260208101919091526040015f20546001600160a01b031614610ca45760405162461bcd60e51b815260206004820152601b60248201527f506f7274616c52656769737472793a20616c726561647920736574000000000060448201526064016108f1565b5f5b610cb360a083018361137e565b9050816001600160401b03161015610d81575f610cd360a084018461137e565b836001600160401b0316818110610cec57610cec6112a0565b9050602002016020810190610d019190611178565b90508060ff16816001600160401b0316148015610d225750610d2281610f46565b610d6e5760405162461bcd60e51b815260206004820152601d60248201527f506f7274616c52656769737472793a20696e76616c696420736861726400000060448201526064016108f1565b5080610d79816112e6565b915050610ca6565b508060015f610d966040840160208501611178565b6001600160401b0316815260208101919091526040015f20610db88282611619565b505f9050610dcc6040830160208401611178565b81546001810183555f9283526020928390206004820401805460039092166008026101000a6001600160401b03818102199093169390921691909102919091179055610e1a90820182611271565b6001600160a01b0316610e336040830160208401611178565b6001600160401b03167fb08d1911b978b0c040fa5e01711aa326770a97c5f00039d45e7ae8dec7409e73610e6d6060850160408601611178565b610e7d6080860160608701611178565b610e8d60a0870160808801611178565b610e9a60a088018861137e565b610ea760c08a018a611335565b604051610eba9796959493929190611762565b60405180910390a350565b7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c19930080546001600160a01b031981166001600160a01b03848116918217845560405192169182907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0905f90a3505050565b610f3d610f62565b61049c81610fab565b5f60ff821660011480610f5c575060ff82166004145b92915050565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a0054600160401b900460ff166105f857604051631afcd79f60e31b815260040160405180910390fd5b6108cc610f62565b5f81518084525f5b81811015610fd757602081850181015186830182015201610fbb565b505f602082860101526020601f19601f83011685010191505092915050565b5f60e0830160018060a01b0383511684526020808401516001600160401b03808216602088015280604087015116604088015280606087015116606088015280608087015116608088015260a0860151915060e060a0880152838251808652610100890191506020840195505f93505b8084101561108857855183168252948401946001939093019290840190611066565b5060c0870151945087810360c08901526110a28186610fb3565b98975050505050505050565b5f60208083016020845280855180835260408601915060408160051b8701019250602087015f5b8281101561110357603f198886030184526110f1858351610ff6565b945092850192908501906001016110d5565b5092979650505050505050565b5f60208284031215611120575f80fd5b5035919050565b5f60208284031215611137575f80fd5b81356001600160401b0381111561114c575f80fd5b820160e0818503121561115d575f80fd5b9392505050565b6001600160401b038116811461049c575f80fd5b5f60208284031215611188575f80fd5b813561115d81611164565b6001600160a01b03871681526001600160401b038681166020830152858116604083015284811660608301528316608082015260c060a082018190525f906110a290830184610fb3565b5f80602083850312156111ee575f80fd5b82356001600160401b0380821115611204575f80fd5b818501915085601f830112611217575f80fd5b813581811115611225575f80fd5b8660208260051b8501011115611239575f80fd5b60209290920196919550909350505050565b602081525f61115d6020830184610ff6565b6001600160a01b038116811461049c575f80fd5b5f60208284031215611281575f80fd5b813561115d8161125d565b634e487b7160e01b5f52604160045260245ffd5b634e487b7160e01b5f52603260045260245ffd5b600181811c908216806112c857607f821691505b60208210810361044c57634e487b7160e01b5f52602260045260245ffd5b5f6001600160401b0380831681810361130d57634e487b7160e01b5f52601160045260245ffd5b6001019392505050565b5f823560de1983360301811261132b575f80fd5b9190910192915050565b5f808335601e1984360301811261134a575f80fd5b8301803591506001600160401b03821115611363575f80fd5b602001915036819003821315611377575f80fd5b9250929050565b5f808335601e19843603018112611393575f80fd5b8301803591506001600160401b038211156113ac575f80fd5b6020019150600581901b3603821315611377575f80fd5b5f8135610f5c81611164565b5b818110156113e3575f81556001016113d0565b5050565b600160401b8211156113fb576113fb61128c565b8054828255808310156105e257815f5260205f206003840160021c810160188560031b16801561143b575f198083018054828460200360031b1c16815550505b5061144e6003840160021c8301826113cf565b5050505050565b6001600160401b0383111561146c5761146c61128c565b61147683826113e7565b5f8181526020902082908460021c5f5b818110156114e1575f805b60048110156114d4576114c36114a6876113c3565b6001600160401b03908116600684901b90811b91901b1984161790565b602096909601959150600101611491565b5083820155600101611486565b50600319861680870381881461151f575f805b82811015611519576115086114a6886113c3565b6020979097019691506001016114f4565b50848401555b5050505050505050565b601f8211156105e257805f5260205f20601f840160051c8101602085101561154e5750805b61144e601f850160051c8301826113cf565b6001600160401b038311156115775761157761128c565b61158b8361158583546112b4565b83611529565b5f601f8411600181146115bc575f85156115a55750838201355b5f19600387901b1c1916600186901b17835561144e565b5f83815260208120601f198716915b828110156115eb57868501358255602094850194600190920191016115cb565b5086821015611607575f1960f88860031b161c19848701351681555b505060018560011b0183555050505050565b81356116248161125d565b81546001600160a01b031981166001600160a01b03929092169182178355602084013561165081611164565b6001600160e01b03199190911690911760a09190911b67ffffffffffffffff60a01b16178155600181016116a7611689604085016113c3565b825467ffffffffffffffff19166001600160401b0391909116178255565b6116e86116b6606085016113c3565b82546fffffffffffffffff0000000000000000191660409190911b6fffffffffffffffff000000000000000016178255565b6117236116f7608085016113c3565b82805467ffffffffffffffff60801b191660809290921b67ffffffffffffffff60801b16919091179055565b5061173160a083018361137e565b61173f818360028601611455565b505061174e60c0830183611335565b61175c818360038601611560565b50505050565b5f60a082016001600160401b03808b1684526020818b1681860152818a16604086015260a060608601528288845260c0860190508993505f5b898110156117c25784356117ae81611164565b84168252938201939082019060010161179b565b5085810360808701528681528688838301375f818801830152601f909601601f19169095019094019a995050505050505050505056fea2646970667358221220e86905efdf0c516a922bbfa801771f34f618a2e34b129da1915da518b66afb9364736f6c63430008180033",
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
// Solidity: function deployments(uint64 ) view returns(address addr, uint64 chainId, uint64 deployHeight, uint64 attestInterval, uint64 blockPeriodNs, string name)
func (_PortalRegistry *PortalRegistryCaller) Deployments(opts *bind.CallOpts, arg0 uint64) (struct {
	Addr           common.Address
	ChainId        uint64
	DeployHeight   uint64
	AttestInterval uint64
	BlockPeriodNs  uint64
	Name           string
}, error) {
	var out []interface{}
	err := _PortalRegistry.contract.Call(opts, &out, "deployments", arg0)

	outstruct := new(struct {
		Addr           common.Address
		ChainId        uint64
		DeployHeight   uint64
		AttestInterval uint64
		BlockPeriodNs  uint64
		Name           string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Addr = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.ChainId = *abi.ConvertType(out[1], new(uint64)).(*uint64)
	outstruct.DeployHeight = *abi.ConvertType(out[2], new(uint64)).(*uint64)
	outstruct.AttestInterval = *abi.ConvertType(out[3], new(uint64)).(*uint64)
	outstruct.BlockPeriodNs = *abi.ConvertType(out[4], new(uint64)).(*uint64)
	outstruct.Name = *abi.ConvertType(out[5], new(string)).(*string)

	return *outstruct, err

}

// Deployments is a free data retrieval call binding the contract method 0x473d0452.
//
// Solidity: function deployments(uint64 ) view returns(address addr, uint64 chainId, uint64 deployHeight, uint64 attestInterval, uint64 blockPeriodNs, string name)
func (_PortalRegistry *PortalRegistrySession) Deployments(arg0 uint64) (struct {
	Addr           common.Address
	ChainId        uint64
	DeployHeight   uint64
	AttestInterval uint64
	BlockPeriodNs  uint64
	Name           string
}, error) {
	return _PortalRegistry.Contract.Deployments(&_PortalRegistry.CallOpts, arg0)
}

// Deployments is a free data retrieval call binding the contract method 0x473d0452.
//
// Solidity: function deployments(uint64 ) view returns(address addr, uint64 chainId, uint64 deployHeight, uint64 attestInterval, uint64 blockPeriodNs, string name)
func (_PortalRegistry *PortalRegistryCallerSession) Deployments(arg0 uint64) (struct {
	Addr           common.Address
	ChainId        uint64
	DeployHeight   uint64
	AttestInterval uint64
	BlockPeriodNs  uint64
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
// Solidity: function bulkRegister((address,uint64,uint64,uint64,uint64,uint64[],string)[] deps) returns()
func (_PortalRegistry *PortalRegistryTransactor) BulkRegister(opts *bind.TransactOpts, deps []PortalRegistryDeployment) (*types.Transaction, error) {
	return _PortalRegistry.contract.Transact(opts, "bulkRegister", deps)
}

// BulkRegister is a paid mutator transaction binding the contract method 0x52d482e2.
//
// Solidity: function bulkRegister((address,uint64,uint64,uint64,uint64,uint64[],string)[] deps) returns()
func (_PortalRegistry *PortalRegistrySession) BulkRegister(deps []PortalRegistryDeployment) (*types.Transaction, error) {
	return _PortalRegistry.Contract.BulkRegister(&_PortalRegistry.TransactOpts, deps)
}

// BulkRegister is a paid mutator transaction binding the contract method 0x52d482e2.
//
// Solidity: function bulkRegister((address,uint64,uint64,uint64,uint64,uint64[],string)[] deps) returns()
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
	BlockPeriodNs  uint64
	Shards         []uint64
	Name           string
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterPortalRegistered is a free log retrieval operation binding the contract event 0xb08d1911b978b0c040fa5e01711aa326770a97c5f00039d45e7ae8dec7409e73.
//
// Solidity: event PortalRegistered(uint64 indexed chainId, address indexed addr, uint64 deployHeight, uint64 attestInterval, uint64 blockPeriodNs, uint64[] shards, string name)
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
// Solidity: event PortalRegistered(uint64 indexed chainId, address indexed addr, uint64 deployHeight, uint64 attestInterval, uint64 blockPeriodNs, uint64[] shards, string name)
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
// Solidity: event PortalRegistered(uint64 indexed chainId, address indexed addr, uint64 deployHeight, uint64 attestInterval, uint64 blockPeriodNs, uint64[] shards, string name)
func (_PortalRegistry *PortalRegistryFilterer) ParsePortalRegistered(log types.Log) (*PortalRegistryPortalRegistered, error) {
	event := new(PortalRegistryPortalRegistered)
	if err := _PortalRegistry.contract.UnpackLog(event, "PortalRegistered", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
