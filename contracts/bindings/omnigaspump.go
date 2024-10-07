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

// OmniGasPumpInitParams is an auto generated low-level Go binding around an user-defined struct.
type OmniGasPumpInitParams struct {
	GasStation common.Address
	Oracle     common.Address
	Portal     common.Address
	Owner      common.Address
	MaxSwap    *big.Int
	Toll       *big.Int
}

// OmniGasPumpMetaData contains all meta data concerning the OmniGasPump contract.
var OmniGasPumpMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"SETTLE_GAS\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"TOLL_DENOM\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"defaultConfLevel\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"dryFillUp\",\"inputs\":[{\"name\":\"amtETH\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"fillUp\",\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"gasStation\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"p\",\"type\":\"tuple\",\"internalType\":\"structOmniGasPump.InitParams\",\"components\":[{\"name\":\"gasStation\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"oracle\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"portal\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"maxSwap\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"toll\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"maxSwap\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"omni\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIOmniPortal\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"oracle\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIConversionRateOracle\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owed\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"quote\",\"inputs\":[{\"name\":\"amtOMNI\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setGasStation\",\"inputs\":[{\"name\":\"station\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setMaxSwap\",\"inputs\":[{\"name\":\"max\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setOracle\",\"inputs\":[{\"name\":\"oracle_\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setToll\",\"inputs\":[{\"name\":\"pct\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"toll\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"withdraw\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"xfee\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"DefaultConfLevelSet\",\"inputs\":[{\"name\":\"conf\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FilledUp\",\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"owed\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"amtETH\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"fee\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"toll\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"amtOMNI\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"GasStationSet\",\"inputs\":[{\"name\":\"station\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MaxSwapSet\",\"inputs\":[{\"name\":\"max\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OmniPortalSet\",\"inputs\":[{\"name\":\"omni\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OracleSet\",\"inputs\":[{\"name\":\"oracle\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TollSet\",\"inputs\":[{\"name\":\"pct\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"EnforcedPause\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ExpectedPause\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnableInvalidOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnableUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]}]",
	Bin: "0x608060405234801561001057600080fd5b5061001961001e565b6100d0565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000900460ff161561006e5760405163f92ee8a960e01b815260040160405180910390fd5b80546001600160401b03908116146100cd5780546001600160401b0319166001600160401b0390811782556040519081527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b50565b6117e5806100df6000396000f3fe6080604052600436106101665760003560e01c80637adbf973116100d1578063a67265b11161008a578063dbb602fd11610064578063dbb602fd14610440578063df18e04714610460578063ed1bd76c1461048d578063f2fde38b146104ad57600080fd5b8063a67265b1146103da578063c4918b4e146103fa578063ca48b20b1461041057600080fd5b80637adbf973146103135780637dc0d1d0146103335780638456cb59146103535780638aec67fe146103685780638da5cb5b1461037d578063a3dace5d146103ba57600080fd5b80634b260981116101235780634b2609811461023657806351cff8d91461024c57806355e0af6b1461026c5780635c975abb1461029b578063715018a6146102cb57806374eeb847146102e057600080fd5b806308a957a91461016b5780630e6e91d81461018d578063285aaa20146101ad57806339acf9f1146101d65780633f4ba83a1461020e5780634ae809ee14610223575b600080fd5b34801561017757600080fd5b5061018b610186366004611586565b6104cd565b005b34801561019957600080fd5b5061018b6101a83660046115af565b6104e1565b3480156101b957600080fd5b506101c360355481565b6040519081526020015b60405180910390f35b3480156101e257600080fd5b506000546101f6906001600160a01b031681565b6040516001600160a01b0390911681526020016101cd565b34801561021a57600080fd5b5061018b6104f2565b6101c3610231366004611586565b610504565b34801561024257600080fd5b506101c36103e881565b34801561025857600080fd5b5061018b610267366004611586565b61071c565b34801561027857600080fd5b5061028c6102873660046115af565b6107cb565b6040516101cd9392919061160e565b3480156102a757600080fd5b506000805160206117908339815191525460ff1660405190151581526020016101cd565b3480156102d757600080fd5b5061018b6108ad565b3480156102ec57600080fd5b5060005461030190600160a01b900460ff1681565b60405160ff90911681526020016101cd565b34801561031f57600080fd5b5061018b61032e366004611586565b6108bf565b34801561033f57600080fd5b506032546101f6906001600160a01b031681565b34801561035f57600080fd5b5061018b6108d0565b34801561037457600080fd5b506101c36108e0565b34801561038957600080fd5b507f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c199300546001600160a01b03166101f6565b3480156103c657600080fd5b5061018b6103d53660046115af565b61094e565b3480156103e657600080fd5b5061018b6103f5366004611638565b61095f565b34801561040657600080fd5b506101c360345481565b34801561041c57600080fd5b50610427620222e081565b60405167ffffffffffffffff90911681526020016101cd565b34801561044c57600080fd5b506033546101f6906001600160a01b031681565b34801561046c57600080fd5b506101c361047b366004611586565b60366020526000908152604090205481565b34801561049957600080fd5b506101c36104a83660046115af565b610adf565b3480156104b957600080fd5b5061018b6104c8366004611586565b610b2e565b6104d5610b69565b6104de81610bc4565b50565b6104e9610b69565b6104de81610c6f565b6104fa610b69565b610502610cec565b565b600061050e610d46565b60006105186108e0565b90508034101561056f5760405162461bcd60e51b815260206004820152601d60248201527f4f6d6e6947617350756d703a20696e73756666696369656e742066656500000060448201526064015b60405180910390fd5b600061057b8234611666565b90506034548111156105c75760405162461bcd60e51b815260206004820152601560248201527409edadcd28ec2e6a0eadae07440deeccae440dac2f605b1b6044820152606401610566565b60006103e8603554836105da9190611679565b6105e49190611690565b90506105f08183611666565b915060006105fd83610d77565b6001600160a01b03871660009081526036602052604081208054929350839290919061062a9084906116b2565b909155506106a7905061063b610e8d565b6033546001600160a01b038981166000818152603660205260409081902054905160248101929092526044820152600192919091169060640160408051601f198184030181529190526020810180516001600160e01b0316631decdcfb60e11b179052620222e0610f0a565b506001600160a01b038616600081815260366020908152604091829020548251908152349181019190915290810186905260608101849052608081018390527f7737fe59897f758714c24688a6470bb05235f01af1f4293edd0c290e651dd8319060a00160405180910390a295945050505050565b610724610b69565b6000816001600160a01b03164760405160006040518083038185875af1925050503d8060008114610771576040519150601f19603f3d011682016040523d82523d6000602084013e610776565b606091505b50509050806107c75760405162461bcd60e51b815260206004820152601c60248201527f4f6d6e6947617350756d703a207769746864726177206661696c6564000000006044820152606401610566565b5050565b600080606060006107da6108e0565b90508085101561081957505060408051808201909152601081526f696e73756666696369656e742066656560801b6020820152600092508291506108a6565b6108238186611666565b945060345485111561085c5750506040805180820190915260088152670deeccae440dac2f60c31b6020820152600092508291506108a6565b6103e86035548661086d9190611679565b6108779190611690565b6108819086611666565b945061088c85610d77565b600160405180602001604052806000815250935093509350505b9193909250565b6108b5610b69565b6105026000611054565b6108c7610b69565b6104de816110c5565b6108d8610b69565b610502611169565b60006001600160a01b036000196109476108f8610e8d565b6040516001600160a01b03851660248201526044810184905260640160408051601f198184030181529190526020810180516001600160e01b0316631decdcfb60e11b179052620222e06111b2565b9250505090565b610956610b69565b6104de81611230565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a008054600160401b810460ff16159067ffffffffffffffff166000811580156109a55750825b905060008267ffffffffffffffff1660011480156109c25750303b155b9050811580156109d0575080155b156109ee5760405163f92ee8a960e01b815260040160405180910390fd5b845467ffffffffffffffff191660011785558315610a1857845460ff60401b1916600160401b1785555b610a30610a2b6040880160208901611586565b6110c5565b610a45610a406020880188611586565b610bc4565b610a528660800135610c6f565b610a5f8660a00135611230565b610a79610a726060880160408901611586565b60016112b6565b610a91610a8c6080880160608901611586565b6112d0565b8315610ad757845460ff60401b19168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b505050505050565b600080610aeb836112e1565b90506035546103e8610afd9190611666565b610b096103e883611679565b610b139190611690565b9050610b1d6108e0565b610b2790826116b2565b9392505050565b610b36610b69565b6001600160a01b038116610b6057604051631e4fbdf760e01b815260006004820152602401610566565b6104de81611054565b33610b9b7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c199300546001600160a01b031690565b6001600160a01b0316146105025760405163118cdaa760e01b8152336004820152602401610566565b6001600160a01b038116610c1a5760405162461bcd60e51b815260206004820152601960248201527f4f6d6e6947617350756d703a207a65726f2061646472657373000000000000006044820152606401610566565b603380546001600160a01b0319166001600160a01b0383169081179091556040519081527ffd263e3b7583e8397be8a61710d1105cf8c0f111bbac1014d0ec7dbcd1e422f1906020015b60405180910390a150565b60008111610cb75760405162461bcd60e51b815260206004820152601560248201527409edadcd28ec2e6a0eadae07440f4cae4de40dac2f605b1b6044820152606401610566565b60348190556040518181527f343ecf9262f8cafd2e9b0ffdab9f14bf18a17899eeef3e41133e84c354e5298390602001610c64565b610cf46113ad565b600080516020611790833981519152805460ff191681557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa335b6040516001600160a01b039091168152602001610c64565b6000805160206117908339815191525460ff16156105025760405163d93c066560e01b815260040160405180910390fd5b6032546000906001600160a01b0316638b7bfd70610d93610e8d565b6040516001600160e01b031960e084901b16815267ffffffffffffffff9091166004820152602401602060405180830381865afa158015610dd8573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610dfc91906116c5565b603260009054906101000a90046001600160a01b03166001600160a01b0316638f9d6ace6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610e4f573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610e7391906116c5565b610e7d9084611679565b610e879190611690565b92915050565b60008060009054906101000a90046001600160a01b03166001600160a01b031663110ff5f16040518163ffffffff1660e01b8152600401602060405180830381865afa158015610ee1573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610f0591906116de565b905090565b60008054604051632376548f60e21b815282916001600160a01b031690638dd9523c90610f3f908a9088908890600401611708565b602060405180830381865afa158015610f5c573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610f8091906116c5565b90508047101580610f915750803410155b610fdd5760405162461bcd60e51b815260206004820152601860248201527f584170703a20696e73756666696369656e742066756e647300000000000000006044820152606401610566565b60005460405163c21dda4f60e01b81526001600160a01b039091169063c21dda4f908390611017908b908b908b908b908b9060040161173f565b6000604051808303818588803b15801561103057600080fd5b505af1158015611044573d6000803e3d6000fd5b50939a9950505050505050505050565b7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c19930080546001600160a01b031981166001600160a01b03848116918217845560405192169182907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a3505050565b6001600160a01b03811661111b5760405162461bcd60e51b815260206004820152601860248201527f4f6d6e6947617350756d703a207a65726f206f7261636c6500000000000000006044820152606401610566565b603280546001600160a01b0319166001600160a01b0383169081179091556040519081527f3f32684a32a11dabdbb8c0177de80aa3ae36a004d75210335b49e544e48cd0aa90602001610c64565b611171610d46565b600080516020611790833981519152805460ff191660011781557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a25833610d2e565b60008054604051632376548f60e21b81526001600160a01b0390911690638dd9523c906111e790879087908790600401611708565b602060405180830381865afa158015611204573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061122891906116c5565b949350505050565b6103e881106112815760405162461bcd60e51b815260206004820152601960248201527f4f6d6e6947617350756d703a2070637420746f6f2068696768000000000000006044820152606401610566565b60358190556040518181527f0b3d400288f60ce0f5632cd941b5748faa91ebea844cbe78c5180b7838a0933f90602001610c64565b6112be6113dd565b6112c782611426565b6107c7816114bf565b6112d86113dd565b6104de81611562565b603254604080516347ceb56760e11b815290516000926001600160a01b031691638f9d6ace9160048083019260209291908290030181865afa15801561132b573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061134f91906116c5565b6032546001600160a01b0316638b7bfd70611368610e8d565b6040516001600160e01b031960e084901b16815267ffffffffffffffff9091166004820152602401602060405180830381865afa158015610e4f573d6000803e3d6000fd5b6000805160206117908339815191525460ff1661050257604051638dfc202b60e01b815260040160405180910390fd5b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a0054600160401b900460ff1661050257604051631afcd79f60e31b815260040160405180910390fd5b6001600160a01b0381166114715760405162461bcd60e51b8152602060048201526012602482015271584170703a206e6f207a65726f206f6d6e6960701b6044820152606401610566565b600080546001600160a01b0319166001600160a01b0383169081179091556040519081527f79162c8d053a07e70cdc1ccc536f0440b571f8508377d2bef51094fadab98f4790602001610c64565b6114c88161156a565b6115145760405162461bcd60e51b815260206004820152601860248201527f584170703a20696e76616c696420636f6e66206c6576656c00000000000000006044820152606401610566565b6000805460ff60a01b1916600160a01b60ff8416908102919091179091556040519081527f8de08a798b4e50b4f351c1eaa91a11530043802be3ffac2df87db0c45a2e848390602001610c64565b610b366113dd565b600060ff821660011480610e87575060ff821660041492915050565b60006020828403121561159857600080fd5b81356001600160a01b0381168114610b2757600080fd5b6000602082840312156115c157600080fd5b5035919050565b6000815180845260005b818110156115ee576020818501810151868301820152016115d2565b506000602082860101526020601f19601f83011685010191505092915050565b838152821515602082015260606040820152600061162f60608301846115c8565b95945050505050565b600060c0828403121561164a57600080fd5b50919050565b634e487b7160e01b600052601160045260246000fd5b81810381811115610e8757610e87611650565b8082028115828204841417610e8757610e87611650565b6000826116ad57634e487b7160e01b600052601260045260246000fd5b500490565b80820180821115610e8757610e87611650565b6000602082840312156116d757600080fd5b5051919050565b6000602082840312156116f057600080fd5b815167ffffffffffffffff81168114610b2757600080fd5b600067ffffffffffffffff80861683526060602084015261172c60608401866115c8565b9150808416604084015250949350505050565b600067ffffffffffffffff808816835260ff8716602084015260018060a01b038616604084015260a0606084015261177a60a08401866115c8565b9150808416608084015250969550505050505056fecd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f03300a2646970667358221220cc3b1bdeb6e4f7efc1b982591c7d80bbccd23cf9ab913ec6b372fc0ae400303564736f6c63430008180033",
}

// OmniGasPumpABI is the input ABI used to generate the binding from.
// Deprecated: Use OmniGasPumpMetaData.ABI instead.
var OmniGasPumpABI = OmniGasPumpMetaData.ABI

// OmniGasPumpBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use OmniGasPumpMetaData.Bin instead.
var OmniGasPumpBin = OmniGasPumpMetaData.Bin

// DeployOmniGasPump deploys a new Ethereum contract, binding an instance of OmniGasPump to it.
func DeployOmniGasPump(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *OmniGasPump, error) {
	parsed, err := OmniGasPumpMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(OmniGasPumpBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &OmniGasPump{OmniGasPumpCaller: OmniGasPumpCaller{contract: contract}, OmniGasPumpTransactor: OmniGasPumpTransactor{contract: contract}, OmniGasPumpFilterer: OmniGasPumpFilterer{contract: contract}}, nil
}

// OmniGasPump is an auto generated Go binding around an Ethereum contract.
type OmniGasPump struct {
	OmniGasPumpCaller     // Read-only binding to the contract
	OmniGasPumpTransactor // Write-only binding to the contract
	OmniGasPumpFilterer   // Log filterer for contract events
}

// OmniGasPumpCaller is an auto generated read-only Go binding around an Ethereum contract.
type OmniGasPumpCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OmniGasPumpTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OmniGasPumpTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OmniGasPumpFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OmniGasPumpFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OmniGasPumpSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OmniGasPumpSession struct {
	Contract     *OmniGasPump      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OmniGasPumpCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OmniGasPumpCallerSession struct {
	Contract *OmniGasPumpCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// OmniGasPumpTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OmniGasPumpTransactorSession struct {
	Contract     *OmniGasPumpTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// OmniGasPumpRaw is an auto generated low-level Go binding around an Ethereum contract.
type OmniGasPumpRaw struct {
	Contract *OmniGasPump // Generic contract binding to access the raw methods on
}

// OmniGasPumpCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OmniGasPumpCallerRaw struct {
	Contract *OmniGasPumpCaller // Generic read-only contract binding to access the raw methods on
}

// OmniGasPumpTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OmniGasPumpTransactorRaw struct {
	Contract *OmniGasPumpTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOmniGasPump creates a new instance of OmniGasPump, bound to a specific deployed contract.
func NewOmniGasPump(address common.Address, backend bind.ContractBackend) (*OmniGasPump, error) {
	contract, err := bindOmniGasPump(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OmniGasPump{OmniGasPumpCaller: OmniGasPumpCaller{contract: contract}, OmniGasPumpTransactor: OmniGasPumpTransactor{contract: contract}, OmniGasPumpFilterer: OmniGasPumpFilterer{contract: contract}}, nil
}

// NewOmniGasPumpCaller creates a new read-only instance of OmniGasPump, bound to a specific deployed contract.
func NewOmniGasPumpCaller(address common.Address, caller bind.ContractCaller) (*OmniGasPumpCaller, error) {
	contract, err := bindOmniGasPump(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OmniGasPumpCaller{contract: contract}, nil
}

// NewOmniGasPumpTransactor creates a new write-only instance of OmniGasPump, bound to a specific deployed contract.
func NewOmniGasPumpTransactor(address common.Address, transactor bind.ContractTransactor) (*OmniGasPumpTransactor, error) {
	contract, err := bindOmniGasPump(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OmniGasPumpTransactor{contract: contract}, nil
}

// NewOmniGasPumpFilterer creates a new log filterer instance of OmniGasPump, bound to a specific deployed contract.
func NewOmniGasPumpFilterer(address common.Address, filterer bind.ContractFilterer) (*OmniGasPumpFilterer, error) {
	contract, err := bindOmniGasPump(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OmniGasPumpFilterer{contract: contract}, nil
}

// bindOmniGasPump binds a generic wrapper to an already deployed contract.
func bindOmniGasPump(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := OmniGasPumpMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OmniGasPump *OmniGasPumpRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OmniGasPump.Contract.OmniGasPumpCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OmniGasPump *OmniGasPumpRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OmniGasPump.Contract.OmniGasPumpTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OmniGasPump *OmniGasPumpRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OmniGasPump.Contract.OmniGasPumpTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OmniGasPump *OmniGasPumpCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OmniGasPump.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OmniGasPump *OmniGasPumpTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OmniGasPump.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OmniGasPump *OmniGasPumpTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OmniGasPump.Contract.contract.Transact(opts, method, params...)
}

// SETTLEGAS is a free data retrieval call binding the contract method 0xca48b20b.
//
// Solidity: function SETTLE_GAS() view returns(uint64)
func (_OmniGasPump *OmniGasPumpCaller) SETTLEGAS(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _OmniGasPump.contract.Call(opts, &out, "SETTLE_GAS")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// SETTLEGAS is a free data retrieval call binding the contract method 0xca48b20b.
//
// Solidity: function SETTLE_GAS() view returns(uint64)
func (_OmniGasPump *OmniGasPumpSession) SETTLEGAS() (uint64, error) {
	return _OmniGasPump.Contract.SETTLEGAS(&_OmniGasPump.CallOpts)
}

// SETTLEGAS is a free data retrieval call binding the contract method 0xca48b20b.
//
// Solidity: function SETTLE_GAS() view returns(uint64)
func (_OmniGasPump *OmniGasPumpCallerSession) SETTLEGAS() (uint64, error) {
	return _OmniGasPump.Contract.SETTLEGAS(&_OmniGasPump.CallOpts)
}

// TOLLDENOM is a free data retrieval call binding the contract method 0x4b260981.
//
// Solidity: function TOLL_DENOM() view returns(uint256)
func (_OmniGasPump *OmniGasPumpCaller) TOLLDENOM(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OmniGasPump.contract.Call(opts, &out, "TOLL_DENOM")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TOLLDENOM is a free data retrieval call binding the contract method 0x4b260981.
//
// Solidity: function TOLL_DENOM() view returns(uint256)
func (_OmniGasPump *OmniGasPumpSession) TOLLDENOM() (*big.Int, error) {
	return _OmniGasPump.Contract.TOLLDENOM(&_OmniGasPump.CallOpts)
}

// TOLLDENOM is a free data retrieval call binding the contract method 0x4b260981.
//
// Solidity: function TOLL_DENOM() view returns(uint256)
func (_OmniGasPump *OmniGasPumpCallerSession) TOLLDENOM() (*big.Int, error) {
	return _OmniGasPump.Contract.TOLLDENOM(&_OmniGasPump.CallOpts)
}

// DefaultConfLevel is a free data retrieval call binding the contract method 0x74eeb847.
//
// Solidity: function defaultConfLevel() view returns(uint8)
func (_OmniGasPump *OmniGasPumpCaller) DefaultConfLevel(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _OmniGasPump.contract.Call(opts, &out, "defaultConfLevel")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// DefaultConfLevel is a free data retrieval call binding the contract method 0x74eeb847.
//
// Solidity: function defaultConfLevel() view returns(uint8)
func (_OmniGasPump *OmniGasPumpSession) DefaultConfLevel() (uint8, error) {
	return _OmniGasPump.Contract.DefaultConfLevel(&_OmniGasPump.CallOpts)
}

// DefaultConfLevel is a free data retrieval call binding the contract method 0x74eeb847.
//
// Solidity: function defaultConfLevel() view returns(uint8)
func (_OmniGasPump *OmniGasPumpCallerSession) DefaultConfLevel() (uint8, error) {
	return _OmniGasPump.Contract.DefaultConfLevel(&_OmniGasPump.CallOpts)
}

// DryFillUp is a free data retrieval call binding the contract method 0x55e0af6b.
//
// Solidity: function dryFillUp(uint256 amtETH) view returns(uint256, bool, string)
func (_OmniGasPump *OmniGasPumpCaller) DryFillUp(opts *bind.CallOpts, amtETH *big.Int) (*big.Int, bool, string, error) {
	var out []interface{}
	err := _OmniGasPump.contract.Call(opts, &out, "dryFillUp", amtETH)

	if err != nil {
		return *new(*big.Int), *new(bool), *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(bool)).(*bool)
	out2 := *abi.ConvertType(out[2], new(string)).(*string)

	return out0, out1, out2, err

}

// DryFillUp is a free data retrieval call binding the contract method 0x55e0af6b.
//
// Solidity: function dryFillUp(uint256 amtETH) view returns(uint256, bool, string)
func (_OmniGasPump *OmniGasPumpSession) DryFillUp(amtETH *big.Int) (*big.Int, bool, string, error) {
	return _OmniGasPump.Contract.DryFillUp(&_OmniGasPump.CallOpts, amtETH)
}

// DryFillUp is a free data retrieval call binding the contract method 0x55e0af6b.
//
// Solidity: function dryFillUp(uint256 amtETH) view returns(uint256, bool, string)
func (_OmniGasPump *OmniGasPumpCallerSession) DryFillUp(amtETH *big.Int) (*big.Int, bool, string, error) {
	return _OmniGasPump.Contract.DryFillUp(&_OmniGasPump.CallOpts, amtETH)
}

// GasStation is a free data retrieval call binding the contract method 0xdbb602fd.
//
// Solidity: function gasStation() view returns(address)
func (_OmniGasPump *OmniGasPumpCaller) GasStation(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OmniGasPump.contract.Call(opts, &out, "gasStation")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GasStation is a free data retrieval call binding the contract method 0xdbb602fd.
//
// Solidity: function gasStation() view returns(address)
func (_OmniGasPump *OmniGasPumpSession) GasStation() (common.Address, error) {
	return _OmniGasPump.Contract.GasStation(&_OmniGasPump.CallOpts)
}

// GasStation is a free data retrieval call binding the contract method 0xdbb602fd.
//
// Solidity: function gasStation() view returns(address)
func (_OmniGasPump *OmniGasPumpCallerSession) GasStation() (common.Address, error) {
	return _OmniGasPump.Contract.GasStation(&_OmniGasPump.CallOpts)
}

// MaxSwap is a free data retrieval call binding the contract method 0xc4918b4e.
//
// Solidity: function maxSwap() view returns(uint256)
func (_OmniGasPump *OmniGasPumpCaller) MaxSwap(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OmniGasPump.contract.Call(opts, &out, "maxSwap")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MaxSwap is a free data retrieval call binding the contract method 0xc4918b4e.
//
// Solidity: function maxSwap() view returns(uint256)
func (_OmniGasPump *OmniGasPumpSession) MaxSwap() (*big.Int, error) {
	return _OmniGasPump.Contract.MaxSwap(&_OmniGasPump.CallOpts)
}

// MaxSwap is a free data retrieval call binding the contract method 0xc4918b4e.
//
// Solidity: function maxSwap() view returns(uint256)
func (_OmniGasPump *OmniGasPumpCallerSession) MaxSwap() (*big.Int, error) {
	return _OmniGasPump.Contract.MaxSwap(&_OmniGasPump.CallOpts)
}

// Omni is a free data retrieval call binding the contract method 0x39acf9f1.
//
// Solidity: function omni() view returns(address)
func (_OmniGasPump *OmniGasPumpCaller) Omni(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OmniGasPump.contract.Call(opts, &out, "omni")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Omni is a free data retrieval call binding the contract method 0x39acf9f1.
//
// Solidity: function omni() view returns(address)
func (_OmniGasPump *OmniGasPumpSession) Omni() (common.Address, error) {
	return _OmniGasPump.Contract.Omni(&_OmniGasPump.CallOpts)
}

// Omni is a free data retrieval call binding the contract method 0x39acf9f1.
//
// Solidity: function omni() view returns(address)
func (_OmniGasPump *OmniGasPumpCallerSession) Omni() (common.Address, error) {
	return _OmniGasPump.Contract.Omni(&_OmniGasPump.CallOpts)
}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() view returns(address)
func (_OmniGasPump *OmniGasPumpCaller) Oracle(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OmniGasPump.contract.Call(opts, &out, "oracle")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() view returns(address)
func (_OmniGasPump *OmniGasPumpSession) Oracle() (common.Address, error) {
	return _OmniGasPump.Contract.Oracle(&_OmniGasPump.CallOpts)
}

// Oracle is a free data retrieval call binding the contract method 0x7dc0d1d0.
//
// Solidity: function oracle() view returns(address)
func (_OmniGasPump *OmniGasPumpCallerSession) Oracle() (common.Address, error) {
	return _OmniGasPump.Contract.Oracle(&_OmniGasPump.CallOpts)
}

// Owed is a free data retrieval call binding the contract method 0xdf18e047.
//
// Solidity: function owed(address ) view returns(uint256)
func (_OmniGasPump *OmniGasPumpCaller) Owed(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _OmniGasPump.contract.Call(opts, &out, "owed", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Owed is a free data retrieval call binding the contract method 0xdf18e047.
//
// Solidity: function owed(address ) view returns(uint256)
func (_OmniGasPump *OmniGasPumpSession) Owed(arg0 common.Address) (*big.Int, error) {
	return _OmniGasPump.Contract.Owed(&_OmniGasPump.CallOpts, arg0)
}

// Owed is a free data retrieval call binding the contract method 0xdf18e047.
//
// Solidity: function owed(address ) view returns(uint256)
func (_OmniGasPump *OmniGasPumpCallerSession) Owed(arg0 common.Address) (*big.Int, error) {
	return _OmniGasPump.Contract.Owed(&_OmniGasPump.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_OmniGasPump *OmniGasPumpCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OmniGasPump.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_OmniGasPump *OmniGasPumpSession) Owner() (common.Address, error) {
	return _OmniGasPump.Contract.Owner(&_OmniGasPump.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_OmniGasPump *OmniGasPumpCallerSession) Owner() (common.Address, error) {
	return _OmniGasPump.Contract.Owner(&_OmniGasPump.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_OmniGasPump *OmniGasPumpCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _OmniGasPump.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_OmniGasPump *OmniGasPumpSession) Paused() (bool, error) {
	return _OmniGasPump.Contract.Paused(&_OmniGasPump.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_OmniGasPump *OmniGasPumpCallerSession) Paused() (bool, error) {
	return _OmniGasPump.Contract.Paused(&_OmniGasPump.CallOpts)
}

// Quote is a free data retrieval call binding the contract method 0xed1bd76c.
//
// Solidity: function quote(uint256 amtOMNI) view returns(uint256)
func (_OmniGasPump *OmniGasPumpCaller) Quote(opts *bind.CallOpts, amtOMNI *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _OmniGasPump.contract.Call(opts, &out, "quote", amtOMNI)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Quote is a free data retrieval call binding the contract method 0xed1bd76c.
//
// Solidity: function quote(uint256 amtOMNI) view returns(uint256)
func (_OmniGasPump *OmniGasPumpSession) Quote(amtOMNI *big.Int) (*big.Int, error) {
	return _OmniGasPump.Contract.Quote(&_OmniGasPump.CallOpts, amtOMNI)
}

// Quote is a free data retrieval call binding the contract method 0xed1bd76c.
//
// Solidity: function quote(uint256 amtOMNI) view returns(uint256)
func (_OmniGasPump *OmniGasPumpCallerSession) Quote(amtOMNI *big.Int) (*big.Int, error) {
	return _OmniGasPump.Contract.Quote(&_OmniGasPump.CallOpts, amtOMNI)
}

// Toll is a free data retrieval call binding the contract method 0x285aaa20.
//
// Solidity: function toll() view returns(uint256)
func (_OmniGasPump *OmniGasPumpCaller) Toll(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OmniGasPump.contract.Call(opts, &out, "toll")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Toll is a free data retrieval call binding the contract method 0x285aaa20.
//
// Solidity: function toll() view returns(uint256)
func (_OmniGasPump *OmniGasPumpSession) Toll() (*big.Int, error) {
	return _OmniGasPump.Contract.Toll(&_OmniGasPump.CallOpts)
}

// Toll is a free data retrieval call binding the contract method 0x285aaa20.
//
// Solidity: function toll() view returns(uint256)
func (_OmniGasPump *OmniGasPumpCallerSession) Toll() (*big.Int, error) {
	return _OmniGasPump.Contract.Toll(&_OmniGasPump.CallOpts)
}

// Xfee is a free data retrieval call binding the contract method 0x8aec67fe.
//
// Solidity: function xfee() view returns(uint256)
func (_OmniGasPump *OmniGasPumpCaller) Xfee(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _OmniGasPump.contract.Call(opts, &out, "xfee")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Xfee is a free data retrieval call binding the contract method 0x8aec67fe.
//
// Solidity: function xfee() view returns(uint256)
func (_OmniGasPump *OmniGasPumpSession) Xfee() (*big.Int, error) {
	return _OmniGasPump.Contract.Xfee(&_OmniGasPump.CallOpts)
}

// Xfee is a free data retrieval call binding the contract method 0x8aec67fe.
//
// Solidity: function xfee() view returns(uint256)
func (_OmniGasPump *OmniGasPumpCallerSession) Xfee() (*big.Int, error) {
	return _OmniGasPump.Contract.Xfee(&_OmniGasPump.CallOpts)
}

// FillUp is a paid mutator transaction binding the contract method 0x4ae809ee.
//
// Solidity: function fillUp(address recipient) payable returns(uint256)
func (_OmniGasPump *OmniGasPumpTransactor) FillUp(opts *bind.TransactOpts, recipient common.Address) (*types.Transaction, error) {
	return _OmniGasPump.contract.Transact(opts, "fillUp", recipient)
}

// FillUp is a paid mutator transaction binding the contract method 0x4ae809ee.
//
// Solidity: function fillUp(address recipient) payable returns(uint256)
func (_OmniGasPump *OmniGasPumpSession) FillUp(recipient common.Address) (*types.Transaction, error) {
	return _OmniGasPump.Contract.FillUp(&_OmniGasPump.TransactOpts, recipient)
}

// FillUp is a paid mutator transaction binding the contract method 0x4ae809ee.
//
// Solidity: function fillUp(address recipient) payable returns(uint256)
func (_OmniGasPump *OmniGasPumpTransactorSession) FillUp(recipient common.Address) (*types.Transaction, error) {
	return _OmniGasPump.Contract.FillUp(&_OmniGasPump.TransactOpts, recipient)
}

// Initialize is a paid mutator transaction binding the contract method 0xa67265b1.
//
// Solidity: function initialize((address,address,address,address,uint256,uint256) p) returns()
func (_OmniGasPump *OmniGasPumpTransactor) Initialize(opts *bind.TransactOpts, p OmniGasPumpInitParams) (*types.Transaction, error) {
	return _OmniGasPump.contract.Transact(opts, "initialize", p)
}

// Initialize is a paid mutator transaction binding the contract method 0xa67265b1.
//
// Solidity: function initialize((address,address,address,address,uint256,uint256) p) returns()
func (_OmniGasPump *OmniGasPumpSession) Initialize(p OmniGasPumpInitParams) (*types.Transaction, error) {
	return _OmniGasPump.Contract.Initialize(&_OmniGasPump.TransactOpts, p)
}

// Initialize is a paid mutator transaction binding the contract method 0xa67265b1.
//
// Solidity: function initialize((address,address,address,address,uint256,uint256) p) returns()
func (_OmniGasPump *OmniGasPumpTransactorSession) Initialize(p OmniGasPumpInitParams) (*types.Transaction, error) {
	return _OmniGasPump.Contract.Initialize(&_OmniGasPump.TransactOpts, p)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_OmniGasPump *OmniGasPumpTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OmniGasPump.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_OmniGasPump *OmniGasPumpSession) Pause() (*types.Transaction, error) {
	return _OmniGasPump.Contract.Pause(&_OmniGasPump.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_OmniGasPump *OmniGasPumpTransactorSession) Pause() (*types.Transaction, error) {
	return _OmniGasPump.Contract.Pause(&_OmniGasPump.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_OmniGasPump *OmniGasPumpTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OmniGasPump.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_OmniGasPump *OmniGasPumpSession) RenounceOwnership() (*types.Transaction, error) {
	return _OmniGasPump.Contract.RenounceOwnership(&_OmniGasPump.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_OmniGasPump *OmniGasPumpTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _OmniGasPump.Contract.RenounceOwnership(&_OmniGasPump.TransactOpts)
}

// SetGasStation is a paid mutator transaction binding the contract method 0x08a957a9.
//
// Solidity: function setGasStation(address station) returns()
func (_OmniGasPump *OmniGasPumpTransactor) SetGasStation(opts *bind.TransactOpts, station common.Address) (*types.Transaction, error) {
	return _OmniGasPump.contract.Transact(opts, "setGasStation", station)
}

// SetGasStation is a paid mutator transaction binding the contract method 0x08a957a9.
//
// Solidity: function setGasStation(address station) returns()
func (_OmniGasPump *OmniGasPumpSession) SetGasStation(station common.Address) (*types.Transaction, error) {
	return _OmniGasPump.Contract.SetGasStation(&_OmniGasPump.TransactOpts, station)
}

// SetGasStation is a paid mutator transaction binding the contract method 0x08a957a9.
//
// Solidity: function setGasStation(address station) returns()
func (_OmniGasPump *OmniGasPumpTransactorSession) SetGasStation(station common.Address) (*types.Transaction, error) {
	return _OmniGasPump.Contract.SetGasStation(&_OmniGasPump.TransactOpts, station)
}

// SetMaxSwap is a paid mutator transaction binding the contract method 0x0e6e91d8.
//
// Solidity: function setMaxSwap(uint256 max) returns()
func (_OmniGasPump *OmniGasPumpTransactor) SetMaxSwap(opts *bind.TransactOpts, max *big.Int) (*types.Transaction, error) {
	return _OmniGasPump.contract.Transact(opts, "setMaxSwap", max)
}

// SetMaxSwap is a paid mutator transaction binding the contract method 0x0e6e91d8.
//
// Solidity: function setMaxSwap(uint256 max) returns()
func (_OmniGasPump *OmniGasPumpSession) SetMaxSwap(max *big.Int) (*types.Transaction, error) {
	return _OmniGasPump.Contract.SetMaxSwap(&_OmniGasPump.TransactOpts, max)
}

// SetMaxSwap is a paid mutator transaction binding the contract method 0x0e6e91d8.
//
// Solidity: function setMaxSwap(uint256 max) returns()
func (_OmniGasPump *OmniGasPumpTransactorSession) SetMaxSwap(max *big.Int) (*types.Transaction, error) {
	return _OmniGasPump.Contract.SetMaxSwap(&_OmniGasPump.TransactOpts, max)
}

// SetOracle is a paid mutator transaction binding the contract method 0x7adbf973.
//
// Solidity: function setOracle(address oracle_) returns()
func (_OmniGasPump *OmniGasPumpTransactor) SetOracle(opts *bind.TransactOpts, oracle_ common.Address) (*types.Transaction, error) {
	return _OmniGasPump.contract.Transact(opts, "setOracle", oracle_)
}

// SetOracle is a paid mutator transaction binding the contract method 0x7adbf973.
//
// Solidity: function setOracle(address oracle_) returns()
func (_OmniGasPump *OmniGasPumpSession) SetOracle(oracle_ common.Address) (*types.Transaction, error) {
	return _OmniGasPump.Contract.SetOracle(&_OmniGasPump.TransactOpts, oracle_)
}

// SetOracle is a paid mutator transaction binding the contract method 0x7adbf973.
//
// Solidity: function setOracle(address oracle_) returns()
func (_OmniGasPump *OmniGasPumpTransactorSession) SetOracle(oracle_ common.Address) (*types.Transaction, error) {
	return _OmniGasPump.Contract.SetOracle(&_OmniGasPump.TransactOpts, oracle_)
}

// SetToll is a paid mutator transaction binding the contract method 0xa3dace5d.
//
// Solidity: function setToll(uint256 pct) returns()
func (_OmniGasPump *OmniGasPumpTransactor) SetToll(opts *bind.TransactOpts, pct *big.Int) (*types.Transaction, error) {
	return _OmniGasPump.contract.Transact(opts, "setToll", pct)
}

// SetToll is a paid mutator transaction binding the contract method 0xa3dace5d.
//
// Solidity: function setToll(uint256 pct) returns()
func (_OmniGasPump *OmniGasPumpSession) SetToll(pct *big.Int) (*types.Transaction, error) {
	return _OmniGasPump.Contract.SetToll(&_OmniGasPump.TransactOpts, pct)
}

// SetToll is a paid mutator transaction binding the contract method 0xa3dace5d.
//
// Solidity: function setToll(uint256 pct) returns()
func (_OmniGasPump *OmniGasPumpTransactorSession) SetToll(pct *big.Int) (*types.Transaction, error) {
	return _OmniGasPump.Contract.SetToll(&_OmniGasPump.TransactOpts, pct)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_OmniGasPump *OmniGasPumpTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _OmniGasPump.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_OmniGasPump *OmniGasPumpSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _OmniGasPump.Contract.TransferOwnership(&_OmniGasPump.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_OmniGasPump *OmniGasPumpTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _OmniGasPump.Contract.TransferOwnership(&_OmniGasPump.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_OmniGasPump *OmniGasPumpTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OmniGasPump.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_OmniGasPump *OmniGasPumpSession) Unpause() (*types.Transaction, error) {
	return _OmniGasPump.Contract.Unpause(&_OmniGasPump.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_OmniGasPump *OmniGasPumpTransactorSession) Unpause() (*types.Transaction, error) {
	return _OmniGasPump.Contract.Unpause(&_OmniGasPump.TransactOpts)
}

// Withdraw is a paid mutator transaction binding the contract method 0x51cff8d9.
//
// Solidity: function withdraw(address to) returns()
func (_OmniGasPump *OmniGasPumpTransactor) Withdraw(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _OmniGasPump.contract.Transact(opts, "withdraw", to)
}

// Withdraw is a paid mutator transaction binding the contract method 0x51cff8d9.
//
// Solidity: function withdraw(address to) returns()
func (_OmniGasPump *OmniGasPumpSession) Withdraw(to common.Address) (*types.Transaction, error) {
	return _OmniGasPump.Contract.Withdraw(&_OmniGasPump.TransactOpts, to)
}

// Withdraw is a paid mutator transaction binding the contract method 0x51cff8d9.
//
// Solidity: function withdraw(address to) returns()
func (_OmniGasPump *OmniGasPumpTransactorSession) Withdraw(to common.Address) (*types.Transaction, error) {
	return _OmniGasPump.Contract.Withdraw(&_OmniGasPump.TransactOpts, to)
}

// OmniGasPumpDefaultConfLevelSetIterator is returned from FilterDefaultConfLevelSet and is used to iterate over the raw logs and unpacked data for DefaultConfLevelSet events raised by the OmniGasPump contract.
type OmniGasPumpDefaultConfLevelSetIterator struct {
	Event *OmniGasPumpDefaultConfLevelSet // Event containing the contract specifics and raw log

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
func (it *OmniGasPumpDefaultConfLevelSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OmniGasPumpDefaultConfLevelSet)
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
		it.Event = new(OmniGasPumpDefaultConfLevelSet)
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
func (it *OmniGasPumpDefaultConfLevelSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OmniGasPumpDefaultConfLevelSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OmniGasPumpDefaultConfLevelSet represents a DefaultConfLevelSet event raised by the OmniGasPump contract.
type OmniGasPumpDefaultConfLevelSet struct {
	Conf uint8
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterDefaultConfLevelSet is a free log retrieval operation binding the contract event 0x8de08a798b4e50b4f351c1eaa91a11530043802be3ffac2df87db0c45a2e8483.
//
// Solidity: event DefaultConfLevelSet(uint8 conf)
func (_OmniGasPump *OmniGasPumpFilterer) FilterDefaultConfLevelSet(opts *bind.FilterOpts) (*OmniGasPumpDefaultConfLevelSetIterator, error) {

	logs, sub, err := _OmniGasPump.contract.FilterLogs(opts, "DefaultConfLevelSet")
	if err != nil {
		return nil, err
	}
	return &OmniGasPumpDefaultConfLevelSetIterator{contract: _OmniGasPump.contract, event: "DefaultConfLevelSet", logs: logs, sub: sub}, nil
}

// WatchDefaultConfLevelSet is a free log subscription operation binding the contract event 0x8de08a798b4e50b4f351c1eaa91a11530043802be3ffac2df87db0c45a2e8483.
//
// Solidity: event DefaultConfLevelSet(uint8 conf)
func (_OmniGasPump *OmniGasPumpFilterer) WatchDefaultConfLevelSet(opts *bind.WatchOpts, sink chan<- *OmniGasPumpDefaultConfLevelSet) (event.Subscription, error) {

	logs, sub, err := _OmniGasPump.contract.WatchLogs(opts, "DefaultConfLevelSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OmniGasPumpDefaultConfLevelSet)
				if err := _OmniGasPump.contract.UnpackLog(event, "DefaultConfLevelSet", log); err != nil {
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
func (_OmniGasPump *OmniGasPumpFilterer) ParseDefaultConfLevelSet(log types.Log) (*OmniGasPumpDefaultConfLevelSet, error) {
	event := new(OmniGasPumpDefaultConfLevelSet)
	if err := _OmniGasPump.contract.UnpackLog(event, "DefaultConfLevelSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OmniGasPumpFilledUpIterator is returned from FilterFilledUp and is used to iterate over the raw logs and unpacked data for FilledUp events raised by the OmniGasPump contract.
type OmniGasPumpFilledUpIterator struct {
	Event *OmniGasPumpFilledUp // Event containing the contract specifics and raw log

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
func (it *OmniGasPumpFilledUpIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OmniGasPumpFilledUp)
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
		it.Event = new(OmniGasPumpFilledUp)
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
func (it *OmniGasPumpFilledUpIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OmniGasPumpFilledUpIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OmniGasPumpFilledUp represents a FilledUp event raised by the OmniGasPump contract.
type OmniGasPumpFilledUp struct {
	Recipient common.Address
	Owed      *big.Int
	AmtETH    *big.Int
	Fee       *big.Int
	Toll      *big.Int
	AmtOMNI   *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterFilledUp is a free log retrieval operation binding the contract event 0x7737fe59897f758714c24688a6470bb05235f01af1f4293edd0c290e651dd831.
//
// Solidity: event FilledUp(address indexed recipient, uint256 owed, uint256 amtETH, uint256 fee, uint256 toll, uint256 amtOMNI)
func (_OmniGasPump *OmniGasPumpFilterer) FilterFilledUp(opts *bind.FilterOpts, recipient []common.Address) (*OmniGasPumpFilledUpIterator, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _OmniGasPump.contract.FilterLogs(opts, "FilledUp", recipientRule)
	if err != nil {
		return nil, err
	}
	return &OmniGasPumpFilledUpIterator{contract: _OmniGasPump.contract, event: "FilledUp", logs: logs, sub: sub}, nil
}

// WatchFilledUp is a free log subscription operation binding the contract event 0x7737fe59897f758714c24688a6470bb05235f01af1f4293edd0c290e651dd831.
//
// Solidity: event FilledUp(address indexed recipient, uint256 owed, uint256 amtETH, uint256 fee, uint256 toll, uint256 amtOMNI)
func (_OmniGasPump *OmniGasPumpFilterer) WatchFilledUp(opts *bind.WatchOpts, sink chan<- *OmniGasPumpFilledUp, recipient []common.Address) (event.Subscription, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _OmniGasPump.contract.WatchLogs(opts, "FilledUp", recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OmniGasPumpFilledUp)
				if err := _OmniGasPump.contract.UnpackLog(event, "FilledUp", log); err != nil {
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

// ParseFilledUp is a log parse operation binding the contract event 0x7737fe59897f758714c24688a6470bb05235f01af1f4293edd0c290e651dd831.
//
// Solidity: event FilledUp(address indexed recipient, uint256 owed, uint256 amtETH, uint256 fee, uint256 toll, uint256 amtOMNI)
func (_OmniGasPump *OmniGasPumpFilterer) ParseFilledUp(log types.Log) (*OmniGasPumpFilledUp, error) {
	event := new(OmniGasPumpFilledUp)
	if err := _OmniGasPump.contract.UnpackLog(event, "FilledUp", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OmniGasPumpGasStationSetIterator is returned from FilterGasStationSet and is used to iterate over the raw logs and unpacked data for GasStationSet events raised by the OmniGasPump contract.
type OmniGasPumpGasStationSetIterator struct {
	Event *OmniGasPumpGasStationSet // Event containing the contract specifics and raw log

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
func (it *OmniGasPumpGasStationSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OmniGasPumpGasStationSet)
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
		it.Event = new(OmniGasPumpGasStationSet)
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
func (it *OmniGasPumpGasStationSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OmniGasPumpGasStationSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OmniGasPumpGasStationSet represents a GasStationSet event raised by the OmniGasPump contract.
type OmniGasPumpGasStationSet struct {
	Station common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterGasStationSet is a free log retrieval operation binding the contract event 0xfd263e3b7583e8397be8a61710d1105cf8c0f111bbac1014d0ec7dbcd1e422f1.
//
// Solidity: event GasStationSet(address station)
func (_OmniGasPump *OmniGasPumpFilterer) FilterGasStationSet(opts *bind.FilterOpts) (*OmniGasPumpGasStationSetIterator, error) {

	logs, sub, err := _OmniGasPump.contract.FilterLogs(opts, "GasStationSet")
	if err != nil {
		return nil, err
	}
	return &OmniGasPumpGasStationSetIterator{contract: _OmniGasPump.contract, event: "GasStationSet", logs: logs, sub: sub}, nil
}

// WatchGasStationSet is a free log subscription operation binding the contract event 0xfd263e3b7583e8397be8a61710d1105cf8c0f111bbac1014d0ec7dbcd1e422f1.
//
// Solidity: event GasStationSet(address station)
func (_OmniGasPump *OmniGasPumpFilterer) WatchGasStationSet(opts *bind.WatchOpts, sink chan<- *OmniGasPumpGasStationSet) (event.Subscription, error) {

	logs, sub, err := _OmniGasPump.contract.WatchLogs(opts, "GasStationSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OmniGasPumpGasStationSet)
				if err := _OmniGasPump.contract.UnpackLog(event, "GasStationSet", log); err != nil {
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

// ParseGasStationSet is a log parse operation binding the contract event 0xfd263e3b7583e8397be8a61710d1105cf8c0f111bbac1014d0ec7dbcd1e422f1.
//
// Solidity: event GasStationSet(address station)
func (_OmniGasPump *OmniGasPumpFilterer) ParseGasStationSet(log types.Log) (*OmniGasPumpGasStationSet, error) {
	event := new(OmniGasPumpGasStationSet)
	if err := _OmniGasPump.contract.UnpackLog(event, "GasStationSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OmniGasPumpInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the OmniGasPump contract.
type OmniGasPumpInitializedIterator struct {
	Event *OmniGasPumpInitialized // Event containing the contract specifics and raw log

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
func (it *OmniGasPumpInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OmniGasPumpInitialized)
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
		it.Event = new(OmniGasPumpInitialized)
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
func (it *OmniGasPumpInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OmniGasPumpInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OmniGasPumpInitialized represents a Initialized event raised by the OmniGasPump contract.
type OmniGasPumpInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_OmniGasPump *OmniGasPumpFilterer) FilterInitialized(opts *bind.FilterOpts) (*OmniGasPumpInitializedIterator, error) {

	logs, sub, err := _OmniGasPump.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &OmniGasPumpInitializedIterator{contract: _OmniGasPump.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_OmniGasPump *OmniGasPumpFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *OmniGasPumpInitialized) (event.Subscription, error) {

	logs, sub, err := _OmniGasPump.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OmniGasPumpInitialized)
				if err := _OmniGasPump.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_OmniGasPump *OmniGasPumpFilterer) ParseInitialized(log types.Log) (*OmniGasPumpInitialized, error) {
	event := new(OmniGasPumpInitialized)
	if err := _OmniGasPump.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OmniGasPumpMaxSwapSetIterator is returned from FilterMaxSwapSet and is used to iterate over the raw logs and unpacked data for MaxSwapSet events raised by the OmniGasPump contract.
type OmniGasPumpMaxSwapSetIterator struct {
	Event *OmniGasPumpMaxSwapSet // Event containing the contract specifics and raw log

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
func (it *OmniGasPumpMaxSwapSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OmniGasPumpMaxSwapSet)
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
		it.Event = new(OmniGasPumpMaxSwapSet)
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
func (it *OmniGasPumpMaxSwapSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OmniGasPumpMaxSwapSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OmniGasPumpMaxSwapSet represents a MaxSwapSet event raised by the OmniGasPump contract.
type OmniGasPumpMaxSwapSet struct {
	Max *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterMaxSwapSet is a free log retrieval operation binding the contract event 0x343ecf9262f8cafd2e9b0ffdab9f14bf18a17899eeef3e41133e84c354e52983.
//
// Solidity: event MaxSwapSet(uint256 max)
func (_OmniGasPump *OmniGasPumpFilterer) FilterMaxSwapSet(opts *bind.FilterOpts) (*OmniGasPumpMaxSwapSetIterator, error) {

	logs, sub, err := _OmniGasPump.contract.FilterLogs(opts, "MaxSwapSet")
	if err != nil {
		return nil, err
	}
	return &OmniGasPumpMaxSwapSetIterator{contract: _OmniGasPump.contract, event: "MaxSwapSet", logs: logs, sub: sub}, nil
}

// WatchMaxSwapSet is a free log subscription operation binding the contract event 0x343ecf9262f8cafd2e9b0ffdab9f14bf18a17899eeef3e41133e84c354e52983.
//
// Solidity: event MaxSwapSet(uint256 max)
func (_OmniGasPump *OmniGasPumpFilterer) WatchMaxSwapSet(opts *bind.WatchOpts, sink chan<- *OmniGasPumpMaxSwapSet) (event.Subscription, error) {

	logs, sub, err := _OmniGasPump.contract.WatchLogs(opts, "MaxSwapSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OmniGasPumpMaxSwapSet)
				if err := _OmniGasPump.contract.UnpackLog(event, "MaxSwapSet", log); err != nil {
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

// ParseMaxSwapSet is a log parse operation binding the contract event 0x343ecf9262f8cafd2e9b0ffdab9f14bf18a17899eeef3e41133e84c354e52983.
//
// Solidity: event MaxSwapSet(uint256 max)
func (_OmniGasPump *OmniGasPumpFilterer) ParseMaxSwapSet(log types.Log) (*OmniGasPumpMaxSwapSet, error) {
	event := new(OmniGasPumpMaxSwapSet)
	if err := _OmniGasPump.contract.UnpackLog(event, "MaxSwapSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OmniGasPumpOmniPortalSetIterator is returned from FilterOmniPortalSet and is used to iterate over the raw logs and unpacked data for OmniPortalSet events raised by the OmniGasPump contract.
type OmniGasPumpOmniPortalSetIterator struct {
	Event *OmniGasPumpOmniPortalSet // Event containing the contract specifics and raw log

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
func (it *OmniGasPumpOmniPortalSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OmniGasPumpOmniPortalSet)
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
		it.Event = new(OmniGasPumpOmniPortalSet)
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
func (it *OmniGasPumpOmniPortalSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OmniGasPumpOmniPortalSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OmniGasPumpOmniPortalSet represents a OmniPortalSet event raised by the OmniGasPump contract.
type OmniGasPumpOmniPortalSet struct {
	Omni common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterOmniPortalSet is a free log retrieval operation binding the contract event 0x79162c8d053a07e70cdc1ccc536f0440b571f8508377d2bef51094fadab98f47.
//
// Solidity: event OmniPortalSet(address omni)
func (_OmniGasPump *OmniGasPumpFilterer) FilterOmniPortalSet(opts *bind.FilterOpts) (*OmniGasPumpOmniPortalSetIterator, error) {

	logs, sub, err := _OmniGasPump.contract.FilterLogs(opts, "OmniPortalSet")
	if err != nil {
		return nil, err
	}
	return &OmniGasPumpOmniPortalSetIterator{contract: _OmniGasPump.contract, event: "OmniPortalSet", logs: logs, sub: sub}, nil
}

// WatchOmniPortalSet is a free log subscription operation binding the contract event 0x79162c8d053a07e70cdc1ccc536f0440b571f8508377d2bef51094fadab98f47.
//
// Solidity: event OmniPortalSet(address omni)
func (_OmniGasPump *OmniGasPumpFilterer) WatchOmniPortalSet(opts *bind.WatchOpts, sink chan<- *OmniGasPumpOmniPortalSet) (event.Subscription, error) {

	logs, sub, err := _OmniGasPump.contract.WatchLogs(opts, "OmniPortalSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OmniGasPumpOmniPortalSet)
				if err := _OmniGasPump.contract.UnpackLog(event, "OmniPortalSet", log); err != nil {
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
func (_OmniGasPump *OmniGasPumpFilterer) ParseOmniPortalSet(log types.Log) (*OmniGasPumpOmniPortalSet, error) {
	event := new(OmniGasPumpOmniPortalSet)
	if err := _OmniGasPump.contract.UnpackLog(event, "OmniPortalSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OmniGasPumpOracleSetIterator is returned from FilterOracleSet and is used to iterate over the raw logs and unpacked data for OracleSet events raised by the OmniGasPump contract.
type OmniGasPumpOracleSetIterator struct {
	Event *OmniGasPumpOracleSet // Event containing the contract specifics and raw log

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
func (it *OmniGasPumpOracleSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OmniGasPumpOracleSet)
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
		it.Event = new(OmniGasPumpOracleSet)
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
func (it *OmniGasPumpOracleSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OmniGasPumpOracleSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OmniGasPumpOracleSet represents a OracleSet event raised by the OmniGasPump contract.
type OmniGasPumpOracleSet struct {
	Oracle common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterOracleSet is a free log retrieval operation binding the contract event 0x3f32684a32a11dabdbb8c0177de80aa3ae36a004d75210335b49e544e48cd0aa.
//
// Solidity: event OracleSet(address oracle)
func (_OmniGasPump *OmniGasPumpFilterer) FilterOracleSet(opts *bind.FilterOpts) (*OmniGasPumpOracleSetIterator, error) {

	logs, sub, err := _OmniGasPump.contract.FilterLogs(opts, "OracleSet")
	if err != nil {
		return nil, err
	}
	return &OmniGasPumpOracleSetIterator{contract: _OmniGasPump.contract, event: "OracleSet", logs: logs, sub: sub}, nil
}

// WatchOracleSet is a free log subscription operation binding the contract event 0x3f32684a32a11dabdbb8c0177de80aa3ae36a004d75210335b49e544e48cd0aa.
//
// Solidity: event OracleSet(address oracle)
func (_OmniGasPump *OmniGasPumpFilterer) WatchOracleSet(opts *bind.WatchOpts, sink chan<- *OmniGasPumpOracleSet) (event.Subscription, error) {

	logs, sub, err := _OmniGasPump.contract.WatchLogs(opts, "OracleSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OmniGasPumpOracleSet)
				if err := _OmniGasPump.contract.UnpackLog(event, "OracleSet", log); err != nil {
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

// ParseOracleSet is a log parse operation binding the contract event 0x3f32684a32a11dabdbb8c0177de80aa3ae36a004d75210335b49e544e48cd0aa.
//
// Solidity: event OracleSet(address oracle)
func (_OmniGasPump *OmniGasPumpFilterer) ParseOracleSet(log types.Log) (*OmniGasPumpOracleSet, error) {
	event := new(OmniGasPumpOracleSet)
	if err := _OmniGasPump.contract.UnpackLog(event, "OracleSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OmniGasPumpOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the OmniGasPump contract.
type OmniGasPumpOwnershipTransferredIterator struct {
	Event *OmniGasPumpOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *OmniGasPumpOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OmniGasPumpOwnershipTransferred)
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
		it.Event = new(OmniGasPumpOwnershipTransferred)
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
func (it *OmniGasPumpOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OmniGasPumpOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OmniGasPumpOwnershipTransferred represents a OwnershipTransferred event raised by the OmniGasPump contract.
type OmniGasPumpOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_OmniGasPump *OmniGasPumpFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*OmniGasPumpOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _OmniGasPump.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &OmniGasPumpOwnershipTransferredIterator{contract: _OmniGasPump.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_OmniGasPump *OmniGasPumpFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *OmniGasPumpOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _OmniGasPump.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OmniGasPumpOwnershipTransferred)
				if err := _OmniGasPump.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_OmniGasPump *OmniGasPumpFilterer) ParseOwnershipTransferred(log types.Log) (*OmniGasPumpOwnershipTransferred, error) {
	event := new(OmniGasPumpOwnershipTransferred)
	if err := _OmniGasPump.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OmniGasPumpPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the OmniGasPump contract.
type OmniGasPumpPausedIterator struct {
	Event *OmniGasPumpPaused // Event containing the contract specifics and raw log

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
func (it *OmniGasPumpPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OmniGasPumpPaused)
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
		it.Event = new(OmniGasPumpPaused)
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
func (it *OmniGasPumpPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OmniGasPumpPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OmniGasPumpPaused represents a Paused event raised by the OmniGasPump contract.
type OmniGasPumpPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_OmniGasPump *OmniGasPumpFilterer) FilterPaused(opts *bind.FilterOpts) (*OmniGasPumpPausedIterator, error) {

	logs, sub, err := _OmniGasPump.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &OmniGasPumpPausedIterator{contract: _OmniGasPump.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_OmniGasPump *OmniGasPumpFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *OmniGasPumpPaused) (event.Subscription, error) {

	logs, sub, err := _OmniGasPump.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OmniGasPumpPaused)
				if err := _OmniGasPump.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_OmniGasPump *OmniGasPumpFilterer) ParsePaused(log types.Log) (*OmniGasPumpPaused, error) {
	event := new(OmniGasPumpPaused)
	if err := _OmniGasPump.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OmniGasPumpTollSetIterator is returned from FilterTollSet and is used to iterate over the raw logs and unpacked data for TollSet events raised by the OmniGasPump contract.
type OmniGasPumpTollSetIterator struct {
	Event *OmniGasPumpTollSet // Event containing the contract specifics and raw log

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
func (it *OmniGasPumpTollSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OmniGasPumpTollSet)
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
		it.Event = new(OmniGasPumpTollSet)
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
func (it *OmniGasPumpTollSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OmniGasPumpTollSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OmniGasPumpTollSet represents a TollSet event raised by the OmniGasPump contract.
type OmniGasPumpTollSet struct {
	Pct *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterTollSet is a free log retrieval operation binding the contract event 0x0b3d400288f60ce0f5632cd941b5748faa91ebea844cbe78c5180b7838a0933f.
//
// Solidity: event TollSet(uint256 pct)
func (_OmniGasPump *OmniGasPumpFilterer) FilterTollSet(opts *bind.FilterOpts) (*OmniGasPumpTollSetIterator, error) {

	logs, sub, err := _OmniGasPump.contract.FilterLogs(opts, "TollSet")
	if err != nil {
		return nil, err
	}
	return &OmniGasPumpTollSetIterator{contract: _OmniGasPump.contract, event: "TollSet", logs: logs, sub: sub}, nil
}

// WatchTollSet is a free log subscription operation binding the contract event 0x0b3d400288f60ce0f5632cd941b5748faa91ebea844cbe78c5180b7838a0933f.
//
// Solidity: event TollSet(uint256 pct)
func (_OmniGasPump *OmniGasPumpFilterer) WatchTollSet(opts *bind.WatchOpts, sink chan<- *OmniGasPumpTollSet) (event.Subscription, error) {

	logs, sub, err := _OmniGasPump.contract.WatchLogs(opts, "TollSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OmniGasPumpTollSet)
				if err := _OmniGasPump.contract.UnpackLog(event, "TollSet", log); err != nil {
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

// ParseTollSet is a log parse operation binding the contract event 0x0b3d400288f60ce0f5632cd941b5748faa91ebea844cbe78c5180b7838a0933f.
//
// Solidity: event TollSet(uint256 pct)
func (_OmniGasPump *OmniGasPumpFilterer) ParseTollSet(log types.Log) (*OmniGasPumpTollSet, error) {
	event := new(OmniGasPumpTollSet)
	if err := _OmniGasPump.contract.UnpackLog(event, "TollSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OmniGasPumpUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the OmniGasPump contract.
type OmniGasPumpUnpausedIterator struct {
	Event *OmniGasPumpUnpaused // Event containing the contract specifics and raw log

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
func (it *OmniGasPumpUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OmniGasPumpUnpaused)
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
		it.Event = new(OmniGasPumpUnpaused)
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
func (it *OmniGasPumpUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OmniGasPumpUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OmniGasPumpUnpaused represents a Unpaused event raised by the OmniGasPump contract.
type OmniGasPumpUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_OmniGasPump *OmniGasPumpFilterer) FilterUnpaused(opts *bind.FilterOpts) (*OmniGasPumpUnpausedIterator, error) {

	logs, sub, err := _OmniGasPump.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &OmniGasPumpUnpausedIterator{contract: _OmniGasPump.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_OmniGasPump *OmniGasPumpFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *OmniGasPumpUnpaused) (event.Subscription, error) {

	logs, sub, err := _OmniGasPump.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OmniGasPumpUnpaused)
				if err := _OmniGasPump.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_OmniGasPump *OmniGasPumpFilterer) ParseUnpaused(log types.Log) (*OmniGasPumpUnpaused, error) {
	event := new(OmniGasPumpUnpaused)
	if err := _OmniGasPump.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
