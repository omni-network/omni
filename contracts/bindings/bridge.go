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
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"receiveDefaultGasLimit\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"receiveLockboxGasLimit\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"DEFAULT_ADMIN_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"PAUSER_ROLE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"bridgeFee\",\"inputs\":[{\"name\":\"destChainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"fee\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"claimable\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"defaultConfLevel\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRoleAdmin\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getRoute\",\"inputs\":[{\"name\":\"destChainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"bridge\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"hasLockbox\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"grantRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"hasRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"admin_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"pauser_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"omni_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"token_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"lockbox_\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"lockbox\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"omni\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIOmniPortal\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"receiveToken\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"callerConfirmation\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"retryTransfer\",\"inputs\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"revokeRole\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"sendToken\",\"inputs\":[{\"name\":\"destChainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"wrap\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"setRoutes\",\"inputs\":[{\"name\":\"chainIds\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"},{\"name\":\"routes\",\"type\":\"tuple[]\",\"internalType\":\"structIBridge.Route[]\",\"components\":[{\"name\":\"bridge\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"hasLockbox\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportsInterface\",\"inputs\":[{\"name\":\"interfaceId\",\"type\":\"bytes4\",\"internalType\":\"bytes4\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"token\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"DefaultConfLevelSet\",\"inputs\":[{\"name\":\"conf\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"LockboxWithdrawalFailed\",\"inputs\":[{\"name\":\"badLockbox\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OmniPortalSet\",\"inputs\":[{\"name\":\"omni\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RetrySuccessful\",\"inputs\":[{\"name\":\"caller\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleAdminChanged\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"previousAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"newAdminRole\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleGranted\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RoleRevoked\",\"inputs\":[{\"name\":\"role\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"},{\"name\":\"account\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"RouteConfigured\",\"inputs\":[{\"name\":\"destChainId\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"bridge\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"hasLockbox\",\"type\":\"bool\",\"indexed\":true,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokenMintFailed\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokenReceived\",\"inputs\":[{\"name\":\"srcChainId\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokenSent\",\"inputs\":[{\"name\":\"destChainId\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"from\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AccessControlBadConfirmation\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AccessControlUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"neededRole\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"ArrayLengthMismatch\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"CannotWrap\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"EnforcedPause\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ExpectedPause\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientPayment\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidRoute\",\"inputs\":[{\"name\":\"chainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"type\":\"error\",\"name\":\"NoClaimable\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"Unauthorized\",\"inputs\":[{\"name\":\"chainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ZeroAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ZeroAmount\",\"inputs\":[]}]",
	Bin: "0x60c060405234801561001057600080fd5b5060405161234a38038061234a83398101604081905261002f9161011f565b6001600160401b03808316608052811660a05261004a610051565b5050610152565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000900460ff16156100a15760405163f92ee8a960e01b815260040160405180910390fd5b80546001600160401b03908116146101005780546001600160401b0319166001600160401b0390811782556040519081527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b50565b80516001600160401b038116811461011a57600080fd5b919050565b6000806040838503121561013257600080fd5b61013b83610103565b915061014960208401610103565b90509250929050565b60805160a0516121c561018560003960008181610b0101526119a8015260008181610adb015261198201526121c56000f3fe6080604052600436106101405760003560e01c806374eeb847116100b657806397235a1e1161006f57806397235a1e14610413578063a217fddf14610433578063a48fcd8e14610448578063d547741f1461045b578063e63ab1e91461047b578063fc0c546a1461049d57600080fd5b806374eeb8471461034b57806378f8d76e1461037e5780638456cb591461039e5780638ad201f6146103b357806391d14854146103d357806392099d44146103f357600080fd5b806339acf9f11161010857806339acf9f11461020a5780633f4ba83a14610242578063402914f5146102575780634acd8d82146102845780635c975abb1461030657806366cc57021461032b57600080fd5b806301ffc9a7146101455780631459457a1461017a578063248a9ca31461019c5780632f2ff15d146101ca57806336568abe146101ea575b600080fd5b34801561015157600080fd5b50610165610160366004611c2d565b6104bd565b60405190151581526020015b60405180910390f35b34801561018657600080fd5b5061019a610195366004611c6c565b6104f4565b005b3480156101a857600080fd5b506101bc6101b7366004611cdd565b6107ba565b604051908152602001610171565b3480156101d657600080fd5b5061019a6101e5366004611cf6565b6107dc565b3480156101f657600080fd5b5061019a610205366004611cf6565b6107fe565b34801561021657600080fd5b5060005461022a906001600160a01b031681565b6040516001600160a01b039091168152602001610171565b34801561024e57600080fd5b5061019a610836565b34801561026357600080fd5b506101bc610272366004611d26565b60356020526000908152604090205481565b34801561029057600080fd5b506102e761029f366004611d58565b6001600160401b03166000908152603460209081526040918290208251808401909352546001600160a01b038116808452600160a01b90910460ff1615159290910182905291565b604080516001600160a01b039093168352901515602083015201610171565b34801561031257600080fd5b506000805160206121708339815191525460ff16610165565b34801561033757600080fd5b5060335461022a906001600160a01b031681565b34801561035757600080fd5b5060005461036c90600160a01b900460ff1681565b60405160ff9091168152602001610171565b34801561038a57600080fd5b5061019a610399366004611d26565b610859565b3480156103aa57600080fd5b5061019a610a2b565b3480156103bf57600080fd5b506101bc6103ce366004611d58565b610a4b565b3480156103df57600080fd5b506101656103ee366004611cf6565b610b2c565b3480156103ff57600080fd5b5061019a61040e366004611dc0565b610b64565b34801561041f57600080fd5b5061019a61042e366004611e60565b610d1d565b34801561043f57600080fd5b506101bc600081565b61019a610456366004611e9a565b610e92565b34801561046757600080fd5b5061019a610476366004611cf6565b610eb2565b34801561048757600080fd5b506101bc60008051602061213083398151915281565b3480156104a957600080fd5b5060325461022a906001600160a01b031681565b60006001600160e01b03198216637965db0b60e01b14806104ee57506301ffc9a760e01b6001600160e01b03198316145b92915050565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a008054600160401b810460ff1615906001600160401b03166000811580156105395750825b90506000826001600160401b031660011480156105555750303b155b905081158015610563575080155b156105815760405163f92ee8a960e01b815260040160405180910390fd5b845467ffffffffffffffff1916600117855583156105ab57845460ff60401b1916600160401b1785555b6001600160a01b038a166105d25760405163d92e233d60e01b815260040160405180910390fd5b6001600160a01b0389166105f95760405163d92e233d60e01b815260040160405180910390fd5b6001600160a01b0388166106205760405163d92e233d60e01b815260040160405180910390fd5b6001600160a01b0387166106475760405163d92e233d60e01b815260040160405180910390fd5b61064f610ece565b610657610ed8565b610662886004610ee8565b61066d60008b610f06565b506106866000805160206121308339815191528a610f06565b50603280546001600160a01b0319166001600160a01b03898116919091179091558616156106ca57603380546001600160a01b0319166001600160a01b0388161790555b6001600160a01b038616156107685761075286600019886001600160a01b031663fc0c546a6040518163ffffffff1660e01b8152600401602060405180830381865afa15801561071e573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906107429190611eed565b6001600160a01b03169190610fab565b6107686001600160a01b03881687600019610fab565b83156107ae57845460ff60401b19168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b50505050505050505050565b6000908152600080516020612150833981519152602052604090206001015490565b6107e5826107ba565b6107ee81611042565b6107f88383610f06565b50505050565b6001600160a01b03811633146108275760405163334bd91960e11b815260040160405180910390fd5b610831828261104c565b505050565b60008051602061213083398151915261084e81611042565b6108566110c8565b50565b6001600160a01b0381166000908152603560205260408120549081900361089357604051631129777360e21b815260040160405180910390fd5b6001600160a01b038083166000908152603560205260408120556033541661091e576032546040516340c10f1960e01b81526001600160a01b03909116906340c10f19906108e79085908590600401611f0a565b600060405180830381600087803b15801561090157600080fd5b505af1158015610915573d6000803e3d6000fd5b505050506109e7565b6032546040516340c10f1960e01b81526001600160a01b03909116906340c10f19906109509030908590600401611f0a565b600060405180830381600087803b15801561096a57600080fd5b505af115801561097e573d6000803e3d6000fd5b505060335460405163040b850f60e31b81526001600160a01b03909116925063205c287891506109b49085908590600401611f0a565b600060405180830381600087803b1580156109ce57600080fd5b505af11580156109e2573d6000803e3d6000fd5b505050505b6040518181526001600160a01b0383169033907fafc6034e5bc12b75c8fd712cc3306dba0afd7d2c5156fe40015ff2b3551f86c09060200160405180910390a35050565b600080516020612130833981519152610a4381611042565b610856611129565b6001600160401b03811660009081526034602090815260408083208151808301835290546001600160a01b0381168252600160a01b900460ff1615159281019290925251610b25908490610aa790600019908190602401611f0a565b60408051601f19818403018152919052602080820180516001600160e01b0316634b91ad0f60e11b179052840151610aff577f0000000000000000000000000000000000000000000000000000000000000000611172565b7f0000000000000000000000000000000000000000000000000000000000000000611172565b9392505050565b6000918252600080516020612150833981519152602090815260408084206001600160a01b0393909316845291905290205460ff1690565b6000610b6f81611042565b838214610b8f5760405163512509d360e11b815260040160405180910390fd5b60005b84811015610d15576000848483818110610bae57610bae611f23565b610bc49260206040909202019081019150611d26565b6001600160a01b031603610beb5760405163d92e233d60e01b815260040160405180910390fd5b838382818110610bfd57610bfd611f23565b90506040020160346000888885818110610c1957610c19611f23565b9050602002016020810190610c2e9190611d58565b6001600160401b031681526020810191909152604001600020610c518282611f39565b905050838382818110610c6657610c66611f23565b9050604002016020016020810190610c7e9190611f92565b1515848483818110610c9257610c92611f23565b610ca89260206040909202019081019150611d26565b6001600160a01b0316878784818110610cc357610cc3611f23565b9050602002016020810190610cd89190611d58565b6001600160401b03167fab19e99f8223191275fefd1410893ea2b3001028e27ab75b975987c1c8c4320760405160405180910390a4600101610b92565b505050505050565b60005460408051631799380760e11b815281516001600160a01b0390931692632f32700e926004808401939192918290030181865afa158015610d64573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610d889190611faf565b8051600180546020909301516001600160401b039092166001600160e01b031990931692909217600160401b6001600160a01b039283160217909155600054163303610e4a576001546001600160401b0381166000908152603460205260409020546001600160a01b03908116600160401b9092041614610e4557600154604051633dfc334560e01b81526001600160401b0382166004820152600160401b9091046001600160a01b031660248201526044015b60405180910390fd5b610e74565b604051633dfc334560e01b81526001600160401b0346166004820152336024820152604401610e3c565b610e7e82826111f0565b5050600180546001600160e01b0319169055565b610e9a611290565b610ea6848484846112c1565b6107f884848484611387565b610ebb826107ba565b610ec481611042565b6107f8838361104c565b610ed66114e8565b565b610ee06114e8565b610ed6611531565b610ef06114e8565b610ef982611552565b610f02816115eb565b5050565b6000600080516020612150833981519152610f218484610b2c565b610fa1576000848152602082815260408083206001600160a01b03871684529091529020805460ff19166001179055610f573390565b6001600160a01b0316836001600160a01b0316857f2f8788117e7eff1d82e926ec794901d17c78024a50270940304540a733656f0d60405160405180910390a460019150506104ee565b60009150506104ee565b816014528060345263095ea7b360601b60005260206000604460106000875af1806001600051141661103757803d853b15171061103757600060345263095ea7b360601b600052600038604460106000885af1508160345260206000604460106000885af19050806001600051141661103757803d853b15171061103757633e3f8f736000526004601cfd5b506000603452505050565b610856813361168e565b60006000805160206121508339815191526110678484610b2c565b15610fa1576000848152602082815260408083206001600160a01b0387168085529252808320805460ff1916905551339287917ff6391f5c32d9c69d2a47ea670b442974b53935d1edc7fd64eb21e047a839171b9190a460019150506104ee565b6110d06116b9565b600080516020612170833981519152805460ff191681557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa335b6040516001600160a01b0390911681526020015b60405180910390a150565b611131611290565b600080516020612170833981519152805460ff191660011781557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a2583361110a565b60008054604051632376548f60e21b81526001600160a01b0390911690638dd9523c906111a790879087908790600401612063565b602060405180830381865afa1580156111c4573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906111e8919061209e565b949350505050565b6001546040518281526001600160a01b038416916001600160401b0316907ffd6bb19aab6d8ed3fbc8848d3b1a76c95135e46d05bec2b84424219f9681929f9060200160405180910390a36032546033546001600160a01b039182169116806112665761126082858560006116e9565b506107f8565b600061127583868660016116e9565b9050801561128957611289838387876117c5565b5050505050565b6000805160206121708339815191525460ff1615610ed65760405163d93c066560e01b815260040160405180910390fd5b6001600160401b0384166000908152603460205260409020546001600160a01b031661130b5760405163f6e829e160e01b81526001600160401b0385166004820152602401610e3c565b6001600160a01b0383166113325760405163d92e233d60e01b815260040160405180910390fd5b8160000361135357604051631f2a200560e01b815260040160405180910390fd5b80801561136957506033546001600160a01b0316155b156107f857604051637d25d4c960e11b815260040160405180910390fd5b6032546033546001600160a01b039182169116821561147d57611419333086846001600160a01b031663fc0c546a6040518163ffffffff1660e01b8152600401602060405180830381865afa1580156113e4573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906114089190611eed565b6001600160a01b0316929190611888565b60405160016255295b60e01b031981526001600160a01b0382169063ffaad6a59061144a9033908890600401611f0a565b600060405180830381600087803b15801561146457600080fd5b505af1158015611478573d6000803e3d6000fd5b505050505b6040516379ef98bf60e11b81526001600160a01b0383169063f3df317e906114ab9033908890600401611f0a565b600060405180830381600087803b1580156114c557600080fd5b505af11580156114d9573d6000803e3d6000fd5b50505050610d158686866118e6565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a0054600160401b900460ff16610ed657604051631afcd79f60e31b815260040160405180910390fd5b6115396114e8565b600080516020612170833981519152805460ff19169055565b6001600160a01b03811661159d5760405162461bcd60e51b8152602060048201526012602482015271584170703a206e6f207a65726f206f6d6e6960701b6044820152606401610e3c565b600080546001600160a01b0319166001600160a01b0383169081179091556040519081527f79162c8d053a07e70cdc1ccc536f0440b571f8508377d2bef51094fadab98f479060200161111e565b6115f481611a6a565b6116405760405162461bcd60e51b815260206004820152601860248201527f584170703a20696e76616c696420636f6e66206c6576656c00000000000000006044820152606401610e3c565b6000805460ff60a01b1916600160a01b60ff8416908102919091179091556040519081527f8de08a798b4e50b4f351c1eaa91a11530043802be3ffac2df87db0c45a2e84839060200161111e565b6116988282610b2c565b610f0257808260405163e2517d3f60e01b8152600401610e3c929190611f0a565b6000805160206121708339815191525460ff16610ed657604051638dfc202b60e01b815260040160405180910390fd5b6000846001600160a01b03166340c10f19836117055785611707565b305b856040518363ffffffff1660e01b8152600401611725929190611f0a565b600060405180830381600087803b15801561173f57600080fd5b505af1925050508015611750575060015b6117ba576001600160a01b03808516600081815260356020526040908190208054870190555190918716907f259005b50cf55d190d280dfa8480709f3e27c4e6ecd89e3d491dfd288f6ce385906117aa9087815260200190565b60405180910390a35060006111e8565b506001949350505050565b60405163040b850f60e31b81526001600160a01b0384169063205c2878906117f39085908590600401611f0a565b600060405180830381600087803b15801561180d57600080fd5b505af192505050801561181e575060015b6107f8576118366001600160a01b0385168383611a86565b816001600160a01b0316836001600160a01b03167fd5684d2e31a0d2443b1102269ffecfa6d05a6a36e312219f9108c0e95f8bc9af8360405161187b91815260200190565b60405180910390a36107f8565b60405181606052826040528360601b602c526323b872dd60601b600c52602060006064601c6000895af180600160005114166118d757803d873b1517106118d757637939f4246000526004601cfd5b50600060605260405250505050565b6001600160401b03831660009081526034602090815260408083208151808301835290546001600160a01b0381168252600160a01b900460ff161515928101929092525190919061193d9085908590602401611f0a565b60408051601f19818403018152919052602080820180516001600160e01b0316634b91ad0f60e11b1790528351908401519192506000916119cc91889185906119a6577f0000000000000000000000000000000000000000000000000000000000000000611acb565b7f0000000000000000000000000000000000000000000000000000000000000000611acb565b9050803410156119ef5760405163cd1c886760e01b815260040160405180910390fd5b80341115611a0b57611a0b611a0482346120b7565b3390611c11565b846001600160a01b0316336001600160a01b0316876001600160401b03167fc0464e720761b6de8643eed8d1cbf17ec66c3eb60179efaed5a58cf75580a4dc87604051611a5a91815260200190565b60405180910390a4505050505050565b600060ff8216600114806104ee575060ff821660041492915050565b816014528060345263a9059cbb60601b60005260206000604460106000875af1806001600051141661103757803d853b151710611037576390b8ec186000526004601cfd5b60008054604051632376548f60e21b815282916001600160a01b031690638dd9523c90611b0090899088908890600401612063565b602060405180830381865afa158015611b1d573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611b41919061209e565b905080471015611b935760405162461bcd60e51b815260206004820152601860248201527f584170703a20696e73756666696369656e742066756e647300000000000000006044820152606401610e3c565b60005460405163c21dda4f60e01b81526001600160a01b0382169163c21dda4f918491611bd5918b91600160a01b900460ff16908b908b908b906004016120d8565b6000604051808303818588803b158015611bee57600080fd5b505af1158015611c02573d6000803e3d6000fd5b50939998505050505050505050565b60003860003884865af1610f025763b12d13eb6000526004601cfd5b600060208284031215611c3f57600080fd5b81356001600160e01b031981168114610b2557600080fd5b6001600160a01b038116811461085657600080fd5b600080600080600060a08688031215611c8457600080fd5b8535611c8f81611c57565b94506020860135611c9f81611c57565b93506040860135611caf81611c57565b92506060860135611cbf81611c57565b91506080860135611ccf81611c57565b809150509295509295909350565b600060208284031215611cef57600080fd5b5035919050565b60008060408385031215611d0957600080fd5b823591506020830135611d1b81611c57565b809150509250929050565b600060208284031215611d3857600080fd5b8135610b2581611c57565b6001600160401b038116811461085657600080fd5b600060208284031215611d6a57600080fd5b8135610b2581611d43565b60008083601f840112611d8757600080fd5b5081356001600160401b03811115611d9e57600080fd5b6020830191508360208260061b8501011115611db957600080fd5b9250929050565b60008060008060408587031215611dd657600080fd5b84356001600160401b03811115611dec57600080fd5b8501601f81018713611dfd57600080fd5b80356001600160401b03811115611e1357600080fd5b8760208260051b8401011115611e2857600080fd5b6020918201955093508501356001600160401b03811115611e4857600080fd5b611e5487828801611d75565b95989497509550505050565b60008060408385031215611e7357600080fd5b8235611e7e81611c57565b946020939093013593505050565b801515811461085657600080fd5b60008060008060808587031215611eb057600080fd5b8435611ebb81611d43565b93506020850135611ecb81611c57565b9250604085013591506060850135611ee281611e8c565b939692955090935050565b600060208284031215611eff57600080fd5b8151610b2581611c57565b6001600160a01b03929092168252602082015260400190565b634e487b7160e01b600052603260045260246000fd5b8135611f4481611c57565b81546001600160a01b031981166001600160a01b039290921691821783556020840135611f7081611e8c565b6001600160a81b03199190911690911790151560a01b60ff60a01b1617905550565b600060208284031215611fa457600080fd5b8135610b2581611e8c565b60006040828403128015611fc257600080fd5b50604080519081016001600160401b0381118282101715611ff357634e487b7160e01b600052604160045260246000fd5b604052825161200181611d43565b8152602083015161201181611c57565b60208201529392505050565b6000815180845260005b8181101561204357602081850181015186830182015201612027565b506000602082860101526020601f19601f83011685010191505092915050565b6001600160401b0384168152606060208201526000612085606083018561201d565b90506001600160401b0383166040830152949350505050565b6000602082840312156120b057600080fd5b5051919050565b818103818111156104ee57634e487b7160e01b600052601160045260246000fd5b6001600160401b038616815260ff851660208201526001600160a01b038416604082015260a0606082018190526000906121149083018561201d565b90506001600160401b0383166080830152969550505050505056fe539440820030c4994db4e31b6b800deafd503688728f932addfe7a410515c14c02dd7bc7dec4dceedda775e58dd541e08a116c6c53815c0bd028192f7b626800cd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f03300a264697066735822122070ef0a1d7929ac75e007228c1a4da329162f62dd41c1b4fddd996cea71f2221564736f6c634300081a0033",
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

// RetryTransfer is a paid mutator transaction binding the contract method 0x78f8d76e.
//
// Solidity: function retryTransfer(address addr) returns()
func (_Bridge *BridgeTransactor) RetryTransfer(opts *bind.TransactOpts, addr common.Address) (*types.Transaction, error) {
	return _Bridge.contract.Transact(opts, "retryTransfer", addr)
}

// RetryTransfer is a paid mutator transaction binding the contract method 0x78f8d76e.
//
// Solidity: function retryTransfer(address addr) returns()
func (_Bridge *BridgeSession) RetryTransfer(addr common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.RetryTransfer(&_Bridge.TransactOpts, addr)
}

// RetryTransfer is a paid mutator transaction binding the contract method 0x78f8d76e.
//
// Solidity: function retryTransfer(address addr) returns()
func (_Bridge *BridgeTransactorSession) RetryTransfer(addr common.Address) (*types.Transaction, error) {
	return _Bridge.Contract.RetryTransfer(&_Bridge.TransactOpts, addr)
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
