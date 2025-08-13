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
	Bin: "0x608060405234801561000f575f80fd5b5061001861001d565b6100cf565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000900460ff161561006d5760405163f92ee8a960e01b815260040160405180910390fd5b80546001600160401b03908116146100cc5780546001600160401b0319166001600160401b0390811782556040519081527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b50565b611812806100dc5f395ff3fe608060405260043610610161575f3560e01c80637adbf973116100cd578063a67265b111610087578063dbb602fd11610062578063dbb602fd14610423578063df18e04714610442578063ed1bd76c1461046d578063f2fde38b1461048c575f80fd5b8063a67265b1146103c0578063c4918b4e146103df578063ca48b20b146103f4575f80fd5b80637adbf973146102ff5780637dc0d1d01461031e5780638456cb591461033d5780638aec67fe146103515780638da5cb5b14610365578063a3dace5d146103a1575f80fd5b80634b2609811161011e5780634b2609811461022a57806351cff8d91461023f57806355e0af6b1461025e5780635c975abb1461028c578063715018a6146102ba57806374eeb847146102ce575f80fd5b806308a957a9146101655780630e6e91d814610186578063285aaa20146101a557806339acf9f1146101cd5780633f4ba83a146102035780634ae809ee14610217575b5f80fd5b348015610170575f80fd5b5061018461017f3660046115ca565b6104ab565b005b348015610191575f80fd5b506101846101a03660046115f0565b6104bf565b3480156101b0575f80fd5b506101ba60355481565b6040519081526020015b60405180910390f35b3480156101d8575f80fd5b505f546101eb906001600160a01b031681565b6040516001600160a01b0390911681526020016101c4565b34801561020e575f80fd5b506101846104d0565b6101ba6102253660046115ca565b6104e2565b348015610235575f80fd5b506101ba6103e881565b34801561024a575f80fd5b506101846102593660046115ca565b610744565b348015610269575f80fd5b5061027d6102783660046115f0565b610841565b6040516101c49392919061164a565b348015610297575f80fd5b505f805160206117bd8339815191525460ff1660405190151581526020016101c4565b3480156102c5575f80fd5b5061018461091e565b3480156102d9575f80fd5b505f546102ed90600160a01b900460ff1681565b60405160ff90911681526020016101c4565b34801561030a575f80fd5b506101846103193660046115ca565b61092f565b348015610329575f80fd5b506032546101eb906001600160a01b031681565b348015610348575f80fd5b50610184610940565b34801561035c575f80fd5b506101ba610950565b348015610370575f80fd5b507f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c199300546001600160a01b03166101eb565b3480156103ac575f80fd5b506101846103bb3660046115f0565b6109bc565b3480156103cb575f80fd5b506101846103da366004611673565b6109cd565b3480156103ea575f80fd5b506101ba60345481565b3480156103ff575f80fd5b5061040a620222e081565b60405167ffffffffffffffff90911681526020016101c4565b34801561042e575f80fd5b506033546101eb906001600160a01b031681565b34801561044d575f80fd5b506101ba61045c3660046115ca565b60366020525f908152604090205481565b348015610478575f80fd5b506101ba6104873660046115f0565b610b36565b348015610497575f80fd5b506101846104a63660046115ca565b610b84565b6104b3610bbe565b6104bc81610c19565b50565b6104c7610bbe565b6104bc81610cc4565b6104d8610bbe565b6104e0610d40565b565b5f6104eb610d99565b6001600160a01b0382166105425760405162461bcd60e51b815260206004820152601960248201527827b6b734a3b0b9a83ab6b81d103737903d32b9379030b2323960391b60448201526064015b60405180910390fd5b5f61054b610950565b90508034101561059d5760405162461bcd60e51b815260206004820152601d60248201527f4f6d6e6947617350756d703a20696e73756666696369656e74206665650000006044820152606401610539565b5f6105a8823461169d565b90506034548111156105f45760405162461bcd60e51b815260206004820152601560248201527409edadcd28ec2e6a0eadae07440deeccae440dac2f605b1b6044820152606401610539565b5f6103e86035548361060691906116b0565b61061091906116c7565b905061061c818361169d565b91505f61062883610dc9565b6001600160a01b0387165f908152603660205260408120805492935083929091906106549084906116e6565b909155506106d09050610665610ed9565b6033546001600160a01b038981165f818152603660205260409081902054905160248101929092526044820152600192919091169060640160408051601f198184030181529190526020810180516001600160e01b0316631decdcfb60e11b179052620222e0610f52565b506001600160a01b0386165f81815260366020908152604091829020548251908152349181019190915290810186905260608101849052608081018390527f7737fe59897f758714c24688a6470bb05235f01af1f4293edd0c290e651dd8319060a00160405180910390a295945050505050565b61074c610bbe565b6001600160a01b03811661079e5760405162461bcd60e51b815260206004820152601960248201527827b6b734a3b0b9a83ab6b81d103737903d32b9379030b2323960391b6044820152606401610539565b5f816001600160a01b0316476040515f6040518083038185875af1925050503d805f81146107e7576040519150601f19603f3d011682016040523d82523d5f602084013e6107ec565b606091505b505090508061083d5760405162461bcd60e51b815260206004820152601c60248201527f4f6d6e6947617350756d703a207769746864726177206661696c6564000000006044820152606401610539565b5050565b5f8060605f61084e610950565b90508085101561088c57505060408051808201909152601081526f696e73756666696369656e742066656560801b60208201525f9250829150610917565b610896818661169d565b94506034548511156108ce5750506040805180820190915260088152670deeccae440dac2f60c31b60208201525f9250829150610917565b6103e8603554866108df91906116b0565b6108e991906116c7565b6108f3908661169d565b94506108fe85610dc9565b600160405180602001604052805f815250935093509350505b9193909250565b610926610bbe565b6104e05f611089565b610937610bbe565b6104bc816110f9565b610948610bbe565b6104e061119d565b5f6001600160a01b035f196109b5610966610ed9565b6040516001600160a01b03851660248201526044810184905260640160408051601f198184030181529190526020810180516001600160e01b0316631decdcfb60e11b179052620222e06111e5565b9250505090565b6109c4610bbe565b6104bc81611260565b5f6109d66112e6565b805490915060ff600160401b820416159067ffffffffffffffff165f811580156109fd5750825b90505f8267ffffffffffffffff166001148015610a195750303b155b905081158015610a27575080155b15610a455760405163f92ee8a960e01b815260040160405180910390fd5b845467ffffffffffffffff191660011785558315610a6f57845460ff60401b1916600160401b1785555b610a87610a8260408801602089016115ca565b6110f9565b610a9c610a9760208801886115ca565b610c19565b610aa98660800135610cc4565b610ab68660a00135611260565b610ad0610ac960608801604089016115ca565b600161130e565b610ae8610ae360808801606089016115ca565b611328565b8315610b2e57845460ff60401b19168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b505050505050565b5f80610b4183611339565b90506035546103e8610b53919061169d565b610b5f6103e8836116b0565b610b6991906116c7565b9050610b73610950565b610b7d90826116e6565b9392505050565b610b8c610bbe565b6001600160a01b038116610bb557604051631e4fbdf760e01b81525f6004820152602401610539565b6104bc81611089565b33610bf07f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c199300546001600160a01b031690565b6001600160a01b0316146104e05760405163118cdaa760e01b8152336004820152602401610539565b6001600160a01b038116610c6f5760405162461bcd60e51b815260206004820152601960248201527f4f6d6e6947617350756d703a207a65726f2061646472657373000000000000006044820152606401610539565b603380546001600160a01b0319166001600160a01b0383169081179091556040519081527ffd263e3b7583e8397be8a61710d1105cf8c0f111bbac1014d0ec7dbcd1e422f1906020015b60405180910390a150565b5f8111610d0b5760405162461bcd60e51b815260206004820152601560248201527409edadcd28ec2e6a0eadae07440f4cae4de40dac2f605b1b6044820152606401610539565b60348190556040518181527f343ecf9262f8cafd2e9b0ffdab9f14bf18a17899eeef3e41133e84c354e5298390602001610cb9565b610d48611400565b5f805160206117bd833981519152805460ff191681557f5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa335b6040516001600160a01b039091168152602001610cb9565b5f805160206117bd8339815191525460ff16156104e05760405163d93c066560e01b815260040160405180910390fd5b6032545f906001600160a01b0316638b7bfd70610de4610ed9565b6040516001600160e01b031960e084901b16815267ffffffffffffffff9091166004820152602401602060405180830381865afa158015610e27573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610e4b91906116f9565b60325f9054906101000a90046001600160a01b03166001600160a01b0316638f9d6ace6040518163ffffffff1660e01b8152600401602060405180830381865afa158015610e9b573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610ebf91906116f9565b610ec990846116b0565b610ed391906116c7565b92915050565b5f805f9054906101000a90046001600160a01b03166001600160a01b031663110ff5f16040518163ffffffff1660e01b8152600401602060405180830381865afa158015610f29573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610f4d9190611710565b905090565b5f8054604051632376548f60e21b815282916001600160a01b031690638dd9523c90610f86908a9088908890600401611737565b602060405180830381865afa158015610fa1573d5f803e3d5ffd5b505050506040513d601f19601f82011682018060405250810190610fc591906116f9565b9050804710156110175760405162461bcd60e51b815260206004820152601860248201527f584170703a20696e73756666696369656e742066756e647300000000000000006044820152606401610539565b5f5460405163c21dda4f60e01b81526001600160a01b039091169063c21dda4f908390611050908b908b908b908b908b9060040161176d565b5f604051808303818588803b158015611067575f80fd5b505af1158015611079573d5f803e3d5ffd5b50939a9950505050505050505050565b7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c19930080546001600160a01b031981166001600160a01b03848116918217845560405192169182907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0905f90a3505050565b6001600160a01b03811661114f5760405162461bcd60e51b815260206004820152601860248201527f4f6d6e6947617350756d703a207a65726f206f7261636c6500000000000000006044820152606401610539565b603280546001600160a01b0319166001600160a01b0383169081179091556040519081527f3f32684a32a11dabdbb8c0177de80aa3ae36a004d75210335b49e544e48cd0aa90602001610cb9565b6111a5610d99565b5f805160206117bd833981519152805460ff191660011781557f62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a25833610d81565b5f8054604051632376548f60e21b81526001600160a01b0390911690638dd9523c9061121990879087908790600401611737565b602060405180830381865afa158015611234573d5f803e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061125891906116f9565b949350505050565b6103e881106112b15760405162461bcd60e51b815260206004820152601960248201527f4f6d6e6947617350756d703a2070637420746f6f2068696768000000000000006044820152606401610539565b60358190556040518181527f0b3d400288f60ce0f5632cd941b5748faa91ebea844cbe78c5180b7838a0933f90602001610cb9565b5f807ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00610ed3565b61131661142f565b61131f82611454565b61083d816114ec565b61133061142f565b6104bc8161158e565b603254604080516347ceb56760e11b815290515f926001600160a01b031691638f9d6ace9160048083019260209291908290030181865afa158015611380573d5f803e3d5ffd5b505050506040513d601f19601f820116820180604052508101906113a491906116f9565b6032546001600160a01b0316638b7bfd706113bd610ed9565b6040516001600160e01b031960e084901b16815267ffffffffffffffff9091166004820152602401602060405180830381865afa158015610e9b573d5f803e3d5ffd5b5f805160206117bd8339815191525460ff166104e057604051638dfc202b60e01b815260040160405180910390fd5b611437611596565b6104e057604051631afcd79f60e31b815260040160405180910390fd5b6001600160a01b03811661149f5760405162461bcd60e51b8152602060048201526012602482015271584170703a206e6f207a65726f206f6d6e6960701b6044820152606401610539565b5f80546001600160a01b0319166001600160a01b0383169081179091556040519081527f79162c8d053a07e70cdc1ccc536f0440b571f8508377d2bef51094fadab98f4790602001610cb9565b6114f5816115af565b6115415760405162461bcd60e51b815260206004820152601860248201527f584170703a20696e76616c696420636f6e66206c6576656c00000000000000006044820152606401610539565b5f805460ff60a01b1916600160a01b60ff8416908102919091179091556040519081527f8de08a798b4e50b4f351c1eaa91a11530043802be3ffac2df87db0c45a2e848390602001610cb9565b610b8c61142f565b5f61159f6112e6565b54600160401b900460ff16919050565b5f60ff821660011480610ed3575060ff821660041492915050565b5f602082840312156115da575f80fd5b81356001600160a01b0381168114610b7d575f80fd5b5f60208284031215611600575f80fd5b5035919050565b5f81518084525f5b8181101561162b5760208185018101518683018201520161160f565b505f602082860101526020601f19601f83011685010191505092915050565b8381528215156020820152606060408201525f61166a6060830184611607565b95945050505050565b5f60c08284031215611683575f80fd5b50919050565b634e487b7160e01b5f52601160045260245ffd5b81810381811115610ed357610ed3611689565b8082028115828204841417610ed357610ed3611689565b5f826116e157634e487b7160e01b5f52601260045260245ffd5b500490565b80820180821115610ed357610ed3611689565b5f60208284031215611709575f80fd5b5051919050565b5f60208284031215611720575f80fd5b815167ffffffffffffffff81168114610b7d575f80fd5b5f67ffffffffffffffff80861683526060602084015261175a6060840186611607565b9150808416604084015250949350505050565b5f67ffffffffffffffff808816835260ff8716602084015260018060a01b038616604084015260a060608401526117a760a0840186611607565b9150808416608084015250969550505050505056fecd5ed15c6e187e77e9aee88184c21f4f2182ab5827cb3b7e07fbedcd63f03300a26469706673582212200b46a416aee6ac204fae71db01909cc217135e004ea605eb3f54e4d3f203f34c64736f6c63430008180033",
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
