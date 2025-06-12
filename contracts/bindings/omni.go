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

// OmniMetaData contains all meta data concerning the Omni contract.
var OmniMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"initialSupply\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"DOMAIN_SEPARATOR\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"allowance\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"approve\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"balanceOf\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"decimals\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"decreaseAllowance\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"subtractedValue\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"eip712Domain\",\"inputs\":[],\"outputs\":[{\"name\":\"fields\",\"type\":\"bytes1\",\"internalType\":\"bytes1\"},{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"version\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"verifyingContract\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"extensions\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"increaseAllowance\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"addedValue\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"name\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"nonces\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"permit\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"deadline\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"v\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"r\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"s\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"symbol\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalSupply\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transfer\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferFrom\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Approval\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"EIP712DomainChanged\",\"inputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Transfer\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"InvalidShortString\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"StringTooLong\",\"inputs\":[{\"name\":\"str\",\"type\":\"string\",\"internalType\":\"string\"}]}]",
	Bin: "0x61016060405234801562000011575f80fd5b506040516200161c3803806200161c8339810160408190526200003491620002e8565b6040518060400160405280600c81526020016b4f6d6e69204e6574776f726b60a01b81525080604051806040016040528060018152602001603160f81b8152506040518060400160405280600c81526020016b4f6d6e69204e6574776f726b60a01b815250604051806040016040528060048152602001634f4d4e4960e01b8152508160039081620000c79190620003c1565b506004620000d68282620003c1565b50620000e891508390506005620001a4565b61012052620000f9816006620001a4565b61014052815160208084019190912060e052815190820120610100524660a0526200018660e05161010051604080517f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f60208201529081019290925260608201524660808201523060a08201525f9060c00160405160208183030381529060405280519060200120905090565b60805250503060c052506200019c8183620001dc565b50506200051f565b5f602083511015620001c357620001bb83620002a1565b9050620001d6565b81620001d08482620003c1565b5060ff90505b92915050565b6001600160a01b038216620002385760405162461bcd60e51b815260206004820152601f60248201527f45524332303a206d696e7420746f20746865207a65726f20616464726573730060448201526064015b60405180910390fd5b8060025f8282546200024b91906200048d565b90915550506001600160a01b0382165f81815260208181526040808320805486019055518481527fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef910160405180910390a35050565b5f80829050601f81511115620002ce578260405163305a27a960e01b81526004016200022f9190620004ad565b8051620002db82620004fb565b179392505050565b505050565b5f8060408385031215620002fa575f80fd5b825160208401519092506001600160a01b038116811462000319575f80fd5b809150509250929050565b634e487b7160e01b5f52604160045260245ffd5b600181811c908216806200034d57607f821691505b6020821081036200036c57634e487b7160e01b5f52602260045260245ffd5b50919050565b601f821115620002e357805f5260205f20601f840160051c81016020851015620003995750805b601f840160051c820191505b81811015620003ba575f8155600101620003a5565b5050505050565b81516001600160401b03811115620003dd57620003dd62000324565b620003f581620003ee845462000338565b8462000372565b602080601f8311600181146200042b575f8415620004135750858301515b5f19600386901b1c1916600185901b17855562000485565b5f85815260208120601f198616915b828110156200045b578886015182559484019460019091019084016200043a565b50858210156200047957878501515f19600388901b60f8161c191681555b505060018460011b0185555b505050505050565b80820180821115620001d657634e487b7160e01b5f52601160045260245ffd5b5f602080835283518060208501525f5b81811015620004db57858101830151858201604001528201620004bd565b505f604082860101526040601f19601f8301168501019250505092915050565b805160208083015191908110156200036c575f1960209190910360031b1b16919050565b60805160a05160c05160e0516101005161012051610140516110ab620005715f395f61038001525f61035501525f610a0b01525f6109e301525f61093e01525f61096801525f61099201526110ab5ff3fe608060405234801561000f575f80fd5b50600436106100f0575f3560e01c806370a0823111610093578063a457c2d711610063578063a457c2d7146101e2578063a9059cbb146101f5578063d505accf14610208578063dd62ed3e1461021d575f80fd5b806370a08231146101845780637ecebe00146101ac57806384b0196e146101bf57806395d89b41146101da575f80fd5b806323b872dd116100ce57806323b872dd14610147578063313ce5671461015a5780633644e515146101695780633950935114610171575f80fd5b806306fdde03146100f4578063095ea7b31461011257806318160ddd14610135575b5f80fd5b6100fc610230565b6040516101099190610e2d565b60405180910390f35b610125610120366004610e61565b6102c0565b6040519015158152602001610109565b6002545b604051908152602001610109565b610125610155366004610e89565b6102d9565b60405160128152602001610109565b6101396102fc565b61012561017f366004610e61565b61030a565b610139610192366004610ec2565b6001600160a01b03165f9081526020819052604090205490565b6101396101ba366004610ec2565b61032b565b6101c7610348565b6040516101099796959493929190610edb565b6100fc6103cf565b6101256101f0366004610e61565b6103de565b610125610203366004610e61565b61045d565b61021b610216366004610f72565b61046a565b005b61013961022b366004610fdf565b6105cb565b60606003805461023f90611010565b80601f016020809104026020016040519081016040528092919081815260200182805461026b90611010565b80156102b65780601f1061028d576101008083540402835291602001916102b6565b820191905f5260205f20905b81548152906001019060200180831161029957829003601f168201915b5050505050905090565b5f336102cd8185856105f5565b60019150505b92915050565b5f336102e6858285610718565b6102f1858585610790565b506001949350505050565b5f610305610932565b905090565b5f336102cd81858561031c83836105cb565b6103269190611042565b6105f5565b6001600160a01b0381165f908152600760205260408120546102d3565b5f6060808280808361037b7f00000000000000000000000000000000000000000000000000000000000000006005610a5b565b6103a67f00000000000000000000000000000000000000000000000000000000000000006006610a5b565b604080515f80825260208201909252600f60f81b9b939a50919850469750309650945092509050565b60606004805461023f90611010565b5f33816103eb82866105cb565b9050838110156104505760405162461bcd60e51b815260206004820152602560248201527f45524332303a2064656372656173656420616c6c6f77616e63652062656c6f77604482015264207a65726f60d81b60648201526084015b60405180910390fd5b6102f182868684036105f5565b5f336102cd818585610790565b834211156104ba5760405162461bcd60e51b815260206004820152601d60248201527f45524332305065726d69743a206578706972656420646561646c696e650000006044820152606401610447565b5f7f6e71edae12b1b97f4d1f60370fef10105fa2faae0126114a169c64845d6126c98888886104e88c610b04565b6040805160208101969096526001600160a01b0394851690860152929091166060840152608083015260a082015260c0810186905260e0016040516020818303038152906040528051906020012090505f61054282610b2b565b90505f61055182878787610b57565b9050896001600160a01b0316816001600160a01b0316146105b45760405162461bcd60e51b815260206004820152601e60248201527f45524332305065726d69743a20696e76616c6964207369676e617475726500006044820152606401610447565b6105bf8a8a8a6105f5565b50505050505050505050565b6001600160a01b039182165f90815260016020908152604080832093909416825291909152205490565b6001600160a01b0383166106575760405162461bcd60e51b8152602060048201526024808201527f45524332303a20617070726f76652066726f6d20746865207a65726f206164646044820152637265737360e01b6064820152608401610447565b6001600160a01b0382166106b85760405162461bcd60e51b815260206004820152602260248201527f45524332303a20617070726f766520746f20746865207a65726f206164647265604482015261737360f01b6064820152608401610447565b6001600160a01b038381165f8181526001602090815260408083209487168084529482529182902085905590518481527f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925910160405180910390a3505050565b5f61072384846105cb565b90505f19811461078a578181101561077d5760405162461bcd60e51b815260206004820152601d60248201527f45524332303a20696e73756666696369656e7420616c6c6f77616e63650000006044820152606401610447565b61078a84848484036105f5565b50505050565b6001600160a01b0383166107f45760405162461bcd60e51b815260206004820152602560248201527f45524332303a207472616e736665722066726f6d20746865207a65726f206164604482015264647265737360d81b6064820152608401610447565b6001600160a01b0382166108565760405162461bcd60e51b815260206004820152602360248201527f45524332303a207472616e7366657220746f20746865207a65726f206164647260448201526265737360e81b6064820152608401610447565b6001600160a01b0383165f90815260208190526040902054818110156108cd5760405162461bcd60e51b815260206004820152602660248201527f45524332303a207472616e7366657220616d6f756e7420657863656564732062604482015265616c616e636560d01b6064820152608401610447565b6001600160a01b038481165f81815260208181526040808320878703905593871680835291849020805487019055925185815290927fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef910160405180910390a361078a565b5f306001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614801561098a57507f000000000000000000000000000000000000000000000000000000000000000046145b156109b457507f000000000000000000000000000000000000000000000000000000000000000090565b610305604080517f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f60208201527f0000000000000000000000000000000000000000000000000000000000000000918101919091527f000000000000000000000000000000000000000000000000000000000000000060608201524660808201523060a08201525f9060c00160405160208183030381529060405280519060200120905090565b606060ff8314610a7557610a6e83610b7d565b90506102d3565b818054610a8190611010565b80601f0160208091040260200160405190810160405280929190818152602001828054610aad90611010565b8015610af85780601f10610acf57610100808354040283529160200191610af8565b820191905f5260205f20905b815481529060010190602001808311610adb57829003601f168201915b505050505090506102d3565b6001600160a01b0381165f9081526007602052604090208054600181018255905b50919050565b5f6102d3610b37610932565b8360405161190160f01b8152600281019290925260228201526042902090565b5f805f610b6687878787610bba565b91509150610b7381610c77565b5095945050505050565b60605f610b8983610dc3565b6040805160208082528183019092529192505f91906020820181803683375050509182525060208101929092525090565b5f807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a0831115610bef57505f90506003610c6e565b604080515f8082526020820180845289905260ff881692820192909252606081018690526080810185905260019060a0016020604051602081039080840390855afa158015610c40573d5f803e3d5ffd5b5050604051601f1901519150506001600160a01b038116610c68575f60019250925050610c6e565b91505f90505b94509492505050565b5f816004811115610c8a57610c8a611061565b03610c925750565b6001816004811115610ca657610ca6611061565b03610cf35760405162461bcd60e51b815260206004820152601860248201527f45434453413a20696e76616c6964207369676e617475726500000000000000006044820152606401610447565b6002816004811115610d0757610d07611061565b03610d545760405162461bcd60e51b815260206004820152601f60248201527f45434453413a20696e76616c6964207369676e6174757265206c656e677468006044820152606401610447565b6003816004811115610d6857610d68611061565b03610dc05760405162461bcd60e51b815260206004820152602260248201527f45434453413a20696e76616c6964207369676e6174757265202773272076616c604482015261756560f01b6064820152608401610447565b50565b5f60ff8216601f8111156102d357604051632cd44ac360e21b815260040160405180910390fd5b5f81518084525f5b81811015610e0e57602081850181015186830182015201610df2565b505f602082860101526020601f19601f83011685010191505092915050565b602081525f610e3f6020830184610dea565b9392505050565b80356001600160a01b0381168114610e5c575f80fd5b919050565b5f8060408385031215610e72575f80fd5b610e7b83610e46565b946020939093013593505050565b5f805f60608486031215610e9b575f80fd5b610ea484610e46565b9250610eb260208501610e46565b9150604084013590509250925092565b5f60208284031215610ed2575f80fd5b610e3f82610e46565b60ff60f81b881681525f602060e06020840152610efb60e084018a610dea565b8381036040850152610f0d818a610dea565b606085018990526001600160a01b038816608086015260a0850187905284810360c0860152855180825260208088019350909101905f5b81811015610f6057835183529284019291840191600101610f44565b50909c9b505050505050505050505050565b5f805f805f805f60e0888a031215610f88575f80fd5b610f9188610e46565b9650610f9f60208901610e46565b95506040880135945060608801359350608088013560ff81168114610fc2575f80fd5b9699959850939692959460a0840135945060c09093013592915050565b5f8060408385031215610ff0575f80fd5b610ff983610e46565b915061100760208401610e46565b90509250929050565b600181811c9082168061102457607f821691505b602082108103610b2557634e487b7160e01b5f52602260045260245ffd5b808201808211156102d357634e487b7160e01b5f52601160045260245ffd5b634e487b7160e01b5f52602160045260245ffdfea26469706673582212204cf6fec89d39dab924781d6072e10e0002b6ef2d7a5ce713487efcd10feac41064736f6c63430008180033",
}

// OmniABI is the input ABI used to generate the binding from.
// Deprecated: Use OmniMetaData.ABI instead.
var OmniABI = OmniMetaData.ABI

// OmniBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use OmniMetaData.Bin instead.
var OmniBin = OmniMetaData.Bin

// DeployOmni deploys a new Ethereum contract, binding an instance of Omni to it.
func DeployOmni(auth *bind.TransactOpts, backend bind.ContractBackend, initialSupply *big.Int, recipient common.Address) (common.Address, *types.Transaction, *Omni, error) {
	parsed, err := OmniMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(OmniBin), backend, initialSupply, recipient)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Omni{OmniCaller: OmniCaller{contract: contract}, OmniTransactor: OmniTransactor{contract: contract}, OmniFilterer: OmniFilterer{contract: contract}}, nil
}

// Omni is an auto generated Go binding around an Ethereum contract.
type Omni struct {
	OmniCaller     // Read-only binding to the contract
	OmniTransactor // Write-only binding to the contract
	OmniFilterer   // Log filterer for contract events
}

// OmniCaller is an auto generated read-only Go binding around an Ethereum contract.
type OmniCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OmniTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OmniTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OmniFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OmniFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OmniSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OmniSession struct {
	Contract     *Omni             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OmniCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OmniCallerSession struct {
	Contract *OmniCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// OmniTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OmniTransactorSession struct {
	Contract     *OmniTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OmniRaw is an auto generated low-level Go binding around an Ethereum contract.
type OmniRaw struct {
	Contract *Omni // Generic contract binding to access the raw methods on
}

// OmniCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OmniCallerRaw struct {
	Contract *OmniCaller // Generic read-only contract binding to access the raw methods on
}

// OmniTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OmniTransactorRaw struct {
	Contract *OmniTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOmni creates a new instance of Omni, bound to a specific deployed contract.
func NewOmni(address common.Address, backend bind.ContractBackend) (*Omni, error) {
	contract, err := bindOmni(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Omni{OmniCaller: OmniCaller{contract: contract}, OmniTransactor: OmniTransactor{contract: contract}, OmniFilterer: OmniFilterer{contract: contract}}, nil
}

// NewOmniCaller creates a new read-only instance of Omni, bound to a specific deployed contract.
func NewOmniCaller(address common.Address, caller bind.ContractCaller) (*OmniCaller, error) {
	contract, err := bindOmni(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OmniCaller{contract: contract}, nil
}

// NewOmniTransactor creates a new write-only instance of Omni, bound to a specific deployed contract.
func NewOmniTransactor(address common.Address, transactor bind.ContractTransactor) (*OmniTransactor, error) {
	contract, err := bindOmni(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OmniTransactor{contract: contract}, nil
}

// NewOmniFilterer creates a new log filterer instance of Omni, bound to a specific deployed contract.
func NewOmniFilterer(address common.Address, filterer bind.ContractFilterer) (*OmniFilterer, error) {
	contract, err := bindOmni(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OmniFilterer{contract: contract}, nil
}

// bindOmni binds a generic wrapper to an already deployed contract.
func bindOmni(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := OmniMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Omni *OmniRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Omni.Contract.OmniCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Omni *OmniRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Omni.Contract.OmniTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Omni *OmniRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Omni.Contract.OmniTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Omni *OmniCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Omni.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Omni *OmniTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Omni.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Omni *OmniTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Omni.Contract.contract.Transact(opts, method, params...)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_Omni *OmniCaller) DOMAINSEPARATOR(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Omni.contract.Call(opts, &out, "DOMAIN_SEPARATOR")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_Omni *OmniSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _Omni.Contract.DOMAINSEPARATOR(&_Omni.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_Omni *OmniCallerSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _Omni.Contract.DOMAINSEPARATOR(&_Omni.CallOpts)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Omni *OmniCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Omni.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Omni *OmniSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Omni.Contract.Allowance(&_Omni.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_Omni *OmniCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Omni.Contract.Allowance(&_Omni.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Omni *OmniCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Omni.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Omni *OmniSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Omni.Contract.BalanceOf(&_Omni.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_Omni *OmniCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _Omni.Contract.BalanceOf(&_Omni.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Omni *OmniCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Omni.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Omni *OmniSession) Decimals() (uint8, error) {
	return _Omni.Contract.Decimals(&_Omni.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Omni *OmniCallerSession) Decimals() (uint8, error) {
	return _Omni.Contract.Decimals(&_Omni.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_Omni *OmniCaller) Eip712Domain(opts *bind.CallOpts) (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	var out []interface{}
	err := _Omni.contract.Call(opts, &out, "eip712Domain")

	outstruct := new(struct {
		Fields            [1]byte
		Name              string
		Version           string
		ChainId           *big.Int
		VerifyingContract common.Address
		Salt              [32]byte
		Extensions        []*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Fields = *abi.ConvertType(out[0], new([1]byte)).(*[1]byte)
	outstruct.Name = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Version = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.ChainId = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.VerifyingContract = *abi.ConvertType(out[4], new(common.Address)).(*common.Address)
	outstruct.Salt = *abi.ConvertType(out[5], new([32]byte)).(*[32]byte)
	outstruct.Extensions = *abi.ConvertType(out[6], new([]*big.Int)).(*[]*big.Int)

	return *outstruct, err

}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_Omni *OmniSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _Omni.Contract.Eip712Domain(&_Omni.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_Omni *OmniCallerSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _Omni.Contract.Eip712Domain(&_Omni.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Omni *OmniCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Omni.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Omni *OmniSession) Name() (string, error) {
	return _Omni.Contract.Name(&_Omni.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Omni *OmniCallerSession) Name() (string, error) {
	return _Omni.Contract.Name(&_Omni.CallOpts)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_Omni *OmniCaller) Nonces(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Omni.contract.Call(opts, &out, "nonces", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_Omni *OmniSession) Nonces(owner common.Address) (*big.Int, error) {
	return _Omni.Contract.Nonces(&_Omni.CallOpts, owner)
}

// Nonces is a free data retrieval call binding the contract method 0x7ecebe00.
//
// Solidity: function nonces(address owner) view returns(uint256)
func (_Omni *OmniCallerSession) Nonces(owner common.Address) (*big.Int, error) {
	return _Omni.Contract.Nonces(&_Omni.CallOpts, owner)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Omni *OmniCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Omni.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Omni *OmniSession) Symbol() (string, error) {
	return _Omni.Contract.Symbol(&_Omni.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Omni *OmniCallerSession) Symbol() (string, error) {
	return _Omni.Contract.Symbol(&_Omni.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Omni *OmniCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Omni.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Omni *OmniSession) TotalSupply() (*big.Int, error) {
	return _Omni.Contract.TotalSupply(&_Omni.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_Omni *OmniCallerSession) TotalSupply() (*big.Int, error) {
	return _Omni.Contract.TotalSupply(&_Omni.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Omni *OmniTransactor) Approve(opts *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Omni.contract.Transact(opts, "approve", spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Omni *OmniSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Omni.Contract.Approve(&_Omni.TransactOpts, spender, amount)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 amount) returns(bool)
func (_Omni *OmniTransactorSession) Approve(spender common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Omni.Contract.Approve(&_Omni.TransactOpts, spender, amount)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Omni *OmniTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Omni.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Omni *OmniSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Omni.Contract.DecreaseAllowance(&_Omni.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(address spender, uint256 subtractedValue) returns(bool)
func (_Omni *OmniTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _Omni.Contract.DecreaseAllowance(&_Omni.TransactOpts, spender, subtractedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Omni *OmniTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Omni.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Omni *OmniSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Omni.Contract.IncreaseAllowance(&_Omni.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(address spender, uint256 addedValue) returns(bool)
func (_Omni *OmniTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _Omni.Contract.IncreaseAllowance(&_Omni.TransactOpts, spender, addedValue)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_Omni *OmniTransactor) Permit(opts *bind.TransactOpts, owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Omni.contract.Transact(opts, "permit", owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_Omni *OmniSession) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Omni.Contract.Permit(&_Omni.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// Permit is a paid mutator transaction binding the contract method 0xd505accf.
//
// Solidity: function permit(address owner, address spender, uint256 value, uint256 deadline, uint8 v, bytes32 r, bytes32 s) returns()
func (_Omni *OmniTransactorSession) Permit(owner common.Address, spender common.Address, value *big.Int, deadline *big.Int, v uint8, r [32]byte, s [32]byte) (*types.Transaction, error) {
	return _Omni.Contract.Permit(&_Omni.TransactOpts, owner, spender, value, deadline, v, r, s)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_Omni *OmniTransactor) Transfer(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Omni.contract.Transact(opts, "transfer", to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_Omni *OmniSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Omni.Contract.Transfer(&_Omni.TransactOpts, to, amount)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 amount) returns(bool)
func (_Omni *OmniTransactorSession) Transfer(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Omni.Contract.Transfer(&_Omni.TransactOpts, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_Omni *OmniTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Omni.contract.Transact(opts, "transferFrom", from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_Omni *OmniSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Omni.Contract.TransferFrom(&_Omni.TransactOpts, from, to, amount)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 amount) returns(bool)
func (_Omni *OmniTransactorSession) TransferFrom(from common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _Omni.Contract.TransferFrom(&_Omni.TransactOpts, from, to, amount)
}

// OmniApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Omni contract.
type OmniApprovalIterator struct {
	Event *OmniApproval // Event containing the contract specifics and raw log

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
func (it *OmniApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OmniApproval)
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
		it.Event = new(OmniApproval)
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
func (it *OmniApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OmniApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OmniApproval represents a Approval event raised by the Omni contract.
type OmniApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Omni *OmniFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*OmniApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Omni.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &OmniApprovalIterator{contract: _Omni.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Omni *OmniFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *OmniApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Omni.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OmniApproval)
				if err := _Omni.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Omni *OmniFilterer) ParseApproval(log types.Log) (*OmniApproval, error) {
	event := new(OmniApproval)
	if err := _Omni.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OmniEIP712DomainChangedIterator is returned from FilterEIP712DomainChanged and is used to iterate over the raw logs and unpacked data for EIP712DomainChanged events raised by the Omni contract.
type OmniEIP712DomainChangedIterator struct {
	Event *OmniEIP712DomainChanged // Event containing the contract specifics and raw log

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
func (it *OmniEIP712DomainChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OmniEIP712DomainChanged)
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
		it.Event = new(OmniEIP712DomainChanged)
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
func (it *OmniEIP712DomainChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OmniEIP712DomainChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OmniEIP712DomainChanged represents a EIP712DomainChanged event raised by the Omni contract.
type OmniEIP712DomainChanged struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEIP712DomainChanged is a free log retrieval operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_Omni *OmniFilterer) FilterEIP712DomainChanged(opts *bind.FilterOpts) (*OmniEIP712DomainChangedIterator, error) {

	logs, sub, err := _Omni.contract.FilterLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return &OmniEIP712DomainChangedIterator{contract: _Omni.contract, event: "EIP712DomainChanged", logs: logs, sub: sub}, nil
}

// WatchEIP712DomainChanged is a free log subscription operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_Omni *OmniFilterer) WatchEIP712DomainChanged(opts *bind.WatchOpts, sink chan<- *OmniEIP712DomainChanged) (event.Subscription, error) {

	logs, sub, err := _Omni.contract.WatchLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OmniEIP712DomainChanged)
				if err := _Omni.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
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

// ParseEIP712DomainChanged is a log parse operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_Omni *OmniFilterer) ParseEIP712DomainChanged(log types.Log) (*OmniEIP712DomainChanged, error) {
	event := new(OmniEIP712DomainChanged)
	if err := _Omni.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OmniTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Omni contract.
type OmniTransferIterator struct {
	Event *OmniTransfer // Event containing the contract specifics and raw log

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
func (it *OmniTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OmniTransfer)
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
		it.Event = new(OmniTransfer)
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
func (it *OmniTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OmniTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OmniTransfer represents a Transfer event raised by the Omni contract.
type OmniTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Omni *OmniFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*OmniTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Omni.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &OmniTransferIterator{contract: _Omni.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Omni *OmniFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *OmniTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Omni.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OmniTransfer)
				if err := _Omni.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Omni *OmniFilterer) ParseTransfer(log types.Log) (*OmniTransfer, error) {
	event := new(OmniTransfer)
	if err := _Omni.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
