package uintconv

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common/hexutil"
)

func ToBig(u uint64) (big.Int, error) {
	b, err := hexutil.DecodeBig(hexutil.EncodeUint64(u))
	if err != nil {
		return big.Int{}, err
	}

	return *b, nil
}

func ToHex(u uint64) (hexutil.Big, error) {
	b, err := hexutil.DecodeBig(hexutil.EncodeUint64(u))
	if err != nil {
		return hexutil.Big{}, err
	}
	hex := hexutil.Big(*b)

	return hex, nil
}
