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

// StablecoinProxyMetaData contains all meta data concerning the StablecoinProxy contract.
var StablecoinProxyMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_delegate\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"fallback\",\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"getImplementation\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"Upgraded\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AddressEmptyCode\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC1967InvalidImplementation\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC1967NonPayable\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FailedCall\",\"inputs\":[]}]",
	Bin: "0x608060405234801561001057600080fd5b5060405161046f38038061046f83398101604081905261002f91610279565b818161003b8282610044565b50505050610369565b61004d826100a3565b6040516001600160a01b038316907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b90600090a280511561009757610092828261011f565b505050565b61009f610196565b5050565b806001600160a01b03163b6000036100de57604051634c9c8ce360e01b81526001600160a01b03821660048201526024015b60405180910390fd5b7f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc80546001600160a01b0319166001600160a01b0392909216919091179055565b6060600080846001600160a01b03168460405161013c919061034d565b600060405180830381855af49150503d8060008114610177576040519150601f19603f3d011682016040523d82523d6000602084013e61017c565b606091505b50909250905061018d8583836101b7565b95945050505050565b34156101b55760405163b398979f60e01b815260040160405180910390fd5b565b6060826101cc576101c782610216565b61020f565b81511580156101e357506001600160a01b0384163b155b1561020c57604051639996b31560e01b81526001600160a01b03851660048201526024016100d5565b50805b9392505050565b8051156102265780518082602001fd5b60405163d6bda27560e01b815260040160405180910390fd5b634e487b7160e01b600052604160045260246000fd5b60005b83811015610270578181015183820152602001610258565b50506000910152565b6000806040838503121561028c57600080fd5b82516001600160a01b03811681146102a357600080fd5b60208401519092506001600160401b038111156102bf57600080fd5b8301601f810185136102d057600080fd5b80516001600160401b038111156102e9576102e961023f565b604051601f8201601f19908116603f011681016001600160401b03811182821017156103175761031761023f565b60405281815282820160200187101561032f57600080fd5b610340826020830160208601610255565b8093505050509250929050565b6000825161035f818460208701610255565b9190910192915050565b60f8806103776000396000f3fe608060405260043610601c5760003560e01c8063aaf10f42146024575b60226052565b005b348015602f57600080fd5b5060366060565b6040516001600160a01b03909116815260200160405180910390f35b605e605a606d565b609f565b565b60006068606d565b905090565b600060687f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc546001600160a01b031690565b3660008037600080366000845af43d6000803e80801560bd573d6000f35b3d6000fdfea26469706673582212208cc014af6997832f15e03cae4a6370d707a2cee6ca035726aa20e4089a2f29ed64736f6c634300081a0033",
}

// StablecoinProxyABI is the input ABI used to generate the binding from.
// Deprecated: Use StablecoinProxyMetaData.ABI instead.
var StablecoinProxyABI = StablecoinProxyMetaData.ABI

// StablecoinProxyBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use StablecoinProxyMetaData.Bin instead.
var StablecoinProxyBin = StablecoinProxyMetaData.Bin

// DeployStablecoinProxy deploys a new Ethereum contract, binding an instance of StablecoinProxy to it.
func DeployStablecoinProxy(auth *bind.TransactOpts, backend bind.ContractBackend, _delegate common.Address, _data []byte) (common.Address, *types.Transaction, *StablecoinProxy, error) {
	parsed, err := StablecoinProxyMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(StablecoinProxyBin), backend, _delegate, _data)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &StablecoinProxy{StablecoinProxyCaller: StablecoinProxyCaller{contract: contract}, StablecoinProxyTransactor: StablecoinProxyTransactor{contract: contract}, StablecoinProxyFilterer: StablecoinProxyFilterer{contract: contract}}, nil
}

// StablecoinProxy is an auto generated Go binding around an Ethereum contract.
type StablecoinProxy struct {
	StablecoinProxyCaller     // Read-only binding to the contract
	StablecoinProxyTransactor // Write-only binding to the contract
	StablecoinProxyFilterer   // Log filterer for contract events
}

// StablecoinProxyCaller is an auto generated read-only Go binding around an Ethereum contract.
type StablecoinProxyCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StablecoinProxyTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StablecoinProxyTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StablecoinProxyFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StablecoinProxyFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StablecoinProxySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StablecoinProxySession struct {
	Contract     *StablecoinProxy  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StablecoinProxyCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StablecoinProxyCallerSession struct {
	Contract *StablecoinProxyCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// StablecoinProxyTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StablecoinProxyTransactorSession struct {
	Contract     *StablecoinProxyTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// StablecoinProxyRaw is an auto generated low-level Go binding around an Ethereum contract.
type StablecoinProxyRaw struct {
	Contract *StablecoinProxy // Generic contract binding to access the raw methods on
}

// StablecoinProxyCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StablecoinProxyCallerRaw struct {
	Contract *StablecoinProxyCaller // Generic read-only contract binding to access the raw methods on
}

// StablecoinProxyTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StablecoinProxyTransactorRaw struct {
	Contract *StablecoinProxyTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStablecoinProxy creates a new instance of StablecoinProxy, bound to a specific deployed contract.
func NewStablecoinProxy(address common.Address, backend bind.ContractBackend) (*StablecoinProxy, error) {
	contract, err := bindStablecoinProxy(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StablecoinProxy{StablecoinProxyCaller: StablecoinProxyCaller{contract: contract}, StablecoinProxyTransactor: StablecoinProxyTransactor{contract: contract}, StablecoinProxyFilterer: StablecoinProxyFilterer{contract: contract}}, nil
}

// NewStablecoinProxyCaller creates a new read-only instance of StablecoinProxy, bound to a specific deployed contract.
func NewStablecoinProxyCaller(address common.Address, caller bind.ContractCaller) (*StablecoinProxyCaller, error) {
	contract, err := bindStablecoinProxy(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StablecoinProxyCaller{contract: contract}, nil
}

// NewStablecoinProxyTransactor creates a new write-only instance of StablecoinProxy, bound to a specific deployed contract.
func NewStablecoinProxyTransactor(address common.Address, transactor bind.ContractTransactor) (*StablecoinProxyTransactor, error) {
	contract, err := bindStablecoinProxy(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StablecoinProxyTransactor{contract: contract}, nil
}

// NewStablecoinProxyFilterer creates a new log filterer instance of StablecoinProxy, bound to a specific deployed contract.
func NewStablecoinProxyFilterer(address common.Address, filterer bind.ContractFilterer) (*StablecoinProxyFilterer, error) {
	contract, err := bindStablecoinProxy(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StablecoinProxyFilterer{contract: contract}, nil
}

// bindStablecoinProxy binds a generic wrapper to an already deployed contract.
func bindStablecoinProxy(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := StablecoinProxyMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StablecoinProxy *StablecoinProxyRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StablecoinProxy.Contract.StablecoinProxyCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StablecoinProxy *StablecoinProxyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StablecoinProxy.Contract.StablecoinProxyTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StablecoinProxy *StablecoinProxyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StablecoinProxy.Contract.StablecoinProxyTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StablecoinProxy *StablecoinProxyCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StablecoinProxy.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StablecoinProxy *StablecoinProxyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StablecoinProxy.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StablecoinProxy *StablecoinProxyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StablecoinProxy.Contract.contract.Transact(opts, method, params...)
}

// GetImplementation is a free data retrieval call binding the contract method 0xaaf10f42.
//
// Solidity: function getImplementation() view returns(address)
func (_StablecoinProxy *StablecoinProxyCaller) GetImplementation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StablecoinProxy.contract.Call(opts, &out, "getImplementation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetImplementation is a free data retrieval call binding the contract method 0xaaf10f42.
//
// Solidity: function getImplementation() view returns(address)
func (_StablecoinProxy *StablecoinProxySession) GetImplementation() (common.Address, error) {
	return _StablecoinProxy.Contract.GetImplementation(&_StablecoinProxy.CallOpts)
}

// GetImplementation is a free data retrieval call binding the contract method 0xaaf10f42.
//
// Solidity: function getImplementation() view returns(address)
func (_StablecoinProxy *StablecoinProxyCallerSession) GetImplementation() (common.Address, error) {
	return _StablecoinProxy.Contract.GetImplementation(&_StablecoinProxy.CallOpts)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_StablecoinProxy *StablecoinProxyTransactor) Fallback(opts *bind.TransactOpts, calldata []byte) (*types.Transaction, error) {
	return _StablecoinProxy.contract.RawTransact(opts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_StablecoinProxy *StablecoinProxySession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _StablecoinProxy.Contract.Fallback(&_StablecoinProxy.TransactOpts, calldata)
}

// Fallback is a paid mutator transaction binding the contract fallback function.
//
// Solidity: fallback() payable returns()
func (_StablecoinProxy *StablecoinProxyTransactorSession) Fallback(calldata []byte) (*types.Transaction, error) {
	return _StablecoinProxy.Contract.Fallback(&_StablecoinProxy.TransactOpts, calldata)
}

// StablecoinProxyUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the StablecoinProxy contract.
type StablecoinProxyUpgradedIterator struct {
	Event *StablecoinProxyUpgraded // Event containing the contract specifics and raw log

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
func (it *StablecoinProxyUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StablecoinProxyUpgraded)
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
		it.Event = new(StablecoinProxyUpgraded)
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
func (it *StablecoinProxyUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StablecoinProxyUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StablecoinProxyUpgraded represents a Upgraded event raised by the StablecoinProxy contract.
type StablecoinProxyUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_StablecoinProxy *StablecoinProxyFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*StablecoinProxyUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _StablecoinProxy.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &StablecoinProxyUpgradedIterator{contract: _StablecoinProxy.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_StablecoinProxy *StablecoinProxyFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *StablecoinProxyUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _StablecoinProxy.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StablecoinProxyUpgraded)
				if err := _StablecoinProxy.contract.UnpackLog(event, "Upgraded", log); err != nil {
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

// ParseUpgraded is a log parse operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_StablecoinProxy *StablecoinProxyFilterer) ParseUpgraded(log types.Log) (*StablecoinProxyUpgraded, error) {
	event := new(StablecoinProxyUpgraded)
	if err := _StablecoinProxy.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
