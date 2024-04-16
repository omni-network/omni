package data

import (
	"context"

	"github.com/omni-network/omni/explorer/db/ent"
	"github.com/omni-network/omni/explorer/db/ent/block"
	"github.com/omni-network/omni/explorer/db/ent/msg"
	"github.com/omni-network/omni/explorer/db/ent/receipt"
	"github.com/omni-network/omni/explorer/graphql/resolvers"
	"github.com/omni-network/omni/lib/log"

	"github.com/ethereum/go-ethereum/common"
)

// This searches for matches for the following:
// - Block (finds matching block hash)
// - Message (finds matching tx hash)
// - Receipt (finds matching tx hash)
// TODO (Dan): This is a very naive search implementation. It should be improved. We also should search by address?
func (p Provider) Search(ctx context.Context, query string) (*resolvers.SearchResult, bool, error) {
	hash := []byte(query)
	searchResult := &resolvers.SearchResult{}

	blockQuery, err := p.EntClient.Block.Query().Where(block.BlockHash(hash)).First(ctx)
	if err != nil && !ent.IsNotFound(err) {
		log.Error(ctx, "Block query graphql provider err", err)

		return nil, false, err
	}

	if blockQuery != nil {
		searchResult.BlockHeight, err = Uint2Hex(blockQuery.BlockHeight)
		if err != nil {
			log.Error(ctx, "Uint2Hex err", err)

			return nil, false, err
		}

		chainID, err := Uint2Hex(blockQuery.SourceChainID)
		if err != nil {
			log.Error(ctx, "Uint2Hex err", err)

			return nil, false, err
		}

		searchResult.SourceChainID = chainID
		searchResult.Type = "block"

		return nil, true, nil
	}

	msgQuery, err := p.EntClient.Msg.Query().Where(msg.TxHash(hash)).First(ctx)
	if err != nil && !ent.IsNotFound(err) {
		log.Error(ctx, "Search msg graphql provider err", err)

		return nil, false, err
	}

	if msgQuery != nil {
		searchResult.TxHash = common.Hash(msgQuery.TxHash)
		searchResult.Type = "message"

		return searchResult, true, nil
	}

	receiptQuery, err := p.EntClient.Receipt.Query().Where(receipt.TxHash(hash)).First(ctx)
	if err != nil && !ent.IsNotFound(err) {
		log.Error(ctx, "Search receipt graphql provider err", err)

		return nil, false, err
	}

	if receiptQuery != nil {
		searchResult.TxHash = common.Hash(receiptQuery.TxHash)
		searchResult.Type = "receipt"

		return searchResult, true, nil
	}

	return nil, true, nil
}
