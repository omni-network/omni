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

// OmniPortalInitParams is an auto generated low-level Go binding around an user-defined struct.
type OmniPortalInitParams struct {
	Owner                common.Address
	FeeOracle            common.Address
	OmniChainId          uint64
	OmniCChainId         uint64
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
type XTypesBlockHeader struct {
	SourceChainId     uint64
	ConfLevel         uint8
	Offset            uint64
	SourceBlockHeight uint64
	SourceBlockHash   [32]byte
}

// XTypesChain is an auto generated low-level Go binding around an user-defined struct.
type XTypesChain struct {
	ChainId uint64
	Shards  []uint64
}

// XTypesMsg is an auto generated low-level Go binding around an user-defined struct.
type XTypesMsg struct {
	SourceChainId uint64
	DestChainId   uint64
	ShardId       uint64
	Offset        uint64
	Sender        common.Address
	To            common.Address
	Data          []byte
	GasLimit      uint64
}

// XTypesMsgShort is an auto generated low-level Go binding around an user-defined struct.
type XTypesMsgShort struct {
	SourceChainId uint64
	Sender        common.Address
}

// XTypesSigTuple is an auto generated low-level Go binding around an user-defined struct.
type XTypesSigTuple struct {
	ValidatorAddr common.Address
	Signature     []byte
}

// XTypesSubmission is an auto generated low-level Go binding around an user-defined struct.
type XTypesSubmission struct {
	AttestationRoot [32]byte
	ValidatorSetId  uint64
	BlockHeader     XTypesBlockHeader
	Msgs            []XTypesMsg
	Proof           [][32]byte
	ProofFlags      []bool
	Signatures      []XTypesSigTuple
}

// XTypesValidator is an auto generated low-level Go binding around an user-defined struct.
type XTypesValidator struct {
	Addr  common.Address
	Power uint64
}

// OmniPortalMetaData contains all meta data concerning the OmniPortal contract.
var OmniPortalMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"ActionXCall\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"ActionXSubmit\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"KeyPauseAll\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"XSubQuorumDenominator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"XSubQuorumNumerator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"addValidatorSet\",\"inputs\":[{\"name\":\"valSetId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"validators\",\"type\":\"tuple[]\",\"internalType\":\"structXTypes.Validator[]\",\"components\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"power\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"chainId\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"collectFees\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"feeFor\",\"inputs\":[{\"name\":\"destChainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"gasLimit\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"feeOracle\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"inXBlockOffset\",\"inputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"inXMsgOffset\",\"inputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"initialize\",\"inputs\":[{\"name\":\"p\",\"type\":\"tuple\",\"internalType\":\"structOmniPortal.InitParams\",\"components\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"feeOracle\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"omniChainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"omniCChainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"xmsgMaxGasLimit\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"xmsgMinGasLimit\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"xmsgMaxDataSize\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"xreceiptMaxErrorSize\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"xsubValsetCutoff\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"cChainXMsgOffset\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"cChainXBlockOffset\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"valSetId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"validators\",\"type\":\"tuple[]\",\"internalType\":\"structXTypes.Validator[]\",\"components\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"power\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isPaused\",\"inputs\":[{\"name\":\"actionId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isPaused\",\"inputs\":[{\"name\":\"actionId\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"chainId_\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isPaused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isSupportedDest\",\"inputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isSupportedShard\",\"inputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isXCall\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"latestValSetId\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"network\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"chainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"omniCChainId\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"omniChainId\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"outXMsgOffset\",\"inputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pauseXCall\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pauseXCallTo\",\"inputs\":[{\"name\":\"chainId_\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pauseXSubmit\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pauseXSubmitFrom\",\"inputs\":[{\"name\":\"chainId_\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setFeeOracle\",\"inputs\":[{\"name\":\"feeOracle_\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setInXBlockOffset\",\"inputs\":[{\"name\":\"sourceChainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"shardId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"offset\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setInXMsgOffset\",\"inputs\":[{\"name\":\"sourceChainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"shardId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"offset\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setNetwork\",\"inputs\":[{\"name\":\"network_\",\"type\":\"tuple[]\",\"internalType\":\"structXTypes.Chain[]\",\"components\":[{\"name\":\"chainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"shards\",\"type\":\"uint64[]\",\"internalType\":\"uint64[]\"}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setXMsgMaxDataSize\",\"inputs\":[{\"name\":\"numBytes\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setXMsgMaxGasLimit\",\"inputs\":[{\"name\":\"gasLimit\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setXMsgMinGasLimit\",\"inputs\":[{\"name\":\"gasLimit\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setXReceiptMaxErrorSize\",\"inputs\":[{\"name\":\"numBytes\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setXSubValsetCutoff\",\"inputs\":[{\"name\":\"xsubValsetCutoff_\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unpauseXCall\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unpauseXCallTo\",\"inputs\":[{\"name\":\"chainId_\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unpauseXSubmit\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unpauseXSubmitFrom\",\"inputs\":[{\"name\":\"chainId_\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"valSet\",\"inputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"valSetTotalPower\",\"inputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"xcall\",\"inputs\":[{\"name\":\"destChainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"conf\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"gasLimit\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"xmsg\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structXTypes.MsgShort\",\"components\":[{\"name\":\"sourceChainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"xmsgMaxDataSize\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"xmsgMaxGasLimit\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"xmsgMinGasLimit\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"xreceiptMaxErrorSize\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"xsubValsetCutoff\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"uint8\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"xsubmit\",\"inputs\":[{\"name\":\"xsub\",\"type\":\"tuple\",\"internalType\":\"structXTypes.Submission\",\"components\":[{\"name\":\"attestationRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"validatorSetId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"blockHeader\",\"type\":\"tuple\",\"internalType\":\"structXTypes.BlockHeader\",\"components\":[{\"name\":\"sourceChainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"confLevel\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"offset\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"sourceBlockHeight\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"sourceBlockHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"msgs\",\"type\":\"tuple[]\",\"internalType\":\"structXTypes.Msg[]\",\"components\":[{\"name\":\"sourceChainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"destChainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"shardId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"offset\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"sender\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"gasLimit\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"proof\",\"type\":\"bytes32[]\",\"internalType\":\"bytes32[]\"},{\"name\":\"proofFlags\",\"type\":\"bool[]\",\"internalType\":\"bool[]\"},{\"name\":\"signatures\",\"type\":\"tuple[]\",\"internalType\":\"structXTypes.SigTuple[]\",\"components\":[{\"name\":\"validatorAddr\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}]}]}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"FeeOracleSet\",\"inputs\":[{\"name\":\"oracle\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"FeesCollected\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"InXBlockOffsetSet\",\"inputs\":[{\"name\":\"srcChainId\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"shardId\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"offset\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"InXMsgOffsetSet\",\"inputs\":[{\"name\":\"srcChainId\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"shardId\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"offset\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ValidatorSetAdded\",\"inputs\":[{\"name\":\"setId\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"XCallPaused\",\"inputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"XCallToPaused\",\"inputs\":[{\"name\":\"chainId\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"XCallToUnpaused\",\"inputs\":[{\"name\":\"chainId\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"XCallUnpaused\",\"inputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"XMsg\",\"inputs\":[{\"name\":\"destChainId\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"shardId\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"offset\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"sender\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"to\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"},{\"name\":\"gasLimit\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"},{\"name\":\"fees\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"XMsgMaxDataSizeSet\",\"inputs\":[{\"name\":\"size\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"XMsgMaxGasLimitSet\",\"inputs\":[{\"name\":\"gasLimit\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"XMsgMinGasLimitSet\",\"inputs\":[{\"name\":\"gasLimit\",\"type\":\"uint64\",\"indexed\":false,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"XReceipt\",\"inputs\":[{\"name\":\"sourceChainId\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"shardId\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"offset\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"gasUsed\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"relayer\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"success\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"},{\"name\":\"error\",\"type\":\"bytes\",\"indexed\":false,\"internalType\":\"bytes\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"XReceiptMaxErrorSizeSet\",\"inputs\":[{\"name\":\"size\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"XSubValsetCutoffSet\",\"inputs\":[{\"name\":\"cutoff\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"XSubmitFromPaused\",\"inputs\":[{\"name\":\"chainId\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"XSubmitFromUnpaused\",\"inputs\":[{\"name\":\"chainId\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"XSubmitPaused\",\"inputs\":[],\"anonymous\":false},{\"type\":\"event\",\"name\":\"XSubmitUnpaused\",\"inputs\":[],\"anonymous\":false},{\"type\":\"error\",\"name\":\"ECDSAInvalidSignature\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ECDSAInvalidSignatureLength\",\"inputs\":[{\"name\":\"length\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"type\":\"error\",\"name\":\"ECDSAInvalidSignatureS\",\"inputs\":[{\"name\":\"s\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"type\":\"error\",\"name\":\"InvalidInitialization\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MerkleProofInvalidMultiproof\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotInitializing\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnableInvalidOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnableUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ReentrancyGuardReentrantCall\",\"inputs\":[]}]",
	Bin: "0x60806040523480156200001157600080fd5b506200001c62000022565b620000d6565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a00805468010000000000000000900460ff1615620000735760405163f92ee8a960e01b815260040160405180910390fd5b80546001600160401b0390811614620000d35780546001600160401b0319166001600160401b0390811782556040519081527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b50565b614b8480620000e66000396000f3fe60806040526004361061036b5760003560e01c80638532eb9f116101c6578063b4d5afd1116100f7578063c3d8ad6711610095578063d051c97d1161006f578063d051c97d14610af6578063d533b44514610b37578063f2fde38b14610b57578063f45cc7b814610b7757600080fd5b8063c3d8ad6714610a98578063c4ab80bc14610aad578063cf84c81814610acd57600080fd5b8063bff0e84d116100d1578063bff0e84d14610a25578063c21dda4f14610a45578063c26dfc0514610a58578063c2f9b96814610a7857600080fd5b8063b4d5afd1146109b0578063b521466d146109e5578063bb8590ad14610a0557600080fd5b8063a480ca7911610164578063afe821981161013e578063afe8219814610923578063afe8af9c14610943578063b187bd2614610979578063b2b2f5bd1461098e57600080fd5b8063a480ca79146108b3578063a8a98962146108d3578063aaf1bc97146108f357600080fd5b806397b52062116101a057806397b520621461083c5780639a8a05921461085c578063a10ac97a1461086f578063a32eb7c61461089157600080fd5b80638532eb9f146107b15780638da5cb5b146107d15780638dd9523c1461080e57600080fd5b80633aa87330116102a057806355e2448e1161023e57806374eba9391161021857806374eba9391461074057806378fe53071461076057806383d0cbd9146107875780638456cb591461079c57600080fd5b806355e2448e146106ca57806357542050146106ea578063715018a61461072b57600080fd5b8063461ab4881161027a578063461ab4881461064357806349cc3bf614610663578063500b19e71461067d57806354d26bba146106b557600080fd5b80633aa87330146105ac5780633f4ba83a146105ed5780633fd3b15e1461060257600080fd5b806323f2ca201161030d5780632f32700e116102e75780632f32700e146104e057806330632e8b1461054557806336d219121461056557806336d853f91461058c57600080fd5b806323f2ca2014610460578063241b71bb1461048057806324278bbe146104b057600080fd5b806310a5a7f71161034957806310a5a7f7146103d3578063110ff5f1146103f35780631d3eb6e31461042b57806323dbce501461044b57600080fd5b80630360d20f1461037057806306c3dc5f1461039c578063103ba701146103b1575b600080fd5b34801561037c57600080fd5b50610385600281565b60405160ff90911681526020015b60405180910390f35b3480156103a857600080fd5b50610385600381565b3480156103bd57600080fd5b506103d16103cc366004613d40565b610b9e565b005b3480156103df57600080fd5b506103d16103ee366004613d7b565b610bb2565b3480156103ff57600080fd5b50600154610413906001600160401b031681565b6040516001600160401b039091168152602001610393565b34801561043757600080fd5b506103d1610446366004613d98565b610c11565b34801561045757600080fd5b506103d1610d2c565b34801561046c57600080fd5b506103d161047b366004613e0c565b610d76565b34801561048c57600080fd5b506104a061049b366004613e47565b6110ae565b6040519015158152602001610393565b3480156104bc57600080fd5b506104a06104cb366004613d7b565b60056020526000908152604090205460ff1681565b3480156104ec57600080fd5b50604080518082018252600080825260209182015281518083018352600b546001600160401b0381168083526001600160a01b03600160401b909204821692840192835284519081529151169181019190915201610393565b34801561055157600080fd5b506103d1610560366004613e60565b6110bf565b34801561057157600080fd5b5060015461041390600160401b90046001600160401b031681565b34801561059857600080fd5b506103d16105a7366004613d7b565b6113d6565b3480156105b857600080fd5b506104136105c7366004613e9b565b60066020908152600092835260408084209091529082529020546001600160401b031681565b3480156105f957600080fd5b506103d16113e7565b34801561060e57600080fd5b5061041361061d366004613e9b565b60086020908152600092835260408084209091529082529020546001600160401b031681565b34801561064f57600080fd5b506104a061065e366004613ed4565b611422565b34801561066f57600080fd5b506000546103859060ff1681565b34801561068957600080fd5b5060025461069d906001600160a01b031681565b6040516001600160a01b039091168152602001610393565b3480156106c157600080fd5b506103d1611439565b3480156106d657600080fd5b50600b546001600160401b031615156104a0565b3480156106f657600080fd5b50610413610705366004613f10565b600a6020908152600092835260408084209091529082529020546001600160401b031681565b34801561073757600080fd5b506103d1611483565b34801561074c57600080fd5b5061041361075b366004613e47565b611497565b34801561076c57600080fd5b5060005461041390600160681b90046001600160401b031681565b34801561079357600080fd5b506103d16114c6565b3480156107a857600080fd5b506103d1611510565b3480156107bd57600080fd5b506103d16107cc366004613f45565b61154b565b3480156107dd57600080fd5b507f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c199300546001600160a01b031661069d565b34801561081a57600080fd5b5061082e610829366004614014565b61165e565b604051908152602001610393565b34801561084857600080fd5b506103d161085736600461407b565b6116df565b34801561086857600080fd5b5046610413565b34801561087b57600080fd5b5061082e600080516020614aef83398151915281565b34801561089d57600080fd5b5061082e600080516020614b2f83398151915281565b3480156108bf57600080fd5b506103d16108ce3660046140c6565b61175c565b3480156108df57600080fd5b506103d16108ee3660046140c6565b6117e4565b3480156108ff57600080fd5b506104a061090e366004613d7b565b60046020526000908152604090205460ff1681565b34801561092f57600080fd5b506103d161093e366004613d7b565b6117f5565b34801561094f57600080fd5b5061041361095e366004613d7b565b6009602052600090815260409020546001600160401b031681565b34801561098557600080fd5b506104a061184f565b34801561099a57600080fd5b5061082e600080516020614acf83398151915281565b3480156109bc57600080fd5b506000546109d2906301000000900461ffff1681565b60405161ffff9091168152602001610393565b3480156109f157600080fd5b506103d1610a003660046140e1565b6118a5565b348015610a1157600080fd5b506103d1610a20366004613d7b565b6118b6565b348015610a3157600080fd5b506103d1610a403660046140e1565b6118c7565b6103d1610a53366004614105565b6118d8565b348015610a6457600080fd5b506000546109d290610100900461ffff1681565b348015610a8457600080fd5b506103d1610a93366004613d7b565b611cb2565b348015610aa457600080fd5b506103d1611d11565b348015610ab957600080fd5b506103d1610ac836600461407b565b611d5b565b348015610ad957600080fd5b50600054610413906501000000000090046001600160401b031681565b348015610b0257600080fd5b50610413610b11366004613e9b565b60076020908152600092835260408084209091529082529020546001600160401b031681565b348015610b4357600080fd5b506103d1610b52366004613d7b565b611dcf565b348015610b6357600080fd5b506103d1610b723660046140c6565b611e29565b348015610b8357600080fd5b5060005461041390600160a81b90046001600160401b031681565b610ba6611e64565b610baf81611ebf565b50565b610bba611e64565b610bda610bd5600080516020614acf83398151915283611f5b565b611fa4565b6040516001600160401b038216907fcd7910e1c5569d8433ce4ef8e5d51c1bdc03168f614b576da47dc3d2b51d033a90600090a250565b333014610c5d5760405162461bcd60e51b815260206004820152601560248201527427b6b734a837b93a30b61d1037b7363c9039b2b63360591b60448201526064015b60405180910390fd5b600154600b546001600160401b03908116600160401b9092041614610cbe5760405162461bcd60e51b815260206004820152601760248201527627b6b734a837b93a30b61d1037b7363c9031b1b430b4b760491b6044820152606401610c54565b600b54600160401b90046001600160a01b031615610d1e5760405162461bcd60e51b815260206004820152601e60248201527f4f6d6e69506f7274616c3a206f6e6c792063636861696e2073656e64657200006044820152606401610c54565b610d28828261201f565b5050565b610d34611e64565b610d4b600080516020614b2f833981519152611fa4565b6040517f3d0f9c56dac46156a2db0aa09ee7804770ad9fc9549d21023164f22d69475ed890600090a1565b600080516020614b2f833981519152610d956060830160408401613d7b565b610da882610da38484611f5b565b612199565b15610dea5760405162461bcd60e51b815260206004820152601260248201527113db5b9a541bdc9d185b0e881c185d5cd95960721b6044820152606401610c54565b610df2612220565b366000610e0260e086018661418e565b9092509050604085016000610e1a8260208901613d7b565b905082610e605760405162461bcd60e51b81526020600482015260146024820152734f6d6e69506f7274616c3a206e6f20786d73677360601b6044820152606401610c54565b6001600160401b0380821660009081526009602052604090205416610ec75760405162461bcd60e51b815260206004820152601b60248201527f4f6d6e69506f7274616c3a20756e6b6e6f776e2076616c2073657400000000006044820152606401610c54565b610ecf61226a565b6001600160401b0316816001600160401b03161015610f305760405162461bcd60e51b815260206004820152601760248201527f4f6d6e69506f7274616c3a206f6c642076616c207365740000000000000000006044820152606401610c54565b610f748735610f436101408a018a61418e565b6001600160401b038086166000908152600a60209081526040808320600990925290912054909116600260036122ba565b610fb85760405162461bcd60e51b81526020600482015260156024820152744f6d6e69506f7274616c3a206e6f2071756f72756d60581b6044820152606401610c54565b610fe18735838686610fce6101008d018d61418e565b610fdc6101208f018f61418e565b6124dc565b61102d5760405162461bcd60e51b815260206004820152601960248201527f4f6d6e69506f7274616c3a20696e76616c69642070726f6f66000000000000006044820152606401610c54565b60005b8381101561107b5761107361104a36859003850185614245565b86868481811061105c5761105c6142d3565b905060200281019061106e91906142e9565b612557565b600101611030565b50505050506110a960017f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f0055565b505050565b60006110b982612abd565b92915050565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a008054600160401b810460ff1615906001600160401b03166000811580156111045750825b90506000826001600160401b031660011480156111205750303b155b90508115801561112e575080155b1561114c5760405163f92ee8a960e01b815260040160405180910390fd5b845467ffffffffffffffff19166001178555831561117657845460ff60401b1916600160401b1785555b61118b61118660208801886140c6565b612b24565b6111a361119e60408801602089016140c6565b612b35565b6111bb6111b660a0880160808901613d7b565b612bd9565b6111d36111ce60c0880160a08901613d7b565b612c91565b6111eb6111e660e0880160c089016140e1565b612d45565b6112046111ff610100880160e089016140e1565b612de9565b61121e61121961012088016101008901613d40565b611ebf565b61124661123361018088016101608901613d7b565b611241610180890189614309565b612e89565b6112566060870160408801613d7b565b6001805467ffffffffffffffff19166001600160401b03929092169190911790556112876080870160608801613d7b565b600180546001600160401b0392909216600160401b026fffffffffffffffff0000000000000000199092169190911790556101046112cd61014088016101208901613d7b565b600760006112e160808b0160608c01613d7b565b6001600160401b0390811682526020808301939093526040918201600090812086831682529093529120805467ffffffffffffffff19169290911691909117905561133461016088016101408901613d7b565b6008600061134860808b0160608c01613d7b565b6001600160401b03908116825260208083019390935260409182016000908120958216815294909252909220805467ffffffffffffffff19169190921617905583156113ce57845460ff60401b19168555604051600181527fc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d29060200160405180910390a15b505050505050565b6113de611e64565b610baf81612bd9565b6113ef611e64565b6113f76131b8565b6040517fa45f47fdea8a1efdd9029a5691c7f759c32b7c698632b563573e155625d1693390600090a1565b600061143283610da38585611f5b565b9392505050565b611441611e64565b611458600080516020614acf8339815191526131cb565b6040517f4c48c7b71557216a3192842746bdfc381f98d7536d9eb1c6764f3b45e679482790600090a1565b61148b611e64565b6114956000613246565b565b600381815481106114a757600080fd5b60009182526020909120600290910201546001600160401b0316905081565b6114ce611e64565b6114e5600080516020614acf833981519152611fa4565b6040517f5f335a4032d4cfb6aca7835b0c2225f36d4d9eaa4ed43ee59ed537e02dff6b3990600090a1565b611518611e64565b6115206132b7565b6040517f9e87fac88ff661f02d44f95383c817fece4bce600a3dab7a54406878b965e75290600090a1565b3330146115925760405162461bcd60e51b815260206004820152601560248201527427b6b734a837b93a30b61d1037b7363c9039b2b63360591b6044820152606401610c54565b600154600b546001600160401b03908116600160401b90920416146115f35760405162461bcd60e51b815260206004820152601760248201527627b6b734a837b93a30b61d1037b7363c9031b1b430b4b760491b6044820152606401610c54565b600b54600160401b90046001600160a01b0316156116535760405162461bcd60e51b815260206004820152601e60248201527f4f6d6e69506f7274616c3a206f6e6c792063636861696e2073656e64657200006044820152606401610c54565b6110a9838383612e89565b600254604051632376548f60e21b81526000916001600160a01b031690638dd9523c9061169590889088908890889060040161437b565b602060405180830381865afa1580156116b2573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906116d691906143b3565b95945050505050565b6116e7611e64565b6001600160401b03838116600081815260086020908152604080832087861680855290835292819020805467ffffffffffffffff191695871695861790555193845290927fe070f08cae8464c91238e8cbea64ccee5e7b48dd79a843f144e3721ee6bdd9b591015b60405180910390a3505050565b611764611e64565b60405147906001600160a01b0383169082156108fc029083906000818181858888f1935050505015801561179c573d6000803e3d6000fd5b50816001600160a01b03167f9dc46f23cfb5ddcad0ae7ea2be38d47fec07bb9382ec7e564efc69e036dd66ce826040516117d891815260200190565b60405180910390a25050565b6117ec611e64565b610baf81612b35565b6117fd611e64565b611818610bd5600080516020614b2f83398151915283611f5b565b6040516001600160401b038216907fab78810a0515df65f9f10bfbcb92d03d5df71d9fd3b9414e9ad831a5117d6daa90600090a250565b60006118a0600080516020614aef833981519152600052600080516020614b0f8339815191526020527ffae9838a178d7f201aa98e2ce5340158edda60bb1e8f168f46503bf3e99f13be5460ff1690565b905090565b6118ad611e64565b610baf81612d45565b6118be611e64565b610baf81612c91565b6118cf611e64565b610baf81612de9565b600080516020614acf833981519152866118f682610da38484611f5b565b156119385760405162461bcd60e51b815260206004820152601260248201527113db5b9a541bdc9d185b0e881c185d5cd95960721b6044820152606401610c54565b6001600160401b03881660009081526005602052604090205460ff166119a05760405162461bcd60e51b815260206004820152601c60248201527f4f6d6e69506f7274616c3a20756e737570706f727465642064657374000000006044820152606401610c54565b6001600160a01b0386166119f65760405162461bcd60e51b815260206004820152601b60248201527f4f6d6e69506f7274616c3a206e6f20706f7274616c207863616c6c00000000006044820152606401610c54565b6000546001600160401b036501000000000090910481169084161115611a5e5760405162461bcd60e51b815260206004820152601d60248201527f4f6d6e69506f7274616c3a206761734c696d697420746f6f20686967680000006044820152606401610c54565b6000546001600160401b03600160681b90910481169084161015611ac45760405162461bcd60e51b815260206004820152601c60248201527f4f6d6e69506f7274616c3a206761734c696d697420746f6f206c6f77000000006044820152606401610c54565b6000546301000000900461ffff16841115611b215760405162461bcd60e51b815260206004820152601a60248201527f4f6d6e69506f7274616c3a206461746120746f6f206c617267650000000000006044820152606401610c54565b60ff808816600081815260046020526040902054909116611b845760405162461bcd60e51b815260206004820152601d60248201527f4f6d6e69506f7274616c3a20756e737570706f727465642073686172640000006044820152606401610c54565b6000611b928a88888861165e565b905080341015611be45760405162461bcd60e51b815260206004820152601c60248201527f4f6d6e69506f7274616c3a20696e73756666696369656e7420666565000000006044820152606401610c54565b6001600160401b03808b166000908152600660209081526040808320868516845290915281208054600193919291611c1e918591166143e2565b82546101009290920a6001600160401b038181021990931691831602179091558b811660008181526006602090815260408083208886168085529252918290205491519190931693507fb7c8eb9d7a7fbcdab809ab7b8a7c41701eb3115e3fe99d30ff490d8552f72bfa90611c9e9033908e908e908e908e908b90614409565b60405180910390a450505050505050505050565b611cba611e64565b611cda611cd5600080516020614b2f83398151915283611f5b565b6131cb565b6040516001600160401b038216907fc551305d9bd408be4327b7f8aba28b04ccf6b6c76925392d195ecf9cc764294d90600090a250565b611d19611e64565b611d30600080516020614b2f8339815191526131cb565b6040517f2cb9d71d4c31860b70e9b707c69aa2f5953e03474f00cfcfff205c4745f8287590600090a1565b611d63611e64565b6001600160401b03838116600081815260076020908152604080832087861680855290835292819020805467ffffffffffffffff191695871695861790555193845290927f8647aae68c8456a1dcbfaf5eaadc94278ae423526d3f09c7b972bff7355d55c7910161174f565b611dd7611e64565b611df2611cd5600080516020614acf83398151915283611f5b565b6040516001600160401b038216907f1ed9223556fb0971076c30172f1f00630efd313b6a05290a562aef95928e712590600090a250565b611e31611e64565b6001600160a01b038116611e5b57604051631e4fbdf760e01b815260006004820152602401610c54565b610baf81613246565b33611e967f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c199300546001600160a01b031690565b6001600160a01b0316146114955760405163118cdaa760e01b8152336004820152602401610c54565b60008160ff1611611f125760405162461bcd60e51b815260206004820152601a60248201527f4f6d6e69506f7274616c3a206e6f207a65726f206375746f66660000000000006044820152606401610c54565b6000805460ff191660ff83169081179091556040519081527f1683dc51426224f6e37a3b41dd5849e2db1bfe22366d1d913fa0ef6f757e828f906020015b60405180910390a150565b60008282604051602001611f8692919091825260c01b6001600160c01b031916602082015260280190565b60405160208183030381529060405280519060200120905092915050565b6000818152600080516020614b0f833981519152602081905260409091205460ff16156120065760405162461bcd60e51b815260206004820152601060248201526f14185d5cd8589b194e881c185d5cd95960821b6044820152606401610c54565b600091825260205260409020805460ff19166001179055565b6120276132ce565b3660005b8281101561219357838382818110612045576120456142d3565b90506020028101906120579190614454565b6003805460018101825560009190915290925082906002027fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b0161209b82826144f0565b50506120a44690565b6001600160401b03166120ba6020840184613d7b565b6001600160401b031614612108576001600560006120db6020860186613d7b565b6001600160401b031681526020810191909152604001600020805460ff191691151591909117905561218b565b60005b612118602084018461418e565b905081101561218957600160046000612134602087018761418e565b85818110612144576121446142d3565b90506020020160208101906121599190613d7b565b6001600160401b031681526020810191909152604001600020805460ff191691151591909117905560010161210b565b505b60010161202b565b50505050565b600080516020614aef8339815191526000908152600080516020614b0f83398151915260208190527ffae9838a178d7f201aa98e2ce5340158edda60bb1e8f168f46503bf3e99f13be5460ff16806121ff575060008481526020829052604090205460ff165b80612218575060008381526020829052604090205460ff165b949350505050565b7f9b779b17422d0df92223018b32b4d1fa46e071723d6817e2486d003becc55f0080546001190161226457604051633ee5aeb560e01b815260040160405180910390fd5b60029055565b6000805460ff8116600160a81b9091046001600160401b03161161228e5750600190565b6000546122af9060ff811690600160a81b90046001600160401b0316614618565b6118a09060016143e2565b6000803660005b888110156124c9578989828181106122db576122db6142d3565b90506020028101906122ed9190614454565b9150801561240f5760008a8a612304600185614638565b818110612313576123136142d3565b90506020028101906123259190614454565b61232e9061464b565b80519091506001600160a01b031661234960208501856140c6565b6001600160a01b03160361239f5760405162461bcd60e51b815260206004820152601b60248201527f51756f72756d3a206475706c69636174652076616c696461746f7200000000006044820152606401610c54565b80516001600160a01b03166123b760208501856140c6565b6001600160a01b03161161240d5760405162461bcd60e51b815260206004820152601760248201527f51756f72756d3a2073696773206e6f7420736f727465640000000000000000006044820152606401610c54565b505b612419828c6133cd565b6124655760405162461bcd60e51b815260206004820152601960248201527f51756f72756d3a20696e76616c6964207369676e6174757265000000000000006044820152606401610c54565b87600061247560208501856140c6565b6001600160a01b031681526020810191909152604001600020546124a2906001600160401b0316846143e2565b92506124b083888888613441565b156124c157600193505050506124d1565b6001016122c1565b506000925050505b979650505050505050565b60408051600180825281830190925260009182919060208083019080368337019050509050612517868686866125128d8d61347e565b61354b565b8160008151811061252a5761252a6142d3565b602002602001018181525050612549818b6125448c6137ac565b6137c4565b9a9950505050505050505050565b60006125666020830183613d7b565b9050600061257a6040840160208501613d7b565b9050600061258e6060850160408601613d7b565b905060006125a26080860160608701613d7b565b905085600001516001600160401b0316846001600160401b0316146126095760405162461bcd60e51b815260206004820152601e60248201527f4f6d6e69506f7274616c3a2077726f6e6720736f7572636520636861696e00006044820152606401610c54565b466001600160401b0316836001600160401b0316148061263057506001600160401b038316155b61267c5760405162461bcd60e51b815260206004820152601c60248201527f4f6d6e69506f7274616c3a2077726f6e67206465737420636861696e000000006044820152606401610c54565b6001600160401b03808516600090815260076020908152604080832086851684529091529020546126af911660016143e2565b6001600160401b0316816001600160401b03161461270f5760405162461bcd60e51b815260206004820152601860248201527f4f6d6e69506f7274616c3a2077726f6e67206f666673657400000000000000006044820152606401610c54565b856020015160ff16600460ff16148061273157508160ff16866020015160ff16145b61277d5760405162461bcd60e51b815260206004820152601c60248201527f4f6d6e69506f7274616c3a2077726f6e6720636f6e66206c6576656c000000006044820152606401610c54565b6040808701516001600160401b03868116600090815260086020908152848220878416835290529290922054908216911610156127f7576040868101516001600160401b038681166000908152600860209081528482208784168352905292909220805467ffffffffffffffff1916929091169190911790555b6001600160401b038085166000908152600760209081526040808320868516845290915281208054600193919291612831918591166143e2565b92506101000a8154816001600160401b0302191690836001600160401b03160217905550306001600160a01b03168560a001602081019061287291906140c6565b6001600160a01b03160361294c57806001600160401b0316826001600160401b0316856001600160401b03167f8277cab1f0fa69b34674f64a7d43f242b0bacece6f5b7e8652f1e0d88a9b873b6000336000604051602401612905906020808252601e908201527f4f6d6e69506f7274616c3a206e6f207863616c6c20746f20706f7274616c0000604082015260600190565b60408051601f198184030181529181526020820180516001600160e01b031662461bcd60e51b1790525161293c9493929190614746565b60405180910390a4505050505050565b604080518082019091526001600160401b03851681526020810161297660a08801608089016140c6565b6001600160a01b039081169091528151600b8054602090940151909216600160401b026001600160e01b03199093166001600160401b039091161791909117905560008080806129cc60c08a0160a08b016140c6565b6001600160a01b031614612a1e57612a196129ed60c08a0160a08b016140c6565b6129fe6101008b0160e08c01613d7b565b6001600160401b0316612a1460c08c018c614782565b6137da565b612a33565b612a33612a2e60c08a018a614782565b61389a565b600b80546001600160e01b031916905591945092509050600083612a575782612a68565b604051806020016040528060008152505b9050846001600160401b0316866001600160401b0316896001600160401b03167f8277cab1f0fa69b34674f64a7d43f242b0bacece6f5b7e8652f1e0d88a9b873b85338987604051611c9e9493929190614746565b600080516020614aef8339815191526000908152600080516020614b0f83398151915260208190527ffae9838a178d7f201aa98e2ce5340158edda60bb1e8f168f46503bf3e99f13be5460ff16806114325750600092835260205250604090205460ff1690565b612b2c613931565b610baf8161397a565b6001600160a01b038116612b8b5760405162461bcd60e51b815260206004820152601d60248201527f4f6d6e69506f7274616c3a206e6f207a65726f206665654f7261636c650000006044820152606401610c54565b600280546001600160a01b0319166001600160a01b0383169081179091556040519081527fd97bdb0db82b52a85aa07f8da78033b1d6e159d94f1e3cbd4109d946c3bcfd3290602001611f50565b6000816001600160401b031611612c325760405162461bcd60e51b815260206004820152601b60248201527f4f6d6e69506f7274616c3a206e6f207a65726f206d61782067617300000000006044820152606401610c54565b600080546cffffffffffffffff00000000001916650100000000006001600160401b038416908102919091179091556040519081527f1153561ac5effc2926ba6c612f86a397c997bc43dfbfc718da08065be0c5fe4d90602001611f50565b6000816001600160401b031611612cea5760405162461bcd60e51b815260206004820152601b60248201527f4f6d6e69506f7274616c3a206e6f207a65726f206d696e2067617300000000006044820152606401610c54565b6000805467ffffffffffffffff60681b1916600160681b6001600160401b038416908102919091179091556040519081527f8c852a6291aa436654b167353bca4a4b0c3d024c7562cb5082e7c869bddabf3e90602001611f50565b60008161ffff1611612d995760405162461bcd60e51b815260206004820152601c60248201527f4f6d6e69506f7274616c3a206e6f207a65726f206d61782073697a65000000006044820152606401610c54565b6000805464ffff0000001916630100000061ffff8416908102919091179091556040519081527f65923e04419dc810d0ea08a94a7f608d4c4d949818d95c3788f895e575dd206490602001611f50565b60008161ffff1611612e3d5760405162461bcd60e51b815260206004820152601c60248201527f4f6d6e69506f7274616c3a206e6f207a65726f206d61782073697a65000000006044820152606401610c54565b6000805462ffff00191661010061ffff8416908102919091179091556040519081527f620bbea084306b66a8cc6b5b63830d6b3874f9d2438914e259ffd5065c33f7b090602001611f50565b8080612ed75760405162461bcd60e51b815260206004820152601960248201527f4f6d6e69506f7274616c3a206e6f2076616c696461746f7273000000000000006044820152606401610c54565b6001600160401b038085166000908152600960205260409020541615612f3f5760405162461bcd60e51b815260206004820152601d60248201527f4f6d6e69506f7274616c3a206475706c69636174652076616c207365740000006044820152606401610c54565b604080518082018252600080825260208083018290526001600160401b0388168252600a9052918220825b8481101561311757868682818110612f8457612f846142d3565b905060400201803603810190612f9a91906147c8565b80519093506001600160a01b0316612ff45760405162461bcd60e51b815260206004820152601d60248201527f4f6d6e69506f7274616c3a206e6f207a65726f2076616c696461746f720000006044820152606401610c54565b600083602001516001600160401b0316116130515760405162461bcd60e51b815260206004820152601960248201527f4f6d6e69506f7274616c3a206e6f207a65726f20706f776572000000000000006044820152606401610c54565b82516001600160a01b03166000908152602083905260409020546001600160401b0316156130c15760405162461bcd60e51b815260206004820152601f60248201527f4f6d6e69506f7274616c3a206475706c69636174652076616c696461746f72006044820152606401610c54565b60208301516130d090856143e2565b60208481015185516001600160a01b03166000908152918590526040909120805467ffffffffffffffff19166001600160401b039092169190911790559350600101612f6a565b506001600160401b038781166000818152600960205260408120805467ffffffffffffffff191687851617905554600160a81b9004909116101561317b576000805467ffffffffffffffff60a81b1916600160a81b6001600160401b038a16021790555b6040516001600160401b038816907f3a7c2f997a87ba92aedaecd1127f4129cae1283e2809ebf5304d321b943fd10790600090a250505050505050565b611495600080516020614aef8339815191525b6000818152600080516020614b0f833981519152602081905260409091205460ff166132305760405162461bcd60e51b815260206004820152601460248201527314185d5cd8589b194e881b9bdd081c185d5cd95960621b6044820152606401610c54565b600091825260205260409020805460ff19169055565b7f9016d09d72d40fdae2fd8ceac6b6234c7706214fd39c1cd1e609a0528c19930080546001600160a01b031981166001600160a01b03848116918217845560405192169182907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a3505050565b611495600080516020614aef833981519152611fa4565b6000805b6003548110156133c057600381815481106132ef576132ef6142d3565b906000526020600020906002020191506133064690565b82546001600160401b039081169116146133405781546001600160401b03166000908152600560205260409020805460ff191690556133b8565b60005b60018301548110156133b65760006004600085600101848154811061336a5761336a6142d3565b6000918252602080832060048304015460039092166008026101000a9091046001600160401b031683528201929092526040019020805460ff1916911515919091179055600101613343565b505b6001016132d2565b50610baf60036000613ca5565b60006133dc60208401846140c6565b6001600160a01b0316613430836133f66020870187614782565b8080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061398292505050565b6001600160a01b0316149392505050565b60008160ff168360ff16856134569190614807565b6134609190614848565b6001600160401b0316856001600160401b0316119050949350505050565b60606000826001600160401b0381111561349a5761349a6141d7565b6040519080825280602002602001820160405280156134c3578160200160208202803683370190505b50905060005b838110156135435761351e60028686848181106134e8576134e86142d3565b90506020028101906134fa91906142e9565b60405160200161350a91906148b3565b6040516020818303038152906040526139ac565b828281518110613530576135306142d3565b60209081029190910101526001016134c9565b509392505050565b8051600090858461355d8160016149a9565b61356783856149a9565b1461358557604051631a8a024960e11b815260040160405180910390fd5b6000816001600160401b0381111561359f5761359f6141d7565b6040519080825280602002602001820160405280156135c8578160200160208202803683370190505b5090506000806000805b8581101561371557600088851061360d5785846135ee816149bc565b955081518110613600576136006142d3565b6020026020010151613633565b8a85613618816149bc565b96508151811061362a5761362a6142d3565b60200260200101515b905060008d8d84818110613649576136496142d3565b905060200201602081019061365e91906149d5565b61368b578f8f8561366e816149bc565b965081811061367f5761367f6142d3565b905060200201356136e2565b8986106136bc57868561369d816149bc565b9650815181106136af576136af6142d3565b60200260200101516136e2565b8b866136c7816149bc565b9750815181106136d9576136d96142d3565b60200260200101515b90506136ee82826139e3565b878481518110613700576137006142d3565b602090810291909101015250506001016135d2565b5084156137675785811461373c57604051631a8a024960e11b815260040160405180910390fd5b836001860381518110613751576137516142d3565b60200260200101519750505050505050506116d6565b86156137805788600081518110613751576137516142d3565b8c8c6000818110613793576137936142d3565b9050602002013597505050505050505095945050505050565b60006110b960018360405160200161350a91906149f7565b6000826137d18584613a12565b14949350505050565b600060606000805a905060008061385d8960008060019054906101000a900461ffff168b8b8080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f820116905080830192505050505050508e6001600160a01b0316613a4d90949392919063ffffffff16565b9150915060005a9050613871603f8b614a65565b811161387957fe5b82826138858387614638565b965096509650505050505b9450945094915050565b600060606000805a9050600080306001600160a01b031688886040516138c1929190614a79565b6000604051808303816000865af19150503d80600081146138fe576040519150601f19603f3d011682016040523d82523d6000602084013e613903565b606091505b50915091505a6139139084614638565b92508161392257805160208201fd5b909450925090505b9250925092565b7ff0c57e16840df040f15088dc2f81fe391c3923bec73e23a9662efc9c229c6a0054600160401b900460ff1661149557604051631afcd79f60e31b815260040160405180910390fd5b611e31613931565b6000806000806139928686613ad7565b9250925092506139a28282613b21565b5090949350505050565b600082826040516020016139c1929190614a89565b60408051601f1981840301815282825280516020918201209083015201611f86565b60008183106139ff576000828152602084905260409020611432565b6000838152602083905260409020611432565b600081815b845181101561354357613a4382868381518110613a3657613a366142d3565b60200260200101516139e3565b9150600101613a17565b6000606060008060008661ffff166001600160401b03811115613a7257613a726141d7565b6040519080825280601f01601f191660200182016040528015613a9c576020820181803683370190505b5090506000808751602089018b8e8ef191503d925086831115613abd578692505b828152826000602083013e90999098509650505050505050565b60008060008351604103613b115760208401516040850151606086015160001a613b0388828585613bda565b95509550955050505061392a565b505081516000915060029061392a565b6000826003811115613b3557613b35614ab8565b03613b3e575050565b6001826003811115613b5257613b52614ab8565b03613b705760405163f645eedf60e01b815260040160405180910390fd5b6002826003811115613b8457613b84614ab8565b03613ba55760405163fce698f760e01b815260048101829052602401610c54565b6003826003811115613bb957613bb9614ab8565b03610d28576040516335e2f38360e21b815260048101829052602401610c54565b600080807f7fffffffffffffffffffffffffffffff5d576e7357a4501ddfe92f46681b20a0841115613c155750600091506003905082613890565b604080516000808252602082018084528a905260ff891692820192909252606081018790526080810186905260019060a0016020604051602081039080840390855afa158015613c69573d6000803e3d6000fd5b5050604051601f1901519150506001600160a01b038116613c9557506000925060019150829050613890565b9760009750879650945050505050565b5080546000825560020290600052602060002090810190610baf91905b80821115613cf157805467ffffffffffffffff191681556000613ce86001830182613cf5565b50600201613cc2565b5090565b508054600082556003016004900490600052602060002090810190610baf91905b80821115613cf15760008155600101613d16565b803560ff81168114613d3b57600080fd5b919050565b600060208284031215613d5257600080fd5b61143282613d2a565b6001600160401b0381168114610baf57600080fd5b8035613d3b81613d5b565b600060208284031215613d8d57600080fd5b813561143281613d5b565b60008060208385031215613dab57600080fd5b82356001600160401b0380821115613dc257600080fd5b818501915085601f830112613dd657600080fd5b813581811115613de557600080fd5b8660208260051b8501011115613dfa57600080fd5b60209290920196919550909350505050565b600060208284031215613e1e57600080fd5b81356001600160401b03811115613e3457600080fd5b8201610160818503121561143257600080fd5b600060208284031215613e5957600080fd5b5035919050565b600060208284031215613e7257600080fd5b81356001600160401b03811115613e8857600080fd5b82016101a0818503121561143257600080fd5b60008060408385031215613eae57600080fd5b8235613eb981613d5b565b91506020830135613ec981613d5b565b809150509250929050565b60008060408385031215613ee757600080fd5b823591506020830135613ec981613d5b565b80356001600160a01b0381168114613d3b57600080fd5b60008060408385031215613f2357600080fd5b8235613f2e81613d5b565b9150613f3c60208401613ef9565b90509250929050565b600080600060408486031215613f5a57600080fd5b8335613f6581613d5b565b925060208401356001600160401b0380821115613f8157600080fd5b818601915086601f830112613f9557600080fd5b813581811115613fa457600080fd5b8760208260061b8501011115613fb957600080fd5b6020830194508093505050509250925092565b60008083601f840112613fde57600080fd5b5081356001600160401b03811115613ff557600080fd5b60208301915083602082850101111561400d57600080fd5b9250929050565b6000806000806060858703121561402a57600080fd5b843561403581613d5b565b935060208501356001600160401b0381111561405057600080fd5b61405c87828801613fcc565b909450925050604085013561407081613d5b565b939692955090935050565b60008060006060848603121561409057600080fd5b833561409b81613d5b565b925060208401356140ab81613d5b565b915060408401356140bb81613d5b565b809150509250925092565b6000602082840312156140d857600080fd5b61143282613ef9565b6000602082840312156140f357600080fd5b813561ffff8116811461143257600080fd5b60008060008060008060a0878903121561411e57600080fd5b863561412981613d5b565b955061413760208801613d2a565b945061414560408801613ef9565b935060608701356001600160401b0381111561416057600080fd5b61416c89828a01613fcc565b909450925050608087013561418081613d5b565b809150509295509295509295565b6000808335601e198436030181126141a557600080fd5b8301803591506001600160401b038211156141bf57600080fd5b6020019150600581901b360382131561400d57600080fd5b634e487b7160e01b600052604160045260246000fd5b604080519081016001600160401b038111828210171561420f5761420f6141d7565b60405290565b604051601f8201601f191681016001600160401b038111828210171561423d5761423d6141d7565b604052919050565b600060a0828403121561425757600080fd5b60405160a081018181106001600160401b0382111715614279576142796141d7565b604052823561428781613d5b565b815261429560208401613d2a565b602082015260408301356142a881613d5b565b604082015260608301356142bb81613d5b565b60608201526080928301359281019290925250919050565b634e487b7160e01b600052603260045260246000fd5b6000823560fe198336030181126142ff57600080fd5b9190910192915050565b6000808335601e1984360301811261432057600080fd5b8301803591506001600160401b0382111561433a57600080fd5b6020019150600681901b360382131561400d57600080fd5b81835281816020850137506000828201602090810191909152601f909101601f19169091010190565b60006001600160401b0380871683526060602084015261439f606084018688614352565b915080841660408401525095945050505050565b6000602082840312156143c557600080fd5b5051919050565b634e487b7160e01b600052601160045260246000fd5b6001600160401b03818116838216019080821115614402576144026143cc565b5092915050565b6001600160a01b0387811682528616602082015260a0604082018190526000906144369083018688614352565b6001600160401b039490941660608301525060800152949350505050565b60008235603e198336030181126142ff57600080fd5b600081356110b981613d5b565b600160401b82111561448b5761448b6141d7565b8054828255808310156110a95760008260005260206000206003850160021c81016003840160021c8201915060188660031b1680156144db576000198083018054828460200360031b1c16815550505b505b818110156113ce578281556001016144dd565b81356144fb81613d5b565b815467ffffffffffffffff19166001600160401b0391821617825560019081830160208581013536879003601e1901811261453557600080fd5b860180358481111561454657600080fd5b6020820194508060051b360385131561455e57600080fd5b6145688185614477565b60009384526020842093600282901c92505b828110156145d1576000805b60048110156145c5576145b861459b8961446a565b6001600160401b03908116600684901b90811b91901b1984161790565b9786019791508801614586565b5085820155860161457a565b50600319811680820381831461460c576000805b82811015614606576145f961459b8a61446a565b98870198915089016145e5565b50868501555b50505050505050505050565b6001600160401b03828116828216039080821115614402576144026143cc565b818103818111156110b9576110b96143cc565b60006040823603121561465d57600080fd5b6146656141ed565b61466e83613ef9565b81526020808401356001600160401b038082111561468b57600080fd5b9085019036601f83011261469e57600080fd5b8135818111156146b0576146b06141d7565b6146c2601f8201601f19168501614215565b915080825236848285010111156146d857600080fd5b80848401858401376000908201840152918301919091525092915050565b60005b838110156147115781810151838201526020016146f9565b50506000910152565b600081518084526147328160208601602086016146f6565b601f01601f19169290920160200192915050565b8481526001600160a01b038416602082015282151560408201526080606082018190526000906147789083018461471a565b9695505050505050565b6000808335601e1984360301811261479957600080fd5b8301803591506001600160401b038211156147b357600080fd5b60200191503681900382131561400d57600080fd5b6000604082840312156147da57600080fd5b6147e26141ed565b6147eb83613ef9565b815260208301356147fb81613d5b565b60208201529392505050565b6001600160401b0381811683821602808216919082811461482a5761482a6143cc565b505092915050565b634e487b7160e01b600052601260045260246000fd5b60006001600160401b038084168061486257614862614832565b92169190910492915050565b6000808335601e1984360301811261488557600080fd5b83016020810192503590506001600160401b038111156148a457600080fd5b80360382131561400d57600080fd5b60208152600082356148c481613d5b565b6001600160401b0381166020840152506148e060208401613d70565b6001600160401b0381166040840152506148fc60408401613d70565b6001600160401b03811660608401525061491860608401613d70565b6001600160401b03811660808401525061493460808401613ef9565b6001600160a01b03811660a08401525061495060a08401613ef9565b6001600160a01b03811660c08401525061496d60c084018461486e565b6101008060e086015261498561012086018385614352565b925061499360e08701613d70565b6001600160401b03169401939093529392505050565b808201808211156110b9576110b96143cc565b6000600182016149ce576149ce6143cc565b5060010190565b6000602082840312156149e757600080fd5b8135801515811461143257600080fd5b60a081018235614a0681613d5b565b6001600160401b03808216845260ff614a2160208701613d2a565b16602085015260408501359150614a3782613d5b565b9081166040840152606084013590614a4e82613d5b565b166060830152608092830135929091019190915290565b600082614a7457614a74614832565b500490565b8183823760009101908152919050565b60ff60f81b8360f81b16815260008251614aaa8160018501602087016146f6565b919091016001019392505050565b634e487b7160e01b600052602160045260246000fdfea06a0c1264badca141841b5f52470407dac9adaaa539dd445540986341b73a6876e8952e4b09b8d505aa08998d716721a1dbf0884ac74202e33985da1ed005e9ff37105740f03695c8f3597f3aff2b92fbe1c80abea3c28731ecff2efd693400feccba1cfc4544bf9cd83b76f36ae5c464750b6c43f682e26744ee21ec31fc1ea2646970667358221220677909ee34f3e258451d9b2e392d0c14cb9a04e359816dd2c050b1455c1984a164736f6c63430008180033",
}

// OmniPortalABI is the input ABI used to generate the binding from.
// Deprecated: Use OmniPortalMetaData.ABI instead.
var OmniPortalABI = OmniPortalMetaData.ABI

// OmniPortalBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use OmniPortalMetaData.Bin instead.
var OmniPortalBin = OmniPortalMetaData.Bin

// DeployOmniPortal deploys a new Ethereum contract, binding an instance of OmniPortal to it.
func DeployOmniPortal(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *OmniPortal, error) {
	parsed, err := OmniPortalMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(OmniPortalBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &OmniPortal{OmniPortalCaller: OmniPortalCaller{contract: contract}, OmniPortalTransactor: OmniPortalTransactor{contract: contract}, OmniPortalFilterer: OmniPortalFilterer{contract: contract}}, nil
}

// OmniPortal is an auto generated Go binding around an Ethereum contract.
type OmniPortal struct {
	OmniPortalCaller     // Read-only binding to the contract
	OmniPortalTransactor // Write-only binding to the contract
	OmniPortalFilterer   // Log filterer for contract events
}

// OmniPortalCaller is an auto generated read-only Go binding around an Ethereum contract.
type OmniPortalCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OmniPortalTransactor is an auto generated write-only Go binding around an Ethereum contract.
type OmniPortalTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OmniPortalFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type OmniPortalFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// OmniPortalSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type OmniPortalSession struct {
	Contract     *OmniPortal       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// OmniPortalCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type OmniPortalCallerSession struct {
	Contract *OmniPortalCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// OmniPortalTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type OmniPortalTransactorSession struct {
	Contract     *OmniPortalTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// OmniPortalRaw is an auto generated low-level Go binding around an Ethereum contract.
type OmniPortalRaw struct {
	Contract *OmniPortal // Generic contract binding to access the raw methods on
}

// OmniPortalCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type OmniPortalCallerRaw struct {
	Contract *OmniPortalCaller // Generic read-only contract binding to access the raw methods on
}

// OmniPortalTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type OmniPortalTransactorRaw struct {
	Contract *OmniPortalTransactor // Generic write-only contract binding to access the raw methods on
}

// NewOmniPortal creates a new instance of OmniPortal, bound to a specific deployed contract.
func NewOmniPortal(address common.Address, backend bind.ContractBackend) (*OmniPortal, error) {
	contract, err := bindOmniPortal(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &OmniPortal{OmniPortalCaller: OmniPortalCaller{contract: contract}, OmniPortalTransactor: OmniPortalTransactor{contract: contract}, OmniPortalFilterer: OmniPortalFilterer{contract: contract}}, nil
}

// NewOmniPortalCaller creates a new read-only instance of OmniPortal, bound to a specific deployed contract.
func NewOmniPortalCaller(address common.Address, caller bind.ContractCaller) (*OmniPortalCaller, error) {
	contract, err := bindOmniPortal(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &OmniPortalCaller{contract: contract}, nil
}

// NewOmniPortalTransactor creates a new write-only instance of OmniPortal, bound to a specific deployed contract.
func NewOmniPortalTransactor(address common.Address, transactor bind.ContractTransactor) (*OmniPortalTransactor, error) {
	contract, err := bindOmniPortal(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &OmniPortalTransactor{contract: contract}, nil
}

// NewOmniPortalFilterer creates a new log filterer instance of OmniPortal, bound to a specific deployed contract.
func NewOmniPortalFilterer(address common.Address, filterer bind.ContractFilterer) (*OmniPortalFilterer, error) {
	contract, err := bindOmniPortal(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &OmniPortalFilterer{contract: contract}, nil
}

// bindOmniPortal binds a generic wrapper to an already deployed contract.
func bindOmniPortal(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := OmniPortalMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OmniPortal *OmniPortalRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OmniPortal.Contract.OmniPortalCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OmniPortal *OmniPortalRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OmniPortal.Contract.OmniPortalTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OmniPortal *OmniPortalRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OmniPortal.Contract.OmniPortalTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_OmniPortal *OmniPortalCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _OmniPortal.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_OmniPortal *OmniPortalTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OmniPortal.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_OmniPortal *OmniPortalTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _OmniPortal.Contract.contract.Transact(opts, method, params...)
}

// ActionXCall is a free data retrieval call binding the contract method 0xb2b2f5bd.
//
// Solidity: function ActionXCall() view returns(bytes32)
func (_OmniPortal *OmniPortalCaller) ActionXCall(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _OmniPortal.contract.Call(opts, &out, "ActionXCall")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ActionXCall is a free data retrieval call binding the contract method 0xb2b2f5bd.
//
// Solidity: function ActionXCall() view returns(bytes32)
func (_OmniPortal *OmniPortalSession) ActionXCall() ([32]byte, error) {
	return _OmniPortal.Contract.ActionXCall(&_OmniPortal.CallOpts)
}

// ActionXCall is a free data retrieval call binding the contract method 0xb2b2f5bd.
//
// Solidity: function ActionXCall() view returns(bytes32)
func (_OmniPortal *OmniPortalCallerSession) ActionXCall() ([32]byte, error) {
	return _OmniPortal.Contract.ActionXCall(&_OmniPortal.CallOpts)
}

// ActionXSubmit is a free data retrieval call binding the contract method 0xa32eb7c6.
//
// Solidity: function ActionXSubmit() view returns(bytes32)
func (_OmniPortal *OmniPortalCaller) ActionXSubmit(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _OmniPortal.contract.Call(opts, &out, "ActionXSubmit")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ActionXSubmit is a free data retrieval call binding the contract method 0xa32eb7c6.
//
// Solidity: function ActionXSubmit() view returns(bytes32)
func (_OmniPortal *OmniPortalSession) ActionXSubmit() ([32]byte, error) {
	return _OmniPortal.Contract.ActionXSubmit(&_OmniPortal.CallOpts)
}

// ActionXSubmit is a free data retrieval call binding the contract method 0xa32eb7c6.
//
// Solidity: function ActionXSubmit() view returns(bytes32)
func (_OmniPortal *OmniPortalCallerSession) ActionXSubmit() ([32]byte, error) {
	return _OmniPortal.Contract.ActionXSubmit(&_OmniPortal.CallOpts)
}

// KeyPauseAll is a free data retrieval call binding the contract method 0xa10ac97a.
//
// Solidity: function KeyPauseAll() view returns(bytes32)
func (_OmniPortal *OmniPortalCaller) KeyPauseAll(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _OmniPortal.contract.Call(opts, &out, "KeyPauseAll")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// KeyPauseAll is a free data retrieval call binding the contract method 0xa10ac97a.
//
// Solidity: function KeyPauseAll() view returns(bytes32)
func (_OmniPortal *OmniPortalSession) KeyPauseAll() ([32]byte, error) {
	return _OmniPortal.Contract.KeyPauseAll(&_OmniPortal.CallOpts)
}

// KeyPauseAll is a free data retrieval call binding the contract method 0xa10ac97a.
//
// Solidity: function KeyPauseAll() view returns(bytes32)
func (_OmniPortal *OmniPortalCallerSession) KeyPauseAll() ([32]byte, error) {
	return _OmniPortal.Contract.KeyPauseAll(&_OmniPortal.CallOpts)
}

// XSubQuorumDenominator is a free data retrieval call binding the contract method 0x06c3dc5f.
//
// Solidity: function XSubQuorumDenominator() view returns(uint8)
func (_OmniPortal *OmniPortalCaller) XSubQuorumDenominator(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _OmniPortal.contract.Call(opts, &out, "XSubQuorumDenominator")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// XSubQuorumDenominator is a free data retrieval call binding the contract method 0x06c3dc5f.
//
// Solidity: function XSubQuorumDenominator() view returns(uint8)
func (_OmniPortal *OmniPortalSession) XSubQuorumDenominator() (uint8, error) {
	return _OmniPortal.Contract.XSubQuorumDenominator(&_OmniPortal.CallOpts)
}

// XSubQuorumDenominator is a free data retrieval call binding the contract method 0x06c3dc5f.
//
// Solidity: function XSubQuorumDenominator() view returns(uint8)
func (_OmniPortal *OmniPortalCallerSession) XSubQuorumDenominator() (uint8, error) {
	return _OmniPortal.Contract.XSubQuorumDenominator(&_OmniPortal.CallOpts)
}

// XSubQuorumNumerator is a free data retrieval call binding the contract method 0x0360d20f.
//
// Solidity: function XSubQuorumNumerator() view returns(uint8)
func (_OmniPortal *OmniPortalCaller) XSubQuorumNumerator(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _OmniPortal.contract.Call(opts, &out, "XSubQuorumNumerator")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// XSubQuorumNumerator is a free data retrieval call binding the contract method 0x0360d20f.
//
// Solidity: function XSubQuorumNumerator() view returns(uint8)
func (_OmniPortal *OmniPortalSession) XSubQuorumNumerator() (uint8, error) {
	return _OmniPortal.Contract.XSubQuorumNumerator(&_OmniPortal.CallOpts)
}

// XSubQuorumNumerator is a free data retrieval call binding the contract method 0x0360d20f.
//
// Solidity: function XSubQuorumNumerator() view returns(uint8)
func (_OmniPortal *OmniPortalCallerSession) XSubQuorumNumerator() (uint8, error) {
	return _OmniPortal.Contract.XSubQuorumNumerator(&_OmniPortal.CallOpts)
}

// ChainId is a free data retrieval call binding the contract method 0x9a8a0592.
//
// Solidity: function chainId() view returns(uint64)
func (_OmniPortal *OmniPortalCaller) ChainId(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _OmniPortal.contract.Call(opts, &out, "chainId")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// ChainId is a free data retrieval call binding the contract method 0x9a8a0592.
//
// Solidity: function chainId() view returns(uint64)
func (_OmniPortal *OmniPortalSession) ChainId() (uint64, error) {
	return _OmniPortal.Contract.ChainId(&_OmniPortal.CallOpts)
}

// ChainId is a free data retrieval call binding the contract method 0x9a8a0592.
//
// Solidity: function chainId() view returns(uint64)
func (_OmniPortal *OmniPortalCallerSession) ChainId() (uint64, error) {
	return _OmniPortal.Contract.ChainId(&_OmniPortal.CallOpts)
}

// FeeFor is a free data retrieval call binding the contract method 0x8dd9523c.
//
// Solidity: function feeFor(uint64 destChainId, bytes data, uint64 gasLimit) view returns(uint256)
func (_OmniPortal *OmniPortalCaller) FeeFor(opts *bind.CallOpts, destChainId uint64, data []byte, gasLimit uint64) (*big.Int, error) {
	var out []interface{}
	err := _OmniPortal.contract.Call(opts, &out, "feeFor", destChainId, data, gasLimit)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// FeeFor is a free data retrieval call binding the contract method 0x8dd9523c.
//
// Solidity: function feeFor(uint64 destChainId, bytes data, uint64 gasLimit) view returns(uint256)
func (_OmniPortal *OmniPortalSession) FeeFor(destChainId uint64, data []byte, gasLimit uint64) (*big.Int, error) {
	return _OmniPortal.Contract.FeeFor(&_OmniPortal.CallOpts, destChainId, data, gasLimit)
}

// FeeFor is a free data retrieval call binding the contract method 0x8dd9523c.
//
// Solidity: function feeFor(uint64 destChainId, bytes data, uint64 gasLimit) view returns(uint256)
func (_OmniPortal *OmniPortalCallerSession) FeeFor(destChainId uint64, data []byte, gasLimit uint64) (*big.Int, error) {
	return _OmniPortal.Contract.FeeFor(&_OmniPortal.CallOpts, destChainId, data, gasLimit)
}

// FeeOracle is a free data retrieval call binding the contract method 0x500b19e7.
//
// Solidity: function feeOracle() view returns(address)
func (_OmniPortal *OmniPortalCaller) FeeOracle(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OmniPortal.contract.Call(opts, &out, "feeOracle")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// FeeOracle is a free data retrieval call binding the contract method 0x500b19e7.
//
// Solidity: function feeOracle() view returns(address)
func (_OmniPortal *OmniPortalSession) FeeOracle() (common.Address, error) {
	return _OmniPortal.Contract.FeeOracle(&_OmniPortal.CallOpts)
}

// FeeOracle is a free data retrieval call binding the contract method 0x500b19e7.
//
// Solidity: function feeOracle() view returns(address)
func (_OmniPortal *OmniPortalCallerSession) FeeOracle() (common.Address, error) {
	return _OmniPortal.Contract.FeeOracle(&_OmniPortal.CallOpts)
}

// InXBlockOffset is a free data retrieval call binding the contract method 0x3fd3b15e.
//
// Solidity: function inXBlockOffset(uint64 , uint64 ) view returns(uint64)
func (_OmniPortal *OmniPortalCaller) InXBlockOffset(opts *bind.CallOpts, arg0 uint64, arg1 uint64) (uint64, error) {
	var out []interface{}
	err := _OmniPortal.contract.Call(opts, &out, "inXBlockOffset", arg0, arg1)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// InXBlockOffset is a free data retrieval call binding the contract method 0x3fd3b15e.
//
// Solidity: function inXBlockOffset(uint64 , uint64 ) view returns(uint64)
func (_OmniPortal *OmniPortalSession) InXBlockOffset(arg0 uint64, arg1 uint64) (uint64, error) {
	return _OmniPortal.Contract.InXBlockOffset(&_OmniPortal.CallOpts, arg0, arg1)
}

// InXBlockOffset is a free data retrieval call binding the contract method 0x3fd3b15e.
//
// Solidity: function inXBlockOffset(uint64 , uint64 ) view returns(uint64)
func (_OmniPortal *OmniPortalCallerSession) InXBlockOffset(arg0 uint64, arg1 uint64) (uint64, error) {
	return _OmniPortal.Contract.InXBlockOffset(&_OmniPortal.CallOpts, arg0, arg1)
}

// InXMsgOffset is a free data retrieval call binding the contract method 0xd051c97d.
//
// Solidity: function inXMsgOffset(uint64 , uint64 ) view returns(uint64)
func (_OmniPortal *OmniPortalCaller) InXMsgOffset(opts *bind.CallOpts, arg0 uint64, arg1 uint64) (uint64, error) {
	var out []interface{}
	err := _OmniPortal.contract.Call(opts, &out, "inXMsgOffset", arg0, arg1)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// InXMsgOffset is a free data retrieval call binding the contract method 0xd051c97d.
//
// Solidity: function inXMsgOffset(uint64 , uint64 ) view returns(uint64)
func (_OmniPortal *OmniPortalSession) InXMsgOffset(arg0 uint64, arg1 uint64) (uint64, error) {
	return _OmniPortal.Contract.InXMsgOffset(&_OmniPortal.CallOpts, arg0, arg1)
}

// InXMsgOffset is a free data retrieval call binding the contract method 0xd051c97d.
//
// Solidity: function inXMsgOffset(uint64 , uint64 ) view returns(uint64)
func (_OmniPortal *OmniPortalCallerSession) InXMsgOffset(arg0 uint64, arg1 uint64) (uint64, error) {
	return _OmniPortal.Contract.InXMsgOffset(&_OmniPortal.CallOpts, arg0, arg1)
}

// IsPaused is a free data retrieval call binding the contract method 0x241b71bb.
//
// Solidity: function isPaused(bytes32 actionId) view returns(bool)
func (_OmniPortal *OmniPortalCaller) IsPaused(opts *bind.CallOpts, actionId [32]byte) (bool, error) {
	var out []interface{}
	err := _OmniPortal.contract.Call(opts, &out, "isPaused", actionId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPaused is a free data retrieval call binding the contract method 0x241b71bb.
//
// Solidity: function isPaused(bytes32 actionId) view returns(bool)
func (_OmniPortal *OmniPortalSession) IsPaused(actionId [32]byte) (bool, error) {
	return _OmniPortal.Contract.IsPaused(&_OmniPortal.CallOpts, actionId)
}

// IsPaused is a free data retrieval call binding the contract method 0x241b71bb.
//
// Solidity: function isPaused(bytes32 actionId) view returns(bool)
func (_OmniPortal *OmniPortalCallerSession) IsPaused(actionId [32]byte) (bool, error) {
	return _OmniPortal.Contract.IsPaused(&_OmniPortal.CallOpts, actionId)
}

// IsPaused0 is a free data retrieval call binding the contract method 0x461ab488.
//
// Solidity: function isPaused(bytes32 actionId, uint64 chainId_) view returns(bool)
func (_OmniPortal *OmniPortalCaller) IsPaused0(opts *bind.CallOpts, actionId [32]byte, chainId_ uint64) (bool, error) {
	var out []interface{}
	err := _OmniPortal.contract.Call(opts, &out, "isPaused0", actionId, chainId_)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPaused0 is a free data retrieval call binding the contract method 0x461ab488.
//
// Solidity: function isPaused(bytes32 actionId, uint64 chainId_) view returns(bool)
func (_OmniPortal *OmniPortalSession) IsPaused0(actionId [32]byte, chainId_ uint64) (bool, error) {
	return _OmniPortal.Contract.IsPaused0(&_OmniPortal.CallOpts, actionId, chainId_)
}

// IsPaused0 is a free data retrieval call binding the contract method 0x461ab488.
//
// Solidity: function isPaused(bytes32 actionId, uint64 chainId_) view returns(bool)
func (_OmniPortal *OmniPortalCallerSession) IsPaused0(actionId [32]byte, chainId_ uint64) (bool, error) {
	return _OmniPortal.Contract.IsPaused0(&_OmniPortal.CallOpts, actionId, chainId_)
}

// IsPaused1 is a free data retrieval call binding the contract method 0xb187bd26.
//
// Solidity: function isPaused() view returns(bool)
func (_OmniPortal *OmniPortalCaller) IsPaused1(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _OmniPortal.contract.Call(opts, &out, "isPaused1")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPaused1 is a free data retrieval call binding the contract method 0xb187bd26.
//
// Solidity: function isPaused() view returns(bool)
func (_OmniPortal *OmniPortalSession) IsPaused1() (bool, error) {
	return _OmniPortal.Contract.IsPaused1(&_OmniPortal.CallOpts)
}

// IsPaused1 is a free data retrieval call binding the contract method 0xb187bd26.
//
// Solidity: function isPaused() view returns(bool)
func (_OmniPortal *OmniPortalCallerSession) IsPaused1() (bool, error) {
	return _OmniPortal.Contract.IsPaused1(&_OmniPortal.CallOpts)
}

// IsSupportedDest is a free data retrieval call binding the contract method 0x24278bbe.
//
// Solidity: function isSupportedDest(uint64 ) view returns(bool)
func (_OmniPortal *OmniPortalCaller) IsSupportedDest(opts *bind.CallOpts, arg0 uint64) (bool, error) {
	var out []interface{}
	err := _OmniPortal.contract.Call(opts, &out, "isSupportedDest", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsSupportedDest is a free data retrieval call binding the contract method 0x24278bbe.
//
// Solidity: function isSupportedDest(uint64 ) view returns(bool)
func (_OmniPortal *OmniPortalSession) IsSupportedDest(arg0 uint64) (bool, error) {
	return _OmniPortal.Contract.IsSupportedDest(&_OmniPortal.CallOpts, arg0)
}

// IsSupportedDest is a free data retrieval call binding the contract method 0x24278bbe.
//
// Solidity: function isSupportedDest(uint64 ) view returns(bool)
func (_OmniPortal *OmniPortalCallerSession) IsSupportedDest(arg0 uint64) (bool, error) {
	return _OmniPortal.Contract.IsSupportedDest(&_OmniPortal.CallOpts, arg0)
}

// IsSupportedShard is a free data retrieval call binding the contract method 0xaaf1bc97.
//
// Solidity: function isSupportedShard(uint64 ) view returns(bool)
func (_OmniPortal *OmniPortalCaller) IsSupportedShard(opts *bind.CallOpts, arg0 uint64) (bool, error) {
	var out []interface{}
	err := _OmniPortal.contract.Call(opts, &out, "isSupportedShard", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsSupportedShard is a free data retrieval call binding the contract method 0xaaf1bc97.
//
// Solidity: function isSupportedShard(uint64 ) view returns(bool)
func (_OmniPortal *OmniPortalSession) IsSupportedShard(arg0 uint64) (bool, error) {
	return _OmniPortal.Contract.IsSupportedShard(&_OmniPortal.CallOpts, arg0)
}

// IsSupportedShard is a free data retrieval call binding the contract method 0xaaf1bc97.
//
// Solidity: function isSupportedShard(uint64 ) view returns(bool)
func (_OmniPortal *OmniPortalCallerSession) IsSupportedShard(arg0 uint64) (bool, error) {
	return _OmniPortal.Contract.IsSupportedShard(&_OmniPortal.CallOpts, arg0)
}

// IsXCall is a free data retrieval call binding the contract method 0x55e2448e.
//
// Solidity: function isXCall() view returns(bool)
func (_OmniPortal *OmniPortalCaller) IsXCall(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _OmniPortal.contract.Call(opts, &out, "isXCall")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsXCall is a free data retrieval call binding the contract method 0x55e2448e.
//
// Solidity: function isXCall() view returns(bool)
func (_OmniPortal *OmniPortalSession) IsXCall() (bool, error) {
	return _OmniPortal.Contract.IsXCall(&_OmniPortal.CallOpts)
}

// IsXCall is a free data retrieval call binding the contract method 0x55e2448e.
//
// Solidity: function isXCall() view returns(bool)
func (_OmniPortal *OmniPortalCallerSession) IsXCall() (bool, error) {
	return _OmniPortal.Contract.IsXCall(&_OmniPortal.CallOpts)
}

// LatestValSetId is a free data retrieval call binding the contract method 0xf45cc7b8.
//
// Solidity: function latestValSetId() view returns(uint64)
func (_OmniPortal *OmniPortalCaller) LatestValSetId(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _OmniPortal.contract.Call(opts, &out, "latestValSetId")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// LatestValSetId is a free data retrieval call binding the contract method 0xf45cc7b8.
//
// Solidity: function latestValSetId() view returns(uint64)
func (_OmniPortal *OmniPortalSession) LatestValSetId() (uint64, error) {
	return _OmniPortal.Contract.LatestValSetId(&_OmniPortal.CallOpts)
}

// LatestValSetId is a free data retrieval call binding the contract method 0xf45cc7b8.
//
// Solidity: function latestValSetId() view returns(uint64)
func (_OmniPortal *OmniPortalCallerSession) LatestValSetId() (uint64, error) {
	return _OmniPortal.Contract.LatestValSetId(&_OmniPortal.CallOpts)
}

// Network is a free data retrieval call binding the contract method 0x74eba939.
//
// Solidity: function network(uint256 ) view returns(uint64 chainId)
func (_OmniPortal *OmniPortalCaller) Network(opts *bind.CallOpts, arg0 *big.Int) (uint64, error) {
	var out []interface{}
	err := _OmniPortal.contract.Call(opts, &out, "network", arg0)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// Network is a free data retrieval call binding the contract method 0x74eba939.
//
// Solidity: function network(uint256 ) view returns(uint64 chainId)
func (_OmniPortal *OmniPortalSession) Network(arg0 *big.Int) (uint64, error) {
	return _OmniPortal.Contract.Network(&_OmniPortal.CallOpts, arg0)
}

// Network is a free data retrieval call binding the contract method 0x74eba939.
//
// Solidity: function network(uint256 ) view returns(uint64 chainId)
func (_OmniPortal *OmniPortalCallerSession) Network(arg0 *big.Int) (uint64, error) {
	return _OmniPortal.Contract.Network(&_OmniPortal.CallOpts, arg0)
}

// OmniCChainId is a free data retrieval call binding the contract method 0x36d21912.
//
// Solidity: function omniCChainId() view returns(uint64)
func (_OmniPortal *OmniPortalCaller) OmniCChainId(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _OmniPortal.contract.Call(opts, &out, "omniCChainId")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// OmniCChainId is a free data retrieval call binding the contract method 0x36d21912.
//
// Solidity: function omniCChainId() view returns(uint64)
func (_OmniPortal *OmniPortalSession) OmniCChainId() (uint64, error) {
	return _OmniPortal.Contract.OmniCChainId(&_OmniPortal.CallOpts)
}

// OmniCChainId is a free data retrieval call binding the contract method 0x36d21912.
//
// Solidity: function omniCChainId() view returns(uint64)
func (_OmniPortal *OmniPortalCallerSession) OmniCChainId() (uint64, error) {
	return _OmniPortal.Contract.OmniCChainId(&_OmniPortal.CallOpts)
}

// OmniChainId is a free data retrieval call binding the contract method 0x110ff5f1.
//
// Solidity: function omniChainId() view returns(uint64)
func (_OmniPortal *OmniPortalCaller) OmniChainId(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _OmniPortal.contract.Call(opts, &out, "omniChainId")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// OmniChainId is a free data retrieval call binding the contract method 0x110ff5f1.
//
// Solidity: function omniChainId() view returns(uint64)
func (_OmniPortal *OmniPortalSession) OmniChainId() (uint64, error) {
	return _OmniPortal.Contract.OmniChainId(&_OmniPortal.CallOpts)
}

// OmniChainId is a free data retrieval call binding the contract method 0x110ff5f1.
//
// Solidity: function omniChainId() view returns(uint64)
func (_OmniPortal *OmniPortalCallerSession) OmniChainId() (uint64, error) {
	return _OmniPortal.Contract.OmniChainId(&_OmniPortal.CallOpts)
}

// OutXMsgOffset is a free data retrieval call binding the contract method 0x3aa87330.
//
// Solidity: function outXMsgOffset(uint64 , uint64 ) view returns(uint64)
func (_OmniPortal *OmniPortalCaller) OutXMsgOffset(opts *bind.CallOpts, arg0 uint64, arg1 uint64) (uint64, error) {
	var out []interface{}
	err := _OmniPortal.contract.Call(opts, &out, "outXMsgOffset", arg0, arg1)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// OutXMsgOffset is a free data retrieval call binding the contract method 0x3aa87330.
//
// Solidity: function outXMsgOffset(uint64 , uint64 ) view returns(uint64)
func (_OmniPortal *OmniPortalSession) OutXMsgOffset(arg0 uint64, arg1 uint64) (uint64, error) {
	return _OmniPortal.Contract.OutXMsgOffset(&_OmniPortal.CallOpts, arg0, arg1)
}

// OutXMsgOffset is a free data retrieval call binding the contract method 0x3aa87330.
//
// Solidity: function outXMsgOffset(uint64 , uint64 ) view returns(uint64)
func (_OmniPortal *OmniPortalCallerSession) OutXMsgOffset(arg0 uint64, arg1 uint64) (uint64, error) {
	return _OmniPortal.Contract.OutXMsgOffset(&_OmniPortal.CallOpts, arg0, arg1)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_OmniPortal *OmniPortalCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _OmniPortal.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_OmniPortal *OmniPortalSession) Owner() (common.Address, error) {
	return _OmniPortal.Contract.Owner(&_OmniPortal.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_OmniPortal *OmniPortalCallerSession) Owner() (common.Address, error) {
	return _OmniPortal.Contract.Owner(&_OmniPortal.CallOpts)
}

// ValSet is a free data retrieval call binding the contract method 0x57542050.
//
// Solidity: function valSet(uint64 , address ) view returns(uint64)
func (_OmniPortal *OmniPortalCaller) ValSet(opts *bind.CallOpts, arg0 uint64, arg1 common.Address) (uint64, error) {
	var out []interface{}
	err := _OmniPortal.contract.Call(opts, &out, "valSet", arg0, arg1)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// ValSet is a free data retrieval call binding the contract method 0x57542050.
//
// Solidity: function valSet(uint64 , address ) view returns(uint64)
func (_OmniPortal *OmniPortalSession) ValSet(arg0 uint64, arg1 common.Address) (uint64, error) {
	return _OmniPortal.Contract.ValSet(&_OmniPortal.CallOpts, arg0, arg1)
}

// ValSet is a free data retrieval call binding the contract method 0x57542050.
//
// Solidity: function valSet(uint64 , address ) view returns(uint64)
func (_OmniPortal *OmniPortalCallerSession) ValSet(arg0 uint64, arg1 common.Address) (uint64, error) {
	return _OmniPortal.Contract.ValSet(&_OmniPortal.CallOpts, arg0, arg1)
}

// ValSetTotalPower is a free data retrieval call binding the contract method 0xafe8af9c.
//
// Solidity: function valSetTotalPower(uint64 ) view returns(uint64)
func (_OmniPortal *OmniPortalCaller) ValSetTotalPower(opts *bind.CallOpts, arg0 uint64) (uint64, error) {
	var out []interface{}
	err := _OmniPortal.contract.Call(opts, &out, "valSetTotalPower", arg0)

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// ValSetTotalPower is a free data retrieval call binding the contract method 0xafe8af9c.
//
// Solidity: function valSetTotalPower(uint64 ) view returns(uint64)
func (_OmniPortal *OmniPortalSession) ValSetTotalPower(arg0 uint64) (uint64, error) {
	return _OmniPortal.Contract.ValSetTotalPower(&_OmniPortal.CallOpts, arg0)
}

// ValSetTotalPower is a free data retrieval call binding the contract method 0xafe8af9c.
//
// Solidity: function valSetTotalPower(uint64 ) view returns(uint64)
func (_OmniPortal *OmniPortalCallerSession) ValSetTotalPower(arg0 uint64) (uint64, error) {
	return _OmniPortal.Contract.ValSetTotalPower(&_OmniPortal.CallOpts, arg0)
}

// Xmsg is a free data retrieval call binding the contract method 0x2f32700e.
//
// Solidity: function xmsg() view returns((uint64,address))
func (_OmniPortal *OmniPortalCaller) Xmsg(opts *bind.CallOpts) (XTypesMsgShort, error) {
	var out []interface{}
	err := _OmniPortal.contract.Call(opts, &out, "xmsg")

	if err != nil {
		return *new(XTypesMsgShort), err
	}

	out0 := *abi.ConvertType(out[0], new(XTypesMsgShort)).(*XTypesMsgShort)

	return out0, err

}

// Xmsg is a free data retrieval call binding the contract method 0x2f32700e.
//
// Solidity: function xmsg() view returns((uint64,address))
func (_OmniPortal *OmniPortalSession) Xmsg() (XTypesMsgShort, error) {
	return _OmniPortal.Contract.Xmsg(&_OmniPortal.CallOpts)
}

// Xmsg is a free data retrieval call binding the contract method 0x2f32700e.
//
// Solidity: function xmsg() view returns((uint64,address))
func (_OmniPortal *OmniPortalCallerSession) Xmsg() (XTypesMsgShort, error) {
	return _OmniPortal.Contract.Xmsg(&_OmniPortal.CallOpts)
}

// XmsgMaxDataSize is a free data retrieval call binding the contract method 0xb4d5afd1.
//
// Solidity: function xmsgMaxDataSize() view returns(uint16)
func (_OmniPortal *OmniPortalCaller) XmsgMaxDataSize(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _OmniPortal.contract.Call(opts, &out, "xmsgMaxDataSize")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// XmsgMaxDataSize is a free data retrieval call binding the contract method 0xb4d5afd1.
//
// Solidity: function xmsgMaxDataSize() view returns(uint16)
func (_OmniPortal *OmniPortalSession) XmsgMaxDataSize() (uint16, error) {
	return _OmniPortal.Contract.XmsgMaxDataSize(&_OmniPortal.CallOpts)
}

// XmsgMaxDataSize is a free data retrieval call binding the contract method 0xb4d5afd1.
//
// Solidity: function xmsgMaxDataSize() view returns(uint16)
func (_OmniPortal *OmniPortalCallerSession) XmsgMaxDataSize() (uint16, error) {
	return _OmniPortal.Contract.XmsgMaxDataSize(&_OmniPortal.CallOpts)
}

// XmsgMaxGasLimit is a free data retrieval call binding the contract method 0xcf84c818.
//
// Solidity: function xmsgMaxGasLimit() view returns(uint64)
func (_OmniPortal *OmniPortalCaller) XmsgMaxGasLimit(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _OmniPortal.contract.Call(opts, &out, "xmsgMaxGasLimit")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// XmsgMaxGasLimit is a free data retrieval call binding the contract method 0xcf84c818.
//
// Solidity: function xmsgMaxGasLimit() view returns(uint64)
func (_OmniPortal *OmniPortalSession) XmsgMaxGasLimit() (uint64, error) {
	return _OmniPortal.Contract.XmsgMaxGasLimit(&_OmniPortal.CallOpts)
}

// XmsgMaxGasLimit is a free data retrieval call binding the contract method 0xcf84c818.
//
// Solidity: function xmsgMaxGasLimit() view returns(uint64)
func (_OmniPortal *OmniPortalCallerSession) XmsgMaxGasLimit() (uint64, error) {
	return _OmniPortal.Contract.XmsgMaxGasLimit(&_OmniPortal.CallOpts)
}

// XmsgMinGasLimit is a free data retrieval call binding the contract method 0x78fe5307.
//
// Solidity: function xmsgMinGasLimit() view returns(uint64)
func (_OmniPortal *OmniPortalCaller) XmsgMinGasLimit(opts *bind.CallOpts) (uint64, error) {
	var out []interface{}
	err := _OmniPortal.contract.Call(opts, &out, "xmsgMinGasLimit")

	if err != nil {
		return *new(uint64), err
	}

	out0 := *abi.ConvertType(out[0], new(uint64)).(*uint64)

	return out0, err

}

// XmsgMinGasLimit is a free data retrieval call binding the contract method 0x78fe5307.
//
// Solidity: function xmsgMinGasLimit() view returns(uint64)
func (_OmniPortal *OmniPortalSession) XmsgMinGasLimit() (uint64, error) {
	return _OmniPortal.Contract.XmsgMinGasLimit(&_OmniPortal.CallOpts)
}

// XmsgMinGasLimit is a free data retrieval call binding the contract method 0x78fe5307.
//
// Solidity: function xmsgMinGasLimit() view returns(uint64)
func (_OmniPortal *OmniPortalCallerSession) XmsgMinGasLimit() (uint64, error) {
	return _OmniPortal.Contract.XmsgMinGasLimit(&_OmniPortal.CallOpts)
}

// XreceiptMaxErrorSize is a free data retrieval call binding the contract method 0xc26dfc05.
//
// Solidity: function xreceiptMaxErrorSize() view returns(uint16)
func (_OmniPortal *OmniPortalCaller) XreceiptMaxErrorSize(opts *bind.CallOpts) (uint16, error) {
	var out []interface{}
	err := _OmniPortal.contract.Call(opts, &out, "xreceiptMaxErrorSize")

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// XreceiptMaxErrorSize is a free data retrieval call binding the contract method 0xc26dfc05.
//
// Solidity: function xreceiptMaxErrorSize() view returns(uint16)
func (_OmniPortal *OmniPortalSession) XreceiptMaxErrorSize() (uint16, error) {
	return _OmniPortal.Contract.XreceiptMaxErrorSize(&_OmniPortal.CallOpts)
}

// XreceiptMaxErrorSize is a free data retrieval call binding the contract method 0xc26dfc05.
//
// Solidity: function xreceiptMaxErrorSize() view returns(uint16)
func (_OmniPortal *OmniPortalCallerSession) XreceiptMaxErrorSize() (uint16, error) {
	return _OmniPortal.Contract.XreceiptMaxErrorSize(&_OmniPortal.CallOpts)
}

// XsubValsetCutoff is a free data retrieval call binding the contract method 0x49cc3bf6.
//
// Solidity: function xsubValsetCutoff() view returns(uint8)
func (_OmniPortal *OmniPortalCaller) XsubValsetCutoff(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _OmniPortal.contract.Call(opts, &out, "xsubValsetCutoff")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// XsubValsetCutoff is a free data retrieval call binding the contract method 0x49cc3bf6.
//
// Solidity: function xsubValsetCutoff() view returns(uint8)
func (_OmniPortal *OmniPortalSession) XsubValsetCutoff() (uint8, error) {
	return _OmniPortal.Contract.XsubValsetCutoff(&_OmniPortal.CallOpts)
}

// XsubValsetCutoff is a free data retrieval call binding the contract method 0x49cc3bf6.
//
// Solidity: function xsubValsetCutoff() view returns(uint8)
func (_OmniPortal *OmniPortalCallerSession) XsubValsetCutoff() (uint8, error) {
	return _OmniPortal.Contract.XsubValsetCutoff(&_OmniPortal.CallOpts)
}

// AddValidatorSet is a paid mutator transaction binding the contract method 0x8532eb9f.
//
// Solidity: function addValidatorSet(uint64 valSetId, (address,uint64)[] validators) returns()
func (_OmniPortal *OmniPortalTransactor) AddValidatorSet(opts *bind.TransactOpts, valSetId uint64, validators []XTypesValidator) (*types.Transaction, error) {
	return _OmniPortal.contract.Transact(opts, "addValidatorSet", valSetId, validators)
}

// AddValidatorSet is a paid mutator transaction binding the contract method 0x8532eb9f.
//
// Solidity: function addValidatorSet(uint64 valSetId, (address,uint64)[] validators) returns()
func (_OmniPortal *OmniPortalSession) AddValidatorSet(valSetId uint64, validators []XTypesValidator) (*types.Transaction, error) {
	return _OmniPortal.Contract.AddValidatorSet(&_OmniPortal.TransactOpts, valSetId, validators)
}

// AddValidatorSet is a paid mutator transaction binding the contract method 0x8532eb9f.
//
// Solidity: function addValidatorSet(uint64 valSetId, (address,uint64)[] validators) returns()
func (_OmniPortal *OmniPortalTransactorSession) AddValidatorSet(valSetId uint64, validators []XTypesValidator) (*types.Transaction, error) {
	return _OmniPortal.Contract.AddValidatorSet(&_OmniPortal.TransactOpts, valSetId, validators)
}

// CollectFees is a paid mutator transaction binding the contract method 0xa480ca79.
//
// Solidity: function collectFees(address to) returns()
func (_OmniPortal *OmniPortalTransactor) CollectFees(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _OmniPortal.contract.Transact(opts, "collectFees", to)
}

// CollectFees is a paid mutator transaction binding the contract method 0xa480ca79.
//
// Solidity: function collectFees(address to) returns()
func (_OmniPortal *OmniPortalSession) CollectFees(to common.Address) (*types.Transaction, error) {
	return _OmniPortal.Contract.CollectFees(&_OmniPortal.TransactOpts, to)
}

// CollectFees is a paid mutator transaction binding the contract method 0xa480ca79.
//
// Solidity: function collectFees(address to) returns()
func (_OmniPortal *OmniPortalTransactorSession) CollectFees(to common.Address) (*types.Transaction, error) {
	return _OmniPortal.Contract.CollectFees(&_OmniPortal.TransactOpts, to)
}

// Initialize is a paid mutator transaction binding the contract method 0x30632e8b.
//
// Solidity: function initialize((address,address,uint64,uint64,uint64,uint64,uint16,uint16,uint8,uint64,uint64,uint64,(address,uint64)[]) p) returns()
func (_OmniPortal *OmniPortalTransactor) Initialize(opts *bind.TransactOpts, p OmniPortalInitParams) (*types.Transaction, error) {
	return _OmniPortal.contract.Transact(opts, "initialize", p)
}

// Initialize is a paid mutator transaction binding the contract method 0x30632e8b.
//
// Solidity: function initialize((address,address,uint64,uint64,uint64,uint64,uint16,uint16,uint8,uint64,uint64,uint64,(address,uint64)[]) p) returns()
func (_OmniPortal *OmniPortalSession) Initialize(p OmniPortalInitParams) (*types.Transaction, error) {
	return _OmniPortal.Contract.Initialize(&_OmniPortal.TransactOpts, p)
}

// Initialize is a paid mutator transaction binding the contract method 0x30632e8b.
//
// Solidity: function initialize((address,address,uint64,uint64,uint64,uint64,uint16,uint16,uint8,uint64,uint64,uint64,(address,uint64)[]) p) returns()
func (_OmniPortal *OmniPortalTransactorSession) Initialize(p OmniPortalInitParams) (*types.Transaction, error) {
	return _OmniPortal.Contract.Initialize(&_OmniPortal.TransactOpts, p)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_OmniPortal *OmniPortalTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OmniPortal.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_OmniPortal *OmniPortalSession) Pause() (*types.Transaction, error) {
	return _OmniPortal.Contract.Pause(&_OmniPortal.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_OmniPortal *OmniPortalTransactorSession) Pause() (*types.Transaction, error) {
	return _OmniPortal.Contract.Pause(&_OmniPortal.TransactOpts)
}

// PauseXCall is a paid mutator transaction binding the contract method 0x83d0cbd9.
//
// Solidity: function pauseXCall() returns()
func (_OmniPortal *OmniPortalTransactor) PauseXCall(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OmniPortal.contract.Transact(opts, "pauseXCall")
}

// PauseXCall is a paid mutator transaction binding the contract method 0x83d0cbd9.
//
// Solidity: function pauseXCall() returns()
func (_OmniPortal *OmniPortalSession) PauseXCall() (*types.Transaction, error) {
	return _OmniPortal.Contract.PauseXCall(&_OmniPortal.TransactOpts)
}

// PauseXCall is a paid mutator transaction binding the contract method 0x83d0cbd9.
//
// Solidity: function pauseXCall() returns()
func (_OmniPortal *OmniPortalTransactorSession) PauseXCall() (*types.Transaction, error) {
	return _OmniPortal.Contract.PauseXCall(&_OmniPortal.TransactOpts)
}

// PauseXCallTo is a paid mutator transaction binding the contract method 0x10a5a7f7.
//
// Solidity: function pauseXCallTo(uint64 chainId_) returns()
func (_OmniPortal *OmniPortalTransactor) PauseXCallTo(opts *bind.TransactOpts, chainId_ uint64) (*types.Transaction, error) {
	return _OmniPortal.contract.Transact(opts, "pauseXCallTo", chainId_)
}

// PauseXCallTo is a paid mutator transaction binding the contract method 0x10a5a7f7.
//
// Solidity: function pauseXCallTo(uint64 chainId_) returns()
func (_OmniPortal *OmniPortalSession) PauseXCallTo(chainId_ uint64) (*types.Transaction, error) {
	return _OmniPortal.Contract.PauseXCallTo(&_OmniPortal.TransactOpts, chainId_)
}

// PauseXCallTo is a paid mutator transaction binding the contract method 0x10a5a7f7.
//
// Solidity: function pauseXCallTo(uint64 chainId_) returns()
func (_OmniPortal *OmniPortalTransactorSession) PauseXCallTo(chainId_ uint64) (*types.Transaction, error) {
	return _OmniPortal.Contract.PauseXCallTo(&_OmniPortal.TransactOpts, chainId_)
}

// PauseXSubmit is a paid mutator transaction binding the contract method 0x23dbce50.
//
// Solidity: function pauseXSubmit() returns()
func (_OmniPortal *OmniPortalTransactor) PauseXSubmit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OmniPortal.contract.Transact(opts, "pauseXSubmit")
}

// PauseXSubmit is a paid mutator transaction binding the contract method 0x23dbce50.
//
// Solidity: function pauseXSubmit() returns()
func (_OmniPortal *OmniPortalSession) PauseXSubmit() (*types.Transaction, error) {
	return _OmniPortal.Contract.PauseXSubmit(&_OmniPortal.TransactOpts)
}

// PauseXSubmit is a paid mutator transaction binding the contract method 0x23dbce50.
//
// Solidity: function pauseXSubmit() returns()
func (_OmniPortal *OmniPortalTransactorSession) PauseXSubmit() (*types.Transaction, error) {
	return _OmniPortal.Contract.PauseXSubmit(&_OmniPortal.TransactOpts)
}

// PauseXSubmitFrom is a paid mutator transaction binding the contract method 0xafe82198.
//
// Solidity: function pauseXSubmitFrom(uint64 chainId_) returns()
func (_OmniPortal *OmniPortalTransactor) PauseXSubmitFrom(opts *bind.TransactOpts, chainId_ uint64) (*types.Transaction, error) {
	return _OmniPortal.contract.Transact(opts, "pauseXSubmitFrom", chainId_)
}

// PauseXSubmitFrom is a paid mutator transaction binding the contract method 0xafe82198.
//
// Solidity: function pauseXSubmitFrom(uint64 chainId_) returns()
func (_OmniPortal *OmniPortalSession) PauseXSubmitFrom(chainId_ uint64) (*types.Transaction, error) {
	return _OmniPortal.Contract.PauseXSubmitFrom(&_OmniPortal.TransactOpts, chainId_)
}

// PauseXSubmitFrom is a paid mutator transaction binding the contract method 0xafe82198.
//
// Solidity: function pauseXSubmitFrom(uint64 chainId_) returns()
func (_OmniPortal *OmniPortalTransactorSession) PauseXSubmitFrom(chainId_ uint64) (*types.Transaction, error) {
	return _OmniPortal.Contract.PauseXSubmitFrom(&_OmniPortal.TransactOpts, chainId_)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_OmniPortal *OmniPortalTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OmniPortal.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_OmniPortal *OmniPortalSession) RenounceOwnership() (*types.Transaction, error) {
	return _OmniPortal.Contract.RenounceOwnership(&_OmniPortal.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_OmniPortal *OmniPortalTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _OmniPortal.Contract.RenounceOwnership(&_OmniPortal.TransactOpts)
}

// SetFeeOracle is a paid mutator transaction binding the contract method 0xa8a98962.
//
// Solidity: function setFeeOracle(address feeOracle_) returns()
func (_OmniPortal *OmniPortalTransactor) SetFeeOracle(opts *bind.TransactOpts, feeOracle_ common.Address) (*types.Transaction, error) {
	return _OmniPortal.contract.Transact(opts, "setFeeOracle", feeOracle_)
}

// SetFeeOracle is a paid mutator transaction binding the contract method 0xa8a98962.
//
// Solidity: function setFeeOracle(address feeOracle_) returns()
func (_OmniPortal *OmniPortalSession) SetFeeOracle(feeOracle_ common.Address) (*types.Transaction, error) {
	return _OmniPortal.Contract.SetFeeOracle(&_OmniPortal.TransactOpts, feeOracle_)
}

// SetFeeOracle is a paid mutator transaction binding the contract method 0xa8a98962.
//
// Solidity: function setFeeOracle(address feeOracle_) returns()
func (_OmniPortal *OmniPortalTransactorSession) SetFeeOracle(feeOracle_ common.Address) (*types.Transaction, error) {
	return _OmniPortal.Contract.SetFeeOracle(&_OmniPortal.TransactOpts, feeOracle_)
}

// SetInXBlockOffset is a paid mutator transaction binding the contract method 0x97b52062.
//
// Solidity: function setInXBlockOffset(uint64 sourceChainId, uint64 shardId, uint64 offset) returns()
func (_OmniPortal *OmniPortalTransactor) SetInXBlockOffset(opts *bind.TransactOpts, sourceChainId uint64, shardId uint64, offset uint64) (*types.Transaction, error) {
	return _OmniPortal.contract.Transact(opts, "setInXBlockOffset", sourceChainId, shardId, offset)
}

// SetInXBlockOffset is a paid mutator transaction binding the contract method 0x97b52062.
//
// Solidity: function setInXBlockOffset(uint64 sourceChainId, uint64 shardId, uint64 offset) returns()
func (_OmniPortal *OmniPortalSession) SetInXBlockOffset(sourceChainId uint64, shardId uint64, offset uint64) (*types.Transaction, error) {
	return _OmniPortal.Contract.SetInXBlockOffset(&_OmniPortal.TransactOpts, sourceChainId, shardId, offset)
}

// SetInXBlockOffset is a paid mutator transaction binding the contract method 0x97b52062.
//
// Solidity: function setInXBlockOffset(uint64 sourceChainId, uint64 shardId, uint64 offset) returns()
func (_OmniPortal *OmniPortalTransactorSession) SetInXBlockOffset(sourceChainId uint64, shardId uint64, offset uint64) (*types.Transaction, error) {
	return _OmniPortal.Contract.SetInXBlockOffset(&_OmniPortal.TransactOpts, sourceChainId, shardId, offset)
}

// SetInXMsgOffset is a paid mutator transaction binding the contract method 0xc4ab80bc.
//
// Solidity: function setInXMsgOffset(uint64 sourceChainId, uint64 shardId, uint64 offset) returns()
func (_OmniPortal *OmniPortalTransactor) SetInXMsgOffset(opts *bind.TransactOpts, sourceChainId uint64, shardId uint64, offset uint64) (*types.Transaction, error) {
	return _OmniPortal.contract.Transact(opts, "setInXMsgOffset", sourceChainId, shardId, offset)
}

// SetInXMsgOffset is a paid mutator transaction binding the contract method 0xc4ab80bc.
//
// Solidity: function setInXMsgOffset(uint64 sourceChainId, uint64 shardId, uint64 offset) returns()
func (_OmniPortal *OmniPortalSession) SetInXMsgOffset(sourceChainId uint64, shardId uint64, offset uint64) (*types.Transaction, error) {
	return _OmniPortal.Contract.SetInXMsgOffset(&_OmniPortal.TransactOpts, sourceChainId, shardId, offset)
}

// SetInXMsgOffset is a paid mutator transaction binding the contract method 0xc4ab80bc.
//
// Solidity: function setInXMsgOffset(uint64 sourceChainId, uint64 shardId, uint64 offset) returns()
func (_OmniPortal *OmniPortalTransactorSession) SetInXMsgOffset(sourceChainId uint64, shardId uint64, offset uint64) (*types.Transaction, error) {
	return _OmniPortal.Contract.SetInXMsgOffset(&_OmniPortal.TransactOpts, sourceChainId, shardId, offset)
}

// SetNetwork is a paid mutator transaction binding the contract method 0x1d3eb6e3.
//
// Solidity: function setNetwork((uint64,uint64[])[] network_) returns()
func (_OmniPortal *OmniPortalTransactor) SetNetwork(opts *bind.TransactOpts, network_ []XTypesChain) (*types.Transaction, error) {
	return _OmniPortal.contract.Transact(opts, "setNetwork", network_)
}

// SetNetwork is a paid mutator transaction binding the contract method 0x1d3eb6e3.
//
// Solidity: function setNetwork((uint64,uint64[])[] network_) returns()
func (_OmniPortal *OmniPortalSession) SetNetwork(network_ []XTypesChain) (*types.Transaction, error) {
	return _OmniPortal.Contract.SetNetwork(&_OmniPortal.TransactOpts, network_)
}

// SetNetwork is a paid mutator transaction binding the contract method 0x1d3eb6e3.
//
// Solidity: function setNetwork((uint64,uint64[])[] network_) returns()
func (_OmniPortal *OmniPortalTransactorSession) SetNetwork(network_ []XTypesChain) (*types.Transaction, error) {
	return _OmniPortal.Contract.SetNetwork(&_OmniPortal.TransactOpts, network_)
}

// SetXMsgMaxDataSize is a paid mutator transaction binding the contract method 0xb521466d.
//
// Solidity: function setXMsgMaxDataSize(uint16 numBytes) returns()
func (_OmniPortal *OmniPortalTransactor) SetXMsgMaxDataSize(opts *bind.TransactOpts, numBytes uint16) (*types.Transaction, error) {
	return _OmniPortal.contract.Transact(opts, "setXMsgMaxDataSize", numBytes)
}

// SetXMsgMaxDataSize is a paid mutator transaction binding the contract method 0xb521466d.
//
// Solidity: function setXMsgMaxDataSize(uint16 numBytes) returns()
func (_OmniPortal *OmniPortalSession) SetXMsgMaxDataSize(numBytes uint16) (*types.Transaction, error) {
	return _OmniPortal.Contract.SetXMsgMaxDataSize(&_OmniPortal.TransactOpts, numBytes)
}

// SetXMsgMaxDataSize is a paid mutator transaction binding the contract method 0xb521466d.
//
// Solidity: function setXMsgMaxDataSize(uint16 numBytes) returns()
func (_OmniPortal *OmniPortalTransactorSession) SetXMsgMaxDataSize(numBytes uint16) (*types.Transaction, error) {
	return _OmniPortal.Contract.SetXMsgMaxDataSize(&_OmniPortal.TransactOpts, numBytes)
}

// SetXMsgMaxGasLimit is a paid mutator transaction binding the contract method 0x36d853f9.
//
// Solidity: function setXMsgMaxGasLimit(uint64 gasLimit) returns()
func (_OmniPortal *OmniPortalTransactor) SetXMsgMaxGasLimit(opts *bind.TransactOpts, gasLimit uint64) (*types.Transaction, error) {
	return _OmniPortal.contract.Transact(opts, "setXMsgMaxGasLimit", gasLimit)
}

// SetXMsgMaxGasLimit is a paid mutator transaction binding the contract method 0x36d853f9.
//
// Solidity: function setXMsgMaxGasLimit(uint64 gasLimit) returns()
func (_OmniPortal *OmniPortalSession) SetXMsgMaxGasLimit(gasLimit uint64) (*types.Transaction, error) {
	return _OmniPortal.Contract.SetXMsgMaxGasLimit(&_OmniPortal.TransactOpts, gasLimit)
}

// SetXMsgMaxGasLimit is a paid mutator transaction binding the contract method 0x36d853f9.
//
// Solidity: function setXMsgMaxGasLimit(uint64 gasLimit) returns()
func (_OmniPortal *OmniPortalTransactorSession) SetXMsgMaxGasLimit(gasLimit uint64) (*types.Transaction, error) {
	return _OmniPortal.Contract.SetXMsgMaxGasLimit(&_OmniPortal.TransactOpts, gasLimit)
}

// SetXMsgMinGasLimit is a paid mutator transaction binding the contract method 0xbb8590ad.
//
// Solidity: function setXMsgMinGasLimit(uint64 gasLimit) returns()
func (_OmniPortal *OmniPortalTransactor) SetXMsgMinGasLimit(opts *bind.TransactOpts, gasLimit uint64) (*types.Transaction, error) {
	return _OmniPortal.contract.Transact(opts, "setXMsgMinGasLimit", gasLimit)
}

// SetXMsgMinGasLimit is a paid mutator transaction binding the contract method 0xbb8590ad.
//
// Solidity: function setXMsgMinGasLimit(uint64 gasLimit) returns()
func (_OmniPortal *OmniPortalSession) SetXMsgMinGasLimit(gasLimit uint64) (*types.Transaction, error) {
	return _OmniPortal.Contract.SetXMsgMinGasLimit(&_OmniPortal.TransactOpts, gasLimit)
}

// SetXMsgMinGasLimit is a paid mutator transaction binding the contract method 0xbb8590ad.
//
// Solidity: function setXMsgMinGasLimit(uint64 gasLimit) returns()
func (_OmniPortal *OmniPortalTransactorSession) SetXMsgMinGasLimit(gasLimit uint64) (*types.Transaction, error) {
	return _OmniPortal.Contract.SetXMsgMinGasLimit(&_OmniPortal.TransactOpts, gasLimit)
}

// SetXReceiptMaxErrorSize is a paid mutator transaction binding the contract method 0xbff0e84d.
//
// Solidity: function setXReceiptMaxErrorSize(uint16 numBytes) returns()
func (_OmniPortal *OmniPortalTransactor) SetXReceiptMaxErrorSize(opts *bind.TransactOpts, numBytes uint16) (*types.Transaction, error) {
	return _OmniPortal.contract.Transact(opts, "setXReceiptMaxErrorSize", numBytes)
}

// SetXReceiptMaxErrorSize is a paid mutator transaction binding the contract method 0xbff0e84d.
//
// Solidity: function setXReceiptMaxErrorSize(uint16 numBytes) returns()
func (_OmniPortal *OmniPortalSession) SetXReceiptMaxErrorSize(numBytes uint16) (*types.Transaction, error) {
	return _OmniPortal.Contract.SetXReceiptMaxErrorSize(&_OmniPortal.TransactOpts, numBytes)
}

// SetXReceiptMaxErrorSize is a paid mutator transaction binding the contract method 0xbff0e84d.
//
// Solidity: function setXReceiptMaxErrorSize(uint16 numBytes) returns()
func (_OmniPortal *OmniPortalTransactorSession) SetXReceiptMaxErrorSize(numBytes uint16) (*types.Transaction, error) {
	return _OmniPortal.Contract.SetXReceiptMaxErrorSize(&_OmniPortal.TransactOpts, numBytes)
}

// SetXSubValsetCutoff is a paid mutator transaction binding the contract method 0x103ba701.
//
// Solidity: function setXSubValsetCutoff(uint8 xsubValsetCutoff_) returns()
func (_OmniPortal *OmniPortalTransactor) SetXSubValsetCutoff(opts *bind.TransactOpts, xsubValsetCutoff_ uint8) (*types.Transaction, error) {
	return _OmniPortal.contract.Transact(opts, "setXSubValsetCutoff", xsubValsetCutoff_)
}

// SetXSubValsetCutoff is a paid mutator transaction binding the contract method 0x103ba701.
//
// Solidity: function setXSubValsetCutoff(uint8 xsubValsetCutoff_) returns()
func (_OmniPortal *OmniPortalSession) SetXSubValsetCutoff(xsubValsetCutoff_ uint8) (*types.Transaction, error) {
	return _OmniPortal.Contract.SetXSubValsetCutoff(&_OmniPortal.TransactOpts, xsubValsetCutoff_)
}

// SetXSubValsetCutoff is a paid mutator transaction binding the contract method 0x103ba701.
//
// Solidity: function setXSubValsetCutoff(uint8 xsubValsetCutoff_) returns()
func (_OmniPortal *OmniPortalTransactorSession) SetXSubValsetCutoff(xsubValsetCutoff_ uint8) (*types.Transaction, error) {
	return _OmniPortal.Contract.SetXSubValsetCutoff(&_OmniPortal.TransactOpts, xsubValsetCutoff_)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_OmniPortal *OmniPortalTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _OmniPortal.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_OmniPortal *OmniPortalSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _OmniPortal.Contract.TransferOwnership(&_OmniPortal.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_OmniPortal *OmniPortalTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _OmniPortal.Contract.TransferOwnership(&_OmniPortal.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_OmniPortal *OmniPortalTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OmniPortal.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_OmniPortal *OmniPortalSession) Unpause() (*types.Transaction, error) {
	return _OmniPortal.Contract.Unpause(&_OmniPortal.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_OmniPortal *OmniPortalTransactorSession) Unpause() (*types.Transaction, error) {
	return _OmniPortal.Contract.Unpause(&_OmniPortal.TransactOpts)
}

// UnpauseXCall is a paid mutator transaction binding the contract method 0x54d26bba.
//
// Solidity: function unpauseXCall() returns()
func (_OmniPortal *OmniPortalTransactor) UnpauseXCall(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OmniPortal.contract.Transact(opts, "unpauseXCall")
}

// UnpauseXCall is a paid mutator transaction binding the contract method 0x54d26bba.
//
// Solidity: function unpauseXCall() returns()
func (_OmniPortal *OmniPortalSession) UnpauseXCall() (*types.Transaction, error) {
	return _OmniPortal.Contract.UnpauseXCall(&_OmniPortal.TransactOpts)
}

// UnpauseXCall is a paid mutator transaction binding the contract method 0x54d26bba.
//
// Solidity: function unpauseXCall() returns()
func (_OmniPortal *OmniPortalTransactorSession) UnpauseXCall() (*types.Transaction, error) {
	return _OmniPortal.Contract.UnpauseXCall(&_OmniPortal.TransactOpts)
}

// UnpauseXCallTo is a paid mutator transaction binding the contract method 0xd533b445.
//
// Solidity: function unpauseXCallTo(uint64 chainId_) returns()
func (_OmniPortal *OmniPortalTransactor) UnpauseXCallTo(opts *bind.TransactOpts, chainId_ uint64) (*types.Transaction, error) {
	return _OmniPortal.contract.Transact(opts, "unpauseXCallTo", chainId_)
}

// UnpauseXCallTo is a paid mutator transaction binding the contract method 0xd533b445.
//
// Solidity: function unpauseXCallTo(uint64 chainId_) returns()
func (_OmniPortal *OmniPortalSession) UnpauseXCallTo(chainId_ uint64) (*types.Transaction, error) {
	return _OmniPortal.Contract.UnpauseXCallTo(&_OmniPortal.TransactOpts, chainId_)
}

// UnpauseXCallTo is a paid mutator transaction binding the contract method 0xd533b445.
//
// Solidity: function unpauseXCallTo(uint64 chainId_) returns()
func (_OmniPortal *OmniPortalTransactorSession) UnpauseXCallTo(chainId_ uint64) (*types.Transaction, error) {
	return _OmniPortal.Contract.UnpauseXCallTo(&_OmniPortal.TransactOpts, chainId_)
}

// UnpauseXSubmit is a paid mutator transaction binding the contract method 0xc3d8ad67.
//
// Solidity: function unpauseXSubmit() returns()
func (_OmniPortal *OmniPortalTransactor) UnpauseXSubmit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _OmniPortal.contract.Transact(opts, "unpauseXSubmit")
}

// UnpauseXSubmit is a paid mutator transaction binding the contract method 0xc3d8ad67.
//
// Solidity: function unpauseXSubmit() returns()
func (_OmniPortal *OmniPortalSession) UnpauseXSubmit() (*types.Transaction, error) {
	return _OmniPortal.Contract.UnpauseXSubmit(&_OmniPortal.TransactOpts)
}

// UnpauseXSubmit is a paid mutator transaction binding the contract method 0xc3d8ad67.
//
// Solidity: function unpauseXSubmit() returns()
func (_OmniPortal *OmniPortalTransactorSession) UnpauseXSubmit() (*types.Transaction, error) {
	return _OmniPortal.Contract.UnpauseXSubmit(&_OmniPortal.TransactOpts)
}

// UnpauseXSubmitFrom is a paid mutator transaction binding the contract method 0xc2f9b968.
//
// Solidity: function unpauseXSubmitFrom(uint64 chainId_) returns()
func (_OmniPortal *OmniPortalTransactor) UnpauseXSubmitFrom(opts *bind.TransactOpts, chainId_ uint64) (*types.Transaction, error) {
	return _OmniPortal.contract.Transact(opts, "unpauseXSubmitFrom", chainId_)
}

// UnpauseXSubmitFrom is a paid mutator transaction binding the contract method 0xc2f9b968.
//
// Solidity: function unpauseXSubmitFrom(uint64 chainId_) returns()
func (_OmniPortal *OmniPortalSession) UnpauseXSubmitFrom(chainId_ uint64) (*types.Transaction, error) {
	return _OmniPortal.Contract.UnpauseXSubmitFrom(&_OmniPortal.TransactOpts, chainId_)
}

// UnpauseXSubmitFrom is a paid mutator transaction binding the contract method 0xc2f9b968.
//
// Solidity: function unpauseXSubmitFrom(uint64 chainId_) returns()
func (_OmniPortal *OmniPortalTransactorSession) UnpauseXSubmitFrom(chainId_ uint64) (*types.Transaction, error) {
	return _OmniPortal.Contract.UnpauseXSubmitFrom(&_OmniPortal.TransactOpts, chainId_)
}

// Xcall is a paid mutator transaction binding the contract method 0xc21dda4f.
//
// Solidity: function xcall(uint64 destChainId, uint8 conf, address to, bytes data, uint64 gasLimit) payable returns()
func (_OmniPortal *OmniPortalTransactor) Xcall(opts *bind.TransactOpts, destChainId uint64, conf uint8, to common.Address, data []byte, gasLimit uint64) (*types.Transaction, error) {
	return _OmniPortal.contract.Transact(opts, "xcall", destChainId, conf, to, data, gasLimit)
}

// Xcall is a paid mutator transaction binding the contract method 0xc21dda4f.
//
// Solidity: function xcall(uint64 destChainId, uint8 conf, address to, bytes data, uint64 gasLimit) payable returns()
func (_OmniPortal *OmniPortalSession) Xcall(destChainId uint64, conf uint8, to common.Address, data []byte, gasLimit uint64) (*types.Transaction, error) {
	return _OmniPortal.Contract.Xcall(&_OmniPortal.TransactOpts, destChainId, conf, to, data, gasLimit)
}

// Xcall is a paid mutator transaction binding the contract method 0xc21dda4f.
//
// Solidity: function xcall(uint64 destChainId, uint8 conf, address to, bytes data, uint64 gasLimit) payable returns()
func (_OmniPortal *OmniPortalTransactorSession) Xcall(destChainId uint64, conf uint8, to common.Address, data []byte, gasLimit uint64) (*types.Transaction, error) {
	return _OmniPortal.Contract.Xcall(&_OmniPortal.TransactOpts, destChainId, conf, to, data, gasLimit)
}

// Xsubmit is a paid mutator transaction binding the contract method 0x23f2ca20.
//
// Solidity: function xsubmit((bytes32,uint64,(uint64,uint8,uint64,uint64,bytes32),(uint64,uint64,uint64,uint64,address,address,bytes,uint64)[],bytes32[],bool[],(address,bytes)[]) xsub) returns()
func (_OmniPortal *OmniPortalTransactor) Xsubmit(opts *bind.TransactOpts, xsub XTypesSubmission) (*types.Transaction, error) {
	return _OmniPortal.contract.Transact(opts, "xsubmit", xsub)
}

// Xsubmit is a paid mutator transaction binding the contract method 0x23f2ca20.
//
// Solidity: function xsubmit((bytes32,uint64,(uint64,uint8,uint64,uint64,bytes32),(uint64,uint64,uint64,uint64,address,address,bytes,uint64)[],bytes32[],bool[],(address,bytes)[]) xsub) returns()
func (_OmniPortal *OmniPortalSession) Xsubmit(xsub XTypesSubmission) (*types.Transaction, error) {
	return _OmniPortal.Contract.Xsubmit(&_OmniPortal.TransactOpts, xsub)
}

// Xsubmit is a paid mutator transaction binding the contract method 0x23f2ca20.
//
// Solidity: function xsubmit((bytes32,uint64,(uint64,uint8,uint64,uint64,bytes32),(uint64,uint64,uint64,uint64,address,address,bytes,uint64)[],bytes32[],bool[],(address,bytes)[]) xsub) returns()
func (_OmniPortal *OmniPortalTransactorSession) Xsubmit(xsub XTypesSubmission) (*types.Transaction, error) {
	return _OmniPortal.Contract.Xsubmit(&_OmniPortal.TransactOpts, xsub)
}

// OmniPortalFeeOracleSetIterator is returned from FilterFeeOracleSet and is used to iterate over the raw logs and unpacked data for FeeOracleSet events raised by the OmniPortal contract.
type OmniPortalFeeOracleSetIterator struct {
	Event *OmniPortalFeeOracleSet // Event containing the contract specifics and raw log

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
func (it *OmniPortalFeeOracleSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OmniPortalFeeOracleSet)
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
		it.Event = new(OmniPortalFeeOracleSet)
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
func (it *OmniPortalFeeOracleSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OmniPortalFeeOracleSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OmniPortalFeeOracleSet represents a FeeOracleSet event raised by the OmniPortal contract.
type OmniPortalFeeOracleSet struct {
	Oracle common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterFeeOracleSet is a free log retrieval operation binding the contract event 0xd97bdb0db82b52a85aa07f8da78033b1d6e159d94f1e3cbd4109d946c3bcfd32.
//
// Solidity: event FeeOracleSet(address oracle)
func (_OmniPortal *OmniPortalFilterer) FilterFeeOracleSet(opts *bind.FilterOpts) (*OmniPortalFeeOracleSetIterator, error) {

	logs, sub, err := _OmniPortal.contract.FilterLogs(opts, "FeeOracleSet")
	if err != nil {
		return nil, err
	}
	return &OmniPortalFeeOracleSetIterator{contract: _OmniPortal.contract, event: "FeeOracleSet", logs: logs, sub: sub}, nil
}

// WatchFeeOracleSet is a free log subscription operation binding the contract event 0xd97bdb0db82b52a85aa07f8da78033b1d6e159d94f1e3cbd4109d946c3bcfd32.
//
// Solidity: event FeeOracleSet(address oracle)
func (_OmniPortal *OmniPortalFilterer) WatchFeeOracleSet(opts *bind.WatchOpts, sink chan<- *OmniPortalFeeOracleSet) (event.Subscription, error) {

	logs, sub, err := _OmniPortal.contract.WatchLogs(opts, "FeeOracleSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OmniPortalFeeOracleSet)
				if err := _OmniPortal.contract.UnpackLog(event, "FeeOracleSet", log); err != nil {
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
func (_OmniPortal *OmniPortalFilterer) ParseFeeOracleSet(log types.Log) (*OmniPortalFeeOracleSet, error) {
	event := new(OmniPortalFeeOracleSet)
	if err := _OmniPortal.contract.UnpackLog(event, "FeeOracleSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OmniPortalFeesCollectedIterator is returned from FilterFeesCollected and is used to iterate over the raw logs and unpacked data for FeesCollected events raised by the OmniPortal contract.
type OmniPortalFeesCollectedIterator struct {
	Event *OmniPortalFeesCollected // Event containing the contract specifics and raw log

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
func (it *OmniPortalFeesCollectedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OmniPortalFeesCollected)
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
		it.Event = new(OmniPortalFeesCollected)
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
func (it *OmniPortalFeesCollectedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OmniPortalFeesCollectedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OmniPortalFeesCollected represents a FeesCollected event raised by the OmniPortal contract.
type OmniPortalFeesCollected struct {
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterFeesCollected is a free log retrieval operation binding the contract event 0x9dc46f23cfb5ddcad0ae7ea2be38d47fec07bb9382ec7e564efc69e036dd66ce.
//
// Solidity: event FeesCollected(address indexed to, uint256 amount)
func (_OmniPortal *OmniPortalFilterer) FilterFeesCollected(opts *bind.FilterOpts, to []common.Address) (*OmniPortalFeesCollectedIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _OmniPortal.contract.FilterLogs(opts, "FeesCollected", toRule)
	if err != nil {
		return nil, err
	}
	return &OmniPortalFeesCollectedIterator{contract: _OmniPortal.contract, event: "FeesCollected", logs: logs, sub: sub}, nil
}

// WatchFeesCollected is a free log subscription operation binding the contract event 0x9dc46f23cfb5ddcad0ae7ea2be38d47fec07bb9382ec7e564efc69e036dd66ce.
//
// Solidity: event FeesCollected(address indexed to, uint256 amount)
func (_OmniPortal *OmniPortalFilterer) WatchFeesCollected(opts *bind.WatchOpts, sink chan<- *OmniPortalFeesCollected, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _OmniPortal.contract.WatchLogs(opts, "FeesCollected", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OmniPortalFeesCollected)
				if err := _OmniPortal.contract.UnpackLog(event, "FeesCollected", log); err != nil {
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
func (_OmniPortal *OmniPortalFilterer) ParseFeesCollected(log types.Log) (*OmniPortalFeesCollected, error) {
	event := new(OmniPortalFeesCollected)
	if err := _OmniPortal.contract.UnpackLog(event, "FeesCollected", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OmniPortalInXBlockOffsetSetIterator is returned from FilterInXBlockOffsetSet and is used to iterate over the raw logs and unpacked data for InXBlockOffsetSet events raised by the OmniPortal contract.
type OmniPortalInXBlockOffsetSetIterator struct {
	Event *OmniPortalInXBlockOffsetSet // Event containing the contract specifics and raw log

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
func (it *OmniPortalInXBlockOffsetSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OmniPortalInXBlockOffsetSet)
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
		it.Event = new(OmniPortalInXBlockOffsetSet)
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
func (it *OmniPortalInXBlockOffsetSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OmniPortalInXBlockOffsetSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OmniPortalInXBlockOffsetSet represents a InXBlockOffsetSet event raised by the OmniPortal contract.
type OmniPortalInXBlockOffsetSet struct {
	SrcChainId uint64
	ShardId    uint64
	Offset     uint64
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterInXBlockOffsetSet is a free log retrieval operation binding the contract event 0xe070f08cae8464c91238e8cbea64ccee5e7b48dd79a843f144e3721ee6bdd9b5.
//
// Solidity: event InXBlockOffsetSet(uint64 indexed srcChainId, uint64 indexed shardId, uint64 offset)
func (_OmniPortal *OmniPortalFilterer) FilterInXBlockOffsetSet(opts *bind.FilterOpts, srcChainId []uint64, shardId []uint64) (*OmniPortalInXBlockOffsetSetIterator, error) {

	var srcChainIdRule []interface{}
	for _, srcChainIdItem := range srcChainId {
		srcChainIdRule = append(srcChainIdRule, srcChainIdItem)
	}
	var shardIdRule []interface{}
	for _, shardIdItem := range shardId {
		shardIdRule = append(shardIdRule, shardIdItem)
	}

	logs, sub, err := _OmniPortal.contract.FilterLogs(opts, "InXBlockOffsetSet", srcChainIdRule, shardIdRule)
	if err != nil {
		return nil, err
	}
	return &OmniPortalInXBlockOffsetSetIterator{contract: _OmniPortal.contract, event: "InXBlockOffsetSet", logs: logs, sub: sub}, nil
}

// WatchInXBlockOffsetSet is a free log subscription operation binding the contract event 0xe070f08cae8464c91238e8cbea64ccee5e7b48dd79a843f144e3721ee6bdd9b5.
//
// Solidity: event InXBlockOffsetSet(uint64 indexed srcChainId, uint64 indexed shardId, uint64 offset)
func (_OmniPortal *OmniPortalFilterer) WatchInXBlockOffsetSet(opts *bind.WatchOpts, sink chan<- *OmniPortalInXBlockOffsetSet, srcChainId []uint64, shardId []uint64) (event.Subscription, error) {

	var srcChainIdRule []interface{}
	for _, srcChainIdItem := range srcChainId {
		srcChainIdRule = append(srcChainIdRule, srcChainIdItem)
	}
	var shardIdRule []interface{}
	for _, shardIdItem := range shardId {
		shardIdRule = append(shardIdRule, shardIdItem)
	}

	logs, sub, err := _OmniPortal.contract.WatchLogs(opts, "InXBlockOffsetSet", srcChainIdRule, shardIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OmniPortalInXBlockOffsetSet)
				if err := _OmniPortal.contract.UnpackLog(event, "InXBlockOffsetSet", log); err != nil {
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
func (_OmniPortal *OmniPortalFilterer) ParseInXBlockOffsetSet(log types.Log) (*OmniPortalInXBlockOffsetSet, error) {
	event := new(OmniPortalInXBlockOffsetSet)
	if err := _OmniPortal.contract.UnpackLog(event, "InXBlockOffsetSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OmniPortalInXMsgOffsetSetIterator is returned from FilterInXMsgOffsetSet and is used to iterate over the raw logs and unpacked data for InXMsgOffsetSet events raised by the OmniPortal contract.
type OmniPortalInXMsgOffsetSetIterator struct {
	Event *OmniPortalInXMsgOffsetSet // Event containing the contract specifics and raw log

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
func (it *OmniPortalInXMsgOffsetSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OmniPortalInXMsgOffsetSet)
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
		it.Event = new(OmniPortalInXMsgOffsetSet)
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
func (it *OmniPortalInXMsgOffsetSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OmniPortalInXMsgOffsetSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OmniPortalInXMsgOffsetSet represents a InXMsgOffsetSet event raised by the OmniPortal contract.
type OmniPortalInXMsgOffsetSet struct {
	SrcChainId uint64
	ShardId    uint64
	Offset     uint64
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterInXMsgOffsetSet is a free log retrieval operation binding the contract event 0x8647aae68c8456a1dcbfaf5eaadc94278ae423526d3f09c7b972bff7355d55c7.
//
// Solidity: event InXMsgOffsetSet(uint64 indexed srcChainId, uint64 indexed shardId, uint64 offset)
func (_OmniPortal *OmniPortalFilterer) FilterInXMsgOffsetSet(opts *bind.FilterOpts, srcChainId []uint64, shardId []uint64) (*OmniPortalInXMsgOffsetSetIterator, error) {

	var srcChainIdRule []interface{}
	for _, srcChainIdItem := range srcChainId {
		srcChainIdRule = append(srcChainIdRule, srcChainIdItem)
	}
	var shardIdRule []interface{}
	for _, shardIdItem := range shardId {
		shardIdRule = append(shardIdRule, shardIdItem)
	}

	logs, sub, err := _OmniPortal.contract.FilterLogs(opts, "InXMsgOffsetSet", srcChainIdRule, shardIdRule)
	if err != nil {
		return nil, err
	}
	return &OmniPortalInXMsgOffsetSetIterator{contract: _OmniPortal.contract, event: "InXMsgOffsetSet", logs: logs, sub: sub}, nil
}

// WatchInXMsgOffsetSet is a free log subscription operation binding the contract event 0x8647aae68c8456a1dcbfaf5eaadc94278ae423526d3f09c7b972bff7355d55c7.
//
// Solidity: event InXMsgOffsetSet(uint64 indexed srcChainId, uint64 indexed shardId, uint64 offset)
func (_OmniPortal *OmniPortalFilterer) WatchInXMsgOffsetSet(opts *bind.WatchOpts, sink chan<- *OmniPortalInXMsgOffsetSet, srcChainId []uint64, shardId []uint64) (event.Subscription, error) {

	var srcChainIdRule []interface{}
	for _, srcChainIdItem := range srcChainId {
		srcChainIdRule = append(srcChainIdRule, srcChainIdItem)
	}
	var shardIdRule []interface{}
	for _, shardIdItem := range shardId {
		shardIdRule = append(shardIdRule, shardIdItem)
	}

	logs, sub, err := _OmniPortal.contract.WatchLogs(opts, "InXMsgOffsetSet", srcChainIdRule, shardIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OmniPortalInXMsgOffsetSet)
				if err := _OmniPortal.contract.UnpackLog(event, "InXMsgOffsetSet", log); err != nil {
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
func (_OmniPortal *OmniPortalFilterer) ParseInXMsgOffsetSet(log types.Log) (*OmniPortalInXMsgOffsetSet, error) {
	event := new(OmniPortalInXMsgOffsetSet)
	if err := _OmniPortal.contract.UnpackLog(event, "InXMsgOffsetSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OmniPortalInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the OmniPortal contract.
type OmniPortalInitializedIterator struct {
	Event *OmniPortalInitialized // Event containing the contract specifics and raw log

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
func (it *OmniPortalInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OmniPortalInitialized)
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
		it.Event = new(OmniPortalInitialized)
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
func (it *OmniPortalInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OmniPortalInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OmniPortalInitialized represents a Initialized event raised by the OmniPortal contract.
type OmniPortalInitialized struct {
	Version uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_OmniPortal *OmniPortalFilterer) FilterInitialized(opts *bind.FilterOpts) (*OmniPortalInitializedIterator, error) {

	logs, sub, err := _OmniPortal.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &OmniPortalInitializedIterator{contract: _OmniPortal.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0xc7f505b2f371ae2175ee4913f4499e1f2633a7b5936321eed1cdaeb6115181d2.
//
// Solidity: event Initialized(uint64 version)
func (_OmniPortal *OmniPortalFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *OmniPortalInitialized) (event.Subscription, error) {

	logs, sub, err := _OmniPortal.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OmniPortalInitialized)
				if err := _OmniPortal.contract.UnpackLog(event, "Initialized", log); err != nil {
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
func (_OmniPortal *OmniPortalFilterer) ParseInitialized(log types.Log) (*OmniPortalInitialized, error) {
	event := new(OmniPortalInitialized)
	if err := _OmniPortal.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OmniPortalOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the OmniPortal contract.
type OmniPortalOwnershipTransferredIterator struct {
	Event *OmniPortalOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *OmniPortalOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OmniPortalOwnershipTransferred)
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
		it.Event = new(OmniPortalOwnershipTransferred)
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
func (it *OmniPortalOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OmniPortalOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OmniPortalOwnershipTransferred represents a OwnershipTransferred event raised by the OmniPortal contract.
type OmniPortalOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_OmniPortal *OmniPortalFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*OmniPortalOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _OmniPortal.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &OmniPortalOwnershipTransferredIterator{contract: _OmniPortal.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_OmniPortal *OmniPortalFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *OmniPortalOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _OmniPortal.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OmniPortalOwnershipTransferred)
				if err := _OmniPortal.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_OmniPortal *OmniPortalFilterer) ParseOwnershipTransferred(log types.Log) (*OmniPortalOwnershipTransferred, error) {
	event := new(OmniPortalOwnershipTransferred)
	if err := _OmniPortal.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OmniPortalPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the OmniPortal contract.
type OmniPortalPausedIterator struct {
	Event *OmniPortalPaused // Event containing the contract specifics and raw log

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
func (it *OmniPortalPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OmniPortalPaused)
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
		it.Event = new(OmniPortalPaused)
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
func (it *OmniPortalPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OmniPortalPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OmniPortalPaused represents a Paused event raised by the OmniPortal contract.
type OmniPortalPaused struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x9e87fac88ff661f02d44f95383c817fece4bce600a3dab7a54406878b965e752.
//
// Solidity: event Paused()
func (_OmniPortal *OmniPortalFilterer) FilterPaused(opts *bind.FilterOpts) (*OmniPortalPausedIterator, error) {

	logs, sub, err := _OmniPortal.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &OmniPortalPausedIterator{contract: _OmniPortal.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x9e87fac88ff661f02d44f95383c817fece4bce600a3dab7a54406878b965e752.
//
// Solidity: event Paused()
func (_OmniPortal *OmniPortalFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *OmniPortalPaused) (event.Subscription, error) {

	logs, sub, err := _OmniPortal.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OmniPortalPaused)
				if err := _OmniPortal.contract.UnpackLog(event, "Paused", log); err != nil {
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
func (_OmniPortal *OmniPortalFilterer) ParsePaused(log types.Log) (*OmniPortalPaused, error) {
	event := new(OmniPortalPaused)
	if err := _OmniPortal.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OmniPortalUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the OmniPortal contract.
type OmniPortalUnpausedIterator struct {
	Event *OmniPortalUnpaused // Event containing the contract specifics and raw log

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
func (it *OmniPortalUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OmniPortalUnpaused)
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
		it.Event = new(OmniPortalUnpaused)
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
func (it *OmniPortalUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OmniPortalUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OmniPortalUnpaused represents a Unpaused event raised by the OmniPortal contract.
type OmniPortalUnpaused struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0xa45f47fdea8a1efdd9029a5691c7f759c32b7c698632b563573e155625d16933.
//
// Solidity: event Unpaused()
func (_OmniPortal *OmniPortalFilterer) FilterUnpaused(opts *bind.FilterOpts) (*OmniPortalUnpausedIterator, error) {

	logs, sub, err := _OmniPortal.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &OmniPortalUnpausedIterator{contract: _OmniPortal.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0xa45f47fdea8a1efdd9029a5691c7f759c32b7c698632b563573e155625d16933.
//
// Solidity: event Unpaused()
func (_OmniPortal *OmniPortalFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *OmniPortalUnpaused) (event.Subscription, error) {

	logs, sub, err := _OmniPortal.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OmniPortalUnpaused)
				if err := _OmniPortal.contract.UnpackLog(event, "Unpaused", log); err != nil {
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
func (_OmniPortal *OmniPortalFilterer) ParseUnpaused(log types.Log) (*OmniPortalUnpaused, error) {
	event := new(OmniPortalUnpaused)
	if err := _OmniPortal.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OmniPortalValidatorSetAddedIterator is returned from FilterValidatorSetAdded and is used to iterate over the raw logs and unpacked data for ValidatorSetAdded events raised by the OmniPortal contract.
type OmniPortalValidatorSetAddedIterator struct {
	Event *OmniPortalValidatorSetAdded // Event containing the contract specifics and raw log

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
func (it *OmniPortalValidatorSetAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OmniPortalValidatorSetAdded)
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
		it.Event = new(OmniPortalValidatorSetAdded)
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
func (it *OmniPortalValidatorSetAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OmniPortalValidatorSetAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OmniPortalValidatorSetAdded represents a ValidatorSetAdded event raised by the OmniPortal contract.
type OmniPortalValidatorSetAdded struct {
	SetId uint64
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterValidatorSetAdded is a free log retrieval operation binding the contract event 0x3a7c2f997a87ba92aedaecd1127f4129cae1283e2809ebf5304d321b943fd107.
//
// Solidity: event ValidatorSetAdded(uint64 indexed setId)
func (_OmniPortal *OmniPortalFilterer) FilterValidatorSetAdded(opts *bind.FilterOpts, setId []uint64) (*OmniPortalValidatorSetAddedIterator, error) {

	var setIdRule []interface{}
	for _, setIdItem := range setId {
		setIdRule = append(setIdRule, setIdItem)
	}

	logs, sub, err := _OmniPortal.contract.FilterLogs(opts, "ValidatorSetAdded", setIdRule)
	if err != nil {
		return nil, err
	}
	return &OmniPortalValidatorSetAddedIterator{contract: _OmniPortal.contract, event: "ValidatorSetAdded", logs: logs, sub: sub}, nil
}

// WatchValidatorSetAdded is a free log subscription operation binding the contract event 0x3a7c2f997a87ba92aedaecd1127f4129cae1283e2809ebf5304d321b943fd107.
//
// Solidity: event ValidatorSetAdded(uint64 indexed setId)
func (_OmniPortal *OmniPortalFilterer) WatchValidatorSetAdded(opts *bind.WatchOpts, sink chan<- *OmniPortalValidatorSetAdded, setId []uint64) (event.Subscription, error) {

	var setIdRule []interface{}
	for _, setIdItem := range setId {
		setIdRule = append(setIdRule, setIdItem)
	}

	logs, sub, err := _OmniPortal.contract.WatchLogs(opts, "ValidatorSetAdded", setIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OmniPortalValidatorSetAdded)
				if err := _OmniPortal.contract.UnpackLog(event, "ValidatorSetAdded", log); err != nil {
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
func (_OmniPortal *OmniPortalFilterer) ParseValidatorSetAdded(log types.Log) (*OmniPortalValidatorSetAdded, error) {
	event := new(OmniPortalValidatorSetAdded)
	if err := _OmniPortal.contract.UnpackLog(event, "ValidatorSetAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OmniPortalXCallPausedIterator is returned from FilterXCallPaused and is used to iterate over the raw logs and unpacked data for XCallPaused events raised by the OmniPortal contract.
type OmniPortalXCallPausedIterator struct {
	Event *OmniPortalXCallPaused // Event containing the contract specifics and raw log

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
func (it *OmniPortalXCallPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OmniPortalXCallPaused)
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
		it.Event = new(OmniPortalXCallPaused)
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
func (it *OmniPortalXCallPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OmniPortalXCallPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OmniPortalXCallPaused represents a XCallPaused event raised by the OmniPortal contract.
type OmniPortalXCallPaused struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterXCallPaused is a free log retrieval operation binding the contract event 0x5f335a4032d4cfb6aca7835b0c2225f36d4d9eaa4ed43ee59ed537e02dff6b39.
//
// Solidity: event XCallPaused()
func (_OmniPortal *OmniPortalFilterer) FilterXCallPaused(opts *bind.FilterOpts) (*OmniPortalXCallPausedIterator, error) {

	logs, sub, err := _OmniPortal.contract.FilterLogs(opts, "XCallPaused")
	if err != nil {
		return nil, err
	}
	return &OmniPortalXCallPausedIterator{contract: _OmniPortal.contract, event: "XCallPaused", logs: logs, sub: sub}, nil
}

// WatchXCallPaused is a free log subscription operation binding the contract event 0x5f335a4032d4cfb6aca7835b0c2225f36d4d9eaa4ed43ee59ed537e02dff6b39.
//
// Solidity: event XCallPaused()
func (_OmniPortal *OmniPortalFilterer) WatchXCallPaused(opts *bind.WatchOpts, sink chan<- *OmniPortalXCallPaused) (event.Subscription, error) {

	logs, sub, err := _OmniPortal.contract.WatchLogs(opts, "XCallPaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OmniPortalXCallPaused)
				if err := _OmniPortal.contract.UnpackLog(event, "XCallPaused", log); err != nil {
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
func (_OmniPortal *OmniPortalFilterer) ParseXCallPaused(log types.Log) (*OmniPortalXCallPaused, error) {
	event := new(OmniPortalXCallPaused)
	if err := _OmniPortal.contract.UnpackLog(event, "XCallPaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OmniPortalXCallToPausedIterator is returned from FilterXCallToPaused and is used to iterate over the raw logs and unpacked data for XCallToPaused events raised by the OmniPortal contract.
type OmniPortalXCallToPausedIterator struct {
	Event *OmniPortalXCallToPaused // Event containing the contract specifics and raw log

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
func (it *OmniPortalXCallToPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OmniPortalXCallToPaused)
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
		it.Event = new(OmniPortalXCallToPaused)
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
func (it *OmniPortalXCallToPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OmniPortalXCallToPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OmniPortalXCallToPaused represents a XCallToPaused event raised by the OmniPortal contract.
type OmniPortalXCallToPaused struct {
	ChainId uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterXCallToPaused is a free log retrieval operation binding the contract event 0xcd7910e1c5569d8433ce4ef8e5d51c1bdc03168f614b576da47dc3d2b51d033a.
//
// Solidity: event XCallToPaused(uint64 indexed chainId)
func (_OmniPortal *OmniPortalFilterer) FilterXCallToPaused(opts *bind.FilterOpts, chainId []uint64) (*OmniPortalXCallToPausedIterator, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}

	logs, sub, err := _OmniPortal.contract.FilterLogs(opts, "XCallToPaused", chainIdRule)
	if err != nil {
		return nil, err
	}
	return &OmniPortalXCallToPausedIterator{contract: _OmniPortal.contract, event: "XCallToPaused", logs: logs, sub: sub}, nil
}

// WatchXCallToPaused is a free log subscription operation binding the contract event 0xcd7910e1c5569d8433ce4ef8e5d51c1bdc03168f614b576da47dc3d2b51d033a.
//
// Solidity: event XCallToPaused(uint64 indexed chainId)
func (_OmniPortal *OmniPortalFilterer) WatchXCallToPaused(opts *bind.WatchOpts, sink chan<- *OmniPortalXCallToPaused, chainId []uint64) (event.Subscription, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}

	logs, sub, err := _OmniPortal.contract.WatchLogs(opts, "XCallToPaused", chainIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OmniPortalXCallToPaused)
				if err := _OmniPortal.contract.UnpackLog(event, "XCallToPaused", log); err != nil {
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
func (_OmniPortal *OmniPortalFilterer) ParseXCallToPaused(log types.Log) (*OmniPortalXCallToPaused, error) {
	event := new(OmniPortalXCallToPaused)
	if err := _OmniPortal.contract.UnpackLog(event, "XCallToPaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OmniPortalXCallToUnpausedIterator is returned from FilterXCallToUnpaused and is used to iterate over the raw logs and unpacked data for XCallToUnpaused events raised by the OmniPortal contract.
type OmniPortalXCallToUnpausedIterator struct {
	Event *OmniPortalXCallToUnpaused // Event containing the contract specifics and raw log

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
func (it *OmniPortalXCallToUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OmniPortalXCallToUnpaused)
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
		it.Event = new(OmniPortalXCallToUnpaused)
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
func (it *OmniPortalXCallToUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OmniPortalXCallToUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OmniPortalXCallToUnpaused represents a XCallToUnpaused event raised by the OmniPortal contract.
type OmniPortalXCallToUnpaused struct {
	ChainId uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterXCallToUnpaused is a free log retrieval operation binding the contract event 0x1ed9223556fb0971076c30172f1f00630efd313b6a05290a562aef95928e7125.
//
// Solidity: event XCallToUnpaused(uint64 indexed chainId)
func (_OmniPortal *OmniPortalFilterer) FilterXCallToUnpaused(opts *bind.FilterOpts, chainId []uint64) (*OmniPortalXCallToUnpausedIterator, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}

	logs, sub, err := _OmniPortal.contract.FilterLogs(opts, "XCallToUnpaused", chainIdRule)
	if err != nil {
		return nil, err
	}
	return &OmniPortalXCallToUnpausedIterator{contract: _OmniPortal.contract, event: "XCallToUnpaused", logs: logs, sub: sub}, nil
}

// WatchXCallToUnpaused is a free log subscription operation binding the contract event 0x1ed9223556fb0971076c30172f1f00630efd313b6a05290a562aef95928e7125.
//
// Solidity: event XCallToUnpaused(uint64 indexed chainId)
func (_OmniPortal *OmniPortalFilterer) WatchXCallToUnpaused(opts *bind.WatchOpts, sink chan<- *OmniPortalXCallToUnpaused, chainId []uint64) (event.Subscription, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}

	logs, sub, err := _OmniPortal.contract.WatchLogs(opts, "XCallToUnpaused", chainIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OmniPortalXCallToUnpaused)
				if err := _OmniPortal.contract.UnpackLog(event, "XCallToUnpaused", log); err != nil {
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
func (_OmniPortal *OmniPortalFilterer) ParseXCallToUnpaused(log types.Log) (*OmniPortalXCallToUnpaused, error) {
	event := new(OmniPortalXCallToUnpaused)
	if err := _OmniPortal.contract.UnpackLog(event, "XCallToUnpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OmniPortalXCallUnpausedIterator is returned from FilterXCallUnpaused and is used to iterate over the raw logs and unpacked data for XCallUnpaused events raised by the OmniPortal contract.
type OmniPortalXCallUnpausedIterator struct {
	Event *OmniPortalXCallUnpaused // Event containing the contract specifics and raw log

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
func (it *OmniPortalXCallUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OmniPortalXCallUnpaused)
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
		it.Event = new(OmniPortalXCallUnpaused)
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
func (it *OmniPortalXCallUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OmniPortalXCallUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OmniPortalXCallUnpaused represents a XCallUnpaused event raised by the OmniPortal contract.
type OmniPortalXCallUnpaused struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterXCallUnpaused is a free log retrieval operation binding the contract event 0x4c48c7b71557216a3192842746bdfc381f98d7536d9eb1c6764f3b45e6794827.
//
// Solidity: event XCallUnpaused()
func (_OmniPortal *OmniPortalFilterer) FilterXCallUnpaused(opts *bind.FilterOpts) (*OmniPortalXCallUnpausedIterator, error) {

	logs, sub, err := _OmniPortal.contract.FilterLogs(opts, "XCallUnpaused")
	if err != nil {
		return nil, err
	}
	return &OmniPortalXCallUnpausedIterator{contract: _OmniPortal.contract, event: "XCallUnpaused", logs: logs, sub: sub}, nil
}

// WatchXCallUnpaused is a free log subscription operation binding the contract event 0x4c48c7b71557216a3192842746bdfc381f98d7536d9eb1c6764f3b45e6794827.
//
// Solidity: event XCallUnpaused()
func (_OmniPortal *OmniPortalFilterer) WatchXCallUnpaused(opts *bind.WatchOpts, sink chan<- *OmniPortalXCallUnpaused) (event.Subscription, error) {

	logs, sub, err := _OmniPortal.contract.WatchLogs(opts, "XCallUnpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OmniPortalXCallUnpaused)
				if err := _OmniPortal.contract.UnpackLog(event, "XCallUnpaused", log); err != nil {
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
func (_OmniPortal *OmniPortalFilterer) ParseXCallUnpaused(log types.Log) (*OmniPortalXCallUnpaused, error) {
	event := new(OmniPortalXCallUnpaused)
	if err := _OmniPortal.contract.UnpackLog(event, "XCallUnpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OmniPortalXMsgIterator is returned from FilterXMsg and is used to iterate over the raw logs and unpacked data for XMsg events raised by the OmniPortal contract.
type OmniPortalXMsgIterator struct {
	Event *OmniPortalXMsg // Event containing the contract specifics and raw log

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
func (it *OmniPortalXMsgIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OmniPortalXMsg)
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
		it.Event = new(OmniPortalXMsg)
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
func (it *OmniPortalXMsgIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OmniPortalXMsgIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OmniPortalXMsg represents a XMsg event raised by the OmniPortal contract.
type OmniPortalXMsg struct {
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
func (_OmniPortal *OmniPortalFilterer) FilterXMsg(opts *bind.FilterOpts, destChainId []uint64, shardId []uint64, offset []uint64) (*OmniPortalXMsgIterator, error) {

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

	logs, sub, err := _OmniPortal.contract.FilterLogs(opts, "XMsg", destChainIdRule, shardIdRule, offsetRule)
	if err != nil {
		return nil, err
	}
	return &OmniPortalXMsgIterator{contract: _OmniPortal.contract, event: "XMsg", logs: logs, sub: sub}, nil
}

// WatchXMsg is a free log subscription operation binding the contract event 0xb7c8eb9d7a7fbcdab809ab7b8a7c41701eb3115e3fe99d30ff490d8552f72bfa.
//
// Solidity: event XMsg(uint64 indexed destChainId, uint64 indexed shardId, uint64 indexed offset, address sender, address to, bytes data, uint64 gasLimit, uint256 fees)
func (_OmniPortal *OmniPortalFilterer) WatchXMsg(opts *bind.WatchOpts, sink chan<- *OmniPortalXMsg, destChainId []uint64, shardId []uint64, offset []uint64) (event.Subscription, error) {

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

	logs, sub, err := _OmniPortal.contract.WatchLogs(opts, "XMsg", destChainIdRule, shardIdRule, offsetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OmniPortalXMsg)
				if err := _OmniPortal.contract.UnpackLog(event, "XMsg", log); err != nil {
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
func (_OmniPortal *OmniPortalFilterer) ParseXMsg(log types.Log) (*OmniPortalXMsg, error) {
	event := new(OmniPortalXMsg)
	if err := _OmniPortal.contract.UnpackLog(event, "XMsg", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OmniPortalXMsgMaxDataSizeSetIterator is returned from FilterXMsgMaxDataSizeSet and is used to iterate over the raw logs and unpacked data for XMsgMaxDataSizeSet events raised by the OmniPortal contract.
type OmniPortalXMsgMaxDataSizeSetIterator struct {
	Event *OmniPortalXMsgMaxDataSizeSet // Event containing the contract specifics and raw log

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
func (it *OmniPortalXMsgMaxDataSizeSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OmniPortalXMsgMaxDataSizeSet)
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
		it.Event = new(OmniPortalXMsgMaxDataSizeSet)
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
func (it *OmniPortalXMsgMaxDataSizeSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OmniPortalXMsgMaxDataSizeSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OmniPortalXMsgMaxDataSizeSet represents a XMsgMaxDataSizeSet event raised by the OmniPortal contract.
type OmniPortalXMsgMaxDataSizeSet struct {
	Size uint16
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterXMsgMaxDataSizeSet is a free log retrieval operation binding the contract event 0x65923e04419dc810d0ea08a94a7f608d4c4d949818d95c3788f895e575dd2064.
//
// Solidity: event XMsgMaxDataSizeSet(uint16 size)
func (_OmniPortal *OmniPortalFilterer) FilterXMsgMaxDataSizeSet(opts *bind.FilterOpts) (*OmniPortalXMsgMaxDataSizeSetIterator, error) {

	logs, sub, err := _OmniPortal.contract.FilterLogs(opts, "XMsgMaxDataSizeSet")
	if err != nil {
		return nil, err
	}
	return &OmniPortalXMsgMaxDataSizeSetIterator{contract: _OmniPortal.contract, event: "XMsgMaxDataSizeSet", logs: logs, sub: sub}, nil
}

// WatchXMsgMaxDataSizeSet is a free log subscription operation binding the contract event 0x65923e04419dc810d0ea08a94a7f608d4c4d949818d95c3788f895e575dd2064.
//
// Solidity: event XMsgMaxDataSizeSet(uint16 size)
func (_OmniPortal *OmniPortalFilterer) WatchXMsgMaxDataSizeSet(opts *bind.WatchOpts, sink chan<- *OmniPortalXMsgMaxDataSizeSet) (event.Subscription, error) {

	logs, sub, err := _OmniPortal.contract.WatchLogs(opts, "XMsgMaxDataSizeSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OmniPortalXMsgMaxDataSizeSet)
				if err := _OmniPortal.contract.UnpackLog(event, "XMsgMaxDataSizeSet", log); err != nil {
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
func (_OmniPortal *OmniPortalFilterer) ParseXMsgMaxDataSizeSet(log types.Log) (*OmniPortalXMsgMaxDataSizeSet, error) {
	event := new(OmniPortalXMsgMaxDataSizeSet)
	if err := _OmniPortal.contract.UnpackLog(event, "XMsgMaxDataSizeSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OmniPortalXMsgMaxGasLimitSetIterator is returned from FilterXMsgMaxGasLimitSet and is used to iterate over the raw logs and unpacked data for XMsgMaxGasLimitSet events raised by the OmniPortal contract.
type OmniPortalXMsgMaxGasLimitSetIterator struct {
	Event *OmniPortalXMsgMaxGasLimitSet // Event containing the contract specifics and raw log

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
func (it *OmniPortalXMsgMaxGasLimitSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OmniPortalXMsgMaxGasLimitSet)
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
		it.Event = new(OmniPortalXMsgMaxGasLimitSet)
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
func (it *OmniPortalXMsgMaxGasLimitSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OmniPortalXMsgMaxGasLimitSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OmniPortalXMsgMaxGasLimitSet represents a XMsgMaxGasLimitSet event raised by the OmniPortal contract.
type OmniPortalXMsgMaxGasLimitSet struct {
	GasLimit uint64
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterXMsgMaxGasLimitSet is a free log retrieval operation binding the contract event 0x1153561ac5effc2926ba6c612f86a397c997bc43dfbfc718da08065be0c5fe4d.
//
// Solidity: event XMsgMaxGasLimitSet(uint64 gasLimit)
func (_OmniPortal *OmniPortalFilterer) FilterXMsgMaxGasLimitSet(opts *bind.FilterOpts) (*OmniPortalXMsgMaxGasLimitSetIterator, error) {

	logs, sub, err := _OmniPortal.contract.FilterLogs(opts, "XMsgMaxGasLimitSet")
	if err != nil {
		return nil, err
	}
	return &OmniPortalXMsgMaxGasLimitSetIterator{contract: _OmniPortal.contract, event: "XMsgMaxGasLimitSet", logs: logs, sub: sub}, nil
}

// WatchXMsgMaxGasLimitSet is a free log subscription operation binding the contract event 0x1153561ac5effc2926ba6c612f86a397c997bc43dfbfc718da08065be0c5fe4d.
//
// Solidity: event XMsgMaxGasLimitSet(uint64 gasLimit)
func (_OmniPortal *OmniPortalFilterer) WatchXMsgMaxGasLimitSet(opts *bind.WatchOpts, sink chan<- *OmniPortalXMsgMaxGasLimitSet) (event.Subscription, error) {

	logs, sub, err := _OmniPortal.contract.WatchLogs(opts, "XMsgMaxGasLimitSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OmniPortalXMsgMaxGasLimitSet)
				if err := _OmniPortal.contract.UnpackLog(event, "XMsgMaxGasLimitSet", log); err != nil {
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
func (_OmniPortal *OmniPortalFilterer) ParseXMsgMaxGasLimitSet(log types.Log) (*OmniPortalXMsgMaxGasLimitSet, error) {
	event := new(OmniPortalXMsgMaxGasLimitSet)
	if err := _OmniPortal.contract.UnpackLog(event, "XMsgMaxGasLimitSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OmniPortalXMsgMinGasLimitSetIterator is returned from FilterXMsgMinGasLimitSet and is used to iterate over the raw logs and unpacked data for XMsgMinGasLimitSet events raised by the OmniPortal contract.
type OmniPortalXMsgMinGasLimitSetIterator struct {
	Event *OmniPortalXMsgMinGasLimitSet // Event containing the contract specifics and raw log

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
func (it *OmniPortalXMsgMinGasLimitSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OmniPortalXMsgMinGasLimitSet)
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
		it.Event = new(OmniPortalXMsgMinGasLimitSet)
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
func (it *OmniPortalXMsgMinGasLimitSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OmniPortalXMsgMinGasLimitSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OmniPortalXMsgMinGasLimitSet represents a XMsgMinGasLimitSet event raised by the OmniPortal contract.
type OmniPortalXMsgMinGasLimitSet struct {
	GasLimit uint64
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterXMsgMinGasLimitSet is a free log retrieval operation binding the contract event 0x8c852a6291aa436654b167353bca4a4b0c3d024c7562cb5082e7c869bddabf3e.
//
// Solidity: event XMsgMinGasLimitSet(uint64 gasLimit)
func (_OmniPortal *OmniPortalFilterer) FilterXMsgMinGasLimitSet(opts *bind.FilterOpts) (*OmniPortalXMsgMinGasLimitSetIterator, error) {

	logs, sub, err := _OmniPortal.contract.FilterLogs(opts, "XMsgMinGasLimitSet")
	if err != nil {
		return nil, err
	}
	return &OmniPortalXMsgMinGasLimitSetIterator{contract: _OmniPortal.contract, event: "XMsgMinGasLimitSet", logs: logs, sub: sub}, nil
}

// WatchXMsgMinGasLimitSet is a free log subscription operation binding the contract event 0x8c852a6291aa436654b167353bca4a4b0c3d024c7562cb5082e7c869bddabf3e.
//
// Solidity: event XMsgMinGasLimitSet(uint64 gasLimit)
func (_OmniPortal *OmniPortalFilterer) WatchXMsgMinGasLimitSet(opts *bind.WatchOpts, sink chan<- *OmniPortalXMsgMinGasLimitSet) (event.Subscription, error) {

	logs, sub, err := _OmniPortal.contract.WatchLogs(opts, "XMsgMinGasLimitSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OmniPortalXMsgMinGasLimitSet)
				if err := _OmniPortal.contract.UnpackLog(event, "XMsgMinGasLimitSet", log); err != nil {
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
func (_OmniPortal *OmniPortalFilterer) ParseXMsgMinGasLimitSet(log types.Log) (*OmniPortalXMsgMinGasLimitSet, error) {
	event := new(OmniPortalXMsgMinGasLimitSet)
	if err := _OmniPortal.contract.UnpackLog(event, "XMsgMinGasLimitSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OmniPortalXReceiptIterator is returned from FilterXReceipt and is used to iterate over the raw logs and unpacked data for XReceipt events raised by the OmniPortal contract.
type OmniPortalXReceiptIterator struct {
	Event *OmniPortalXReceipt // Event containing the contract specifics and raw log

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
func (it *OmniPortalXReceiptIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OmniPortalXReceipt)
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
		it.Event = new(OmniPortalXReceipt)
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
func (it *OmniPortalXReceiptIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OmniPortalXReceiptIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OmniPortalXReceipt represents a XReceipt event raised by the OmniPortal contract.
type OmniPortalXReceipt struct {
	SourceChainId uint64
	ShardId       uint64
	Offset        uint64
	GasUsed       *big.Int
	Relayer       common.Address
	Success       bool
	Error         []byte
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterXReceipt is a free log retrieval operation binding the contract event 0x8277cab1f0fa69b34674f64a7d43f242b0bacece6f5b7e8652f1e0d88a9b873b.
//
// Solidity: event XReceipt(uint64 indexed sourceChainId, uint64 indexed shardId, uint64 indexed offset, uint256 gasUsed, address relayer, bool success, bytes error)
func (_OmniPortal *OmniPortalFilterer) FilterXReceipt(opts *bind.FilterOpts, sourceChainId []uint64, shardId []uint64, offset []uint64) (*OmniPortalXReceiptIterator, error) {

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

	logs, sub, err := _OmniPortal.contract.FilterLogs(opts, "XReceipt", sourceChainIdRule, shardIdRule, offsetRule)
	if err != nil {
		return nil, err
	}
	return &OmniPortalXReceiptIterator{contract: _OmniPortal.contract, event: "XReceipt", logs: logs, sub: sub}, nil
}

// WatchXReceipt is a free log subscription operation binding the contract event 0x8277cab1f0fa69b34674f64a7d43f242b0bacece6f5b7e8652f1e0d88a9b873b.
//
// Solidity: event XReceipt(uint64 indexed sourceChainId, uint64 indexed shardId, uint64 indexed offset, uint256 gasUsed, address relayer, bool success, bytes error)
func (_OmniPortal *OmniPortalFilterer) WatchXReceipt(opts *bind.WatchOpts, sink chan<- *OmniPortalXReceipt, sourceChainId []uint64, shardId []uint64, offset []uint64) (event.Subscription, error) {

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

	logs, sub, err := _OmniPortal.contract.WatchLogs(opts, "XReceipt", sourceChainIdRule, shardIdRule, offsetRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OmniPortalXReceipt)
				if err := _OmniPortal.contract.UnpackLog(event, "XReceipt", log); err != nil {
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
// Solidity: event XReceipt(uint64 indexed sourceChainId, uint64 indexed shardId, uint64 indexed offset, uint256 gasUsed, address relayer, bool success, bytes error)
func (_OmniPortal *OmniPortalFilterer) ParseXReceipt(log types.Log) (*OmniPortalXReceipt, error) {
	event := new(OmniPortalXReceipt)
	if err := _OmniPortal.contract.UnpackLog(event, "XReceipt", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OmniPortalXReceiptMaxErrorSizeSetIterator is returned from FilterXReceiptMaxErrorSizeSet and is used to iterate over the raw logs and unpacked data for XReceiptMaxErrorSizeSet events raised by the OmniPortal contract.
type OmniPortalXReceiptMaxErrorSizeSetIterator struct {
	Event *OmniPortalXReceiptMaxErrorSizeSet // Event containing the contract specifics and raw log

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
func (it *OmniPortalXReceiptMaxErrorSizeSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OmniPortalXReceiptMaxErrorSizeSet)
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
		it.Event = new(OmniPortalXReceiptMaxErrorSizeSet)
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
func (it *OmniPortalXReceiptMaxErrorSizeSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OmniPortalXReceiptMaxErrorSizeSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OmniPortalXReceiptMaxErrorSizeSet represents a XReceiptMaxErrorSizeSet event raised by the OmniPortal contract.
type OmniPortalXReceiptMaxErrorSizeSet struct {
	Size uint16
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterXReceiptMaxErrorSizeSet is a free log retrieval operation binding the contract event 0x620bbea084306b66a8cc6b5b63830d6b3874f9d2438914e259ffd5065c33f7b0.
//
// Solidity: event XReceiptMaxErrorSizeSet(uint16 size)
func (_OmniPortal *OmniPortalFilterer) FilterXReceiptMaxErrorSizeSet(opts *bind.FilterOpts) (*OmniPortalXReceiptMaxErrorSizeSetIterator, error) {

	logs, sub, err := _OmniPortal.contract.FilterLogs(opts, "XReceiptMaxErrorSizeSet")
	if err != nil {
		return nil, err
	}
	return &OmniPortalXReceiptMaxErrorSizeSetIterator{contract: _OmniPortal.contract, event: "XReceiptMaxErrorSizeSet", logs: logs, sub: sub}, nil
}

// WatchXReceiptMaxErrorSizeSet is a free log subscription operation binding the contract event 0x620bbea084306b66a8cc6b5b63830d6b3874f9d2438914e259ffd5065c33f7b0.
//
// Solidity: event XReceiptMaxErrorSizeSet(uint16 size)
func (_OmniPortal *OmniPortalFilterer) WatchXReceiptMaxErrorSizeSet(opts *bind.WatchOpts, sink chan<- *OmniPortalXReceiptMaxErrorSizeSet) (event.Subscription, error) {

	logs, sub, err := _OmniPortal.contract.WatchLogs(opts, "XReceiptMaxErrorSizeSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OmniPortalXReceiptMaxErrorSizeSet)
				if err := _OmniPortal.contract.UnpackLog(event, "XReceiptMaxErrorSizeSet", log); err != nil {
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
func (_OmniPortal *OmniPortalFilterer) ParseXReceiptMaxErrorSizeSet(log types.Log) (*OmniPortalXReceiptMaxErrorSizeSet, error) {
	event := new(OmniPortalXReceiptMaxErrorSizeSet)
	if err := _OmniPortal.contract.UnpackLog(event, "XReceiptMaxErrorSizeSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OmniPortalXSubValsetCutoffSetIterator is returned from FilterXSubValsetCutoffSet and is used to iterate over the raw logs and unpacked data for XSubValsetCutoffSet events raised by the OmniPortal contract.
type OmniPortalXSubValsetCutoffSetIterator struct {
	Event *OmniPortalXSubValsetCutoffSet // Event containing the contract specifics and raw log

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
func (it *OmniPortalXSubValsetCutoffSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OmniPortalXSubValsetCutoffSet)
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
		it.Event = new(OmniPortalXSubValsetCutoffSet)
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
func (it *OmniPortalXSubValsetCutoffSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OmniPortalXSubValsetCutoffSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OmniPortalXSubValsetCutoffSet represents a XSubValsetCutoffSet event raised by the OmniPortal contract.
type OmniPortalXSubValsetCutoffSet struct {
	Cutoff uint8
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterXSubValsetCutoffSet is a free log retrieval operation binding the contract event 0x1683dc51426224f6e37a3b41dd5849e2db1bfe22366d1d913fa0ef6f757e828f.
//
// Solidity: event XSubValsetCutoffSet(uint8 cutoff)
func (_OmniPortal *OmniPortalFilterer) FilterXSubValsetCutoffSet(opts *bind.FilterOpts) (*OmniPortalXSubValsetCutoffSetIterator, error) {

	logs, sub, err := _OmniPortal.contract.FilterLogs(opts, "XSubValsetCutoffSet")
	if err != nil {
		return nil, err
	}
	return &OmniPortalXSubValsetCutoffSetIterator{contract: _OmniPortal.contract, event: "XSubValsetCutoffSet", logs: logs, sub: sub}, nil
}

// WatchXSubValsetCutoffSet is a free log subscription operation binding the contract event 0x1683dc51426224f6e37a3b41dd5849e2db1bfe22366d1d913fa0ef6f757e828f.
//
// Solidity: event XSubValsetCutoffSet(uint8 cutoff)
func (_OmniPortal *OmniPortalFilterer) WatchXSubValsetCutoffSet(opts *bind.WatchOpts, sink chan<- *OmniPortalXSubValsetCutoffSet) (event.Subscription, error) {

	logs, sub, err := _OmniPortal.contract.WatchLogs(opts, "XSubValsetCutoffSet")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OmniPortalXSubValsetCutoffSet)
				if err := _OmniPortal.contract.UnpackLog(event, "XSubValsetCutoffSet", log); err != nil {
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
func (_OmniPortal *OmniPortalFilterer) ParseXSubValsetCutoffSet(log types.Log) (*OmniPortalXSubValsetCutoffSet, error) {
	event := new(OmniPortalXSubValsetCutoffSet)
	if err := _OmniPortal.contract.UnpackLog(event, "XSubValsetCutoffSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OmniPortalXSubmitFromPausedIterator is returned from FilterXSubmitFromPaused and is used to iterate over the raw logs and unpacked data for XSubmitFromPaused events raised by the OmniPortal contract.
type OmniPortalXSubmitFromPausedIterator struct {
	Event *OmniPortalXSubmitFromPaused // Event containing the contract specifics and raw log

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
func (it *OmniPortalXSubmitFromPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OmniPortalXSubmitFromPaused)
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
		it.Event = new(OmniPortalXSubmitFromPaused)
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
func (it *OmniPortalXSubmitFromPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OmniPortalXSubmitFromPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OmniPortalXSubmitFromPaused represents a XSubmitFromPaused event raised by the OmniPortal contract.
type OmniPortalXSubmitFromPaused struct {
	ChainId uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterXSubmitFromPaused is a free log retrieval operation binding the contract event 0xab78810a0515df65f9f10bfbcb92d03d5df71d9fd3b9414e9ad831a5117d6daa.
//
// Solidity: event XSubmitFromPaused(uint64 indexed chainId)
func (_OmniPortal *OmniPortalFilterer) FilterXSubmitFromPaused(opts *bind.FilterOpts, chainId []uint64) (*OmniPortalXSubmitFromPausedIterator, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}

	logs, sub, err := _OmniPortal.contract.FilterLogs(opts, "XSubmitFromPaused", chainIdRule)
	if err != nil {
		return nil, err
	}
	return &OmniPortalXSubmitFromPausedIterator{contract: _OmniPortal.contract, event: "XSubmitFromPaused", logs: logs, sub: sub}, nil
}

// WatchXSubmitFromPaused is a free log subscription operation binding the contract event 0xab78810a0515df65f9f10bfbcb92d03d5df71d9fd3b9414e9ad831a5117d6daa.
//
// Solidity: event XSubmitFromPaused(uint64 indexed chainId)
func (_OmniPortal *OmniPortalFilterer) WatchXSubmitFromPaused(opts *bind.WatchOpts, sink chan<- *OmniPortalXSubmitFromPaused, chainId []uint64) (event.Subscription, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}

	logs, sub, err := _OmniPortal.contract.WatchLogs(opts, "XSubmitFromPaused", chainIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OmniPortalXSubmitFromPaused)
				if err := _OmniPortal.contract.UnpackLog(event, "XSubmitFromPaused", log); err != nil {
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
func (_OmniPortal *OmniPortalFilterer) ParseXSubmitFromPaused(log types.Log) (*OmniPortalXSubmitFromPaused, error) {
	event := new(OmniPortalXSubmitFromPaused)
	if err := _OmniPortal.contract.UnpackLog(event, "XSubmitFromPaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OmniPortalXSubmitFromUnpausedIterator is returned from FilterXSubmitFromUnpaused and is used to iterate over the raw logs and unpacked data for XSubmitFromUnpaused events raised by the OmniPortal contract.
type OmniPortalXSubmitFromUnpausedIterator struct {
	Event *OmniPortalXSubmitFromUnpaused // Event containing the contract specifics and raw log

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
func (it *OmniPortalXSubmitFromUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OmniPortalXSubmitFromUnpaused)
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
		it.Event = new(OmniPortalXSubmitFromUnpaused)
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
func (it *OmniPortalXSubmitFromUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OmniPortalXSubmitFromUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OmniPortalXSubmitFromUnpaused represents a XSubmitFromUnpaused event raised by the OmniPortal contract.
type OmniPortalXSubmitFromUnpaused struct {
	ChainId uint64
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterXSubmitFromUnpaused is a free log retrieval operation binding the contract event 0xc551305d9bd408be4327b7f8aba28b04ccf6b6c76925392d195ecf9cc764294d.
//
// Solidity: event XSubmitFromUnpaused(uint64 indexed chainId)
func (_OmniPortal *OmniPortalFilterer) FilterXSubmitFromUnpaused(opts *bind.FilterOpts, chainId []uint64) (*OmniPortalXSubmitFromUnpausedIterator, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}

	logs, sub, err := _OmniPortal.contract.FilterLogs(opts, "XSubmitFromUnpaused", chainIdRule)
	if err != nil {
		return nil, err
	}
	return &OmniPortalXSubmitFromUnpausedIterator{contract: _OmniPortal.contract, event: "XSubmitFromUnpaused", logs: logs, sub: sub}, nil
}

// WatchXSubmitFromUnpaused is a free log subscription operation binding the contract event 0xc551305d9bd408be4327b7f8aba28b04ccf6b6c76925392d195ecf9cc764294d.
//
// Solidity: event XSubmitFromUnpaused(uint64 indexed chainId)
func (_OmniPortal *OmniPortalFilterer) WatchXSubmitFromUnpaused(opts *bind.WatchOpts, sink chan<- *OmniPortalXSubmitFromUnpaused, chainId []uint64) (event.Subscription, error) {

	var chainIdRule []interface{}
	for _, chainIdItem := range chainId {
		chainIdRule = append(chainIdRule, chainIdItem)
	}

	logs, sub, err := _OmniPortal.contract.WatchLogs(opts, "XSubmitFromUnpaused", chainIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OmniPortalXSubmitFromUnpaused)
				if err := _OmniPortal.contract.UnpackLog(event, "XSubmitFromUnpaused", log); err != nil {
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
func (_OmniPortal *OmniPortalFilterer) ParseXSubmitFromUnpaused(log types.Log) (*OmniPortalXSubmitFromUnpaused, error) {
	event := new(OmniPortalXSubmitFromUnpaused)
	if err := _OmniPortal.contract.UnpackLog(event, "XSubmitFromUnpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OmniPortalXSubmitPausedIterator is returned from FilterXSubmitPaused and is used to iterate over the raw logs and unpacked data for XSubmitPaused events raised by the OmniPortal contract.
type OmniPortalXSubmitPausedIterator struct {
	Event *OmniPortalXSubmitPaused // Event containing the contract specifics and raw log

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
func (it *OmniPortalXSubmitPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OmniPortalXSubmitPaused)
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
		it.Event = new(OmniPortalXSubmitPaused)
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
func (it *OmniPortalXSubmitPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OmniPortalXSubmitPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OmniPortalXSubmitPaused represents a XSubmitPaused event raised by the OmniPortal contract.
type OmniPortalXSubmitPaused struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterXSubmitPaused is a free log retrieval operation binding the contract event 0x3d0f9c56dac46156a2db0aa09ee7804770ad9fc9549d21023164f22d69475ed8.
//
// Solidity: event XSubmitPaused()
func (_OmniPortal *OmniPortalFilterer) FilterXSubmitPaused(opts *bind.FilterOpts) (*OmniPortalXSubmitPausedIterator, error) {

	logs, sub, err := _OmniPortal.contract.FilterLogs(opts, "XSubmitPaused")
	if err != nil {
		return nil, err
	}
	return &OmniPortalXSubmitPausedIterator{contract: _OmniPortal.contract, event: "XSubmitPaused", logs: logs, sub: sub}, nil
}

// WatchXSubmitPaused is a free log subscription operation binding the contract event 0x3d0f9c56dac46156a2db0aa09ee7804770ad9fc9549d21023164f22d69475ed8.
//
// Solidity: event XSubmitPaused()
func (_OmniPortal *OmniPortalFilterer) WatchXSubmitPaused(opts *bind.WatchOpts, sink chan<- *OmniPortalXSubmitPaused) (event.Subscription, error) {

	logs, sub, err := _OmniPortal.contract.WatchLogs(opts, "XSubmitPaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OmniPortalXSubmitPaused)
				if err := _OmniPortal.contract.UnpackLog(event, "XSubmitPaused", log); err != nil {
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
func (_OmniPortal *OmniPortalFilterer) ParseXSubmitPaused(log types.Log) (*OmniPortalXSubmitPaused, error) {
	event := new(OmniPortalXSubmitPaused)
	if err := _OmniPortal.contract.UnpackLog(event, "XSubmitPaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// OmniPortalXSubmitUnpausedIterator is returned from FilterXSubmitUnpaused and is used to iterate over the raw logs and unpacked data for XSubmitUnpaused events raised by the OmniPortal contract.
type OmniPortalXSubmitUnpausedIterator struct {
	Event *OmniPortalXSubmitUnpaused // Event containing the contract specifics and raw log

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
func (it *OmniPortalXSubmitUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(OmniPortalXSubmitUnpaused)
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
		it.Event = new(OmniPortalXSubmitUnpaused)
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
func (it *OmniPortalXSubmitUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *OmniPortalXSubmitUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// OmniPortalXSubmitUnpaused represents a XSubmitUnpaused event raised by the OmniPortal contract.
type OmniPortalXSubmitUnpaused struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterXSubmitUnpaused is a free log retrieval operation binding the contract event 0x2cb9d71d4c31860b70e9b707c69aa2f5953e03474f00cfcfff205c4745f82875.
//
// Solidity: event XSubmitUnpaused()
func (_OmniPortal *OmniPortalFilterer) FilterXSubmitUnpaused(opts *bind.FilterOpts) (*OmniPortalXSubmitUnpausedIterator, error) {

	logs, sub, err := _OmniPortal.contract.FilterLogs(opts, "XSubmitUnpaused")
	if err != nil {
		return nil, err
	}
	return &OmniPortalXSubmitUnpausedIterator{contract: _OmniPortal.contract, event: "XSubmitUnpaused", logs: logs, sub: sub}, nil
}

// WatchXSubmitUnpaused is a free log subscription operation binding the contract event 0x2cb9d71d4c31860b70e9b707c69aa2f5953e03474f00cfcfff205c4745f82875.
//
// Solidity: event XSubmitUnpaused()
func (_OmniPortal *OmniPortalFilterer) WatchXSubmitUnpaused(opts *bind.WatchOpts, sink chan<- *OmniPortalXSubmitUnpaused) (event.Subscription, error) {

	logs, sub, err := _OmniPortal.contract.WatchLogs(opts, "XSubmitUnpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(OmniPortalXSubmitUnpaused)
				if err := _OmniPortal.contract.UnpackLog(event, "XSubmitUnpaused", log); err != nil {
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
func (_OmniPortal *OmniPortalFilterer) ParseXSubmitUnpaused(log types.Log) (*OmniPortalXSubmitUnpaused, error) {
	event := new(OmniPortalXSubmitUnpaused)
	if err := _OmniPortal.contract.UnpackLog(event, "XSubmitUnpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
