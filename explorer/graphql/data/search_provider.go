package data

import (
	"context"

	"github.com/omni-network/omni/explorer/db/ent"
	"github.com/omni-network/omni/explorer/db/ent/block"
	"github.com/omni-network/omni/explorer/db/ent/msg"
	"github.com/omni-network/omni/explorer/db/ent/receipt"
	"github.com/omni-network/omni/explorer/graphql/resolvers"
	"github.com/omni-network/omni/explorer/graphql/utils"
	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

// This searches for matches for the following:
// - Block (finds matching block hash)
// - Message (finds matching tx hash)
// - Receipt (finds matching tx hash)
// TODO (Dan): This is a very naive search implementation. It should be improved. We also should search by address?
func (p Provider) Search(ctx context.Context, query string) (*resolvers.SearchResult, bool, error) {
	searchResult := &resolvers.SearchResult{}
	hash, err := hexutil.Decode(query)
	if err != nil {
		return nil, false, errors.Wrap(err, "search hexutil.Decode")
	}

	blockQuery, err := p.EntClient.Block.Query().Where(block.BlockHash(hash)).First(ctx)
	if err != nil && !ent.IsNotFound(err) {
		return nil, false, errors.Wrap(err, "search block graphql provider")
	}

	if blockQuery != nil {
		searchResult.BlockHeight, err = utils.Uint2Hex(blockQuery.BlockHeight)
		if err != nil {
			log.Error(ctx, "Uint2Hex err", err)

			return nil, false, err
		}

		chainID, err := utils.Uint2Hex(blockQuery.SourceChainID)
		if err != nil {
			log.Error(ctx, "Uint2Hex err", err)

			return nil, false, err
		}

		searchResult.SourceChainID = chainID
		searchResult.Type = resolvers.BLOCK

		return searchResult, true, nil
	}

	msgQuery, err := p.EntClient.Msg.Query().Where(msg.TxHash(hash)).First(ctx)
	if err != nil && !ent.IsNotFound(err) {
		return nil, false, errors.Wrap(err, "search msg graphql provider")
	}

	if msgQuery != nil {
		searchResult.TxHash = common.Hash(msgQuery.TxHash)
		searchResult.Type = resolvers.MESSAGE

		return searchResult, true, nil
	}

	receiptQuery, err := p.EntClient.Receipt.Query().Where(receipt.TxHash(hash)).First(ctx)
	if err != nil && !ent.IsNotFound(err) {
		return nil, false, errors.Wrap(err, "search receipt graphql provider")
	}

	if receiptQuery != nil {
		searchResult.TxHash = common.Hash(receiptQuery.TxHash)
		searchResult.Type = resolvers.RECEIPT

		return searchResult, true, nil
	}

	return nil, true, nil
}
