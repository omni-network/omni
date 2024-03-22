package fireblocks

import (
	"context"
	"encoding/hex"
	"net/http"
	"strconv"
	"time"

	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
)

func (c Client) createRawSignTransaction(ctx context.Context, digest common.Hash) (string, error) {
	request := c.newRawSignRequest(digest)
	headers, err := c.authHeaders(endpointTransactions, request)
	if err != nil {
		return "", err
	}

	var res createTransactionResponse
	var errRes errorResponse
	ok, err := c.jsonHTTP.Send(
		ctx,
		endpointTransactions,
		http.MethodPost,
		request,
		headers,
		&res,
		&errRes,
	)
	if err != nil {
		return "", err
	} else if !ok {
		return "", errors.New("failed to create transaction", "msg", errRes.Message, "code", errRes.Code)
	}

	return res.ID, nil
}

// Sign creates a raw sign transaction and waits for it to complete, returning the resulting signature (Ethereum RSV format).
// The signer address is checked against the resulting signed address.
func (c Client) Sign(ctx context.Context, digest common.Hash, signer common.Address) ([65]byte, error) {
	id, err := c.createRawSignTransaction(ctx, digest)
	if err != nil {
		return [65]byte{}, err
	}

	// First try immediately.
	if resp, ok, err := c.maybeGetSignature(ctx, id, digest, signer); err != nil {
		return [65]byte{}, err
	} else if ok {
		return resp, nil
	}

	// Then poll every QueryInterval
	queryTicker := time.NewTicker(c.opts.QueryInterval)
	defer queryTicker.Stop()

	var attempt int
	for {
		select {
		case <-ctx.Done():
			return [65]byte{}, errors.Wrap(ctx.Err(), "context canceled")
		case <-queryTicker.C:
			if resp, ok, err := c.maybeGetSignature(ctx, id, digest, signer); err != nil {
				return [65]byte{}, err
			} else if ok {
				return resp, nil
			}

			attempt++
			if attempt%c.opts.LogFreqFactor == 0 {
				log.Warn(ctx, "Fireblocks transaction not signed yet (will retry)", nil,
					"attempt", attempt,
					"id", id,
				)
			}
		}
	}
}

func (c Client) maybeGetSignature(ctx context.Context, txID string, digest common.Hash, signer common.Address) ([65]byte, bool, error) {
	tx, err := c.getTransactionByID(ctx, txID)
	if err != nil {
		return [65]byte{}, false, err
	}

	ok, err := isComplete(tx)
	if err != nil {
		return [65]byte{}, false, err
	} else if !ok {
		return [65]byte{}, false, nil
	}

	// Get the resulting signature.
	sig, err := tx.Sig0()
	if err != nil {
		return [65]byte{}, false, errors.Wrap(err, "get signature")
	}

	// Get the signer pubkey.
	pubkey, err := tx.Pubkey0()
	if err != nil {
		return [65]byte{}, false, err
	}
	addr := crypto.PubkeyToAddress(*pubkey)
	if addr != signer { // Ensure it matches the expected signer.
		return [65]byte{}, false, errors.New("signed address mismatch", "expect", signer, "actual", addr)
	}

	// Ensure the signature is valid.
	if !crypto.VerifySignature(crypto.CompressPubkey(pubkey), digest[:], sig[:64]) {
		return [65]byte{}, false, errors.New("signature verification failed")
	}

	return sig, true, nil
}

// isComplete returns true if the transaction is complete, false if still pending, or an error if it failed.
func isComplete(tx transaction) (bool, error) {
	switch tx.Status {
	case "COMPLETED":
		return true, nil
	case "CANCELED", "BLOCKED_BY_POLICY", "REJECTED", "FAILED":
		return false, errors.New("transaction failed", "status", tx.Status)
	default:
		return false, nil
	}
}

// newRawSignRequest creates a new transaction request.
func (c Client) newRawSignRequest(digest common.Hash) createTransactionRequest {
	return createTransactionRequest{
		Operation: "RAW",
		Source: source{
			Type: "VAULT_ACCOUNT",
			ID:   strconv.FormatUint(c.opts.VaultAccountID, 10),
		},
		AssetID: c.getAssetID(),
		ExtraParameters: &extraParameters{
			RawMessageData: rawMessageData{
				Messages: []unsignedRawMessage{{
					Content: hex.EncodeToString(digest[:]), // No 0x prefix, just hex.
				}},
			},
		},
	}
}
