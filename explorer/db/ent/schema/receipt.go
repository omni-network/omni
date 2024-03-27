package schema

import (
	"context"
	gen "github.com/omni-network/omni/explorer/db/ent"
	"github.com/omni-network/omni/explorer/db/ent/hook"
	"github.com/omni-network/omni/explorer/db/ent/msg"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// Receipt holds the schema definition for the Receipt entity.
type Receipt struct {
	ent.Schema
}

// Fields of the Receipt.
func (Receipt) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("UUID", uuid.UUID{}).
			Default(uuid.New),
		field.Uint64("GasUsed"),
		field.Bool("Success"),
		field.Bytes("RelayerAddress").
			MaxLen(20),
		field.Uint64("SourceChainID"),
		field.Uint64("DestChainID"),
		field.Uint64("StreamOffset"),
		field.Bytes("TxHash").
			MaxLen(32),
		field.Time("CreatedAt").
			Default(time.Now()),
	}
}

// Edges of the XReceipt.
func (Receipt) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("Block", Block.Type).
			Ref("Receipts").
			Unique(),
		edge.From("Msgs", Msg.Type).
			Ref("Receipts"),
	}
}

// Hooks of the Msg.
func (Receipt) Hooks() []ent.Hook {
	return []ent.Hook{
		// First hook.
		hook.On(
			func(next ent.Mutator) ent.Mutator {
				return hook.MsgFunc(func(ctx context.Context, m *gen.MsgMutation) (ent.Value, error) {
					// go and find the associated message using source chain id, dest chain id and txhash
					sourceChainID, ok := m.SourceChainID()
					if !ok {
						return nil, nil
					}

					destChainID, ok := m.DestChainID()
					if !ok {
						return nil, nil
					}

					txHash, ok := m.TxHash()
					if !ok {
						return nil, nil
					}
					matches, err := m.Client().Msg.Query().Where(
						msg.SourceChainID(sourceChainID),
						msg.DestChainID(destChainID),
						msg.TxHash(txHash),
					).All(ctx)
					if err != nil {
						return nil, err
					}

					for _, match := range matches {
						m.AddReceiptIDs(match.ID)
					}
					return next.Mutate(ctx, m)
				})
			},
			// Limit the hook only for the create operation.
			// If we added update, we would get an infinite loop here
			// This mutation checks for the existence of messages that it should be associated with when we create a receipt
			// We also do the inverse of this, checking for receipts when we create a message
			ent.OpCreate,
		),
	}
}
