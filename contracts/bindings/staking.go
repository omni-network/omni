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

// StakingMetaData contains all meta data concerning the Staking contract.
var StakingMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"MinDelegation\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MinDeposit\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"allowValidators\",\"inputs\":[{\"name\":\"validators\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"createValidator\",\"inputs\":[{\"name\":\"pubkey\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"createValidator\",\"inputs\":[{\"name\":\"pubkey\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"delegate\",\"inputs\":[{\"name\":\"validator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"delegateFor\",\"inputs\":[{\"name\":\"delegator\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"validator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"disableAllowlist\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"disallowValidators\",\"inputs\":[{\"name\":\"validators\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"eip712Domain\",\"inputs\":[],\"outputs\":[{\"name\":\"fields\",\"type\":\"bytes1\",\"internalType\":\"bytes1\"},{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"version\",\"type\":\"string\",\"internalType\":\"string\"},{\"name\":\"chainId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"verifyingContract\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"salt\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"extensions\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"enableAllowlist\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getConsPubkeyDigest\",\"inputs\":[{\"name\":\"validator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"owner_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"isAllowlistEnabled_\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"initializeV1\",\"inputs\":[{\"name\":\"owner_\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"isAllowlistEnabled_\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"initializeV2\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isAllowedValidator\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isAllowlistEnabled\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"AllowlistDisabled\",\"inputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"AllowlistEnabled\",\"inputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CreateValidator\",\"inputs\":[{\"name\":\"validator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"pubkey\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"deposit\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Delegate\",\"inputs\":[{\"name\":\"delegator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"validator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"EIP712DomainChanged\",\"inputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ValidatorAllowed\",\"inputs\":[{\"name\":\"validator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ValidatorDisallowed\",\"inputs\":[{\"name\":\"validator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnableInvalidOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnableUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]}]",
	Bin: "0x608060405234801561001057600080fd5b5061001961001e565b6100d0565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000900460ff161561006e5760405163f92ee8a960e01b815260040160405180910390fd5b80546001600160401b03908116146100cd5780546001600160401b0319166001600160401b0390811782556040519081527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b50565b611d44806100df6000396000f3fe60806040526004361061011f5760003560e01c806384768b7a116100a0578063b62ac12c11610064578063b62ac12c14610317578063c6a2aac814610337578063cf8e629a1461034c578063d146fd1b14610361578063f2fde38b1461037b57600080fd5b806384768b7a1461023557806384b0196e146102755780638da5cb5b1461029d5780638f38fae8146102e4578063a5a470ad1461030457600080fd5b806359bcddde116100e757806359bcddde146101c95780635c19a95c146101e55780635cd8a76b146101f8578063715018a61461020d57806378fcbe5b1461022257600080fd5b8063117407e31461012457806311bcd830146101465780631fa10bcc146101765780633f0b1edf14610189578063400ada75146101a9575b600080fd5b34801561013057600080fd5b5061014461013f36600461178c565b61039b565b005b34801561015257600080fd5b5061016368056bc75e2d6310000081565b6040519081526020015b60405180910390f35b610144610184366004611843565b61046b565b34801561019557600080fd5b506101446101a436600461178c565b610568565b3480156101b557600080fd5b506101446101c43660046118cb565b610634565b3480156101d557600080fd5b50610163670de0b6b3a764000081565b6101446101f3366004611907565b610783565b34801561020457600080fd5b50610144610790565b34801561021957600080fd5b50610144610893565b610144610230366004611922565b6108a7565b34801561024157600080fd5b50610265610250366004611907565b60016020526000908152604090205460ff1681565b604051901515815260200161016d565b34801561028157600080fd5b5061028a6108b5565b60405161016d979695949392919061199b565b3480156102a957600080fd5b507f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c199300546040516001600160a01b03909116815260200161016d565b3480156102f057600080fd5b506101446102ff3660046118cb565b610961565b610144610312366004611a34565b610a11565b34801561032357600080fd5b50610163610332366004611907565b610af1565b34801561034357600080fd5b50610144610b5c565b34801561035857600080fd5b50610144610b9a565b34801561036d57600080fd5b506000546102659060ff1681565b34801561038757600080fd5b50610144610396366004611907565b610bd5565b6103a3610c10565b60005b818110156104665760018060008585858181106103c5576103c5611a76565b90506020020160208101906103da9190611907565b6001600160a01b031681526020810191909152604001600020805460ff191691151591909117905582828281811061041457610414611a76565b90506020020160208101906104299190611907565b6001600160a01b03167fc6bdfc1f9b9f1f30ad26b86a7c623e58400512467a50e0c80439bfdaf3a2de9860405160405180910390a26001016103a6565b505050565b60005460ff16158061048c57503360009081526001602052604090205460ff165b6104d45760405162461bcd60e51b815260206004820152601460248201527314dd185ada5b99ce881b9bdd08185b1b1bddd95960621b60448201526064015b60405180910390fd5b68056bc75e2d631000003410156104fd5760405162461bcd60e51b81526004016104cb90611a8c565b60008061050a8686610c6b565b9150915061051b8282338787610e05565b336001600160a01b03167fc7abef7b73f049da6a9bc2349ba5066a39e316eabc9f671b6f9406aa9490a45387873460405161055893929190611ac3565b60405180910390a2505050505050565b610570610c10565b60005b818110156104665760006001600085858581811061059357610593611a76565b90506020020160208101906105a89190611907565b6001600160a01b031681526020810191909152604001600020805460ff19169115159190911790558282828181106105e2576105e2611a76565b90506020020160208101906105f79190611907565b6001600160a01b03167f3df1f5fcca9e1ece84ca685a63062905d8fe97ddb23246224be416f2d3c8613f60405160405180910390a2600101610573565b600080516020611cef8339815191528054600160401b810460ff16159067ffffffffffffffff166000811580156106685750825b905060008267ffffffffffffffff1660011480156106855750303b155b905081158015610693575080155b156106b15760405163f92ee8a960e01b815260040160405180910390fd5b845467ffffffffffffffff1916600117855583156106db57845460ff60401b1916600160401b1785555b6106e487610ec5565b610726604051806040016040528060078152602001665374616b696e6760c81b815250604051806040016040528060018152602001603160f81b815250610ed6565b6000805460ff1916871515179055831561077a57845460ff60401b19168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b50505050505050565b61078d3382610ee8565b50565b600080516020611cef833981519152805460029190600160401b900460ff16806107c85750805467ffffffffffffffff808416911610155b156107e65760405163f92ee8a960e01b815260040160405180910390fd5b805468ffffffffffffffffff191667ffffffffffffffff831617600160401b17815560408051808201825260078152665374616b696e6760c81b602080830191909152825180840190935260018352603160f81b9083015261084791610ed6565b805460ff60401b1916815560405167ffffffffffffffff831681527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15050565b61089b610c10565b6108a56000610fd7565b565b6108b18282610ee8565b5050565b60006060808280808381600080516020611ccf83398151915280549091501580156108e257506001810154155b6109265760405162461bcd60e51b81526020600482015260156024820152741152540dcc4c8e88155b9a5b9a5d1a585b1a5e9959605a1b60448201526064016104cb565b61092e611048565b61093661110b565b60408051600080825260208201909252600f60f81b9c939b5091995046985030975095509350915050565b600080516020611cef8339815191528054600160401b810460ff16159067ffffffffffffffff166000811580156109955750825b905060008267ffffffffffffffff1660011480156109b25750303b155b9050811580156109c0575080155b156109de5760405163f92ee8a960e01b815260040160405180910390fd5b845467ffffffffffffffff191660011785558315610a0857845460ff60401b1916600160401b1785555b61072687610ec5565b60005460ff161580610a3257503360009081526001602052604090205460ff165b610a755760405162461bcd60e51b815260206004820152601460248201527314dd185ada5b99ce881b9bdd08185b1b1bddd95960621b60448201526064016104cb565b68056bc75e2d63100000341015610a9e5760405162461bcd60e51b81526004016104cb90611a8c565b610aa8828261114a565b336001600160a01b03167fc7abef7b73f049da6a9bc2349ba5066a39e316eabc9f671b6f9406aa9490a453838334604051610ae593929190611ac3565b60405180910390a25050565b6000610b567fc9a51567e61a6d1a243a60e57bf4560e7e543694b79349ce2cba3a14fe21b04283604051602001610b3b9291909182526001600160a01b0316602082015260400190565b6040516020818303038152906040528051906020012061115a565b92915050565b610b64610c10565b6000805460ff191660011781556040517f8a943acd5f4e6d3df7565a4a08a93f6b04cc31bb6c01ca4aef7abd6baf455ec39190a1565b610ba2610c10565b6000805460ff191681556040517f2d35c8d348a345fd7b3b03b7cfcf7ad0b60c2d46742d5ca536342e4185becb079190a1565b610bdd610c10565b6001600160a01b038116610c0757604051631e4fbdf760e01b8152600060048201526024016104cb565b61078d81610fd7565b33610c427f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c199300546001600160a01b031690565b6001600160a01b0316146108a55760405163118cdaa760e01b81523360048201526024016104cb565b60008060218314610cbe5760405162461bcd60e51b815260206004820152601e60248201527f536563703235366b313a207075626b6579206e6f74203333206279746573000060448201526064016104cb565b83836000818110610cd157610cd1611a76565b9050013560f81c60f81b6001600160f81b031916600260f81b1480610d1f575083836000818110610d0457610d04611a76565b9050013560f81c60f81b6001600160f81b031916600360f81b145b610d6b5760405162461bcd60e51b815260206004820181905260248201527f536563703235366b313a20696e76616c6964207075626b65792070726566697860448201526064016104cb565b60018401356000610da086868381610d8557610d85611a76565b919091013560f81c905083600060076401000003d019611187565b9050610dac82826112bb565b610df85760405162461bcd60e51b815260206004820152601e60248201527f536563703235366b313a207075626b6579206e6f74206f6e206375727665000060448201526064016104cb565b90925090505b9250929050565b6000610e4f610e1385610af1565b84848080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152506112d992505050565b50509050610e5d8686611326565b6001600160a01b0316816001600160a01b031614610ebd5760405162461bcd60e51b815260206004820152601a60248201527f5374616b696e673a20696e76616c6964207369676e617475726500000000000060448201526064016104cb565b505050505050565b610ecd61135c565b61078d81611393565b610ede61135c565b6108b1828261139b565b60005460ff161580610f1257506001600160a01b03811660009081526001602052604090205460ff165b610f5e5760405162461bcd60e51b815260206004820152601860248201527f5374616b696e673a206e6f7420616c6c6f7765642076616c000000000000000060448201526064016104cb565b670de0b6b3a7640000341015610f865760405162461bcd60e51b81526004016104cb90611a8c565b806001600160a01b0316826001600160a01b03167f510b11bb3f3c799b11307c01ab7db0d335683ef5b2da98f7697de744f465eacc34604051610fcb91815260200190565b60405180910390a35050565b7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c19930080546001600160a01b031981166001600160a01b03848116918217845560405192169182907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a3505050565b7fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d1028054606091600080516020611ccf8339815191529161108790611b12565b80601f01602080910402602001604051908101604052809291908181526020018280546110b390611b12565b80156111005780601f106110d557610100808354040283529160200191611100565b820191906000526020600020905b8154815290600101906020018083116110e357829003601f168201915b505050505091505090565b7fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d1038054606091600080516020611ccf8339815191529161108790611b12565b6111548282610c6b565b50505050565b6000610b566111676113fc565b8360405161190160f01b8152600281019290925260228201526042902090565b60008560ff166002148061119e57508560ff166003145b6112045760405162461bcd60e51b815260206004820152603160248201527f456c6c697074696343757276653a696e6e76616c696420636f6d7072657373656044820152700c8408a8640e0ded2dce840e0e4caccd2f607b1b60648201526084016104cb565b6000828061121457611214611b4c565b838061122257611222611b4c565b85858061123157611231611b4c565b888a0908848061124357611243611b4c565b858061125157611251611b4c565b898a09890908905061127a81600461126a866001611b78565b6112749190611b8b565b8561140b565b90506000600261128d60ff8a1684611b78565b6112979190611b9f565b156112ab576112a68285611bb3565b6112ad565b815b925050505b95945050505050565b60006112d28383600060076401000003d0196114e2565b9392505050565b600080600083516041036113135760208401516040850151606086015160001a6113058882858561159b565b95509550955050505061131f565b50508151600091506002905b9250925092565b60408051818152606081018252600091829190602082018180368337505050602081019485526040810193909352505051902090565b600080516020611cef83398151915254600160401b900460ff166108a557604051631afcd79f60e31b815260040160405180910390fd5b610bdd61135c565b6113a361135c565b600080516020611ccf8339815191527fa16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d1026113dd8482611c0e565b50600381016113ec8382611c0e565b5060008082556001909101555050565b600061140661166a565b905090565b60008160000361145d5760405162461bcd60e51b815260206004820152601e60248201527f456c6c697074696343757276653a206d6f64756c7573206973207a65726f000060448201526064016104cb565b8360000361146d575060006112d2565b8260000361147d575060016112d2565b6001600160ff1b5b80156114d957838186161515870a85848509099150836002820486161515870a85848509099150836004820486161515870a85848509099150836008820486161515870a8584850909915060109004611485565b50949350505050565b60008515806114f15750818610155b806114fa575084155b806115055750818510155b15611512575060006112b2565b6000828061152257611522611b4c565b86870990506000838061153757611537611b4c565b88858061154657611546611b4c565b8a8b09099050851561157657838061156057611560611b4c565b848061156e5761156e611b4c565b878a09820890505b841561159057838061158a5761158a611b4c565b85820890505b149695505050505050565b600080807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a08411156115d65750600091506003905082611660565b604080516000808252602082018084528a905260ff891692820192909252606081018790526080810186905260019060a0016020604051602081039080840390855afa15801561162a573d6000803e3d6000fd5b5050604051601f1901519150506001600160a01b03811661165657506000925060019150829050611660565b9250600091508190505b9450945094915050565b60007f8b73c3c69bb8fe3d512ecc4cf759cc79239f7b179b0ffacaa9a75d522b39400f6116956116de565b61169d611748565b60408051602081019490945283019190915260608201524660808201523060a082015260c00160405160208183030381529060405280519060200120905090565b6000600080516020611ccf833981519152816116f8611048565b80519091501561171057805160209091012092915050565b8154801561171f579392505050565b7fc5d2460186f7233c927e7db2dcc703c0e500b653ca82273b7bfad8045d85a470935050505090565b6000600080516020611ccf8339815191528161176261110b565b80519091501561177a57805160209091012092915050565b6001820154801561171f579392505050565b6000806020838503121561179f57600080fd5b823567ffffffffffffffff808211156117b757600080fd5b818501915085601f8301126117cb57600080fd5b8135818111156117da57600080fd5b8660208260051b85010111156117ef57600080fd5b60209290920196919550909350505050565b60008083601f84011261181357600080fd5b50813567ffffffffffffffff81111561182b57600080fd5b602083019150836020828501011115610dfe57600080fd5b6000806000806040858703121561185957600080fd5b843567ffffffffffffffff8082111561187157600080fd5b61187d88838901611801565b9096509450602087013591508082111561189657600080fd5b506118a387828801611801565b95989497509550505050565b80356001600160a01b03811681146118c657600080fd5b919050565b600080604083850312156118de57600080fd5b6118e7836118af565b9150602083013580151581146118fc57600080fd5b809150509250929050565b60006020828403121561191957600080fd5b6112d2826118af565b6000806040838503121561193557600080fd5b61193e836118af565b915061194c602084016118af565b90509250929050565b6000815180845260005b8181101561197b5760208185018101518683018201520161195f565b506000602082860101526020601f19601f83011685010191505092915050565b60ff60f81b881681526000602060e060208401526119bc60e084018a611955565b83810360408501526119ce818a611955565b606085018990526001600160a01b038816608086015260a0850187905284810360c08601528551808252602080880193509091019060005b81811015611a2257835183529284019291840191600101611a06565b50909c9b505050505050505050505050565b60008060208385031215611a4757600080fd5b823567ffffffffffffffff811115611a5e57600080fd5b611a6a85828601611801565b90969095509350505050565b634e487b7160e01b600052603260045260246000fd5b6020808252601d908201527f5374616b696e673a20696e73756666696369656e74206465706f736974000000604082015260600190565b604081528260408201528284606083013760006060848301015260006060601f19601f8601168301019050826020830152949350505050565b634e487b7160e01b600052604160045260246000fd5b600181811c90821680611b2657607f821691505b602082108103611b4657634e487b7160e01b600052602260045260246000fd5b50919050565b634e487b7160e01b600052601260045260246000fd5b634e487b7160e01b600052601160045260246000fd5b80820180821115610b5657610b56611b62565b600082611b9a57611b9a611b4c565b500490565b600082611bae57611bae611b4c565b500690565b81810381811115610b5657610b56611b62565b601f821115610466576000816000526020600020601f850160051c81016020861015611bef5750805b601f850160051c820191505b81811015610ebd57828155600101611bfb565b815167ffffffffffffffff811115611c2857611c28611afc565b611c3c81611c368454611b12565b84611bc6565b602080601f831160018114611c715760008415611c595750858301515b600019600386901b1c1916600185901b178555610ebd565b600085815260208120601f198616915b82811015611ca057888601518255948401946001909101908401611c81565b5085821015611cbe5787850151600019600388901b60f8161c191681555b5050505050600190811b0190555056fea16a46d94261c7517cc8ff89f61c0ce93598e3c849801011dee649a6a557d100f0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00a264697066735822122046de98416a0be9edf91c0cb513e1d8c65d98ab4a0637cd0c8d9f89bb24f26de464736f6c63430008180033",
}

// StakingABI is the input ABI used to generate the binding from.
// Deprecated: Use StakingMetaData.ABI instead.
var StakingABI = StakingMetaData.ABI

// StakingBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use StakingMetaData.Bin instead.
var StakingBin = StakingMetaData.Bin

// DeployStaking deploys a new Ethereum contract, binding an instance of Staking to it.
func DeployStaking(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Staking, error) {
	parsed, err := StakingMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(StakingBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Staking{StakingCaller: StakingCaller{contract: contract}, StakingTransactor: StakingTransactor{contract: contract}, StakingFilterer: StakingFilterer{contract: contract}}, nil
}

// Staking is an auto generated Go binding around an Ethereum contract.
type Staking struct {
	StakingCaller     // Read-only binding to the contract
	StakingTransactor // Write-only binding to the contract
	StakingFilterer   // Log filterer for contract events
}

// StakingCaller is an auto generated read-only Go binding around an Ethereum contract.
type StakingCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StakingTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StakingFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakingSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StakingSession struct {
	Contract     *Staking          // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StakingCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StakingCallerSession struct {
	Contract *StakingCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts  // Call options to use throughout this session
}

// StakingTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StakingTransactorSession struct {
	Contract     *StakingTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// StakingRaw is an auto generated low-level Go binding around an Ethereum contract.
type StakingRaw struct {
	Contract *Staking // Generic contract binding to access the raw methods on
}

// StakingCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StakingCallerRaw struct {
	Contract *StakingCaller // Generic read-only contract binding to access the raw methods on
}

// StakingTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StakingTransactorRaw struct {
	Contract *StakingTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStaking creates a new instance of Staking, bound to a specific deployed contract.
func NewStaking(address common.Address, backend bind.ContractBackend) (*Staking, error) {
	contract, err := bindStaking(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Staking{StakingCaller: StakingCaller{contract: contract}, StakingTransactor: StakingTransactor{contract: contract}, StakingFilterer: StakingFilterer{contract: contract}}, nil
}

// NewStakingCaller creates a new read-only instance of Staking, bound to a specific deployed contract.
func NewStakingCaller(address common.Address, caller bind.ContractCaller) (*StakingCaller, error) {
	contract, err := bindStaking(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StakingCaller{contract: contract}, nil
}

// NewStakingTransactor creates a new write-only instance of Staking, bound to a specific deployed contract.
func NewStakingTransactor(address common.Address, transactor bind.ContractTransactor) (*StakingTransactor, error) {
	contract, err := bindStaking(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StakingTransactor{contract: contract}, nil
}

// NewStakingFilterer creates a new log filterer instance of Staking, bound to a specific deployed contract.
func NewStakingFilterer(address common.Address, filterer bind.ContractFilterer) (*StakingFilterer, error) {
	contract, err := bindStaking(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StakingFilterer{contract: contract}, nil
}

// bindStaking binds a generic wrapper to an already deployed contract.
func bindStaking(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := StakingMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Staking *StakingRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Staking.Contract.StakingCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Staking *StakingRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Staking.Contract.StakingTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Staking *StakingRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Staking.Contract.StakingTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Staking *StakingCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Staking.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Staking *StakingTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Staking.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Staking *StakingTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Staking.Contract.contract.Transact(opts, method, params...)
}

// MinDelegation is a free data retrieval call binding the contract method 0x59bcddde.
//
// Solidity: function MinDelegation() view returns(uint256)
func (_Staking *StakingCaller) MinDelegation(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "MinDelegation")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinDelegation is a free data retrieval call binding the contract method 0x59bcddde.
//
// Solidity: function MinDelegation() view returns(uint256)
func (_Staking *StakingSession) MinDelegation() (*big.Int, error) {
	return _Staking.Contract.MinDelegation(&_Staking.CallOpts)
}

// MinDelegation is a free data retrieval call binding the contract method 0x59bcddde.
//
// Solidity: function MinDelegation() view returns(uint256)
func (_Staking *StakingCallerSession) MinDelegation() (*big.Int, error) {
	return _Staking.Contract.MinDelegation(&_Staking.CallOpts)
}

// MinDeposit is a free data retrieval call binding the contract method 0x11bcd830.
//
// Solidity: function MinDeposit() view returns(uint256)
func (_Staking *StakingCaller) MinDeposit(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "MinDeposit")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinDeposit is a free data retrieval call binding the contract method 0x11bcd830.
//
// Solidity: function MinDeposit() view returns(uint256)
func (_Staking *StakingSession) MinDeposit() (*big.Int, error) {
	return _Staking.Contract.MinDeposit(&_Staking.CallOpts)
}

// MinDeposit is a free data retrieval call binding the contract method 0x11bcd830.
//
// Solidity: function MinDeposit() view returns(uint256)
func (_Staking *StakingCallerSession) MinDeposit() (*big.Int, error) {
	return _Staking.Contract.MinDeposit(&_Staking.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_Staking *StakingCaller) Eip712Domain(opts *bind.CallOpts) (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "eip712Domain")

	outstruct := new(struct {
		Fields            [1]byte
		Name              string
		Version           string
		ChainId           *big.Int
		VerifyingContract common.Address
		Salt              [32]byte
		Extensions        []*big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Fields = *abi.ConvertType(out[0], new([1]byte)).(*[1]byte)
	outstruct.Name = *abi.ConvertType(out[1], new(string)).(*string)
	outstruct.Version = *abi.ConvertType(out[2], new(string)).(*string)
	outstruct.ChainId = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.VerifyingContract = *abi.ConvertType(out[4], new(common.Address)).(*common.Address)
	outstruct.Salt = *abi.ConvertType(out[5], new([32]byte)).(*[32]byte)
	outstruct.Extensions = *abi.ConvertType(out[6], new([]*big.Int)).(*[]*big.Int)

	return *outstruct, err

}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_Staking *StakingSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _Staking.Contract.Eip712Domain(&_Staking.CallOpts)
}

// Eip712Domain is a free data retrieval call binding the contract method 0x84b0196e.
//
// Solidity: function eip712Domain() view returns(bytes1 fields, string name, string version, uint256 chainId, address verifyingContract, bytes32 salt, uint256[] extensions)
func (_Staking *StakingCallerSession) Eip712Domain() (struct {
	Fields            [1]byte
	Name              string
	Version           string
	ChainId           *big.Int
	VerifyingContract common.Address
	Salt              [32]byte
	Extensions        []*big.Int
}, error) {
	return _Staking.Contract.Eip712Domain(&_Staking.CallOpts)
}

// GetConsPubkeyDigest is a free data retrieval call binding the contract method 0xb62ac12c.
//
// Solidity: function getConsPubkeyDigest(address validator) view returns(bytes32)
func (_Staking *StakingCaller) GetConsPubkeyDigest(opts *bind.CallOpts, validator common.Address) ([32]byte, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "getConsPubkeyDigest", validator)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetConsPubkeyDigest is a free data retrieval call binding the contract method 0xb62ac12c.
//
// Solidity: function getConsPubkeyDigest(address validator) view returns(bytes32)
func (_Staking *StakingSession) GetConsPubkeyDigest(validator common.Address) ([32]byte, error) {
	return _Staking.Contract.GetConsPubkeyDigest(&_Staking.CallOpts, validator)
}

// GetConsPubkeyDigest is a free data retrieval call binding the contract method 0xb62ac12c.
//
// Solidity: function getConsPubkeyDigest(address validator) view returns(bytes32)
func (_Staking *StakingCallerSession) GetConsPubkeyDigest(validator common.Address) ([32]byte, error) {
	return _Staking.Contract.GetConsPubkeyDigest(&_Staking.CallOpts, validator)
}

// IsAllowedValidator is a free data retrieval call binding the contract method 0x84768b7a.
//
// Solidity: function isAllowedValidator(address ) view returns(bool)
func (_Staking *StakingCaller) IsAllowedValidator(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "isAllowedValidator", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAllowedValidator is a free data retrieval call binding the contract method 0x84768b7a.
//
// Solidity: function isAllowedValidator(address ) view returns(bool)
func (_Staking *StakingSession) IsAllowedValidator(arg0 common.Address) (bool, error) {
	return _Staking.Contract.IsAllowedValidator(&_Staking.CallOpts, arg0)
}

// IsAllowedValidator is a free data retrieval call binding the contract method 0x84768b7a.
//
// Solidity: function isAllowedValidator(address ) view returns(bool)
func (_Staking *StakingCallerSession) IsAllowedValidator(arg0 common.Address) (bool, error) {
	return _Staking.Contract.IsAllowedValidator(&_Staking.CallOpts, arg0)
}

// IsAllowlistEnabled is a free data retrieval call binding the contract method 0xd146fd1b.
//
// Solidity: function isAllowlistEnabled() view returns(bool)
func (_Staking *StakingCaller) IsAllowlistEnabled(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "isAllowlistEnabled")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsAllowlistEnabled is a free data retrieval call binding the contract method 0xd146fd1b.
//
// Solidity: function isAllowlistEnabled() view returns(bool)
func (_Staking *StakingSession) IsAllowlistEnabled() (bool, error) {
	return _Staking.Contract.IsAllowlistEnabled(&_Staking.CallOpts)
}

// IsAllowlistEnabled is a free data retrieval call binding the contract method 0xd146fd1b.
//
// Solidity: function isAllowlistEnabled() view returns(bool)
func (_Staking *StakingCallerSession) IsAllowlistEnabled() (bool, error) {
	return _Staking.Contract.IsAllowlistEnabled(&_Staking.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Staking *StakingCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Staking.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Staking *StakingSession) Owner() (common.Address, error) {
	return _Staking.Contract.Owner(&_Staking.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_Staking *StakingCallerSession) Owner() (common.Address, error) {
	return _Staking.Contract.Owner(&_Staking.CallOpts)
}

// AllowValidators is a paid mutator transaction binding the contract method 0x117407e3.
//
// Solidity: function allowValidators(address[] validators) returns()
func (_Staking *StakingTransactor) AllowValidators(opts *bind.TransactOpts, validators []common.Address) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "allowValidators", validators)
}

// AllowValidators is a paid mutator transaction binding the contract method 0x117407e3.
//
// Solidity: function allowValidators(address[] validators) returns()
func (_Staking *StakingSession) AllowValidators(validators []common.Address) (*types.Transaction, error) {
	return _Staking.Contract.AllowValidators(&_Staking.TransactOpts, validators)
}

// AllowValidators is a paid mutator transaction binding the contract method 0x117407e3.
//
// Solidity: function allowValidators(address[] validators) returns()
func (_Staking *StakingTransactorSession) AllowValidators(validators []common.Address) (*types.Transaction, error) {
	return _Staking.Contract.AllowValidators(&_Staking.TransactOpts, validators)
}

// CreateValidator is a paid mutator transaction binding the contract method 0x1fa10bcc.
//
// Solidity: function createValidator(bytes pubkey, bytes signature) payable returns()
func (_Staking *StakingTransactor) CreateValidator(opts *bind.TransactOpts, pubkey []byte, signature []byte) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "createValidator", pubkey, signature)
}

// CreateValidator is a paid mutator transaction binding the contract method 0x1fa10bcc.
//
// Solidity: function createValidator(bytes pubkey, bytes signature) payable returns()
func (_Staking *StakingSession) CreateValidator(pubkey []byte, signature []byte) (*types.Transaction, error) {
	return _Staking.Contract.CreateValidator(&_Staking.TransactOpts, pubkey, signature)
}

// CreateValidator is a paid mutator transaction binding the contract method 0x1fa10bcc.
//
// Solidity: function createValidator(bytes pubkey, bytes signature) payable returns()
func (_Staking *StakingTransactorSession) CreateValidator(pubkey []byte, signature []byte) (*types.Transaction, error) {
	return _Staking.Contract.CreateValidator(&_Staking.TransactOpts, pubkey, signature)
}

// CreateValidator0 is a paid mutator transaction binding the contract method 0xa5a470ad.
//
// Solidity: function createValidator(bytes pubkey) payable returns()
func (_Staking *StakingTransactor) CreateValidator0(opts *bind.TransactOpts, pubkey []byte) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "createValidator0", pubkey)
}

// CreateValidator0 is a paid mutator transaction binding the contract method 0xa5a470ad.
//
// Solidity: function createValidator(bytes pubkey) payable returns()
func (_Staking *StakingSession) CreateValidator0(pubkey []byte) (*types.Transaction, error) {
	return _Staking.Contract.CreateValidator0(&_Staking.TransactOpts, pubkey)
}

// CreateValidator0 is a paid mutator transaction binding the contract method 0xa5a470ad.
//
// Solidity: function createValidator(bytes pubkey) payable returns()
func (_Staking *StakingTransactorSession) CreateValidator0(pubkey []byte) (*types.Transaction, error) {
	return _Staking.Contract.CreateValidator0(&_Staking.TransactOpts, pubkey)
}

// Delegate is a paid mutator transaction binding the contract method 0x5c19a95c.
//
// Solidity: function delegate(address validator) payable returns()
func (_Staking *StakingTransactor) Delegate(opts *bind.TransactOpts, validator common.Address) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "delegate", validator)
}

// Delegate is a paid mutator transaction binding the contract method 0x5c19a95c.
//
// Solidity: function delegate(address validator) payable returns()
func (_Staking *StakingSession) Delegate(validator common.Address) (*types.Transaction, error) {
	return _Staking.Contract.Delegate(&_Staking.TransactOpts, validator)
}

// Delegate is a paid mutator transaction binding the contract method 0x5c19a95c.
//
// Solidity: function delegate(address validator) payable returns()
func (_Staking *StakingTransactorSession) Delegate(validator common.Address) (*types.Transaction, error) {
	return _Staking.Contract.Delegate(&_Staking.TransactOpts, validator)
}

// DelegateFor is a paid mutator transaction binding the contract method 0x78fcbe5b.
//
// Solidity: function delegateFor(address delegator, address validator) payable returns()
func (_Staking *StakingTransactor) DelegateFor(opts *bind.TransactOpts, delegator common.Address, validator common.Address) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "delegateFor", delegator, validator)
}

// DelegateFor is a paid mutator transaction binding the contract method 0x78fcbe5b.
//
// Solidity: function delegateFor(address delegator, address validator) payable returns()
func (_Staking *StakingSession) DelegateFor(delegator common.Address, validator common.Address) (*types.Transaction, error) {
	return _Staking.Contract.DelegateFor(&_Staking.TransactOpts, delegator, validator)
}

// DelegateFor is a paid mutator transaction binding the contract method 0x78fcbe5b.
//
// Solidity: function delegateFor(address delegator, address validator) payable returns()
func (_Staking *StakingTransactorSession) DelegateFor(delegator common.Address, validator common.Address) (*types.Transaction, error) {
	return _Staking.Contract.DelegateFor(&_Staking.TransactOpts, delegator, validator)
}

// DisableAllowlist is a paid mutator transaction binding the contract method 0xcf8e629a.
//
// Solidity: function disableAllowlist() returns()
func (_Staking *StakingTransactor) DisableAllowlist(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "disableAllowlist")
}

// DisableAllowlist is a paid mutator transaction binding the contract method 0xcf8e629a.
//
// Solidity: function disableAllowlist() returns()
func (_Staking *StakingSession) DisableAllowlist() (*types.Transaction, error) {
	return _Staking.Contract.DisableAllowlist(&_Staking.TransactOpts)
}

// DisableAllowlist is a paid mutator transaction binding the contract method 0xcf8e629a.
//
// Solidity: function disableAllowlist() returns()
func (_Staking *StakingTransactorSession) DisableAllowlist() (*types.Transaction, error) {
	return _Staking.Contract.DisableAllowlist(&_Staking.TransactOpts)
}

// DisallowValidators is a paid mutator transaction binding the contract method 0x3f0b1edf.
//
// Solidity: function disallowValidators(address[] validators) returns()
func (_Staking *StakingTransactor) DisallowValidators(opts *bind.TransactOpts, validators []common.Address) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "disallowValidators", validators)
}

// DisallowValidators is a paid mutator transaction binding the contract method 0x3f0b1edf.
//
// Solidity: function disallowValidators(address[] validators) returns()
func (_Staking *StakingSession) DisallowValidators(validators []common.Address) (*types.Transaction, error) {
	return _Staking.Contract.DisallowValidators(&_Staking.TransactOpts, validators)
}

// DisallowValidators is a paid mutator transaction binding the contract method 0x3f0b1edf.
//
// Solidity: function disallowValidators(address[] validators) returns()
func (_Staking *StakingTransactorSession) DisallowValidators(validators []common.Address) (*types.Transaction, error) {
	return _Staking.Contract.DisallowValidators(&_Staking.TransactOpts, validators)
}

// EnableAllowlist is a paid mutator transaction binding the contract method 0xc6a2aac8.
//
// Solidity: function enableAllowlist() returns()
func (_Staking *StakingTransactor) EnableAllowlist(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "enableAllowlist")
}

// EnableAllowlist is a paid mutator transaction binding the contract method 0xc6a2aac8.
//
// Solidity: function enableAllowlist() returns()
func (_Staking *StakingSession) EnableAllowlist() (*types.Transaction, error) {
	return _Staking.Contract.EnableAllowlist(&_Staking.TransactOpts)
}

// EnableAllowlist is a paid mutator transaction binding the contract method 0xc6a2aac8.
//
// Solidity: function enableAllowlist() returns()
func (_Staking *StakingTransactorSession) EnableAllowlist() (*types.Transaction, error) {
	return _Staking.Contract.EnableAllowlist(&_Staking.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x400ada75.
//
// Solidity: function initialize(address owner_, bool isAllowlistEnabled_) returns()
func (_Staking *StakingTransactor) Initialize(opts *bind.TransactOpts, owner_ common.Address, isAllowlistEnabled_ bool) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "initialize", owner_, isAllowlistEnabled_)
}

// Initialize is a paid mutator transaction binding the contract method 0x400ada75.
//
// Solidity: function initialize(address owner_, bool isAllowlistEnabled_) returns()
func (_Staking *StakingSession) Initialize(owner_ common.Address, isAllowlistEnabled_ bool) (*types.Transaction, error) {
	return _Staking.Contract.Initialize(&_Staking.TransactOpts, owner_, isAllowlistEnabled_)
}

// Initialize is a paid mutator transaction binding the contract method 0x400ada75.
//
// Solidity: function initialize(address owner_, bool isAllowlistEnabled_) returns()
func (_Staking *StakingTransactorSession) Initialize(owner_ common.Address, isAllowlistEnabled_ bool) (*types.Transaction, error) {
	return _Staking.Contract.Initialize(&_Staking.TransactOpts, owner_, isAllowlistEnabled_)
}

// InitializeV1 is a paid mutator transaction binding the contract method 0x8f38fae8.
//
// Solidity: function initializeV1(address owner_, bool isAllowlistEnabled_) returns()
func (_Staking *StakingTransactor) InitializeV1(opts *bind.TransactOpts, owner_ common.Address, isAllowlistEnabled_ bool) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "initializeV1", owner_, isAllowlistEnabled_)
}

// InitializeV1 is a paid mutator transaction binding the contract method 0x8f38fae8.
//
// Solidity: function initializeV1(address owner_, bool isAllowlistEnabled_) returns()
func (_Staking *StakingSession) InitializeV1(owner_ common.Address, isAllowlistEnabled_ bool) (*types.Transaction, error) {
	return _Staking.Contract.InitializeV1(&_Staking.TransactOpts, owner_, isAllowlistEnabled_)
}

// InitializeV1 is a paid mutator transaction binding the contract method 0x8f38fae8.
//
// Solidity: function initializeV1(address owner_, bool isAllowlistEnabled_) returns()
func (_Staking *StakingTransactorSession) InitializeV1(owner_ common.Address, isAllowlistEnabled_ bool) (*types.Transaction, error) {
	return _Staking.Contract.InitializeV1(&_Staking.TransactOpts, owner_, isAllowlistEnabled_)
}

// InitializeV2 is a paid mutator transaction binding the contract method 0x5cd8a76b.
//
// Solidity: function initializeV2() returns()
func (_Staking *StakingTransactor) InitializeV2(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "initializeV2")
}

// InitializeV2 is a paid mutator transaction binding the contract method 0x5cd8a76b.
//
// Solidity: function initializeV2() returns()
func (_Staking *StakingSession) InitializeV2() (*types.Transaction, error) {
	return _Staking.Contract.InitializeV2(&_Staking.TransactOpts)
}

// InitializeV2 is a paid mutator transaction binding the contract method 0x5cd8a76b.
//
// Solidity: function initializeV2() returns()
func (_Staking *StakingTransactorSession) InitializeV2() (*types.Transaction, error) {
	return _Staking.Contract.InitializeV2(&_Staking.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Staking *StakingTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Staking *StakingSession) RenounceOwnership() (*types.Transaction, error) {
	return _Staking.Contract.RenounceOwnership(&_Staking.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Staking *StakingTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Staking.Contract.RenounceOwnership(&_Staking.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Staking *StakingTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _Staking.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Staking *StakingSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Staking.Contract.TransferOwnership(&_Staking.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_Staking *StakingTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _Staking.Contract.TransferOwnership(&_Staking.TransactOpts, newOwner)
}

// StakingAllowlistDisabledIterator is returned from FilterAllowlistDisabled and is used to iterate over the raw logs and unpacked data for AllowlistDisabled events raised by the Staking contract.
type StakingAllowlistDisabledIterator struct {
	Event *StakingAllowlistDisabled // Event containing the contract specifics and raw log

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
func (it *StakingAllowlistDisabledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingAllowlistDisabled)
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
		it.Event = new(StakingAllowlistDisabled)
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
func (it *StakingAllowlistDisabledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingAllowlistDisabledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingAllowlistDisabled represents a AllowlistDisabled event raised by the Staking contract.
type StakingAllowlistDisabled struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterAllowlistDisabled is a free log retrieval operation binding the contract event 0x2d35c8d348a345fd7b3b03b7cfcf7ad0b60c2d46742d5ca536342e4185becb07.
//
// Solidity: event AllowlistDisabled()
func (_Staking *StakingFilterer) FilterAllowlistDisabled(opts *bind.FilterOpts) (*StakingAllowlistDisabledIterator, error) {

	logs, sub, err := _Staking.contract.FilterLogs(opts, "AllowlistDisabled")
	if err != nil {
		return nil, err
	}
	return &StakingAllowlistDisabledIterator{contract: _Staking.contract, event: "AllowlistDisabled", logs: logs, sub: sub}, nil
}

// WatchAllowlistDisabled is a free log subscription operation binding the contract event 0x2d35c8d348a345fd7b3b03b7cfcf7ad0b60c2d46742d5ca536342e4185becb07.
//
// Solidity: event AllowlistDisabled()
func (_Staking *StakingFilterer) WatchAllowlistDisabled(opts *bind.WatchOpts, sink chan<- *StakingAllowlistDisabled) (event.Subscription, error) {

	logs, sub, err := _Staking.contract.WatchLogs(opts, "AllowlistDisabled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingAllowlistDisabled)
				if err := _Staking.contract.UnpackLog(event, "AllowlistDisabled", log); err != nil {
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

// ParseAllowlistDisabled is a log parse operation binding the contract event 0x2d35c8d348a345fd7b3b03b7cfcf7ad0b60c2d46742d5ca536342e4185becb07.
//
// Solidity: event AllowlistDisabled()
func (_Staking *StakingFilterer) ParseAllowlistDisabled(log types.Log) (*StakingAllowlistDisabled, error) {
	event := new(StakingAllowlistDisabled)
	if err := _Staking.contract.UnpackLog(event, "AllowlistDisabled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingAllowlistEnabledIterator is returned from FilterAllowlistEnabled and is used to iterate over the raw logs and unpacked data for AllowlistEnabled events raised by the Staking contract.
type StakingAllowlistEnabledIterator struct {
	Event *StakingAllowlistEnabled // Event containing the contract specifics and raw log

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
func (it *StakingAllowlistEnabledIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingAllowlistEnabled)
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
		it.Event = new(StakingAllowlistEnabled)
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
func (it *StakingAllowlistEnabledIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingAllowlistEnabledIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingAllowlistEnabled represents a AllowlistEnabled event raised by the Staking contract.
type StakingAllowlistEnabled struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterAllowlistEnabled is a free log retrieval operation binding the contract event 0x8a943acd5f4e6d3df7565a4a08a93f6b04cc31bb6c01ca4aef7abd6baf455ec3.
//
// Solidity: event AllowlistEnabled()
func (_Staking *StakingFilterer) FilterAllowlistEnabled(opts *bind.FilterOpts) (*StakingAllowlistEnabledIterator, error) {

	logs, sub, err := _Staking.contract.FilterLogs(opts, "AllowlistEnabled")
	if err != nil {
		return nil, err
	}
	return &StakingAllowlistEnabledIterator{contract: _Staking.contract, event: "AllowlistEnabled", logs: logs, sub: sub}, nil
}

// WatchAllowlistEnabled is a free log subscription operation binding the contract event 0x8a943acd5f4e6d3df7565a4a08a93f6b04cc31bb6c01ca4aef7abd6baf455ec3.
//
// Solidity: event AllowlistEnabled()
func (_Staking *StakingFilterer) WatchAllowlistEnabled(opts *bind.WatchOpts, sink chan<- *StakingAllowlistEnabled) (event.Subscription, error) {

	logs, sub, err := _Staking.contract.WatchLogs(opts, "AllowlistEnabled")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingAllowlistEnabled)
				if err := _Staking.contract.UnpackLog(event, "AllowlistEnabled", log); err != nil {
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

// ParseAllowlistEnabled is a log parse operation binding the contract event 0x8a943acd5f4e6d3df7565a4a08a93f6b04cc31bb6c01ca4aef7abd6baf455ec3.
//
// Solidity: event AllowlistEnabled()
func (_Staking *StakingFilterer) ParseAllowlistEnabled(log types.Log) (*StakingAllowlistEnabled, error) {
	event := new(StakingAllowlistEnabled)
	if err := _Staking.contract.UnpackLog(event, "AllowlistEnabled", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingCreateValidatorIterator is returned from FilterCreateValidator and is used to iterate over the raw logs and unpacked data for CreateValidator events raised by the Staking contract.
type StakingCreateValidatorIterator struct {
	Event *StakingCreateValidator // Event containing the contract specifics and raw log

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
func (it *StakingCreateValidatorIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingCreateValidator)
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
		it.Event = new(StakingCreateValidator)
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
func (it *StakingCreateValidatorIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingCreateValidatorIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingCreateValidator represents a CreateValidator event raised by the Staking contract.
type StakingCreateValidator struct {
	Validator common.Address
	Pubkey    []byte
	Deposit   *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterCreateValidator is a free log retrieval operation binding the contract event 0xc7abef7b73f049da6a9bc2349ba5066a39e316eabc9f671b6f9406aa9490a453.
//
// Solidity: event CreateValidator(address indexed validator, bytes pubkey, uint256 deposit)
func (_Staking *StakingFilterer) FilterCreateValidator(opts *bind.FilterOpts, validator []common.Address) (*StakingCreateValidatorIterator, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _Staking.contract.FilterLogs(opts, "CreateValidator", validatorRule)
	if err != nil {
		return nil, err
	}
	return &StakingCreateValidatorIterator{contract: _Staking.contract, event: "CreateValidator", logs: logs, sub: sub}, nil
}

// WatchCreateValidator is a free log subscription operation binding the contract event 0xc7abef7b73f049da6a9bc2349ba5066a39e316eabc9f671b6f9406aa9490a453.
//
// Solidity: event CreateValidator(address indexed validator, bytes pubkey, uint256 deposit)
func (_Staking *StakingFilterer) WatchCreateValidator(opts *bind.WatchOpts, sink chan<- *StakingCreateValidator, validator []common.Address) (event.Subscription, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _Staking.contract.WatchLogs(opts, "CreateValidator", validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingCreateValidator)
				if err := _Staking.contract.UnpackLog(event, "CreateValidator", log); err != nil {
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

// ParseCreateValidator is a log parse operation binding the contract event 0xc7abef7b73f049da6a9bc2349ba5066a39e316eabc9f671b6f9406aa9490a453.
//
// Solidity: event CreateValidator(address indexed validator, bytes pubkey, uint256 deposit)
func (_Staking *StakingFilterer) ParseCreateValidator(log types.Log) (*StakingCreateValidator, error) {
	event := new(StakingCreateValidator)
	if err := _Staking.contract.UnpackLog(event, "CreateValidator", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingDelegateIterator is returned from FilterDelegate and is used to iterate over the raw logs and unpacked data for Delegate events raised by the Staking contract.
type StakingDelegateIterator struct {
	Event *StakingDelegate // Event containing the contract specifics and raw log

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
func (it *StakingDelegateIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingDelegate)
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
		it.Event = new(StakingDelegate)
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
func (it *StakingDelegateIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingDelegateIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingDelegate represents a Delegate event raised by the Staking contract.
type StakingDelegate struct {
	Delegator common.Address
	Validator common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDelegate is a free log retrieval operation binding the contract event 0x510b11bb3f3c799b11307c01ab7db0d335683ef5b2da98f7697de744f465eacc.
//
// Solidity: event Delegate(address indexed delegator, address indexed validator, uint256 amount)
func (_Staking *StakingFilterer) FilterDelegate(opts *bind.FilterOpts, delegator []common.Address, validator []common.Address) (*StakingDelegateIterator, error) {

	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _Staking.contract.FilterLogs(opts, "Delegate", delegatorRule, validatorRule)
	if err != nil {
		return nil, err
	}
	return &StakingDelegateIterator{contract: _Staking.contract, event: "Delegate", logs: logs, sub: sub}, nil
}

// WatchDelegate is a free log subscription operation binding the contract event 0x510b11bb3f3c799b11307c01ab7db0d335683ef5b2da98f7697de744f465eacc.
//
// Solidity: event Delegate(address indexed delegator, address indexed validator, uint256 amount)
func (_Staking *StakingFilterer) WatchDelegate(opts *bind.WatchOpts, sink chan<- *StakingDelegate, delegator []common.Address, validator []common.Address) (event.Subscription, error) {

	var delegatorRule []interface{}
	for _, delegatorItem := range delegator {
		delegatorRule = append(delegatorRule, delegatorItem)
	}
	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _Staking.contract.WatchLogs(opts, "Delegate", delegatorRule, validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingDelegate)
				if err := _Staking.contract.UnpackLog(event, "Delegate", log); err != nil {
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

// ParseDelegate is a log parse operation binding the contract event 0x510b11bb3f3c799b11307c01ab7db0d335683ef5b2da98f7697de744f465eacc.
//
// Solidity: event Delegate(address indexed delegator, address indexed validator, uint256 amount)
func (_Staking *StakingFilterer) ParseDelegate(log types.Log) (*StakingDelegate, error) {
	event := new(StakingDelegate)
	if err := _Staking.contract.UnpackLog(event, "Delegate", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingEIP712DomainChangedIterator is returned from FilterEIP712DomainChanged and is used to iterate over the raw logs and unpacked data for EIP712DomainChanged events raised by the Staking contract.
type StakingEIP712DomainChangedIterator struct {
	Event *StakingEIP712DomainChanged // Event containing the contract specifics and raw log

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
func (it *StakingEIP712DomainChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingEIP712DomainChanged)
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
		it.Event = new(StakingEIP712DomainChanged)
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
func (it *StakingEIP712DomainChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingEIP712DomainChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingEIP712DomainChanged represents a EIP712DomainChanged event raised by the Staking contract.
type StakingEIP712DomainChanged struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterEIP712DomainChanged is a free log retrieval operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_Staking *StakingFilterer) FilterEIP712DomainChanged(opts *bind.FilterOpts) (*StakingEIP712DomainChangedIterator, error) {

	logs, sub, err := _Staking.contract.FilterLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return &StakingEIP712DomainChangedIterator{contract: _Staking.contract, event: "EIP712DomainChanged", logs: logs, sub: sub}, nil
}

// WatchEIP712DomainChanged is a free log subscription operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_Staking *StakingFilterer) WatchEIP712DomainChanged(opts *bind.WatchOpts, sink chan<- *StakingEIP712DomainChanged) (event.Subscription, error) {

	logs, sub, err := _Staking.contract.WatchLogs(opts, "EIP712DomainChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingEIP712DomainChanged)
				if err := _Staking.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
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

// ParseEIP712DomainChanged is a log parse operation binding the contract event 0x0a6387c9ea3628b88a633bb4f3b151770f70085117a15f9bf3787cda53f13d31.
//
// Solidity: event EIP712DomainChanged()
func (_Staking *StakingFilterer) ParseEIP712DomainChanged(log types.Log) (*StakingEIP712DomainChanged, error) {
	event := new(StakingEIP712DomainChanged)
	if err := _Staking.contract.UnpackLog(event, "EIP712DomainChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the Staking contract.
type StakingInitializedIterator struct {
	Event *StakingInitialized // Event containing the contract specifics and raw log

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
func (it *StakingInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingInitialized)
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
		it.Event = new(StakingInitialized)
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
func (it *StakingInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingInitialized represents a Initialized event raised by the Staking contract.
type StakingInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Staking *StakingFilterer) FilterInitialized(opts *bind.FilterOpts) (*StakingInitializedIterator, error) {

	logs, sub, err := _Staking.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &StakingInitializedIterator{contract: _Staking.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_Staking *StakingFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *StakingInitialized) (event.Subscription, error) {

	logs, sub, err := _Staking.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingInitialized)
				if err := _Staking.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_Staking *StakingFilterer) ParseInitialized(log types.Log) (*StakingInitialized, error) {
	event := new(StakingInitialized)
	if err := _Staking.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Staking contract.
type StakingOwnershipTransferredIterator struct {
	Event *StakingOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *StakingOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingOwnershipTransferred)
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
		it.Event = new(StakingOwnershipTransferred)
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
func (it *StakingOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingOwnershipTransferred represents a OwnershipTransferred event raised by the Staking contract.
type StakingOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Staking *StakingFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*StakingOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Staking.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &StakingOwnershipTransferredIterator{contract: _Staking.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Staking *StakingFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *StakingOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Staking.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingOwnershipTransferred)
				if err := _Staking.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Staking *StakingFilterer) ParseOwnershipTransferred(log types.Log) (*StakingOwnershipTransferred, error) {
	event := new(StakingOwnershipTransferred)
	if err := _Staking.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingValidatorAllowedIterator is returned from FilterValidatorAllowed and is used to iterate over the raw logs and unpacked data for ValidatorAllowed events raised by the Staking contract.
type StakingValidatorAllowedIterator struct {
	Event *StakingValidatorAllowed // Event containing the contract specifics and raw log

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
func (it *StakingValidatorAllowedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingValidatorAllowed)
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
		it.Event = new(StakingValidatorAllowed)
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
func (it *StakingValidatorAllowedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingValidatorAllowedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingValidatorAllowed represents a ValidatorAllowed event raised by the Staking contract.
type StakingValidatorAllowed struct {
	Validator common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterValidatorAllowed is a free log retrieval operation binding the contract event 0xc6bdfc1f9b9f1f30ad26b86a7c623e58400512467a50e0c80439bfdaf3a2de98.
//
// Solidity: event ValidatorAllowed(address indexed validator)
func (_Staking *StakingFilterer) FilterValidatorAllowed(opts *bind.FilterOpts, validator []common.Address) (*StakingValidatorAllowedIterator, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _Staking.contract.FilterLogs(opts, "ValidatorAllowed", validatorRule)
	if err != nil {
		return nil, err
	}
	return &StakingValidatorAllowedIterator{contract: _Staking.contract, event: "ValidatorAllowed", logs: logs, sub: sub}, nil
}

// WatchValidatorAllowed is a free log subscription operation binding the contract event 0xc6bdfc1f9b9f1f30ad26b86a7c623e58400512467a50e0c80439bfdaf3a2de98.
//
// Solidity: event ValidatorAllowed(address indexed validator)
func (_Staking *StakingFilterer) WatchValidatorAllowed(opts *bind.WatchOpts, sink chan<- *StakingValidatorAllowed, validator []common.Address) (event.Subscription, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _Staking.contract.WatchLogs(opts, "ValidatorAllowed", validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingValidatorAllowed)
				if err := _Staking.contract.UnpackLog(event, "ValidatorAllowed", log); err != nil {
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

// ParseValidatorAllowed is a log parse operation binding the contract event 0xc6bdfc1f9b9f1f30ad26b86a7c623e58400512467a50e0c80439bfdaf3a2de98.
//
// Solidity: event ValidatorAllowed(address indexed validator)
func (_Staking *StakingFilterer) ParseValidatorAllowed(log types.Log) (*StakingValidatorAllowed, error) {
	event := new(StakingValidatorAllowed)
	if err := _Staking.contract.UnpackLog(event, "ValidatorAllowed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakingValidatorDisallowedIterator is returned from FilterValidatorDisallowed and is used to iterate over the raw logs and unpacked data for ValidatorDisallowed events raised by the Staking contract.
type StakingValidatorDisallowedIterator struct {
	Event *StakingValidatorDisallowed // Event containing the contract specifics and raw log

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
func (it *StakingValidatorDisallowedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakingValidatorDisallowed)
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
		it.Event = new(StakingValidatorDisallowed)
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
func (it *StakingValidatorDisallowedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakingValidatorDisallowedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakingValidatorDisallowed represents a ValidatorDisallowed event raised by the Staking contract.
type StakingValidatorDisallowed struct {
	Validator common.Address
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterValidatorDisallowed is a free log retrieval operation binding the contract event 0x3df1f5fcca9e1ece84ca685a63062905d8fe97ddb23246224be416f2d3c8613f.
//
// Solidity: event ValidatorDisallowed(address indexed validator)
func (_Staking *StakingFilterer) FilterValidatorDisallowed(opts *bind.FilterOpts, validator []common.Address) (*StakingValidatorDisallowedIterator, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _Staking.contract.FilterLogs(opts, "ValidatorDisallowed", validatorRule)
	if err != nil {
		return nil, err
	}
	return &StakingValidatorDisallowedIterator{contract: _Staking.contract, event: "ValidatorDisallowed", logs: logs, sub: sub}, nil
}

// WatchValidatorDisallowed is a free log subscription operation binding the contract event 0x3df1f5fcca9e1ece84ca685a63062905d8fe97ddb23246224be416f2d3c8613f.
//
// Solidity: event ValidatorDisallowed(address indexed validator)
func (_Staking *StakingFilterer) WatchValidatorDisallowed(opts *bind.WatchOpts, sink chan<- *StakingValidatorDisallowed, validator []common.Address) (event.Subscription, error) {

	var validatorRule []interface{}
	for _, validatorItem := range validator {
		validatorRule = append(validatorRule, validatorItem)
	}

	logs, sub, err := _Staking.contract.WatchLogs(opts, "ValidatorDisallowed", validatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakingValidatorDisallowed)
				if err := _Staking.contract.UnpackLog(event, "ValidatorDisallowed", log); err != nil {
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

// ParseValidatorDisallowed is a log parse operation binding the contract event 0x3df1f5fcca9e1ece84ca685a63062905d8fe97ddb23246224be416f2d3c8613f.
//
// Solidity: event ValidatorDisallowed(address indexed validator)
func (_Staking *StakingFilterer) ParseValidatorDisallowed(log types.Log) (*StakingValidatorDisallowed, error) {
	event := new(StakingValidatorDisallowed)
	if err := _Staking.contract.UnpackLog(event, "ValidatorDisallowed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
