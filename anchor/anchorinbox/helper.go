package anchorinbox

import (
	"context"
	"crypto/rand"
	"crypto/sha256"
	"encoding/binary"
	"time"

	"github.com/omni-network/omni/contracts/bindings"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/svmutil"
	"github.com/omni-network/omni/lib/umath"

	"github.com/gagliardetto/solana-go"
	"github.com/gagliardetto/solana-go/programs/system"
	"github.com/gagliardetto/solana-go/programs/token"
	"github.com/gagliardetto/solana-go/rpc"
)

var EventNameUpdated = "EventUpdated"

func DecodeLogEvents(logMsgs []string) ([]*Event, error) {
	decoded, err := decodeEventsFromLogMessage(logMsgs)
	if err != nil {
		return nil, err
	}

	return parseEvents(decoded)
}

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
func FindOrderStateAddress(orderID solana.PublicKey) (solana.PublicKey, uint8, error) {
	return (&OpenInstruction{Params: &OpenParams{OrderId: orderID}}).FindOrderStateAddress()
}

// FindOrderTokenAddress returns the address of the order token account.
//
// This is equivalent to anchor: `seeds = [ORDER_TOKEN_SEED_PREFIX, params.order_id.as_ref()]`.
func FindOrderTokenAddress(orderID solana.PublicKey) (solana.PublicKey, uint8, error) {
	return (&OpenInstruction{Params: &OpenParams{OrderId: orderID}}).FindOrderTokenAddress()
}

// FindInboxStateAddress returns the address of the inbox state account.
//
// This is equivalent to anchor: `seeds = [InitState::SEED_PREFIX]`.
func FindInboxStateAddress() (solana.PublicKey, uint8, error) {
	return new(InitInstruction).FindInboxStateAddress()
}

// NewInit returns a new Init instruction with the given parameters.
func NewInit(chainID uint64, closeBuffer time.Duration, admin solana.PublicKey) (*InitInstruction, error) {
	inboxState, _, err := FindInboxStateAddress()
	if err != nil {
		return nil, err
	}

	return NewInitInstruction(chainID, int64(closeBuffer.Seconds()), inboxState, admin, system.ProgramID), nil
}

type OpenOrder struct {
	*OpenInstruction

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
func NewOpenOrder(params OpenParams, owner, mint solana.PublicKey) (OpenOrder, error) {
	ownerToken, _, err := solana.FindAssociatedTokenAddress(owner, mint)
	if err != nil {
		return OpenOrder{}, errors.Wrap(err, "find owner ata")
	}

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

	inboxState, _, err := FindInboxStateAddress()
	if err != nil {
		return OpenOrder{}, err
	}

	open := NewOpenInstruction(params, orderState, owner, mint, ownerToken, orderToken, token.ProgramID, inboxState, system.ProgramID)

	return OpenOrder{
		OpenInstruction: open,
		ID:              params.OrderId,
		StateAddress:    orderState,
		StateBump:       stateBump,
		TokenAddress:    orderToken,
		TokenBump:       tokenBump,
	}, nil
}

func randU64() uint64 {
	var b [8]byte
	_, _ = rand.Read(b[:])

	return binary.LittleEndian.Uint64(b[:])
}

// GetInboxState returns the inbox state account data or false if it does not exist.
func GetInboxState(ctx context.Context, cl *rpc.Client) (InboxStateAccount, bool, error) {
	// Find the PDA address for the inbox state account.
	inboxState, _, err := FindInboxStateAddress()
	if err != nil {
		return InboxStateAccount{}, false, errors.Wrap(err, "find inbox state address")
	}

	// Decode the account data into an InboxState struct.
	var inboxStateData InboxStateAccount
	_, err = svmutil.GetAccountDataInto(ctx, cl, inboxState, &inboxStateData)
	if errors.Is(err, rpc.ErrNotFound) {
		return InboxStateAccount{}, false, nil
	} else if err != nil {
		return InboxStateAccount{}, false, errors.Wrap(err, "get inbox state account data")
	}

	return inboxStateData, true, nil
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
	_, err = svmutil.GetAccountDataInto(ctx, cl, orderState, &orderStateData)
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

func NewMarkFilledOrder(ctx context.Context, cl *rpc.Client, claimableBy, admin, orderID solana.PublicKey) (*MarkFilledInstruction, error) {
	state, ok, err := GetOrderState(ctx, cl, orderID)
	if err != nil {
		return nil, err
	} else if !ok {
		return nil, errors.New("order not found")
	}

	chainID, err := svmutil.ChainID(ctx, cl)
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

func NewClaimOrder(ctx context.Context, cl *rpc.Client, claimer, orderID solana.PublicKey) (*ClaimInstruction, error) {
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
func NewRejectOrder(ctx context.Context, cl *rpc.Client, admin, orderID solana.PublicKey, reason uint8) (*RejectInstruction, error) {
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

	resp, err := svmutil.FillHash(
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

// GetInitSig retrieves the init instruction signature of the anchor inbox program.
func GetInitSig(ctx context.Context, cl *rpc.Client) (solana.Signature, error) {
	state, ok, err := GetInboxState(ctx, cl)
	if err != nil {
		return solana.Signature{}, errors.Wrap(err, "get inbox state")
	} else if !ok {
		return solana.Signature{}, errors.New("inbox state not found")
	}

	blockResp, ok, err := svmutil.GetBlock(ctx, cl, state.DeployedAt, rpc.TransactionDetailsFull)
	if err != nil {
		return solana.Signature{}, errors.Wrap(err, "get block")
	} else if !ok {
		return solana.Signature{}, errors.New("no deployed_at block", "slot", state.DeployedAt)
	}

	for _, txMeta := range blockResp.Transactions {
		tx, err := txMeta.GetTransaction()
		if err != nil {
			return solana.Signature{}, errors.Wrap(err, "get transaction")
		}

		if ok, err := tx.HasAccount(ProgramID); err != nil {
			return solana.Signature{}, errors.Wrap(err, "has account")
		} else if ok {
			return tx.Signatures[0], nil
		}
	}

	return solana.Signature{}, errors.New("no init instruction found", "slot", state.DeployedAt)
}
