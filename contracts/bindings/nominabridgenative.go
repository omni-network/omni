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

// NominaBridgeNativeMetaData contains all meta data concerning the NominaBridgeNative contract.
var NominaBridgeNativeMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"ACTION_BRIDGE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"ACTION_WITHDRAW\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"KeyPauseAll\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"XCALL_WITHDRAW_GAS_LIMIT\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"bridge\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"bridgeFee\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"claim\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"claimable\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"owner_\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"initializeV2\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isPaused\",\"inputs\":[{\"name\":\"action\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"l1Bridge\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"l1ChainId\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"l1Deposits\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[{\"name\":\"action\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"portal\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIOmniPortal\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setup\",\"inputs\":[{\"name\":\"l1ChainId_\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"portal_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"l1Bridge_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"l1Deposits_\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[{\"name\":\"action\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdraw\",\"inputs\":[{\"name\":\"payor\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Bridge\",\"inputs\":[{\"name\":\"payor\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Claimed\",\"inputs\":[{\"name\":\"claimant\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"key\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Setup\",\"inputs\":[{\"name\":\"l1ChainId\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"portal\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"l1Bridge\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"l1Deposits\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"key\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Withdraw\",\"inputs\":[{\"name\":\"payor\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"success\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnableInvalidOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnableUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]}]",
	Bin: "0x608060405234801561000f575f80fd5b5061001861001d565b6100cf565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000900460ff161561006d5760405163f92ee8a960e01b815260040160405180910390fd5b80546001600160401b03908116146100cc5780546001600160401b0319166001600160401b0390811782556040519081527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b50565b6117d6806100dc5f395ff3fe608060405260043610610147575f3560e01c80636425666b116100b3578063a10ac97a1161006d578063a10ac97a146103e5578063c3de453d14610405578063c4d66de814610418578063d9caed1214610437578063ed56531a14610456578063f2fde38b14610475575f80fd5b80636425666b1461030f578063715018a61461034c5780638456cb59146103605780638da5cb5b146103745780638fdcb4c9146103b0578063969b53da146103c6575f80fd5b806332c8bb771161010457806332c8bb77146102695780633abfe55f1461027e5780633f4ba83a1461029d578063402914f5146102b1578063499e85cd146102dc5780635cd8a76b146102fb575f80fd5b806309839a931461014b57806312622e5b146101915780631e83409a146101c7578063241b71bb146101e857806325d70f78146102175780632f4dae9f1461024a575b5f80fd5b348015610156575f80fd5b5061017e7f0683d1c283a672fc58eb7940a0dba83ea98b96966a9ca1b030dec2c60cea4d1e81565b6040519081526020015b60405180910390f35b34801561019c575f80fd5b505f546101af906001600160401b031681565b6040516001600160401b039091168152602001610188565b3480156101d2575f80fd5b506101e66101e13660046114c0565b610494565b005b3480156101f3575f80fd5b506102076102023660046114db565b6107d4565b6040519015158152602001610188565b348015610222575f80fd5b5061017e7f855511cc3694f64379908437d6d64458dc76d02482052bfb8a5b33a72c054c7781565b348015610255575f80fd5b506101e66102643660046114db565b6107e4565b348015610274575f80fd5b5061017e60015481565b348015610289575f80fd5b5061017e6102983660046114f2565b6107f8565b3480156102a8575f80fd5b506101e66108c8565b3480156102bc575f80fd5b5061017e6102cb3660046114c0565b60036020525f908152604090205481565b3480156102e7575f80fd5b506101e66102f6366004611530565b6108da565b348015610306575f80fd5b506101e661097f565b34801561031a575f80fd5b505f5461033490600160401b90046001600160a01b031681565b6040516001600160a01b039091168152602001610188565b348015610357575f80fd5b506101e6610a55565b34801561036b575f80fd5b506101e6610a66565b34801561037f575f80fd5b507f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c199300546001600160a01b0316610334565b3480156103bb575f80fd5b506101af6201388081565b3480156103d1575f80fd5b50600254610334906001600160a01b031681565b3480156103f0575f80fd5b5061017e5f8051602061178183398151915281565b6101e66104133660046114f2565b610a76565b348015610423575f80fd5b506101e66104323660046114c0565b610acc565b348015610442575f80fd5b506101e661045136600461157e565b610bc3565b348015610461575f80fd5b506101e66104703660046114db565b610e82565b348015610480575f80fd5b506101e661048f3660046114c0565b610e93565b7f855511cc3694f64379908437d6d64458dc76d02482052bfb8a5b33a72c054c776104be81610ecd565b156104e45760405162461bcd60e51b81526004016104db906115bc565b60405180910390fd5b5f8060089054906101000a90046001600160a01b03166001600160a01b0316632f32700e6040518163ffffffff1660e01b81526004016040805180830381865afa158015610534573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061055891906115ea565b5f54909150600160401b90046001600160a01b031633146105b55760405162461bcd60e51b8152602060048201526017602482015276139bdb5a5b98509c9a5919d94e881b9bdd081e18d85b1b604a1b60448201526064016104db565b5f5481516001600160401b0390811691161461060a5760405162461bcd60e51b81526020600482015260146024820152734e6f6d696e614272696467653a206e6f74204c3160601b60448201526064016104db565b6001600160a01b0383166106605760405162461bcd60e51b815260206004820152601e60248201527f4e6f6d696e614272696467653a206e6f20636c61696d20746f207a65726f000060448201526064016104db565b6020808201516001600160a01b0381165f90815260039092526040909120546106cb5760405162461bcd60e51b815260206004820152601e60248201527f4e6f6d696e614272696467653a206e6f7468696e6720746f20636c61696d000060448201526064016104db565b6001600160a01b038181165f908152600360205260408082208054908390559051909287169083908381818185875af1925050503d805f8114610729576040519150601f19603f3d011682016040523d82523d5f602084013e61072e565b606091505b505090508061077f5760405162461bcd60e51b815260206004820152601d60248201527f4e6f6d696e614272696467653a207472616e73666572206661696c656400000060448201526064016104db565b856001600160a01b0316836001600160a01b03167ff7a40077ff7a04c7e61f6f26fb13774259ddf1b6bce9ecf26a8276cdd3992683846040516107c491815260200190565b60405180910390a3505050505050565b5f6107de82610ecd565b92915050565b6107ec610f43565b6107f581610f9e565b50565b5f80546040516001600160a01b03858116602483015260448201859052600160401b83041691638dd9523c916001600160401b039091169060640160408051601f198184030181529181526020820180516001600160e01b031663f3fef3a360e01b179052516001600160e01b031960e085901b1681526108829291906201388090600401611695565b602060405180830381865afa15801561089d573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906108c191906116ca565b9392505050565b6108d0610f43565b6108d8611054565b565b6108e2610f43565b5f80546001600160a01b03858116600160401b81026001600160e01b03199093166001600160401b03891690811793909317909355600280549186166001600160a01b0319909216821790556001849055604080519283526020830193909352818301526060810183905290517fc12d4a2db17193df88185c2dc087fa9536c12710f1381b359b80553e5d9a12939181900360800190a150505050565b60025f61098a61106a565b8054909150600160401b900460ff16806109b1575080546001600160401b03808416911610155b156109cf5760405163f92ee8a960e01b815260040160405180910390fd5b805468ffffffffffffffffff19166001600160401b03831617600160401b17815560018054604b91905f90610a059084906116f5565b9091555050805460ff60401b191681556040516001600160401b03831681527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15050565b610a5d610f43565b6108d85f611092565b610a6e610f43565b6108d8611102565b7f0683d1c283a672fc58eb7940a0dba83ea98b96966a9ca1b030dec2c60cea4d1e610aa081610ecd565b15610abd5760405162461bcd60e51b81526004016104db906115bc565b610ac78383611118565b505050565b5f610ad561106a565b805490915060ff600160401b82041615906001600160401b03165f81158015610afb5750825b90505f826001600160401b03166001148015610b165750303b155b905081158015610b24575080155b15610b425760405163f92ee8a960e01b815260040160405180910390fd5b845467ffffffffffffffff191660011785558315610b6c57845460ff60401b1916600160401b1785555b610b758661139f565b8315610bbb57845460ff60401b19168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b505050505050565b7f855511cc3694f64379908437d6d64458dc76d02482052bfb8a5b33a72c054c77610bed81610ecd565b15610c0a5760405162461bcd60e51b81526004016104db906115bc565b5f8060089054906101000a90046001600160a01b03166001600160a01b0316632f32700e6040518163ffffffff1660e01b81526004016040805180830381865afa158015610c5a573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610c7e91906115ea565b5f54909150600160401b90046001600160a01b03163314610cdb5760405162461bcd60e51b8152602060048201526017602482015276139bdb5a5b98509c9a5919d94e881b9bdd081e18d85b1b604a1b60448201526064016104db565b60025460208201516001600160a01b03908116911614610d3d5760405162461bcd60e51b815260206004820152601860248201527f4e6f6d696e614272696467653a206e6f7420627269646765000000000000000060448201526064016104db565b5f5481516001600160401b03908116911614610d925760405162461bcd60e51b81526020600482015260146024820152734e6f6d696e614272696467653a206e6f74204c3160601b60448201526064016104db565b8260015f828254610da3919061170c565b909155505f90506001600160a01b03851684617d005a610dc3919061171f565b6040519091905f818181858888f193505050503d805f8114610e00576040519150601f19603f3d011682016040523d82523d5f602084013e610e05565b606091505b5050905080610e3b576001600160a01b0386165f9081526003602052604081208054869290610e3590849061170c565b90915550505b6040805185815282151560208201526001600160a01b0380881692908916917f2b664ab52fe561d3ace376046aea39744dd736ec1f67d89d504ffd2192825f6191016107c4565b610e8a610f43565b6107f5816113b0565b610e9b610f43565b6001600160a01b038116610ec457604051631e4fbdf760e01b81525f60048201526024016104db565b6107f581611092565b5f805160206117818339815191525f9081527fff37105740f03695c8f3597f3aff2b92fbe1c80abea3c28731ecff2efd69340060208190527ffae9838a178d7f201aa98e2ce5340158edda60bb1e8f168f46503bf3e99f13be5460ff16806108c157505f92835260205250604090205460ff1690565b33610f757f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c199300546001600160a01b031690565b6001600160a01b0316146108d85760405163118cdaa760e01b81523360048201526024016104db565b5f8181527fff37105740f03695c8f3597f3aff2b92fbe1c80abea3c28731ecff2efd693400602081905260409091205460ff166110145760405162461bcd60e51b815260206004820152601460248201527314185d5cd8589b194e881b9bdd081c185d5cd95960621b60448201526064016104db565b5f82815260208290526040808220805460ff191690555183917fd05bfc2250abb0f8fd265a54c53a24359c5484af63cad2e4ce87c78ab751395a91a25050565b6108d85f80516020611781833981519152610f9e565b5f807ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a006107de565b7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c19930080546001600160a01b031981166001600160a01b03848116918217845560405192169182907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0905f90a3505050565b6108d85f805160206117818339815191526113b0565b6001600160a01b03821661116e5760405162461bcd60e51b815260206004820152601f60248201527f4e6f6d696e614272696467653a206e6f2062726964676520746f207a65726f0060448201526064016104db565b5f81116111bd5760405162461bcd60e51b815260206004820181905260248201527f4e6f6d696e614272696467653a20616d6f756e74206d757374206265203e203060448201526064016104db565b60015481111561120f5760405162461bcd60e51b815260206004820152601a60248201527f4e6f6d696e614272696467653a206e6f206c697175696469747900000000000060448201526064016104db565b61121982826107f8565b611223908261170c565b3410156112725760405162461bcd60e51b815260206004820181905260248201527f4e6f6d696e614272696467653a20696e73756666696369656e742066756e647360448201526064016104db565b8060015f828254611283919061171f565b90915550505f54600160401b90046001600160a01b031663c21dda4f6112a9833461171f565b5f546002546040516001600160a01b038881166024830152604482018890526001600160401b0390931692600492169060640160408051601f198184030181529181526020820180516001600160e01b031663f3fef3a360e01b179052516001600160e01b031960e088901b16815261132d94939291906201388090600401611732565b5f604051808303818588803b158015611344575f80fd5b505af1158015611356573d5f803e3d5ffd5b50506040518481526001600160a01b03861693503392507f59bc8a913d49a9626dd6ba5def7fcf12804061c1bb9b8b6db077e1a12cb4b422915060200160405180910390a35050565b6113a7611466565b6107f58161148b565b5f8181527fff37105740f03695c8f3597f3aff2b92fbe1c80abea3c28731ecff2efd693400602081905260409091205460ff16156114235760405162461bcd60e51b815260206004820152601060248201526f14185d5cd8589b194e881c185d5cd95960821b60448201526064016104db565b5f82815260208290526040808220805460ff191660011790555183917f0cb09dc71d57eeec2046f6854976717e4874a3cf2d6ddeddde337e5b6de6ba3191a25050565b61146e611493565b6108d857604051631afcd79f60e31b815260040160405180910390fd5b610e9b611466565b5f61149c61106a565b54600160401b900460ff16919050565b6001600160a01b03811681146107f5575f80fd5b5f602082840312156114d0575f80fd5b81356108c1816114ac565b5f602082840312156114eb575f80fd5b5035919050565b5f8060408385031215611503575f80fd5b823561150e816114ac565b946020939093013593505050565b6001600160401b03811681146107f5575f80fd5b5f805f8060808587031215611543575f80fd5b843561154e8161151c565b9350602085013561155e816114ac565b9250604085013561156e816114ac565b9396929550929360600135925050565b5f805f60608486031215611590575f80fd5b833561159b816114ac565b925060208401356115ab816114ac565b929592945050506040919091013590565b602080825260149082015273139bdb5a5b98509c9a5919d94e881c185d5cd95960621b604082015260600190565b5f604082840312156115fa575f80fd5b604051604081018181106001600160401b038211171561162857634e487b7160e01b5f52604160045260245ffd5b60405282516116368161151c565b81526020830151611646816114ac565b60208201529392505050565b5f81518084525f5b818110156116765760208185018101518683018201520161165a565b505f602082860101526020601f19601f83011685010191505092915050565b5f6001600160401b038086168352606060208401526116b76060840186611652565b9150808416604084015250949350505050565b5f602082840312156116da575f80fd5b5051919050565b634e487b7160e01b5f52601160045260245ffd5b80820281158282048414176107de576107de6116e1565b808201808211156107de576107de6116e1565b818103818111156107de576107de6116e1565b5f6001600160401b03808816835260ff8716602084015260018060a01b038616604084015260a0606084015261176b60a0840186611652565b9150808416608084015250969550505050505056fe76e8952e4b09b8d505aa08998d716721a1dbf0884ac74202e33985da1ed005e9a2646970667358221220915006044505918c3a7429032cd5d2451b78b6e82e0cb9ed1744bf29f3e3013464736f6c63430008180033",
}

// NominaBridgeNativeABI is the input ABI used to generate the binding from.
// Deprecated: Use NominaBridgeNativeMetaData.ABI instead.
var NominaBridgeNativeABI = NominaBridgeNativeMetaData.ABI

// NominaBridgeNativeBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use NominaBridgeNativeMetaData.Bin instead.
var NominaBridgeNativeBin = NominaBridgeNativeMetaData.Bin

// DeployNominaBridgeNative deploys a new Ethereum contract, binding an instance of NominaBridgeNative to it.
func DeployNominaBridgeNative(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *NominaBridgeNative, error) {
	parsed, err := NominaBridgeNativeMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(NominaBridgeNativeBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &NominaBridgeNative{NominaBridgeNativeCaller: NominaBridgeNativeCaller{contract: contract}, NominaBridgeNativeTransactor: NominaBridgeNativeTransactor{contract: contract}, NominaBridgeNativeFilterer: NominaBridgeNativeFilterer{contract: contract}}, nil
}

// NominaBridgeNative is an auto generated Go binding around an Ethereum contract.
type NominaBridgeNative struct {
	NominaBridgeNativeCaller     // Read-only binding to the contract
	NominaBridgeNativeTransactor // Write-only binding to the contract
	NominaBridgeNativeFilterer   // Log filterer for contract events
}

// NominaBridgeNativeCaller is an auto generated read-only Go binding around an Ethereum contract.
type NominaBridgeNativeCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NominaBridgeNativeTransactor is an auto generated write-only Go binding around an Ethereum contract.
type NominaBridgeNativeTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NominaBridgeNativeFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type NominaBridgeNativeFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NominaBridgeNativeSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NominaBridgeNativeSession struct {
	Contract     *NominaBridgeNative // Generic contract binding to set the session for
	CallOpts     bind.CallOpts       // Call options to use throughout this session
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// NominaBridgeNativeCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NominaBridgeNativeCallerSession struct {
	Contract *NominaBridgeNativeCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts             // Call options to use throughout this session
}

// NominaBridgeNativeTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NominaBridgeNativeTransactorSession struct {
	Contract     *NominaBridgeNativeTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts             // Transaction auth options to use throughout this session
}

// NominaBridgeNativeRaw is an auto generated low-level Go binding around an Ethereum contract.
type NominaBridgeNativeRaw struct {
	Contract *NominaBridgeNative // Generic contract binding to access the raw methods on
}

// NominaBridgeNativeCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NominaBridgeNativeCallerRaw struct {
	Contract *NominaBridgeNativeCaller // Generic read-only contract binding to access the raw methods on
}

// NominaBridgeNativeTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NominaBridgeNativeTransactorRaw struct {
	Contract *NominaBridgeNativeTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNominaBridgeNative creates a new instance of NominaBridgeNative, bound to a specific deployed contract.
func NewNominaBridgeNative(address common.Address, backend bind.ContractBackend) (*NominaBridgeNative, error) {
	contract, err := bindNominaBridgeNative(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &NominaBridgeNative{NominaBridgeNativeCaller: NominaBridgeNativeCaller{contract: contract}, NominaBridgeNativeTransactor: NominaBridgeNativeTransactor{contract: contract}, NominaBridgeNativeFilterer: NominaBridgeNativeFilterer{contract: contract}}, nil
}

// NewNominaBridgeNativeCaller creates a new read-only instance of NominaBridgeNative, bound to a specific deployed contract.
func NewNominaBridgeNativeCaller(address common.Address, caller bind.ContractCaller) (*NominaBridgeNativeCaller, error) {
	contract, err := bindNominaBridgeNative(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NominaBridgeNativeCaller{contract: contract}, nil
}

// NewNominaBridgeNativeTransactor creates a new write-only instance of NominaBridgeNative, bound to a specific deployed contract.
func NewNominaBridgeNativeTransactor(address common.Address, transactor bind.ContractTransactor) (*NominaBridgeNativeTransactor, error) {
	contract, err := bindNominaBridgeNative(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NominaBridgeNativeTransactor{contract: contract}, nil
}

// NewNominaBridgeNativeFilterer creates a new log filterer instance of NominaBridgeNative, bound to a specific deployed contract.
func NewNominaBridgeNativeFilterer(address common.Address, filterer bind.ContractFilterer) (*NominaBridgeNativeFilterer, error) {
	contract, err := bindNominaBridgeNative(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NominaBridgeNativeFilterer{contract: contract}, nil
}

// bindNominaBridgeNative binds a generic wrapper to an already deployed contract.
func bindNominaBridgeNative(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := NominaBridgeNativeMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NominaBridgeNative *NominaBridgeNativeRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NominaBridgeNative.Contract.NominaBridgeNativeCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NominaBridgeNative *NominaBridgeNativeRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NominaBridgeNative.Contract.NominaBridgeNativeTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NominaBridgeNative *NominaBridgeNativeRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NominaBridgeNative.Contract.NominaBridgeNativeTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NominaBridgeNative *NominaBridgeNativeCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NominaBridgeNative.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NominaBridgeNative *NominaBridgeNativeTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NominaBridgeNative.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NominaBridgeNative *NominaBridgeNativeTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NominaBridgeNative.Contract.contract.Transact(opts, method, params...)
}

// ACTIONBRIDGE is a free data retrieval call binding the contract method 0x09839a93.
//
// Solidity: function ACTION_BRIDGE() view returns(bytes32)
func (_NominaBridgeNative *NominaBridgeNativeCaller) ACTIONBRIDGE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _NominaBridgeNative.contract.Call(opts, &out, "ACTION_BRIDGE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ACTIONBRIDGE is a free data retrieval call binding the contract method 0x09839a93.
//
// Solidity: function ACTION_BRIDGE() view returns(bytes32)
func (_NominaBridgeNative *NominaBridgeNativeSession) ACTIONBRIDGE() ([32]byte, error) {
	return _NominaBridgeNative.Contract.ACTIONBRIDGE(&_NominaBridgeNative.CallOpts)
}

// ACTIONBRIDGE is a free data retrieval call binding the contract method 0x09839a93.
//
// Solidity: function ACTION_BRIDGE() view returns(bytes32)
func (_NominaBridgeNative *NominaBridgeNativeCallerSession) ACTIONBRIDGE() ([32]byte, error) {
	return _NominaBridgeNative.Contract.ACTIONBRIDGE(&_NominaBridgeNative.CallOpts)
}

// ACTIONWITHDRAW is a free data retrieval call binding the contract method 0x25d70f78.
//
// Solidity: function ACTION_WITHDRAW() view returns(bytes32)
func (_NominaBridgeNative *NominaBridgeNativeCaller) ACTIONWITHDRAW(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _NominaBridgeNative.contract.Call(opts, &out, "ACTION_WITHDRAW")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ACTIONWITHDRAW is a free data retrieval call binding the contract method 0x25d70f78.
//
// Solidity: function ACTION_WITHDRAW() view returns(bytes32)
func (_NominaBridgeNative *NominaBridgeNativeSession) ACTIONWITHDRAW() ([32]byte, error) {
	return _NominaBridgeNative.Contract.ACTIONWITHDRAW(&_NominaBridgeNative.CallOpts)
}

// ACTIONWITHDRAW is a free data retrieval call binding the contract method 0x25d70f78.
//
// Solidity: function ACTION_WITHDRAW() view returns(bytes32)
func (_NominaBridgeNative *NominaBridgeNativeCallerSession) ACTIONWITHDRAW() ([32]byte, error) {
	return _NominaBridgeNative.Contract.ACTIONWITHDRAW(&_NominaBridgeNative.CallOpts)
}

// KeyPauseAll is a free data retrieval call binding the contract method 0xa10ac97a.
//
// Solidity: function KeyPauseAll() view returns(bytes32)
func (_NominaBridgeNative *NominaBridgeNativeCaller) KeyPauseAll(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _NominaBridgeNative.contract.Call(opts, &out, "KeyPauseAll")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// KeyPauseAll is a free data retrieval call binding the contract method 0xa10ac97a.
//
// Solidity: function KeyPauseAll() view returns(bytes32)
func (_NominaBridgeNative *NominaBridgeNativeSession) KeyPauseAll() ([32]byte, error) {
	return _NominaBridgeNative.Contract.KeyPauseAll(&_NominaBridgeNative.CallOpts)
}

// KeyPauseAll is a free data retrieval call binding the contract method 0xa10ac97a.
//
// Solidity: function KeyPauseAll() view returns(bytes32)
func (_NominaBridgeNative *NominaBridgeNativeCallerSession) KeyPauseAll() ([32]byte, error) {
	return _NominaBridgeNative.Contract.KeyPauseAll(&_NominaBridgeNative.CallOpts)
}

// XCALLWITHDRAWGASLIMIT is a free data retrieval call binding the contract method 0x8fdcb4c9.
//
// Solidity: function XCALL_WITHDRAW_GAS_LIMIT() view returns(uint64)
func (_NominaBridgeNative *NominaBridgeNativeCaller) XCALLWITHDRAWGASLIMIT(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _NominaBridgeNative.contract.Call(opts, &out, "XCALL_WITHDRAW_GAS_LIMIT")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// XCALLWITHDRAWGASLIMIT is a free data retrieval call binding the contract method 0x8fdcb4c9.
//
// Solidity: function XCALL_WITHDRAW_GAS_LIMIT() view returns(uint64)
func (_NominaBridgeNative *NominaBridgeNativeSession) XCALLWITHDRAWGASLIMIT() (uint64, error) {
	return _NominaBridgeNative.Contract.XCALLWITHDRAWGASLIMIT(&_NominaBridgeNative.CallOpts)
}

// XCALLWITHDRAWGASLIMIT is a free data retrieval call binding the contract method 0x8fdcb4c9.
//
// Solidity: function XCALL_WITHDRAW_GAS_LIMIT() view returns(uint64)
func (_NominaBridgeNative *NominaBridgeNativeCallerSession) XCALLWITHDRAWGASLIMIT() (uint64, error) {
	return _NominaBridgeNative.Contract.XCALLWITHDRAWGASLIMIT(&_NominaBridgeNative.CallOpts)
}

// BridgeFee is a free data retrieval call binding the contract method 0x3abfe55f.
//
// Solidity: function bridgeFee(address to, uint256 amount) view returns(uint256)
func (_NominaBridgeNative *NominaBridgeNativeCaller) BridgeFee(opts *bind.CallOpts, to common.Address, amount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _NominaBridgeNative.contract.Call(opts, &out, "bridgeFee", to, amount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BridgeFee is a free data retrieval call binding the contract method 0x3abfe55f.
//
// Solidity: function bridgeFee(address to, uint256 amount) view returns(uint256)
func (_NominaBridgeNative *NominaBridgeNativeSession) BridgeFee(to common.Address, amount *big.Int) (*big.Int, error) {
	return _NominaBridgeNative.Contract.BridgeFee(&_NominaBridgeNative.CallOpts, to, amount)
}

// BridgeFee is a free data retrieval call binding the contract method 0x3abfe55f.
//
// Solidity: function bridgeFee(address to, uint256 amount) view returns(uint256)
func (_NominaBridgeNative *NominaBridgeNativeCallerSession) BridgeFee(to common.Address, amount *big.Int) (*big.Int, error) {
	return _NominaBridgeNative.Contract.BridgeFee(&_NominaBridgeNative.CallOpts, to, amount)
}

// Claimable is a free data retrieval call binding the contract method 0x402914f5.
//
// Solidity: function claimable(address ) view returns(uint256)
func (_NominaBridgeNative *NominaBridgeNativeCaller) Claimable(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _NominaBridgeNative.contract.Call(opts, &out, "claimable", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Claimable is a free data retrieval call binding the contract method 0x402914f5.
//
// Solidity: function claimable(address ) view returns(uint256)
func (_NominaBridgeNative *NominaBridgeNativeSession) Claimable(arg0 common.Address) (*big.Int, error) {
	return _NominaBridgeNative.Contract.Claimable(&_NominaBridgeNative.CallOpts, arg0)
}

// Claimable is a free data retrieval call binding the contract method 0x402914f5.
//
// Solidity: function claimable(address ) view returns(uint256)
func (_NominaBridgeNative *NominaBridgeNativeCallerSession) Claimable(arg0 common.Address) (*big.Int, error) {
	return _NominaBridgeNative.Contract.Claimable(&_NominaBridgeNative.CallOpts, arg0)
}

// IsPaused is a free data retrieval call binding the contract method 0x241b71bb.
//
// Solidity: function isPaused(bytes32 action) view returns(bool)
func (_NominaBridgeNative *NominaBridgeNativeCaller) IsPaused(opts *bind.CallOpts, action [32]byte) (bool, error) {
	var out []interface{}
	err := _NominaBridgeNative.contract.Call(opts, &out, "isPaused", action)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPaused is a free data retrieval call binding the contract method 0x241b71bb.
//
// Solidity: function isPaused(bytes32 action) view returns(bool)
func (_NominaBridgeNative *NominaBridgeNativeSession) IsPaused(action [32]byte) (bool, error) {
	return _NominaBridgeNative.Contract.IsPaused(&_NominaBridgeNative.CallOpts, action)
}

// IsPaused is a free data retrieval call binding the contract method 0x241b71bb.
//
// Solidity: function isPaused(bytes32 action) view returns(bool)
func (_NominaBridgeNative *NominaBridgeNativeCallerSession) IsPaused(action [32]byte) (bool, error) {
	return _NominaBridgeNative.Contract.IsPaused(&_NominaBridgeNative.CallOpts, action)
}

// L1Bridge is a free data retrieval call binding the contract method 0x969b53da.
//
// Solidity: function l1Bridge() view returns(address)
func (_NominaBridgeNative *NominaBridgeNativeCaller) L1Bridge(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NominaBridgeNative.contract.Call(opts, &out, "l1Bridge")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// L1Bridge is a free data retrieval call binding the contract method 0x969b53da.
//
// Solidity: function l1Bridge() view returns(address)
func (_NominaBridgeNative *NominaBridgeNativeSession) L1Bridge() (common.Address, error) {
	return _NominaBridgeNative.Contract.L1Bridge(&_NominaBridgeNative.CallOpts)
}

// L1Bridge is a free data retrieval call binding the contract method 0x969b53da.
//
// Solidity: function l1Bridge() view returns(address)
func (_NominaBridgeNative *NominaBridgeNativeCallerSession) L1Bridge() (common.Address, error) {
	return _NominaBridgeNative.Contract.L1Bridge(&_NominaBridgeNative.CallOpts)
}

// L1ChainId is a free data retrieval call binding the contract method 0x12622e5b.
//
// Solidity: function l1ChainId() view returns(uint64)
func (_NominaBridgeNative *NominaBridgeNativeCaller) L1ChainId(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _NominaBridgeNative.contract.Call(opts, &out, "l1ChainId")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// L1ChainId is a free data retrieval call binding the contract method 0x12622e5b.
//
// Solidity: function l1ChainId() view returns(uint64)
func (_NominaBridgeNative *NominaBridgeNativeSession) L1ChainId() (uint64, error) {
	return _NominaBridgeNative.Contract.L1ChainId(&_NominaBridgeNative.CallOpts)
}

// L1ChainId is a free data retrieval call binding the contract method 0x12622e5b.
//
// Solidity: function l1ChainId() view returns(uint64)
func (_NominaBridgeNative *NominaBridgeNativeCallerSession) L1ChainId() (uint64, error) {
	return _NominaBridgeNative.Contract.L1ChainId(&_NominaBridgeNative.CallOpts)
}

// L1Deposits is a free data retrieval call binding the contract method 0x32c8bb77.
//
// Solidity: function l1Deposits() view returns(uint256)
func (_NominaBridgeNative *NominaBridgeNativeCaller) L1Deposits(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _NominaBridgeNative.contract.Call(opts, &out, "l1Deposits")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// L1Deposits is a free data retrieval call binding the contract method 0x32c8bb77.
//
// Solidity: function l1Deposits() view returns(uint256)
func (_NominaBridgeNative *NominaBridgeNativeSession) L1Deposits() (*big.Int, error) {
	return _NominaBridgeNative.Contract.L1Deposits(&_NominaBridgeNative.CallOpts)
}

// L1Deposits is a free data retrieval call binding the contract method 0x32c8bb77.
//
// Solidity: function l1Deposits() view returns(uint256)
func (_NominaBridgeNative *NominaBridgeNativeCallerSession) L1Deposits() (*big.Int, error) {
	return _NominaBridgeNative.Contract.L1Deposits(&_NominaBridgeNative.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_NominaBridgeNative *NominaBridgeNativeCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NominaBridgeNative.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_NominaBridgeNative *NominaBridgeNativeSession) Owner() (common.Address, error) {
	return _NominaBridgeNative.Contract.Owner(&_NominaBridgeNative.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_NominaBridgeNative *NominaBridgeNativeCallerSession) Owner() (common.Address, error) {
	return _NominaBridgeNative.Contract.Owner(&_NominaBridgeNative.CallOpts)
}

// Portal is a free data retrieval call binding the contract method 0x6425666b.
//
// Solidity: function portal() view returns(address)
func (_NominaBridgeNative *NominaBridgeNativeCaller) Portal(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NominaBridgeNative.contract.Call(opts, &out, "portal")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Portal is a free data retrieval call binding the contract method 0x6425666b.
//
// Solidity: function portal() view returns(address)
func (_NominaBridgeNative *NominaBridgeNativeSession) Portal() (common.Address, error) {
	return _NominaBridgeNative.Contract.Portal(&_NominaBridgeNative.CallOpts)
}

// Portal is a free data retrieval call binding the contract method 0x6425666b.
//
// Solidity: function portal() view returns(address)
func (_NominaBridgeNative *NominaBridgeNativeCallerSession) Portal() (common.Address, error) {
	return _NominaBridgeNative.Contract.Portal(&_NominaBridgeNative.CallOpts)
}

// Bridge is a paid mutator transaction binding the contract method 0xc3de453d.
//
// Solidity: function bridge(address to, uint256 amount) payable returns()
func (_NominaBridgeNative *NominaBridgeNativeTransactor) Bridge(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _NominaBridgeNative.contract.Transact(opts, "bridge", to, amount)
}

// Bridge is a paid mutator transaction binding the contract method 0xc3de453d.
//
// Solidity: function bridge(address to, uint256 amount) payable returns()
func (_NominaBridgeNative *NominaBridgeNativeSession) Bridge(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _NominaBridgeNative.Contract.Bridge(&_NominaBridgeNative.TransactOpts, to, amount)
}

// Bridge is a paid mutator transaction binding the contract method 0xc3de453d.
//
// Solidity: function bridge(address to, uint256 amount) payable returns()
func (_NominaBridgeNative *NominaBridgeNativeTransactorSession) Bridge(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _NominaBridgeNative.Contract.Bridge(&_NominaBridgeNative.TransactOpts, to, amount)
}

// Claim is a paid mutator transaction binding the contract method 0x1e83409a.
//
// Solidity: function claim(address to) returns()
func (_NominaBridgeNative *NominaBridgeNativeTransactor) Claim(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _NominaBridgeNative.contract.Transact(opts, "claim", to)
}

// Claim is a paid mutator transaction binding the contract method 0x1e83409a.
//
// Solidity: function claim(address to) returns()
func (_NominaBridgeNative *NominaBridgeNativeSession) Claim(to common.Address) (*types.Transaction, error) {
	return _NominaBridgeNative.Contract.Claim(&_NominaBridgeNative.TransactOpts, to)
}

// Claim is a paid mutator transaction binding the contract method 0x1e83409a.
//
// Solidity: function claim(address to) returns()
func (_NominaBridgeNative *NominaBridgeNativeTransactorSession) Claim(to common.Address) (*types.Transaction, error) {
	return _NominaBridgeNative.Contract.Claim(&_NominaBridgeNative.TransactOpts, to)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address owner_) returns()
func (_NominaBridgeNative *NominaBridgeNativeTransactor) Initialize(opts *bind.TransactOpts, owner_ common.Address) (*types.Transaction, error) {
	return _NominaBridgeNative.contract.Transact(opts, "initialize", owner_)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address owner_) returns()
func (_NominaBridgeNative *NominaBridgeNativeSession) Initialize(owner_ common.Address) (*types.Transaction, error) {
	return _NominaBridgeNative.Contract.Initialize(&_NominaBridgeNative.TransactOpts, owner_)
}

// Initialize is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address owner_) returns()
func (_NominaBridgeNative *NominaBridgeNativeTransactorSession) Initialize(owner_ common.Address) (*types.Transaction, error) {
	return _NominaBridgeNative.Contract.Initialize(&_NominaBridgeNative.TransactOpts, owner_)
}

// InitializeV2 is a paid mutator transaction binding the contract method 0x5cd8a76b.
//
// Solidity: function initializeV2() returns()
func (_NominaBridgeNative *NominaBridgeNativeTransactor) InitializeV2(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NominaBridgeNative.contract.Transact(opts, "initializeV2")
}

// InitializeV2 is a paid mutator transaction binding the contract method 0x5cd8a76b.
//
// Solidity: function initializeV2() returns()
func (_NominaBridgeNative *NominaBridgeNativeSession) InitializeV2() (*types.Transaction, error) {
	return _NominaBridgeNative.Contract.InitializeV2(&_NominaBridgeNative.TransactOpts)
}

// InitializeV2 is a paid mutator transaction binding the contract method 0x5cd8a76b.
//
// Solidity: function initializeV2() returns()
func (_NominaBridgeNative *NominaBridgeNativeTransactorSession) InitializeV2() (*types.Transaction, error) {
	return _NominaBridgeNative.Contract.InitializeV2(&_NominaBridgeNative.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_NominaBridgeNative *NominaBridgeNativeTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NominaBridgeNative.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_NominaBridgeNative *NominaBridgeNativeSession) Pause() (*types.Transaction, error) {
	return _NominaBridgeNative.Contract.Pause(&_NominaBridgeNative.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_NominaBridgeNative *NominaBridgeNativeTransactorSession) Pause() (*types.Transaction, error) {
	return _NominaBridgeNative.Contract.Pause(&_NominaBridgeNative.TransactOpts)
}

// Pause0 is a paid mutator transaction binding the contract method 0xed56531a.
//
// Solidity: function pause(bytes32 action) returns()
func (_NominaBridgeNative *NominaBridgeNativeTransactor) Pause0(opts *bind.TransactOpts, action [32]byte) (*types.Transaction, error) {
	return _NominaBridgeNative.contract.Transact(opts, "pause0", action)
}

// Pause0 is a paid mutator transaction binding the contract method 0xed56531a.
//
// Solidity: function pause(bytes32 action) returns()
func (_NominaBridgeNative *NominaBridgeNativeSession) Pause0(action [32]byte) (*types.Transaction, error) {
	return _NominaBridgeNative.Contract.Pause0(&_NominaBridgeNative.TransactOpts, action)
}

// Pause0 is a paid mutator transaction binding the contract method 0xed56531a.
//
// Solidity: function pause(bytes32 action) returns()
func (_NominaBridgeNative *NominaBridgeNativeTransactorSession) Pause0(action [32]byte) (*types.Transaction, error) {
	return _NominaBridgeNative.Contract.Pause0(&_NominaBridgeNative.TransactOpts, action)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_NominaBridgeNative *NominaBridgeNativeTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NominaBridgeNative.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_NominaBridgeNative *NominaBridgeNativeSession) RenounceOwnership() (*types.Transaction, error) {
	return _NominaBridgeNative.Contract.RenounceOwnership(&_NominaBridgeNative.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_NominaBridgeNative *NominaBridgeNativeTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _NominaBridgeNative.Contract.RenounceOwnership(&_NominaBridgeNative.TransactOpts)
}

// Setup is a paid mutator transaction binding the contract method 0x499e85cd.
//
// Solidity: function setup(uint64 l1ChainId_, address portal_, address l1Bridge_, uint256 l1Deposits_) returns()
func (_NominaBridgeNative *NominaBridgeNativeTransactor) Setup(opts *bind.TransactOpts, l1ChainId_ uint64, portal_ common.Address, l1Bridge_ common.Address, l1Deposits_ *big.Int) (*types.Transaction, error) {
	return _NominaBridgeNative.contract.Transact(opts, "setup", l1ChainId_, portal_, l1Bridge_, l1Deposits_)
}

// Setup is a paid mutator transaction binding the contract method 0x499e85cd.
//
// Solidity: function setup(uint64 l1ChainId_, address portal_, address l1Bridge_, uint256 l1Deposits_) returns()
func (_NominaBridgeNative *NominaBridgeNativeSession) Setup(l1ChainId_ uint64, portal_ common.Address, l1Bridge_ common.Address, l1Deposits_ *big.Int) (*types.Transaction, error) {
	return _NominaBridgeNative.Contract.Setup(&_NominaBridgeNative.TransactOpts, l1ChainId_, portal_, l1Bridge_, l1Deposits_)
}

// Setup is a paid mutator transaction binding the contract method 0x499e85cd.
//
// Solidity: function setup(uint64 l1ChainId_, address portal_, address l1Bridge_, uint256 l1Deposits_) returns()
func (_NominaBridgeNative *NominaBridgeNativeTransactorSession) Setup(l1ChainId_ uint64, portal_ common.Address, l1Bridge_ common.Address, l1Deposits_ *big.Int) (*types.Transaction, error) {
	return _NominaBridgeNative.Contract.Setup(&_NominaBridgeNative.TransactOpts, l1ChainId_, portal_, l1Bridge_, l1Deposits_)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_NominaBridgeNative *NominaBridgeNativeTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _NominaBridgeNative.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_NominaBridgeNative *NominaBridgeNativeSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _NominaBridgeNative.Contract.TransferOwnership(&_NominaBridgeNative.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_NominaBridgeNative *NominaBridgeNativeTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _NominaBridgeNative.Contract.TransferOwnership(&_NominaBridgeNative.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x2f4dae9f.
//
// Solidity: function unpause(bytes32 action) returns()
func (_NominaBridgeNative *NominaBridgeNativeTransactor) Unpause(opts *bind.TransactOpts, action [32]byte) (*types.Transaction, error) {
	return _NominaBridgeNative.contract.Transact(opts, "unpause", action)
}

// Unpause is a paid mutator transaction binding the contract method 0x2f4dae9f.
//
// Solidity: function unpause(bytes32 action) returns()
func (_NominaBridgeNative *NominaBridgeNativeSession) Unpause(action [32]byte) (*types.Transaction, error) {
	return _NominaBridgeNative.Contract.Unpause(&_NominaBridgeNative.TransactOpts, action)
}

// Unpause is a paid mutator transaction binding the contract method 0x2f4dae9f.
//
// Solidity: function unpause(bytes32 action) returns()
func (_NominaBridgeNative *NominaBridgeNativeTransactorSession) Unpause(action [32]byte) (*types.Transaction, error) {
	return _NominaBridgeNative.Contract.Unpause(&_NominaBridgeNative.TransactOpts, action)
}

// Unpause0 is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_NominaBridgeNative *NominaBridgeNativeTransactor) Unpause0(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NominaBridgeNative.contract.Transact(opts, "unpause0")
}

// Unpause0 is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_NominaBridgeNative *NominaBridgeNativeSession) Unpause0() (*types.Transaction, error) {
	return _NominaBridgeNative.Contract.Unpause0(&_NominaBridgeNative.TransactOpts)
}

// Unpause0 is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_NominaBridgeNative *NominaBridgeNativeTransactorSession) Unpause0() (*types.Transaction, error) {
	return _NominaBridgeNative.Contract.Unpause0(&_NominaBridgeNative.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0xd9caed12.
//
// Solidity: function withdraw(address payor, address to, uint256 amount) returns()
func (_NominaBridgeNative *NominaBridgeNativeTransactor) Withdraw(opts *bind.TransactOpts, payor common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _NominaBridgeNative.contract.Transact(opts, "withdraw", payor, to, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xd9caed12.
//
// Solidity: function withdraw(address payor, address to, uint256 amount) returns()
func (_NominaBridgeNative *NominaBridgeNativeSession) Withdraw(payor common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _NominaBridgeNative.Contract.Withdraw(&_NominaBridgeNative.TransactOpts, payor, to, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xd9caed12.
//
// Solidity: function withdraw(address payor, address to, uint256 amount) returns()
func (_NominaBridgeNative *NominaBridgeNativeTransactorSession) Withdraw(payor common.Address, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _NominaBridgeNative.Contract.Withdraw(&_NominaBridgeNative.TransactOpts, payor, to, amount)
}

// NominaBridgeNativeBridgeIterator is returned from FilterBridge and is used to iterate over the raw logs and unpacked data for Bridge events raised by the NominaBridgeNative contract.
type NominaBridgeNativeBridgeIterator struct {
	Event *NominaBridgeNativeBridge // Event containing the contract specifics and raw log

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
func (it *NominaBridgeNativeBridgeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NominaBridgeNativeBridge)
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
		it.Event = new(NominaBridgeNativeBridge)
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
func (it *NominaBridgeNativeBridgeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NominaBridgeNativeBridgeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NominaBridgeNativeBridge represents a Bridge event raised by the NominaBridgeNative contract.
type NominaBridgeNativeBridge struct {
	Payor  common.Address
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBridge is a free log retrieval operation binding the contract event 0x59bc8a913d49a9626dd6ba5def7fcf12804061c1bb9b8b6db077e1a12cb4b422.
//
// Solidity: event Bridge(address indexed payor, address indexed to, uint256 amount)
func (_NominaBridgeNative *NominaBridgeNativeFilterer) FilterBridge(opts *bind.FilterOpts, payor []common.Address, to []common.Address) (*NominaBridgeNativeBridgeIterator, error) {

	var payorRule []interface{}
	for _, payorItem := range payor {
		payorRule = append(payorRule, payorItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _NominaBridgeNative.contract.FilterLogs(opts, "Bridge", payorRule, toRule)
	if err != nil {
		return nil, err
	}
	return &NominaBridgeNativeBridgeIterator{contract: _NominaBridgeNative.contract, event: "Bridge", logs: logs, sub: sub}, nil
}

// WatchBridge is a free log subscription operation binding the contract event 0x59bc8a913d49a9626dd6ba5def7fcf12804061c1bb9b8b6db077e1a12cb4b422.
//
// Solidity: event Bridge(address indexed payor, address indexed to, uint256 amount)
func (_NominaBridgeNative *NominaBridgeNativeFilterer) WatchBridge(opts *bind.WatchOpts, sink chan<- *NominaBridgeNativeBridge, payor []common.Address, to []common.Address) (event.Subscription, error) {

	var payorRule []interface{}
	for _, payorItem := range payor {
		payorRule = append(payorRule, payorItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _NominaBridgeNative.contract.WatchLogs(opts, "Bridge", payorRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NominaBridgeNativeBridge)
				if err := _NominaBridgeNative.contract.UnpackLog(event, "Bridge", log); err != nil {
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

// ParseBridge is a log parse operation binding the contract event 0x59bc8a913d49a9626dd6ba5def7fcf12804061c1bb9b8b6db077e1a12cb4b422.
//
// Solidity: event Bridge(address indexed payor, address indexed to, uint256 amount)
func (_NominaBridgeNative *NominaBridgeNativeFilterer) ParseBridge(log types.Log) (*NominaBridgeNativeBridge, error) {
	event := new(NominaBridgeNativeBridge)
	if err := _NominaBridgeNative.contract.UnpackLog(event, "Bridge", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NominaBridgeNativeClaimedIterator is returned from FilterClaimed and is used to iterate over the raw logs and unpacked data for Claimed events raised by the NominaBridgeNative contract.
type NominaBridgeNativeClaimedIterator struct {
	Event *NominaBridgeNativeClaimed // Event containing the contract specifics and raw log

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
func (it *NominaBridgeNativeClaimedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NominaBridgeNativeClaimed)
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
		it.Event = new(NominaBridgeNativeClaimed)
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
func (it *NominaBridgeNativeClaimedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NominaBridgeNativeClaimedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NominaBridgeNativeClaimed represents a Claimed event raised by the NominaBridgeNative contract.
type NominaBridgeNativeClaimed struct {
	Claimant common.Address
	To       common.Address
	Amount   *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterClaimed is a free log retrieval operation binding the contract event 0xf7a40077ff7a04c7e61f6f26fb13774259ddf1b6bce9ecf26a8276cdd3992683.
//
// Solidity: event Claimed(address indexed claimant, address indexed to, uint256 amount)
func (_NominaBridgeNative *NominaBridgeNativeFilterer) FilterClaimed(opts *bind.FilterOpts, claimant []common.Address, to []common.Address) (*NominaBridgeNativeClaimedIterator, error) {

	var claimantRule []interface{}
	for _, claimantItem := range claimant {
		claimantRule = append(claimantRule, claimantItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _NominaBridgeNative.contract.FilterLogs(opts, "Claimed", claimantRule, toRule)
	if err != nil {
		return nil, err
	}
	return &NominaBridgeNativeClaimedIterator{contract: _NominaBridgeNative.contract, event: "Claimed", logs: logs, sub: sub}, nil
}

// WatchClaimed is a free log subscription operation binding the contract event 0xf7a40077ff7a04c7e61f6f26fb13774259ddf1b6bce9ecf26a8276cdd3992683.
//
// Solidity: event Claimed(address indexed claimant, address indexed to, uint256 amount)
func (_NominaBridgeNative *NominaBridgeNativeFilterer) WatchClaimed(opts *bind.WatchOpts, sink chan<- *NominaBridgeNativeClaimed, claimant []common.Address, to []common.Address) (event.Subscription, error) {

	var claimantRule []interface{}
	for _, claimantItem := range claimant {
		claimantRule = append(claimantRule, claimantItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _NominaBridgeNative.contract.WatchLogs(opts, "Claimed", claimantRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NominaBridgeNativeClaimed)
				if err := _NominaBridgeNative.contract.UnpackLog(event, "Claimed", log); err != nil {
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

// ParseClaimed is a log parse operation binding the contract event 0xf7a40077ff7a04c7e61f6f26fb13774259ddf1b6bce9ecf26a8276cdd3992683.
//
// Solidity: event Claimed(address indexed claimant, address indexed to, uint256 amount)
func (_NominaBridgeNative *NominaBridgeNativeFilterer) ParseClaimed(log types.Log) (*NominaBridgeNativeClaimed, error) {
	event := new(NominaBridgeNativeClaimed)
	if err := _NominaBridgeNative.contract.UnpackLog(event, "Claimed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NominaBridgeNativeInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the NominaBridgeNative contract.
type NominaBridgeNativeInitializedIterator struct {
	Event *NominaBridgeNativeInitialized // Event containing the contract specifics and raw log

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
func (it *NominaBridgeNativeInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NominaBridgeNativeInitialized)
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
		it.Event = new(NominaBridgeNativeInitialized)
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
func (it *NominaBridgeNativeInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NominaBridgeNativeInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NominaBridgeNativeInitialized represents a Initialized event raised by the NominaBridgeNative contract.
type NominaBridgeNativeInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_NominaBridgeNative *NominaBridgeNativeFilterer) FilterInitialized(opts *bind.FilterOpts) (*NominaBridgeNativeInitializedIterator, error) {

	logs, sub, err := _NominaBridgeNative.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &NominaBridgeNativeInitializedIterator{contract: _NominaBridgeNative.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_NominaBridgeNative *NominaBridgeNativeFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *NominaBridgeNativeInitialized) (event.Subscription, error) {

	logs, sub, err := _NominaBridgeNative.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NominaBridgeNativeInitialized)
				if err := _NominaBridgeNative.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_NominaBridgeNative *NominaBridgeNativeFilterer) ParseInitialized(log types.Log) (*NominaBridgeNativeInitialized, error) {
	event := new(NominaBridgeNativeInitialized)
	if err := _NominaBridgeNative.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NominaBridgeNativeOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the NominaBridgeNative contract.
type NominaBridgeNativeOwnershipTransferredIterator struct {
	Event *NominaBridgeNativeOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *NominaBridgeNativeOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NominaBridgeNativeOwnershipTransferred)
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
		it.Event = new(NominaBridgeNativeOwnershipTransferred)
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
func (it *NominaBridgeNativeOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NominaBridgeNativeOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NominaBridgeNativeOwnershipTransferred represents a OwnershipTransferred event raised by the NominaBridgeNative contract.
type NominaBridgeNativeOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_NominaBridgeNative *NominaBridgeNativeFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*NominaBridgeNativeOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _NominaBridgeNative.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &NominaBridgeNativeOwnershipTransferredIterator{contract: _NominaBridgeNative.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_NominaBridgeNative *NominaBridgeNativeFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *NominaBridgeNativeOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _NominaBridgeNative.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NominaBridgeNativeOwnershipTransferred)
				if err := _NominaBridgeNative.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_NominaBridgeNative *NominaBridgeNativeFilterer) ParseOwnershipTransferred(log types.Log) (*NominaBridgeNativeOwnershipTransferred, error) {
	event := new(NominaBridgeNativeOwnershipTransferred)
	if err := _NominaBridgeNative.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NominaBridgeNativePausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the NominaBridgeNative contract.
type NominaBridgeNativePausedIterator struct {
	Event *NominaBridgeNativePaused // Event containing the contract specifics and raw log

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
func (it *NominaBridgeNativePausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NominaBridgeNativePaused)
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
		it.Event = new(NominaBridgeNativePaused)
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
func (it *NominaBridgeNativePausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NominaBridgeNativePausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NominaBridgeNativePaused represents a Paused event raised by the NominaBridgeNative contract.
type NominaBridgeNativePaused struct {
	Key [32]byte
	Raw types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x0cb09dc71d57eeec2046f6854976717e4874a3cf2d6ddeddde337e5b6de6ba31.
//
// Solidity: event Paused(bytes32 indexed key)
func (_NominaBridgeNative *NominaBridgeNativeFilterer) FilterPaused(opts *bind.FilterOpts, key [][32]byte) (*NominaBridgeNativePausedIterator, error) {

	var keyRule []interface{}
	for _, keyItem := range key {
		keyRule = append(keyRule, keyItem)
	}

	logs, sub, err := _NominaBridgeNative.contract.FilterLogs(opts, "Paused", keyRule)
	if err != nil {
		return nil, err
	}
	return &NominaBridgeNativePausedIterator{contract: _NominaBridgeNative.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x0cb09dc71d57eeec2046f6854976717e4874a3cf2d6ddeddde337e5b6de6ba31.
//
// Solidity: event Paused(bytes32 indexed key)
func (_NominaBridgeNative *NominaBridgeNativeFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *NominaBridgeNativePaused, key [][32]byte) (event.Subscription, error) {

	var keyRule []interface{}
	for _, keyItem := range key {
		keyRule = append(keyRule, keyItem)
	}

	logs, sub, err := _NominaBridgeNative.contract.WatchLogs(opts, "Paused", keyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NominaBridgeNativePaused)
				if err := _NominaBridgeNative.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x0cb09dc71d57eeec2046f6854976717e4874a3cf2d6ddeddde337e5b6de6ba31.
//
// Solidity: event Paused(bytes32 indexed key)
func (_NominaBridgeNative *NominaBridgeNativeFilterer) ParsePaused(log types.Log) (*NominaBridgeNativePaused, error) {
	event := new(NominaBridgeNativePaused)
	if err := _NominaBridgeNative.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NominaBridgeNativeSetupIterator is returned from FilterSetup and is used to iterate over the raw logs and unpacked data for Setup events raised by the NominaBridgeNative contract.
type NominaBridgeNativeSetupIterator struct {
	Event *NominaBridgeNativeSetup // Event containing the contract specifics and raw log

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
func (it *NominaBridgeNativeSetupIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NominaBridgeNativeSetup)
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
		it.Event = new(NominaBridgeNativeSetup)
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
func (it *NominaBridgeNativeSetupIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NominaBridgeNativeSetupIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NominaBridgeNativeSetup represents a Setup event raised by the NominaBridgeNative contract.
type NominaBridgeNativeSetup struct {
	L1ChainId  uint64
	Portal     common.Address
	L1Bridge   common.Address
	L1Deposits *big.Int
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterSetup is a free log retrieval operation binding the contract event 0xc12d4a2db17193df88185c2dc087fa9536c12710f1381b359b80553e5d9a1293.
//
// Solidity: event Setup(uint64 l1ChainId, address portal, address l1Bridge, uint256 l1Deposits)
func (_NominaBridgeNative *NominaBridgeNativeFilterer) FilterSetup(opts *bind.FilterOpts) (*NominaBridgeNativeSetupIterator, error) {

	logs, sub, err := _NominaBridgeNative.contract.FilterLogs(opts, "Setup")
	if err != nil {
		return nil, err
	}
	return &NominaBridgeNativeSetupIterator{contract: _NominaBridgeNative.contract, event: "Setup", logs: logs, sub: sub}, nil
}

// WatchSetup is a free log subscription operation binding the contract event 0xc12d4a2db17193df88185c2dc087fa9536c12710f1381b359b80553e5d9a1293.
//
// Solidity: event Setup(uint64 l1ChainId, address portal, address l1Bridge, uint256 l1Deposits)
func (_NominaBridgeNative *NominaBridgeNativeFilterer) WatchSetup(opts *bind.WatchOpts, sink chan<- *NominaBridgeNativeSetup) (event.Subscription, error) {

	logs, sub, err := _NominaBridgeNative.contract.WatchLogs(opts, "Setup")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NominaBridgeNativeSetup)
				if err := _NominaBridgeNative.contract.UnpackLog(event, "Setup", log); err != nil {
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

// ParseSetup is a log parse operation binding the contract event 0xc12d4a2db17193df88185c2dc087fa9536c12710f1381b359b80553e5d9a1293.
//
// Solidity: event Setup(uint64 l1ChainId, address portal, address l1Bridge, uint256 l1Deposits)
func (_NominaBridgeNative *NominaBridgeNativeFilterer) ParseSetup(log types.Log) (*NominaBridgeNativeSetup, error) {
	event := new(NominaBridgeNativeSetup)
	if err := _NominaBridgeNative.contract.UnpackLog(event, "Setup", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NominaBridgeNativeUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the NominaBridgeNative contract.
type NominaBridgeNativeUnpausedIterator struct {
	Event *NominaBridgeNativeUnpaused // Event containing the contract specifics and raw log

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
func (it *NominaBridgeNativeUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NominaBridgeNativeUnpaused)
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
		it.Event = new(NominaBridgeNativeUnpaused)
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
func (it *NominaBridgeNativeUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NominaBridgeNativeUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NominaBridgeNativeUnpaused represents a Unpaused event raised by the NominaBridgeNative contract.
type NominaBridgeNativeUnpaused struct {
	Key [32]byte
	Raw types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0xd05bfc2250abb0f8fd265a54c53a24359c5484af63cad2e4ce87c78ab751395a.
//
// Solidity: event Unpaused(bytes32 indexed key)
func (_NominaBridgeNative *NominaBridgeNativeFilterer) FilterUnpaused(opts *bind.FilterOpts, key [][32]byte) (*NominaBridgeNativeUnpausedIterator, error) {

	var keyRule []interface{}
	for _, keyItem := range key {
		keyRule = append(keyRule, keyItem)
	}

	logs, sub, err := _NominaBridgeNative.contract.FilterLogs(opts, "Unpaused", keyRule)
	if err != nil {
		return nil, err
	}
	return &NominaBridgeNativeUnpausedIterator{contract: _NominaBridgeNative.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0xd05bfc2250abb0f8fd265a54c53a24359c5484af63cad2e4ce87c78ab751395a.
//
// Solidity: event Unpaused(bytes32 indexed key)
func (_NominaBridgeNative *NominaBridgeNativeFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *NominaBridgeNativeUnpaused, key [][32]byte) (event.Subscription, error) {

	var keyRule []interface{}
	for _, keyItem := range key {
		keyRule = append(keyRule, keyItem)
	}

	logs, sub, err := _NominaBridgeNative.contract.WatchLogs(opts, "Unpaused", keyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NominaBridgeNativeUnpaused)
				if err := _NominaBridgeNative.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0xd05bfc2250abb0f8fd265a54c53a24359c5484af63cad2e4ce87c78ab751395a.
//
// Solidity: event Unpaused(bytes32 indexed key)
func (_NominaBridgeNative *NominaBridgeNativeFilterer) ParseUnpaused(log types.Log) (*NominaBridgeNativeUnpaused, error) {
	event := new(NominaBridgeNativeUnpaused)
	if err := _NominaBridgeNative.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NominaBridgeNativeWithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the NominaBridgeNative contract.
type NominaBridgeNativeWithdrawIterator struct {
	Event *NominaBridgeNativeWithdraw // Event containing the contract specifics and raw log

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
func (it *NominaBridgeNativeWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NominaBridgeNativeWithdraw)
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
		it.Event = new(NominaBridgeNativeWithdraw)
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
func (it *NominaBridgeNativeWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NominaBridgeNativeWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NominaBridgeNativeWithdraw represents a Withdraw event raised by the NominaBridgeNative contract.
type NominaBridgeNativeWithdraw struct {
	Payor   common.Address
	To      common.Address
	Amount  *big.Int
	Success bool
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0x2b664ab52fe561d3ace376046aea39744dd736ec1f67d89d504ffd2192825f61.
//
// Solidity: event Withdraw(address indexed payor, address indexed to, uint256 amount, bool success)
func (_NominaBridgeNative *NominaBridgeNativeFilterer) FilterWithdraw(opts *bind.FilterOpts, payor []common.Address, to []common.Address) (*NominaBridgeNativeWithdrawIterator, error) {

	var payorRule []interface{}
	for _, payorItem := range payor {
		payorRule = append(payorRule, payorItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _NominaBridgeNative.contract.FilterLogs(opts, "Withdraw", payorRule, toRule)
	if err != nil {
		return nil, err
	}
	return &NominaBridgeNativeWithdrawIterator{contract: _NominaBridgeNative.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0x2b664ab52fe561d3ace376046aea39744dd736ec1f67d89d504ffd2192825f61.
//
// Solidity: event Withdraw(address indexed payor, address indexed to, uint256 amount, bool success)
func (_NominaBridgeNative *NominaBridgeNativeFilterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *NominaBridgeNativeWithdraw, payor []common.Address, to []common.Address) (event.Subscription, error) {

	var payorRule []interface{}
	for _, payorItem := range payor {
		payorRule = append(payorRule, payorItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _NominaBridgeNative.contract.WatchLogs(opts, "Withdraw", payorRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NominaBridgeNativeWithdraw)
				if err := _NominaBridgeNative.contract.UnpackLog(event, "Withdraw", log); err != nil {
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

// ParseWithdraw is a log parse operation binding the contract event 0x2b664ab52fe561d3ace376046aea39744dd736ec1f67d89d504ffd2192825f61.
//
// Solidity: event Withdraw(address indexed payor, address indexed to, uint256 amount, bool success)
func (_NominaBridgeNative *NominaBridgeNativeFilterer) ParseWithdraw(log types.Log) (*NominaBridgeNativeWithdraw, error) {
	event := new(NominaBridgeNativeWithdraw)
	if err := _NominaBridgeNative.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
