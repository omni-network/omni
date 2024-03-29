package schema

import (
	"context"
	"time"

	gen "github.com/omni-network/omni/explorer/db/ent"
	"github.com/omni-network/omni/explorer/db/ent/hook"
	"github.com/omni-network/omni/explorer/db/ent/receipt"
	"github.com/omni-network/omni/lib/errors"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Msg holds the schema definition for the Msg entity.
type Msg struct {
	ent.Schema
}

// Fields of the XMsg.
func (Msg) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("UUID", uuid.UUID{}).
			Default(uuid.New),
		field.Bytes("SourceMsgSender").
			MaxLen(20),
		field.Bytes("DestAddress").
			MaxLen(20),
		field.Bytes("Data"),
		field.Uint64("DestGasLimit"),
		field.Uint64("SourceChainID"),
		field.Uint64("DestChainID"),
		field.Uint64("StreamOffset"),
		field.Bytes("TxHash").
			MaxLen(32),
		field.Time("CreatedAt").
			Default(time.Now()),
	}
}

// Edges of the XMsg.
func (Msg) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("Block", Block.Type).
			Ref("Msgs").
			Unique(),
		edge.To("Receipts", Receipt.Type),
	}
}

// Hooks of the Msg.
func (Msg) Hooks() []ent.Hook {
	return []ent.Hook{
		// First hook.
		hook.On(
			func(next ent.Mutator) ent.Mutator {
				return hook.MsgFunc(func(ctx context.Context, m *gen.MsgMutation) (ent.Value, error) {
					// go and find the associated receipt using source chain id, dest chain id and txhash
					sourceChainID, ok := m.SourceChainID()
					if !ok {
						return nil, errors.New("source chain id missing")
					}

					destChainID, ok := m.DestChainID()
					if !ok {
						return nil, errors.New("dest chain id missing")
					}

					streamOffset, ok := m.StreamOffset()
					if !ok {
						return nil, errors.New("stream offset missing")
					}
					receipts, err := m.Client().Receipt.Query().Where(
						receipt.SourceChainID(sourceChainID),
						receipt.DestChainID(destChainID),
						receipt.StreamOffset(streamOffset),
					).All(ctx)
					if err != nil {
						return nil, err
					}

					for _, r := range receipts {
						m.AddReceiptIDs(r.ID)
					}

					return next.Mutate(ctx, m)
				})
			},
			// Limit the hook only for these operations.
			// If we added update, we would get an infinite loop here
			// This mutation checks for the existence of a receipt in the case where we received the receipt prior to the message.
			// We also check for the message when we create a receipt
			ent.OpCreate,
		),
	}
}
