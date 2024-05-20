package schema

import (
	"context"
	"time"

	gen "github.com/omni-network/omni/explorer/db/ent"
	"github.com/omni-network/omni/explorer/db/ent/hook"
	"github.com/omni-network/omni/explorer/db/ent/msg"
	"github.com/omni-network/omni/lib/errors"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
)

// Receipt holds the schema definition for the Receipt entity.
type Receipt struct {
	ent.Schema
}

// Fields of the Receipt.
func (Receipt) Fields() []ent.Field {
	return []ent.Field{
		field.Bytes("block_hash").MaxLen(32),
		field.Uint64("gas_used"),
		field.Bool("success"),
		field.Bytes("relayer_address").MaxLen(20),
		field.Uint64("source_chain_id"),
		field.Uint64("dest_chain_id"),
		field.Uint64("offset"),
		field.Bytes("tx_hash").MaxLen(32),
		field.Time("created_at").Default(time.Now()),
	}
}

// Edges of the Receipt.
func (Receipt) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("block", Block.Type).Ref("receipts"),
		edge.From("msgs", Msg.Type).Ref("receipts"),
	}
}

// Indexes of the Receipt.
func (Receipt) Indexes() []ent.Index {
	return []ent.Index{
		index.Fields("tx_hash"),
		index.Fields("source_chain_id", "dest_chain_id", "offset"),
	}
}

// Hooks of the Msg.
func (Receipt) Hooks() []ent.Hook {
	return []ent.Hook{
		// Hook for setting edges to messages.
		hook.On(
			func(next ent.Mutator) ent.Mutator {
				return hook.ReceiptFunc(func(ctx context.Context, r *gen.ReceiptMutation) (ent.Value, error) {
					// go and find the associated message using source chain id, dest chain id and txhash
					sourceChainID, ok := r.SourceChainID()
					if !ok {
						return nil, errors.New("source chain id missing")
					}

					destChainID, ok := r.DestChainID()
					if !ok {
						return nil, errors.New("dest chain id missing")
					}

					offset, ok := r.Offset()
					if !ok {
						return nil, errors.New("stream offset missing")
					}
					matches, err := r.Client().Msg.Query().Where(
						msg.SourceChainID(sourceChainID),
						msg.DestChainID(destChainID),
						msg.Offset(offset),
					).All(ctx)
					if err != nil {
						return nil, err
					}

					status := "SUCCESS"
					success, ok := r.Success()
					if !ok {
						return nil, errors.New("success missing")
					}

					if !success {
						status = "FAILED"
					}

					txHash, ok := r.TxHash()
					if !ok {
						return nil, errors.New("tx hash missing")
					}

					for _, match := range matches {
						r.AddMsgIDs(match.ID)
						// Setting the status of the message and the receipt hash (always going to be most recent)
						r.Client().Msg.UpdateOne(match).
							SetStatus(status).
							SetReceiptHash(txHash).
							SaveX(ctx)
					}

					return next.Mutate(ctx, r)
				})
			},
			// Limit the hook only for the create operation.
			// If we added update, we would get an infinite loop here if we ever updated messages which we don't do currently
			// This mutation checks for the existence of messages that it should be associated with when we create a receipt
			// We also do the inverse of this, checking for receipts when we create a message
			ent.OpCreate,
		),
	}
}
