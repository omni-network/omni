package types

import (
	"math/big"

	"github.com/omni-network/omni/lib/errors"

	"github.com/ethereum/go-ethereum/beacon/engine"
	"github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
)

func PayloadToProto(p *engine.ExecutableData) (*ExecutionPayloadDeneb, error) {
	if p == nil {
		return nil, errors.New("nil payload")
	} else if p.BlobGasUsed == nil {
		return nil, errors.New("nil payload BlobGasUsed")
	} else if p.ExcessBlobGas == nil {
		return nil, errors.New("nil payload ExcessBlobGas")
	} else if p.ExecutionWitness != nil {
		return nil, errors.New("payload has ExecutionWitness")
	}

	var baseFeePerGas [32]byte
	p.BaseFeePerGas.FillBytes(baseFeePerGas[:])

	var withdrawals []Withdrawal
	for _, w := range p.Withdrawals {
		withdrawal, err := WithdrawalToProto(w)
		if err != nil {
			return nil, errors.Wrap(err, "withdrawal to proto")
		}
		withdrawals = append(withdrawals, withdrawal)
	}

	return &ExecutionPayloadDeneb{
		ParentHash:    Hash(p.ParentHash),
		FeeRecipient:  Address(p.FeeRecipient),
		StateRoot:     Hash(p.StateRoot),
		ReceiptsRoot:  Hash(p.ReceiptsRoot),
		LogsBloom:     p.LogsBloom,
		PrevRandao:    Hash(p.Random),
		BlockNumber:   p.Number,
		GasLimit:      p.GasLimit,
		GasUsed:       p.GasUsed,
		Timestamp:     p.Timestamp,
		ExtraData:     p.ExtraData,
		BaseFeePerGas: baseFeePerGas[:],
		BlockHash:     Hash(p.BlockHash),
		Transactions:  p.Transactions,
		Withdrawals:   withdrawals,
		BlobGasUsed:   *p.BlobGasUsed,
		ExcessBlobGas: *p.ExcessBlobGas,
	}, nil
}

func PayloadFromProto(p *ExecutionPayloadDeneb) (engine.ExecutableData, error) {
	if p == nil {
		return engine.ExecutableData{}, errors.New("nil payload")
	} else if len(p.BaseFeePerGas) != 32 {
		return engine.ExecutableData{}, errors.New("invalid BaseFeePerGas length")
	}

	withdrawals := make([]*ethtypes.Withdrawal, 0) // Geth requires an empty json array, not null.
	for _, w := range p.Withdrawals {
		withdrawals = append(withdrawals, WithdrawalFromProto(w))
	}

	transactions := p.Transactions
	if transactions == nil {
		// Geth requires an empty json array, not null.
		transactions = make([][]byte, 0)
	}

	// No need to verify other fields here:
	// - they are either fixed length (hash,address,uint64)
	// - or bytes (or slices of bytes) which are verified by geth.

	return engine.ExecutableData{
		ParentHash:    common.Hash(p.ParentHash),
		FeeRecipient:  common.Address(p.FeeRecipient),
		StateRoot:     common.Hash(p.StateRoot),
		ReceiptsRoot:  common.Hash(p.ReceiptsRoot),
		LogsBloom:     p.LogsBloom,
		Random:        common.Hash(p.PrevRandao),
		Number:        p.BlockNumber,
		GasLimit:      p.GasLimit,
		GasUsed:       p.GasUsed,
		Timestamp:     p.Timestamp,
		ExtraData:     p.ExtraData,
		BaseFeePerGas: new(big.Int).SetBytes(p.BaseFeePerGas),
		BlockHash:     common.Hash(p.BlockHash),
		Transactions:  transactions,
		Withdrawals:   withdrawals,
		BlobGasUsed:   &p.BlobGasUsed,
		ExcessBlobGas: &p.ExcessBlobGas,
	}, nil
}

func WithdrawalToProto(w *ethtypes.Withdrawal) (Withdrawal, error) {
	if w == nil {
		return Withdrawal{}, errors.New("nil withdrawal")
	} else if w.Validator != 0 {
		return Withdrawal{}, errors.New("non-zero validator index")
	}

	return Withdrawal{
		Index:      w.Index,
		Address:    Address(w.Address),
		AmountGwei: w.Amount,
	}, nil
}

func WithdrawalFromProto(w Withdrawal) *ethtypes.Withdrawal {
	return &ethtypes.Withdrawal{
		Index:   w.Index,
		Address: common.Address(w.Address),
		Amount:  w.AmountGwei,
	}
}
