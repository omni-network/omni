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
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"omni\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"nomina\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"ACTION_BRIDGE\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"ACTION_WITHDRAW\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"KeyPauseAll\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"NOMINA\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIERC20\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"OMNI\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIERC20\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"XCALL_WITHDRAW_GAS_LIMIT\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"bridge\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"bridgeFee\",\"inputs\":[{\"name\":\"payor\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"owner_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"portal_\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"initializeV2\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"initializeV3\",\"inputs\":[{\"name\":\"postHaltRoot_\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isPaused\",\"inputs\":[{\"name\":\"action\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[{\"name\":\"action\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"portal\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIOmniPortal\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"postHaltClaimed\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"postHaltRoot\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"postHaltWithdraw\",\"inputs\":[{\"name\":\"accounts\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"amounts\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"},{\"name\":\"multiProof\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"multiProofFlags\",\"type\":\"bool[]\",\"internalType\":\"bool[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[{\"name\":\"action\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdraw\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Bridge\",\"inputs\":[{\"name\":\"payor\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"key\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"PostHaltWithdraw\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"key\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Withdraw\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MerkleProofInvalidMultiproof\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnableInvalidOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnableUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]}]",
	Bin: "0x60c060405234801562000010575f80fd5b506040516200241338038062002413833981016040819052620000339162000128565b6001600160a01b03808316608052811660a0526200005062000058565b50506200015e565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000900460ff1615620000a95760405163f92ee8a960e01b815260040160405180910390fd5b80546001600160401b0390811614620001095780546001600160401b0319166001600160401b0390811782556040519081527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b50565b80516001600160a01b038116811462000123575f80fd5b919050565b5f80604083850312156200013a575f80fd5b62000145836200010c565b915062000155602084016200010c565b90509250929050565b60805160a051612262620001b15f395f818161032a015281816109c501528181610d3001528181610de90152818161139a015261191d01525f818161015c01528181610d5f0152610e1401526122625ff3fe608060405260043610610147575f3560e01c8063715018a6116100b3578063c3de453d1161006d578063c3de453d146103d6578063e4cbdeed146103e9578063e849f0eb14610408578063ed56531a14610436578063f2fde38b14610455578063f3fef3a314610474575f80fd5b8063715018a6146102f15780638456cb59146103055780638ccc0d18146103195780638da5cb5b1461034c5780638fdcb4c914610388578063a10ac97a146103b6575f80fd5b80633794999d116101045780633794999d1461024e5780633cb174801461026d5780633f4ba83a1461028c578063485cc955146102a05780635cd8a76b146102bf5780636425666b146102d3575f80fd5b8063063bdf281461014b57806309839a931461019b5780631d822629146101c9578063241b71bb146101de57806325d70f781461020d5780632f4dae9f1461022d575b5f80fd5b348015610156575f80fd5b5061017e7f000000000000000000000000000000000000000000000000000000000000000081565b6040516001600160a01b0390911681526020015b60405180910390f35b3480156101a6575f80fd5b506101bb5f805160206121ed83398151915281565b604051908152602001610192565b3480156101d4575f80fd5b506101bb60015481565b3480156101e9575f80fd5b506101fd6101f8366004611de9565b610493565b6040519015158152602001610192565b348015610218575f80fd5b506101bb5f8051602061220d83398151915281565b348015610238575f80fd5b5061024c610247366004611de9565b6104a3565b005b348015610259575f80fd5b506101bb610268366004611e14565b6104b7565b348015610278575f80fd5b5061024c610287366004611e99565b6105e7565b348015610297575f80fd5b5061024c610b2e565b3480156102ab575f80fd5b5061024c6102ba366004611f53565b610b40565b3480156102ca575f80fd5b5061024c610ca8565b3480156102de575f80fd5b505f5461017e906001600160a01b031681565b3480156102fc575f80fd5b5061024c610f21565b348015610310575f80fd5b5061024c610f32565b348015610324575f80fd5b5061017e7f000000000000000000000000000000000000000000000000000000000000000081565b348015610357575f80fd5b507f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c199300546001600160a01b031661017e565b348015610393575f80fd5b5061039e6201388081565b6040516001600160401b039091168152602001610192565b3480156103c1575f80fd5b506101bb5f805160206121cd83398151915281565b61024c6103e4366004611f8a565b610f42565b3480156103f4575f80fd5b5061024c610403366004611de9565b610fad565b348015610413575f80fd5b506101fd610422366004611fb4565b60026020525f908152604090205460ff1681565b348015610441575f80fd5b5061024c610450366004611de9565b6110cf565b348015610460575f80fd5b5061024c61046f366004611fb4565b6110e0565b34801561047f575f80fd5b5061024c61048e366004611f8a565b61111a565b5f61049d8261144e565b92915050565b6104ab6114c4565b6104b48161151f565b50565b5f80546040805163110ff5f160e01b815290516001600160a01b0390921691638dd9523c91839163110ff5f1916004808201926020929091908290030181865afa158015610507573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061052b9190611fea565b6040516001600160a01b038089166024830152871660448201526064810186905260840160408051601f198184030181529181526020820180516001600160e01b0316636ce5768960e11b179052516001600160e01b031960e085901b16815261059e9291906201388090600401612046565b602060405180830381865afa1580156105b9573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906105dd919061207b565b90505b9392505050565b86851461063b5760405162461bcd60e51b815260206004820152601d60248201527f4e6f6d696e614272696467653a206c656e677468206d69736d6174636800000060448201526064015b60405180910390fd5b60015461068a5760405162461bcd60e51b815260206004820152601960248201527f4e6f6d696e614272696467653a206e6f20726f6f7420736574000000000000006044820152606401610632565b5f876001600160401b038111156106a3576106a3612092565b6040519080825280602002602001820160405280156106cc578160200160208202803683370190505b5090505f5b888110156108a6575f8a8a838181106106ec576106ec6120a6565b90506020020160208101906107019190611fb4565b90505f898984818110610716576107166120a6565b6001600160a01b0385165f908152600260209081526040909120549102929092013592505060ff161561078b5760405162461bcd60e51b815260206004820152601d60248201527f4e6f6d696e614272696467653a20616c726561647920636c61696d65640000006044820152606401610632565b6001600160a01b0382166107e15760405162461bcd60e51b815260206004820152601a60248201527f4e6f6d696e614272696467653a206e6f207a65726f20616464720000000000006044820152606401610632565b5f81116108305760405162461bcd60e51b815260206004820181905260248201527f4e6f6d696e614272696467653a20616d6f756e74206d757374206265203e20306044820152606401610632565b604080516001600160a01b038416602082015290810182905260600160408051601f198184030181528282528051602091820120908301520160405160208183030381529060405280519060200120848481518110610891576108916120a6565b602090810291909101015250506001016106d1565b506109178585808060200260200160405190810160405280939291908181526020018383602002808284375f92019190915250506040805160208089028281018201909352888252909350889250879182918501908490808284375f920191909152505060015491508590506115d5565b6109635760405162461bcd60e51b815260206004820152601b60248201527f4e6f6d696e614272696467653a20696e76616c69642070726f6f6600000000006044820152606401610632565b5f5b88811015610b2257600160025f8c8c85818110610984576109846120a6565b90506020020160208101906109999190611fb4565b6001600160a01b03908116825260208201929092526040015f20805460ff1916921515929092179091557f00000000000000000000000000000000000000000000000000000000000000001663a9059cbb8b8b848181106109fc576109fc6120a6565b9050602002016020810190610a119190611fb4565b8a8a85818110610a2357610a236120a6565b6040516001600160e01b031960e087901b1681526001600160a01b03909416600485015260200291909101356024830152506044016020604051808303815f875af1158015610a74573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610a9891906120ba565b50898982818110610aab57610aab6120a6565b9050602002016020810190610ac09190611fb4565b6001600160a01b03167f0c8025c30a7d66106b1ae8ed6f3dcbbf8d9ed725bb599f3bc855d11fa9203687898984818110610afc57610afc6120a6565b90506020020135604051610b1291815260200190565b60405180910390a2600101610965565b50505050505050505050565b610b366114c4565b610b3e6115ec565b565b5f610b49611602565b805490915060ff600160401b82041615906001600160401b03165f81158015610b6f5750825b90505f826001600160401b03166001148015610b8a5750303b155b905081158015610b98575080155b15610bb65760405163f92ee8a960e01b815260040160405180910390fd5b845467ffffffffffffffff191660011785558315610be057845460ff60401b1916600160401b1785555b6001600160a01b038616610c365760405162461bcd60e51b815260206004820152601a60248201527f4e6f6d696e614272696467653a206e6f207a65726f20616464720000000000006044820152606401610632565b610c3f8761162a565b5f80546001600160a01b0319166001600160a01b0388161790558315610c9f57845460ff60401b19168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b50505050505050565b60025f610cb3611602565b8054909150600160401b900460ff1680610cda575080546001600160401b03808416911610155b15610cf85760405163f92ee8a960e01b815260040160405180910390fd5b805468ffffffffffffffffff19166001600160401b03831617600160401b17815560405163095ea7b360e01b81526001600160a01b037f0000000000000000000000000000000000000000000000000000000000000000811660048301525f1960248301527f0000000000000000000000000000000000000000000000000000000000000000169063095ea7b3906044016020604051808303815f875af1158015610da5573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610dc991906120ba565b506040516370a0823160e01b815230600482018190526001600160a01b037f00000000000000000000000000000000000000000000000000000000000000008116926367c6e39c92917f000000000000000000000000000000000000000000000000000000000000000016906370a0823190602401602060405180830381865afa158015610e59573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610e7d919061207b565b6040516001600160e01b031960e085901b1681526001600160a01b03909216600483015260248201526044015f604051808303815f87803b158015610ec0575f80fd5b505af1158015610ed2573d5f803e3d5ffd5b5050825460ff60401b1916835550506040516001600160401b03831681527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15050565b610f296114c4565b610b3e5f61163b565b610f3a6114c4565b610b3e6116ab565b5f805160206121ed833981519152610f598161144e565b15610f9d5760405162461bcd60e51b8152602060048201526014602482015273139bdb5a5b98509c9a5919d94e881c185d5cd95960621b6044820152606401610632565b610fa83384846116c1565b505050565b60035f610fb8611602565b8054909150600160401b900460ff1680610fdf575080546001600160401b03808416911610155b15610ffd5760405163f92ee8a960e01b815260040160405180910390fd5b805468ffffffffffffffffff19166001600160401b03831617600160401b17815560018390556110395f8051602061220d83398151915261144e565b611053576110535f8051602061220d833981519152611a9a565b6110695f805160206121ed83398151915261144e565b611083576110835f805160206121ed833981519152611a9a565b805460ff60401b191681556040516001600160401b03831681527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a1505050565b6110d76114c4565b6104b481611a9a565b6110e86114c4565b6001600160a01b03811661111157604051631e4fbdf760e01b81525f6004820152602401610632565b6104b48161163b565b5f8051602061220d8339815191526111318161144e565b156111755760405162461bcd60e51b8152602060048201526014602482015273139bdb5a5b98509c9a5919d94e881c185d5cd95960621b6044820152606401610632565b5f805460408051631799380760e11b815281516001600160a01b0390931692632f32700e926004808401939192918290030181865afa1580156111ba573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906111de91906120d9565b5f549091506001600160a01b0316331461123a5760405162461bcd60e51b815260206004820152601760248201527f4e6f6d696e614272696467653a206e6f74207863616c6c0000000000000000006044820152606401610632565b60208101516001600160a01b0316600262048789608a1b011461129f5760405162461bcd60e51b815260206004820152601860248201527f4e6f6d696e614272696467653a206e6f742062726964676500000000000000006044820152606401610632565b5f8054906101000a90046001600160a01b03166001600160a01b031663110ff5f16040518163ffffffff1660e01b8152600401602060405180830381865afa1580156112ed573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906113119190611fea565b6001600160401b0316815f01516001600160401b0316146113745760405162461bcd60e51b815260206004820152601d60248201527f4e6f6d696e614272696467653a206e6f74206f6d6e6920706f7274616c0000006044820152606401610632565b60405163a9059cbb60e01b81526001600160a01b038581166004830152602482018590527f0000000000000000000000000000000000000000000000000000000000000000169063a9059cbb906044016020604051808303815f875af11580156113e0573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061140491906120ba565b50836001600160a01b03167f884edad9ce6fa2440d8a54cc123490eb96d2768479d49ff9c7366125a94243648460405161144091815260200190565b60405180910390a250505050565b5f805160206121cd8339815191525f9081527fff37105740f03695c8f3597f3aff2b92fbe1c80abea3c28731ecff2efd69340060208190527ffae9838a178d7f201aa98e2ce5340158edda60bb1e8f168f46503bf3e99f13be5460ff16806105e057505f92835260205250604090205460ff1690565b336114f67f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c199300546001600160a01b031690565b6001600160a01b031614610b3e5760405163118cdaa760e01b8152336004820152602401610632565b5f8181527fff37105740f03695c8f3597f3aff2b92fbe1c80abea3c28731ecff2efd693400602081905260409091205460ff166115955760405162461bcd60e51b815260206004820152601460248201527314185d5cd8589b194e881b9bdd081c185d5cd95960621b6044820152606401610632565b5f82815260208290526040808220805460ff191690555183917fd05bfc2250abb0f8fd265a54c53a24359c5484af63cad2e4ce87c78ab751395a91a25050565b5f826115e2868685611b50565b1495945050505050565b610b3e5f805160206121cd83398151915261151f565b5f807ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a0061049d565b611632611d7a565b6104b481611d9f565b7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c19930080546001600160a01b031981166001600160a01b03848116918217845560405192169182907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0905f90a3505050565b610b3e5f805160206121cd833981519152611a9a565b5f81116117105760405162461bcd60e51b815260206004820181905260248201527f4e6f6d696e614272696467653a20616d6f756e74206d757374206265203e20306044820152606401610632565b6001600160a01b0382166117665760405162461bcd60e51b815260206004820152601f60248201527f4e6f6d696e614272696467653a206e6f2062726964676520746f207a65726f006044820152606401610632565b5f805f9054906101000a90046001600160a01b03166001600160a01b031663110ff5f16040518163ffffffff1660e01b8152600401602060405180830381865afa1580156117b6573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906117da9190611fea565b6040516001600160a01b03808716602483015285166044820152606481018490529091505f9060840160408051601f198184030181529181526020820180516001600160e01b0316636ce5768960e11b1790525f549051632376548f60e21b81529192506001600160a01b031690638dd9523c9061186390859085906201388090600401612046565b602060405180830381865afa15801561187e573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906118a2919061207b565b3410156118f15760405162461bcd60e51b815260206004820152601e60248201527f4e6f6d696e614272696467653a20696e73756666696369656e742066656500006044820152606401610632565b6040516323b872dd60e01b81526001600160a01b038681166004830152306024830152604482018590527f000000000000000000000000000000000000000000000000000000000000000016906323b872dd906064016020604051808303815f875af1158015611963573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061198791906120ba565b6119d35760405162461bcd60e51b815260206004820152601d60248201527f4e6f6d696e614272696467653a207472616e73666572206661696c65640000006044820152606401610632565b5f5460405163c21dda4f60e01b81526001600160a01b039091169063c21dda4f903490611a18908690600490600262048789608a1b019088906201388090840161213f565b5f604051808303818588803b158015611a2f575f80fd5b505af1158015611a41573d5f803e3d5ffd5b5050505050836001600160a01b0316856001600160a01b03167f59bc8a913d49a9626dd6ba5def7fcf12804061c1bb9b8b6db077e1a12cb4b42285604051611a8b91815260200190565b60405180910390a35050505050565b5f8181527fff37105740f03695c8f3597f3aff2b92fbe1c80abea3c28731ecff2efd693400602081905260409091205460ff1615611b0d5760405162461bcd60e51b815260206004820152601060248201526f14185d5cd8589b194e881c185d5cd95960821b6044820152606401610632565b5f82815260208290526040808220805460ff191660011790555183917f0cb09dc71d57eeec2046f6854976717e4874a3cf2d6ddeddde337e5b6de6ba3191a25050565b805182515f9190611b628160016121a1565b8651611b6e90846121a1565b14611b8c57604051631a8a024960e11b815260040160405180910390fd5b5f816001600160401b03811115611ba557611ba5612092565b604051908082528060200260200182016040528015611bce578160200160208202803683370190505b5090505f805f805b85811015611cfe575f878510611c10578584611bf1816121b4565b955081518110611c0357611c036120a6565b6020026020010151611c36565b8985611c1b816121b4565b965081518110611c2d57611c2d6120a6565b60200260200101515b90505f8b8381518110611c4b57611c4b6120a6565b6020026020010151611c81578c84611c62816121b4565b955081518110611c7457611c746120a6565b6020026020010151611ccb565b888610611ca5578685611c93816121b4565b965081518110611c7457611c746120a6565b8a86611cb0816121b4565b975081518110611cc257611cc26120a6565b60200260200101515b9050611cd78282611da7565b878481518110611ce957611ce96120a6565b60209081029190910101525050600101611bd6565b508415611d505789518114611d2657604051631a8a024960e11b815260040160405180910390fd5b836001860381518110611d3b57611d3b6120a6565b602002602001015196505050505050506105e0565b8515611d6857875f81518110611d3b57611d3b6120a6565b895f81518110611d3b57611d3b6120a6565b611d82611dd0565b610b3e57604051631afcd79f60e31b815260040160405180910390fd5b6110e8611d7a565b5f818310611dc1575f8281526020849052604090206105e0565b505f9182526020526040902090565b5f611dd9611602565b54600160401b900460ff16919050565b5f60208284031215611df9575f80fd5b5035919050565b6001600160a01b03811681146104b4575f80fd5b5f805f60608486031215611e26575f80fd5b8335611e3181611e00565b92506020840135611e4181611e00565b929592945050506040919091013590565b5f8083601f840112611e62575f80fd5b5081356001600160401b03811115611e78575f80fd5b6020830191508360208260051b8501011115611e92575f80fd5b9250929050565b5f805f805f805f806080898b031215611eb0575f80fd5b88356001600160401b0380821115611ec6575f80fd5b611ed28c838d01611e52565b909a50985060208b0135915080821115611eea575f80fd5b611ef68c838d01611e52565b909850965060408b0135915080821115611f0e575f80fd5b611f1a8c838d01611e52565b909650945060608b0135915080821115611f32575f80fd5b50611f3f8b828c01611e52565b999c989b5096995094979396929594505050565b5f8060408385031215611f64575f80fd5b8235611f6f81611e00565b91506020830135611f7f81611e00565b809150509250929050565b5f8060408385031215611f9b575f80fd5b8235611fa681611e00565b946020939093013593505050565b5f60208284031215611fc4575f80fd5b81356105e081611e00565b80516001600160401b0381168114611fe5575f80fd5b919050565b5f60208284031215611ffa575f80fd5b6105e082611fcf565b5f81518084525f5b818110156120275760208185018101518683018201520161200b565b505f602082860101526020601f19601f83011685010191505092915050565b5f6001600160401b038086168352606060208401526120686060840186612003565b9150808416604084015250949350505050565b5f6020828403121561208b575f80fd5b5051919050565b634e487b7160e01b5f52604160045260245ffd5b634e487b7160e01b5f52603260045260245ffd5b5f602082840312156120ca575f80fd5b815180151581146105e0575f80fd5b5f604082840312156120e9575f80fd5b604051604081018181106001600160401b038211171561211757634e487b7160e01b5f52604160045260245ffd5b60405261212383611fcf565b8152602083015161213381611e00565b60208201529392505050565b5f6001600160401b03808816835260ff8716602084015260018060a01b038616604084015260a0606084015261217860a0840186612003565b91508084166080840152509695505050505050565b634e487b7160e01b5f52601160045260245ffd5b8082018082111561049d5761049d61218d565b5f600182016121c5576121c561218d565b506001019056fe76e8952e4b09b8d505aa08998d716721a1dbf0884ac74202e33985da1ed005e90683d1c283a672fc58eb7940a0dba83ea98b96966a9ca1b030dec2c60cea4d1e855511cc3694f64379908437d6d64458dc76d02482052bfb8a5b33a72c054c77a26469706673582212206d789df73bc9176d34950f7179226278843d8aaabc7c732ed4b117e2b37d93e064736f6c63430008180033",
}

// NominaBridgeL1ABI is the input ABI used to generate the binding from.
// Deprecated: Use NominaBridgeL1MetaData.ABI instead.
var NominaBridgeL1ABI = NominaBridgeL1MetaData.ABI

// NominaBridgeL1Bin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use NominaBridgeL1MetaData.Bin instead.
var NominaBridgeL1Bin = NominaBridgeL1MetaData.Bin

// DeployNominaBridgeL1 deploys a new Ethereum contract, binding an instance of NominaBridgeL1 to it.
func DeployNominaBridgeL1(auth *bind.TransactOpts, backend bind.ContractBackend, omni common.Address, nomina common.Address) (common.Address, *types.Transaction, *NominaBridgeL1, error) {
	parsed, err := NominaBridgeL1MetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(NominaBridgeL1Bin), backend, omni, nomina)
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

// NOMINA is a free data retrieval call binding the contract method 0x8ccc0d18.
//
// Solidity: function NOMINA() view returns(address)
func (_NominaBridgeL1 *NominaBridgeL1Caller) NOMINA(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NominaBridgeL1.contract.Call(opts, &out, "NOMINA")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NOMINA is a free data retrieval call binding the contract method 0x8ccc0d18.
//
// Solidity: function NOMINA() view returns(address)
func (_NominaBridgeL1 *NominaBridgeL1Session) NOMINA() (common.Address, error) {
	return _NominaBridgeL1.Contract.NOMINA(&_NominaBridgeL1.CallOpts)
}

// NOMINA is a free data retrieval call binding the contract method 0x8ccc0d18.
//
// Solidity: function NOMINA() view returns(address)
func (_NominaBridgeL1 *NominaBridgeL1CallerSession) NOMINA() (common.Address, error) {
	return _NominaBridgeL1.Contract.NOMINA(&_NominaBridgeL1.CallOpts)
}

// OMNI is a free data retrieval call binding the contract method 0x063bdf28.
//
// Solidity: function OMNI() view returns(address)
func (_NominaBridgeL1 *NominaBridgeL1Caller) OMNI(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NominaBridgeL1.contract.Call(opts, &out, "OMNI")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OMNI is a free data retrieval call binding the contract method 0x063bdf28.
//
// Solidity: function OMNI() view returns(address)
func (_NominaBridgeL1 *NominaBridgeL1Session) OMNI() (common.Address, error) {
	return _NominaBridgeL1.Contract.OMNI(&_NominaBridgeL1.CallOpts)
}

// OMNI is a free data retrieval call binding the contract method 0x063bdf28.
//
// Solidity: function OMNI() view returns(address)
func (_NominaBridgeL1 *NominaBridgeL1CallerSession) OMNI() (common.Address, error) {
	return _NominaBridgeL1.Contract.OMNI(&_NominaBridgeL1.CallOpts)
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

// PostHaltClaimed is a free data retrieval call binding the contract method 0xe849f0eb.
//
// Solidity: function postHaltClaimed(address ) view returns(bool)
func (_NominaBridgeL1 *NominaBridgeL1Caller) PostHaltClaimed(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _NominaBridgeL1.contract.Call(opts, &out, "postHaltClaimed", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// PostHaltClaimed is a free data retrieval call binding the contract method 0xe849f0eb.
//
// Solidity: function postHaltClaimed(address ) view returns(bool)
func (_NominaBridgeL1 *NominaBridgeL1Session) PostHaltClaimed(arg0 common.Address) (bool, error) {
	return _NominaBridgeL1.Contract.PostHaltClaimed(&_NominaBridgeL1.CallOpts, arg0)
}

// PostHaltClaimed is a free data retrieval call binding the contract method 0xe849f0eb.
//
// Solidity: function postHaltClaimed(address ) view returns(bool)
func (_NominaBridgeL1 *NominaBridgeL1CallerSession) PostHaltClaimed(arg0 common.Address) (bool, error) {
	return _NominaBridgeL1.Contract.PostHaltClaimed(&_NominaBridgeL1.CallOpts, arg0)
}

// PostHaltRoot is a free data retrieval call binding the contract method 0x1d822629.
//
// Solidity: function postHaltRoot() view returns(bytes32)
func (_NominaBridgeL1 *NominaBridgeL1Caller) PostHaltRoot(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _NominaBridgeL1.contract.Call(opts, &out, "postHaltRoot")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PostHaltRoot is a free data retrieval call binding the contract method 0x1d822629.
//
// Solidity: function postHaltRoot() view returns(bytes32)
func (_NominaBridgeL1 *NominaBridgeL1Session) PostHaltRoot() ([32]byte, error) {
	return _NominaBridgeL1.Contract.PostHaltRoot(&_NominaBridgeL1.CallOpts)
}

// PostHaltRoot is a free data retrieval call binding the contract method 0x1d822629.
//
// Solidity: function postHaltRoot() view returns(bytes32)
func (_NominaBridgeL1 *NominaBridgeL1CallerSession) PostHaltRoot() ([32]byte, error) {
	return _NominaBridgeL1.Contract.PostHaltRoot(&_NominaBridgeL1.CallOpts)
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

// InitializeV3 is a paid mutator transaction binding the contract method 0xe4cbdeed.
//
// Solidity: function initializeV3(bytes32 postHaltRoot_) returns()
func (_NominaBridgeL1 *NominaBridgeL1Transactor) InitializeV3(opts *bind.TransactOpts, postHaltRoot_ [32]byte) (*types.Transaction, error) {
	return _NominaBridgeL1.contract.Transact(opts, "initializeV3", postHaltRoot_)
}

// InitializeV3 is a paid mutator transaction binding the contract method 0xe4cbdeed.
//
// Solidity: function initializeV3(bytes32 postHaltRoot_) returns()
func (_NominaBridgeL1 *NominaBridgeL1Session) InitializeV3(postHaltRoot_ [32]byte) (*types.Transaction, error) {
	return _NominaBridgeL1.Contract.InitializeV3(&_NominaBridgeL1.TransactOpts, postHaltRoot_)
}

// InitializeV3 is a paid mutator transaction binding the contract method 0xe4cbdeed.
//
// Solidity: function initializeV3(bytes32 postHaltRoot_) returns()
func (_NominaBridgeL1 *NominaBridgeL1TransactorSession) InitializeV3(postHaltRoot_ [32]byte) (*types.Transaction, error) {
	return _NominaBridgeL1.Contract.InitializeV3(&_NominaBridgeL1.TransactOpts, postHaltRoot_)
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

// PostHaltWithdraw is a paid mutator transaction binding the contract method 0x3cb17480.
//
// Solidity: function postHaltWithdraw(address[] accounts, uint256[] amounts, bytes32[] multiProof, bool[] multiProofFlags) returns()
func (_NominaBridgeL1 *NominaBridgeL1Transactor) PostHaltWithdraw(opts *bind.TransactOpts, accounts []common.Address, amounts []*big.Int, multiProof [][32]byte, multiProofFlags []bool) (*types.Transaction, error) {
	return _NominaBridgeL1.contract.Transact(opts, "postHaltWithdraw", accounts, amounts, multiProof, multiProofFlags)
}

// PostHaltWithdraw is a paid mutator transaction binding the contract method 0x3cb17480.
//
// Solidity: function postHaltWithdraw(address[] accounts, uint256[] amounts, bytes32[] multiProof, bool[] multiProofFlags) returns()
func (_NominaBridgeL1 *NominaBridgeL1Session) PostHaltWithdraw(accounts []common.Address, amounts []*big.Int, multiProof [][32]byte, multiProofFlags []bool) (*types.Transaction, error) {
	return _NominaBridgeL1.Contract.PostHaltWithdraw(&_NominaBridgeL1.TransactOpts, accounts, amounts, multiProof, multiProofFlags)
}

// PostHaltWithdraw is a paid mutator transaction binding the contract method 0x3cb17480.
//
// Solidity: function postHaltWithdraw(address[] accounts, uint256[] amounts, bytes32[] multiProof, bool[] multiProofFlags) returns()
func (_NominaBridgeL1 *NominaBridgeL1TransactorSession) PostHaltWithdraw(accounts []common.Address, amounts []*big.Int, multiProof [][32]byte, multiProofFlags []bool) (*types.Transaction, error) {
	return _NominaBridgeL1.Contract.PostHaltWithdraw(&_NominaBridgeL1.TransactOpts, accounts, amounts, multiProof, multiProofFlags)
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

// NominaBridgeL1PostHaltWithdrawIterator is returned from FilterPostHaltWithdraw and is used to iterate over the raw logs and unpacked data for PostHaltWithdraw events raised by the NominaBridgeL1 contract.
type NominaBridgeL1PostHaltWithdrawIterator struct {
	Event *NominaBridgeL1PostHaltWithdraw // Event containing the contract specifics and raw log

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
func (it *NominaBridgeL1PostHaltWithdrawIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NominaBridgeL1PostHaltWithdraw)
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
		it.Event = new(NominaBridgeL1PostHaltWithdraw)
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
func (it *NominaBridgeL1PostHaltWithdrawIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NominaBridgeL1PostHaltWithdrawIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NominaBridgeL1PostHaltWithdraw represents a PostHaltWithdraw event raised by the NominaBridgeL1 contract.
type NominaBridgeL1PostHaltWithdraw struct {
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterPostHaltWithdraw is a free log retrieval operation binding the contract event 0x0c8025c30a7d66106b1ae8ed6f3dcbbf8d9ed725bb599f3bc855d11fa9203687.
//
// Solidity: event PostHaltWithdraw(address indexed to, uint256 amount)
func (_NominaBridgeL1 *NominaBridgeL1Filterer) FilterPostHaltWithdraw(opts *bind.FilterOpts, to []common.Address) (*NominaBridgeL1PostHaltWithdrawIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _NominaBridgeL1.contract.FilterLogs(opts, "PostHaltWithdraw", toRule)
	if err != nil {
		return nil, err
	}
	return &NominaBridgeL1PostHaltWithdrawIterator{contract: _NominaBridgeL1.contract, event: "PostHaltWithdraw", logs: logs, sub: sub}, nil
}

// WatchPostHaltWithdraw is a free log subscription operation binding the contract event 0x0c8025c30a7d66106b1ae8ed6f3dcbbf8d9ed725bb599f3bc855d11fa9203687.
//
// Solidity: event PostHaltWithdraw(address indexed to, uint256 amount)
func (_NominaBridgeL1 *NominaBridgeL1Filterer) WatchPostHaltWithdraw(opts *bind.WatchOpts, sink chan<- *NominaBridgeL1PostHaltWithdraw, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _NominaBridgeL1.contract.WatchLogs(opts, "PostHaltWithdraw", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NominaBridgeL1PostHaltWithdraw)
				if err := _NominaBridgeL1.contract.UnpackLog(event, "PostHaltWithdraw", log); err != nil {
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

// ParsePostHaltWithdraw is a log parse operation binding the contract event 0x0c8025c30a7d66106b1ae8ed6f3dcbbf8d9ed725bb599f3bc855d11fa9203687.
//
// Solidity: event PostHaltWithdraw(address indexed to, uint256 amount)
func (_NominaBridgeL1 *NominaBridgeL1Filterer) ParsePostHaltWithdraw(log types.Log) (*NominaBridgeL1PostHaltWithdraw, error) {
	event := new(NominaBridgeL1PostHaltWithdraw)
	if err := _NominaBridgeL1.contract.UnpackLog(event, "PostHaltWithdraw", log); err != nil {
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
