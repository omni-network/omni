package anchorinbox

import (
	"context"
	"crypto/rand"
	"crypto/sha256"
	"encoding/binary"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/solutil"
	"github.com/omni-network/omni/lib/umath"

	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/programs/system"
	"github.com/gagliardetto/solana-go/programs/token"
	"github.com/gagliardetto/solana-go/rpc"
)

var EventNameUpdated = "EventUpdated"

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

// GetOrderState retrieves the order state account data for a given order ID.
func GetOrderState(ctx context.Context, cl *rpc.Client, orderID solana.PublicKey) (OrderStateAccount, bool, error) {
	// Find the PDA address for the order state account.
	orderState, _, err := FindOrderStateAddress(orderID)
	if err != nil {
		return OrderStateAccount{}, false, errors.Wrap(err, "find order state address")
	}

	// Decode the account data into an OrderState struct.
	var orderStateData OrderStateAccount
	_, err = solutil.GetAccountDataInto(ctx, cl, orderState, &orderStateData)
	if errors.Is(err, rpc.ErrNotFound) {
		return OrderStateAccount{}, false, nil
	} else if err != nil {
		return OrderStateAccount{}, false, errors.Wrap(err, "get order state account data")
	}

	return orderStateData, true, nil
}

func FillData(chainID uint64, state OrderStateAccount) (bindings.SolverNetFillOriginData, error) {
	deadline, err := umath.ToUint32(state.ClosableAt)
	if err != nil {
		return bindings.SolverNetFillOriginData{}, err
	}

	return bindings.SolverNetFillOriginData{
		SrcChainId:   chainID,
		DestChainId:  state.DestChainId,
		FillDeadline: deadline,
		Calls: []bindings.SolverNetCall{
			{
				Target:   state.DestCall.Target,
				Selector: state.DestCall.Selector,
				Value:    state.DestCall.Value.BigInt(),
				Params:   state.DestCall.Params,
			},
		},
		Expenses: []bindings.SolverNetTokenExpense{
			{
				Spender: state.DestExpense.Spender,
				Token:   state.DestExpense.Token,
				Amount:  state.DestExpense.Amount.BigInt(),
			},
		},
	}, nil
}

func NewMarkFilledOrder(ctx context.Context, cl *rpc.Client, claimableBy, admin, orderID solana.PublicKey) (*MarkFilled, error) {
	state, ok, err := GetOrderState(ctx, cl, orderID)
	if err != nil {
		return nil, err
	} else if !ok {
		return nil, errors.New("order not found")
	}

	chainID, err := solutil.ChainID(ctx, cl)
	if err != nil {
		return nil, err
	}

	fillHash, err := fillHash(chainID, state)
	if err != nil {
		return nil, err
	}

	orderState, _, err := FindOrderStateAddress(orderID)
	if err != nil {
		return nil, err
	}

	inboxState, _, err := FindInboxStateAddress()
	if err != nil {
		return nil, err
	}

	return NewMarkFilledInstruction(
		orderID,
		fillHash,
		claimableBy,
		orderState,
		inboxState,
		admin,
	), nil
}

func NewClaimOrder(ctx context.Context, cl *rpc.Client, claimer, orderID solana.PublicKey) (*Claim, error) {
	state, ok, err := GetOrderState(ctx, cl, orderID)
	if err != nil {
		return nil, err
	} else if !ok {
		return nil, errors.New("order not found")
	}

	claimerToken, _, err := solana.FindAssociatedTokenAddress(claimer, state.DepositMint)
	if err != nil {
		return nil, errors.Wrap(err, "find ata")
	}

	ownerToken, _, err := solana.FindAssociatedTokenAddress(state.Owner, state.DepositMint)
	if err != nil {
		return nil, errors.Wrap(err, "find ata")
	}

	orderState, _, err := FindOrderStateAddress(orderID)
	if err != nil {
		return nil, err
	}

	orderToken, _, err := FindOrderTokenAddress(orderID)
	if err != nil {
		return nil, err
	}

	return NewClaimInstruction(
		orderID,
		orderState,
		orderToken,
		ownerToken,
		claimer,
		claimerToken,
		token.ProgramID,
	), nil
}

// NewRejectOrder returns a new order rejection instruction.
func NewRejectOrder(ctx context.Context, cl *rpc.Client, admin, orderID solana.PublicKey, reason uint8) (*Reject, error) {
	state, ok, err := GetOrderState(ctx, cl, orderID)
	if err != nil {
		return nil, err
	} else if !ok {
		return nil, errors.New("order not found")
	}

	ownerToken, _, err := solana.FindAssociatedTokenAddress(state.Owner, state.DepositMint)
	if err != nil {
		return nil, errors.Wrap(err, "find ata")
	}

	orderState, _, err := FindOrderStateAddress(orderID)
	if err != nil {
		return nil, err
	}

	orderToken, _, err := FindOrderTokenAddress(orderID)
	if err != nil {
		return nil, err
	}

	inboxState, _, err := FindInboxStateAddress()
	if err != nil {
		return nil, err
	}

	return NewRejectInstruction(
		orderID,
		reason,
		orderState,
		orderToken,
		ownerToken,
		inboxState,
		admin,
		token.ProgramID,
	), nil
}

func fillHash(chainID uint64, state OrderStateAccount) (solana.PublicKey, error) {
	fillDeadline, err := umath.ToUint32(state.ClosableAt)
	if err != nil {
		return solana.PublicKey{}, err
	}

	resp, err := solutil.FillHash(
		state.OrderId,
		chainID,
		state.DestChainId,
		fillDeadline,
		state.DestCall.Target,
		state.DestCall.Selector,
		state.DestCall.Value.BigInt(),
		state.DestCall.Params,
		state.DestExpense.Spender,
		state.DestExpense.Token,
		state.DestExpense.Amount.BigInt(),
	)
	if err != nil {
		return solana.PublicKey{}, err
	}

	return solana.PublicKey(resp), nil
}

// DecodeMetaError returns a CustomError if encoded in the transaction meta.
func DecodeMetaError(res *rpc.GetTransactionResult) CustomError {
	if res == nil || res.Meta.Err == nil {
		return nil
	}

	m1, ok := res.Meta.Err.(map[string]any)
	if !ok {
		return nil
	}

	m2, ok := m1["InstructionError"].([]any)
	if !ok || len(m2) < 2 {
		return nil
	}

	m3, ok := m2[1].(map[string]any)
	if !ok {
		return nil
	}

	m4, ok := m3["Custom"].(float64)
	if !ok {
		return nil
	}

	return Errors[int(m4)]
}
