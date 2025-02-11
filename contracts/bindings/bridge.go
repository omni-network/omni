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

// IBridgeRoute is an auto generated low-level Go binding around an user-defined struct.
type IBridgeRoute struct {
	Bridge     common.Address
	HasLockbox bool
}

// BridgeMetaData contains all meta data concerning the Bridge contract.
var BridgeMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"DEFAULT_ADMIN_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"PAUSER_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"bridgeFee\",\"inputs\":[{\"name\":\"destChainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"fee\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"defaultConfLevel\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRoleAdmin\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRoute\",\"inputs\":[{\"name\":\"destChainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"bridge\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"hasLockbox\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"grantRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"hasRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"admin_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"pauser_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"omni_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"token_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"lockbox_\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"lockbox\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"omni\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIOmniPortal\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"receiveToken\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"callerConfirmation\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"revokeRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"sendToken\",\"inputs\":[{\"name\":\"destChainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"wrap\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"setRoutes\",\"inputs\":[{\"name\":\"chainIds\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"},{\"name\":\"routes\",\"type\":\"tuple[]\",\"internalType\":\"structIBridge.Route[]\",\"components\":[{\"name\":\"bridge\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"hasLockbox\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"token\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"DefaultConfLevelSet\",\"inputs\":[{\"name\":\"conf\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"LockboxWithdrawalFailed\",\"inputs\":[{\"name\":\"badLockbox\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OmniPortalSet\",\"inputs\":[{\"name\":\"omni\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleAdminChanged\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"previousAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"newAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleGranted\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleRevoked\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RouteConfigured\",\"inputs\":[{\"name\":\"destChainId\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"bridge\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"hasLockbox\",\"type\":\"bool\",\"indexed\":true,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokenReceived\",\"inputs\":[{\"name\":\"srcChainId\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokenSent\",\"inputs\":[{\"name\":\"destChainId\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AccessControlBadConfirmation\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AccessControlUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"neededRole\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"ArrayLengthMismatch\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CannotWrap\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"EnforcedPause\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ExpectedPause\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientPayment\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidRoute\",\"inputs\":[{\"name\":\"chainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"Unauthorized\",\"inputs\":[{\"name\":\"chainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ZeroAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroAmount\",\"inputs\":[]}]",
	Bin: "0x6080604052348015600f57600080fd5b506016601a565b60ca565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000900460ff161560695760405163f92ee8a960e01b815260040160405180910390fd5b80546001600160401b039081161460c75780546001600160401b0319166001600160401b0390811782556040519081527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b50565b611ecc806100d96000396000f3fe60806040526004361061012a5760003560e01c806374eeb847116100ab57806397235a1e1161006f57806397235a1e146103b0578063a217fddf146103d0578063a48fcd8e146103e5578063d547741f146103f8578063e63ab1e914610418578063fc0c546a1461043a57600080fd5b806374eeb847146103085780638456cb591461033b5780638ad201f61461035057806391d148541461037057806392099d441461039057600080fd5b806339acf9f1116100f257806339acf9f1146101f45780633f4ba83a1461022c5780634acd8d82146102415780635c975abb146102c357806366cc5702146102e857600080fd5b806301ffc9a71461012f5780631459457a14610164578063248a9ca3146101865780632f2ff15d146101b457806336568abe146101d4575b600080fd5b34801561013b57600080fd5b5061014f61014a36600461194d565b61045a565b60405190151581526020015b60405180910390f35b34801561017057600080fd5b5061018461017f36600461198c565b610491565b005b34801561019257600080fd5b506101a66101a13660046119fd565b610757565b60405190815260200161015b565b3480156101c057600080fd5b506101846101cf366004611a16565b610779565b3480156101e057600080fd5b506101846101ef366004611a16565b61079b565b34801561020057600080fd5b50600054610214906001600160a01b031681565b6040516001600160a01b03909116815260200161015b565b34801561023857600080fd5b506101846107d3565b34801561024d57600080fd5b506102a461025c366004611a5b565b6001600160401b03166000908152603460209081526040918290208251808401909352546001600160a01b038116808452600160a01b90910460ff1615159290910182905291565b604080516001600160a01b03909316835290151560208301520161015b565b3480156102cf57600080fd5b50600080516020611e778339815191525460ff1661014f565b3480156102f457600080fd5b50603354610214906001600160a01b031681565b34801561031457600080fd5b5060005461032990600160a01b900460ff1681565b60405160ff909116815260200161015b565b34801561034757600080fd5b506101846107f6565b34801561035c57600080fd5b506101a661036b366004611a5b565b610816565b34801561037c57600080fd5b5061014f61038b366004611a16565b6108c2565b34801561039c57600080fd5b506101846103ab366004611ac3565b6108fa565b3480156103bc57600080fd5b506101846103cb366004611b63565b610ab3565b3480156103dc57600080fd5b506101a6600081565b6101846103f3366004611b9d565b610c28565b34801561040457600080fd5b50610184610413366004611a16565b610c48565b34801561042457600080fd5b506101a6600080516020611e3783398151915281565b34801561044657600080fd5b50603254610214906001600160a01b031681565b60006001600160e01b03198216637965db0b60e01b148061048b57506301ffc9a760e01b6001600160e01b03198316145b92915050565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a008054600160401b810460ff1615906001600160401b03166000811580156104d65750825b90506000826001600160401b031660011480156104f25750303b155b905081158015610500575080155b1561051e5760405163f92ee8a960e01b815260040160405180910390fd5b845467ffffffffffffffff19166001178555831561054857845460ff60401b1916600160401b1785555b6001600160a01b038a1661056f5760405163d92e233d60e01b815260040160405180910390fd5b6001600160a01b0389166105965760405163d92e233d60e01b815260040160405180910390fd5b6001600160a01b0388166105bd5760405163d92e233d60e01b815260040160405180910390fd5b6001600160a01b0387166105e45760405163d92e233d60e01b815260040160405180910390fd5b6105ec610c64565b6105f4610c6e565b6105ff886004610c7e565b61060a60008b610c9c565b50610623600080516020611e378339815191528a610c9c565b50603280546001600160a01b0319166001600160a01b038981169190911790915586161561066757603380546001600160a01b0319166001600160a01b0388161790555b6001600160a01b03861615610705576106ef86600019886001600160a01b031663fc0c546a6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156106bb573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906106df9190611bf0565b6001600160a01b03169190610d41565b6107056001600160a01b03881687600019610d41565b831561074b57845460ff60401b19168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b50505050505050505050565b6000908152600080516020611e57833981519152602052604090206001015490565b61078282610757565b61078b81610dd8565b6107958383610c9c565b50505050565b6001600160a01b03811633146107c45760405163334bd91960e11b815260040160405180910390fd5b6107ce8282610de2565b505050565b600080516020611e378339815191526107eb81610dd8565b6107f3610e5e565b50565b600080516020611e3783398151915261080e81610dd8565b6107f3610ebf565b6001600160401b03811660009081526034602090815260408083208151808301835290546001600160a01b038082168352600160a01b90910460ff16151593820193909352905160248101929092526000196044830152906108bb90849060640160408051601f19818403018152919052602080820180516001600160e01b0316634b91ad0f60e11b1790528401516108b2576201e848610f08565b6202bf20610f08565b9392505050565b6000918252600080516020611e57833981519152602090815260408084206001600160a01b0393909316845291905290205460ff1690565b600061090581610dd8565b8382146109255760405163512509d360e11b815260040160405180910390fd5b60005b84811015610aab57600084848381811061094457610944611c0d565b61095a9260206040909202019081019150611c23565b6001600160a01b0316036109815760405163d92e233d60e01b815260040160405180910390fd5b83838281811061099357610993611c0d565b905060400201603460008888858181106109af576109af611c0d565b90506020020160208101906109c49190611a5b565b6001600160401b0316815260208101919091526040016000206109e78282611c40565b9050508383828181106109fc576109fc611c0d565b9050604002016020016020810190610a149190611c99565b1515848483818110610a2857610a28611c0d565b610a3e9260206040909202019081019150611c23565b6001600160a01b0316878784818110610a5957610a59611c0d565b9050602002016020810190610a6e9190611a5b565b6001600160401b03167fab19e99f8223191275fefd1410893ea2b3001028e27ab75b975987c1c8c4320760405160405180910390a4600101610928565b505050505050565b60005460408051631799380760e11b815281516001600160a01b0390931692632f32700e926004808401939192918290030181865afa158015610afa573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610b1e9190611cb6565b8051600180546020909301516001600160401b039092166001600160e01b031990931692909217600160401b6001600160a01b039283160217909155600054163303610be0576001546001600160401b0381166000908152603460205260409020546001600160a01b03908116600160401b9092041614610bdb57600154604051633dfc334560e01b81526001600160401b0382166004820152600160401b9091046001600160a01b031660248201526044015b60405180910390fd5b610c0a565b604051633dfc334560e01b81526001600160401b0346166004820152336024820152604401610bd2565b610c148282610f86565b5050600180546001600160e01b0319169055565b610c30611177565b610c3c848484846111a8565b6107958484848461126e565b610c5182610757565b610c5a81610dd8565b6107958383610de2565b610c6c6113cf565b565b610c766113cf565b610c6c611418565b610c866113cf565b610c8f82611439565b610c98816114d2565b5050565b6000600080516020611e57833981519152610cb784846108c2565b610d37576000848152602082815260408083206001600160a01b03871684529091529020805460ff19166001179055610ced3390565b6001600160a01b0316836001600160a01b0316857f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a4600191505061048b565b600091505061048b565b816014528060345263095ea7b360601b60005260206000604460106000875af18060016000511416610dcd57803d853b151710610dcd57600060345263095ea7b360601b600052600038604460106000885af1508160345260206000604460106000885af190508060016000511416610dcd57803d853b151710610dcd57633e3f8f736000526004601cfd5b506000603452505050565b6107f38133611575565b6000600080516020611e57833981519152610dfd84846108c2565b15610d37576000848152602082815260408083206001600160a01b0387168085529252808320805460ff1916905551339287917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a4600191505061048b565b610e666115ae565b600080516020611e77833981519152805460ff191681557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa335b6040516001600160a01b0390911681526020015b60405180910390a150565b610ec7611177565b600080516020611e77833981519152805460ff191660011781557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a25833610ea0565b60008054604051632376548f60e21b81526001600160a01b0390911690638dd9523c90610f3d90879087908790600401611d6a565b602060405180830381865afa158015610f5a573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610f7e9190611da5565b949350505050565b6032546033546001600160a01b03918216911680611005576040516340c10f1960e01b81526001600160a01b038581166004830152602482018590528316906340c10f1990604401600060405180830381600087803b158015610fe857600080fd5b505af1158015610ffc573d6000803e3d6000fd5b50505050611126565b6040516340c10f1960e01b8152306004820152602481018490526001600160a01b038316906340c10f1990604401600060405180830381600087803b15801561104d57600080fd5b505af1158015611061573d6000803e3d6000fd5b505060405163040b850f60e31b81526001600160a01b038781166004830152602482018790528416925063205c28789150604401600060405180830381600087803b1580156110af57600080fd5b505af19250505080156110c0575060015b611126576110d86001600160a01b03831685856115de565b836001600160a01b0316816001600160a01b03167fd5684d2e31a0d2443b1102269ffecfa6d05a6a36e312219f9108c0e95f8bc9af8560405161111d91815260200190565b60405180910390a35b6001546040518481526001600160a01b038616916001600160401b0316907ffd6bb19aab6d8ed3fbc8848d3b1a76c95135e46d05bec2b84424219f9681929f9060200160405180910390a350505050565b600080516020611e778339815191525460ff1615610c6c5760405163d93c066560e01b815260040160405180910390fd5b6001600160401b0384166000908152603460205260409020546001600160a01b03166111f25760405163f6e829e160e01b81526001600160401b0385166004820152602401610bd2565b6001600160a01b0383166112195760405163d92e233d60e01b815260040160405180910390fd5b8160000361123a57604051631f2a200560e01b815260040160405180910390fd5b80801561125057506033546001600160a01b0316155b1561079557604051637d25d4c960e11b815260040160405180910390fd5b6032546033546001600160a01b039182169116821561136457611300333086846001600160a01b031663fc0c546a6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156112cb573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906112ef9190611bf0565b6001600160a01b0316929190611623565b60405160016255295b60e01b03198152336004820152602481018590526001600160a01b0382169063ffaad6a590604401600060405180830381600087803b15801561134b57600080fd5b505af115801561135f573d6000803e3d6000fd5b505050505b6040516379ef98bf60e11b8152336004820152602481018590526001600160a01b0383169063f3df317e90604401600060405180830381600087803b1580156113ac57600080fd5b505af11580156113c0573d6000803e3d6000fd5b50505050610aab868686611681565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a0054600160401b900460ff16610c6c57604051631afcd79f60e31b815260040160405180910390fd5b6114206113cf565b600080516020611e77833981519152805460ff19169055565b6001600160a01b0381166114845760405162461bcd60e51b8152602060048201526012602482015271584170703a206e6f207a65726f206f6d6e6960701b6044820152606401610bd2565b600080546001600160a01b0319166001600160a01b0383169081179091556040519081527f79162c8d053a07e70cdc1ccc536f0440b571f8508377d2bef51094fadab98f4790602001610eb4565b6114db816117cf565b6115275760405162461bcd60e51b815260206004820152601860248201527f584170703a20696e76616c696420636f6e66206c6576656c00000000000000006044820152606401610bd2565b6000805460ff60a01b1916600160a01b60ff8416908102919091179091556040519081527f8de08a798b4e50b4f351c1eaa91a11530043802be3ffac2df87db0c45a2e848390602001610eb4565b61157f82826108c2565b610c985760405163e2517d3f60e01b81526001600160a01b038216600482015260248101839052604401610bd2565b600080516020611e778339815191525460ff16610c6c57604051638dfc202b60e01b815260040160405180910390fd5b816014528060345263a9059cbb60601b60005260206000604460106000875af18060016000511416610dcd57803d853b151710610dcd576390b8ec186000526004601cfd5b60405181606052826040528360601b602c526323b872dd60601b600c52602060006064601c6000895af1806001600051141661167257803d873b15171061167257637939f4246000526004601cfd5b50600060605260405250505050565b6001600160401b03831660009081526034602090815260408083208151808301835290546001600160a01b038082168352600160a01b90910460ff161515938201939093529051918516602483015260448201849052919060640160408051601f19818403018152919052602080820180516001600160e01b0316634b91ad0f60e11b1790528351908401519192506000916117319188918590611728576201e8486117eb565b6202bf206117eb565b9050803410156117545760405163cd1c886760e01b815260040160405180910390fd5b80341115611770576117706117698234611dbe565b3390611931565b846001600160a01b0316336001600160a01b0316876001600160401b03167fc0464e720761b6de8643eed8d1cbf17ec66c3eb60179efaed5a58cf75580a4dc876040516117bf91815260200190565b60405180910390a4505050505050565b600060ff82166001148061048b575060ff821660041492915050565b60008054604051632376548f60e21b815282916001600160a01b031690638dd9523c9061182090899088908890600401611d6a565b602060405180830381865afa15801561183d573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906118619190611da5565b9050804710156118b35760405162461bcd60e51b815260206004820152601860248201527f584170703a20696e73756666696369656e742066756e647300000000000000006044820152606401610bd2565b60005460405163c21dda4f60e01b81526001600160a01b0382169163c21dda4f9184916118f5918b91600160a01b900460ff16908b908b908b90600401611ddf565b6000604051808303818588803b15801561190e57600080fd5b505af1158015611922573d6000803e3d6000fd5b50939998505050505050505050565b60003860003884865af1610c985763b12d13eb6000526004601cfd5b60006020828403121561195f57600080fd5b81356001600160e01b0319811681146108bb57600080fd5b6001600160a01b03811681146107f357600080fd5b600080600080600060a086880312156119a457600080fd5b85356119af81611977565b945060208601356119bf81611977565b935060408601356119cf81611977565b925060608601356119df81611977565b915060808601356119ef81611977565b809150509295509295909350565b600060208284031215611a0f57600080fd5b5035919050565b60008060408385031215611a2957600080fd5b823591506020830135611a3b81611977565b809150509250929050565b6001600160401b03811681146107f357600080fd5b600060208284031215611a6d57600080fd5b81356108bb81611a46565b60008083601f840112611a8a57600080fd5b5081356001600160401b03811115611aa157600080fd5b6020830191508360208260061b8501011115611abc57600080fd5b9250929050565b60008060008060408587031215611ad957600080fd5b84356001600160401b03811115611aef57600080fd5b8501601f81018713611b0057600080fd5b80356001600160401b03811115611b1657600080fd5b8760208260051b8401011115611b2b57600080fd5b6020918201955093508501356001600160401b03811115611b4b57600080fd5b611b5787828801611a78565b95989497509550505050565b60008060408385031215611b7657600080fd5b8235611b8181611977565b946020939093013593505050565b80151581146107f357600080fd5b60008060008060808587031215611bb357600080fd5b8435611bbe81611a46565b93506020850135611bce81611977565b9250604085013591506060850135611be581611b8f565b939692955090935050565b600060208284031215611c0257600080fd5b81516108bb81611977565b634e487b7160e01b600052603260045260246000fd5b600060208284031215611c3557600080fd5b81356108bb81611977565b8135611c4b81611977565b81546001600160a01b031981166001600160a01b039290921691821783556020840135611c7781611b8f565b6001600160a81b03199190911690911790151560a01b60ff60a01b1617905550565b600060208284031215611cab57600080fd5b81356108bb81611b8f565b60006040828403128015611cc957600080fd5b50604080519081016001600160401b0381118282101715611cfa57634e487b7160e01b600052604160045260246000fd5b6040528251611d0881611a46565b81526020830151611d1881611977565b60208201529392505050565b6000815180845260005b81811015611d4a57602081850181015186830182015201611d2e565b506000602082860101526020601f19601f83011685010191505092915050565b6001600160401b0384168152606060208201526000611d8c6060830185611d24565b90506001600160401b0383166040830152949350505050565b600060208284031215611db757600080fd5b5051919050565b8181038181111561048b57634e487b7160e01b600052601160045260246000fd5b6001600160401b038616815260ff851660208201526001600160a01b038416604082015260a060608201819052600090611e1b90830185611d24565b90506001600160401b0383166080830152969550505050505056fe539440820030c4994db4e31b6b800deafd503688728f932addfe7a410515c14c02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800cd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f03300a264697066735822122033a2497ad4bc1114803fc2ac85b79fc27ca9b5ac2d56a63156e561ed5b9cecb664736f6c634300081a0033",
}

// BridgeABI is the input ABI used to generate the binding from.
// Deprecated: Use BridgeMetaData.ABI instead.
var BridgeABI = BridgeMetaData.ABI

// BridgeBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use BridgeMetaData.Bin instead.
var BridgeBin = BridgeMetaData.Bin

// DeployBridge deploys a new Ethereum contract, binding an instance of Bridge to it.
func DeployBridge(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Bridge, error) {
	parsed, err := BridgeMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BridgeBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Bridge{BridgeCaller: BridgeCaller{contract: contract}, BridgeTransactor: BridgeTransactor{contract: contract}, BridgeFilterer: BridgeFilterer{contract: contract}}, nil
}

// Bridge is an auto generated Go binding around an Ethereum contract.
type Bridge struct {
	BridgeCaller     // Read-only binding to the contract
	BridgeTransactor // Write-only binding to the contract
	BridgeFilterer   // Log filterer for contract events
}

// BridgeCaller is an auto generated read-only Go binding around an Ethereum contract.
type BridgeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BridgeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BridgeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BridgeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BridgeSession struct {
	Contract     *Bridge           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BridgeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BridgeCallerSession struct {
	Contract *BridgeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// BridgeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BridgeTransactorSession struct {
	Contract     *BridgeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BridgeRaw is an auto generated low-level Go binding around an Ethereum contract.
type BridgeRaw struct {
	Contract *Bridge // Generic contract binding to access the raw methods on
}

// BridgeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BridgeCallerRaw struct {
	Contract *BridgeCaller // Generic read-only contract binding to access the raw methods on
}

// BridgeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BridgeTransactorRaw struct {
	Contract *BridgeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBridge creates a new instance of Bridge, bound to a specific deployed contract.
func NewBridge(address common.Address, backend bind.ContractBackend) (*Bridge, error) {
	contract, err := bindBridge(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Bridge{BridgeCaller: BridgeCaller{contract: contract}, BridgeTransactor: BridgeTransactor{contract: contract}, BridgeFilterer: BridgeFilterer{contract: contract}}, nil
}

// NewBridgeCaller creates a new read-only instance of Bridge, bound to a specific deployed contract.
func NewBridgeCaller(address common.Address, caller bind.ContractCaller) (*BridgeCaller, error) {
	contract, err := bindBridge(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BridgeCaller{contract: contract}, nil
}

// NewBridgeTransactor creates a new write-only instance of Bridge, bound to a specific deployed contract.
func NewBridgeTransactor(address common.Address, transactor bind.ContractTransactor) (*BridgeTransactor, error) {
	contract, err := bindBridge(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BridgeTransactor{contract: contract}, nil
}

// NewBridgeFilterer creates a new log filterer instance of Bridge, bound to a specific deployed contract.
func NewBridgeFilterer(address common.Address, filterer bind.ContractFilterer) (*BridgeFilterer, error) {
	contract, err := bindBridge(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BridgeFilterer{contract: contract}, nil
}

// bindBridge binds a generic wrapper to an already deployed contract.
func bindBridge(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BridgeMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bridge *BridgeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bridge.Contract.BridgeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bridge *BridgeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bridge.Contract.BridgeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bridge *BridgeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bridge.Contract.BridgeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Bridge *BridgeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Bridge.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Bridge *BridgeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bridge.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Bridge *BridgeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Bridge.Contract.contract.Transact(opts, method, params...)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Bridge *BridgeCaller) DEFAULTADMINROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "DEFAULT_ADMIN_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Bridge *BridgeSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Bridge.Contract.DEFAULTADMINROLE(&_Bridge.CallOpts)
}

// DEFAULTADMINROLE is a free data retrieval call binding the contract method 0xa217fddf.
//
// Solidity: function DEFAULT_ADMIN_ROLE() view returns(bytes32)
func (_Bridge *BridgeCallerSession) DEFAULTADMINROLE() ([32]byte, error) {
	return _Bridge.Contract.DEFAULTADMINROLE(&_Bridge.CallOpts)
}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_Bridge *BridgeCaller) PAUSERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "PAUSER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_Bridge *BridgeSession) PAUSERROLE() ([32]byte, error) {
	return _Bridge.Contract.PAUSERROLE(&_Bridge.CallOpts)
}

// PAUSERROLE is a free data retrieval call binding the contract method 0xe63ab1e9.
//
// Solidity: function PAUSER_ROLE() view returns(bytes32)
func (_Bridge *BridgeCallerSession) PAUSERROLE() ([32]byte, error) {
	return _Bridge.Contract.PAUSERROLE(&_Bridge.CallOpts)
}

// BridgeFee is a free data retrieval call binding the contract method 0x8ad201f6.
//
// Solidity: function bridgeFee(uint64 destChainId) view returns(uint256 fee)
func (_Bridge *BridgeCaller) BridgeFee(opts *bind.CallOpts, destChainId uint64) (*big.Int, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "bridgeFee", destChainId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BridgeFee is a free data retrieval call binding the contract method 0x8ad201f6.
//
// Solidity: function bridgeFee(uint64 destChainId) view returns(uint256 fee)
func (_Bridge *BridgeSession) BridgeFee(destChainId uint64) (*big.Int, error) {
	return _Bridge.Contract.BridgeFee(&_Bridge.CallOpts, destChainId)
}

// BridgeFee is a free data retrieval call binding the contract method 0x8ad201f6.
//
// Solidity: function bridgeFee(uint64 destChainId) view returns(uint256 fee)
func (_Bridge *BridgeCallerSession) BridgeFee(destChainId uint64) (*big.Int, error) {
	return _Bridge.Contract.BridgeFee(&_Bridge.CallOpts, destChainId)
}

// DefaultConfLevel is a free data retrieval call binding the contract method 0x74eeb847.
//
// Solidity: function defaultConfLevel() view returns(uint8)
func (_Bridge *BridgeCaller) DefaultConfLevel(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "defaultConfLevel")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// DefaultConfLevel is a free data retrieval call binding the contract method 0x74eeb847.
//
// Solidity: function defaultConfLevel() view returns(uint8)
func (_Bridge *BridgeSession) DefaultConfLevel() (uint8, error) {
	return _Bridge.Contract.DefaultConfLevel(&_Bridge.CallOpts)
}

// DefaultConfLevel is a free data retrieval call binding the contract method 0x74eeb847.
//
// Solidity: function defaultConfLevel() view returns(uint8)
func (_Bridge *BridgeCallerSession) DefaultConfLevel() (uint8, error) {
	return _Bridge.Contract.DefaultConfLevel(&_Bridge.CallOpts)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Bridge *BridgeCaller) GetRoleAdmin(opts *bind.CallOpts, role [32]byte) ([32]byte, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "getRoleAdmin", role)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Bridge *BridgeSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Bridge.Contract.GetRoleAdmin(&_Bridge.CallOpts, role)
}

// GetRoleAdmin is a free data retrieval call binding the contract method 0x248a9ca3.
//
// Solidity: function getRoleAdmin(bytes32 role) view returns(bytes32)
func (_Bridge *BridgeCallerSession) GetRoleAdmin(role [32]byte) ([32]byte, error) {
	return _Bridge.Contract.GetRoleAdmin(&_Bridge.CallOpts, role)
}

// GetRoute is a free data retrieval call binding the contract method 0x4acd8d82.
//
// Solidity: function getRoute(uint64 destChainId) view returns(address bridge, bool hasLockbox)
func (_Bridge *BridgeCaller) GetRoute(opts *bind.CallOpts, destChainId uint64) (struct {
	Bridge     common.Address
	HasLockbox bool
}, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "getRoute", destChainId)

	outstruct := new(struct {
		Bridge     common.Address
		HasLockbox bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Bridge = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.HasLockbox = *abi.ConvertType(out[1], new(bool)).(*bool)

	return *outstruct, err

}

// GetRoute is a free data retrieval call binding the contract method 0x4acd8d82.
//
// Solidity: function getRoute(uint64 destChainId) view returns(address bridge, bool hasLockbox)
func (_Bridge *BridgeSession) GetRoute(destChainId uint64) (struct {
	Bridge     common.Address
	HasLockbox bool
}, error) {
	return _Bridge.Contract.GetRoute(&_Bridge.CallOpts, destChainId)
}

// GetRoute is a free data retrieval call binding the contract method 0x4acd8d82.
//
// Solidity: function getRoute(uint64 destChainId) view returns(address bridge, bool hasLockbox)
func (_Bridge *BridgeCallerSession) GetRoute(destChainId uint64) (struct {
	Bridge     common.Address
	HasLockbox bool
}, error) {
	return _Bridge.Contract.GetRoute(&_Bridge.CallOpts, destChainId)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Bridge *BridgeCaller) HasRole(opts *bind.CallOpts, role [32]byte, account common.Address) (bool, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "hasRole", role, account)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Bridge *BridgeSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Bridge.Contract.HasRole(&_Bridge.CallOpts, role, account)
}

// HasRole is a free data retrieval call binding the contract method 0x91d14854.
//
// Solidity: function hasRole(bytes32 role, address account) view returns(bool)
func (_Bridge *BridgeCallerSession) HasRole(role [32]byte, account common.Address) (bool, error) {
	return _Bridge.Contract.HasRole(&_Bridge.CallOpts, role, account)
}

// Lockbox is a free data retrieval call binding the contract method 0x66cc5702.
//
// Solidity: function lockbox() view returns(address)
func (_Bridge *BridgeCaller) Lockbox(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "lockbox")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Lockbox is a free data retrieval call binding the contract method 0x66cc5702.
//
// Solidity: function lockbox() view returns(address)
func (_Bridge *BridgeSession) Lockbox() (common.Address, error) {
	return _Bridge.Contract.Lockbox(&_Bridge.CallOpts)
}

// Lockbox is a free data retrieval call binding the contract method 0x66cc5702.
//
// Solidity: function lockbox() view returns(address)
func (_Bridge *BridgeCallerSession) Lockbox() (common.Address, error) {
	return _Bridge.Contract.Lockbox(&_Bridge.CallOpts)
}

// Omni is a free data retrieval call binding the contract method 0x39acf9f1.
//
// Solidity: function omni() view returns(address)
func (_Bridge *BridgeCaller) Omni(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "omni")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Omni is a free data retrieval call binding the contract method 0x39acf9f1.
//
// Solidity: function omni() view returns(address)
func (_Bridge *BridgeSession) Omni() (common.Address, error) {
	return _Bridge.Contract.Omni(&_Bridge.CallOpts)
}

// Omni is a free data retrieval call binding the contract method 0x39acf9f1.
//
// Solidity: function omni() view returns(address)
func (_Bridge *BridgeCallerSession) Omni() (common.Address, error) {
	return _Bridge.Contract.Omni(&_Bridge.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Bridge *BridgeCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Bridge *BridgeSession) Paused() (bool, error) {
	return _Bridge.Contract.Paused(&_Bridge.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_Bridge *BridgeCallerSession) Paused() (bool, error) {
	return _Bridge.Contract.Paused(&_Bridge.CallOpts)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Bridge *BridgeCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Bridge *BridgeSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Bridge.Contract.SupportsInterface(&_Bridge.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_Bridge *BridgeCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _Bridge.Contract.SupportsInterface(&_Bridge.CallOpts, interfaceId)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Bridge *BridgeCaller) Token(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "token")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Bridge *BridgeSession) Token() (common.Address, error) {
	return _Bridge.Contract.Token(&_Bridge.CallOpts)
}

// Token is a free data retrieval call binding the contract method 0xfc0c546a.
//
// Solidity: function token() view returns(address)
func (_Bridge *BridgeCallerSession) Token() (common.Address, error) {
	return _Bridge.Contract.Token(&_Bridge.CallOpts)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Bridge *BridgeTransactor) GrantRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "grantRole", role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Bridge *BridgeSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.GrantRole(&_Bridge.TransactOpts, role, account)
}

// GrantRole is a paid mutator transaction binding the contract method 0x2f2ff15d.
//
// Solidity: function grantRole(bytes32 role, address account) returns()
func (_Bridge *BridgeTransactorSession) GrantRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.GrantRole(&_Bridge.TransactOpts, role, account)
}

// Initialize is a paid mutator transaction binding the contract method 0x1459457a.
//
// Solidity: function initialize(address admin_, address pauser_, address omni_, address token_, address lockbox_) returns()
func (_Bridge *BridgeTransactor) Initialize(opts *bind.TransactOpts, admin_ common.Address, pauser_ common.Address, omni_ common.Address, token_ common.Address, lockbox_ common.Address) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "initialize", admin_, pauser_, omni_, token_, lockbox_)
}

// Initialize is a paid mutator transaction binding the contract method 0x1459457a.
//
// Solidity: function initialize(address admin_, address pauser_, address omni_, address token_, address lockbox_) returns()
func (_Bridge *BridgeSession) Initialize(admin_ common.Address, pauser_ common.Address, omni_ common.Address, token_ common.Address, lockbox_ common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.Initialize(&_Bridge.TransactOpts, admin_, pauser_, omni_, token_, lockbox_)
}

// Initialize is a paid mutator transaction binding the contract method 0x1459457a.
//
// Solidity: function initialize(address admin_, address pauser_, address omni_, address token_, address lockbox_) returns()
func (_Bridge *BridgeTransactorSession) Initialize(admin_ common.Address, pauser_ common.Address, omni_ common.Address, token_ common.Address, lockbox_ common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.Initialize(&_Bridge.TransactOpts, admin_, pauser_, omni_, token_, lockbox_)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Bridge *BridgeTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Bridge *BridgeSession) Pause() (*types.Transaction, error) {
	return _Bridge.Contract.Pause(&_Bridge.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Bridge *BridgeTransactorSession) Pause() (*types.Transaction, error) {
	return _Bridge.Contract.Pause(&_Bridge.TransactOpts)
}

// ReceiveToken is a paid mutator transaction binding the contract method 0x97235a1e.
//
// Solidity: function receiveToken(address to, uint256 value) returns()
func (_Bridge *BridgeTransactor) ReceiveToken(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "receiveToken", to, value)
}

// ReceiveToken is a paid mutator transaction binding the contract method 0x97235a1e.
//
// Solidity: function receiveToken(address to, uint256 value) returns()
func (_Bridge *BridgeSession) ReceiveToken(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.ReceiveToken(&_Bridge.TransactOpts, to, value)
}

// ReceiveToken is a paid mutator transaction binding the contract method 0x97235a1e.
//
// Solidity: function receiveToken(address to, uint256 value) returns()
func (_Bridge *BridgeTransactorSession) ReceiveToken(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Bridge.Contract.ReceiveToken(&_Bridge.TransactOpts, to, value)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_Bridge *BridgeTransactor) RenounceRole(opts *bind.TransactOpts, role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "renounceRole", role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_Bridge *BridgeSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.RenounceRole(&_Bridge.TransactOpts, role, callerConfirmation)
}

// RenounceRole is a paid mutator transaction binding the contract method 0x36568abe.
//
// Solidity: function renounceRole(bytes32 role, address callerConfirmation) returns()
func (_Bridge *BridgeTransactorSession) RenounceRole(role [32]byte, callerConfirmation common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.RenounceRole(&_Bridge.TransactOpts, role, callerConfirmation)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Bridge *BridgeTransactor) RevokeRole(opts *bind.TransactOpts, role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "revokeRole", role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Bridge *BridgeSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.RevokeRole(&_Bridge.TransactOpts, role, account)
}

// RevokeRole is a paid mutator transaction binding the contract method 0xd547741f.
//
// Solidity: function revokeRole(bytes32 role, address account) returns()
func (_Bridge *BridgeTransactorSession) RevokeRole(role [32]byte, account common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.RevokeRole(&_Bridge.TransactOpts, role, account)
}

// SendToken is a paid mutator transaction binding the contract method 0xa48fcd8e.
//
// Solidity: function sendToken(uint64 destChainId, address to, uint256 value, bool wrap) payable returns()
func (_Bridge *BridgeTransactor) SendToken(opts *bind.TransactOpts, destChainId uint64, to common.Address, value *big.Int, wrap bool) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "sendToken", destChainId, to, value, wrap)
}

// SendToken is a paid mutator transaction binding the contract method 0xa48fcd8e.
//
// Solidity: function sendToken(uint64 destChainId, address to, uint256 value, bool wrap) payable returns()
func (_Bridge *BridgeSession) SendToken(destChainId uint64, to common.Address, value *big.Int, wrap bool) (*types.Transaction, error) {
	return _Bridge.Contract.SendToken(&_Bridge.TransactOpts, destChainId, to, value, wrap)
}

// SendToken is a paid mutator transaction binding the contract method 0xa48fcd8e.
//
// Solidity: function sendToken(uint64 destChainId, address to, uint256 value, bool wrap) payable returns()
func (_Bridge *BridgeTransactorSession) SendToken(destChainId uint64, to common.Address, value *big.Int, wrap bool) (*types.Transaction, error) {
	return _Bridge.Contract.SendToken(&_Bridge.TransactOpts, destChainId, to, value, wrap)
}

// SetRoutes is a paid mutator transaction binding the contract method 0x92099d44.
//
// Solidity: function setRoutes(uint64[] chainIds, (address,bool)[] routes) returns()
func (_Bridge *BridgeTransactor) SetRoutes(opts *bind.TransactOpts, chainIds []uint64, routes []IBridgeRoute) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "setRoutes", chainIds, routes)
}

// SetRoutes is a paid mutator transaction binding the contract method 0x92099d44.
//
// Solidity: function setRoutes(uint64[] chainIds, (address,bool)[] routes) returns()
func (_Bridge *BridgeSession) SetRoutes(chainIds []uint64, routes []IBridgeRoute) (*types.Transaction, error) {
	return _Bridge.Contract.SetRoutes(&_Bridge.TransactOpts, chainIds, routes)
}

// SetRoutes is a paid mutator transaction binding the contract method 0x92099d44.
//
// Solidity: function setRoutes(uint64[] chainIds, (address,bool)[] routes) returns()
func (_Bridge *BridgeTransactorSession) SetRoutes(chainIds []uint64, routes []IBridgeRoute) (*types.Transaction, error) {
	return _Bridge.Contract.SetRoutes(&_Bridge.TransactOpts, chainIds, routes)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Bridge *BridgeTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Bridge *BridgeSession) Unpause() (*types.Transaction, error) {
	return _Bridge.Contract.Unpause(&_Bridge.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_Bridge *BridgeTransactorSession) Unpause() (*types.Transaction, error) {
	return _Bridge.Contract.Unpause(&_Bridge.TransactOpts)
}

// BridgeDefaultConfLevelSetIterator is returned from FilterDefaultConfLevelSet and is used to iterate over the raw logs and unpacked data for DefaultConfLevelSet events raised by the Bridge contract.
type BridgeDefaultConfLevelSetIterator struct {
	Event *BridgeDefaultConfLevelSet // Event containing the contract specifics and raw log

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
func (it *BridgeDefaultConfLevelSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeDefaultConfLevelSet)
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
		it.Event = new(BridgeDefaultConfLevelSet)
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
func (it *BridgeDefaultConfLevelSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeDefaultConfLevelSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeDefaultConfLevelSet represents a DefaultConfLevelSet event raised by the Bridge contract.
type BridgeDefaultConfLevelSet struct {
	Conf uint8
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterDefaultConfLevelSet is a free log retrieval operation binding the contract event 0x8de08a798b4e50b4f351c1eaa91a11530043802be3ffac2df87db0c45a2e8483.
//
// Solidity: event DefaultConfLevelSet(uint8 conf)
func (_Bridge *BridgeFilterer) FilterDefaultConfLevelSet(opts *bind.FilterOpts) (*BridgeDefaultConfLevelSetIterator, error) {

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "DefaultConfLevelSet")
	if err != nil {
		return nil, err
	}
	return &BridgeDefaultConfLevelSetIterator{contract: _Bridge.contract, event: "DefaultConfLevelSet", logs: logs, sub: sub}, nil
}

// WatchDefaultConfLevelSet is a free log subscription operation binding the contract event 0x8de08a798b4e50b4f351c1eaa91a11530043802be3ffac2df87db0c45a2e8483.
//
// Solidity: event DefaultConfLevelSet(uint8 conf)
func (_Bridge *BridgeFilterer) WatchDefaultConfLevelSet(opts *bind.WatchOpts, sink chan<- *BridgeDefaultConfLevelSet) (event.Subscription, error) {

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "DefaultConfLevelSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeDefaultConfLevelSet)
				if err := _Bridge.contract.UnpackLog(event, "DefaultConfLevelSet", log); err != nil {
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
func (_Bridge *BridgeFilterer) ParseDefaultConfLevelSet(log types.Log) (*BridgeDefaultConfLevelSet, error) {
	event := new(BridgeDefaultConfLevelSet)
	if err := _Bridge.contract.UnpackLog(event, "DefaultConfLevelSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Bridge contract.
type BridgeInitializedIterator struct {
	Event *BridgeInitialized // Event containing the contract specifics and raw log

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
func (it *BridgeInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeInitialized)
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
		it.Event = new(BridgeInitialized)
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
func (it *BridgeInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeInitialized represents a Initialized event raised by the Bridge contract.
type BridgeInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Bridge *BridgeFilterer) FilterInitialized(opts *bind.FilterOpts) (*BridgeInitializedIterator, error) {

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &BridgeInitializedIterator{contract: _Bridge.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Bridge *BridgeFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *BridgeInitialized) (event.Subscription, error) {

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeInitialized)
				if err := _Bridge.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Bridge *BridgeFilterer) ParseInitialized(log types.Log) (*BridgeInitialized, error) {
	event := new(BridgeInitialized)
	if err := _Bridge.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeLockboxWithdrawalFailedIterator is returned from FilterLockboxWithdrawalFailed and is used to iterate over the raw logs and unpacked data for LockboxWithdrawalFailed events raised by the Bridge contract.
type BridgeLockboxWithdrawalFailedIterator struct {
	Event *BridgeLockboxWithdrawalFailed // Event containing the contract specifics and raw log

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
func (it *BridgeLockboxWithdrawalFailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeLockboxWithdrawalFailed)
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
		it.Event = new(BridgeLockboxWithdrawalFailed)
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
func (it *BridgeLockboxWithdrawalFailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeLockboxWithdrawalFailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeLockboxWithdrawalFailed represents a LockboxWithdrawalFailed event raised by the Bridge contract.
type BridgeLockboxWithdrawalFailed struct {
	BadLockbox common.Address
	To         common.Address
	Value      *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterLockboxWithdrawalFailed is a free log retrieval operation binding the contract event 0xd5684d2e31a0d2443b1102269ffecfa6d05a6a36e312219f9108c0e95f8bc9af.
//
// Solidity: event LockboxWithdrawalFailed(address indexed badLockbox, address indexed to, uint256 value)
func (_Bridge *BridgeFilterer) FilterLockboxWithdrawalFailed(opts *bind.FilterOpts, badLockbox []common.Address, to []common.Address) (*BridgeLockboxWithdrawalFailedIterator, error) {

	var badLockboxRule []interface{}
	for _, badLockboxItem := range badLockbox {
		badLockboxRule = append(badLockboxRule, badLockboxItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "LockboxWithdrawalFailed", badLockboxRule, toRule)
	if err != nil {
		return nil, err
	}
	return &BridgeLockboxWithdrawalFailedIterator{contract: _Bridge.contract, event: "LockboxWithdrawalFailed", logs: logs, sub: sub}, nil
}

// WatchLockboxWithdrawalFailed is a free log subscription operation binding the contract event 0xd5684d2e31a0d2443b1102269ffecfa6d05a6a36e312219f9108c0e95f8bc9af.
//
// Solidity: event LockboxWithdrawalFailed(address indexed badLockbox, address indexed to, uint256 value)
func (_Bridge *BridgeFilterer) WatchLockboxWithdrawalFailed(opts *bind.WatchOpts, sink chan<- *BridgeLockboxWithdrawalFailed, badLockbox []common.Address, to []common.Address) (event.Subscription, error) {

	var badLockboxRule []interface{}
	for _, badLockboxItem := range badLockbox {
		badLockboxRule = append(badLockboxRule, badLockboxItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "LockboxWithdrawalFailed", badLockboxRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeLockboxWithdrawalFailed)
				if err := _Bridge.contract.UnpackLog(event, "LockboxWithdrawalFailed", log); err != nil {
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

// ParseLockboxWithdrawalFailed is a log parse operation binding the contract event 0xd5684d2e31a0d2443b1102269ffecfa6d05a6a36e312219f9108c0e95f8bc9af.
//
// Solidity: event LockboxWithdrawalFailed(address indexed badLockbox, address indexed to, uint256 value)
func (_Bridge *BridgeFilterer) ParseLockboxWithdrawalFailed(log types.Log) (*BridgeLockboxWithdrawalFailed, error) {
	event := new(BridgeLockboxWithdrawalFailed)
	if err := _Bridge.contract.UnpackLog(event, "LockboxWithdrawalFailed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeOmniPortalSetIterator is returned from FilterOmniPortalSet and is used to iterate over the raw logs and unpacked data for OmniPortalSet events raised by the Bridge contract.
type BridgeOmniPortalSetIterator struct {
	Event *BridgeOmniPortalSet // Event containing the contract specifics and raw log

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
func (it *BridgeOmniPortalSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeOmniPortalSet)
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
		it.Event = new(BridgeOmniPortalSet)
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
func (it *BridgeOmniPortalSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeOmniPortalSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeOmniPortalSet represents a OmniPortalSet event raised by the Bridge contract.
type BridgeOmniPortalSet struct {
	Omni common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterOmniPortalSet is a free log retrieval operation binding the contract event 0x79162c8d053a07e70cdc1ccc536f0440b571f8508377d2bef51094fadab98f47.
//
// Solidity: event OmniPortalSet(address omni)
func (_Bridge *BridgeFilterer) FilterOmniPortalSet(opts *bind.FilterOpts) (*BridgeOmniPortalSetIterator, error) {

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "OmniPortalSet")
	if err != nil {
		return nil, err
	}
	return &BridgeOmniPortalSetIterator{contract: _Bridge.contract, event: "OmniPortalSet", logs: logs, sub: sub}, nil
}

// WatchOmniPortalSet is a free log subscription operation binding the contract event 0x79162c8d053a07e70cdc1ccc536f0440b571f8508377d2bef51094fadab98f47.
//
// Solidity: event OmniPortalSet(address omni)
func (_Bridge *BridgeFilterer) WatchOmniPortalSet(opts *bind.WatchOpts, sink chan<- *BridgeOmniPortalSet) (event.Subscription, error) {

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "OmniPortalSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeOmniPortalSet)
				if err := _Bridge.contract.UnpackLog(event, "OmniPortalSet", log); err != nil {
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
func (_Bridge *BridgeFilterer) ParseOmniPortalSet(log types.Log) (*BridgeOmniPortalSet, error) {
	event := new(BridgeOmniPortalSet)
	if err := _Bridge.contract.UnpackLog(event, "OmniPortalSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgePausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the Bridge contract.
type BridgePausedIterator struct {
	Event *BridgePaused // Event containing the contract specifics and raw log

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
func (it *BridgePausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgePaused)
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
		it.Event = new(BridgePaused)
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
func (it *BridgePausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgePausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgePaused represents a Paused event raised by the Bridge contract.
type BridgePaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Bridge *BridgeFilterer) FilterPaused(opts *bind.FilterOpts) (*BridgePausedIterator, error) {

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &BridgePausedIterator{contract: _Bridge.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_Bridge *BridgeFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *BridgePaused) (event.Subscription, error) {

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgePaused)
				if err := _Bridge.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_Bridge *BridgeFilterer) ParsePaused(log types.Log) (*BridgePaused, error) {
	event := new(BridgePaused)
	if err := _Bridge.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeRoleAdminChangedIterator is returned from FilterRoleAdminChanged and is used to iterate over the raw logs and unpacked data for RoleAdminChanged events raised by the Bridge contract.
type BridgeRoleAdminChangedIterator struct {
	Event *BridgeRoleAdminChanged // Event containing the contract specifics and raw log

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
func (it *BridgeRoleAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeRoleAdminChanged)
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
		it.Event = new(BridgeRoleAdminChanged)
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
func (it *BridgeRoleAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeRoleAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeRoleAdminChanged represents a RoleAdminChanged event raised by the Bridge contract.
type BridgeRoleAdminChanged struct {
	Role              [32]byte
	PreviousAdminRole [32]byte
	NewAdminRole      [32]byte
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterRoleAdminChanged is a free log retrieval operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Bridge *BridgeFilterer) FilterRoleAdminChanged(opts *bind.FilterOpts, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (*BridgeRoleAdminChangedIterator, error) {

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

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return &BridgeRoleAdminChangedIterator{contract: _Bridge.contract, event: "RoleAdminChanged", logs: logs, sub: sub}, nil
}

// WatchRoleAdminChanged is a free log subscription operation binding the contract event 0xbd79b86ffe0ab8e8776151514217cd7cacd52c909f66475c3af44e129f0b00ff.
//
// Solidity: event RoleAdminChanged(bytes32 indexed role, bytes32 indexed previousAdminRole, bytes32 indexed newAdminRole)
func (_Bridge *BridgeFilterer) WatchRoleAdminChanged(opts *bind.WatchOpts, sink chan<- *BridgeRoleAdminChanged, role [][32]byte, previousAdminRole [][32]byte, newAdminRole [][32]byte) (event.Subscription, error) {

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

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "RoleAdminChanged", roleRule, previousAdminRoleRule, newAdminRoleRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeRoleAdminChanged)
				if err := _Bridge.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
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
func (_Bridge *BridgeFilterer) ParseRoleAdminChanged(log types.Log) (*BridgeRoleAdminChanged, error) {
	event := new(BridgeRoleAdminChanged)
	if err := _Bridge.contract.UnpackLog(event, "RoleAdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeRoleGrantedIterator is returned from FilterRoleGranted and is used to iterate over the raw logs and unpacked data for RoleGranted events raised by the Bridge contract.
type BridgeRoleGrantedIterator struct {
	Event *BridgeRoleGranted // Event containing the contract specifics and raw log

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
func (it *BridgeRoleGrantedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeRoleGranted)
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
		it.Event = new(BridgeRoleGranted)
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
func (it *BridgeRoleGrantedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeRoleGrantedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeRoleGranted represents a RoleGranted event raised by the Bridge contract.
type BridgeRoleGranted struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleGranted is a free log retrieval operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Bridge *BridgeFilterer) FilterRoleGranted(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*BridgeRoleGrantedIterator, error) {

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

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &BridgeRoleGrantedIterator{contract: _Bridge.contract, event: "RoleGranted", logs: logs, sub: sub}, nil
}

// WatchRoleGranted is a free log subscription operation binding the contract event 0x2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d.
//
// Solidity: event RoleGranted(bytes32 indexed role, address indexed account, address indexed sender)
func (_Bridge *BridgeFilterer) WatchRoleGranted(opts *bind.WatchOpts, sink chan<- *BridgeRoleGranted, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "RoleGranted", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeRoleGranted)
				if err := _Bridge.contract.UnpackLog(event, "RoleGranted", log); err != nil {
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
func (_Bridge *BridgeFilterer) ParseRoleGranted(log types.Log) (*BridgeRoleGranted, error) {
	event := new(BridgeRoleGranted)
	if err := _Bridge.contract.UnpackLog(event, "RoleGranted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeRoleRevokedIterator is returned from FilterRoleRevoked and is used to iterate over the raw logs and unpacked data for RoleRevoked events raised by the Bridge contract.
type BridgeRoleRevokedIterator struct {
	Event *BridgeRoleRevoked // Event containing the contract specifics and raw log

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
func (it *BridgeRoleRevokedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeRoleRevoked)
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
		it.Event = new(BridgeRoleRevoked)
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
func (it *BridgeRoleRevokedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeRoleRevokedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeRoleRevoked represents a RoleRevoked event raised by the Bridge contract.
type BridgeRoleRevoked struct {
	Role    [32]byte
	Account common.Address
	Sender  common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterRoleRevoked is a free log retrieval operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Bridge *BridgeFilterer) FilterRoleRevoked(opts *bind.FilterOpts, role [][32]byte, account []common.Address, sender []common.Address) (*BridgeRoleRevokedIterator, error) {

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

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return &BridgeRoleRevokedIterator{contract: _Bridge.contract, event: "RoleRevoked", logs: logs, sub: sub}, nil
}

// WatchRoleRevoked is a free log subscription operation binding the contract event 0xf6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b.
//
// Solidity: event RoleRevoked(bytes32 indexed role, address indexed account, address indexed sender)
func (_Bridge *BridgeFilterer) WatchRoleRevoked(opts *bind.WatchOpts, sink chan<- *BridgeRoleRevoked, role [][32]byte, account []common.Address, sender []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "RoleRevoked", roleRule, accountRule, senderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeRoleRevoked)
				if err := _Bridge.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
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
func (_Bridge *BridgeFilterer) ParseRoleRevoked(log types.Log) (*BridgeRoleRevoked, error) {
	event := new(BridgeRoleRevoked)
	if err := _Bridge.contract.UnpackLog(event, "RoleRevoked", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeRouteConfiguredIterator is returned from FilterRouteConfigured and is used to iterate over the raw logs and unpacked data for RouteConfigured events raised by the Bridge contract.
type BridgeRouteConfiguredIterator struct {
	Event *BridgeRouteConfigured // Event containing the contract specifics and raw log

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
func (it *BridgeRouteConfiguredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeRouteConfigured)
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
		it.Event = new(BridgeRouteConfigured)
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
func (it *BridgeRouteConfiguredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeRouteConfiguredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeRouteConfigured represents a RouteConfigured event raised by the Bridge contract.
type BridgeRouteConfigured struct {
	DestChainId uint64
	Bridge      common.Address
	HasLockbox  bool
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterRouteConfigured is a free log retrieval operation binding the contract event 0xab19e99f8223191275fefd1410893ea2b3001028e27ab75b975987c1c8c43207.
//
// Solidity: event RouteConfigured(uint64 indexed destChainId, address indexed bridge, bool indexed hasLockbox)
func (_Bridge *BridgeFilterer) FilterRouteConfigured(opts *bind.FilterOpts, destChainId []uint64, bridge []common.Address, hasLockbox []bool) (*BridgeRouteConfiguredIterator, error) {

	var destChainIdRule []interface{}
	for _, destChainIdItem := range destChainId {
		destChainIdRule = append(destChainIdRule, destChainIdItem)
	}
	var bridgeRule []interface{}
	for _, bridgeItem := range bridge {
		bridgeRule = append(bridgeRule, bridgeItem)
	}
	var hasLockboxRule []interface{}
	for _, hasLockboxItem := range hasLockbox {
		hasLockboxRule = append(hasLockboxRule, hasLockboxItem)
	}

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "RouteConfigured", destChainIdRule, bridgeRule, hasLockboxRule)
	if err != nil {
		return nil, err
	}
	return &BridgeRouteConfiguredIterator{contract: _Bridge.contract, event: "RouteConfigured", logs: logs, sub: sub}, nil
}

// WatchRouteConfigured is a free log subscription operation binding the contract event 0xab19e99f8223191275fefd1410893ea2b3001028e27ab75b975987c1c8c43207.
//
// Solidity: event RouteConfigured(uint64 indexed destChainId, address indexed bridge, bool indexed hasLockbox)
func (_Bridge *BridgeFilterer) WatchRouteConfigured(opts *bind.WatchOpts, sink chan<- *BridgeRouteConfigured, destChainId []uint64, bridge []common.Address, hasLockbox []bool) (event.Subscription, error) {

	var destChainIdRule []interface{}
	for _, destChainIdItem := range destChainId {
		destChainIdRule = append(destChainIdRule, destChainIdItem)
	}
	var bridgeRule []interface{}
	for _, bridgeItem := range bridge {
		bridgeRule = append(bridgeRule, bridgeItem)
	}
	var hasLockboxRule []interface{}
	for _, hasLockboxItem := range hasLockbox {
		hasLockboxRule = append(hasLockboxRule, hasLockboxItem)
	}

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "RouteConfigured", destChainIdRule, bridgeRule, hasLockboxRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeRouteConfigured)
				if err := _Bridge.contract.UnpackLog(event, "RouteConfigured", log); err != nil {
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

// ParseRouteConfigured is a log parse operation binding the contract event 0xab19e99f8223191275fefd1410893ea2b3001028e27ab75b975987c1c8c43207.
//
// Solidity: event RouteConfigured(uint64 indexed destChainId, address indexed bridge, bool indexed hasLockbox)
func (_Bridge *BridgeFilterer) ParseRouteConfigured(log types.Log) (*BridgeRouteConfigured, error) {
	event := new(BridgeRouteConfigured)
	if err := _Bridge.contract.UnpackLog(event, "RouteConfigured", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeTokenReceivedIterator is returned from FilterTokenReceived and is used to iterate over the raw logs and unpacked data for TokenReceived events raised by the Bridge contract.
type BridgeTokenReceivedIterator struct {
	Event *BridgeTokenReceived // Event containing the contract specifics and raw log

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
func (it *BridgeTokenReceivedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeTokenReceived)
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
		it.Event = new(BridgeTokenReceived)
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
func (it *BridgeTokenReceivedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeTokenReceivedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeTokenReceived represents a TokenReceived event raised by the Bridge contract.
type BridgeTokenReceived struct {
	SrcChainId uint64
	To         common.Address
	Value      *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterTokenReceived is a free log retrieval operation binding the contract event 0xfd6bb19aab6d8ed3fbc8848d3b1a76c95135e46d05bec2b84424219f9681929f.
//
// Solidity: event TokenReceived(uint64 indexed srcChainId, address indexed to, uint256 value)
func (_Bridge *BridgeFilterer) FilterTokenReceived(opts *bind.FilterOpts, srcChainId []uint64, to []common.Address) (*BridgeTokenReceivedIterator, error) {

	var srcChainIdRule []interface{}
	for _, srcChainIdItem := range srcChainId {
		srcChainIdRule = append(srcChainIdRule, srcChainIdItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "TokenReceived", srcChainIdRule, toRule)
	if err != nil {
		return nil, err
	}
	return &BridgeTokenReceivedIterator{contract: _Bridge.contract, event: "TokenReceived", logs: logs, sub: sub}, nil
}

// WatchTokenReceived is a free log subscription operation binding the contract event 0xfd6bb19aab6d8ed3fbc8848d3b1a76c95135e46d05bec2b84424219f9681929f.
//
// Solidity: event TokenReceived(uint64 indexed srcChainId, address indexed to, uint256 value)
func (_Bridge *BridgeFilterer) WatchTokenReceived(opts *bind.WatchOpts, sink chan<- *BridgeTokenReceived, srcChainId []uint64, to []common.Address) (event.Subscription, error) {

	var srcChainIdRule []interface{}
	for _, srcChainIdItem := range srcChainId {
		srcChainIdRule = append(srcChainIdRule, srcChainIdItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "TokenReceived", srcChainIdRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeTokenReceived)
				if err := _Bridge.contract.UnpackLog(event, "TokenReceived", log); err != nil {
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

// ParseTokenReceived is a log parse operation binding the contract event 0xfd6bb19aab6d8ed3fbc8848d3b1a76c95135e46d05bec2b84424219f9681929f.
//
// Solidity: event TokenReceived(uint64 indexed srcChainId, address indexed to, uint256 value)
func (_Bridge *BridgeFilterer) ParseTokenReceived(log types.Log) (*BridgeTokenReceived, error) {
	event := new(BridgeTokenReceived)
	if err := _Bridge.contract.UnpackLog(event, "TokenReceived", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeTokenSentIterator is returned from FilterTokenSent and is used to iterate over the raw logs and unpacked data for TokenSent events raised by the Bridge contract.
type BridgeTokenSentIterator struct {
	Event *BridgeTokenSent // Event containing the contract specifics and raw log

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
func (it *BridgeTokenSentIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeTokenSent)
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
		it.Event = new(BridgeTokenSent)
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
func (it *BridgeTokenSentIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeTokenSentIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeTokenSent represents a TokenSent event raised by the Bridge contract.
type BridgeTokenSent struct {
	DestChainId uint64
	From        common.Address
	To          common.Address
	Value       *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterTokenSent is a free log retrieval operation binding the contract event 0xc0464e720761b6de8643eed8d1cbf17ec66c3eb60179efaed5a58cf75580a4dc.
//
// Solidity: event TokenSent(uint64 indexed destChainId, address indexed from, address indexed to, uint256 value)
func (_Bridge *BridgeFilterer) FilterTokenSent(opts *bind.FilterOpts, destChainId []uint64, from []common.Address, to []common.Address) (*BridgeTokenSentIterator, error) {

	var destChainIdRule []interface{}
	for _, destChainIdItem := range destChainId {
		destChainIdRule = append(destChainIdRule, destChainIdItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "TokenSent", destChainIdRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &BridgeTokenSentIterator{contract: _Bridge.contract, event: "TokenSent", logs: logs, sub: sub}, nil
}

// WatchTokenSent is a free log subscription operation binding the contract event 0xc0464e720761b6de8643eed8d1cbf17ec66c3eb60179efaed5a58cf75580a4dc.
//
// Solidity: event TokenSent(uint64 indexed destChainId, address indexed from, address indexed to, uint256 value)
func (_Bridge *BridgeFilterer) WatchTokenSent(opts *bind.WatchOpts, sink chan<- *BridgeTokenSent, destChainId []uint64, from []common.Address, to []common.Address) (event.Subscription, error) {

	var destChainIdRule []interface{}
	for _, destChainIdItem := range destChainId {
		destChainIdRule = append(destChainIdRule, destChainIdItem)
	}
	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "TokenSent", destChainIdRule, fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeTokenSent)
				if err := _Bridge.contract.UnpackLog(event, "TokenSent", log); err != nil {
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

// ParseTokenSent is a log parse operation binding the contract event 0xc0464e720761b6de8643eed8d1cbf17ec66c3eb60179efaed5a58cf75580a4dc.
//
// Solidity: event TokenSent(uint64 indexed destChainId, address indexed from, address indexed to, uint256 value)
func (_Bridge *BridgeFilterer) ParseTokenSent(log types.Log) (*BridgeTokenSent, error) {
	event := new(BridgeTokenSent)
	if err := _Bridge.contract.UnpackLog(event, "TokenSent", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BridgeUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the Bridge contract.
type BridgeUnpausedIterator struct {
	Event *BridgeUnpaused // Event containing the contract specifics and raw log

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
func (it *BridgeUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeUnpaused)
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
		it.Event = new(BridgeUnpaused)
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
func (it *BridgeUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeUnpaused represents a Unpaused event raised by the Bridge contract.
type BridgeUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Bridge *BridgeFilterer) FilterUnpaused(opts *bind.FilterOpts) (*BridgeUnpausedIterator, error) {

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &BridgeUnpausedIterator{contract: _Bridge.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_Bridge *BridgeFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *BridgeUnpaused) (event.Subscription, error) {

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeUnpaused)
				if err := _Bridge.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_Bridge *BridgeFilterer) ParseUnpaused(log types.Log) (*BridgeUnpaused, error) {
	event := new(BridgeUnpaused)
	if err := _Bridge.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
