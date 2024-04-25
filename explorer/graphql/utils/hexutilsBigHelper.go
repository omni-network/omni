package utils

import "github.com/ethereum/go-ethereum/common/hexutil"

func TryParseHexUtilBigToUint64(b *hexutil.Big) *uint64 {
	if b == nil {
		return nil
	}
	val := b.ToInt().Uint64()

	return &val
}
