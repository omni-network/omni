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
	Bin: "0x60a060405230608052348015610013575f80fd5b5061001c610021565b6100d3565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000900460ff16156100715760405163f92ee8a960e01b815260040160405180910390fd5b80546001600160401b03908116146100d05780546001600160401b0319166001600160401b0390811782556040519081527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b50565b6080516120266100f95f395f81816110e30152818161110c015261126d01526120265ff3fe6080604052600436106101f1575f3560e01c806370a0823111610108578063ad3cb1cc1161009d578063d547741f1161006d578063d547741f146105d9578063dd62ed3e146105f8578063e63ab1e914610617578063f3df317e14610637578063f72c0d8b14610656575f80fd5b8063ad3cb1cc14610538578063bc8c4b4f14610568578063c91f0c5314610587578063d5391393146105a6575f80fd5b806394408b9a116100d857806394408b9a146104d357806395d89b41146104f2578063a217fddf14610506578063a9059cbb14610519575f80fd5b806370a08231146104415780638456cb5914610481578063917b1ace1461049557806391d14854146104b4575f80fd5b8063313ce5671161018957806342966c681161015957806342966c68146103a5578063475ca324146103c45780634f1ef286146103f757806352d1902d1461040a5780635c975abb1461041e575f80fd5b8063313ce5671461033857806336568abe146103535780633f4ba83a1461037257806340c10f1914610386575f80fd5b806323b872dd116101c457806323b872dd146102a6578063248a9ca3146102c5578063282c51f3146102e45780632f2ff15d14610317575f80fd5b806301ffc9a7146101f557806306fdde0314610229578063095ea7b31461024a57806318160ddd14610269575b5f80fd5b348015610200575f80fd5b5061021461020f366004611a31565b610689565b60405190151581526020015b60405180910390f35b348015610234575f80fd5b5061023d6106bf565b6040516102209190611a58565b348015610255575f80fd5b50610214610264366004611aa8565b61077f565b348015610274575f80fd5b507f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace02545b604051908152602001610220565b3480156102b1575f80fd5b506102146102c0366004611ad0565b610806565b3480156102d0575f80fd5b506102986102df366004611b0a565b61082b565b3480156102ef575f80fd5b506102987f9667e80708b6eeeb0053fa0cca44e028ff548e2a9f029edfeac87c118b08b7c881565b348015610322575f80fd5b50610336610331366004611b21565b61084b565b005b348015610343575f80fd5b5060405160128152602001610220565b34801561035e575f80fd5b5061033661036d366004611b21565b61086d565b34801561037d575f80fd5b506103366108a5565b348015610391575f80fd5b506103366103a0366004611aa8565b6108c7565b3480156103b0575f80fd5b506103366103bf366004611b0a565b6108fb565b3480156103cf575f80fd5b506102987f715bacafb7a853b9b91b59ae724920a9eb0c006c5b318ac393fa1bc8974edd9881565b610336610405366004611bd6565b610933565b348015610415575f80fd5b5061029861094e565b348015610429575f80fd5b505f80516020611fd18339815191525460ff16610214565b34801561044c575f80fd5b5061029861045b366004611c34565b6001600160a01b03165f9081525f80516020611f51833981519152602052604090205490565b34801561048c575f80fd5b50610336610969565b3480156104a0575f80fd5b506103366104af366004611c4d565b610988565b3480156104bf575f80fd5b506102146104ce366004611b21565b610a9a565b3480156104de575f80fd5b506103366104ed366004611c34565b610ad0565b3480156104fd575f80fd5b5061023d610af0565b348015610511575f80fd5b506102985f81565b348015610524575f80fd5b50610214610533366004611aa8565b610b2e565b348015610543575f80fd5b5061023d604051806040016040528060058152602001640352e302e360dc1b81525081565b348015610573575f80fd5b50610214610582366004611c34565b610b45565b348015610592575f80fd5b506103366105a1366004611cdc565b610b81565b3480156105b1575f80fd5b506102987ff0887ba65ee2024ea881d91b74c2450ef19e1557f03bed3ea9f16b037cbe2dc981565b3480156105e4575f80fd5b506103366105f3366004611b21565b610d5a565b348015610603575f80fd5b50610298610612366004611d91565b610d76565b348015610622575f80fd5b506102985f80516020611f9183398151915281565b348015610642575f80fd5b50610336610651366004611aa8565b610dbf565b348015610661575f80fd5b506102987fa615a8afb6fffcb8c6809ac0997b5c9c12b8cc97651150f14c8f6203168cff4c81565b5f6001600160e01b03198216637965db0b60e01b14806106b957506301ffc9a760e01b6001600160e01b03198316145b92915050565b7f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace0380546060915f80516020611f51833981519152916106fd90611db9565b80601f016020809104026020016040519081016040528092919081815260200182805461072990611db9565b80156107745780601f1061074b57610100808354040283529160200191610774565b820191905f5260205f20905b81548152906001019060200180831161075757829003601f168201915b505050505091505090565b5f8261078a81610b45565b156107b85760405163c8c29b9960e01b81526001600160a01b03821660048201526024015b60405180910390fd5b336107c281610b45565b156107eb5760405163c8c29b9960e01b81526001600160a01b03821660048201526024016107af565b6107f3610df3565b6107fd8585610e25565b95945050505050565b5f33610813858285610e32565b61081e858585610e90565b60019150505b9392505050565b5f9081525f80516020611fb1833981519152602052604090206001015490565b6108548261082b565b61085d81610eed565b6108678383610ef7565b50505050565b6001600160a01b03811633146108965760405163334bd91960e11b815260040160405180910390fd5b6108a08282610f98565b505050565b5f80516020611f918339815191526108bc81610eed565b6108c4611011565b50565b7ff0887ba65ee2024ea881d91b74c2450ef19e1557f03bed3ea9f16b037cbe2dc96108f181610eed565b6108a08383611070565b7f9667e80708b6eeeb0053fa0cca44e028ff548e2a9f029edfeac87c118b08b7c861092581610eed565b61092f33836110a4565b5050565b61093b6110d8565b6109448261117c565b61092f82826111a6565b5f610957611262565b505f80516020611f7183398151915290565b5f80516020611f9183398151915261098081610eed565b6108c46112ab565b5f80516020611f9183398151915261099f81610eed565b5f82815b81811015610a9257826001600160a01b03168686838181106109c7576109c7611df1565b90506020020160208101906109dc9190611c34565b6001600160a01b031611610a325760405162461bcd60e51b815260206004820152601a60248201527f4164647265737365732073686f756c6420626520736f7274656400000000000060448201526064016107af565b610a61868683818110610a4757610a47611df1565b9050602002016020810190610a5c9190611c34565b6112f3565b858582818110610a7357610a73611df1565b9050602002016020810190610a889190611c34565b92506001016109a3565b505050505050565b5f9182525f80516020611fb1833981519152602090815260408084206001600160a01b0393909316845291905290205460ff1690565b5f80516020611f91833981519152610ae781610eed565b61092f826113a2565b7f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace0480546060915f80516020611f51833981519152916106fd90611db9565b5f33610b3b818585610e90565b5060019392505050565b6001600160a01b03165f9081527f345cc2404af916c3db112e7a6103770647a90ed78a5d681e21dc2e1174232900602052604090205460ff1690565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a008054600160401b810460ff16159067ffffffffffffffff165f81158015610bc65750825b90505f8267ffffffffffffffff166001148015610be25750303b155b905081158015610bf0575080155b15610c0e5760405163f92ee8a960e01b815260040160405180910390fd5b845467ffffffffffffffff191660011785558315610c3857845460ff60401b1916600160401b1785555b610c428c8c611444565b610c4a611456565b610c52611456565b610c5c5f8a610ef7565b50610c877ff0887ba65ee2024ea881d91b74c2450ef19e1557f03bed3ea9f16b037cbe2dc98b610ef7565b50610cb27fa615a8afb6fffcb8c6809ac0997b5c9c12b8cc97651150f14c8f6203168cff4c89610ef7565b50610cbb61145e565b610cc3611456565b610cda5f80516020611f9183398151915288610ef7565b50610d057f715bacafb7a853b9b91b59ae724920a9eb0c006c5b318ac393fa1bc8974edd9887610ef7565b508315610d4c57845460ff60401b19168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b505050505050505050505050565b610d638261082b565b610d6c81610eed565b6108678383610f98565b6001600160a01b039182165f9081527f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace016020908152604080832093909416825291909152205490565b7f715bacafb7a853b9b91b59ae724920a9eb0c006c5b318ac393fa1bc8974edd98610de981610eed565b6108a083836110a4565b5f80516020611fd18339815191525460ff1615610e235760405163d93c066560e01b815260040160405180910390fd5b565b5f33610b3b81858561146e565b5f610e3d8484610d76565b90505f198110156108675781811015610e8257604051637dc7a0d960e11b81526001600160a01b038416600482015260248101829052604481018390526064016107af565b61086784848484035f611477565b6001600160a01b038316610eb957604051634b637e8f60e11b81525f60048201526024016107af565b6001600160a01b038216610ee25760405163ec442f0560e01b81525f60048201526024016107af565b6108a083838361155b565b6108c481336115ff565b5f5f80516020611fb1833981519152610f108484610a9a565b610f8f575f848152602082815260408083206001600160a01b03871684529091529020805460ff19166001179055610f453390565b6001600160a01b0316836001600160a01b0316857f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a460019150506106b9565b5f9150506106b9565b5f5f80516020611fb1833981519152610fb18484610a9a565b15610f8f575f848152602082815260408083206001600160a01b0387168085529252808320805460ff1916905551339287917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a460019150506106b9565b611019611638565b5f80516020611fd1833981519152805460ff191681557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa335b6040516001600160a01b03909116815260200160405180910390a150565b6001600160a01b0382166110995760405163ec442f0560e01b81525f60048201526024016107af565b61092f5f838361155b565b6001600160a01b0382166110cd57604051634b637e8f60e11b81525f60048201526024016107af565b61092f825f8361155b565b306001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016148061115e57507f00000000000000000000000000000000000000000000000000000000000000006001600160a01b03166111525f80516020611f71833981519152546001600160a01b031690565b6001600160a01b031614155b15610e235760405163703e46dd60e11b815260040160405180910390fd5b7fa615a8afb6fffcb8c6809ac0997b5c9c12b8cc97651150f14c8f6203168cff4c61092f81610eed565b816001600160a01b03166352d1902d6040518163ffffffff1660e01b8152600401602060405180830381865afa925050508015611200575060408051601f3d908101601f191682019092526111fd91810190611e05565b60015b61122857604051634c9c8ce360e01b81526001600160a01b03831660048201526024016107af565b5f80516020611f71833981519152811461125857604051632a87526960e21b8152600481018290526024016107af565b6108a08383611667565b306001600160a01b037f00000000000000000000000000000000000000000000000000000000000000001614610e235760405163703e46dd60e11b815260040160405180910390fd5b6112b3610df3565b5f80516020611fd1833981519152805460ff191660011781557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a25833611052565b806112fd81610b45565b156113265760405163c8c29b9960e01b81526001600160a01b03821660048201526024016107af565b6001600160a01b0382165f8181527f345cc2404af916c3db112e7a6103770647a90ed78a5d681e21dc2e11742329006020818152604092839020805460ff191660011790559151928352917fae7f60c1b8f645c3beffeb531169cbc446874bbf247698325318879ac850c34691015b60405180910390a1505050565b806113ac81610b45565b6113d457604051638d542b2960e01b81526001600160a01b03821660048201526024016107af565b6001600160a01b0382165f8181527f345cc2404af916c3db112e7a6103770647a90ed78a5d681e21dc2e11742329006020818152604092839020805460ff191690559151928352917f0c18efbde61ac471ead6960a3f1097735c68ecdb685ae8e2a108c28385399a659101611395565b61144c6116bc565b61092f8282611705565b610e236116bc565b6114666116bc565b610e23611755565b6108a083838360015b5f80516020611f518339815191526001600160a01b0385166114ae5760405163e602df0560e01b81525f60048201526024016107af565b6001600160a01b0384166114d757604051634a1406b160e11b81525f60048201526024016107af565b6001600160a01b038086165f9081526001830160209081526040808320938816835292905220839055811561155457836001600160a01b0316856001600160a01b03167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b9258560405161154b91815260200190565b60405180910390a35b5050505050565b8261156581610b45565b1561158e5760405163c8c29b9960e01b81526001600160a01b03821660048201526024016107af565b8261159881610b45565b156115c15760405163c8c29b9960e01b81526001600160a01b03821660048201526024016107af565b336115cb81610b45565b156115f45760405163c8c29b9960e01b81526001600160a01b03821660048201526024016107af565b610a92868686611775565b6116098282610a9a565b61092f5760405163e2517d3f60e01b81526001600160a01b0382166004820152602481018390526044016107af565b5f80516020611fd18339815191525460ff16610e2357604051638dfc202b60e01b815260040160405180910390fd5b61167082611788565b6040516001600160a01b038316907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b905f90a28051156116b4576108a082826117eb565b61092f611854565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a0054600160401b900460ff16610e2357604051631afcd79f60e31b815260040160405180910390fd5b61170d6116bc565b5f80516020611f518339815191527f52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace036117468482611e60565b50600481016108678382611e60565b61175d6116bc565b5f80516020611fd1833981519152805460ff19169055565b61177d610df3565b6108a0838383611873565b806001600160a01b03163b5f036117bd57604051634c9c8ce360e01b81526001600160a01b03821660048201526024016107af565b5f80516020611f7183398151915280546001600160a01b0319166001600160a01b0392909216919091179055565b60605f80846001600160a01b0316846040516118079190611f1b565b5f60405180830381855af49150503d805f811461183f576040519150601f19603f3d011682016040523d82523d5f602084013e611844565b606091505b50915091506107fd8583836119ac565b3415610e235760405163b398979f60e01b815260040160405180910390fd5b5f80516020611f518339815191526001600160a01b0384166118ad5781816002015f8282546118a29190611f31565b9091555061191d9050565b6001600160a01b0384165f90815260208290526040902054828110156118ff5760405163391434e360e21b81526001600160a01b038616600482015260248101829052604481018490526064016107af565b6001600160a01b0385165f9081526020839052604090209083900390555b6001600160a01b03831661193b576002810180548390039055611959565b6001600160a01b0383165f9081526020829052604090208054830190555b826001600160a01b0316846001600160a01b03167fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef8460405161199e91815260200190565b60405180910390a350505050565b6060826119c1576119bc82611a08565b610824565b81511580156119d857506001600160a01b0384163b155b15611a0157604051639996b31560e01b81526001600160a01b03851660048201526024016107af565b5080610824565b805115611a185780518082602001fd5b60405163d6bda27560e01b815260040160405180910390fd5b5f60208284031215611a41575f80fd5b81356001600160e01b031981168114610824575f80fd5b602081525f82518060208401528060208501604085015e5f604082850101526040601f19601f83011684010191505092915050565b80356001600160a01b0381168114611aa3575f80fd5b919050565b5f8060408385031215611ab9575f80fd5b611ac283611a8d565b946020939093013593505050565b5f805f60608486031215611ae2575f80fd5b611aeb84611a8d565b9250611af960208501611a8d565b929592945050506040919091013590565b5f60208284031215611b1a575f80fd5b5035919050565b5f8060408385031215611b32575f80fd5b82359150611b4260208401611a8d565b90509250929050565b634e487b7160e01b5f52604160045260245ffd5b5f8067ffffffffffffffff841115611b7957611b79611b4b565b50604051601f19601f85018116603f0116810181811067ffffffffffffffff82111715611ba857611ba8611b4b565b604052838152905080828401851015611bbf575f80fd5b838360208301375f60208583010152509392505050565b5f8060408385031215611be7575f80fd5b611bf083611a8d565b9150602083013567ffffffffffffffff811115611c0b575f80fd5b8301601f81018513611c1b575f80fd5b611c2a85823560208401611b5f565b9150509250929050565b5f60208284031215611c44575f80fd5b61082482611a8d565b5f8060208385031215611c5e575f80fd5b823567ffffffffffffffff811115611c74575f80fd5b8301601f81018513611c84575f80fd5b803567ffffffffffffffff811115611c9a575f80fd5b8560208260051b8401011115611cae575f80fd5b6020919091019590945092505050565b5f82601f830112611ccd575f80fd5b61082483833560208501611b5f565b5f805f805f805f60e0888a031215611cf2575f80fd5b873567ffffffffffffffff811115611d08575f80fd5b611d148a828b01611cbe565b975050602088013567ffffffffffffffff811115611d30575f80fd5b611d3c8a828b01611cbe565b965050611d4b60408901611a8d565b9450611d5960608901611a8d565b9350611d6760808901611a8d565b9250611d7560a08901611a8d565b9150611d8360c08901611a8d565b905092959891949750929550565b5f8060408385031215611da2575f80fd5b611dab83611a8d565b9150611b4260208401611a8d565b600181811c90821680611dcd57607f821691505b602082108103611deb57634e487b7160e01b5f52602260045260245ffd5b50919050565b634e487b7160e01b5f52603260045260245ffd5b5f60208284031215611e15575f80fd5b5051919050565b601f8211156108a057805f5260205f20601f840160051c81016020851015611e415750805b601f840160051c820191505b81811015611554575f8155600101611e4d565b815167ffffffffffffffff811115611e7a57611e7a611b4b565b611e8e81611e888454611db9565b84611e1c565b6020601f821160018114611ec0575f8315611ea95750848201515b5f19600385901b1c1916600184901b178455611554565b5f84815260208120601f198516915b82811015611eef5787850151825560209485019460019092019101611ecf565b5084821015611f0c57868401515f19600387901b60f8161c191681555b50505050600190811b01905550565b5f82518060208501845e5f920191825250919050565b808201808211156106b957634e487b7160e01b5f52601160045260245ffdfe52c63247e1f47db19d5ce0460030c497f067ca4cebf71ba98eeadabe20bace00360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc539440820030c4994db4e31b6b800deafd503688728f932addfe7a410515c14c02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800cd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f03300a2646970667358221220ed32196bf4d29519cbd67a11c282fb0d3f666f4250eba7f50b9f16feba2cbf0c64736f6c634300081a0033",
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
