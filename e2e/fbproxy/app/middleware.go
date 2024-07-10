package app

import (
	"context"
	"encoding/json"
	"math/big"

	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
)

type txSigner interface {
	Sign(ctx context.Context, digest common.Hash, signer common.Address) ([65]byte, error)
}

// NewSendTxMiddleware returns a middleware func that
//   - intercepts eth_sendTransaction requests, signs them `txsigner`
//     and replaces them with eth_sendRawTransaction requests
//   - leaves all other requests unmodified
func NewSendTxMiddleware(txsigner txSigner, chainID *big.Int) Middleware {
	// signer used to create signaure hash and recover tx sender,
	// not for actual signing which is left to txsigner
	signer := types.LatestSignerForChainID(chainID)

	return func(ctx context.Context, req JSONRPCMessage) (JSONRPCMessage, error) {
		if req.Method != "eth_sendTransaction" {
			return req, nil
		}

		log.Debug(ctx, "Intercepted eth_sendTransaction")

		var paramsIn []TransactionArgs
		err := json.Unmarshal(req.Params, &paramsIn)
		if err != nil {
			return JSONRPCMessage{}, errors.Wrap(err, "unmarshal tx")
		}

		if len(paramsIn) != 1 {
			return JSONRPCMessage{}, errors.New("only one transaction supported")
		}

		args := paramsIn[0]
		tx := args.ToTransaction()

		sig, err := txsigner.Sign(ctx, signer.Hash(tx), *args.From)
		if err != nil {
			return JSONRPCMessage{}, errors.Wrap(err, "sign")
		}

		log.Debug(ctx, "Signed tx", "tx", tx.Hash().Hex(), "from", args.From.Hex())

		signed, err := tx.WithSignature(signer, sig[:])
		if err != nil {
			return JSONRPCMessage{}, errors.Wrap(err, "with signature")
		}

		data, err := signed.MarshalBinary()
		if err != nil {
			return JSONRPCMessage{}, errors.Wrap(err, "marshal binary")
		}

		paramsOut, err := json.Marshal([]string{hexutil.Encode(data)})
		if err != nil {
			return JSONRPCMessage{}, errors.Wrap(err, "marshal hex")
		}

		return JSONRPCMessage{
			Version: req.Version,
			ID:      req.ID,
			Method:  "eth_sendRawTransaction",
			Params:  paramsOut,
		}, nil
	}
}
