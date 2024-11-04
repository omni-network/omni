// This file copies Receipt type from ethereum-optimism/op-geth/core/types/receipt.go
// It omits everything else, refactoring where needed.

package ethclient

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
)

//go:generate go run github.com/fjl/gencodec -type Receipt -field-override receiptMarshaling -out gen_receipt_json.go

const (
	// ReceiptStatusFailed is the status code of a transaction if execution failed.
	ReceiptStatusFailed = uint64(0)

	// ReceiptStatusSuccessful is the status code of a transaction if execution succeeded.
	ReceiptStatusSuccessful = uint64(1)
)

// Receipt represents the results of a transaction.
type Receipt struct {
	// Consensus fields: These fields are defined by the Yellow Paper
	Type              uint8        `json:"type,omitempty"`
	PostState         []byte       `json:"root"`
	Status            uint64       `json:"status"`
	CumulativeGasUsed uint64       `gencodec:"required"   json:"cumulativeGasUsed"`
	Bloom             types.Bloom  `gencodec:"required"   json:"logsBloom"`
	Logs              []*types.Log `gencodec:"required"   json:"logs"`

	// Implementation fields: These fields are added by geth when processing a transaction or retrieving a receipt.
	// gencodec annotated fields: these are stored in the chain database.
	TxHash            common.Hash    `gencodec:"required"           json:"transactionHash"`
	ContractAddress   common.Address `json:"contractAddress"`
	GasUsed           uint64         `gencodec:"required"           json:"gasUsed"`
	EffectiveGasPrice *big.Int       `json:"effectiveGasPrice"` // required, but tag omitted for backwards compatibility
	BlobGasUsed       uint64         `json:"blobGasUsed,omitempty"`
	BlobGasPrice      *big.Int       `json:"blobGasPrice,omitempty"`

	// DepositNonce was introduced in Regolith to store the actual nonce used by deposit transactions
	// The state transition process ensures this is only set for Regolith deposit transactions.
	DepositNonce *uint64 `json:"depositNonce,omitempty"`
	// DepositReceiptVersion was introduced in Canyon to indicate an update to how receipt hashes
	// should be computed when set. The state transition process ensures this is only set for
	// post-Canyon deposit transactions.
	DepositReceiptVersion *uint64 `json:"depositReceiptVersion,omitempty"`

	// Inclusion information: These fields provide information about the inclusion of the
	// transaction corresponding to this receipt.
	BlockHash        common.Hash `json:"blockHash,omitempty"`
	BlockNumber      *big.Int    `json:"blockNumber,omitempty"`
	TransactionIndex uint        `json:"transactionIndex"`

	// Optimism: extend receipts with L1 fee info
	L1GasPrice          *big.Int   `json:"l1GasPrice,omitempty"`          // Present from pre-bedrock. L1 Basefee after Bedrock
	L1BlobBaseFee       *big.Int   `json:"l1BlobBaseFee,omitempty"`       // Always nil prior to the Ecotone hardfork
	L1GasUsed           *big.Int   `json:"l1GasUsed,omitempty"`           // Present from pre-bedrock, deprecated as of Fjord
	L1Fee               *big.Int   `json:"l1Fee,omitempty"`               // Present from pre-bedrock
	FeeScalar           *big.Float `json:"l1FeeScalar,omitempty"`         // Present from pre-bedrock to Ecotone. Nil after Ecotone
	L1BaseFeeScalar     *uint64    `json:"l1BaseFeeScalar,omitempty"`     // Always nil prior to the Ecotone hardfork
	L1BlobBaseFeeScalar *uint64    `json:"l1BlobBaseFeeScalar,omitempty"` // Always nil prior to the Ecotone hardfork
}

type receiptMarshaling struct {
	Type              hexutil.Uint64
	PostState         hexutil.Bytes
	Status            hexutil.Uint64
	CumulativeGasUsed hexutil.Uint64
	GasUsed           hexutil.Uint64
	EffectiveGasPrice *hexutil.Big
	BlobGasUsed       hexutil.Uint64
	BlobGasPrice      *hexutil.Big
	BlockNumber       *hexutil.Big
	TransactionIndex  hexutil.Uint

	// Optimism
	L1GasPrice            *hexutil.Big
	L1BlobBaseFee         *hexutil.Big
	L1GasUsed             *hexutil.Big
	L1Fee                 *hexutil.Big
	FeeScalar             *big.Float
	L1BaseFeeScalar       *hexutil.Uint64
	L1BlobBaseFeeScalar   *hexutil.Uint64
	DepositNonce          *hexutil.Uint64
	DepositReceiptVersion *hexutil.Uint64
}
