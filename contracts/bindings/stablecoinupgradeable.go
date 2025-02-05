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

// StablecoinUpgradeableMetaData contains all meta data concerning the StablecoinUpgradeable contract.
var StablecoinUpgradeableMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"BURNER_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"CLAWBACKER_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"DEFAULT_ADMIN_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MINTER_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"PAUSER_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"UPGRADER_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"UPGRADE_INTERFACE_VERSION\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"accountPaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"allowance\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"approve\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"balanceOf\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"burn\",\"inputs\":[{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"clawback\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"decimals\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRoleAdmin\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"grantRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"hasRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"name_\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"symbol_\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"minter_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"admin_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"upgrader_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"pauser_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"clawbacker_\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"mint\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"name\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pauseAccounts\",\"inputs\":[{\"name\":\"accounts\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"proxiableUUID\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"callerConfirmation\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"revokeRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"symbol\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"totalSupply\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transfer\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferFrom\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unpauseAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"upgradeToAndCall\",\"inputs\":[{\"name\":\"newImplementation\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"event\",\"name\":\"AccountPaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AccountUnpaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Approval\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"spender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleAdminChanged\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"previousAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"newAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleGranted\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleRevoked\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Transfer\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Upgraded\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AccessControlBadConfirmation\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AccessControlUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"neededRole\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"AccountIsNotPaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"AccountIsPaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"AddressEmptyCode\",\"inputs\":[{\"name\":\"target\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC1967InvalidImplementation\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC1967NonPayable\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ERC20InsufficientAllowance\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"allowance\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"needed\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ERC20InsufficientBalance\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"balance\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"needed\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ERC20InvalidApprover\",\"inputs\":[{\"name\":\"approver\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC20InvalidReceiver\",\"inputs\":[{\"name\":\"receiver\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC20InvalidSender\",\"inputs\":[{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ERC20InvalidSpender\",\"inputs\":[{\"name\":\"spender\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"EnforcedPause\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ExpectedPause\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"FailedCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UUPSUnauthorizedCallContext\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"UUPSUnsupportedProxiableUUID\",\"inputs\":[{\"name\":\"slot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}]",
	Bin: "0x60a06040523060805234801561001457600080fd5b5061001d610022565b6100d4565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000900460ff16156100725760405163f92ee8a960e01b815260040160405180910390fd5b80546001600160401b03908116146100d15780546001600160401b0319166001600160401b0390811782556040519081527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b50565b6080516120f96100fd6000396000818161113e0152818161116701526112ca01526120f96000f3fe6080604052600436106101f95760003560e01c806370a082311161010d578063ad3cb1cc116100a0578063d547741f1161006f578063d547741f14610602578063dd62ed3e14610622578063e63ab1e914610642578063f3df317e14610664578063f72c0d8b1461068457600080fd5b8063ad3cb1cc1461055d578063bc8c4b4f1461058e578063c91f0c53146105ae578063d5391393146105ce57600080fd5b806394408b9a116100dc57806394408b9a146104f357806395d89b4114610513578063a217fddf14610528578063a9059cbb1461053d57600080fd5b806370a082311461045b5780638456cb591461049e578063917b1ace146104b357806391d14854146104d357600080fd5b8063313ce5671161019057806342966c681161015f57806342966c68146103ba578063475ca324146103da5780634f1ef2861461040e57806352d1902d146104215780635c975abb1461043657600080fd5b8063313ce5671461034957806336568abe146103655780633f4ba83a1461038557806340c10f191461039a57600080fd5b806323b872dd116101cc57806323b872dd146102b3578063248a9ca3146102d3578063282c51f3146102f35780632f2ff15d1461032757600080fd5b806301ffc9a7146101fe57806306fdde0314610233578063095ea7b31461025557806318160ddd14610275575b600080fd5b34801561020a57600080fd5b5061021e610219366004611aa4565b6106b8565b60405190151581526020015b60405180910390f35b34801561023f57600080fd5b506102486106ef565b60405161022a9190611af2565b34801561026157600080fd5b5061021e610270366004611b41565b6107b2565b34801561028157600080fd5b507f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace02545b60405190815260200161022a565b3480156102bf57600080fd5b5061021e6102ce366004611b6b565b61083a565b3480156102df57600080fd5b506102a56102ee366004611ba8565b610860565b3480156102ff57600080fd5b506102a57f9667e80708b6eeeb0053fa0cca44e028ff548e2a9f029edfeac87c118b08b7c881565b34801561033357600080fd5b50610347610342366004611bc1565b610882565b005b34801561035557600080fd5b506040516012815260200161022a565b34801561037157600080fd5b50610347610380366004611bc1565b6108a4565b34801561039157600080fd5b506103476108dc565b3480156103a657600080fd5b506103476103b5366004611b41565b6108ff565b3480156103c657600080fd5b506103476103d5366004611ba8565b610933565b3480156103e657600080fd5b506102a57f715bacafb7a853b9b91b59ae724920a9eb0c006c5b318ac393fa1bc8974edd9881565b61034761041c366004611c7d565b61096b565b34801561042d57600080fd5b506102a5610986565b34801561044257600080fd5b506000805160206120a48339815191525460ff1661021e565b34801561046757600080fd5b506102a5610476366004611cdf565b6001600160a01b03166000908152600080516020612024833981519152602052604090205490565b3480156104aa57600080fd5b506103476109a3565b3480156104bf57600080fd5b506103476104ce366004611cfa565b6109c3565b3480156104df57600080fd5b5061021e6104ee366004611bc1565b610ad7565b3480156104ff57600080fd5b5061034761050e366004611cdf565b610b0f565b34801561051f57600080fd5b50610248610b30565b34801561053457600080fd5b506102a5600081565b34801561054957600080fd5b5061021e610558366004611b41565b610b6f565b34801561056957600080fd5b50610248604051806040016040528060058152602001640352e302e360dc1b81525081565b34801561059a57600080fd5b5061021e6105a9366004611cdf565b610b87565b3480156105ba57600080fd5b506103476105c9366004611d91565b610bc4565b3480156105da57600080fd5b506102a57ff0887ba65ee2024ea881d91b74c2450ef19e1557f03bed3ea9f16b037cbe2dc981565b34801561060e57600080fd5b5061034761061d366004611bc1565b610da1565b34801561062e57600080fd5b506102a561063d366004611e4d565b610dbd565b34801561064e57600080fd5b506102a560008051602061206483398151915281565b34801561067057600080fd5b5061034761067f366004611b41565b610e07565b34801561069057600080fd5b506102a57fa615a8afb6fffcb8c6809ac0997b5c9c12b8cc97651150f14c8f6203168cff4c81565b60006001600160e01b03198216637965db0b60e01b14806106e957506301ffc9a760e01b6001600160e01b03198316145b92915050565b7f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace0380546060916000805160206120248339815191529161072e90611e77565b80601f016020809104026020016040519081016040528092919081815260200182805461075a90611e77565b80156107a75780601f1061077c576101008083540402835291602001916107a7565b820191906000526020600020905b81548152906001019060200180831161078a57829003601f168201915b505050505091505090565b6000826107be81610b87565b156107ec5760405163c8c29b9960e01b81526001600160a01b03821660048201526024015b60405180910390fd5b336107f681610b87565b1561081f5760405163c8c29b9960e01b81526001600160a01b03821660048201526024016107e3565b610827610e3b565b6108318585610e6e565b95945050505050565b600033610848858285610e7c565b610853858585610edd565b60019150505b9392505050565b6000908152600080516020612084833981519152602052604090206001015490565b61088b82610860565b61089481610f3c565b61089e8383610f46565b50505050565b6001600160a01b03811633146108cd5760405163334bd91960e11b815260040160405180910390fd5b6108d78282610feb565b505050565b6000805160206120648339815191526108f481610f3c565b6108fc611067565b50565b7ff0887ba65ee2024ea881d91b74c2450ef19e1557f03bed3ea9f16b037cbe2dc961092981610f3c565b6108d783836110c7565b7f9667e80708b6eeeb0053fa0cca44e028ff548e2a9f029edfeac87c118b08b7c861095d81610f3c565b61096733836110fd565b5050565b610973611133565b61097c826111d8565b6109678282611202565b60006109906112bf565b5060008051602061204483398151915290565b6000805160206120648339815191526109bb81610f3c565b6108fc611308565b6000805160206120648339815191526109db81610f3c565b600082815b81811015610acf57826001600160a01b0316868683818110610a0457610a04611eb1565b9050602002016020810190610a199190611cdf565b6001600160a01b031611610a6f5760405162461bcd60e51b815260206004820152601a60248201527f4164647265737365732073686f756c6420626520736f7274656400000000000060448201526064016107e3565b610a9e868683818110610a8457610a84611eb1565b9050602002016020810190610a999190611cdf565b611351565b858582818110610ab057610ab0611eb1565b9050602002016020810190610ac59190611cdf565b92506001016109e0565b505050505050565b6000918252600080516020612084833981519152602090815260408084206001600160a01b0393909316845291905290205460ff1690565b600080516020612064833981519152610b2781610f3c565b61096782611401565b7f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace0480546060916000805160206120248339815191529161072e90611e77565b600033610b7d818585610edd565b5060019392505050565b6001600160a01b031660009081527f345cc2404af916c3db112e7a6103770647a90ed78a5d681e21dc2e1174232900602052604090205460ff1690565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a008054600160401b810460ff16159067ffffffffffffffff16600081158015610c0a5750825b905060008267ffffffffffffffff166001148015610c275750303b155b905081158015610c35575080155b15610c535760405163f92ee8a960e01b815260040160405180910390fd5b845467ffffffffffffffff191660011785558315610c7d57845460ff60401b1916600160401b1785555b610c878c8c6114a4565b610c8f6114b6565b610c976114b6565b610ca260008a610f46565b50610ccd7ff0887ba65ee2024ea881d91b74c2450ef19e1557f03bed3ea9f16b037cbe2dc98b610f46565b50610cf87fa615a8afb6fffcb8c6809ac0997b5c9c12b8cc97651150f14c8f6203168cff4c89610f46565b50610d016114be565b610d096114b6565b610d2160008051602061206483398151915288610f46565b50610d4c7f715bacafb7a853b9b91b59ae724920a9eb0c006c5b318ac393fa1bc8974edd9887610f46565b508315610d9357845460ff60401b19168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b505050505050505050505050565b610daa82610860565b610db381610f3c565b61089e8383610feb565b6001600160a01b0391821660009081527f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace016020908152604080832093909416825291909152205490565b7f715bacafb7a853b9b91b59ae724920a9eb0c006c5b318ac393fa1bc8974edd98610e3181610f3c565b6108d783836110fd565b6000805160206120a48339815191525460ff1615610e6c5760405163d93c066560e01b815260040160405180910390fd5b565b600033610b7d8185856114ce565b6000610e888484610dbd565b905060001981101561089e5781811015610ece57604051637dc7a0d960e11b81526001600160a01b038416600482015260248101829052604481018390526064016107e3565b61089e848484840360006114d7565b6001600160a01b038316610f0757604051634b637e8f60e11b8152600060048201526024016107e3565b6001600160a01b038216610f315760405163ec442f0560e01b8152600060048201526024016107e3565b6108d78383836115bf565b6108fc8133611663565b6000600080516020612084833981519152610f618484610ad7565b610fe1576000848152602082815260408083206001600160a01b03871684529091529020805460ff19166001179055610f973390565b6001600160a01b0316836001600160a01b0316857f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a460019150506106e9565b60009150506106e9565b60006000805160206120848339815191526110068484610ad7565b15610fe1576000848152602082815260408083206001600160a01b0387168085529252808320805460ff1916905551339287917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a460019150506106e9565b61106f61169c565b6000805160206120a4833981519152805460ff191681557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa335b6040516001600160a01b03909116815260200160405180910390a150565b6001600160a01b0382166110f15760405163ec442f0560e01b8152600060048201526024016107e3565b610967600083836115bf565b6001600160a01b03821661112757604051634b637e8f60e11b8152600060048201526024016107e3565b610967826000836115bf565b306001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614806111ba57507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166111ae600080516020612044833981519152546001600160a01b031690565b6001600160a01b031614155b15610e6c5760405163703e46dd60e11b815260040160405180910390fd5b7fa615a8afb6fffcb8c6809ac0997b5c9c12b8cc97651150f14c8f6203168cff4c61096781610f3c565b816001600160a01b03166352d1902d6040518163ffffffff1660e01b8152600401602060405180830381865afa92505050801561125c575060408051601f3d908101601f1916820190925261125991810190611ec7565b60015b61128457604051634c9c8ce360e01b81526001600160a01b03831660048201526024016107e3565b60008051602061204483398151915281146112b557604051632a87526960e21b8152600481018290526024016107e3565b6108d783836116cc565b306001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614610e6c5760405163703e46dd60e11b815260040160405180910390fd5b611310610e3b565b6000805160206120a4833981519152805460ff191660011781557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258336110a9565b8061135b81610b87565b156113845760405163c8c29b9960e01b81526001600160a01b03821660048201526024016107e3565b6001600160a01b03821660008181527f345cc2404af916c3db112e7a6103770647a90ed78a5d681e21dc2e11742329006020818152604092839020805460ff191660011790559151928352917fae7f60c1b8f645c3beffeb531169cbc446874bbf247698325318879ac850c34691015b60405180910390a1505050565b8061140b81610b87565b61143357604051638d542b2960e01b81526001600160a01b03821660048201526024016107e3565b6001600160a01b03821660008181527f345cc2404af916c3db112e7a6103770647a90ed78a5d681e21dc2e11742329006020818152604092839020805460ff191690559151928352917f0c18efbde61ac471ead6960a3f1097735c68ecdb685ae8e2a108c28385399a6591016113f4565b6114ac611722565b610967828261176b565b610e6c611722565b6114c6611722565b610e6c6117bc565b6108d783838360015b6000805160206120248339815191526001600160a01b0385166115105760405163e602df0560e01b8152600060048201526024016107e3565b6001600160a01b03841661153a57604051634a1406b160e11b8152600060048201526024016107e3565b6001600160a01b038086166000908152600183016020908152604080832093881683529290522083905581156115b857836001600160a01b0316856001600160a01b03167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925856040516115af91815260200190565b60405180910390a35b5050505050565b826115c981610b87565b156115f25760405163c8c29b9960e01b81526001600160a01b03821660048201526024016107e3565b826115fc81610b87565b156116255760405163c8c29b9960e01b81526001600160a01b03821660048201526024016107e3565b3361162f81610b87565b156116585760405163c8c29b9960e01b81526001600160a01b03821660048201526024016107e3565b610acf8686866117dd565b61166d8282610ad7565b6109675760405163e2517d3f60e01b81526001600160a01b0382166004820152602481018390526044016107e3565b6000805160206120a48339815191525460ff16610e6c57604051638dfc202b60e01b815260040160405180910390fd5b6116d5826117f0565b6040516001600160a01b038316907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b90600090a280511561171a576108d78282611855565b6109676118c2565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a0054600160401b900460ff16610e6c57604051631afcd79f60e31b815260040160405180910390fd5b611773611722565b6000805160206120248339815191527f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace036117ad8482611f27565b506004810161089e8382611f27565b6117c4611722565b6000805160206120a4833981519152805460ff19169055565b6117e5610e3b565b6108d78383836118e1565b806001600160a01b03163b60000361182657604051634c9c8ce360e01b81526001600160a01b03821660048201526024016107e3565b60008051602061204483398151915280546001600160a01b0319166001600160a01b0392909216919091179055565b6060600080846001600160a01b0316846040516118729190611fe6565b600060405180830381855af49150503d80600081146118ad576040519150601f19603f3d011682016040523d82523d6000602084013e6118b2565b606091505b5091509150610831858383611a1f565b3415610e6c5760405163b398979f60e01b815260040160405180910390fd5b6000805160206120248339815191526001600160a01b03841661191d57818160020160008282546119129190612002565b9091555061198f9050565b6001600160a01b038416600090815260208290526040902054828110156119705760405163391434e360e21b81526001600160a01b038616600482015260248101829052604481018490526064016107e3565b6001600160a01b03851660009081526020839052604090209083900390555b6001600160a01b0383166119ad5760028101805483900390556119cc565b6001600160a01b03831660009081526020829052604090208054830190555b826001600160a01b0316846001600160a01b03167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef84604051611a1191815260200190565b60405180910390a350505050565b606082611a3457611a2f82611a7b565b610859565b8151158015611a4b57506001600160a01b0384163b155b15611a7457604051639996b31560e01b81526001600160a01b03851660048201526024016107e3565b5080610859565b805115611a8b5780518082602001fd5b60405163d6bda27560e01b815260040160405180910390fd5b600060208284031215611ab657600080fd5b81356001600160e01b03198116811461085957600080fd5b60005b83811015611ae9578181015183820152602001611ad1565b50506000910152565b6020815260008251806020840152611b11816040850160208701611ace565b601f01601f19169190910160400192915050565b80356001600160a01b0381168114611b3c57600080fd5b919050565b60008060408385031215611b5457600080fd5b611b5d83611b25565b946020939093013593505050565b600080600060608486031215611b8057600080fd5b611b8984611b25565b9250611b9760208501611b25565b929592945050506040919091013590565b600060208284031215611bba57600080fd5b5035919050565b60008060408385031215611bd457600080fd5b82359150611be460208401611b25565b90509250929050565b634e487b7160e01b600052604160045260246000fd5b60008067ffffffffffffffff841115611c1e57611c1e611bed565b50604051601f19601f85018116603f0116810181811067ffffffffffffffff82111715611c4d57611c4d611bed565b604052838152905080828401851015611c6557600080fd5b83836020830137600060208583010152509392505050565b60008060408385031215611c9057600080fd5b611c9983611b25565b9150602083013567ffffffffffffffff811115611cb557600080fd5b8301601f81018513611cc657600080fd5b611cd585823560208401611c03565b9150509250929050565b600060208284031215611cf157600080fd5b61085982611b25565b60008060208385031215611d0d57600080fd5b823567ffffffffffffffff811115611d2457600080fd5b8301601f81018513611d3557600080fd5b803567ffffffffffffffff811115611d4c57600080fd5b8560208260051b8401011115611d6157600080fd5b6020919091019590945092505050565b600082601f830112611d8257600080fd5b61085983833560208501611c03565b600080600080600080600060e0888a031215611dac57600080fd5b873567ffffffffffffffff811115611dc357600080fd5b611dcf8a828b01611d71565b975050602088013567ffffffffffffffff811115611dec57600080fd5b611df88a828b01611d71565b965050611e0760408901611b25565b9450611e1560608901611b25565b9350611e2360808901611b25565b9250611e3160a08901611b25565b9150611e3f60c08901611b25565b905092959891949750929550565b60008060408385031215611e6057600080fd5b611e6983611b25565b9150611be460208401611b25565b600181811c90821680611e8b57607f821691505b602082108103611eab57634e487b7160e01b600052602260045260246000fd5b50919050565b634e487b7160e01b600052603260045260246000fd5b600060208284031215611ed957600080fd5b5051919050565b601f8211156108d757806000526020600020601f840160051c81016020851015611f075750805b601f840160051c820191505b818110156115b85760008155600101611f13565b815167ffffffffffffffff811115611f4157611f41611bed565b611f5581611f4f8454611e77565b84611ee0565b6020601f821160018114611f895760008315611f715750848201515b600019600385901b1c1916600184901b1784556115b8565b600084815260208120601f198516915b82811015611fb95787850151825560209485019460019092019101611f99565b5084821015611fd75786840151600019600387901b60f8161c191681555b50505050600190811b01905550565b60008251611ff8818460208701611ace565b9190910192915050565b808201808211156106e957634e487b7160e01b600052601160045260246000fdfe52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace00360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc539440820030c4994db4e31b6b800deafd503688728f932addfe7a410515c14c02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800cd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f03300a2646970667358221220e000acc2f5b8562207be2104fbf03a8136740f71c21a723e26d412b7ee7b89bf64736f6c634300081a0033",
}

// StablecoinUpgradeableABI is the input ABI used to generate the binding from.
// Deprecated: Use StablecoinUpgradeableMetaData.ABI instead.
var StablecoinUpgradeableABI = StablecoinUpgradeableMetaData.ABI

// StablecoinUpgradeableBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use StablecoinUpgradeableMetaData.Bin instead.
var StablecoinUpgradeableBin = StablecoinUpgradeableMetaData.Bin

// DeployStablecoinUpgradeable deploys a new Ethereum contract, binding an instance of StablecoinUpgradeable to it.
func DeployStablecoinUpgradeable(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *StablecoinUpgradeable, error) {
	parsed, err := StablecoinUpgradeableMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(StablecoinUpgradeableBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &StablecoinUpgradeable{StablecoinUpgradeableCaller: StablecoinUpgradeableCaller{contract: contract}, StablecoinUpgradeableTransactor: StablecoinUpgradeableTransactor{contract: contract}, StablecoinUpgradeableFilterer: StablecoinUpgradeableFilterer{contract: contract}}, nil
}

// StablecoinUpgradeable is an auto generated Go binding around an Ethereum contract.
type StablecoinUpgradeable struct {
	StablecoinUpgradeableCaller     // Read-only binding to the contract
	StablecoinUpgradeableTransactor // Write-only binding to the contract
	StablecoinUpgradeableFilterer   // Log filterer for contract events
}

// StablecoinUpgradeableCaller is an auto generated read-only Go binding around an Ethereum contract.
type StablecoinUpgradeableCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StablecoinUpgradeableTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StablecoinUpgradeableTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StablecoinUpgradeableFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StablecoinUpgradeableFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StablecoinUpgradeableSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StablecoinUpgradeableSession struct {
	Contract     *StablecoinUpgradeable // Generic contract binding to set the session for
	CallOpts     bind.CallOpts          // Call options to use throughout this session
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// StablecoinUpgradeableCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StablecoinUpgradeableCallerSession struct {
	Contract *StablecoinUpgradeableCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                // Call options to use throughout this session
}

// StablecoinUpgradeableTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StablecoinUpgradeableTransactorSession struct {
	Contract     *StablecoinUpgradeableTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                // Transaction auth options to use throughout this session
}

// StablecoinUpgradeableRaw is an auto generated low-level Go binding around an Ethereum contract.
type StablecoinUpgradeableRaw struct {
	Contract *StablecoinUpgradeable // Generic contract binding to access the raw methods on
}

// StablecoinUpgradeableCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StablecoinUpgradeableCallerRaw struct {
	Contract *StablecoinUpgradeableCaller // Generic read-only contract binding to access the raw methods on
}

// StablecoinUpgradeableTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StablecoinUpgradeableTransactorRaw struct {
	Contract *StablecoinUpgradeableTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStablecoinUpgradeable creates a new instance of StablecoinUpgradeable, bound to a specific deployed contract.
func NewStablecoinUpgradeable(address common.Address, backend bind.ContractBackend) (*StablecoinUpgradeable, error) {
	contract, err := bindStablecoinUpgradeable(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StablecoinUpgradeable{StablecoinUpgradeableCaller: StablecoinUpgradeableCaller{contract: contract}, StablecoinUpgradeableTransactor: StablecoinUpgradeableTransactor{contract: contract}, StablecoinUpgradeableFilterer: StablecoinUpgradeableFilterer{contract: contract}}, nil
}

// NewStablecoinUpgradeableCaller creates a new read-only instance of StablecoinUpgradeable, bound to a specific deployed contract.
func NewStablecoinUpgradeableCaller(address common.Address, caller bind.ContractCaller) (*StablecoinUpgradeableCaller, error) {
	contract, err := bindStablecoinUpgradeable(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StablecoinUpgradeableCaller{contract: contract}, nil
}

// NewStablecoinUpgradeableTransactor creates a new write-only instance of StablecoinUpgradeable, bound to a specific deployed contract.
func NewStablecoinUpgradeableTransactor(address common.Address, transactor bind.ContractTransactor) (*StablecoinUpgradeableTransactor, error) {
	contract, err := bindStablecoinUpgradeable(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StablecoinUpgradeableTransactor{contract: contract}, nil
}

// NewStablecoinUpgradeableFilterer creates a new log filterer instance of StablecoinUpgradeable, bound to a specific deployed contract.
func NewStablecoinUpgradeableFilterer(address common.Address, filterer bind.ContractFilterer) (*StablecoinUpgradeableFilterer, error) {
	contract, err := bindStablecoinUpgradeable(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StablecoinUpgradeableFilterer{contract: contract}, nil
}

// bindStablecoinUpgradeable binds a generic wrapper to an already deployed contract.
func bindStablecoinUpgradeable(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := StablecoinUpgradeableMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StablecoinUpgradeable *StablecoinUpgradeableRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StablecoinUpgradeable.Contract.StablecoinUpgradeableCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StablecoinUpgradeable *StablecoinUpgradeableRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StablecoinUpgradeable.Contract.StablecoinUpgradeableTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StablecoinUpgradeable *StablecoinUpgradeableRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StablecoinUpgradeable.Contract.StablecoinUpgradeableTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StablecoinUpgradeable *StablecoinUpgradeableCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StablecoinUpgradeable.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StablecoinUpgradeable *StablecoinUpgradeableTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StablecoinUpgradeable.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StablecoinUpgradeable *StablecoinUpgradeableTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StablecoinUpgradeable.Contract.contract.Transact(opts, method, params...)
}

// BURNERROLE is a free data retrieval call binding the contract method 0x282c51f3.
//
// Solidity: function BURNER_ROLE() view returns(bytes32)
func (_StablecoinUpgradeable *StablecoinUpgradeableCaller) BURNERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _StablecoinUpgradeable.contract.Call(opts, &out, "BURNER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// BURNERROLE is a free data retrieval call binding the contract method 0x282c51f3.
//
// Solidity: function BURNER_ROLE() view returns(bytes32)
func (_StablecoinUpgradeable *StablecoinUpgradeableSession) BURNERROLE() ([32]byte, error) {
	return _StablecoinUpgradeable.Contract.BURNERROLE(&_StablecoinUpgradeable.CallOpts)
}

// BURNERROLE is a free data retrieval call binding the contract method 0x282c51f3.
//
// Solidity: function BURNER_ROLE() view returns(bytes32)
func (_StablecoinUpgradeable *StablecoinUpgradeableCallerSession) BURNERROLE() ([32]byte, error) {
	return _StablecoinUpgradeable.Contract.BURNERROLE(&_StablecoinUpgradeable.CallOpts)
}

// CLAWBACKERROLE is a free data retrieval call binding the contract method 0x475ca324.
//
// Solidity: function CLAWBACKER_ROLE() view returns(bytes32)
func (_StablecoinUpgradeable *StablecoinUpgradeableCaller) CLAWBACKERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _StablecoinUpgradeable.contract.Call(opts, &out, "CLAWBACKER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CLAWBACKERROLE is a free data retrieval call binding the contract method 0x475ca324.
//
// Solidity: function CLAWBACKER_ROLE() view returns(bytes32)
func (_StablecoinUpgradeable *StablecoinUpgradeableSession) CLAWBACKERROLE() ([32]byte, error) {
	return _StablecoinUpgradeable.Contract.CLAWBACKERROLE(&_StablecoinUpgradeable.CallOpts)
}

// CLAWBACKERROLE is a free data retrieval call binding the contract method 0x475ca324.
//
// Solidity: function CLAWBACKER_ROLE() view returns(bytes32)
func (_StablecoinUpgradeable *StablecoinUpgradeableCallerSession) CLAWBACKERROLE() ([32]byte, error) {
	return _StablecoinUpgradeable.Contract.CLAWBACKERROLE(&_StablecoinUpgradeable.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_StablecoinUpgradeable *StablecoinUpgradeableCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _StablecoinUpgradeable.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_StablecoinUpgradeable *StablecoinUpgradeableSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _StablecoinUpgradeable.Contract.DEFAULTADMINROLE(&_StablecoinUpgradeable.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_StablecoinUpgradeable *StablecoinUpgradeableCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _StablecoinUpgradeable.Contract.DEFAULTADMINROLE(&_StablecoinUpgradeable.CallOpts)
}

// MINTERROLE is a free data retrieval call binding the contract method 0xd5391393.
//
// Solidity: function MINTER_ROLE() view returns(bytes32)
func (_StablecoinUpgradeable *StablecoinUpgradeableCaller) MINTERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _StablecoinUpgradeable.contract.Call(opts, &out, "MINTER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// MINTERROLE is a free data retrieval call binding the contract method 0xd5391393.
//
// Solidity: function MINTER_ROLE() view returns(bytes32)
func (_StablecoinUpgradeable *StablecoinUpgradeableSession) MINTERROLE() ([32]byte, error) {
	return _StablecoinUpgradeable.Contract.MINTERROLE(&_StablecoinUpgradeable.CallOpts)
}

// MINTERROLE is a free data retrieval call binding the contract method 0xd5391393.
//
// Solidity: function MINTER_ROLE() view returns(bytes32)
func (_StablecoinUpgradeable *StablecoinUpgradeableCallerSession) MINTERROLE() ([32]byte, error) {
	return _StablecoinUpgradeable.Contract.MINTERROLE(&_StablecoinUpgradeable.CallOpts)
}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_StablecoinUpgradeable *StablecoinUpgradeableCaller) PAUSERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _StablecoinUpgradeable.contract.Call(opts, &out, "PAUSER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_StablecoinUpgradeable *StablecoinUpgradeableSession) PAUSERROLE() ([32]byte, error) {
	return _StablecoinUpgradeable.Contract.PAUSERROLE(&_StablecoinUpgradeable.CallOpts)
}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_StablecoinUpgradeable *StablecoinUpgradeableCallerSession) PAUSERROLE() ([32]byte, error) {
	return _StablecoinUpgradeable.Contract.PAUSERROLE(&_StablecoinUpgradeable.CallOpts)
}

// UPGRADERROLE is a free data retrieval call binding the contract method 0xf72c0d8b.
//
// Solidity: function UPGRADER_ROLE() view returns(bytes32)
func (_StablecoinUpgradeable *StablecoinUpgradeableCaller) UPGRADERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _StablecoinUpgradeable.contract.Call(opts, &out, "UPGRADER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// UPGRADERROLE is a free data retrieval call binding the contract method 0xf72c0d8b.
//
// Solidity: function UPGRADER_ROLE() view returns(bytes32)
func (_StablecoinUpgradeable *StablecoinUpgradeableSession) UPGRADERROLE() ([32]byte, error) {
	return _StablecoinUpgradeable.Contract.UPGRADERROLE(&_StablecoinUpgradeable.CallOpts)
}

// UPGRADERROLE is a free data retrieval call binding the contract method 0xf72c0d8b.
//
// Solidity: function UPGRADER_ROLE() view returns(bytes32)
func (_StablecoinUpgradeable *StablecoinUpgradeableCallerSession) UPGRADERROLE() ([32]byte, error) {
	return _StablecoinUpgradeable.Contract.UPGRADERROLE(&_StablecoinUpgradeable.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_StablecoinUpgradeable *StablecoinUpgradeableCaller) UPGRADEINTERFACEVERSION(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _StablecoinUpgradeable.contract.Call(opts, &out, "UPGRADE_INTERFACE_VERSION")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_StablecoinUpgradeable *StablecoinUpgradeableSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _StablecoinUpgradeable.Contract.UPGRADEINTERFACEVERSION(&_StablecoinUpgradeable.CallOpts)
}

// UPGRADEINTERFACEVERSION is a free data retrieval call binding the contract method 0xad3cb1cc.
//
// Solidity: function UPGRADE_INTERFACE_VERSION() view returns(string)
func (_StablecoinUpgradeable *StablecoinUpgradeableCallerSession) UPGRADEINTERFACEVERSION() (string, error) {
	return _StablecoinUpgradeable.Contract.UPGRADEINTERFACEVERSION(&_StablecoinUpgradeable.CallOpts)
}

// AccountPaused is a free data retrieval call binding the contract method 0xbc8c4b4f.
//
// Solidity: function accountPaused(address account) view returns(bool)
func (_StablecoinUpgradeable *StablecoinUpgradeableCaller) AccountPaused(opts *bind.CallOpts, account common.Address) (bool, error) {
	var out []interface{}
	err := _StablecoinUpgradeable.contract.Call(opts, &out, "accountPaused", account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AccountPaused is a free data retrieval call binding the contract method 0xbc8c4b4f.
//
// Solidity: function accountPaused(address account) view returns(bool)
func (_StablecoinUpgradeable *StablecoinUpgradeableSession) AccountPaused(account common.Address) (bool, error) {
	return _StablecoinUpgradeable.Contract.AccountPaused(&_StablecoinUpgradeable.CallOpts, account)
}

// AccountPaused is a free data retrieval call binding the contract method 0xbc8c4b4f.
//
// Solidity: function accountPaused(address account) view returns(bool)
func (_StablecoinUpgradeable *StablecoinUpgradeableCallerSession) AccountPaused(account common.Address) (bool, error) {
	return _StablecoinUpgradeable.Contract.AccountPaused(&_StablecoinUpgradeable.CallOpts, account)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_StablecoinUpgradeable *StablecoinUpgradeableCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _StablecoinUpgradeable.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_StablecoinUpgradeable *StablecoinUpgradeableSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _StablecoinUpgradeable.Contract.Allowance(&_StablecoinUpgradeable.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256)
func (_StablecoinUpgradeable *StablecoinUpgradeableCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _StablecoinUpgradeable.Contract.Allowance(&_StablecoinUpgradeable.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_StablecoinUpgradeable *StablecoinUpgradeableCaller) BalanceOf(opts *bind.CallOpts, account common.Address) (*big.Int, error) {
	var out []interface{}
	err := _StablecoinUpgradeable.contract.Call(opts, &out, "balanceOf", account)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_StablecoinUpgradeable *StablecoinUpgradeableSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _StablecoinUpgradeable.Contract.BalanceOf(&_StablecoinUpgradeable.CallOpts, account)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address account) view returns(uint256)
func (_StablecoinUpgradeable *StablecoinUpgradeableCallerSession) BalanceOf(account common.Address) (*big.Int, error) {
	return _StablecoinUpgradeable.Contract.BalanceOf(&_StablecoinUpgradeable.CallOpts, account)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_StablecoinUpgradeable *StablecoinUpgradeableCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _StablecoinUpgradeable.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_StablecoinUpgradeable *StablecoinUpgradeableSession) Decimals() (uint8, error) {
	return _StablecoinUpgradeable.Contract.Decimals(&_StablecoinUpgradeable.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_StablecoinUpgradeable *StablecoinUpgradeableCallerSession) Decimals() (uint8, error) {
	return _StablecoinUpgradeable.Contract.Decimals(&_StablecoinUpgradeable.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_StablecoinUpgradeable *StablecoinUpgradeableCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _StablecoinUpgradeable.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_StablecoinUpgradeable *StablecoinUpgradeableSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _StablecoinUpgradeable.Contract.GetRoleAdmin(&_StablecoinUpgradeable.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_StablecoinUpgradeable *StablecoinUpgradeableCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _StablecoinUpgradeable.Contract.GetRoleAdmin(&_StablecoinUpgradeable.CallOpts, role)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_StablecoinUpgradeable *StablecoinUpgradeableCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _StablecoinUpgradeable.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_StablecoinUpgradeable *StablecoinUpgradeableSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _StablecoinUpgradeable.Contract.HasRole(&_StablecoinUpgradeable.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_StablecoinUpgradeable *StablecoinUpgradeableCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _StablecoinUpgradeable.Contract.HasRole(&_StablecoinUpgradeable.CallOpts, role, account)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_StablecoinUpgradeable *StablecoinUpgradeableCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _StablecoinUpgradeable.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_StablecoinUpgradeable *StablecoinUpgradeableSession) Name() (string, error) {
	return _StablecoinUpgradeable.Contract.Name(&_StablecoinUpgradeable.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_StablecoinUpgradeable *StablecoinUpgradeableCallerSession) Name() (string, error) {
	return _StablecoinUpgradeable.Contract.Name(&_StablecoinUpgradeable.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_StablecoinUpgradeable *StablecoinUpgradeableCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _StablecoinUpgradeable.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_StablecoinUpgradeable *StablecoinUpgradeableSession) Paused() (bool, error) {
	return _StablecoinUpgradeable.Contract.Paused(&_StablecoinUpgradeable.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_StablecoinUpgradeable *StablecoinUpgradeableCallerSession) Paused() (bool, error) {
	return _StablecoinUpgradeable.Contract.Paused(&_StablecoinUpgradeable.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_StablecoinUpgradeable *StablecoinUpgradeableCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _StablecoinUpgradeable.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_StablecoinUpgradeable *StablecoinUpgradeableSession) ProxiableUUID() ([32]byte, error) {
	return _StablecoinUpgradeable.Contract.ProxiableUUID(&_StablecoinUpgradeable.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_StablecoinUpgradeable *StablecoinUpgradeableCallerSession) ProxiableUUID() ([32]byte, error) {
	return _StablecoinUpgradeable.Contract.ProxiableUUID(&_StablecoinUpgradeable.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_StablecoinUpgradeable *StablecoinUpgradeableCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _StablecoinUpgradeable.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_StablecoinUpgradeable *StablecoinUpgradeableSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _StablecoinUpgradeable.Contract.SupportsInterface(&_StablecoinUpgradeable.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_StablecoinUpgradeable *StablecoinUpgradeableCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _StablecoinUpgradeable.Contract.SupportsInterface(&_StablecoinUpgradeable.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_StablecoinUpgradeable *StablecoinUpgradeableCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _StablecoinUpgradeable.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_StablecoinUpgradeable *StablecoinUpgradeableSession) Symbol() (string, error) {
	return _StablecoinUpgradeable.Contract.Symbol(&_StablecoinUpgradeable.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_StablecoinUpgradeable *StablecoinUpgradeableCallerSession) Symbol() (string, error) {
	return _StablecoinUpgradeable.Contract.Symbol(&_StablecoinUpgradeable.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_StablecoinUpgradeable *StablecoinUpgradeableCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StablecoinUpgradeable.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_StablecoinUpgradeable *StablecoinUpgradeableSession) TotalSupply() (*big.Int, error) {
	return _StablecoinUpgradeable.Contract.TotalSupply(&_StablecoinUpgradeable.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_StablecoinUpgradeable *StablecoinUpgradeableCallerSession) TotalSupply() (*big.Int, error) {
	return _StablecoinUpgradeable.Contract.TotalSupply(&_StablecoinUpgradeable.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_StablecoinUpgradeable *StablecoinUpgradeableTransactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _StablecoinUpgradeable.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_StablecoinUpgradeable *StablecoinUpgradeableSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _StablecoinUpgradeable.Contract.Approve(&_StablecoinUpgradeable.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_StablecoinUpgradeable *StablecoinUpgradeableTransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _StablecoinUpgradeable.Contract.Approve(&_StablecoinUpgradeable.TransactOpts, spender, value)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 value) returns()
func (_StablecoinUpgradeable *StablecoinUpgradeableTransactor) Burn(opts *bind.TransactOpts, value *big.Int) (*types.Transaction, error) {
	return _StablecoinUpgradeable.contract.Transact(opts, "burn", value)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 value) returns()
func (_StablecoinUpgradeable *StablecoinUpgradeableSession) Burn(value *big.Int) (*types.Transaction, error) {
	return _StablecoinUpgradeable.Contract.Burn(&_StablecoinUpgradeable.TransactOpts, value)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 value) returns()
func (_StablecoinUpgradeable *StablecoinUpgradeableTransactorSession) Burn(value *big.Int) (*types.Transaction, error) {
	return _StablecoinUpgradeable.Contract.Burn(&_StablecoinUpgradeable.TransactOpts, value)
}

// Clawback is a paid mutator transaction binding the contract method 0xf3df317e.
//
// Solidity: function clawback(address from, uint256 value) returns()
func (_StablecoinUpgradeable *StablecoinUpgradeableTransactor) Clawback(opts *bind.TransactOpts, from common.Address, value *big.Int) (*types.Transaction, error) {
	return _StablecoinUpgradeable.contract.Transact(opts, "clawback", from, value)
}

// Clawback is a paid mutator transaction binding the contract method 0xf3df317e.
//
// Solidity: function clawback(address from, uint256 value) returns()
func (_StablecoinUpgradeable *StablecoinUpgradeableSession) Clawback(from common.Address, value *big.Int) (*types.Transaction, error) {
	return _StablecoinUpgradeable.Contract.Clawback(&_StablecoinUpgradeable.TransactOpts, from, value)
}

// Clawback is a paid mutator transaction binding the contract method 0xf3df317e.
//
// Solidity: function clawback(address from, uint256 value) returns()
func (_StablecoinUpgradeable *StablecoinUpgradeableTransactorSession) Clawback(from common.Address, value *big.Int) (*types.Transaction, error) {
	return _StablecoinUpgradeable.Contract.Clawback(&_StablecoinUpgradeable.TransactOpts, from, value)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_StablecoinUpgradeable *StablecoinUpgradeableTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _StablecoinUpgradeable.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_StablecoinUpgradeable *StablecoinUpgradeableSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _StablecoinUpgradeable.Contract.GrantRole(&_StablecoinUpgradeable.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_StablecoinUpgradeable *StablecoinUpgradeableTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _StablecoinUpgradeable.Contract.GrantRole(&_StablecoinUpgradeable.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0xc91f0c53.
//
// Solidity: function initialize(string name_, string symbol_, address minter_, address admin_, address upgrader_, address pauser_, address clawbacker_) returns()
func (_StablecoinUpgradeable *StablecoinUpgradeableTransactor) Initialize(opts *bind.TransactOpts, name_ string, symbol_ string, minter_ common.Address, admin_ common.Address, upgrader_ common.Address, pauser_ common.Address, clawbacker_ common.Address) (*types.Transaction, error) {
	return _StablecoinUpgradeable.contract.Transact(opts, "initialize", name_, symbol_, minter_, admin_, upgrader_, pauser_, clawbacker_)
}

// Initialize is a paid mutator transaction binding the contract method 0xc91f0c53.
//
// Solidity: function initialize(string name_, string symbol_, address minter_, address admin_, address upgrader_, address pauser_, address clawbacker_) returns()
func (_StablecoinUpgradeable *StablecoinUpgradeableSession) Initialize(name_ string, symbol_ string, minter_ common.Address, admin_ common.Address, upgrader_ common.Address, pauser_ common.Address, clawbacker_ common.Address) (*types.Transaction, error) {
	return _StablecoinUpgradeable.Contract.Initialize(&_StablecoinUpgradeable.TransactOpts, name_, symbol_, minter_, admin_, upgrader_, pauser_, clawbacker_)
}

// Initialize is a paid mutator transaction binding the contract method 0xc91f0c53.
//
// Solidity: function initialize(string name_, string symbol_, address minter_, address admin_, address upgrader_, address pauser_, address clawbacker_) returns()
func (_StablecoinUpgradeable *StablecoinUpgradeableTransactorSession) Initialize(name_ string, symbol_ string, minter_ common.Address, admin_ common.Address, upgrader_ common.Address, pauser_ common.Address, clawbacker_ common.Address) (*types.Transaction, error) {
	return _StablecoinUpgradeable.Contract.Initialize(&_StablecoinUpgradeable.TransactOpts, name_, symbol_, minter_, admin_, upgrader_, pauser_, clawbacker_)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 value) returns()
func (_StablecoinUpgradeable *StablecoinUpgradeableTransactor) Mint(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _StablecoinUpgradeable.contract.Transact(opts, "mint", to, value)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 value) returns()
func (_StablecoinUpgradeable *StablecoinUpgradeableSession) Mint(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _StablecoinUpgradeable.Contract.Mint(&_StablecoinUpgradeable.TransactOpts, to, value)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(address to, uint256 value) returns()
func (_StablecoinUpgradeable *StablecoinUpgradeableTransactorSession) Mint(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _StablecoinUpgradeable.Contract.Mint(&_StablecoinUpgradeable.TransactOpts, to, value)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_StablecoinUpgradeable *StablecoinUpgradeableTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StablecoinUpgradeable.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_StablecoinUpgradeable *StablecoinUpgradeableSession) Pause() (*types.Transaction, error) {
	return _StablecoinUpgradeable.Contract.Pause(&_StablecoinUpgradeable.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_StablecoinUpgradeable *StablecoinUpgradeableTransactorSession) Pause() (*types.Transaction, error) {
	return _StablecoinUpgradeable.Contract.Pause(&_StablecoinUpgradeable.TransactOpts)
}

// PauseAccounts is a paid mutator transaction binding the contract method 0x917b1ace.
//
// Solidity: function pauseAccounts(address[] accounts) returns()
func (_StablecoinUpgradeable *StablecoinUpgradeableTransactor) PauseAccounts(opts *bind.TransactOpts, accounts []common.Address) (*types.Transaction, error) {
	return _StablecoinUpgradeable.contract.Transact(opts, "pauseAccounts", accounts)
}

// PauseAccounts is a paid mutator transaction binding the contract method 0x917b1ace.
//
// Solidity: function pauseAccounts(address[] accounts) returns()
func (_StablecoinUpgradeable *StablecoinUpgradeableSession) PauseAccounts(accounts []common.Address) (*types.Transaction, error) {
	return _StablecoinUpgradeable.Contract.PauseAccounts(&_StablecoinUpgradeable.TransactOpts, accounts)
}

// PauseAccounts is a paid mutator transaction binding the contract method 0x917b1ace.
//
// Solidity: function pauseAccounts(address[] accounts) returns()
func (_StablecoinUpgradeable *StablecoinUpgradeableTransactorSession) PauseAccounts(accounts []common.Address) (*types.Transaction, error) {
	return _StablecoinUpgradeable.Contract.PauseAccounts(&_StablecoinUpgradeable.TransactOpts, accounts)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_StablecoinUpgradeable *StablecoinUpgradeableTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _StablecoinUpgradeable.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_StablecoinUpgradeable *StablecoinUpgradeableSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _StablecoinUpgradeable.Contract.RenounceRole(&_StablecoinUpgradeable.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_StablecoinUpgradeable *StablecoinUpgradeableTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _StablecoinUpgradeable.Contract.RenounceRole(&_StablecoinUpgradeable.TransactOpts, role, callerConfirmation)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_StablecoinUpgradeable *StablecoinUpgradeableTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _StablecoinUpgradeable.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_StablecoinUpgradeable *StablecoinUpgradeableSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _StablecoinUpgradeable.Contract.RevokeRole(&_StablecoinUpgradeable.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_StablecoinUpgradeable *StablecoinUpgradeableTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _StablecoinUpgradeable.Contract.RevokeRole(&_StablecoinUpgradeable.TransactOpts, role, account)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_StablecoinUpgradeable *StablecoinUpgradeableTransactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _StablecoinUpgradeable.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_StablecoinUpgradeable *StablecoinUpgradeableSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _StablecoinUpgradeable.Contract.Transfer(&_StablecoinUpgradeable.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_StablecoinUpgradeable *StablecoinUpgradeableTransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _StablecoinUpgradeable.Contract.Transfer(&_StablecoinUpgradeable.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_StablecoinUpgradeable *StablecoinUpgradeableTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _StablecoinUpgradeable.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_StablecoinUpgradeable *StablecoinUpgradeableSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _StablecoinUpgradeable.Contract.TransferFrom(&_StablecoinUpgradeable.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_StablecoinUpgradeable *StablecoinUpgradeableTransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _StablecoinUpgradeable.Contract.TransferFrom(&_StablecoinUpgradeable.TransactOpts, from, to, value)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_StablecoinUpgradeable *StablecoinUpgradeableTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StablecoinUpgradeable.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_StablecoinUpgradeable *StablecoinUpgradeableSession) Unpause() (*types.Transaction, error) {
	return _StablecoinUpgradeable.Contract.Unpause(&_StablecoinUpgradeable.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_StablecoinUpgradeable *StablecoinUpgradeableTransactorSession) Unpause() (*types.Transaction, error) {
	return _StablecoinUpgradeable.Contract.Unpause(&_StablecoinUpgradeable.TransactOpts)
}

// UnpauseAccount is a paid mutator transaction binding the contract method 0x94408b9a.
//
// Solidity: function unpauseAccount(address account) returns()
func (_StablecoinUpgradeable *StablecoinUpgradeableTransactor) UnpauseAccount(opts *bind.TransactOpts, account common.Address) (*types.Transaction, error) {
	return _StablecoinUpgradeable.contract.Transact(opts, "unpauseAccount", account)
}

// UnpauseAccount is a paid mutator transaction binding the contract method 0x94408b9a.
//
// Solidity: function unpauseAccount(address account) returns()
func (_StablecoinUpgradeable *StablecoinUpgradeableSession) UnpauseAccount(account common.Address) (*types.Transaction, error) {
	return _StablecoinUpgradeable.Contract.UnpauseAccount(&_StablecoinUpgradeable.TransactOpts, account)
}

// UnpauseAccount is a paid mutator transaction binding the contract method 0x94408b9a.
//
// Solidity: function unpauseAccount(address account) returns()
func (_StablecoinUpgradeable *StablecoinUpgradeableTransactorSession) UnpauseAccount(account common.Address) (*types.Transaction, error) {
	return _StablecoinUpgradeable.Contract.UnpauseAccount(&_StablecoinUpgradeable.TransactOpts, account)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_StablecoinUpgradeable *StablecoinUpgradeableTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _StablecoinUpgradeable.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_StablecoinUpgradeable *StablecoinUpgradeableSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _StablecoinUpgradeable.Contract.UpgradeToAndCall(&_StablecoinUpgradeable.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_StablecoinUpgradeable *StablecoinUpgradeableTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _StablecoinUpgradeable.Contract.UpgradeToAndCall(&_StablecoinUpgradeable.TransactOpts, newImplementation, data)
}

// StablecoinUpgradeableAccountPausedIterator is returned from FilterAccountPaused and is used to iterate over the raw logs and unpacked data for AccountPaused events raised by the StablecoinUpgradeable contract.
type StablecoinUpgradeableAccountPausedIterator struct {
	Event *StablecoinUpgradeableAccountPaused // Event containing the contract specifics and raw log

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
func (it *StablecoinUpgradeableAccountPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StablecoinUpgradeableAccountPaused)
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
		it.Event = new(StablecoinUpgradeableAccountPaused)
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
func (it *StablecoinUpgradeableAccountPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StablecoinUpgradeableAccountPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StablecoinUpgradeableAccountPaused represents a AccountPaused event raised by the StablecoinUpgradeable contract.
type StablecoinUpgradeableAccountPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAccountPaused is a free log retrieval operation binding the contract event 0xae7f60c1b8f645c3beffeb531169cbc446874bbf247698325318879ac850c346.
//
// Solidity: event AccountPaused(address account)
func (_StablecoinUpgradeable *StablecoinUpgradeableFilterer) FilterAccountPaused(opts *bind.FilterOpts) (*StablecoinUpgradeableAccountPausedIterator, error) {

	logs, sub, err := _StablecoinUpgradeable.contract.FilterLogs(opts, "AccountPaused")
	if err != nil {
		return nil, err
	}
	return &StablecoinUpgradeableAccountPausedIterator{contract: _StablecoinUpgradeable.contract, event: "AccountPaused", logs: logs, sub: sub}, nil
}

// WatchAccountPaused is a free log subscription operation binding the contract event 0xae7f60c1b8f645c3beffeb531169cbc446874bbf247698325318879ac850c346.
//
// Solidity: event AccountPaused(address account)
func (_StablecoinUpgradeable *StablecoinUpgradeableFilterer) WatchAccountPaused(opts *bind.WatchOpts, sink chan<- *StablecoinUpgradeableAccountPaused) (event.Subscription, error) {

	logs, sub, err := _StablecoinUpgradeable.contract.WatchLogs(opts, "AccountPaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StablecoinUpgradeableAccountPaused)
				if err := _StablecoinUpgradeable.contract.UnpackLog(event, "AccountPaused", log); err != nil {
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

// ParseAccountPaused is a log parse operation binding the contract event 0xae7f60c1b8f645c3beffeb531169cbc446874bbf247698325318879ac850c346.
//
// Solidity: event AccountPaused(address account)
func (_StablecoinUpgradeable *StablecoinUpgradeableFilterer) ParseAccountPaused(log types.Log) (*StablecoinUpgradeableAccountPaused, error) {
	event := new(StablecoinUpgradeableAccountPaused)
	if err := _StablecoinUpgradeable.contract.UnpackLog(event, "AccountPaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StablecoinUpgradeableAccountUnpausedIterator is returned from FilterAccountUnpaused and is used to iterate over the raw logs and unpacked data for AccountUnpaused events raised by the StablecoinUpgradeable contract.
type StablecoinUpgradeableAccountUnpausedIterator struct {
	Event *StablecoinUpgradeableAccountUnpaused // Event containing the contract specifics and raw log

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
func (it *StablecoinUpgradeableAccountUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StablecoinUpgradeableAccountUnpaused)
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
		it.Event = new(StablecoinUpgradeableAccountUnpaused)
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
func (it *StablecoinUpgradeableAccountUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StablecoinUpgradeableAccountUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StablecoinUpgradeableAccountUnpaused represents a AccountUnpaused event raised by the StablecoinUpgradeable contract.
type StablecoinUpgradeableAccountUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterAccountUnpaused is a free log retrieval operation binding the contract event 0x0c18efbde61ac471ead6960a3f1097735c68ecdb685ae8e2a108c28385399a65.
//
// Solidity: event AccountUnpaused(address account)
func (_StablecoinUpgradeable *StablecoinUpgradeableFilterer) FilterAccountUnpaused(opts *bind.FilterOpts) (*StablecoinUpgradeableAccountUnpausedIterator, error) {

	logs, sub, err := _StablecoinUpgradeable.contract.FilterLogs(opts, "AccountUnpaused")
	if err != nil {
		return nil, err
	}
	return &StablecoinUpgradeableAccountUnpausedIterator{contract: _StablecoinUpgradeable.contract, event: "AccountUnpaused", logs: logs, sub: sub}, nil
}

// WatchAccountUnpaused is a free log subscription operation binding the contract event 0x0c18efbde61ac471ead6960a3f1097735c68ecdb685ae8e2a108c28385399a65.
//
// Solidity: event AccountUnpaused(address account)
func (_StablecoinUpgradeable *StablecoinUpgradeableFilterer) WatchAccountUnpaused(opts *bind.WatchOpts, sink chan<- *StablecoinUpgradeableAccountUnpaused) (event.Subscription, error) {

	logs, sub, err := _StablecoinUpgradeable.contract.WatchLogs(opts, "AccountUnpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StablecoinUpgradeableAccountUnpaused)
				if err := _StablecoinUpgradeable.contract.UnpackLog(event, "AccountUnpaused", log); err != nil {
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

// ParseAccountUnpaused is a log parse operation binding the contract event 0x0c18efbde61ac471ead6960a3f1097735c68ecdb685ae8e2a108c28385399a65.
//
// Solidity: event AccountUnpaused(address account)
func (_StablecoinUpgradeable *StablecoinUpgradeableFilterer) ParseAccountUnpaused(log types.Log) (*StablecoinUpgradeableAccountUnpaused, error) {
	event := new(StablecoinUpgradeableAccountUnpaused)
	if err := _StablecoinUpgradeable.contract.UnpackLog(event, "AccountUnpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StablecoinUpgradeableApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the StablecoinUpgradeable contract.
type StablecoinUpgradeableApprovalIterator struct {
	Event *StablecoinUpgradeableApproval // Event containing the contract specifics and raw log

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
func (it *StablecoinUpgradeableApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StablecoinUpgradeableApproval)
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
		it.Event = new(StablecoinUpgradeableApproval)
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
func (it *StablecoinUpgradeableApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StablecoinUpgradeableApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StablecoinUpgradeableApproval represents a Approval event raised by the StablecoinUpgradeable contract.
type StablecoinUpgradeableApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_StablecoinUpgradeable *StablecoinUpgradeableFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*StablecoinUpgradeableApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _StablecoinUpgradeable.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &StablecoinUpgradeableApprovalIterator{contract: _StablecoinUpgradeable.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_StablecoinUpgradeable *StablecoinUpgradeableFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *StablecoinUpgradeableApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _StablecoinUpgradeable.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StablecoinUpgradeableApproval)
				if err := _StablecoinUpgradeable.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_StablecoinUpgradeable *StablecoinUpgradeableFilterer) ParseApproval(log types.Log) (*StablecoinUpgradeableApproval, error) {
	event := new(StablecoinUpgradeableApproval)
	if err := _StablecoinUpgradeable.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StablecoinUpgradeableInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the StablecoinUpgradeable contract.
type StablecoinUpgradeableInitializedIterator struct {
	Event *StablecoinUpgradeableInitialized // Event containing the contract specifics and raw log

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
func (it *StablecoinUpgradeableInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StablecoinUpgradeableInitialized)
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
		it.Event = new(StablecoinUpgradeableInitialized)
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
func (it *StablecoinUpgradeableInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StablecoinUpgradeableInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StablecoinUpgradeableInitialized represents a Initialized event raised by the StablecoinUpgradeable contract.
type StablecoinUpgradeableInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_StablecoinUpgradeable *StablecoinUpgradeableFilterer) FilterInitialized(opts *bind.FilterOpts) (*StablecoinUpgradeableInitializedIterator, error) {

	logs, sub, err := _StablecoinUpgradeable.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &StablecoinUpgradeableInitializedIterator{contract: _StablecoinUpgradeable.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_StablecoinUpgradeable *StablecoinUpgradeableFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *StablecoinUpgradeableInitialized) (event.Subscription, error) {

	logs, sub, err := _StablecoinUpgradeable.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StablecoinUpgradeableInitialized)
				if err := _StablecoinUpgradeable.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_StablecoinUpgradeable *StablecoinUpgradeableFilterer) ParseInitialized(log types.Log) (*StablecoinUpgradeableInitialized, error) {
	event := new(StablecoinUpgradeableInitialized)
	if err := _StablecoinUpgradeable.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StablecoinUpgradeablePausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the StablecoinUpgradeable contract.
type StablecoinUpgradeablePausedIterator struct {
	Event *StablecoinUpgradeablePaused // Event containing the contract specifics and raw log

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
func (it *StablecoinUpgradeablePausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StablecoinUpgradeablePaused)
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
		it.Event = new(StablecoinUpgradeablePaused)
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
func (it *StablecoinUpgradeablePausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StablecoinUpgradeablePausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StablecoinUpgradeablePaused represents a Paused event raised by the StablecoinUpgradeable contract.
type StablecoinUpgradeablePaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_StablecoinUpgradeable *StablecoinUpgradeableFilterer) FilterPaused(opts *bind.FilterOpts) (*StablecoinUpgradeablePausedIterator, error) {

	logs, sub, err := _StablecoinUpgradeable.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &StablecoinUpgradeablePausedIterator{contract: _StablecoinUpgradeable.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_StablecoinUpgradeable *StablecoinUpgradeableFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *StablecoinUpgradeablePaused) (event.Subscription, error) {

	logs, sub, err := _StablecoinUpgradeable.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StablecoinUpgradeablePaused)
				if err := _StablecoinUpgradeable.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_StablecoinUpgradeable *StablecoinUpgradeableFilterer) ParsePaused(log types.Log) (*StablecoinUpgradeablePaused, error) {
	event := new(StablecoinUpgradeablePaused)
	if err := _StablecoinUpgradeable.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StablecoinUpgradeableRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the StablecoinUpgradeable contract.
type StablecoinUpgradeableRoleAdminChangedIterator struct {
	Event *StablecoinUpgradeableRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *StablecoinUpgradeableRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StablecoinUpgradeableRoleAdminChanged)
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
		it.Event = new(StablecoinUpgradeableRoleAdminChanged)
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
func (it *StablecoinUpgradeableRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StablecoinUpgradeableRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StablecoinUpgradeableRoleAdminChanged represents a RoleAdminChanged event raised by the StablecoinUpgradeable contract.
type StablecoinUpgradeableRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_StablecoinUpgradeable *StablecoinUpgradeableFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*StablecoinUpgradeableRoleAdminChangedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _StablecoinUpgradeable.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &StablecoinUpgradeableRoleAdminChangedIterator{contract: _StablecoinUpgradeable.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_StablecoinUpgradeable *StablecoinUpgradeableFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *StablecoinUpgradeableRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var previousAdminRoleRule []interface{}
	for _, previousAdminRoleItem := range previousAdminRole {
		previousAdminRoleRule = append(previousAdminRoleRule, previousAdminRoleItem)
	}
	var newAdminRoleRule []interface{}
	for _, newAdminRoleItem := range newAdminRole {
		newAdminRoleRule = append(newAdminRoleRule, newAdminRoleItem)
	}

	logs, sub, err := _StablecoinUpgradeable.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StablecoinUpgradeableRoleAdminChanged)
				if err := _StablecoinUpgradeable.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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

// ParseRoleAdminChanged is a log parse operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_StablecoinUpgradeable *StablecoinUpgradeableFilterer) ParseRoleAdminChanged(log types.Log) (*StablecoinUpgradeableRoleAdminChanged, error) {
	event := new(StablecoinUpgradeableRoleAdminChanged)
	if err := _StablecoinUpgradeable.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StablecoinUpgradeableRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the StablecoinUpgradeable contract.
type StablecoinUpgradeableRoleGrantedIterator struct {
	Event *StablecoinUpgradeableRoleGranted // Event containing the contract specifics and raw log

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
func (it *StablecoinUpgradeableRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StablecoinUpgradeableRoleGranted)
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
		it.Event = new(StablecoinUpgradeableRoleGranted)
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
func (it *StablecoinUpgradeableRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StablecoinUpgradeableRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StablecoinUpgradeableRoleGranted represents a RoleGranted event raised by the StablecoinUpgradeable contract.
type StablecoinUpgradeableRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_StablecoinUpgradeable *StablecoinUpgradeableFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*StablecoinUpgradeableRoleGrantedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _StablecoinUpgradeable.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &StablecoinUpgradeableRoleGrantedIterator{contract: _StablecoinUpgradeable.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_StablecoinUpgradeable *StablecoinUpgradeableFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *StablecoinUpgradeableRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _StablecoinUpgradeable.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StablecoinUpgradeableRoleGranted)
				if err := _StablecoinUpgradeable.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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

// ParseRoleGranted is a log parse operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_StablecoinUpgradeable *StablecoinUpgradeableFilterer) ParseRoleGranted(log types.Log) (*StablecoinUpgradeableRoleGranted, error) {
	event := new(StablecoinUpgradeableRoleGranted)
	if err := _StablecoinUpgradeable.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StablecoinUpgradeableRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the StablecoinUpgradeable contract.
type StablecoinUpgradeableRoleRevokedIterator struct {
	Event *StablecoinUpgradeableRoleRevoked // Event containing the contract specifics and raw log

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
func (it *StablecoinUpgradeableRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StablecoinUpgradeableRoleRevoked)
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
		it.Event = new(StablecoinUpgradeableRoleRevoked)
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
func (it *StablecoinUpgradeableRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StablecoinUpgradeableRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StablecoinUpgradeableRoleRevoked represents a RoleRevoked event raised by the StablecoinUpgradeable contract.
type StablecoinUpgradeableRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_StablecoinUpgradeable *StablecoinUpgradeableFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*StablecoinUpgradeableRoleRevokedIterator, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _StablecoinUpgradeable.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &StablecoinUpgradeableRoleRevokedIterator{contract: _StablecoinUpgradeable.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_StablecoinUpgradeable *StablecoinUpgradeableFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *StablecoinUpgradeableRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

	var roleRule []interface{}
	for _, roleItem := range role {
		roleRule = append(roleRule, roleItem)
	}
	var accountRule []interface{}
	for _, accountItem := range account {
		accountRule = append(accountRule, accountItem)
	}
	var senderRule []interface{}
	for _, senderItem := range sender {
		senderRule = append(senderRule, senderItem)
	}

	logs, sub, err := _StablecoinUpgradeable.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StablecoinUpgradeableRoleRevoked)
				if err := _StablecoinUpgradeable.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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

// ParseRoleRevoked is a log parse operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_StablecoinUpgradeable *StablecoinUpgradeableFilterer) ParseRoleRevoked(log types.Log) (*StablecoinUpgradeableRoleRevoked, error) {
	event := new(StablecoinUpgradeableRoleRevoked)
	if err := _StablecoinUpgradeable.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StablecoinUpgradeableTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the StablecoinUpgradeable contract.
type StablecoinUpgradeableTransferIterator struct {
	Event *StablecoinUpgradeableTransfer // Event containing the contract specifics and raw log

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
func (it *StablecoinUpgradeableTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StablecoinUpgradeableTransfer)
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
		it.Event = new(StablecoinUpgradeableTransfer)
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
func (it *StablecoinUpgradeableTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StablecoinUpgradeableTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StablecoinUpgradeableTransfer represents a Transfer event raised by the StablecoinUpgradeable contract.
type StablecoinUpgradeableTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_StablecoinUpgradeable *StablecoinUpgradeableFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*StablecoinUpgradeableTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _StablecoinUpgradeable.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &StablecoinUpgradeableTransferIterator{contract: _StablecoinUpgradeable.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_StablecoinUpgradeable *StablecoinUpgradeableFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *StablecoinUpgradeableTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _StablecoinUpgradeable.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StablecoinUpgradeableTransfer)
				if err := _StablecoinUpgradeable.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_StablecoinUpgradeable *StablecoinUpgradeableFilterer) ParseTransfer(log types.Log) (*StablecoinUpgradeableTransfer, error) {
	event := new(StablecoinUpgradeableTransfer)
	if err := _StablecoinUpgradeable.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StablecoinUpgradeableUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the StablecoinUpgradeable contract.
type StablecoinUpgradeableUnpausedIterator struct {
	Event *StablecoinUpgradeableUnpaused // Event containing the contract specifics and raw log

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
func (it *StablecoinUpgradeableUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StablecoinUpgradeableUnpaused)
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
		it.Event = new(StablecoinUpgradeableUnpaused)
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
func (it *StablecoinUpgradeableUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StablecoinUpgradeableUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StablecoinUpgradeableUnpaused represents a Unpaused event raised by the StablecoinUpgradeable contract.
type StablecoinUpgradeableUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_StablecoinUpgradeable *StablecoinUpgradeableFilterer) FilterUnpaused(opts *bind.FilterOpts) (*StablecoinUpgradeableUnpausedIterator, error) {

	logs, sub, err := _StablecoinUpgradeable.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &StablecoinUpgradeableUnpausedIterator{contract: _StablecoinUpgradeable.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_StablecoinUpgradeable *StablecoinUpgradeableFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *StablecoinUpgradeableUnpaused) (event.Subscription, error) {

	logs, sub, err := _StablecoinUpgradeable.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StablecoinUpgradeableUnpaused)
				if err := _StablecoinUpgradeable.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_StablecoinUpgradeable *StablecoinUpgradeableFilterer) ParseUnpaused(log types.Log) (*StablecoinUpgradeableUnpaused, error) {
	event := new(StablecoinUpgradeableUnpaused)
	if err := _StablecoinUpgradeable.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StablecoinUpgradeableUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the StablecoinUpgradeable contract.
type StablecoinUpgradeableUpgradedIterator struct {
	Event *StablecoinUpgradeableUpgraded // Event containing the contract specifics and raw log

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
func (it *StablecoinUpgradeableUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StablecoinUpgradeableUpgraded)
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
		it.Event = new(StablecoinUpgradeableUpgraded)
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
func (it *StablecoinUpgradeableUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StablecoinUpgradeableUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StablecoinUpgradeableUpgraded represents a Upgraded event raised by the StablecoinUpgradeable contract.
type StablecoinUpgradeableUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_StablecoinUpgradeable *StablecoinUpgradeableFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*StablecoinUpgradeableUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _StablecoinUpgradeable.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &StablecoinUpgradeableUpgradedIterator{contract: _StablecoinUpgradeable.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_StablecoinUpgradeable *StablecoinUpgradeableFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *StablecoinUpgradeableUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _StablecoinUpgradeable.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StablecoinUpgradeableUpgraded)
				if err := _StablecoinUpgradeable.contract.UnpackLog(event, "Upgraded", log); err != nil {
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
func (_StablecoinUpgradeable *StablecoinUpgradeableFilterer) ParseUpgraded(log types.Log) (*StablecoinUpgradeableUpgraded, error) {
	event := new(StablecoinUpgradeableUpgraded)
	if err := _StablecoinUpgradeable.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
