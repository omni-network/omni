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

// NominaPortalInitParams is an auto generated low-level Go binding around an user-defined struct.
type NominaPortalInitParams struct {
	Owner                common.Address
	FeeOracle            common.Address
	NominaChainId        uint64
	NominaCChainId       uint64
	XmsgMaxGasLimit      uint64
	XmsgMinGasLimit      uint64
	XmsgMaxDataSize      uint16
	XreceiptMaxErrorSize uint16
	XsubValsetCutoff     uint8
	CChainXMsgOffset     uint64
	CChainXBlockOffset   uint64
	ValSetId             uint64
	Validators           []XTypesValidator
}

// XTypesBlockHeader is an auto generated low-level Go binding around an user-defined struct.
// autocommented by commenttypes.go
// type XTypesBlockHeader struct {
// 	SourceChainId     uint64
// 	ConsensusChainId  uint64
// 	ConfLevel         uint8
// 	Offset            uint64
// 	SourceBlockHeight uint64
// 	SourceBlockHash   [32]byte
// }

// XTypesChain is an auto generated low-level Go binding around an user-defined struct.
// autocommented by commenttypes.go
// type XTypesChain struct {
// 	ChainId uint64
// 	Shards  []uint64
// }

// XTypesMsg is an auto generated low-level Go binding around an user-defined struct.
// autocommented by commenttypes.go
// type XTypesMsg struct {
// 	DestChainId uint64
// 	ShardId     uint64
// 	Offset      uint64
// 	Sender      common.Address
// 	To          common.Address
// 	Data        []byte
// 	GasLimit    uint64
// }

// XTypesMsgContext is an auto generated low-level Go binding around an user-defined struct.
// autocommented by commenttypes.go
// type XTypesMsgContext struct {
// 	SourceChainId uint64
// 	Sender        common.Address
// }

// XTypesSigTuple is an auto generated low-level Go binding around an user-defined struct.
// autocommented by commenttypes.go
// type XTypesSigTuple struct {
// 	ValidatorAddr common.Address
// 	Signature     []byte
// }

// XTypesSubmission is an auto generated low-level Go binding around an user-defined struct.
// autocommented by commenttypes.go
// type XTypesSubmission struct {
// 	AttestationRoot [32]byte
// 	ValidatorSetId  uint64
// 	BlockHeader     XTypesBlockHeader
// 	Msgs            []XTypesMsg
// 	Proof           [][32]byte
// 	ProofFlags      []bool
// 	Signatures      []XTypesSigTuple
// }

// XTypesValidator is an auto generated low-level Go binding around an user-defined struct.
// autocommented by commenttypes.go
// type XTypesValidator struct {
// 	Addr  common.Address
// 	Power uint64
// }

// NominaPortalMetaData contains all meta data concerning the NominaPortal contract.
var NominaPortalMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"ActionXCall\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"ActionXSubmit\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"KeyPauseAll\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"XSubQuorumDenominator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"XSubQuorumNumerator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"addValidatorSet\",\"inputs\":[{\"name\":\"valSetId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"validators\",\"type\":\"tuple[]\",\"internalType\":\"structXTypes.Validator[]\",\"components\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"power\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"chainId\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"collectFees\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"feeFor\",\"inputs\":[{\"name\":\"destChainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"gasLimit\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"feeOracle\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"inXBlockOffset\",\"inputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"inXMsgOffset\",\"inputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"p\",\"type\":\"tuple\",\"internalType\":\"structNominaPortal.InitParams\",\"components\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"feeOracle\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"nominaChainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"nominaCChainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"xmsgMaxGasLimit\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"xmsgMinGasLimit\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"xmsgMaxDataSize\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"xreceiptMaxErrorSize\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"xsubValsetCutoff\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"cChainXMsgOffset\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"cChainXBlockOffset\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"valSetId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"validators\",\"type\":\"tuple[]\",\"internalType\":\"structXTypes.Validator[]\",\"components\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"power\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isPaused\",\"inputs\":[{\"name\":\"actionId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isPaused\",\"inputs\":[{\"name\":\"actionId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"chainId_\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isPaused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isSupportedDest\",\"inputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isSupportedShard\",\"inputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isXCall\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"latestValSetId\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"network\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple[]\",\"internalType\":\"structXTypes.Chain[]\",\"components\":[{\"name\":\"chainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"shards\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"nominaCChainId\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"nominaChainId\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"omniCChainId\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"omniChainId\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"outXMsgOffset\",\"inputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pauseXCall\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pauseXCallTo\",\"inputs\":[{\"name\":\"chainId_\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pauseXSubmit\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pauseXSubmitFrom\",\"inputs\":[{\"name\":\"chainId_\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setFeeOracle\",\"inputs\":[{\"name\":\"feeOracle_\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setInXBlockOffset\",\"inputs\":[{\"name\":\"sourceChainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"shardId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"offset\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setInXMsgOffset\",\"inputs\":[{\"name\":\"sourceChainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"shardId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"offset\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setNetwork\",\"inputs\":[{\"name\":\"network_\",\"type\":\"tuple[]\",\"internalType\":\"structXTypes.Chain[]\",\"components\":[{\"name\":\"chainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"shards\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setXMsgMaxDataSize\",\"inputs\":[{\"name\":\"numBytes\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setXMsgMaxGasLimit\",\"inputs\":[{\"name\":\"gasLimit\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setXMsgMinGasLimit\",\"inputs\":[{\"name\":\"gasLimit\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setXReceiptMaxErrorSize\",\"inputs\":[{\"name\":\"numBytes\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setXSubValsetCutoff\",\"inputs\":[{\"name\":\"xsubValsetCutoff_\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unpauseXCall\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unpauseXCallTo\",\"inputs\":[{\"name\":\"chainId_\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unpauseXSubmit\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unpauseXSubmitFrom\",\"inputs\":[{\"name\":\"chainId_\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"valSet\",\"inputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"valSetTotalPower\",\"inputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"xcall\",\"inputs\":[{\"name\":\"destChainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"conf\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"gasLimit\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"xmsg\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structXTypes.MsgContext\",\"components\":[{\"name\":\"sourceChainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"xmsgMaxDataSize\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"xmsgMaxGasLimit\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"xmsgMinGasLimit\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"xreceiptMaxErrorSize\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"xsubValsetCutoff\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"xsubmit\",\"inputs\":[{\"name\":\"xsub\",\"type\":\"tuple\",\"internalType\":\"structXTypes.Submission\",\"components\":[{\"name\":\"attestationRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"validatorSetId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"blockHeader\",\"type\":\"tuple\",\"internalType\":\"structXTypes.BlockHeader\",\"components\":[{\"name\":\"sourceChainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"consensusChainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"confLevel\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"offset\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"sourceBlockHeight\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"sourceBlockHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"msgs\",\"type\":\"tuple[]\",\"internalType\":\"structXTypes.Msg[]\",\"components\":[{\"name\":\"destChainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"shardId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"offset\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"gasLimit\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"proof\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"proofFlags\",\"type\":\"bool[]\",\"internalType\":\"bool[]\"},{\"name\":\"signatures\",\"type\":\"tuple[]\",\"internalType\":\"structXTypes.SigTuple[]\",\"components\":[{\"name\":\"validatorAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"FeeOracleSet\",\"inputs\":[{\"name\":\"oracle\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FeesCollected\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"InXBlockOffsetSet\",\"inputs\":[{\"name\":\"srcChainId\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"shardId\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"offset\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"InXMsgOffsetSet\",\"inputs\":[{\"name\":\"srcChainId\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"shardId\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"offset\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"key\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"key\",\"type\":\"bytes32\",\"indexed\":true,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ValidatorSetAdded\",\"inputs\":[{\"name\":\"setId\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"XCallPaused\",\"inputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"XCallToPaused\",\"inputs\":[{\"name\":\"chainId\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"XCallToUnpaused\",\"inputs\":[{\"name\":\"chainId\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"XCallUnpaused\",\"inputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"XMsg\",\"inputs\":[{\"name\":\"destChainId\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"shardId\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"offset\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"gasLimit\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"fees\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"XMsgMaxDataSizeSet\",\"inputs\":[{\"name\":\"size\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"XMsgMaxGasLimitSet\",\"inputs\":[{\"name\":\"gasLimit\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"XMsgMinGasLimitSet\",\"inputs\":[{\"name\":\"gasLimit\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"XReceipt\",\"inputs\":[{\"name\":\"sourceChainId\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"shardId\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"offset\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"gasUsed\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"relayer\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"success\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"},{\"name\":\"err\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"XReceiptMaxErrorSizeSet\",\"inputs\":[{\"name\":\"size\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"XSubValsetCutoffSet\",\"inputs\":[{\"name\":\"cutoff\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"XSubmitFromPaused\",\"inputs\":[{\"name\":\"chainId\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"XSubmitFromUnpaused\",\"inputs\":[{\"name\":\"chainId\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"XSubmitPaused\",\"inputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"XSubmitUnpaused\",\"inputs\":[],\"anonymous\":false},{\"type\":\"error\",\"name\":\"ECDSAInvalidSignature\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ECDSAInvalidSignatureLength\",\"inputs\":[{\"name\":\"length\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ECDSAInvalidSignatureS\",\"inputs\":[{\"name\":\"s\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MerkleProofInvalidMultiproof\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnableInvalidOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnableUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ReentrancyGuardReentrantCall\",\"inputs\":[]}]",
	Bin: "0x6080604052348015600e575f5ffd5b5060156019565b60c9565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000900460ff161560685760405163f92ee8a960e01b815260040160405180910390fd5b80546001600160401b039081161460c65780546001600160401b0319166001600160401b0390811782556040519081527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b50565b614c95806100d65f395ff3fe608060405260043610610371575f3560e01c80638532eb9f116101c8578063b2b2f5bd116100fd578063c2f9b9681161009d578063d051c97d1161006d578063d051c97d14610af6578063d533b44514610b35578063f2fde38b14610b54578063f45cc7b814610b73575f5ffd5b8063c2f9b96814610a7f578063c3d8ad6714610a9e578063c4ab80bc14610ab2578063cf84c81814610ad1575f5ffd5b8063bb8590ad116100d8578063bb8590ad14610a10578063bff0e84d14610a2f578063c21dda4f14610a4e578063c26dfc0514610a61575f5ffd5b8063b2b2f5bd1461099e578063b4d5afd1146109be578063b521466d146109f1575f5ffd5b8063a32eb7c611610168578063aaf1bc9711610143578063aaf1bc9714610909578063afe8219814610937578063afe8af9c14610956578063b187bd261461098a575f5ffd5b8063a32eb7c6146108ab578063a480ca79146108cb578063a8a98962146108ea575f5ffd5b80638dd9523c116101a35780638dd9523c1461082d57806397b520621461085a5780639a8a059214610879578063a10ac97a1461088b575f5ffd5b80638532eb9f146107ac57806387320ac2146107cb5780638da5cb5b146107f1575f5ffd5b80633f4ba83a116102a957806355e2448e11610249578063715018a611610219578063715018a61461074b57806378fe53071461075f57806383d0cbd9146107845780638456cb5914610798575f5ffd5b806355e2448e146106ad57806357542050146106cc57806366a1eaf31461070b5780636739afca1461072a575f5ffd5b806349cc3bf61161028457806349cc3bf61461062b5780634cc3908814610643578063500b19e71461066257806354d26bba14610699575f5ffd5b80633f4ba83a146105b95780633fd3b15e146105cd578063461ab4881461060c575f5ffd5b8063241b71bb1161031457806330632e8b116102ef57806330632e8b1461051857806336d219121461053757806336d853f91461055b5780633aa873301461057a575f5ffd5b8063241b71bb1461045857806324278bbe146104875780632f32700e146104b5575f5ffd5b806310a5a7f71161034f57806310a5a7f7146103d5578063110ff5f1146103f45780631d3eb6e31461042557806323dbce5014610444575f5ffd5b80630360d20f1461037557806306c3dc5f146103a0578063103ba701146103b4575b5f5ffd5b348015610380575f5ffd5b50610389600281565b60405160ff90911681526020015b60405180910390f35b3480156103ab575f5ffd5b50610389600381565b3480156103bf575f5ffd5b506103d36103ce366004613fba565b610b98565b005b3480156103e0575f5ffd5b506103d36103ef366004613ff2565b610bac565b3480156103ff575f5ffd5b506001546001600160401b03165b6040516001600160401b039091168152602001610397565b348015610430575f5ffd5b506103d361043f36600461400d565b610c09565b34801561044f575f5ffd5b506103d3610d28565b348015610463575f5ffd5b5061047761047236600461407c565b610d70565b6040519015158152602001610397565b348015610492575f5ffd5b506104776104a1366004613ff2565b60056020525f908152604090205460ff1681565b3480156104c0575f5ffd5b506040805180820182525f80825260209182015281518083018352600b546001600160401b0381168083526001600160a01b03600160401b909204821692840192835284519081529151169181019190915201610397565b348015610523575f5ffd5b506103d3610532366004614093565b610d80565b348015610542575f5ffd5b50600154600160401b90046001600160401b031661040d565b348015610566575f5ffd5b506103d3610575366004613ff2565b61100a565b348015610585575f5ffd5b5061040d6105943660046140ca565b600660209081525f92835260408084209091529082529020546001600160401b031681565b3480156105c4575f5ffd5b506103d361101b565b3480156105d8575f5ffd5b5061040d6105e73660046140ca565b600860209081525f92835260408084209091529082529020546001600160401b031681565b348015610617575f5ffd5b50610477610626366004614101565b611055565b348015610636575f5ffd5b505f546103899060ff1681565b34801561064e575f5ffd5b5060015461040d906001600160401b031681565b34801561066d575f5ffd5b50600254610681906001600160a01b031681565b6040516001600160a01b039091168152602001610397565b3480156106a4575f5ffd5b506103d3611070565b3480156106b8575f5ffd5b50600b546001600160401b03161515610477565b3480156106d7575f5ffd5b5061040d6106e636600461413a565b600a60209081525f92835260408084209091529082529020546001600160401b031681565b348015610716575f5ffd5b506103d361072536600461416d565b6110b8565b348015610735575f5ffd5b5061073e61145e565b60405161039791906141a4565b348015610756575f5ffd5b506103d361154d565b34801561076a575f5ffd5b505f5461040d90600160681b90046001600160401b031681565b34801561078f575f5ffd5b506103d3611560565b3480156107a3575f5ffd5b506103d36115a8565b3480156107b7575f5ffd5b506103d36107c636600461425b565b6115e2565b3480156107d6575f5ffd5b5060015461040d90600160401b90046001600160401b031681565b3480156107fc575f5ffd5b507f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c199300546001600160a01b0316610681565b348015610838575f5ffd5b5061084c610847366004614320565b6116f9565b604051908152602001610397565b348015610865575f5ffd5b506103d3610874366004614383565b611777565b348015610884575f5ffd5b504661040d565b348015610896575f5ffd5b5061084c5f516020614c355f395f51905f5281565b3480156108b6575f5ffd5b5061084c5f516020614c755f395f51905f5281565b3480156108d6575f5ffd5b506103d36108e53660046143cb565b61178a565b3480156108f5575f5ffd5b506103d36109043660046143cb565b61180f565b348015610914575f5ffd5b50610477610923366004613ff2565b60046020525f908152604090205460ff1681565b348015610942575f5ffd5b506103d3610951366004613ff2565b611820565b348015610961575f5ffd5b5061040d610970366004613ff2565b60096020525f90815260409020546001600160401b031681565b348015610995575f5ffd5b50610477611878565b3480156109a9575f5ffd5b5061084c5f516020614c155f395f51905f5281565b3480156109c9575f5ffd5b505f546109de906301000000900461ffff1681565b60405161ffff9091168152602001610397565b3480156109fc575f5ffd5b506103d3610a0b3660046143e4565b6118ca565b348015610a1b575f5ffd5b506103d3610a2a366004613ff2565b6118db565b348015610a3a575f5ffd5b506103d3610a493660046143e4565b6118ec565b6103d3610a5c366004614405565b6118fd565b348015610a6c575f5ffd5b505f546109de90610100900461ffff1681565b348015610a8a575f5ffd5b506103d3610a99366004613ff2565b611cce565b348015610aa9575f5ffd5b506103d3611d2b565b348015610abd575f5ffd5b506103d3610acc366004614383565b611d73565b348015610adc575f5ffd5b505f5461040d90600160281b90046001600160401b031681565b348015610b01575f5ffd5b5061040d610b103660046140ca565b600760209081525f92835260408084209091529082529020546001600160401b031681565b348015610b40575f5ffd5b506103d3610b4f366004613ff2565b611d86565b348015610b5f575f5ffd5b506103d3610b6e3660046143cb565b611dde565b348015610b7e575f5ffd5b505f5461040d90600160a81b90046001600160401b031681565b610ba0611e18565b610ba981611e73565b50565b610bb4611e18565b610bd3610bce5f516020614c155f395f51905f5283611f0d565b611f55565b6040516001600160401b038216907fcd7910e1c5569d8433ce4ef8e5d51c1bdc03168f614b576da47dc3d2b51d033a905f90a250565b333014610c575760405162461bcd60e51b81526020600482015260176024820152762737b6b4b730a837b93a30b61d1037b7363c9039b2b63360491b60448201526064015b60405180910390fd5b600154600b546001600160401b03908116600160401b9092041614610cba5760405162461bcd60e51b81526020600482015260196024820152782737b6b4b730a837b93a30b61d1037b7363c9031b1b430b4b760391b6044820152606401610c4e565b600b54600160401b90046001600160a01b031615610d1a5760405162461bcd60e51b815260206004820181905260248201527f4e6f6d696e61506f7274616c3a206f6e6c792063636861696e2073656e6465726044820152606401610c4e565b610d248282611ff8565b5050565b610d30611e18565b610d465f516020614c755f395f51905f52611f55565b6040517f3d0f9c56dac46156a2db0aa09ee7804770ad9fc9549d21023164f22d69475ed8905f90a1565b5f610d7a8261216b565b92915050565b5f610d896121ce565b805490915060ff600160401b82041615906001600160401b03165f81158015610daf5750825b90505f826001600160401b03166001148015610dca5750303b155b905081158015610dd8575080155b15610df65760405163f92ee8a960e01b815260040160405180910390fd5b845467ffffffffffffffff191660011785558315610e2057845460ff60401b1916600160401b1785555b610e35610e3060208801886143cb565b6121f6565b610e4d610e4860408801602089016143cb565b612207565b610e65610e6060a0880160808901613ff2565b6122ab565b610e7d610e7860e0880160c089016143e4565b61236b565b610e95610e9060c0880160a08901613ff2565b61240d565b610eae610ea9610100880160e089016143e4565b612523565b610ec8610ec361012088016101008901613fba565b611e73565b610ef0610edd61018088016101608901613ff2565b610eeb610180890189614489565b6125c1565b610f006060870160408801613ff2565b6001805467ffffffffffffffff19166001600160401b0392909216919091179055610f316080870160608801613ff2565b600180546001600160401b0392909216600160401b026fffffffffffffffff000000000000000019909216919091179055610104610f90610f786080890160608a01613ff2565b82610f8b6101408b016101208c01613ff2565b6128f2565b610fbb610fa36080890160608a01613ff2565b82610fb66101608b016101408c01613ff2565b612966565b50831561100257845460ff60401b19168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b505050505050565b611012611e18565b610ba9816122ab565b611023611e18565b61102b6129d1565b6040517fa45f47fdea8a1efdd9029a5691c7f759c32b7c698632b563573e155625d16933905f90a1565b5f611069836110648585611f0d565b6129e7565b9392505050565b611078611e18565b61108e5f516020614c155f395f51905f52612a69565b6040517f4c48c7b71557216a3192842746bdfc381f98d7536d9eb1c6764f3b45e6794827905f90a1565b5f516020614c755f395f51905f526110d66060830160408401613ff2565b6110e4826110648484611f0d565b156111285760405162461bcd60e51b8152602060048201526014602482015273139bdb5a5b98541bdc9d185b0e881c185d5cd95960621b6044820152606401610c4e565b611130612b0c565b365f6111406101008601866144ce565b9092509050604085015f6111578260208901613ff2565b600154909150600160401b90046001600160401b031661117d6040840160208501613ff2565b6001600160401b0316146111d35760405162461bcd60e51b815260206004820152601d60248201527f4e6f6d696e61506f7274616c3a2077726f6e672063636861696e2049440000006044820152606401610c4e565b826112195760405162461bcd60e51b81526020600482015260166024820152754e6f6d696e61506f7274616c3a206e6f20786d73677360501b6044820152606401610c4e565b6001600160401b038082165f908152600960205260409020541661127f5760405162461bcd60e51b815260206004820152601d60248201527f4e6f6d696e61506f7274616c3a20756e6b6e6f776e2076616c207365740000006044820152606401610c4e565b611287612b56565b6001600160401b0316816001600160401b031610156112e85760405162461bcd60e51b815260206004820152601960248201527f4e6f6d696e61506f7274616c3a206f6c642076616c20736574000000000000006044820152606401610c4e565b61132b87356112fb6101608a018a6144ce565b6001600160401b038086165f908152600a6020908152604080832060099092529091205490911660026003612ba4565b6113775760405162461bcd60e51b815260206004820152601760248201527f4e6f6d696e61506f7274616c3a206e6f2071756f72756d0000000000000000006044820152606401610c4e565b6113a0873583868661138d6101208d018d6144ce565b61139b6101408f018f6144ce565b612d53565b6113ec5760405162461bcd60e51b815260206004820152601b60248201527f4e6f6d696e61506f7274616c3a20696e76616c69642070726f6f6600000000006044820152606401610c4e565b5f5b8381101561142b576114238386868481811061140c5761140c614513565b905060200281019061141e9190614527565b612dcc565b6001016113ee565b505050505061145960017f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f0055565b505050565b60606003805480602002602001604051908101604052809291908181526020015f905b82821015611544575f8481526020908190206040805180820182526002860290920180546001600160401b0316835260018101805483518187028101870190945280845293949193858301939283018282801561152c57602002820191905f5260205f20905f905b82829054906101000a90046001600160401b03166001600160401b0316815260200190600801906020826007010492830192600103820291508084116114e95790505b50505050508152505081526020019060010190611481565b50505050905090565b611555611e18565b61155e5f61351b565b565b611568611e18565b61157e5f516020614c155f395f51905f52611f55565b6040517f5f335a4032d4cfb6aca7835b0c2225f36d4d9eaa4ed43ee59ed537e02dff6b39905f90a1565b6115b0611e18565b6115b861358b565b6040517f9e87fac88ff661f02d44f95383c817fece4bce600a3dab7a54406878b965e752905f90a1565b33301461162b5760405162461bcd60e51b81526020600482015260176024820152762737b6b4b730a837b93a30b61d1037b7363c9039b2b63360491b6044820152606401610c4e565b600154600b546001600160401b03908116600160401b909204161461168e5760405162461bcd60e51b81526020600482015260196024820152782737b6b4b730a837b93a30b61d1037b7363c9031b1b430b4b760391b6044820152606401610c4e565b600b54600160401b90046001600160a01b0316156116ee5760405162461bcd60e51b815260206004820181905260248201527f4e6f6d696e61506f7274616c3a206f6e6c792063636861696e2073656e6465726044820152606401610c4e565b6114598383836125c1565b600254604051632376548f60e21b81525f916001600160a01b031690638dd9523c9061172f90889088908890889060040161456d565b602060405180830381865afa15801561174a573d5f5f3e3d5ffd5b505050506040513d601f19601f8201168201806040525081019061176e91906145a9565b95945050505050565b61177f611e18565b611459838383612966565b611792611e18565b60405147906001600160a01b0383169082156108fc029083905f818181858888f193505050501580156117c7573d5f5f3e3d5ffd5b50816001600160a01b03167f9dc46f23cfb5ddcad0ae7ea2be38d47fec07bb9382ec7e564efc69e036dd66ce8260405161180391815260200190565b60405180910390a25050565b611817611e18565b610ba981612207565b611828611e18565b611842610bce5f516020614c755f395f51905f5283611f0d565b6040516001600160401b038216907fab78810a0515df65f9f10bfbcb92d03d5df71d9fd3b9414e9ad831a5117d6daa905f90a250565b5f6118c55f516020614c355f395f51905f525f525f516020614c555f395f51905f526020527ffae9838a178d7f201aa98e2ce5340158edda60bb1e8f168f46503bf3e99f13be5460ff1690565b905090565b6118d2611e18565b610ba98161236b565b6118e3611e18565b610ba98161240d565b6118f4611e18565b610ba981612523565b5f516020614c155f395f51905f528661191a826110648484611f0d565b1561195e5760405162461bcd60e51b8152602060048201526014602482015273139bdb5a5b98541bdc9d185b0e881c185d5cd95960621b6044820152606401610c4e565b6001600160401b0388165f9081526005602052604090205460ff166119c55760405162461bcd60e51b815260206004820152601e60248201527f4e6f6d696e61506f7274616c3a20756e737570706f72746564206465737400006044820152606401610c4e565b6001600160a01b038616611a1b5760405162461bcd60e51b815260206004820152601d60248201527f4e6f6d696e61506f7274616c3a206e6f20706f7274616c207863616c6c0000006044820152606401610c4e565b5f546001600160401b03600160281b90910481169084161115611a805760405162461bcd60e51b815260206004820152601f60248201527f4e6f6d696e61506f7274616c3a206761734c696d697420746f6f2068696768006044820152606401610c4e565b5f546001600160401b03600160681b90910481169084161015611ae55760405162461bcd60e51b815260206004820152601e60248201527f4e6f6d696e61506f7274616c3a206761734c696d697420746f6f206c6f7700006044820152606401610c4e565b5f546301000000900461ffff16841115611b415760405162461bcd60e51b815260206004820152601c60248201527f4e6f6d696e61506f7274616c3a206461746120746f6f206c61726765000000006044820152606401610c4e565b60ff8088165f81815260046020526040902054909116611ba35760405162461bcd60e51b815260206004820152601f60248201527f4e6f6d696e61506f7274616c3a20756e737570706f72746564207368617264006044820152606401610c4e565b5f611bb08a8888886116f9565b905080341015611c025760405162461bcd60e51b815260206004820152601e60248201527f4e6f6d696e61506f7274616c3a20696e73756666696369656e742066656500006044820152606401610c4e565b6001600160401b03808b165f908152600660209081526040808320868516845290915281208054600193919291611c3b918591166145d4565b82546101009290920a6001600160401b038181021990931691831602179091558b81165f8181526006602090815260408083208886168085529252918290205491519190931693507fb7c8eb9d7a7fbcdab809ab7b8a7c41701eb3115e3fe99d30ff490d8552f72bfa90611cba9033908e908e908e908e908b906145f3565b60405180910390a450505050505050505050565b611cd6611e18565b611cf5611cf05f516020614c755f395f51905f5283611f0d565b612a69565b6040516001600160401b038216907fc551305d9bd408be4327b7f8aba28b04ccf6b6c76925392d195ecf9cc764294d905f90a250565b611d33611e18565b611d495f516020614c755f395f51905f52612a69565b6040517f2cb9d71d4c31860b70e9b707c69aa2f5953e03474f00cfcfff205c4745f82875905f90a1565b611d7b611e18565b6114598383836128f2565b611d8e611e18565b611da8611cf05f516020614c155f395f51905f5283611f0d565b6040516001600160401b038216907f1ed9223556fb0971076c30172f1f00630efd313b6a05290a562aef95928e7125905f90a250565b611de6611e18565b6001600160a01b038116611e0f57604051631e4fbdf760e01b81525f6004820152602401610c4e565b610ba98161351b565b33611e4a7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c199300546001600160a01b031690565b6001600160a01b03161461155e5760405163118cdaa760e01b8152336004820152602401610c4e565b5f8160ff1611611ec55760405162461bcd60e51b815260206004820152601c60248201527f4e6f6d696e61506f7274616c3a206e6f207a65726f206375746f6666000000006044820152606401610c4e565b5f805460ff191660ff83169081179091556040519081527f1683dc51426224f6e37a3b41dd5849e2db1bfe22366d1d913fa0ef6f757e828f906020015b60405180910390a150565b5f8282604051602001611f3792919091825260c01b6001600160c01b031916602082015260280190565b60405160208183030381529060405280519060200120905092915050565b5f8181525f516020614c555f395f51905f52602081905260409091205460ff1615611fb55760405162461bcd60e51b815260206004820152601060248201526f14185d5cd8589b194e881c185d5cd95960821b6044820152606401610c4e565b5f82815260208290526040808220805460ff191660011790555183917f0cb09dc71d57eeec2046f6854976717e4874a3cf2d6ddeddde337e5b6de6ba3191a25050565b6120006135a1565b365f5b828110156121655783838281811061201d5761201d614513565b905060200281019061202f919061463d565b600380546001810182555f9190915290925082906002027fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b016120728282614702565b505061207b4690565b6001600160401b03166120916020840184613ff2565b6001600160401b0316146120dd57600160055f6120b16020860186613ff2565b6001600160401b0316815260208101919091526040015f20805460ff191691151591909117905561215d565b5f5b6120ec60208401846144ce565b905081101561215b57600160045f61210760208701876144ce565b8581811061211757612117614513565b905060200201602081019061212c9190613ff2565b6001600160401b0316815260208101919091526040015f20805460ff19169115159190911790556001016120df565b505b600101612003565b50505050565b5f516020614c355f395f51905f525f9081525f516020614c555f395f51905f5260208190527ffae9838a178d7f201aa98e2ce5340158edda60bb1e8f168f46503bf3e99f13be5460ff168061106957505f92835260205250604090205460ff1690565b5f807ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00610d7a565b6121fe613697565b610ba9816136bc565b6001600160a01b03811661225d5760405162461bcd60e51b815260206004820152601f60248201527f4e6f6d696e61506f7274616c3a206e6f207a65726f206665654f7261636c65006044820152606401610c4e565b600280546001600160a01b0319166001600160a01b0383169081179091556040519081527fd97bdb0db82b52a85aa07f8da78033b1d6e159d94f1e3cbd4109d946c3bcfd3290602001611f02565b5f546001600160401b03600160681b90910481169082161161230f5760405162461bcd60e51b815260206004820152601b60248201527f4e6f6d696e61506f7274616c3a206e6f742061626f7665206d696e00000000006044820152606401610c4e565b5f80546cffffffffffffffff00000000001916600160281b6001600160401b038416908102919091179091556040519081527f1153561ac5effc2926ba6c612f86a397c997bc43dfbfc718da08065be0c5fe4d90602001611f02565b5f8161ffff16116123be5760405162461bcd60e51b815260206004820152601e60248201527f4e6f6d696e61506f7274616c3a206e6f207a65726f206d61782073697a6500006044820152606401610c4e565b5f805464ffff0000001916630100000061ffff8416908102919091179091556040519081527f65923e04419dc810d0ea08a94a7f608d4c4d949818d95c3788f895e575dd206490602001611f02565b5f816001600160401b0316116124655760405162461bcd60e51b815260206004820152601d60248201527f4e6f6d696e61506f7274616c3a206e6f207a65726f206d696e206761730000006044820152606401610c4e565b5f546001600160401b03600160281b9091048116908216106124c95760405162461bcd60e51b815260206004820152601b60248201527f4e6f6d696e61506f7274616c3a206e6f742062656c6f77206d617800000000006044820152606401610c4e565b5f805467ffffffffffffffff60681b1916600160681b6001600160401b038416908102919091179091556040519081527f8c852a6291aa436654b167353bca4a4b0c3d024c7562cb5082e7c869bddabf3e90602001611f02565b5f8161ffff16116125765760405162461bcd60e51b815260206004820152601e60248201527f4e6f6d696e61506f7274616c3a206e6f207a65726f206d61782073697a6500006044820152606401610c4e565b5f805462ffff00191661010061ffff8416908102919091179091556040519081527f620bbea084306b66a8cc6b5b63830d6b3874f9d2438914e259ffd5065c33f7b090602001611f02565b808061260f5760405162461bcd60e51b815260206004820152601b60248201527f4e6f6d696e61506f7274616c3a206e6f2076616c696461746f727300000000006044820152606401610c4e565b6001600160401b038085165f9081526009602052604090205416156126765760405162461bcd60e51b815260206004820152601f60248201527f4e6f6d696e61506f7274616c3a206475706c69636174652076616c20736574006044820152606401610c4e565b6040805180820182525f80825260208083018290526001600160401b0388168252600a9052918220825b84811015612854578686828181106126ba576126ba614513565b9050604002018036038101906126d09190614843565b80519093506001600160a01b031661272a5760405162461bcd60e51b815260206004820152601f60248201527f4e6f6d696e61506f7274616c3a206e6f207a65726f2076616c696461746f72006044820152606401610c4e565b5f83602001516001600160401b0316116127865760405162461bcd60e51b815260206004820152601b60248201527f4e6f6d696e61506f7274616c3a206e6f207a65726f20706f77657200000000006044820152606401610c4e565b82516001600160a01b03165f908152602083905260409020546001600160401b0316156127ff5760405162461bcd60e51b815260206004820152602160248201527f4e6f6d696e61506f7274616c3a206475706c69636174652076616c696461746f6044820152603960f91b6064820152608401610c4e565b602083015161280e90856145d4565b60208481015185516001600160a01b03165f908152918590526040909120805467ffffffffffffffff19166001600160401b0390921691909117905593506001016126a0565b506001600160401b038781165f818152600960205260408120805467ffffffffffffffff191687851617905554600160a81b900490911610156128b6575f805467ffffffffffffffff60a81b1916600160a81b6001600160401b038a16021790555b6040516001600160401b038816907f3a7c2f997a87ba92aedaecd1127f4129cae1283e2809ebf5304d321b943fd107905f90a250505050505050565b6001600160401b038381165f81815260076020908152604080832087861680855290835292819020805467ffffffffffffffff191695871695861790555193845290927f8647aae68c8456a1dcbfaf5eaadc94278ae423526d3f09c7b972bff7355d55c791015b60405180910390a3505050565b6001600160401b038381165f81815260086020908152604080832087861680855290835292819020805467ffffffffffffffff191695871695861790555193845290927fe070f08cae8464c91238e8cbea64ccee5e7b48dd79a843f144e3721ee6bdd9b59101612959565b61155e5f516020614c355f395f51905f52612a69565b5f516020614c355f395f51905f525f9081525f516020614c555f395f51905f5260208190527ffae9838a178d7f201aa98e2ce5340158edda60bb1e8f168f46503bf3e99f13be5460ff1680612a4957505f8481526020829052604090205460ff165b80612a6157505f8381526020829052604090205460ff165b949350505050565b5f8181525f516020614c555f395f51905f52602081905260409091205460ff16612acc5760405162461bcd60e51b815260206004820152601460248201527314185d5cd8589b194e881b9bdd081c185d5cd95960621b6044820152606401610c4e565b5f82815260208290526040808220805460ff191690555183917fd05bfc2250abb0f8fd265a54c53a24359c5484af63cad2e4ce87c78ab751395a91a25050565b7f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f00805460011901612b5057604051633ee5aeb560e01b815260040160405180910390fd5b60029055565b5f805460ff8116600160a81b9091046001600160401b031611612b795750600190565b5f54612b999060ff811690600160a81b90046001600160401b031661489f565b6118c59060016145d4565b5f8036815b88811015612d4157898982818110612bc357612bc3614513565b9050602002810190612bd5919061463d565b91508015612c8957368a8a612beb6001856148be565b818110612bfa57612bfa614513565b9050602002810190612c0c919061463d565b9050612c1b60208201826143cb565b6001600160a01b0316612c3160208501856143cb565b6001600160a01b031611612c875760405162461bcd60e51b815260206004820152601f60248201527f51756f72756d3a2073696773206e6f7420646564757065642f736f72746564006044820152606401610c4e565b505b612c93828c6136c4565b612cdf5760405162461bcd60e51b815260206004820152601960248201527f51756f72756d3a20696e76616c6964207369676e6174757265000000000000006044820152606401610c4e565b875f612cee60208501856143cb565b6001600160a01b0316815260208101919091526040015f2054612d1a906001600160401b0316846145d4565b9250612d2883888888613736565b15612d395760019350505050612d48565b600101612ba9565b505f925050505b979650505050505050565b6040805160018082528183019092525f9182919060208083019080368337019050509050612d8d86868686612d888d8d61376d565b613838565b815f81518110612d9f57612d9f614513565b602002602001018181525050612dbe818b612db98c613a8e565b613aa5565b9a9950505050505050505050565b5f612dda6020840184613ff2565b90505f612dea6020840184613ff2565b90505f612dfd6040850160208601613ff2565b90505f612e106060860160408701613ff2565b9050466001600160401b0316836001600160401b03161480612e3957506001600160401b038316155b612e855760405162461bcd60e51b815260206004820152601e60248201527f4e6f6d696e61506f7274616c3a2077726f6e67206465737420636861696e00006044820152606401610c4e565b6001600160401b038085165f9081526007602090815260408083208685168452909152902054612eb7911660016145d4565b6001600160401b0316816001600160401b031614612f175760405162461bcd60e51b815260206004820152601a60248201527f4e6f6d696e61506f7274616c3a2077726f6e67206f66667365740000000000006044820152606401610c4e565b612f276060870160408801613fba565b60ff16600460ff161480612f4f575060ff8216612f4a6060880160408901613fba565b60ff16145b612f9b5760405162461bcd60e51b815260206004820152601e60248201527f4e6f6d696e61506f7274616c3a2077726f6e6720636f6e66206c6576656c00006044820152606401610c4e565b612fab6080870160608801613ff2565b6001600160401b038581165f90815260086020908152604080832087851684529091529020549181169116101561302957612fec6080870160608801613ff2565b6001600160401b038581165f90815260086020908152604080832087851684529091529020805467ffffffffffffffff1916929091169190911790555b6001600160401b038085165f908152600760209081526040808320868516845290915281208054600193919291613062918591166145d4565b92506101000a8154816001600160401b0302191690836001600160401b03160217905550306001600160a01b03168560800160208101906130a391906143cb565b6001600160a01b03160361317957806001600160401b0316826001600160401b0316856001600160401b03167f8277cab1f0fa69b34674f64a7d43f242b0bacece6f5b7e8652f1e0d88a9b873b5f335f604051602401613132906020808252818101527f4e6f6d696e61506f7274616c3a206e6f207863616c6c20746f20706f7274616c604082015260600190565b60408051601f198184030181529181526020820180516001600160e01b031662461bcd60e51b1790525161316994939291906148ff565b60405180910390a4505050505050565b5f8061318b60a08801608089016143cb565b6001600160a01b031614905080156132d1575f6131ab60a088018861493a565b6131b49161497c565b600154909150600160401b90046001600160401b03166131d760208a018a613ff2565b6001600160401b031614801561320457505f6131f96080890160608a016143cb565b6001600160a01b0316145b801561322457505f6132196020890189613ff2565b6001600160401b0316145b8015613249575061010461323e6040890160208a01613ff2565b6001600160401b0316145b801561327f57506001600160e01b03198116638532eb9f60e01b148061327f57506001600160e01b03198116631d3eb6e360e01b145b6132cb5760405162461bcd60e51b815260206004820152601d60248201527f4e6f6d696e61506f7274616c3a20696e76616c69642073797363616c6c0000006044820152606401610c4e565b506133b3565b600154600160401b90046001600160401b03166132f16020890189613ff2565b6001600160401b03161415801561332057505f61331460808801606089016143cb565b6001600160a01b031614155b801561334157505f6133356020880188613ff2565b6001600160401b031614155b8015613367575061010461335b6040880160208901613ff2565b6001600160401b031614155b6133b35760405162461bcd60e51b815260206004820152601b60248201527f4e6f6d696e61506f7274616c3a20696e76616c6964207863616c6c00000000006044820152606401610c4e565b604080518082019091526001600160401b0386168152602081016133dd6080890160608a016143cb565b6001600160a01b039081169091528151600b8054602090940151909216600160401b026001600160e01b03199093166001600160401b03909116179190911790555f8080836134695761346461343960a08b0160808c016143cb565b61344960e08c0160c08d01613ff2565b6001600160401b031661345f60a08d018d61493a565b613aba565b61347e565b61347e61347960a08b018b61493a565b613b74565b600b80546001600160e01b0319169055919450925090505f836134a157826134b1565b60405180602001604052805f8152505b9050856001600160401b0316876001600160401b03168a6001600160401b03167f8277cab1f0fa69b34674f64a7d43f242b0bacece6f5b7e8652f1e0d88a9b873b8533898760405161350694939291906148ff565b60405180910390a45050505050505050505050565b7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c19930080546001600160a01b031981166001600160a01b03848116918217845560405192169182907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0905f90a3505050565b61155e5f516020614c355f395f51905f52611f55565b5f805b60035481101561368b57600381815481106135c1576135c1614513565b905f5260205f20906002020191506135d64690565b82546001600160401b0390811691161461360f5781546001600160401b03165f908152600560205260409020805460ff19169055613683565b5f5b6001830154811015613681575f60045f85600101848154811061363657613636614513565b5f918252602080832060048304015460039092166008026101000a9091046001600160401b031683528201929092526040019020805460ff1916911515919091179055600101613611565b505b6001016135a4565b50610ba960035f613f28565b61369f613c04565b61155e57604051631afcd79f60e31b815260040160405180910390fd5b611de6613697565b5f6136d260208401846143cb565b6001600160a01b0316613725836136ec602087018761493a565b8080601f0160208091040260200160405190810160405280939291908181526020018383808284375f92019190915250613c1d92505050565b6001600160a01b0316149392505050565b5f61374d60ff84166001600160401b038616614671565b61376360ff84166001600160401b038816614671565b1195945050505050565b60605f826001600160401b038111156137885761378861465d565b6040519080825280602002602001820160405280156137b1578160200160208202803683370190505b5090505f5b838110156138305761380b60028686848181106137d5576137d5614513565b90506020028101906137e79190614527565b6040516020016137f791906149f5565b604051602081830303815290604052613c45565b82828151811061381d5761381d614513565b60209081029190910101526001016137b6565b509392505050565b80515f9083613848816001614acd565b6138528884614acd565b1461387057604051631a8a024960e11b815260040160405180910390fd5b5f816001600160401b038111156138895761388961465d565b6040519080825280602002602001820160405280156138b2578160200160208202803683370190505b5090505f8080805b858110156139fb575f8785106138f45785846138d581614ae0565b9550815181106138e7576138e7614513565b602002602001015161391a565b89856138ff81614ae0565b96508151811061391157613911614513565b60200260200101515b90505f8c8c8481811061392f5761392f614513565b90506020020160208101906139449190614af8565b613971578e8e8561395481614ae0565b965081811061396557613965614513565b905060200201356139c8565b8886106139a257868561398381614ae0565b96508151811061399557613995614513565b60200260200101516139c8565b8a866139ad81614ae0565b9750815181106139bf576139bf614513565b60200260200101515b90506139d48282613c7b565b8784815181106139e6576139e6614513565b602090810291909101015250506001016138ba565b508415613a4c57808b14613a2257604051631a8a024960e11b815260040160405180910390fd5b836001860381518110613a3757613a37614513565b6020026020010151965050505050505061176e565b8515613a6457875f81518110613a3757613a37614513565b8b8b5f818110613a7657613a76614513565b90506020020135965050505050505095945050505050565b5f610d7a6001836040516020016137f79190614b17565b5f82613ab18584613ca7565b14949350505050565b5f60605f5f5a90505f5f613b38895f5f60019054906101000a900461ffff168b8b8080601f0160208091040260200160405190810160405280939291908181526020018383808284375f81840152601f19601f820116905080830192505050505050508e6001600160a01b0316613ce190949392919063ffffffff16565b915091505f5a9050613b4b603f8b614ba9565b8111613b5357fe5b8282613b5f83876148be565b965096509650505050505b9450945094915050565b5f60605f5f5a90505f5f306001600160a01b03168888604051613b98929190614bc8565b5f604051808303815f865af19150503d805f8114613bd1576040519150601f19603f3d011682016040523d82523d5f602084013e613bd6565b606091505b50915091505a613be690846148be565b925081613bf557805160208201fd5b909450925090505b9250925092565b5f613c0d6121ce565b54600160401b900460ff16919050565b5f5f5f5f613c2b8686613d66565b925092509250613c3b8282613dac565b5090949350505050565b5f8282604051602001613c59929190614bd7565b60408051601f1981840301815282825280516020918201209083015201611f37565b5f818310613c95575f828152602084905260409020611069565b5f838152602083905260409020611069565b5f81815b845181101561383057613cd782868381518110613cca57613cca614513565b6020026020010151613c7b565b9150600101613cab565b5f60605f5f5f8661ffff166001600160401b03811115613d0357613d0361465d565b6040519080825280601f01601f191660200182016040528015613d2d576020820181803683370190505b5090505f5f8751602089018b8e8ef191503d925086831115613d4d578692505b828152825f602083013e90999098509650505050505050565b5f5f5f8351604103613d9d576020840151604085015160608601515f1a613d8f88828585613e64565b955095509550505050613bfd565b505081515f9150600290613bfd565b5f826003811115613dbf57613dbf614c00565b03613dc8575050565b6001826003811115613ddc57613ddc614c00565b03613dfa5760405163f645eedf60e01b815260040160405180910390fd5b6002826003811115613e0e57613e0e614c00565b03613e2f5760405163fce698f760e01b815260048101829052602401610c4e565b6003826003811115613e4357613e43614c00565b03610d24576040516335e2f38360e21b815260048101829052602401610c4e565b5f80807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a0841115613e9d57505f91506003905082613b6a565b604080515f808252602082018084528a905260ff891692820192909252606081018790526080810186905260019060a0016020604051602081039080840390855afa158015613eee573d5f5f3e3d5ffd5b5050604051601f1901519150506001600160a01b038116613f1957505f925060019150829050613b6a565b975f9750879650945050505050565b5080545f8255600202905f5260205f2090810190610ba991905b80821115613f7057805467ffffffffffffffff191681555f613f676001830182613f74565b50600201613f42565b5090565b5080545f825560030160049004905f5260205f2090810190610ba991905b80821115613f70575f8155600101613f92565b803560ff81168114613fb5575f5ffd5b919050565b5f60208284031215613fca575f5ffd5b61106982613fa5565b6001600160401b0381168114610ba9575f5ffd5b8035613fb581613fd3565b5f60208284031215614002575f5ffd5b813561106981613fd3565b5f5f6020838503121561401e575f5ffd5b82356001600160401b03811115614033575f5ffd5b8301601f81018513614043575f5ffd5b80356001600160401b03811115614058575f5ffd5b8560208260051b840101111561406c575f5ffd5b6020919091019590945092505050565b5f6020828403121561408c575f5ffd5b5035919050565b5f602082840312156140a3575f5ffd5b81356001600160401b038111156140b8575f5ffd5b82016101a08185031215611069575f5ffd5b5f5f604083850312156140db575f5ffd5b82356140e681613fd3565b915060208301356140f681613fd3565b809150509250929050565b5f5f60408385031215614112575f5ffd5b8235915060208301356140f681613fd3565b80356001600160a01b0381168114613fb5575f5ffd5b5f5f6040838503121561414b575f5ffd5b823561415681613fd3565b915061416460208401614124565b90509250929050565b5f6020828403121561417d575f5ffd5b81356001600160401b03811115614192575f5ffd5b82016101808185031215611069575f5ffd5b5f602082016020835280845180835260408501915060408160051b8601019250602086015f5b8281101561424f57868503603f19018452815180516001600160401b03168652602090810151604082880181905281519088018190529101905f9060608801905b80831015614237576001600160401b03845116825260208201915060208401935060018301925061420b565b509650505060209384019391909101906001016141ca565b50929695505050505050565b5f5f5f6040848603121561426d575f5ffd5b833561427881613fd3565b925060208401356001600160401b03811115614292575f5ffd5b8401601f810186136142a2575f5ffd5b80356001600160401b038111156142b7575f5ffd5b8660208260061b84010111156142cb575f5ffd5b939660209190910195509293505050565b5f5f83601f8401126142ec575f5ffd5b5081356001600160401b03811115614302575f5ffd5b602083019150836020828501011115614319575f5ffd5b9250929050565b5f5f5f5f60608587031215614333575f5ffd5b843561433e81613fd3565b935060208501356001600160401b03811115614358575f5ffd5b614364878288016142dc565b909450925050604085013561437881613fd3565b939692955090935050565b5f5f5f60608486031215614395575f5ffd5b83356143a081613fd3565b925060208401356143b081613fd3565b915060408401356143c081613fd3565b809150509250925092565b5f602082840312156143db575f5ffd5b61106982614124565b5f602082840312156143f4575f5ffd5b813561ffff81168114611069575f5ffd5b5f5f5f5f5f5f60a0878903121561441a575f5ffd5b863561442581613fd3565b955061443360208801613fa5565b945061444160408801614124565b935060608701356001600160401b0381111561445b575f5ffd5b61446789828a016142dc565b909450925050608087013561447b81613fd3565b809150509295509295509295565b5f5f8335601e1984360301811261449e575f5ffd5b8301803591506001600160401b038211156144b7575f5ffd5b6020019150600681901b3603821315614319575f5ffd5b5f5f8335601e198436030181126144e3575f5ffd5b8301803591506001600160401b038211156144fc575f5ffd5b6020019150600581901b3603821315614319575f5ffd5b634e487b7160e01b5f52603260045260245ffd5b5f823560de1983360301811261453b575f5ffd5b9190910192915050565b81835281816020850137505f828201602090810191909152601f909101601f19169091010190565b6001600160401b0385168152606060208201525f61458f606083018587614545565b90506001600160401b038316604083015295945050505050565b5f602082840312156145b9575f5ffd5b5051919050565b634e487b7160e01b5f52601160045260245ffd5b6001600160401b038181168382160190811115610d7a57610d7a6145c0565b6001600160a01b0387811682528616602082015260a0604082018190525f9061461f9083018688614545565b6001600160401b039490941660608301525060800152949350505050565b5f8235603e1983360301811261453b575f5ffd5b5f8135610d7a81613fd3565b634e487b7160e01b5f52604160045260245ffd5b8082028115828204841417610d7a57610d7a6145c0565b600160401b82111561469c5761469c61465d565b80548282558083101561145957815f5260205f206003840160021c81016003830160021c8201915060188560031b1680156146e6575f19820180545f198360200360031b1c168155505b505b818110156146fb575f81556001016146e8565b5050505050565b813561470d81613fd3565b6001600160401b0381166001600160401b0319835416178255506020820135601e1983360301811261473d575f5ffd5b820180356001600160401b03811115614754575f5ffd5b6020820191508060051b360382131561476b575f5ffd5b600183016001600160401b038211156147865761478661465d565b6147908282614688565b5f8181526020902090508160021c5f5b818110156147fb575f5f5b60048110156147ee576147dd6147c088614651565b6001600160401b03908116600684901b90811b91901b1984161790565b6020979097019691506001016147ab565b50838201556001016147a0565b506003198316808403818514614839575f5f5b82811015614833576148226147c089614651565b60209890980197915060010161480e565b50848401555b5050505050505050565b5f6040828403128015614854575f5ffd5b50604080519081016001600160401b03811182821017156148775761487761465d565b60405261488383614124565b8152602083013561489381613fd3565b60208201529392505050565b6001600160401b038281168282160390811115610d7a57610d7a6145c0565b81810381811115610d7a57610d7a6145c0565b5f81518084528060208401602086015e5f602082860101526020601f19601f83011685010191505092915050565b8481526001600160a01b038416602082015282151560408201526080606082018190525f90614930908301846148d1565b9695505050505050565b5f5f8335601e1984360301811261494f575f5ffd5b8301803591506001600160401b03821115614968575f5ffd5b602001915036819003821315614319575f5ffd5b80356001600160e01b031981169060048410156149ad576001600160e01b0319600485900360031b81901b82161691505b5092915050565b5f5f8335601e198436030181126149c9575f5ffd5b83016020810192503590506001600160401b038111156149e7575f5ffd5b803603821315614319575f5ffd5b602081525f8235614a0581613fd3565b6001600160401b0381166020840152506020830135614a2381613fd3565b6001600160401b038116604084015250614a3f60408401613fe7565b6001600160401b038116606084015250614a5b60608401614124565b6001600160a01b038116608084015250614a7760808401614124565b6001600160a01b03811660a084015250614a9460a08401846149b4565b60e060c0850152614aaa61010085018284614545565b915050614ab960c08501613fe7565b6001600160401b03811660e0850152613830565b80820180821115610d7a57610d7a6145c0565b5f60018201614af157614af16145c0565b5060010190565b5f60208284031215614b08575f5ffd5b81358015158114611069575f5ffd5b60c081018235614b2681613fd3565b6001600160401b031682526020830135614b3f81613fd3565b6001600160401b0316602083015260ff614b5b60408501613fa5565b1660408301526060830135614b6f81613fd3565b6001600160401b031660608301526080830135614b8b81613fd3565b6001600160401b03811660808401525060a092830135919092015290565b5f82614bc357634e487b7160e01b5f52601260045260245ffd5b500490565b818382375f9101908152919050565b60ff60f81b8360f81b1681525f82518060208501600185015e5f92016001019182525092915050565b634e487b7160e01b5f52602160045260245ffdfea06a0c1264badca141841b5f52470407dac9adaaa539dd445540986341b73a6876e8952e4b09b8d505aa08998d716721a1dbf0884ac74202e33985da1ed005e9ff37105740f03695c8f3597f3aff2b92fbe1c80abea3c28731ecff2efd693400feccba1cfc4544bf9cd83b76f36ae5c464750b6c43f682e26744ee21ec31fc1e",
}

// NominaPortalABI is the input ABI used to generate the binding from.
// Deprecated: Use NominaPortalMetaData.ABI instead.
var NominaPortalABI = NominaPortalMetaData.ABI

// NominaPortalBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use NominaPortalMetaData.Bin instead.
var NominaPortalBin = NominaPortalMetaData.Bin

// DeployNominaPortal deploys a new Ethereum contract, binding an instance of NominaPortal to it.
func DeployNominaPortal(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *NominaPortal, error) {
	parsed, err := NominaPortalMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(NominaPortalBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &NominaPortal{NominaPortalCaller: NominaPortalCaller{contract: contract}, NominaPortalTransactor: NominaPortalTransactor{contract: contract}, NominaPortalFilterer: NominaPortalFilterer{contract: contract}}, nil
}

// NominaPortal is an auto generated Go binding around an Ethereum contract.
type NominaPortal struct {
	NominaPortalCaller     // Read-only binding to the contract
	NominaPortalTransactor // Write-only binding to the contract
	NominaPortalFilterer   // Log filterer for contract events
}

// NominaPortalCaller is an auto generated read-only Go binding around an Ethereum contract.
type NominaPortalCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NominaPortalTransactor is an auto generated write-only Go binding around an Ethereum contract.
type NominaPortalTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NominaPortalFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type NominaPortalFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NominaPortalSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NominaPortalSession struct {
	Contract     *NominaPortal     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// NominaPortalCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NominaPortalCallerSession struct {
	Contract *NominaPortalCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// NominaPortalTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NominaPortalTransactorSession struct {
	Contract     *NominaPortalTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// NominaPortalRaw is an auto generated low-level Go binding around an Ethereum contract.
type NominaPortalRaw struct {
	Contract *NominaPortal // Generic contract binding to access the raw methods on
}

// NominaPortalCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NominaPortalCallerRaw struct {
	Contract *NominaPortalCaller // Generic read-only contract binding to access the raw methods on
}

// NominaPortalTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NominaPortalTransactorRaw struct {
	Contract *NominaPortalTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNominaPortal creates a new instance of NominaPortal, bound to a specific deployed contract.
func NewNominaPortal(address common.Address, backend bind.ContractBackend) (*NominaPortal, error) {
	contract, err := bindNominaPortal(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &NominaPortal{NominaPortalCaller: NominaPortalCaller{contract: contract}, NominaPortalTransactor: NominaPortalTransactor{contract: contract}, NominaPortalFilterer: NominaPortalFilterer{contract: contract}}, nil
}

// NewNominaPortalCaller creates a new read-only instance of NominaPortal, bound to a specific deployed contract.
func NewNominaPortalCaller(address common.Address, caller bind.ContractCaller) (*NominaPortalCaller, error) {
	contract, err := bindNominaPortal(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NominaPortalCaller{contract: contract}, nil
}

// NewNominaPortalTransactor creates a new write-only instance of NominaPortal, bound to a specific deployed contract.
func NewNominaPortalTransactor(address common.Address, transactor bind.ContractTransactor) (*NominaPortalTransactor, error) {
	contract, err := bindNominaPortal(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NominaPortalTransactor{contract: contract}, nil
}

// NewNominaPortalFilterer creates a new log filterer instance of NominaPortal, bound to a specific deployed contract.
func NewNominaPortalFilterer(address common.Address, filterer bind.ContractFilterer) (*NominaPortalFilterer, error) {
	contract, err := bindNominaPortal(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NominaPortalFilterer{contract: contract}, nil
}

// bindNominaPortal binds a generic wrapper to an already deployed contract.
func bindNominaPortal(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := NominaPortalMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NominaPortal *NominaPortalRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NominaPortal.Contract.NominaPortalCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NominaPortal *NominaPortalRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NominaPortal.Contract.NominaPortalTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NominaPortal *NominaPortalRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NominaPortal.Contract.NominaPortalTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NominaPortal *NominaPortalCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NominaPortal.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NominaPortal *NominaPortalTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NominaPortal.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NominaPortal *NominaPortalTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NominaPortal.Contract.contract.Transact(opts, method, params...)
}

// ActionXCall is a free data retrieval call binding the contract method 0xb2b2f5bd.
//
// Solidity: function ActionXCall() view returns(bytes32)
func (_NominaPortal *NominaPortalCaller) ActionXCall(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _NominaPortal.contract.Call(opts, &out, "ActionXCall")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ActionXCall is a free data retrieval call binding the contract method 0xb2b2f5bd.
//
// Solidity: function ActionXCall() view returns(bytes32)
func (_NominaPortal *NominaPortalSession) ActionXCall() ([32]byte, error) {
	return _NominaPortal.Contract.ActionXCall(&_NominaPortal.CallOpts)
}

// ActionXCall is a free data retrieval call binding the contract method 0xb2b2f5bd.
//
// Solidity: function ActionXCall() view returns(bytes32)
func (_NominaPortal *NominaPortalCallerSession) ActionXCall() ([32]byte, error) {
	return _NominaPortal.Contract.ActionXCall(&_NominaPortal.CallOpts)
}

// ActionXSubmit is a free data retrieval call binding the contract method 0xa32eb7c6.
//
// Solidity: function ActionXSubmit() view returns(bytes32)
func (_NominaPortal *NominaPortalCaller) ActionXSubmit(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _NominaPortal.contract.Call(opts, &out, "ActionXSubmit")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ActionXSubmit is a free data retrieval call binding the contract method 0xa32eb7c6.
//
// Solidity: function ActionXSubmit() view returns(bytes32)
func (_NominaPortal *NominaPortalSession) ActionXSubmit() ([32]byte, error) {
	return _NominaPortal.Contract.ActionXSubmit(&_NominaPortal.CallOpts)
}

// ActionXSubmit is a free data retrieval call binding the contract method 0xa32eb7c6.
//
// Solidity: function ActionXSubmit() view returns(bytes32)
func (_NominaPortal *NominaPortalCallerSession) ActionXSubmit() ([32]byte, error) {
	return _NominaPortal.Contract.ActionXSubmit(&_NominaPortal.CallOpts)
}

// KeyPauseAll is a free data retrieval call binding the contract method 0xa10ac97a.
//
// Solidity: function KeyPauseAll() view returns(bytes32)
func (_NominaPortal *NominaPortalCaller) KeyPauseAll(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _NominaPortal.contract.Call(opts, &out, "KeyPauseAll")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// KeyPauseAll is a free data retrieval call binding the contract method 0xa10ac97a.
//
// Solidity: function KeyPauseAll() view returns(bytes32)
func (_NominaPortal *NominaPortalSession) KeyPauseAll() ([32]byte, error) {
	return _NominaPortal.Contract.KeyPauseAll(&_NominaPortal.CallOpts)
}

// KeyPauseAll is a free data retrieval call binding the contract method 0xa10ac97a.
//
// Solidity: function KeyPauseAll() view returns(bytes32)
func (_NominaPortal *NominaPortalCallerSession) KeyPauseAll() ([32]byte, error) {
	return _NominaPortal.Contract.KeyPauseAll(&_NominaPortal.CallOpts)
}

// XSubQuorumDenominator is a free data retrieval call binding the contract method 0x06c3dc5f.
//
// Solidity: function XSubQuorumDenominator() view returns(uint8)
func (_NominaPortal *NominaPortalCaller) XSubQuorumDenominator(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _NominaPortal.contract.Call(opts, &out, "XSubQuorumDenominator")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// XSubQuorumDenominator is a free data retrieval call binding the contract method 0x06c3dc5f.
//
// Solidity: function XSubQuorumDenominator() view returns(uint8)
func (_NominaPortal *NominaPortalSession) XSubQuorumDenominator() (uint8, error) {
	return _NominaPortal.Contract.XSubQuorumDenominator(&_NominaPortal.CallOpts)
}

// XSubQuorumDenominator is a free data retrieval call binding the contract method 0x06c3dc5f.
//
// Solidity: function XSubQuorumDenominator() view returns(uint8)
func (_NominaPortal *NominaPortalCallerSession) XSubQuorumDenominator() (uint8, error) {
	return _NominaPortal.Contract.XSubQuorumDenominator(&_NominaPortal.CallOpts)
}

// XSubQuorumNumerator is a free data retrieval call binding the contract method 0x0360d20f.
//
// Solidity: function XSubQuorumNumerator() view returns(uint8)
func (_NominaPortal *NominaPortalCaller) XSubQuorumNumerator(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _NominaPortal.contract.Call(opts, &out, "XSubQuorumNumerator")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// XSubQuorumNumerator is a free data retrieval call binding the contract method 0x0360d20f.
//
// Solidity: function XSubQuorumNumerator() view returns(uint8)
func (_NominaPortal *NominaPortalSession) XSubQuorumNumerator() (uint8, error) {
	return _NominaPortal.Contract.XSubQuorumNumerator(&_NominaPortal.CallOpts)
}

// XSubQuorumNumerator is a free data retrieval call binding the contract method 0x0360d20f.
//
// Solidity: function XSubQuorumNumerator() view returns(uint8)
func (_NominaPortal *NominaPortalCallerSession) XSubQuorumNumerator() (uint8, error) {
	return _NominaPortal.Contract.XSubQuorumNumerator(&_NominaPortal.CallOpts)
}

// ChainId is a free data retrieval call binding the contract method 0x9a8a0592.
//
// Solidity: function chainId() view returns(uint64)
func (_NominaPortal *NominaPortalCaller) ChainId(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _NominaPortal.contract.Call(opts, &out, "chainId")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// ChainId is a free data retrieval call binding the contract method 0x9a8a0592.
//
// Solidity: function chainId() view returns(uint64)
func (_NominaPortal *NominaPortalSession) ChainId() (uint64, error) {
	return _NominaPortal.Contract.ChainId(&_NominaPortal.CallOpts)
}

// ChainId is a free data retrieval call binding the contract method 0x9a8a0592.
//
// Solidity: function chainId() view returns(uint64)
func (_NominaPortal *NominaPortalCallerSession) ChainId() (uint64, error) {
	return _NominaPortal.Contract.ChainId(&_NominaPortal.CallOpts)
}

// FeeFor is a free data retrieval call binding the contract method 0x8dd9523c.
//
// Solidity: function feeFor(uint64 destChainId, bytes data, uint64 gasLimit) view returns(uint256)
func (_NominaPortal *NominaPortalCaller) FeeFor(opts *bind.CallOpts, destChainId uint64, data []byte, gasLimit uint64) (*big.Int, error) {
	var out []interface{}
	err := _NominaPortal.contract.Call(opts, &out, "feeFor", destChainId, data, gasLimit)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FeeFor is a free data retrieval call binding the contract method 0x8dd9523c.
//
// Solidity: function feeFor(uint64 destChainId, bytes data, uint64 gasLimit) view returns(uint256)
func (_NominaPortal *NominaPortalSession) FeeFor(destChainId uint64, data []byte, gasLimit uint64) (*big.Int, error) {
	return _NominaPortal.Contract.FeeFor(&_NominaPortal.CallOpts, destChainId, data, gasLimit)
}

// FeeFor is a free data retrieval call binding the contract method 0x8dd9523c.
//
// Solidity: function feeFor(uint64 destChainId, bytes data, uint64 gasLimit) view returns(uint256)
func (_NominaPortal *NominaPortalCallerSession) FeeFor(destChainId uint64, data []byte, gasLimit uint64) (*big.Int, error) {
	return _NominaPortal.Contract.FeeFor(&_NominaPortal.CallOpts, destChainId, data, gasLimit)
}

// FeeOracle is a free data retrieval call binding the contract method 0x500b19e7.
//
// Solidity: function feeOracle() view returns(address)
func (_NominaPortal *NominaPortalCaller) FeeOracle(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NominaPortal.contract.Call(opts, &out, "feeOracle")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeOracle is a free data retrieval call binding the contract method 0x500b19e7.
//
// Solidity: function feeOracle() view returns(address)
func (_NominaPortal *NominaPortalSession) FeeOracle() (common.Address, error) {
	return _NominaPortal.Contract.FeeOracle(&_NominaPortal.CallOpts)
}

// FeeOracle is a free data retrieval call binding the contract method 0x500b19e7.
//
// Solidity: function feeOracle() view returns(address)
func (_NominaPortal *NominaPortalCallerSession) FeeOracle() (common.Address, error) {
	return _NominaPortal.Contract.FeeOracle(&_NominaPortal.CallOpts)
}

// InXBlockOffset is a free data retrieval call binding the contract method 0x3fd3b15e.
//
// Solidity: function inXBlockOffset(uint64 , uint64 ) view returns(uint64)
func (_NominaPortal *NominaPortalCaller) InXBlockOffset(opts *bind.CallOpts, arg0 uint64, arg1 uint64) (uint64, error) {
	var out []interface{}
	err := _NominaPortal.contract.Call(opts, &out, "inXBlockOffset", arg0, arg1)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// InXBlockOffset is a free data retrieval call binding the contract method 0x3fd3b15e.
//
// Solidity: function inXBlockOffset(uint64 , uint64 ) view returns(uint64)
func (_NominaPortal *NominaPortalSession) InXBlockOffset(arg0 uint64, arg1 uint64) (uint64, error) {
	return _NominaPortal.Contract.InXBlockOffset(&_NominaPortal.CallOpts, arg0, arg1)
}

// InXBlockOffset is a free data retrieval call binding the contract method 0x3fd3b15e.
//
// Solidity: function inXBlockOffset(uint64 , uint64 ) view returns(uint64)
func (_NominaPortal *NominaPortalCallerSession) InXBlockOffset(arg0 uint64, arg1 uint64) (uint64, error) {
	return _NominaPortal.Contract.InXBlockOffset(&_NominaPortal.CallOpts, arg0, arg1)
}

// InXMsgOffset is a free data retrieval call binding the contract method 0xd051c97d.
//
// Solidity: function inXMsgOffset(uint64 , uint64 ) view returns(uint64)
func (_NominaPortal *NominaPortalCaller) InXMsgOffset(opts *bind.CallOpts, arg0 uint64, arg1 uint64) (uint64, error) {
	var out []interface{}
	err := _NominaPortal.contract.Call(opts, &out, "inXMsgOffset", arg0, arg1)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// InXMsgOffset is a free data retrieval call binding the contract method 0xd051c97d.
//
// Solidity: function inXMsgOffset(uint64 , uint64 ) view returns(uint64)
func (_NominaPortal *NominaPortalSession) InXMsgOffset(arg0 uint64, arg1 uint64) (uint64, error) {
	return _NominaPortal.Contract.InXMsgOffset(&_NominaPortal.CallOpts, arg0, arg1)
}

// InXMsgOffset is a free data retrieval call binding the contract method 0xd051c97d.
//
// Solidity: function inXMsgOffset(uint64 , uint64 ) view returns(uint64)
func (_NominaPortal *NominaPortalCallerSession) InXMsgOffset(arg0 uint64, arg1 uint64) (uint64, error) {
	return _NominaPortal.Contract.InXMsgOffset(&_NominaPortal.CallOpts, arg0, arg1)
}

// IsPaused is a free data retrieval call binding the contract method 0x241b71bb.
//
// Solidity: function isPaused(bytes32 actionId) view returns(bool)
func (_NominaPortal *NominaPortalCaller) IsPaused(opts *bind.CallOpts, actionId [32]byte) (bool, error) {
	var out []interface{}
	err := _NominaPortal.contract.Call(opts, &out, "isPaused", actionId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPaused is a free data retrieval call binding the contract method 0x241b71bb.
//
// Solidity: function isPaused(bytes32 actionId) view returns(bool)
func (_NominaPortal *NominaPortalSession) IsPaused(actionId [32]byte) (bool, error) {
	return _NominaPortal.Contract.IsPaused(&_NominaPortal.CallOpts, actionId)
}

// IsPaused is a free data retrieval call binding the contract method 0x241b71bb.
//
// Solidity: function isPaused(bytes32 actionId) view returns(bool)
func (_NominaPortal *NominaPortalCallerSession) IsPaused(actionId [32]byte) (bool, error) {
	return _NominaPortal.Contract.IsPaused(&_NominaPortal.CallOpts, actionId)
}

// IsPaused0 is a free data retrieval call binding the contract method 0x461ab488.
//
// Solidity: function isPaused(bytes32 actionId, uint64 chainId_) view returns(bool)
func (_NominaPortal *NominaPortalCaller) IsPaused0(opts *bind.CallOpts, actionId [32]byte, chainId_ uint64) (bool, error) {
	var out []interface{}
	err := _NominaPortal.contract.Call(opts, &out, "isPaused0", actionId, chainId_)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPaused0 is a free data retrieval call binding the contract method 0x461ab488.
//
// Solidity: function isPaused(bytes32 actionId, uint64 chainId_) view returns(bool)
func (_NominaPortal *NominaPortalSession) IsPaused0(actionId [32]byte, chainId_ uint64) (bool, error) {
	return _NominaPortal.Contract.IsPaused0(&_NominaPortal.CallOpts, actionId, chainId_)
}

// IsPaused0 is a free data retrieval call binding the contract method 0x461ab488.
//
// Solidity: function isPaused(bytes32 actionId, uint64 chainId_) view returns(bool)
func (_NominaPortal *NominaPortalCallerSession) IsPaused0(actionId [32]byte, chainId_ uint64) (bool, error) {
	return _NominaPortal.Contract.IsPaused0(&_NominaPortal.CallOpts, actionId, chainId_)
}

// IsPaused1 is a free data retrieval call binding the contract method 0xb187bd26.
//
// Solidity: function isPaused() view returns(bool)
func (_NominaPortal *NominaPortalCaller) IsPaused1(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _NominaPortal.contract.Call(opts, &out, "isPaused1")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPaused1 is a free data retrieval call binding the contract method 0xb187bd26.
//
// Solidity: function isPaused() view returns(bool)
func (_NominaPortal *NominaPortalSession) IsPaused1() (bool, error) {
	return _NominaPortal.Contract.IsPaused1(&_NominaPortal.CallOpts)
}

// IsPaused1 is a free data retrieval call binding the contract method 0xb187bd26.
//
// Solidity: function isPaused() view returns(bool)
func (_NominaPortal *NominaPortalCallerSession) IsPaused1() (bool, error) {
	return _NominaPortal.Contract.IsPaused1(&_NominaPortal.CallOpts)
}

// IsSupportedDest is a free data retrieval call binding the contract method 0x24278bbe.
//
// Solidity: function isSupportedDest(uint64 ) view returns(bool)
func (_NominaPortal *NominaPortalCaller) IsSupportedDest(opts *bind.CallOpts, arg0 uint64) (bool, error) {
	var out []interface{}
	err := _NominaPortal.contract.Call(opts, &out, "isSupportedDest", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsSupportedDest is a free data retrieval call binding the contract method 0x24278bbe.
//
// Solidity: function isSupportedDest(uint64 ) view returns(bool)
func (_NominaPortal *NominaPortalSession) IsSupportedDest(arg0 uint64) (bool, error) {
	return _NominaPortal.Contract.IsSupportedDest(&_NominaPortal.CallOpts, arg0)
}

// IsSupportedDest is a free data retrieval call binding the contract method 0x24278bbe.
//
// Solidity: function isSupportedDest(uint64 ) view returns(bool)
func (_NominaPortal *NominaPortalCallerSession) IsSupportedDest(arg0 uint64) (bool, error) {
	return _NominaPortal.Contract.IsSupportedDest(&_NominaPortal.CallOpts, arg0)
}

// IsSupportedShard is a free data retrieval call binding the contract method 0xaaf1bc97.
//
// Solidity: function isSupportedShard(uint64 ) view returns(bool)
func (_NominaPortal *NominaPortalCaller) IsSupportedShard(opts *bind.CallOpts, arg0 uint64) (bool, error) {
	var out []interface{}
	err := _NominaPortal.contract.Call(opts, &out, "isSupportedShard", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsSupportedShard is a free data retrieval call binding the contract method 0xaaf1bc97.
//
// Solidity: function isSupportedShard(uint64 ) view returns(bool)
func (_NominaPortal *NominaPortalSession) IsSupportedShard(arg0 uint64) (bool, error) {
	return _NominaPortal.Contract.IsSupportedShard(&_NominaPortal.CallOpts, arg0)
}

// IsSupportedShard is a free data retrieval call binding the contract method 0xaaf1bc97.
//
// Solidity: function isSupportedShard(uint64 ) view returns(bool)
func (_NominaPortal *NominaPortalCallerSession) IsSupportedShard(arg0 uint64) (bool, error) {
	return _NominaPortal.Contract.IsSupportedShard(&_NominaPortal.CallOpts, arg0)
}

// IsXCall is a free data retrieval call binding the contract method 0x55e2448e.
//
// Solidity: function isXCall() view returns(bool)
func (_NominaPortal *NominaPortalCaller) IsXCall(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _NominaPortal.contract.Call(opts, &out, "isXCall")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsXCall is a free data retrieval call binding the contract method 0x55e2448e.
//
// Solidity: function isXCall() view returns(bool)
func (_NominaPortal *NominaPortalSession) IsXCall() (bool, error) {
	return _NominaPortal.Contract.IsXCall(&_NominaPortal.CallOpts)
}

// IsXCall is a free data retrieval call binding the contract method 0x55e2448e.
//
// Solidity: function isXCall() view returns(bool)
func (_NominaPortal *NominaPortalCallerSession) IsXCall() (bool, error) {
	return _NominaPortal.Contract.IsXCall(&_NominaPortal.CallOpts)
}

// LatestValSetId is a free data retrieval call binding the contract method 0xf45cc7b8.
//
// Solidity: function latestValSetId() view returns(uint64)
func (_NominaPortal *NominaPortalCaller) LatestValSetId(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _NominaPortal.contract.Call(opts, &out, "latestValSetId")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// LatestValSetId is a free data retrieval call binding the contract method 0xf45cc7b8.
//
// Solidity: function latestValSetId() view returns(uint64)
func (_NominaPortal *NominaPortalSession) LatestValSetId() (uint64, error) {
	return _NominaPortal.Contract.LatestValSetId(&_NominaPortal.CallOpts)
}

// LatestValSetId is a free data retrieval call binding the contract method 0xf45cc7b8.
//
// Solidity: function latestValSetId() view returns(uint64)
func (_NominaPortal *NominaPortalCallerSession) LatestValSetId() (uint64, error) {
	return _NominaPortal.Contract.LatestValSetId(&_NominaPortal.CallOpts)
}

// Network is a free data retrieval call binding the contract method 0x6739afca.
//
// Solidity: function network() view returns((uint64,uint64[])[])
func (_NominaPortal *NominaPortalCaller) Network(opts *bind.CallOpts) ([]XTypesChain, error) {
	var out []interface{}
	err := _NominaPortal.contract.Call(opts, &out, "network")

	if err != nil {
		return *new([]XTypesChain), err
	}

	out0 := *abi.ConvertType(out[0], new([]XTypesChain)).(*[]XTypesChain)

	return out0, err

}

// Network is a free data retrieval call binding the contract method 0x6739afca.
//
// Solidity: function network() view returns((uint64,uint64[])[])
func (_NominaPortal *NominaPortalSession) Network() ([]XTypesChain, error) {
	return _NominaPortal.Contract.Network(&_NominaPortal.CallOpts)
}

// Network is a free data retrieval call binding the contract method 0x6739afca.
//
// Solidity: function network() view returns((uint64,uint64[])[])
func (_NominaPortal *NominaPortalCallerSession) Network() ([]XTypesChain, error) {
	return _NominaPortal.Contract.Network(&_NominaPortal.CallOpts)
}

// NominaCChainId is a free data retrieval call binding the contract method 0x87320ac2.
//
// Solidity: function nominaCChainId() view returns(uint64)
func (_NominaPortal *NominaPortalCaller) NominaCChainId(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _NominaPortal.contract.Call(opts, &out, "nominaCChainId")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// NominaCChainId is a free data retrieval call binding the contract method 0x87320ac2.
//
// Solidity: function nominaCChainId() view returns(uint64)
func (_NominaPortal *NominaPortalSession) NominaCChainId() (uint64, error) {
	return _NominaPortal.Contract.NominaCChainId(&_NominaPortal.CallOpts)
}

// NominaCChainId is a free data retrieval call binding the contract method 0x87320ac2.
//
// Solidity: function nominaCChainId() view returns(uint64)
func (_NominaPortal *NominaPortalCallerSession) NominaCChainId() (uint64, error) {
	return _NominaPortal.Contract.NominaCChainId(&_NominaPortal.CallOpts)
}

// NominaChainId is a free data retrieval call binding the contract method 0x4cc39088.
//
// Solidity: function nominaChainId() view returns(uint64)
func (_NominaPortal *NominaPortalCaller) NominaChainId(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _NominaPortal.contract.Call(opts, &out, "nominaChainId")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// NominaChainId is a free data retrieval call binding the contract method 0x4cc39088.
//
// Solidity: function nominaChainId() view returns(uint64)
func (_NominaPortal *NominaPortalSession) NominaChainId() (uint64, error) {
	return _NominaPortal.Contract.NominaChainId(&_NominaPortal.CallOpts)
}

// NominaChainId is a free data retrieval call binding the contract method 0x4cc39088.
//
// Solidity: function nominaChainId() view returns(uint64)
func (_NominaPortal *NominaPortalCallerSession) NominaChainId() (uint64, error) {
	return _NominaPortal.Contract.NominaChainId(&_NominaPortal.CallOpts)
}

// OmniCChainId is a free data retrieval call binding the contract method 0x36d21912.
//
// Solidity: function omniCChainId() view returns(uint64)
func (_NominaPortal *NominaPortalCaller) OmniCChainId(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _NominaPortal.contract.Call(opts, &out, "omniCChainId")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// OmniCChainId is a free data retrieval call binding the contract method 0x36d21912.
//
// Solidity: function omniCChainId() view returns(uint64)
func (_NominaPortal *NominaPortalSession) OmniCChainId() (uint64, error) {
	return _NominaPortal.Contract.OmniCChainId(&_NominaPortal.CallOpts)
}

// OmniCChainId is a free data retrieval call binding the contract method 0x36d21912.
//
// Solidity: function omniCChainId() view returns(uint64)
func (_NominaPortal *NominaPortalCallerSession) OmniCChainId() (uint64, error) {
	return _NominaPortal.Contract.OmniCChainId(&_NominaPortal.CallOpts)
}

// OmniChainId is a free data retrieval call binding the contract method 0x110ff5f1.
//
// Solidity: function omniChainId() view returns(uint64)
func (_NominaPortal *NominaPortalCaller) OmniChainId(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _NominaPortal.contract.Call(opts, &out, "omniChainId")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// OmniChainId is a free data retrieval call binding the contract method 0x110ff5f1.
//
// Solidity: function omniChainId() view returns(uint64)
func (_NominaPortal *NominaPortalSession) OmniChainId() (uint64, error) {
	return _NominaPortal.Contract.OmniChainId(&_NominaPortal.CallOpts)
}

// OmniChainId is a free data retrieval call binding the contract method 0x110ff5f1.
//
// Solidity: function omniChainId() view returns(uint64)
func (_NominaPortal *NominaPortalCallerSession) OmniChainId() (uint64, error) {
	return _NominaPortal.Contract.OmniChainId(&_NominaPortal.CallOpts)
}

// OutXMsgOffset is a free data retrieval call binding the contract method 0x3aa87330.
//
// Solidity: function outXMsgOffset(uint64 , uint64 ) view returns(uint64)
func (_NominaPortal *NominaPortalCaller) OutXMsgOffset(opts *bind.CallOpts, arg0 uint64, arg1 uint64) (uint64, error) {
	var out []interface{}
	err := _NominaPortal.contract.Call(opts, &out, "outXMsgOffset", arg0, arg1)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// OutXMsgOffset is a free data retrieval call binding the contract method 0x3aa87330.
//
// Solidity: function outXMsgOffset(uint64 , uint64 ) view returns(uint64)
func (_NominaPortal *NominaPortalSession) OutXMsgOffset(arg0 uint64, arg1 uint64) (uint64, error) {
	return _NominaPortal.Contract.OutXMsgOffset(&_NominaPortal.CallOpts, arg0, arg1)
}

// OutXMsgOffset is a free data retrieval call binding the contract method 0x3aa87330.
//
// Solidity: function outXMsgOffset(uint64 , uint64 ) view returns(uint64)
func (_NominaPortal *NominaPortalCallerSession) OutXMsgOffset(arg0 uint64, arg1 uint64) (uint64, error) {
	return _NominaPortal.Contract.OutXMsgOffset(&_NominaPortal.CallOpts, arg0, arg1)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_NominaPortal *NominaPortalCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _NominaPortal.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_NominaPortal *NominaPortalSession) Owner() (common.Address, error) {
	return _NominaPortal.Contract.Owner(&_NominaPortal.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_NominaPortal *NominaPortalCallerSession) Owner() (common.Address, error) {
	return _NominaPortal.Contract.Owner(&_NominaPortal.CallOpts)
}

// ValSet is a free data retrieval call binding the contract method 0x57542050.
//
// Solidity: function valSet(uint64 , address ) view returns(uint64)
func (_NominaPortal *NominaPortalCaller) ValSet(opts *bind.CallOpts, arg0 uint64, arg1 common.Address) (uint64, error) {
	var out []interface{}
	err := _NominaPortal.contract.Call(opts, &out, "valSet", arg0, arg1)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// ValSet is a free data retrieval call binding the contract method 0x57542050.
//
// Solidity: function valSet(uint64 , address ) view returns(uint64)
func (_NominaPortal *NominaPortalSession) ValSet(arg0 uint64, arg1 common.Address) (uint64, error) {
	return _NominaPortal.Contract.ValSet(&_NominaPortal.CallOpts, arg0, arg1)
}

// ValSet is a free data retrieval call binding the contract method 0x57542050.
//
// Solidity: function valSet(uint64 , address ) view returns(uint64)
func (_NominaPortal *NominaPortalCallerSession) ValSet(arg0 uint64, arg1 common.Address) (uint64, error) {
	return _NominaPortal.Contract.ValSet(&_NominaPortal.CallOpts, arg0, arg1)
}

// ValSetTotalPower is a free data retrieval call binding the contract method 0xafe8af9c.
//
// Solidity: function valSetTotalPower(uint64 ) view returns(uint64)
func (_NominaPortal *NominaPortalCaller) ValSetTotalPower(opts *bind.CallOpts, arg0 uint64) (uint64, error) {
	var out []interface{}
	err := _NominaPortal.contract.Call(opts, &out, "valSetTotalPower", arg0)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// ValSetTotalPower is a free data retrieval call binding the contract method 0xafe8af9c.
//
// Solidity: function valSetTotalPower(uint64 ) view returns(uint64)
func (_NominaPortal *NominaPortalSession) ValSetTotalPower(arg0 uint64) (uint64, error) {
	return _NominaPortal.Contract.ValSetTotalPower(&_NominaPortal.CallOpts, arg0)
}

// ValSetTotalPower is a free data retrieval call binding the contract method 0xafe8af9c.
//
// Solidity: function valSetTotalPower(uint64 ) view returns(uint64)
func (_NominaPortal *NominaPortalCallerSession) ValSetTotalPower(arg0 uint64) (uint64, error) {
	return _NominaPortal.Contract.ValSetTotalPower(&_NominaPortal.CallOpts, arg0)
}

// Xmsg is a free data retrieval call binding the contract method 0x2f32700e.
//
// Solidity: function xmsg() view returns((uint64,address))
func (_NominaPortal *NominaPortalCaller) Xmsg(opts *bind.CallOpts) (XTypesMsgContext, error) {
	var out []interface{}
	err := _NominaPortal.contract.Call(opts, &out, "xmsg")

	if err != nil {
		return *new(XTypesMsgContext), err
	}

	out0 := *abi.ConvertType(out[0], new(XTypesMsgContext)).(*XTypesMsgContext)

	return out0, err

}

// Xmsg is a free data retrieval call binding the contract method 0x2f32700e.
//
// Solidity: function xmsg() view returns((uint64,address))
func (_NominaPortal *NominaPortalSession) Xmsg() (XTypesMsgContext, error) {
	return _NominaPortal.Contract.Xmsg(&_NominaPortal.CallOpts)
}

// Xmsg is a free data retrieval call binding the contract method 0x2f32700e.
//
// Solidity: function xmsg() view returns((uint64,address))
func (_NominaPortal *NominaPortalCallerSession) Xmsg() (XTypesMsgContext, error) {
	return _NominaPortal.Contract.Xmsg(&_NominaPortal.CallOpts)
}

// XmsgMaxDataSize is a free data retrieval call binding the contract method 0xb4d5afd1.
//
// Solidity: function xmsgMaxDataSize() view returns(uint16)
func (_NominaPortal *NominaPortalCaller) XmsgMaxDataSize(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _NominaPortal.contract.Call(opts, &out, "xmsgMaxDataSize")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// XmsgMaxDataSize is a free data retrieval call binding the contract method 0xb4d5afd1.
//
// Solidity: function xmsgMaxDataSize() view returns(uint16)
func (_NominaPortal *NominaPortalSession) XmsgMaxDataSize() (uint16, error) {
	return _NominaPortal.Contract.XmsgMaxDataSize(&_NominaPortal.CallOpts)
}

// XmsgMaxDataSize is a free data retrieval call binding the contract method 0xb4d5afd1.
//
// Solidity: function xmsgMaxDataSize() view returns(uint16)
func (_NominaPortal *NominaPortalCallerSession) XmsgMaxDataSize() (uint16, error) {
	return _NominaPortal.Contract.XmsgMaxDataSize(&_NominaPortal.CallOpts)
}

// XmsgMaxGasLimit is a free data retrieval call binding the contract method 0xcf84c818.
//
// Solidity: function xmsgMaxGasLimit() view returns(uint64)
func (_NominaPortal *NominaPortalCaller) XmsgMaxGasLimit(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _NominaPortal.contract.Call(opts, &out, "xmsgMaxGasLimit")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// XmsgMaxGasLimit is a free data retrieval call binding the contract method 0xcf84c818.
//
// Solidity: function xmsgMaxGasLimit() view returns(uint64)
func (_NominaPortal *NominaPortalSession) XmsgMaxGasLimit() (uint64, error) {
	return _NominaPortal.Contract.XmsgMaxGasLimit(&_NominaPortal.CallOpts)
}

// XmsgMaxGasLimit is a free data retrieval call binding the contract method 0xcf84c818.
//
// Solidity: function xmsgMaxGasLimit() view returns(uint64)
func (_NominaPortal *NominaPortalCallerSession) XmsgMaxGasLimit() (uint64, error) {
	return _NominaPortal.Contract.XmsgMaxGasLimit(&_NominaPortal.CallOpts)
}

// XmsgMinGasLimit is a free data retrieval call binding the contract method 0x78fe5307.
//
// Solidity: function xmsgMinGasLimit() view returns(uint64)
func (_NominaPortal *NominaPortalCaller) XmsgMinGasLimit(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _NominaPortal.contract.Call(opts, &out, "xmsgMinGasLimit")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// XmsgMinGasLimit is a free data retrieval call binding the contract method 0x78fe5307.
//
// Solidity: function xmsgMinGasLimit() view returns(uint64)
func (_NominaPortal *NominaPortalSession) XmsgMinGasLimit() (uint64, error) {
	return _NominaPortal.Contract.XmsgMinGasLimit(&_NominaPortal.CallOpts)
}

// XmsgMinGasLimit is a free data retrieval call binding the contract method 0x78fe5307.
//
// Solidity: function xmsgMinGasLimit() view returns(uint64)
func (_NominaPortal *NominaPortalCallerSession) XmsgMinGasLimit() (uint64, error) {
	return _NominaPortal.Contract.XmsgMinGasLimit(&_NominaPortal.CallOpts)
}

// XreceiptMaxErrorSize is a free data retrieval call binding the contract method 0xc26dfc05.
//
// Solidity: function xreceiptMaxErrorSize() view returns(uint16)
func (_NominaPortal *NominaPortalCaller) XreceiptMaxErrorSize(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _NominaPortal.contract.Call(opts, &out, "xreceiptMaxErrorSize")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// XreceiptMaxErrorSize is a free data retrieval call binding the contract method 0xc26dfc05.
//
// Solidity: function xreceiptMaxErrorSize() view returns(uint16)
func (_NominaPortal *NominaPortalSession) XreceiptMaxErrorSize() (uint16, error) {
	return _NominaPortal.Contract.XreceiptMaxErrorSize(&_NominaPortal.CallOpts)
}

// XreceiptMaxErrorSize is a free data retrieval call binding the contract method 0xc26dfc05.
//
// Solidity: function xreceiptMaxErrorSize() view returns(uint16)
func (_NominaPortal *NominaPortalCallerSession) XreceiptMaxErrorSize() (uint16, error) {
	return _NominaPortal.Contract.XreceiptMaxErrorSize(&_NominaPortal.CallOpts)
}

// XsubValsetCutoff is a free data retrieval call binding the contract method 0x49cc3bf6.
//
// Solidity: function xsubValsetCutoff() view returns(uint8)
func (_NominaPortal *NominaPortalCaller) XsubValsetCutoff(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _NominaPortal.contract.Call(opts, &out, "xsubValsetCutoff")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// XsubValsetCutoff is a free data retrieval call binding the contract method 0x49cc3bf6.
//
// Solidity: function xsubValsetCutoff() view returns(uint8)
func (_NominaPortal *NominaPortalSession) XsubValsetCutoff() (uint8, error) {
	return _NominaPortal.Contract.XsubValsetCutoff(&_NominaPortal.CallOpts)
}

// XsubValsetCutoff is a free data retrieval call binding the contract method 0x49cc3bf6.
//
// Solidity: function xsubValsetCutoff() view returns(uint8)
func (_NominaPortal *NominaPortalCallerSession) XsubValsetCutoff() (uint8, error) {
	return _NominaPortal.Contract.XsubValsetCutoff(&_NominaPortal.CallOpts)
}

// AddValidatorSet is a paid mutator transaction binding the contract method 0x8532eb9f.
//
// Solidity: function addValidatorSet(uint64 valSetId, (address,uint64)[] validators) returns()
func (_NominaPortal *NominaPortalTransactor) AddValidatorSet(opts *bind.TransactOpts, valSetId uint64, validators []XTypesValidator) (*types.Transaction, error) {
	return _NominaPortal.contract.Transact(opts, "addValidatorSet", valSetId, validators)
}

// AddValidatorSet is a paid mutator transaction binding the contract method 0x8532eb9f.
//
// Solidity: function addValidatorSet(uint64 valSetId, (address,uint64)[] validators) returns()
func (_NominaPortal *NominaPortalSession) AddValidatorSet(valSetId uint64, validators []XTypesValidator) (*types.Transaction, error) {
	return _NominaPortal.Contract.AddValidatorSet(&_NominaPortal.TransactOpts, valSetId, validators)
}

// AddValidatorSet is a paid mutator transaction binding the contract method 0x8532eb9f.
//
// Solidity: function addValidatorSet(uint64 valSetId, (address,uint64)[] validators) returns()
func (_NominaPortal *NominaPortalTransactorSession) AddValidatorSet(valSetId uint64, validators []XTypesValidator) (*types.Transaction, error) {
	return _NominaPortal.Contract.AddValidatorSet(&_NominaPortal.TransactOpts, valSetId, validators)
}

// CollectFees is a paid mutator transaction binding the contract method 0xa480ca79.
//
// Solidity: function collectFees(address to) returns()
func (_NominaPortal *NominaPortalTransactor) CollectFees(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _NominaPortal.contract.Transact(opts, "collectFees", to)
}

// CollectFees is a paid mutator transaction binding the contract method 0xa480ca79.
//
// Solidity: function collectFees(address to) returns()
func (_NominaPortal *NominaPortalSession) CollectFees(to common.Address) (*types.Transaction, error) {
	return _NominaPortal.Contract.CollectFees(&_NominaPortal.TransactOpts, to)
}

// CollectFees is a paid mutator transaction binding the contract method 0xa480ca79.
//
// Solidity: function collectFees(address to) returns()
func (_NominaPortal *NominaPortalTransactorSession) CollectFees(to common.Address) (*types.Transaction, error) {
	return _NominaPortal.Contract.CollectFees(&_NominaPortal.TransactOpts, to)
}

// Initialize is a paid mutator transaction binding the contract method 0x30632e8b.
//
// Solidity: function initialize((address,address,uint64,uint64,uint64,uint64,uint16,uint16,uint8,uint64,uint64,uint64,(address,uint64)[]) p) returns()
func (_NominaPortal *NominaPortalTransactor) Initialize(opts *bind.TransactOpts, p NominaPortalInitParams) (*types.Transaction, error) {
	return _NominaPortal.contract.Transact(opts, "initialize", p)
}

// Initialize is a paid mutator transaction binding the contract method 0x30632e8b.
//
// Solidity: function initialize((address,address,uint64,uint64,uint64,uint64,uint16,uint16,uint8,uint64,uint64,uint64,(address,uint64)[]) p) returns()
func (_NominaPortal *NominaPortalSession) Initialize(p NominaPortalInitParams) (*types.Transaction, error) {
	return _NominaPortal.Contract.Initialize(&_NominaPortal.TransactOpts, p)
}

// Initialize is a paid mutator transaction binding the contract method 0x30632e8b.
//
// Solidity: function initialize((address,address,uint64,uint64,uint64,uint64,uint16,uint16,uint8,uint64,uint64,uint64,(address,uint64)[]) p) returns()
func (_NominaPortal *NominaPortalTransactorSession) Initialize(p NominaPortalInitParams) (*types.Transaction, error) {
	return _NominaPortal.Contract.Initialize(&_NominaPortal.TransactOpts, p)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_NominaPortal *NominaPortalTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NominaPortal.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_NominaPortal *NominaPortalSession) Pause() (*types.Transaction, error) {
	return _NominaPortal.Contract.Pause(&_NominaPortal.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_NominaPortal *NominaPortalTransactorSession) Pause() (*types.Transaction, error) {
	return _NominaPortal.Contract.Pause(&_NominaPortal.TransactOpts)
}

// PauseXCall is a paid mutator transaction binding the contract method 0x83d0cbd9.
//
// Solidity: function pauseXCall() returns()
func (_NominaPortal *NominaPortalTransactor) PauseXCall(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NominaPortal.contract.Transact(opts, "pauseXCall")
}

// PauseXCall is a paid mutator transaction binding the contract method 0x83d0cbd9.
//
// Solidity: function pauseXCall() returns()
func (_NominaPortal *NominaPortalSession) PauseXCall() (*types.Transaction, error) {
	return _NominaPortal.Contract.PauseXCall(&_NominaPortal.TransactOpts)
}

// PauseXCall is a paid mutator transaction binding the contract method 0x83d0cbd9.
//
// Solidity: function pauseXCall() returns()
func (_NominaPortal *NominaPortalTransactorSession) PauseXCall() (*types.Transaction, error) {
	return _NominaPortal.Contract.PauseXCall(&_NominaPortal.TransactOpts)
}

// PauseXCallTo is a paid mutator transaction binding the contract method 0x10a5a7f7.
//
// Solidity: function pauseXCallTo(uint64 chainId_) returns()
func (_NominaPortal *NominaPortalTransactor) PauseXCallTo(opts *bind.TransactOpts, chainId_ uint64) (*types.Transaction, error) {
	return _NominaPortal.contract.Transact(opts, "pauseXCallTo", chainId_)
}

// PauseXCallTo is a paid mutator transaction binding the contract method 0x10a5a7f7.
//
// Solidity: function pauseXCallTo(uint64 chainId_) returns()
func (_NominaPortal *NominaPortalSession) PauseXCallTo(chainId_ uint64) (*types.Transaction, error) {
	return _NominaPortal.Contract.PauseXCallTo(&_NominaPortal.TransactOpts, chainId_)
}

// PauseXCallTo is a paid mutator transaction binding the contract method 0x10a5a7f7.
//
// Solidity: function pauseXCallTo(uint64 chainId_) returns()
func (_NominaPortal *NominaPortalTransactorSession) PauseXCallTo(chainId_ uint64) (*types.Transaction, error) {
	return _NominaPortal.Contract.PauseXCallTo(&_NominaPortal.TransactOpts, chainId_)
}

// PauseXSubmit is a paid mutator transaction binding the contract method 0x23dbce50.
//
// Solidity: function pauseXSubmit() returns()
func (_NominaPortal *NominaPortalTransactor) PauseXSubmit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NominaPortal.contract.Transact(opts, "pauseXSubmit")
}

// PauseXSubmit is a paid mutator transaction binding the contract method 0x23dbce50.
//
// Solidity: function pauseXSubmit() returns()
func (_NominaPortal *NominaPortalSession) PauseXSubmit() (*types.Transaction, error) {
	return _NominaPortal.Contract.PauseXSubmit(&_NominaPortal.TransactOpts)
}

// PauseXSubmit is a paid mutator transaction binding the contract method 0x23dbce50.
//
// Solidity: function pauseXSubmit() returns()
func (_NominaPortal *NominaPortalTransactorSession) PauseXSubmit() (*types.Transaction, error) {
	return _NominaPortal.Contract.PauseXSubmit(&_NominaPortal.TransactOpts)
}

// PauseXSubmitFrom is a paid mutator transaction binding the contract method 0xafe82198.
//
// Solidity: function pauseXSubmitFrom(uint64 chainId_) returns()
func (_NominaPortal *NominaPortalTransactor) PauseXSubmitFrom(opts *bind.TransactOpts, chainId_ uint64) (*types.Transaction, error) {
	return _NominaPortal.contract.Transact(opts, "pauseXSubmitFrom", chainId_)
}

// PauseXSubmitFrom is a paid mutator transaction binding the contract method 0xafe82198.
//
// Solidity: function pauseXSubmitFrom(uint64 chainId_) returns()
func (_NominaPortal *NominaPortalSession) PauseXSubmitFrom(chainId_ uint64) (*types.Transaction, error) {
	return _NominaPortal.Contract.PauseXSubmitFrom(&_NominaPortal.TransactOpts, chainId_)
}

// PauseXSubmitFrom is a paid mutator transaction binding the contract method 0xafe82198.
//
// Solidity: function pauseXSubmitFrom(uint64 chainId_) returns()
func (_NominaPortal *NominaPortalTransactorSession) PauseXSubmitFrom(chainId_ uint64) (*types.Transaction, error) {
	return _NominaPortal.Contract.PauseXSubmitFrom(&_NominaPortal.TransactOpts, chainId_)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_NominaPortal *NominaPortalTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NominaPortal.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_NominaPortal *NominaPortalSession) RenounceOwnership() (*types.Transaction, error) {
	return _NominaPortal.Contract.RenounceOwnership(&_NominaPortal.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_NominaPortal *NominaPortalTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _NominaPortal.Contract.RenounceOwnership(&_NominaPortal.TransactOpts)
}

// SetFeeOracle is a paid mutator transaction binding the contract method 0xa8a98962.
//
// Solidity: function setFeeOracle(address feeOracle_) returns()
func (_NominaPortal *NominaPortalTransactor) SetFeeOracle(opts *bind.TransactOpts, feeOracle_ common.Address) (*types.Transaction, error) {
	return _NominaPortal.contract.Transact(opts, "setFeeOracle", feeOracle_)
}

// SetFeeOracle is a paid mutator transaction binding the contract method 0xa8a98962.
//
// Solidity: function setFeeOracle(address feeOracle_) returns()
func (_NominaPortal *NominaPortalSession) SetFeeOracle(feeOracle_ common.Address) (*types.Transaction, error) {
	return _NominaPortal.Contract.SetFeeOracle(&_NominaPortal.TransactOpts, feeOracle_)
}

// SetFeeOracle is a paid mutator transaction binding the contract method 0xa8a98962.
//
// Solidity: function setFeeOracle(address feeOracle_) returns()
func (_NominaPortal *NominaPortalTransactorSession) SetFeeOracle(feeOracle_ common.Address) (*types.Transaction, error) {
	return _NominaPortal.Contract.SetFeeOracle(&_NominaPortal.TransactOpts, feeOracle_)
}

// SetInXBlockOffset is a paid mutator transaction binding the contract method 0x97b52062.
//
// Solidity: function setInXBlockOffset(uint64 sourceChainId, uint64 shardId, uint64 offset) returns()
func (_NominaPortal *NominaPortalTransactor) SetInXBlockOffset(opts *bind.TransactOpts, sourceChainId uint64, shardId uint64, offset uint64) (*types.Transaction, error) {
	return _NominaPortal.contract.Transact(opts, "setInXBlockOffset", sourceChainId, shardId, offset)
}

// SetInXBlockOffset is a paid mutator transaction binding the contract method 0x97b52062.
//
// Solidity: function setInXBlockOffset(uint64 sourceChainId, uint64 shardId, uint64 offset) returns()
func (_NominaPortal *NominaPortalSession) SetInXBlockOffset(sourceChainId uint64, shardId uint64, offset uint64) (*types.Transaction, error) {
	return _NominaPortal.Contract.SetInXBlockOffset(&_NominaPortal.TransactOpts, sourceChainId, shardId, offset)
}

// SetInXBlockOffset is a paid mutator transaction binding the contract method 0x97b52062.
//
// Solidity: function setInXBlockOffset(uint64 sourceChainId, uint64 shardId, uint64 offset) returns()
func (_NominaPortal *NominaPortalTransactorSession) SetInXBlockOffset(sourceChainId uint64, shardId uint64, offset uint64) (*types.Transaction, error) {
	return _NominaPortal.Contract.SetInXBlockOffset(&_NominaPortal.TransactOpts, sourceChainId, shardId, offset)
}

// SetInXMsgOffset is a paid mutator transaction binding the contract method 0xc4ab80bc.
//
// Solidity: function setInXMsgOffset(uint64 sourceChainId, uint64 shardId, uint64 offset) returns()
func (_NominaPortal *NominaPortalTransactor) SetInXMsgOffset(opts *bind.TransactOpts, sourceChainId uint64, shardId uint64, offset uint64) (*types.Transaction, error) {
	return _NominaPortal.contract.Transact(opts, "setInXMsgOffset", sourceChainId, shardId, offset)
}

// SetInXMsgOffset is a paid mutator transaction binding the contract method 0xc4ab80bc.
//
// Solidity: function setInXMsgOffset(uint64 sourceChainId, uint64 shardId, uint64 offset) returns()
func (_NominaPortal *NominaPortalSession) SetInXMsgOffset(sourceChainId uint64, shardId uint64, offset uint64) (*types.Transaction, error) {
	return _NominaPortal.Contract.SetInXMsgOffset(&_NominaPortal.TransactOpts, sourceChainId, shardId, offset)
}

// SetInXMsgOffset is a paid mutator transaction binding the contract method 0xc4ab80bc.
//
// Solidity: function setInXMsgOffset(uint64 sourceChainId, uint64 shardId, uint64 offset) returns()
func (_NominaPortal *NominaPortalTransactorSession) SetInXMsgOffset(sourceChainId uint64, shardId uint64, offset uint64) (*types.Transaction, error) {
	return _NominaPortal.Contract.SetInXMsgOffset(&_NominaPortal.TransactOpts, sourceChainId, shardId, offset)
}

// SetNetwork is a paid mutator transaction binding the contract method 0x1d3eb6e3.
//
// Solidity: function setNetwork((uint64,uint64[])[] network_) returns()
func (_NominaPortal *NominaPortalTransactor) SetNetwork(opts *bind.TransactOpts, network_ []XTypesChain) (*types.Transaction, error) {
	return _NominaPortal.contract.Transact(opts, "setNetwork", network_)
}

// SetNetwork is a paid mutator transaction binding the contract method 0x1d3eb6e3.
//
// Solidity: function setNetwork((uint64,uint64[])[] network_) returns()
func (_NominaPortal *NominaPortalSession) SetNetwork(network_ []XTypesChain) (*types.Transaction, error) {
	return _NominaPortal.Contract.SetNetwork(&_NominaPortal.TransactOpts, network_)
}

// SetNetwork is a paid mutator transaction binding the contract method 0x1d3eb6e3.
//
// Solidity: function setNetwork((uint64,uint64[])[] network_) returns()
func (_NominaPortal *NominaPortalTransactorSession) SetNetwork(network_ []XTypesChain) (*types.Transaction, error) {
	return _NominaPortal.Contract.SetNetwork(&_NominaPortal.TransactOpts, network_)
}

// SetXMsgMaxDataSize is a paid mutator transaction binding the contract method 0xb521466d.
//
// Solidity: function setXMsgMaxDataSize(uint16 numBytes) returns()
func (_NominaPortal *NominaPortalTransactor) SetXMsgMaxDataSize(opts *bind.TransactOpts, numBytes uint16) (*types.Transaction, error) {
	return _NominaPortal.contract.Transact(opts, "setXMsgMaxDataSize", numBytes)
}

// SetXMsgMaxDataSize is a paid mutator transaction binding the contract method 0xb521466d.
//
// Solidity: function setXMsgMaxDataSize(uint16 numBytes) returns()
func (_NominaPortal *NominaPortalSession) SetXMsgMaxDataSize(numBytes uint16) (*types.Transaction, error) {
	return _NominaPortal.Contract.SetXMsgMaxDataSize(&_NominaPortal.TransactOpts, numBytes)
}

// SetXMsgMaxDataSize is a paid mutator transaction binding the contract method 0xb521466d.
//
// Solidity: function setXMsgMaxDataSize(uint16 numBytes) returns()
func (_NominaPortal *NominaPortalTransactorSession) SetXMsgMaxDataSize(numBytes uint16) (*types.Transaction, error) {
	return _NominaPortal.Contract.SetXMsgMaxDataSize(&_NominaPortal.TransactOpts, numBytes)
}

// SetXMsgMaxGasLimit is a paid mutator transaction binding the contract method 0x36d853f9.
//
// Solidity: function setXMsgMaxGasLimit(uint64 gasLimit) returns()
func (_NominaPortal *NominaPortalTransactor) SetXMsgMaxGasLimit(opts *bind.TransactOpts, gasLimit uint64) (*types.Transaction, error) {
	return _NominaPortal.contract.Transact(opts, "setXMsgMaxGasLimit", gasLimit)
}

// SetXMsgMaxGasLimit is a paid mutator transaction binding the contract method 0x36d853f9.
//
// Solidity: function setXMsgMaxGasLimit(uint64 gasLimit) returns()
func (_NominaPortal *NominaPortalSession) SetXMsgMaxGasLimit(gasLimit uint64) (*types.Transaction, error) {
	return _NominaPortal.Contract.SetXMsgMaxGasLimit(&_NominaPortal.TransactOpts, gasLimit)
}

// SetXMsgMaxGasLimit is a paid mutator transaction binding the contract method 0x36d853f9.
//
// Solidity: function setXMsgMaxGasLimit(uint64 gasLimit) returns()
func (_NominaPortal *NominaPortalTransactorSession) SetXMsgMaxGasLimit(gasLimit uint64) (*types.Transaction, error) {
	return _NominaPortal.Contract.SetXMsgMaxGasLimit(&_NominaPortal.TransactOpts, gasLimit)
}

// SetXMsgMinGasLimit is a paid mutator transaction binding the contract method 0xbb8590ad.
//
// Solidity: function setXMsgMinGasLimit(uint64 gasLimit) returns()
func (_NominaPortal *NominaPortalTransactor) SetXMsgMinGasLimit(opts *bind.TransactOpts, gasLimit uint64) (*types.Transaction, error) {
	return _NominaPortal.contract.Transact(opts, "setXMsgMinGasLimit", gasLimit)
}

// SetXMsgMinGasLimit is a paid mutator transaction binding the contract method 0xbb8590ad.
//
// Solidity: function setXMsgMinGasLimit(uint64 gasLimit) returns()
func (_NominaPortal *NominaPortalSession) SetXMsgMinGasLimit(gasLimit uint64) (*types.Transaction, error) {
	return _NominaPortal.Contract.SetXMsgMinGasLimit(&_NominaPortal.TransactOpts, gasLimit)
}

// SetXMsgMinGasLimit is a paid mutator transaction binding the contract method 0xbb8590ad.
//
// Solidity: function setXMsgMinGasLimit(uint64 gasLimit) returns()
func (_NominaPortal *NominaPortalTransactorSession) SetXMsgMinGasLimit(gasLimit uint64) (*types.Transaction, error) {
	return _NominaPortal.Contract.SetXMsgMinGasLimit(&_NominaPortal.TransactOpts, gasLimit)
}

// SetXReceiptMaxErrorSize is a paid mutator transaction binding the contract method 0xbff0e84d.
//
// Solidity: function setXReceiptMaxErrorSize(uint16 numBytes) returns()
func (_NominaPortal *NominaPortalTransactor) SetXReceiptMaxErrorSize(opts *bind.TransactOpts, numBytes uint16) (*types.Transaction, error) {
	return _NominaPortal.contract.Transact(opts, "setXReceiptMaxErrorSize", numBytes)
}

// SetXReceiptMaxErrorSize is a paid mutator transaction binding the contract method 0xbff0e84d.
//
// Solidity: function setXReceiptMaxErrorSize(uint16 numBytes) returns()
func (_NominaPortal *NominaPortalSession) SetXReceiptMaxErrorSize(numBytes uint16) (*types.Transaction, error) {
	return _NominaPortal.Contract.SetXReceiptMaxErrorSize(&_NominaPortal.TransactOpts, numBytes)
}

// SetXReceiptMaxErrorSize is a paid mutator transaction binding the contract method 0xbff0e84d.
//
// Solidity: function setXReceiptMaxErrorSize(uint16 numBytes) returns()
func (_NominaPortal *NominaPortalTransactorSession) SetXReceiptMaxErrorSize(numBytes uint16) (*types.Transaction, error) {
	return _NominaPortal.Contract.SetXReceiptMaxErrorSize(&_NominaPortal.TransactOpts, numBytes)
}

// SetXSubValsetCutoff is a paid mutator transaction binding the contract method 0x103ba701.
//
// Solidity: function setXSubValsetCutoff(uint8 xsubValsetCutoff_) returns()
func (_NominaPortal *NominaPortalTransactor) SetXSubValsetCutoff(opts *bind.TransactOpts, xsubValsetCutoff_ uint8) (*types.Transaction, error) {
	return _NominaPortal.contract.Transact(opts, "setXSubValsetCutoff", xsubValsetCutoff_)
}

// SetXSubValsetCutoff is a paid mutator transaction binding the contract method 0x103ba701.
//
// Solidity: function setXSubValsetCutoff(uint8 xsubValsetCutoff_) returns()
func (_NominaPortal *NominaPortalSession) SetXSubValsetCutoff(xsubValsetCutoff_ uint8) (*types.Transaction, error) {
	return _NominaPortal.Contract.SetXSubValsetCutoff(&_NominaPortal.TransactOpts, xsubValsetCutoff_)
}

// SetXSubValsetCutoff is a paid mutator transaction binding the contract method 0x103ba701.
//
// Solidity: function setXSubValsetCutoff(uint8 xsubValsetCutoff_) returns()
func (_NominaPortal *NominaPortalTransactorSession) SetXSubValsetCutoff(xsubValsetCutoff_ uint8) (*types.Transaction, error) {
	return _NominaPortal.Contract.SetXSubValsetCutoff(&_NominaPortal.TransactOpts, xsubValsetCutoff_)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_NominaPortal *NominaPortalTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _NominaPortal.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_NominaPortal *NominaPortalSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _NominaPortal.Contract.TransferOwnership(&_NominaPortal.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_NominaPortal *NominaPortalTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _NominaPortal.Contract.TransferOwnership(&_NominaPortal.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_NominaPortal *NominaPortalTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NominaPortal.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_NominaPortal *NominaPortalSession) Unpause() (*types.Transaction, error) {
	return _NominaPortal.Contract.Unpause(&_NominaPortal.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_NominaPortal *NominaPortalTransactorSession) Unpause() (*types.Transaction, error) {
	return _NominaPortal.Contract.Unpause(&_NominaPortal.TransactOpts)
}

// UnpauseXCall is a paid mutator transaction binding the contract method 0x54d26bba.
//
// Solidity: function unpauseXCall() returns()
func (_NominaPortal *NominaPortalTransactor) UnpauseXCall(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NominaPortal.contract.Transact(opts, "unpauseXCall")
}

// UnpauseXCall is a paid mutator transaction binding the contract method 0x54d26bba.
//
// Solidity: function unpauseXCall() returns()
func (_NominaPortal *NominaPortalSession) UnpauseXCall() (*types.Transaction, error) {
	return _NominaPortal.Contract.UnpauseXCall(&_NominaPortal.TransactOpts)
}

// UnpauseXCall is a paid mutator transaction binding the contract method 0x54d26bba.
//
// Solidity: function unpauseXCall() returns()
func (_NominaPortal *NominaPortalTransactorSession) UnpauseXCall() (*types.Transaction, error) {
	return _NominaPortal.Contract.UnpauseXCall(&_NominaPortal.TransactOpts)
}

// UnpauseXCallTo is a paid mutator transaction binding the contract method 0xd533b445.
//
// Solidity: function unpauseXCallTo(uint64 chainId_) returns()
func (_NominaPortal *NominaPortalTransactor) UnpauseXCallTo(opts *bind.TransactOpts, chainId_ uint64) (*types.Transaction, error) {
	return _NominaPortal.contract.Transact(opts, "unpauseXCallTo", chainId_)
}

// UnpauseXCallTo is a paid mutator transaction binding the contract method 0xd533b445.
//
// Solidity: function unpauseXCallTo(uint64 chainId_) returns()
func (_NominaPortal *NominaPortalSession) UnpauseXCallTo(chainId_ uint64) (*types.Transaction, error) {
	return _NominaPortal.Contract.UnpauseXCallTo(&_NominaPortal.TransactOpts, chainId_)
}

// UnpauseXCallTo is a paid mutator transaction binding the contract method 0xd533b445.
//
// Solidity: function unpauseXCallTo(uint64 chainId_) returns()
func (_NominaPortal *NominaPortalTransactorSession) UnpauseXCallTo(chainId_ uint64) (*types.Transaction, error) {
	return _NominaPortal.Contract.UnpauseXCallTo(&_NominaPortal.TransactOpts, chainId_)
}

// UnpauseXSubmit is a paid mutator transaction binding the contract method 0xc3d8ad67.
//
// Solidity: function unpauseXSubmit() returns()
func (_NominaPortal *NominaPortalTransactor) UnpauseXSubmit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NominaPortal.contract.Transact(opts, "unpauseXSubmit")
}

// UnpauseXSubmit is a paid mutator transaction binding the contract method 0xc3d8ad67.
//
// Solidity: function unpauseXSubmit() returns()
func (_NominaPortal *NominaPortalSession) UnpauseXSubmit() (*types.Transaction, error) {
	return _NominaPortal.Contract.UnpauseXSubmit(&_NominaPortal.TransactOpts)
}

// UnpauseXSubmit is a paid mutator transaction binding the contract method 0xc3d8ad67.
//
// Solidity: function unpauseXSubmit() returns()
func (_NominaPortal *NominaPortalTransactorSession) UnpauseXSubmit() (*types.Transaction, error) {
	return _NominaPortal.Contract.UnpauseXSubmit(&_NominaPortal.TransactOpts)
}

// UnpauseXSubmitFrom is a paid mutator transaction binding the contract method 0xc2f9b968.
//
// Solidity: function unpauseXSubmitFrom(uint64 chainId_) returns()
func (_NominaPortal *NominaPortalTransactor) UnpauseXSubmitFrom(opts *bind.TransactOpts, chainId_ uint64) (*types.Transaction, error) {
	return _NominaPortal.contract.Transact(opts, "unpauseXSubmitFrom", chainId_)
}

// UnpauseXSubmitFrom is a paid mutator transaction binding the contract method 0xc2f9b968.
//
// Solidity: function unpauseXSubmitFrom(uint64 chainId_) returns()
func (_NominaPortal *NominaPortalSession) UnpauseXSubmitFrom(chainId_ uint64) (*types.Transaction, error) {
	return _NominaPortal.Contract.UnpauseXSubmitFrom(&_NominaPortal.TransactOpts, chainId_)
}

// UnpauseXSubmitFrom is a paid mutator transaction binding the contract method 0xc2f9b968.
//
// Solidity: function unpauseXSubmitFrom(uint64 chainId_) returns()
func (_NominaPortal *NominaPortalTransactorSession) UnpauseXSubmitFrom(chainId_ uint64) (*types.Transaction, error) {
	return _NominaPortal.Contract.UnpauseXSubmitFrom(&_NominaPortal.TransactOpts, chainId_)
}

// Xcall is a paid mutator transaction binding the contract method 0xc21dda4f.
//
// Solidity: function xcall(uint64 destChainId, uint8 conf, address to, bytes data, uint64 gasLimit) payable returns()
func (_NominaPortal *NominaPortalTransactor) Xcall(opts *bind.TransactOpts, destChainId uint64, conf uint8, to common.Address, data []byte, gasLimit uint64) (*types.Transaction, error) {
	return _NominaPortal.contract.Transact(opts, "xcall", destChainId, conf, to, data, gasLimit)
}

// Xcall is a paid mutator transaction binding the contract method 0xc21dda4f.
//
// Solidity: function xcall(uint64 destChainId, uint8 conf, address to, bytes data, uint64 gasLimit) payable returns()
func (_NominaPortal *NominaPortalSession) Xcall(destChainId uint64, conf uint8, to common.Address, data []byte, gasLimit uint64) (*types.Transaction, error) {
	return _NominaPortal.Contract.Xcall(&_NominaPortal.TransactOpts, destChainId, conf, to, data, gasLimit)
}

// Xcall is a paid mutator transaction binding the contract method 0xc21dda4f.
//
// Solidity: function xcall(uint64 destChainId, uint8 conf, address to, bytes data, uint64 gasLimit) payable returns()
func (_NominaPortal *NominaPortalTransactorSession) Xcall(destChainId uint64, conf uint8, to common.Address, data []byte, gasLimit uint64) (*types.Transaction, error) {
	return _NominaPortal.Contract.Xcall(&_NominaPortal.TransactOpts, destChainId, conf, to, data, gasLimit)
}

// Xsubmit is a paid mutator transaction binding the contract method 0x66a1eaf3.
//
// Solidity: function xsubmit((bytes32,uint64,(uint64,uint64,uint8,uint64,uint64,bytes32),(uint64,uint64,uint64,address,address,bytes,uint64)[],bytes32[],bool[],(address,bytes)[]) xsub) returns()
func (_NominaPortal *NominaPortalTransactor) Xsubmit(opts *bind.TransactOpts, xsub XTypesSubmission) (*types.Transaction, error) {
	return _NominaPortal.contract.Transact(opts, "xsubmit", xsub)
}

// Xsubmit is a paid mutator transaction binding the contract method 0x66a1eaf3.
//
// Solidity: function xsubmit((bytes32,uint64,(uint64,uint64,uint8,uint64,uint64,bytes32),(uint64,uint64,uint64,address,address,bytes,uint64)[],bytes32[],bool[],(address,bytes)[]) xsub) returns()
func (_NominaPortal *NominaPortalSession) Xsubmit(xsub XTypesSubmission) (*types.Transaction, error) {
	return _NominaPortal.Contract.Xsubmit(&_NominaPortal.TransactOpts, xsub)
}

// Xsubmit is a paid mutator transaction binding the contract method 0x66a1eaf3.
//
// Solidity: function xsubmit((bytes32,uint64,(uint64,uint64,uint8,uint64,uint64,bytes32),(uint64,uint64,uint64,address,address,bytes,uint64)[],bytes32[],bool[],(address,bytes)[]) xsub) returns()
func (_NominaPortal *NominaPortalTransactorSession) Xsubmit(xsub XTypesSubmission) (*types.Transaction, error) {
	return _NominaPortal.Contract.Xsubmit(&_NominaPortal.TransactOpts, xsub)
}

// NominaPortalFeeOracleSetIterator is returned from FilterFeeOracleSet and is used to iterate over the raw logs and unpacked data for FeeOracleSet events raised by the NominaPortal contract.
type NominaPortalFeeOracleSetIterator struct {
	Event *NominaPortalFeeOracleSet // Event containing the contract specifics and raw log

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
func (it *NominaPortalFeeOracleSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NominaPortalFeeOracleSet)
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
		it.Event = new(NominaPortalFeeOracleSet)
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
func (it *NominaPortalFeeOracleSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NominaPortalFeeOracleSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NominaPortalFeeOracleSet represents a FeeOracleSet event raised by the NominaPortal contract.
type NominaPortalFeeOracleSet struct {
	Oracle common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterFeeOracleSet is a free log retrieval operation binding the contract event 0xd97bdb0db82b52a85aa07f8da78033b1d6e159d94f1e3cbd4109d946c3bcfd32.
//
// Solidity: event FeeOracleSet(address oracle)
func (_NominaPortal *NominaPortalFilterer) FilterFeeOracleSet(opts *bind.FilterOpts) (*NominaPortalFeeOracleSetIterator, error) {

	logs, sub, err := _NominaPortal.contract.FilterLogs(opts, "FeeOracleSet")
	if err != nil {
		return nil, err
	}
	return &NominaPortalFeeOracleSetIterator{contract: _NominaPortal.contract, event: "FeeOracleSet", logs: logs, sub: sub}, nil
}

// WatchFeeOracleSet is a free log subscription operation binding the contract event 0xd97bdb0db82b52a85aa07f8da78033b1d6e159d94f1e3cbd4109d946c3bcfd32.
//
// Solidity: event FeeOracleSet(address oracle)
func (_NominaPortal *NominaPortalFilterer) WatchFeeOracleSet(opts *bind.WatchOpts, sink chan<- *NominaPortalFeeOracleSet) (event.Subscription, error) {

	logs, sub, err := _NominaPortal.contract.WatchLogs(opts, "FeeOracleSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NominaPortalFeeOracleSet)
				if err := _NominaPortal.contract.UnpackLog(event, "FeeOracleSet", log); err != nil {
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

// ParseFeeOracleSet is a log parse operation binding the contract event 0xd97bdb0db82b52a85aa07f8da78033b1d6e159d94f1e3cbd4109d946c3bcfd32.
//
// Solidity: event FeeOracleSet(address oracle)
func (_NominaPortal *NominaPortalFilterer) ParseFeeOracleSet(log types.Log) (*NominaPortalFeeOracleSet, error) {
	event := new(NominaPortalFeeOracleSet)
	if err := _NominaPortal.contract.UnpackLog(event, "FeeOracleSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NominaPortalFeesCollectedIterator is returned from FilterFeesCollected and is used to iterate over the raw logs and unpacked data for FeesCollected events raised by the NominaPortal contract.
type NominaPortalFeesCollectedIterator struct {
	Event *NominaPortalFeesCollected // Event containing the contract specifics and raw log

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
func (it *NominaPortalFeesCollectedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NominaPortalFeesCollected)
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
		it.Event = new(NominaPortalFeesCollected)
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
func (it *NominaPortalFeesCollectedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NominaPortalFeesCollectedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NominaPortalFeesCollected represents a FeesCollected event raised by the NominaPortal contract.
type NominaPortalFeesCollected struct {
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterFeesCollected is a free log retrieval operation binding the contract event 0x9dc46f23cfb5ddcad0ae7ea2be38d47fec07bb9382ec7e564efc69e036dd66ce.
//
// Solidity: event FeesCollected(address indexed to, uint256 amount)
func (_NominaPortal *NominaPortalFilterer) FilterFeesCollected(opts *bind.FilterOpts, to []common.Address) (*NominaPortalFeesCollectedIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _NominaPortal.contract.FilterLogs(opts, "FeesCollected", toRule)
	if err != nil {
		return nil, err
	}
	return &NominaPortalFeesCollectedIterator{contract: _NominaPortal.contract, event: "FeesCollected", logs: logs, sub: sub}, nil
}

// WatchFeesCollected is a free log subscription operation binding the contract event 0x9dc46f23cfb5ddcad0ae7ea2be38d47fec07bb9382ec7e564efc69e036dd66ce.
//
// Solidity: event FeesCollected(address indexed to, uint256 amount)
func (_NominaPortal *NominaPortalFilterer) WatchFeesCollected(opts *bind.WatchOpts, sink chan<- *NominaPortalFeesCollected, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _NominaPortal.contract.WatchLogs(opts, "FeesCollected", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NominaPortalFeesCollected)
				if err := _NominaPortal.contract.UnpackLog(event, "FeesCollected", log); err != nil {
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

// ParseFeesCollected is a log parse operation binding the contract event 0x9dc46f23cfb5ddcad0ae7ea2be38d47fec07bb9382ec7e564efc69e036dd66ce.
//
// Solidity: event FeesCollected(address indexed to, uint256 amount)
func (_NominaPortal *NominaPortalFilterer) ParseFeesCollected(log types.Log) (*NominaPortalFeesCollected, error) {
	event := new(NominaPortalFeesCollected)
	if err := _NominaPortal.contract.UnpackLog(event, "FeesCollected", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NominaPortalInXBlockOffsetSetIterator is returned from FilterInXBlockOffsetSet and is used to iterate over the raw logs and unpacked data for InXBlockOffsetSet events raised by the NominaPortal contract.
type NominaPortalInXBlockOffsetSetIterator struct {
	Event *NominaPortalInXBlockOffsetSet // Event containing the contract specifics and raw log

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
func (it *NominaPortalInXBlockOffsetSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NominaPortalInXBlockOffsetSet)
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
		it.Event = new(NominaPortalInXBlockOffsetSet)
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
func (it *NominaPortalInXBlockOffsetSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NominaPortalInXBlockOffsetSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NominaPortalInXBlockOffsetSet represents a InXBlockOffsetSet event raised by the NominaPortal contract.
type NominaPortalInXBlockOffsetSet struct {
	SrcChainId uint64
	ShardId    uint64
	Offset     uint64
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterInXBlockOffsetSet is a free log retrieval operation binding the contract event 0xe070f08cae8464c91238e8cbea64ccee5e7b48dd79a843f144e3721ee6bdd9b5.
//
// Solidity: event InXBlockOffsetSet(uint64 indexed srcChainId, uint64 indexed shardId, uint64 offset)
func (_NominaPortal *NominaPortalFilterer) FilterInXBlockOffsetSet(opts *bind.FilterOpts, srcChainId []uint64, shardId []uint64) (*NominaPortalInXBlockOffsetSetIterator, error) {

	var srcChainIdRule []interface{}
	for _, srcChainIdItem := range srcChainId {
		srcChainIdRule = append(srcChainIdRule, srcChainIdItem)
	}
	var shardIdRule []interface{}
	for _, shardIdItem := range shardId {
		shardIdRule = append(shardIdRule, shardIdItem)
	}

	logs, sub, err := _NominaPortal.contract.FilterLogs(opts, "InXBlockOffsetSet", srcChainIdRule, shardIdRule)
	if err != nil {
		return nil, err
	}
	return &NominaPortalInXBlockOffsetSetIterator{contract: _NominaPortal.contract, event: "InXBlockOffsetSet", logs: logs, sub: sub}, nil
}

// WatchInXBlockOffsetSet is a free log subscription operation binding the contract event 0xe070f08cae8464c91238e8cbea64ccee5e7b48dd79a843f144e3721ee6bdd9b5.
//
// Solidity: event InXBlockOffsetSet(uint64 indexed srcChainId, uint64 indexed shardId, uint64 offset)
func (_NominaPortal *NominaPortalFilterer) WatchInXBlockOffsetSet(opts *bind.WatchOpts, sink chan<- *NominaPortalInXBlockOffsetSet, srcChainId []uint64, shardId []uint64) (event.Subscription, error) {

	var srcChainIdRule []interface{}
	for _, srcChainIdItem := range srcChainId {
		srcChainIdRule = append(srcChainIdRule, srcChainIdItem)
	}
	var shardIdRule []interface{}
	for _, shardIdItem := range shardId {
		shardIdRule = append(shardIdRule, shardIdItem)
	}

	logs, sub, err := _NominaPortal.contract.WatchLogs(opts, "InXBlockOffsetSet", srcChainIdRule, shardIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NominaPortalInXBlockOffsetSet)
				if err := _NominaPortal.contract.UnpackLog(event, "InXBlockOffsetSet", log); err != nil {
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

// ParseInXBlockOffsetSet is a log parse operation binding the contract event 0xe070f08cae8464c91238e8cbea64ccee5e7b48dd79a843f144e3721ee6bdd9b5.
//
// Solidity: event InXBlockOffsetSet(uint64 indexed srcChainId, uint64 indexed shardId, uint64 offset)
func (_NominaPortal *NominaPortalFilterer) ParseInXBlockOffsetSet(log types.Log) (*NominaPortalInXBlockOffsetSet, error) {
	event := new(NominaPortalInXBlockOffsetSet)
	if err := _NominaPortal.contract.UnpackLog(event, "InXBlockOffsetSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NominaPortalInXMsgOffsetSetIterator is returned from FilterInXMsgOffsetSet and is used to iterate over the raw logs and unpacked data for InXMsgOffsetSet events raised by the NominaPortal contract.
type NominaPortalInXMsgOffsetSetIterator struct {
	Event *NominaPortalInXMsgOffsetSet // Event containing the contract specifics and raw log

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
func (it *NominaPortalInXMsgOffsetSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NominaPortalInXMsgOffsetSet)
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
		it.Event = new(NominaPortalInXMsgOffsetSet)
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
func (it *NominaPortalInXMsgOffsetSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NominaPortalInXMsgOffsetSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NominaPortalInXMsgOffsetSet represents a InXMsgOffsetSet event raised by the NominaPortal contract.
type NominaPortalInXMsgOffsetSet struct {
	SrcChainId uint64
	ShardId    uint64
	Offset     uint64
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterInXMsgOffsetSet is a free log retrieval operation binding the contract event 0x8647aae68c8456a1dcbfaf5eaadc94278ae423526d3f09c7b972bff7355d55c7.
//
// Solidity: event InXMsgOffsetSet(uint64 indexed srcChainId, uint64 indexed shardId, uint64 offset)
func (_NominaPortal *NominaPortalFilterer) FilterInXMsgOffsetSet(opts *bind.FilterOpts, srcChainId []uint64, shardId []uint64) (*NominaPortalInXMsgOffsetSetIterator, error) {

	var srcChainIdRule []interface{}
	for _, srcChainIdItem := range srcChainId {
		srcChainIdRule = append(srcChainIdRule, srcChainIdItem)
	}
	var shardIdRule []interface{}
	for _, shardIdItem := range shardId {
		shardIdRule = append(shardIdRule, shardIdItem)
	}

	logs, sub, err := _NominaPortal.contract.FilterLogs(opts, "InXMsgOffsetSet", srcChainIdRule, shardIdRule)
	if err != nil {
		return nil, err
	}
	return &NominaPortalInXMsgOffsetSetIterator{contract: _NominaPortal.contract, event: "InXMsgOffsetSet", logs: logs, sub: sub}, nil
}

// WatchInXMsgOffsetSet is a free log subscription operation binding the contract event 0x8647aae68c8456a1dcbfaf5eaadc94278ae423526d3f09c7b972bff7355d55c7.
//
// Solidity: event InXMsgOffsetSet(uint64 indexed srcChainId, uint64 indexed shardId, uint64 offset)
func (_NominaPortal *NominaPortalFilterer) WatchInXMsgOffsetSet(opts *bind.WatchOpts, sink chan<- *NominaPortalInXMsgOffsetSet, srcChainId []uint64, shardId []uint64) (event.Subscription, error) {

	var srcChainIdRule []interface{}
	for _, srcChainIdItem := range srcChainId {
		srcChainIdRule = append(srcChainIdRule, srcChainIdItem)
	}
	var shardIdRule []interface{}
	for _, shardIdItem := range shardId {
		shardIdRule = append(shardIdRule, shardIdItem)
	}

	logs, sub, err := _NominaPortal.contract.WatchLogs(opts, "InXMsgOffsetSet", srcChainIdRule, shardIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NominaPortalInXMsgOffsetSet)
				if err := _NominaPortal.contract.UnpackLog(event, "InXMsgOffsetSet", log); err != nil {
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

// ParseInXMsgOffsetSet is a log parse operation binding the contract event 0x8647aae68c8456a1dcbfaf5eaadc94278ae423526d3f09c7b972bff7355d55c7.
//
// Solidity: event InXMsgOffsetSet(uint64 indexed srcChainId, uint64 indexed shardId, uint64 offset)
func (_NominaPortal *NominaPortalFilterer) ParseInXMsgOffsetSet(log types.Log) (*NominaPortalInXMsgOffsetSet, error) {
	event := new(NominaPortalInXMsgOffsetSet)
	if err := _NominaPortal.contract.UnpackLog(event, "InXMsgOffsetSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NominaPortalInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the NominaPortal contract.
type NominaPortalInitializedIterator struct {
	Event *NominaPortalInitialized // Event containing the contract specifics and raw log

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
func (it *NominaPortalInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NominaPortalInitialized)
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
		it.Event = new(NominaPortalInitialized)
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
func (it *NominaPortalInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NominaPortalInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NominaPortalInitialized represents a Initialized event raised by the NominaPortal contract.
type NominaPortalInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_NominaPortal *NominaPortalFilterer) FilterInitialized(opts *bind.FilterOpts) (*NominaPortalInitializedIterator, error) {

	logs, sub, err := _NominaPortal.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &NominaPortalInitializedIterator{contract: _NominaPortal.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_NominaPortal *NominaPortalFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *NominaPortalInitialized) (event.Subscription, error) {

	logs, sub, err := _NominaPortal.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NominaPortalInitialized)
				if err := _NominaPortal.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_NominaPortal *NominaPortalFilterer) ParseInitialized(log types.Log) (*NominaPortalInitialized, error) {
	event := new(NominaPortalInitialized)
	if err := _NominaPortal.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NominaPortalOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the NominaPortal contract.
type NominaPortalOwnershipTransferredIterator struct {
	Event *NominaPortalOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *NominaPortalOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NominaPortalOwnershipTransferred)
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
		it.Event = new(NominaPortalOwnershipTransferred)
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
func (it *NominaPortalOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NominaPortalOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NominaPortalOwnershipTransferred represents a OwnershipTransferred event raised by the NominaPortal contract.
type NominaPortalOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_NominaPortal *NominaPortalFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*NominaPortalOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _NominaPortal.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &NominaPortalOwnershipTransferredIterator{contract: _NominaPortal.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_NominaPortal *NominaPortalFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *NominaPortalOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _NominaPortal.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NominaPortalOwnershipTransferred)
				if err := _NominaPortal.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_NominaPortal *NominaPortalFilterer) ParseOwnershipTransferred(log types.Log) (*NominaPortalOwnershipTransferred, error) {
	event := new(NominaPortalOwnershipTransferred)
	if err := _NominaPortal.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NominaPortalPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the NominaPortal contract.
type NominaPortalPausedIterator struct {
	Event *NominaPortalPaused // Event containing the contract specifics and raw log

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
func (it *NominaPortalPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NominaPortalPaused)
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
		it.Event = new(NominaPortalPaused)
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
func (it *NominaPortalPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NominaPortalPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NominaPortalPaused represents a Paused event raised by the NominaPortal contract.
type NominaPortalPaused struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x9e87fac88ff661f02d44f95383c817fece4bce600a3dab7a54406878b965e752.
//
// Solidity: event Paused()
func (_NominaPortal *NominaPortalFilterer) FilterPaused(opts *bind.FilterOpts) (*NominaPortalPausedIterator, error) {

	logs, sub, err := _NominaPortal.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &NominaPortalPausedIterator{contract: _NominaPortal.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x9e87fac88ff661f02d44f95383c817fece4bce600a3dab7a54406878b965e752.
//
// Solidity: event Paused()
func (_NominaPortal *NominaPortalFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *NominaPortalPaused) (event.Subscription, error) {

	logs, sub, err := _NominaPortal.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NominaPortalPaused)
				if err := _NominaPortal.contract.UnpackLog(event, "Paused", log); err != nil {
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

// ParsePaused is a log parse operation binding the contract event 0x9e87fac88ff661f02d44f95383c817fece4bce600a3dab7a54406878b965e752.
//
// Solidity: event Paused()
func (_NominaPortal *NominaPortalFilterer) ParsePaused(log types.Log) (*NominaPortalPaused, error) {
	event := new(NominaPortalPaused)
	if err := _NominaPortal.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NominaPortalPaused0Iterator is returned from FilterPaused0 and is used to iterate over the raw logs and unpacked data for Paused0 events raised by the NominaPortal contract.
type NominaPortalPaused0Iterator struct {
	Event *NominaPortalPaused0 // Event containing the contract specifics and raw log

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
func (it *NominaPortalPaused0Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NominaPortalPaused0)
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
		it.Event = new(NominaPortalPaused0)
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
func (it *NominaPortalPaused0Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NominaPortalPaused0Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NominaPortalPaused0 represents a Paused0 event raised by the NominaPortal contract.
type NominaPortalPaused0 struct {
	Key [32]byte
	Raw types.Log // Blockchain specific contextual infos
}

// FilterPaused0 is a free log retrieval operation binding the contract event 0x0cb09dc71d57eeec2046f6854976717e4874a3cf2d6ddeddde337e5b6de6ba31.
//
// Solidity: event Paused(bytes32 indexed key)
func (_NominaPortal *NominaPortalFilterer) FilterPaused0(opts *bind.FilterOpts, key [][32]byte) (*NominaPortalPaused0Iterator, error) {

	var keyRule []interface{}
	for _, keyItem := range key {
		keyRule = append(keyRule, keyItem)
	}

	logs, sub, err := _NominaPortal.contract.FilterLogs(opts, "Paused0", keyRule)
	if err != nil {
		return nil, err
	}
	return &NominaPortalPaused0Iterator{contract: _NominaPortal.contract, event: "Paused0", logs: logs, sub: sub}, nil
}

// WatchPaused0 is a free log subscription operation binding the contract event 0x0cb09dc71d57eeec2046f6854976717e4874a3cf2d6ddeddde337e5b6de6ba31.
//
// Solidity: event Paused(bytes32 indexed key)
func (_NominaPortal *NominaPortalFilterer) WatchPaused0(opts *bind.WatchOpts, sink chan<- *NominaPortalPaused0, key [][32]byte) (event.Subscription, error) {

	var keyRule []interface{}
	for _, keyItem := range key {
		keyRule = append(keyRule, keyItem)
	}

	logs, sub, err := _NominaPortal.contract.WatchLogs(opts, "Paused0", keyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NominaPortalPaused0)
				if err := _NominaPortal.contract.UnpackLog(event, "Paused0", log); err != nil {
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

// ParsePaused0 is a log parse operation binding the contract event 0x0cb09dc71d57eeec2046f6854976717e4874a3cf2d6ddeddde337e5b6de6ba31.
//
// Solidity: event Paused(bytes32 indexed key)
func (_NominaPortal *NominaPortalFilterer) ParsePaused0(log types.Log) (*NominaPortalPaused0, error) {
	event := new(NominaPortalPaused0)
	if err := _NominaPortal.contract.UnpackLog(event, "Paused0", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NominaPortalUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the NominaPortal contract.
type NominaPortalUnpausedIterator struct {
	Event *NominaPortalUnpaused // Event containing the contract specifics and raw log

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
func (it *NominaPortalUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NominaPortalUnpaused)
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
		it.Event = new(NominaPortalUnpaused)
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
func (it *NominaPortalUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NominaPortalUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NominaPortalUnpaused represents a Unpaused event raised by the NominaPortal contract.
type NominaPortalUnpaused struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0xa45f47fdea8a1efdd9029a5691c7f759c32b7c698632b563573e155625d16933.
//
// Solidity: event Unpaused()
func (_NominaPortal *NominaPortalFilterer) FilterUnpaused(opts *bind.FilterOpts) (*NominaPortalUnpausedIterator, error) {

	logs, sub, err := _NominaPortal.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &NominaPortalUnpausedIterator{contract: _NominaPortal.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0xa45f47fdea8a1efdd9029a5691c7f759c32b7c698632b563573e155625d16933.
//
// Solidity: event Unpaused()
func (_NominaPortal *NominaPortalFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *NominaPortalUnpaused) (event.Subscription, error) {

	logs, sub, err := _NominaPortal.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NominaPortalUnpaused)
				if err := _NominaPortal.contract.UnpackLog(event, "Unpaused", log); err != nil {
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

// ParseUnpaused is a log parse operation binding the contract event 0xa45f47fdea8a1efdd9029a5691c7f759c32b7c698632b563573e155625d16933.
//
// Solidity: event Unpaused()
func (_NominaPortal *NominaPortalFilterer) ParseUnpaused(log types.Log) (*NominaPortalUnpaused, error) {
	event := new(NominaPortalUnpaused)
	if err := _NominaPortal.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NominaPortalUnpaused0Iterator is returned from FilterUnpaused0 and is used to iterate over the raw logs and unpacked data for Unpaused0 events raised by the NominaPortal contract.
type NominaPortalUnpaused0Iterator struct {
	Event *NominaPortalUnpaused0 // Event containing the contract specifics and raw log

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
func (it *NominaPortalUnpaused0Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NominaPortalUnpaused0)
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
		it.Event = new(NominaPortalUnpaused0)
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
func (it *NominaPortalUnpaused0Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NominaPortalUnpaused0Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NominaPortalUnpaused0 represents a Unpaused0 event raised by the NominaPortal contract.
type NominaPortalUnpaused0 struct {
	Key [32]byte
	Raw types.Log // Blockchain specific contextual infos
}

// FilterUnpaused0 is a free log retrieval operation binding the contract event 0xd05bfc2250abb0f8fd265a54c53a24359c5484af63cad2e4ce87c78ab751395a.
//
// Solidity: event Unpaused(bytes32 indexed key)
func (_NominaPortal *NominaPortalFilterer) FilterUnpaused0(opts *bind.FilterOpts, key [][32]byte) (*NominaPortalUnpaused0Iterator, error) {

	var keyRule []interface{}
	for _, keyItem := range key {
		keyRule = append(keyRule, keyItem)
	}

	logs, sub, err := _NominaPortal.contract.FilterLogs(opts, "Unpaused0", keyRule)
	if err != nil {
		return nil, err
	}
	return &NominaPortalUnpaused0Iterator{contract: _NominaPortal.contract, event: "Unpaused0", logs: logs, sub: sub}, nil
}

// WatchUnpaused0 is a free log subscription operation binding the contract event 0xd05bfc2250abb0f8fd265a54c53a24359c5484af63cad2e4ce87c78ab751395a.
//
// Solidity: event Unpaused(bytes32 indexed key)
func (_NominaPortal *NominaPortalFilterer) WatchUnpaused0(opts *bind.WatchOpts, sink chan<- *NominaPortalUnpaused0, key [][32]byte) (event.Subscription, error) {

	var keyRule []interface{}
	for _, keyItem := range key {
		keyRule = append(keyRule, keyItem)
	}

	logs, sub, err := _NominaPortal.contract.WatchLogs(opts, "Unpaused0", keyRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NominaPortalUnpaused0)
				if err := _NominaPortal.contract.UnpackLog(event, "Unpaused0", log); err != nil {
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

// ParseUnpaused0 is a log parse operation binding the contract event 0xd05bfc2250abb0f8fd265a54c53a24359c5484af63cad2e4ce87c78ab751395a.
//
// Solidity: event Unpaused(bytes32 indexed key)
func (_NominaPortal *NominaPortalFilterer) ParseUnpaused0(log types.Log) (*NominaPortalUnpaused0, error) {
	event := new(NominaPortalUnpaused0)
	if err := _NominaPortal.contract.UnpackLog(event, "Unpaused0", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NominaPortalValidatorSetAddedIterator is returned from FilterValidatorSetAdded and is used to iterate over the raw logs and unpacked data for ValidatorSetAdded events raised by the NominaPortal contract.
type NominaPortalValidatorSetAddedIterator struct {
	Event *NominaPortalValidatorSetAdded // Event containing the contract specifics and raw log

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
func (it *NominaPortalValidatorSetAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NominaPortalValidatorSetAdded)
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
		it.Event = new(NominaPortalValidatorSetAdded)
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
func (it *NominaPortalValidatorSetAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NominaPortalValidatorSetAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NominaPortalValidatorSetAdded represents a ValidatorSetAdded event raised by the NominaPortal contract.
type NominaPortalValidatorSetAdded struct {
	SetId uint64
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterValidatorSetAdded is a free log retrieval operation binding the contract event 0x3a7c2f997a87ba92aedaecd1127f4129cae1283e2809ebf5304d321b943fd107.
//
// Solidity: event ValidatorSetAdded(uint64 indexed setId)
func (_NominaPortal *NominaPortalFilterer) FilterValidatorSetAdded(opts *bind.FilterOpts, setId []uint64) (*NominaPortalValidatorSetAddedIterator, error) {

	var setIdRule []interface{}
	for _, setIdItem := range setId {
		setIdRule = append(setIdRule, setIdItem)
	}

	logs, sub, err := _NominaPortal.contract.FilterLogs(opts, "ValidatorSetAdded", setIdRule)
	if err != nil {
		return nil, err
	}
	return &NominaPortalValidatorSetAddedIterator{contract: _NominaPortal.contract, event: "ValidatorSetAdded", logs: logs, sub: sub}, nil
}

// WatchValidatorSetAdded is a free log subscription operation binding the contract event 0x3a7c2f997a87ba92aedaecd1127f4129cae1283e2809ebf5304d321b943fd107.
//
// Solidity: event ValidatorSetAdded(uint64 indexed setId)
func (_NominaPortal *NominaPortalFilterer) WatchValidatorSetAdded(opts *bind.WatchOpts, sink chan<- *NominaPortalValidatorSetAdded, setId []uint64) (event.Subscription, error) {

	var setIdRule []interface{}
	for _, setIdItem := range setId {
		setIdRule = append(setIdRule, setIdItem)
	}

	logs, sub, err := _NominaPortal.contract.WatchLogs(opts, "ValidatorSetAdded", setIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NominaPortalValidatorSetAdded)
				if err := _NominaPortal.contract.UnpackLog(event, "ValidatorSetAdded", log); err != nil {
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

// ParseValidatorSetAdded is a log parse operation binding the contract event 0x3a7c2f997a87ba92aedaecd1127f4129cae1283e2809ebf5304d321b943fd107.
//
// Solidity: event ValidatorSetAdded(uint64 indexed setId)
func (_NominaPortal *NominaPortalFilterer) ParseValidatorSetAdded(log types.Log) (*NominaPortalValidatorSetAdded, error) {
	event := new(NominaPortalValidatorSetAdded)
	if err := _NominaPortal.contract.UnpackLog(event, "ValidatorSetAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NominaPortalXCallPausedIterator is returned from FilterXCallPaused and is used to iterate over the raw logs and unpacked data for XCallPaused events raised by the NominaPortal contract.
type NominaPortalXCallPausedIterator struct {
	Event *NominaPortalXCallPaused // Event containing the contract specifics and raw log

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
func (it *NominaPortalXCallPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NominaPortalXCallPaused)
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
		it.Event = new(NominaPortalXCallPaused)
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
func (it *NominaPortalXCallPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NominaPortalXCallPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NominaPortalXCallPaused represents a XCallPaused event raised by the NominaPortal contract.
type NominaPortalXCallPaused struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterXCallPaused is a free log retrieval operation binding the contract event 0x5f335a4032d4cfb6aca7835b0c2225f36d4d9eaa4ed43ee59ed537e02dff6b39.
//
// Solidity: event XCallPaused()
func (_NominaPortal *NominaPortalFilterer) FilterXCallPaused(opts *bind.FilterOpts) (*NominaPortalXCallPausedIterator, error) {

	logs, sub, err := _NominaPortal.contract.FilterLogs(opts, "XCallPaused")
	if err != nil {
		return nil, err
	}
	return &NominaPortalXCallPausedIterator{contract: _NominaPortal.contract, event: "XCallPaused", logs: logs, sub: sub}, nil
}

// WatchXCallPaused is a free log subscription operation binding the contract event 0x5f335a4032d4cfb6aca7835b0c2225f36d4d9eaa4ed43ee59ed537e02dff6b39.
//
// Solidity: event XCallPaused()
func (_NominaPortal *NominaPortalFilterer) WatchXCallPaused(opts *bind.WatchOpts, sink chan<- *NominaPortalXCallPaused) (event.Subscription, error) {

	logs, sub, err := _NominaPortal.contract.WatchLogs(opts, "XCallPaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NominaPortalXCallPaused)
				if err := _NominaPortal.contract.UnpackLog(event, "XCallPaused", log); err != nil {
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

// ParseXCallPaused is a log parse operation binding the contract event 0x5f335a4032d4cfb6aca7835b0c2225f36d4d9eaa4ed43ee59ed537e02dff6b39.
//
// Solidity: event XCallPaused()
func (_NominaPortal *NominaPortalFilterer) ParseXCallPaused(log types.Log) (*NominaPortalXCallPaused, error) {
	event := new(NominaPortalXCallPaused)
	if err := _NominaPortal.contract.UnpackLog(event, "XCallPaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NominaPortalXCallToPausedIterator is returned from FilterXCallToPaused and is used to iterate over the raw logs and unpacked data for XCallToPaused events raised by the NominaPortal contract.
type NominaPortalXCallToPausedIterator struct {
	Event *NominaPortalXCallToPaused // Event containing the contract specifics and raw log

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
func (it *NominaPortalXCallToPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NominaPortalXCallToPaused)
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
		it.Event = new(NominaPortalXCallToPaused)
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
func (it *NominaPortalXCallToPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NominaPortalXCallToPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NominaPortalXCallToPaused represents a XCallToPaused event raised by the NominaPortal contract.
type NominaPortalXCallToPaused struct {
	ChainId uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterXCallToPaused is a free log retrieval operation binding the contract event 0xcd7910e1c5569d8433ce4ef8e5d51c1bdc03168f614b576da47dc3d2b51d033a.
//
// Solidity: event XCallToPaused(uint64 indexed chainId)
func (_NominaPortal *NominaPortalFilterer) FilterXCallToPaused(opts *bind.FilterOpts, chainId []uint64) (*NominaPortalXCallToPausedIterator, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}

	logs, sub, err := _NominaPortal.contract.FilterLogs(opts, "XCallToPaused", chainIdRule)
	if err != nil {
		return nil, err
	}
	return &NominaPortalXCallToPausedIterator{contract: _NominaPortal.contract, event: "XCallToPaused", logs: logs, sub: sub}, nil
}

// WatchXCallToPaused is a free log subscription operation binding the contract event 0xcd7910e1c5569d8433ce4ef8e5d51c1bdc03168f614b576da47dc3d2b51d033a.
//
// Solidity: event XCallToPaused(uint64 indexed chainId)
func (_NominaPortal *NominaPortalFilterer) WatchXCallToPaused(opts *bind.WatchOpts, sink chan<- *NominaPortalXCallToPaused, chainId []uint64) (event.Subscription, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}

	logs, sub, err := _NominaPortal.contract.WatchLogs(opts, "XCallToPaused", chainIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NominaPortalXCallToPaused)
				if err := _NominaPortal.contract.UnpackLog(event, "XCallToPaused", log); err != nil {
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

// ParseXCallToPaused is a log parse operation binding the contract event 0xcd7910e1c5569d8433ce4ef8e5d51c1bdc03168f614b576da47dc3d2b51d033a.
//
// Solidity: event XCallToPaused(uint64 indexed chainId)
func (_NominaPortal *NominaPortalFilterer) ParseXCallToPaused(log types.Log) (*NominaPortalXCallToPaused, error) {
	event := new(NominaPortalXCallToPaused)
	if err := _NominaPortal.contract.UnpackLog(event, "XCallToPaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NominaPortalXCallToUnpausedIterator is returned from FilterXCallToUnpaused and is used to iterate over the raw logs and unpacked data for XCallToUnpaused events raised by the NominaPortal contract.
type NominaPortalXCallToUnpausedIterator struct {
	Event *NominaPortalXCallToUnpaused // Event containing the contract specifics and raw log

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
func (it *NominaPortalXCallToUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NominaPortalXCallToUnpaused)
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
		it.Event = new(NominaPortalXCallToUnpaused)
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
func (it *NominaPortalXCallToUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NominaPortalXCallToUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NominaPortalXCallToUnpaused represents a XCallToUnpaused event raised by the NominaPortal contract.
type NominaPortalXCallToUnpaused struct {
	ChainId uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterXCallToUnpaused is a free log retrieval operation binding the contract event 0x1ed9223556fb0971076c30172f1f00630efd313b6a05290a562aef95928e7125.
//
// Solidity: event XCallToUnpaused(uint64 indexed chainId)
func (_NominaPortal *NominaPortalFilterer) FilterXCallToUnpaused(opts *bind.FilterOpts, chainId []uint64) (*NominaPortalXCallToUnpausedIterator, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}

	logs, sub, err := _NominaPortal.contract.FilterLogs(opts, "XCallToUnpaused", chainIdRule)
	if err != nil {
		return nil, err
	}
	return &NominaPortalXCallToUnpausedIterator{contract: _NominaPortal.contract, event: "XCallToUnpaused", logs: logs, sub: sub}, nil
}

// WatchXCallToUnpaused is a free log subscription operation binding the contract event 0x1ed9223556fb0971076c30172f1f00630efd313b6a05290a562aef95928e7125.
//
// Solidity: event XCallToUnpaused(uint64 indexed chainId)
func (_NominaPortal *NominaPortalFilterer) WatchXCallToUnpaused(opts *bind.WatchOpts, sink chan<- *NominaPortalXCallToUnpaused, chainId []uint64) (event.Subscription, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}

	logs, sub, err := _NominaPortal.contract.WatchLogs(opts, "XCallToUnpaused", chainIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NominaPortalXCallToUnpaused)
				if err := _NominaPortal.contract.UnpackLog(event, "XCallToUnpaused", log); err != nil {
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

// ParseXCallToUnpaused is a log parse operation binding the contract event 0x1ed9223556fb0971076c30172f1f00630efd313b6a05290a562aef95928e7125.
//
// Solidity: event XCallToUnpaused(uint64 indexed chainId)
func (_NominaPortal *NominaPortalFilterer) ParseXCallToUnpaused(log types.Log) (*NominaPortalXCallToUnpaused, error) {
	event := new(NominaPortalXCallToUnpaused)
	if err := _NominaPortal.contract.UnpackLog(event, "XCallToUnpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NominaPortalXCallUnpausedIterator is returned from FilterXCallUnpaused and is used to iterate over the raw logs and unpacked data for XCallUnpaused events raised by the NominaPortal contract.
type NominaPortalXCallUnpausedIterator struct {
	Event *NominaPortalXCallUnpaused // Event containing the contract specifics and raw log

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
func (it *NominaPortalXCallUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NominaPortalXCallUnpaused)
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
		it.Event = new(NominaPortalXCallUnpaused)
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
func (it *NominaPortalXCallUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NominaPortalXCallUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NominaPortalXCallUnpaused represents a XCallUnpaused event raised by the NominaPortal contract.
type NominaPortalXCallUnpaused struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterXCallUnpaused is a free log retrieval operation binding the contract event 0x4c48c7b71557216a3192842746bdfc381f98d7536d9eb1c6764f3b45e6794827.
//
// Solidity: event XCallUnpaused()
func (_NominaPortal *NominaPortalFilterer) FilterXCallUnpaused(opts *bind.FilterOpts) (*NominaPortalXCallUnpausedIterator, error) {

	logs, sub, err := _NominaPortal.contract.FilterLogs(opts, "XCallUnpaused")
	if err != nil {
		return nil, err
	}
	return &NominaPortalXCallUnpausedIterator{contract: _NominaPortal.contract, event: "XCallUnpaused", logs: logs, sub: sub}, nil
}

// WatchXCallUnpaused is a free log subscription operation binding the contract event 0x4c48c7b71557216a3192842746bdfc381f98d7536d9eb1c6764f3b45e6794827.
//
// Solidity: event XCallUnpaused()
func (_NominaPortal *NominaPortalFilterer) WatchXCallUnpaused(opts *bind.WatchOpts, sink chan<- *NominaPortalXCallUnpaused) (event.Subscription, error) {

	logs, sub, err := _NominaPortal.contract.WatchLogs(opts, "XCallUnpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NominaPortalXCallUnpaused)
				if err := _NominaPortal.contract.UnpackLog(event, "XCallUnpaused", log); err != nil {
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

// ParseXCallUnpaused is a log parse operation binding the contract event 0x4c48c7b71557216a3192842746bdfc381f98d7536d9eb1c6764f3b45e6794827.
//
// Solidity: event XCallUnpaused()
func (_NominaPortal *NominaPortalFilterer) ParseXCallUnpaused(log types.Log) (*NominaPortalXCallUnpaused, error) {
	event := new(NominaPortalXCallUnpaused)
	if err := _NominaPortal.contract.UnpackLog(event, "XCallUnpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NominaPortalXMsgIterator is returned from FilterXMsg and is used to iterate over the raw logs and unpacked data for XMsg events raised by the NominaPortal contract.
type NominaPortalXMsgIterator struct {
	Event *NominaPortalXMsg // Event containing the contract specifics and raw log

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
func (it *NominaPortalXMsgIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NominaPortalXMsg)
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
		it.Event = new(NominaPortalXMsg)
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
func (it *NominaPortalXMsgIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NominaPortalXMsgIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NominaPortalXMsg represents a XMsg event raised by the NominaPortal contract.
type NominaPortalXMsg struct {
	DestChainId uint64
	ShardId     uint64
	Offset      uint64
	Sender      common.Address
	To          common.Address
	Data        []byte
	GasLimit    uint64
	Fees        *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterXMsg is a free log retrieval operation binding the contract event 0xb7c8eb9d7a7fbcdab809ab7b8a7c41701eb3115e3fe99d30ff490d8552f72bfa.
//
// Solidity: event XMsg(uint64 indexed destChainId, uint64 indexed shardId, uint64 indexed offset, address sender, address to, bytes data, uint64 gasLimit, uint256 fees)
func (_NominaPortal *NominaPortalFilterer) FilterXMsg(opts *bind.FilterOpts, destChainId []uint64, shardId []uint64, offset []uint64) (*NominaPortalXMsgIterator, error) {

	var destChainIdRule []interface{}
	for _, destChainIdItem := range destChainId {
		destChainIdRule = append(destChainIdRule, destChainIdItem)
	}
	var shardIdRule []interface{}
	for _, shardIdItem := range shardId {
		shardIdRule = append(shardIdRule, shardIdItem)
	}
	var offsetRule []interface{}
	for _, offsetItem := range offset {
		offsetRule = append(offsetRule, offsetItem)
	}

	logs, sub, err := _NominaPortal.contract.FilterLogs(opts, "XMsg", destChainIdRule, shardIdRule, offsetRule)
	if err != nil {
		return nil, err
	}
	return &NominaPortalXMsgIterator{contract: _NominaPortal.contract, event: "XMsg", logs: logs, sub: sub}, nil
}

// WatchXMsg is a free log subscription operation binding the contract event 0xb7c8eb9d7a7fbcdab809ab7b8a7c41701eb3115e3fe99d30ff490d8552f72bfa.
//
// Solidity: event XMsg(uint64 indexed destChainId, uint64 indexed shardId, uint64 indexed offset, address sender, address to, bytes data, uint64 gasLimit, uint256 fees)
func (_NominaPortal *NominaPortalFilterer) WatchXMsg(opts *bind.WatchOpts, sink chan<- *NominaPortalXMsg, destChainId []uint64, shardId []uint64, offset []uint64) (event.Subscription, error) {

	var destChainIdRule []interface{}
	for _, destChainIdItem := range destChainId {
		destChainIdRule = append(destChainIdRule, destChainIdItem)
	}
	var shardIdRule []interface{}
	for _, shardIdItem := range shardId {
		shardIdRule = append(shardIdRule, shardIdItem)
	}
	var offsetRule []interface{}
	for _, offsetItem := range offset {
		offsetRule = append(offsetRule, offsetItem)
	}

	logs, sub, err := _NominaPortal.contract.WatchLogs(opts, "XMsg", destChainIdRule, shardIdRule, offsetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NominaPortalXMsg)
				if err := _NominaPortal.contract.UnpackLog(event, "XMsg", log); err != nil {
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

// ParseXMsg is a log parse operation binding the contract event 0xb7c8eb9d7a7fbcdab809ab7b8a7c41701eb3115e3fe99d30ff490d8552f72bfa.
//
// Solidity: event XMsg(uint64 indexed destChainId, uint64 indexed shardId, uint64 indexed offset, address sender, address to, bytes data, uint64 gasLimit, uint256 fees)
func (_NominaPortal *NominaPortalFilterer) ParseXMsg(log types.Log) (*NominaPortalXMsg, error) {
	event := new(NominaPortalXMsg)
	if err := _NominaPortal.contract.UnpackLog(event, "XMsg", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NominaPortalXMsgMaxDataSizeSetIterator is returned from FilterXMsgMaxDataSizeSet and is used to iterate over the raw logs and unpacked data for XMsgMaxDataSizeSet events raised by the NominaPortal contract.
type NominaPortalXMsgMaxDataSizeSetIterator struct {
	Event *NominaPortalXMsgMaxDataSizeSet // Event containing the contract specifics and raw log

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
func (it *NominaPortalXMsgMaxDataSizeSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NominaPortalXMsgMaxDataSizeSet)
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
		it.Event = new(NominaPortalXMsgMaxDataSizeSet)
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
func (it *NominaPortalXMsgMaxDataSizeSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NominaPortalXMsgMaxDataSizeSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NominaPortalXMsgMaxDataSizeSet represents a XMsgMaxDataSizeSet event raised by the NominaPortal contract.
type NominaPortalXMsgMaxDataSizeSet struct {
	Size uint16
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterXMsgMaxDataSizeSet is a free log retrieval operation binding the contract event 0x65923e04419dc810d0ea08a94a7f608d4c4d949818d95c3788f895e575dd2064.
//
// Solidity: event XMsgMaxDataSizeSet(uint16 size)
func (_NominaPortal *NominaPortalFilterer) FilterXMsgMaxDataSizeSet(opts *bind.FilterOpts) (*NominaPortalXMsgMaxDataSizeSetIterator, error) {

	logs, sub, err := _NominaPortal.contract.FilterLogs(opts, "XMsgMaxDataSizeSet")
	if err != nil {
		return nil, err
	}
	return &NominaPortalXMsgMaxDataSizeSetIterator{contract: _NominaPortal.contract, event: "XMsgMaxDataSizeSet", logs: logs, sub: sub}, nil
}

// WatchXMsgMaxDataSizeSet is a free log subscription operation binding the contract event 0x65923e04419dc810d0ea08a94a7f608d4c4d949818d95c3788f895e575dd2064.
//
// Solidity: event XMsgMaxDataSizeSet(uint16 size)
func (_NominaPortal *NominaPortalFilterer) WatchXMsgMaxDataSizeSet(opts *bind.WatchOpts, sink chan<- *NominaPortalXMsgMaxDataSizeSet) (event.Subscription, error) {

	logs, sub, err := _NominaPortal.contract.WatchLogs(opts, "XMsgMaxDataSizeSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NominaPortalXMsgMaxDataSizeSet)
				if err := _NominaPortal.contract.UnpackLog(event, "XMsgMaxDataSizeSet", log); err != nil {
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

// ParseXMsgMaxDataSizeSet is a log parse operation binding the contract event 0x65923e04419dc810d0ea08a94a7f608d4c4d949818d95c3788f895e575dd2064.
//
// Solidity: event XMsgMaxDataSizeSet(uint16 size)
func (_NominaPortal *NominaPortalFilterer) ParseXMsgMaxDataSizeSet(log types.Log) (*NominaPortalXMsgMaxDataSizeSet, error) {
	event := new(NominaPortalXMsgMaxDataSizeSet)
	if err := _NominaPortal.contract.UnpackLog(event, "XMsgMaxDataSizeSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NominaPortalXMsgMaxGasLimitSetIterator is returned from FilterXMsgMaxGasLimitSet and is used to iterate over the raw logs and unpacked data for XMsgMaxGasLimitSet events raised by the NominaPortal contract.
type NominaPortalXMsgMaxGasLimitSetIterator struct {
	Event *NominaPortalXMsgMaxGasLimitSet // Event containing the contract specifics and raw log

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
func (it *NominaPortalXMsgMaxGasLimitSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NominaPortalXMsgMaxGasLimitSet)
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
		it.Event = new(NominaPortalXMsgMaxGasLimitSet)
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
func (it *NominaPortalXMsgMaxGasLimitSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NominaPortalXMsgMaxGasLimitSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NominaPortalXMsgMaxGasLimitSet represents a XMsgMaxGasLimitSet event raised by the NominaPortal contract.
type NominaPortalXMsgMaxGasLimitSet struct {
	GasLimit uint64
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterXMsgMaxGasLimitSet is a free log retrieval operation binding the contract event 0x1153561ac5effc2926ba6c612f86a397c997bc43dfbfc718da08065be0c5fe4d.
//
// Solidity: event XMsgMaxGasLimitSet(uint64 gasLimit)
func (_NominaPortal *NominaPortalFilterer) FilterXMsgMaxGasLimitSet(opts *bind.FilterOpts) (*NominaPortalXMsgMaxGasLimitSetIterator, error) {

	logs, sub, err := _NominaPortal.contract.FilterLogs(opts, "XMsgMaxGasLimitSet")
	if err != nil {
		return nil, err
	}
	return &NominaPortalXMsgMaxGasLimitSetIterator{contract: _NominaPortal.contract, event: "XMsgMaxGasLimitSet", logs: logs, sub: sub}, nil
}

// WatchXMsgMaxGasLimitSet is a free log subscription operation binding the contract event 0x1153561ac5effc2926ba6c612f86a397c997bc43dfbfc718da08065be0c5fe4d.
//
// Solidity: event XMsgMaxGasLimitSet(uint64 gasLimit)
func (_NominaPortal *NominaPortalFilterer) WatchXMsgMaxGasLimitSet(opts *bind.WatchOpts, sink chan<- *NominaPortalXMsgMaxGasLimitSet) (event.Subscription, error) {

	logs, sub, err := _NominaPortal.contract.WatchLogs(opts, "XMsgMaxGasLimitSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NominaPortalXMsgMaxGasLimitSet)
				if err := _NominaPortal.contract.UnpackLog(event, "XMsgMaxGasLimitSet", log); err != nil {
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

// ParseXMsgMaxGasLimitSet is a log parse operation binding the contract event 0x1153561ac5effc2926ba6c612f86a397c997bc43dfbfc718da08065be0c5fe4d.
//
// Solidity: event XMsgMaxGasLimitSet(uint64 gasLimit)
func (_NominaPortal *NominaPortalFilterer) ParseXMsgMaxGasLimitSet(log types.Log) (*NominaPortalXMsgMaxGasLimitSet, error) {
	event := new(NominaPortalXMsgMaxGasLimitSet)
	if err := _NominaPortal.contract.UnpackLog(event, "XMsgMaxGasLimitSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NominaPortalXMsgMinGasLimitSetIterator is returned from FilterXMsgMinGasLimitSet and is used to iterate over the raw logs and unpacked data for XMsgMinGasLimitSet events raised by the NominaPortal contract.
type NominaPortalXMsgMinGasLimitSetIterator struct {
	Event *NominaPortalXMsgMinGasLimitSet // Event containing the contract specifics and raw log

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
func (it *NominaPortalXMsgMinGasLimitSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NominaPortalXMsgMinGasLimitSet)
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
		it.Event = new(NominaPortalXMsgMinGasLimitSet)
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
func (it *NominaPortalXMsgMinGasLimitSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NominaPortalXMsgMinGasLimitSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NominaPortalXMsgMinGasLimitSet represents a XMsgMinGasLimitSet event raised by the NominaPortal contract.
type NominaPortalXMsgMinGasLimitSet struct {
	GasLimit uint64
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterXMsgMinGasLimitSet is a free log retrieval operation binding the contract event 0x8c852a6291aa436654b167353bca4a4b0c3d024c7562cb5082e7c869bddabf3e.
//
// Solidity: event XMsgMinGasLimitSet(uint64 gasLimit)
func (_NominaPortal *NominaPortalFilterer) FilterXMsgMinGasLimitSet(opts *bind.FilterOpts) (*NominaPortalXMsgMinGasLimitSetIterator, error) {

	logs, sub, err := _NominaPortal.contract.FilterLogs(opts, "XMsgMinGasLimitSet")
	if err != nil {
		return nil, err
	}
	return &NominaPortalXMsgMinGasLimitSetIterator{contract: _NominaPortal.contract, event: "XMsgMinGasLimitSet", logs: logs, sub: sub}, nil
}

// WatchXMsgMinGasLimitSet is a free log subscription operation binding the contract event 0x8c852a6291aa436654b167353bca4a4b0c3d024c7562cb5082e7c869bddabf3e.
//
// Solidity: event XMsgMinGasLimitSet(uint64 gasLimit)
func (_NominaPortal *NominaPortalFilterer) WatchXMsgMinGasLimitSet(opts *bind.WatchOpts, sink chan<- *NominaPortalXMsgMinGasLimitSet) (event.Subscription, error) {

	logs, sub, err := _NominaPortal.contract.WatchLogs(opts, "XMsgMinGasLimitSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NominaPortalXMsgMinGasLimitSet)
				if err := _NominaPortal.contract.UnpackLog(event, "XMsgMinGasLimitSet", log); err != nil {
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

// ParseXMsgMinGasLimitSet is a log parse operation binding the contract event 0x8c852a6291aa436654b167353bca4a4b0c3d024c7562cb5082e7c869bddabf3e.
//
// Solidity: event XMsgMinGasLimitSet(uint64 gasLimit)
func (_NominaPortal *NominaPortalFilterer) ParseXMsgMinGasLimitSet(log types.Log) (*NominaPortalXMsgMinGasLimitSet, error) {
	event := new(NominaPortalXMsgMinGasLimitSet)
	if err := _NominaPortal.contract.UnpackLog(event, "XMsgMinGasLimitSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NominaPortalXReceiptIterator is returned from FilterXReceipt and is used to iterate over the raw logs and unpacked data for XReceipt events raised by the NominaPortal contract.
type NominaPortalXReceiptIterator struct {
	Event *NominaPortalXReceipt // Event containing the contract specifics and raw log

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
func (it *NominaPortalXReceiptIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NominaPortalXReceipt)
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
		it.Event = new(NominaPortalXReceipt)
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
func (it *NominaPortalXReceiptIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NominaPortalXReceiptIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NominaPortalXReceipt represents a XReceipt event raised by the NominaPortal contract.
type NominaPortalXReceipt struct {
	SourceChainId uint64
	ShardId       uint64
	Offset        uint64
	GasUsed       *big.Int
	Relayer       common.Address
	Success       bool
	Err           []byte
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterXReceipt is a free log retrieval operation binding the contract event 0x8277cab1f0fa69b34674f64a7d43f242b0bacece6f5b7e8652f1e0d88a9b873b.
//
// Solidity: event XReceipt(uint64 indexed sourceChainId, uint64 indexed shardId, uint64 indexed offset, uint256 gasUsed, address relayer, bool success, bytes err)
func (_NominaPortal *NominaPortalFilterer) FilterXReceipt(opts *bind.FilterOpts, sourceChainId []uint64, shardId []uint64, offset []uint64) (*NominaPortalXReceiptIterator, error) {

	var sourceChainIdRule []interface{}
	for _, sourceChainIdItem := range sourceChainId {
		sourceChainIdRule = append(sourceChainIdRule, sourceChainIdItem)
	}
	var shardIdRule []interface{}
	for _, shardIdItem := range shardId {
		shardIdRule = append(shardIdRule, shardIdItem)
	}
	var offsetRule []interface{}
	for _, offsetItem := range offset {
		offsetRule = append(offsetRule, offsetItem)
	}

	logs, sub, err := _NominaPortal.contract.FilterLogs(opts, "XReceipt", sourceChainIdRule, shardIdRule, offsetRule)
	if err != nil {
		return nil, err
	}
	return &NominaPortalXReceiptIterator{contract: _NominaPortal.contract, event: "XReceipt", logs: logs, sub: sub}, nil
}

// WatchXReceipt is a free log subscription operation binding the contract event 0x8277cab1f0fa69b34674f64a7d43f242b0bacece6f5b7e8652f1e0d88a9b873b.
//
// Solidity: event XReceipt(uint64 indexed sourceChainId, uint64 indexed shardId, uint64 indexed offset, uint256 gasUsed, address relayer, bool success, bytes err)
func (_NominaPortal *NominaPortalFilterer) WatchXReceipt(opts *bind.WatchOpts, sink chan<- *NominaPortalXReceipt, sourceChainId []uint64, shardId []uint64, offset []uint64) (event.Subscription, error) {

	var sourceChainIdRule []interface{}
	for _, sourceChainIdItem := range sourceChainId {
		sourceChainIdRule = append(sourceChainIdRule, sourceChainIdItem)
	}
	var shardIdRule []interface{}
	for _, shardIdItem := range shardId {
		shardIdRule = append(shardIdRule, shardIdItem)
	}
	var offsetRule []interface{}
	for _, offsetItem := range offset {
		offsetRule = append(offsetRule, offsetItem)
	}

	logs, sub, err := _NominaPortal.contract.WatchLogs(opts, "XReceipt", sourceChainIdRule, shardIdRule, offsetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NominaPortalXReceipt)
				if err := _NominaPortal.contract.UnpackLog(event, "XReceipt", log); err != nil {
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

// ParseXReceipt is a log parse operation binding the contract event 0x8277cab1f0fa69b34674f64a7d43f242b0bacece6f5b7e8652f1e0d88a9b873b.
//
// Solidity: event XReceipt(uint64 indexed sourceChainId, uint64 indexed shardId, uint64 indexed offset, uint256 gasUsed, address relayer, bool success, bytes err)
func (_NominaPortal *NominaPortalFilterer) ParseXReceipt(log types.Log) (*NominaPortalXReceipt, error) {
	event := new(NominaPortalXReceipt)
	if err := _NominaPortal.contract.UnpackLog(event, "XReceipt", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NominaPortalXReceiptMaxErrorSizeSetIterator is returned from FilterXReceiptMaxErrorSizeSet and is used to iterate over the raw logs and unpacked data for XReceiptMaxErrorSizeSet events raised by the NominaPortal contract.
type NominaPortalXReceiptMaxErrorSizeSetIterator struct {
	Event *NominaPortalXReceiptMaxErrorSizeSet // Event containing the contract specifics and raw log

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
func (it *NominaPortalXReceiptMaxErrorSizeSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NominaPortalXReceiptMaxErrorSizeSet)
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
		it.Event = new(NominaPortalXReceiptMaxErrorSizeSet)
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
func (it *NominaPortalXReceiptMaxErrorSizeSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NominaPortalXReceiptMaxErrorSizeSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NominaPortalXReceiptMaxErrorSizeSet represents a XReceiptMaxErrorSizeSet event raised by the NominaPortal contract.
type NominaPortalXReceiptMaxErrorSizeSet struct {
	Size uint16
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterXReceiptMaxErrorSizeSet is a free log retrieval operation binding the contract event 0x620bbea084306b66a8cc6b5b63830d6b3874f9d2438914e259ffd5065c33f7b0.
//
// Solidity: event XReceiptMaxErrorSizeSet(uint16 size)
func (_NominaPortal *NominaPortalFilterer) FilterXReceiptMaxErrorSizeSet(opts *bind.FilterOpts) (*NominaPortalXReceiptMaxErrorSizeSetIterator, error) {

	logs, sub, err := _NominaPortal.contract.FilterLogs(opts, "XReceiptMaxErrorSizeSet")
	if err != nil {
		return nil, err
	}
	return &NominaPortalXReceiptMaxErrorSizeSetIterator{contract: _NominaPortal.contract, event: "XReceiptMaxErrorSizeSet", logs: logs, sub: sub}, nil
}

// WatchXReceiptMaxErrorSizeSet is a free log subscription operation binding the contract event 0x620bbea084306b66a8cc6b5b63830d6b3874f9d2438914e259ffd5065c33f7b0.
//
// Solidity: event XReceiptMaxErrorSizeSet(uint16 size)
func (_NominaPortal *NominaPortalFilterer) WatchXReceiptMaxErrorSizeSet(opts *bind.WatchOpts, sink chan<- *NominaPortalXReceiptMaxErrorSizeSet) (event.Subscription, error) {

	logs, sub, err := _NominaPortal.contract.WatchLogs(opts, "XReceiptMaxErrorSizeSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NominaPortalXReceiptMaxErrorSizeSet)
				if err := _NominaPortal.contract.UnpackLog(event, "XReceiptMaxErrorSizeSet", log); err != nil {
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

// ParseXReceiptMaxErrorSizeSet is a log parse operation binding the contract event 0x620bbea084306b66a8cc6b5b63830d6b3874f9d2438914e259ffd5065c33f7b0.
//
// Solidity: event XReceiptMaxErrorSizeSet(uint16 size)
func (_NominaPortal *NominaPortalFilterer) ParseXReceiptMaxErrorSizeSet(log types.Log) (*NominaPortalXReceiptMaxErrorSizeSet, error) {
	event := new(NominaPortalXReceiptMaxErrorSizeSet)
	if err := _NominaPortal.contract.UnpackLog(event, "XReceiptMaxErrorSizeSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NominaPortalXSubValsetCutoffSetIterator is returned from FilterXSubValsetCutoffSet and is used to iterate over the raw logs and unpacked data for XSubValsetCutoffSet events raised by the NominaPortal contract.
type NominaPortalXSubValsetCutoffSetIterator struct {
	Event *NominaPortalXSubValsetCutoffSet // Event containing the contract specifics and raw log

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
func (it *NominaPortalXSubValsetCutoffSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NominaPortalXSubValsetCutoffSet)
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
		it.Event = new(NominaPortalXSubValsetCutoffSet)
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
func (it *NominaPortalXSubValsetCutoffSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NominaPortalXSubValsetCutoffSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NominaPortalXSubValsetCutoffSet represents a XSubValsetCutoffSet event raised by the NominaPortal contract.
type NominaPortalXSubValsetCutoffSet struct {
	Cutoff uint8
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterXSubValsetCutoffSet is a free log retrieval operation binding the contract event 0x1683dc51426224f6e37a3b41dd5849e2db1bfe22366d1d913fa0ef6f757e828f.
//
// Solidity: event XSubValsetCutoffSet(uint8 cutoff)
func (_NominaPortal *NominaPortalFilterer) FilterXSubValsetCutoffSet(opts *bind.FilterOpts) (*NominaPortalXSubValsetCutoffSetIterator, error) {

	logs, sub, err := _NominaPortal.contract.FilterLogs(opts, "XSubValsetCutoffSet")
	if err != nil {
		return nil, err
	}
	return &NominaPortalXSubValsetCutoffSetIterator{contract: _NominaPortal.contract, event: "XSubValsetCutoffSet", logs: logs, sub: sub}, nil
}

// WatchXSubValsetCutoffSet is a free log subscription operation binding the contract event 0x1683dc51426224f6e37a3b41dd5849e2db1bfe22366d1d913fa0ef6f757e828f.
//
// Solidity: event XSubValsetCutoffSet(uint8 cutoff)
func (_NominaPortal *NominaPortalFilterer) WatchXSubValsetCutoffSet(opts *bind.WatchOpts, sink chan<- *NominaPortalXSubValsetCutoffSet) (event.Subscription, error) {

	logs, sub, err := _NominaPortal.contract.WatchLogs(opts, "XSubValsetCutoffSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NominaPortalXSubValsetCutoffSet)
				if err := _NominaPortal.contract.UnpackLog(event, "XSubValsetCutoffSet", log); err != nil {
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

// ParseXSubValsetCutoffSet is a log parse operation binding the contract event 0x1683dc51426224f6e37a3b41dd5849e2db1bfe22366d1d913fa0ef6f757e828f.
//
// Solidity: event XSubValsetCutoffSet(uint8 cutoff)
func (_NominaPortal *NominaPortalFilterer) ParseXSubValsetCutoffSet(log types.Log) (*NominaPortalXSubValsetCutoffSet, error) {
	event := new(NominaPortalXSubValsetCutoffSet)
	if err := _NominaPortal.contract.UnpackLog(event, "XSubValsetCutoffSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NominaPortalXSubmitFromPausedIterator is returned from FilterXSubmitFromPaused and is used to iterate over the raw logs and unpacked data for XSubmitFromPaused events raised by the NominaPortal contract.
type NominaPortalXSubmitFromPausedIterator struct {
	Event *NominaPortalXSubmitFromPaused // Event containing the contract specifics and raw log

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
func (it *NominaPortalXSubmitFromPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NominaPortalXSubmitFromPaused)
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
		it.Event = new(NominaPortalXSubmitFromPaused)
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
func (it *NominaPortalXSubmitFromPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NominaPortalXSubmitFromPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NominaPortalXSubmitFromPaused represents a XSubmitFromPaused event raised by the NominaPortal contract.
type NominaPortalXSubmitFromPaused struct {
	ChainId uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterXSubmitFromPaused is a free log retrieval operation binding the contract event 0xab78810a0515df65f9f10bfbcb92d03d5df71d9fd3b9414e9ad831a5117d6daa.
//
// Solidity: event XSubmitFromPaused(uint64 indexed chainId)
func (_NominaPortal *NominaPortalFilterer) FilterXSubmitFromPaused(opts *bind.FilterOpts, chainId []uint64) (*NominaPortalXSubmitFromPausedIterator, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}

	logs, sub, err := _NominaPortal.contract.FilterLogs(opts, "XSubmitFromPaused", chainIdRule)
	if err != nil {
		return nil, err
	}
	return &NominaPortalXSubmitFromPausedIterator{contract: _NominaPortal.contract, event: "XSubmitFromPaused", logs: logs, sub: sub}, nil
}

// WatchXSubmitFromPaused is a free log subscription operation binding the contract event 0xab78810a0515df65f9f10bfbcb92d03d5df71d9fd3b9414e9ad831a5117d6daa.
//
// Solidity: event XSubmitFromPaused(uint64 indexed chainId)
func (_NominaPortal *NominaPortalFilterer) WatchXSubmitFromPaused(opts *bind.WatchOpts, sink chan<- *NominaPortalXSubmitFromPaused, chainId []uint64) (event.Subscription, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}

	logs, sub, err := _NominaPortal.contract.WatchLogs(opts, "XSubmitFromPaused", chainIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NominaPortalXSubmitFromPaused)
				if err := _NominaPortal.contract.UnpackLog(event, "XSubmitFromPaused", log); err != nil {
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

// ParseXSubmitFromPaused is a log parse operation binding the contract event 0xab78810a0515df65f9f10bfbcb92d03d5df71d9fd3b9414e9ad831a5117d6daa.
//
// Solidity: event XSubmitFromPaused(uint64 indexed chainId)
func (_NominaPortal *NominaPortalFilterer) ParseXSubmitFromPaused(log types.Log) (*NominaPortalXSubmitFromPaused, error) {
	event := new(NominaPortalXSubmitFromPaused)
	if err := _NominaPortal.contract.UnpackLog(event, "XSubmitFromPaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NominaPortalXSubmitFromUnpausedIterator is returned from FilterXSubmitFromUnpaused and is used to iterate over the raw logs and unpacked data for XSubmitFromUnpaused events raised by the NominaPortal contract.
type NominaPortalXSubmitFromUnpausedIterator struct {
	Event *NominaPortalXSubmitFromUnpaused // Event containing the contract specifics and raw log

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
func (it *NominaPortalXSubmitFromUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NominaPortalXSubmitFromUnpaused)
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
		it.Event = new(NominaPortalXSubmitFromUnpaused)
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
func (it *NominaPortalXSubmitFromUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NominaPortalXSubmitFromUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NominaPortalXSubmitFromUnpaused represents a XSubmitFromUnpaused event raised by the NominaPortal contract.
type NominaPortalXSubmitFromUnpaused struct {
	ChainId uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterXSubmitFromUnpaused is a free log retrieval operation binding the contract event 0xc551305d9bd408be4327b7f8aba28b04ccf6b6c76925392d195ecf9cc764294d.
//
// Solidity: event XSubmitFromUnpaused(uint64 indexed chainId)
func (_NominaPortal *NominaPortalFilterer) FilterXSubmitFromUnpaused(opts *bind.FilterOpts, chainId []uint64) (*NominaPortalXSubmitFromUnpausedIterator, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}

	logs, sub, err := _NominaPortal.contract.FilterLogs(opts, "XSubmitFromUnpaused", chainIdRule)
	if err != nil {
		return nil, err
	}
	return &NominaPortalXSubmitFromUnpausedIterator{contract: _NominaPortal.contract, event: "XSubmitFromUnpaused", logs: logs, sub: sub}, nil
}

// WatchXSubmitFromUnpaused is a free log subscription operation binding the contract event 0xc551305d9bd408be4327b7f8aba28b04ccf6b6c76925392d195ecf9cc764294d.
//
// Solidity: event XSubmitFromUnpaused(uint64 indexed chainId)
func (_NominaPortal *NominaPortalFilterer) WatchXSubmitFromUnpaused(opts *bind.WatchOpts, sink chan<- *NominaPortalXSubmitFromUnpaused, chainId []uint64) (event.Subscription, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}

	logs, sub, err := _NominaPortal.contract.WatchLogs(opts, "XSubmitFromUnpaused", chainIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NominaPortalXSubmitFromUnpaused)
				if err := _NominaPortal.contract.UnpackLog(event, "XSubmitFromUnpaused", log); err != nil {
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

// ParseXSubmitFromUnpaused is a log parse operation binding the contract event 0xc551305d9bd408be4327b7f8aba28b04ccf6b6c76925392d195ecf9cc764294d.
//
// Solidity: event XSubmitFromUnpaused(uint64 indexed chainId)
func (_NominaPortal *NominaPortalFilterer) ParseXSubmitFromUnpaused(log types.Log) (*NominaPortalXSubmitFromUnpaused, error) {
	event := new(NominaPortalXSubmitFromUnpaused)
	if err := _NominaPortal.contract.UnpackLog(event, "XSubmitFromUnpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NominaPortalXSubmitPausedIterator is returned from FilterXSubmitPaused and is used to iterate over the raw logs and unpacked data for XSubmitPaused events raised by the NominaPortal contract.
type NominaPortalXSubmitPausedIterator struct {
	Event *NominaPortalXSubmitPaused // Event containing the contract specifics and raw log

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
func (it *NominaPortalXSubmitPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NominaPortalXSubmitPaused)
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
		it.Event = new(NominaPortalXSubmitPaused)
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
func (it *NominaPortalXSubmitPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NominaPortalXSubmitPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NominaPortalXSubmitPaused represents a XSubmitPaused event raised by the NominaPortal contract.
type NominaPortalXSubmitPaused struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterXSubmitPaused is a free log retrieval operation binding the contract event 0x3d0f9c56dac46156a2db0aa09ee7804770ad9fc9549d21023164f22d69475ed8.
//
// Solidity: event XSubmitPaused()
func (_NominaPortal *NominaPortalFilterer) FilterXSubmitPaused(opts *bind.FilterOpts) (*NominaPortalXSubmitPausedIterator, error) {

	logs, sub, err := _NominaPortal.contract.FilterLogs(opts, "XSubmitPaused")
	if err != nil {
		return nil, err
	}
	return &NominaPortalXSubmitPausedIterator{contract: _NominaPortal.contract, event: "XSubmitPaused", logs: logs, sub: sub}, nil
}

// WatchXSubmitPaused is a free log subscription operation binding the contract event 0x3d0f9c56dac46156a2db0aa09ee7804770ad9fc9549d21023164f22d69475ed8.
//
// Solidity: event XSubmitPaused()
func (_NominaPortal *NominaPortalFilterer) WatchXSubmitPaused(opts *bind.WatchOpts, sink chan<- *NominaPortalXSubmitPaused) (event.Subscription, error) {

	logs, sub, err := _NominaPortal.contract.WatchLogs(opts, "XSubmitPaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NominaPortalXSubmitPaused)
				if err := _NominaPortal.contract.UnpackLog(event, "XSubmitPaused", log); err != nil {
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

// ParseXSubmitPaused is a log parse operation binding the contract event 0x3d0f9c56dac46156a2db0aa09ee7804770ad9fc9549d21023164f22d69475ed8.
//
// Solidity: event XSubmitPaused()
func (_NominaPortal *NominaPortalFilterer) ParseXSubmitPaused(log types.Log) (*NominaPortalXSubmitPaused, error) {
	event := new(NominaPortalXSubmitPaused)
	if err := _NominaPortal.contract.UnpackLog(event, "XSubmitPaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// NominaPortalXSubmitUnpausedIterator is returned from FilterXSubmitUnpaused and is used to iterate over the raw logs and unpacked data for XSubmitUnpaused events raised by the NominaPortal contract.
type NominaPortalXSubmitUnpausedIterator struct {
	Event *NominaPortalXSubmitUnpaused // Event containing the contract specifics and raw log

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
func (it *NominaPortalXSubmitUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(NominaPortalXSubmitUnpaused)
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
		it.Event = new(NominaPortalXSubmitUnpaused)
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
func (it *NominaPortalXSubmitUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *NominaPortalXSubmitUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// NominaPortalXSubmitUnpaused represents a XSubmitUnpaused event raised by the NominaPortal contract.
type NominaPortalXSubmitUnpaused struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterXSubmitUnpaused is a free log retrieval operation binding the contract event 0x2cb9d71d4c31860b70e9b707c69aa2f5953e03474f00cfcfff205c4745f82875.
//
// Solidity: event XSubmitUnpaused()
func (_NominaPortal *NominaPortalFilterer) FilterXSubmitUnpaused(opts *bind.FilterOpts) (*NominaPortalXSubmitUnpausedIterator, error) {

	logs, sub, err := _NominaPortal.contract.FilterLogs(opts, "XSubmitUnpaused")
	if err != nil {
		return nil, err
	}
	return &NominaPortalXSubmitUnpausedIterator{contract: _NominaPortal.contract, event: "XSubmitUnpaused", logs: logs, sub: sub}, nil
}

// WatchXSubmitUnpaused is a free log subscription operation binding the contract event 0x2cb9d71d4c31860b70e9b707c69aa2f5953e03474f00cfcfff205c4745f82875.
//
// Solidity: event XSubmitUnpaused()
func (_NominaPortal *NominaPortalFilterer) WatchXSubmitUnpaused(opts *bind.WatchOpts, sink chan<- *NominaPortalXSubmitUnpaused) (event.Subscription, error) {

	logs, sub, err := _NominaPortal.contract.WatchLogs(opts, "XSubmitUnpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(NominaPortalXSubmitUnpaused)
				if err := _NominaPortal.contract.UnpackLog(event, "XSubmitUnpaused", log); err != nil {
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

// ParseXSubmitUnpaused is a log parse operation binding the contract event 0x2cb9d71d4c31860b70e9b707c69aa2f5953e03474f00cfcfff205c4745f82875.
//
// Solidity: event XSubmitUnpaused()
func (_NominaPortal *NominaPortalFilterer) ParseXSubmitUnpaused(log types.Log) (*NominaPortalXSubmitUnpaused, error) {
	event := new(NominaPortalXSubmitUnpaused)
	if err := _NominaPortal.contract.UnpackLog(event, "XSubmitUnpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
