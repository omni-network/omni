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
	"entgo.io/ent/schema/index"
)

// Msg holds the schema definition for the Msg entity.
type Msg struct {
	ent.Schema
}

// Fields of the XMsg.
func (Msg) Fields() []ent.Field {
	return []ent.Field{
		field.Bytes("sender").MaxLen(20),
		field.Bytes("to").MaxLen(20),
		field.Bytes("data"),
		field.Uint64("gas_limit"),
		field.Uint64("source_chain_id"),
		field.Uint64("dest_chain_id"),
		field.Uint64("offset"),
		field.Bytes("tx_hash").MaxLen(32),
		field.Bytes("receipt_hash").MaxLen(32).Optional(),
		field.String("status").Optional().Default("PENDING"),
		field.Time("created_at").Default(time.Now()),
	}
}

// Indexes of the Msg.
func (Msg) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("sender"),
		index.Fields("to"),
		index.Fields("status"),
		index.Fields("tx_hash"),
		index.Fields("source_chain_id", "dest_chain_id", "offset").Unique(),
	}
}

// Edges of the XMsg.
func (Msg) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("block", Block.Type).Ref("msgs"),
		edge.To("receipts", Receipt.Type),
	}
}

// Hooks of the Msg.
func (Msg) Hooks() []ent.Hook {
	return []ent.Hook{
		// Hook for setting edges to receipts.
		hook.On(
			func(next ent.Mutator) ent.Mutator {
				return hook.MsgFunc(func(ctx context.Context, m *gen.MsgMutation) (ent.Value, error) {
					// find the associated receipt using source chain id, dest chain id and txhash
					sourceChainID, ok := m.SourceChainID()
					if !ok {
						return nil, errors.New("source chain id missing")
					}

					destChainID, ok := m.DestChainID()
					if !ok {
						return nil, errors.New("dest chain id missing")
					}

					offset, ok := m.Offset()
					if !ok {
						return nil, errors.New("stream offset missing")
					}
					receipts, err := m.Client().Receipt.Query().
						Where(
							receipt.SourceChainID(sourceChainID),
							receipt.DestChainID(destChainID),
							receipt.Offset(offset),
						).
						Order(gen.Desc(receipt.FieldCreatedAt)).
						All(ctx)
					if err != nil {
						return nil, err
					}

					if len(receipts) == 0 {
						return next.Mutate(ctx, m)
					}
					status := "SUCCESS"
					if !receipts[0].Success {
						status = "FAILED"
					}
					m.SetStatus(status)

					for _, r := range receipts {
						m.AddReceiptIDs(r.ID)
					}

					return next.Mutate(ctx, m)
				})
			},
			// Limit the hook only for these operations.
			// If we added update, we would get an infinite loop here, if we ever updated receipts which we don't do currently
			// This mutation checks for the existence of a receipt in the case where we received the receipt prior to the message.
			// We also check for the message when we create a receipt
			ent.OpCreate,
		),
	}
}
