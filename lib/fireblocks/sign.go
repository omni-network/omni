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

func (c Client) createRawSignTransaction(ctx context.Context, account uint64, digest common.Hash) (string, error) {
	request := c.newRawSignRequest(account, digest)
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
		return "", errors.New("failed to create transaction", "resp_msg", errRes.Message, "resp_code", errRes.Code)
	}

	return res.ID, nil
}

// Sign creates a raw sign transaction and waits for it to complete, returning the resulting signature (Ethereum RSV format).
// The signer address is checked against the resulting signed address.
func (c Client) Sign(ctx context.Context, digest common.Hash, signer common.Address) ([65]byte, error) {
	account, err := c.getAccount(ctx, signer)
	if err != nil {
		return [65]byte{}, err
	}

	id, err := c.createRawSignTransaction(ctx, account, digest)
	if err != nil {
		return [65]byte{}, errors.Wrap(err, "create raw sign tx")
	}

	log.Debug(ctx, "Awaiting fireblocks signature", "id", id)

	// First try immediately.
	resp, status, err := c.maybeGetSignature(ctx, id, digest, signer)
	if err != nil {
		return [65]byte{}, errors.Wrap(err, "get sig")
	} else if status.Completed() {
		return resp, nil
	}

	// Then poll every QueryInterval
	queryTicker := time.NewTicker(c.cfg.QueryInterval)
	defer queryTicker.Stop()

	var attempt int
	prevStatus := status
	for {
		select {
		case <-ctx.Done():
			return [65]byte{}, errors.Wrap(ctx.Err(), "timeout waiting", "prev_status", prevStatus)
		case <-queryTicker.C:
			resp, status, err = c.maybeGetSignature(ctx, id, digest, signer)
			if err != nil {
				return [65]byte{}, errors.Wrap(err, "get sig", "prev_status", prevStatus)
			} else if status.Completed() {
				return resp, nil
			}

			prevStatus = status

			attempt++
			if attempt%c.cfg.LogFreqFactor == 0 {
				log.Warn(ctx, "Fireblocks transaction not signed yet (will retry)", nil,
					"sender", signer,
					"attempt", attempt,
					"status", status,
					"id", id,
				)
			}
		}
	}
}

// maybeGetSignature returns the resulting signature and "completed" status if the transaction has been signed.
// If the transaction is still pending, it returns an empty signature and the "pending" status.
// If the transaction has failed, it returns an empty signature and the "failed" status and an error.
func (c Client) maybeGetSignature(ctx context.Context, txID string, digest common.Hash, signer common.Address) ([65]byte, Status, error) {
	tx, err := c.getTransactionByID(ctx, txID)
	if err != nil {
		return [65]byte{}, "", errors.Wrap(err, "get tx")
	}

	if tx.Status.Failed() {
		return [65]byte{}, tx.Status, errors.New("transaction failed", "status", tx.Status)
	} else if !tx.Status.Completed() {
		return [65]byte{}, tx.Status, nil
	}
	// Get the resulting signature.
	sig, err := tx.Sig0()
	if err != nil {
		return [65]byte{}, "", errors.Wrap(err, "get signature")
	}

	// Get the signer pubkey.
	pubkey, err := tx.Pubkey0()
	if err != nil {
		return [65]byte{}, "", err
	}
	addr := crypto.PubkeyToAddress(*pubkey)
	if addr != signer { // Ensure it matches the expected signer.
		return [65]byte{}, "", errors.New("signed address mismatch", "expect", signer, "actual", addr)
	}

	// Ensure the signature is valid.
	if !crypto.VerifySignature(crypto.CompressPubkey(pubkey), digest[:], sig[:64]) {
		return [65]byte{}, "", errors.New("signature verification failed")
	}

	return sig, tx.Status, nil
}

// newRawSignRequest creates a new transaction request.
func (c Client) newRawSignRequest(account uint64, digest common.Hash) createTransactionRequest {
	return createTransactionRequest{
		Operation: "RAW",
		Source: source{
			Type: "VAULT_ACCOUNT",
			ID:   strconv.FormatUint(account, 10),
		},
		AssetID: c.getAssetID(),
		ExtraParameters: &extraParameters{
			RawMessageData: rawMessageData{
				Messages: []unsignedRawMessage{{
					Content: hex.EncodeToString(digest[:]), // No 0x prefix, just hex.
				}},
			},
		},
		Note: c.cfg.SignNote,
	}
}
