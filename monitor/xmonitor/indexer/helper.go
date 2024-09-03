package indexer

import (
	"encoding/json"

	"github.com/omni-network/omni/lib/errors"
	"github.com/omni-network/omni/lib/xchain"

	"github.com/ethereum/go-ethereum/common"
)

func (b *Block) XChainBlock() (xchain.Block, error) {
	var resp xchain.Block
	if err := json.Unmarshal(b.GetBlockJson(), &resp); err != nil { //nolint:musttag // TODO: Rather use protobuf
		return xchain.Block{}, errors.Wrap(err, "unmarshal block")
	}

	return resp, nil
}

func (l *MsgLink) Hash() common.Hash {
	return common.BytesToHash(l.GetIdHash())
}
