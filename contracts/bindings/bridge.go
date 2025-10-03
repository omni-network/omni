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
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"receiveDefaultGasLimit\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"receiveLockboxGasLimit\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"AUTHORIZER_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"CONFIGURER_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"DEFAULT_ADMIN_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"PAUSER_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"UNPAUSER_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"authorizeRoutes\",\"inputs\":[{\"name\":\"chainIds\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"},{\"name\":\"expectedRoutes\",\"type\":\"tuple[]\",\"internalType\":\"structIBridge.Route[]\",\"components\":[{\"name\":\"bridge\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"hasLockbox\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"bridgeFee\",\"inputs\":[{\"name\":\"destChainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"fee\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"claimFailedReceive\",\"inputs\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"claimable\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"configureRoutes\",\"inputs\":[{\"name\":\"chainIds\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"},{\"name\":\"routes\",\"type\":\"tuple[]\",\"internalType\":\"structIBridge.Route[]\",\"components\":[{\"name\":\"bridge\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"hasLockbox\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"defaultConfLevel\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRoleAdmin\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRoute\",\"inputs\":[{\"name\":\"destChainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"bridge\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"hasLockbox\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"grantRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"hasRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"admin_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"configurer_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"authorizer_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"pauser_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"unpauser_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"omni_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"token_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"lockbox_\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"lockbox\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"omni\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIOmniPortal\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"receiveToken\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"callerConfirmation\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"revokeRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"sendToken\",\"inputs\":[{\"name\":\"destChainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"wrap\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"refundTo\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"token\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"DefaultConfLevelSet\",\"inputs\":[{\"name\":\"conf\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"LockboxWithdrawalFailed\",\"inputs\":[{\"name\":\"badLockbox\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OmniPortalSet\",\"inputs\":[{\"name\":\"omni\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RetrySuccessful\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleAdminChanged\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"previousAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"newAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleGranted\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleRevoked\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RouteAuthorized\",\"inputs\":[{\"name\":\"destChainId\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"bridge\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"hasLockbox\",\"type\":\"bool\",\"indexed\":true,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RouteConfigured\",\"inputs\":[{\"name\":\"destChainId\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"bridge\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"hasLockbox\",\"type\":\"bool\",\"indexed\":true,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokenMintFailed\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokenReceived\",\"inputs\":[{\"name\":\"srcChainId\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"success\",\"type\":\"bool\",\"indexed\":true,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokenSent\",\"inputs\":[{\"name\":\"destChainId\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokenTransferFailed\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AccessControlBadConfirmation\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AccessControlUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"neededRole\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"ArrayLengthMismatch\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CannotWrap\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"EnforcedPause\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ExpectedPause\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientPayment\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidRoute\",\"inputs\":[{\"name\":\"chainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"NoClaimable\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"Unauthorized\",\"inputs\":[{\"name\":\"chainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ZeroAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroAmount\",\"inputs\":[]}]",
	Bin: "0x60c060405234801561000f575f80fd5b506040516128a93803806128a983398101604081905261002e9161011d565b6001600160401b03808316608052811660a052610049610050565b505061014e565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000900460ff16156100a05760405163f92ee8a960e01b815260040160405180910390fd5b80546001600160401b03908116146100ff5780546001600160401b0319166001600160401b0390811782556040519081527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b50565b80516001600160401b0381168114610118575f80fd5b919050565b5f806040838503121561012e575f80fd5b61013783610102565b915061014560208401610102565b90509250929050565b60805160a05161272c61017d5f395f818161103e0152611b5801525f81816110180152611b32015261272c5ff3fe608060405260043610610195575f3560e01c806374eeb847116100e7578063a217fddf11610087578063d547741f11610062578063d547741f146104d9578063e63ab1e9146104f8578063fb1bb9de1461052b578063fc0c546a1461055e575f80fd5b8063a217fddf14610460578063abbb9f4c14610473578063bc06946f146104a6575f80fd5b80638ad201f6116100c25780638ad201f6146103e457806391d148541461040357806397235a1e14610422578063a1cae82a14610441575f80fd5b806374eeb847146103805780638456cb59146103b15780638a29e2de146103c5575f80fd5b806339acf9f1116101525780634acd8d821161012d5780634acd8d82146102ed5780635c975abb1461032b5780635e23aa5f1461034e57806366cc570214610361575f80fd5b806339acf9f1146102785780633f4ba83a146102ae578063402914f5146102c2575f80fd5b806301ffc9a71461019957806322af16ec146101cd578063248a9ca3146101ee5780632ea7dfb71461021b5780632f2ff15d1461023a57806336568abe14610259575b5f80fd5b3480156101a4575f80fd5b506101b86101b3366004612199565b61057d565b60405190151581526020015b60405180910390f35b3480156101d8575f80fd5b506101ec6101e73660046121d4565b6105b3565b005b3480156101f9575f80fd5b5061020d6102083660046121ef565b61077c565b6040519081526020016101c4565b348015610226575f80fd5b506101ec61023536600461224d565b61079c565b348015610245575f80fd5b506101ec6102543660046122e5565b610a5a565b348015610264575f80fd5b506101ec6102733660046122e5565b610a7c565b348015610283575f80fd5b505f54610296906001600160a01b031681565b6040516001600160a01b0390911681526020016101c4565b3480156102b9575f80fd5b506101ec610ab4565b3480156102cd575f80fd5b5061020d6102dc3660046121d4565b60366020525f908152604090205481565b3480156102f8575f80fd5b5061030c610307366004612327565b610ae9565b604080516001600160a01b0390931683529015156020830152016101c4565b348015610336575f80fd5b505f805160206126d78339815191525460ff166101b8565b6101ec61035c36600461234f565b610b69565b34801561036c575f80fd5b50603354610296906001600160a01b031681565b34801561038b575f80fd5b505f5461039f90600160a01b900460ff1681565b60405160ff90911681526020016101c4565b3480156103bc575f80fd5b506101ec610b92565b3480156103d0575f80fd5b506101ec6103df3660046123b3565b610bc4565b3480156103ef575f80fd5b5061020d6103fe366004612327565b610f5f565b34801561040e575f80fd5b506101b861041d3660046122e5565b611069565b34801561042d575f80fd5b506101ec61043c366004612457565b61109f565b34801561044c575f80fd5b506101ec61045b36600461224d565b61120a565b34801561046b575f80fd5b5061020d5f81565b34801561047e575f80fd5b5061020d7f527e2c92bb6983874717bce74818faf5a9be45b6e85909ee478af653c6d9875581565b3480156104b1575f80fd5b5061020d7f94858e5561d6a33fcce848f16acfe1514fe5166e32b456aff42d7fb50e8c52ad81565b3480156104e4575f80fd5b506101ec6104f33660046122e5565b611386565b348015610503575f80fd5b5061020d7f539440820030c4994db4e31b6b800deafd503688728f932addfe7a410515c14c81565b348015610536575f80fd5b5061020d7f82b32d9ab5100db08aeb9a0e08b422d14851ec118736590462bf9c085a6e944881565b348015610569575f80fd5b50603254610296906001600160a01b031681565b5f6001600160e01b03198216637965db0b60e01b14806105ad57506301ffc9a760e01b6001600160e01b03198316145b92915050565b6105bb6113a2565b6001600160a01b0381165f90815260366020526040812054908190036105f457604051631129777360e21b815260040160405180910390fd5b6001600160a01b038083165f9081526036602052604081205560335416610679576032546040516340c10f1960e01b81526001600160a01b03909116906340c10f19906106479085908590600401612481565b5f604051808303815f87803b15801561065e575f80fd5b505af1158015610670573d5f803e3d5ffd5b50505050610738565b6032546040516340c10f1960e01b81526001600160a01b03909116906340c10f19906106ab9030908590600401612481565b5f604051808303815f87803b1580156106c2575f80fd5b505af11580156106d4573d5f803e3d5ffd5b505060335460405163040b850f60e31b81526001600160a01b03909116925063205c2878915061070a9085908590600401612481565b5f604051808303815f87803b158015610721575f80fd5b505af1158015610733573d5f803e3d5ffd5b505050505b6040518181526001600160a01b0383169033907fafc6034e5bc12b75c8fd712cc3306dba0afd7d2c5156fe40015ff2b3551f86c09060200160405180910390a35050565b5f9081525f805160206126b7833981519152602052604090206001015490565b7f94858e5561d6a33fcce848f16acfe1514fe5166e32b456aff42d7fb50e8c52ad6107c6816113d4565b5f5b84811015610a52575f60355f8888858181106107e6576107e661249a565b90506020020160208101906107fb9190612327565b6001600160401b0316815260208082019290925260409081015f208151808301909252546001600160a01b0381168252600160a01b900460ff1615159181019190915290508484838181106108525761085261249a565b61086892602060409092020190810191506121d4565b6001600160a01b0316815f01516001600160a01b03161415806108bc57508484838181106108985761089861249a565b90506040020160200160208101906108b091906124ae565b15158160200151151514155b15610912578686838181106108d3576108d361249a565b90506020020160208101906108e89190612327565b60405163f6e829e160e01b81526001600160401b0390911660048201526024015b60405180910390fd5b8060345f8989868181106109285761092861249a565b905060200201602081019061093d9190612327565b6001600160401b031681526020808201929092526040015f9081208351815494909301511515600160a01b026001600160a81b03199094166001600160a01b039093169290921792909217905560359088888581811061099f5761099f61249a565b90506020020160208101906109b49190612327565b6001600160401b031681526020808201929092526040015f2080546001600160a81b03191690558101518151901515906001600160a01b03168888858181106109ff576109ff61249a565b9050602002016020810190610a149190612327565b6001600160401b03167f520fd445d84479cc5d10c3b3a468e84b5f8d6069143aa225be61e6dae8d5e38c60405160405180910390a4506001016107c8565b505050505050565b610a638261077c565b610a6c816113d4565b610a7683836113de565b50505050565b6001600160a01b0381163314610aa55760405163334bd91960e11b815260040160405180910390fd5b610aaf828261147f565b505050565b7f82b32d9ab5100db08aeb9a0e08b422d14851ec118736590462bf9c085a6e9448610ade816113d4565b610ae66114f8565b50565b6001600160401b0381165f9081526034602090815260408083208151808301909252546001600160a01b038116808352600160a01b90910460ff161515928201929092528291610b575760405163f6e829e160e01b81526001600160401b0385166004820152602401610909565b80516020909101519094909350915050565b610b716113a2565b610b7e8585858585611558565b610b8b8585858585611643565b5050505050565b7f539440820030c4994db4e31b6b800deafd503688728f932addfe7a410515c14c610bbc816113d4565b610ae66117a2565b5f610bcd6117ea565b805490915060ff600160401b82041615906001600160401b03165f81158015610bf35750825b90505f826001600160401b03166001148015610c0e5750303b155b905081158015610c1c575080155b15610c3a5760405163f92ee8a960e01b815260040160405180910390fd5b845467ffffffffffffffff191660011785558315610c6457845460ff60401b1916600160401b1785555b6001600160a01b038d16610c8b5760405163d92e233d60e01b815260040160405180910390fd5b6001600160a01b038c16610cb25760405163d92e233d60e01b815260040160405180910390fd5b6001600160a01b038b16610cd95760405163d92e233d60e01b815260040160405180910390fd5b6001600160a01b038a16610d005760405163d92e233d60e01b815260040160405180910390fd5b6001600160a01b038916610d275760405163d92e233d60e01b815260040160405180910390fd5b6001600160a01b038816610d4e5760405163d92e233d60e01b815260040160405180910390fd5b6001600160a01b038716610d755760405163d92e233d60e01b815260040160405180910390fd5b610d7d611812565b610d85611812565b610d9088600461181a565b610d9a5f8e6113de565b50610dc57f527e2c92bb6983874717bce74818faf5a9be45b6e85909ee478af653c6d987558d6113de565b50610df07f94858e5561d6a33fcce848f16acfe1514fe5166e32b456aff42d7fb50e8c52ad8c6113de565b50610e1b7f539440820030c4994db4e31b6b800deafd503688728f932addfe7a410515c14c8b6113de565b50610e467f82b32d9ab5100db08aeb9a0e08b422d14851ec118736590462bf9c085a6e94488a6113de565b50603280546001600160a01b0319166001600160a01b0389811691909117909155861615610f0a578560335f6101000a8154816001600160a01b0302191690836001600160a01b03160217905550610f0a865f19886001600160a01b031663fc0c546a6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610ed6573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610efa91906124c9565b6001600160a01b03169190611838565b8315610f5057845460ff60401b19168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b50505050505050505050505050565b6001600160401b0381165f9081526034602090815260408083208151808301909252546001600160a01b038116808352600160a01b90910460ff1615159282019290925290610fcc5760405163f6e829e160e01b81526001600160401b0384166004820152602401610909565b611062835f1980604051602401610fe4929190612481565b60408051601f19818403018152919052602080820180516001600160e01b0316634b91ad0f60e11b17905284015161103c577f00000000000000000000000000000000000000000000000000000000000000006118c2565b7f00000000000000000000000000000000000000000000000000000000000000006118c2565b9392505050565b5f9182525f805160206126b7833981519152602090815260408084206001600160a01b0393909316845291905290205460ff1690565b5f5460408051631799380760e11b815281516001600160a01b0390931692632f32700e926004808401939192918290030181865afa1580156110e3573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061110791906124e4565b8051600180546020909301516001600160401b039092166001600160e01b031990931692909217600160401b6001600160a01b0392831602179091555f541633036111c2576001546001600160401b0381165f908152603460205260409020546001600160a01b03908116600160401b90920416146111bd57600154604051633dfc334560e01b81526001600160401b0382166004820152600160401b9091046001600160a01b03166024820152604401610909565b6111ec565b604051633dfc334560e01b81526001600160401b0346166004820152336024820152604401610909565b6111f6828261193d565b5050600180546001600160e01b0319169055565b7f527e2c92bb6983874717bce74818faf5a9be45b6e85909ee478af653c6d98755611234816113d4565b8382146112545760405163512509d360e11b815260040160405180910390fd5b5f5b84811015610a52578383828181106112705761127061249a565b90506040020160355f88888581811061128b5761128b61249a565b90506020020160208101906112a09190612327565b6001600160401b0316815260208101919091526040015f206112c2828261254e565b9050508383828181106112d7576112d761249a565b90506040020160200160208101906112ef91906124ae565b15158484838181106113035761130361249a565b61131992602060409092020190810191506121d4565b6001600160a01b03168787848181106113345761133461249a565b90506020020160208101906113499190612327565b6001600160401b03167fab19e99f8223191275fefd1410893ea2b3001028e27ab75b975987c1c8c4320760405160405180910390a4600101611256565b61138f8261077c565b611398816113d4565b610a76838361147f565b5f805160206126d78339815191525460ff16156113d25760405163d93c066560e01b815260040160405180910390fd5b565b610ae681336119e5565b5f5f805160206126b78339815191526113f78484611069565b611476575f848152602082815260408083206001600160a01b03871684529091529020805460ff1916600117905561142c3390565b6001600160a01b0316836001600160a01b0316857f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a460019150506105ad565b5f9150506105ad565b5f5f805160206126b78339815191526114988484611069565b15611476575f848152602082815260408083206001600160a01b0387168085529252808320805460ff1916905551339287917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a460019150506105ad565b611500611a10565b5f805160206126d7833981519152805460ff191681557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa335b6040516001600160a01b0390911681526020015b60405180910390a150565b6001600160401b0385165f908152603460205260409020546001600160a01b03166115a15760405163f6e829e160e01b81526001600160401b0386166004820152602401610909565b6001600160a01b0384166115c85760405163d92e233d60e01b815260040160405180910390fd5b825f036115e857604051631f2a200560e01b815260040160405180910390fd5b8180156115fe57506033546001600160a01b0316155b1561161c57604051637d25d4c960e11b815260040160405180910390fd5b6001600160a01b038116610b8b5760405163d92e233d60e01b815260040160405180910390fd5b6032546033546001600160a01b0391821691168315611732576116d3333087846001600160a01b031663fc0c546a6040518163ffffffff1660e01b8152600401602060405180830381865afa15801561169e573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906116c291906124c9565b6001600160a01b0316929190611a3f565b60405160016255295b60e01b031981526001600160a01b0382169063ffaad6a5906117049033908990600401612481565b5f604051808303815f87803b15801561171b575f80fd5b505af115801561172d573d5f803e3d5ffd5b505050505b6040516379ef98bf60e11b81526001600160a01b0383169063f3df317e906117609033908990600401612481565b5f604051808303815f87803b158015611777575f80fd5b505af1158015611789573d5f803e3d5ffd5b5050505061179987878786611a98565b50505050505050565b6117aa6113a2565b5f805160206126d7833981519152805460ff191660011781557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a25833611539565b5f807ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a006105ad565b6113d2611c24565b611822611c24565b61182b82611c49565b61183481611ce1565b5050565b816014528060345263095ea7b360601b5f5260205f604460105f875af18060015f5114166118b857803d853b1517106118b8575f60345263095ea7b360601b5f525f38604460105f885af1508160345260205f604460105f885af190508060015f5114166118b857803d853b1517106118b857633e3f8f735f526004601cfd5b505f603452505050565b5f8054604051632376548f60e21b81526001600160a01b0390911690638dd9523c906118f6908790879087906004016125d5565b602060405180830381865afa158015611911573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190611935919061260f565b949350505050565b6032546033546001600160a01b0391821691165f81611969576119628386865f611d83565b905061198d565b6119768386866001611d83565b9050801561198d5761198a83838787611e66565b90505b600154604051858152821515916001600160a01b038816916001600160401b03909116907f7fd41495160762948bbb964edcce550c26992b895a329f399e047786f20973669060200160405180910390a45050505050565b6119ef8282611069565b61183457808260405163e2517d3f60e01b8152600401610909929190612481565b5f805160206126d78339815191525460ff166113d257604051638dfc202b60e01b815260040160405180910390fd5b60405181606052826040528360601b602c526323b872dd60601b600c5260205f6064601c5f895af18060015f511416611a8a57803d873b151710611a8a57637939f4245f526004601cfd5b505f60605260405250505050565b6001600160401b0384165f9081526034602090815260408083208151808301835290546001600160a01b0381168252600160a01b900460ff1615159281019290925251909190611aee9086908690602401612481565b60408051601f19818403018152919052602080820180516001600160e01b0316634b91ad0f60e11b1790528351908401519192505f91611b7c9189918590611b56577f000000000000000000000000000000000000000000000000000000000000000061200e565b7f000000000000000000000000000000000000000000000000000000000000000061200e565b905080341015611b9f5760405163cd1c886760e01b815260040160405180910390fd5b80341115611bc457611bc4611bb48234612626565b6001600160a01b0386169061214c565b856001600160a01b0316336001600160a01b0316886001600160401b03167fc0464e720761b6de8643eed8d1cbf17ec66c3eb60179efaed5a58cf75580a4dc88604051611c1391815260200190565b60405180910390a450505050505050565b611c2c612165565b6113d257604051631afcd79f60e31b815260040160405180910390fd5b6001600160a01b038116611c945760405162461bcd60e51b8152602060048201526012602482015271584170703a206e6f207a65726f206f6d6e6960701b6044820152606401610909565b5f80546001600160a01b0319166001600160a01b0383169081179091556040519081527f79162c8d053a07e70cdc1ccc536f0440b571f8508377d2bef51094fadab98f479060200161154d565b611cea8161217e565b611d365760405162461bcd60e51b815260206004820152601860248201527f584170703a20696e76616c696420636f6e66206c6576656c00000000000000006044820152606401610909565b5f805460ff60a01b1916600160a01b60ff8416908102919091179091556040519081527f8de08a798b4e50b4f351c1eaa91a11530043802be3ffac2df87db0c45a2e84839060200161154d565b5f846001600160a01b03166340c10f1983611d9e5785611da0565b305b856040518363ffffffff1660e01b8152600401611dbe929190612481565b5f604051808303815f87803b158015611dd5575f80fd5b505af1925050508015611de6575060015b611e5b576001600160a01b0384165f908152603660205260409020805484019055836001600160a01b0316856001600160a01b03167f259005b50cf55d190d280dfa8480709f3e27c4e6ecd89e3d491dfd288f6ce38585604051611e4c91815260200190565b60405180910390a3505f611935565b506001949350505050565b60405163040b850f60e31b81525f906001600160a01b0385169063205c287890611e969086908690600401612481565b5f604051808303815f87803b158015611ead575f80fd5b505af1925050508015611ebe575060015b611e5b5760405163a9059cbb60e01b81526001600160a01b0386169063a9059cbb90611ef09086908690600401612481565b6020604051808303815f875af1925050508015611f2a575060408051601f3d908101601f19168201909252611f2791810190612645565b60015b611f90576001600160a01b0383165f908152603660205260409020805483019055826001600160a01b0316846001600160a01b03167fd5684d2e31a0d2443b1102269ffecfa6d05a6a36e312219f9108c0e95f8bc9af84604051611e4c91815260200190565b80612008576001600160a01b0384165f908152603660205260409020805484019055836001600160a01b0316866001600160a01b03167fc87767983e580cd51a7614924de0506ff919e5220d94509794cd03bb6564b0bf85604051611ff791815260200190565b60405180910390a35f915050611935565b50611e5b565b5f8054604051632376548f60e21b815282916001600160a01b031690638dd9523c90612042908990889088906004016125d5565b602060405180830381865afa15801561205d573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190612081919061260f565b9050804710156120d35760405162461bcd60e51b815260206004820152601860248201527f584170703a20696e73756666696369656e742066756e647300000000000000006044820152606401610909565b5f5460405163c21dda4f60e01b81526001600160a01b0382169163c21dda4f918491612114918b91600160a01b900460ff16908b908b908b90600401612660565b5f604051808303818588803b15801561212b575f80fd5b505af115801561213d573d5f803e3d5ffd5b50939998505050505050505050565b5f385f3884865af16118345763b12d13eb5f526004601cfd5b5f61216e6117ea565b54600160401b900460ff16919050565b5f60ff8216600114806105ad575060ff821660041492915050565b5f602082840312156121a9575f80fd5b81356001600160e01b031981168114611062575f80fd5b6001600160a01b0381168114610ae6575f80fd5b5f602082840312156121e4575f80fd5b8135611062816121c0565b5f602082840312156121ff575f80fd5b5035919050565b5f8083601f840112612216575f80fd5b5081356001600160401b0381111561222c575f80fd5b6020830191508360208260061b8501011115612246575f80fd5b9250929050565b5f805f8060408587031215612260575f80fd5b84356001600160401b03811115612275575f80fd5b8501601f81018713612285575f80fd5b80356001600160401b0381111561229a575f80fd5b8760208260051b84010111156122ae575f80fd5b6020918201955093508501356001600160401b038111156122cd575f80fd5b6122d987828801612206565b95989497509550505050565b5f80604083850312156122f6575f80fd5b823591506020830135612308816121c0565b809150509250929050565b6001600160401b0381168114610ae6575f80fd5b5f60208284031215612337575f80fd5b813561106281612313565b8015158114610ae6575f80fd5b5f805f805f60a08688031215612363575f80fd5b853561236e81612313565b9450602086013561237e816121c0565b935060408601359250606086013561239581612342565b915060808601356123a5816121c0565b809150509295509295909350565b5f805f805f805f80610100898b0312156123cb575f80fd5b88356123d6816121c0565b975060208901356123e6816121c0565b965060408901356123f6816121c0565b95506060890135612406816121c0565b94506080890135612416816121c0565b935060a0890135612426816121c0565b925060c0890135612436816121c0565b915060e0890135612446816121c0565b809150509295985092959890939650565b5f8060408385031215612468575f80fd5b8235612473816121c0565b946020939093013593505050565b6001600160a01b03929092168252602082015260400190565b634e487b7160e01b5f52603260045260245ffd5b5f602082840312156124be575f80fd5b813561106281612342565b5f602082840312156124d9575f80fd5b8151611062816121c0565b5f60408284031280156124f5575f80fd5b50604080519081016001600160401b038111828210171561252457634e487b7160e01b5f52604160045260245ffd5b604052825161253281612313565b81526020830151612542816121c0565b60208201529392505050565b8135612559816121c0565b81546001600160a01b031981166001600160a01b03929092169182178355602084013561258581612342565b6001600160a81b03199190911690911790151560a01b60ff60a01b1617905550565b5f81518084528060208401602086015e5f602082860101526020601f19601f83011685010191505092915050565b6001600160401b0384168152606060208201525f6125f660608301856125a7565b90506001600160401b0383166040830152949350505050565b5f6020828403121561261f575f80fd5b5051919050565b818103818111156105ad57634e487b7160e01b5f52601160045260245ffd5b5f60208284031215612655575f80fd5b815161106281612342565b6001600160401b038616815260ff851660208201526001600160a01b038416604082015260a0606082018190525f9061269b908301856125a7565b90506001600160401b0383166080830152969550505050505056fe02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800cd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f03300a26469706673582212208e1ee4bb3cc8d1c51acd56f8234b25a22181d564989661742f750f592102c52b64736f6c634300081a0033",
}

// BridgeABI is the input ABI used to generate the binding from.
// Deprecated: Use BridgeMetaData.ABI instead.
var BridgeABI = BridgeMetaData.ABI

// BridgeBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use BridgeMetaData.Bin instead.
var BridgeBin = BridgeMetaData.Bin

// DeployBridge deploys a new Ethereum contract, binding an instance of Bridge to it.
func DeployBridge(auth *bind.TransactOpts, backend bind.ContractBackend, receiveDefaultGasLimit uint64, receiveLockboxGasLimit uint64) (common.Address, *types.Transaction, *Bridge, error) {
	parsed, err := BridgeMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BridgeBin), backend, receiveDefaultGasLimit, receiveLockboxGasLimit)
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

// AUTHORIZERROLE is a free data retrieval call binding the contract method 0xbc06946f.
//
// Solidity: function AUTHORIZER_ROLE() view returns(bytes32)
func (_Bridge *BridgeCaller) AUTHORIZERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "AUTHORIZER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// AUTHORIZERROLE is a free data retrieval call binding the contract method 0xbc06946f.
//
// Solidity: function AUTHORIZER_ROLE() view returns(bytes32)
func (_Bridge *BridgeSession) AUTHORIZERROLE() ([32]byte, error) {
	return _Bridge.Contract.AUTHORIZERROLE(&_Bridge.CallOpts)
}

// AUTHORIZERROLE is a free data retrieval call binding the contract method 0xbc06946f.
//
// Solidity: function AUTHORIZER_ROLE() view returns(bytes32)
func (_Bridge *BridgeCallerSession) AUTHORIZERROLE() ([32]byte, error) {
	return _Bridge.Contract.AUTHORIZERROLE(&_Bridge.CallOpts)
}

// CONFIGURERROLE is a free data retrieval call binding the contract method 0xabbb9f4c.
//
// Solidity: function CONFIGURER_ROLE() view returns(bytes32)
func (_Bridge *BridgeCaller) CONFIGURERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "CONFIGURER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// CONFIGURERROLE is a free data retrieval call binding the contract method 0xabbb9f4c.
//
// Solidity: function CONFIGURER_ROLE() view returns(bytes32)
func (_Bridge *BridgeSession) CONFIGURERROLE() ([32]byte, error) {
	return _Bridge.Contract.CONFIGURERROLE(&_Bridge.CallOpts)
}

// CONFIGURERROLE is a free data retrieval call binding the contract method 0xabbb9f4c.
//
// Solidity: function CONFIGURER_ROLE() view returns(bytes32)
func (_Bridge *BridgeCallerSession) CONFIGURERROLE() ([32]byte, error) {
	return _Bridge.Contract.CONFIGURERROLE(&_Bridge.CallOpts)
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

// UNPAUSERROLE is a free data retrieval call binding the contract method 0xfb1bb9de.
//
// Solidity: function UNPAUSER_ROLE() view returns(bytes32)
func (_Bridge *BridgeCaller) UNPAUSERROLE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "UNPAUSER_ROLE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// UNPAUSERROLE is a free data retrieval call binding the contract method 0xfb1bb9de.
//
// Solidity: function UNPAUSER_ROLE() view returns(bytes32)
func (_Bridge *BridgeSession) UNPAUSERROLE() ([32]byte, error) {
	return _Bridge.Contract.UNPAUSERROLE(&_Bridge.CallOpts)
}

// UNPAUSERROLE is a free data retrieval call binding the contract method 0xfb1bb9de.
//
// Solidity: function UNPAUSER_ROLE() view returns(bytes32)
func (_Bridge *BridgeCallerSession) UNPAUSERROLE() ([32]byte, error) {
	return _Bridge.Contract.UNPAUSERROLE(&_Bridge.CallOpts)
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

// Claimable is a free data retrieval call binding the contract method 0x402914f5.
//
// Solidity: function claimable(address ) view returns(uint256)
func (_Bridge *BridgeCaller) Claimable(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Bridge.contract.Call(opts, &out, "claimable", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Claimable is a free data retrieval call binding the contract method 0x402914f5.
//
// Solidity: function claimable(address ) view returns(uint256)
func (_Bridge *BridgeSession) Claimable(arg0 common.Address) (*big.Int, error) {
	return _Bridge.Contract.Claimable(&_Bridge.CallOpts, arg0)
}

// Claimable is a free data retrieval call binding the contract method 0x402914f5.
//
// Solidity: function claimable(address ) view returns(uint256)
func (_Bridge *BridgeCallerSession) Claimable(arg0 common.Address) (*big.Int, error) {
	return _Bridge.Contract.Claimable(&_Bridge.CallOpts, arg0)
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

// AuthorizeRoutes is a paid mutator transaction binding the contract method 0x2ea7dfb7.
//
// Solidity: function authorizeRoutes(uint64[] chainIds, (address,bool)[] expectedRoutes) returns()
func (_Bridge *BridgeTransactor) AuthorizeRoutes(opts *bind.TransactOpts, chainIds []uint64, expectedRoutes []IBridgeRoute) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "authorizeRoutes", chainIds, expectedRoutes)
}

// AuthorizeRoutes is a paid mutator transaction binding the contract method 0x2ea7dfb7.
//
// Solidity: function authorizeRoutes(uint64[] chainIds, (address,bool)[] expectedRoutes) returns()
func (_Bridge *BridgeSession) AuthorizeRoutes(chainIds []uint64, expectedRoutes []IBridgeRoute) (*types.Transaction, error) {
	return _Bridge.Contract.AuthorizeRoutes(&_Bridge.TransactOpts, chainIds, expectedRoutes)
}

// AuthorizeRoutes is a paid mutator transaction binding the contract method 0x2ea7dfb7.
//
// Solidity: function authorizeRoutes(uint64[] chainIds, (address,bool)[] expectedRoutes) returns()
func (_Bridge *BridgeTransactorSession) AuthorizeRoutes(chainIds []uint64, expectedRoutes []IBridgeRoute) (*types.Transaction, error) {
	return _Bridge.Contract.AuthorizeRoutes(&_Bridge.TransactOpts, chainIds, expectedRoutes)
}

// ClaimFailedReceive is a paid mutator transaction binding the contract method 0x22af16ec.
//
// Solidity: function claimFailedReceive(address addr) returns()
func (_Bridge *BridgeTransactor) ClaimFailedReceive(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "claimFailedReceive", addr)
}

// ClaimFailedReceive is a paid mutator transaction binding the contract method 0x22af16ec.
//
// Solidity: function claimFailedReceive(address addr) returns()
func (_Bridge *BridgeSession) ClaimFailedReceive(addr common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.ClaimFailedReceive(&_Bridge.TransactOpts, addr)
}

// ClaimFailedReceive is a paid mutator transaction binding the contract method 0x22af16ec.
//
// Solidity: function claimFailedReceive(address addr) returns()
func (_Bridge *BridgeTransactorSession) ClaimFailedReceive(addr common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.ClaimFailedReceive(&_Bridge.TransactOpts, addr)
}

// ConfigureRoutes is a paid mutator transaction binding the contract method 0xa1cae82a.
//
// Solidity: function configureRoutes(uint64[] chainIds, (address,bool)[] routes) returns()
func (_Bridge *BridgeTransactor) ConfigureRoutes(opts *bind.TransactOpts, chainIds []uint64, routes []IBridgeRoute) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "configureRoutes", chainIds, routes)
}

// ConfigureRoutes is a paid mutator transaction binding the contract method 0xa1cae82a.
//
// Solidity: function configureRoutes(uint64[] chainIds, (address,bool)[] routes) returns()
func (_Bridge *BridgeSession) ConfigureRoutes(chainIds []uint64, routes []IBridgeRoute) (*types.Transaction, error) {
	return _Bridge.Contract.ConfigureRoutes(&_Bridge.TransactOpts, chainIds, routes)
}

// ConfigureRoutes is a paid mutator transaction binding the contract method 0xa1cae82a.
//
// Solidity: function configureRoutes(uint64[] chainIds, (address,bool)[] routes) returns()
func (_Bridge *BridgeTransactorSession) ConfigureRoutes(chainIds []uint64, routes []IBridgeRoute) (*types.Transaction, error) {
	return _Bridge.Contract.ConfigureRoutes(&_Bridge.TransactOpts, chainIds, routes)
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

// Initialize is a paid mutator transaction binding the contract method 0x8a29e2de.
//
// Solidity: function initialize(address admin_, address configurer_, address authorizer_, address pauser_, address unpauser_, address omni_, address token_, address lockbox_) returns()
func (_Bridge *BridgeTransactor) Initialize(opts *bind.TransactOpts, admin_ common.Address, configurer_ common.Address, authorizer_ common.Address, pauser_ common.Address, unpauser_ common.Address, omni_ common.Address, token_ common.Address, lockbox_ common.Address) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "initialize", admin_, configurer_, authorizer_, pauser_, unpauser_, omni_, token_, lockbox_)
}

// Initialize is a paid mutator transaction binding the contract method 0x8a29e2de.
//
// Solidity: function initialize(address admin_, address configurer_, address authorizer_, address pauser_, address unpauser_, address omni_, address token_, address lockbox_) returns()
func (_Bridge *BridgeSession) Initialize(admin_ common.Address, configurer_ common.Address, authorizer_ common.Address, pauser_ common.Address, unpauser_ common.Address, omni_ common.Address, token_ common.Address, lockbox_ common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.Initialize(&_Bridge.TransactOpts, admin_, configurer_, authorizer_, pauser_, unpauser_, omni_, token_, lockbox_)
}

// Initialize is a paid mutator transaction binding the contract method 0x8a29e2de.
//
// Solidity: function initialize(address admin_, address configurer_, address authorizer_, address pauser_, address unpauser_, address omni_, address token_, address lockbox_) returns()
func (_Bridge *BridgeTransactorSession) Initialize(admin_ common.Address, configurer_ common.Address, authorizer_ common.Address, pauser_ common.Address, unpauser_ common.Address, omni_ common.Address, token_ common.Address, lockbox_ common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.Initialize(&_Bridge.TransactOpts, admin_, configurer_, authorizer_, pauser_, unpauser_, omni_, token_, lockbox_)
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

// SendToken is a paid mutator transaction binding the contract method 0x5e23aa5f.
//
// Solidity: function sendToken(uint64 destChainId, address to, uint256 value, bool wrap, address refundTo) payable returns()
func (_Bridge *BridgeTransactor) SendToken(opts *bind.TransactOpts, destChainId uint64, to common.Address, value *big.Int, wrap bool, refundTo common.Address) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "sendToken", destChainId, to, value, wrap, refundTo)
}

// SendToken is a paid mutator transaction binding the contract method 0x5e23aa5f.
//
// Solidity: function sendToken(uint64 destChainId, address to, uint256 value, bool wrap, address refundTo) payable returns()
func (_Bridge *BridgeSession) SendToken(destChainId uint64, to common.Address, value *big.Int, wrap bool, refundTo common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.SendToken(&_Bridge.TransactOpts, destChainId, to, value, wrap, refundTo)
}

// SendToken is a paid mutator transaction binding the contract method 0x5e23aa5f.
//
// Solidity: function sendToken(uint64 destChainId, address to, uint256 value, bool wrap, address refundTo) payable returns()
func (_Bridge *BridgeTransactorSession) SendToken(destChainId uint64, to common.Address, value *big.Int, wrap bool, refundTo common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.SendToken(&_Bridge.TransactOpts, destChainId, to, value, wrap, refundTo)
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

// BridgeRetrySuccessfulIterator is returned from FilterRetrySuccessful and is used to iterate over the raw logs and unpacked data for RetrySuccessful events raised by the Bridge contract.
type BridgeRetrySuccessfulIterator struct {
	Event *BridgeRetrySuccessful // Event containing the contract specifics and raw log

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
func (it *BridgeRetrySuccessfulIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeRetrySuccessful)
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
		it.Event = new(BridgeRetrySuccessful)
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
func (it *BridgeRetrySuccessfulIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeRetrySuccessfulIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeRetrySuccessful represents a RetrySuccessful event raised by the Bridge contract.
type BridgeRetrySuccessful struct {
	Caller common.Address
	To     common.Address
	Value  *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterRetrySuccessful is a free log retrieval operation binding the contract event 0xafc6034e5bc12b75c8fd712cc3306dba0afd7d2c5156fe40015ff2b3551f86c0.
//
// Solidity: event RetrySuccessful(address indexed caller, address indexed to, uint256 value)
func (_Bridge *BridgeFilterer) FilterRetrySuccessful(opts *bind.FilterOpts, caller []common.Address, to []common.Address) (*BridgeRetrySuccessfulIterator, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "RetrySuccessful", callerRule, toRule)
	if err != nil {
		return nil, err
	}
	return &BridgeRetrySuccessfulIterator{contract: _Bridge.contract, event: "RetrySuccessful", logs: logs, sub: sub}, nil
}

// WatchRetrySuccessful is a free log subscription operation binding the contract event 0xafc6034e5bc12b75c8fd712cc3306dba0afd7d2c5156fe40015ff2b3551f86c0.
//
// Solidity: event RetrySuccessful(address indexed caller, address indexed to, uint256 value)
func (_Bridge *BridgeFilterer) WatchRetrySuccessful(opts *bind.WatchOpts, sink chan<- *BridgeRetrySuccessful, caller []common.Address, to []common.Address) (event.Subscription, error) {

	var callerRule []interface{}
	for _, callerItem := range caller {
		callerRule = append(callerRule, callerItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "RetrySuccessful", callerRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeRetrySuccessful)
				if err := _Bridge.contract.UnpackLog(event, "RetrySuccessful", log); err != nil {
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

// ParseRetrySuccessful is a log parse operation binding the contract event 0xafc6034e5bc12b75c8fd712cc3306dba0afd7d2c5156fe40015ff2b3551f86c0.
//
// Solidity: event RetrySuccessful(address indexed caller, address indexed to, uint256 value)
func (_Bridge *BridgeFilterer) ParseRetrySuccessful(log types.Log) (*BridgeRetrySuccessful, error) {
	event := new(BridgeRetrySuccessful)
	if err := _Bridge.contract.UnpackLog(event, "RetrySuccessful", log); err != nil {
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

// BridgeRouteAuthorizedIterator is returned from FilterRouteAuthorized and is used to iterate over the raw logs and unpacked data for RouteAuthorized events raised by the Bridge contract.
type BridgeRouteAuthorizedIterator struct {
	Event *BridgeRouteAuthorized // Event containing the contract specifics and raw log

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
func (it *BridgeRouteAuthorizedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeRouteAuthorized)
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
		it.Event = new(BridgeRouteAuthorized)
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
func (it *BridgeRouteAuthorizedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeRouteAuthorizedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeRouteAuthorized represents a RouteAuthorized event raised by the Bridge contract.
type BridgeRouteAuthorized struct {
	DestChainId uint64
	Bridge      common.Address
	HasLockbox  bool
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterRouteAuthorized is a free log retrieval operation binding the contract event 0x520fd445d84479cc5d10c3b3a468e84b5f8d6069143aa225be61e6dae8d5e38c.
//
// Solidity: event RouteAuthorized(uint64 indexed destChainId, address indexed bridge, bool indexed hasLockbox)
func (_Bridge *BridgeFilterer) FilterRouteAuthorized(opts *bind.FilterOpts, destChainId []uint64, bridge []common.Address, hasLockbox []bool) (*BridgeRouteAuthorizedIterator, error) {

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

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "RouteAuthorized", destChainIdRule, bridgeRule, hasLockboxRule)
	if err != nil {
		return nil, err
	}
	return &BridgeRouteAuthorizedIterator{contract: _Bridge.contract, event: "RouteAuthorized", logs: logs, sub: sub}, nil
}

// WatchRouteAuthorized is a free log subscription operation binding the contract event 0x520fd445d84479cc5d10c3b3a468e84b5f8d6069143aa225be61e6dae8d5e38c.
//
// Solidity: event RouteAuthorized(uint64 indexed destChainId, address indexed bridge, bool indexed hasLockbox)
func (_Bridge *BridgeFilterer) WatchRouteAuthorized(opts *bind.WatchOpts, sink chan<- *BridgeRouteAuthorized, destChainId []uint64, bridge []common.Address, hasLockbox []bool) (event.Subscription, error) {

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

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "RouteAuthorized", destChainIdRule, bridgeRule, hasLockboxRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeRouteAuthorized)
				if err := _Bridge.contract.UnpackLog(event, "RouteAuthorized", log); err != nil {
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

// ParseRouteAuthorized is a log parse operation binding the contract event 0x520fd445d84479cc5d10c3b3a468e84b5f8d6069143aa225be61e6dae8d5e38c.
//
// Solidity: event RouteAuthorized(uint64 indexed destChainId, address indexed bridge, bool indexed hasLockbox)
func (_Bridge *BridgeFilterer) ParseRouteAuthorized(log types.Log) (*BridgeRouteAuthorized, error) {
	event := new(BridgeRouteAuthorized)
	if err := _Bridge.contract.UnpackLog(event, "RouteAuthorized", log); err != nil {
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

// BridgeTokenMintFailedIterator is returned from FilterTokenMintFailed and is used to iterate over the raw logs and unpacked data for TokenMintFailed events raised by the Bridge contract.
type BridgeTokenMintFailedIterator struct {
	Event *BridgeTokenMintFailed // Event containing the contract specifics and raw log

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
func (it *BridgeTokenMintFailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeTokenMintFailed)
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
		it.Event = new(BridgeTokenMintFailed)
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
func (it *BridgeTokenMintFailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeTokenMintFailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeTokenMintFailed represents a TokenMintFailed event raised by the Bridge contract.
type BridgeTokenMintFailed struct {
	Token common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTokenMintFailed is a free log retrieval operation binding the contract event 0x259005b50cf55d190d280dfa8480709f3e27c4e6ecd89e3d491dfd288f6ce385.
//
// Solidity: event TokenMintFailed(address indexed token, address indexed to, uint256 value)
func (_Bridge *BridgeFilterer) FilterTokenMintFailed(opts *bind.FilterOpts, token []common.Address, to []common.Address) (*BridgeTokenMintFailedIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "TokenMintFailed", tokenRule, toRule)
	if err != nil {
		return nil, err
	}
	return &BridgeTokenMintFailedIterator{contract: _Bridge.contract, event: "TokenMintFailed", logs: logs, sub: sub}, nil
}

// WatchTokenMintFailed is a free log subscription operation binding the contract event 0x259005b50cf55d190d280dfa8480709f3e27c4e6ecd89e3d491dfd288f6ce385.
//
// Solidity: event TokenMintFailed(address indexed token, address indexed to, uint256 value)
func (_Bridge *BridgeFilterer) WatchTokenMintFailed(opts *bind.WatchOpts, sink chan<- *BridgeTokenMintFailed, token []common.Address, to []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "TokenMintFailed", tokenRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeTokenMintFailed)
				if err := _Bridge.contract.UnpackLog(event, "TokenMintFailed", log); err != nil {
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

// ParseTokenMintFailed is a log parse operation binding the contract event 0x259005b50cf55d190d280dfa8480709f3e27c4e6ecd89e3d491dfd288f6ce385.
//
// Solidity: event TokenMintFailed(address indexed token, address indexed to, uint256 value)
func (_Bridge *BridgeFilterer) ParseTokenMintFailed(log types.Log) (*BridgeTokenMintFailed, error) {
	event := new(BridgeTokenMintFailed)
	if err := _Bridge.contract.UnpackLog(event, "TokenMintFailed", log); err != nil {
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
	Success    bool
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterTokenReceived is a free log retrieval operation binding the contract event 0x7fd41495160762948bbb964edcce550c26992b895a329f399e047786f2097366.
//
// Solidity: event TokenReceived(uint64 indexed srcChainId, address indexed to, uint256 value, bool indexed success)
func (_Bridge *BridgeFilterer) FilterTokenReceived(opts *bind.FilterOpts, srcChainId []uint64, to []common.Address, success []bool) (*BridgeTokenReceivedIterator, error) {

	var srcChainIdRule []interface{}
	for _, srcChainIdItem := range srcChainId {
		srcChainIdRule = append(srcChainIdRule, srcChainIdItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	var successRule []interface{}
	for _, successItem := range success {
		successRule = append(successRule, successItem)
	}

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "TokenReceived", srcChainIdRule, toRule, successRule)
	if err != nil {
		return nil, err
	}
	return &BridgeTokenReceivedIterator{contract: _Bridge.contract, event: "TokenReceived", logs: logs, sub: sub}, nil
}

// WatchTokenReceived is a free log subscription operation binding the contract event 0x7fd41495160762948bbb964edcce550c26992b895a329f399e047786f2097366.
//
// Solidity: event TokenReceived(uint64 indexed srcChainId, address indexed to, uint256 value, bool indexed success)
func (_Bridge *BridgeFilterer) WatchTokenReceived(opts *bind.WatchOpts, sink chan<- *BridgeTokenReceived, srcChainId []uint64, to []common.Address, success []bool) (event.Subscription, error) {

	var srcChainIdRule []interface{}
	for _, srcChainIdItem := range srcChainId {
		srcChainIdRule = append(srcChainIdRule, srcChainIdItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	var successRule []interface{}
	for _, successItem := range success {
		successRule = append(successRule, successItem)
	}

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "TokenReceived", srcChainIdRule, toRule, successRule)
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

// ParseTokenReceived is a log parse operation binding the contract event 0x7fd41495160762948bbb964edcce550c26992b895a329f399e047786f2097366.
//
// Solidity: event TokenReceived(uint64 indexed srcChainId, address indexed to, uint256 value, bool indexed success)
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

// BridgeTokenTransferFailedIterator is returned from FilterTokenTransferFailed and is used to iterate over the raw logs and unpacked data for TokenTransferFailed events raised by the Bridge contract.
type BridgeTokenTransferFailedIterator struct {
	Event *BridgeTokenTransferFailed // Event containing the contract specifics and raw log

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
func (it *BridgeTokenTransferFailedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BridgeTokenTransferFailed)
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
		it.Event = new(BridgeTokenTransferFailed)
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
func (it *BridgeTokenTransferFailedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BridgeTokenTransferFailedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BridgeTokenTransferFailed represents a TokenTransferFailed event raised by the Bridge contract.
type BridgeTokenTransferFailed struct {
	Token common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTokenTransferFailed is a free log retrieval operation binding the contract event 0xc87767983e580cd51a7614924de0506ff919e5220d94509794cd03bb6564b0bf.
//
// Solidity: event TokenTransferFailed(address indexed token, address indexed to, uint256 value)
func (_Bridge *BridgeFilterer) FilterTokenTransferFailed(opts *bind.FilterOpts, token []common.Address, to []common.Address) (*BridgeTokenTransferFailedIterator, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Bridge.contract.FilterLogs(opts, "TokenTransferFailed", tokenRule, toRule)
	if err != nil {
		return nil, err
	}
	return &BridgeTokenTransferFailedIterator{contract: _Bridge.contract, event: "TokenTransferFailed", logs: logs, sub: sub}, nil
}

// WatchTokenTransferFailed is a free log subscription operation binding the contract event 0xc87767983e580cd51a7614924de0506ff919e5220d94509794cd03bb6564b0bf.
//
// Solidity: event TokenTransferFailed(address indexed token, address indexed to, uint256 value)
func (_Bridge *BridgeFilterer) WatchTokenTransferFailed(opts *bind.WatchOpts, sink chan<- *BridgeTokenTransferFailed, token []common.Address, to []common.Address) (event.Subscription, error) {

	var tokenRule []interface{}
	for _, tokenItem := range token {
		tokenRule = append(tokenRule, tokenItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Bridge.contract.WatchLogs(opts, "TokenTransferFailed", tokenRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BridgeTokenTransferFailed)
				if err := _Bridge.contract.UnpackLog(event, "TokenTransferFailed", log); err != nil {
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

// ParseTokenTransferFailed is a log parse operation binding the contract event 0xc87767983e580cd51a7614924de0506ff919e5220d94509794cd03bb6564b0bf.
//
// Solidity: event TokenTransferFailed(address indexed token, address indexed to, uint256 value)
func (_Bridge *BridgeFilterer) ParseTokenTransferFailed(log types.Log) (*BridgeTokenTransferFailed, error) {
	event := new(BridgeTokenTransferFailed)
	if err := _Bridge.contract.UnpackLog(event, "TokenTransferFailed", log); err != nil {
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
