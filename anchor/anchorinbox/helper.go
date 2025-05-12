package anchorinbox

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/binary"

	"github.com/omni-network/omni/lib/errors"

	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/programs/system"
	"github.com/gagliardetto/solana-go/programs/token"
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

type OpenOrder struct {
	*Open

	ID           solana.PublicKey
	StateAddress solana.PublicKey
	StateBump    uint8
	TokenAddress solana.PublicKey
	TokenBump    uint8
}

// NewOpenOrder returns a convenient OpenOrder struct that extends *Open
// and includes inbox pda accounts.
//
// Note that OpenParams Nonce and OrderID will be generated if omitted.
func NewOpenOrder(params OpenParams, owner, mint, ownerToken solana.PublicKey) (OpenOrder, error) {
	for params.Nonce == 0 {
		params.Nonce = randU64()
	}
	params.OrderId = NewOrderID(owner, params.Nonce)

	orderState, stateBump, err := FindOrderStateAddress(params.OrderId)
	if err != nil {
		return OpenOrder{}, err
	}

	orderToken, tokenBump, err := FindOrderTokenAddress(params.OrderId)
	if err != nil {
		return OpenOrder{}, err
	}

	inboxAddr, _, err := FindInboxStateAddress()
	if err != nil {
		return OpenOrder{}, err
	}

	open := NewOpenInstruction(params, orderState, owner, mint, ownerToken, orderToken, token.ProgramID, inboxAddr, system.ProgramID)

	return OpenOrder{
		Open:         open,
		ID:           params.OrderId,
		StateAddress: orderState,
		StateBump:    stateBump,
		TokenAddress: orderToken,
		TokenBump:    tokenBump,
	}, nil
}

func randU64() uint64 {
	var b [8]byte
	_, _ = rand.Read(b[:])

	return binary.LittleEndian.Uint64(b[:])
}
