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

// NominaBridgeL1MetaData contains all meta data concerning the NominaBridgeL1 contract.
var NominaBridgeL1MetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"omni_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"nomina_\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"ACTION_BRIDGE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"ACTION_WITHDRAW\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"KeyPauseAll\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"XCALL_WITHDRAW_GAS_LIMIT\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"bridge\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"bridgeFee\",\"inputs\":[{\"name\":\"payor\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"owner_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"portal_\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"initializeV2\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isPaused\",\"inputs\":[{\"name\":\"action\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"nomina\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIERC20\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"omni\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIERC20\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[{\"name\":\"action\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"portal\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractINominaPortal\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[{\"name\":\"action\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdraw\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Bridge\",\"inputs\":[{\"name\":\"payor\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"key\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"key\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Withdraw\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnableInvalidOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnableUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]}]",
	Bin: "0x60c060405234801561000f575f5ffd5b5060405161191238038061191283398101604081905261002e9161011d565b6001600160a01b03808316608052811660a052610049610050565b505061014e565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000900460ff16156100a05760405163f92ee8a960e01b815260040160405180910390fd5b80546001600160401b03908116146100ff5780546001600160401b0319166001600160401b0390811782556040519081527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b50565b80516001600160a01b0381168114610118575f5ffd5b919050565b5f5f6040838503121561012e575f5ffd5b61013783610102565b915061014560208401610102565b90509250929050565b60805160a0516117806101925f395f81816102c80152818161075d01528181610cca015261123c01525f818161021801528181610787015261083e01526117805ff3fe60806040526004361061011b575f3560e01c8063646310c81161009d578063a10ac97a11610062578063a10ac97a1461037c578063c3de453d1461039c578063ed56531a146103af578063f2fde38b146103ce578063f3fef3a3146103ed575f5ffd5b8063646310c8146102b7578063715018a6146102ea5780638456cb59146102fe5780638da5cb5b146103125780638fdcb4c91461034e575f5ffd5b806339acf9f1116100e357806339acf9f1146102075780633f4ba83a14610252578063485cc955146102665780635cd8a76b146102855780636425666b14610299575f5ffd5b806309839a931461011f578063241b71bb1461016557806325d70f78146101945780632f4dae9f146101c75780633794999d146101e8575b5f5ffd5b34801561012a575f5ffd5b506101527f0683d1c283a672fc58eb7940a0dba83ea98b96966a9ca1b030dec2c60cea4d1e81565b6040519081526020015b60405180910390f35b348015610170575f5ffd5b5061018461017f3660046114b5565b61040c565b604051901515815260200161015c565b34801561019f575f5ffd5b506101527f855511cc3694f64379908437d6d64458dc76d02482052bfb8a5b33a72c054c7781565b3480156101d2575f5ffd5b506101e66101e13660046114b5565b61041c565b005b3480156101f3575f5ffd5b506101526102023660046114e0565b610430565b348015610212575f5ffd5b5061023a7f000000000000000000000000000000000000000000000000000000000000000081565b6040516001600160a01b03909116815260200161015c565b34801561025d575f5ffd5b506101e661055e565b348015610271575f5ffd5b506101e661028036600461151e565b610570565b348015610290575f5ffd5b506101e66106dd565b3480156102a4575f5ffd5b505f5461023a906001600160a01b031681565b3480156102c2575f5ffd5b5061023a7f000000000000000000000000000000000000000000000000000000000000000081565b3480156102f5575f5ffd5b506101e661094c565b348015610309575f5ffd5b506101e661095d565b34801561031d575f5ffd5b507f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c199300546001600160a01b031661023a565b348015610359575f5ffd5b506103646201388081565b6040516001600160401b03909116815260200161015c565b348015610387575f5ffd5b506101525f51602061172b5f395f51905f5281565b6101e66103aa366004611555565b61096d565b3480156103ba575f5ffd5b506101e66103c93660046114b5565b6109eb565b3480156103d9575f5ffd5b506101e66103e836600461157f565b6109fc565b3480156103f8575f5ffd5b506101e6610407366004611555565b610a36565b5f61041682610d7e565b92915050565b610424610dfa565b61042d81610e55565b50565b5f805460408051630998721160e31b815290516001600160a01b0390921691638dd9523c918391634cc39088916004808201926020929091908290030181865afa158015610480573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906104a491906115b5565b6040516001600160a01b038089166024830152871660448201526064810186905260840160408051601f198184030181529181526020820180516001600160e01b0316636ce5768960e11b179052516001600160e01b031960e085901b16815261051792919062013880906004016115fc565b602060405180830381865afa158015610532573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906105569190611636565b949350505050565b610566610dfa565b61056e610f0b565b565b5f610579610f21565b805490915060ff600160401b82041615906001600160401b03165f8115801561059f5750825b90505f826001600160401b031660011480156105ba5750303b155b9050811580156105c8575080155b156105e65760405163f92ee8a960e01b815260040160405180910390fd5b845467ffffffffffffffff19166001178555831561061057845460ff60401b1916600160401b1785555b6001600160a01b03861661066b5760405162461bcd60e51b815260206004820152601a60248201527f4e6f6d696e614272696467653a206e6f207a65726f206164647200000000000060448201526064015b60405180910390fd5b61067487610f49565b5f80546001600160a01b0319166001600160a01b03881617905583156106d457845460ff60401b19168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b50505050505050565b60025f6106e8610f21565b8054909150600160401b900460ff168061070f575080546001600160401b03808416911610155b1561072d5760405163f92ee8a960e01b815260040160405180910390fd5b805468ffffffffffffffffff19166001600160401b03831617600160401b17815560405163095ea7b360e01b81527f0000000000000000000000000000000000000000000000000000000000000000906001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000169063095ea7b3906107d39084905f19906004016001600160a01b03929092168252602082015260400190565b6020604051808303815f875af11580156107ef573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610813919061164d565b506040516370a0823160e01b815230600482018190526001600160a01b03808416926367c6e39c92917f000000000000000000000000000000000000000000000000000000000000000016906370a0823190602401602060405180830381865afa158015610883573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906108a79190611636565b6040516001600160e01b031960e085901b1681526001600160a01b03909216600483015260248201526044015f604051808303815f87803b1580156108ea575f5ffd5b505af11580156108fc573d5f5f3e3d5ffd5b5050835460ff60401b1916845550506040516001600160401b03841681527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2915060200160405180910390a15050565b610954610dfa565b61056e5f610f5a565b610965610dfa565b61056e610fca565b7f0683d1c283a672fc58eb7940a0dba83ea98b96966a9ca1b030dec2c60cea4d1e61099781610d7e565b156109db5760405162461bcd60e51b8152602060048201526014602482015273139bdb5a5b98509c9a5919d94e881c185d5cd95960621b6044820152606401610662565b6109e6338484610fe0565b505050565b6109f3610dfa565b61042d816113b9565b610a04610dfa565b6001600160a01b038116610a2d57604051631e4fbdf760e01b81525f6004820152602401610662565b61042d81610f5a565b7f855511cc3694f64379908437d6d64458dc76d02482052bfb8a5b33a72c054c77610a6081610d7e565b15610aa45760405162461bcd60e51b8152602060048201526014602482015273139bdb5a5b98509c9a5919d94e881c185d5cd95960621b6044820152606401610662565b5f805460408051631799380760e11b815281516001600160a01b0390931692632f32700e926004808401939192918290030181865afa158015610ae9573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610b0d919061166c565b5f549091506001600160a01b03163314610b695760405162461bcd60e51b815260206004820152601760248201527f4e6f6d696e614272696467653a206e6f74207863616c6c0000000000000000006044820152606401610662565b60208101516001600160a01b0316600262048789608a1b0114610bce5760405162461bcd60e51b815260206004820152601860248201527f4e6f6d696e614272696467653a206e6f742062726964676500000000000000006044820152606401610662565b5f5f9054906101000a90046001600160a01b03166001600160a01b0316634cc390886040518163ffffffff1660e01b8152600401602060405180830381865afa158015610c1d573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610c4191906115b5565b6001600160401b0316815f01516001600160401b031614610ca45760405162461bcd60e51b815260206004820152601f60248201527f4e6f6d696e614272696467653a206e6f74206e6f6d696e6120706f7274616c006044820152606401610662565b60405163a9059cbb60e01b81526001600160a01b038581166004830152602482018590527f0000000000000000000000000000000000000000000000000000000000000000169063a9059cbb906044016020604051808303815f875af1158015610d10573d5f5f3e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610d34919061164d565b50836001600160a01b03167f884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a942436484604051610d7091815260200190565b60405180910390a250505050565b5f51602061172b5f395f51905f525f9081527fff37105740f03695c8f3597f3aff2b92fbe1c80abea3c28731ecff2efd69340060208190527ffae9838a178d7f201aa98e2ce5340158edda60bb1e8f168f46503bf3e99f13be5460ff1680610df357505f8381526020829052604090205460ff165b9392505050565b33610e2c7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c199300546001600160a01b031690565b6001600160a01b03161461056e5760405163118cdaa760e01b8152336004820152602401610662565b5f8181527fff37105740f03695c8f3597f3aff2b92fbe1c80abea3c28731ecff2efd693400602081905260409091205460ff16610ecb5760405162461bcd60e51b815260206004820152601460248201527314185d5cd8589b194e881b9bdd081c185d5cd95960621b6044820152606401610662565b5f82815260208290526040808220805460ff191690555183917fd05bfc2250abb0f8fd265a54c53a24359c5484af63cad2e4ce87c78ab751395a91a25050565b61056e5f51602061172b5f395f51905f52610e55565b5f807ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00610416565b610f5161146f565b61042d81611494565b7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c19930080546001600160a01b031981166001600160a01b03848116918217845560405192169182907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0905f90a3505050565b61056e5f51602061172b5f395f51905f526113b9565b5f811161102f5760405162461bcd60e51b815260206004820181905260248201527f4e6f6d696e614272696467653a20616d6f756e74206d757374206265203e20306044820152606401610662565b6001600160a01b0382166110855760405162461bcd60e51b815260206004820152601f60248201527f4e6f6d696e614272696467653a206e6f2062726964676520746f207a65726f006044820152606401610662565b5f5f5f9054906101000a90046001600160a01b03166001600160a01b0316634cc390886040518163ffffffff1660e01b8152600401602060405180830381865afa1580156110d5573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906110f991906115b5565b6040516001600160a01b03808716602483015285166044820152606481018490529091505f9060840160408051601f198184030181529181526020820180516001600160e01b0316636ce5768960e11b1790525f549051632376548f60e21b81529192506001600160a01b031690638dd9523c90611182908590859062013880906004016115fc565b602060405180830381865afa15801561119d573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906111c19190611636565b3410156112105760405162461bcd60e51b815260206004820152601e60248201527f4e6f6d696e614272696467653a20696e73756666696369656e742066656500006044820152606401610662565b6040516323b872dd60e01b81526001600160a01b038681166004830152306024830152604482018590527f000000000000000000000000000000000000000000000000000000000000000016906323b872dd906064016020604051808303815f875af1158015611282573d5f5f3e3d5ffd5b505050506040513d601f19601f820116820180604052508101906112a6919061164d565b6112f25760405162461bcd60e51b815260206004820152601d60248201527f4e6f6d696e614272696467653a207472616e73666572206661696c65640000006044820152606401610662565b5f5460405163c21dda4f60e01b81526001600160a01b039091169063c21dda4f903490611337908690600490600262048789608a1b01908890620138809084016116d4565b5f604051808303818588803b15801561134e575f5ffd5b505af1158015611360573d5f5f3e3d5ffd5b5050505050836001600160a01b0316856001600160a01b03167f59bc8a913d49a9626dd6ba5def7fcf12804061c1bb9b8b6db077e1a12cb4b422856040516113aa91815260200190565b60405180910390a35050505050565b5f8181527fff37105740f03695c8f3597f3aff2b92fbe1c80abea3c28731ecff2efd693400602081905260409091205460ff161561142c5760405162461bcd60e51b815260206004820152601060248201526f14185d5cd8589b194e881c185d5cd95960821b6044820152606401610662565b5f82815260208290526040808220805460ff191660011790555183917f0cb09dc71d57eeec2046f6854976717e4874a3cf2d6ddeddde337e5b6de6ba3191a25050565b61147761149c565b61056e57604051631afcd79f60e31b815260040160405180910390fd5b610a0461146f565b5f6114a5610f21565b54600160401b900460ff16919050565b5f602082840312156114c5575f5ffd5b5035919050565b6001600160a01b038116811461042d575f5ffd5b5f5f5f606084860312156114f2575f5ffd5b83356114fd816114cc565b9250602084013561150d816114cc565b929592945050506040919091013590565b5f5f6040838503121561152f575f5ffd5b823561153a816114cc565b9150602083013561154a816114cc565b809150509250929050565b5f5f60408385031215611566575f5ffd5b8235611571816114cc565b946020939093013593505050565b5f6020828403121561158f575f5ffd5b8135610df3816114cc565b80516001600160401b03811681146115b0575f5ffd5b919050565b5f602082840312156115c5575f5ffd5b610df38261159a565b5f81518084528060208401602086015e5f602082860101526020601f19601f83011685010191505092915050565b6001600160401b0384168152606060208201525f61161d60608301856115ce565b90506001600160401b0383166040830152949350505050565b5f60208284031215611646575f5ffd5b5051919050565b5f6020828403121561165d575f5ffd5b81518015158114610df3575f5ffd5b5f604082840312801561167d575f5ffd5b50604080519081016001600160401b03811182821017156116ac57634e487b7160e01b5f52604160045260245ffd5b6040526116b88361159a565b815260208301516116c8816114cc565b60208201529392505050565b6001600160401b038616815260ff851660208201526001600160a01b038416604082015260a0606082018190525f9061170f908301856115ce565b90506001600160401b0383166080830152969550505050505056fe76e8952e4b09b8d505aa08998d716721a1dbf0884ac74202e33985da1ed005e9a2646970667358221220453213437e224fa478e3700ab07965af43f987ca488777375e51c13d65a5022d64736f6c634300081e0033",
}

// NominaBridgeL1ABI is the input ABI used to generate the binding from.
// Deprecated: Use NominaBridgeL1MetaData.ABI instead.
var NominaBridgeL1ABI = NominaBridgeL1MetaData.ABI

// NominaBridgeL1Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use NominaBridgeL1MetaData.Bin instead.
var NominaBridgeL1Bin = NominaBridgeL1MetaData.Bin

// DeployNominaBridgeL1 deploys a new Ethereum contract, binding an instance of NominaBridgeL1 to it.
func DeployNominaBridgeL1(auth *bind.TransactOpts, backend bind.ContractBackend, omni_ common.Address, nomina_ common.Address) (common.Address, *types.Transaction, *NominaBridgeL1, error) {
	parsed, err := NominaBridgeL1MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(NominaBridgeL1Bin), backend, omni_, nomina_)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &NominaBridgeL1{NominaBridgeL1Caller: NominaBridgeL1Caller{contract: contract}, NominaBridgeL1Transactor: NominaBridgeL1Transactor{contract: contract}, NominaBridgeL1Filterer: NominaBridgeL1Filterer{contract: contract}}, nil
}

// NominaBridgeL1 is an auto generated Go binding around an Ethereum contract.
type NominaBridgeL1 struct {
	NominaBridgeL1Caller     // Read-only binding to the contract
	NominaBridgeL1Transactor // Write-only binding to the contract
	NominaBridgeL1Filterer   // Log filterer for contract events
}

// NominaBridgeL1Caller is an auto generated read-only Go binding around an Ethereum contract.
type NominaBridgeL1Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NominaBridgeL1Transactor is an auto generated write-only Go binding around an Ethereum contract.
type NominaBridgeL1Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NominaBridgeL1Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type NominaBridgeL1Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NominaBridgeL1Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NominaBridgeL1Session struct {
	Contract     *NominaBridgeL1   // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// NominaBridgeL1CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NominaBridgeL1CallerSession struct {
	Contract *NominaBridgeL1Caller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts         // Call options to use throughout this session
}

// NominaBridgeL1TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NominaBridgeL1TransactorSession struct {
	Contract     *NominaBridgeL1Transactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts         // Transaction auth options to use throughout this session
}

// NominaBridgeL1Raw is an auto generated low-level Go binding around an Ethereum contract.
type NominaBridgeL1Raw struct {
	Contract *NominaBridgeL1 // Generic contract binding to access the raw methods on
}

// NominaBridgeL1CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NominaBridgeL1CallerRaw struct {
	Contract *NominaBridgeL1Caller // Generic read-only contract binding to access the raw methods on
}

// NominaBridgeL1TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NominaBridgeL1TransactorRaw struct {
	Contract *NominaBridgeL1Transactor // Generic write-only contract binding to access the raw methods on
}

// NewNominaBridgeL1 creates a new instance of NominaBridgeL1, bound to a specific deployed contract.
func NewNominaBridgeL1(address common.Address, backend bind.ContractBackend) (*NominaBridgeL1, error) {
	contract, err := bindNominaBridgeL1(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &NominaBridgeL1{NominaBridgeL1Caller: NominaBridgeL1Caller{contract: contract}, NominaBridgeL1Transactor: NominaBridgeL1Transactor{contract: contract}, NominaBridgeL1Filterer: NominaBridgeL1Filterer{contract: contract}}, nil
}

// NewNominaBridgeL1Caller creates a new read-only instance of NominaBridgeL1, bound to a specific deployed contract.
func NewNominaBridgeL1Caller(address common.Address, caller bind.ContractCaller) (*NominaBridgeL1Caller, error) {
	contract, err := bindNominaBridgeL1(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NominaBridgeL1Caller{contract: contract}, nil
}

// NewNominaBridgeL1Transactor creates a new write-only instance of NominaBridgeL1, bound to a specific deployed contract.
func NewNominaBridgeL1Transactor(address common.Address, transactor bind.ContractTransactor) (*NominaBridgeL1Transactor, error) {
	contract, err := bindNominaBridgeL1(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NominaBridgeL1Transactor{contract: contract}, nil
}

// NewNominaBridgeL1Filterer creates a new log filterer instance of NominaBridgeL1, bound to a specific deployed contract.
func NewNominaBridgeL1Filterer(address common.Address, filterer bind.ContractFilterer) (*NominaBridgeL1Filterer, error) {
	contract, err := bindNominaBridgeL1(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NominaBridgeL1Filterer{contract: contract}, nil
}

// bindNominaBridgeL1 binds a generic wrapper to an already deployed contract.
func bindNominaBridgeL1(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := NominaBridgeL1MetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NominaBridgeL1 *NominaBridgeL1Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NominaBridgeL1.Contract.NominaBridgeL1Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NominaBridgeL1 *NominaBridgeL1Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NominaBridgeL1.Contract.NominaBridgeL1Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NominaBridgeL1 *NominaBridgeL1Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NominaBridgeL1.Contract.NominaBridgeL1Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NominaBridgeL1 *NominaBridgeL1CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NominaBridgeL1.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NominaBridgeL1 *NominaBridgeL1TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NominaBridgeL1.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NominaBridgeL1 *NominaBridgeL1TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NominaBridgeL1.Contract.contract.Transact(opts, method, params...)
}

// ACTIONBRIDGE is a free data retrieval call binding the contract method 0x09839a93.
//
// Solidity: function ACTION_BRIDGE() view returns(bytes32)
func (_NominaBridgeL1 *NominaBridgeL1Caller) ACTIONBRIDGE(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _NominaBridgeL1.contract.Call(opts, &out, "ACTION_BRIDGE")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ACTIONBRIDGE is a free data retrieval call binding the contract method 0x09839a93.
//
// Solidity: function ACTION_BRIDGE() view returns(bytes32)
func (_NominaBridgeL1 *NominaBridgeL1Session) ACTIONBRIDGE() ([32]byte, error) {
	return _NominaBridgeL1.Contract.ACTIONBRIDGE(&_NominaBridgeL1.CallOpts)
}

// ACTIONBRIDGE is a free data retrieval call binding the contract method 0x09839a93.
//
// Solidity: function ACTION_BRIDGE() view returns(bytes32)
func (_NominaBridgeL1 *NominaBridgeL1CallerSession) ACTIONBRIDGE() ([32]byte, error) {
	return _NominaBridgeL1.Contract.ACTIONBRIDGE(&_NominaBridgeL1.CallOpts)
}

// ACTIONWITHDRAW is a free data retrieval call binding the contract method 0x25d70f78.
//
// Solidity: function ACTION_WITHDRAW() view returns(bytes32)
func (_NominaBridgeL1 *NominaBridgeL1Caller) ACTIONWITHDRAW(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _NominaBridgeL1.contract.Call(opts, &out, "ACTION_WITHDRAW")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ACTIONWITHDRAW is a free data retrieval call binding the contract method 0x25d70f78.
//
// Solidity: function ACTION_WITHDRAW() view returns(bytes32)
func (_NominaBridgeL1 *NominaBridgeL1Session) ACTIONWITHDRAW() ([32]byte, error) {
	return _NominaBridgeL1.Contract.ACTIONWITHDRAW(&_NominaBridgeL1.CallOpts)
}

// ACTIONWITHDRAW is a free data retrieval call binding the contract method 0x25d70f78.
//
// Solidity: function ACTION_WITHDRAW() view returns(bytes32)
func (_NominaBridgeL1 *NominaBridgeL1CallerSession) ACTIONWITHDRAW() ([32]byte, error) {
	return _NominaBridgeL1.Contract.ACTIONWITHDRAW(&_NominaBridgeL1.CallOpts)
}

// KeyPauseAll is a free data retrieval call binding the contract method 0xa10ac97a.
//
// Solidity: function KeyPauseAll() view returns(bytes32)
func (_NominaBridgeL1 *NominaBridgeL1Caller) KeyPauseAll(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _NominaBridgeL1.contract.Call(opts, &out, "KeyPauseAll")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// KeyPauseAll is a free data retrieval call binding the contract method 0xa10ac97a.
//
// Solidity: function KeyPauseAll() view returns(bytes32)
func (_NominaBridgeL1 *NominaBridgeL1Session) KeyPauseAll() ([32]byte, error) {
	return _NominaBridgeL1.Contract.KeyPauseAll(&_NominaBridgeL1.CallOpts)
}

// KeyPauseAll is a free data retrieval call binding the contract method 0xa10ac97a.
//
// Solidity: function KeyPauseAll() view returns(bytes32)
func (_NominaBridgeL1 *NominaBridgeL1CallerSession) KeyPauseAll() ([32]byte, error) {
	return _NominaBridgeL1.Contract.KeyPauseAll(&_NominaBridgeL1.CallOpts)
}

// XCALLWITHDRAWGASLIMIT is a free data retrieval call binding the contract method 0x8fdcb4c9.
//
// Solidity: function XCALL_WITHDRAW_GAS_LIMIT() view returns(uint64)
func (_NominaBridgeL1 *NominaBridgeL1Caller) XCALLWITHDRAWGASLIMIT(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _NominaBridgeL1.contract.Call(opts, &out, "XCALL_WITHDRAW_GAS_LIMIT")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// XCALLWITHDRAWGASLIMIT is a free data retrieval call binding the contract method 0x8fdcb4c9.
//
// Solidity: function XCALL_WITHDRAW_GAS_LIMIT() view returns(uint64)
func (_NominaBridgeL1 *NominaBridgeL1Session) XCALLWITHDRAWGASLIMIT() (uint64, error) {
	return _NominaBridgeL1.Contract.XCALLWITHDRAWGASLIMIT(&_NominaBridgeL1.CallOpts)
}

// XCALLWITHDRAWGASLIMIT is a free data retrieval call binding the contract method 0x8fdcb4c9.
//
// Solidity: function XCALL_WITHDRAW_GAS_LIMIT() view returns(uint64)
func (_NominaBridgeL1 *NominaBridgeL1CallerSession) XCALLWITHDRAWGASLIMIT() (uint64, error) {
	return _NominaBridgeL1.Contract.XCALLWITHDRAWGASLIMIT(&_NominaBridgeL1.CallOpts)
}

// BridgeFee is a free data retrieval call binding the contract method 0x3794999d.
//
// Solidity: function bridgeFee(address payor, address to, uint256 amount) view returns(uint256)
func (_NominaBridgeL1 *NominaBridgeL1Caller) BridgeFee(opts *bind.CallOpts, payor common.Address, to common.Address, amount *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _NominaBridgeL1.contract.Call(opts, &out, "bridgeFee", payor, to, amount)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BridgeFee is a free data retrieval call binding the contract method 0x3794999d.
//
// Solidity: function bridgeFee(address payor, address to, uint256 amount) view returns(uint256)
func (_NominaBridgeL1 *NominaBridgeL1Session) BridgeFee(payor common.Address, to common.Address, amount *big.Int) (*big.Int, error) {
	return _NominaBridgeL1.Contract.BridgeFee(&_NominaBridgeL1.CallOpts, payor, to, amount)
}

// BridgeFee is a free data retrieval call binding the contract method 0x3794999d.
//
// Solidity: function bridgeFee(address payor, address to, uint256 amount) view returns(uint256)
func (_NominaBridgeL1 *NominaBridgeL1CallerSession) BridgeFee(payor common.Address, to common.Address, amount *big.Int) (*big.Int, error) {
	return _NominaBridgeL1.Contract.BridgeFee(&_NominaBridgeL1.CallOpts, payor, to, amount)
}

// IsPaused is a free data retrieval call binding the contract method 0x241b71bb.
//
// Solidity: function isPaused(bytes32 action) view returns(bool)
func (_NominaBridgeL1 *NominaBridgeL1Caller) IsPaused(opts *bind.CallOpts, action [32]byte) (bool, error) {
	var out []interface{}
	err := _NominaBridgeL1.contract.Call(opts, &out, "isPaused", action)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPaused is a free data retrieval call binding the contract method 0x241b71bb.
//
// Solidity: function isPaused(bytes32 action) view returns(bool)
func (_NominaBridgeL1 *NominaBridgeL1Session) IsPaused(action [32]byte) (bool, error) {
	return _NominaBridgeL1.Contract.IsPaused(&_NominaBridgeL1.CallOpts, action)
}

// IsPaused is a free data retrieval call binding the contract method 0x241b71bb.
//
// Solidity: function isPaused(bytes32 action) view returns(bool)
func (_NominaBridgeL1 *NominaBridgeL1CallerSession) IsPaused(action [32]byte) (bool, error) {
	return _NominaBridgeL1.Contract.IsPaused(&_NominaBridgeL1.CallOpts, action)
}

// Nomina is a free data retrieval call binding the contract method 0x646310c8.
//
// Solidity: function nomina() view returns(address)
func (_NominaBridgeL1 *NominaBridgeL1Caller) Nomina(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NominaBridgeL1.contract.Call(opts, &out, "nomina")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Nomina is a free data retrieval call binding the contract method 0x646310c8.
//
// Solidity: function nomina() view returns(address)
func (_NominaBridgeL1 *NominaBridgeL1Session) Nomina() (common.Address, error) {
	return _NominaBridgeL1.Contract.Nomina(&_NominaBridgeL1.CallOpts)
}

// Nomina is a free data retrieval call binding the contract method 0x646310c8.
//
// Solidity: function nomina() view returns(address)
func (_NominaBridgeL1 *NominaBridgeL1CallerSession) Nomina() (common.Address, error) {
	return _NominaBridgeL1.Contract.Nomina(&_NominaBridgeL1.CallOpts)
}

// Omni is a free data retrieval call binding the contract method 0x39acf9f1.
//
// Solidity: function omni() view returns(address)
func (_NominaBridgeL1 *NominaBridgeL1Caller) Omni(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NominaBridgeL1.contract.Call(opts, &out, "omni")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Omni is a free data retrieval call binding the contract method 0x39acf9f1.
//
// Solidity: function omni() view returns(address)
func (_NominaBridgeL1 *NominaBridgeL1Session) Omni() (common.Address, error) {
	return _NominaBridgeL1.Contract.Omni(&_NominaBridgeL1.CallOpts)
}

// Omni is a free data retrieval call binding the contract method 0x39acf9f1.
//
// Solidity: function omni() view returns(address)
func (_NominaBridgeL1 *NominaBridgeL1CallerSession) Omni() (common.Address, error) {
	return _NominaBridgeL1.Contract.Omni(&_NominaBridgeL1.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_NominaBridgeL1 *NominaBridgeL1Caller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NominaBridgeL1.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_NominaBridgeL1 *NominaBridgeL1Session) Owner() (common.Address, error) {
	return _NominaBridgeL1.Contract.Owner(&_NominaBridgeL1.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_NominaBridgeL1 *NominaBridgeL1CallerSession) Owner() (common.Address, error) {
	return _NominaBridgeL1.Contract.Owner(&_NominaBridgeL1.CallOpts)
}

// Portal is a free data retrieval call binding the contract method 0x6425666b.
//
// Solidity: function portal() view returns(address)
func (_NominaBridgeL1 *NominaBridgeL1Caller) Portal(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NominaBridgeL1.contract.Call(opts, &out, "portal")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Portal is a free data retrieval call binding the contract method 0x6425666b.
//
// Solidity: function portal() view returns(address)
func (_NominaBridgeL1 *NominaBridgeL1Session) Portal() (common.Address, error) {
	return _NominaBridgeL1.Contract.Portal(&_NominaBridgeL1.CallOpts)
}

// Portal is a free data retrieval call binding the contract method 0x6425666b.
//
// Solidity: function portal() view returns(address)
func (_NominaBridgeL1 *NominaBridgeL1CallerSession) Portal() (common.Address, error) {
	return _NominaBridgeL1.Contract.Portal(&_NominaBridgeL1.CallOpts)
}

// Bridge is a paid mutator transaction binding the contract method 0xc3de453d.
//
// Solidity: function bridge(address to, uint256 amount) payable returns()
func (_NominaBridgeL1 *NominaBridgeL1Transactor) Bridge(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _NominaBridgeL1.contract.Transact(opts, "bridge", to, amount)
}

// Bridge is a paid mutator transaction binding the contract method 0xc3de453d.
//
// Solidity: function bridge(address to, uint256 amount) payable returns()
func (_NominaBridgeL1 *NominaBridgeL1Session) Bridge(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _NominaBridgeL1.Contract.Bridge(&_NominaBridgeL1.TransactOpts, to, amount)
}

// Bridge is a paid mutator transaction binding the contract method 0xc3de453d.
//
// Solidity: function bridge(address to, uint256 amount) payable returns()
func (_NominaBridgeL1 *NominaBridgeL1TransactorSession) Bridge(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _NominaBridgeL1.Contract.Bridge(&_NominaBridgeL1.TransactOpts, to, amount)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address owner_, address portal_) returns()
func (_NominaBridgeL1 *NominaBridgeL1Transactor) Initialize(opts *bind.TransactOpts, owner_ common.Address, portal_ common.Address) (*types.Transaction, error) {
	return _NominaBridgeL1.contract.Transact(opts, "initialize", owner_, portal_)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address owner_, address portal_) returns()
func (_NominaBridgeL1 *NominaBridgeL1Session) Initialize(owner_ common.Address, portal_ common.Address) (*types.Transaction, error) {
	return _NominaBridgeL1.Contract.Initialize(&_NominaBridgeL1.TransactOpts, owner_, portal_)
}

// Initialize is a paid mutator transaction binding the contract method 0x485cc955.
//
// Solidity: function initialize(address owner_, address portal_) returns()
func (_NominaBridgeL1 *NominaBridgeL1TransactorSession) Initialize(owner_ common.Address, portal_ common.Address) (*types.Transaction, error) {
	return _NominaBridgeL1.Contract.Initialize(&_NominaBridgeL1.TransactOpts, owner_, portal_)
}

// InitializeV2 is a paid mutator transaction binding the contract method 0x5cd8a76b.
//
// Solidity: function initializeV2() returns()
func (_NominaBridgeL1 *NominaBridgeL1Transactor) InitializeV2(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NominaBridgeL1.contract.Transact(opts, "initializeV2")
}

// InitializeV2 is a paid mutator transaction binding the contract method 0x5cd8a76b.
//
// Solidity: function initializeV2() returns()
func (_NominaBridgeL1 *NominaBridgeL1Session) InitializeV2() (*types.Transaction, error) {
	return _NominaBridgeL1.Contract.InitializeV2(&_NominaBridgeL1.TransactOpts)
}

// InitializeV2 is a paid mutator transaction binding the contract method 0x5cd8a76b.
//
// Solidity: function initializeV2() returns()
func (_NominaBridgeL1 *NominaBridgeL1TransactorSession) InitializeV2() (*types.Transaction, error) {
	return _NominaBridgeL1.Contract.InitializeV2(&_NominaBridgeL1.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_NominaBridgeL1 *NominaBridgeL1Transactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NominaBridgeL1.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_NominaBridgeL1 *NominaBridgeL1Session) Pause() (*types.Transaction, error) {
	return _NominaBridgeL1.Contract.Pause(&_NominaBridgeL1.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_NominaBridgeL1 *NominaBridgeL1TransactorSession) Pause() (*types.Transaction, error) {
	return _NominaBridgeL1.Contract.Pause(&_NominaBridgeL1.TransactOpts)
}

// Pause0 is a paid mutator transaction binding the contract method 0xed56531a.
//
// Solidity: function pause(bytes32 action) returns()
func (_NominaBridgeL1 *NominaBridgeL1Transactor) Pause0(opts *bind.TransactOpts, action [32]byte) (*types.Transaction, error) {
	return _NominaBridgeL1.contract.Transact(opts, "pause0", action)
}

// Pause0 is a paid mutator transaction binding the contract method 0xed56531a.
//
// Solidity: function pause(bytes32 action) returns()
func (_NominaBridgeL1 *NominaBridgeL1Session) Pause0(action [32]byte) (*types.Transaction, error) {
	return _NominaBridgeL1.Contract.Pause0(&_NominaBridgeL1.TransactOpts, action)
}

// Pause0 is a paid mutator transaction binding the contract method 0xed56531a.
//
// Solidity: function pause(bytes32 action) returns()
func (_NominaBridgeL1 *NominaBridgeL1TransactorSession) Pause0(action [32]byte) (*types.Transaction, error) {
	return _NominaBridgeL1.Contract.Pause0(&_NominaBridgeL1.TransactOpts, action)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_NominaBridgeL1 *NominaBridgeL1Transactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NominaBridgeL1.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_NominaBridgeL1 *NominaBridgeL1Session) RenounceOwnership() (*types.Transaction, error) {
	return _NominaBridgeL1.Contract.RenounceOwnership(&_NominaBridgeL1.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_NominaBridgeL1 *NominaBridgeL1TransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _NominaBridgeL1.Contract.RenounceOwnership(&_NominaBridgeL1.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_NominaBridgeL1 *NominaBridgeL1Transactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _NominaBridgeL1.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_NominaBridgeL1 *NominaBridgeL1Session) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _NominaBridgeL1.Contract.TransferOwnership(&_NominaBridgeL1.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_NominaBridgeL1 *NominaBridgeL1TransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _NominaBridgeL1.Contract.TransferOwnership(&_NominaBridgeL1.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x2f4dae9f.
//
// Solidity: function unpause(bytes32 action) returns()
func (_NominaBridgeL1 *NominaBridgeL1Transactor) Unpause(opts *bind.TransactOpts, action [32]byte) (*types.Transaction, error) {
	return _NominaBridgeL1.contract.Transact(opts, "unpause", action)
}

// Unpause is a paid mutator transaction binding the contract method 0x2f4dae9f.
//
// Solidity: function unpause(bytes32 action) returns()
func (_NominaBridgeL1 *NominaBridgeL1Session) Unpause(action [32]byte) (*types.Transaction, error) {
	return _NominaBridgeL1.Contract.Unpause(&_NominaBridgeL1.TransactOpts, action)
}

// Unpause is a paid mutator transaction binding the contract method 0x2f4dae9f.
//
// Solidity: function unpause(bytes32 action) returns()
func (_NominaBridgeL1 *NominaBridgeL1TransactorSession) Unpause(action [32]byte) (*types.Transaction, error) {
	return _NominaBridgeL1.Contract.Unpause(&_NominaBridgeL1.TransactOpts, action)
}

// Unpause0 is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_NominaBridgeL1 *NominaBridgeL1Transactor) Unpause0(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NominaBridgeL1.contract.Transact(opts, "unpause0")
}

// Unpause0 is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_NominaBridgeL1 *NominaBridgeL1Session) Unpause0() (*types.Transaction, error) {
	return _NominaBridgeL1.Contract.Unpause0(&_NominaBridgeL1.TransactOpts)
}

// Unpause0 is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_NominaBridgeL1 *NominaBridgeL1TransactorSession) Unpause0() (*types.Transaction, error) {
	return _NominaBridgeL1.Contract.Unpause0(&_NominaBridgeL1.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address to, uint256 amount) returns()
func (_NominaBridgeL1 *NominaBridgeL1Transactor) Withdraw(opts *bind.TransactOpts, to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _NominaBridgeL1.contract.Transact(opts, "withdraw", to, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address to, uint256 amount) returns()
func (_NominaBridgeL1 *NominaBridgeL1Session) Withdraw(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _NominaBridgeL1.Contract.Withdraw(&_NominaBridgeL1.TransactOpts, to, amount)
}

// Withdraw is a paid mutator transaction binding the contract method 0xf3fef3a3.
//
// Solidity: function withdraw(address to, uint256 amount) returns()
func (_NominaBridgeL1 *NominaBridgeL1TransactorSession) Withdraw(to common.Address, amount *big.Int) (*types.Transaction, error) {
	return _NominaBridgeL1.Contract.Withdraw(&_NominaBridgeL1.TransactOpts, to, amount)
}

// NominaBridgeL1BridgeIterator is returned from FilterBridge and is used to iterate over the raw logs and unpacked data for Bridge events raised by the NominaBridgeL1 contract.
type NominaBridgeL1BridgeIterator struct {
	Event *NominaBridgeL1Bridge // Event containing the contract specifics and raw log

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
func (it *NominaBridgeL1BridgeIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NominaBridgeL1Bridge)
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
		it.Event = new(NominaBridgeL1Bridge)
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
func (it *NominaBridgeL1BridgeIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NominaBridgeL1BridgeIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NominaBridgeL1Bridge represents a Bridge event raised by the NominaBridgeL1 contract.
type NominaBridgeL1Bridge struct {
	Payor  common.Address
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBridge is a free log retrieval operation binding the contract event 0x59bc8a913d49a9626dd6ba5def7fcf12804061c1bb9b8b6db077e1a12cb4b422.
//
// Solidity: event Bridge(address indexed payor, address indexed to, uint256 amount)
func (_NominaBridgeL1 *NominaBridgeL1Filterer) FilterBridge(opts *bind.FilterOpts, payor []common.Address, to []common.Address) (*NominaBridgeL1BridgeIterator, error) {

	var payorRule []interface{}
	for _, payorItem := range payor {
		payorRule = append(payorRule, payorItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _NominaBridgeL1.contract.FilterLogs(opts, "Bridge", payorRule, toRule)
	if err != nil {
		return nil, err
	}
	return &NominaBridgeL1BridgeIterator{contract: _NominaBridgeL1.contract, event: "Bridge", logs: logs, sub: sub}, nil
}

// WatchBridge is a free log subscription operation binding the contract event 0x59bc8a913d49a9626dd6ba5def7fcf12804061c1bb9b8b6db077e1a12cb4b422.
//
// Solidity: event Bridge(address indexed payor, address indexed to, uint256 amount)
func (_NominaBridgeL1 *NominaBridgeL1Filterer) WatchBridge(opts *bind.WatchOpts, sink chan<- *NominaBridgeL1Bridge, payor []common.Address, to []common.Address) (event.Subscription, error) {

	var payorRule []interface{}
	for _, payorItem := range payor {
		payorRule = append(payorRule, payorItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _NominaBridgeL1.contract.WatchLogs(opts, "Bridge", payorRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NominaBridgeL1Bridge)
				if err := _NominaBridgeL1.contract.UnpackLog(event, "Bridge", log); err != nil {
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
func (_NominaBridgeL1 *NominaBridgeL1Filterer) ParseBridge(log types.Log) (*NominaBridgeL1Bridge, error) {
	event := new(NominaBridgeL1Bridge)
	if err := _NominaBridgeL1.contract.UnpackLog(event, "Bridge", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NominaBridgeL1InitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the NominaBridgeL1 contract.
type NominaBridgeL1InitializedIterator struct {
	Event *NominaBridgeL1Initialized // Event containing the contract specifics and raw log

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
func (it *NominaBridgeL1InitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NominaBridgeL1Initialized)
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
		it.Event = new(NominaBridgeL1Initialized)
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
func (it *NominaBridgeL1InitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NominaBridgeL1InitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NominaBridgeL1Initialized represents a Initialized event raised by the NominaBridgeL1 contract.
type NominaBridgeL1Initialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_NominaBridgeL1 *NominaBridgeL1Filterer) FilterInitialized(opts *bind.FilterOpts) (*NominaBridgeL1InitializedIterator, error) {

	logs, sub, err := _NominaBridgeL1.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &NominaBridgeL1InitializedIterator{contract: _NominaBridgeL1.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_NominaBridgeL1 *NominaBridgeL1Filterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *NominaBridgeL1Initialized) (event.Subscription, error) {

	logs, sub, err := _NominaBridgeL1.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NominaBridgeL1Initialized)
				if err := _NominaBridgeL1.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_NominaBridgeL1 *NominaBridgeL1Filterer) ParseInitialized(log types.Log) (*NominaBridgeL1Initialized, error) {
	event := new(NominaBridgeL1Initialized)
	if err := _NominaBridgeL1.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NominaBridgeL1OwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the NominaBridgeL1 contract.
type NominaBridgeL1OwnershipTransferredIterator struct {
	Event *NominaBridgeL1OwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *NominaBridgeL1OwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NominaBridgeL1OwnershipTransferred)
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
		it.Event = new(NominaBridgeL1OwnershipTransferred)
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
func (it *NominaBridgeL1OwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NominaBridgeL1OwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NominaBridgeL1OwnershipTransferred represents a OwnershipTransferred event raised by the NominaBridgeL1 contract.
type NominaBridgeL1OwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_NominaBridgeL1 *NominaBridgeL1Filterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*NominaBridgeL1OwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _NominaBridgeL1.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &NominaBridgeL1OwnershipTransferredIterator{contract: _NominaBridgeL1.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_NominaBridgeL1 *NominaBridgeL1Filterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *NominaBridgeL1OwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _NominaBridgeL1.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NominaBridgeL1OwnershipTransferred)
				if err := _NominaBridgeL1.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_NominaBridgeL1 *NominaBridgeL1Filterer) ParseOwnershipTransferred(log types.Log) (*NominaBridgeL1OwnershipTransferred, error) {
	event := new(NominaBridgeL1OwnershipTransferred)
	if err := _NominaBridgeL1.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NominaBridgeL1PausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the NominaBridgeL1 contract.
type NominaBridgeL1PausedIterator struct {
	Event *NominaBridgeL1Paused // Event containing the contract specifics and raw log

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
func (it *NominaBridgeL1PausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NominaBridgeL1Paused)
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
		it.Event = new(NominaBridgeL1Paused)
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
func (it *NominaBridgeL1PausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NominaBridgeL1PausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NominaBridgeL1Paused represents a Paused event raised by the NominaBridgeL1 contract.
type NominaBridgeL1Paused struct {
	Key [32]byte
	Raw types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x0cb09dc71d57eeec2046f6854976717e4874a3cf2d6ddeddde337e5b6de6ba31.
//
// Solidity: event Paused(bytes32 indexed key)
func (_NominaBridgeL1 *NominaBridgeL1Filterer) FilterPaused(opts *bind.FilterOpts, key [][32]byte) (*NominaBridgeL1PausedIterator, error) {

	var keyRule []interface{}
	for _, keyItem := range key {
		keyRule = append(keyRule, keyItem)
	}

	logs, sub, err := _NominaBridgeL1.contract.FilterLogs(opts, "Paused", keyRule)
	if err != nil {
		return nil, err
	}
	return &NominaBridgeL1PausedIterator{contract: _NominaBridgeL1.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x0cb09dc71d57eeec2046f6854976717e4874a3cf2d6ddeddde337e5b6de6ba31.
//
// Solidity: event Paused(bytes32 indexed key)
func (_NominaBridgeL1 *NominaBridgeL1Filterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *NominaBridgeL1Paused, key [][32]byte) (event.Subscription, error) {

	var keyRule []interface{}
	for _, keyItem := range key {
		keyRule = append(keyRule, keyItem)
	}

	logs, sub, err := _NominaBridgeL1.contract.WatchLogs(opts, "Paused", keyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NominaBridgeL1Paused)
				if err := _NominaBridgeL1.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_NominaBridgeL1 *NominaBridgeL1Filterer) ParsePaused(log types.Log) (*NominaBridgeL1Paused, error) {
	event := new(NominaBridgeL1Paused)
	if err := _NominaBridgeL1.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NominaBridgeL1UnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the NominaBridgeL1 contract.
type NominaBridgeL1UnpausedIterator struct {
	Event *NominaBridgeL1Unpaused // Event containing the contract specifics and raw log

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
func (it *NominaBridgeL1UnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NominaBridgeL1Unpaused)
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
		it.Event = new(NominaBridgeL1Unpaused)
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
func (it *NominaBridgeL1UnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NominaBridgeL1UnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NominaBridgeL1Unpaused represents a Unpaused event raised by the NominaBridgeL1 contract.
type NominaBridgeL1Unpaused struct {
	Key [32]byte
	Raw types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0xd05bfc2250abb0f8fd265a54c53a24359c5484af63cad2e4ce87c78ab751395a.
//
// Solidity: event Unpaused(bytes32 indexed key)
func (_NominaBridgeL1 *NominaBridgeL1Filterer) FilterUnpaused(opts *bind.FilterOpts, key [][32]byte) (*NominaBridgeL1UnpausedIterator, error) {

	var keyRule []interface{}
	for _, keyItem := range key {
		keyRule = append(keyRule, keyItem)
	}

	logs, sub, err := _NominaBridgeL1.contract.FilterLogs(opts, "Unpaused", keyRule)
	if err != nil {
		return nil, err
	}
	return &NominaBridgeL1UnpausedIterator{contract: _NominaBridgeL1.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0xd05bfc2250abb0f8fd265a54c53a24359c5484af63cad2e4ce87c78ab751395a.
//
// Solidity: event Unpaused(bytes32 indexed key)
func (_NominaBridgeL1 *NominaBridgeL1Filterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *NominaBridgeL1Unpaused, key [][32]byte) (event.Subscription, error) {

	var keyRule []interface{}
	for _, keyItem := range key {
		keyRule = append(keyRule, keyItem)
	}

	logs, sub, err := _NominaBridgeL1.contract.WatchLogs(opts, "Unpaused", keyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NominaBridgeL1Unpaused)
				if err := _NominaBridgeL1.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_NominaBridgeL1 *NominaBridgeL1Filterer) ParseUnpaused(log types.Log) (*NominaBridgeL1Unpaused, error) {
	event := new(NominaBridgeL1Unpaused)
	if err := _NominaBridgeL1.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NominaBridgeL1WithdrawIterator is returned from FilterWithdraw and is used to iterate over the raw logs and unpacked data for Withdraw events raised by the NominaBridgeL1 contract.
type NominaBridgeL1WithdrawIterator struct {
	Event *NominaBridgeL1Withdraw // Event containing the contract specifics and raw log

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
func (it *NominaBridgeL1WithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NominaBridgeL1Withdraw)
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
		it.Event = new(NominaBridgeL1Withdraw)
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
func (it *NominaBridgeL1WithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NominaBridgeL1WithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NominaBridgeL1Withdraw represents a Withdraw event raised by the NominaBridgeL1 contract.
type NominaBridgeL1Withdraw struct {
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWithdraw is a free log retrieval operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address indexed to, uint256 amount)
func (_NominaBridgeL1 *NominaBridgeL1Filterer) FilterWithdraw(opts *bind.FilterOpts, to []common.Address) (*NominaBridgeL1WithdrawIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _NominaBridgeL1.contract.FilterLogs(opts, "Withdraw", toRule)
	if err != nil {
		return nil, err
	}
	return &NominaBridgeL1WithdrawIterator{contract: _NominaBridgeL1.contract, event: "Withdraw", logs: logs, sub: sub}, nil
}

// WatchWithdraw is a free log subscription operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address indexed to, uint256 amount)
func (_NominaBridgeL1 *NominaBridgeL1Filterer) WatchWithdraw(opts *bind.WatchOpts, sink chan<- *NominaBridgeL1Withdraw, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _NominaBridgeL1.contract.WatchLogs(opts, "Withdraw", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NominaBridgeL1Withdraw)
				if err := _NominaBridgeL1.contract.UnpackLog(event, "Withdraw", log); err != nil {
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

// ParseWithdraw is a log parse operation binding the contract event 0x884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a9424364.
//
// Solidity: event Withdraw(address indexed to, uint256 amount)
func (_NominaBridgeL1 *NominaBridgeL1Filterer) ParseWithdraw(log types.Log) (*NominaBridgeL1Withdraw, error) {
	event := new(NominaBridgeL1Withdraw)
	if err := _NominaBridgeL1.contract.UnpackLog(event, "Withdraw", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
