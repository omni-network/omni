package anchorinbox

import (
	"crypto/sha256"
	"encoding/binary"

	"github.com/omni-network/omni/lib/errors"

	"github.com/gagliardetto/solana-go"
)

var (
	EventNameOpened     = "EventOpened"
	EventNameMarkFilled = "EventMarkFilled"
)

// NewOrderID returns the order ID (32 byte array Pubkey),
// by hashing the account (pubkey) and nonce.
func NewOrderID(owner solana.PublicKey, nonce uint64) solana.PublicKey {
	// Convenience function to convert uint64 to bytes.
	uintToBytes := func(n uint64) []byte {
		b := make([]byte, 8)
		binary.LittleEndian.PutUint64(b, n)

		return b
	}

	h := sha256.New()
	_, _ = h.Write(owner[:])
	_, _ = h.Write(uintToBytes(nonce))

	return solana.PublicKeyFromBytes(h.Sum(nil))
}

// FindOrderStateAddress returns the address of the order state account.
//
// This is equivalent to anchor: `seeds = [OrderState::SEED_PREFIX, params.order_id.as_ref()]`.
//
// Note the generated code variant produces invalid results since it assumes MsgPack encoding instruction args.
// `(&Open{Params: &OpenParams{OrderId: orderID}}).FindOrderStateAddress()`.
func FindOrderStateAddress(orderID solana.PublicKey) (solana.PublicKey, uint8, error) {
	seeds := [][]byte{
		[]byte("order_state"),
		orderID[:],
	}

	account, bump, err := solana.FindProgramAddress(seeds, ProgramID)
	if err != nil {
		return solana.PublicKey{}, 0, errors.Wrap(err, "find program address")
	}

	return account, bump, nil
}

// FindOrderTokenAddress returns the address of the order token account.
//
// This is equivalent to anchor: `seeds = [ORDER_TOKEN_SEED_PREFIX, params.order_id.as_ref()]`.
//
// Note the generated code variant produces invalid results since it assumes MsgPack encoding instruction args.
// `(&Open{Params: &OpenParams{OrderId: orderID}}).FindOrderTokenAddress()`.
func FindOrderTokenAddress(orderID solana.PublicKey) (solana.PublicKey, uint8, error) {
	seeds := [][]byte{
		[]byte("order_token"),
		orderID[:],
	}

	account, bump, err := solana.FindProgramAddress(seeds, ProgramID)
	if err != nil {
		return solana.PublicKey{}, 0, errors.Wrap(err, "find program address")
	}

	return account, bump, nil
}

// FindInboxStateAddress returns the address of the inbox state account.
//
// This is equivalent to anchor: `seeds = [InitState::SEED_PREFIX]`.
//
// Generated code variant is fine, since no instruction arguments required.
func FindInboxStateAddress() (solana.PublicKey, uint8, error) {
	return new(Init).FindInboxStateAddress()
}
